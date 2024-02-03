package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ provider.Provider = &IPHelperProvider{}

type IPHelperProvider struct {
	version string
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &IPHelperProvider{
			version: version,
		}
	}
}

func (p *IPHelperProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "iphelpers"
	resp.Version = p.version
}

func (p *IPHelperProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
}

func (p *IPHelperProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
}

func (p *IPHelperProvider) Resources(ctx context.Context) []func() resource.Resource {
	return nil
}

func (p *IPHelperProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewEUIDataSource,
	}
}
