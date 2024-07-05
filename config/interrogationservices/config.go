package interrogationservices

import "github.com/crossplane/upjet/pkg/config"

// Configure harbor_interrogation_serivces resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("harbor_interrogation_services", func(r *config.Resource) {
		r.ShortGroup = "interrogationservices"
		r.Kind = "InterrogationServices"
	})
}
