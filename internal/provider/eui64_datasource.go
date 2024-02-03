package provider

import (
	"context"
	"fmt"
	"net"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/mdlayher/netx/eui64"
)

var _ datasource.DataSource = EUI64DataSource{}

type EUI64DataSource struct{}

type EUI64DataSourceModel struct {
	MacAddress  types.String `tfsdk:"mac_address"`
	Prefix      types.String `tfsdk:"prefix"`
	IPV6Address types.String `tfsdk:"ipv6_address"`
}

func NewEUIDataSource() datasource.DataSource {
	return EUI64DataSource{}
}

func (d EUI64DataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_eui64_address"
}

func (d EUI64DataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Generates a EUI64 IPv6 address from a mac address",

		Attributes: map[string]schema.Attribute{
			"mac_address": schema.StringAttribute{
				MarkdownDescription: "MAC address to compute the EUI64 for",
				Required:            true,
			},
			"prefix": schema.StringAttribute{
				MarkdownDescription: "Prefix without prefixlength to use as the network part of the address. Default to `fe80::` if missing",
				Optional:            true,
			},
			"ipv6_address": schema.StringAttribute{
				MarkdownDescription: "Resulting IPv6 address of the computation",
				Computed:            true,
			},
		},
	}
}

func (d EUI64DataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
}

func (d EUI64DataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data EUI64DataSourceModel

	// read config
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	mac, err := net.ParseMAC(data.MacAddress.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to parse mac address, got error: %v", err))
		return
	}

	prefixInput := data.Prefix.ValueString()
	if prefixInput == "" {
		prefixInput = "fe80::"
	}

	prefix := net.ParseIP(prefixInput)
	if prefix == nil {
		resp.Diagnostics.AddError("Client Error", "Unable to parse prefix")
		return
	}

	addr, err := eui64.ParseMAC(prefix, mac)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to generate EUI-64, got error: %v", err))
		return
	}
	data.IPV6Address = types.StringValue(addr.String())

	// set result
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
