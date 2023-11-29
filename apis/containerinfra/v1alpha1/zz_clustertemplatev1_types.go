// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

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

type ClustertemplateV1InitParameters struct {

	// The API server port for the Container
	// Orchestration Engine for this cluster template. Changing this updates the
	// API server port of the existing cluster template.
	ApiserverPort *float64 `json:"apiserverPort,omitempty" tf:"apiserver_port,omitempty"`

	// The distro for the cluster (fedora-atomic,
	// coreos, etc.). Changing this updates the cluster distro of the existing
	// cluster template.
	ClusterDistro *string `json:"clusterDistro,omitempty" tf:"cluster_distro,omitempty"`

	// The Container Orchestration Engine for this cluster
	// template. Changing this updates the engine of the existing cluster
	// template.
	Coe *string `json:"coe,omitempty" tf:"coe,omitempty"`

	// Address of the DNS nameserver that is used in
	// nodes of the cluster. Changing this updates the DNS nameserver of the
	// existing cluster template.
	DNSNameserver *string `json:"dnsNameserver,omitempty" tf:"dns_nameserver,omitempty"`

	// Docker storage driver. Changing this
	// updates the Docker storage driver of the existing cluster template.
	DockerStorageDriver *string `json:"dockerStorageDriver,omitempty" tf:"docker_storage_driver,omitempty"`

	// The size (in GB) of the Docker volume.
	// Changing this updates the Docker volume size of the existing cluster
	// template.
	DockerVolumeSize *float64 `json:"dockerVolumeSize,omitempty" tf:"docker_volume_size,omitempty"`

	// The ID of the external network that will
	// be used for the cluster. Changing this updates the external network ID of
	// the existing cluster template.
	ExternalNetworkID *string `json:"externalNetworkId,omitempty" tf:"external_network_id,omitempty"`

	// The fixed network that will be attached to the
	// cluster. Changing this updates the fixed network of the existing cluster
	// template.
	FixedNetwork *string `json:"fixedNetwork,omitempty" tf:"fixed_network,omitempty"`

	// The fixed subnet that will be attached to the
	// cluster. Changing this updates the fixed subnet of the existing cluster
	// template.
	FixedSubnet *string `json:"fixedSubnet,omitempty" tf:"fixed_subnet,omitempty"`

	// The flavor for the nodes of the cluster. Can be set via
	// the OS_MAGNUM_FLAVOR environment variable. Changing this updates the
	// flavor of the existing cluster template.
	Flavor *string `json:"flavor,omitempty" tf:"flavor,omitempty"`

	// Indicates whether created cluster should
	// create floating IP for every node or not. Changing this updates the
	// floating IP enabled attribute of the existing cluster template.
	FloatingIPEnabled *bool `json:"floatingIpEnabled,omitempty" tf:"floating_ip_enabled,omitempty"`

	// The address of a proxy for receiving all HTTP
	// requests and relay them. Changing this updates the HTTP proxy address of
	// the existing cluster template.
	HTTPProxy *string `json:"httpProxy,omitempty" tf:"http_proxy,omitempty"`

	// The address of a proxy for receiving all HTTPS
	// requests and relay them. Changing this updates the HTTPS proxy address of
	// the existing cluster template.
	HTTPSProxy *string `json:"httpsProxy,omitempty" tf:"https_proxy,omitempty"`

	// Indicates whether the ClusterTemplate is hidden or not.
	// Changing this updates the hidden attribute of the existing cluster
	// template.
	Hidden *bool `json:"hidden,omitempty" tf:"hidden,omitempty"`

	// The reference to an image that is used for nodes of the
	// cluster. Can be set via the OS_MAGNUM_IMAGE environment variable.
	// Changing this updates the image attribute of the existing cluster template.
	Image *string `json:"image,omitempty" tf:"image,omitempty"`

	// The insecure registry URL for the cluster
	// template. Changing this updates the insecure registry attribute of the
	// existing cluster template.
	InsecureRegistry *string `json:"insecureRegistry,omitempty" tf:"insecure_registry,omitempty"`

	// The name of the Compute service SSH keypair.
	// Changing this updates the keypair of the existing cluster template.
	KeypairID *string `json:"keypairId,omitempty" tf:"keypair_id,omitempty"`

	// The list of key value pairs representing additional
	// properties of the cluster template. Changing this updates the labels of the
	// existing cluster template.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The flavor for the master nodes. Can be set via
	// the OS_MAGNUM_MASTER_FLAVOR environment variable. Changing this updates
	// the master flavor of the existing cluster template.
	MasterFlavor *string `json:"masterFlavor,omitempty" tf:"master_flavor,omitempty"`

	// Indicates whether created cluster should
	// has a loadbalancer for master nodes or not. Changing this updates the
	// attribute of the existing cluster template.
	MasterLBEnabled *bool `json:"masterLbEnabled,omitempty" tf:"master_lb_enabled,omitempty"`

	// The name of the cluster template. Changing this updates
	// the name of the existing cluster template.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of the driver for the container
	// network. Changing this updates the network driver of the existing cluster
	// template.
	NetworkDriver *string `json:"networkDriver,omitempty" tf:"network_driver,omitempty"`

	// A comma-separated list of IP addresses that shouldn't
	// be used in the cluster. Changing this updates the no proxy list of the
	// existing cluster template.
	NoProxy *string `json:"noProxy,omitempty" tf:"no_proxy,omitempty"`

	// Indicates whether cluster template should be public.
	// Changing this updates the public attribute of the existing cluster
	// template.
	Public *bool `json:"public,omitempty" tf:"public,omitempty"`

	// The region in which to obtain the V1 Container Infra
	// client. A Container Infra client is needed to create a cluster template. If
	// omitted,the region argument of the provider is used. Changing this
	// creates a new cluster template.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Indicates whether Docker registry is enabled
	// in the cluster. Changing this updates the registry enabled attribute of the
	// existing cluster template.
	RegistryEnabled *bool `json:"registryEnabled,omitempty" tf:"registry_enabled,omitempty"`

	// The server type for the cluster template. Changing
	// this updates the server type of the existing cluster template.
	ServerType *string `json:"serverType,omitempty" tf:"server_type,omitempty"`

	// Indicates whether the TLS should be disabled in
	// the cluster. Changing this updates the attribute of the existing cluster.
	TLSDisabled *bool `json:"tlsDisabled,omitempty" tf:"tls_disabled,omitempty"`

	// The name of the driver that is used for the
	// volumes of the cluster nodes. Changing this updates the volume driver of
	// the existing cluster template.
	VolumeDriver *string `json:"volumeDriver,omitempty" tf:"volume_driver,omitempty"`
}

