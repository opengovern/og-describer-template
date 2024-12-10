package render

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-render/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableRenderService(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "render_service",
		Description: "Information about service descriptions, including ID, name, environment, and deployment configuration.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListService,
		},
		Get: &plugin.GetConfig{
			Hydrate: opengovernance.GetService,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "The unique identifier for the service."},
			{Name: "autoDeploy", Type: proto.ColumnType_STRING, Description: "Indicates whether the service deploys automatically."},
			{Name: "branch", Type: proto.ColumnType_STRING, Description: "The branch associated with the service."},
			{Name: "buildFilter", Type: proto.ColumnType_JSON, Description: "The build filter associated with the service."},
			{Name: "createdAt", Type: proto.ColumnType_TIMESTAMP, Description: "The timestamp of when the service was created."},
			{Name: "dashboardUrl", Type: proto.ColumnType_STRING, Description: "The URL to the service's dashboard."},
			{Name: "environmentId", Type: proto.ColumnType_STRING, Description: "The ID of the environment associated with the service."},
			{Name: "imagePath", Type: proto.ColumnType_STRING, Description: "The image path used by the service."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The name of the service."},
			{Name: "notifyOnFail", Type: proto.ColumnType_STRING, Description: "Indicates whether to notify on build failure."},
			{Name: "ownerId", Type: proto.ColumnType_STRING, Description: "The ID of the owner of the service."},
			{Name: "registryCredential", Type: proto.ColumnType_JSON, Description: "The registry credentials associated with the service."},
			{Name: "repo", Type: proto.ColumnType_STRING, Description: "The repository associated with the service."},
			{Name: "rootDir", Type: proto.ColumnType_STRING, Description: "The root directory for the service."},
			{Name: "slug", Type: proto.ColumnType_STRING, Description: "The slug associated with the service."},
			{Name: "suspended", Type: proto.ColumnType_STRING, Description: "Indicates whether the service is suspended."},
			{Name: "suspenders", Type: proto.ColumnType_JSON, Description: "A list of suspenders associated with the service."},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "The type of the service."},
			{Name: "updatedAt", Type: proto.ColumnType_TIMESTAMP, Description: "The timestamp of the last update to the service."},
			{Name: "serviceDetails", Type: proto.ColumnType_JSON, Description: "The details of the service."},
		},
	}
}
