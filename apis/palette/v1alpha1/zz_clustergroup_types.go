// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ClusterGroupClusterProfileInitParameters struct {

	// (String) The ID of this resource.
	// The ID of the cluster profile.
	// +crossplane:generate:reference:type=github.com/vrabbi/provider-palette/apis/palette/v1alpha1.ClusterProfile
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Reference to a ClusterProfile in palette to populate id.
	// +kubebuilder:validation:Optional
	IDRef *v1.Reference `json:"idRef,omitempty" tf:"-"`

	// Selector for a ClusterProfile in palette to populate id.
	// +kubebuilder:validation:Optional
	IDSelector *v1.Selector `json:"idSelector,omitempty" tf:"-"`

	// (Block List) For packs of type spectro, helm, and manifest, at least one pack must be specified. (see below for nested schema)
	// For packs of type `spectro`, `helm`, and `manifest`, at least one pack must be specified.
	Pack []ClusterGroupClusterProfilePackInitParameters `json:"pack,omitempty" tf:"pack,omitempty"`

	// value pairs. For example: priority = "5".
	// A map of cluster profile variables, specified as key-value pairs. For example: `priority = "5"`.
	// +mapType=granular
	Variables map[string]*string `json:"variables,omitempty" tf:"variables,omitempty"`
}

type ClusterGroupClusterProfileObservation struct {

	// (String) The ID of this resource.
	// The ID of the cluster profile.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Block List) For packs of type spectro, helm, and manifest, at least one pack must be specified. (see below for nested schema)
	// For packs of type `spectro`, `helm`, and `manifest`, at least one pack must be specified.
	Pack []ClusterGroupClusterProfilePackObservation `json:"pack,omitempty" tf:"pack,omitempty"`

	// value pairs. For example: priority = "5".
	// A map of cluster profile variables, specified as key-value pairs. For example: `priority = "5"`.
	// +mapType=granular
	Variables map[string]*string `json:"variables,omitempty" tf:"variables,omitempty"`
}

type ClusterGroupClusterProfilePackInitParameters struct {

	// (Block List) (see below for nested schema)
	Manifest []ClusterGroupClusterProfilePackManifestInitParameters `json:"manifest,omitempty" tf:"manifest,omitempty"`

	// (String) Name of the cluster group
	// The name of the pack. The name must be unique within the cluster profile.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The registry UID of the pack. The registry UID is the unique identifier of the registry. This attribute is required if there is more than one registry that contains a pack with the same name.
	// The registry UID of the pack. The registry UID is the unique identifier of the registry. This attribute is required if there is more than one registry that contains a pack with the same name.
	// +crossplane:generate:reference:type=github.com/vrabbi/provider-palette/apis/palette/v1alpha1.OCIRegistry
	RegistryUID *string `json:"registryUid,omitempty" tf:"registry_uid,omitempty"`

	// Reference to a OCIRegistry in palette to populate registryUid.
	// +kubebuilder:validation:Optional
	RegistryUIDRef *v1.Reference `json:"registryUidRef,omitempty" tf:"-"`

	// Selector for a OCIRegistry in palette to populate registryUid.
	// +kubebuilder:validation:Optional
	RegistryUIDSelector *v1.Selector `json:"registryUidSelector,omitempty" tf:"-"`

	// (String) The tag of the pack. The tag is the version of the pack. This attribute is required if the pack type is spectro or helm.
	// The tag of the pack. The tag is the version of the pack. This attribute is required if the pack type is `spectro` or `helm`.
	Tag *string `json:"tag,omitempty" tf:"tag,omitempty"`

	// (String) The type of the pack. Allowed values are spectro, manifest, helm, or oci. The default value is spectro. If using an OCI registry for pack, set the type to oci.
	// The type of the pack. Allowed values are `spectro`, `manifest`, `helm`, or `oci`. The default value is spectro. If using an OCI registry for pack, set the type to `oci`.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// (String) The unique identifier of the pack. The value can be looked up using the spectrocloud_pack data source. This value is required if the pack type is spectro and for helm if the chart is from a public helm registry.
	// The unique identifier of the pack. The value can be looked up using the [`spectrocloud_pack`](https://registry.io/providers/spectrocloud/spectrocloud/latest/docs/data-sources/pack) data source. This value is required if the pack type is `spectro` and for `helm` if the chart is from a public helm registry.
	UID *string `json:"uid,omitempty" tf:"uid,omitempty"`

	// (String)
	// The values of the pack. The values are the configuration values of the pack. The values are specified in YAML format.
	Values *string `json:"values,omitempty" tf:"values,omitempty"`
}

