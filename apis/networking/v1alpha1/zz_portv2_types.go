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

type AllowedAddressPairsInitParameters struct {

	// The additional IP address.
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// The additional MAC address.
	MacAddress *string `json:"macAddress,omitempty" tf:"mac_address,omitempty"`
}

type AllowedAddressPairsObservation struct {

	// The additional IP address.
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// The additional MAC address.
	MacAddress *string `json:"macAddress,omitempty" tf:"mac_address,omitempty"`
}

type AllowedAddressPairsParameters struct {

	// The additional IP address.
	// +kubebuilder:validation:Optional
	IPAddress *string `json:"ipAddress" tf:"ip_address,omitempty"`

	// The additional MAC address.
	// +kubebuilder:validation:Optional
	MacAddress *string `json:"macAddress,omitempty" tf:"mac_address,omitempty"`
}

type BindingInitParameters struct {

	// The ID of the host to allocate port on.
	HostID *string `json:"hostId,omitempty" tf:"host_id,omitempty"`

	// Custom data to be passed as binding:profile. Data
	// must be passed as JSON.
	Profile *string `json:"profile,omitempty" tf:"profile,omitempty"`

	// VNIC type for the port. Can either be direct,
	// direct-physical, macvtap, normal, baremetal or virtio-forwarder.
	// Default value is normal. It can be updated on unbound ports only.
	VnicType *string `json:"vnicType,omitempty" tf:"vnic_type,omitempty"`
}

type BindingObservation struct {

	// The ID of the host to allocate port on.
	HostID *string `json:"hostId,omitempty" tf:"host_id,omitempty"`

	// Custom data to be passed as binding:profile. Data
	// must be passed as JSON.
	Profile *string `json:"profile,omitempty" tf:"profile,omitempty"`

	// (Computed) A map of JSON strings containing additional
	// details for this specific binding.
	// +mapType=granular
	VifDetails map[string]*string `json:"vifDetails,omitempty" tf:"vif_details,omitempty"`

	// (Computed) The VNIC type of the port binding.
	VifType *string `json:"vifType,omitempty" tf:"vif_type,omitempty"`

	// VNIC type for the port. Can either be direct,
	// direct-physical, macvtap, normal, baremetal or virtio-forwarder.
	// Default value is normal. It can be updated on unbound ports only.
	VnicType *string `json:"vnicType,omitempty" tf:"vnic_type,omitempty"`
}

type BindingParameters struct {

	// The ID of the host to allocate port on.
	// +kubebuilder:validation:Optional
	HostID *string `json:"hostId,omitempty" tf:"host_id,omitempty"`

	// Custom data to be passed as binding:profile. Data
	// must be passed as JSON.
	// +kubebuilder:validation:Optional
	Profile *string `json:"profile,omitempty" tf:"profile,omitempty"`

	// VNIC type for the port. Can either be direct,
	// direct-physical, macvtap, normal, baremetal or virtio-forwarder.
	// Default value is normal. It can be updated on unbound ports only.
	// +kubebuilder:validation:Optional
	VnicType *string `json:"vnicType,omitempty" tf:"vnic_type,omitempty"`
}

type ExtraDHCPOptionInitParameters struct {

	// IP protocol version. Defaults to 4.
	IPVersion *float64 `json:"ipVersion,omitempty" tf:"ip_version,omitempty"`

	// Name of the DHCP option.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Value of the DHCP option.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ExtraDHCPOptionObservation struct {

	// IP protocol version. Defaults to 4.
	IPVersion *float64 `json:"ipVersion,omitempty" tf:"ip_version,omitempty"`

	// Name of the DHCP option.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Value of the DHCP option.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ExtraDHCPOptionParameters struct {

	// IP protocol version. Defaults to 4.
	// +kubebuilder:validation:Optional
	IPVersion *float64 `json:"ipVersion,omitempty" tf:"ip_version,omitempty"`

	// Name of the DHCP option.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// Value of the DHCP option.
	// +kubebuilder:validation:Optional
	Value *string `json:"value" tf:"value,omitempty"`
}

