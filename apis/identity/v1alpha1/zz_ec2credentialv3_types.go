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

type EC2CredentialV3InitParameters struct {

	// The ID of the project the EC2 credential is created
	// for and that authentication requests using this EC2 credential will
	// be scoped to. Only administrative users can specify a project ID different
	// from the current auth scope.
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// The region in which to obtain the V3 Keystone client.
	// If omitted, the region argument of the provider is used. Changing this
	// creates a new EC2 credential.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The ID of the user the EC2 credential is created for.
	// Only administrative users can specify a user ID different from the current
	// auth scope.
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`
}

type EC2CredentialV3Observation struct {

	// contains an EC2 credential access UUID
	Access *string `json:"access,omitempty" tf:"access,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the project the EC2 credential is created
	// for and that authentication requests using this EC2 credential will
	// be scoped to. Only administrative users can specify a project ID different
	// from the current auth scope.
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// The region in which to obtain the V3 Keystone client.
	// If omitted, the region argument of the provider is used. Changing this
	// creates a new EC2 credential.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// contains an EC2 credential trust ID scope
	TrustID *string `json:"trustId,omitempty" tf:"trust_id,omitempty"`

	// The ID of the user the EC2 credential is created for.
	// Only administrative users can specify a user ID different from the current
	// auth scope.
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`
}

type EC2CredentialV3Parameters struct {

	// The ID of the project the EC2 credential is created
	// for and that authentication requests using this EC2 credential will
	// be scoped to. Only administrative users can specify a project ID different
	// from the current auth scope.
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// The region in which to obtain the V3 Keystone client.
	// If omitted, the region argument of the provider is used. Changing this
	// creates a new EC2 credential.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The ID of the user the EC2 credential is created for.
	// Only administrative users can specify a user ID different from the current
	// auth scope.
	// +kubebuilder:validation:Optional
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`
}

// EC2CredentialV3Spec defines the desired state of EC2CredentialV3
type EC2CredentialV3Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EC2CredentialV3Parameters `json:"forProvider"`
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
	InitProvider EC2CredentialV3InitParameters `json:"initProvider,omitempty"`
}

// EC2CredentialV3Status defines the observed state of EC2CredentialV3.
type EC2CredentialV3Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EC2CredentialV3Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EC2CredentialV3 is the Schema for the EC2CredentialV3s API. Manages a V3 EC2 Credential resource within OpenStack Keystone.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,openstack}
type EC2CredentialV3 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EC2CredentialV3Spec   `json:"spec"`
	Status            EC2CredentialV3Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EC2CredentialV3List contains a list of EC2CredentialV3s
type EC2CredentialV3List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EC2CredentialV3 `json:"items"`
}

// Repository type metadata.
var (
	EC2CredentialV3_Kind             = "EC2CredentialV3"
	EC2CredentialV3_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EC2CredentialV3_Kind}.String()
	EC2CredentialV3_KindAPIVersion   = EC2CredentialV3_Kind + "." + CRDGroupVersion.String()
	EC2CredentialV3_GroupVersionKind = CRDGroupVersion.WithKind(EC2CredentialV3_Kind)
)

func init() {
	SchemeBuilder.Register(&EC2CredentialV3{}, &EC2CredentialV3List{})
}