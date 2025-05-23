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

type OutputsInitParameters struct {

	// The description of the stack resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	OutputKey *string `json:"outputKey,omitempty" tf:"output_key,omitempty"`

	OutputValue *string `json:"outputValue,omitempty" tf:"output_value,omitempty"`
}

type OutputsObservation struct {

	// The description of the stack resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	OutputKey *string `json:"outputKey,omitempty" tf:"output_key,omitempty"`

	OutputValue *string `json:"outputValue,omitempty" tf:"output_value,omitempty"`
}

type OutputsParameters struct {

	// The description of the stack resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	OutputKey *string `json:"outputKey" tf:"output_key,omitempty"`

	// +kubebuilder:validation:Optional
	OutputValue *string `json:"outputValue" tf:"output_value,omitempty"`
}

type StackV1InitParameters struct {

	// List of stack capabilities for stack.
	Capabilities []*string `json:"capabilities,omitempty" tf:"capabilities,omitempty"`

	// The date and time when the resource was created. The date
	// and time stamp format is ISO 8601: CCYY-MM-DDThh:mm:ss±hh:mm
	// For example, 2015-08-27T09:49:58-05:00. The ±hh:mm value, if included,
	// is the time zone as an offset from UTC.
	CreationTime *string `json:"creationTime,omitempty" tf:"creation_time,omitempty"`

	// The description of the stack resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Enables or disables deletion of all stack
	// resources when a stack creation fails. Default is true, meaning all
	// resources are not deleted when stack creation fails.
	DisableRollback *bool `json:"disableRollback,omitempty" tf:"disable_rollback,omitempty"`

	// Environment key/value pairs to associate with
	// the stack which contains details for the environment of the stack.
	// Allowed keys: Bin, URL, Files. Changing this updates the existing stack
	// Environment Opts.
	// +mapType=granular
	EnvironmentOpts map[string]*string `json:"environmentOpts,omitempty" tf:"environment_opts,omitempty"`

	// A unique name for the stack. It must start with an
	// alphabetic character. Changing this updates the stack's name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// List of notification topics for stack.
	NotificationTopics []*string `json:"notificationTopics,omitempty" tf:"notification_topics,omitempty"`

	// A list of stack outputs.
	Outputs []OutputsInitParameters `json:"outputs,omitempty" tf:"outputs,omitempty"`

	// User-defined key/value pairs as parameters to pass
	// to the template. Changing this updates the existing stack parameters.
	// +mapType=granular
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// The region in which to create the stack. If
	// omitted, the region argument of the provider is used. Changing this
	// creates a new stack.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The status of the stack.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// The reason for the current status of the stack.
	StatusReason *string `json:"statusReason,omitempty" tf:"status_reason,omitempty"`

	// A list of tags to assosciate with the Stack
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The description of the stack template.
	TemplateDescription *string `json:"templateDescription,omitempty" tf:"template_description,omitempty"`

	// Template key/value pairs to associate with the
	// stack which contains either the template file or url.
	// Allowed keys: Bin, URL, Files. Changing this updates the existing stack
	// Template Opts.
	// +mapType=granular
	TemplateOpts map[string]*string `json:"templateOpts,omitempty" tf:"template_opts,omitempty"`

	// The timeout for stack action in minutes.
	Timeout *float64 `json:"timeout,omitempty" tf:"timeout,omitempty"`

	// The date and time when the resource was updated. The date
	// and time stamp format is ISO 8601: CCYY-MM-DDThh:mm:ss±hh:mm
	// For example, 2015-08-27T09:49:58-05:00. The ±hh:mm value, if included,
	// is the time zone as an offset from UTC.
	UpdatedTime *string `json:"updatedTime,omitempty" tf:"updated_time,omitempty"`
}

