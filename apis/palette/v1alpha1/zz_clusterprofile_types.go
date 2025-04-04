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

type ClusterProfileInitParameters_2 struct {

	// all, aws, azure, gcp, vsphere, openstack, maas, virtual, baremetal, eks, aks, edge-native, generic, and gke or any custom cloud provider registered in Palette, e.g., nutanix.If the value is set to all, then the type must be set to add-on. Otherwise, the cluster profile may be incompatible with other providers. Default value is all.
	// Specify the infrastructure provider the cluster profile is for. Only Palette supported infrastructure providers can be used. The supported cloud types are - `all, aws, azure, gcp, vsphere, openstack, maas, virtual, baremetal, eks, aks, edge-native, generic, and gke` or any custom cloud provider registered in Palette, e.g., `nutanix`.If the value is set to `all`, then the type must be set to `add-on`. Otherwise, the cluster profile may be incompatible with other providers. Default value is `all`.
	Cloud *string `json:"cloud,omitempty" tf:"cloud,omitempty"`

	// (String) The context of the cluster profile. Allowed values are project or tenant. Default value is project. If  the project context is specified, the project name will sourced from the provider configuration parameter project_name.
	// The context of the cluster profile. Allowed values are `project` or `tenant`. Default value is `project`. If  the `project` context is specified, the project name will sourced from the provider configuration parameter [`project_name`](https://registry.io/providers/spectrocloud/spectrocloud/latest/docs#schema).
	Context *string `json:"context,omitempty" tf:"context,omitempty"`

	// (String)
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String)
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Block List) For packs of type spectro, helm, and manifest, at least one pack must be specified. (see below for nested schema)
	// For packs of type `spectro`, `helm`, and `manifest`, at least one pack must be specified.
	Pack []ClusterProfilePackInitParameters_2 `json:"pack,omitempty" tf:"pack,omitempty"`

	// on profiles. (see below for nested schema)
	// List of variables for the cluster profile. Note: This is a preview feature and is currently only supported for the `edge_native` cloud type and general `add-on` profiles.
	ProfileVariables []ProfileVariablesInitParameters `json:"profileVariables,omitempty" tf:"profile_variables,omitempty"`

	// (Set of String) A list of tags to be applied to the cluster. Tags must be in the form of key:value.
	// A list of tags to be applied to the cluster. Tags must be in the form of `key:value`.
	// +listType=set
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// on, and system. These values map to the following User Interface (UI) labels. Use the value ' cluster ' for a Full cluster profile.For an Infrastructure cluster profile, use the value infra; for an Add-on cluster profile, use the value add-on.System cluster profiles can be specified using the value system. To learn more about cluster profiles, refer to the Cluster Profile documentation. Default value is add-on.
	// Specify the cluster profile type to use. Allowed values are `cluster`, `infra`, `add-on`, and `system`. These values map to the following User Interface (UI) labels. Use the value ' cluster ' for a **Full** cluster profile.For an Infrastructure cluster profile, use the value `infra`; for an Add-on cluster profile, use the value `add-on`.System cluster profiles can be specified using the value `system`. To learn more about cluster profiles, refer to the [Cluster Profile](https://docs.spectrocloud.com/cluster-profiles) documentation. Default value is `add-on`.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// (String) Version of the cluster profile. Defaults to '1.0.0'.
	// Version of the cluster profile. Defaults to '1.0.0'.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type ClusterProfileObservation_2 struct {

	// all, aws, azure, gcp, vsphere, openstack, maas, virtual, baremetal, eks, aks, edge-native, generic, and gke or any custom cloud provider registered in Palette, e.g., nutanix.If the value is set to all, then the type must be set to add-on. Otherwise, the cluster profile may be incompatible with other providers. Default value is all.
	// Specify the infrastructure provider the cluster profile is for. Only Palette supported infrastructure providers can be used. The supported cloud types are - `all, aws, azure, gcp, vsphere, openstack, maas, virtual, baremetal, eks, aks, edge-native, generic, and gke` or any custom cloud provider registered in Palette, e.g., `nutanix`.If the value is set to `all`, then the type must be set to `add-on`. Otherwise, the cluster profile may be incompatible with other providers. Default value is `all`.
	Cloud *string `json:"cloud,omitempty" tf:"cloud,omitempty"`

	// (String) The context of the cluster profile. Allowed values are project or tenant. Default value is project. If  the project context is specified, the project name will sourced from the provider configuration parameter project_name.
	// The context of the cluster profile. Allowed values are `project` or `tenant`. Default value is `project`. If  the `project` context is specified, the project name will sourced from the provider configuration parameter [`project_name`](https://registry.io/providers/spectrocloud/spectrocloud/latest/docs#schema).
	Context *string `json:"context,omitempty" tf:"context,omitempty"`

	// (String)
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String)
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Block List) For packs of type spectro, helm, and manifest, at least one pack must be specified. (see below for nested schema)
	// For packs of type `spectro`, `helm`, and `manifest`, at least one pack must be specified.
	Pack []ClusterProfilePackObservation_2 `json:"pack,omitempty" tf:"pack,omitempty"`

	// on profiles. (see below for nested schema)
	// List of variables for the cluster profile. Note: This is a preview feature and is currently only supported for the `edge_native` cloud type and general `add-on` profiles.
	ProfileVariables []ProfileVariablesObservation `json:"profileVariables,omitempty" tf:"profile_variables,omitempty"`

	// (Set of String) A list of tags to be applied to the cluster. Tags must be in the form of key:value.
	// A list of tags to be applied to the cluster. Tags must be in the form of `key:value`.
	// +listType=set
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// on, and system. These values map to the following User Interface (UI) labels. Use the value ' cluster ' for a Full cluster profile.For an Infrastructure cluster profile, use the value infra; for an Add-on cluster profile, use the value add-on.System cluster profiles can be specified using the value system. To learn more about cluster profiles, refer to the Cluster Profile documentation. Default value is add-on.
	// Specify the cluster profile type to use. Allowed values are `cluster`, `infra`, `add-on`, and `system`. These values map to the following User Interface (UI) labels. Use the value ' cluster ' for a **Full** cluster profile.For an Infrastructure cluster profile, use the value `infra`; for an Add-on cluster profile, use the value `add-on`.System cluster profiles can be specified using the value `system`. To learn more about cluster profiles, refer to the [Cluster Profile](https://docs.spectrocloud.com/cluster-profiles) documentation. Default value is `add-on`.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// (String) Version of the cluster profile. Defaults to '1.0.0'.
	// Version of the cluster profile. Defaults to '1.0.0'.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type ClusterProfilePackInitParameters_2 struct {

	// (Block List) (see below for nested schema)
	Manifest []ClusterProfilePackManifestInitParameters_2 `json:"manifest,omitempty" tf:"manifest,omitempty"`

	// (String)
	// The name of the pack. The name must be unique within the cluster profile.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The registry UID of the pack. The registry UID is the unique identifier of the registry. This attribute is required if there is more than one registry that contains a pack with the same name.
	// The registry UID of the pack. The registry UID is the unique identifier of the registry. This attribute is required if there is more than one registry that contains a pack with the same name.
	RegistryUID *string `json:"registryUid,omitempty" tf:"registry_uid,omitempty"`

	// (String) The tag of the pack. The tag is the version of the pack. This attribute is required if the pack type is spectro or helm.
	// The tag of the pack. The tag is the version of the pack. This attribute is required if the pack type is `spectro` or `helm`.
	Tag *string `json:"tag,omitempty" tf:"tag,omitempty"`

	// on, and system. These values map to the following User Interface (UI) labels. Use the value ' cluster ' for a Full cluster profile.For an Infrastructure cluster profile, use the value infra; for an Add-on cluster profile, use the value add-on.System cluster profiles can be specified using the value system. To learn more about cluster profiles, refer to the Cluster Profile documentation. Default value is add-on.
	// The type of the pack. Allowed values are `spectro`, `manifest`, `helm`, or `oci`. The default value is spectro. If using an OCI registry for pack, set the type to `oci`.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// (String) The unique identifier of the pack. The value can be looked up using the spectrocloud_pack data source. This value is required if the pack type is spectro and for helm if the chart is from a public helm registry.
	// The unique identifier of the pack. The value can be looked up using the [`spectrocloud_pack`](https://registry.io/providers/spectrocloud/spectrocloud/latest/docs/data-sources/pack) data source. This value is required if the pack type is `spectro` and for `helm` if the chart is from a public helm registry.
	UID *string `json:"uid,omitempty" tf:"uid,omitempty"`

	// (String) The values of the pack. The values are the configuration values of the pack. The values are specified in YAML format.
	// The values of the pack. The values are the configuration values of the pack. The values are specified in YAML format.
	Values *string `json:"values,omitempty" tf:"values,omitempty"`
}

