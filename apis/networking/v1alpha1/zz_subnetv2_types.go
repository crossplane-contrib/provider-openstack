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

type AllocationPoolInitParameters struct {

	// The ending address.
	End *string `json:"end,omitempty" tf:"end,omitempty"`

	// The starting address.
	Start *string `json:"start,omitempty" tf:"start,omitempty"`
}

type AllocationPoolObservation struct {

	// The ending address.
	End *string `json:"end,omitempty" tf:"end,omitempty"`

	// The starting address.
	Start *string `json:"start,omitempty" tf:"start,omitempty"`
}

type AllocationPoolParameters struct {

	// The ending address.
	// +kubebuilder:validation:Optional
	End *string `json:"end" tf:"end,omitempty"`

	// The starting address.
	// +kubebuilder:validation:Optional
	Start *string `json:"start" tf:"start,omitempty"`
}

type AllocationPoolsInitParameters struct {

	// The ending address.
	End *string `json:"end,omitempty" tf:"end,omitempty"`

	// The starting address.
	Start *string `json:"start,omitempty" tf:"start,omitempty"`
}

type AllocationPoolsObservation struct {

	// The ending address.
	End *string `json:"end,omitempty" tf:"end,omitempty"`

	// The starting address.
	Start *string `json:"start,omitempty" tf:"start,omitempty"`
}

type AllocationPoolsParameters struct {

	// The ending address.
	// +kubebuilder:validation:Optional
	End *string `json:"end" tf:"end,omitempty"`

	// The starting address.
	// +kubebuilder:validation:Optional
	Start *string `json:"start" tf:"start,omitempty"`
}

type HostRoutesInitParameters struct {

	// The destination CIDR.
	DestinationCidr *string `json:"destinationCidr,omitempty" tf:"destination_cidr,omitempty"`

	// The next hop in the route.
	NextHop *string `json:"nextHop,omitempty" tf:"next_hop,omitempty"`
}

type HostRoutesObservation struct {

	// The destination CIDR.
	DestinationCidr *string `json:"destinationCidr,omitempty" tf:"destination_cidr,omitempty"`

	// The next hop in the route.
	NextHop *string `json:"nextHop,omitempty" tf:"next_hop,omitempty"`
}

type HostRoutesParameters struct {

	// The destination CIDR.
	// +kubebuilder:validation:Optional
	DestinationCidr *string `json:"destinationCidr" tf:"destination_cidr,omitempty"`

	// The next hop in the route.
	// +kubebuilder:validation:Optional
	NextHop *string `json:"nextHop" tf:"next_hop,omitempty"`
}

