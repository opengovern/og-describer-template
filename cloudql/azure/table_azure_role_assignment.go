package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAzureIamRoleAssignment(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_role_assignment",
		Description: "Azure Role Assignment",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListRoleAssignment,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The friendly name that identifies the role assignment.",
				Transform:   transform.FromField("Description.RoleAssignment.Name")},
			{
				Name:        "id",
				Description: "Contains ID to identify a role assignment uniquely.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RoleAssignment.ID")},
			{
				Name:        "scope",
				Description: "Current state of the role assignment.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RoleAssignment.Properties.Scope")},
			{
				Name:        "type",
				Description: "Contains the resource type.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RoleAssignment.Type")},
			{
				Name:        "principal_id",
				Description: "Contains the principal id.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RoleAssignment.Properties.PrincipalID")},
			{
				Name:        "principal_type",
				Description: "Principal type of the assigned principal ID.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.RoleAssignment.Properties.PrincipalType"),
			},
			{
				Name:        "created_on",
				Description: "Time it was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.RoleAssignment.Properties.CreatedOn"),
			},
			{
				Name:        "updated_on",
				Description: "Time it was updated.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.RoleAssignment.Properties.UpdatedOn"),
			},
			{
				Name:        "role_definition_id",
				Description: "Name of the assigned role definition.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RoleAssignment.Properties.RoleDefinitionID")},
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RoleAssignment.Name")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				//// LIST FUNCTION

				Transform: transform.

					// Check if context has been cancelled or if the limit has been hit (if specified)
					// if there is a limit, it will return the number of rows required to reach this limit
					FromField("Description.RoleAssignment.ID").Transform(idToAkas),
			},
		}),
	}
}
