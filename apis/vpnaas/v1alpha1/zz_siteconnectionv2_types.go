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

type DpdInitParameters struct {

	// The dead peer detection (DPD) action.
	// A valid value is clear, hold, restart, disabled, or restart-by-peer.
	// Default value is hold.
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// The dead peer detection (DPD) interval, in seconds.
	// A valid value is a positive integer.
	// Default is 30.
	Interval *float64 `json:"interval,omitempty" tf:"interval,omitempty"`

	// The dead peer detection (DPD) timeout in seconds.
	// A valid value is a positive integer that is greater than the DPD interval value.
	// Default is 120.
	Timeout *float64 `json:"timeout,omitempty" tf:"timeout,omitempty"`
}

type DpdObservation struct {

	// The dead peer detection (DPD) action.
	// A valid value is clear, hold, restart, disabled, or restart-by-peer.
	// Default value is hold.
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// The dead peer detection (DPD) interval, in seconds.
	// A valid value is a positive integer.
	// Default is 30.
	Interval *float64 `json:"interval,omitempty" tf:"interval,omitempty"`

	// The dead peer detection (DPD) timeout in seconds.
	// A valid value is a positive integer that is greater than the DPD interval value.
	// Default is 120.
	Timeout *float64 `json:"timeout,omitempty" tf:"timeout,omitempty"`
}

type DpdParameters struct {

	// The dead peer detection (DPD) action.
	// A valid value is clear, hold, restart, disabled, or restart-by-peer.
	// Default value is hold.
	// +kubebuilder:validation:Optional
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// The dead peer detection (DPD) interval, in seconds.
	// A valid value is a positive integer.
	// Default is 30.
	// +kubebuilder:validation:Optional
	Interval *float64 `json:"interval,omitempty" tf:"interval,omitempty"`

	// The dead peer detection (DPD) timeout in seconds.
	// A valid value is a positive integer that is greater than the DPD interval value.
	// Default is 120.
	// +kubebuilder:validation:Optional
	Timeout *float64 `json:"timeout,omitempty" tf:"timeout,omitempty"`
}

type SiteConnectionV2InitParameters struct {

	// The administrative state of the resource. Can either be up(true) or down(false).
	// Changing this updates the administrative state of the existing connection.
	AdminStateUp *bool `json:"adminStateUp,omitempty" tf:"admin_state_up,omitempty"`

	// The human-readable description for the connection.
	// Changing this updates the description of the existing connection.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A dictionary with dead peer detection (DPD) protocol controls.
	Dpd []DpdInitParameters `json:"dpd,omitempty" tf:"dpd,omitempty"`

	// The ID of the IKE policy. Changing this creates a new connection.
	IkepolicyID *string `json:"ikepolicyId,omitempty" tf:"ikepolicy_id,omitempty"`

	// A valid value is response-only or bi-directional. Default is bi-directional.
	Initiator *string `json:"initiator,omitempty" tf:"initiator,omitempty"`

	// The ID of the IPsec policy. Changing this creates a new connection.
	IpsecpolicyID *string `json:"ipsecpolicyId,omitempty" tf:"ipsecpolicy_id,omitempty"`

	// The ID for the endpoint group that contains private subnets for the local side of the connection.
	// You must specify this parameter with the peer_ep_group_id parameter unless
	// in backward- compatible mode where peer_cidrs is provided with a subnet_id for the VPN service.
	// Changing this updates the existing connection.
	LocalEpGroupID *string `json:"localEpGroupId,omitempty" tf:"local_ep_group_id,omitempty"`

	// An ID to be used instead of the external IP address for a virtual router used in traffic between instances on different networks in east-west traffic.
	// Most often, local ID would be domain name, email address, etc.
	// If this is not configured then the external IP address will be used as the ID.
	LocalID *string `json:"localId,omitempty" tf:"local_id,omitempty"`

	// The maximum transmission unit (MTU) value to address fragmentation.
	// Minimum value is 68 for IPv4, and 1280 for IPv6.
	Mtu *float64 `json:"mtu,omitempty" tf:"mtu,omitempty"`

	// The name of the connection. Changing this updates the name of
	// the existing connection.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The peer gateway public IPv4 or IPv6 address or FQDN.
	PeerAddress *string `json:"peerAddress,omitempty" tf:"peer_address,omitempty"`

	// Unique list of valid peer private CIDRs in the form < net_address > / < prefix > .
	PeerCidrs []*string `json:"peerCidrs,omitempty" tf:"peer_cidrs,omitempty"`

	// The ID for the endpoint group that contains private CIDRs in the form < net_address > / < prefix > for the peer side of the connection.
	// You must specify this parameter with the local_ep_group_id parameter unless in backward-compatible mode
	// where peer_cidrs is provided with a subnet_id for the VPN service.
	PeerEpGroupID *string `json:"peerEpGroupId,omitempty" tf:"peer_ep_group_id,omitempty"`

	// The peer router identity for authentication. A valid value is an IPv4 address, IPv6 address, e-mail address, key ID, or FQDN.
	// Typically, this value matches the peer_address value.
	// Changing this updates the existing policy.
	PeerID *string `json:"peerId,omitempty" tf:"peer_id,omitempty"`

	// The pre-shared key. A valid value is any string.
	Psk *string `json:"psk,omitempty" tf:"psk,omitempty"`

	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create an IPSec site connection. If omitted, the
	// region argument of the provider is used. Changing this creates a new
	// site connection.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The owner of the connection. Required if admin wants to
	// create a connection for another project. Changing this creates a new connection.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// Map of additional options.
	// +mapType=granular
	ValueSpecs map[string]*string `json:"valueSpecs,omitempty" tf:"value_specs,omitempty"`

	// The ID of the VPN service. Changing this creates a new connection.
	VpnserviceID *string `json:"vpnserviceId,omitempty" tf:"vpnservice_id,omitempty"`
}

