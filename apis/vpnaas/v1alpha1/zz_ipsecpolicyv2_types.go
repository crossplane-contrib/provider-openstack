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

type IpsecPolicyV2InitParameters struct {

	// The authentication hash algorithm. Valid values are sha1, sha256, sha384, sha512.
	// Default is sha1. Changing this updates the algorithm of the existing policy.
	AuthAlgorithm *string `json:"authAlgorithm,omitempty" tf:"auth_algorithm,omitempty"`

	// The human-readable description for the policy.
	// Changing this updates the description of the existing policy.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The encapsulation mode. Valid values are tunnel and transport. Default is tunnel.
	// Changing this updates the existing policy.
	EncapsulationMode *string `json:"encapsulationMode,omitempty" tf:"encapsulation_mode,omitempty"`

	// The encryption algorithm. Valid values are 3des, aes-128, aes-192 and so on.
	// The default value is aes-128. Changing this updates the existing policy.
	EncryptionAlgorithm *string `json:"encryptionAlgorithm,omitempty" tf:"encryption_algorithm,omitempty"`

	// The lifetime of the security association. Consists of Unit and Value.
	Lifetime []IpsecPolicyV2LifetimeInitParameters `json:"lifetime,omitempty" tf:"lifetime,omitempty"`

	// The name of the policy. Changing this updates the name of
	// the existing policy.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The perfect forward secrecy mode. Valid values are group2, group5 and group14. Default
	// is group5. Changing this updates the existing policy.
	Pfs *string `json:"pfs,omitempty" tf:"pfs,omitempty"`

	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create an IPSec policy. If omitted, the
	// region argument of the provider is used. Changing this creates a new
	// policy.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The owner of the policy. Required if admin wants to
	// create a policy for another project. Changing this creates a new policy.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// The transform protocol. Valid values are esp, ah and ah-esp.
	// Changing this updates the existing policy. Default is ESP.
	TransformProtocol *string `json:"transformProtocol,omitempty" tf:"transform_protocol,omitempty"`

	// Map of additional options.
	// +mapType=granular
	ValueSpecs map[string]*string `json:"valueSpecs,omitempty" tf:"value_specs,omitempty"`
}

type IpsecPolicyV2LifetimeInitParameters struct {
	Units *string `json:"units,omitempty" tf:"units,omitempty"`

	// The value for the lifetime of the security association. Must be a positive integer.
	// Default is 3600.
	Value *float64 `json:"value,omitempty" tf:"value,omitempty"`
}

type IpsecPolicyV2LifetimeObservation struct {
	Units *string `json:"units,omitempty" tf:"units,omitempty"`

	// The value for the lifetime of the security association. Must be a positive integer.
	// Default is 3600.
	Value *float64 `json:"value,omitempty" tf:"value,omitempty"`
}

type IpsecPolicyV2LifetimeParameters struct {

	// +kubebuilder:validation:Optional
	Units *string `json:"units,omitempty" tf:"units,omitempty"`

	// The value for the lifetime of the security association. Must be a positive integer.
	// Default is 3600.
	// +kubebuilder:validation:Optional
	Value *float64 `json:"value,omitempty" tf:"value,omitempty"`
}