type FixedIPInitParameters struct {

	// IP address desired in the subnet for this port. If
	// you don't specify ip_address, an available IP address from the specified
	// subnet will be allocated to this port. This field will not be populated if it
	// is left blank or omitted. To retrieve the assigned IP address, use the
	// all_fixed_ips attribute.
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// Subnet in which to allocate IP address for
	// this port.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-openstack/apis/networking/v1alpha1.SubnetV2
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a SubnetV2 in networking to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a SubnetV2 in networking to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`
}

type FixedIPObservation struct {

	// IP address desired in the subnet for this port. If
	// you don't specify ip_address, an available IP address from the specified
	// subnet will be allocated to this port. This field will not be populated if it
	// is left blank or omitted. To retrieve the assigned IP address, use the
	// all_fixed_ips attribute.
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// Subnet in which to allocate IP address for
	// this port.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`
}

type FixedIPParameters struct {

	// IP address desired in the subnet for this port. If
	// you don't specify ip_address, an available IP address from the specified
	// subnet will be allocated to this port. This field will not be populated if it
	// is left blank or omitted. To retrieve the assigned IP address, use the
	// all_fixed_ips attribute.
	// +kubebuilder:validation:Optional
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// Subnet in which to allocate IP address for
	// this port.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-openstack/apis/networking/v1alpha1.SubnetV2
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a SubnetV2 in networking to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a SubnetV2 in networking to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`
}

type PortV2InitParameters struct {

	// Administrative up/down status for the port
	// (must be true or false if provided). Changing this updates the
	// admin_state_up of an existing port.
	AdminStateUp *bool `json:"adminStateUp,omitempty" tf:"admin_state_up,omitempty"`

	// An IP/MAC Address pair of additional IP
	// addresses that can be active on this port. The structure is described
	// below.
	AllowedAddressPairs []AllowedAddressPairsInitParameters `json:"allowedAddressPairs,omitempty" tf:"allowed_address_pairs,omitempty"`

	// The port binding allows to specify binding information
	// for the port. The structure is described below.
	Binding *BindingInitParameters `json:"binding,omitempty" tf:"binding,omitempty"`

	// The port DNS name. Available, when Neutron DNS extension
	// is enabled.
	DNSName *string `json:"dnsName,omitempty" tf:"dns_name,omitempty"`

	// Human-readable description of the port. Changing
	// this updates the description of an existing port.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the device attached to the port. Changing this
	// creates a new port.
	DeviceID *string `json:"deviceId,omitempty" tf:"device_id,omitempty"`

	// The device owner of the port. Changing this creates
	// a new port.
	DeviceOwner *string `json:"deviceOwner,omitempty" tf:"device_owner,omitempty"`

	// An extra DHCP option that needs to be configured
	// on the port. The structure is described below. Can be specified multiple
	// times.
	ExtraDHCPOption []ExtraDHCPOptionInitParameters `json:"extraDhcpOption,omitempty" tf:"extra_dhcp_option,omitempty"`

	// An array of desired IPs for
	// this port. The structure is described below.
	FixedIP []FixedIPInitParameters `json:"fixedIp,omitempty" tf:"fixed_ip,omitempty"`

	// Specify a specific MAC address for the port. Changing
	// this creates a new port.
	MacAddress *string `json:"macAddress,omitempty" tf:"mac_address,omitempty"`

	// A unique name for the port. Changing this
	// updates the name of an existing port.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the network to attach the port to. Changing
	// this creates a new port.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-openstack/apis/networking/v1alpha1.NetworkV2
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// Reference to a NetworkV2 in networking to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDRef *v1.Reference `json:"networkIdRef,omitempty" tf:"-"`

	// Selector for a NetworkV2 in networking to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDSelector *v1.Selector `json:"networkIdSelector,omitempty" tf:"-"`

	// Create a port with no fixed
	// IP address. This will also remove any fixed IPs previously set on a port. true
	// is the only valid value for this argument.
	NoFixedIP *bool `json:"noFixedIp,omitempty" tf:"no_fixed_ip,omitempty"`

	// If set to
	// true, then no security groups are applied to the port. If set to false and
	// no security_group_ids are specified, then the port will yield to the default
	// behavior of the Networking service, which is to usually apply the "default"
	// security group.
	NoSecurityGroups *bool `json:"noSecurityGroups,omitempty" tf:"no_security_groups,omitempty"`

	// Whether to explicitly enable or disable
	// port security on the port. Port Security is usually enabled by default, so
	// omitting argument will usually result in a value of true. Setting this
	// explicitly to false will disable port security. In order to disable port
	// security, the port must not have any security groups. Valid values are true
	// and false.
	PortSecurityEnabled *bool `json:"portSecurityEnabled,omitempty" tf:"port_security_enabled,omitempty"`

	// Reference to the associated QoS policy.
	QosPolicyID *string `json:"qosPolicyId,omitempty" tf:"qos_policy_id,omitempty"`

	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a port. If omitted, the
	// region argument of the provider is used. Changing this creates a new
	// port.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// A list
	// of security group IDs to apply to the port. The security groups must be
	// specified by ID and not name (as opposed to how they are configured with
	// the Compute Instance).
	// +listType=set
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// A set of string tags for the port.
	// +listType=set
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The owner of the port. Required if admin wants
	// to create a port for another tenant. Changing this creates a new port.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// Map of additional options.
	// +mapType=granular
	ValueSpecs map[string]*string `json:"valueSpecs,omitempty" tf:"value_specs,omitempty"`
}

type PortV2Observation struct {

	// Administrative up/down status for the port
	// (must be true or false if provided). Changing this updates the
	// admin_state_up of an existing port.
	AdminStateUp *bool `json:"adminStateUp,omitempty" tf:"admin_state_up,omitempty"`

	// The collection of Fixed IP addresses on the port in the
	// order returned by the Network v2 API.
	AllFixedIps []*string `json:"allFixedIps,omitempty" tf:"all_fixed_ips,omitempty"`

	// The collection of Security Group IDs on the port
	// which have been explicitly and implicitly added.
	// +listType=set
	AllSecurityGroupIds []*string `json:"allSecurityGroupIds,omitempty" tf:"all_security_group_ids,omitempty"`

	// The collection of tags assigned on the port, which have been
	// explicitly and implicitly added.
	// +listType=set
	AllTags []*string `json:"allTags,omitempty" tf:"all_tags,omitempty"`

	// An IP/MAC Address pair of additional IP
	// addresses that can be active on this port. The structure is described
	// below.
	AllowedAddressPairs []AllowedAddressPairsObservation `json:"allowedAddressPairs,omitempty" tf:"allowed_address_pairs,omitempty"`

	// The port binding allows to specify binding information
	// for the port. The structure is described below.
	Binding *BindingObservation `json:"binding,omitempty" tf:"binding,omitempty"`

	// The list of maps representing port DNS assignments.
	DNSAssignment []map[string]*string `json:"dnsAssignment,omitempty" tf:"dns_assignment,omitempty"`

	// The port DNS name. Available, when Neutron DNS extension
	// is enabled.
	DNSName *string `json:"dnsName,omitempty" tf:"dns_name,omitempty"`

	// Human-readable description of the port. Changing
	// this updates the description of an existing port.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the device attached to the port. Changing this
	// creates a new port.
	DeviceID *string `json:"deviceId,omitempty" tf:"device_id,omitempty"`

	// The device owner of the port. Changing this creates
	// a new port.
	DeviceOwner *string `json:"deviceOwner,omitempty" tf:"device_owner,omitempty"`

	// An extra DHCP option that needs to be configured
	// on the port. The structure is described below. Can be specified multiple
	// times.
	ExtraDHCPOption []ExtraDHCPOptionObservation `json:"extraDhcpOption,omitempty" tf:"extra_dhcp_option,omitempty"`

	// An array of desired IPs for
	// this port. The structure is described below.
	FixedIP []FixedIPObservation `json:"fixedIp,omitempty" tf:"fixed_ip,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specify a specific MAC address for the port. Changing
	// this creates a new port.
	MacAddress *string `json:"macAddress,omitempty" tf:"mac_address,omitempty"`

	// A unique name for the port. Changing this
	// updates the name of an existing port.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the network to attach the port to. Changing
	// this creates a new port.
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// Create a port with no fixed
	// IP address. This will also remove any fixed IPs previously set on a port. true
	// is the only valid value for this argument.
	NoFixedIP *bool `json:"noFixedIp,omitempty" tf:"no_fixed_ip,omitempty"`

	// If set to
	// true, then no security groups are applied to the port. If set to false and
	// no security_group_ids are specified, then the port will yield to the default
	// behavior of the Networking service, which is to usually apply the "default"
	// security group.
	NoSecurityGroups *bool `json:"noSecurityGroups,omitempty" tf:"no_security_groups,omitempty"`

	// Whether to explicitly enable or disable
	// port security on the port. Port Security is usually enabled by default, so
	// omitting argument will usually result in a value of true. Setting this
	// explicitly to false will disable port security. In order to disable port
	// security, the port must not have any security groups. Valid values are true
	// and false.
	PortSecurityEnabled *bool `json:"portSecurityEnabled,omitempty" tf:"port_security_enabled,omitempty"`

	// Reference to the associated QoS policy.
	QosPolicyID *string `json:"qosPolicyId,omitempty" tf:"qos_policy_id,omitempty"`

	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a port. If omitted, the
	// region argument of the provider is used. Changing this creates a new
	// port.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// A list
	// of security group IDs to apply to the port. The security groups must be
	// specified by ID and not name (as opposed to how they are configured with
	// the Compute Instance).
	// +listType=set
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// A set of string tags for the port.
	// +listType=set
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The owner of the port. Required if admin wants
	// to create a port for another tenant. Changing this creates a new port.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// Map of additional options.
	// +mapType=granular
	ValueSpecs map[string]*string `json:"valueSpecs,omitempty" tf:"value_specs,omitempty"`
}

