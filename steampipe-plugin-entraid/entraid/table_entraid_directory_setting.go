package entraid

import (
	"context"
	"github.com/opengovern/og-describer-entraid/pkg/sdk/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableEntraIdDirectorySetting(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "entraid_directory_setting",
		Description: "Represents the configurations that can be used to customize the tenant-wide and object-specific restrictions and allowed behavior",
		Get: &plugin.GetConfig{
			Hydrate:    opengovernance.GetAdDirectorySetting,
			KeyColumns: plugin.AllColumns([]string{"id", "name"}),
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListAdDirectorySetting,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "display_name",
				Type:        proto.ColumnType_STRING,
				Description: "Display name of this group of settings, which comes from the associated template.",
				Transform:   transform.FromField("Description.DisplayName")},
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "Unique identifier for these settings.",
				Transform:   transform.FromField("Description.Id")},
			{
				Name:        "template_id",
				Type:        proto.ColumnType_STRING,
				Description: "Unique identifier for the template used to create this group of settings.",
				Transform:   transform.FromField("Description.TemplateId")},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "Name of the setting.",
				Transform:   transform.FromField("Description.Name")},
			{
				Name:        "value",
				Type:        proto.ColumnType_STRING,
				Description: "Value of the setting.",
				Transform:   transform.FromField("Description.Value")},

			// Standard columns
			{
				Name:        "title",
				Type:        proto.ColumnType_STRING,
				Description: ColumnDescriptionTitle,
				Transform:   transform.FromField("Description.Name")},
			{
				Name:        "tenant_id",
				Type:        proto.ColumnType_STRING,
				Description: ColumnDescriptionTenant,
				Transform:   transform.FromField("Description.TenantID")},
		}),
	}
}
