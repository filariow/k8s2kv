/*
Copyright 2021 filario.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"bytes"
	"context"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"

	akvv1 "github.com/filariow/k8s2kv/api/v1"
	"github.com/filariow/k8s2kv/pkg/kv"
)

const (
	secretField = ".spec.secret"
)

// KeyVaultCertificateSyncReconciler reconciles a KeyVaultCertificateSync object
type KeyVaultCertificateSyncReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=akv.fil.it,resources=keyvaultcertificatesyncs,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=akv.fil.it,resources=keyvaultcertificatesyncs/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=akv.fil.it,resources=keyvaultcertificatesyncs/finalizers,verbs=update
//+kubebuilder:rbac:groups="",resources=secrets,verbs=get;list;watch

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the KeyVaultCertificateSync object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.10.0/pkg/reconcile
func (r *KeyVaultCertificateSyncReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	l := log.FromContext(ctx)

	// get key vault certificate sync resource
	l.Info("retrieving KeyVaultCertificateSync", "request", req.String())
	kvs := &akvv1.KeyVaultCertificateSync{}
	if err := r.Get(ctx, types.NamespacedName{Namespace: req.Namespace, Name: req.Name}, kvs); err != nil {
		l.Info("error retrieving KeyVaultCertificateSync", "request", req.String())
		return ctrl.Result{}, nil
	}

	// get secret from k8s
	l.Info("retrieving secret", "secret namespace", kvs.Namespace, "secret name", kvs.Spec.SecretName, "request", req.String())
	s := &corev1.Secret{}
	if err := r.Get(ctx, types.NamespacedName{Namespace: kvs.Namespace, Name: kvs.Spec.SecretName}, s); err != nil {
		l.Info("error retrieving secret", "secret namespace", kvs.Namespace, "secret name", kvs.Spec.SecretName, "request", req.String())
		return ctrl.Result{}, nil
	}

	sp := kvs.Spec
	crt, tkey := s.Data["tls.crt"], s.Data["tls.key"]
	ec := []byte("-----END CERTIFICATE-----\n")

	kp, _ := pem.Decode(tkey)
	if kp == nil {
		l.Info("error parsing Private Key")
		return ctrl.Result{}, nil
	}

	kd, err := x509.ParsePKCS1PrivateKey(kp.Bytes)
	if err != nil {
		l.Info("error decrypting private key", "error", err)
		return ctrl.Result{}, nil
	}

	b, err := x509.MarshalPKCS8PrivateKey(kd)
	if err != nil {
		l.Info("error converting key to PKCS8 Private Key", "error", err, "key", string(tkey))
		return ctrl.Result{}, nil
	}

	key := pem.EncodeToMemory(&pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: b,
	})

	crt = bytes.Split(crt, ec)[0]
	value := append(append(crt, ec...), key...)

	// get secret form kv
	sv, err := kv.GetCertificate(ctx, sp.KeyVaultName, sp.KeyVaultCertificateName)
	if err != nil && !errors.Is(err, kv.ErrNotFound) {
		l.Info("error retrieving key vault certificate", "key vault", sp.KeyVaultName)
		return ctrl.Result{}, nil
	}

	// if k8s's secret is different, update the version in kv
	if sv != nil && bytes.Equal(*sv, value) {
		l.Info("certificate was up to date", "secret namespace", kvs.Namespace, "secret name", kvs.Spec.SecretName, "request", req.String())
		return ctrl.Result{}, nil
	}

	l.Info("updating certificate", "secret namespace", kvs.Namespace, "secret name", kvs.Spec.SecretName, "request", req.String())
	b64cert := base64.StdEncoding.EncodeToString(value)
	if err := kv.ImportCertificate(ctx, sp.KeyVaultName, sp.KeyVaultCertificateName, b64cert); err != nil {
		l.Info("error updating key vault certificate", "key vault", sp.KeyVaultName, "error", err, "pem", string(value))
		return ctrl.Result{}, nil
	}

	l.Info("certificate updated", "secret namespace", kvs.Namespace, "secret name", kvs.Spec.SecretName, "request", req.String())
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *KeyVaultCertificateSyncReconciler) SetupWithManager(mgr ctrl.Manager) error {
	if err := mgr.GetFieldIndexer().IndexField(context.Background(), &akvv1.KeyVaultCertificateSync{}, secretField, func(rawObj client.Object) []string {
		kvcs := rawObj.(*akvv1.KeyVaultCertificateSync)
		if kvcs.Spec.SecretName == "" {
			return nil
		}
		return []string{kvcs.Spec.SecretName}
	}); err != nil {
		return err
	}

	return ctrl.NewControllerManagedBy(mgr).
		For(&akvv1.KeyVaultCertificateSync{}).
		Watches(
			&source.Kind{Type: &corev1.Secret{}},
			handler.EnqueueRequestsFromMapFunc(r.findSecret),
			builder.WithPredicates(predicate.ResourceVersionChangedPredicate{}),
		).
		Complete(r)
}

func (r *KeyVaultCertificateSyncReconciler) findSecret(secret client.Object) []reconcile.Request {
	attachedSecrets := &akvv1.KeyVaultCertificateSyncList{}
	listOps := &client.ListOptions{
		FieldSelector: fields.OneTermEqualSelector(secretField, secret.GetName()),
		Namespace:     secret.GetNamespace(),
	}
	err := r.List(context.TODO(), attachedSecrets, listOps)
	if err != nil {
		return []reconcile.Request{}
	}

	rr := make([]reconcile.Request, len(attachedSecrets.Items))
	for i, item := range attachedSecrets.Items {
		rr[i] = reconcile.Request{
			NamespacedName: types.NamespacedName{
				Name:      item.GetName(),
				Namespace: item.GetNamespace(),
			},
		}
	}
	return rr
}
