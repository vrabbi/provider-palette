/*
Copyright 2022 Upbound Inc.
*/

package config

import "github.com/crossplane/upjet/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider (24).
var ExternalNameConfigs = map[string]config.ExternalName{
	// TENANCY
	"spectrocloud_project": config.IdentifierFromProvider,
	// SECURITY
	"spectrocloud_ssh_key": config.IdentifierFromProvider,
	// REGISTRY
	"spectrocloud_registry_helm": config.IdentifierFromProvider,
	"spectrocloud_registry_oci":  config.IdentifierFromProvider,
	// RBAC
	"spectrocloud_user": config.IdentifierFromProvider,
	"spectrocloud_team": config.IdentifierFromProvider,
	"spectrocloud_role": config.IdentifierFromProvider,
	// CLOUD ACCOUNTS
	"spectrocloud_cloudaccount_aws":     config.IdentifierFromProvider,
	"spectrocloud_cloudaccount_azure":   config.IdentifierFromProvider,
	"spectrocloud_cloudaccount_custom":  config.IdentifierFromProvider,
	"spectrocloud_cloudaccount_gcp":     config.IdentifierFromProvider,
	"spectrocloud_cloudaccount_vsphere": config.IdentifierFromProvider,
	// PROFILES
	"spectrocloud_cluster_profile": config.TemplatedStringAsIdentifier("id", "{{ .external_name }}:{{ .parameters.context }}"),
	// CLUSTERS
	"spectrocloud_cluster_vsphere": config.IdentifierFromProvider,
	"spectrocloud_cluster_eks":     config.IdentifierFromProvider,
	// MACROS
	"spectrocloud_macros": config.IdentifierFromProvider,
	// SETTINGS
	"spectrocloud_developer_setting": config.IdentifierFromProvider,
	"spectrocloud_platform_setting":  config.IdentifierFromProvider,
	"spectrocloud_filter":            config.IdentifierFromProvider,
	// IPAM + DNS
	"spectrocloud_privatecloudgateway_dns_map": config.IdentifierFromProvider,
	"spectrocloud_privatecloudgateway_ippool":  config.IdentifierFromProvider,
	"spectrocloud_backup_storage_location":     config.IdentifierFromProvider,
	// POLICIES
	"spectrocloud_password_policy": config.IdentifierFromProvider,
	"spectrocloud_resource_limit":  config.IdentifierFromProvider,

	/* BELLOW RESOURCES (22) ARE NOT IMPLEMENTED YET

	// HIGH PRIORITY (7)
	"spectrocloud_cluster_custom_cloud": config.IdentifierFromProvider,
	"spectrocloud_cluster_aks": config.IdentifierFromProvider,
	"spectrocloud_cluster_aws": config.IdentifierFromProvider,
	"spectrocloud_cluster_azure": config.IdentifierFromProvider,
	"spectrocloud_cluster_group": config.IdentifierFromProvider,
	"spectrocloud_virtual_cluster": config.IdentifierFromProvider,
	"spectrocloud_workspace": config.IdentifierFromProvider,

	// MEDIUM PRIORITY (4)
	"spectrocloud_cloudaccount_maas": config.IdentifierFromProvider,
	"spectrocloud_cluster_gcp": config.IdentifierFromProvider,
	"spectrocloud_cluster_gke": config.IdentifierFromProvider,
	"spectrocloud_cluster_maas": config.IdentifierFromProvider,

	// LOW PRIORITY (11)
	"spectrocloud_cloudaccount_openstack": config.IdentifierFromProvider,
	"spectrocloud_cluster_edge_native": config.IdentifierFromProvider,
	"spectrocloud_cluster_edge_vsphere": config.IdentifierFromProvider,
	"spectrocloud_cluster_openstack": config.IdentifierFromProvider,
	"spectrocloud_datavolume": config.IdentifierFromProvider,
	"spectrocloud_registration_token": config.IdentifierFromProvider,
	"spectrocloud_virtual_machine": config.IdentifierFromProvider,
	"spectrocloud_addon_deployment": config.IdentifierFromProvider,
	"spectrocloud_appliance": config.IdentifierFromProvider,
	"spectrocloud_application": config.IdentifierFromProvider,
	"spectrocloud_application_profile": config.IdentifierFromProvider,
	*/
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
