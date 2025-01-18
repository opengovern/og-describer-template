package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAzureLocation(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_location",
		Description: "Azure Location",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListLocation,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The location name",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Location.Name")},
			{
				Name:        "display_name",
				Description: "The display name of the location.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Location.DisplayName")},
			{
				Name:        "id",
				Description: "The fully qualified ID of the location.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Location.ID")},
			{
				Name:        "latitude",
				Description: "The latitude of the location.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Location.Latitude")},
			{
				Name:        "longitude",
				Description: "The longitude of the location",
				Type:        proto.ColumnType_JSON,

				// Steampipe standard columns
				Transform: transform.FromField("Description.Location.Longitude")},

			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Location.Name")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				//// LIST FUNCTION

				Transform: transform.FromField("Description.Location.ID").Transform(idToAkas),
			},
		}),
	}
}
