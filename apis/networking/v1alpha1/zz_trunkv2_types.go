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

type SubPortInitParameters struct {

	// The ID of the port to be made a subport of the trunk.
	PortID *string `json:"portId,omitempty" tf:"port_id,omitempty"`

	// The numeric id of the subport segment.
	SegmentationID *float64 `json:"segmentationId,omitempty" tf:"segmentation_id,omitempty"`

	// The segmentation technology to use, e.g., "vlan".
	SegmentationType *string `json:"segmentationType,omitempty" tf:"segmentation_type,omitempty"`
}

type SubPortObservation struct {

	// The ID of the port to be made a subport of the trunk.
	PortID *string `json:"portId,omitempty" tf:"port_id,omitempty"`

	// The numeric id of the subport segment.
	SegmentationID *float64 `json:"segmentationId,omitempty" tf:"segmentation_id,omitempty"`

	// The segmentation technology to use, e.g., "vlan".
	SegmentationType *string `json:"segmentationType,omitempty" tf:"segmentation_type,omitempty"`
}

type SubPortParameters struct {

	// The ID of the port to be made a subport of the trunk.
	// +kubebuilder:validation:Optional
	PortID *string `json:"portId" tf:"port_id,omitempty"`

	// The numeric id of the subport segment.
	// +kubebuilder:validation:Optional
	SegmentationID *float64 `json:"segmentationId" tf:"segmentation_id,omitempty"`

	// The segmentation technology to use, e.g., "vlan".
	// +kubebuilder:validation:Optional
	SegmentationType *string `json:"segmentationType" tf:"segmentation_type,omitempty"`
}

type TrunkV2InitParameters struct {

	// Administrative up/down status for the trunk
	// (must be "true" or "false" if provided). Changing this updates the
	// admin_state_up of an existing trunk.
	AdminStateUp *bool `json:"adminStateUp,omitempty" tf:"admin_state_up,omitempty"`

	// Human-readable description of the trunk. Changing this
	// updates the name of the existing trunk.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A unique name for the trunk. Changing this
	// updates the name of an existing trunk.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the port to be used as the parent port of the
	// trunk. This is the port that should be used as the compute instance network
	// port. Changing this creates a new trunk.
	PortID *string `json:"portId,omitempty" tf:"port_id,omitempty"`

	// The region in which to obtain the V2 networking client.
	// A networking client is needed to create a trunk. If omitted, the
	// region argument of the provider is used. Changing this creates a new
	// trunk.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The set of ports that will be made subports of the trunk.
	// The structure of each subport is described below.
	SubPort []SubPortInitParameters `json:"subPort,omitempty" tf:"sub_port,omitempty"`

	// A set of string tags for the port.
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The owner of the Trunk. Required if admin wants
	// to create a trunk on behalf of another tenant. Changing this creates a new trunk.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type TrunkV2Observation struct {

	// Administrative up/down status for the trunk
	// (must be "true" or "false" if provided). Changing this updates the
	// admin_state_up of an existing trunk.
	AdminStateUp *bool `json:"adminStateUp,omitempty" tf:"admin_state_up,omitempty"`

	// The collection of tags assigned on the trunk, which have been
	// explicitly and implicitly added.
	AllTags []*string `json:"allTags,omitempty" tf:"all_tags,omitempty"`

	// Human-readable description of the trunk. Changing this
	// updates the name of the existing trunk.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A unique name for the trunk. Changing this
	// updates the name of an existing trunk.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the port to be used as the parent port of the
	// trunk. This is the port that should be used as the compute instance network
	// port. Changing this creates a new trunk.
	PortID *string `json:"portId,omitempty" tf:"port_id,omitempty"`

	// The region in which to obtain the V2 networking client.
	// A networking client is needed to create a trunk. If omitted, the
	// region argument of the provider is used. Changing this creates a new
	// trunk.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The set of ports that will be made subports of the trunk.
	// The structure of each subport is described below.
	SubPort []SubPortObservation `json:"subPort,omitempty" tf:"sub_port,omitempty"`

	// A set of string tags for the port.
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The owner of the Trunk. Required if admin wants
	// to create a trunk on behalf of another tenant. Changing this creates a new trunk.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type TrunkV2Parameters struct {

	// Administrative up/down status for the trunk
	// (must be "true" or "false" if provided). Changing this updates the
	// admin_state_up of an existing trunk.
	// +kubebuilder:validation:Optional
	AdminStateUp *bool `json:"adminStateUp,omitempty" tf:"admin_state_up,omitempty"`

	// Human-readable description of the trunk. Changing this
	// updates the name of the existing trunk.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A unique name for the trunk. Changing this
	// updates the name of an existing trunk.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the port to be used as the parent port of the
	// trunk. This is the port that should be used as the compute instance network
	// port. Changing this creates a new trunk.
	// +kubebuilder:validation:Optional
	PortID *string `json:"portId,omitempty" tf:"port_id,omitempty"`

	// The region in which to obtain the V2 networking client.
	// A networking client is needed to create a trunk. If omitted, the
	// region argument of the provider is used. Changing this creates a new
	// trunk.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The set of ports that will be made subports of the trunk.
	// The structure of each subport is described below.
	// +kubebuilder:validation:Optional
	SubPort []SubPortParameters `json:"subPort,omitempty" tf:"sub_port,omitempty"`

	// A set of string tags for the port.
	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The owner of the Trunk. Required if admin wants
	// to create a trunk on behalf of another tenant. Changing this creates a new trunk.
	// +kubebuilder:validation:Optional
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

// TrunkV2Spec defines the desired state of TrunkV2
type TrunkV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TrunkV2Parameters `json:"forProvider"`
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
	InitProvider TrunkV2InitParameters `json:"initProvider,omitempty"`
}

// TrunkV2Status defines the observed state of TrunkV2.
type TrunkV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TrunkV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TrunkV2 is the Schema for the TrunkV2s API. Manages a networking V2 trunk resource within OpenStack.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,openstack}
type TrunkV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.portId) || (has(self.initProvider) && has(self.initProvider.portId))",message="spec.forProvider.portId is a required parameter"
	Spec   TrunkV2Spec   `json:"spec"`
	Status TrunkV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TrunkV2List contains a list of TrunkV2s
type TrunkV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TrunkV2 `json:"items"`
}

// Repository type metadata.
var (
	TrunkV2_Kind             = "TrunkV2"
	TrunkV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TrunkV2_Kind}.String()
	TrunkV2_KindAPIVersion   = TrunkV2_Kind + "." + CRDGroupVersion.String()
	TrunkV2_GroupVersionKind = CRDGroupVersion.WithKind(TrunkV2_Kind)
)

func init() {
	SchemeBuilder.Register(&TrunkV2{}, &TrunkV2List{})
}
