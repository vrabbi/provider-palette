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
		r.References["cloud_account_id"] = config.Reference{
			TerraformName: "spectrocloud_cloudaccount_vsphere",
		}
		r.References["cluster_profile.id"] = config.Reference{
			TerraformName: "spectrocloud_cluster_profile",
		}
		r.References["cluster_profile.pack.registry_uid"] = config.Reference{
			TerraformName: "spectrocloud_registry_oci",
		}
		r.References["backup_policy.backup_location_id"] = config.Reference{
			TerraformName: "spectrocloud_backup_storage_location",
		}
		r.References["machine_pool.placement.static_ip_pool_id"] = config.Reference{
			TerraformName: "spectrocloud_privatecloudgateway_ippool",
		}
	})
	p.AddResourceConfigurator("spectrocloud_cluster_eks", func(r *config.Resource) {
		r.ShortGroup = "palette"
		r.Kind = "EKSCluster"
		r.References["cloud_account_id"] = config.Reference{
			TerraformName: "spectrocloud_cloudaccount_aws",
		}
		r.References["cluster_profile.id"] = config.Reference{
			TerraformName: "spectrocloud_cluster_profile",
		}
		r.References["cluster_profile.pack.registry_uid"] = config.Reference{
			TerraformName: "spectrocloud_registry_oci",
		}
		r.References["backup_policy.backup_location_id"] = config.Reference{
			TerraformName: "spectrocloud_backup_storage_location",
		}
	})
	p.AddResourceConfigurator("spectrocloud_virtual_cluster", func(r *config.Resource) {
		r.ShortGroup = "palette"
		r.Kind = "VirtualCluster"
		r.References["cluster_group_uid"] = config.Reference{
			TerraformName: "spectrocloud_cluster_group",
		}
		r.References["backup_policy.backup_location_id"] = config.Reference{
			TerraformName: "spectrocloud_backup_storage_location",
		}
		r.References["cluster_profile.id"] = config.Reference{
			TerraformName: "spectrocloud_cluster_profile",
		}
		r.References["cluster_profile.pack.registry_uid"] = config.Reference{
			TerraformName: "spectrocloud_registry_oci",
		}
	})
	p.AddResourceConfigurator("spectrocloud_cluster_aks", func(r *config.Resource) {
		r.ShortGroup = "palette"
		r.Kind = "AKSCluster"
		r.References["cloud_account_id"] = config.Reference{
			TerraformName: "spectrocloud_cloudaccount_azure",
		}
		r.References["cluster_profile.id"] = config.Reference{
			TerraformName: "spectrocloud_cluster_profile",
		}
		r.References["cluster_profile.pack.registry_uid"] = config.Reference{
			TerraformName: "spectrocloud_registry_oci",
		}
		r.References["backup_policy.backup_location_id"] = config.Reference{
			TerraformName: "spectrocloud_backup_storage_location",
		}
	})
	p.AddResourceConfigurator("spectrocloud_cluster_gke", func(r *config.Resource) {
		r.ShortGroup = "palette"
		r.Kind = "GKECluster"
		r.References["cloud_account_id"] = config.Reference{
			TerraformName: "spectrocloud_cloudaccount_gcp",
		}
		r.References["cluster_profile.id"] = config.Reference{
			TerraformName: "spectrocloud_cluster_profile",
		}
		r.References["cluster_profile.pack.registry_uid"] = config.Reference{
			TerraformName: "spectrocloud_registry_oci",
		}
		r.References["backup_policy.backup_location_id"] = config.Reference{
			TerraformName: "spectrocloud_backup_storage_location",
		}

	})
	p.AddResourceConfigurator("spectrocloud_cluster_gcp", func(r *config.Resource) {
		r.ShortGroup = "palette"
		r.Kind = "GCPCluster"
		r.References["cloud_account_id"] = config.Reference{
			TerraformName: "spectrocloud_cloudaccount_gcp",
		}
		r.References["cluster_profile.id"] = config.Reference{
			TerraformName: "spectrocloud_cluster_profile",
		}
		r.References["cluster_profile.pack.registry_uid"] = config.Reference{
			TerraformName: "spectrocloud_registry_oci",
		}
		r.References["backup_policy.backup_location_id"] = config.Reference{
			TerraformName: "spectrocloud_backup_storage_location",
		}
	})
	p.AddResourceConfigurator("spectrocloud_cluster_azure", func(r *config.Resource) {
		r.ShortGroup = "palette"
		r.Kind = "AzureCluster"
		r.References["cloud_account_id"] = config.Reference{
			TerraformName: "spectrocloud_cloudaccount_azure",
		}
		r.References["cluster_profile.id"] = config.Reference{
			TerraformName: "spectrocloud_cluster_profile",
		}
		r.References["cluster_profile.pack.registry_uid"] = config.Reference{
			TerraformName: "spectrocloud_registry_oci",
		}
		r.References["backup_policy.backup_location_id"] = config.Reference{
			TerraformName: "spectrocloud_backup_storage_location",
		}
	})
	p.AddResourceConfigurator("spectrocloud_cluster_aws", func(r *config.Resource) {
		r.ShortGroup = "palette"
		r.Kind = "AWSCluster"
		r.References["cloud_account_id"] = config.Reference{
			TerraformName: "spectrocloud_cloudaccount_aws",
		}
		r.References["cluster_profile.id"] = config.Reference{
			TerraformName: "spectrocloud_cluster_profile",
		}
		r.References["cluster_profile.pack.registry_uid"] = config.Reference{
			TerraformName: "spectrocloud_registry_oci",
		}
		r.References["backup_policy.backup_location_id"] = config.Reference{
			TerraformName: "spectrocloud_backup_storage_location",
		}
	})
	p.AddResourceConfigurator("spectrocloud_cluster_custom_cloud", func(r *config.Resource) {
		r.ShortGroup = "palette"
		r.Kind = "CustomCloudCluster"
		r.References["cloud_account_id"] = config.Reference{
			TerraformName: "spectrocloud_cloudaccount_custom",
		}
		r.References["cluster_profile.id"] = config.Reference{
			TerraformName: "spectrocloud_cluster_profile",
		}
		r.References["cluster_profile.pack.registry_uid"] = config.Reference{
			TerraformName: "spectrocloud_registry_oci",
		}
		r.References["backup_policy.backup_location_id"] = config.Reference{
			TerraformName: "spectrocloud_backup_storage_location",
		}
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
		r.References["project_role.role_ids"] = config.Reference{
			TerraformName: "spectrocloud_role",
		}
		r.References["project_role.project_id"] = config.Reference{
			TerraformName: "spectrocloud_project",
		}
		r.References["tenant_role"] = config.Reference{
			TerraformName: "spectrocloud_role",
		}
		r.References["workspace_role.project_id"] = config.Reference{
			TerraformName: "spectrocloud_project",
		}
		r.References["workspace_role.workspace.role_ids"] = config.Reference{
			TerraformName: "spectrocloud_role",
		}
		r.References["workspace_role.workspace.id"] = config.Reference{
			TerraformName: "spectrocloud_workspace",
		}
		r.References["resource_role.project_ids"] = config.Reference{
			TerraformName: "spectrocloud_project",
		}
		r.References["resource_role.role_ids"] = config.Reference{
			TerraformName: "spectrocloud_role",
		}
		r.References["resource_role.filter_ids"] = config.Reference{
			TerraformName: "spectrocloud_filter",
		}
		r.References["team_ids"] = config.Reference{
			TerraformName: "spectrocloud_team",
		}
	})
	p.AddResourceConfigurator("spectrocloud_team", func(r *config.Resource) {
		r.ShortGroup = "palette"
		r.Kind = "Team"
		r.References["project_role_mapping.roles"] = config.Reference{
			TerraformName: "spectrocloud_role",
		}
		r.References["project_role_mapping.id"] = config.Reference{
			TerraformName: "spectrocloud_project",
		}
		r.References["tenant_role_mapping"] = config.Reference{
			TerraformName: "spectrocloud_role",
		}
		r.References["workspace_role_mapping.id"] = config.Reference{
			TerraformName: "spectrocloud_project",
		}
		r.References["workspace_role_mapping.workspace.roles"] = config.Reference{
			TerraformName: "spectrocloud_role",
		}
		r.References["workspace_role_mapping.workspace.id"] = config.Reference{
			TerraformName: "spectrocloud_workspace",
		}
		r.References["users"] = config.Reference{
			TerraformName: "spectrocloud_user",
		}
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
		r.References["cluster_profile.id"] = config.Reference{
			TerraformName: "spectrocloud_cluster_profile",
		}
		r.References["cluster_profile.pack.registry_uid"] = config.Reference{
			TerraformName: "spectrocloud_registry_oci",
		}
	})
	p.AddResourceConfigurator("spectrocloud_workspace", func(r *config.Resource) {
		r.ShortGroup = "palette"
		r.Kind = "Workspace"
		r.References["backup_policy.backup_location_id"] = config.Reference{
			TerraformName: "spectrocloud_backup_storage_location",
		}
	})

	// Edge Management
	p.AddResourceConfigurator("spectrocloud_registration_token", func(r *config.Resource) {
		r.ShortGroup = "palette"
		r.Kind = "RegistrationToken"
		r.References["project_uid"] = config.Reference{
			TerraformName: "spectrocloud_project",
		}
	})
	p.AddResourceConfigurator("spectrocloud_appliance", func(r *config.Resource) {
		r.ShortGroup = "palette"
		r.Kind = "Appliance"
		r.ExternalName = config.IdentifierFromProvider
	})

	// App Mode
	p.AddResourceConfigurator("spectrocloud_application", func(r *config.Resource) {
		r.ShortGroup = "palette"
		r.Kind = "Application"
		r.References["application_profile_uid"] = config.Reference{
			TerraformName: "spectrocloud_application_profile",
		}
		r.References["config.cluster_group_uid"] = config.Reference{
			TerraformName: "spectrocloud_cluster_group",
		}
		r.References["config.cluster_uid"] = config.Reference{
			TerraformName: "spectrocloud_virtual_cluster",
		}
	})
}
