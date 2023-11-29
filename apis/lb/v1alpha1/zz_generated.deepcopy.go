//go:build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
Copyright 2023 Jakob Schlagenhaufer, Jan Dittrich
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QuotaV2) DeepCopyInto(out *QuotaV2) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QuotaV2.
func (in *QuotaV2) DeepCopy() *QuotaV2 {
	if in == nil {
		return nil
	}
	out := new(QuotaV2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *QuotaV2) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QuotaV2InitParameters) DeepCopyInto(out *QuotaV2InitParameters) {
	*out = *in
	if in.HealthMonitor != nil {
		in, out := &in.HealthMonitor, &out.HealthMonitor
		*out = new(float64)
		**out = **in
	}
	if in.L7Policy != nil {
		in, out := &in.L7Policy, &out.L7Policy
		*out = new(float64)
		**out = **in
	}
	if in.L7Rule != nil {
		in, out := &in.L7Rule, &out.L7Rule
		*out = new(float64)
		**out = **in
	}
	if in.Listener != nil {
		in, out := &in.Listener, &out.Listener
		*out = new(float64)
		**out = **in
	}
	if in.Loadbalancer != nil {
		in, out := &in.Loadbalancer, &out.Loadbalancer
		*out = new(float64)
		**out = **in
	}
	if in.Member != nil {
		in, out := &in.Member, &out.Member
		*out = new(float64)
		**out = **in
	}
	if in.Pool != nil {
		in, out := &in.Pool, &out.Pool
		*out = new(float64)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QuotaV2InitParameters.
func (in *QuotaV2InitParameters) DeepCopy() *QuotaV2InitParameters {
	if in == nil {
		return nil
	}
	out := new(QuotaV2InitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QuotaV2List) DeepCopyInto(out *QuotaV2List) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]QuotaV2, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QuotaV2List.
func (in *QuotaV2List) DeepCopy() *QuotaV2List {
	if in == nil {
		return nil
	}
	out := new(QuotaV2List)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *QuotaV2List) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QuotaV2Observation) DeepCopyInto(out *QuotaV2Observation) {
	*out = *in
	if in.HealthMonitor != nil {
		in, out := &in.HealthMonitor, &out.HealthMonitor
		*out = new(float64)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.L7Policy != nil {
		in, out := &in.L7Policy, &out.L7Policy
		*out = new(float64)
		**out = **in
	}
	if in.L7Rule != nil {
		in, out := &in.L7Rule, &out.L7Rule
		*out = new(float64)
		**out = **in
	}
	if in.Listener != nil {
		in, out := &in.Listener, &out.Listener
		*out = new(float64)
		**out = **in
	}
	if in.Loadbalancer != nil {
		in, out := &in.Loadbalancer, &out.Loadbalancer
		*out = new(float64)
		**out = **in
	}
	if in.Member != nil {
		in, out := &in.Member, &out.Member
		*out = new(float64)
		**out = **in
	}
	if in.Pool != nil {
		in, out := &in.Pool, &out.Pool
		*out = new(float64)
		**out = **in
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QuotaV2Observation.
func (in *QuotaV2Observation) DeepCopy() *QuotaV2Observation {
	if in == nil {
		return nil
	}
	out := new(QuotaV2Observation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QuotaV2Parameters) DeepCopyInto(out *QuotaV2Parameters) {
	*out = *in
	if in.HealthMonitor != nil {
		in, out := &in.HealthMonitor, &out.HealthMonitor
		*out = new(float64)
		**out = **in
	}
	if in.L7Policy != nil {
		in, out := &in.L7Policy, &out.L7Policy
		*out = new(float64)
		**out = **in
	}
	if in.L7Rule != nil {
		in, out := &in.L7Rule, &out.L7Rule
		*out = new(float64)
		**out = **in
	}
	if in.Listener != nil {
		in, out := &in.Listener, &out.Listener
		*out = new(float64)
		**out = **in
	}
	if in.Loadbalancer != nil {
		in, out := &in.Loadbalancer, &out.Loadbalancer
		*out = new(float64)
		**out = **in
	}
	if in.Member != nil {
		in, out := &in.Member, &out.Member
		*out = new(float64)
		**out = **in
	}
	if in.Pool != nil {
		in, out := &in.Pool, &out.Pool
		*out = new(float64)
		**out = **in
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(string)
		**out = **in
	}
	if in.ProjectIDRef != nil {
		in, out := &in.ProjectIDRef, &out.ProjectIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ProjectIDSelector != nil {
		in, out := &in.ProjectIDSelector, &out.ProjectIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QuotaV2Parameters.
func (in *QuotaV2Parameters) DeepCopy() *QuotaV2Parameters {
	if in == nil {
		return nil
	}
	out := new(QuotaV2Parameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QuotaV2Spec) DeepCopyInto(out *QuotaV2Spec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QuotaV2Spec.
func (in *QuotaV2Spec) DeepCopy() *QuotaV2Spec {
	if in == nil {
		return nil
	}
	out := new(QuotaV2Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QuotaV2Status) DeepCopyInto(out *QuotaV2Status) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QuotaV2Status.
func (in *QuotaV2Status) DeepCopy() *QuotaV2Status {
	if in == nil {
		return nil
	}
	out := new(QuotaV2Status)
	in.DeepCopyInto(out)
	return out
}
