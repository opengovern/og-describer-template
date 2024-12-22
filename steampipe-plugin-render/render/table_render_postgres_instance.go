package render

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-render/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableRenderPostgres(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "render_postgres_instance",
		Description: "Information about PostgreSQL database descriptions, including ID, configuration, status, and timestamps.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListPostgres,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    opengovernance.GetPostgres,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "The unique identifier for the PostgreSQL instance.", Transform: transform.FromField("Description.ID")},
			{Name: "ipAllowList", Type: proto.ColumnType_JSON, Description: "A list of IP addresses allowed to access the PostgreSQL instance.", Transform: transform.FromField("Description.IPAllowList")},
			{Name: "createdAt", Type: proto.ColumnType_TIMESTAMP, Description: "The timestamp of when the PostgreSQL instance was created.", Transform: transform.FromField("Description.CreatedAt")},
			{Name: "updatedAt", Type: proto.ColumnType_TIMESTAMP, Description: "The timestamp of the last update to the PostgreSQL instance.", Transform: transform.FromField("Description.UpdatedAt")},
			{Name: "expiresAt", Type: proto.ColumnType_TIMESTAMP, Description: "The timestamp of when the PostgreSQL instance expires.", Transform: transform.FromField("Description.ExpiresAt")},
			{Name: "databaseName", Type: proto.ColumnType_STRING, Description: "The name of the PostgreSQL database.", Transform: transform.FromField("Description.DatabaseName")},
			{Name: "databaseUser", Type: proto.ColumnType_STRING, Description: "The username for the PostgreSQL database.", Transform: transform.FromField("Description.DatabaseUser")},
			{Name: "environmentId", Type: proto.ColumnType_STRING, Description: "The ID of the environment associated with the PostgreSQL instance.", Transform: transform.FromField("Description.EnvironmentID")},
			{Name: "highAvailabilityEnabled", Type: proto.ColumnType_BOOL, Description: "Indicates whether high availability is enabled for the PostgreSQL instance.", Transform: transform.FromField("Description.HighAvailabilityEnabled")},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The name of the PostgreSQL instance.", Transform: transform.FromField("Description.Name")},
			{Name: "owner", Type: proto.ColumnType_JSON, Description: "Information about the owner of the PostgreSQL instance.", Transform: transform.FromField("Description.Owner")},
			{Name: "plan", Type: proto.ColumnType_STRING, Description: "The plan associated with the PostgreSQL instance.", Transform: transform.FromField("Description.Plan")},
			{Name: "diskSizeGB", Type: proto.ColumnType_INT, Description: "The disk size of the PostgreSQL instance in gigabytes.", Transform: transform.FromField("Description.DiskSizeGB")},
			{Name: "primaryPostgresID", Type: proto.ColumnType_STRING, Description: "The ID of the primary PostgreSQL instance.", Transform: transform.FromField("Description.PrimaryPostgresID")},
			{Name: "region", Type: proto.ColumnType_STRING, Description: "The region where the PostgreSQL instance is located.", Transform: transform.FromField("Description.Region")},
			{Name: "readReplicas", Type: proto.ColumnType_JSON, Description: "A list of read replicas associated with the PostgreSQL instance.", Transform: transform.FromField("Description.ReadReplicas")},
			{Name: "role", Type: proto.ColumnType_STRING, Description: "The role of the PostgreSQL instance (e.g., primary, replica).", Transform: transform.FromField("Description.Role")},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "The status of the PostgreSQL instance.", Transform: transform.FromField("Description.Status")},
			{Name: "version", Type: proto.ColumnType_STRING, Description: "The version of PostgreSQL used in the instance.", Transform: transform.FromField("Description.Version")},
			{Name: "suspended", Type: proto.ColumnType_STRING, Description: "Indicates whether the PostgreSQL instance is suspended.", Transform: transform.FromField("Description.Suspended")},
			{Name: "suspenders", Type: proto.ColumnType_JSON, Description: "A list of suspenders associated with the PostgreSQL instance.", Transform: transform.FromField("Description.Suspenders")},
			{Name: "dashboardUrl", Type: proto.ColumnType_STRING, Description: "The URL of the PostgreSQL instance's dashboard.", Transform: transform.FromField("Description.DashboardURL")},
		},
	}
}
