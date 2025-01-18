package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION ////

func tableAzureAKSVersion(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_kubernetes_service_version",
		Description: "Azure Kubernetes Service Version",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListKubernetesServiceVersion,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "version",
				Type:        proto.ColumnType_STRING,
				Description: "The major.minor version of Kubernetes release.",
				Transform:   transform.FromField("Description.Version.Version")},
			{
				Name:        "is_preview",
				Description: "Whether Kubernetes version is currently in preview.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Version.IsPreview")},
			{
				Name:        "capabilities",
				Description: "Capabilities on this Kubernetes version.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Version.Capabilities")},
			{
				Name:        "patch_versions",
				Description: "Patch versions of Kubernetes release.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Version.PatchVersions")},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Version.Version")},

			// Azure standard columns
			{
				Name:        "location",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromQual("Location"),
			},
		}),
	}
}
