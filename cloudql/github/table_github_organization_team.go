package github

import (
	opengovernance "github.com/opengovern/og-describer-github/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func gitHubTeamColumns() []*plugin.Column {
	return []*plugin.Column{
		{Name: "id", Type: proto.ColumnType_INT, Description: "The ID of the team.",
			Transform: transform.FromField("Description.ID")},
		{Name: "name",
			Transform: transform.FromField("Description.Name"),
			Type:      proto.ColumnType_STRING, Description: "The name of the team."},
		{Name: "node_id", Type: proto.ColumnType_STRING, Description: "The node id of the team.",
			Transform: transform.FromField("Description.NodeID")},
		{Name: "slug", Type: proto.ColumnType_STRING, Description: "The node id of the team.",
			Transform: transform.FromField("Description.Slug")},
		{Name: "description", Type: proto.ColumnType_STRING, Description: "The description of the team.",
			Transform: transform.FromField("Description.Description")},
		{Name: "privacy", Type: proto.ColumnType_STRING, Description: "The description of the team.",
			Transform: transform.FromField("Description.Privacy")},
		{Name: "notification_setting", Type: proto.ColumnType_STRING, Description: "The description of the team.",
			Transform: transform.FromField("Description.NotificationSetting")},
		{Name: "url", Type: proto.ColumnType_STRING, Description: "The description of the team.",
			Transform: transform.FromField("Description.URL")},
		{Name: "html_url", Type: proto.ColumnType_STRING, Description: "The description of the team.",
			Transform: transform.FromField("Description.HTMLURL")},
		{Name: "permission", Type: proto.ColumnType_STRING, Description: "The description of the team.",
			Transform: transform.FromField("Description.Permission")},
		{Name: "members_count", Type: proto.ColumnType_STRING, Description: "The description of the team.",
			Transform: transform.FromField("Description.MembersCount")},
		{Name: "organization_id", Type: proto.ColumnType_STRING, Description: "The description of the team.",
			Transform: transform.FromField("Description.OrganizationID")},
		{Name: "repos_count", Type: proto.ColumnType_STRING, Description: "The description of the team.",
			Transform: transform.FromField("Description.ReposCount")},
		{Name: "parent_team_id", Type: proto.ColumnType_INT, Description: "The description of the team.",
			Transform: transform.FromField("Description.ParentTeamID")},
		{Name: "team_sync", Type: proto.ColumnType_JSON, Description: "The description of the team.",
			Transform: transform.FromField("Description.TeamSync")},
	}
}

func tableGitHubOrganizationTeam() *plugin.Table {
	return &plugin.Table{
		Name:        "github_organization_team",
		Description: "GitHub Teams in a given organization. GitHub Teams are groups of organization members that reflect your company or group's structure with cascading access permissions and mentions.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListTeamMember,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"organization_name", "slug"}),
			Hydrate:    opengovernance.GetTeamMember,
		},
		Columns: commonColumns(gitHubTeamColumns()),
	}
}
