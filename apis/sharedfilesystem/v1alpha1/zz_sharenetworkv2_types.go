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

type SharenetworkV2InitParameters struct {

	// The human-readable description for the share network.
	// Changing this updates the description of the existing share network.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name for the share network. Changing this updates the name
	// of the existing share network.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The UUID of a neutron network when setting up or updating
	// a share network. Changing this updates the existing share network if it's not used by
	// shares.
	NeutronNetID *string `json:"neutronNetId,omitempty" tf:"neutron_net_id,omitempty"`

	// The UUID of the neutron subnet when setting up or
	// updating a share network. Changing this updates the existing share network if it's
	// not used by shares.
	NeutronSubnetID *string `json:"neutronSubnetId,omitempty" tf:"neutron_subnet_id,omitempty"`

	// The region in which to obtain the V2 Shared File System client.
	// A Shared File System client is needed to create a share network. If omitted, the
	// region argument of the provider is used. Changing this creates a new
	// share network.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The list of security service IDs to associate with
	// the share network. The security service must be specified by ID and not name.
	SecurityServiceIds []*string `json:"securityServiceIds,omitempty" tf:"security_service_ids,omitempty"`
}

type SharenetworkV2Observation struct {

	// The share network CIDR.
	Cidr *string `json:"cidr,omitempty" tf:"cidr,omitempty"`

	// The human-readable description for the share network.
	// Changing this updates the description of the existing share network.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The unique ID for the Share Network.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The IP version of the share network. Can either be 4 or 6.
	IPVersion *float64 `json:"ipVersion,omitempty" tf:"ip_version,omitempty"`

	// The name for the share network. Changing this updates the name
	// of the existing share network.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The share network type. Can either be VLAN, VXLAN, GRE, or flat.
	NetworkType *string `json:"networkType,omitempty" tf:"network_type,omitempty"`

	// The UUID of a neutron network when setting up or updating
	// a share network. Changing this updates the existing share network if it's not used by
	// shares.
	NeutronNetID *string `json:"neutronNetId,omitempty" tf:"neutron_net_id,omitempty"`

	// The UUID of the neutron subnet when setting up or
	// updating a share network. Changing this updates the existing share network if it's
	// not used by shares.
	NeutronSubnetID *string `json:"neutronSubnetId,omitempty" tf:"neutron_subnet_id,omitempty"`

	// The owner of the Share Network.
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// The region in which to obtain the V2 Shared File System client.
	// A Shared File System client is needed to create a share network. If omitted, the
	// region argument of the provider is used. Changing this creates a new
	// share network.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The list of security service IDs to associate with
	// the share network. The security service must be specified by ID and not name.
	SecurityServiceIds []*string `json:"securityServiceIds,omitempty" tf:"security_service_ids,omitempty"`

	// The share network segmentation ID.
	SegmentationID *float64 `json:"segmentationId,omitempty" tf:"segmentation_id,omitempty"`
}

type SharenetworkV2Parameters struct {

	// The human-readable description for the share network.
	// Changing this updates the description of the existing share network.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name for the share network. Changing this updates the name
	// of the existing share network.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The UUID of a neutron network when setting up or updating
	// a share network. Changing this updates the existing share network if it's not used by
	// shares.
	// +kubebuilder:validation:Optional
	NeutronNetID *string `json:"neutronNetId,omitempty" tf:"neutron_net_id,omitempty"`

	// The UUID of the neutron subnet when setting up or
	// updating a share network. Changing this updates the existing share network if it's
	// not used by shares.
	// +kubebuilder:validation:Optional
	NeutronSubnetID *string `json:"neutronSubnetId,omitempty" tf:"neutron_subnet_id,omitempty"`

	// The region in which to obtain the V2 Shared File System client.
	// A Shared File System client is needed to create a share network. If omitted, the
	// region argument of the provider is used. Changing this creates a new
	// share network.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The list of security service IDs to associate with
	// the share network. The security service must be specified by ID and not name.
	// +kubebuilder:validation:Optional
	SecurityServiceIds []*string `json:"securityServiceIds,omitempty" tf:"security_service_ids,omitempty"`
}

// SharenetworkV2Spec defines the desired state of SharenetworkV2
type SharenetworkV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SharenetworkV2Parameters `json:"forProvider"`
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
	InitProvider SharenetworkV2InitParameters `json:"initProvider,omitempty"`
}

// SharenetworkV2Status defines the observed state of SharenetworkV2.
type SharenetworkV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SharenetworkV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SharenetworkV2 is the Schema for the SharenetworkV2s API. Configure a Shared File System share network.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,openstack}
type SharenetworkV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.neutronNetId) || (has(self.initProvider) && has(self.initProvider.neutronNetId))",message="spec.forProvider.neutronNetId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.neutronSubnetId) || (has(self.initProvider) && has(self.initProvider.neutronSubnetId))",message="spec.forProvider.neutronSubnetId is a required parameter"
	Spec   SharenetworkV2Spec   `json:"spec"`
	Status SharenetworkV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SharenetworkV2List contains a list of SharenetworkV2s
type SharenetworkV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SharenetworkV2 `json:"items"`
}

// Repository type metadata.
var (
	SharenetworkV2_Kind             = "SharenetworkV2"
	SharenetworkV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SharenetworkV2_Kind}.String()
	SharenetworkV2_KindAPIVersion   = SharenetworkV2_Kind + "." + CRDGroupVersion.String()
	SharenetworkV2_GroupVersionKind = CRDGroupVersion.WithKind(SharenetworkV2_Kind)
)

func init() {
	SchemeBuilder.Register(&SharenetworkV2{}, &SharenetworkV2List{})
}
