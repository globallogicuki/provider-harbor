// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	group "github.com/globallogicuki/provider-harbor/internal/controller/group/group"
	label "github.com/globallogicuki/provider-harbor/internal/controller/label/label"
	membergroup "github.com/globallogicuki/provider-harbor/internal/controller/project/membergroup"
	project "github.com/globallogicuki/provider-harbor/internal/controller/project/project"
	retentionpolicy "github.com/globallogicuki/provider-harbor/internal/controller/project/retentionpolicy"
	providerconfig "github.com/globallogicuki/provider-harbor/internal/controller/providerconfig"
	robotaccount "github.com/globallogicuki/provider-harbor/internal/controller/robotaccount/robotaccount"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		group.Setup,
		label.Setup,
		membergroup.Setup,
		project.Setup,
		retentionpolicy.Setup,
		providerconfig.Setup,
		robotaccount.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
