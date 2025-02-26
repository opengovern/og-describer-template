package github

import (
	opengovernance "github.com/opengovern/og-describer-github/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func gitHubOrganizationRoleDefinition() []*plugin.Column {
	tableCols := []*plugin.Column{
		{
			Name:        "id",
			Type:        proto.ColumnType_INT,
			Description: "Affiliation filter - valid values 'ALL' (default), 'OUTSIDE', 'DIRECT'.",
			Transform:   transform.FromField("Description.Id")},
		{
			Name:        "name",
			Type:        proto.ColumnType_STRING,
			Description: "Affiliation filter - valid values 'ALL' (default), 'OUTSIDE', 'DIRECT'.",
			Transform:   transform.FromField("Description.Name")},
		{
			Name:        "description",
			Type:        proto.ColumnType_STRING,
			Description: "Affiliation filter - valid values 'ALL' (default), 'OUTSIDE', 'DIRECT'.",
			Transform:   transform.FromField("Description.Description")},
		{
			Name:        "permissions",
			Type:        proto.ColumnType_JSON,
			Description: "Affiliation filter - valid values 'ALL' (default), 'OUTSIDE', 'DIRECT'.",
			Transform:   transform.FromField("Description.Permissions")},
		{
			Name:        "organization_id",
			Type:        proto.ColumnType_INT,
			Description: "The name of the repository",
			Transform:   transform.FromField("Description.OrganizationId")},
		{
			Name:        "created_at",
			Type:        proto.ColumnType_TIMESTAMP,
			Description: "The permission the collaborator has on the repository.",
			Transform:   transform.FromField("Description.CreatedAt")},
		{
			Name:        "updated_at",
			Type:        proto.ColumnType_TIMESTAMP,
			Description: "The permission the collaborator has on the repository.",
			Transform:   transform.FromField("Description.UpdatedAt")},
		{
			Name:        "source",
			Type:        proto.ColumnType_STRING,
			Description: "The login details of the collaborator.",
			Transform:   transform.FromField("Description.Source")},
		{
			Name:        "base_role",
			Type:        proto.ColumnType_STRING,
			Description: "The login details of the collaborator.",
			Transform:   transform.FromField("Description.BaseRole")},
		{
			Name:        "type",
			Type:        proto.ColumnType_STRING,
			Description: "The login details of the collaborator.",
			Transform:   transform.FromField("Description.Type")},
	}

	return tableCols
}

func tableGitHubOrganizationRoleDefinition() *plugin.Table {
	return &plugin.Table{
		Name:        "github_organization_role_definition",
		Description: "GitHub members for a given organization. GitHub Users are user accounts in GitHub.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListOrganizationRoleDefinition,
		},
		Columns: commonColumns(gitHubOrganizationRoleDefinition()),
	}
}