type ClusterProfilePackManifestInitParameters_2 struct {

	// (String) The content of the manifest. The content is the YAML content of the manifest.
	// The content of the manifest. The content is the YAML content of the manifest.
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// (String)
	// The name of the manifest. The name must be unique within the pack.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type ClusterProfilePackManifestObservation_2 struct {

	// (String) The content of the manifest. The content is the YAML content of the manifest.
	// The content of the manifest. The content is the YAML content of the manifest.
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// (String)
	// The name of the manifest. The name must be unique within the pack.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The unique identifier of the pack. The value can be looked up using the spectrocloud_pack data source. This value is required if the pack type is spectro and for helm if the chart is from a public helm registry.
	UID *string `json:"uid,omitempty" tf:"uid,omitempty"`
}

type ClusterProfilePackManifestParameters_2 struct {

	// (String) The content of the manifest. The content is the YAML content of the manifest.
	// The content of the manifest. The content is the YAML content of the manifest.
	// +kubebuilder:validation:Optional
	Content *string `json:"content" tf:"content,omitempty"`

	// (String)
	// The name of the manifest. The name must be unique within the pack.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`
}

type ClusterProfilePackObservation_2 struct {

	// (Block List) (see below for nested schema)
	Manifest []ClusterProfilePackManifestObservation_2 `json:"manifest,omitempty" tf:"manifest,omitempty"`

	// (String)
	// The name of the pack. The name must be unique within the cluster profile.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The registry UID of the pack. The registry UID is the unique identifier of the registry. This attribute is required if there is more than one registry that contains a pack with the same name.
	// The registry UID of the pack. The registry UID is the unique identifier of the registry. This attribute is required if there is more than one registry that contains a pack with the same name.
	RegistryUID *string `json:"registryUid,omitempty" tf:"registry_uid,omitempty"`

	// (String) The tag of the pack. The tag is the version of the pack. This attribute is required if the pack type is spectro or helm.
	// The tag of the pack. The tag is the version of the pack. This attribute is required if the pack type is `spectro` or `helm`.
	Tag *string `json:"tag,omitempty" tf:"tag,omitempty"`

	// on, and system. These values map to the following User Interface (UI) labels. Use the value ' cluster ' for a Full cluster profile.For an Infrastructure cluster profile, use the value infra; for an Add-on cluster profile, use the value add-on.System cluster profiles can be specified using the value system. To learn more about cluster profiles, refer to the Cluster Profile documentation. Default value is add-on.
	// The type of the pack. Allowed values are `spectro`, `manifest`, `helm`, or `oci`. The default value is spectro. If using an OCI registry for pack, set the type to `oci`.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// (String) The unique identifier of the pack. The value can be looked up using the spectrocloud_pack data source. This value is required if the pack type is spectro and for helm if the chart is from a public helm registry.
	// The unique identifier of the pack. The value can be looked up using the [`spectrocloud_pack`](https://registry.io/providers/spectrocloud/spectrocloud/latest/docs/data-sources/pack) data source. This value is required if the pack type is `spectro` and for `helm` if the chart is from a public helm registry.
	UID *string `json:"uid,omitempty" tf:"uid,omitempty"`

	// (String) The values of the pack. The values are the configuration values of the pack. The values are specified in YAML format.
	// The values of the pack. The values are the configuration values of the pack. The values are specified in YAML format.
	Values *string `json:"values,omitempty" tf:"values,omitempty"`
}

