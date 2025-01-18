package render

import (
	"context"
	"github.com/opengovern/og-describer-render/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableRenderEnvironment(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "render_environment",
		Description: "Information about environment descriptions, including ID, name, project details, and associated resources.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListEnvironment,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    opengovernance.GetEnvironment,
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "The unique identifier for the environment.", Transform: transform.FromField("Description.ID")},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The name of the environment.", Transform: transform.FromField("Description.Name")},
			{Name: "projectId", Type: proto.ColumnType_STRING, Description: "The ID of the project associated with the environment.", Transform: transform.FromField("Description.ProjectID")},
			{Name: "databasesIds", Type: proto.ColumnType_JSON, Description: "A list of database IDs associated with the environment.", Transform: transform.FromField("Description.DatabasesIDs")},
			{Name: "redisIds", Type: proto.ColumnType_JSON, Description: "A list of Redis instance IDs associated with the environment.", Transform: transform.FromField("Description.RedisIDs")},
			{Name: "serviceIds", Type: proto.ColumnType_JSON, Description: "A list of service IDs associated with the environment.", Transform: transform.FromField("Description.ServiceIDs")},
			{Name: "envGroupIds", Type: proto.ColumnType_JSON, Description: "A list of environment group IDs associated with the environment.", Transform: transform.FromField("Description.EnvGroupIDs")},
			{Name: "protectedStatus", Type: proto.ColumnType_STRING, Description: "The protected status of the environment.", Transform: transform.FromField("Description.ProtectedStatus")},
		}),
	}
}
