package describer

import (
	"context"
	"github.com/google/go-github/v55/github"
	"github.com/shurcooL/githubv4"
	"slices"
)

const (
	maxPagesCount = 100
	pageSize      = 50
)

func appendRepoColumnIncludes(m *map[string]interface{}, cols []string) {
	optionals := map[string]string{
		"allow_update_branch":              "includeAllowUpdateBranch",
		"archived_at":                      "includeArchivedAt",
		"auto_merge_allowed":               "includeAutoMergeAllowed",
		"can_administer":                   "includeCanAdminister",
		"can_create_projects":              "includeCanCreateProjects",
		"can_subscribe":                    "includeCanSubscribe",
		"can_update_topics":                "includeCanUpdateTopics",
		"code_of_conduct":                  "includeCodeOfConduct",
		"contact_links":                    "includeContactLinks",
		"created_at":                       "includeCreatedAt",
		"default_branch_ref":               "includeDefaultBranchRef",
		"delete_branch_on_merge":           "includeDeleteBranchOnMerge",
		"description":                      "includeDescription",
		"disk_usage":                       "includeDiskUsage",
		"fork_count":                       "includeForkCount",
		"forking_allowed":                  "includeForkingAllowed",
		"funding_links":                    "includeFundingLinks",
		"has_discussions_enabled":          "includeHasDiscussionsEnabled",
		"has_issues_enabled":               "includeHasIssuesEnabled",
		"has_projects_enabled":             "includeHasProjectsEnabled",
		"has_starred":                      "includeHasStarred",
		"has_vulnerability_alerts_enabled": "includeHasVulnerabilityAlertsEnabled",
		"has_wiki_enabled":                 "includeHasWikiEnabled",
		"homepage_url":                     "includeHomepageUrl",
		"interaction_ability":              "includeInteractionAbility",
		"is_archived":                      "includeIsArchived",
		"is_blank_issues_enabled":          "includeIsBlankIssuesEnabled",
		"is_disabled":                      "includeIsDisabled",
		"is_empty":                         "includeIsEmpty",
		"is_fork":                          "includeIsFork",
		"is_in_organization":               "includeIsInOrganization",
		"is_locked":                        "includeIsLocked",
		"is_mirror":                        "includeIsMirror",
		"is_private":                       "includeIsPrivate",
		"is_security_policy_enabled":       "includeIsSecurityPolicyEnabled",
		"is_template":                      "includeIsTemplate",
		"is_user_configuration_repository": "includeIsUserConfigurationRepository",
		"issue_templates":                  "includeIssueTemplates",
		"license_info":                     "includeLicenseInfo",
		"lock_reason":                      "includeLockReason",
		"merge_commit_allowed":             "includeMergeCommitAllowed",
		"merge_commit_message":             "includeMergeCommitMessage",
		"merge_commit_title":               "includeMergeCommitTitle",
		"mirror_url":                       "includeMirrorUrl",
		"open_graph_image_url":             "includeOpenGraphImageUrl",
		"open_issues_total_count":          "includeOpenIssues",
		"possible_commit_emails":           "includePossibleCommitEmails",
		"primary_language":                 "includePrimaryLanguage",
		"projects_url":                     "includeProjectsUrl",
		"pull_request_templates":           "includePullRequestTemplates",
		"pushed_at":                        "includePushedAt",
		"rebase_merge_allowed":             "includeRebaseMergeAllowed",
		"repository_topics_total_count":    "includeRepositoryTopics",
		"security_policy_url":              "includeSecurityPolicyUrl",
		"squash_merge_allowed":             "includeSquashMergeAllowed",
		"squash_merge_commit_message":      "includeSquashMergeCommitMessage",
		"squash_merge_commit_title":        "includeSquashMergeCommitTitle",
		"ssh_url":                          "includeSshUrl",
		"stargazer_count":                  "includeStargazerCount",
		"subscription":                     "includeSubscription",
		"updated_at":                       "includeUpdatedAt",
		"url":                              "includeUrl",
		"uses_custom_open_graph_image":     "includeUsesCustomOpenGraphImage",
		"visibility":                       "includeVisibility",
		"watchers_total_count":             "includeWatchers",
		"web_commit_signoff_required":      "includeWebCommitSignoffRequired",
		"your_permission":                  "includeYourPermission",
	}
	for key, value := range optionals {
		(*m)[value] = githubv4.Boolean(slices.Contains(cols, key))
	}
}

func tableCols() []string {
	return []string{
		"id",
		"node_id",
		"name",
		"allow_update_branch",
		"archived_at",
		"auto_merge_allowed",
		"code_of_conduct",
		"contact_links",
		"created_at",
		"default_branch_ref",
		"delete_branch_on_merge",
		"description",
		"disk_usage",
		"fork_count",
		"forking_allowed",
		"funding_links",
		"has_discussions_enabled",
		"has_issues_enabled",
		"has_projects_enabled",
		"has_vulnerability_alerts_enabled",
		"has_wiki_enabled",
		"homepage_url",
		"interaction_ability",
		"is_archived",
		"is_blank_issues_enabled",
		"is_disabled",
		"is_empty",
		"is_fork",
		"is_in_organization",
		"is_locked",
		"is_mirror",
		"is_private",
		"is_security_policy_enabled",
		"is_template",
		"is_user_configuration_repository",
		"issue_templates",
		"license_info",
		"lock_reason",
		"merge_commit_allowed",
		"merge_commit_message",
		"merge_commit_title",
		"mirror_url",
		"name_with_owner",
		"open_graph_image_url",
		"owner_login",
		"primary_language",
		"projects_url",
		"pull_request_templates",
		"pushed_at",
		"rebase_merge_allowed",
		"security_policy_url",
		"squash_merge_allowed",
		"squash_merge_commit_message",
		"squash_merge_commit_title",
		"ssh_url",
		"stargazer_count",
		"updated_at",
		"url",
		"uses_custom_open_graph_image",
		"can_administer",
		"can_create_projects",
		"can_subscribe",
		"can_update_topics",
		"has_starred",
		"possible_commit_emails",
		"subscription",
		"visibility",
		"your_permission",
		"web_commit_signoff_required",
		"repository_topics_total_count",
		"open_issues_total_count",
		"watchers_total_count",
		"hooks",
		"topics",
		"subscribers_count",
		"has_downloads",
		"has_pages",
		"network_count",
	}
}

func getOwnerName(ctx context.Context, client *github.Client) (string, error) {
	owner, _, err := client.Users.Get(ctx, "")
	if err != nil {
		return "", err
	}
	ownerName := *owner.Name
	return ownerName, err
}

func getRepositoriesName(ctx context.Context, client *github.Client, owner string) ([]string, error) {
	opt := &github.RepositoryListOptions{
		ListOptions: github.ListOptions{PerPage: maxPagesCount},
	}
	var repositories []string
	for {
		repos, resp, err := client.Repositories.List(ctx, owner, opt)
		if err != nil {
			return nil, err
		}
		for _, repo := range repos {
			repositories = append(repositories, repo.GetName())
		}
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}
	return repositories, nil
}
