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

type UserMembershipV3InitParameters struct {

	// The UUID of group to which the user will be added.
	// Changing this creates a new user membership.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-openstack/apis/identity/v1alpha1.GroupV3
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	GroupID *string `json:"groupId,omitempty" tf:"group_id,omitempty"`

	// Reference to a GroupV3 in identity to populate groupId.
	// +kubebuilder:validation:Optional
	GroupIDRef *v1.Reference `json:"groupIdRef,omitempty" tf:"-"`

	// Selector for a GroupV3 in identity to populate groupId.
	// +kubebuilder:validation:Optional
	GroupIDSelector *v1.Selector `json:"groupIdSelector,omitempty" tf:"-"`

	// The region in which to obtain the V3 Identity client.
	// If omitted, the region argument of the provider is used.
	// Changing this creates a new user membership.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The UUID of user to use. Changing this creates a new user membership.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-openstack/apis/identity/v1alpha1.UserV3
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`

	// Reference to a UserV3 in identity to populate userId.
	// +kubebuilder:validation:Optional
	UserIDRef *v1.Reference `json:"userIdRef,omitempty" tf:"-"`

	// Selector for a UserV3 in identity to populate userId.
	// +kubebuilder:validation:Optional
	UserIDSelector *v1.Selector `json:"userIdSelector,omitempty" tf:"-"`
}

type UserMembershipV3Observation struct {

	// The UUID of group to which the user will be added.
	// Changing this creates a new user membership.
	GroupID *string `json:"groupId,omitempty" tf:"group_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The region in which to obtain the V3 Identity client.
	// If omitted, the region argument of the provider is used.
	// Changing this creates a new user membership.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The UUID of user to use. Changing this creates a new user membership.
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`
}

type UserMembershipV3Parameters struct {

	// The UUID of group to which the user will be added.
	// Changing this creates a new user membership.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-openstack/apis/identity/v1alpha1.GroupV3
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	GroupID *string `json:"groupId,omitempty" tf:"group_id,omitempty"`

	// Reference to a GroupV3 in identity to populate groupId.
	// +kubebuilder:validation:Optional
	GroupIDRef *v1.Reference `json:"groupIdRef,omitempty" tf:"-"`

	// Selector for a GroupV3 in identity to populate groupId.
	// +kubebuilder:validation:Optional
	GroupIDSelector *v1.Selector `json:"groupIdSelector,omitempty" tf:"-"`

	// The region in which to obtain the V3 Identity client.
	// If omitted, the region argument of the provider is used.
	// Changing this creates a new user membership.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The UUID of user to use. Changing this creates a new user membership.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-openstack/apis/identity/v1alpha1.UserV3
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`

	// Reference to a UserV3 in identity to populate userId.
	// +kubebuilder:validation:Optional
	UserIDRef *v1.Reference `json:"userIdRef,omitempty" tf:"-"`

	// Selector for a UserV3 in identity to populate userId.
	// +kubebuilder:validation:Optional
	UserIDSelector *v1.Selector `json:"userIdSelector,omitempty" tf:"-"`
}

// UserMembershipV3Spec defines the desired state of UserMembershipV3
type UserMembershipV3Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     UserMembershipV3Parameters `json:"forProvider"`
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
	InitProvider UserMembershipV3InitParameters `json:"initProvider,omitempty"`
}

// UserMembershipV3Status defines the observed state of UserMembershipV3.
type UserMembershipV3Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        UserMembershipV3Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// UserMembershipV3 is the Schema for the UserMembershipV3s API. Manages a user membership to group V3 resource within OpenStack.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,openstack}
type UserMembershipV3 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              UserMembershipV3Spec   `json:"spec"`
	Status            UserMembershipV3Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UserMembershipV3List contains a list of UserMembershipV3s
type UserMembershipV3List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UserMembershipV3 `json:"items"`
}

// Repository type metadata.
var (
	UserMembershipV3_Kind             = "UserMembershipV3"
	UserMembershipV3_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: UserMembershipV3_Kind}.String()
	UserMembershipV3_KindAPIVersion   = UserMembershipV3_Kind + "." + CRDGroupVersion.String()
	UserMembershipV3_GroupVersionKind = CRDGroupVersion.WithKind(UserMembershipV3_Kind)
)

func init() {
	SchemeBuilder.Register(&UserMembershipV3{}, &UserMembershipV3List{})
}
