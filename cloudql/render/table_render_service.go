package render

import (
	"context"
	"github.com/opengovern/og-describer-render/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableRenderService(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "render_service",
		Description: "Information about service descriptions, including ID, name, environment, and deployment configuration.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListService,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    opengovernance.GetService,
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "The unique identifier for the service.", Transform: transform.FromField("Description.ID")},
			{Name: "auto_deploy", Type: proto.ColumnType_STRING, Description: "Indicates whether the service deploys automatically.", Transform: transform.FromField("Description.AutoDeploy")},
			{Name: "branch", Type: proto.ColumnType_STRING, Description: "The branch associated with the service.", Transform: transform.FromField("Description.Branch")},
			{Name: "build_filter", Type: proto.ColumnType_JSON, Description: "The build filter associated with the service.", Transform: transform.FromField("Description.BuildFilter")},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "The timestamp of when the service was created.", Transform: transform.FromField("Description.CreatedAt")},
			{Name: "dashboard_url", Type: proto.ColumnType_STRING, Description: "The URL to the service's dashboard.", Transform: transform.FromField("Description.DashboardURL")},
			{Name: "environment_id", Type: proto.ColumnType_STRING, Description: "The ID of the environment associated with the service.", Transform: transform.FromField("Description.EnvironmentID")},
			{Name: "image_path", Type: proto.ColumnType_STRING, Description: "The image path used by the service.", Transform: transform.FromField("Description.ImagePath")},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The name of the service.", Transform: transform.FromField("Description.Name")},
			{Name: "notify_on_fail", Type: proto.ColumnType_STRING, Description: "Indicates whether to notify on build failure.", Transform: transform.FromField("Description.NotifyOnFail")},
			{Name: "owner_id", Type: proto.ColumnType_STRING, Description: "The ID of the owner of the service.", Transform: transform.FromField("Description.OwnerID")},
			{Name: "registry_credential", Type: proto.ColumnType_JSON, Description: "The registry credentials associated with the service.", Transform: transform.FromField("Description.RegistryCredential")},
			{Name: "repo", Type: proto.ColumnType_STRING, Description: "The repository associated with the service.", Transform: transform.FromField("Description.Repo")},
			{Name: "root_dir", Type: proto.ColumnType_STRING, Description: "The root directory for the service.", Transform: transform.FromField("Description.RootDir")},
			{Name: "slug", Type: proto.ColumnType_STRING, Description: "The slug associated with the service.", Transform: transform.FromField("Description.Slug")},
			{Name: "suspended", Type: proto.ColumnType_STRING, Description: "Indicates whether the service is suspended.", Transform: transform.FromField("Description.Suspended")},
			{Name: "suspenders", Type: proto.ColumnType_JSON, Description: "A list of suspenders associated with the service.", Transform: transform.FromField("Description.Suspenders")},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "The type of the service.", Transform: transform.FromField("Description.Type")},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Description: "The timestamp of the last update to the service.", Transform: transform.FromField("Description.UpdatedAt")},
			{Name: "service_details", Type: proto.ColumnType_JSON, Description: "The details of the service.", Transform: transform.FromField("Description.ServiceDetails")},
		}),
	}
}
