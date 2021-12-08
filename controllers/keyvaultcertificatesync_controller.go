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
	"context"

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

	// TODO(user): your logic here
	l.Info("Hello", "request", req)

	// get key vault certificate sync resource
	l.Info("retrieving KeyVaultCertificateSync", "request", req.String())
	kvs := &akvv1.KeyVaultCertificateSync{}
	if err := r.Get(ctx, types.NamespacedName{Namespace: req.Namespace, Name: req.Name}, kvs); err != nil {
		l.Info("error retrieving KeyVaultCertificateSync", "request", req.String())
		return ctrl.Result{}, nil
	}

	// get secret from k8s
	l.Info("retrieving secret", "secret namespace", kvs.Namespace, "secret name", kvs.Spec.Secret, "request", req.String())
	s := &corev1.Secret{}
	if err := r.Get(ctx, types.NamespacedName{Namespace: kvs.Namespace, Name: kvs.Spec.Secret}, s); err != nil {
		l.Info("error retrieving secret", "secret namespace", kvs.Namespace, "secret name", kvs.Spec.Secret, "request", req.String())
		return ctrl.Result{}, nil
	}

	// get secret form kv

	// if k8s's secret is different, update the version in kv

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *KeyVaultCertificateSyncReconciler) SetupWithManager(mgr ctrl.Manager) error {
	if err := mgr.GetFieldIndexer().IndexField(context.Background(), &akvv1.KeyVaultCertificateSync{}, secretField, func(rawObj client.Object) []string {
		kvcs := rawObj.(*akvv1.KeyVaultCertificateSync)
		if kvcs.Spec.Secret == "" {
			return nil
		}
		return []string{kvcs.Spec.Secret}
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
