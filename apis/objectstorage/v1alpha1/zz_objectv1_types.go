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

type ObjectV1InitParameters struct {

	// A unique (within an account) name for the container.
	// The container name must be from 1 to 256 characters long and can start
	// with any character and contain any pattern. Character set must be UTF-8.
	// The container name cannot contain a slash (/) character because this
	// character delimits the container and object name. For example, the path
	// /v1/account/www/pages specifies the www container, not the www/pages container.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-openstack/apis/objectstorage/v1alpha1.ContainerV1
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("name",false)
	ContainerName *string `json:"containerName,omitempty" tf:"container_name,omitempty"`

	// Reference to a ContainerV1 in objectstorage to populate containerName.
	// +kubebuilder:validation:Optional
	ContainerNameRef *v1.Reference `json:"containerNameRef,omitempty" tf:"-"`

	// Selector for a ContainerV1 in objectstorage to populate containerName.
	// +kubebuilder:validation:Optional
	ContainerNameSelector *v1.Selector `json:"containerNameSelector,omitempty" tf:"-"`

	// A string representing the content of the object. Conflicts with
	// source and copy_from.
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// A string which specifies the override behavior for
	// the browser. For example, this header might specify that the browser use a download
	// program to save this file rather than show the file, which is the default.
	ContentDisposition *string `json:"contentDisposition,omitempty" tf:"content_disposition,omitempty"`

	// A string representing the value of the Content-Encoding
	// metadata.
	ContentEncoding *string `json:"contentEncoding,omitempty" tf:"content_encoding,omitempty"`

	// A string which sets the MIME type for the object.
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// A string representing the name of an object
	// used to create the new object by copying the copy_from object. The value is in form
	// {container}/{object}. You must UTF-8-encode and then URL-encode the names of the
	// container and object before you include them in the header. Conflicts with source and
	// content.
	CopyFrom *string `json:"copyFrom,omitempty" tf:"copy_from,omitempty"`

	// An integer representing the number of seconds after which the
	// system removes the object. Internally, the Object Storage system stores this value in
	// the X-Delete-At metadata item.
	DeleteAfter *float64 `json:"deleteAfter,omitempty" tf:"delete_after,omitempty"`

	// An string representing the date when the system removes the object.
	// For example, "2015-08-26" is equivalent to Mon, Wed, 26 Aug 2015 00:00:00 GMT.
	DeleteAt *string `json:"deleteAt,omitempty" tf:"delete_at,omitempty"`

	// If set to true, Object Storage guesses the content
	// type based on the file extension and ignores the value sent in the Content-Type
	// header, if present.
	DetectContentType *bool `json:"detectContentType,omitempty" tf:"detect_content_type,omitempty"`

	// Used to trigger updates. The only meaningful value is ${md5(file("path/to/file"))}.
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	// +mapType=granular
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// A unique name for the object.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A string set to specify that this is a dynamic large
	// object manifest object. The value is the container and object name prefix of the
	// segment objects in the form container/prefix. You must UTF-8-encode and then
	// URL-encode the names of the container and prefix before you include them in this
	// header.
	ObjectManifest *string `json:"objectManifest,omitempty" tf:"object_manifest,omitempty"`

	// The region in which to create the container. If
	// omitted, the region argument of the provider is used. Changing this
	// creates a new container.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// A string representing the local path of a file which will be used
	// as the object's content. Conflicts with source and copy_from.
	Source *string `json:"source,omitempty" tf:"source,omitempty"`
}

