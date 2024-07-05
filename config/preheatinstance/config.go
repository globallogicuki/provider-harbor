package preheatinstance

import "github.com/crossplane/upjet/pkg/config"

// Configure harbor_preheat_instance resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("harbor_preheat_instance", func(r *config.Resource) {
		r.ShortGroup = "preheatinstance"
		r.Kind = "PreheatInstance"
	})
}
