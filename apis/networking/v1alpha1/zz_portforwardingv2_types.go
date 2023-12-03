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

type PortforwardingV2InitParameters struct {

	// A text describing the port forwarding. Changing this
	// updates the description of an existing port forwarding.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The TCP/UDP/other protocol port number of the port forwarding. Changing this
	// updates the external_port of an existing port forwarding.
	ExternalPort *float64 `json:"externalPort,omitempty" tf:"external_port,omitempty"`

	// The ID of the Neutron floating IP address. Changing this creates a new port forwarding.
	FloatingipID *string `json:"floatingipId,omitempty" tf:"floatingip_id,omitempty"`

	// The fixed IPv4 address of the Neutron port associated with the port forwarding.
	// Changing this updates the internal_ip_address of an existing port forwarding.
	InternalIPAddress *string `json:"internalIpAddress,omitempty" tf:"internal_ip_address,omitempty"`

	// The TCP/UDP/other protocol port number of the Neutron port fixed IP address associated to the
	// port forwarding. Changing this updates the internal_port of an existing port forwarding.
	InternalPort *float64 `json:"internalPort,omitempty" tf:"internal_port,omitempty"`

	// The ID of the Neutron port associated with the port forwarding. Changing
	// this updates the internal_port_id of an existing port forwarding.
	InternalPortID *string `json:"internalPortId,omitempty" tf:"internal_port_id,omitempty"`

	// The IP protocol used in the port forwarding. Changing this updates the protocol
	// of an existing port forwarding.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// The region in which to obtain the V2 networking client.
	// A networking client is needed to create a port forwarding. If omitted, the
	// region argument of the provider is used. Changing this creates a new
	// port forwarding.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type PortforwardingV2Observation struct {

	// A text describing the port forwarding. Changing this
	// updates the description of an existing port forwarding.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The TCP/UDP/other protocol port number of the port forwarding. Changing this
	// updates the external_port of an existing port forwarding.
	ExternalPort *float64 `json:"externalPort,omitempty" tf:"external_port,omitempty"`

	// The ID of the Neutron floating IP address. Changing this creates a new port forwarding.
	FloatingipID *string `json:"floatingipId,omitempty" tf:"floatingip_id,omitempty"`

	// The ID of the floating IP port forwarding.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The fixed IPv4 address of the Neutron port associated with the port forwarding.
	// Changing this updates the internal_ip_address of an existing port forwarding.
	InternalIPAddress *string `json:"internalIpAddress,omitempty" tf:"internal_ip_address,omitempty"`

	// The TCP/UDP/other protocol port number of the Neutron port fixed IP address associated to the
	// port forwarding. Changing this updates the internal_port of an existing port forwarding.
	InternalPort *float64 `json:"internalPort,omitempty" tf:"internal_port,omitempty"`

	// The ID of the Neutron port associated with the port forwarding. Changing
	// this updates the internal_port_id of an existing port forwarding.
	InternalPortID *string `json:"internalPortId,omitempty" tf:"internal_port_id,omitempty"`

	// The IP protocol used in the port forwarding. Changing this updates the protocol
	// of an existing port forwarding.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// The region in which to obtain the V2 networking client.
	// A networking client is needed to create a port forwarding. If omitted, the
	// region argument of the provider is used. Changing this creates a new
	// port forwarding.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type PortforwardingV2Parameters struct {

	// A text describing the port forwarding. Changing this
	// updates the description of an existing port forwarding.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The TCP/UDP/other protocol port number of the port forwarding. Changing this
	// updates the external_port of an existing port forwarding.
	// +kubebuilder:validation:Optional
	ExternalPort *float64 `json:"externalPort,omitempty" tf:"external_port,omitempty"`

	// The ID of the Neutron floating IP address. Changing this creates a new port forwarding.
	// +kubebuilder:validation:Optional
	FloatingipID *string `json:"floatingipId,omitempty" tf:"floatingip_id,omitempty"`

	// The fixed IPv4 address of the Neutron port associated with the port forwarding.
	// Changing this updates the internal_ip_address of an existing port forwarding.
	// +kubebuilder:validation:Optional
	InternalIPAddress *string `json:"internalIpAddress,omitempty" tf:"internal_ip_address,omitempty"`

	// The TCP/UDP/other protocol port number of the Neutron port fixed IP address associated to the
	// port forwarding. Changing this updates the internal_port of an existing port forwarding.
	// +kubebuilder:validation:Optional
	InternalPort *float64 `json:"internalPort,omitempty" tf:"internal_port,omitempty"`

	// The ID of the Neutron port associated with the port forwarding. Changing
	// this updates the internal_port_id of an existing port forwarding.
	// +kubebuilder:validation:Optional
	InternalPortID *string `json:"internalPortId,omitempty" tf:"internal_port_id,omitempty"`

	// The IP protocol used in the port forwarding. Changing this updates the protocol
	// of an existing port forwarding.
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// The region in which to obtain the V2 networking client.
	// A networking client is needed to create a port forwarding. If omitted, the
	// region argument of the provider is used. Changing this creates a new
	// port forwarding.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

// PortforwardingV2Spec defines the desired state of PortforwardingV2
type PortforwardingV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PortforwardingV2Parameters `json:"forProvider"`
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
	InitProvider PortforwardingV2InitParameters `json:"initProvider,omitempty"`
}

// PortforwardingV2Status defines the observed state of PortforwardingV2.
type PortforwardingV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PortforwardingV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PortforwardingV2 is the Schema for the PortforwardingV2s API. Manages a V2 port forwarding resource within OpenStack.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,openstack}
type PortforwardingV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.externalPort) || (has(self.initProvider) && has(self.initProvider.externalPort))",message="spec.forProvider.externalPort is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.floatingipId) || (has(self.initProvider) && has(self.initProvider.floatingipId))",message="spec.forProvider.floatingipId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.internalIpAddress) || (has(self.initProvider) && has(self.initProvider.internalIpAddress))",message="spec.forProvider.internalIpAddress is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.internalPort) || (has(self.initProvider) && has(self.initProvider.internalPort))",message="spec.forProvider.internalPort is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.internalPortId) || (has(self.initProvider) && has(self.initProvider.internalPortId))",message="spec.forProvider.internalPortId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.protocol) || (has(self.initProvider) && has(self.initProvider.protocol))",message="spec.forProvider.protocol is a required parameter"
	Spec   PortforwardingV2Spec   `json:"spec"`
	Status PortforwardingV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PortforwardingV2List contains a list of PortforwardingV2s
type PortforwardingV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PortforwardingV2 `json:"items"`
}

// Repository type metadata.
var (
	PortforwardingV2_Kind             = "PortforwardingV2"
	PortforwardingV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PortforwardingV2_Kind}.String()
	PortforwardingV2_KindAPIVersion   = PortforwardingV2_Kind + "." + CRDGroupVersion.String()
	PortforwardingV2_GroupVersionKind = CRDGroupVersion.WithKind(PortforwardingV2_Kind)
)

func init() {
	SchemeBuilder.Register(&PortforwardingV2{}, &PortforwardingV2List{})
}
