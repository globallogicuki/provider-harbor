package registry

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure harbor_registry resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("harbor_registry", func(r *config.Resource) {
		r.ShortGroup = "registry"
		r.Kind = "Registry"

		// Configure the providerName field to provide better documentation
		if s, ok := r.TerraformResource.Schema["provider_name"]; ok {
			s.Description = "Can be one of: alibaba, artifact-hub, aws, azure, docker-hub, docker-registry, gitlab, github, google, harbor, huawei, jfrog, quay"
		}
	})
}
