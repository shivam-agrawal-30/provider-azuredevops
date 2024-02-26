/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/pkg/config"

	"github.com/shivam-agrawal-30/provider-azuredevops/config/gitrepository"
	"github.com/shivam-agrawal-30/provider-azuredevops/config/project"
)

const (
	resourcePrefix = "azuredevops"
	modulePath     = "github.com/shivam-agrawal-30/provider-azuredevops"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		project.Configure,
		gitrepository.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
