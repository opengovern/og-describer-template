package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAzureTenant(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_tenant",
		Description: "Azure Tenant",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListTenant,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The display name of the tenant.",
				Transform:   transform.FromField("Description.TenantIDDescription.Name"),
			},
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "The fully qualified ID of the tenant. For example, /tenants/00000000-0000-0000-0000-000000000000.",
				Transform:   transform.FromField("Description.TenantIDDescription.ID"),
			},
			{
				Name:        "tenant_id",
				Type:        proto.ColumnType_STRING,
				Description: "The tenant ID. For example, 00000000-0000-0000-0000-000000000000.",
				Transform:   transform.FromField("Description.TenantIDDescription.TenantID"),
			},
			{
				Name:        "tenant_category",
				Type:        proto.ColumnType_STRING,
				Description: "The tenant category. Possible values include: 'Home', 'ProjectedBy', 'ManagedBy'.",
				Transform:   transform.FromField("TenantCategory"),
			},
			{
				Name:        "country",
				Type:        proto.ColumnType_STRING,
				Description: "Country/region name of the address for the tenant.",
			},
			{
				Name:        "country_code",
				Type:        proto.ColumnType_STRING,
				Description: "Country/region abbreviation for the tenant.",
			},
			{
				Name:        "display_name",
				Type:        proto.ColumnType_STRING,
				Description: "The list of domains for the tenant.",
				Transform:   transform.FromField("Description.TenantIDDescription.Name"),
			},
			{
				Name:        "domains",
				Type:        proto.ColumnType_JSON,
				Description: "The list of domains for the tenant.",
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TenantIDDescription.Name"),
			},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.TenantIDDescription.ID").Transform(idToAkas),
			},
		}),
	}
}
