package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAzureResourceGroup(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_resource_group",
		Description: "Azure Resource Group",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			Hydrate:    opengovernance.GetResourceGroup,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceGroupNotFound"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListResourceGroup,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The friendly name that identifies the resource group.",
				Transform:   transform.FromField("Description.Group.Name")},
			{
				Name:        "id",
				Description: "Contains ID to identify a resource group uniquely.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Group.ID")},
			{
				Name:        "provisioning_state",
				Description: "Current state of the resource group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Group.Properties.ProvisioningState")},
			{
				Name:        "managed_by",
				Description: "Contains ID of the resource that manages this resource group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Group.ManagedBy")},
			{
				Name:        "type",
				Description: "Type of the resource group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Group.Type")},
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Group.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Group.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.Group.ID").Transform(idToAkas),
			},
			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Group.Location").Transform(toLower),
			},
		}),
	}
}