type ClusterGroupClusterProfilePackManifestInitParameters struct {

	// (String) The content of the manifest. The content is the YAML content of the manifest.
	// The content of the manifest. The content is the YAML content of the manifest.
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// (String) Name of the cluster group
	// The name of the manifest. The name must be unique within the pack.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type ClusterGroupClusterProfilePackManifestObservation struct {

	// (String) The content of the manifest. The content is the YAML content of the manifest.
	// The content of the manifest. The content is the YAML content of the manifest.
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// (String) Name of the cluster group
	// The name of the manifest. The name must be unique within the pack.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The unique identifier of the pack. The value can be looked up using the spectrocloud_pack data source. This value is required if the pack type is spectro and for helm if the chart is from a public helm registry.
	UID *string `json:"uid,omitempty" tf:"uid,omitempty"`
}

type ClusterGroupClusterProfilePackManifestParameters struct {

	// (String) The content of the manifest. The content is the YAML content of the manifest.
	// The content of the manifest. The content is the YAML content of the manifest.
	// +kubebuilder:validation:Optional
	Content *string `json:"content" tf:"content,omitempty"`

	// (String) Name of the cluster group
	// The name of the manifest. The name must be unique within the pack.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`
}

type ClusterGroupClusterProfilePackObservation struct {

	// (Block List) (see below for nested schema)
	Manifest []ClusterGroupClusterProfilePackManifestObservation `json:"manifest,omitempty" tf:"manifest,omitempty"`

	// (String) Name of the cluster group
	// The name of the pack. The name must be unique within the cluster profile.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The registry UID of the pack. The registry UID is the unique identifier of the registry. This attribute is required if there is more than one registry that contains a pack with the same name.
	// The registry UID of the pack. The registry UID is the unique identifier of the registry. This attribute is required if there is more than one registry that contains a pack with the same name.
	RegistryUID *string `json:"registryUid,omitempty" tf:"registry_uid,omitempty"`

	// (String) The tag of the pack. The tag is the version of the pack. This attribute is required if the pack type is spectro or helm.
	// The tag of the pack. The tag is the version of the pack. This attribute is required if the pack type is `spectro` or `helm`.
	Tag *string `json:"tag,omitempty" tf:"tag,omitempty"`

	// (String) The type of the pack. Allowed values are spectro, manifest, helm, or oci. The default value is spectro. If using an OCI registry for pack, set the type to oci.
	// The type of the pack. Allowed values are `spectro`, `manifest`, `helm`, or `oci`. The default value is spectro. If using an OCI registry for pack, set the type to `oci`.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// (String) The unique identifier of the pack. The value can be looked up using the spectrocloud_pack data source. This value is required if the pack type is spectro and for helm if the chart is from a public helm registry.
	// The unique identifier of the pack. The value can be looked up using the [`spectrocloud_pack`](https://registry.io/providers/spectrocloud/spectrocloud/latest/docs/data-sources/pack) data source. This value is required if the pack type is `spectro` and for `helm` if the chart is from a public helm registry.
	UID *string `json:"uid,omitempty" tf:"uid,omitempty"`

	// (String)
	// The values of the pack. The values are the configuration values of the pack. The values are specified in YAML format.
	Values *string `json:"values,omitempty" tf:"values,omitempty"`
}

