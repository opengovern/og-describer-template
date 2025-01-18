package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAzureSecurityCenterContact(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_security_center_contact",
		Description: "Azure Security Center Contact",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			Hydrate:    opengovernance.GetSecurityCenterContact,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListSecurityCenterContact,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "The resource id.",
				Transform:   transform.FromField("Description.Contact.ID")},
			{
				Name:        "name",
				Description: "The resource name.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Contact.Name")},
			{
				Name:        "type",
				Description: "The resource type.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Contact.Type")},
			{
				Name:        "email",
				Description: "The email of this security contact.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Contact.Properties.Emails")},
			{
				Name:        "phone",
				Description: "The phone number of this security contact.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Contact.Properties.Phone")},
			{
				Name:        "alert_notifications",
				Description: "Whether to send security alerts notifications to the security contact.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Contact.Properties.AlertNotifications")},
			{
				Name:        "alerts_to_admins",
				Description: "Whether to send security alerts notifications to subscription admins.",
				Type:        proto.ColumnType_STRING,

				// Steampipe standard columns
				Transform: transform.FromField("Description.Contact.Properties.AlertNotifications")},

			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Contact.Name")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				//// LIST FUNCTION

				Transform: transform.

					// Check if context has been cancelled or if the limit has been hit (if specified)
					// if there is a limit, it will return the number of rows required to reach this limit
					FromField("Description.Contact.ID").Transform(idToAkas),
			},
		}),
	}
}

// Check if context has been cancelled or if the limit has been hit (if specified)
// if there is a limit, it will return the number of rows required to reach this limit

//// HYDRATE FUNCTIONS
