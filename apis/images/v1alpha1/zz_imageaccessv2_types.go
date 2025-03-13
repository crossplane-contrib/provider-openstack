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

type ImageAccessV2InitParameters struct {

	// The image ID.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-openstack/apis/images/v1alpha1.ImageV2
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	ImageID *string `json:"imageId,omitempty" tf:"image_id,omitempty"`

	// Reference to a ImageV2 in images to populate imageId.
	// +kubebuilder:validation:Optional
	ImageIDRef *v1.Reference `json:"imageIdRef,omitempty" tf:"-"`

	// Selector for a ImageV2 in images to populate imageId.
	// +kubebuilder:validation:Optional
	ImageIDSelector *v1.Selector `json:"imageIdSelector,omitempty" tf:"-"`

	// The member ID, e.g. the target project ID.
	MemberID *string `json:"memberId,omitempty" tf:"member_id,omitempty"`

	// The region in which to obtain the V2 Glance client.
	// A Glance client is needed to manage Image members. If omitted, the region
	// argument of the provider is used. Changing this creates a new resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The member proposal status. Optional if admin wants to
	// force the member proposal acceptance. Can either be accepted, rejected or
	// pending. Defaults to pending. Foridden for non-admin users.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type ImageAccessV2Observation struct {

	// The date the image access was created.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The image ID.
	ImageID *string `json:"imageId,omitempty" tf:"image_id,omitempty"`

	// The member ID, e.g. the target project ID.
	MemberID *string `json:"memberId,omitempty" tf:"member_id,omitempty"`

	// The region in which to obtain the V2 Glance client.
	// A Glance client is needed to manage Image members. If omitted, the region
	// argument of the provider is used. Changing this creates a new resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The member schema.
	Schema *string `json:"schema,omitempty" tf:"schema,omitempty"`

	// The member proposal status. Optional if admin wants to
	// force the member proposal acceptance. Can either be accepted, rejected or
	// pending. Defaults to pending. Foridden for non-admin users.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// The date the image access was last updated.
	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`
}

type ImageAccessV2Parameters struct {

	// The image ID.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-openstack/apis/images/v1alpha1.ImageV2
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ImageID *string `json:"imageId,omitempty" tf:"image_id,omitempty"`

	// Reference to a ImageV2 in images to populate imageId.
	// +kubebuilder:validation:Optional
	ImageIDRef *v1.Reference `json:"imageIdRef,omitempty" tf:"-"`

	// Selector for a ImageV2 in images to populate imageId.
	// +kubebuilder:validation:Optional
	ImageIDSelector *v1.Selector `json:"imageIdSelector,omitempty" tf:"-"`

	// The member ID, e.g. the target project ID.
	// +kubebuilder:validation:Optional
	MemberID *string `json:"memberId,omitempty" tf:"member_id,omitempty"`

	// The region in which to obtain the V2 Glance client.
	// A Glance client is needed to manage Image members. If omitted, the region
	// argument of the provider is used. Changing this creates a new resource.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The member proposal status. Optional if admin wants to
	// force the member proposal acceptance. Can either be accepted, rejected or
	// pending. Defaults to pending. Foridden for non-admin users.
	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

// ImageAccessV2Spec defines the desired state of ImageAccessV2
type ImageAccessV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ImageAccessV2Parameters `json:"forProvider"`
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
	InitProvider ImageAccessV2InitParameters `json:"initProvider,omitempty"`
}

// ImageAccessV2Status defines the observed state of ImageAccessV2.
type ImageAccessV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ImageAccessV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ImageAccessV2 is the Schema for the ImageAccessV2s API. Manages a V2 Image member resource within OpenStack Glance.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,openstack}
type ImageAccessV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.memberId) || (has(self.initProvider) && has(self.initProvider.memberId))",message="spec.forProvider.memberId is a required parameter"
	Spec   ImageAccessV2Spec   `json:"spec"`
	Status ImageAccessV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ImageAccessV2List contains a list of ImageAccessV2s
type ImageAccessV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ImageAccessV2 `json:"items"`
}

// Repository type metadata.
var (
	ImageAccessV2_Kind             = "ImageAccessV2"
	ImageAccessV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ImageAccessV2_Kind}.String()
	ImageAccessV2_KindAPIVersion   = ImageAccessV2_Kind + "." + CRDGroupVersion.String()
	ImageAccessV2_GroupVersionKind = CRDGroupVersion.WithKind(ImageAccessV2_Kind)
)

func init() {
	SchemeBuilder.Register(&ImageAccessV2{}, &ImageAccessV2List{})
}