type ClusterGroupClusterProfilePackParameters struct {

	// (Block List) (see below for nested schema)
	// +kubebuilder:validation:Optional
	Manifest []ClusterGroupClusterProfilePackManifestParameters `json:"manifest,omitempty" tf:"manifest,omitempty"`

	// (String) Name of the cluster group
	// The name of the pack. The name must be unique within the cluster profile.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// (String) The registry UID of the pack. The registry UID is the unique identifier of the registry. This attribute is required if there is more than one registry that contains a pack with the same name.
	// The registry UID of the pack. The registry UID is the unique identifier of the registry. This attribute is required if there is more than one registry that contains a pack with the same name.
	// +crossplane:generate:reference:type=github.com/vrabbi/provider-palette/apis/palette/v1alpha1.OCIRegistry
	// +kubebuilder:validation:Optional
	RegistryUID *string `json:"registryUid,omitempty" tf:"registry_uid,omitempty"`

	// Reference to a OCIRegistry in palette to populate registryUid.
	// +kubebuilder:validation:Optional
	RegistryUIDRef *v1.Reference `json:"registryUidRef,omitempty" tf:"-"`

	// Selector for a OCIRegistry in palette to populate registryUid.
	// +kubebuilder:validation:Optional
	RegistryUIDSelector *v1.Selector `json:"registryUidSelector,omitempty" tf:"-"`

	// (String) The tag of the pack. The tag is the version of the pack. This attribute is required if the pack type is spectro or helm.
	// The tag of the pack. The tag is the version of the pack. This attribute is required if the pack type is `spectro` or `helm`.
	// +kubebuilder:validation:Optional
	Tag *string `json:"tag,omitempty" tf:"tag,omitempty"`

	// (String) The type of the pack. Allowed values are spectro, manifest, helm, or oci. The default value is spectro. If using an OCI registry for pack, set the type to oci.
	// The type of the pack. Allowed values are `spectro`, `manifest`, `helm`, or `oci`. The default value is spectro. If using an OCI registry for pack, set the type to `oci`.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// (String) The unique identifier of the pack. The value can be looked up using the spectrocloud_pack data source. This value is required if the pack type is spectro and for helm if the chart is from a public helm registry.
	// The unique identifier of the pack. The value can be looked up using the [`spectrocloud_pack`](https://registry.io/providers/spectrocloud/spectrocloud/latest/docs/data-sources/pack) data source. This value is required if the pack type is `spectro` and for `helm` if the chart is from a public helm registry.
	// +kubebuilder:validation:Optional
	UID *string `json:"uid,omitempty" tf:"uid,omitempty"`

	// (String)
	// The values of the pack. The values are the configuration values of the pack. The values are specified in YAML format.
	// +kubebuilder:validation:Optional
	Values *string `json:"values,omitempty" tf:"values,omitempty"`
}

type ClusterGroupClusterProfileParameters struct {

	// (String) The ID of this resource.
	// The ID of the cluster profile.
	// +crossplane:generate:reference:type=github.com/vrabbi/provider-palette/apis/palette/v1alpha1.ClusterProfile
	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Reference to a ClusterProfile in palette to populate id.
	// +kubebuilder:validation:Optional
	IDRef *v1.Reference `json:"idRef,omitempty" tf:"-"`

	// Selector for a ClusterProfile in palette to populate id.
	// +kubebuilder:validation:Optional
	IDSelector *v1.Selector `json:"idSelector,omitempty" tf:"-"`

	// (Block List) For packs of type spectro, helm, and manifest, at least one pack must be specified. (see below for nested schema)
	// For packs of type `spectro`, `helm`, and `manifest`, at least one pack must be specified.
	// +kubebuilder:validation:Optional
	Pack []ClusterGroupClusterProfilePackParameters `json:"pack,omitempty" tf:"pack,omitempty"`

	// value pairs. For example: priority = "5".
	// A map of cluster profile variables, specified as key-value pairs. For example: `priority = "5"`.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Variables map[string]*string `json:"variables,omitempty" tf:"variables,omitempty"`
}

