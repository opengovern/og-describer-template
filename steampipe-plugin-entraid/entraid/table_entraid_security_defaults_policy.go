package entraid

import (
	"context"
	"github.com/opengovern/og-describer-entraid/pkg/sdk/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableEntraIdSecurityDefaultsPolicy(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "entraid_security_defaults_policy",
		Description: "Represents the Azure Active Directory security defaults policy",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListAdSecurityDefaultsPolicy,
		},

		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "display_name",
				Type:        proto.ColumnType_STRING,
				Description: "Display name for this policy.",
				Transform:   transform.FromField("Description.DisplayName")},
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "Identifier for this policy.",
				Transform:   transform.FromField("Description.Id")},
			{
				Name:        "is_enabled",
				Type:        proto.ColumnType_BOOL,
				Description: "If set to true, Azure Active Directory security defaults is enabled for the tenant.",
				Transform:   transform.FromField("Description.IsEnabled")},
			{
				Name:        "description",
				Type:        proto.ColumnType_STRING,
				Description: "Description for this policy.",
				Transform:   transform.FromField("Description.Description")},

			// Standard columns
			{
				Name:        "title",
				Type:        proto.ColumnType_STRING,
				Description: ColumnDescriptionTitle,
				Transform:   transform.FromField("Description.DisplayName")},
			{
				Name:        "tenant_id",
				Type:        proto.ColumnType_STRING,
				Description: ColumnDescriptionTenant,
				Transform:   transform.FromField("Description.TenantID")},
		}),
	}
}
