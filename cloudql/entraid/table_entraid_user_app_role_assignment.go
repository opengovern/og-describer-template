package entraid

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-entraid/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableEntraIdUserAppRoleAssignment(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "entraid_user_app_role_assignment",
		Description: "Represents an application role granted for a specific application. Includes the users and groups assigned app roles for this application.",
		Get: &plugin.GetConfig{
			Hydrate: opengovernance.GetUserAppRoleAssignment,
			KeyColumns: plugin.KeyColumnSlice{
				{Name: "id", Require: plugin.Required},
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListUserAppRoleAssignment,
		},

		Columns: []*plugin.Column{
			{Name: "id",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Id"),
				Description: "A unique identifier for the appRoleAssignment key."},
			{Name: "app_role_id", Type: proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AppRoleId"),
				Description: "The identifier (id) for the app role which is assigned to the principal. This app role must be exposed in the appRoles property on the resource application's service principal (resourceId). If the resource application has not declared any app roles, a default app role ID of 00000000-0000-0000-0000-000000000000 can be specified to signal that the principal is assigned to the resource app without any specific app roles."},
			{Name: "resource_id", Type: proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResourceId"),
				Description: "The unique identifier (id) for the resource service principal for which the assignment is made."},
			{Name: "resource_display_name", Type: proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResourceDisplayName"),
				Description: "The display name of the resource app's service principal to which the assignment is made."},

			// Other fields
			{Name: "created_date_time", Type: proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.CreatedDateTime"),
				Description: "The time when the app role assignment was created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z."},
			{Name: "deleted_date_time", Type: proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.DeletedDateTime"),
				Description: "The date and time when the app role assignment was deleted. Always null for an appRoleAssignment object that hasn't been deleted."},

			{Name: "principal_id",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PrincipalId"),
				Description: "The unique identifier (id) for the user, security group, or service principal being granted the app role."},
			{
				Name:        "principal_display_name",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PrincipalDisplayName"),
				Description: "The display name of the user, group, or service principal that was granted the app role assignment."},
			{
				Name:        "principal_type",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PrincipalType"),
				Description: "The type of the assigned principal. This can either be User, Group, or ServicePrincipal.",
			},

			// Standard columns
			{
				Name:        "user_id",
				Type:        proto.ColumnType_STRING,
				Description: "The identifier (id) of the user principal.",
				Transform:   transform.FromField("Description.UserId"),
			},
		},
	}
}