type ClusterGroupConfigInitParameters struct {

	// (Number) The CPU limit in millicores.
	// The CPU limit in millicores.
	CPUMillicore *float64 `json:"cpuMillicore,omitempty" tf:"cpu_millicore,omitempty"`

	// (String) The host endpoint type. Allowed values are 'Ingress' or 'LoadBalancer'. Defaults to 'Ingress'.
	// The host endpoint type. Allowed values are 'Ingress' or 'LoadBalancer'. Defaults to 'Ingress'.
	HostEndpointType *string `json:"hostEndpointType,omitempty" tf:"host_endpoint_type,omitempty"`

	// (String) The Kubernetes distribution, allowed values are k3s and cncf_k8s.
	// The Kubernetes distribution, allowed values are `k3s` and `cncf_k8s`.
	K8SDistribution *string `json:"k8sDistribution,omitempty" tf:"k8s_distribution,omitempty"`

	// (Number) The memory limit in megabytes (MB).
	// The memory limit in megabytes (MB).
	MemoryInMb *float64 `json:"memoryInMb,omitempty" tf:"memory_in_mb,omitempty"`

	// (Number) The allowed oversubscription percentage.
	// The allowed oversubscription percentage.
	OversubscriptionPercent *float64 `json:"oversubscriptionPercent,omitempty" tf:"oversubscription_percent,omitempty"`

	// (Number) The storage limit in gigabytes (GB).
	// The storage limit in gigabytes (GB).
	StorageInGb *float64 `json:"storageInGb,omitempty" tf:"storage_in_gb,omitempty"`

	// (String)
	Values *string `json:"values,omitempty" tf:"values,omitempty"`
}

type ClusterGroupConfigObservation struct {

	// (Number) The CPU limit in millicores.
	// The CPU limit in millicores.
	CPUMillicore *float64 `json:"cpuMillicore,omitempty" tf:"cpu_millicore,omitempty"`

	// (String) The host endpoint type. Allowed values are 'Ingress' or 'LoadBalancer'. Defaults to 'Ingress'.
	// The host endpoint type. Allowed values are 'Ingress' or 'LoadBalancer'. Defaults to 'Ingress'.
	HostEndpointType *string `json:"hostEndpointType,omitempty" tf:"host_endpoint_type,omitempty"`

	// (String) The Kubernetes distribution, allowed values are k3s and cncf_k8s.
	// The Kubernetes distribution, allowed values are `k3s` and `cncf_k8s`.
	K8SDistribution *string `json:"k8sDistribution,omitempty" tf:"k8s_distribution,omitempty"`

	// (Number) The memory limit in megabytes (MB).
	// The memory limit in megabytes (MB).
	MemoryInMb *float64 `json:"memoryInMb,omitempty" tf:"memory_in_mb,omitempty"`

	// (Number) The allowed oversubscription percentage.
	// The allowed oversubscription percentage.
	OversubscriptionPercent *float64 `json:"oversubscriptionPercent,omitempty" tf:"oversubscription_percent,omitempty"`

	// (Number) The storage limit in gigabytes (GB).
	// The storage limit in gigabytes (GB).
	StorageInGb *float64 `json:"storageInGb,omitempty" tf:"storage_in_gb,omitempty"`

	// (String)
	Values *string `json:"values,omitempty" tf:"values,omitempty"`
}

