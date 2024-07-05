package garbagecollection

import "github.com/crossplane/upjet/pkg/config"

// Configure harbor_garbage_collection resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("harbor_garbage_collection", func(r *config.Resource) {
		r.ShortGroup = "garbagecollection"
		r.Kind = "GarbageCollection"
	})
}
