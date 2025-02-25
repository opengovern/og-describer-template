package github

import (
	opengovernance "github.com/opengovern/og-describer-github/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func gitHubOrganizationRoleColumns() []*plugin.Column {
	tableCols := []*plugin.Column{
		{
			Name:        "id",
			Type:        proto.ColumnType_STRING,
			Description: "The organization the member is associated with.",
			Transform:   transform.FromField("Description.ID")},
		{
			Name:        "name",
			Type:        proto.ColumnType_STRING,
			Description: "The role this user has in the organization. Returns null if information is not available to viewer.",
			Transform:   transform.FromField("Description.Name")},
		{
			Name:      "description",
			Type:      proto.ColumnType_STRING,
			Transform: transform.FromField("Description.Description")},
		{
			Name:        "permissions",
			Type:        proto.ColumnType_JSON,
			Description: "permissions",
			Transform:   transform.FromField("Description.Permissions")},
		{
			Name:        "source",
			Description: "",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("Description.Source")},
		{
			Name:        "base_role",
			Type:        proto.ColumnType_STRING,
			Description: "",
			Transform:   transform.FromField("Description.BaseRole")},
		{
			Name:        "created_at",
			Type:        proto.ColumnType_TIMESTAMP,
			Description: "",
			Transform:   transform.FromField("Description.CreatedAt")},
		{
			Name:        "updated_at",
			Type:        proto.ColumnType_TIMESTAMP,
			Description: "",
			Transform:   transform.FromField("Description.UpdatedAt")},
	}

	return tableCols
}

func tableGitHubOrganizationRole() *plugin.Table {
	return &plugin.Table{
		Name:        "github_organization_role",
		Description: "GitHub roles for a given organization. GitHub Users are user accounts in GitHub.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListOrganizationRole,
		},
		Columns: commonColumns(gitHubOrganizationRoleColumns()),
	}
}
