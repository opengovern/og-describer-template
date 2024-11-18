package azuread

import (
	"context"
	"github.com/opengovern/og-describer-entraid/pkg/sdk/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAzureAdDirectoryRole(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azuread_directory_role",
		Description: "Represents an Azure Active Directory (Azure AD) directory role.",
		Get: &plugin.GetConfig{
			Hydrate: opengovernance.GetAdDirectoryRole,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isIgnorableErrorPredicate([]string{"Request_ResourceNotFound", "Invalid object identifier"}),
			},
			KeyColumns: plugin.SingleColumn("id"),
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListAdDirectoryRole,
		},

		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "The unique identifier for the directory role.",
				Transform:   transform.FromField("Description.Id")},
			{
				Name:        "description",
				Type:        proto.ColumnType_STRING,
				Description: "The description for the directory role.",
				Transform:   transform.FromField("Description.Description")},
			{
				Name:        "display_name",
				Type:        proto.ColumnType_STRING,
				Description: "The display name for the directory role.",
				Transform:   transform.FromField("Description.DisplayName")},

			// Other fields
			{
				Name:        "role_template_id",
				Type:        proto.ColumnType_STRING,
				Description: "The id of the directoryRoleTemplate that this role is based on. The property must be specified when activating a directory role in a tenant with a POST operation. After the directory role has been activated, the property is read only.",
				Transform:   transform.FromField("Description.RoleTemplateId")},

			// Json fields
			{
				Name:        "member_ids",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.MemberIds"),
				Description: "Id of the owners of the application. The owners are a set of non-admin users who are allowed to modify this object."},

			// Standard columns
			{
				Name:        "title",
				Type:        proto.ColumnType_STRING,
				Description: ColumnDescriptionTitle,
				Transform:   transform.From(adDirectoryRoleTitle)},
			{
				Name:        "tenant_id",
				Type:        proto.ColumnType_STRING,
				Description: ColumnDescriptionTenant,
				Transform:   transform.FromField("Description.TenantID")},
		}),
	}
}

func adDirectoryRoleTitle(_ context.Context, d *transform.TransformData) (interface{}, error) {
	data := d.HydrateItem.(*opengovernance.AdDirectoryRole)
	if data == nil {
		return nil, nil
	}

	title := data.Description.DisplayName
	if title == nil {
		title = data.Description.Id
	}

	return title, nil
}
