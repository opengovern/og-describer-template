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

func appendBranchColumnIncludes(m *map[string]interface{}, cols []string) {
	protectionIncluded := githubv4.Boolean(slices.Contains(cols, "protected") || slices.Contains(cols, "branch_protection_rule"))
	(*m)["includeBranchProtectionRule"] = protectionIncluded
	(*m)["includeAllowsDeletions"] = protectionIncluded
	(*m)["includeAllowsForcePushes"] = protectionIncluded
	(*m)["includeBlocksCreations"] = protectionIncluded
	(*m)["includeCreator"] = protectionIncluded
	(*m)["includeBranchProtectionRuleId"] = protectionIncluded
	(*m)["includeDismissesStaleReviews"] = protectionIncluded
	(*m)["includeIsAdminEnforced"] = protectionIncluded
	(*m)["includeLockAllowsFetchAndMerge"] = protectionIncluded
	(*m)["includeLockBranch"] = protectionIncluded
	(*m)["includePattern"] = protectionIncluded
	(*m)["includeRequireLastPushApproval"] = protectionIncluded
	(*m)["includeRequiredApprovingReviewCount"] = protectionIncluded
	(*m)["includeRequiredDeploymentEnvironments"] = protectionIncluded
	(*m)["includeRequiredStatusChecks"] = protectionIncluded
	(*m)["includeRequiresApprovingReviews"] = protectionIncluded
	(*m)["includeRequiresConversationResolution"] = protectionIncluded
	(*m)["includeRequiresCodeOwnerReviews"] = protectionIncluded
	(*m)["includeRequiresCommitSignatures"] = protectionIncluded
	(*m)["includeRequiresDeployments"] = protectionIncluded
	(*m)["includeRequiresLinearHistory"] = protectionIncluded
	(*m)["includeRequiresStatusChecks"] = protectionIncluded
	(*m)["includeRequiresStrictStatusChecks"] = protectionIncluded
	(*m)["includeRestrictsPushes"] = protectionIncluded
	(*m)["includeRestrictsReviewDismissals"] = protectionIncluded
	(*m)["includeMatchingBranches"] = protectionIncluded
}

func appendBranchProtectionRuleColumnIncludes(m *map[string]interface{}, cols []string) {
	(*m)["includeAllowsDeletions"] = githubv4.Boolean(slices.Contains(cols, "allows_deletions"))
	(*m)["includeAllowsForcePushes"] = githubv4.Boolean(slices.Contains(cols, "allows_force_pushes"))
	(*m)["includeBlocksCreations"] = githubv4.Boolean(slices.Contains(cols, "blocks_creations"))
	(*m)["includeCreator"] = githubv4.Boolean(slices.Contains(cols, "creator") || slices.Contains(cols, "creator_login"))
	(*m)["includeBranchProtectionRuleId"] = githubv4.Boolean(slices.Contains(cols, "id"))
	(*m)["includeDismissesStaleReviews"] = githubv4.Boolean(slices.Contains(cols, "dismisses_stale_reviews"))
	(*m)["includeIsAdminEnforced"] = githubv4.Boolean(slices.Contains(cols, "is_admin_enforced"))
	(*m)["includeLockAllowsFetchAndMerge"] = githubv4.Boolean(slices.Contains(cols, "lock_allows_fetch_and_merge"))
	(*m)["includeLockBranch"] = githubv4.Boolean(slices.Contains(cols, "lock_branch"))
	(*m)["includePattern"] = githubv4.Boolean(slices.Contains(cols, "pattern"))
	(*m)["includeRequireLastPushApproval"] = githubv4.Boolean(slices.Contains(cols, "require_last_push_approval"))
	(*m)["includeRequiredApprovingReviewCount"] = githubv4.Boolean(slices.Contains(cols, "required_approving_review_count"))
	(*m)["includeRequiredDeploymentEnvironments"] = githubv4.Boolean(slices.Contains(cols, "required_deployment_environments"))
	(*m)["includeRequiredStatusChecks"] = githubv4.Boolean(slices.Contains(cols, "required_status_checks"))
	(*m)["includeRequiresApprovingReviews"] = githubv4.Boolean(slices.Contains(cols, "requires_approving_reviews"))
	(*m)["includeRequiresConversationResolution"] = githubv4.Boolean(slices.Contains(cols, "requires_conversation_resolution"))
	(*m)["includeRequiresCodeOwnerReviews"] = githubv4.Boolean(slices.Contains(cols, "requires_code_owner_reviews"))
	(*m)["includeRequiresCommitSignatures"] = githubv4.Boolean(slices.Contains(cols, "requires_commit_signatures"))
	(*m)["includeRequiresDeployments"] = githubv4.Boolean(slices.Contains(cols, "requires_deployments"))
	(*m)["includeRequiresLinearHistory"] = githubv4.Boolean(slices.Contains(cols, "requires_linear_history"))
	(*m)["includeRequiresStatusChecks"] = githubv4.Boolean(slices.Contains(cols, "requires_status_checks"))
	(*m)["includeRequiresStrictStatusChecks"] = githubv4.Boolean(slices.Contains(cols, "requires_strict_status_checks"))
	(*m)["includeRestrictsPushes"] = githubv4.Boolean(slices.Contains(cols, "restricts_pushes"))
	(*m)["includeRestrictsReviewDismissals"] = githubv4.Boolean(slices.Contains(cols, "restricts_review_dismissals"))
	(*m)["includeMatchingBranches"] = githubv4.Boolean(slices.Contains(cols, "matching_branches"))
}

