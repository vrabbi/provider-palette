/*
Copyright 2022 Upbound Inc.
*/

package config

import "github.com/crossplane/upjet/pkg/config"

// ExternalNameConfigs contains all external name configurations for this provider
// Currently 37 out of 45 resources are supported.
var ExternalNameConfigs = map[string]config.ExternalName{
	// TENANCY (3)
	"spectrocloud_project":       config.IdentifierFromProvider,
	"spectrocloud_cluster_group": config.IdentifierFromProvider,
	"spectrocloud_workspace":     config.IdentifierFromProvider,
	// SECURITY (1)
	"spectrocloud_ssh_key": config.IdentifierFromProvider,
	// REGISTRY (2)
	"spectrocloud_registry_helm": config.IdentifierFromProvider,
	"spectrocloud_registry_oci":  config.IdentifierFromProvider,
	// RBAC (3)
	"spectrocloud_user": config.IdentifierFromProvider,
	"spectrocloud_team": config.IdentifierFromProvider,
	"spectrocloud_role": config.IdentifierFromProvider,
	// CLOUD ACCOUNTS (6)
	"spectrocloud_cloudaccount_aws":     config.IdentifierFromProvider,
	"spectrocloud_cloudaccount_azure":   config.IdentifierFromProvider,
	"spectrocloud_cloudaccount_custom":  config.IdentifierFromProvider,
	"spectrocloud_cloudaccount_gcp":     config.IdentifierFromProvider,
	"spectrocloud_cloudaccount_vsphere": config.IdentifierFromProvider,
	"spectrocloud_cloudaccount_maas":    config.IdentifierFromProvider,
	// PROFILES (2)
	"spectrocloud_cluster_profile":     config.TemplatedStringAsIdentifier("id", "{{ .external_name }}:{{ .parameters.context }}"),
	"spectrocloud_application_profile": config.IdentifierFromProvider,
	// CLUSTERS (9)
	"spectrocloud_cluster_vsphere":      config.IdentifierFromProvider,
	"spectrocloud_cluster_eks":          config.IdentifierFromProvider,
	"spectrocloud_cluster_aws":          config.IdentifierFromProvider,
	"spectrocloud_virtual_cluster":      config.IdentifierFromProvider,
	"spectrocloud_cluster_aks":          config.IdentifierFromProvider,
	"spectrocloud_cluster_azure":        config.IdentifierFromProvider,
	"spectrocloud_cluster_gke":          config.IdentifierFromProvider,
	"spectrocloud_cluster_gcp":          config.IdentifierFromProvider,
	"spectrocloud_cluster_custom_cloud": config.IdentifierFromProvider,
	// MACROS (1)
	"spectrocloud_macros": config.IdentifierFromProvider,
	// SETTINGS (3)
	"spectrocloud_developer_setting": config.IdentifierFromProvider,
	"spectrocloud_platform_setting":  config.IdentifierFromProvider,
	"spectrocloud_filter":            config.IdentifierFromProvider,
	// IPAM + DNS (2)
	"spectrocloud_privatecloudgateway_dns_map": config.IdentifierFromProvider,
	"spectrocloud_privatecloudgateway_ippool":  config.IdentifierFromProvider,
	// Backups (1)
	"spectrocloud_backup_storage_location": config.IdentifierFromProvider,
	// POLICIES (2)
	"spectrocloud_password_policy": config.IdentifierFromProvider,
	"spectrocloud_resource_limit":  config.IdentifierFromProvider,
	// Edge Management (1)
	"spectrocloud_registration_token": config.IdentifierFromProvider,
	"spectrocloud_appliance":          config.IdentifierFromProvider,
	// App Mode (1)
	"spectrocloud_application": config.IdentifierFromProvider,

	/* 8 RESOURCES ARE NOT IMPLEMENTED YET
	// MAAS Support
	"spectrocloud_cluster_maas": config.IdentifierFromProvider,
	// OpenStack Support
	"spectrocloud_cloudaccount_openstack": config.IdentifierFromProvider,
	"spectrocloud_cluster_openstack": config.IdentifierFromProvider,
	// VMO Support
	"spectrocloud_datavolume": config.IdentifierFromProvider,
	"spectrocloud_virtual_machine": config.IdentifierFromProvider,
	// Edge Support

	"spectrocloud_cluster_edge_native": config.IdentifierFromProvider,
	"spectrocloud_cluster_edge_vsphere": config.IdentifierFromProvider,
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
