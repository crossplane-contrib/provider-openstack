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

type MemberInitParameters struct {

	// The IP address of the members to receive traffic from
	// the load balancer.
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// The administrative state of the member.
	// A valid value is true (UP) or false (DOWN). Defaults to true.
	AdminStateUp *bool `json:"adminStateUp,omitempty" tf:"admin_state_up,omitempty"`

	// A bool that indicates whether the member is
	// backup. Requires octavia minor version 2.1 or later.
	Backup *bool `json:"backup,omitempty" tf:"backup,omitempty"`

	// An alternate IP address used for health
	// monitoring a backend member.
	MonitorAddress *string `json:"monitorAddress,omitempty" tf:"monitor_address,omitempty"`

	// An alternate protocol port used for health
	// monitoring a backend member.
	MonitorPort *float64 `json:"monitorPort,omitempty" tf:"monitor_port,omitempty"`

	// Human-readable name for the member.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The port on which to listen for client traffic.
	ProtocolPort *float64 `json:"protocolPort,omitempty" tf:"protocol_port,omitempty"`

	// The subnet in which to access the member.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// A positive integer value that indicates the relative
	// portion of traffic that this members should receive from the pool. For
	// example, a member with a weight of 10 receives five times as much traffic
	// as a member with a weight of 2. Defaults to 1.
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type MemberObservation struct {

	// The IP address of the members to receive traffic from
	// the load balancer.
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// The administrative state of the member.
	// A valid value is true (UP) or false (DOWN). Defaults to true.
	AdminStateUp *bool `json:"adminStateUp,omitempty" tf:"admin_state_up,omitempty"`

	// A bool that indicates whether the member is
	// backup. Requires octavia minor version 2.1 or later.
	Backup *bool `json:"backup,omitempty" tf:"backup,omitempty"`

	// The unique ID for the members.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// An alternate IP address used for health
	// monitoring a backend member.
	MonitorAddress *string `json:"monitorAddress,omitempty" tf:"monitor_address,omitempty"`

	// An alternate protocol port used for health
	// monitoring a backend member.
	MonitorPort *float64 `json:"monitorPort,omitempty" tf:"monitor_port,omitempty"`

	// Human-readable name for the member.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The port on which to listen for client traffic.
	ProtocolPort *float64 `json:"protocolPort,omitempty" tf:"protocol_port,omitempty"`

	// The subnet in which to access the member.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// A positive integer value that indicates the relative
	// portion of traffic that this members should receive from the pool. For
	// example, a member with a weight of 10 receives five times as much traffic
	// as a member with a weight of 2. Defaults to 1.
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type MemberParameters struct {

	// The IP address of the members to receive traffic from
	// the load balancer.
	// +kubebuilder:validation:Optional
	Address *string `json:"address" tf:"address,omitempty"`

	// The administrative state of the member.
	// A valid value is true (UP) or false (DOWN). Defaults to true.
	// +kubebuilder:validation:Optional
	AdminStateUp *bool `json:"adminStateUp,omitempty" tf:"admin_state_up,omitempty"`

	// A bool that indicates whether the member is
	// backup. Requires octavia minor version 2.1 or later.
	// +kubebuilder:validation:Optional
	Backup *bool `json:"backup,omitempty" tf:"backup,omitempty"`

	// An alternate IP address used for health
	// monitoring a backend member.
	// +kubebuilder:validation:Optional
	MonitorAddress *string `json:"monitorAddress,omitempty" tf:"monitor_address,omitempty"`

	// An alternate protocol port used for health
	// monitoring a backend member.
	// +kubebuilder:validation:Optional
	MonitorPort *float64 `json:"monitorPort,omitempty" tf:"monitor_port,omitempty"`

	// Human-readable name for the member.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The port on which to listen for client traffic.
	// +kubebuilder:validation:Optional
	ProtocolPort *float64 `json:"protocolPort" tf:"protocol_port,omitempty"`

	// The subnet in which to access the member.
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// A positive integer value that indicates the relative
	// portion of traffic that this members should receive from the pool. For
	// example, a member with a weight of 10 receives five times as much traffic
	// as a member with a weight of 2. Defaults to 1.
	// +kubebuilder:validation:Optional
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type MembersV2InitParameters struct {

	// A set of dictionaries containing member parameters. The
	// structure is described below.
	Member []MemberInitParameters `json:"member,omitempty" tf:"member,omitempty"`

	// The id of the pool that members will be assigned to.
	// Changing this creates a new members resource.
	PoolID *string `json:"poolId,omitempty" tf:"pool_id,omitempty"`

	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create pool members. If omitted, the
	// region argument of the provider is used. Changing this creates a new
	// members resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type MembersV2Observation struct {

	// The unique ID for the members.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A set of dictionaries containing member parameters. The
	// structure is described below.
	Member []MemberObservation `json:"member,omitempty" tf:"member,omitempty"`

	// The id of the pool that members will be assigned to.
	// Changing this creates a new members resource.
	PoolID *string `json:"poolId,omitempty" tf:"pool_id,omitempty"`

	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create pool members. If omitted, the
	// region argument of the provider is used. Changing this creates a new
	// members resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type MembersV2Parameters struct {

	// A set of dictionaries containing member parameters. The
	// structure is described below.
	// +kubebuilder:validation:Optional
	Member []MemberParameters `json:"member,omitempty" tf:"member,omitempty"`

	// The id of the pool that members will be assigned to.
	// Changing this creates a new members resource.
	// +kubebuilder:validation:Optional
	PoolID *string `json:"poolId,omitempty" tf:"pool_id,omitempty"`

	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create pool members. If omitted, the
	// region argument of the provider is used. Changing this creates a new
	// members resource.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

// MembersV2Spec defines the desired state of MembersV2
type MembersV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MembersV2Parameters `json:"forProvider"`
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
	InitProvider MembersV2InitParameters `json:"initProvider,omitempty"`
}

// MembersV2Status defines the observed state of MembersV2.
type MembersV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MembersV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MembersV2 is the Schema for the MembersV2s API. Manages a V2 members resource within OpenStack (batch members update).
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,openstack}
type MembersV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.poolId) || (has(self.initProvider) && has(self.initProvider.poolId))",message="spec.forProvider.poolId is a required parameter"
	Spec   MembersV2Spec   `json:"spec"`
	Status MembersV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MembersV2List contains a list of MembersV2s
type MembersV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MembersV2 `json:"items"`
}

// Repository type metadata.
var (
	MembersV2_Kind             = "MembersV2"
	MembersV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MembersV2_Kind}.String()
	MembersV2_KindAPIVersion   = MembersV2_Kind + "." + CRDGroupVersion.String()
	MembersV2_GroupVersionKind = CRDGroupVersion.WithKind(MembersV2_Kind)
)

func init() {
	SchemeBuilder.Register(&MembersV2{}, &MembersV2List{})
}
