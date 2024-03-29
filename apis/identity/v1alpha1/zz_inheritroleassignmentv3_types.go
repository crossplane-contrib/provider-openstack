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

type InheritRoleAssignmentV3InitParameters struct {

	// The domain to assign the role in.
	DomainID *string `json:"domainId,omitempty" tf:"domain_id,omitempty"`

	// The group to assign the role to.
	GroupID *string `json:"groupId,omitempty" tf:"group_id,omitempty"`

	// The project to assign the role in.
	// The project should be able to containt child projects.
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The role to assign.
	RoleID *string `json:"roleId,omitempty" tf:"role_id,omitempty"`

	// The user to assign the role to.
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`
}

type InheritRoleAssignmentV3Observation struct {

	// The domain to assign the role in.
	DomainID *string `json:"domainId,omitempty" tf:"domain_id,omitempty"`

	// The group to assign the role to.
	GroupID *string `json:"groupId,omitempty" tf:"group_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The project to assign the role in.
	// The project should be able to containt child projects.
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The role to assign.
	RoleID *string `json:"roleId,omitempty" tf:"role_id,omitempty"`

	// The user to assign the role to.
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`
}

type InheritRoleAssignmentV3Parameters struct {

	// The domain to assign the role in.
	// +kubebuilder:validation:Optional
	DomainID *string `json:"domainId,omitempty" tf:"domain_id,omitempty"`

	// The group to assign the role to.
	// +kubebuilder:validation:Optional
	GroupID *string `json:"groupId,omitempty" tf:"group_id,omitempty"`

	// The project to assign the role in.
	// The project should be able to containt child projects.
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The role to assign.
	// +kubebuilder:validation:Optional
	RoleID *string `json:"roleId,omitempty" tf:"role_id,omitempty"`

	// The user to assign the role to.
	// +kubebuilder:validation:Optional
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`
}

// InheritRoleAssignmentV3Spec defines the desired state of InheritRoleAssignmentV3
type InheritRoleAssignmentV3Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InheritRoleAssignmentV3Parameters `json:"forProvider"`
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
	InitProvider InheritRoleAssignmentV3InitParameters `json:"initProvider,omitempty"`
}

// InheritRoleAssignmentV3Status defines the observed state of InheritRoleAssignmentV3.
type InheritRoleAssignmentV3Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InheritRoleAssignmentV3Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// InheritRoleAssignmentV3 is the Schema for the InheritRoleAssignmentV3s API. Manages a V3 Inherit Role assignment within OpenStack Keystone.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,openstack}
type InheritRoleAssignmentV3 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.roleId) || (has(self.initProvider) && has(self.initProvider.roleId))",message="spec.forProvider.roleId is a required parameter"
	Spec   InheritRoleAssignmentV3Spec   `json:"spec"`
	Status InheritRoleAssignmentV3Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InheritRoleAssignmentV3List contains a list of InheritRoleAssignmentV3s
type InheritRoleAssignmentV3List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InheritRoleAssignmentV3 `json:"items"`
}

// Repository type metadata.
var (
	InheritRoleAssignmentV3_Kind             = "InheritRoleAssignmentV3"
	InheritRoleAssignmentV3_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: InheritRoleAssignmentV3_Kind}.String()
	InheritRoleAssignmentV3_KindAPIVersion   = InheritRoleAssignmentV3_Kind + "." + CRDGroupVersion.String()
	InheritRoleAssignmentV3_GroupVersionKind = CRDGroupVersion.WithKind(InheritRoleAssignmentV3_Kind)
)

func init() {
	SchemeBuilder.Register(&InheritRoleAssignmentV3{}, &InheritRoleAssignmentV3List{})
}
