package azuread

import (
	"context"
	"github.com/opengovern/og-describer-entraid/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAzureAdDevice(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azuread_device",
		Description: "Represents an Azure AD device.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListAdDevice,
			KeyColumns: plugin.KeyColumnSlice{
				// Key fields
				{Name: "display_name", Require: plugin.Optional},
				{Name: "account_enabled", Require: plugin.Optional, Operators: []string{"<>", "="}},
				{Name: "operating_system", Require: plugin.Optional},
				{Name: "operating_system_version", Require: plugin.Optional},
				{Name: "profile_type", Require: plugin.Optional},
				{Name: "trust_type", Require: plugin.Optional},
				{Name: "filter", Require: plugin.Optional},
			},
		},

		Columns: azureOGColumns([]*plugin.Column{

			{Name: "id", Type: proto.ColumnType_STRING, Description: "The unique identifier for the device. Inherited from directoryObject.", Transform: transform.FromField("Description.Id")},
			{Name: "display_name", Type: proto.ColumnType_STRING, Description: "The name displayed for the device.", Transform: transform.FromField("Description.DisplayName")},
			{Name: "account_enabled", Type: proto.ColumnType_BOOL, Description: "True if the account is enabled; otherwise, false.", Transform: transform.FromField("Description.AccountEnabled")},
			{Name: "device_id", Type: proto.ColumnType_STRING, Description: "Unique identifier set by Azure Device Registration Service at the time of registration.", Transform: transform.FromField("Description.DeviceId")},
			{Name: "approximate_last_sign_in_date_time", Type: proto.ColumnType_TIMESTAMP, Description: "The timestamp type represents date and time information using ISO 8601 format and is always in UTC time.", Transform: transform.FromField("Description.ApproximateLastSignInDateTime")},

			// Other fields
			{Name: "filter", Type: proto.ColumnType_STRING, Transform: transform.FromQual("filter"), Description: "Odata query to search for resources."},
			{Name: "is_compliant", Type: proto.ColumnType_BOOL, Description: "True if the device is compliant; otherwise, false.", Transform: transform.FromField("Description.IsCompliant")},
			{Name: "is_managed", Type: proto.ColumnType_BOOL, Description: "True if the device is managed; otherwise, false.", Transform: transform.FromField("Description.IsManaged")},
			{Name: "mdm_app_id", Type: proto.ColumnType_STRING, Description: "Application identifier used to register device into MDM.", Transform: transform.FromField("Description.MdmAppId")},
			{Name: "operating_system", Type: proto.ColumnType_STRING, Description: "The type of operating system on the device.", Transform: transform.FromField("Description.OperatingSystem")},
			{Name: "operating_system_version", Type: proto.ColumnType_STRING, Description: "The version of the operating system on the device.", Transform: transform.FromField("Description.OperatingSystemVersion")},
			{Name: "profile_type", Type: proto.ColumnType_STRING, Description: "A string value that can be used to classify device types.", Transform: transform.FromField("Description.ProfileType")},
			{Name: "trust_type", Type: proto.ColumnType_STRING, Description: "Type of trust for the joined device. Possible values: Workplace (indicates bring your own personal devices), AzureAd (Cloud only joined devices), ServerAd (on-premises domain joined devices joined to Azure AD).", Transform: transform.FromField("Description.TrustType")},

			// JSON fields
			{Name: "extension_attributes", Type: proto.ColumnType_JSON, Description: "Contains extension attributes 1-15 for the device. The individual extension attributes are not selectable. These properties are mastered in cloud and can be set during creation or update of a device object in Azure AD.", Transform: transform.FromField("Description.ExtensionAttributes")},
			{Name: "member_of", Type: proto.ColumnType_JSON, Description: "A list the groups and directory roles that the device is a direct member of.", Transform: transform.FromField("Description.MemberOf")},

			// Standard columns
			{Name: "title", Type: proto.ColumnType_STRING, Description: ColumnDescriptionTitle, Transform: transform.From(adDeviceTitle)},
			{Name: "tenant_id", Type: proto.ColumnType_STRING, Description: ColumnDescriptionTenant, Transform: transform.FromField("Description.TenantID")},
		}),
	}
}

//// TRANSFORM FUNCTIONS

func adDeviceTitle(_ context.Context, d *transform.TransformData) (interface{}, error) {
	data := d.HydrateItem.(opengovernance.AdDevice).Description
	title := data.DisplayName
	if title == nil {
		title = data.DeviceId
	}

	return title, nil
}