type SubnetV2InitParameters struct {

	// A block declaring the start and end range of
	// the IP addresses available for use with DHCP in this subnet. Multiple
	// allocation_pool blocks can be declared, providing the subnet with more
	// than one range of IP addresses to use with DHCP. However, each IP range
	// must be from the same CIDR that the subnet is part of.
	// The allocation_pool block is documented below.
	AllocationPool []AllocationPoolInitParameters `json:"allocationPool,omitempty" tf:"allocation_pool,omitempty"`

	// (Deprecated - use allocation_pool instead)
	// A block declaring the start and end range of the IP addresses available for
	// use with DHCP in this subnet.
	// The allocation_pools block is documented below.
	AllocationPools []AllocationPoolsInitParameters `json:"allocationPools,omitempty" tf:"allocation_pools,omitempty"`

	// CIDR representing IP range for this subnet, based on IP
	// version. You can omit this option if you are creating a subnet from a
	// subnet pool.
	Cidr *string `json:"cidr,omitempty" tf:"cidr,omitempty"`

	// An array of DNS name server names used by hosts
	// in this subnet. Changing this updates the DNS name servers for the existing
	// subnet.
	DNSNameservers []*string `json:"dnsNameservers,omitempty" tf:"dns_nameservers,omitempty"`

	// Human-readable description of the subnet. Changing this
	// updates the name of the existing subnet.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The administrative state of the network.
	// Acceptable values are "true" and "false". Changing this value enables or
	// disables the DHCP capabilities of the existing subnet. Defaults to true.
	EnableDHCP *bool `json:"enableDhcp,omitempty" tf:"enable_dhcp,omitempty"`

	// Default gateway used by devices in this subnet.
	// Leaving this blank and not setting no_gateway will cause a default
	// gateway of .1 to be used. Changing this updates the gateway IP of the
	// existing subnet.
	GatewayIP *string `json:"gatewayIp,omitempty" tf:"gateway_ip,omitempty"`

	// (Deprecated - use openstack_networking_subnet_route_v2
	// instead) An array of routes that should be used by devices
	// with IPs from this subnet (not including local subnet route). The host_route
	// object structure is documented below. Changing this updates the host routes
	// for the existing subnet.
	HostRoutes []HostRoutesInitParameters `json:"hostRoutes,omitempty" tf:"host_routes,omitempty"`

	// IP version, either 4 (default) or 6. Changing this creates a
	// new subnet.
	IPVersion *float64 `json:"ipVersion,omitempty" tf:"ip_version,omitempty"`

	// The IPv6 address mode. Valid values are
	// dhcpv6-stateful, dhcpv6-stateless, or slaac.
	IPv6AddressMode *string `json:"ipv6AddressMode,omitempty" tf:"ipv6_address_mode,omitempty"`

	// The IPv6 Router Advertisement mode. Valid values
	// are dhcpv6-stateful, dhcpv6-stateless, or slaac.
	IPv6RaMode *string `json:"ipv6RaMode,omitempty" tf:"ipv6_ra_mode,omitempty"`

	// The name of the subnet. Changing this updates the name of
	// the existing subnet.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The UUID of the parent network. Changing this
	// creates a new subnet.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-openstack/apis/networking/v1alpha1.SubnetV2
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// Reference to a SubnetV2 in networking to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDRef *v1.Reference `json:"networkIdRef,omitempty" tf:"-"`

	// Selector for a SubnetV2 in networking to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDSelector *v1.Selector `json:"networkIdSelector,omitempty" tf:"-"`

	// Do not set a gateway IP on this subnet. Changing
	// this removes or adds a default gateway IP of the existing subnet.
	NoGateway *bool `json:"noGateway,omitempty" tf:"no_gateway,omitempty"`

	// The prefix length to use when creating a subnet
	// from a subnet pool. The default subnet pool prefix length that was defined
	// when creating the subnet pool will be used if not provided. Changing this
	// creates a new subnet.
	PrefixLength *float64 `json:"prefixLength,omitempty" tf:"prefix_length,omitempty"`

	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a Neutron subnet. If omitted, the
	// region argument of the provider is used. Changing this creates a new
	// subnet.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// An array of service types used by the subnet.
	// Changing this updates the service types for the existing subnet.
	ServiceTypes []*string `json:"serviceTypes,omitempty" tf:"service_types,omitempty"`

	// The ID of the subnetpool associated with the subnet.
	SubnetpoolID *string `json:"subnetpoolId,omitempty" tf:"subnetpool_id,omitempty"`

	// A set of string tags for the subnet.
	// +listType=set
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The owner of the subnet. Required if admin wants to
	// create a subnet for another tenant. Changing this creates a new subnet.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// Map of additional options.
	// +mapType=granular
	ValueSpecs map[string]*string `json:"valueSpecs,omitempty" tf:"value_specs,omitempty"`
}

