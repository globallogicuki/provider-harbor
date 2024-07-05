package configauth

import "github.com/crossplane/upjet/pkg/config"

// Configure harbor_config_auth resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("harbor_config_auth", func(r *config.Resource) {
		r.ShortGroup = "config"
		r.Kind = "ConfigAuth"
	})
}
