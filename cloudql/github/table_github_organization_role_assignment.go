package github

import (
	opengovernance "github.com/opengovern/og-describer-github/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func gitHubOrganizationRoleAssignment() []*plugin.Column {
	tableCols := []*plugin.Column{
		{
			Name:        "role_id",
			Type:        proto.ColumnType_INT,
			Description: "Affiliation filter - valid values 'ALL' (default), 'OUTSIDE', 'DIRECT'.",
			Transform:   transform.FromField("Description.RoleId")},
		{
			Name:        "organization_id",
			Type:        proto.ColumnType_INT,
			Description: "The name of the repository",
			Transform:   transform.FromField("Description.OrganizationId")},
		{
			Name:        "principal_type",
			Type:        proto.ColumnType_STRING,
			Description: "The name of the repository",
			Transform:   transform.FromField("Description.PrincipalType")},
		{
			Name:        "principal_id",
			Type:        proto.ColumnType_INT,
			Description: "The name of the repository",
			Transform:   transform.FromField("Description.PrincipalId")},
	}

	return tableCols
}

func tableGitHubOrganizationRoleAssignment() *plugin.Table {
	return &plugin.Table{
		Name:        "github_organization_role_assignment",
		Description: "GitHub members for a given organization. GitHub Users are user accounts in GitHub.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListOrganizationRoleAssignment,
		},
		Columns: commonColumns(gitHubOrganizationRoleAssignment()),
	}
}
