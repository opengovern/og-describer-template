package github

import (
	opengovernance "github.com/opengovern/og-describer-github/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableGitHubTeamMember() *plugin.Table {
	return &plugin.Table{
		Name:        "github_team_member",
		Description: "GitHub members for a given team. GitHub Users are user accounts in GitHub.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListTeamMember,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"id"}),
			Hydrate:    opengovernance.GetTeamMember,
		},
		Columns: commonColumns(gitHubTeamMemberColumns()),
	}
}

func gitHubTeamMemberColumns() []*plugin.Column {
	cols := []*plugin.Column{
		{Name: "slug", Type: proto.ColumnType_STRING, Description: "The team slug name.",
			Transform: transform.FromField("Description.Slug")},
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
		{Name: "company", Type: proto.ColumnType_STRING, Description: "The company on the users profile.",
			Transform: transform.FromField("Description.Company")},
		{Name: "interaction_ability", Type: proto.ColumnType_JSON, Description: "The interaction ability settings for this user.",
			Transform: transform.FromField("Description.InteractionAbility")},
		{Name: "is_site_admin", Type: proto.ColumnType_BOOL, Description: "If true, user is a site administrator.",
			Transform: transform.FromField("Description.IsSiteAdmin")},
		{Name: "location", Type: proto.ColumnType_STRING, Description: "The location of the user.",
			Transform: transform.FromField("Description.Location")},
	}

	return cols
}
