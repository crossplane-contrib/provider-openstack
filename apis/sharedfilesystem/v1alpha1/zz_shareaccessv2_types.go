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

type ShareAccessV2InitParameters struct {

	// The access level to the share. Can either be rw or ro.
	AccessLevel *string `json:"accessLevel,omitempty" tf:"access_level,omitempty"`

	// The value that defines the access. Can either be an IP
	// address or a username verified by configured Security Service of the Share Network.
	AccessTo *string `json:"accessTo,omitempty" tf:"access_to,omitempty"`

	// The access rule type. Can either be an ip, user,
	// cert, or cephx. cephx support requires an OpenStack environment that supports
	// Shared Filesystem microversion 2.13 (Mitaka) or later.
	AccessType *string `json:"accessType,omitempty" tf:"access_type,omitempty"`

	// The region in which to obtain the V2 Shared File System client.
	// A Shared File System client is needed to create a share access. Changing this
	// creates a new share access.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The UUID of the share to which you are granted access.
	ShareID *string `json:"shareId,omitempty" tf:"share_id,omitempty"`
}

type ShareAccessV2Observation struct {

	// The access level to the share. Can either be rw or ro.
	AccessLevel *string `json:"accessLevel,omitempty" tf:"access_level,omitempty"`

	// The value that defines the access. Can either be an IP
	// address or a username verified by configured Security Service of the Share Network.
	AccessTo *string `json:"accessTo,omitempty" tf:"access_to,omitempty"`

	// The access rule type. Can either be an ip, user,
	// cert, or cephx. cephx support requires an OpenStack environment that supports
	// Shared Filesystem microversion 2.13 (Mitaka) or later.
	AccessType *string `json:"accessType,omitempty" tf:"access_type,omitempty"`

	// The unique ID for the Share Access.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The region in which to obtain the V2 Shared File System client.
	// A Shared File System client is needed to create a share access. Changing this
	// creates a new share access.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The UUID of the share to which you are granted access.
	ShareID *string `json:"shareId,omitempty" tf:"share_id,omitempty"`

	// The share access state.
	State *string `json:"state,omitempty" tf:"state,omitempty"`
}

type ShareAccessV2Parameters struct {

	// The access level to the share. Can either be rw or ro.
	// +kubebuilder:validation:Optional
	AccessLevel *string `json:"accessLevel,omitempty" tf:"access_level,omitempty"`

	// The value that defines the access. Can either be an IP
	// address or a username verified by configured Security Service of the Share Network.
	// +kubebuilder:validation:Optional
	AccessTo *string `json:"accessTo,omitempty" tf:"access_to,omitempty"`

	// The access rule type. Can either be an ip, user,
	// cert, or cephx. cephx support requires an OpenStack environment that supports
	// Shared Filesystem microversion 2.13 (Mitaka) or later.
	// +kubebuilder:validation:Optional
	AccessType *string `json:"accessType,omitempty" tf:"access_type,omitempty"`

	// The region in which to obtain the V2 Shared File System client.
	// A Shared File System client is needed to create a share access. Changing this
	// creates a new share access.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The UUID of the share to which you are granted access.
	// +kubebuilder:validation:Optional
	ShareID *string `json:"shareId,omitempty" tf:"share_id,omitempty"`
}

// ShareAccessV2Spec defines the desired state of ShareAccessV2
type ShareAccessV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ShareAccessV2Parameters `json:"forProvider"`
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
	InitProvider ShareAccessV2InitParameters `json:"initProvider,omitempty"`
}

// ShareAccessV2Status defines the observed state of ShareAccessV2.
type ShareAccessV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ShareAccessV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ShareAccessV2 is the Schema for the ShareAccessV2s API. Configure a Shared File System share access list.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,openstack}
type ShareAccessV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.accessLevel) || (has(self.initProvider) && has(self.initProvider.accessLevel))",message="spec.forProvider.accessLevel is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.accessTo) || (has(self.initProvider) && has(self.initProvider.accessTo))",message="spec.forProvider.accessTo is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.accessType) || (has(self.initProvider) && has(self.initProvider.accessType))",message="spec.forProvider.accessType is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.shareId) || (has(self.initProvider) && has(self.initProvider.shareId))",message="spec.forProvider.shareId is a required parameter"
	Spec   ShareAccessV2Spec   `json:"spec"`
	Status ShareAccessV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ShareAccessV2List contains a list of ShareAccessV2s
type ShareAccessV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ShareAccessV2 `json:"items"`
}

// Repository type metadata.
var (
	ShareAccessV2_Kind             = "ShareAccessV2"
	ShareAccessV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ShareAccessV2_Kind}.String()
	ShareAccessV2_KindAPIVersion   = ShareAccessV2_Kind + "." + CRDGroupVersion.String()
	ShareAccessV2_GroupVersionKind = CRDGroupVersion.WithKind(ShareAccessV2_Kind)
)

func init() {
	SchemeBuilder.Register(&ShareAccessV2{}, &ShareAccessV2List{})
}
