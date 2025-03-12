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

type VolumeTypeV3InitParameters struct {

	// Human-readable description of the port. Changing
	// this updates the description of an existing volume type.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Key/Value pairs of metadata for the volume type.
	// +mapType=granular
	ExtraSpecs map[string]*string `json:"extraSpecs,omitempty" tf:"extra_specs,omitempty"`

	// Whether the volume type is public. Changing
	// this updates the is_public of an existing volume type.
	IsPublic *bool `json:"isPublic,omitempty" tf:"is_public,omitempty"`

	// Name of the volume type.  Changing this
	// updates the name of an existing volume type.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The region in which to create the volume. If
	// omitted, the region argument of the provider is used. Changing this
	// creates a new quotaset.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type VolumeTypeV3Observation struct {

	// Human-readable description of the port. Changing
	// this updates the description of an existing volume type.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Key/Value pairs of metadata for the volume type.
	// +mapType=granular
	ExtraSpecs map[string]*string `json:"extraSpecs,omitempty" tf:"extra_specs,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Whether the volume type is public. Changing
	// this updates the is_public of an existing volume type.
	IsPublic *bool `json:"isPublic,omitempty" tf:"is_public,omitempty"`

	// Name of the volume type.  Changing this
	// updates the name of an existing volume type.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The region in which to create the volume. If
	// omitted, the region argument of the provider is used. Changing this
	// creates a new quotaset.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type VolumeTypeV3Parameters struct {

	// Human-readable description of the port. Changing
	// this updates the description of an existing volume type.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Key/Value pairs of metadata for the volume type.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	ExtraSpecs map[string]*string `json:"extraSpecs,omitempty" tf:"extra_specs,omitempty"`

	// Whether the volume type is public. Changing
	// this updates the is_public of an existing volume type.
	// +kubebuilder:validation:Optional
	IsPublic *bool `json:"isPublic,omitempty" tf:"is_public,omitempty"`

	// Name of the volume type.  Changing this
	// updates the name of an existing volume type.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The region in which to create the volume. If
	// omitted, the region argument of the provider is used. Changing this
	// creates a new quotaset.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

// VolumeTypeV3Spec defines the desired state of VolumeTypeV3
type VolumeTypeV3Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VolumeTypeV3Parameters `json:"forProvider"`
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
	InitProvider VolumeTypeV3InitParameters `json:"initProvider,omitempty"`
}

// VolumeTypeV3Status defines the observed state of VolumeTypeV3.
type VolumeTypeV3Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VolumeTypeV3Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// VolumeTypeV3 is the Schema for the VolumeTypeV3s API. Manages a V3 volume type resource within OpenStack.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,openstack}
type VolumeTypeV3 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   VolumeTypeV3Spec   `json:"spec"`
	Status VolumeTypeV3Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VolumeTypeV3List contains a list of VolumeTypeV3s
type VolumeTypeV3List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VolumeTypeV3 `json:"items"`
}

// Repository type metadata.
var (
	VolumeTypeV3_Kind             = "VolumeTypeV3"
	VolumeTypeV3_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VolumeTypeV3_Kind}.String()
	VolumeTypeV3_KindAPIVersion   = VolumeTypeV3_Kind + "." + CRDGroupVersion.String()
	VolumeTypeV3_GroupVersionKind = CRDGroupVersion.WithKind(VolumeTypeV3_Kind)
)

func init() {
	SchemeBuilder.Register(&VolumeTypeV3{}, &VolumeTypeV3List{})
}
