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

type RouterRouteV2InitParameters struct {

	// CIDR block to match on the packet’s destination IP. Changing
	// this creates a new routing entry.
	DestinationCidr *string `json:"destinationCidr,omitempty" tf:"destination_cidr,omitempty"`

	// IP address of the next hop gateway.  Changing
	// this creates a new routing entry.
	NextHop *string `json:"nextHop,omitempty" tf:"next_hop,omitempty"`

	// The region in which to obtain the V2 networking client.
	// A networking client is needed to configure a routing entry on a router. If omitted, the
	// region argument of the provider is used. Changing this creates a new
	// routing entry.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// ID of the router this routing entry belongs to. Changing
	// this creates a new routing entry.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-openstack/apis/networking/v1alpha1.RouterV2
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	RouterID *string `json:"routerId,omitempty" tf:"router_id,omitempty"`

	// Reference to a RouterV2 in networking to populate routerId.
	// +kubebuilder:validation:Optional
	RouterIDRef *v1.Reference `json:"routerIdRef,omitempty" tf:"-"`

	// Selector for a RouterV2 in networking to populate routerId.
	// +kubebuilder:validation:Optional
	RouterIDSelector *v1.Selector `json:"routerIdSelector,omitempty" tf:"-"`
}

type RouterRouteV2Observation struct {

	// CIDR block to match on the packet’s destination IP. Changing
	// this creates a new routing entry.
	DestinationCidr *string `json:"destinationCidr,omitempty" tf:"destination_cidr,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// IP address of the next hop gateway.  Changing
	// this creates a new routing entry.
	NextHop *string `json:"nextHop,omitempty" tf:"next_hop,omitempty"`

	// The region in which to obtain the V2 networking client.
	// A networking client is needed to configure a routing entry on a router. If omitted, the
	// region argument of the provider is used. Changing this creates a new
	// routing entry.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// ID of the router this routing entry belongs to. Changing
	// this creates a new routing entry.
	RouterID *string `json:"routerId,omitempty" tf:"router_id,omitempty"`
}

type RouterRouteV2Parameters struct {

	// CIDR block to match on the packet’s destination IP. Changing
	// this creates a new routing entry.
	// +kubebuilder:validation:Optional
	DestinationCidr *string `json:"destinationCidr,omitempty" tf:"destination_cidr,omitempty"`

	// IP address of the next hop gateway.  Changing
	// this creates a new routing entry.
	// +kubebuilder:validation:Optional
	NextHop *string `json:"nextHop,omitempty" tf:"next_hop,omitempty"`

	// The region in which to obtain the V2 networking client.
	// A networking client is needed to configure a routing entry on a router. If omitted, the
	// region argument of the provider is used. Changing this creates a new
	// routing entry.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// ID of the router this routing entry belongs to. Changing
	// this creates a new routing entry.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-openstack/apis/networking/v1alpha1.RouterV2
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	RouterID *string `json:"routerId,omitempty" tf:"router_id,omitempty"`

	// Reference to a RouterV2 in networking to populate routerId.
	// +kubebuilder:validation:Optional
	RouterIDRef *v1.Reference `json:"routerIdRef,omitempty" tf:"-"`

	// Selector for a RouterV2 in networking to populate routerId.
	// +kubebuilder:validation:Optional
	RouterIDSelector *v1.Selector `json:"routerIdSelector,omitempty" tf:"-"`
}

// RouterRouteV2Spec defines the desired state of RouterRouteV2
type RouterRouteV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RouterRouteV2Parameters `json:"forProvider"`
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
	InitProvider RouterRouteV2InitParameters `json:"initProvider,omitempty"`
}

// RouterRouteV2Status defines the observed state of RouterRouteV2.
type RouterRouteV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RouterRouteV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// RouterRouteV2 is the Schema for the RouterRouteV2s API. Creates a routing entry on a OpenStack V2 router.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,openstack}
type RouterRouteV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.destinationCidr) || (has(self.initProvider) && has(self.initProvider.destinationCidr))",message="spec.forProvider.destinationCidr is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.nextHop) || (has(self.initProvider) && has(self.initProvider.nextHop))",message="spec.forProvider.nextHop is a required parameter"
	Spec   RouterRouteV2Spec   `json:"spec"`
	Status RouterRouteV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RouterRouteV2List contains a list of RouterRouteV2s
type RouterRouteV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RouterRouteV2 `json:"items"`
}

// Repository type metadata.
var (
	RouterRouteV2_Kind             = "RouterRouteV2"
	RouterRouteV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RouterRouteV2_Kind}.String()
	RouterRouteV2_KindAPIVersion   = RouterRouteV2_Kind + "." + CRDGroupVersion.String()
	RouterRouteV2_GroupVersionKind = CRDGroupVersion.WithKind(RouterRouteV2_Kind)
)

func init() {
	SchemeBuilder.Register(&RouterRouteV2{}, &RouterRouteV2List{})
}
