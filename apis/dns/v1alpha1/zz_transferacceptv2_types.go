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

type TransferAcceptV2InitParameters struct {

	// Disable wait for zone to reach ACTIVE
	// status. The check is enabled by default. If this argument is true, zone
	// will be considered as created/updated if OpenStack accept returned success.
	DisableStatusCheck *bool `json:"disableStatusCheck,omitempty" tf:"disable_status_check,omitempty"`

	// The transfer key.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// The region in which to obtain the V2 Compute client.
	// Keypairs are associated with accounts, but a Compute client is needed to
	// create one. If omitted, the region argument of the provider is used.
	// Changing this creates a new DNS zone.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Map of additional options. Changing this creates a
	// new transfer accept.
	// +mapType=granular
	ValueSpecs map[string]*string `json:"valueSpecs,omitempty" tf:"value_specs,omitempty"`

	// The ID of the zone transfer request.
	ZoneTransferRequestID *string `json:"zoneTransferRequestId,omitempty" tf:"zone_transfer_request_id,omitempty"`
}

type TransferAcceptV2Observation struct {

	// Disable wait for zone to reach ACTIVE
	// status. The check is enabled by default. If this argument is true, zone
	// will be considered as created/updated if OpenStack accept returned success.
	DisableStatusCheck *bool `json:"disableStatusCheck,omitempty" tf:"disable_status_check,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The transfer key.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// The region in which to obtain the V2 Compute client.
	// Keypairs are associated with accounts, but a Compute client is needed to
	// create one. If omitted, the region argument of the provider is used.
	// Changing this creates a new DNS zone.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Map of additional options. Changing this creates a
	// new transfer accept.
	// +mapType=granular
	ValueSpecs map[string]*string `json:"valueSpecs,omitempty" tf:"value_specs,omitempty"`

	// The ID of the zone transfer request.
	ZoneTransferRequestID *string `json:"zoneTransferRequestId,omitempty" tf:"zone_transfer_request_id,omitempty"`
}

type TransferAcceptV2Parameters struct {

	// Disable wait for zone to reach ACTIVE
	// status. The check is enabled by default. If this argument is true, zone
	// will be considered as created/updated if OpenStack accept returned success.
	// +kubebuilder:validation:Optional
	DisableStatusCheck *bool `json:"disableStatusCheck,omitempty" tf:"disable_status_check,omitempty"`

	// The transfer key.
	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// The region in which to obtain the V2 Compute client.
	// Keypairs are associated with accounts, but a Compute client is needed to
	// create one. If omitted, the region argument of the provider is used.
	// Changing this creates a new DNS zone.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Map of additional options. Changing this creates a
	// new transfer accept.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	ValueSpecs map[string]*string `json:"valueSpecs,omitempty" tf:"value_specs,omitempty"`

	// The ID of the zone transfer request.
	// +kubebuilder:validation:Optional
	ZoneTransferRequestID *string `json:"zoneTransferRequestId,omitempty" tf:"zone_transfer_request_id,omitempty"`
}

// TransferAcceptV2Spec defines the desired state of TransferAcceptV2
type TransferAcceptV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TransferAcceptV2Parameters `json:"forProvider"`
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
	InitProvider TransferAcceptV2InitParameters `json:"initProvider,omitempty"`
}

// TransferAcceptV2Status defines the observed state of TransferAcceptV2.
type TransferAcceptV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TransferAcceptV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// TransferAcceptV2 is the Schema for the TransferAcceptV2s API. Manages a DNS zone Transfer accept in the OpenStack DNS Service
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,openstack}
type TransferAcceptV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.key) || (has(self.initProvider) && has(self.initProvider.key))",message="spec.forProvider.key is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.zoneTransferRequestId) || (has(self.initProvider) && has(self.initProvider.zoneTransferRequestId))",message="spec.forProvider.zoneTransferRequestId is a required parameter"
	Spec   TransferAcceptV2Spec   `json:"spec"`
	Status TransferAcceptV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TransferAcceptV2List contains a list of TransferAcceptV2s
type TransferAcceptV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TransferAcceptV2 `json:"items"`
}

// Repository type metadata.
var (
	TransferAcceptV2_Kind             = "TransferAcceptV2"
	TransferAcceptV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TransferAcceptV2_Kind}.String()
	TransferAcceptV2_KindAPIVersion   = TransferAcceptV2_Kind + "." + CRDGroupVersion.String()
	TransferAcceptV2_GroupVersionKind = CRDGroupVersion.WithKind(TransferAcceptV2_Kind)
)

func init() {
	SchemeBuilder.Register(&TransferAcceptV2{}, &TransferAcceptV2List{})
}
