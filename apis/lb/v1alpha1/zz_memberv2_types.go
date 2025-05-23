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

type MemberV2InitParameters struct {

	// The IP address of the member to receive traffic from
	// the load balancer. Changing this creates a new member.
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// The administrative state of the member.
	// A valid value is true (UP) or false (DOWN). Defaults to true.
	AdminStateUp *bool `json:"adminStateUp,omitempty" tf:"admin_state_up,omitempty"`

	// Boolean that indicates whether that member works as a backup or not. Available
	// only for Octavia >= 2.1.
	Backup *bool `json:"backup,omitempty" tf:"backup,omitempty"`

	// An alternate IP address used for health monitoring a backend member.
	// Available only for Octavia
	MonitorAddress *string `json:"monitorAddress,omitempty" tf:"monitor_address,omitempty"`

	// An alternate protocol port used for health monitoring a backend member.
	// Available only for Octavia
	MonitorPort *float64 `json:"monitorPort,omitempty" tf:"monitor_port,omitempty"`

	// Human-readable name for the member.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The id of the pool that this member will be assigned
	// to. Changing this creates a new member.
	PoolID *string `json:"poolId,omitempty" tf:"pool_id,omitempty"`

	// The port on which to listen for client traffic.
	// Changing this creates a new member.
	ProtocolPort *float64 `json:"protocolPort,omitempty" tf:"protocol_port,omitempty"`

	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a member. If omitted, the region
	// argument of the provider is used. Changing this creates a new member.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The subnet in which to access the member. Changing
	// this creates a new member.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// A list of simple strings assigned to the member.
	// Available only for Octavia >= 2.5.
	// +listType=set
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Required for admins. The UUID of the tenant who owns
	// the member.  Only administrative users can specify a tenant UUID
	// other than their own. Changing this creates a new member.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// A positive integer value that indicates the relative
	// portion of traffic that this member should receive from the pool. For
	// example, a member with a weight of 10 receives five times as much traffic
	// as a member with a weight of 2. Defaults to 1.
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type MemberV2Observation struct {

	// The IP address of the member to receive traffic from
	// the load balancer. Changing this creates a new member.
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// The administrative state of the member.
	// A valid value is true (UP) or false (DOWN). Defaults to true.
	AdminStateUp *bool `json:"adminStateUp,omitempty" tf:"admin_state_up,omitempty"`

	// Boolean that indicates whether that member works as a backup or not. Available
	// only for Octavia >= 2.1.
	Backup *bool `json:"backup,omitempty" tf:"backup,omitempty"`

	// The unique ID for the member.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// An alternate IP address used for health monitoring a backend member.
	// Available only for Octavia
	MonitorAddress *string `json:"monitorAddress,omitempty" tf:"monitor_address,omitempty"`

	// An alternate protocol port used for health monitoring a backend member.
	// Available only for Octavia
	MonitorPort *float64 `json:"monitorPort,omitempty" tf:"monitor_port,omitempty"`

	// Human-readable name for the member.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The id of the pool that this member will be assigned
	// to. Changing this creates a new member.
	PoolID *string `json:"poolId,omitempty" tf:"pool_id,omitempty"`

	// The port on which to listen for client traffic.
	// Changing this creates a new member.
	ProtocolPort *float64 `json:"protocolPort,omitempty" tf:"protocol_port,omitempty"`

	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a member. If omitted, the region
	// argument of the provider is used. Changing this creates a new member.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The subnet in which to access the member. Changing
	// this creates a new member.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// A list of simple strings assigned to the member.
	// Available only for Octavia >= 2.5.
	// +listType=set
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Required for admins. The UUID of the tenant who owns
	// the member.  Only administrative users can specify a tenant UUID
	// other than their own. Changing this creates a new member.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// A positive integer value that indicates the relative
	// portion of traffic that this member should receive from the pool. For
	// example, a member with a weight of 10 receives five times as much traffic
	// as a member with a weight of 2. Defaults to 1.
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type MemberV2Parameters struct {

	// The IP address of the member to receive traffic from
	// the load balancer. Changing this creates a new member.
	// +kubebuilder:validation:Optional
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// The administrative state of the member.
	// A valid value is true (UP) or false (DOWN). Defaults to true.
	// +kubebuilder:validation:Optional
	AdminStateUp *bool `json:"adminStateUp,omitempty" tf:"admin_state_up,omitempty"`

	// Boolean that indicates whether that member works as a backup or not. Available
	// only for Octavia >= 2.1.
	// +kubebuilder:validation:Optional
	Backup *bool `json:"backup,omitempty" tf:"backup,omitempty"`

	// An alternate IP address used for health monitoring a backend member.
	// Available only for Octavia
	// +kubebuilder:validation:Optional
	MonitorAddress *string `json:"monitorAddress,omitempty" tf:"monitor_address,omitempty"`

	// An alternate protocol port used for health monitoring a backend member.
	// Available only for Octavia
	// +kubebuilder:validation:Optional
	MonitorPort *float64 `json:"monitorPort,omitempty" tf:"monitor_port,omitempty"`

	// Human-readable name for the member.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The id of the pool that this member will be assigned
	// to. Changing this creates a new member.
	// +kubebuilder:validation:Optional
	PoolID *string `json:"poolId,omitempty" tf:"pool_id,omitempty"`

	// The port on which to listen for client traffic.
	// Changing this creates a new member.
	// +kubebuilder:validation:Optional
	ProtocolPort *float64 `json:"protocolPort,omitempty" tf:"protocol_port,omitempty"`

	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a member. If omitted, the region
	// argument of the provider is used. Changing this creates a new member.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The subnet in which to access the member. Changing
	// this creates a new member.
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// A list of simple strings assigned to the member.
	// Available only for Octavia >= 2.5.
	// +kubebuilder:validation:Optional
	// +listType=set
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Required for admins. The UUID of the tenant who owns
	// the member.  Only administrative users can specify a tenant UUID
	// other than their own. Changing this creates a new member.
	// +kubebuilder:validation:Optional
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// A positive integer value that indicates the relative
	// portion of traffic that this member should receive from the pool. For
	// example, a member with a weight of 10 receives five times as much traffic
	// as a member with a weight of 2. Defaults to 1.
	// +kubebuilder:validation:Optional
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

// MemberV2Spec defines the desired state of MemberV2
type MemberV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MemberV2Parameters `json:"forProvider"`
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
	InitProvider MemberV2InitParameters `json:"initProvider,omitempty"`
}

// MemberV2Status defines the observed state of MemberV2.
type MemberV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MemberV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// MemberV2 is the Schema for the MemberV2s API. Manages a V2 member resource within OpenStack.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,openstack}
type MemberV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.address) || (has(self.initProvider) && has(self.initProvider.address))",message="spec.forProvider.address is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.poolId) || (has(self.initProvider) && has(self.initProvider.poolId))",message="spec.forProvider.poolId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.protocolPort) || (has(self.initProvider) && has(self.initProvider.protocolPort))",message="spec.forProvider.protocolPort is a required parameter"
	Spec   MemberV2Spec   `json:"spec"`
	Status MemberV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MemberV2List contains a list of MemberV2s
type MemberV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MemberV2 `json:"items"`
}

// Repository type metadata.
var (
	MemberV2_Kind             = "MemberV2"
	MemberV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MemberV2_Kind}.String()
	MemberV2_KindAPIVersion   = MemberV2_Kind + "." + CRDGroupVersion.String()
	MemberV2_GroupVersionKind = CRDGroupVersion.WithKind(MemberV2_Kind)
)

func init() {
	SchemeBuilder.Register(&MemberV2{}, &MemberV2List{})
}
