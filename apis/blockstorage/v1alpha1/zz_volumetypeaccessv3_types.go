// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
Copyright 2023 Jakob Schlagenhaufer, Jan Dittrich
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type VolumeTypeAccessV3InitParameters struct {

	// ID of the project to give access to. Changing this
	// creates a new resource.
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// The region in which to create the volume. If
	// omitted, the region argument of the provider is used. Changing this
	// creates a new quotaset.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// ID of the volume type to give access to. Changing
	// this creates a new resource.
	VolumeTypeID *string `json:"volumeTypeId,omitempty" tf:"volume_type_id,omitempty"`
}

type VolumeTypeAccessV3Observation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// ID of the project to give access to. Changing this
	// creates a new resource.
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// The region in which to create the volume. If
	// omitted, the region argument of the provider is used. Changing this
	// creates a new quotaset.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// ID of the volume type to give access to. Changing
	// this creates a new resource.
	VolumeTypeID *string `json:"volumeTypeId,omitempty" tf:"volume_type_id,omitempty"`
}

type VolumeTypeAccessV3Parameters struct {

	// ID of the project to give access to. Changing this
	// creates a new resource.
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// The region in which to create the volume. If
	// omitted, the region argument of the provider is used. Changing this
	// creates a new quotaset.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// ID of the volume type to give access to. Changing
	// this creates a new resource.
	// +kubebuilder:validation:Optional
	VolumeTypeID *string `json:"volumeTypeId,omitempty" tf:"volume_type_id,omitempty"`
}

// VolumeTypeAccessV3Spec defines the desired state of VolumeTypeAccessV3
type VolumeTypeAccessV3Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VolumeTypeAccessV3Parameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider VolumeTypeAccessV3InitParameters `json:"initProvider,omitempty"`
}

// VolumeTypeAccessV3Status defines the observed state of VolumeTypeAccessV3.
type VolumeTypeAccessV3Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VolumeTypeAccessV3Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VolumeTypeAccessV3 is the Schema for the VolumeTypeAccessV3s API. Manages a V3 volume type access resource within OpenStack.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,openstack}
type VolumeTypeAccessV3 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.projectId) || (has(self.initProvider) && has(self.initProvider.projectId))",message="spec.forProvider.projectId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.volumeTypeId) || (has(self.initProvider) && has(self.initProvider.volumeTypeId))",message="spec.forProvider.volumeTypeId is a required parameter"
	Spec   VolumeTypeAccessV3Spec   `json:"spec"`
	Status VolumeTypeAccessV3Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VolumeTypeAccessV3List contains a list of VolumeTypeAccessV3s
type VolumeTypeAccessV3List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VolumeTypeAccessV3 `json:"items"`
}

// Repository type metadata.
var (
	VolumeTypeAccessV3_Kind             = "VolumeTypeAccessV3"
	VolumeTypeAccessV3_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VolumeTypeAccessV3_Kind}.String()
	VolumeTypeAccessV3_KindAPIVersion   = VolumeTypeAccessV3_Kind + "." + CRDGroupVersion.String()
	VolumeTypeAccessV3_GroupVersionKind = CRDGroupVersion.WithKind(VolumeTypeAccessV3_Kind)
)

func init() {
	SchemeBuilder.Register(&VolumeTypeAccessV3{}, &VolumeTypeAccessV3List{})
}
