/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// (lornest) embedding schema and metadata files
	_ "embed"

	ujconfig "github.com/crossplane/upjet/pkg/config"

	"github.com/globallogicuki/provider-harbor/config/configauth"
	"github.com/globallogicuki/provider-harbor/config/configemail"
	configsystem "github.com/globallogicuki/provider-harbor/config/configsecurity"
	configsecurity "github.com/globallogicuki/provider-harbor/config/configsystem"
	"github.com/globallogicuki/provider-harbor/config/garbagecollection"
	"github.com/globallogicuki/provider-harbor/config/group"
	"github.com/globallogicuki/provider-harbor/config/immutabletagrule"
	"github.com/globallogicuki/provider-harbor/config/interrogationservices"
	"github.com/globallogicuki/provider-harbor/config/label"
	"github.com/globallogicuki/provider-harbor/config/membergroup"
	"github.com/globallogicuki/provider-harbor/config/memberuser"
	"github.com/globallogicuki/provider-harbor/config/preheatinstance"
	"github.com/globallogicuki/provider-harbor/config/project"
	"github.com/globallogicuki/provider-harbor/config/purgeauditlog"
	"github.com/globallogicuki/provider-harbor/config/registry"
	"github.com/globallogicuki/provider-harbor/config/replication"
	"github.com/globallogicuki/provider-harbor/config/retentionpolicy"
	"github.com/globallogicuki/provider-harbor/config/robotaccount"
	"github.com/globallogicuki/provider-harbor/config/tasks"
	"github.com/globallogicuki/provider-harbor/config/user"
	"github.com/globallogicuki/provider-harbor/config/webhook"
)

const (
	resourcePrefix = "harbor"
	modulePath     = "github.com/globallogicuki/provider-harbor"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider(
		[]byte(providerSchema),
		resourcePrefix,
		modulePath,
		[]byte(providerMetadata),
		ujconfig.WithRootGroup("harbor.crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		),
	)

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		configauth.Configure,
		configemail.Configure,
		configsecurity.Configure,
		configsystem.Configure,
		garbagecollection.Configure,
		group.Configure,
		immutabletagrule.Configure,
		interrogationservices.Configure,
		label.Configure,
		preheatinstance.Configure,
		project.Configure,
		membergroup.Configure,
		memberuser.Configure,
		webhook.Configure,
		purgeauditlog.Configure,
		registry.Configure,
		replication.Configure,
		retentionpolicy.Configure,
		robotaccount.Configure,
		tasks.Configure,
		user.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
