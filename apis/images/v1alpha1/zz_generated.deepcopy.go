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
func (in *ImageAccessAcceptV2) DeepCopyInto(out *ImageAccessAcceptV2) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageAccessAcceptV2.
func (in *ImageAccessAcceptV2) DeepCopy() *ImageAccessAcceptV2 {
	if in == nil {
		return nil
	}
	out := new(ImageAccessAcceptV2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ImageAccessAcceptV2) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageAccessAcceptV2InitParameters) DeepCopyInto(out *ImageAccessAcceptV2InitParameters) {
	*out = *in
	if in.ImageID != nil {
		in, out := &in.ImageID, &out.ImageID
		*out = new(string)
		**out = **in
	}
	if in.MemberID != nil {
		in, out := &in.MemberID, &out.MemberID
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageAccessAcceptV2InitParameters.
func (in *ImageAccessAcceptV2InitParameters) DeepCopy() *ImageAccessAcceptV2InitParameters {
	if in == nil {
		return nil
	}
	out := new(ImageAccessAcceptV2InitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageAccessAcceptV2List) DeepCopyInto(out *ImageAccessAcceptV2List) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ImageAccessAcceptV2, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageAccessAcceptV2List.
func (in *ImageAccessAcceptV2List) DeepCopy() *ImageAccessAcceptV2List {
	if in == nil {
		return nil
	}
	out := new(ImageAccessAcceptV2List)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ImageAccessAcceptV2List) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageAccessAcceptV2Observation) DeepCopyInto(out *ImageAccessAcceptV2Observation) {
	*out = *in
	if in.CreatedAt != nil {
		in, out := &in.CreatedAt, &out.CreatedAt
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.ImageID != nil {
		in, out := &in.ImageID, &out.ImageID
		*out = new(string)
		**out = **in
	}
	if in.MemberID != nil {
		in, out := &in.MemberID, &out.MemberID
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Schema != nil {
		in, out := &in.Schema, &out.Schema
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.UpdatedAt != nil {
		in, out := &in.UpdatedAt, &out.UpdatedAt
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageAccessAcceptV2Observation.
func (in *ImageAccessAcceptV2Observation) DeepCopy() *ImageAccessAcceptV2Observation {
	if in == nil {
		return nil
	}
	out := new(ImageAccessAcceptV2Observation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageAccessAcceptV2Parameters) DeepCopyInto(out *ImageAccessAcceptV2Parameters) {
	*out = *in
	if in.ImageID != nil {
		in, out := &in.ImageID, &out.ImageID
		*out = new(string)
		**out = **in
	}
	if in.MemberID != nil {
		in, out := &in.MemberID, &out.MemberID
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageAccessAcceptV2Parameters.
func (in *ImageAccessAcceptV2Parameters) DeepCopy() *ImageAccessAcceptV2Parameters {
	if in == nil {
		return nil
	}
	out := new(ImageAccessAcceptV2Parameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageAccessAcceptV2Spec) DeepCopyInto(out *ImageAccessAcceptV2Spec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageAccessAcceptV2Spec.
func (in *ImageAccessAcceptV2Spec) DeepCopy() *ImageAccessAcceptV2Spec {
	if in == nil {
		return nil
	}
	out := new(ImageAccessAcceptV2Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageAccessAcceptV2Status) DeepCopyInto(out *ImageAccessAcceptV2Status) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageAccessAcceptV2Status.
func (in *ImageAccessAcceptV2Status) DeepCopy() *ImageAccessAcceptV2Status {
	if in == nil {
		return nil
	}
	out := new(ImageAccessAcceptV2Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageAccessV2) DeepCopyInto(out *ImageAccessV2) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageAccessV2.
func (in *ImageAccessV2) DeepCopy() *ImageAccessV2 {
	if in == nil {
		return nil
	}
	out := new(ImageAccessV2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ImageAccessV2) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageAccessV2InitParameters) DeepCopyInto(out *ImageAccessV2InitParameters) {
	*out = *in
	if in.ImageID != nil {
		in, out := &in.ImageID, &out.ImageID
		*out = new(string)
		**out = **in
	}
	if in.ImageIDRef != nil {
		in, out := &in.ImageIDRef, &out.ImageIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ImageIDSelector != nil {
		in, out := &in.ImageIDSelector, &out.ImageIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.MemberID != nil {
		in, out := &in.MemberID, &out.MemberID
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageAccessV2InitParameters.
func (in *ImageAccessV2InitParameters) DeepCopy() *ImageAccessV2InitParameters {
	if in == nil {
		return nil
	}
	out := new(ImageAccessV2InitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageAccessV2List) DeepCopyInto(out *ImageAccessV2List) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ImageAccessV2, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageAccessV2List.
func (in *ImageAccessV2List) DeepCopy() *ImageAccessV2List {
	if in == nil {
		return nil
	}
	out := new(ImageAccessV2List)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ImageAccessV2List) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageAccessV2Observation) DeepCopyInto(out *ImageAccessV2Observation) {
	*out = *in
	if in.CreatedAt != nil {
		in, out := &in.CreatedAt, &out.CreatedAt
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.ImageID != nil {
		in, out := &in.ImageID, &out.ImageID
		*out = new(string)
		**out = **in
	}
	if in.MemberID != nil {
		in, out := &in.MemberID, &out.MemberID
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Schema != nil {
		in, out := &in.Schema, &out.Schema
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.UpdatedAt != nil {
		in, out := &in.UpdatedAt, &out.UpdatedAt
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageAccessV2Observation.
func (in *ImageAccessV2Observation) DeepCopy() *ImageAccessV2Observation {
	if in == nil {
		return nil
	}
	out := new(ImageAccessV2Observation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageAccessV2Parameters) DeepCopyInto(out *ImageAccessV2Parameters) {
	*out = *in
	if in.ImageID != nil {
		in, out := &in.ImageID, &out.ImageID
		*out = new(string)
		**out = **in
	}
	if in.ImageIDRef != nil {
		in, out := &in.ImageIDRef, &out.ImageIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ImageIDSelector != nil {
		in, out := &in.ImageIDSelector, &out.ImageIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.MemberID != nil {
		in, out := &in.MemberID, &out.MemberID
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageAccessV2Parameters.
func (in *ImageAccessV2Parameters) DeepCopy() *ImageAccessV2Parameters {
	if in == nil {
		return nil
	}
	out := new(ImageAccessV2Parameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageAccessV2Spec) DeepCopyInto(out *ImageAccessV2Spec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageAccessV2Spec.
func (in *ImageAccessV2Spec) DeepCopy() *ImageAccessV2Spec {
	if in == nil {
		return nil
	}
	out := new(ImageAccessV2Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageAccessV2Status) DeepCopyInto(out *ImageAccessV2Status) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageAccessV2Status.
func (in *ImageAccessV2Status) DeepCopy() *ImageAccessV2Status {
	if in == nil {
		return nil
	}
	out := new(ImageAccessV2Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageV2) DeepCopyInto(out *ImageV2) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageV2.
func (in *ImageV2) DeepCopy() *ImageV2 {
	if in == nil {
		return nil
	}
	out := new(ImageV2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ImageV2) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageV2InitParameters) DeepCopyInto(out *ImageV2InitParameters) {
	*out = *in
	if in.ContainerFormat != nil {
		in, out := &in.ContainerFormat, &out.ContainerFormat
		*out = new(string)
		**out = **in
	}
	if in.Decompress != nil {
		in, out := &in.Decompress, &out.Decompress
		*out = new(bool)
		**out = **in
	}
	if in.DiskFormat != nil {
		in, out := &in.DiskFormat, &out.DiskFormat
		*out = new(string)
		**out = **in
	}
	if in.Hidden != nil {
		in, out := &in.Hidden, &out.Hidden
		*out = new(bool)
		**out = **in
	}
	if in.ImageCachePath != nil {
		in, out := &in.ImageCachePath, &out.ImageCachePath
		*out = new(string)
		**out = **in
	}
	if in.ImageID != nil {
		in, out := &in.ImageID, &out.ImageID
		*out = new(string)
		**out = **in
	}
	if in.ImageSourcePasswordSecretRef != nil {
		in, out := &in.ImageSourcePasswordSecretRef, &out.ImageSourcePasswordSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.ImageSourceURL != nil {
		in, out := &in.ImageSourceURL, &out.ImageSourceURL
		*out = new(string)
		**out = **in
	}
	if in.ImageSourceUsername != nil {
		in, out := &in.ImageSourceUsername, &out.ImageSourceUsername
		*out = new(string)
		**out = **in
	}
	if in.LocalFilePath != nil {
		in, out := &in.LocalFilePath, &out.LocalFilePath
		*out = new(string)
		**out = **in
	}
	if in.MinDiskGb != nil {
		in, out := &in.MinDiskGb, &out.MinDiskGb
		*out = new(float64)
		**out = **in
	}
	if in.MinRAMMb != nil {
		in, out := &in.MinRAMMb, &out.MinRAMMb
		*out = new(float64)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Protected != nil {
		in, out := &in.Protected, &out.Protected
		*out = new(bool)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.VerifyChecksum != nil {
		in, out := &in.VerifyChecksum, &out.VerifyChecksum
		*out = new(bool)
		**out = **in
	}
	if in.Visibility != nil {
		in, out := &in.Visibility, &out.Visibility
		*out = new(string)
		**out = **in
	}
	if in.WebDownload != nil {
		in, out := &in.WebDownload, &out.WebDownload
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageV2InitParameters.
func (in *ImageV2InitParameters) DeepCopy() *ImageV2InitParameters {
	if in == nil {
		return nil
	}
	out := new(ImageV2InitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageV2List) DeepCopyInto(out *ImageV2List) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ImageV2, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageV2List.
func (in *ImageV2List) DeepCopy() *ImageV2List {
	if in == nil {
		return nil
	}
	out := new(ImageV2List)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ImageV2List) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageV2Observation) DeepCopyInto(out *ImageV2Observation) {
	*out = *in
	if in.Checksum != nil {
		in, out := &in.Checksum, &out.Checksum
		*out = new(string)
		**out = **in
	}
	if in.ContainerFormat != nil {
		in, out := &in.ContainerFormat, &out.ContainerFormat
		*out = new(string)
		**out = **in
	}
	if in.CreatedAt != nil {
		in, out := &in.CreatedAt, &out.CreatedAt
		*out = new(string)
		**out = **in
	}
	if in.Decompress != nil {
		in, out := &in.Decompress, &out.Decompress
		*out = new(bool)
		**out = **in
	}
	if in.DiskFormat != nil {
		in, out := &in.DiskFormat, &out.DiskFormat
		*out = new(string)
		**out = **in
	}
	if in.File != nil {
		in, out := &in.File, &out.File
		*out = new(string)
		**out = **in
	}
	if in.Hidden != nil {
		in, out := &in.Hidden, &out.Hidden
		*out = new(bool)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.ImageCachePath != nil {
		in, out := &in.ImageCachePath, &out.ImageCachePath
		*out = new(string)
		**out = **in
	}
	if in.ImageID != nil {
		in, out := &in.ImageID, &out.ImageID
		*out = new(string)
		**out = **in
	}
	if in.ImageSourceURL != nil {
		in, out := &in.ImageSourceURL, &out.ImageSourceURL
		*out = new(string)
		**out = **in
	}
	if in.ImageSourceUsername != nil {
		in, out := &in.ImageSourceUsername, &out.ImageSourceUsername
		*out = new(string)
		**out = **in
	}
	if in.LocalFilePath != nil {
		in, out := &in.LocalFilePath, &out.LocalFilePath
		*out = new(string)
		**out = **in
	}
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.MinDiskGb != nil {
		in, out := &in.MinDiskGb, &out.MinDiskGb
		*out = new(float64)
		**out = **in
	}
	if in.MinRAMMb != nil {
		in, out := &in.MinRAMMb, &out.MinRAMMb
		*out = new(float64)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Owner != nil {
		in, out := &in.Owner, &out.Owner
		*out = new(string)
		**out = **in
	}
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Protected != nil {
		in, out := &in.Protected, &out.Protected
		*out = new(bool)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Schema != nil {
		in, out := &in.Schema, &out.Schema
		*out = new(string)
		**out = **in
	}
	if in.SizeBytes != nil {
		in, out := &in.SizeBytes, &out.SizeBytes
		*out = new(float64)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.UpdatedAt != nil {
		in, out := &in.UpdatedAt, &out.UpdatedAt
		*out = new(string)
		**out = **in
	}
	if in.VerifyChecksum != nil {
		in, out := &in.VerifyChecksum, &out.VerifyChecksum
		*out = new(bool)
		**out = **in
	}
	if in.Visibility != nil {
		in, out := &in.Visibility, &out.Visibility
		*out = new(string)
		**out = **in
	}
	if in.WebDownload != nil {
		in, out := &in.WebDownload, &out.WebDownload
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageV2Observation.
func (in *ImageV2Observation) DeepCopy() *ImageV2Observation {
	if in == nil {
		return nil
	}
	out := new(ImageV2Observation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageV2Parameters) DeepCopyInto(out *ImageV2Parameters) {
	*out = *in
	if in.ContainerFormat != nil {
		in, out := &in.ContainerFormat, &out.ContainerFormat
		*out = new(string)
		**out = **in
	}
	if in.Decompress != nil {
		in, out := &in.Decompress, &out.Decompress
		*out = new(bool)
		**out = **in
	}
	if in.DiskFormat != nil {
		in, out := &in.DiskFormat, &out.DiskFormat
		*out = new(string)
		**out = **in
	}
	if in.Hidden != nil {
		in, out := &in.Hidden, &out.Hidden
		*out = new(bool)
		**out = **in
	}
	if in.ImageCachePath != nil {
		in, out := &in.ImageCachePath, &out.ImageCachePath
		*out = new(string)
		**out = **in
	}
	if in.ImageID != nil {
		in, out := &in.ImageID, &out.ImageID
		*out = new(string)
		**out = **in
	}
	if in.ImageSourcePasswordSecretRef != nil {
		in, out := &in.ImageSourcePasswordSecretRef, &out.ImageSourcePasswordSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.ImageSourceURL != nil {
		in, out := &in.ImageSourceURL, &out.ImageSourceURL
		*out = new(string)
		**out = **in
	}
	if in.ImageSourceUsername != nil {
		in, out := &in.ImageSourceUsername, &out.ImageSourceUsername
		*out = new(string)
		**out = **in
	}
	if in.LocalFilePath != nil {
		in, out := &in.LocalFilePath, &out.LocalFilePath
		*out = new(string)
		**out = **in
	}
	if in.MinDiskGb != nil {
		in, out := &in.MinDiskGb, &out.MinDiskGb
		*out = new(float64)
		**out = **in
	}
	if in.MinRAMMb != nil {
		in, out := &in.MinRAMMb, &out.MinRAMMb
		*out = new(float64)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Protected != nil {
		in, out := &in.Protected, &out.Protected
		*out = new(bool)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.VerifyChecksum != nil {
		in, out := &in.VerifyChecksum, &out.VerifyChecksum
		*out = new(bool)
		**out = **in
	}
	if in.Visibility != nil {
		in, out := &in.Visibility, &out.Visibility
		*out = new(string)
		**out = **in
	}
	if in.WebDownload != nil {
		in, out := &in.WebDownload, &out.WebDownload
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageV2Parameters.
func (in *ImageV2Parameters) DeepCopy() *ImageV2Parameters {
	if in == nil {
		return nil
	}
	out := new(ImageV2Parameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageV2Spec) DeepCopyInto(out *ImageV2Spec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageV2Spec.
func (in *ImageV2Spec) DeepCopy() *ImageV2Spec {
	if in == nil {
		return nil
	}
	out := new(ImageV2Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageV2Status) DeepCopyInto(out *ImageV2Status) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageV2Status.
func (in *ImageV2Status) DeepCopy() *ImageV2Status {
	if in == nil {
		return nil
	}
	out := new(ImageV2Status)
	in.DeepCopyInto(out)
	return out
}
