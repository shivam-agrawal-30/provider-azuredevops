package gitrepository

//import "github.com/upbound/upjet/pkg/config"

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azuredevops_git_repository", func(r *config.Resource) {
		r.ShortGroup = "gitrepository"
		r.References["project_id"] = config.Reference{
			Type: "github.com/shivam-agrawal-30/provider-azuredevops/apis/project/v1alpha1.Project",
		}
	})
}