type ClusterProfilePackParameters_2 struct {

	// (Block List) (see below for nested schema)
	// +kubebuilder:validation:Optional
	Manifest []ClusterProfilePackManifestParameters_2 `json:"manifest,omitempty" tf:"manifest,omitempty"`

	// (String)
	// The name of the pack. The name must be unique within the cluster profile.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// (String) The registry UID of the pack. The registry UID is the unique identifier of the registry. This attribute is required if there is more than one registry that contains a pack with the same name.
	// The registry UID of the pack. The registry UID is the unique identifier of the registry. This attribute is required if there is more than one registry that contains a pack with the same name.
	// +kubebuilder:validation:Optional
	RegistryUID *string `json:"registryUid,omitempty" tf:"registry_uid,omitempty"`

	// (String) The tag of the pack. The tag is the version of the pack. This attribute is required if the pack type is spectro or helm.
	// The tag of the pack. The tag is the version of the pack. This attribute is required if the pack type is `spectro` or `helm`.
	// +kubebuilder:validation:Optional
	Tag *string `json:"tag,omitempty" tf:"tag,omitempty"`

	// on, and system. These values map to the following User Interface (UI) labels. Use the value ' cluster ' for a Full cluster profile.For an Infrastructure cluster profile, use the value infra; for an Add-on cluster profile, use the value add-on.System cluster profiles can be specified using the value system. To learn more about cluster profiles, refer to the Cluster Profile documentation. Default value is add-on.
	// The type of the pack. Allowed values are `spectro`, `manifest`, `helm`, or `oci`. The default value is spectro. If using an OCI registry for pack, set the type to `oci`.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// (String) The unique identifier of the pack. The value can be looked up using the spectrocloud_pack data source. This value is required if the pack type is spectro and for helm if the chart is from a public helm registry.
	// The unique identifier of the pack. The value can be looked up using the [`spectrocloud_pack`](https://registry.io/providers/spectrocloud/spectrocloud/latest/docs/data-sources/pack) data source. This value is required if the pack type is `spectro` and for `helm` if the chart is from a public helm registry.
	// +kubebuilder:validation:Optional
	UID *string `json:"uid,omitempty" tf:"uid,omitempty"`

	// (String) The values of the pack. The values are the configuration values of the pack. The values are specified in YAML format.
	// The values of the pack. The values are the configuration values of the pack. The values are specified in YAML format.
	// +kubebuilder:validation:Optional
	Values *string `json:"values,omitempty" tf:"values,omitempty"`
}

