package github

import (
	opengovernance "github.com/opengovern/og-describer-github/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func gitHubRepositoryColumns() []*plugin.Column {
	repoColumns := []*plugin.Column{
		{Name: "full_name", Type: proto.ColumnType_STRING,
			Description: "The full name of the repository, including the owner and repo name.",
			Transform:   transform.FromField("Description.NameWithOwner")},
	}
	return append(repoColumns, sharedRepositoryColumns()...)
}

func sharedRepositoryColumns() []*plugin.Column {
	return []*plugin.Column{
		{Name: "id", Type: proto.ColumnType_INT,
			Description: "The numeric ID of the repository.",
			Transform:   transform.FromField("Description.Id")},
		{Name: "node_id", Type: proto.ColumnType_STRING,
			Description: "The node ID of the repository.",
			Transform:   transform.FromField("Description.NodeId")},
		{Name: "name", Type: proto.ColumnType_STRING,
			Description: "The name of the repository.",
			Transform:   transform.FromField("Description.Name")},
		{Name: "allow_update_branch", Type: proto.ColumnType_BOOL,
			Description: "If true, a pull request head branch that is behind its base branch can always be updated even if it is not required to be up to date before merging.",
			Transform:   transform.FromField("Description.AllowUpdateBranch"),
		},
		{Name: "archived_at", Type: proto.ColumnType_TIMESTAMP,
			Description: "Timestamp when repository was archived.",
			Transform:   transform.FromField("Description.ArchivedAt").Transform(convertTimestamp)},
		{Name: "auto_merge_allowed", Type: proto.ColumnType_BOOL,
			Transform:   transform.FromField("Description.AutoMergeAllowed"),
			Description: "If true, auto-merge can be enabled on pull requests in this repository."},
		{Name: "code_of_conduct", Type: proto.ColumnType_JSON,
			Transform:   transform.FromField("Description.CodeOfConduct"),
			Description: "The code of conduct for this repository."},
		{Name: "contact_links", Type: proto.ColumnType_JSON,
			Transform:   transform.FromField("Description.ContactLinks"),
			Description: "List of contact links associated to the repository."},
		{Name: "created_at", Type: proto.ColumnType_TIMESTAMP,
			Transform:   transform.FromField("Description.CreatedAt"),
			Description: "Timestamp when the repository was created."},
		{Name: "default_branch_ref", Type: proto.ColumnType_JSON,
			Transform:   transform.FromField("Description.DefaultBranchRef"),
			Description: "Default ref information."},
		{Name: "delete_branch_on_merge", Type: proto.ColumnType_BOOL,
			Transform:   transform.FromField("Description.DeleteBranchOnMerge"),
			Description: "If true, branches are automatically deleted when merged in this repository."},
		{Name: "description", Type: proto.ColumnType_STRING,
			Transform:   transform.FromField("Description.Description"),
			Description: "The description of the repository."},
		{Name: "disk_usage", Type: proto.ColumnType_INT,
			Transform:   transform.FromField("Description.DiskUsage"),
			Description: "Number of kilobytes this repository occupies on disk."},
		{Name: "fork_count", Type: proto.ColumnType_INT,
			Transform:   transform.FromField("Description.ForkCount"),
			Description: "Number of forks there are of this repository in the whole network."},
		{Name: "forking_allowed", Type: proto.ColumnType_BOOL,
			Transform:   transform.FromField("Description.ForkingAllowed"),
			Description: "If true, repository allows forks."},
		{Name: "funding_links", Type: proto.ColumnType_JSON,
			Transform:   transform.FromField("Description.FundingLinks"),
			Description: "The funding links for this repository."},
		{Name: "has_discussions_enabled", Type: proto.ColumnType_BOOL,
			Transform:   transform.FromField("Description.HasDiscussionsEnabled"),
			Description: "If true, the repository has the Discussions feature enabled."},
		{Name: "has_issues_enabled", Type: proto.ColumnType_BOOL,
			Transform:   transform.FromField("Description.HasIssuesEnabled"),
			Description: "If true, the repository has issues feature enabled."},
		{Name: "has_projects_enabled", Type: proto.ColumnType_BOOL,
			Transform:   transform.FromField("Description.HasProjectsEnabled"),
			Description: "If true, the repository has the Projects feature enabled."},
		{Name: "has_vulnerability_alerts_enabled", Type: proto.ColumnType_BOOL,
			Transform:   transform.FromField("Description.HasVulnerabilityAlertsEnabled"),
			Description: "If true, vulnerability alerts are enabled for the repository."},
		{Name: "has_wiki_enabled", Type: proto.ColumnType_BOOL,
			Transform:   transform.FromField("Description.HasWikiEnabled"),
			Description: "If true, the repository has wiki feature enabled."},
		{Name: "homepage_url", Type: proto.ColumnType_STRING,
			Transform:   transform.FromField("Description.HomepageURL"),
			Description: "The external URL of the repository if set."},
		{Name: "interaction_ability", Type: proto.ColumnType_JSON,
			Transform:   transform.FromField("Description.InteractionAbility"),
			Description: "The interaction ability settings for this repository."},
		{Name: "is_archived", Type: proto.ColumnType_BOOL,
			Transform:   transform.FromField("Description.IsArchived"),
			Description: "If true, the repository is unmaintained (archived)."},
		{Name: "is_blank_issues_enabled", Type: proto.ColumnType_BOOL,
			Transform:   transform.FromField("Description.IsBlankIssuesEnabled"),
			Description: "If true, blank issue creation is allowed."},
		{Name: "is_disabled", Type: proto.ColumnType_BOOL,
			Transform:   transform.FromField("Description.IsDisabled"),
			Description: "If true, this repository disabled."},
		{Name: "is_empty", Type: proto.ColumnType_BOOL,
			Transform:   transform.FromField("Description.IsEmpty"),
			Description: "If true, this repository is empty."},
		{Name: "is_fork", Type: proto.ColumnType_BOOL,
			Transform:   transform.FromField("Description.IsFork"),
			Description: "If true, the repository is a fork."},
		{Name: "is_in_organization", Type: proto.ColumnType_BOOL,
			Transform:   transform.FromField("Description.IsInOrganization"),
			Description: "If true, repository is either owned by an organization, or is a private fork of an organization repository."},
		{Name: "is_locked", Type: proto.ColumnType_BOOL,
			Transform:   transform.FromField("Description.IsLocked"),
			Description: "If true, repository is locked."},
		{Name: "is_mirror", Type: proto.ColumnType_BOOL,
			Transform:   transform.FromField("Description.IsMirror"),
			Description: "If true, the repository is a mirror."},
		{Name: "is_private", Type: proto.ColumnType_BOOL,
			Transform:   transform.FromField("Description.IsPrivate"),
			Description: "If true, the repository is private or internal."},
		{Name: "is_security_policy_enabled", Type: proto.ColumnType_BOOL,
			Transform:   transform.FromField("Description.IsSecurityPolicyEnabled"),
			Description: "If true, repository has a security policy."},
		{Name: "is_template", Type: proto.ColumnType_BOOL,
			Transform:   transform.FromField("Description.IsTemplate"),
			Description: "If true, the repository is a template that can be used to generate new repositories."},
		{Name: "is_user_configuration_repository", Type: proto.ColumnType_BOOL,
			Transform:   transform.FromField("Description.IsUserConfigurationRepository"),
			Description: "If true, this is a user configuration repository."},
		{Name: "issue_templates", Type: proto.ColumnType_JSON,
			Transform:   transform.FromField("Description.IssueTemplates"),
			Description: "A list of issue templates associated to the repository."},
		{Name: "license_info", Type: proto.ColumnType_JSON,
			Transform:   transform.FromField("Description.LicenseInfo"),
			Description: "The license associated with the repository."},
		{Name: "lock_reason", Type: proto.ColumnType_STRING,
			Transform:   transform.FromField("Description.LockReason"),
			Description: "The reason the repository has been locked."},
		{Name: "merge_commit_allowed", Type: proto.ColumnType_BOOL,
			Transform:   transform.FromField("Description.MergeCommitAllowed"),
			Description: "If true, PRs are merged with a merge commit on this repository."},
		{Name: "merge_commit_message", Type: proto.ColumnType_STRING,
			Transform:   transform.FromField("Description.MergeCommitMessage"),
			Description: "How the default commit message will be generated when merging a pull request."},
		{Name: "merge_commit_title", Type: proto.ColumnType_STRING,
			Transform:   transform.FromField("Description.MergeCommitTitle"),
			Description: "How the default commit title will be generated when merging a pull request."},
		{Name: "mirror_url", Type: proto.ColumnType_STRING,
			Transform:   transform.FromField("Description.MirrorURL"),
			Description: "The repository's original mirror URL."},
		{Name: "name_with_owner", Type: proto.ColumnType_STRING,
			Transform:   transform.FromField("Description.NameWithOwner"),
			Description: "The repository's name with owner."},
		{Name: "open_graph_image_url", Type: proto.ColumnType_STRING,
			Transform:   transform.FromField("Description.OpenGraphImageURL"),
			Description: "The image used to represent this repository in Open Graph data."},
		{Name: "owner_login", Type: proto.ColumnType_STRING,
			Transform:   transform.FromField("Description.OwnerLogin"),
			Description: "Login of the repository owner."},
		{Name: "primary_language", Type: proto.ColumnType_JSON,
			Transform:   transform.FromField("Description.PrimaryLanguage"),
			Description: "The primary language of the repository's code."},
		{Name: "projects_url", Type: proto.ColumnType_STRING,
			Transform:   transform.FromField("Description.ProjectsURL"),
			Description: "The URL listing the repository's projects."},
		{Name: "pull_request_templates", Type: proto.ColumnType_JSON,
			Transform:   transform.FromField("Description.PullRequestTemplates"),
			Description: "Returns a list of pull request templates associated to the repository."},
		{Name: "pushed_at", Type: proto.ColumnType_TIMESTAMP,
			Transform:   transform.FromField("Description.PushedAt").NullIfZero().Transform(convertTimestamp),
			Description: "Timestamp when the repository was last pushed to."},
		{Name: "rebase_merge_allowed", Type: proto.ColumnType_BOOL,
			Transform:   transform.FromField("Description.RebaseMergeAllowed"),
			Description: "If true, rebase-merging is enabled on this repository."},
		{Name: "security_policy_url", Type: proto.ColumnType_STRING,
			Transform:   transform.FromField("Description.SecurityPolicyURL"),
			Description: "The security policy URL."},
		{Name: "squash_merge_allowed", Type: proto.ColumnType_BOOL,
			Transform:   transform.FromField("Description.SquashMergeAllowed"),
			Description: "If true, squash-merging is enabled on this repository."},
		{Name: "squash_merge_commit_message", Type: proto.ColumnType_STRING,
			Transform:   transform.FromField("Description.SquashMergeCommitMessage"),
			Description: "How the default commit message will be generated when squash merging a pull request."},
		{Name: "squash_merge_commit_title", Type: proto.ColumnType_STRING,
			Transform:   transform.FromField("Description.SquashMergeCommitTitle"),
			Description: "How the default commit title will be generated when squash merging a pull request."},
		{Name: "ssh_url", Type: proto.ColumnType_STRING,
			Transform:   transform.FromField("Description.SSHURL"),
			Description: "The SSH URL to clone this repository."},
		{Name: "stargazer_count", Type: proto.ColumnType_INT,
			Transform:   transform.FromField("Description.StargazerCount"),
			Description: "Returns a count of how many stargazers there are on this repository."},
		{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP,
			Transform:   transform.FromField("Description.UpdatedAt").NullIfZero().Transform(convertTimestamp),
			Description: "Timestamp when repository was last updated."},
		{Name: "url", Type: proto.ColumnType_STRING,
			Transform:   transform.FromField("Description.URL"),
			Description: "The URL of the repository."},
		{Name: "uses_custom_open_graph_image", Type: proto.ColumnType_BOOL,
			Transform:   transform.FromField("Description.UsesCustomOpenGraphImage"),
			Description: "if true, this repository has a custom image to use with Open Graph as opposed to being represented by the owner's avatar."},
		{Name: "can_administer", Type: proto.ColumnType_BOOL,
			Transform:   transform.FromField("Description.CanAdminister"),
			Description: "If true, you can administer this repository."},
		{Name: "can_create_projects", Type: proto.ColumnType_BOOL,
			Transform:   transform.FromField("Description.CanCreateProjects"),
			Description: "If true, you can create projects in this repository."},
		{Name: "can_subscribe", Type: proto.ColumnType_BOOL,
			Transform:   transform.FromField("Description.CanSubscribe"),
			Description: "If true, you can subscribe to this repository."},
		{Name: "can_update_topics", Type: proto.ColumnType_BOOL,
			Transform:   transform.FromField("Description.CanUpdateTopics"),
			Description: "If true, you can update topics on this repository."},
		{Name: "has_starred", Type: proto.ColumnType_BOOL,
			Transform:   transform.FromField("Description.HasStarred"),
			Description: "If true, you have starred this repository."},
		{Name: "possible_commit_emails", Type: proto.ColumnType_JSON,
			Transform:   transform.FromField("Description.PossibleCommitEmails"),
			Description: "A list of emails you can commit to this repository with."},
		{Name: "subscription", Type: proto.ColumnType_STRING,
			Transform:   transform.FromField("Description.Subscription"),
			Description: "Identifies if the current user is watching, not watching, or ignoring the repository."},
		{Name: "visibility", Type: proto.ColumnType_STRING,
			Transform:   transform.FromField("Description.Visibility"),
			Description: "Indicates the repository's visibility level."},
		{Name: "your_permission", Type: proto.ColumnType_STRING,
			Transform:   transform.FromField("Description.YourPermission"),
			Description: "Your permission level on the repository. Will return null if authenticated as an GitHub App."},
		{Name: "web_commit_signoff_required", Type: proto.ColumnType_BOOL,
			Transform:   transform.FromField("Description.WebCommitSignOffRequired"),
			Description: "If true, contributors are required to sign off on web-based commits in this repository."},
		{Name: "repository_topics_total_count", Type: proto.ColumnType_INT,
			Transform:   transform.FromField("Description.RepositoryTopicsTotalCount"),
			Description: "Count of topics associated with the repository."},
		{Name: "open_issues_total_count", Type: proto.ColumnType_INT,
			Transform:   transform.FromField("Description.OpenIssuesTotalCount"),
			Description: "Count of issues open on the repository."},
		{Name: "watchers_total_count", Type: proto.ColumnType_INT,
			Transform:   transform.FromField("Description.WatchersTotalCount"),
			Description: "Count of watchers on the repository."},
		// Columns from v3 api - hydrates
		{Name: "hooks", Type: proto.ColumnType_JSON, Description: "The API Hooks URL.",
			Transform: transform.FromField("Description.Hooks"),
		},
		{Name: "topics", Type: proto.ColumnType_JSON, Description: "The topics (similar to tags or labels) associated with the repository.",
			Transform: transform.FromField("Description.Topics"),
		},
		{Name: "subscribers_count", Type: proto.ColumnType_INT, Description: "The number of users who have subscribed to the repository.",
			Transform: transform.FromField("Description.SubscribersCount"),
		},
		{Name: "has_downloads", Type: proto.ColumnType_BOOL, Description: "If true, the GitHub Downloads feature is enabled on the repository.",
			Transform: transform.FromField("Description.HasDownloads"),
		},
		{Name: "has_pages", Type: proto.ColumnType_BOOL, Description: "If true, the GitHub Pages feature is enabled on the repository.",
			Transform: transform.FromField("Description.HasPages"),
		},
		{Name: "network_count", Type: proto.ColumnType_INT, Description: "The number of member repositories in the network.",
			Transform: transform.FromField("Description.NetworkCount"),
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
		Columns: commonColumns(gitHubRepositoryColumns()),
	}
}