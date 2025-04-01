package resources

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("spectrocloud_project", func(r *config.Resource) {
		r.ShortGroup = "palette"
		r.Kind = "Project"
	})
	p.AddResourceConfigurator("spectrocloud_ssh_key", func(r *config.Resource) {
		r.ShortGroup = "palette"
		r.Kind = "SSHKey"
	})
	p.AddResourceConfigurator("spectrocloud_registry_helm", func(r *config.Resource) {
		r.ShortGroup = "palette"
		r.Kind = "HelmRegistry"
	})
	p.AddResourceConfigurator("spectrocloud_registry_oci", func(r *config.Resource) {
		r.ShortGroup = "palette"
		r.Kind = "OCIRegistry"
	})
	p.AddResourceConfigurator("spectrocloud_user", func(r *config.Resource) {
		r.ShortGroup = "palette"
		r.Kind = "User"
	})
	p.AddResourceConfigurator("spectrocloud_team", func(r *config.Resource) {
		r.ShortGroup = "palette"
		r.Kind = "Team"
	})
	p.AddResourceConfigurator("spectrocloud_cloudaccount_custom", func(r *config.Resource) {
		r.ShortGroup = "palette"
		r.Kind = "CustomCloudAccount"
	})
	p.AddResourceConfigurator("spectrocloud_cluster_profile", func(r *config.Resource) {
		r.ShortGroup = "palette"
		r.Kind = "ClusterProfile"
		r.ExternalName = config.IdentifierFromProvider
	})
	p.AddResourceConfigurator("spectrocloud_cluster_vsphere", func(r *config.Resource) {
		r.ShortGroup = "palette"
		r.Kind = "VSphereCluster"
	})
	p.AddResourceConfigurator("spectrocloud_cloudaccount_vsphere", func(r *config.Resource) {
		r.ShortGroup = "palette"
		r.Kind = "VSphereCloudAccount"
	})
	p.AddResourceConfigurator("spectrocloud_cloudaccount_aws", func(r *config.Resource) {
		r.ShortGroup = "palette"
		r.Kind = "AWSCloudAccount"
	})
	p.AddResourceConfigurator("spectrocloud_macros", func(r *config.Resource) {
		r.ShortGroup = "palette"
		r.Kind = "Macros"
	})
	p.AddResourceConfigurator("spectrocloud_cloudaccount_azure", func(r *config.Resource) {
		r.ShortGroup = "palette"
		r.Kind = "AzureCloudAccount"
	})
	p.AddResourceConfigurator("spectrocloud_developer_setting", func(r *config.Resource) {
		r.ShortGroup = "palette"
		r.Kind = "DeveloperSetting"
	})
	p.AddResourceConfigurator("spectrocloud_platform_setting", func(r *config.Resource) {
		r.ShortGroup = "palette"
		r.Kind = "PlatformSetting"
	})
	p.AddResourceConfigurator("spectrocloud_cluster_eks", func(r *config.Resource) {
		r.ShortGroup = "palette"
		r.Kind = "EKSCluster"
	})
	p.AddResourceConfigurator("spectrocloud_privatecloudgateway_dns_map", func(r *config.Resource) {
		r.ShortGroup = "palette"
		r.Kind = "PrivateCloudGatewayDNSMap"
	})
	p.AddResourceConfigurator("spectrocloud_privatecloudgateway_ippool", func(r *config.Resource) {
		r.ShortGroup = "palette"
		r.Kind = "PrivateCloudGatewayIPPool"
	})

}
