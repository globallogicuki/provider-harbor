package configsecurity

import "github.com/crossplane/upjet/pkg/config"

// Configure harbor_config_security resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("harbor_config_security", func(r *config.Resource) {
		r.ShortGroup = "config"
		r.Kind = "ConfigSecurity"
	})
}
