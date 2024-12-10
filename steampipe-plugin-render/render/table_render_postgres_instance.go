package render

import (
	"context"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableRenderPostgres(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "render_postgres_instance",
		Description: "Information about PostgreSQL database descriptions, including ID, configuration, status, and timestamps.",
		List: &plugin.ListConfig{
			Hydrate: nil,
		},
		Get: &plugin.GetConfig{
			Hydrate: nil,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "The unique identifier for the PostgreSQL instance."},
			{Name: "ipAllowList", Type: proto.ColumnType_JSON, Description: "A list of IP addresses allowed to access the PostgreSQL instance."},
			{Name: "createdAt", Type: proto.ColumnType_TIMESTAMP, Description: "The timestamp of when the PostgreSQL instance was created."},
			{Name: "updatedAt", Type: proto.ColumnType_TIMESTAMP, Description: "The timestamp of the last update to the PostgreSQL instance."},
			{Name: "expiresAt", Type: proto.ColumnType_TIMESTAMP, Description: "The timestamp of when the PostgreSQL instance expires."},
			{Name: "databaseName", Type: proto.ColumnType_STRING, Description: "The name of the PostgreSQL database."},
			{Name: "databaseUser", Type: proto.ColumnType_STRING, Description: "The username for the PostgreSQL database."},
			{Name: "environmentId", Type: proto.ColumnType_STRING, Description: "The ID of the environment associated with the PostgreSQL instance."},
			{Name: "highAvailabilityEnabled", Type: proto.ColumnType_BOOL, Description: "Indicates whether high availability is enabled for the PostgreSQL instance."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The name of the PostgreSQL instance."},
			{Name: "owner", Type: proto.ColumnType_JSON, Description: "Information about the owner of the PostgreSQL instance."},
			{Name: "plan", Type: proto.ColumnType_STRING, Description: "The plan associated with the PostgreSQL instance."},
			{Name: "diskSizeGB", Type: proto.ColumnType_INT, Description: "The disk size of the PostgreSQL instance in gigabytes."},
			{Name: "primaryPostgresID", Type: proto.ColumnType_STRING, Description: "The ID of the primary PostgreSQL instance."},
			{Name: "region", Type: proto.ColumnType_STRING, Description: "The region where the PostgreSQL instance is located."},
			{Name: "readReplicas", Type: proto.ColumnType_JSON, Description: "A list of read replicas associated with the PostgreSQL instance."},
			{Name: "role", Type: proto.ColumnType_STRING, Description: "The role of the PostgreSQL instance (e.g., primary, replica)."},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "The status of the PostgreSQL instance."},
			{Name: "version", Type: proto.ColumnType_STRING, Description: "The version of PostgreSQL used in the instance."},
			{Name: "suspended", Type: proto.ColumnType_STRING, Description: "Indicates whether the PostgreSQL instance is suspended."},
			{Name: "suspenders", Type: proto.ColumnType_JSON, Description: "A list of suspenders associated with the PostgreSQL instance."},
			{Name: "dashboardUrl", Type: proto.ColumnType_STRING, Description: "The URL of the PostgreSQL instance's dashboard."},
		},
	}
}