type ClusterProfileParameters_2 struct {

	// all, aws, azure, gcp, vsphere, openstack, maas, virtual, baremetal, eks, aks, edge-native, generic, and gke or any custom cloud provider registered in Palette, e.g., nutanix.If the value is set to all, then the type must be set to add-on. Otherwise, the cluster profile may be incompatible with other providers. Default value is all.
	// Specify the infrastructure provider the cluster profile is for. Only Palette supported infrastructure providers can be used. The supported cloud types are - `all, aws, azure, gcp, vsphere, openstack, maas, virtual, baremetal, eks, aks, edge-native, generic, and gke` or any custom cloud provider registered in Palette, e.g., `nutanix`.If the value is set to `all`, then the type must be set to `add-on`. Otherwise, the cluster profile may be incompatible with other providers. Default value is `all`.
	// +kubebuilder:validation:Optional
	Cloud *string `json:"cloud,omitempty" tf:"cloud,omitempty"`

	// (String) The context of the cluster profile. Allowed values are project or tenant. Default value is project. If  the project context is specified, the project name will sourced from the provider configuration parameter project_name.
	// The context of the cluster profile. Allowed values are `project` or `tenant`. Default value is `project`. If  the `project` context is specified, the project name will sourced from the provider configuration parameter [`project_name`](https://registry.io/providers/spectrocloud/spectrocloud/latest/docs#schema).
	// +kubebuilder:validation:Optional
	Context *string `json:"context,omitempty" tf:"context,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Block List) For packs of type spectro, helm, and manifest, at least one pack must be specified. (see below for nested schema)
	// For packs of type `spectro`, `helm`, and `manifest`, at least one pack must be specified.
	// +kubebuilder:validation:Optional
	Pack []ClusterProfilePackParameters_2 `json:"pack,omitempty" tf:"pack,omitempty"`

	// on profiles. (see below for nested schema)
	// List of variables for the cluster profile. Note: This is a preview feature and is currently only supported for the `edge_native` cloud type and general `add-on` profiles.
	// +kubebuilder:validation:Optional
	ProfileVariables []ProfileVariablesParameters `json:"profileVariables,omitempty" tf:"profile_variables,omitempty"`

	// (Set of String) A list of tags to be applied to the cluster. Tags must be in the form of key:value.
	// A list of tags to be applied to the cluster. Tags must be in the form of `key:value`.
	// +kubebuilder:validation:Optional
	// +listType=set
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// on, and system. These values map to the following User Interface (UI) labels. Use the value ' cluster ' for a Full cluster profile.For an Infrastructure cluster profile, use the value infra; for an Add-on cluster profile, use the value add-on.System cluster profiles can be specified using the value system. To learn more about cluster profiles, refer to the Cluster Profile documentation. Default value is add-on.
	// Specify the cluster profile type to use. Allowed values are `cluster`, `infra`, `add-on`, and `system`. These values map to the following User Interface (UI) labels. Use the value ' cluster ' for a **Full** cluster profile.For an Infrastructure cluster profile, use the value `infra`; for an Add-on cluster profile, use the value `add-on`.System cluster profiles can be specified using the value `system`. To learn more about cluster profiles, refer to the [Cluster Profile](https://docs.spectrocloud.com/cluster-profiles) documentation. Default value is `add-on`.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// (String) Version of the cluster profile. Defaults to '1.0.0'.
	// Version of the cluster profile. Defaults to '1.0.0'.
	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type ProfileVariablesInitParameters struct {

	// (Block List, Min: 1) (see below for nested schema)
	Variable []VariableInitParameters `json:"variable,omitempty" tf:"variable,omitempty"`
}

type ProfileVariablesObservation struct {

	// (Block List, Min: 1) (see below for nested schema)
	Variable []VariableObservation `json:"variable,omitempty" tf:"variable,omitempty"`
}

type ProfileVariablesParameters struct {

	// (Block List, Min: 1) (see below for nested schema)
	// +kubebuilder:validation:Optional
	Variable []VariableParameters `json:"variable" tf:"variable,omitempty"`
}

type VariableInitParameters struct {

	// (String) The default value of the variable.
	// The default value of the variable.
	DefaultValue *string `json:"defaultValue,omitempty" tf:"default_value,omitempty"`

	// (String)
	// The description of the variable.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) The display name of the variable should be unique among variables.
	// The display name of the variable should be unique among variables.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// (String) The format of the variable. Default is string, format field can only be set during cluster profile creation. Allowed formats include string, number, boolean, ipv4, ipv4cidr, ipv6, version.
	// The format of the variable. Default is `string`, `format` field can only be set during cluster profile creation. Allowed formats include `string`, `number`, `boolean`, `ipv4`, `ipv4cidr`, `ipv6`, `version`.
	Format *string `json:"format,omitempty" tf:"format,omitempty"`

	// (Boolean) If hidden is set to true, then variable will be hidden for overriding the value. By default the hidden flag will be set to false.
	// If `hidden` is set to `true`, then variable will be hidden for overriding the value. By default the `hidden` flag will be set to `false`.
	Hidden *bool `json:"hidden,omitempty" tf:"hidden,omitempty"`

	// (Boolean) If immutable is set to true, then variable value can't be editable. By default the immutable flag will be set to false.
	// If `immutable` is set to `true`, then variable value can't be editable. By default the `immutable` flag will be set to `false`.
	Immutable *bool `json:"immutable,omitempty" tf:"immutable,omitempty"`

	// (Boolean) If is_sensitive is set to true, then default value will be masked. By default the is_sensitive flag will be set to false.
	// If `is_sensitive` is set to `true`, then default value will be masked. By default the `is_sensitive` flag will be set to false.
	IsSensitive *bool `json:"isSensitive,omitempty" tf:"is_sensitive,omitempty"`

	// (String)
	// The name of the variable should be unique among variables.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) Regular expression pattern which the variable value must match.
	// Regular expression pattern which the variable value must match.
	Regex *string `json:"regex,omitempty" tf:"regex,omitempty"`

	// (Boolean) The required to specify if the variable is optional or mandatory. If it is mandatory then default value must be provided.
	// The `required` to specify if the variable is optional or mandatory. If it is mandatory then default value must be provided.
	Required *bool `json:"required,omitempty" tf:"required,omitempty"`
}

type VariableObservation struct {

	// (String) The default value of the variable.
	// The default value of the variable.
	DefaultValue *string `json:"defaultValue,omitempty" tf:"default_value,omitempty"`

	// (String)
	// The description of the variable.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) The display name of the variable should be unique among variables.
	// The display name of the variable should be unique among variables.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// (String) The format of the variable. Default is string, format field can only be set during cluster profile creation. Allowed formats include string, number, boolean, ipv4, ipv4cidr, ipv6, version.
	// The format of the variable. Default is `string`, `format` field can only be set during cluster profile creation. Allowed formats include `string`, `number`, `boolean`, `ipv4`, `ipv4cidr`, `ipv6`, `version`.
	Format *string `json:"format,omitempty" tf:"format,omitempty"`

	// (Boolean) If hidden is set to true, then variable will be hidden for overriding the value. By default the hidden flag will be set to false.
	// If `hidden` is set to `true`, then variable will be hidden for overriding the value. By default the `hidden` flag will be set to `false`.
	Hidden *bool `json:"hidden,omitempty" tf:"hidden,omitempty"`

	// (Boolean) If immutable is set to true, then variable value can't be editable. By default the immutable flag will be set to false.
	// If `immutable` is set to `true`, then variable value can't be editable. By default the `immutable` flag will be set to `false`.
	Immutable *bool `json:"immutable,omitempty" tf:"immutable,omitempty"`

	// (Boolean) If is_sensitive is set to true, then default value will be masked. By default the is_sensitive flag will be set to false.
	// If `is_sensitive` is set to `true`, then default value will be masked. By default the `is_sensitive` flag will be set to false.
	IsSensitive *bool `json:"isSensitive,omitempty" tf:"is_sensitive,omitempty"`

	// (String)
	// The name of the variable should be unique among variables.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) Regular expression pattern which the variable value must match.
	// Regular expression pattern which the variable value must match.
	Regex *string `json:"regex,omitempty" tf:"regex,omitempty"`

	// (Boolean) The required to specify if the variable is optional or mandatory. If it is mandatory then default value must be provided.
	// The `required` to specify if the variable is optional or mandatory. If it is mandatory then default value must be provided.
	Required *bool `json:"required,omitempty" tf:"required,omitempty"`
}

type VariableParameters struct {

	// (String) The default value of the variable.
	// The default value of the variable.
	// +kubebuilder:validation:Optional
	DefaultValue *string `json:"defaultValue,omitempty" tf:"default_value,omitempty"`

	// (String)
	// The description of the variable.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) The display name of the variable should be unique among variables.
	// The display name of the variable should be unique among variables.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName" tf:"display_name,omitempty"`

	// (String) The format of the variable. Default is string, format field can only be set during cluster profile creation. Allowed formats include string, number, boolean, ipv4, ipv4cidr, ipv6, version.
	// The format of the variable. Default is `string`, `format` field can only be set during cluster profile creation. Allowed formats include `string`, `number`, `boolean`, `ipv4`, `ipv4cidr`, `ipv6`, `version`.
	// +kubebuilder:validation:Optional
	Format *string `json:"format,omitempty" tf:"format,omitempty"`

	// (Boolean) If hidden is set to true, then variable will be hidden for overriding the value. By default the hidden flag will be set to false.
	// If `hidden` is set to `true`, then variable will be hidden for overriding the value. By default the `hidden` flag will be set to `false`.
	// +kubebuilder:validation:Optional
	Hidden *bool `json:"hidden,omitempty" tf:"hidden,omitempty"`

	// (Boolean) If immutable is set to true, then variable value can't be editable. By default the immutable flag will be set to false.
	// If `immutable` is set to `true`, then variable value can't be editable. By default the `immutable` flag will be set to `false`.
	// +kubebuilder:validation:Optional
	Immutable *bool `json:"immutable,omitempty" tf:"immutable,omitempty"`

	// (Boolean) If is_sensitive is set to true, then default value will be masked. By default the is_sensitive flag will be set to false.
	// If `is_sensitive` is set to `true`, then default value will be masked. By default the `is_sensitive` flag will be set to false.
	// +kubebuilder:validation:Optional
	IsSensitive *bool `json:"isSensitive,omitempty" tf:"is_sensitive,omitempty"`

	// (String)
	// The name of the variable should be unique among variables.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// (String) Regular expression pattern which the variable value must match.
	// Regular expression pattern which the variable value must match.
	// +kubebuilder:validation:Optional
	Regex *string `json:"regex,omitempty" tf:"regex,omitempty"`

	// (Boolean) The required to specify if the variable is optional or mandatory. If it is mandatory then default value must be provided.
	// The `required` to specify if the variable is optional or mandatory. If it is mandatory then default value must be provided.
	// +kubebuilder:validation:Optional
	Required *bool `json:"required,omitempty" tf:"required,omitempty"`
}

// ClusterProfileSpec defines the desired state of ClusterProfile
type ClusterProfileSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ClusterProfileParameters_2 `json:"forProvider"`
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
	InitProvider ClusterProfileInitParameters_2 `json:"initProvider,omitempty"`
}

// ClusterProfileStatus defines the observed state of ClusterProfile.
type ClusterProfileStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ClusterProfileObservation_2 `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ClusterProfile is the Schema for the ClusterProfiles API. The Cluster Profile resource allows you to create and manage cluster profiles.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,palette}
type ClusterProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   ClusterProfileSpec   `json:"spec"`
	Status ClusterProfileStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterProfileList contains a list of ClusterProfiles
type ClusterProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterProfile `json:"items"`
}

// Repository type metadata.
var (
	ClusterProfile_Kind             = "ClusterProfile"
	ClusterProfile_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ClusterProfile_Kind}.String()
	ClusterProfile_KindAPIVersion   = ClusterProfile_Kind + "." + CRDGroupVersion.String()
	ClusterProfile_GroupVersionKind = CRDGroupVersion.WithKind(ClusterProfile_Kind)
)

func init() {
	SchemeBuilder.Register(&ClusterProfile{}, &ClusterProfileList{})
}