type ClusterGroupConfigParameters struct {

	// (Number) The CPU limit in millicores.
	// The CPU limit in millicores.
	// +kubebuilder:validation:Optional
	CPUMillicore *float64 `json:"cpuMillicore,omitempty" tf:"cpu_millicore,omitempty"`

	// (String) The host endpoint type. Allowed values are 'Ingress' or 'LoadBalancer'. Defaults to 'Ingress'.
	// The host endpoint type. Allowed values are 'Ingress' or 'LoadBalancer'. Defaults to 'Ingress'.
	// +kubebuilder:validation:Optional
	HostEndpointType *string `json:"hostEndpointType,omitempty" tf:"host_endpoint_type,omitempty"`

	// (String) The Kubernetes distribution, allowed values are k3s and cncf_k8s.
	// The Kubernetes distribution, allowed values are `k3s` and `cncf_k8s`.
	// +kubebuilder:validation:Optional
	K8SDistribution *string `json:"k8sDistribution,omitempty" tf:"k8s_distribution,omitempty"`

	// (Number) The memory limit in megabytes (MB).
	// The memory limit in megabytes (MB).
	// +kubebuilder:validation:Optional
	MemoryInMb *float64 `json:"memoryInMb,omitempty" tf:"memory_in_mb,omitempty"`

	// (Number) The allowed oversubscription percentage.
	// The allowed oversubscription percentage.
	// +kubebuilder:validation:Optional
	OversubscriptionPercent *float64 `json:"oversubscriptionPercent,omitempty" tf:"oversubscription_percent,omitempty"`

	// (Number) The storage limit in gigabytes (GB).
	// The storage limit in gigabytes (GB).
	// +kubebuilder:validation:Optional
	StorageInGb *float64 `json:"storageInGb,omitempty" tf:"storage_in_gb,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Values *string `json:"values,omitempty" tf:"values,omitempty"`
}

type ClusterGroupInitParameters struct {

	// (Block List) (see below for nested schema)
	ClusterProfile []ClusterGroupClusterProfileInitParameters `json:"clusterProfile,omitempty" tf:"cluster_profile,omitempty"`

	// (Block List) A list of clusters to include in the cluster group. (see below for nested schema)
	// A list of clusters to include in the cluster group.
	Clusters []ClustersInitParameters `json:"clusters,omitempty" tf:"clusters,omitempty"`

	// (Block List, Min: 1, Max: 1) (see below for nested schema)
	Config []ClusterGroupConfigInitParameters `json:"config,omitempty" tf:"config,omitempty"`

	// (String) The context of the Cluster group. Allowed values are project or tenant. Defaults to tenant. If  the project context is specified, the project name will sourced from the provider configuration parameter project_name.
	// The context of the Cluster group. Allowed values are `project` or `tenant`. Defaults to `tenant`. If  the `project` context is specified, the project name will sourced from the provider configuration parameter [`project_name`](https://registry.io/providers/spectrocloud/spectrocloud/latest/docs#schema).
	Context *string `json:"context,omitempty" tf:"context,omitempty"`

	// (String) The description of the cluster. Default value is empty string.
	// The description of the cluster. Default value is empty string.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) Name of the cluster group
	// Name of the cluster group
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Set of String) A list of tags to be applied to the cluster group. Tags must be in the form of key:value.
	// A list of tags to be applied to the cluster group. Tags must be in the form of `key:value`.
	// +listType=set
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ClusterGroupObservation struct {

	// (Block List) (see below for nested schema)
	ClusterProfile []ClusterGroupClusterProfileObservation `json:"clusterProfile,omitempty" tf:"cluster_profile,omitempty"`

	// (Block List) A list of clusters to include in the cluster group. (see below for nested schema)
	// A list of clusters to include in the cluster group.
	Clusters []ClustersObservation `json:"clusters,omitempty" tf:"clusters,omitempty"`

	// (Block List, Min: 1, Max: 1) (see below for nested schema)
	Config []ClusterGroupConfigObservation `json:"config,omitempty" tf:"config,omitempty"`

	// (String) The context of the Cluster group. Allowed values are project or tenant. Defaults to tenant. If  the project context is specified, the project name will sourced from the provider configuration parameter project_name.
	// The context of the Cluster group. Allowed values are `project` or `tenant`. Defaults to `tenant`. If  the `project` context is specified, the project name will sourced from the provider configuration parameter [`project_name`](https://registry.io/providers/spectrocloud/spectrocloud/latest/docs#schema).
	Context *string `json:"context,omitempty" tf:"context,omitempty"`

	// (String) The description of the cluster. Default value is empty string.
	// The description of the cluster. Default value is empty string.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) Name of the cluster group
	// Name of the cluster group
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Set of String) A list of tags to be applied to the cluster group. Tags must be in the form of key:value.
	// A list of tags to be applied to the cluster group. Tags must be in the form of `key:value`.
	// +listType=set
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ClusterGroupParameters struct {

	// (Block List) (see below for nested schema)
	// +kubebuilder:validation:Optional
	ClusterProfile []ClusterGroupClusterProfileParameters `json:"clusterProfile,omitempty" tf:"cluster_profile,omitempty"`

	// (Block List) A list of clusters to include in the cluster group. (see below for nested schema)
	// A list of clusters to include in the cluster group.
	// +kubebuilder:validation:Optional
	Clusters []ClustersParameters `json:"clusters,omitempty" tf:"clusters,omitempty"`

	// (Block List, Min: 1, Max: 1) (see below for nested schema)
	// +kubebuilder:validation:Optional
	Config []ClusterGroupConfigParameters `json:"config,omitempty" tf:"config,omitempty"`

	// (String) The context of the Cluster group. Allowed values are project or tenant. Defaults to tenant. If  the project context is specified, the project name will sourced from the provider configuration parameter project_name.
	// The context of the Cluster group. Allowed values are `project` or `tenant`. Defaults to `tenant`. If  the `project` context is specified, the project name will sourced from the provider configuration parameter [`project_name`](https://registry.io/providers/spectrocloud/spectrocloud/latest/docs#schema).
	// +kubebuilder:validation:Optional
	Context *string `json:"context,omitempty" tf:"context,omitempty"`

	// (String) The description of the cluster. Default value is empty string.
	// The description of the cluster. Default value is empty string.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) Name of the cluster group
	// Name of the cluster group
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Set of String) A list of tags to be applied to the cluster group. Tags must be in the form of key:value.
	// A list of tags to be applied to the cluster group. Tags must be in the form of `key:value`.
	// +kubebuilder:validation:Optional
	// +listType=set
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ClustersInitParameters struct {

	// (String) The UID of the host cluster.
	// The UID of the host cluster.
	ClusterUID *string `json:"clusterUid,omitempty" tf:"cluster_uid,omitempty"`

	// (String) The host DNS wildcard for the cluster. i.e. *.dev or *test.com
	// The host DNS wildcard for the cluster. i.e. `*.dev` or `*test.com`
	HostDNS *string `json:"hostDns,omitempty" tf:"host_dns,omitempty"`
}

type ClustersObservation struct {

	// (String) The UID of the host cluster.
	// The UID of the host cluster.
	ClusterUID *string `json:"clusterUid,omitempty" tf:"cluster_uid,omitempty"`

	// (String) The host DNS wildcard for the cluster. i.e. *.dev or *test.com
	// The host DNS wildcard for the cluster. i.e. `*.dev` or `*test.com`
	HostDNS *string `json:"hostDns,omitempty" tf:"host_dns,omitempty"`
}

type ClustersParameters struct {

	// (String) The UID of the host cluster.
	// The UID of the host cluster.
	// +kubebuilder:validation:Optional
	ClusterUID *string `json:"clusterUid" tf:"cluster_uid,omitempty"`

	// (String) The host DNS wildcard for the cluster. i.e. *.dev or *test.com
	// The host DNS wildcard for the cluster. i.e. `*.dev` or `*test.com`
	// +kubebuilder:validation:Optional
	HostDNS *string `json:"hostDns,omitempty" tf:"host_dns,omitempty"`
}

// ClusterGroupSpec defines the desired state of ClusterGroup
type ClusterGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ClusterGroupParameters `json:"forProvider"`
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
	InitProvider ClusterGroupInitParameters `json:"initProvider,omitempty"`
}

// ClusterGroupStatus defines the observed state of ClusterGroup.
type ClusterGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ClusterGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ClusterGroup is the Schema for the ClusterGroups API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,palette}
type ClusterGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.config) || (has(self.initProvider) && has(self.initProvider.config))",message="spec.forProvider.config is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   ClusterGroupSpec   `json:"spec"`
	Status ClusterGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterGroupList contains a list of ClusterGroups
type ClusterGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterGroup `json:"items"`
}

// Repository type metadata.
var (
	ClusterGroup_Kind             = "ClusterGroup"
	ClusterGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ClusterGroup_Kind}.String()
	ClusterGroup_KindAPIVersion   = ClusterGroup_Kind + "." + CRDGroupVersion.String()
	ClusterGroup_GroupVersionKind = CRDGroupVersion.WithKind(ClusterGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&ClusterGroup{}, &ClusterGroupList{})
}
