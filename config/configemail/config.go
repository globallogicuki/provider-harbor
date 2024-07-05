package configemail

import "github.com/crossplane/upjet/pkg/config"

// Configure harbor_config_email resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("harbor_config_email", func(r *config.Resource) {
		r.ShortGroup = "config"
		r.Kind = "ConfigEmail"
	})
}
