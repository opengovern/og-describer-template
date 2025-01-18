package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAzureSecurityCenterSetting(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_security_center_setting",
		Description: "Azure Security Center Setting",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			Hydrate:    opengovernance.GetSecurityCenterSetting,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListSecurityCenterSetting,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "The resource id.",
				Transform:   transform.FromField("Description.Setting.ID")},
			{
				Name:        "name",
				Description: "The resource name.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Setting.Name")},
			{
				Name:        "enabled",
				Description: "Data export setting status.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.ExportSettingStatus")},
			{
				Name:        "type",
				Description: "The resource type.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Setting.Type")},
			{
				Name:        "kind",
				Description: "The kind of the settings string (DataExportSettings).",
				Type:        proto.ColumnType_STRING,

				// Steampipe standard columns
				Transform: transform.FromField("Description.Setting.Kind")},

			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Setting.Name")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Setting.ID").Transform(idToAkas),
			},
		}),
	}
}
