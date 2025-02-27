package github

import (
	opengovernance "github.com/opengovern/og-describer-github/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func gitHubOrganizationMemberColumns() []*plugin.Column {
	cols := []*plugin.Column{
		{Name: "has_two_factor_enabled", Type: proto.ColumnType_BOOL,
			Transform: transform.FromField("Description.HasTwoFactorEnabled")},
		{Name: "role", Type: proto.ColumnType_STRING, Description: "The team member's role (MEMBER, MAINTAINER).",
			Transform: transform.FromField("Description.Role")},
		{Name: "login_id", Type: proto.ColumnType_STRING, Description: "Unique identifier for the user login.",
			Transform: transform.FromField("Description.LoginID")},
		{Name: "login", Type: proto.ColumnType_STRING, Description: "The login name of the user.",
			Transform: transform.FromField("Description.Login")},
		{Name: "id", Type: proto.ColumnType_INT, Description: "The ID of the user.",
			Transform: transform.FromField("Description.ID")},
		{Name: "name", Type: proto.ColumnType_STRING, Description: "The name of the user.",
			Transform: transform.FromField("Description.Name")},
		{Name: "node_id", Type: proto.ColumnType_STRING, Description: "The node ID of the user.",
			Transform: transform.FromField("Description.NodeID")},
		{Name: "email", Type: proto.ColumnType_STRING, Description: "The email of the user.",
			Transform: transform.FromField("Description.Email")},
		{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "Timestamp when user was created.",
			Transform: transform.FromField("Description.CreatedAt")},
		{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Description: "Timestamp when user was updated.",
			Transform: transform.FromField("Description.UpdatedAt")},
		{Name: "company", Type: proto.ColumnType_STRING, Description: "The company on the users profile.",
			Transform: transform.FromField("Description.Company")},
		{Name: "is_site_admin", Type: proto.ColumnType_BOOL, Description: "If true, user is a site administrator.",
			Transform: transform.FromField("Description.IsSiteAdmin")},
		{Name: "location", Type: proto.ColumnType_STRING, Description: "The location of the user.",
			Transform: transform.FromField("Description.Location")},
		{Name: "website_url", Type: proto.ColumnType_STRING, Description: "",
			Transform: transform.FromField("Description.WebsiteURL")},
		{Name: "status", Type: proto.ColumnType_BOOL, Description: "",
			Transform: transform.FromField("Description.Status")},
		{Name: "organization_id", Type: proto.ColumnType_INT, Description: "The unique identifier of the app.",
			Transform: transform.FromField("Description.OrganizationID")},
	}

	return cols
}

func tableGitHubOrganizationMember() *plugin.Table {
	return &plugin.Table{
		Name:        "github_organization_member",
		Description: "GitHub members for a given organization. GitHub Users are user accounts in GitHub.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListOrgMembers,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    opengovernance.GetOrgMembers,
		},
		Columns: commonColumns(gitHubOrganizationMemberColumns()),
	}
}
