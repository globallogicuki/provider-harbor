/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// (lornest) embedding schema and metadata files
	_ "embed"

	ujconfig "github.com/crossplane/upjet/pkg/config"

	"github.com/globallogicuki/provider-harbor/config/project"
	"github.com/globallogicuki/provider-harbor/config/projectmembergroup"
	"github.com/globallogicuki/provider-harbor/config/retentionpolicy"
	"github.com/globallogicuki/provider-harbor/config/robotaccount"
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
		ujconfig.WithRootGroup("crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		),
	)

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		project.Configure,
		retentionpolicy.Configure,
		robotaccount.Configure,
		projectmembergroup.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
