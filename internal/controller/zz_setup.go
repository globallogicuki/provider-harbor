// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	configauth "github.com/globallogicuki/provider-harbor/internal/controller/config/configauth"
	configsecurity "github.com/globallogicuki/provider-harbor/internal/controller/config/configsecurity"
	configsystem "github.com/globallogicuki/provider-harbor/internal/controller/config/configsystem"
	garbagecollection "github.com/globallogicuki/provider-harbor/internal/controller/garbagecollection/garbagecollection"
	group "github.com/globallogicuki/provider-harbor/internal/controller/group/group"
	interrogationservices "github.com/globallogicuki/provider-harbor/internal/controller/interrogationservices/interrogationservices"
	label "github.com/globallogicuki/provider-harbor/internal/controller/label/label"
	preheatinstance "github.com/globallogicuki/provider-harbor/internal/controller/preheatinstance/preheatinstance"
	immutabletagrule "github.com/globallogicuki/provider-harbor/internal/controller/project/immutabletagrule"
	membergroup "github.com/globallogicuki/provider-harbor/internal/controller/project/membergroup"
	memberuser "github.com/globallogicuki/provider-harbor/internal/controller/project/memberuser"
	project "github.com/globallogicuki/provider-harbor/internal/controller/project/project"
	retentionpolicy "github.com/globallogicuki/provider-harbor/internal/controller/project/retentionpolicy"
	webhook "github.com/globallogicuki/provider-harbor/internal/controller/project/webhook"
	providerconfig "github.com/globallogicuki/provider-harbor/internal/controller/providerconfig"
	purgeauditlog "github.com/globallogicuki/provider-harbor/internal/controller/purgeauditlog/purgeauditlog"
	registry "github.com/globallogicuki/provider-harbor/internal/controller/registry/registry"
	replication "github.com/globallogicuki/provider-harbor/internal/controller/registry/replication"
	robotaccount "github.com/globallogicuki/provider-harbor/internal/controller/robotaccount/robotaccount"
	task "github.com/globallogicuki/provider-harbor/internal/controller/tasks/task"
	user "github.com/globallogicuki/provider-harbor/internal/controller/user/user"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		configauth.Setup,
		configsecurity.Setup,
		configsystem.Setup,
		garbagecollection.Setup,
		group.Setup,
		interrogationservices.Setup,
		label.Setup,
		preheatinstance.Setup,
		immutabletagrule.Setup,
		membergroup.Setup,
		memberuser.Setup,
		project.Setup,
		retentionpolicy.Setup,
		webhook.Setup,
		providerconfig.Setup,
		purgeauditlog.Setup,
		registry.Setup,
		replication.Setup,
		robotaccount.Setup,
		task.Setup,
		user.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
