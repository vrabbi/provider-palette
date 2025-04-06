package resources

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	// Backups
	p.AddResourceConfigurator("spectrocloud_backup_storage_location", func(r *config.Resource) {
		r.ShortGroup = "palette"
		r.Kind = "BackupStorageLocation"
	})

	//Cloud Accounts
	p.AddResourceConfigurator("spectrocloud_cloudaccount_custom", func(r *config.Resource) {
		r.ShortGroup = "palette"
		r.Kind = "CustomCloudAccount"
	})
	p.AddResourceConfigurator("spectrocloud_cloudaccount_maas", func(r *config.Resource) {
		r.ShortGroup = "palette"
		r.Kind = "MAASCloudAccount"
	})
	p.AddResourceConfigurator("spectrocloud_cloudaccount_vsphere", func(r *config.Resource) {
		r.ShortGroup = "palette"
		r.Kind = "VSphereCloudAccount"
	})
	p.AddResourceConfigurator("spectrocloud_cloudaccount_aws", func(r *config.Resource) {
		r.ShortGroup = "palette"
		r.Kind = "AWSCloudAccount"
	})
	p.AddResourceConfigurator("spectrocloud_cloudaccount_azure", func(r *config.Resource) {
		r.ShortGroup = "palette"
		r.Kind = "AzureCloudAccount"
	})
	p.AddResourceConfigurator("spectrocloud_cloudaccount_gcp", func(r *config.Resource) {
		r.ShortGroup = "palette"
		r.Kind = "GCPCloudAccount"
		if s, ok := r.TerraformResource.Schema["gcp_json_credentials"]; ok {
			s.Sensitive = false
			s.Computed = false
			s.Optional = false

		}
	})

	// Clusters
	p.AddResourceConfigurator("spectrocloud_cluster_vsphere", func(r *config.Resource) {
		r.ShortGroup = "palette"
		r.Kind = "VSphereCluster"
	})
	p.AddResourceConfigurator("spectrocloud_cluster_eks", func(r *config.Resource) {
		r.ShortGroup = "palette"
		r.Kind = "EKSCluster"
	})
	p.AddResourceConfigurator("spectrocloud_virtual_cluster", func(r *config.Resource) {
		r.ShortGroup = "palette"
		r.Kind = "VirtualCluster"
	})
	p.AddResourceConfigurator("spectrocloud_cluster_aks", func(r *config.Resource) {
		r.ShortGroup = "palette"
		r.Kind = "AKSCluster"
	})
	p.AddResourceConfigurator("spectrocloud_cluster_gke", func(r *config.Resource) {
		r.ShortGroup = "palette"
		r.Kind = "GKECluster"
	})
	p.AddResourceConfigurator("spectrocloud_cluster_gcp", func(r *config.Resource) {
		r.ShortGroup = "palette"
		r.Kind = "GCPCluster"
	})
	p.AddResourceConfigurator("spectrocloud_cluster_azure", func(r *config.Resource) {
		r.ShortGroup = "palette"
		r.Kind = "AzureCluster"
	})
	p.AddResourceConfigurator("spectrocloud_cluster_aws", func(r *config.Resource) {
		r.ShortGroup = "palette"
		r.Kind = "AWSCluster"
	})
	p.AddResourceConfigurator("spectrocloud_cluster_custom_cloud", func(r *config.Resource) {
		r.ShortGroup = "palette"
		r.Kind = "CustomCloudCluster"
	})

	// IPAM
	p.AddResourceConfigurator("spectrocloud_privatecloudgateway_dns_map", func(r *config.Resource) {
		r.ShortGroup = "palette"
		r.Kind = "PrivateCloudGatewayDNSMap"
	})
	p.AddResourceConfigurator("spectrocloud_privatecloudgateway_ippool", func(r *config.Resource) {
		r.ShortGroup = "palette"
		r.Kind = "PrivateCloudGatewayIPPool"
	})

	// Macros
	p.AddResourceConfigurator("spectrocloud_macros", func(r *config.Resource) {
		r.ShortGroup = "palette"
		r.Kind = "Macros"
	})

	// Policy
	p.AddResourceConfigurator("spectrocloud_password_policy", func(r *config.Resource) {
		r.ShortGroup = "palette"
		r.Kind = "PasswordPolicy"
	})
	p.AddResourceConfigurator("spectrocloud_resource_limit", func(r *config.Resource) {
		r.ShortGroup = "palette"
		r.Kind = "ResourceLimit"
	})

	// Profiles
	p.AddResourceConfigurator("spectrocloud_cluster_profile", func(r *config.Resource) {
		r.ShortGroup = "palette"
		r.Kind = "ClusterProfile"
		r.ExternalName = config.IdentifierFromProvider
	})
	p.AddResourceConfigurator("spectrocloud_application_profile", func(r *config.Resource) {
		r.ShortGroup = "palette"
		r.Kind = "ApplicationProfile"
		r.ExternalName = config.IdentifierFromProvider
	})

	// RBAC
	p.AddResourceConfigurator("spectrocloud_role", func(r *config.Resource) {
		r.ShortGroup = "palette"
		r.Kind = "Role"
	})
	p.AddResourceConfigurator("spectrocloud_user", func(r *config.Resource) {
		r.ShortGroup = "palette"
		r.Kind = "User"
	})
	p.AddResourceConfigurator("spectrocloud_team", func(r *config.Resource) {
		r.ShortGroup = "palette"
		r.Kind = "Team"
	})

	// Registries
	p.AddResourceConfigurator("spectrocloud_registry_helm", func(r *config.Resource) {
		r.ShortGroup = "palette"
		r.Kind = "HelmRegistry"
	})
	p.AddResourceConfigurator("spectrocloud_registry_oci", func(r *config.Resource) {
		r.ShortGroup = "palette"
		r.Kind = "OCIRegistry"
	})

	// Security
	p.AddResourceConfigurator("spectrocloud_ssh_key", func(r *config.Resource) {
		r.ShortGroup = "palette"
		r.Kind = "SSHKey"
	})

	// Settings
	p.AddResourceConfigurator("spectrocloud_developer_setting", func(r *config.Resource) {
		r.ShortGroup = "palette"
		r.Kind = "DeveloperSetting"
	})
	p.AddResourceConfigurator("spectrocloud_platform_setting", func(r *config.Resource) {
		r.ShortGroup = "palette"
		r.Kind = "PlatformSetting"
	})
	p.AddResourceConfigurator("spectrocloud_filter", func(r *config.Resource) {
		r.ShortGroup = "palette"
		r.Kind = "Filter"
	})

	// Tenancy
	p.AddResourceConfigurator("spectrocloud_project", func(r *config.Resource) {
		r.ShortGroup = "palette"
		r.Kind = "Project"
	})
	p.AddResourceConfigurator("spectrocloud_cluster_group", func(r *config.Resource) {
		r.ShortGroup = "palette"
		r.Kind = "ClusterGroup"
	})
	p.AddResourceConfigurator("spectrocloud_workspace", func(r *config.Resource) {
		r.ShortGroup = "palette"
		r.Kind = "Workspace"
	})

	// Edge Management
	p.AddResourceConfigurator("spectrocloud_registration_token", func(r *config.Resource) {
		r.ShortGroup = "palette"
		r.Kind = "RegistrationToken"
	})
}
