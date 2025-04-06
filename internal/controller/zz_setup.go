// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	akscluster "github.com/vrabbi/provider-palette/internal/controller/palette/akscluster"
	applicationprofile "github.com/vrabbi/provider-palette/internal/controller/palette/applicationprofile"
	awscloudaccount "github.com/vrabbi/provider-palette/internal/controller/palette/awscloudaccount"
	azurecloudaccount "github.com/vrabbi/provider-palette/internal/controller/palette/azurecloudaccount"
	backupstoragelocation "github.com/vrabbi/provider-palette/internal/controller/palette/backupstoragelocation"
	clustergroup "github.com/vrabbi/provider-palette/internal/controller/palette/clustergroup"
	clusterprofile "github.com/vrabbi/provider-palette/internal/controller/palette/clusterprofile"
	customcloudaccount "github.com/vrabbi/provider-palette/internal/controller/palette/customcloudaccount"
	developersetting "github.com/vrabbi/provider-palette/internal/controller/palette/developersetting"
	ekscluster "github.com/vrabbi/provider-palette/internal/controller/palette/ekscluster"
	filter "github.com/vrabbi/provider-palette/internal/controller/palette/filter"
	gcpcloudaccount "github.com/vrabbi/provider-palette/internal/controller/palette/gcpcloudaccount"
	helmregistry "github.com/vrabbi/provider-palette/internal/controller/palette/helmregistry"
	maascloudaccount "github.com/vrabbi/provider-palette/internal/controller/palette/maascloudaccount"
	macros "github.com/vrabbi/provider-palette/internal/controller/palette/macros"
	ociregistry "github.com/vrabbi/provider-palette/internal/controller/palette/ociregistry"
	passwordpolicy "github.com/vrabbi/provider-palette/internal/controller/palette/passwordpolicy"
	platformsetting "github.com/vrabbi/provider-palette/internal/controller/palette/platformsetting"
	privatecloudgatewaydnsmap "github.com/vrabbi/provider-palette/internal/controller/palette/privatecloudgatewaydnsmap"
	privatecloudgatewayippool "github.com/vrabbi/provider-palette/internal/controller/palette/privatecloudgatewayippool"
	project "github.com/vrabbi/provider-palette/internal/controller/palette/project"
	resourcelimit "github.com/vrabbi/provider-palette/internal/controller/palette/resourcelimit"
	role "github.com/vrabbi/provider-palette/internal/controller/palette/role"
	sshkey "github.com/vrabbi/provider-palette/internal/controller/palette/sshkey"
	team "github.com/vrabbi/provider-palette/internal/controller/palette/team"
	user "github.com/vrabbi/provider-palette/internal/controller/palette/user"
	virtualcluster "github.com/vrabbi/provider-palette/internal/controller/palette/virtualcluster"
	vspherecloudaccount "github.com/vrabbi/provider-palette/internal/controller/palette/vspherecloudaccount"
	vspherecluster "github.com/vrabbi/provider-palette/internal/controller/palette/vspherecluster"
	workspace "github.com/vrabbi/provider-palette/internal/controller/palette/workspace"
	providerconfig "github.com/vrabbi/provider-palette/internal/controller/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		akscluster.Setup,
		applicationprofile.Setup,
		awscloudaccount.Setup,
		azurecloudaccount.Setup,
		backupstoragelocation.Setup,
		clustergroup.Setup,
		clusterprofile.Setup,
		customcloudaccount.Setup,
		developersetting.Setup,
		ekscluster.Setup,
		filter.Setup,
		gcpcloudaccount.Setup,
		helmregistry.Setup,
		maascloudaccount.Setup,
		macros.Setup,
		ociregistry.Setup,
		passwordpolicy.Setup,
		platformsetting.Setup,
		privatecloudgatewaydnsmap.Setup,
		privatecloudgatewayippool.Setup,
		project.Setup,
		resourcelimit.Setup,
		role.Setup,
		sshkey.Setup,
		team.Setup,
		user.Setup,
		virtualcluster.Setup,
		vspherecloudaccount.Setup,
		vspherecluster.Setup,
		workspace.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
