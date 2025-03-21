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

type GroupV3InitParameters struct {

	// A description of the group.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The domain the group belongs to.
	DomainID *string `json:"domainId,omitempty" tf:"domain_id,omitempty"`

	// The name of the group.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The region in which to obtain the V3 Keystone client.
	// If omitted, the region argument of the provider is used. Changing this
	// creates a new group.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type GroupV3Observation struct {

	// A description of the group.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The domain the group belongs to.
	DomainID *string `json:"domainId,omitempty" tf:"domain_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the group.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The region in which to obtain the V3 Keystone client.
	// If omitted, the region argument of the provider is used. Changing this
	// creates a new group.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type GroupV3Parameters struct {

	// A description of the group.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The domain the group belongs to.
	// +kubebuilder:validation:Optional
	DomainID *string `json:"domainId,omitempty" tf:"domain_id,omitempty"`

	// The name of the group.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The region in which to obtain the V3 Keystone client.
	// If omitted, the region argument of the provider is used. Changing this
	// creates a new group.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

// GroupV3Spec defines the desired state of GroupV3
type GroupV3Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GroupV3Parameters `json:"forProvider"`
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
	InitProvider GroupV3InitParameters `json:"initProvider,omitempty"`
}

// GroupV3Status defines the observed state of GroupV3.
type GroupV3Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GroupV3Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// GroupV3 is the Schema for the GroupV3s API. Manages a V3 group resource within OpenStack Keystone.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,openstack}
type GroupV3 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   GroupV3Spec   `json:"spec"`
	Status GroupV3Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GroupV3List contains a list of GroupV3s
type GroupV3List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GroupV3 `json:"items"`
}

// Repository type metadata.
var (
	GroupV3_Kind             = "GroupV3"
	GroupV3_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GroupV3_Kind}.String()
	GroupV3_KindAPIVersion   = GroupV3_Kind + "." + CRDGroupVersion.String()
	GroupV3_GroupVersionKind = CRDGroupVersion.WithKind(GroupV3_Kind)
)

func init() {
	SchemeBuilder.Register(&GroupV3{}, &GroupV3List{})
}
