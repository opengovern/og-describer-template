package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITON

func tableAzureResourceLink(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_resource_link",
		Description: "Azure Resource Link",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    opengovernance.GetResourceLink,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"MissingSubscription", "404"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListResourceLink,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the resource link.",
				Transform:   transform.FromField("Description.ResourceLink.Name")},
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "The fully qualified ID of the resource link.",
				Transform:   transform.FromField("Description.ResourceLink.ID")},
			{
				Name:        "type",
				Type:        proto.ColumnType_STRING,
				Description: "The resource link type.",
				Transform:   transform.FromField("Description.ResourceLink.Type")},
			{
				Name:        "source_id",
				Type:        proto.ColumnType_STRING,
				Description: "The fully qualified ID of the source resource in the link.",
				Transform:   transform.FromField("Description.ResourceLink.Properties.SourceID")},
			{
				Name:        "target_id",
				Type:        proto.ColumnType_STRING,
				Description: "The fully qualified ID of the target resource in the link.",
				Transform:   transform.FromField("Description.ResourceLink.Properties.TargetID")},
			{
				Name:        "notes",
				Type:        proto.ColumnType_STRING,
				Description: "Notes about the resource link.",

				// Steampipe standard columns
				Transform: transform.FromField("Description.ResourceLink.Properties.Notes")},

			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResourceLink.Name")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				// Azure standard columns

				Transform: transform.FromField("Description.ResourceLink.ID").Transform(idToAkas),
			},

			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,

				Transform: transform.

					//// LIST FUNCTION
					FromField("Description.ResourceLink.Properties.SourceID")},
		}),
	}
}

//// HYDRATE FUNCTION
