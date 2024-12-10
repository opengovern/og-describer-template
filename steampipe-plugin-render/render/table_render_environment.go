package render

import (
	"context"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableRenderEnvironment(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "render_environment",
		Description: "Information about environment descriptions, including ID, name, project details, and associated resources.",
		List: &plugin.ListConfig{
			Hydrate: nil,
		},
		Get: &plugin.GetConfig{
			Hydrate: nil,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "The unique identifier for the environment."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The name of the environment."},
			{Name: "projectId", Type: proto.ColumnType_STRING, Description: "The ID of the project associated with the environment."},
			{Name: "databasesIds", Type: proto.ColumnType_JSON, Description: "A list of database IDs associated with the environment."},
			{Name: "redisIds", Type: proto.ColumnType_JSON, Description: "A list of Redis instance IDs associated with the environment."},
			{Name: "serviceIds", Type: proto.ColumnType_JSON, Description: "A list of service IDs associated with the environment."},
			{Name: "envGroupIds", Type: proto.ColumnType_JSON, Description: "A list of environment group IDs associated with the environment."},
			{Name: "protectedStatus", Type: proto.ColumnType_STRING, Description: "The protected status of the environment."},
		},
	}
}
