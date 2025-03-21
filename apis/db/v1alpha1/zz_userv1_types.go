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

type UserV1InitParameters struct {

	// A list of database user should have access to.
	// +listType=set
	Databases []*string `json:"databases,omitempty" tf:"databases,omitempty"`

	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	// The ID for the database instance.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-openstack/apis/db/v1alpha1.InstanceV1
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Reference to a InstanceV1 in db to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDRef *v1.Reference `json:"instanceIdRef,omitempty" tf:"-"`

	// Selector for a InstanceV1 in db to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDSelector *v1.Selector `json:"instanceIdSelector,omitempty" tf:"-"`

	// A unique name for the resource.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// User's password.
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// Openstack region resource is created in.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type UserV1Observation struct {

	// A list of database user should have access to.
	// +listType=set
	Databases []*string `json:"databases,omitempty" tf:"databases,omitempty"`

	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID for the database instance.
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// A unique name for the resource.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Openstack region resource is created in.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type UserV1Parameters struct {

	// A list of database user should have access to.
	// +kubebuilder:validation:Optional
	// +listType=set
	Databases []*string `json:"databases,omitempty" tf:"databases,omitempty"`

	// +kubebuilder:validation:Optional
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	// The ID for the database instance.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-openstack/apis/db/v1alpha1.InstanceV1
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Reference to a InstanceV1 in db to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDRef *v1.Reference `json:"instanceIdRef,omitempty" tf:"-"`

	// Selector for a InstanceV1 in db to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDSelector *v1.Selector `json:"instanceIdSelector,omitempty" tf:"-"`

	// A unique name for the resource.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// User's password.
	// +kubebuilder:validation:Optional
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// Openstack region resource is created in.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

// UserV1Spec defines the desired state of UserV1
type UserV1Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     UserV1Parameters `json:"forProvider"`
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
	InitProvider UserV1InitParameters `json:"initProvider,omitempty"`
}

// UserV1Status defines the observed state of UserV1.
type UserV1Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        UserV1Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// UserV1 is the Schema for the UserV1s API. Manages a V1 database user resource within OpenStack.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,openstack}
type UserV1 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.passwordSecretRef)",message="spec.forProvider.passwordSecretRef is a required parameter"
	Spec   UserV1Spec   `json:"spec"`
	Status UserV1Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UserV1List contains a list of UserV1s
type UserV1List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UserV1 `json:"items"`
}

// Repository type metadata.
var (
	UserV1_Kind             = "UserV1"
	UserV1_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: UserV1_Kind}.String()
	UserV1_KindAPIVersion   = UserV1_Kind + "." + CRDGroupVersion.String()
	UserV1_GroupVersionKind = CRDGroupVersion.WithKind(UserV1_Kind)
)

func init() {
	SchemeBuilder.Register(&UserV1{}, &UserV1List{})
}
