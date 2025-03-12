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

type ACLReadInitParameters struct {

	// Whether the secret is accessible project wide.
	// Defaults to true.
	ProjectAccess *bool `json:"projectAccess,omitempty" tf:"project_access,omitempty"`

	// The list of user IDs, which are allowed to access the
	// secret, when project_access is set to false.
	// +listType=set
	Users []*string `json:"users,omitempty" tf:"users,omitempty"`
}

type ACLReadObservation struct {

	// The date the secret was created.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// Whether the secret is accessible project wide.
	// Defaults to true.
	ProjectAccess *bool `json:"projectAccess,omitempty" tf:"project_access,omitempty"`

	// The date the secret was last updated.
	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`

	// The list of user IDs, which are allowed to access the
	// secret, when project_access is set to false.
	// +listType=set
	Users []*string `json:"users,omitempty" tf:"users,omitempty"`
}

type ACLReadParameters struct {

	// Whether the secret is accessible project wide.
	// Defaults to true.
	// +kubebuilder:validation:Optional
	ProjectAccess *bool `json:"projectAccess,omitempty" tf:"project_access,omitempty"`

	// The list of user IDs, which are allowed to access the
	// secret, when project_access is set to false.
	// +kubebuilder:validation:Optional
	// +listType=set
	Users []*string `json:"users,omitempty" tf:"users,omitempty"`
}

type SecretV1ACLInitParameters struct {
	Read *ACLReadInitParameters `json:"read,omitempty" tf:"read,omitempty"`
}

type SecretV1ACLObservation struct {
	Read *ACLReadObservation `json:"read,omitempty" tf:"read,omitempty"`
}

type SecretV1ACLParameters struct {

	// +kubebuilder:validation:Optional
	Read *ACLReadParameters `json:"read,omitempty" tf:"read,omitempty"`
}

type SecretV1InitParameters struct {

	// Allows to control an access to a secret. Currently only the
	// read operation is supported. If not specified, the secret is accessible
	// project wide.
	ACL *SecretV1ACLInitParameters `json:"acl,omitempty" tf:"acl,omitempty"`

	// Metadata provided by a user or system for informational purposes.
	Algorithm *string `json:"algorithm,omitempty" tf:"algorithm,omitempty"`

	// Metadata provided by a user or system for informational purposes.
	BitLength *float64 `json:"bitLength,omitempty" tf:"bit_length,omitempty"`

	// The expiration time of the secret in the RFC3339 timestamp format (e.g. 2019-03-09T12:58:49Z). If omitted, a secret will never expire. Changing this creates a new secret.
	Expiration *string `json:"expiration,omitempty" tf:"expiration,omitempty"`

	// Additional Metadata for the secret.
	// +mapType=granular
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// Metadata provided by a user or system for informational purposes.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// Human-readable name for the Secret. Does not have
	// to be unique.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The encoding used for the payload to be able to include it in the JSON request. Must be either base64 or binary.
	PayloadContentEncoding *string `json:"payloadContentEncoding,omitempty" tf:"payload_content_encoding,omitempty"`

	// The media type for the content of the payload. Must be one of text/plain, text/plain;charset=utf-8, text/plain; charset=utf-8, application/octet-stream, application/pkcs8.
	PayloadContentType *string `json:"payloadContentType,omitempty" tf:"payload_content_type,omitempty"`

	// The secret's data to be stored. payload_content_type must also be supplied if payload is included.
	PayloadSecretRef *v1.SecretKeySelector `json:"payloadSecretRef,omitempty" tf:"-"`

	// The region in which to obtain the V1 KeyManager client.
	// A KeyManager client is needed to create a secret. If omitted, the
	// region argument of the provider is used. Changing this creates a new
	// V1 secret.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Used to indicate the type of secret being stored. For more information see Secret types.
	SecretType *string `json:"secretType,omitempty" tf:"secret_type,omitempty"`
}

type SecretV1Observation struct {

	// Allows to control an access to a secret. Currently only the
	// read operation is supported. If not specified, the secret is accessible
	// project wide.
	ACL *SecretV1ACLObservation `json:"acl,omitempty" tf:"acl,omitempty"`

	// Metadata provided by a user or system for informational purposes.
	Algorithm *string `json:"algorithm,omitempty" tf:"algorithm,omitempty"`

	// The map of metadata, assigned on the secret, which has been
	// explicitly and implicitly added.
	// +mapType=granular
	AllMetadata map[string]*string `json:"allMetadata,omitempty" tf:"all_metadata,omitempty"`

	// Metadata provided by a user or system for informational purposes.
	BitLength *float64 `json:"bitLength,omitempty" tf:"bit_length,omitempty"`

	// The map of the content types, assigned on the secret.
	// +mapType=granular
	ContentTypes map[string]*string `json:"contentTypes,omitempty" tf:"content_types,omitempty"`

	// The date the secret was created.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// The creator of the secret.
	CreatorID *string `json:"creatorId,omitempty" tf:"creator_id,omitempty"`

	// The expiration time of the secret in the RFC3339 timestamp format (e.g. 2019-03-09T12:58:49Z). If omitted, a secret will never expire. Changing this creates a new secret.
	Expiration *string `json:"expiration,omitempty" tf:"expiration,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Additional Metadata for the secret.
	// +mapType=granular
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// Metadata provided by a user or system for informational purposes.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// Human-readable name for the Secret. Does not have
	// to be unique.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The encoding used for the payload to be able to include it in the JSON request. Must be either base64 or binary.
	PayloadContentEncoding *string `json:"payloadContentEncoding,omitempty" tf:"payload_content_encoding,omitempty"`

	// The media type for the content of the payload. Must be one of text/plain, text/plain;charset=utf-8, text/plain; charset=utf-8, application/octet-stream, application/pkcs8.
	PayloadContentType *string `json:"payloadContentType,omitempty" tf:"payload_content_type,omitempty"`

	// The region in which to obtain the V1 KeyManager client.
	// A KeyManager client is needed to create a secret. If omitted, the
	// region argument of the provider is used. Changing this creates a new
	// V1 secret.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The secret reference / where to find the secret.
	SecretRef *string `json:"secretRef,omitempty" tf:"secret_ref,omitempty"`

	// Used to indicate the type of secret being stored. For more information see Secret types.
	SecretType *string `json:"secretType,omitempty" tf:"secret_type,omitempty"`

	// The status of the secret.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// The date the secret was last updated.
	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`
}

type SecretV1Parameters struct {

	// Allows to control an access to a secret. Currently only the
	// read operation is supported. If not specified, the secret is accessible
	// project wide.
	// +kubebuilder:validation:Optional
	ACL *SecretV1ACLParameters `json:"acl,omitempty" tf:"acl,omitempty"`

	// Metadata provided by a user or system for informational purposes.
	// +kubebuilder:validation:Optional
	Algorithm *string `json:"algorithm,omitempty" tf:"algorithm,omitempty"`

	// Metadata provided by a user or system for informational purposes.
	// +kubebuilder:validation:Optional
	BitLength *float64 `json:"bitLength,omitempty" tf:"bit_length,omitempty"`

	// The expiration time of the secret in the RFC3339 timestamp format (e.g. 2019-03-09T12:58:49Z). If omitted, a secret will never expire. Changing this creates a new secret.
	// +kubebuilder:validation:Optional
	Expiration *string `json:"expiration,omitempty" tf:"expiration,omitempty"`

	// Additional Metadata for the secret.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// Metadata provided by a user or system for informational purposes.
	// +kubebuilder:validation:Optional
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// Human-readable name for the Secret. Does not have
	// to be unique.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The encoding used for the payload to be able to include it in the JSON request. Must be either base64 or binary.
	// +kubebuilder:validation:Optional
	PayloadContentEncoding *string `json:"payloadContentEncoding,omitempty" tf:"payload_content_encoding,omitempty"`

	// The media type for the content of the payload. Must be one of text/plain, text/plain;charset=utf-8, text/plain; charset=utf-8, application/octet-stream, application/pkcs8.
	// +kubebuilder:validation:Optional
	PayloadContentType *string `json:"payloadContentType,omitempty" tf:"payload_content_type,omitempty"`

	// The secret's data to be stored. payload_content_type must also be supplied if payload is included.
	// +kubebuilder:validation:Optional
	PayloadSecretRef *v1.SecretKeySelector `json:"payloadSecretRef,omitempty" tf:"-"`

	// The region in which to obtain the V1 KeyManager client.
	// A KeyManager client is needed to create a secret. If omitted, the
	// region argument of the provider is used. Changing this creates a new
	// V1 secret.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Used to indicate the type of secret being stored. For more information see Secret types.
	// +kubebuilder:validation:Optional
	SecretType *string `json:"secretType,omitempty" tf:"secret_type,omitempty"`
}

// SecretV1Spec defines the desired state of SecretV1
type SecretV1Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecretV1Parameters `json:"forProvider"`
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
	InitProvider SecretV1InitParameters `json:"initProvider,omitempty"`
}

// SecretV1Status defines the observed state of SecretV1.
type SecretV1Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecretV1Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SecretV1 is the Schema for the SecretV1s API. Manages a V1 Barbican secret resource within OpenStack.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,openstack}
type SecretV1 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SecretV1Spec   `json:"spec"`
	Status            SecretV1Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecretV1List contains a list of SecretV1s
type SecretV1List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecretV1 `json:"items"`
}

// Repository type metadata.
var (
	SecretV1_Kind             = "SecretV1"
	SecretV1_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SecretV1_Kind}.String()
	SecretV1_KindAPIVersion   = SecretV1_Kind + "." + CRDGroupVersion.String()
	SecretV1_GroupVersionKind = CRDGroupVersion.WithKind(SecretV1_Kind)
)

func init() {
	SchemeBuilder.Register(&SecretV1{}, &SecretV1List{})
}