type SubnetV2Observation struct {

	// The collection of ags assigned on the subnet, which have been
	// explicitly and implicitly added.
	// +listType=set
	AllTags []*string `json:"allTags,omitempty" tf:"all_tags,omitempty"`

	// A block declaring the start and end range of
	// the IP addresses available for use with DHCP in this subnet. Multiple
	// allocation_pool blocks can be declared, providing the subnet with more
	// than one range of IP addresses to use with DHCP. However, each IP range
	// must be from the same CIDR that the subnet is part of.
	// The allocation_pool block is documented below.
	AllocationPool []AllocationPoolObservation `json:"allocationPool,omitempty" tf:"allocation_pool,omitempty"`

	// (Deprecated - use allocation_pool instead)
	// A block declaring the start and end range of the IP addresses available for
	// use with DHCP in this subnet.
	// The allocation_pools block is documented below.
	AllocationPools []AllocationPoolsObservation `json:"allocationPools,omitempty" tf:"allocation_pools,omitempty"`

	// CIDR representing IP range for this subnet, based on IP
	// version. You can omit this option if you are creating a subnet from a
	// subnet pool.
	Cidr *string `json:"cidr,omitempty" tf:"cidr,omitempty"`

	// An array of DNS name server names used by hosts
	// in this subnet. Changing this updates the DNS name servers for the existing
	// subnet.
	DNSNameservers []*string `json:"dnsNameservers,omitempty" tf:"dns_nameservers,omitempty"`

	// Human-readable description of the subnet. Changing this
	// updates the name of the existing subnet.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The administrative state of the network.
	// Acceptable values are "true" and "false". Changing this value enables or
	// disables the DHCP capabilities of the existing subnet. Defaults to true.
	EnableDHCP *bool `json:"enableDhcp,omitempty" tf:"enable_dhcp,omitempty"`

	// Default gateway used by devices in this subnet.
	// Leaving this blank and not setting no_gateway will cause a default
	// gateway of .1 to be used. Changing this updates the gateway IP of the
	// existing subnet.
	GatewayIP *string `json:"gatewayIp,omitempty" tf:"gateway_ip,omitempty"`

	// (Deprecated - use openstack_networking_subnet_route_v2
	// instead) An array of routes that should be used by devices
	// with IPs from this subnet (not including local subnet route). The host_route
	// object structure is documented below. Changing this updates the host routes
	// for the existing subnet.
	HostRoutes []HostRoutesObservation `json:"hostRoutes,omitempty" tf:"host_routes,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// IP version, either 4 (default) or 6. Changing this creates a
	// new subnet.
	IPVersion *float64 `json:"ipVersion,omitempty" tf:"ip_version,omitempty"`

	// The IPv6 address mode. Valid values are
	// dhcpv6-stateful, dhcpv6-stateless, or slaac.
	IPv6AddressMode *string `json:"ipv6AddressMode,omitempty" tf:"ipv6_address_mode,omitempty"`

	// The IPv6 Router Advertisement mode. Valid values
	// are dhcpv6-stateful, dhcpv6-stateless, or slaac.
	IPv6RaMode *string `json:"ipv6RaMode,omitempty" tf:"ipv6_ra_mode,omitempty"`

	// The name of the subnet. Changing this updates the name of
	// the existing subnet.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The UUID of the parent network. Changing this
	// creates a new subnet.
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// Do not set a gateway IP on this subnet. Changing
	// this removes or adds a default gateway IP of the existing subnet.
	NoGateway *bool `json:"noGateway,omitempty" tf:"no_gateway,omitempty"`

	// The prefix length to use when creating a subnet
	// from a subnet pool. The default subnet pool prefix length that was defined
	// when creating the subnet pool will be used if not provided. Changing this
	// creates a new subnet.
	PrefixLength *float64 `json:"prefixLength,omitempty" tf:"prefix_length,omitempty"`

	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a Neutron subnet. If omitted, the
	// region argument of the provider is used. Changing this creates a new
	// subnet.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// An array of service types used by the subnet.
	// Changing this updates the service types for the existing subnet.
	ServiceTypes []*string `json:"serviceTypes,omitempty" tf:"service_types,omitempty"`

	// The ID of the subnetpool associated with the subnet.
	SubnetpoolID *string `json:"subnetpoolId,omitempty" tf:"subnetpool_id,omitempty"`

	// A set of string tags for the subnet.
	// +listType=set
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The owner of the subnet. Required if admin wants to
	// create a subnet for another tenant. Changing this creates a new subnet.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// Map of additional options.
	// +mapType=granular
	ValueSpecs map[string]*string `json:"valueSpecs,omitempty" tf:"value_specs,omitempty"`
}

type SubnetV2Parameters struct {

	// A block declaring the start and end range of
	// the IP addresses available for use with DHCP in this subnet. Multiple
	// allocation_pool blocks can be declared, providing the subnet with more
	// than one range of IP addresses to use with DHCP. However, each IP range
	// must be from the same CIDR that the subnet is part of.
	// The allocation_pool block is documented below.
	// +kubebuilder:validation:Optional
	AllocationPool []AllocationPoolParameters `json:"allocationPool,omitempty" tf:"allocation_pool,omitempty"`

	// (Deprecated - use allocation_pool instead)
	// A block declaring the start and end range of the IP addresses available for
	// use with DHCP in this subnet.
	// The allocation_pools block is documented below.
	// +kubebuilder:validation:Optional
	AllocationPools []AllocationPoolsParameters `json:"allocationPools,omitempty" tf:"allocation_pools,omitempty"`

	// CIDR representing IP range for this subnet, based on IP
	// version. You can omit this option if you are creating a subnet from a
	// subnet pool.
	// +kubebuilder:validation:Optional
	Cidr *string `json:"cidr,omitempty" tf:"cidr,omitempty"`

	// An array of DNS name server names used by hosts
	// in this subnet. Changing this updates the DNS name servers for the existing
	// subnet.
	// +kubebuilder:validation:Optional
	DNSNameservers []*string `json:"dnsNameservers,omitempty" tf:"dns_nameservers,omitempty"`

	// Human-readable description of the subnet. Changing this
	// updates the name of the existing subnet.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The administrative state of the network.
	// Acceptable values are "true" and "false". Changing this value enables or
	// disables the DHCP capabilities of the existing subnet. Defaults to true.
	// +kubebuilder:validation:Optional
	EnableDHCP *bool `json:"enableDhcp,omitempty" tf:"enable_dhcp,omitempty"`

	// Default gateway used by devices in this subnet.
	// Leaving this blank and not setting no_gateway will cause a default
	// gateway of .1 to be used. Changing this updates the gateway IP of the
	// existing subnet.
	// +kubebuilder:validation:Optional
	GatewayIP *string `json:"gatewayIp,omitempty" tf:"gateway_ip,omitempty"`

	// (Deprecated - use openstack_networking_subnet_route_v2
	// instead) An array of routes that should be used by devices
	// with IPs from this subnet (not including local subnet route). The host_route
	// object structure is documented below. Changing this updates the host routes
	// for the existing subnet.
	// +kubebuilder:validation:Optional
	HostRoutes []HostRoutesParameters `json:"hostRoutes,omitempty" tf:"host_routes,omitempty"`

	// IP version, either 4 (default) or 6. Changing this creates a
	// new subnet.
	// +kubebuilder:validation:Optional
	IPVersion *float64 `json:"ipVersion,omitempty" tf:"ip_version,omitempty"`

	// The IPv6 address mode. Valid values are
	// dhcpv6-stateful, dhcpv6-stateless, or slaac.
	// +kubebuilder:validation:Optional
	IPv6AddressMode *string `json:"ipv6AddressMode,omitempty" tf:"ipv6_address_mode,omitempty"`

	// The IPv6 Router Advertisement mode. Valid values
	// are dhcpv6-stateful, dhcpv6-stateless, or slaac.
	// +kubebuilder:validation:Optional
	IPv6RaMode *string `json:"ipv6RaMode,omitempty" tf:"ipv6_ra_mode,omitempty"`

	// The name of the subnet. Changing this updates the name of
	// the existing subnet.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The UUID of the parent network. Changing this
	// creates a new subnet.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-openstack/apis/networking/v1alpha1.SubnetV2
	// +kubebuilder:validation:Optional
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// Reference to a SubnetV2 in networking to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDRef *v1.Reference `json:"networkIdRef,omitempty" tf:"-"`

	// Selector for a SubnetV2 in networking to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDSelector *v1.Selector `json:"networkIdSelector,omitempty" tf:"-"`

	// Do not set a gateway IP on this subnet. Changing
	// this removes or adds a default gateway IP of the existing subnet.
	// +kubebuilder:validation:Optional
	NoGateway *bool `json:"noGateway,omitempty" tf:"no_gateway,omitempty"`

	// The prefix length to use when creating a subnet
	// from a subnet pool. The default subnet pool prefix length that was defined
	// when creating the subnet pool will be used if not provided. Changing this
	// creates a new subnet.
	// +kubebuilder:validation:Optional
	PrefixLength *float64 `json:"prefixLength,omitempty" tf:"prefix_length,omitempty"`

	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a Neutron subnet. If omitted, the
	// region argument of the provider is used. Changing this creates a new
	// subnet.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// An array of service types used by the subnet.
	// Changing this updates the service types for the existing subnet.
	// +kubebuilder:validation:Optional
	ServiceTypes []*string `json:"serviceTypes,omitempty" tf:"service_types,omitempty"`

	// The ID of the subnetpool associated with the subnet.
	// +kubebuilder:validation:Optional
	SubnetpoolID *string `json:"subnetpoolId,omitempty" tf:"subnetpool_id,omitempty"`

	// A set of string tags for the subnet.
	// +kubebuilder:validation:Optional
	// +listType=set
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The owner of the subnet. Required if admin wants to
	// create a subnet for another tenant. Changing this creates a new subnet.
	// +kubebuilder:validation:Optional
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// Map of additional options.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	ValueSpecs map[string]*string `json:"valueSpecs,omitempty" tf:"value_specs,omitempty"`
}

// SubnetV2Spec defines the desired state of SubnetV2
type SubnetV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SubnetV2Parameters `json:"forProvider"`
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
	InitProvider SubnetV2InitParameters `json:"initProvider,omitempty"`
}

// SubnetV2Status defines the observed state of SubnetV2.
type SubnetV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SubnetV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SubnetV2 is the Schema for the SubnetV2s API. Manages a V2 Neutron subnet resource within OpenStack.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,openstack}
type SubnetV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SubnetV2Spec   `json:"spec"`
	Status            SubnetV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SubnetV2List contains a list of SubnetV2s
type SubnetV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SubnetV2 `json:"items"`
}

// Repository type metadata.
var (
	SubnetV2_Kind             = "SubnetV2"
	SubnetV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SubnetV2_Kind}.String()
	SubnetV2_KindAPIVersion   = SubnetV2_Kind + "." + CRDGroupVersion.String()
	SubnetV2_GroupVersionKind = CRDGroupVersion.WithKind(SubnetV2_Kind)
)

func init() {
	SchemeBuilder.Register(&SubnetV2{}, &SubnetV2List{})
}
