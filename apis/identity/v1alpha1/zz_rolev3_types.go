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

type RoleV3InitParameters struct {

	// The domain the role belongs to.
	DomainID *string `json:"domainId,omitempty" tf:"domain_id,omitempty"`

	// The name of the role.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The region in which to obtain the V3 Keystone client.
	// If omitted, the region argument of the provider is used. Changing this
	// creates a new Role.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type RoleV3Observation struct {

	// The domain the role belongs to.
	DomainID *string `json:"domainId,omitempty" tf:"domain_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the role.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The region in which to obtain the V3 Keystone client.
	// If omitted, the region argument of the provider is used. Changing this
	// creates a new Role.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type RoleV3Parameters struct {

	// The domain the role belongs to.
	// +kubebuilder:validation:Optional
	DomainID *string `json:"domainId,omitempty" tf:"domain_id,omitempty"`

	// The name of the role.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The region in which to obtain the V3 Keystone client.
	// If omitted, the region argument of the provider is used. Changing this
	// creates a new Role.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

// RoleV3Spec defines the desired state of RoleV3
type RoleV3Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RoleV3Parameters `json:"forProvider"`
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
	InitProvider RoleV3InitParameters `json:"initProvider,omitempty"`
}

// RoleV3Status defines the observed state of RoleV3.
type RoleV3Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RoleV3Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// RoleV3 is the Schema for the RoleV3s API. Manages a V3 Role resource within OpenStack Keystone.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,openstack}
type RoleV3 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   RoleV3Spec   `json:"spec"`
	Status RoleV3Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RoleV3List contains a list of RoleV3s
type RoleV3List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RoleV3 `json:"items"`
}

// Repository type metadata.
var (
	RoleV3_Kind             = "RoleV3"
	RoleV3_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RoleV3_Kind}.String()
	RoleV3_KindAPIVersion   = RoleV3_Kind + "." + CRDGroupVersion.String()
	RoleV3_GroupVersionKind = CRDGroupVersion.WithKind(RoleV3_Kind)
)

func init() {
	SchemeBuilder.Register(&RoleV3{}, &RoleV3List{})
}