type ObjectV1Observation struct {

	// A unique (within an account) name for the container.
	// The container name must be from 1 to 256 characters long and can start
	// with any character and contain any pattern. Character set must be UTF-8.
	// The container name cannot contain a slash (/) character because this
	// character delimits the container and object name. For example, the path
	// /v1/account/www/pages specifies the www container, not the www/pages container.
	ContainerName *string `json:"containerName,omitempty" tf:"container_name,omitempty"`

	// A string representing the content of the object. Conflicts with
	// source and copy_from.
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// A string which specifies the override behavior for
	// the browser. For example, this header might specify that the browser use a download
	// program to save this file rather than show the file, which is the default.
	ContentDisposition *string `json:"contentDisposition,omitempty" tf:"content_disposition,omitempty"`

	// A string representing the value of the Content-Encoding
	// metadata.
	ContentEncoding *string `json:"contentEncoding,omitempty" tf:"content_encoding,omitempty"`

	// If the operation succeeds, this value is zero (0) or the
	// length of informational or error text in the response body.
	ContentLength *float64 `json:"contentLength,omitempty" tf:"content_length,omitempty"`

	// A string which sets the MIME type for the object.
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// A string representing the name of an object
	// used to create the new object by copying the copy_from object. The value is in form
	// {container}/{object}. You must UTF-8-encode and then URL-encode the names of the
	// container and object before you include them in the header. Conflicts with source and
	// content.
	CopyFrom *string `json:"copyFrom,omitempty" tf:"copy_from,omitempty"`

	// The date and time the system responded to the request, using the preferred
	// format of RFC 7231 as shown in this example Thu, 16 Jun 2016 15:10:38 GMT. The
	// time is always in UTC.
	Date *string `json:"date,omitempty" tf:"date,omitempty"`

	// An integer representing the number of seconds after which the
	// system removes the object. Internally, the Object Storage system stores this value in
	// the X-Delete-At metadata item.
	DeleteAfter *float64 `json:"deleteAfter,omitempty" tf:"delete_after,omitempty"`

	// An string representing the date when the system removes the object.
	// For example, "2015-08-26" is equivalent to Mon, Wed, 26 Aug 2015 00:00:00 GMT.
	DeleteAt *string `json:"deleteAt,omitempty" tf:"delete_at,omitempty"`

	// If set to true, Object Storage guesses the content
	// type based on the file extension and ignores the value sent in the Content-Type
	// header, if present.
	DetectContentType *bool `json:"detectContentType,omitempty" tf:"detect_content_type,omitempty"`

	// Used to trigger updates. The only meaningful value is ${md5(file("path/to/file"))}.
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The date and time when the object was last modified. The date and time
	// stamp format is ISO 8601:
	// CCYY-MM-DDThh:mm:ss±hh:mm
	// For example, 2015-08-27T09:49:58-05:00.
	// The ±hh:mm value, if included, is the time zone as an offset from UTC. In the previous
	// example, the offset value is -05:00.
	LastModified *string `json:"lastModified,omitempty" tf:"last_modified,omitempty"`

	// +mapType=granular
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// A unique name for the object.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A string set to specify that this is a dynamic large
	// object manifest object. The value is the container and object name prefix of the
	// segment objects in the form container/prefix. You must UTF-8-encode and then
	// URL-encode the names of the container and prefix before you include them in this
	// header.
	ObjectManifest *string `json:"objectManifest,omitempty" tf:"object_manifest,omitempty"`

	// The region in which to create the container. If
	// omitted, the region argument of the provider is used. Changing this
	// creates a new container.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// A string representing the local path of a file which will be used
	// as the object's content. Conflicts with source and copy_from.
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// A unique transaction ID for this request. Your service provider might
	// need this value if you report a problem.
	TransID *string `json:"transId,omitempty" tf:"trans_id,omitempty"`
}