func appendCommitColumnIncludes(m *map[string]interface{}, cols []string) {
	// For BasicCommit struct
	(*m)["includeCommitShortSha"] = githubv4.Boolean(slices.Contains(cols, "short_sha"))
	(*m)["includeCommitAuthoredDate"] = githubv4.Boolean(slices.Contains(cols, "authored_date"))
	(*m)["includeCommitAuthor"] = githubv4.Boolean(slices.Contains(cols, "author") || slices.Contains(cols, "author_login"))
	(*m)["includeCommitCommittedDate"] = githubv4.Boolean(slices.Contains(cols, "committed_date"))
	(*m)["includeCommitCommitter"] = githubv4.Boolean(slices.Contains(cols, "committer") || slices.Contains(cols, "committer_login"))
	(*m)["includeCommitMessage"] = githubv4.Boolean(slices.Contains(cols, "message"))
	(*m)["includeCommitUrl"] = githubv4.Boolean(slices.Contains(cols, "url"))
	// For Commit struct
	(*m)["includeCommitAdditions"] = githubv4.Boolean(slices.Contains(cols, "additions"))
	(*m)["includeCommitAuthoredByCommitter"] = githubv4.Boolean(slices.Contains(cols, "authored_by_committer"))
	(*m)["includeCommitChangedFiles"] = githubv4.Boolean(slices.Contains(cols, "changed_files"))
	(*m)["includeCommitCommittedViaWeb"] = githubv4.Boolean(slices.Contains(cols, "committed_via_web"))
	(*m)["includeCommitCommitUrl"] = githubv4.Boolean(slices.Contains(cols, "commit_url"))
	(*m)["includeCommitDeletions"] = githubv4.Boolean(slices.Contains(cols, "deletions"))
	(*m)["includeCommitSignature"] = githubv4.Boolean(slices.Contains(cols, "signature"))
	(*m)["includeCommitTarballUrl"] = githubv4.Boolean(slices.Contains(cols, "tarball_url"))
	(*m)["includeCommitTreeUrl"] = githubv4.Boolean(slices.Contains(cols, "tree_url"))
	(*m)["includeCommitCanSubscribe"] = githubv4.Boolean(slices.Contains(cols, "can_subscribe"))
	(*m)["includeCommitSubscription"] = githubv4.Boolean(slices.Contains(cols, "subscription"))
	(*m)["includeCommitZipballUrl"] = githubv4.Boolean(slices.Contains(cols, "zipball_url"))
	(*m)["includeCommitMessageHeadline"] = githubv4.Boolean(slices.Contains(cols, "message_headline"))
	(*m)["includeCommitStatus"] = githubv4.Boolean(slices.Contains(cols, "status"))
	(*m)["includeCommitNodeId"] = githubv4.Boolean(slices.Contains(cols, "node_id"))
}

func repositoryCols() []string {
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

func branchCols() []string {
	return []string{
		"repository_full_name",
		"name",
		"commit",
		"protected",
		"branch_protection_rule",
	}
}

func branchProtectionCols() []string {
	return []string{
		"repository_full_name",
		"id",
		"node_id",
		"matching_branches",
		"is_admin_enforced",
		"allows_deletions",
		"allows_force_pushes",
		"blocks_creations",
		"creator_login",
		"dismisses_stale_reviews",
		"lock_allows_fetch_and_merge",
		"lock_branch",
		"pattern",
		"require_last_push_approval",
		"requires_approving_reviews",
		"required_approving_review_count",
		"requires_conversation_resolution",
		"requires_code_owner_reviews",
		"requires_commit_signatures",
		"requires_deployments",
		"required_deployment_environments",
		"requires_linear_history",
		"requires_status_checks",
		"required_status_checks",
		"requires_strict_status_checks",
		"restricts_review_dismissals",
		"restricts_pushes",
		"push_allowance_apps",
		"push_allowance_teams",
		"push_allowance_users",
		"bypass_force_push_allowance_apps",
		"bypass_force_push_allowance_teams",
		"bypass_force_push_allowance_users",
		"bypass_pull_request_allowance_apps",
		"bypass_pull_request_allowance_teams",
		"bypass_pull_request_allowance_users",
		"repository_full_name",
		"name",
		"commit",
		"protected",
		"branch_protection_rule",
	}
}

func commitCols() []string {
	return []string{
		"repository_full_name",
		"sha",
		"short_sha",
		"message",
		"author_login",
		"authored_date",
		"author",
		"committer_login",
		"committed_date",
		"committer",
		"additions",
		"authored_by_committer",
		"deletions",
		"changed_files",
		"committed_via_web",
		"commit_url",
		"signature",
		"status",
		"tarball_url",
		"zipball_url",
		"tree_url",
		"can_subscribe",
		"subscription",
		"url",
		"node_id",
		"message_headline",
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
