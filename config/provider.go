/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// (lornest) embedding schema and metadata files
	_ "embed"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"

	configauthCluster "github.com/buttahtoast/provider-harbor/config/cluster/configauth"
	configsystemCluster "github.com/buttahtoast/provider-harbor/config/cluster/configsecurity"
	configsecurityCluster "github.com/buttahtoast/provider-harbor/config/cluster/configsystem"
	garbagecollectionCluster "github.com/buttahtoast/provider-harbor/config/cluster/garbagecollection"
	groupCluster "github.com/buttahtoast/provider-harbor/config/cluster/group"
	immutabletagruleCluster "github.com/buttahtoast/provider-harbor/config/cluster/immutabletagrule"
	interrogationservicesCluster "github.com/buttahtoast/provider-harbor/config/cluster/interrogationservices"
	labelCluster "github.com/buttahtoast/provider-harbor/config/cluster/label"
	membergroupCluster "github.com/buttahtoast/provider-harbor/config/cluster/membergroup"
	memberuserCluster "github.com/buttahtoast/provider-harbor/config/cluster/memberuser"
	preheatinstanceCluster "github.com/buttahtoast/provider-harbor/config/cluster/preheatinstance"
	projectCluster "github.com/buttahtoast/provider-harbor/config/cluster/project"
	purgeauditlogCluster "github.com/buttahtoast/provider-harbor/config/cluster/purgeauditlog"
	registryCluster "github.com/buttahtoast/provider-harbor/config/cluster/registry"
	replicationCluster "github.com/buttahtoast/provider-harbor/config/cluster/replication"
	retentionpolicyCluster "github.com/buttahtoast/provider-harbor/config/cluster/retentionpolicy"
	robotaccountCluster "github.com/buttahtoast/provider-harbor/config/cluster/robotaccount"
	tasksCluster "github.com/buttahtoast/provider-harbor/config/cluster/tasks"
	userCluster "github.com/buttahtoast/provider-harbor/config/cluster/user"
	webhookCluster "github.com/buttahtoast/provider-harbor/config/cluster/webhook"

	configauthNamespaced "github.com/buttahtoast/provider-harbor/config/namespaced/configauth"
	configsystemNamespaced "github.com/buttahtoast/provider-harbor/config/namespaced/configsecurity"
	configsecurityNamespaced "github.com/buttahtoast/provider-harbor/config/namespaced/configsystem"
	garbagecollectionNamespaced "github.com/buttahtoast/provider-harbor/config/namespaced/garbagecollection"
	groupNamespaced "github.com/buttahtoast/provider-harbor/config/namespaced/group"
	immutabletagruleNamespaced "github.com/buttahtoast/provider-harbor/config/namespaced/immutabletagrule"
	interrogationservicesNamespaced "github.com/buttahtoast/provider-harbor/config/namespaced/interrogationservices"
	labelNamespaced "github.com/buttahtoast/provider-harbor/config/namespaced/label"
	membergroupNamespaced "github.com/buttahtoast/provider-harbor/config/namespaced/membergroup"
	memberuserNamespaced "github.com/buttahtoast/provider-harbor/config/namespaced/memberuser"
	preheatinstanceNamespaced "github.com/buttahtoast/provider-harbor/config/namespaced/preheatinstance"
	projectNamespaced "github.com/buttahtoast/provider-harbor/config/namespaced/project"
	purgeauditlogNamespaced "github.com/buttahtoast/provider-harbor/config/namespaced/purgeauditlog"
	registryNamespaced "github.com/buttahtoast/provider-harbor/config/namespaced/registry"
	replicationNamespaced "github.com/buttahtoast/provider-harbor/config/namespaced/replication"
	retentionpolicyNamespaced "github.com/buttahtoast/provider-harbor/config/namespaced/retentionpolicy"
	robotaccountNamespaced "github.com/buttahtoast/provider-harbor/config/namespaced/robotaccount"
	tasksNamespaced "github.com/buttahtoast/provider-harbor/config/namespaced/tasks"
	userNamespaced "github.com/buttahtoast/provider-harbor/config/namespaced/user"
	webhookNamespaced "github.com/buttahtoast/provider-harbor/config/namespaced/webhook"
)

const (
	resourcePrefix = "harbor"
	modulePath     = "github.com/buttahtoast/provider-harbor"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns cluster-scoped provider configuration.
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider(
		[]byte(providerSchema),
		resourcePrefix,
		modulePath,
		[]byte(providerMetadata),
		ujconfig.WithRootGroup("buttah.cloud"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		),
	)

	for _, configure := range []func(provider *ujconfig.Provider){
		configauthCluster.Configure,
		configsecurityCluster.Configure,
		configsystemCluster.Configure,
		garbagecollectionCluster.Configure,
		groupCluster.Configure,
		immutabletagruleCluster.Configure,
		interrogationservicesCluster.Configure,
		labelCluster.Configure,
		preheatinstanceCluster.Configure,
		projectCluster.Configure,
		membergroupCluster.Configure,
		memberuserCluster.Configure,
		webhookCluster.Configure,
		purgeauditlogCluster.Configure,
		registryCluster.Configure,
		replicationCluster.Configure,
		retentionpolicyCluster.Configure,
		robotaccountCluster.Configure,
		tasksCluster.Configure,
		userCluster.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}

// GetProviderNamespaced returns namespaced provider configuration.
func GetProviderNamespaced() *ujconfig.Provider {
	pc := ujconfig.NewProvider(
		[]byte(providerSchema),
		resourcePrefix,
		modulePath,
		[]byte(providerMetadata),
		ujconfig.WithRootGroup("buttah.m.cloud"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		),
		ujconfig.WithExampleManifestConfiguration(ujconfig.ExampleManifestConfiguration{
			ManagedResourceNamespace: "crossplane-system",
		}),
	)

	for _, configure := range []func(provider *ujconfig.Provider){
		configauthNamespaced.Configure,
		configsecurityNamespaced.Configure,
		configsystemNamespaced.Configure,
		garbagecollectionNamespaced.Configure,
		groupNamespaced.Configure,
		immutabletagruleNamespaced.Configure,
		interrogationservicesNamespaced.Configure,
		labelNamespaced.Configure,
		preheatinstanceNamespaced.Configure,
		projectNamespaced.Configure,
		membergroupNamespaced.Configure,
		memberuserNamespaced.Configure,
		webhookNamespaced.Configure,
		purgeauditlogNamespaced.Configure,
		registryNamespaced.Configure,
		replicationNamespaced.Configure,
		retentionpolicyNamespaced.Configure,
		robotaccountNamespaced.Configure,
		tasksNamespaced.Configure,
		userNamespaced.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}