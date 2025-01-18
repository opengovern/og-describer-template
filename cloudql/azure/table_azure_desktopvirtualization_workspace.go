package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAzureDesktopVirtualizationWorkspace(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_desktopvirtualization_workspace",
		Description: "Azure DesktopVirtualization Workspace",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"), //TODO: change this to the primary key columns in model.go
			Hydrate:    opengovernance.GetDesktopVirtualizationWorkspace,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListDesktopVirtualizationWorkspace,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "The id of the workspace.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workspace.ID")},
			{
				Name:        "name",
				Description: "The name of the workspace.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workspace.Name")},
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workspace.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Workspace.Tags"), // probably needs a transform function
			},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Workspace.ID").Transform(idToAkas), // or generate it below (keep the Transform(arnToTurbotAkas) or use Transform(transform.EnsureStringArray))
			},
		}),
	}
}
