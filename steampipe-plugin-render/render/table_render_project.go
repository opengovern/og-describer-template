package render

import (
	"context"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableRenderProject(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "render_project",
		Description: "Information about project descriptions, including ID, name, owner, and associated environments.",
		List: &plugin.ListConfig{
			Hydrate: nil,
		},
		Get: &plugin.GetConfig{
			Hydrate: nil,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "The unique identifier for the project."},
			{Name: "createdAt", Type: proto.ColumnType_TIMESTAMP, Description: "The timestamp of when the project was created."},
			{Name: "updatedAt", Type: proto.ColumnType_TIMESTAMP, Description: "The timestamp of the last update to the project."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The name of the project."},
			{Name: "owner", Type: proto.ColumnType_JSON, Description: "Information about the owner of the project."},
			{Name: "environmentIds", Type: proto.ColumnType_JSON, Description: "A list of environment IDs associated with the project."},
		},
	}
}
