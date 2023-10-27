/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type RecordsetV2Observation struct {

	// A description of the  record set.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Disable wait for recordset to reach ACTIVE
	// status. This argumen is disabled by default. If it is set to true, the recordset
	// will be considered as created/updated/deleted if OpenStack request returned success.
	DisableStatusCheck *bool `json:"disableStatusCheck,omitempty" tf:"disable_status_check,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the record set. Note the . at the end of the name.
	// Changing this creates a new DNS  record set.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the project DNS zone is created
	// for, sets X-Auth-Sudo-Tenant-ID header (requires an assigned
	// user role in target project)
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// An array of DNS records.
	Records []*string `json:"records,omitempty" tf:"records,omitempty"`

	// The region in which to obtain the V2 DNS client.
	// If omitted, the region argument of the provider is used.
	// Changing this creates a new DNS  record set.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The time to live (TTL) of the record set.
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// The type of record set. Examples: "A", "MX".
	// Changing this creates a new DNS  record set.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Map of additional options. Changing this creates a
	// new record set.
	ValueSpecs map[string]*string `json:"valueSpecs,omitempty" tf:"value_specs,omitempty"`

	// The ID of the zone in which to create the record set.
	// Changing this creates a new DNS  record set.
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`
}

type RecordsetV2Parameters struct {

	// A description of the  record set.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Disable wait for recordset to reach ACTIVE
	// status. This argumen is disabled by default. If it is set to true, the recordset
	// will be considered as created/updated/deleted if OpenStack request returned success.
	// +kubebuilder:validation:Optional
	DisableStatusCheck *bool `json:"disableStatusCheck,omitempty" tf:"disable_status_check,omitempty"`

	// The name of the record set. Note the . at the end of the name.
	// Changing this creates a new DNS  record set.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the project DNS zone is created
	// for, sets X-Auth-Sudo-Tenant-ID header (requires an assigned
	// user role in target project)
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// An array of DNS records.
	// +kubebuilder:validation:Optional
	Records []*string `json:"records,omitempty" tf:"records,omitempty"`

	// The region in which to obtain the V2 DNS client.
	// If omitted, the region argument of the provider is used.
	// Changing this creates a new DNS  record set.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The time to live (TTL) of the record set.
	// +kubebuilder:validation:Optional
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// The type of record set. Examples: "A", "MX".
	// Changing this creates a new DNS  record set.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Map of additional options. Changing this creates a
	// new record set.
	// +kubebuilder:validation:Optional
	ValueSpecs map[string]*string `json:"valueSpecs,omitempty" tf:"value_specs,omitempty"`

	// The ID of the zone in which to create the record set.
	// Changing this creates a new DNS  record set.
	// +crossplane:generate:reference:type=ZoneV2
	// +kubebuilder:validation:Optional
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`

	// Reference to a ZoneV2 to populate zoneId.
	// +kubebuilder:validation:Optional
	ZoneIDRef *v1.Reference `json:"zoneIdRef,omitempty" tf:"-"`

	// Selector for a ZoneV2 to populate zoneId.
	// +kubebuilder:validation:Optional
	ZoneIDSelector *v1.Selector `json:"zoneIdSelector,omitempty" tf:"-"`
}

// RecordsetV2Spec defines the desired state of RecordsetV2
type RecordsetV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RecordsetV2Parameters `json:"forProvider"`
}

// RecordsetV2Status defines the observed state of RecordsetV2.
type RecordsetV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RecordsetV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RecordsetV2 is the Schema for the RecordsetV2s API. Manages a DNS record set in the OpenStack DNS Service
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,openstack}
type RecordsetV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.records)",message="records is a required parameter"
	Spec   RecordsetV2Spec   `json:"spec"`
	Status RecordsetV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RecordsetV2List contains a list of RecordsetV2s
type RecordsetV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RecordsetV2 `json:"items"`
}

// Repository type metadata.
var (
	RecordsetV2_Kind             = "RecordsetV2"
	RecordsetV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RecordsetV2_Kind}.String()
	RecordsetV2_KindAPIVersion   = RecordsetV2_Kind + "." + CRDGroupVersion.String()
	RecordsetV2_GroupVersionKind = CRDGroupVersion.WithKind(RecordsetV2_Kind)
)

func init() {
	SchemeBuilder.Register(&RecordsetV2{}, &RecordsetV2List{})
}
