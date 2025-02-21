package github

import (
	opengovernance "github.com/opengovern/og-describer-github/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func sharedRepositoryColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "repository_full_name",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("Description.RepositoryFullName"),
			Description: "repository full name",
		},

		{
			Name:        "github_repo_id",
			Type:        proto.ColumnType_INT,
			Transform:   transform.FromField("Description.GitHubRepoID"),
			Description: "Unique identifier of the GitHub repository.",
		},
		{
			Name:        "id",
			Type:        proto.ColumnType_INT,
			Transform:   transform.FromField("Description.GitHubRepoID"),
			Description: "Unique identifier of the GitHub repository.",
		},
		{
			Name:        "node_id",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("Description.NodeID"),
			Description: "Node ID of the repository.",
		},
		{
			Name:        "name",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("Description.Name"),
			Description: "Name of the repository.",
		},
		{
			Name:        "name_with_owner",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("Description.NameWithOwner"),
			Description: "Full name of the repository including the owner.",
		},
		{
			Name:        "description",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("Description.Description"),
			Description: "Description of the repository.",
		},
		{
			Name:        "created_at",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("Description.CreatedAt"),
			Description: "Timestamp when the repository was created.",
		},
		{
			Name:        "updated_at",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("Description.UpdatedAt"),
			Description: "Timestamp when the repository was last updated.",
		},
		{
			Name:        "archived_at",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("Description.ArchivedAt"),
			Description: "Timestamp when the repository was archived.",
		},
		{
			Name:        "pushed_at",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("Description.PushedAt"),
			Description: "Timestamp when the repository was last pushed.",
		},
		{
			Name:        "is_active",
			Type:        proto.ColumnType_BOOL,
			Transform:   transform.FromField("Description.IsActive"),
			Description: "Indicates if the repository is active.",
		},
		{
			Name:        "is_empty",
			Type:        proto.ColumnType_BOOL,
			Transform:   transform.FromField("Description.IsEmpty"),
			Description: "Indicates if the repository is empty.",
		},
		{
			Name:        "is_fork",
			Type:        proto.ColumnType_BOOL,
			Transform:   transform.FromField("Description.IsFork"),
			Description: "Indicates if the repository is a fork.",
		},
		{
			Name:        "is_security_policy_enabled",
			Type:        proto.ColumnType_BOOL,
			Transform:   transform.FromField("Description.IsSecurityPolicyEnabled"),
			Description: "Indicates if the repository has a security policy enabled.",
		},
		{
			Name:        "owner_login",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("Description.Owner.Login"),
			Description: "Owner login.",
		},
		{
			Name:        "contact_links",
			Type:        proto.ColumnType_JSON,
			Transform:   transform.FromField("Description.ContactLinks"),
			Description: "ContactLinks",
		},
		{
			Name:        "homepage_url",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("Description.HomepageURL"),
			Description: "Homepage URL of the repository.",
		},
		{
			Name:        "license_info",
			Type:        proto.ColumnType_JSON,
			Transform:   transform.FromField("Description.LicenseInfo"),
			Description: "License information of the repository.",
		},
		{
			Name:        "topics",
			Type:        proto.ColumnType_JSON,
			Transform:   transform.FromField("Description.Topics"),
			Description: "List of topics associated with the repository.",
		},
		{
			Name:        "visibility",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("Description.Visibility"),
			Description: "Visibility status of the repository.",
		},
		{
			Name:        "default_branch_ref",
			Type:        proto.ColumnType_JSON,
			Transform:   transform.FromField("Description.DefaultBranchRef"),
			Description: "Details of the default branch of the repository.",
		},
		{
			Name:        "permissions",
			Type:        proto.ColumnType_JSON,
			Transform:   transform.FromField("Description.Permissions"),
			Description: "Permissions associated with the repository.",
		},
		{
			Name:        "organization_name",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("Description.Organization"),
			Description: "organization name",
		},
		{
			Name:        "organization",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("Description.Organization"),
			Description: "Organization details of the repository.",
		},

		{
			Name:        "parent",
			Type:        proto.ColumnType_JSON,
			Transform:   transform.FromField("Description.Parent"),
			Description: "Parent repository details if the repository is forked.",
		},
		{
			Name:        "source",
			Type:        proto.ColumnType_JSON,
			Transform:   transform.FromField("Description.Source"),
			Description: "Source repository details if the repository is forked.",
		},
		{
			Name:        "primary_language",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("Description.PrimaryLanguage"),
			Description: "Primary language used in the repository.",
		},
		{
			Name:        "languages",
			Type:        proto.ColumnType_JSON,
			Transform:   transform.FromField("Description.Languages"),
			Description: "Languages used in the repository along with their usage statistics.",
		},
		{
			Name:        "hooks",
			Type:        proto.ColumnType_JSON,
			Transform:   transform.FromField("Description.Hooks"),
			Description: "Hooks.",
		},
		{
			Name:        "security_settings",
			Type:        proto.ColumnType_JSON,
			Transform:   transform.FromField("Description.SecuritySettings"),
			Description: "Security settings of the repository.",
		},
		{
			Name:        "interaction_ability",
			Type:        proto.ColumnType_JSON,
			Transform:   transform.FromField("Description.InteractionAbility"),
			Description: "InteractionAbility.",
		},
		{
			Name:        "code_of_conduct",
			Type:        proto.ColumnType_JSON,
			Transform:   transform.FromField("Description.CodeOfConduct"),
			Description: "CodeOfConduct.",
		},
		{
			Name:        "issue_templates",
			Type:        proto.ColumnType_JSON,
			Transform:   transform.FromField("Description.IssueTemplates"),
			Description: "issue_templates.",
		},
		{
			Name:        "possible_commit_emails",
			Type:        proto.ColumnType_JSON,
			Transform:   transform.FromField("Description.PossibleCommitEmails"),
			Description: "PossibleCommitEmails.",
		},
		{
			Name:        "pull_request_templates",
			Type:        proto.ColumnType_JSON,
			Transform:   transform.FromField("Description.PullRequestTemplates"),
			Description: "PullRequestTemplates.",
		},
		{
			Name:        "security_policy_url",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("Description.SecurityPolicyUrl"),
			Description: "SecurityPolicyUrl.",
		},
		{
			Name:        "projects_url",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("Description.ProjectsUrl"),
			Description: "ProjectsUrl.",
		},
		{
			Name:        "lock_reason",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("Description.LockReason"),
			Description: "LockReason.",
		},
		{
			Name:        "ssh_url",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("Description.RepoURLs.SSHURL"),
			Description: "ssh_url.",
		},
		{
			Name:        "url",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("Description.RepoURLs.GitURL"),
			Description: "GitURL.",
		},
		{
			Name:        "allow_update_branch",
			Type:        proto.ColumnType_BOOL,
			Transform:   transform.FromField("Description.RepositorySettings.AllowUpdateBranch"),
			Description: "Allow update branch.",
		},
		{
			Name:        "uses_custom_open_graph_image",
			Type:        proto.ColumnType_BOOL,
			Transform:   transform.FromField("Description.UsesCustomOpenGraphImage"),
			Description: "UsesCustomOpenGraphImage.",
		},
		{
			Name:        "auto_merge_allowed",
			Type:        proto.ColumnType_BOOL,
			Transform:   transform.FromField("Description.RepositorySettings.AllowAutoMerge"),
			Description: "Allow auto merge.",
		},
		{
			Name:        "is_user_configuration_repository",
			Type:        proto.ColumnType_BOOL,
			Transform:   transform.FromField("Description.IsUserConfigurationRepository"),
			Description: "IsUserConfigurationRepository.",
		},
		{
			Name:        "rebase_merge_allowed",
			Type:        proto.ColumnType_BOOL,
			Transform:   transform.FromField("Description.RepositorySettings.AllowRebaseMerge"),
			Description: "Allow rebase merge.",
		},
		{
			Name:        "delete_branch_on_merge",
			Type:        proto.ColumnType_BOOL,
			Transform:   transform.FromField("Description.RepositorySettings.DeleteBranchOnMerge"),
			Description: "Delete branch on merge.",
		},
		{
			Name:        "fork_count",
			Type:        proto.ColumnType_INT,
			Transform:   transform.FromField("Description.Metrics.Forks"),
			Description: "Forks.",
		},
		{
			Name:        "disk_usage",
			Type:        proto.ColumnType_INT,
			Transform:   transform.FromField("Description.DiskUsage"),
			Description: "DiskUsage.",
		},
		{
			Name:        "network_count",
			Type:        proto.ColumnType_INT,
			Transform:   transform.FromField("Description.NetworkCount"),
			Description: "NetworkCount.",
		},
		{
			Name:        "open_issues_total_count",
			Type:        proto.ColumnType_INT,
			Transform:   transform.FromField("Description.OpenIssuesCount"),
			Description: "OpenIssuesCount.",
		},
		{
			Name:        "watchers_total_count",
			Type:        proto.ColumnType_INT,
			Transform:   transform.FromField("Description.WatchersCount"),
			Description: "WatchersCount.",
		},
		{
			Name:        "repository_topics_total_count",
			Type:        proto.ColumnType_INT,
			Transform:   transform.FromField("Description.TopicsTotalCount"),
			Description: "WatchersCount.",
		},
		{
			Name:        "forking_allowed",
			Type:        proto.ColumnType_BOOL,
			Transform:   transform.FromField("Description.RepositorySettings.ForkingAllowed"),
			Description: "ForkingAllowed.",
		},
		{
			Name:        "has_discussions_enabled",
			Type:        proto.ColumnType_BOOL,
			Transform:   transform.FromField("Description.RepositorySettings.HasDiscussionsEnabled"),
			Description: "ForkingAllowed.",
		},
		{
			Name:        "has_downloads",
			Type:        proto.ColumnType_BOOL,
			Transform:   transform.FromField("Description.RepositorySettings.HasDownloads"),
			Description: "HasDownloads.",
		},
		{
			Name:        "has_issues_enabled",
			Type:        proto.ColumnType_BOOL,
			Transform:   transform.FromField("Description.RepositorySettings.HasIssuesEnabled"),
			Description: "HasIssuesEnabled.",
		},
		{
			Name:        "has_pages",
			Type:        proto.ColumnType_BOOL,
			Transform:   transform.FromField("Description.RepositorySettings.HasPages"),
			Description: "HasPages.",
		},
		{
			Name:        "has_projects_enabled",
			Type:        proto.ColumnType_BOOL,
			Transform:   transform.FromField("Description.RepositorySettings.HasProjectsEnabled"),
			Description: "HasProjectsEnabled.",
		},
		{
			Name:        "has_vulnerability_alerts_enabled",
			Type:        proto.ColumnType_BOOL,
			Transform:   transform.FromField("Description.SecuritySettings.VulnerabilityAlertsEnabled"),
			Description: "HasProjectsEnabled.",
		},
		{
			Name:        "has_wiki_enabled",
			Type:        proto.ColumnType_BOOL,
			Transform:   transform.FromField("Description.RepositorySettings.HasWikiEnabled"),
			Description: "HasWikiEnabled.",
		},
		{
			Name:        "is_archived",
			Type:        proto.ColumnType_BOOL,
			Transform:   transform.FromField("Description.RepositorySettings.Archived"),
			Description: "Archived.",
		},
		{
			Name:        "is_disabled",
			Type:        proto.ColumnType_BOOL,
			Transform:   transform.FromField("Description.RepositorySettings.Disabled"),
			Description: "Archived.",
		},
		{
			Name:        "is_locked",
			Type:        proto.ColumnType_BOOL,
			Transform:   transform.FromField("Description.RepositorySettings.Locked"),
			Description: "Locked.",
		},
		{
			Name:        "is_template",
			Type:        proto.ColumnType_BOOL,
			Transform:   transform.FromField("Description.RepositorySettings.IsTemplate"),
			Description: "IsTemplate.",
		},
		{
			Name:        "is_private",
			Type:        proto.ColumnType_BOOL,
			Transform:   transform.FromField("Description.RepositorySettings.IsPrivate"),
			Description: "IsPrivate.",
		},
		{
			Name:        "is_mirror",
			Type:        proto.ColumnType_BOOL,
			Transform:   transform.FromField("Description.RepositorySettings.IsMirror"),
			Description: "IsMirror.",
		},
		{
			Name:        "is_in_organization",
			Type:        proto.ColumnType_BOOL,
			Transform:   transform.FromField("Description.RepositorySettings.IsInOrganization"),
			Description: "IsInOrganization.",
		},
		{
			Name:        "is_blank_issues_enabled",
			Type:        proto.ColumnType_BOOL,
			Transform:   transform.FromField("Description.RepositorySettings.BlankIssuesEnabled"),
			Description: "BlankIssuesEnabled.",
		},
		{
			Name:        "merge_commit_allowed",
			Type:        proto.ColumnType_BOOL,
			Transform:   transform.FromField("Description.RepositorySettings.MergeCommitAllowed"),
			Description: "MergeCommitAllowed.",
		},
		{
			Name:        "merge_commit_message",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("Description.RepositorySettings.MergeCommitMessage"),
			Description: "MergeCommitAllowed.",
		},
		{
			Name:        "merge_commit_title",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("Description.RepositorySettings.MergeCommitTitle"),
			Description: "MergeCommitAllowed.",
		},
		{
			Name:        "mirror_url",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("Description.RepositorySettings.MirrorURL"),
			Description: "MirrorURL.",
		},
		{
			Name:        "squash_merge_allowed",
			Type:        proto.ColumnType_BOOL,
			Transform:   transform.FromField("Description.RepositorySettings.SquashMergeAllowed"),
			Description: "MirrorURL.",
		},
		{
			Name:        "web_commit_signoff_required",
			Type:        proto.ColumnType_BOOL,
			Transform:   transform.FromField("Description.RepositorySettings.WebCommitSignoffRequired"),
			Description: "WebCommitSignoffRequired.",
		},
		{
			Name:        "squash_merge_commit_message",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("Description.RepositorySettings.SquashMergeCommitMessage"),
			Description: "MirrorURL.",
		},
		{
			Name:        "squash_merge_commit_title",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("Description.RepositorySettings.SquashMergeCommitTitle"),
			Description: "MirrorURL.",
		},
		{
			Name:        "stargazer_count",
			Type:        proto.ColumnType_INT,
			Transform:   transform.FromField("Description.Metrics.Stargazers"),
			Description: "Stargazers.",
		},
		{
			Name:        "subscribers_count",
			Type:        proto.ColumnType_INT,
			Transform:   transform.FromField("Description.Metrics.Subscribers"),
			Description: "Subscribers.",
		},
	}
}

func tableGitHubRepository() *plugin.Table {
	return &plugin.Table{
		Name:        "github_repository",
		Description: "GitHub Repositories contain all of your project's files and each file's revision history.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListRepository,
		},
		Columns: commonColumns(sharedRepositoryColumns()),
	}
}
