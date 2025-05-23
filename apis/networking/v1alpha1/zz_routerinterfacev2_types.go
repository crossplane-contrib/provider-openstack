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

type RouterInterfaceV2InitParameters struct {

	// A boolean indicating whether the routes from the
	// corresponding router ID should be deleted so that the router interface can
	// be destroyed without any errors. The default value is false.
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// ID of the port this interface connects to. Changing
	// this creates a new router interface.
	PortID *string `json:"portId,omitempty" tf:"port_id,omitempty"`

	// The region in which to obtain the V2 networking client.
	// A networking client is needed to create a router. If omitted, the
	// region argument of the provider is used. Changing this creates a new
	// router interface.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// ID of the router this interface belongs to. Changing
	// this creates a new router interface.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-openstack/apis/networking/v1alpha1.RouterV2
	RouterID *string `json:"routerId,omitempty" tf:"router_id,omitempty"`

	// Reference to a RouterV2 in networking to populate routerId.
	// +kubebuilder:validation:Optional
	RouterIDRef *v1.Reference `json:"routerIdRef,omitempty" tf:"-"`

	// Selector for a RouterV2 in networking to populate routerId.
	// +kubebuilder:validation:Optional
	RouterIDSelector *v1.Selector `json:"routerIdSelector,omitempty" tf:"-"`

	// ID of the subnet this interface connects to. Changing
	// this creates a new router interface.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-openstack/apis/networking/v1alpha1.SubnetV2
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a SubnetV2 in networking to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a SubnetV2 in networking to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`
}

type RouterInterfaceV2Observation struct {

	// A boolean indicating whether the routes from the
	// corresponding router ID should be deleted so that the router interface can
	// be destroyed without any errors. The default value is false.
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// ID of the port this interface connects to. Changing
	// this creates a new router interface.
	PortID *string `json:"portId,omitempty" tf:"port_id,omitempty"`

	// The region in which to obtain the V2 networking client.
	// A networking client is needed to create a router. If omitted, the
	// region argument of the provider is used. Changing this creates a new
	// router interface.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// ID of the router this interface belongs to. Changing
	// this creates a new router interface.
	RouterID *string `json:"routerId,omitempty" tf:"router_id,omitempty"`

	// ID of the subnet this interface connects to. Changing
	// this creates a new router interface.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`
}

type RouterInterfaceV2Parameters struct {

	// A boolean indicating whether the routes from the
	// corresponding router ID should be deleted so that the router interface can
	// be destroyed without any errors. The default value is false.
	// +kubebuilder:validation:Optional
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// ID of the port this interface connects to. Changing
	// this creates a new router interface.
	// +kubebuilder:validation:Optional
	PortID *string `json:"portId,omitempty" tf:"port_id,omitempty"`

	// The region in which to obtain the V2 networking client.
	// A networking client is needed to create a router. If omitted, the
	// region argument of the provider is used. Changing this creates a new
	// router interface.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// ID of the router this interface belongs to. Changing
	// this creates a new router interface.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-openstack/apis/networking/v1alpha1.RouterV2
	// +kubebuilder:validation:Optional
	RouterID *string `json:"routerId,omitempty" tf:"router_id,omitempty"`

	// Reference to a RouterV2 in networking to populate routerId.
	// +kubebuilder:validation:Optional
	RouterIDRef *v1.Reference `json:"routerIdRef,omitempty" tf:"-"`

	// Selector for a RouterV2 in networking to populate routerId.
	// +kubebuilder:validation:Optional
	RouterIDSelector *v1.Selector `json:"routerIdSelector,omitempty" tf:"-"`

	// ID of the subnet this interface connects to. Changing
	// this creates a new router interface.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-openstack/apis/networking/v1alpha1.SubnetV2
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a SubnetV2 in networking to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a SubnetV2 in networking to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`
}

// RouterInterfaceV2Spec defines the desired state of RouterInterfaceV2
type RouterInterfaceV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RouterInterfaceV2Parameters `json:"forProvider"`
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
	InitProvider RouterInterfaceV2InitParameters `json:"initProvider,omitempty"`
}

// RouterInterfaceV2Status defines the observed state of RouterInterfaceV2.
type RouterInterfaceV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RouterInterfaceV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// RouterInterfaceV2 is the Schema for the RouterInterfaceV2s API. Manages a V2 router interface resource within OpenStack.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,openstack}
type RouterInterfaceV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RouterInterfaceV2Spec   `json:"spec"`
	Status            RouterInterfaceV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RouterInterfaceV2List contains a list of RouterInterfaceV2s
type RouterInterfaceV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RouterInterfaceV2 `json:"items"`
}

// Repository type metadata.
var (
	RouterInterfaceV2_Kind             = "RouterInterfaceV2"
	RouterInterfaceV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RouterInterfaceV2_Kind}.String()
	RouterInterfaceV2_KindAPIVersion   = RouterInterfaceV2_Kind + "." + CRDGroupVersion.String()
	RouterInterfaceV2_GroupVersionKind = CRDGroupVersion.WithKind(RouterInterfaceV2_Kind)
)

func init() {
	SchemeBuilder.Register(&RouterInterfaceV2{}, &RouterInterfaceV2List{})
}