type ClustertemplateV1Observation struct {

	// The API server port for the Container
	// Orchestration Engine for this cluster template. Changing this updates the
	// API server port of the existing cluster template.
	ApiserverPort *float64 `json:"apiserverPort,omitempty" tf:"apiserver_port,omitempty"`

	// The distro for the cluster (fedora-atomic,
	// coreos, etc.). Changing this updates the cluster distro of the existing
	// cluster template.
	ClusterDistro *string `json:"clusterDistro,omitempty" tf:"cluster_distro,omitempty"`

	// The Container Orchestration Engine for this cluster
	// template. Changing this updates the engine of the existing cluster
	// template.
	Coe *string `json:"coe,omitempty" tf:"coe,omitempty"`

	// The time at which cluster template was created.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// Address of the DNS nameserver that is used in
	// nodes of the cluster. Changing this updates the DNS nameserver of the
	// existing cluster template.
	DNSNameserver *string `json:"dnsNameserver,omitempty" tf:"dns_nameserver,omitempty"`

	// Docker storage driver. Changing this
	// updates the Docker storage driver of the existing cluster template.
	DockerStorageDriver *string `json:"dockerStorageDriver,omitempty" tf:"docker_storage_driver,omitempty"`

	// The size (in GB) of the Docker volume.
	// Changing this updates the Docker volume size of the existing cluster
	// template.
	DockerVolumeSize *float64 `json:"dockerVolumeSize,omitempty" tf:"docker_volume_size,omitempty"`

	// The ID of the external network that will
	// be used for the cluster. Changing this updates the external network ID of
	// the existing cluster template.
	ExternalNetworkID *string `json:"externalNetworkId,omitempty" tf:"external_network_id,omitempty"`

	// The fixed network that will be attached to the
	// cluster. Changing this updates the fixed network of the existing cluster
	// template.
	FixedNetwork *string `json:"fixedNetwork,omitempty" tf:"fixed_network,omitempty"`

	// The fixed subnet that will be attached to the
	// cluster. Changing this updates the fixed subnet of the existing cluster
	// template.
	FixedSubnet *string `json:"fixedSubnet,omitempty" tf:"fixed_subnet,omitempty"`

	// The flavor for the nodes of the cluster. Can be set via
	// the OS_MAGNUM_FLAVOR environment variable. Changing this updates the
	// flavor of the existing cluster template.
	Flavor *string `json:"flavor,omitempty" tf:"flavor,omitempty"`

	// Indicates whether created cluster should
	// create floating IP for every node or not. Changing this updates the
	// floating IP enabled attribute of the existing cluster template.
	FloatingIPEnabled *bool `json:"floatingIpEnabled,omitempty" tf:"floating_ip_enabled,omitempty"`

	// The address of a proxy for receiving all HTTP
	// requests and relay them. Changing this updates the HTTP proxy address of
	// the existing cluster template.
	HTTPProxy *string `json:"httpProxy,omitempty" tf:"http_proxy,omitempty"`

	// The address of a proxy for receiving all HTTPS
	// requests and relay them. Changing this updates the HTTPS proxy address of
	// the existing cluster template.
	HTTPSProxy *string `json:"httpsProxy,omitempty" tf:"https_proxy,omitempty"`

	// Indicates whether the ClusterTemplate is hidden or not.
	// Changing this updates the hidden attribute of the existing cluster
	// template.
	Hidden *bool `json:"hidden,omitempty" tf:"hidden,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The reference to an image that is used for nodes of the
	// cluster. Can be set via the OS_MAGNUM_IMAGE environment variable.
	// Changing this updates the image attribute of the existing cluster template.
	Image *string `json:"image,omitempty" tf:"image,omitempty"`

	// The insecure registry URL for the cluster
	// template. Changing this updates the insecure registry attribute of the
	// existing cluster template.
	InsecureRegistry *string `json:"insecureRegistry,omitempty" tf:"insecure_registry,omitempty"`

	// The name of the Compute service SSH keypair.
	// Changing this updates the keypair of the existing cluster template.
	KeypairID *string `json:"keypairId,omitempty" tf:"keypair_id,omitempty"`

	// The list of key value pairs representing additional
	// properties of the cluster template. Changing this updates the labels of the
	// existing cluster template.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The flavor for the master nodes. Can be set via
	// the OS_MAGNUM_MASTER_FLAVOR environment variable. Changing this updates
	// the master flavor of the existing cluster template.
	MasterFlavor *string `json:"masterFlavor,omitempty" tf:"master_flavor,omitempty"`

	// Indicates whether created cluster should
	// has a loadbalancer for master nodes or not. Changing this updates the
	// attribute of the existing cluster template.
	MasterLBEnabled *bool `json:"masterLbEnabled,omitempty" tf:"master_lb_enabled,omitempty"`

	// The name of the cluster template. Changing this updates
	// the name of the existing cluster template.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of the driver for the container
	// network. Changing this updates the network driver of the existing cluster
	// template.
	NetworkDriver *string `json:"networkDriver,omitempty" tf:"network_driver,omitempty"`

	// A comma-separated list of IP addresses that shouldn't
	// be used in the cluster. Changing this updates the no proxy list of the
	// existing cluster template.
	NoProxy *string `json:"noProxy,omitempty" tf:"no_proxy,omitempty"`

	// The project of the cluster template. Required if
	// admin wants to create a cluster template in another project. Changing this
	// creates a new cluster template.
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Indicates whether cluster template should be public.
	// Changing this updates the public attribute of the existing cluster
	// template.
	Public *bool `json:"public,omitempty" tf:"public,omitempty"`

	// The region in which to obtain the V1 Container Infra
	// client. A Container Infra client is needed to create a cluster template. If
	// omitted,the region argument of the provider is used. Changing this
	// creates a new cluster template.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Indicates whether Docker registry is enabled
	// in the cluster. Changing this updates the registry enabled attribute of the
	// existing cluster template.
	RegistryEnabled *bool `json:"registryEnabled,omitempty" tf:"registry_enabled,omitempty"`

	// The server type for the cluster template. Changing
	// this updates the server type of the existing cluster template.
	ServerType *string `json:"serverType,omitempty" tf:"server_type,omitempty"`

	// Indicates whether the TLS should be disabled in
	// the cluster. Changing this updates the attribute of the existing cluster.
	TLSDisabled *bool `json:"tlsDisabled,omitempty" tf:"tls_disabled,omitempty"`

	// The time at which cluster template was created.
	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`

	// The user of the cluster template. Required if admin
	// wants to create a cluster template for another user. Changing this creates
	// a new cluster template.
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`

	// The name of the driver that is used for the
	// volumes of the cluster nodes. Changing this updates the volume driver of
	// the existing cluster template.
	VolumeDriver *string `json:"volumeDriver,omitempty" tf:"volume_driver,omitempty"`
}

type ClustertemplateV1Parameters struct {

	// The API server port for the Container
	// Orchestration Engine for this cluster template. Changing this updates the
	// API server port of the existing cluster template.
	// +kubebuilder:validation:Optional
	ApiserverPort *float64 `json:"apiserverPort,omitempty" tf:"apiserver_port,omitempty"`

	// The distro for the cluster (fedora-atomic,
	// coreos, etc.). Changing this updates the cluster distro of the existing
	// cluster template.
	// +kubebuilder:validation:Optional
	ClusterDistro *string `json:"clusterDistro,omitempty" tf:"cluster_distro,omitempty"`

	// The Container Orchestration Engine for this cluster
	// template. Changing this updates the engine of the existing cluster
	// template.
	// +kubebuilder:validation:Optional
	Coe *string `json:"coe,omitempty" tf:"coe,omitempty"`

	// Address of the DNS nameserver that is used in
	// nodes of the cluster. Changing this updates the DNS nameserver of the
	// existing cluster template.
	// +kubebuilder:validation:Optional
	DNSNameserver *string `json:"dnsNameserver,omitempty" tf:"dns_nameserver,omitempty"`

	// Docker storage driver. Changing this
	// updates the Docker storage driver of the existing cluster template.
	// +kubebuilder:validation:Optional
	DockerStorageDriver *string `json:"dockerStorageDriver,omitempty" tf:"docker_storage_driver,omitempty"`

	// The size (in GB) of the Docker volume.
	// Changing this updates the Docker volume size of the existing cluster
	// template.
	// +kubebuilder:validation:Optional
	DockerVolumeSize *float64 `json:"dockerVolumeSize,omitempty" tf:"docker_volume_size,omitempty"`

	// The ID of the external network that will
	// be used for the cluster. Changing this updates the external network ID of
	// the existing cluster template.
	// +kubebuilder:validation:Optional
	ExternalNetworkID *string `json:"externalNetworkId,omitempty" tf:"external_network_id,omitempty"`

	// The fixed network that will be attached to the
	// cluster. Changing this updates the fixed network of the existing cluster
	// template.
	// +kubebuilder:validation:Optional
	FixedNetwork *string `json:"fixedNetwork,omitempty" tf:"fixed_network,omitempty"`

	// The fixed subnet that will be attached to the
	// cluster. Changing this updates the fixed subnet of the existing cluster
	// template.
	// +kubebuilder:validation:Optional
	FixedSubnet *string `json:"fixedSubnet,omitempty" tf:"fixed_subnet,omitempty"`

	// The flavor for the nodes of the cluster. Can be set via
	// the OS_MAGNUM_FLAVOR environment variable. Changing this updates the
	// flavor of the existing cluster template.
	// +kubebuilder:validation:Optional
	Flavor *string `json:"flavor,omitempty" tf:"flavor,omitempty"`

	// Indicates whether created cluster should
	// create floating IP for every node or not. Changing this updates the
	// floating IP enabled attribute of the existing cluster template.
	// +kubebuilder:validation:Optional
	FloatingIPEnabled *bool `json:"floatingIpEnabled,omitempty" tf:"floating_ip_enabled,omitempty"`

	// The address of a proxy for receiving all HTTP
	// requests and relay them. Changing this updates the HTTP proxy address of
	// the existing cluster template.
	// +kubebuilder:validation:Optional
	HTTPProxy *string `json:"httpProxy,omitempty" tf:"http_proxy,omitempty"`

	// The address of a proxy for receiving all HTTPS
	// requests and relay them. Changing this updates the HTTPS proxy address of
	// the existing cluster template.
	// +kubebuilder:validation:Optional
	HTTPSProxy *string `json:"httpsProxy,omitempty" tf:"https_proxy,omitempty"`

	// Indicates whether the ClusterTemplate is hidden or not.
	// Changing this updates the hidden attribute of the existing cluster
	// template.
	// +kubebuilder:validation:Optional
	Hidden *bool `json:"hidden,omitempty" tf:"hidden,omitempty"`

	// The reference to an image that is used for nodes of the
	// cluster. Can be set via the OS_MAGNUM_IMAGE environment variable.
	// Changing this updates the image attribute of the existing cluster template.
	// +kubebuilder:validation:Optional
	Image *string `json:"image,omitempty" tf:"image,omitempty"`

	// The insecure registry URL for the cluster
	// template. Changing this updates the insecure registry attribute of the
	// existing cluster template.
	// +kubebuilder:validation:Optional
	InsecureRegistry *string `json:"insecureRegistry,omitempty" tf:"insecure_registry,omitempty"`

	// The name of the Compute service SSH keypair.
	// Changing this updates the keypair of the existing cluster template.
	// +kubebuilder:validation:Optional
	KeypairID *string `json:"keypairId,omitempty" tf:"keypair_id,omitempty"`

	// The list of key value pairs representing additional
	// properties of the cluster template. Changing this updates the labels of the
	// existing cluster template.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The flavor for the master nodes. Can be set via
	// the OS_MAGNUM_MASTER_FLAVOR environment variable. Changing this updates
	// the master flavor of the existing cluster template.
	// +kubebuilder:validation:Optional
	MasterFlavor *string `json:"masterFlavor,omitempty" tf:"master_flavor,omitempty"`

	// Indicates whether created cluster should
	// has a loadbalancer for master nodes or not. Changing this updates the
	// attribute of the existing cluster template.
	// +kubebuilder:validation:Optional
	MasterLBEnabled *bool `json:"masterLbEnabled,omitempty" tf:"master_lb_enabled,omitempty"`

	// The name of the cluster template. Changing this updates
	// the name of the existing cluster template.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of the driver for the container
	// network. Changing this updates the network driver of the existing cluster
	// template.
	// +kubebuilder:validation:Optional
	NetworkDriver *string `json:"networkDriver,omitempty" tf:"network_driver,omitempty"`

	// A comma-separated list of IP addresses that shouldn't
	// be used in the cluster. Changing this updates the no proxy list of the
	// existing cluster template.
	// +kubebuilder:validation:Optional
	NoProxy *string `json:"noProxy,omitempty" tf:"no_proxy,omitempty"`

	// Indicates whether cluster template should be public.
	// Changing this updates the public attribute of the existing cluster
	// template.
	// +kubebuilder:validation:Optional
	Public *bool `json:"public,omitempty" tf:"public,omitempty"`

	// The region in which to obtain the V1 Container Infra
	// client. A Container Infra client is needed to create a cluster template. If
	// omitted,the region argument of the provider is used. Changing this
	// creates a new cluster template.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Indicates whether Docker registry is enabled
	// in the cluster. Changing this updates the registry enabled attribute of the
	// existing cluster template.
	// +kubebuilder:validation:Optional
	RegistryEnabled *bool `json:"registryEnabled,omitempty" tf:"registry_enabled,omitempty"`

	// The server type for the cluster template. Changing
	// this updates the server type of the existing cluster template.
	// +kubebuilder:validation:Optional
	ServerType *string `json:"serverType,omitempty" tf:"server_type,omitempty"`

	// Indicates whether the TLS should be disabled in
	// the cluster. Changing this updates the attribute of the existing cluster.
	// +kubebuilder:validation:Optional
	TLSDisabled *bool `json:"tlsDisabled,omitempty" tf:"tls_disabled,omitempty"`

	// The name of the driver that is used for the
	// volumes of the cluster nodes. Changing this updates the volume driver of
	// the existing cluster template.
	// +kubebuilder:validation:Optional
	VolumeDriver *string `json:"volumeDriver,omitempty" tf:"volume_driver,omitempty"`
}

// ClustertemplateV1Spec defines the desired state of ClustertemplateV1
type ClustertemplateV1Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ClustertemplateV1Parameters `json:"forProvider"`
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
	InitProvider ClustertemplateV1InitParameters `json:"initProvider,omitempty"`
}

// ClustertemplateV1Status defines the observed state of ClustertemplateV1.
type ClustertemplateV1Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ClustertemplateV1Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ClustertemplateV1 is the Schema for the ClustertemplateV1s API. Manages a V1 Magnum cluster template resource within OpenStack.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,openstack}
type ClustertemplateV1 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.coe) || (has(self.initProvider) && has(self.initProvider.coe))",message="spec.forProvider.coe is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.image) || (has(self.initProvider) && has(self.initProvider.image))",message="spec.forProvider.image is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   ClustertemplateV1Spec   `json:"spec"`
	Status ClustertemplateV1Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClustertemplateV1List contains a list of ClustertemplateV1s
type ClustertemplateV1List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClustertemplateV1 `json:"items"`
}

// Repository type metadata.
var (
	ClustertemplateV1_Kind             = "ClustertemplateV1"
	ClustertemplateV1_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ClustertemplateV1_Kind}.String()
	ClustertemplateV1_KindAPIVersion   = ClustertemplateV1_Kind + "." + CRDGroupVersion.String()
	ClustertemplateV1_GroupVersionKind = CRDGroupVersion.WithKind(ClustertemplateV1_Kind)
)

func init() {
	SchemeBuilder.Register(&ClustertemplateV1{}, &ClustertemplateV1List{})
}
