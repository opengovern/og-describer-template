package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION ////

func tableAzureRouteTable(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_route_table",
		Description: "Azure Route Table",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "resource_group"}),
			Hydrate:    opengovernance.GetRouteTables,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "ResourceGroupNotFound", "404"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListRouteTables,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The friendly name that identifies the route table",
				Transform:   transform.FromField("Description.RouteTable.Name")},
			{
				Name:        "id",
				Description: "Contains ID to identify a route table uniquely",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromGo(),
			},
			{
				Name:        "etag",
				Description: "An unique read-only string that changes whenever the resource is updated",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RouteTable.Etag")},
			{
				Name:        "type",
				Description: "Type of the resource",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RouteTable.Type")},
			{
				Name:        "disable_bgp_route_propagation",
				Description: "Indicates Whether to disable the routes learned by BGP on that route table. True means disable.",
				Type:        proto.ColumnType_BOOL,

				Transform: transform.FromField("Description.RouteTable.Properties.DisableBgpRoutePropagation")},
			{
				Name:        "provisioning_state",
				Description: "The provisioning state of the route table resource",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.RouteTable.Properties.ProvisioningState"),
			},
			{
				Name:        "routes",
				Description: "A list of routes contained within a route table",
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.RouteTable.Properties.Routes")},
			{
				Name:        "subnets",
				Description: "A list of references to subnets",
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.RouteTable.Properties.Subnets")},

			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.RouteTable.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.RouteTable.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.RouteTable.ID").Transform(idToAkas),
			},

			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.RouteTable.Location").Transform(toLower),
			},
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,

				Transform: transform.

					// Check if context has been cancelled or if the limit has been hit (if specified)
					// if there is a limit, it will return the number of rows required to reach this limit
					FromField("Description.ResourceGroup").Transform(toLower),
			},
		}),
	}
}

// Check if context has been cancelled or if the limit has been hit (if specified)
// if there is a limit, it will return the number of rows required to reach this limit

//// HYDRATE FUNCTIONS ////

// In some cases resource does not give any notFound error
// instead of notFound error, it returns empty data
