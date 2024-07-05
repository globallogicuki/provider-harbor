package configsystem

import "github.com/crossplane/upjet/pkg/config"

// Configure harbor_config_system resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("harbor_config_system", func(r *config.Resource) {
		r.ShortGroup = "config"
		r.Kind = "ConfigSystem"
	})
}
