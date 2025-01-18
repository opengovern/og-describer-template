package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAzureSecurityCenterAutomation(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_security_center_automation",
		Description: "Azure Security Center Automation",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "resource_group"}),
			Hydrate:    opengovernance.GetSecurityCenterAutomation,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListSecurityCenterAutomation,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "The resource id.",
				Transform:   transform.FromField("Description.Automation.ID"),
			},
			{
				Name:        "name",
				Description: "The resource name.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Automation.Name")},
			{
				Name:        "type",
				Description: "The type of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Automation.Type")},
			{
				Name:        "description",
				Description: "The security automation description.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Automation.Properties.Description")},
			{
				Name:        "is_enabled",
				Description: "Indicates whether the security automation is enabled.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Automation.Properties.IsEnabled")},
			{
				Name:        "kind",
				Description: "Kind of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Automation.Kind")},
			{
				Name:        "etag",
				Description: "Entity tag is used for comparing two or more entities from the same requested resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Automation.Etag")},
			{
				Name:        "actions",
				Description: "A collection of the actions which are triggered if all the configured rules evaluations, within at least one rule set, are true.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Automation.Properties.Actions")},
			{
				Name:        "scopes",
				Description: "A collection of scopes on which the security automations logic is applied. Supported scopes are the subscription itself or a resource group under that subscription. The automation will only apply on defined scopes.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Automation.Properties.Scopes")},
			{
				Name:        "sources",
				Description: "A collection of the source event types which evaluate the security automation set of rules.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Automation.Properties.Sources")},

			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Automation.Name")},
			{
				Name:        "tags",
				Description: "A list of key value pairs that describe the resource.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Automation.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.Automation.ID").Transform(idToAkas),
			},
			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Automation.Location").Transform(toLower),
			},
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ResourceGroup")},
		}),
	}
}
