package github

import (
	opengovernance "github.com/opengovern/og-describer-github/pkg/sdk/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableGitHubCommunityProfile() *plugin.Table {
	return &plugin.Table{
		Name:        "github_community_profile",
		Description: "Community profile information for the given repository.",
		List: &plugin.ListConfig{
			KeyColumns:        plugin.SingleColumn("repository_full_name"),
			Hydrate:           opengovernance.ListCommunityProfile,
			ShouldIgnoreError: isNotFoundError([]string{"404"}),
		},
		Columns: commonColumns([]*plugin.Column{
			{
				Name:        "repository_full_name",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RepoFullName"),
				Description: "Full name of the repository that contains the tag."},
			{
				Name:        "code_of_conduct",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.CodeOfConduct"),
				Description: "Code of conduct for the repository."},
			{
				Name:        "contributing",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Contributing"),
				Description: "Contributing guidelines for the repository."},
			{
				Name:        "issue_templates",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.IssueTemplates"),
				Description: "Issue template for the repository."},
			{
				Name:        "pull_request_templates",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.PullRequestTemplates"),
				Description: "Pull request template for the repository."},
			{
				Name:        "license_info",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.LicenseInfo"),
				Description: "License for the repository."},
			{
				Name:        "readme",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ReadMe"),
				Description: "README for the repository."},
			{
				Name:        "security",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Security"),
				Description: "Security for the repository."},
		}),
	}
}
