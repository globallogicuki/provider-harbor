/*
Copyright 2022 Upbound Inc.
*/

package config

import "github.com/crossplane/upjet/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// Import requires using a randomly generated ID from provider: nl-2e21sda
	"harbor_config_auth":            config.IdentifierFromProvider,
	"harbor_config_email":           config.IdentifierFromProvider,
	"harbor_config_security":        config.IdentifierFromProvider,
	"harbor_config_system":          config.IdentifierFromProvider,
	"harbor_garbage_collection":     config.IdentifierFromProvider,
	"harbor_group":                  config.IdentifierFromProvider,
	"harbor_immutable_tag_rule":     config.IdentifierFromProvider,
	"harbor_interrogation_services": config.IdentifierFromProvider,
	"harbor_label":                  config.IdentifierFromProvider,
	"harbor_preheat_instance":       config.IdentifierFromProvider,
	"harbor_project":                config.IdentifierFromProvider,
	"harbor_project_member_group":   config.IdentifierFromProvider,
	"harbor_project_member_user":    config.IdentifierFromProvider,
	"harbor_project_webhook":        config.IdentifierFromProvider,
	"harbor_purge_audit_log":        config.IdentifierFromProvider,
	"harbor_registry":               config.IdentifierFromProvider,
	"harbor_replication":            config.IdentifierFromProvider,
	"harbor_retention_policy":       config.IdentifierFromProvider,
	"harbor_robot_account":          config.IdentifierFromProvider,
	"harbor_tasks":                  config.IdentifierFromProvider,
	"harbor_user":                   config.IdentifierFromProvider,
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
