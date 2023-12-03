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

type QosPolicyV2InitParameters struct {

	// The human-readable description for the QoS policy.
	// Changing this updates the description of the existing QoS policy.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Indicates whether the QoS policy is default
	// QoS policy or not. Changing this updates the default status of the existing
	// QoS policy.
	IsDefault *bool `json:"isDefault,omitempty" tf:"is_default,omitempty"`

	// The name of the QoS policy. Changing this updates the name of
	// the existing QoS policy.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The owner of the QoS policy. Required if admin wants to
	// create a QoS policy for another project. Changing this creates a new QoS policy.
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a Neutron Qos policy. If omitted, the
	// region argument of the provider is used. Changing this creates a new
	// QoS policy.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Indicates whether this QoS policy is shared across
	// all projects. Changing this updates the shared status of the existing
	// QoS policy.
	Shared *bool `json:"shared,omitempty" tf:"shared,omitempty"`

	// A set of string tags for the QoS policy.
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Map of additional options.
	ValueSpecs map[string]*string `json:"valueSpecs,omitempty" tf:"value_specs,omitempty"`
}

type QosPolicyV2Observation struct {

	// The collection of tags assigned on the QoS policy, which have been
	// explicitly and implicitly added.
	AllTags []*string `json:"allTags,omitempty" tf:"all_tags,omitempty"`

	// The time at which QoS policy was created.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// The human-readable description for the QoS policy.
	// Changing this updates the description of the existing QoS policy.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Indicates whether the QoS policy is default
	// QoS policy or not. Changing this updates the default status of the existing
	// QoS policy.
	IsDefault *bool `json:"isDefault,omitempty" tf:"is_default,omitempty"`

	// The name of the QoS policy. Changing this updates the name of
	// the existing QoS policy.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The owner of the QoS policy. Required if admin wants to
	// create a QoS policy for another project. Changing this creates a new QoS policy.
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a Neutron Qos policy. If omitted, the
	// region argument of the provider is used. Changing this creates a new
	// QoS policy.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The revision number of the QoS policy.
	RevisionNumber *float64 `json:"revisionNumber,omitempty" tf:"revision_number,omitempty"`

	// Indicates whether this QoS policy is shared across
	// all projects. Changing this updates the shared status of the existing
	// QoS policy.
	Shared *bool `json:"shared,omitempty" tf:"shared,omitempty"`

	// A set of string tags for the QoS policy.
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The time at which QoS policy was created.
	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`

	// Map of additional options.
	ValueSpecs map[string]*string `json:"valueSpecs,omitempty" tf:"value_specs,omitempty"`
}

type QosPolicyV2Parameters struct {

	// The human-readable description for the QoS policy.
	// Changing this updates the description of the existing QoS policy.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Indicates whether the QoS policy is default
	// QoS policy or not. Changing this updates the default status of the existing
	// QoS policy.
	// +kubebuilder:validation:Optional
	IsDefault *bool `json:"isDefault,omitempty" tf:"is_default,omitempty"`

	// The name of the QoS policy. Changing this updates the name of
	// the existing QoS policy.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The owner of the QoS policy. Required if admin wants to
	// create a QoS policy for another project. Changing this creates a new QoS policy.
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a Neutron Qos policy. If omitted, the
	// region argument of the provider is used. Changing this creates a new
	// QoS policy.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Indicates whether this QoS policy is shared across
	// all projects. Changing this updates the shared status of the existing
	// QoS policy.
	// +kubebuilder:validation:Optional
	Shared *bool `json:"shared,omitempty" tf:"shared,omitempty"`

	// A set of string tags for the QoS policy.
	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Map of additional options.
	// +kubebuilder:validation:Optional
	ValueSpecs map[string]*string `json:"valueSpecs,omitempty" tf:"value_specs,omitempty"`
}

// QosPolicyV2Spec defines the desired state of QosPolicyV2
type QosPolicyV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     QosPolicyV2Parameters `json:"forProvider"`
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
	InitProvider QosPolicyV2InitParameters `json:"initProvider,omitempty"`
}

// QosPolicyV2Status defines the observed state of QosPolicyV2.
type QosPolicyV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        QosPolicyV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// QosPolicyV2 is the Schema for the QosPolicyV2s API. Manages a V2 Neutron QoS policy resource within OpenStack.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,openstack}
type QosPolicyV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              QosPolicyV2Spec   `json:"spec"`
	Status            QosPolicyV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// QosPolicyV2List contains a list of QosPolicyV2s
type QosPolicyV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []QosPolicyV2 `json:"items"`
}

// Repository type metadata.
var (
	QosPolicyV2_Kind             = "QosPolicyV2"
	QosPolicyV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: QosPolicyV2_Kind}.String()
	QosPolicyV2_KindAPIVersion   = QosPolicyV2_Kind + "." + CRDGroupVersion.String()
	QosPolicyV2_GroupVersionKind = CRDGroupVersion.WithKind(QosPolicyV2_Kind)
)

func init() {
	SchemeBuilder.Register(&QosPolicyV2{}, &QosPolicyV2List{})
}
