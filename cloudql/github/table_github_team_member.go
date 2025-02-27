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
			KeyColumns: plugin.AllColumns([]string{"team_id", "member_principal_id"}),
			Hydrate:    opengovernance.GetTeamMember,
		},
		Columns: commonColumns(gitHubTeamMemberColumns()),
	}
}

func gitHubTeamMemberColumns() []*plugin.Column {
	cols := []*plugin.Column{
		{Name: "team_id", Type: proto.ColumnType_STRING, Description: "The team slug name.",
			Transform: transform.FromField("Description.TeamID")},
		{Name: "member_principal_type", Type: proto.ColumnType_STRING, Description: "The team slug name.",
			Transform: transform.FromField("Description.MemberPrincipalType")},
		{Name: "member_principal_id", Type: proto.ColumnType_STRING, Description: "The team slug name.",
			Transform: transform.FromField("Description.MemberPrincipalID")},
	}

	return cols
}
