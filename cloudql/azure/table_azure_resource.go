package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITON

func tableAzureResourceResource(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_resource",
		Description: "Azure Resource",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    opengovernance.GetGenericResource,
			// No error is returned if the resource is not found
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListGenericResource,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "Resource ID.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.GenericResource.ID"),
			},
			{
				Name:        "name",
				Description: "Resource name.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.GenericResource.Name"),
			},
			{
				Name:        "type",
				Description: "Resource type.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.GenericResource.Type"),
			},
			{
				Name:        "created_time",
				Description: "The created time of the resource.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.GenericResource.CreatedTime"),
			},
			{
				Name:        "changed_time",
				Description: "The changed time of the resource.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.GenericResource.ChangedTime"),
			},
			{
				Name:        "identity_principal_id",
				Description: "The principal ID of resource identity.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.GenericResource.Identity.PrincipalID"),
			},
			{
				Name:        "provisioning_state",
				Description: "The provisioning state of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.GenericResource.Identity.ProvisioningState"),
			},
			{
				Name:        "plan_publisher",
				Description: "The plan publisher ID.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.GenericResource.Plan.Publisher"),
			},
			{
				Name:        "plan_name",
				Description: "The plan ID.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.GenericResource.Plan.Name"),
			},
			{
				Name:        "plan_product",
				Description: "The plan offer ID.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.GenericResource.Plan.Product"),
			},
			{
				Name:        "plan_promotion_code",
				Description: "The plan promotion code.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.GenericResource.Plan.PromotionCode"),
			},
			{
				Name:        "plan_version",
				Description: "The plan's version.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.GenericResource.Plan.Version"),
			},
			{
				Name:        "kind",
				Description: "The kind of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.GenericResource.Kind"),
			},
			{
				Name:        "managed_by",
				Description: "ID of the resource that manages this resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.GenericResource.ManagedBy"),
			},
			{
				Name:        "sku",
				Description: "The SKU of the resource.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.GenericResource.SKU"),
			},
			{
				Name:        "identity",
				Description: "The identity of the resource.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.GenericResource.Identity"),
			},
			{
				Name:        "extended_location",
				Description: "Resource extended location.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.GenericResource.ExtendedLocation"),
			},
			{
				Name:        "properties",
				Description: "The resource properties.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.GenericResource.Properties"),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.GenericResource.Name"),
			},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.GenericResource.Tags"),
			},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.GenericResource.ID").Transform(idToAkas),
			},

			// Azure standard columns
			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.GenericResource.Location").Transform(toLower),
			},
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResourceGroup").Transform(extractResourceGroupFromID),
			},
		}),
	}
}