type IpsecPolicyV2Observation struct {

	// The authentication hash algorithm. Valid values are sha1, sha256, sha384, sha512.
	// Default is sha1. Changing this updates the algorithm of the existing policy.
	AuthAlgorithm *string `json:"authAlgorithm,omitempty" tf:"auth_algorithm,omitempty"`

	// The human-readable description for the policy.
	// Changing this updates the description of the existing policy.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The encapsulation mode. Valid values are tunnel and transport. Default is tunnel.
	// Changing this updates the existing policy.
	EncapsulationMode *string `json:"encapsulationMode,omitempty" tf:"encapsulation_mode,omitempty"`

	// The encryption algorithm. Valid values are 3des, aes-128, aes-192 and so on.
	// The default value is aes-128. Changing this updates the existing policy.
	EncryptionAlgorithm *string `json:"encryptionAlgorithm,omitempty" tf:"encryption_algorithm,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The lifetime of the security association. Consists of Unit and Value.
	Lifetime []IpsecPolicyV2LifetimeObservation `json:"lifetime,omitempty" tf:"lifetime,omitempty"`

	// The name of the policy. Changing this updates the name of
	// the existing policy.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The perfect forward secrecy mode. Valid values are group2, group5 and group14. Default
	// is group5. Changing this updates the existing policy.
	Pfs *string `json:"pfs,omitempty" tf:"pfs,omitempty"`

	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create an IPSec policy. If omitted, the
	// region argument of the provider is used. Changing this creates a new
	// policy.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The owner of the policy. Required if admin wants to
	// create a policy for another project. Changing this creates a new policy.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// The transform protocol. Valid values are esp, ah and ah-esp.
	// Changing this updates the existing policy. Default is ESP.
	TransformProtocol *string `json:"transformProtocol,omitempty" tf:"transform_protocol,omitempty"`

	// Map of additional options.
	// +mapType=granular
	ValueSpecs map[string]*string `json:"valueSpecs,omitempty" tf:"value_specs,omitempty"`
}

type IpsecPolicyV2Parameters struct {

	// The authentication hash algorithm. Valid values are sha1, sha256, sha384, sha512.
	// Default is sha1. Changing this updates the algorithm of the existing policy.
	// +kubebuilder:validation:Optional
	AuthAlgorithm *string `json:"authAlgorithm,omitempty" tf:"auth_algorithm,omitempty"`

	// The human-readable description for the policy.
	// Changing this updates the description of the existing policy.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The encapsulation mode. Valid values are tunnel and transport. Default is tunnel.
	// Changing this updates the existing policy.
	// +kubebuilder:validation:Optional
	EncapsulationMode *string `json:"encapsulationMode,omitempty" tf:"encapsulation_mode,omitempty"`

	// The encryption algorithm. Valid values are 3des, aes-128, aes-192 and so on.
	// The default value is aes-128. Changing this updates the existing policy.
	// +kubebuilder:validation:Optional
	EncryptionAlgorithm *string `json:"encryptionAlgorithm,omitempty" tf:"encryption_algorithm,omitempty"`

	// The lifetime of the security association. Consists of Unit and Value.
	// +kubebuilder:validation:Optional
	Lifetime []IpsecPolicyV2LifetimeParameters `json:"lifetime,omitempty" tf:"lifetime,omitempty"`

	// The name of the policy. Changing this updates the name of
	// the existing policy.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The perfect forward secrecy mode. Valid values are group2, group5 and group14. Default
	// is group5. Changing this updates the existing policy.
	// +kubebuilder:validation:Optional
	Pfs *string `json:"pfs,omitempty" tf:"pfs,omitempty"`

	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create an IPSec policy. If omitted, the
	// region argument of the provider is used. Changing this creates a new
	// policy.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The owner of the policy. Required if admin wants to
	// create a policy for another project. Changing this creates a new policy.
	// +kubebuilder:validation:Optional
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// The transform protocol. Valid values are esp, ah and ah-esp.
	// Changing this updates the existing policy. Default is ESP.
	// +kubebuilder:validation:Optional
	TransformProtocol *string `json:"transformProtocol,omitempty" tf:"transform_protocol,omitempty"`

	// Map of additional options.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	ValueSpecs map[string]*string `json:"valueSpecs,omitempty" tf:"value_specs,omitempty"`
}

// IpsecPolicyV2Spec defines the desired state of IpsecPolicyV2
type IpsecPolicyV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IpsecPolicyV2Parameters `json:"forProvider"`
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
	InitProvider IpsecPolicyV2InitParameters `json:"initProvider,omitempty"`
}

// IpsecPolicyV2Status defines the observed state of IpsecPolicyV2.
type IpsecPolicyV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IpsecPolicyV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// IpsecPolicyV2 is the Schema for the IpsecPolicyV2s API. Manages a V2 Neutron IPSec policy resource within OpenStack.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,openstack}
type IpsecPolicyV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IpsecPolicyV2Spec   `json:"spec"`
	Status            IpsecPolicyV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IpsecPolicyV2List contains a list of IpsecPolicyV2s
type IpsecPolicyV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IpsecPolicyV2 `json:"items"`
}

// Repository type metadata.
var (
	IpsecPolicyV2_Kind             = "IpsecPolicyV2"
	IpsecPolicyV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IpsecPolicyV2_Kind}.String()
	IpsecPolicyV2_KindAPIVersion   = IpsecPolicyV2_Kind + "." + CRDGroupVersion.String()
	IpsecPolicyV2_GroupVersionKind = CRDGroupVersion.WithKind(IpsecPolicyV2_Kind)
)

func init() {
	SchemeBuilder.Register(&IpsecPolicyV2{}, &IpsecPolicyV2List{})
}
