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

type FloatingipAssociateV2InitParameters struct {

	// The specific IP address to direct traffic to.
	FixedIP *string `json:"fixedIp,omitempty" tf:"fixed_ip,omitempty"`

	// The floating IP to associate.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-openstack/apis/networking/v1alpha1.FloatingipV2
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("address",false)
	FloatingIP *string `json:"floatingIp,omitempty" tf:"floating_ip,omitempty"`

	// Reference to a FloatingipV2 in networking to populate floatingIp.
	// +kubebuilder:validation:Optional
	FloatingIPRef *v1.Reference `json:"floatingIpRef,omitempty" tf:"-"`

	// Selector for a FloatingipV2 in networking to populate floatingIp.
	// +kubebuilder:validation:Optional
	FloatingIPSelector *v1.Selector `json:"floatingIpSelector,omitempty" tf:"-"`

	// The instance to associte the floating IP with.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-openstack/apis/compute/v1alpha1.InstanceV2
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Reference to a InstanceV2 in compute to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDRef *v1.Reference `json:"instanceIdRef,omitempty" tf:"-"`

	// Selector for a InstanceV2 in compute to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDSelector *v1.Selector `json:"instanceIdSelector,omitempty" tf:"-"`

	// The region in which to obtain the V2 Compute client.
	// Keypairs are associated with accounts, but a Compute client is needed to
	// create one. If omitted, the region argument of the provider is used.
	// Changing this creates a new floatingip_associate.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Defaults to false.
	WaitUntilAssociated *bool `json:"waitUntilAssociated,omitempty" tf:"wait_until_associated,omitempty"`
}

type FloatingipAssociateV2Observation struct {

	// The specific IP address to direct traffic to.
	FixedIP *string `json:"fixedIp,omitempty" tf:"fixed_ip,omitempty"`

	// The floating IP to associate.
	FloatingIP *string `json:"floatingIp,omitempty" tf:"floating_ip,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The instance to associte the floating IP with.
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// The region in which to obtain the V2 Compute client.
	// Keypairs are associated with accounts, but a Compute client is needed to
	// create one. If omitted, the region argument of the provider is used.
	// Changing this creates a new floatingip_associate.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Defaults to false.
	WaitUntilAssociated *bool `json:"waitUntilAssociated,omitempty" tf:"wait_until_associated,omitempty"`
}

type FloatingipAssociateV2Parameters struct {

	// The specific IP address to direct traffic to.
	// +kubebuilder:validation:Optional
	FixedIP *string `json:"fixedIp,omitempty" tf:"fixed_ip,omitempty"`

	// The floating IP to associate.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-openstack/apis/networking/v1alpha1.FloatingipV2
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("address",false)
	// +kubebuilder:validation:Optional
	FloatingIP *string `json:"floatingIp,omitempty" tf:"floating_ip,omitempty"`

	// Reference to a FloatingipV2 in networking to populate floatingIp.
	// +kubebuilder:validation:Optional
	FloatingIPRef *v1.Reference `json:"floatingIpRef,omitempty" tf:"-"`

	// Selector for a FloatingipV2 in networking to populate floatingIp.
	// +kubebuilder:validation:Optional
	FloatingIPSelector *v1.Selector `json:"floatingIpSelector,omitempty" tf:"-"`

	// The instance to associte the floating IP with.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-openstack/apis/compute/v1alpha1.InstanceV2
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Reference to a InstanceV2 in compute to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDRef *v1.Reference `json:"instanceIdRef,omitempty" tf:"-"`

	// Selector for a InstanceV2 in compute to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDSelector *v1.Selector `json:"instanceIdSelector,omitempty" tf:"-"`

	// The region in which to obtain the V2 Compute client.
	// Keypairs are associated with accounts, but a Compute client is needed to
	// create one. If omitted, the region argument of the provider is used.
	// Changing this creates a new floatingip_associate.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Defaults to false.
	// +kubebuilder:validation:Optional
	WaitUntilAssociated *bool `json:"waitUntilAssociated,omitempty" tf:"wait_until_associated,omitempty"`
}

// FloatingipAssociateV2Spec defines the desired state of FloatingipAssociateV2
type FloatingipAssociateV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FloatingipAssociateV2Parameters `json:"forProvider"`
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
	InitProvider FloatingipAssociateV2InitParameters `json:"initProvider,omitempty"`
}

// FloatingipAssociateV2Status defines the observed state of FloatingipAssociateV2.
type FloatingipAssociateV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FloatingipAssociateV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// FloatingipAssociateV2 is the Schema for the FloatingipAssociateV2s API. Associate a floating IP to an instance
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,openstack}
type FloatingipAssociateV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FloatingipAssociateV2Spec   `json:"spec"`
	Status            FloatingipAssociateV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FloatingipAssociateV2List contains a list of FloatingipAssociateV2s
type FloatingipAssociateV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FloatingipAssociateV2 `json:"items"`
}

// Repository type metadata.
var (
	FloatingipAssociateV2_Kind             = "FloatingipAssociateV2"
	FloatingipAssociateV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FloatingipAssociateV2_Kind}.String()
	FloatingipAssociateV2_KindAPIVersion   = FloatingipAssociateV2_Kind + "." + CRDGroupVersion.String()
	FloatingipAssociateV2_GroupVersionKind = CRDGroupVersion.WithKind(FloatingipAssociateV2_Kind)
)

func init() {
	SchemeBuilder.Register(&FloatingipAssociateV2{}, &FloatingipAssociateV2List{})
}
