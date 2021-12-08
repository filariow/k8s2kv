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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// KeyVaultSecretSyncSpec defines the desired state of KeyVaultSecretSync
type KeyVaultSecretSyncSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of KeyVaultSecretSync. Edit keyvaultsecretsync_types.go to remove/update
	Secret             string `json:"secret"`
	Data               string `json:"data"`
	KeyVaultSecretName string `json:"keyVaultSecretName"`
	KeyVaultName       string `json:"keyVaultName"`
}

// KeyVaultSecretSyncStatus defines the observed state of KeyVaultSecretSync
type KeyVaultSecretSyncStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// KeyVaultSecretSync is the Schema for the keyvaultsecretsyncs API
type KeyVaultSecretSync struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KeyVaultSecretSyncSpec   `json:"spec,omitempty"`
	Status KeyVaultSecretSyncStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// KeyVaultSecretSyncList contains a list of KeyVaultSecretSync
type KeyVaultSecretSyncList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KeyVaultSecretSync `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KeyVaultSecretSync{}, &KeyVaultSecretSyncList{})
}
