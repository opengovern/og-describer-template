package azure

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAzureAppManagedEnvironments(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_app_managedenvironments",
		Description: "Azure App ManagedEnvironments",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"), //TODO: change this to the primary key columns in model.go
			//Hydrate:    opengovernance.GetAppManagedEnvironment,
		},
		List: &plugin.ListConfig{
			//Hydrate: opengovernance.ListAppManagedEnvironment,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "The id of the managedenvironments.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.HostingEnvironment.ID")},
			{
				Name:        "name",
				Description: "The name of the managedenvironments.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.HostingEnvironment.Name")},
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.HostingEnvironment.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				// probably needs a transform function
				Transform: transform.FromField("Description.HostingEnvironment.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				// or generate it below (keep the Transform(arnToTurbotAkas) or use Transform(transform.EnsureStringArray))
				Transform: transform.FromField("Description.HostingEnvironment.ID").Transform(idToAkas),
			},
		}),
	}
}
