package retentionpolicy

import "github.com/crossplane/upjet/pkg/config"

// Configure harbor_retention_policy resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("harbor_retention_policy", func(r *config.Resource) {
		r.ShortGroup = "retentionpolicy"
		r.Kind = "RetentionPolicy"
	})
}