type SiteConnectionV2Observation struct {

	// The administrative state of the resource. Can either be up(true) or down(false).
	// Changing this updates the administrative state of the existing connection.
	AdminStateUp *bool `json:"adminStateUp,omitempty" tf:"admin_state_up,omitempty"`

	// The human-readable description for the connection.
	// Changing this updates the description of the existing connection.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A dictionary with dead peer detection (DPD) protocol controls.
	Dpd []DpdObservation `json:"dpd,omitempty" tf:"dpd,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the IKE policy. Changing this creates a new connection.
	IkepolicyID *string `json:"ikepolicyId,omitempty" tf:"ikepolicy_id,omitempty"`

	// A valid value is response-only or bi-directional. Default is bi-directional.
	Initiator *string `json:"initiator,omitempty" tf:"initiator,omitempty"`

	// The ID of the IPsec policy. Changing this creates a new connection.
	IpsecpolicyID *string `json:"ipsecpolicyId,omitempty" tf:"ipsecpolicy_id,omitempty"`

	// The ID for the endpoint group that contains private subnets for the local side of the connection.
	// You must specify this parameter with the peer_ep_group_id parameter unless
	// in backward- compatible mode where peer_cidrs is provided with a subnet_id for the VPN service.
	// Changing this updates the existing connection.
	LocalEpGroupID *string `json:"localEpGroupId,omitempty" tf:"local_ep_group_id,omitempty"`

	// An ID to be used instead of the external IP address for a virtual router used in traffic between instances on different networks in east-west traffic.
	// Most often, local ID would be domain name, email address, etc.
	// If this is not configured then the external IP address will be used as the ID.
	LocalID *string `json:"localId,omitempty" tf:"local_id,omitempty"`

	// The maximum transmission unit (MTU) value to address fragmentation.
	// Minimum value is 68 for IPv4, and 1280 for IPv6.
	Mtu *float64 `json:"mtu,omitempty" tf:"mtu,omitempty"`

	// The name of the connection. Changing this updates the name of
	// the existing connection.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The peer gateway public IPv4 or IPv6 address or FQDN.
	PeerAddress *string `json:"peerAddress,omitempty" tf:"peer_address,omitempty"`

	// Unique list of valid peer private CIDRs in the form < net_address > / < prefix > .
	PeerCidrs []*string `json:"peerCidrs,omitempty" tf:"peer_cidrs,omitempty"`

	// The ID for the endpoint group that contains private CIDRs in the form < net_address > / < prefix > for the peer side of the connection.
	// You must specify this parameter with the local_ep_group_id parameter unless in backward-compatible mode
	// where peer_cidrs is provided with a subnet_id for the VPN service.
	PeerEpGroupID *string `json:"peerEpGroupId,omitempty" tf:"peer_ep_group_id,omitempty"`

	// The peer router identity for authentication. A valid value is an IPv4 address, IPv6 address, e-mail address, key ID, or FQDN.
	// Typically, this value matches the peer_address value.
	// Changing this updates the existing policy.
	PeerID *string `json:"peerId,omitempty" tf:"peer_id,omitempty"`

	// The pre-shared key. A valid value is any string.
	Psk *string `json:"psk,omitempty" tf:"psk,omitempty"`

	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create an IPSec site connection. If omitted, the
	// region argument of the provider is used. Changing this creates a new
	// site connection.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The owner of the connection. Required if admin wants to
	// create a connection for another project. Changing this creates a new connection.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// Map of additional options.
	// +mapType=granular
	ValueSpecs map[string]*string `json:"valueSpecs,omitempty" tf:"value_specs,omitempty"`

	// The ID of the VPN service. Changing this creates a new connection.
	VpnserviceID *string `json:"vpnserviceId,omitempty" tf:"vpnservice_id,omitempty"`
}

type SiteConnectionV2Parameters struct {

	// The administrative state of the resource. Can either be up(true) or down(false).
	// Changing this updates the administrative state of the existing connection.
	// +kubebuilder:validation:Optional
	AdminStateUp *bool `json:"adminStateUp,omitempty" tf:"admin_state_up,omitempty"`

	// The human-readable description for the connection.
	// Changing this updates the description of the existing connection.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A dictionary with dead peer detection (DPD) protocol controls.
	// +kubebuilder:validation:Optional
	Dpd []DpdParameters `json:"dpd,omitempty" tf:"dpd,omitempty"`

	// The ID of the IKE policy. Changing this creates a new connection.
	// +kubebuilder:validation:Optional
	IkepolicyID *string `json:"ikepolicyId,omitempty" tf:"ikepolicy_id,omitempty"`

	// A valid value is response-only or bi-directional. Default is bi-directional.
	// +kubebuilder:validation:Optional
	Initiator *string `json:"initiator,omitempty" tf:"initiator,omitempty"`

	// The ID of the IPsec policy. Changing this creates a new connection.
	// +kubebuilder:validation:Optional
	IpsecpolicyID *string `json:"ipsecpolicyId,omitempty" tf:"ipsecpolicy_id,omitempty"`

	// The ID for the endpoint group that contains private subnets for the local side of the connection.
	// You must specify this parameter with the peer_ep_group_id parameter unless
	// in backward- compatible mode where peer_cidrs is provided with a subnet_id for the VPN service.
	// Changing this updates the existing connection.
	// +kubebuilder:validation:Optional
	LocalEpGroupID *string `json:"localEpGroupId,omitempty" tf:"local_ep_group_id,omitempty"`

	// An ID to be used instead of the external IP address for a virtual router used in traffic between instances on different networks in east-west traffic.
	// Most often, local ID would be domain name, email address, etc.
	// If this is not configured then the external IP address will be used as the ID.
	// +kubebuilder:validation:Optional
	LocalID *string `json:"localId,omitempty" tf:"local_id,omitempty"`

	// The maximum transmission unit (MTU) value to address fragmentation.
	// Minimum value is 68 for IPv4, and 1280 for IPv6.
	// +kubebuilder:validation:Optional
	Mtu *float64 `json:"mtu,omitempty" tf:"mtu,omitempty"`

	// The name of the connection. Changing this updates the name of
	// the existing connection.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The peer gateway public IPv4 or IPv6 address or FQDN.
	// +kubebuilder:validation:Optional
	PeerAddress *string `json:"peerAddress,omitempty" tf:"peer_address,omitempty"`

	// Unique list of valid peer private CIDRs in the form < net_address > / < prefix > .
	// +kubebuilder:validation:Optional
	PeerCidrs []*string `json:"peerCidrs,omitempty" tf:"peer_cidrs,omitempty"`

	// The ID for the endpoint group that contains private CIDRs in the form < net_address > / < prefix > for the peer side of the connection.
	// You must specify this parameter with the local_ep_group_id parameter unless in backward-compatible mode
	// where peer_cidrs is provided with a subnet_id for the VPN service.
	// +kubebuilder:validation:Optional
	PeerEpGroupID *string `json:"peerEpGroupId,omitempty" tf:"peer_ep_group_id,omitempty"`

	// The peer router identity for authentication. A valid value is an IPv4 address, IPv6 address, e-mail address, key ID, or FQDN.
	// Typically, this value matches the peer_address value.
	// Changing this updates the existing policy.
	// +kubebuilder:validation:Optional
	PeerID *string `json:"peerId,omitempty" tf:"peer_id,omitempty"`

	// The pre-shared key. A valid value is any string.
	// +kubebuilder:validation:Optional
	Psk *string `json:"psk,omitempty" tf:"psk,omitempty"`

	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create an IPSec site connection. If omitted, the
	// region argument of the provider is used. Changing this creates a new
	// site connection.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The owner of the connection. Required if admin wants to
	// create a connection for another project. Changing this creates a new connection.
	// +kubebuilder:validation:Optional
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// Map of additional options.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	ValueSpecs map[string]*string `json:"valueSpecs,omitempty" tf:"value_specs,omitempty"`

	// The ID of the VPN service. Changing this creates a new connection.
	// +kubebuilder:validation:Optional
	VpnserviceID *string `json:"vpnserviceId,omitempty" tf:"vpnservice_id,omitempty"`
}

// SiteConnectionV2Spec defines the desired state of SiteConnectionV2
type SiteConnectionV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SiteConnectionV2Parameters `json:"forProvider"`
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
	InitProvider SiteConnectionV2InitParameters `json:"initProvider,omitempty"`
}

// SiteConnectionV2Status defines the observed state of SiteConnectionV2.
type SiteConnectionV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SiteConnectionV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SiteConnectionV2 is the Schema for the SiteConnectionV2s API. Manages a V2 Neutron IPSec site connection resource within OpenStack.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,openstack}
type SiteConnectionV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.ikepolicyId) || (has(self.initProvider) && has(self.initProvider.ikepolicyId))",message="spec.forProvider.ikepolicyId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.ipsecpolicyId) || (has(self.initProvider) && has(self.initProvider.ipsecpolicyId))",message="spec.forProvider.ipsecpolicyId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.peerAddress) || (has(self.initProvider) && has(self.initProvider.peerAddress))",message="spec.forProvider.peerAddress is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.peerId) || (has(self.initProvider) && has(self.initProvider.peerId))",message="spec.forProvider.peerId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.psk) || (has(self.initProvider) && has(self.initProvider.psk))",message="spec.forProvider.psk is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.vpnserviceId) || (has(self.initProvider) && has(self.initProvider.vpnserviceId))",message="spec.forProvider.vpnserviceId is a required parameter"
	Spec   SiteConnectionV2Spec   `json:"spec"`
	Status SiteConnectionV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SiteConnectionV2List contains a list of SiteConnectionV2s
type SiteConnectionV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SiteConnectionV2 `json:"items"`
}

// Repository type metadata.
var (
	SiteConnectionV2_Kind             = "SiteConnectionV2"
	SiteConnectionV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SiteConnectionV2_Kind}.String()
	SiteConnectionV2_KindAPIVersion   = SiteConnectionV2_Kind + "." + CRDGroupVersion.String()
	SiteConnectionV2_GroupVersionKind = CRDGroupVersion.WithKind(SiteConnectionV2_Kind)
)

func init() {
	SchemeBuilder.Register(&SiteConnectionV2{}, &SiteConnectionV2List{})
}
