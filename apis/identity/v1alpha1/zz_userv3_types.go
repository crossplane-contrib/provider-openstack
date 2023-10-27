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

type MultiFactorAuthRuleObservation struct {

	// A list of authentication plugins that the user must
	// authenticate with.
	Rule []*string `json:"rule,omitempty" tf:"rule,omitempty"`
}

type MultiFactorAuthRuleParameters struct {

	// A list of authentication plugins that the user must
	// authenticate with.
	// +kubebuilder:validation:Required
	Rule []*string `json:"rule" tf:"rule,omitempty"`
}

type UserV3Observation struct {

	// The default project this user belongs to.
	DefaultProjectID *string `json:"defaultProjectId,omitempty" tf:"default_project_id,omitempty"`

	// A description of the user.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The domain this user belongs to.
	DomainID *string `json:"domainId,omitempty" tf:"domain_id,omitempty"`

	// Whether the user is enabled or disabled. Valid
	// values are true and false.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Free-form key/value pairs of extra information.
	Extra map[string]*string `json:"extra,omitempty" tf:"extra,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// User will not have to
	// change their password upon first use. Valid values are true and false.
	IgnoreChangePasswordUponFirstUse *bool `json:"ignoreChangePasswordUponFirstUse,omitempty" tf:"ignore_change_password_upon_first_use,omitempty"`

	// User will not have a failure
	// lockout placed on their account. Valid values are true and false.
	IgnoreLockoutFailureAttempts *bool `json:"ignoreLockoutFailureAttempts,omitempty" tf:"ignore_lockout_failure_attempts,omitempty"`

	// User's password will not expire.
	// Valid values are true and false.
	IgnorePasswordExpiry *bool `json:"ignorePasswordExpiry,omitempty" tf:"ignore_password_expiry,omitempty"`

	// Whether to enable multi-factor
	// authentication. Valid values are true and false.
	MultiFactorAuthEnabled *bool `json:"multiFactorAuthEnabled,omitempty" tf:"multi_factor_auth_enabled,omitempty"`

	// A multi-factor authentication rule.
	// The structure is documented below. Please see the
	// Ocata release notes
	// for more information on how to use mulit-factor rules.
	MultiFactorAuthRule []MultiFactorAuthRuleObservation `json:"multiFactorAuthRule,omitempty" tf:"multi_factor_auth_rule,omitempty"`

	// The name of the user.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The region in which to obtain the V3 Keystone client.
	// If omitted, the region argument of the provider is used. Changing this
	// creates a new User.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type UserV3Parameters struct {

	// The default project this user belongs to.
	// +kubebuilder:validation:Optional
	DefaultProjectID *string `json:"defaultProjectId,omitempty" tf:"default_project_id,omitempty"`

	// A description of the user.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The domain this user belongs to.
	// +kubebuilder:validation:Optional
	DomainID *string `json:"domainId,omitempty" tf:"domain_id,omitempty"`

	// Whether the user is enabled or disabled. Valid
	// values are true and false.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Free-form key/value pairs of extra information.
	// +kubebuilder:validation:Optional
	Extra map[string]*string `json:"extra,omitempty" tf:"extra,omitempty"`

	// User will not have to
	// change their password upon first use. Valid values are true and false.
	// +kubebuilder:validation:Optional
	IgnoreChangePasswordUponFirstUse *bool `json:"ignoreChangePasswordUponFirstUse,omitempty" tf:"ignore_change_password_upon_first_use,omitempty"`

	// User will not have a failure
	// lockout placed on their account. Valid values are true and false.
	// +kubebuilder:validation:Optional
	IgnoreLockoutFailureAttempts *bool `json:"ignoreLockoutFailureAttempts,omitempty" tf:"ignore_lockout_failure_attempts,omitempty"`

	// User's password will not expire.
	// Valid values are true and false.
	// +kubebuilder:validation:Optional
	IgnorePasswordExpiry *bool `json:"ignorePasswordExpiry,omitempty" tf:"ignore_password_expiry,omitempty"`

	// Whether to enable multi-factor
	// authentication. Valid values are true and false.
	// +kubebuilder:validation:Optional
	MultiFactorAuthEnabled *bool `json:"multiFactorAuthEnabled,omitempty" tf:"multi_factor_auth_enabled,omitempty"`

	// A multi-factor authentication rule.
	// The structure is documented below. Please see the
	// Ocata release notes
	// for more information on how to use mulit-factor rules.
	// +kubebuilder:validation:Optional
	MultiFactorAuthRule []MultiFactorAuthRuleParameters `json:"multiFactorAuthRule,omitempty" tf:"multi_factor_auth_rule,omitempty"`

	// The name of the user.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The password for the user.
	// +kubebuilder:validation:Optional
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// The region in which to obtain the V3 Keystone client.
	// If omitted, the region argument of the provider is used. Changing this
	// creates a new User.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

// UserV3Spec defines the desired state of UserV3
type UserV3Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     UserV3Parameters `json:"forProvider"`
}

// UserV3Status defines the observed state of UserV3.
type UserV3Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        UserV3Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// UserV3 is the Schema for the UserV3s API. Manages a V3 User resource within OpenStack Keystone.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,openstack}
type UserV3 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              UserV3Spec   `json:"spec"`
	Status            UserV3Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UserV3List contains a list of UserV3s
type UserV3List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UserV3 `json:"items"`
}

// Repository type metadata.
var (
	UserV3_Kind             = "UserV3"
	UserV3_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: UserV3_Kind}.String()
	UserV3_KindAPIVersion   = UserV3_Kind + "." + CRDGroupVersion.String()
	UserV3_GroupVersionKind = CRDGroupVersion.WithKind(UserV3_Kind)
)

func init() {
	SchemeBuilder.Register(&UserV3{}, &UserV3List{})
}
