//go:build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FiltersInitParameters) DeepCopyInto(out *FiltersInitParameters) {
	*out = *in
	if in.Decoration != nil {
		in, out := &in.Decoration, &out.Decoration
		*out = new(string)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Resource != nil {
		in, out := &in.Resource, &out.Resource
		*out = new(string)
		**out = **in
	}
	if in.Tag != nil {
		in, out := &in.Tag, &out.Tag
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FiltersInitParameters.
func (in *FiltersInitParameters) DeepCopy() *FiltersInitParameters {
	if in == nil {
		return nil
	}
	out := new(FiltersInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FiltersObservation) DeepCopyInto(out *FiltersObservation) {
	*out = *in
	if in.Decoration != nil {
		in, out := &in.Decoration, &out.Decoration
		*out = new(string)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Resource != nil {
		in, out := &in.Resource, &out.Resource
		*out = new(string)
		**out = **in
	}
	if in.Tag != nil {
		in, out := &in.Tag, &out.Tag
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FiltersObservation.
func (in *FiltersObservation) DeepCopy() *FiltersObservation {
	if in == nil {
		return nil
	}
	out := new(FiltersObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FiltersParameters) DeepCopyInto(out *FiltersParameters) {
	*out = *in
	if in.Decoration != nil {
		in, out := &in.Decoration, &out.Decoration
		*out = new(string)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Resource != nil {
		in, out := &in.Resource, &out.Resource
		*out = new(string)
		**out = **in
	}
	if in.Tag != nil {
		in, out := &in.Tag, &out.Tag
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FiltersParameters.
func (in *FiltersParameters) DeepCopy() *FiltersParameters {
	if in == nil {
		return nil
	}
	out := new(FiltersParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Registry) DeepCopyInto(out *Registry) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Registry.
func (in *Registry) DeepCopy() *Registry {
	if in == nil {
		return nil
	}
	out := new(Registry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Registry) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryInitParameters) DeepCopyInto(out *RegistryInitParameters) {
	*out = *in
	if in.AccessID != nil {
		in, out := &in.AccessID, &out.AccessID
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.EndpointURL != nil {
		in, out := &in.EndpointURL, &out.EndpointURL
		*out = new(string)
		**out = **in
	}
	if in.Insecure != nil {
		in, out := &in.Insecure, &out.Insecure
		*out = new(bool)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ProviderName != nil {
		in, out := &in.ProviderName, &out.ProviderName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryInitParameters.
func (in *RegistryInitParameters) DeepCopy() *RegistryInitParameters {
	if in == nil {
		return nil
	}
	out := new(RegistryInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryList) DeepCopyInto(out *RegistryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Registry, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryList.
func (in *RegistryList) DeepCopy() *RegistryList {
	if in == nil {
		return nil
	}
	out := new(RegistryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RegistryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryObservation) DeepCopyInto(out *RegistryObservation) {
	*out = *in
	if in.AccessID != nil {
		in, out := &in.AccessID, &out.AccessID
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.EndpointURL != nil {
		in, out := &in.EndpointURL, &out.EndpointURL
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Insecure != nil {
		in, out := &in.Insecure, &out.Insecure
		*out = new(bool)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ProviderName != nil {
		in, out := &in.ProviderName, &out.ProviderName
		*out = new(string)
		**out = **in
	}
	if in.RegistryID != nil {
		in, out := &in.RegistryID, &out.RegistryID
		*out = new(float64)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryObservation.
func (in *RegistryObservation) DeepCopy() *RegistryObservation {
	if in == nil {
		return nil
	}
	out := new(RegistryObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryParameters) DeepCopyInto(out *RegistryParameters) {
	*out = *in
	if in.AccessID != nil {
		in, out := &in.AccessID, &out.AccessID
		*out = new(string)
		**out = **in
	}
	if in.AccessSecretSecretRef != nil {
		in, out := &in.AccessSecretSecretRef, &out.AccessSecretSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.EndpointURL != nil {
		in, out := &in.EndpointURL, &out.EndpointURL
		*out = new(string)
		**out = **in
	}
	if in.Insecure != nil {
		in, out := &in.Insecure, &out.Insecure
		*out = new(bool)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ProviderName != nil {
		in, out := &in.ProviderName, &out.ProviderName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryParameters.
func (in *RegistryParameters) DeepCopy() *RegistryParameters {
	if in == nil {
		return nil
	}
	out := new(RegistryParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistrySpec) DeepCopyInto(out *RegistrySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistrySpec.
func (in *RegistrySpec) DeepCopy() *RegistrySpec {
	if in == nil {
		return nil
	}
	out := new(RegistrySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryStatus) DeepCopyInto(out *RegistryStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryStatus.
func (in *RegistryStatus) DeepCopy() *RegistryStatus {
	if in == nil {
		return nil
	}
	out := new(RegistryStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Replication) DeepCopyInto(out *Replication) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Replication.
func (in *Replication) DeepCopy() *Replication {
	if in == nil {
		return nil
	}
	out := new(Replication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Replication) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationInitParameters) DeepCopyInto(out *ReplicationInitParameters) {
	*out = *in
	if in.Action != nil {
		in, out := &in.Action, &out.Action
		*out = new(string)
		**out = **in
	}
	if in.CopyByChunk != nil {
		in, out := &in.CopyByChunk, &out.CopyByChunk
		*out = new(bool)
		**out = **in
	}
	if in.Deletion != nil {
		in, out := &in.Deletion, &out.Deletion
		*out = new(bool)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DestNamespace != nil {
		in, out := &in.DestNamespace, &out.DestNamespace
		*out = new(string)
		**out = **in
	}
	if in.DestNamespaceReplace != nil {
		in, out := &in.DestNamespaceReplace, &out.DestNamespaceReplace
		*out = new(float64)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.ExecuteOnChanged != nil {
		in, out := &in.ExecuteOnChanged, &out.ExecuteOnChanged
		*out = new(bool)
		**out = **in
	}
	if in.Filters != nil {
		in, out := &in.Filters, &out.Filters
		*out = make([]FiltersInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Override != nil {
		in, out := &in.Override, &out.Override
		*out = new(bool)
		**out = **in
	}
	if in.Schedule != nil {
		in, out := &in.Schedule, &out.Schedule
		*out = new(string)
		**out = **in
	}
	if in.Speed != nil {
		in, out := &in.Speed, &out.Speed
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationInitParameters.
func (in *ReplicationInitParameters) DeepCopy() *ReplicationInitParameters {
	if in == nil {
		return nil
	}
	out := new(ReplicationInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationList) DeepCopyInto(out *ReplicationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Replication, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationList.
func (in *ReplicationList) DeepCopy() *ReplicationList {
	if in == nil {
		return nil
	}
	out := new(ReplicationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReplicationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationObservation) DeepCopyInto(out *ReplicationObservation) {
	*out = *in
	if in.Action != nil {
		in, out := &in.Action, &out.Action
		*out = new(string)
		**out = **in
	}
	if in.CopyByChunk != nil {
		in, out := &in.CopyByChunk, &out.CopyByChunk
		*out = new(bool)
		**out = **in
	}
	if in.Deletion != nil {
		in, out := &in.Deletion, &out.Deletion
		*out = new(bool)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DestNamespace != nil {
		in, out := &in.DestNamespace, &out.DestNamespace
		*out = new(string)
		**out = **in
	}
	if in.DestNamespaceReplace != nil {
		in, out := &in.DestNamespaceReplace, &out.DestNamespaceReplace
		*out = new(float64)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.ExecuteOnChanged != nil {
		in, out := &in.ExecuteOnChanged, &out.ExecuteOnChanged
		*out = new(bool)
		**out = **in
	}
	if in.Filters != nil {
		in, out := &in.Filters, &out.Filters
		*out = make([]FiltersObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Override != nil {
		in, out := &in.Override, &out.Override
		*out = new(bool)
		**out = **in
	}
	if in.RegistryID != nil {
		in, out := &in.RegistryID, &out.RegistryID
		*out = new(float64)
		**out = **in
	}
	if in.ReplicationPolicyID != nil {
		in, out := &in.ReplicationPolicyID, &out.ReplicationPolicyID
		*out = new(float64)
		**out = **in
	}
	if in.Schedule != nil {
		in, out := &in.Schedule, &out.Schedule
		*out = new(string)
		**out = **in
	}
	if in.Speed != nil {
		in, out := &in.Speed, &out.Speed
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationObservation.
func (in *ReplicationObservation) DeepCopy() *ReplicationObservation {
	if in == nil {
		return nil
	}
	out := new(ReplicationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationParameters) DeepCopyInto(out *ReplicationParameters) {
	*out = *in
	if in.Action != nil {
		in, out := &in.Action, &out.Action
		*out = new(string)
		**out = **in
	}
	if in.CopyByChunk != nil {
		in, out := &in.CopyByChunk, &out.CopyByChunk
		*out = new(bool)
		**out = **in
	}
	if in.Deletion != nil {
		in, out := &in.Deletion, &out.Deletion
		*out = new(bool)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DestNamespace != nil {
		in, out := &in.DestNamespace, &out.DestNamespace
		*out = new(string)
		**out = **in
	}
	if in.DestNamespaceReplace != nil {
		in, out := &in.DestNamespaceReplace, &out.DestNamespaceReplace
		*out = new(float64)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.ExecuteOnChanged != nil {
		in, out := &in.ExecuteOnChanged, &out.ExecuteOnChanged
		*out = new(bool)
		**out = **in
	}
	if in.Filters != nil {
		in, out := &in.Filters, &out.Filters
		*out = make([]FiltersParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Override != nil {
		in, out := &in.Override, &out.Override
		*out = new(bool)
		**out = **in
	}
	if in.RegistryID != nil {
		in, out := &in.RegistryID, &out.RegistryID
		*out = new(float64)
		**out = **in
	}
	if in.RegistryIDRef != nil {
		in, out := &in.RegistryIDRef, &out.RegistryIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.RegistryIDSelector != nil {
		in, out := &in.RegistryIDSelector, &out.RegistryIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Schedule != nil {
		in, out := &in.Schedule, &out.Schedule
		*out = new(string)
		**out = **in
	}
	if in.Speed != nil {
		in, out := &in.Speed, &out.Speed
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationParameters.
func (in *ReplicationParameters) DeepCopy() *ReplicationParameters {
	if in == nil {
		return nil
	}
	out := new(ReplicationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationSpec) DeepCopyInto(out *ReplicationSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationSpec.
func (in *ReplicationSpec) DeepCopy() *ReplicationSpec {
	if in == nil {
		return nil
	}
	out := new(ReplicationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationStatus) DeepCopyInto(out *ReplicationStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationStatus.
func (in *ReplicationStatus) DeepCopy() *ReplicationStatus {
	if in == nil {
		return nil
	}
	out := new(ReplicationStatus)
	in.DeepCopyInto(out)
	return out
}
