package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAzureSecurityCenterAutoProvisioning(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_security_center_auto_provisioning",
		Description: "Azure Security Center Auto Provisioning",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			Hydrate:    opengovernance.GetSecurityCenterAutoProvisioning,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListSecurityCenterAutoProvisioning,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "The resource id.",
				Transform:   transform.FromField("Description.AutoProvisioningSetting.ID")},
			{
				Name:        "name",
				Description: "The resource name.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AutoProvisioningSetting.Name")},
			{
				Name:        "type",
				Description: "The resource type.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AutoProvisioningSetting.Type")},
			{
				Name:        "auto_provision",
				Description: "Describes what kind of security agent provisioning action to take. Possible values include: On, Off",
				Type:        proto.ColumnType_STRING,

				// Steampipe standard columns
				Transform: transform.FromField("Description.AutoProvisioningSetting.Properties.AutoProvision")},

			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AutoProvisioningSetting.Name")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				//// LIST FUNCTION

				Transform: transform.

					// Check if context has been cancelled or if the limit has been hit (if specified)
					// if there is a limit, it will return the number of rows required to reach this limit
					FromField("Description.AutoProvisioningSetting.ID").Transform(idToAkas),
			},
		}),
	}
}

// Check if context has been cancelled or if the limit has been hit (if specified)
// if there is a limit, it will return the number of rows required to reach this limit

//// HYDRATE FUNCTIONS
