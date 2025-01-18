package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAzureUserEffectiveAccess(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_user_effective_access",
		Description: "Azure User Effective Access",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    opengovernance.GetUserEffectiveAccess,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListUserEffectiveAccess,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "Contains ID to identify a role assignment uniquely.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ResourceID")},
			{
				Name:        "principal_id",
				Description: "User ID",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PrincipalId")},
			{
				Name:        "principal_name",
				Type:        proto.ColumnType_STRING,
				Description: "The friendly name that identifies the role assignment.",
				Transform:   transform.FromField("Description.PrincipalName")},
			{
				Name:        "principal_type",
				Description: "Contains the resource type.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PrincipalType")},
			{
				Name:        "assignment_type",
				Description: "Assignment type (Explicit, GroupAssignment)",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AssignmentType")},
			{
				Name:        "role_definition_id",
				Description: "Name of the assigned role definition.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RoleAssignment.Properties.RoleDefinitionID")},
			{
				Name:        "parent_principal_id",
				Description: "Parent group principal id",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ParentPrincipalId")},
			{
				Name:        "scope",
				Description: "Role scope",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Scope")},
			{
				Name:        "scope_type",
				Description: "Role scope type",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ScopeType")},
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
