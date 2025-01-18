package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAzureSecurityCenterJITNetworkAccessPolicy(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_security_center_jit_network_access_policy",
		Description: "Azure Security Center JIT Network Access Policy",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListSecurityCenterJitNetworkAccessPolicy,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The resource name.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.JitNetworkAccessPolicy.Name")},
			{
				Name:        "id",
				Description: "The resource id.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.JitNetworkAccessPolicy.ID")},
			{
				Name:        "type",
				Description: "The resource type.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.JitNetworkAccessPolicy.Type")},
			{
				Name:        "kind",
				Description: "Kind of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.JitNetworkAccessPolicy.Kind")},
			{
				Name:        "provisioning_state",
				Description: "The provisioning state of the Just-in-Time policy.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.JitNetworkAccessPolicy.Properties.ProvisioningState")},
			{
				Name:        "virtual_machines",
				Description: "Configurations for Microsoft.Compute/virtualMachines resource type.",
				Type:        proto.ColumnType_JSON,

				// Steampipe standard columns
				Transform: transform.FromField("Description.JitNetworkAccessPolicy.Properties.VirtualMachines")},

			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.JitNetworkAccessPolicy.Name")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				//// LIST FUNCTION

				Transform: transform.FromField("Description.JitNetworkAccessPolicy.ID").Transform(idToAkas),
			},
		}),
	}
}
