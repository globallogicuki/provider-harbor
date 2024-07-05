//go:build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PurgeAuditLog) DeepCopyInto(out *PurgeAuditLog) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PurgeAuditLog.
func (in *PurgeAuditLog) DeepCopy() *PurgeAuditLog {
	if in == nil {
		return nil
	}
	out := new(PurgeAuditLog)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PurgeAuditLog) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PurgeAuditLogInitParameters) DeepCopyInto(out *PurgeAuditLogInitParameters) {
	*out = *in
	if in.AuditRetentionHour != nil {
		in, out := &in.AuditRetentionHour, &out.AuditRetentionHour
		*out = new(float64)
		**out = **in
	}
	if in.IncludeOperations != nil {
		in, out := &in.IncludeOperations, &out.IncludeOperations
		*out = new(string)
		**out = **in
	}
	if in.Schedule != nil {
		in, out := &in.Schedule, &out.Schedule
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PurgeAuditLogInitParameters.
func (in *PurgeAuditLogInitParameters) DeepCopy() *PurgeAuditLogInitParameters {
	if in == nil {
		return nil
	}
	out := new(PurgeAuditLogInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PurgeAuditLogList) DeepCopyInto(out *PurgeAuditLogList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PurgeAuditLog, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PurgeAuditLogList.
func (in *PurgeAuditLogList) DeepCopy() *PurgeAuditLogList {
	if in == nil {
		return nil
	}
	out := new(PurgeAuditLogList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PurgeAuditLogList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PurgeAuditLogObservation) DeepCopyInto(out *PurgeAuditLogObservation) {
	*out = *in
	if in.AuditRetentionHour != nil {
		in, out := &in.AuditRetentionHour, &out.AuditRetentionHour
		*out = new(float64)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.IncludeOperations != nil {
		in, out := &in.IncludeOperations, &out.IncludeOperations
		*out = new(string)
		**out = **in
	}
	if in.Schedule != nil {
		in, out := &in.Schedule, &out.Schedule
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PurgeAuditLogObservation.
func (in *PurgeAuditLogObservation) DeepCopy() *PurgeAuditLogObservation {
	if in == nil {
		return nil
	}
	out := new(PurgeAuditLogObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PurgeAuditLogParameters) DeepCopyInto(out *PurgeAuditLogParameters) {
	*out = *in
	if in.AuditRetentionHour != nil {
		in, out := &in.AuditRetentionHour, &out.AuditRetentionHour
		*out = new(float64)
		**out = **in
	}
	if in.IncludeOperations != nil {
		in, out := &in.IncludeOperations, &out.IncludeOperations
		*out = new(string)
		**out = **in
	}
	if in.Schedule != nil {
		in, out := &in.Schedule, &out.Schedule
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PurgeAuditLogParameters.
func (in *PurgeAuditLogParameters) DeepCopy() *PurgeAuditLogParameters {
	if in == nil {
		return nil
	}
	out := new(PurgeAuditLogParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PurgeAuditLogSpec) DeepCopyInto(out *PurgeAuditLogSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PurgeAuditLogSpec.
func (in *PurgeAuditLogSpec) DeepCopy() *PurgeAuditLogSpec {
	if in == nil {
		return nil
	}
	out := new(PurgeAuditLogSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PurgeAuditLogStatus) DeepCopyInto(out *PurgeAuditLogStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PurgeAuditLogStatus.
func (in *PurgeAuditLogStatus) DeepCopy() *PurgeAuditLogStatus {
	if in == nil {
		return nil
	}
	out := new(PurgeAuditLogStatus)
	in.DeepCopyInto(out)
	return out
}