type PortV2Parameters struct {

	// Administrative up/down status for the port
	// (must be true or false if provided). Changing this updates the
	// admin_state_up of an existing port.
	// +kubebuilder:validation:Optional
	AdminStateUp *bool `json:"adminStateUp,omitempty" tf:"admin_state_up,omitempty"`

	// An IP/MAC Address pair of additional IP
	// addresses that can be active on this port. The structure is described
	// below.
	// +kubebuilder:validation:Optional
	AllowedAddressPairs []AllowedAddressPairsParameters `json:"allowedAddressPairs,omitempty" tf:"allowed_address_pairs,omitempty"`

	// The port binding allows to specify binding information
	// for the port. The structure is described below.
	// +kubebuilder:validation:Optional
	Binding *BindingParameters `json:"binding,omitempty" tf:"binding,omitempty"`

	// The port DNS name. Available, when Neutron DNS extension
	// is enabled.
	// +kubebuilder:validation:Optional
	DNSName *string `json:"dnsName,omitempty" tf:"dns_name,omitempty"`

	// Human-readable description of the port. Changing
	// this updates the description of an existing port.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the device attached to the port. Changing this
	// creates a new port.
	// +kubebuilder:validation:Optional
	DeviceID *string `json:"deviceId,omitempty" tf:"device_id,omitempty"`

	// The device owner of the port. Changing this creates
	// a new port.
	// +kubebuilder:validation:Optional
	DeviceOwner *string `json:"deviceOwner,omitempty" tf:"device_owner,omitempty"`

	// An extra DHCP option that needs to be configured
	// on the port. The structure is described below. Can be specified multiple
	// times.
	// +kubebuilder:validation:Optional
	ExtraDHCPOption []ExtraDHCPOptionParameters `json:"extraDhcpOption,omitempty" tf:"extra_dhcp_option,omitempty"`

	// An array of desired IPs for
	// this port. The structure is described below.
	// +kubebuilder:validation:Optional
	FixedIP []FixedIPParameters `json:"fixedIp,omitempty" tf:"fixed_ip,omitempty"`

	// Specify a specific MAC address for the port. Changing
	// this creates a new port.
	// +kubebuilder:validation:Optional
	MacAddress *string `json:"macAddress,omitempty" tf:"mac_address,omitempty"`

	// A unique name for the port. Changing this
	// updates the name of an existing port.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the network to attach the port to. Changing
	// this creates a new port.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-openstack/apis/networking/v1alpha1.NetworkV2
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// Reference to a NetworkV2 in networking to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDRef *v1.Reference `json:"networkIdRef,omitempty" tf:"-"`

	// Selector for a NetworkV2 in networking to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDSelector *v1.Selector `json:"networkIdSelector,omitempty" tf:"-"`

	// Create a port with no fixed
	// IP address. This will also remove any fixed IPs previously set on a port. true
	// is the only valid value for this argument.
	// +kubebuilder:validation:Optional
	NoFixedIP *bool `json:"noFixedIp,omitempty" tf:"no_fixed_ip,omitempty"`

	// If set to
	// true, then no security groups are applied to the port. If set to false and
	// no security_group_ids are specified, then the port will yield to the default
	// behavior of the Networking service, which is to usually apply the "default"
	// security group.
	// +kubebuilder:validation:Optional
	NoSecurityGroups *bool `json:"noSecurityGroups,omitempty" tf:"no_security_groups,omitempty"`

	// Whether to explicitly enable or disable
	// port security on the port. Port Security is usually enabled by default, so
	// omitting argument will usually result in a value of true. Setting this
	// explicitly to false will disable port security. In order to disable port
	// security, the port must not have any security groups. Valid values are true
	// and false.
	// +kubebuilder:validation:Optional
	PortSecurityEnabled *bool `json:"portSecurityEnabled,omitempty" tf:"port_security_enabled,omitempty"`

	// Reference to the associated QoS policy.
	// +kubebuilder:validation:Optional
	QosPolicyID *string `json:"qosPolicyId,omitempty" tf:"qos_policy_id,omitempty"`

	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a port. If omitted, the
	// region argument of the provider is used. Changing this creates a new
	// port.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// A list
	// of security group IDs to apply to the port. The security groups must be
	// specified by ID and not name (as opposed to how they are configured with
	// the Compute Instance).
	// +kubebuilder:validation:Optional
	// +listType=set
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// A set of string tags for the port.
	// +kubebuilder:validation:Optional
	// +listType=set
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The owner of the port. Required if admin wants
	// to create a port for another tenant. Changing this creates a new port.
	// +kubebuilder:validation:Optional
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// Map of additional options.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	ValueSpecs map[string]*string `json:"valueSpecs,omitempty" tf:"value_specs,omitempty"`
}

// PortV2Spec defines the desired state of PortV2
type PortV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PortV2Parameters `json:"forProvider"`
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
	InitProvider PortV2InitParameters `json:"initProvider,omitempty"`
}

// PortV2Status defines the observed state of PortV2.
type PortV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PortV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// PortV2 is the Schema for the PortV2s API. Manages a V2 port resource within OpenStack.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,openstack}
type PortV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PortV2Spec   `json:"spec"`
	Status            PortV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PortV2List contains a list of PortV2s
type PortV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PortV2 `json:"items"`
}

// Repository type metadata.
var (
	PortV2_Kind             = "PortV2"
	PortV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PortV2_Kind}.String()
	PortV2_KindAPIVersion   = PortV2_Kind + "." + CRDGroupVersion.String()
	PortV2_GroupVersionKind = CRDGroupVersion.WithKind(PortV2_Kind)
)

func init() {
	SchemeBuilder.Register(&PortV2{}, &PortV2List{})
}
