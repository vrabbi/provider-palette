package resources

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("spectrocloud_project", func(r *config.Resource) {
		r.Kind = "Project"
	})
	p.AddResourceConfigurator("spectrocloud_ssh_key", func(r *config.Resource) {
		r.Kind = "SSHKey"
	})
	p.AddResourceConfigurator("spectrocloud_registry_helm", func(r *config.Resource) {
		r.Kind = "HelmRegistry"
	})
	p.AddResourceConfigurator("spectrocloud_user", func(r *config.Resource) {
		r.Kind = "User"
	})
	p.AddResourceConfigurator("spectrocloud_team", func(r *config.Resource) {
		r.Kind = "Team"
	})
}
