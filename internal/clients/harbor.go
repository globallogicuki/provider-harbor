/*
Copyright 2021 Upbound Inc.
*/

package clients

import (
	"context"
	"encoding/json"

	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/crossplane/upjet/pkg/terraform"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/globallogicuki/provider-harbor/apis/v1beta1"
)

const (
	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal harbor credentials as JSON"

	// provider config variables
	url         = "url"
	username    = "username"
	password    = "password"
	apiVersion  = "api_version"
	bearerToken = "bearer_token"
	insecure    = "insecure"
)

type harborConfig struct {
	URL         *string `json:"url,omitempty"`
	Username    *string `json:"username,omitempty"`
	Password    *string `json:"password,omitempty"`
	ApiVersion  *int    `json:"api_version,omitempty"`
	BearerToken *string `json:"bearer_token,omitempty"`
	Insecure    *bool   `json:"insecure,omitempty"`
}

func terraformProviderConfigurationBuilder(
	creds harborConfig,
) (terraform.ProviderConfiguration, error) {
	cnf := terraform.ProviderConfiguration{}

	if creds.URL != nil {
		cnf[url] = *creds.URL
	}

	if creds.Username != nil {
		cnf[username] = *creds.Username
	}

	if creds.Password != nil {
		cnf[password] = *creds.Password
	}

	if creds.ApiVersion != nil {
		cnf[apiVersion] = *creds.ApiVersion
	}

	if creds.BearerToken != nil {
		cnf[bearerToken] = *creds.BearerToken
	}

	return cnf, nil
}

// TerraformSetupBuilder builds Terraform a terraform.SetupFn function which
// returns Terraform provider setup configuration
func TerraformSetupBuilder(version, providerSource, providerVersion string) terraform.SetupFn {
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{
			Version: version,
			Requirement: terraform.ProviderRequirement{
				Source:  providerSource,
				Version: providerVersion,
			},
		}

		configRef := mg.GetProviderConfigReference()
		if configRef == nil {
			return ps, errors.New(errNoProviderConfig)
		}
		pc := &v1beta1.ProviderConfig{}
		if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
			return ps, errors.Wrap(err, errGetProviderConfig)
		}

		t := resource.NewProviderConfigUsageTracker(client, &v1beta1.ProviderConfigUsage{})
		if err := t.Track(ctx, mg); err != nil {
			return ps, errors.Wrap(err, errTrackUsage)
		}

		data, err := resource.CommonCredentialExtractor(
			ctx,
			pc.Spec.Credentials.Source,
			client,
			pc.Spec.Credentials.CommonCredentialSelectors,
		)
		if err != nil {
			return ps, errors.Wrap(err, errExtractCredentials)
		}
		creds := map[string]string{}
		if err := json.Unmarshal(data, &creds); err != nil {
			return ps, errors.Wrap(err, errUnmarshalCredentials)
		}

		ps.Configuration = map[string]any{}
		if v, ok := creds[url]; ok {
			ps.Configuration[url] = v
		}
		if v, ok := creds[username]; ok {
			ps.Configuration[username] = v
		}
		if v, ok := creds[password]; ok {
			ps.Configuration[password] = v
		}
		return ps, nil
	}
}
