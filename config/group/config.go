package group

import "github.com/crossplane/upjet/pkg/config"

// Configure harbor_group resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("harbor_group", func(r *config.Resource) {
		r.ShortGroup = "group"
	})
}