type StackV1Observation struct {

	// List of stack capabilities for stack.
	Capabilities []*string `json:"capabilities,omitempty" tf:"capabilities,omitempty"`

	// The date and time when the resource was created. The date
	// and time stamp format is ISO 8601: CCYY-MM-DDThh:mm:ss±hh:mm
	// For example, 2015-08-27T09:49:58-05:00. The ±hh:mm value, if included,
	// is the time zone as an offset from UTC.
	CreationTime *string `json:"creationTime,omitempty" tf:"creation_time,omitempty"`

	// The description of the stack resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Enables or disables deletion of all stack
	// resources when a stack creation fails. Default is true, meaning all
	// resources are not deleted when stack creation fails.
	DisableRollback *bool `json:"disableRollback,omitempty" tf:"disable_rollback,omitempty"`

	// Environment key/value pairs to associate with
	// the stack which contains details for the environment of the stack.
	// Allowed keys: Bin, URL, Files. Changing this updates the existing stack
	// Environment Opts.
	// +mapType=granular
	EnvironmentOpts map[string]*string `json:"environmentOpts,omitempty" tf:"environment_opts,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A unique name for the stack. It must start with an
	// alphabetic character. Changing this updates the stack's name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// List of notification topics for stack.
	NotificationTopics []*string `json:"notificationTopics,omitempty" tf:"notification_topics,omitempty"`

	// A list of stack outputs.
	Outputs []OutputsObservation `json:"outputs,omitempty" tf:"outputs,omitempty"`

	// User-defined key/value pairs as parameters to pass
	// to the template. Changing this updates the existing stack parameters.
	// +mapType=granular
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// The region in which to create the stack. If
	// omitted, the region argument of the provider is used. Changing this
	// creates a new stack.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The status of the stack.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// The reason for the current status of the stack.
	StatusReason *string `json:"statusReason,omitempty" tf:"status_reason,omitempty"`

	// A list of tags to assosciate with the Stack
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The description of the stack template.
	TemplateDescription *string `json:"templateDescription,omitempty" tf:"template_description,omitempty"`

	// Template key/value pairs to associate with the
	// stack which contains either the template file or url.
	// Allowed keys: Bin, URL, Files. Changing this updates the existing stack
	// Template Opts.
	// +mapType=granular
	TemplateOpts map[string]*string `json:"templateOpts,omitempty" tf:"template_opts,omitempty"`

	// The timeout for stack action in minutes.
	Timeout *float64 `json:"timeout,omitempty" tf:"timeout,omitempty"`

	// The date and time when the resource was updated. The date
	// and time stamp format is ISO 8601: CCYY-MM-DDThh:mm:ss±hh:mm
	// For example, 2015-08-27T09:49:58-05:00. The ±hh:mm value, if included,
	// is the time zone as an offset from UTC.
	UpdatedTime *string `json:"updatedTime,omitempty" tf:"updated_time,omitempty"`
}

type StackV1Parameters struct {

	// List of stack capabilities for stack.
	// +kubebuilder:validation:Optional
	Capabilities []*string `json:"capabilities,omitempty" tf:"capabilities,omitempty"`

	// The date and time when the resource was created. The date
	// and time stamp format is ISO 8601: CCYY-MM-DDThh:mm:ss±hh:mm
	// For example, 2015-08-27T09:49:58-05:00. The ±hh:mm value, if included,
	// is the time zone as an offset from UTC.
	// +kubebuilder:validation:Optional
	CreationTime *string `json:"creationTime,omitempty" tf:"creation_time,omitempty"`

	// The description of the stack resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Enables or disables deletion of all stack
	// resources when a stack creation fails. Default is true, meaning all
	// resources are not deleted when stack creation fails.
	// +kubebuilder:validation:Optional
	DisableRollback *bool `json:"disableRollback,omitempty" tf:"disable_rollback,omitempty"`

	// Environment key/value pairs to associate with
	// the stack which contains details for the environment of the stack.
	// Allowed keys: Bin, URL, Files. Changing this updates the existing stack
	// Environment Opts.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	EnvironmentOpts map[string]*string `json:"environmentOpts,omitempty" tf:"environment_opts,omitempty"`

	// A unique name for the stack. It must start with an
	// alphabetic character. Changing this updates the stack's name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// List of notification topics for stack.
	// +kubebuilder:validation:Optional
	NotificationTopics []*string `json:"notificationTopics,omitempty" tf:"notification_topics,omitempty"`

	// A list of stack outputs.
	// +kubebuilder:validation:Optional
	Outputs []OutputsParameters `json:"outputs,omitempty" tf:"outputs,omitempty"`

	// User-defined key/value pairs as parameters to pass
	// to the template. Changing this updates the existing stack parameters.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// The region in which to create the stack. If
	// omitted, the region argument of the provider is used. Changing this
	// creates a new stack.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The status of the stack.
	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// The reason for the current status of the stack.
	// +kubebuilder:validation:Optional
	StatusReason *string `json:"statusReason,omitempty" tf:"status_reason,omitempty"`

	// A list of tags to assosciate with the Stack
	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The description of the stack template.
	// +kubebuilder:validation:Optional
	TemplateDescription *string `json:"templateDescription,omitempty" tf:"template_description,omitempty"`

	// Template key/value pairs to associate with the
	// stack which contains either the template file or url.
	// Allowed keys: Bin, URL, Files. Changing this updates the existing stack
	// Template Opts.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	TemplateOpts map[string]*string `json:"templateOpts,omitempty" tf:"template_opts,omitempty"`

	// The timeout for stack action in minutes.
	// +kubebuilder:validation:Optional
	Timeout *float64 `json:"timeout,omitempty" tf:"timeout,omitempty"`

	// The date and time when the resource was updated. The date
	// and time stamp format is ISO 8601: CCYY-MM-DDThh:mm:ss±hh:mm
	// For example, 2015-08-27T09:49:58-05:00. The ±hh:mm value, if included,
	// is the time zone as an offset from UTC.
	// +kubebuilder:validation:Optional
	UpdatedTime *string `json:"updatedTime,omitempty" tf:"updated_time,omitempty"`
}

// StackV1Spec defines the desired state of StackV1
type StackV1Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     StackV1Parameters `json:"forProvider"`
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
	InitProvider StackV1InitParameters `json:"initProvider,omitempty"`
}

// StackV1Status defines the observed state of StackV1.
type StackV1Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        StackV1Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// StackV1 is the Schema for the StackV1s API. Manages a V1 stack resource within OpenStack.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,openstack}
type StackV1 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.templateOpts) || (has(self.initProvider) && has(self.initProvider.templateOpts))",message="spec.forProvider.templateOpts is a required parameter"
	Spec   StackV1Spec   `json:"spec"`
	Status StackV1Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StackV1List contains a list of StackV1s
type StackV1List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StackV1 `json:"items"`
}

// Repository type metadata.
var (
	StackV1_Kind             = "StackV1"
	StackV1_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: StackV1_Kind}.String()
	StackV1_KindAPIVersion   = StackV1_Kind + "." + CRDGroupVersion.String()
	StackV1_GroupVersionKind = CRDGroupVersion.WithKind(StackV1_Kind)
)

func init() {
	SchemeBuilder.Register(&StackV1{}, &StackV1List{})
}
