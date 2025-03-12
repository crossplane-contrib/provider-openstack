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

type SecgroupRuleV2InitParameters struct {

	// A description of the rule. Changing this creates a new security group rule.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The direction of the rule, valid values are ingress
	// or egress. Changing this creates a new security group rule.
	Direction *string `json:"direction,omitempty" tf:"direction,omitempty"`

	// The layer 3 protocol type, valid values are IPv4
	// or IPv6. Changing this creates a new security group rule.
	Ethertype *string `json:"ethertype,omitempty" tf:"ethertype,omitempty"`

	// The higher part of the allowed port range, valid
	// integer value needs to be between 1 and 65535. Changing this creates a new
	// security group rule.
	PortRangeMax *float64 `json:"portRangeMax,omitempty" tf:"port_range_max,omitempty"`

	// The lower part of the allowed port range, valid
	// integer value needs to be between 1 and 65535. Changing this creates a new
	// security group rule.
	PortRangeMin *float64 `json:"portRangeMin,omitempty" tf:"port_range_min,omitempty"`

	// The layer 4 protocol type, valid values are following. Changing this creates a new security group rule. This is required if you want to specify a port range.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// The region in which to obtain the V2 networking client.
	// A networking client is needed to create a port. If omitted, the
	// region argument of the provider is used. Changing this creates a new
	// security group rule.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The remote group id, the value needs to be an
	// Openstack ID of a security group in the same tenant. Changing this creates
	// a new security group rule.
	RemoteGroupID *string `json:"remoteGroupId,omitempty" tf:"remote_group_id,omitempty"`

	// The remote CIDR, the value needs to be a valid
	// CIDR (i.e. 192.168.0.0/16). Changing this creates a new security group rule.
	RemoteIPPrefix *string `json:"remoteIpPrefix,omitempty" tf:"remote_ip_prefix,omitempty"`

	// The security group id the rule should belong
	// to, the value needs to be an Openstack ID of a security group in the same
	// tenant. Changing this creates a new security group rule.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-openstack/apis/networking/v1alpha1.SecgroupV2
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	SecurityGroupID *string `json:"securityGroupId,omitempty" tf:"security_group_id,omitempty"`

	// Reference to a SecgroupV2 in networking to populate securityGroupId.
	// +kubebuilder:validation:Optional
	SecurityGroupIDRef *v1.Reference `json:"securityGroupIdRef,omitempty" tf:"-"`

	// Selector for a SecgroupV2 in networking to populate securityGroupId.
	// +kubebuilder:validation:Optional
	SecurityGroupIDSelector *v1.Selector `json:"securityGroupIdSelector,omitempty" tf:"-"`

	// The owner of the security group. Required if admin
	// wants to create a port for another tenant. Changing this creates a new
	// security group rule.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type SecgroupRuleV2Observation struct {

	// A description of the rule. Changing this creates a new security group rule.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The direction of the rule, valid values are ingress
	// or egress. Changing this creates a new security group rule.
	Direction *string `json:"direction,omitempty" tf:"direction,omitempty"`

	// The layer 3 protocol type, valid values are IPv4
	// or IPv6. Changing this creates a new security group rule.
	Ethertype *string `json:"ethertype,omitempty" tf:"ethertype,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The higher part of the allowed port range, valid
	// integer value needs to be between 1 and 65535. Changing this creates a new
	// security group rule.
	PortRangeMax *float64 `json:"portRangeMax,omitempty" tf:"port_range_max,omitempty"`

	// The lower part of the allowed port range, valid
	// integer value needs to be between 1 and 65535. Changing this creates a new
	// security group rule.
	PortRangeMin *float64 `json:"portRangeMin,omitempty" tf:"port_range_min,omitempty"`

	// The layer 4 protocol type, valid values are following. Changing this creates a new security group rule. This is required if you want to specify a port range.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// The region in which to obtain the V2 networking client.
	// A networking client is needed to create a port. If omitted, the
	// region argument of the provider is used. Changing this creates a new
	// security group rule.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The remote group id, the value needs to be an
	// Openstack ID of a security group in the same tenant. Changing this creates
	// a new security group rule.
	RemoteGroupID *string `json:"remoteGroupId,omitempty" tf:"remote_group_id,omitempty"`

	// The remote CIDR, the value needs to be a valid
	// CIDR (i.e. 192.168.0.0/16). Changing this creates a new security group rule.
	RemoteIPPrefix *string `json:"remoteIpPrefix,omitempty" tf:"remote_ip_prefix,omitempty"`

	// The security group id the rule should belong
	// to, the value needs to be an Openstack ID of a security group in the same
	// tenant. Changing this creates a new security group rule.
	SecurityGroupID *string `json:"securityGroupId,omitempty" tf:"security_group_id,omitempty"`

	// The owner of the security group. Required if admin
	// wants to create a port for another tenant. Changing this creates a new
	// security group rule.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type SecgroupRuleV2Parameters struct {

	// A description of the rule. Changing this creates a new security group rule.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The direction of the rule, valid values are ingress
	// or egress. Changing this creates a new security group rule.
	// +kubebuilder:validation:Optional
	Direction *string `json:"direction,omitempty" tf:"direction,omitempty"`

	// The layer 3 protocol type, valid values are IPv4
	// or IPv6. Changing this creates a new security group rule.
	// +kubebuilder:validation:Optional
	Ethertype *string `json:"ethertype,omitempty" tf:"ethertype,omitempty"`

	// The higher part of the allowed port range, valid
	// integer value needs to be between 1 and 65535. Changing this creates a new
	// security group rule.
	// +kubebuilder:validation:Optional
	PortRangeMax *float64 `json:"portRangeMax,omitempty" tf:"port_range_max,omitempty"`

	// The lower part of the allowed port range, valid
	// integer value needs to be between 1 and 65535. Changing this creates a new
	// security group rule.
	// +kubebuilder:validation:Optional
	PortRangeMin *float64 `json:"portRangeMin,omitempty" tf:"port_range_min,omitempty"`

	// The layer 4 protocol type, valid values are following. Changing this creates a new security group rule. This is required if you want to specify a port range.
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// The region in which to obtain the V2 networking client.
	// A networking client is needed to create a port. If omitted, the
	// region argument of the provider is used. Changing this creates a new
	// security group rule.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The remote group id, the value needs to be an
	// Openstack ID of a security group in the same tenant. Changing this creates
	// a new security group rule.
	// +kubebuilder:validation:Optional
	RemoteGroupID *string `json:"remoteGroupId,omitempty" tf:"remote_group_id,omitempty"`

	// The remote CIDR, the value needs to be a valid
	// CIDR (i.e. 192.168.0.0/16). Changing this creates a new security group rule.
	// +kubebuilder:validation:Optional
	RemoteIPPrefix *string `json:"remoteIpPrefix,omitempty" tf:"remote_ip_prefix,omitempty"`

	// The security group id the rule should belong
	// to, the value needs to be an Openstack ID of a security group in the same
	// tenant. Changing this creates a new security group rule.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-openstack/apis/networking/v1alpha1.SecgroupV2
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SecurityGroupID *string `json:"securityGroupId,omitempty" tf:"security_group_id,omitempty"`

	// Reference to a SecgroupV2 in networking to populate securityGroupId.
	// +kubebuilder:validation:Optional
	SecurityGroupIDRef *v1.Reference `json:"securityGroupIdRef,omitempty" tf:"-"`

	// Selector for a SecgroupV2 in networking to populate securityGroupId.
	// +kubebuilder:validation:Optional
	SecurityGroupIDSelector *v1.Selector `json:"securityGroupIdSelector,omitempty" tf:"-"`

	// The owner of the security group. Required if admin
	// wants to create a port for another tenant. Changing this creates a new
	// security group rule.
	// +kubebuilder:validation:Optional
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

// SecgroupRuleV2Spec defines the desired state of SecgroupRuleV2
type SecgroupRuleV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecgroupRuleV2Parameters `json:"forProvider"`
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
	InitProvider SecgroupRuleV2InitParameters `json:"initProvider,omitempty"`
}

// SecgroupRuleV2Status defines the observed state of SecgroupRuleV2.
type SecgroupRuleV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecgroupRuleV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SecgroupRuleV2 is the Schema for the SecgroupRuleV2s API. Manages a V2 Neutron security group rule resource within OpenStack.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,openstack}
type SecgroupRuleV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.direction) || (has(self.initProvider) && has(self.initProvider.direction))",message="spec.forProvider.direction is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.ethertype) || (has(self.initProvider) && has(self.initProvider.ethertype))",message="spec.forProvider.ethertype is a required parameter"
	Spec   SecgroupRuleV2Spec   `json:"spec"`
	Status SecgroupRuleV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecgroupRuleV2List contains a list of SecgroupRuleV2s
type SecgroupRuleV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecgroupRuleV2 `json:"items"`
}

// Repository type metadata.
var (
	SecgroupRuleV2_Kind             = "SecgroupRuleV2"
	SecgroupRuleV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SecgroupRuleV2_Kind}.String()
	SecgroupRuleV2_KindAPIVersion   = SecgroupRuleV2_Kind + "." + CRDGroupVersion.String()
	SecgroupRuleV2_GroupVersionKind = CRDGroupVersion.WithKind(SecgroupRuleV2_Kind)
)

func init() {
	SchemeBuilder.Register(&SecgroupRuleV2{}, &SecgroupRuleV2List{})
}
