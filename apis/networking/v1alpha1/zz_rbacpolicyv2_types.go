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

type RbacPolicyV2InitParameters struct {

	// Action for the RBAC policy. Can either be
	// access_as_external or access_as_shared.
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// The ID of the object_type resource. An
	// object_type of network returns a network ID and an object_type of
	// qos_policy returns a QoS ID.
	ObjectID *string `json:"objectId,omitempty" tf:"object_id,omitempty"`

	// The type of the object that the RBAC policy
	// affects. Can be one of the following: address_scope, address_group,
	// network, qos_policy, security_group or subnetpool.
	ObjectType *string `json:"objectType,omitempty" tf:"object_type,omitempty"`

	// The region in which to obtain the V2 networking client.
	// A networking client is needed to configure a routing entry on a subnet. If omitted, the
	// region argument of the provider is used. Changing this creates a new
	// routing entry.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The ID of the tenant to which the RBAC policy
	// will be enforced.
	TargetTenant *string `json:"targetTenant,omitempty" tf:"target_tenant,omitempty"`
}

type RbacPolicyV2Observation struct {

	// Action for the RBAC policy. Can either be
	// access_as_external or access_as_shared.
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the object_type resource. An
	// object_type of network returns a network ID and an object_type of
	// qos_policy returns a QoS ID.
	ObjectID *string `json:"objectId,omitempty" tf:"object_id,omitempty"`

	// The type of the object that the RBAC policy
	// affects. Can be one of the following: address_scope, address_group,
	// network, qos_policy, security_group or subnetpool.
	ObjectType *string `json:"objectType,omitempty" tf:"object_type,omitempty"`

	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// The region in which to obtain the V2 networking client.
	// A networking client is needed to configure a routing entry on a subnet. If omitted, the
	// region argument of the provider is used. Changing this creates a new
	// routing entry.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The ID of the tenant to which the RBAC policy
	// will be enforced.
	TargetTenant *string `json:"targetTenant,omitempty" tf:"target_tenant,omitempty"`
}

type RbacPolicyV2Parameters struct {

	// Action for the RBAC policy. Can either be
	// access_as_external or access_as_shared.
	// +kubebuilder:validation:Optional
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// The ID of the object_type resource. An
	// object_type of network returns a network ID and an object_type of
	// qos_policy returns a QoS ID.
	// +kubebuilder:validation:Optional
	ObjectID *string `json:"objectId,omitempty" tf:"object_id,omitempty"`

	// The type of the object that the RBAC policy
	// affects. Can be one of the following: address_scope, address_group,
	// network, qos_policy, security_group or subnetpool.
	// +kubebuilder:validation:Optional
	ObjectType *string `json:"objectType,omitempty" tf:"object_type,omitempty"`

	// The region in which to obtain the V2 networking client.
	// A networking client is needed to configure a routing entry on a subnet. If omitted, the
	// region argument of the provider is used. Changing this creates a new
	// routing entry.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The ID of the tenant to which the RBAC policy
	// will be enforced.
	// +kubebuilder:validation:Optional
	TargetTenant *string `json:"targetTenant,omitempty" tf:"target_tenant,omitempty"`
}

// RbacPolicyV2Spec defines the desired state of RbacPolicyV2
type RbacPolicyV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RbacPolicyV2Parameters `json:"forProvider"`
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
	InitProvider RbacPolicyV2InitParameters `json:"initProvider,omitempty"`
}

// RbacPolicyV2Status defines the observed state of RbacPolicyV2.
type RbacPolicyV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RbacPolicyV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// RbacPolicyV2 is the Schema for the RbacPolicyV2s API. Creates an RBAC policy for an OpenStack V2 resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,openstack}
type RbacPolicyV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.action) || (has(self.initProvider) && has(self.initProvider.action))",message="spec.forProvider.action is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.objectId) || (has(self.initProvider) && has(self.initProvider.objectId))",message="spec.forProvider.objectId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.objectType) || (has(self.initProvider) && has(self.initProvider.objectType))",message="spec.forProvider.objectType is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.targetTenant) || (has(self.initProvider) && has(self.initProvider.targetTenant))",message="spec.forProvider.targetTenant is a required parameter"
	Spec   RbacPolicyV2Spec   `json:"spec"`
	Status RbacPolicyV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RbacPolicyV2List contains a list of RbacPolicyV2s
type RbacPolicyV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RbacPolicyV2 `json:"items"`
}

// Repository type metadata.
var (
	RbacPolicyV2_Kind             = "RbacPolicyV2"
	RbacPolicyV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RbacPolicyV2_Kind}.String()
	RbacPolicyV2_KindAPIVersion   = RbacPolicyV2_Kind + "." + CRDGroupVersion.String()
	RbacPolicyV2_GroupVersionKind = CRDGroupVersion.WithKind(RbacPolicyV2_Kind)
)

func init() {
	SchemeBuilder.Register(&RbacPolicyV2{}, &RbacPolicyV2List{})
}
