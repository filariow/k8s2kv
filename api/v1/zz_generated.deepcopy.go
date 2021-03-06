//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyVaultCertificateSync) DeepCopyInto(out *KeyVaultCertificateSync) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyVaultCertificateSync.
func (in *KeyVaultCertificateSync) DeepCopy() *KeyVaultCertificateSync {
	if in == nil {
		return nil
	}
	out := new(KeyVaultCertificateSync)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KeyVaultCertificateSync) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyVaultCertificateSyncList) DeepCopyInto(out *KeyVaultCertificateSyncList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KeyVaultCertificateSync, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyVaultCertificateSyncList.
func (in *KeyVaultCertificateSyncList) DeepCopy() *KeyVaultCertificateSyncList {
	if in == nil {
		return nil
	}
	out := new(KeyVaultCertificateSyncList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KeyVaultCertificateSyncList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyVaultCertificateSyncSpec) DeepCopyInto(out *KeyVaultCertificateSyncSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyVaultCertificateSyncSpec.
func (in *KeyVaultCertificateSyncSpec) DeepCopy() *KeyVaultCertificateSyncSpec {
	if in == nil {
		return nil
	}
	out := new(KeyVaultCertificateSyncSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyVaultCertificateSyncStatus) DeepCopyInto(out *KeyVaultCertificateSyncStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyVaultCertificateSyncStatus.
func (in *KeyVaultCertificateSyncStatus) DeepCopy() *KeyVaultCertificateSyncStatus {
	if in == nil {
		return nil
	}
	out := new(KeyVaultCertificateSyncStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyVaultSecretSync) DeepCopyInto(out *KeyVaultSecretSync) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyVaultSecretSync.
func (in *KeyVaultSecretSync) DeepCopy() *KeyVaultSecretSync {
	if in == nil {
		return nil
	}
	out := new(KeyVaultSecretSync)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KeyVaultSecretSync) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyVaultSecretSyncList) DeepCopyInto(out *KeyVaultSecretSyncList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KeyVaultSecretSync, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyVaultSecretSyncList.
func (in *KeyVaultSecretSyncList) DeepCopy() *KeyVaultSecretSyncList {
	if in == nil {
		return nil
	}
	out := new(KeyVaultSecretSyncList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KeyVaultSecretSyncList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyVaultSecretSyncSpec) DeepCopyInto(out *KeyVaultSecretSyncSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyVaultSecretSyncSpec.
func (in *KeyVaultSecretSyncSpec) DeepCopy() *KeyVaultSecretSyncSpec {
	if in == nil {
		return nil
	}
	out := new(KeyVaultSecretSyncSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyVaultSecretSyncStatus) DeepCopyInto(out *KeyVaultSecretSyncStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyVaultSecretSyncStatus.
func (in *KeyVaultSecretSyncStatus) DeepCopy() *KeyVaultSecretSyncStatus {
	if in == nil {
		return nil
	}
	out := new(KeyVaultSecretSyncStatus)
	in.DeepCopyInto(out)
	return out
}
