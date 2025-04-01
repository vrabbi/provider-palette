// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	awscloudaccount "github.com/vrabbi/provider-palette/internal/controller/palette/awscloudaccount"
	azurecloudaccount "github.com/vrabbi/provider-palette/internal/controller/palette/azurecloudaccount"
	clusterprofile "github.com/vrabbi/provider-palette/internal/controller/palette/clusterprofile"
	customcloudaccount "github.com/vrabbi/provider-palette/internal/controller/palette/customcloudaccount"
	developersetting "github.com/vrabbi/provider-palette/internal/controller/palette/developersetting"
	helmregistry "github.com/vrabbi/provider-palette/internal/controller/palette/helmregistry"
	macros "github.com/vrabbi/provider-palette/internal/controller/palette/macros"
	ociregistry "github.com/vrabbi/provider-palette/internal/controller/palette/ociregistry"
	platformsetting "github.com/vrabbi/provider-palette/internal/controller/palette/platformsetting"
	project "github.com/vrabbi/provider-palette/internal/controller/palette/project"
	sshkey "github.com/vrabbi/provider-palette/internal/controller/palette/sshkey"
	team "github.com/vrabbi/provider-palette/internal/controller/palette/team"
	user "github.com/vrabbi/provider-palette/internal/controller/palette/user"
	vspherecloudaccount "github.com/vrabbi/provider-palette/internal/controller/palette/vspherecloudaccount"
	vspherecluster "github.com/vrabbi/provider-palette/internal/controller/palette/vspherecluster"
	providerconfig "github.com/vrabbi/provider-palette/internal/controller/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		awscloudaccount.Setup,
		azurecloudaccount.Setup,
		clusterprofile.Setup,
		customcloudaccount.Setup,
		developersetting.Setup,
		helmregistry.Setup,
		macros.Setup,
		ociregistry.Setup,
		platformsetting.Setup,
		project.Setup,
		sshkey.Setup,
		team.Setup,
		user.Setup,
		vspherecloudaccount.Setup,
		vspherecluster.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
