package github

import (
	opengovernance "github.com/opengovern/og-describer-github/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func gitHubTeamColumns() []*plugin.Column {
	return []*plugin.Column{
		{Name: "organization", Type: proto.ColumnType_STRING, Description: "The organization the team is associated with.",
			Transform: transform.FromField("Organization.Organization")},
		{Name: "slug",
			Transform: transform.FromField("Organization.Slug"),
			Type:      proto.ColumnType_STRING, Description: "The team slug name."},
		{Name: "name",
			Transform: transform.FromField("Organization.Name"),
			Type:      proto.ColumnType_STRING, Description: "The name of the team."},
		{Name: "id", Type: proto.ColumnType_INT, Description: "The ID of the team.",
			Transform: transform.FromField("Organization.ID")},
		{Name: "node_id", Type: proto.ColumnType_STRING, Description: "The node id of the team.",
			Transform: transform.FromField("Organization.NodeID")},
		{Name: "description", Type: proto.ColumnType_STRING, Description: "The description of the team.",
			Transform: transform.FromField("Organization.Description")},
		{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "Timestamp when team was created.",
			Transform: transform.FromField("Organization.CreatedAt").Transform(convertTimestamp)},
		{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Description: "Timestamp when team was last updated.",
			Transform: transform.FromField("Organization.UpdatedAt").Transform(convertTimestamp)},
		{Name: "combined_slug", Type: proto.ColumnType_STRING, Description: "The slug corresponding to the organization and the team.",
			Transform: transform.FromField("Organization.CombinedSlug")},
		{Name: "parent_team", Type: proto.ColumnType_JSON, Description: "The teams parent team.",
			Transform: transform.FromField("Organization.ParentTeam")},
		{Name: "privacy", Type: proto.ColumnType_STRING, Description: "The privacy setting of the team (VISIBLE or SECRET).",
			Transform: transform.FromField("Organization.Privacy")},
		{Name: "ancestors_total_count", Type: proto.ColumnType_INT, Description: "Count of ancestors this team has.",
			Transform: transform.FromField("Organization.AncestorsTotalCount")},
		{Name: "child_teams_total_count", Type: proto.ColumnType_INT, Description: "Count of children teams this team has.",
			Transform: transform.FromField("Organization.ChildTeamsTotalCount")},
		{Name: "discussions_total_count", Type: proto.ColumnType_INT, Description: "Count of team discussions.",
			Transform: transform.FromField("Organization.DiscussionsTotalCount")},
		{Name: "invitations_total_count", Type: proto.ColumnType_INT, Description: "Count of outstanding team member invitations for the team.",
			Transform: transform.FromField("Organization.InvitationsTotalCount")},
		{Name: "members_total_count", Type: proto.ColumnType_INT, Description: "Count of team members.",
			Transform: transform.FromField("Organization.MembersTotalCount")},
		{Name: "projects_v2_total_count", Type: proto.ColumnType_INT, Description: "Count of the teams v2 projects.",
			Transform: transform.FromField("Organization.ProjectsV2TotalCount")},
		{Name: "repositories_total_count", Type: proto.ColumnType_INT, Description: "Count of repositories the team has.",
			Transform: transform.FromField("Organization.RepositoriesTotalCount")},
		{Name: "url", Type: proto.ColumnType_STRING, Description: "URL for the team page in GitHub.",
			Transform: transform.FromField("Organization.URL")},
		{Name: "avatar_url", Type: proto.ColumnType_STRING, Description: "URL for teams avatar.",
			Transform: transform.FromField("Organization.AvatarURL")},
		{Name: "discussions_url", Type: proto.ColumnType_STRING, Description: "URL for team discussions.",
			Transform: transform.FromField("Organization.DiscussionsURL")},
		{Name: "edit_team_url", Type: proto.ColumnType_STRING, Description: "URL for editing this team.",
			Transform: transform.FromField("Organization.EditTeamURL")},
		{Name: "members_url", Type: proto.ColumnType_STRING, Description: "URL for team members.",
			Transform: transform.FromField("Organization.MembersURL")},
		{Name: "new_team_url", Type: proto.ColumnType_STRING, Description: "The HTTP URL creating a new team.",
			Transform: transform.FromField("Organization.NewTeamURL")},
		{Name: "repositories_url", Type: proto.ColumnType_STRING, Description: "URL for team repositories.",
			Transform: transform.FromField("Organization.RepositoriesURL")},
		{Name: "teams_url", Type: proto.ColumnType_STRING, Description: "URL for this team's teams.",
			Transform: transform.FromField("Organization.TeamsURL")},
		{Name: "can_administer", Type: proto.ColumnType_BOOL, Description: "If true, current user can administer the team.",
			Transform: transform.FromField("Organization.CanAdminister")},
		{Name: "can_subscribe", Type: proto.ColumnType_BOOL, Description: "If true, current user can subscribe to the team.",
			Transform: transform.FromField("Organization.CanSubscribe")},
		{Name: "subscription", Type: proto.ColumnType_STRING, Description: "Subscription status of the current user to the team.",
			Transform: transform.FromField("Organization.Subscription")},
	}
}

func tableGitHubTeam() *plugin.Table {
	return &plugin.Table{
		Name:        "github_team",
		Description: "GitHub Teams in a given organization. GitHub Teams are groups of organization members that reflect your company or group's structure with cascading access permissions and mentions.",
		List: &plugin.ListConfig{
			KeyColumns:        plugin.SingleColumn("organization"),
			ShouldIgnoreError: isNotFoundError([]string{"404"}),
			Hydrate:           opengovernance.ListTeamMembers(),
		},
		Get: &plugin.GetConfig{
			KeyColumns:        plugin.AllColumns([]string{"organization", "slug"}),
			ShouldIgnoreError: isNotFoundError([]string{"404"}),
			Hydrate:           tableGitHubTeamGet,
		},
		Columns: commonColumns(gitHubTeamColumns()),
	}
}
