/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	configauth "github.com/buttahtoast/provider-harbor/internal/controller/cluster/config/configauth"
	configsecurity "github.com/buttahtoast/provider-harbor/internal/controller/cluster/config/configsecurity"
	configsystem "github.com/buttahtoast/provider-harbor/internal/controller/cluster/config/configsystem"
	garbagecollection "github.com/buttahtoast/provider-harbor/internal/controller/cluster/garbagecollection/garbagecollection"
	group "github.com/buttahtoast/provider-harbor/internal/controller/cluster/group/group"
	interrogationservices "github.com/buttahtoast/provider-harbor/internal/controller/cluster/interrogationservices/interrogationservices"
	label "github.com/buttahtoast/provider-harbor/internal/controller/cluster/label/label"
	preheatinstance "github.com/buttahtoast/provider-harbor/internal/controller/cluster/preheatinstance/preheatinstance"
	immutabletagrule "github.com/buttahtoast/provider-harbor/internal/controller/cluster/project/immutabletagrule"
	membergroup "github.com/buttahtoast/provider-harbor/internal/controller/cluster/project/membergroup"
	memberuser "github.com/buttahtoast/provider-harbor/internal/controller/cluster/project/memberuser"
	project "github.com/buttahtoast/provider-harbor/internal/controller/cluster/project/project"
	retentionpolicy "github.com/buttahtoast/provider-harbor/internal/controller/cluster/project/retentionpolicy"
	webhook "github.com/buttahtoast/provider-harbor/internal/controller/cluster/project/webhook"
	providerconfig "github.com/buttahtoast/provider-harbor/internal/controller/cluster/providerconfig"
	purgeauditlog "github.com/buttahtoast/provider-harbor/internal/controller/cluster/purgeauditlog/purgeauditlog"
	registry "github.com/buttahtoast/provider-harbor/internal/controller/cluster/registry/registry"
	replication "github.com/buttahtoast/provider-harbor/internal/controller/cluster/replication/replication"
	robotaccount "github.com/buttahtoast/provider-harbor/internal/controller/cluster/robotaccount/robotaccount"
	task "github.com/buttahtoast/provider-harbor/internal/controller/cluster/tasks/task"
	user "github.com/buttahtoast/provider-harbor/internal/controller/cluster/user/user"
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

// SetupGated creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		configauth.SetupGated,
		configsecurity.SetupGated,
		configsystem.SetupGated,
		garbagecollection.SetupGated,
		group.SetupGated,
		interrogationservices.SetupGated,
		label.SetupGated,
		preheatinstance.SetupGated,
		immutabletagrule.SetupGated,
		membergroup.SetupGated,
		memberuser.SetupGated,
		project.SetupGated,
		retentionpolicy.SetupGated,
		webhook.SetupGated,
		providerconfig.SetupGated,
		purgeauditlog.SetupGated,
		registry.SetupGated,
		replication.SetupGated,
		robotaccount.SetupGated,
		task.SetupGated,
		user.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