type ObjectV1Parameters struct {

	// A unique (within an account) name for the container.
	// The container name must be from 1 to 256 characters long and can start
	// with any character and contain any pattern. Character set must be UTF-8.
	// The container name cannot contain a slash (/) character because this
	// character delimits the container and object name. For example, the path
	// /v1/account/www/pages specifies the www container, not the www/pages container.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-openstack/apis/objectstorage/v1alpha1.ContainerV1
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("name",false)
	// +kubebuilder:validation:Optional
	ContainerName *string `json:"containerName,omitempty" tf:"container_name,omitempty"`

	// Reference to a ContainerV1 in objectstorage to populate containerName.
	// +kubebuilder:validation:Optional
	ContainerNameRef *v1.Reference `json:"containerNameRef,omitempty" tf:"-"`

	// Selector for a ContainerV1 in objectstorage to populate containerName.
	// +kubebuilder:validation:Optional
	ContainerNameSelector *v1.Selector `json:"containerNameSelector,omitempty" tf:"-"`

	// A string representing the content of the object. Conflicts with
	// source and copy_from.
	// +kubebuilder:validation:Optional
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// A string which specifies the override behavior for
	// the browser. For example, this header might specify that the browser use a download
	// program to save this file rather than show the file, which is the default.
	// +kubebuilder:validation:Optional
	ContentDisposition *string `json:"contentDisposition,omitempty" tf:"content_disposition,omitempty"`

	// A string representing the value of the Content-Encoding
	// metadata.
	// +kubebuilder:validation:Optional
	ContentEncoding *string `json:"contentEncoding,omitempty" tf:"content_encoding,omitempty"`

	// A string which sets the MIME type for the object.
	// +kubebuilder:validation:Optional
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// A string representing the name of an object
	// used to create the new object by copying the copy_from object. The value is in form
	// {container}/{object}. You must UTF-8-encode and then URL-encode the names of the
	// container and object before you include them in the header. Conflicts with source and
	// content.
	// +kubebuilder:validation:Optional
	CopyFrom *string `json:"copyFrom,omitempty" tf:"copy_from,omitempty"`

	// An integer representing the number of seconds after which the
	// system removes the object. Internally, the Object Storage system stores this value in
	// the X-Delete-At metadata item.
	// +kubebuilder:validation:Optional
	DeleteAfter *float64 `json:"deleteAfter,omitempty" tf:"delete_after,omitempty"`

	// An string representing the date when the system removes the object.
	// For example, "2015-08-26" is equivalent to Mon, Wed, 26 Aug 2015 00:00:00 GMT.
	// +kubebuilder:validation:Optional
	DeleteAt *string `json:"deleteAt,omitempty" tf:"delete_at,omitempty"`

	// If set to true, Object Storage guesses the content
	// type based on the file extension and ignores the value sent in the Content-Type
	// header, if present.
	// +kubebuilder:validation:Optional
	DetectContentType *bool `json:"detectContentType,omitempty" tf:"detect_content_type,omitempty"`

	// Used to trigger updates. The only meaningful value is ${md5(file("path/to/file"))}.
	// +kubebuilder:validation:Optional
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	// +kubebuilder:validation:Optional
	// +mapType=granular
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// A unique name for the object.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A string set to specify that this is a dynamic large
	// object manifest object. The value is the container and object name prefix of the
	// segment objects in the form container/prefix. You must UTF-8-encode and then
	// URL-encode the names of the container and prefix before you include them in this
	// header.
	// +kubebuilder:validation:Optional
	ObjectManifest *string `json:"objectManifest,omitempty" tf:"object_manifest,omitempty"`

	// The region in which to create the container. If
	// omitted, the region argument of the provider is used. Changing this
	// creates a new container.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// A string representing the local path of a file which will be used
	// as the object's content. Conflicts with source and copy_from.
	// +kubebuilder:validation:Optional
	Source *string `json:"source,omitempty" tf:"source,omitempty"`
}

// ObjectV1Spec defines the desired state of ObjectV1
type ObjectV1Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ObjectV1Parameters `json:"forProvider"`
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
	InitProvider ObjectV1InitParameters `json:"initProvider,omitempty"`
}

// ObjectV1Status defines the observed state of ObjectV1.
type ObjectV1Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ObjectV1Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ObjectV1 is the Schema for the ObjectV1s API. Manages a V1 container object resource within OpenStack.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,openstack}
type ObjectV1 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   ObjectV1Spec   `json:"spec"`
	Status ObjectV1Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ObjectV1List contains a list of ObjectV1s
type ObjectV1List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ObjectV1 `json:"items"`
}

// Repository type metadata.
var (
	ObjectV1_Kind             = "ObjectV1"
	ObjectV1_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ObjectV1_Kind}.String()
	ObjectV1_KindAPIVersion   = ObjectV1_Kind + "." + CRDGroupVersion.String()
	ObjectV1_GroupVersionKind = CRDGroupVersion.WithKind(ObjectV1_Kind)
)

func init() {
	SchemeBuilder.Register(&ObjectV1{}, &ObjectV1List{})
}
