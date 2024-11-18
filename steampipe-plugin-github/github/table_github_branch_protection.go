package github

import (
	opengovernance "github.com/opengovern/og-describer-github/pkg/sdk/es"
	"github.com/opengovern/og-describer-github/steampipe-plugin-github/github/models"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableGitHubBranchProtection() *plugin.Table {
	return &plugin.Table{
		Name:        "github_branch_protection",
		Description: "Branch protection defines rules for pushing to and managing a branch.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListBranchProtection,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("node_id"),
			Hydrate:    opengovernance.GetBranchProtection,
		},
		Columns: commonColumns([]*plugin.Column{
			{
				Name:        "repository_full_name",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RepoFullName"),
				Description: "The full name of the repository (login/repo-name).",
			},
			{
				Name:        "id",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Id"),
				Description: "The ID of the branch protection rule.",
			},
			{
				Name:        "node_id",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.NodeId"),
				Description: "The Node ID of the branch protection rule."},
			{
				Name:        "matching_branches",
				Type:        proto.ColumnType_INT,
				Hydrate:     branchProtectionRuleHydrateMatchingBranchesTotalCount,
				Transform:   transform.FromValue(),
				Description: "Count of branches which match this rule."},
			{
				Name:        "is_admin_enforced",
				Type:        proto.ColumnType_BOOL,
				Hydrate:     branchProtectionRuleHydrateIsAdminEnforced,
				Transform:   transform.FromValue(),
				Description: "If true, enforce all configured restrictions for administrators.",
			},
			{
				Name:        "allows_deletions",
				Type:        proto.ColumnType_BOOL,
				Hydrate:     branchProtectionRuleHydrateAllowsDeletions,
				Transform:   transform.FromValue(),
				Description: "If true, allow users with push access to delete matching branches."},
			{Name: "allows_force_pushes", Type: proto.ColumnType_BOOL, Hydrate: branchProtectionRuleHydrateAllowsForcePushes, Transform: transform.FromValue(), Description: "If true, permit force pushes for all users with push access."},
			{Name: "blocks_creations", Type: proto.ColumnType_BOOL, Hydrate: branchProtectionRuleHydrateBlocksCreations, Transform: transform.FromValue(), Description: "If true, indicates that branch creation is a protected operation."},
			{Name: "creator_login", Type: proto.ColumnType_STRING, Hydrate: branchProtectionRuleHydrateCreatorLogin, Transform: transform.FromValue(), Description: "The login of the user whom created the branch protection rule."},
			{Name: "dismisses_stale_reviews", Type: proto.ColumnType_BOOL, Hydrate: branchProtectionRuleHydrateDismissesStaleReviews, Transform: transform.FromValue(), Description: "If true, new commits pushed to matching branches dismiss pull request review approvals."},
			{Name: "lock_allows_fetch_and_merge", Type: proto.ColumnType_BOOL, Hydrate: branchProtectionRuleHydrateLockAllowsFetchAndMerge, Transform: transform.FromValue(), Description: "If true, users can pull changes from upstream when the branch is locked."},
			{Name: "lock_branch", Type: proto.ColumnType_BOOL, Hydrate: branchProtectionRuleHydrateLockBranch, Transform: transform.FromValue(), Description: "If true, matching branches are read-only and cannot be pushed to."},
			{Name: "pattern", Type: proto.ColumnType_STRING, Hydrate: branchProtectionRuleHydratePattern, Transform: transform.FromValue(), Description: "The protection rule pattern."},
			{Name: "require_last_push_approval", Type: proto.ColumnType_BOOL, Hydrate: branchProtectionRuleHydrateRequireLastPushApproval, Transform: transform.FromValue(), Description: "If true, the most recent push must be approved by someone other than the person who pushed it."},
			{Name: "requires_approving_reviews", Type: proto.ColumnType_BOOL, Hydrate: branchProtectionRuleHydrateRequiresApprovingReviews, Transform: transform.FromValue(), Description: "If true, approving reviews required to update matching branches."},
			{Name: "required_approving_review_count", Type: proto.ColumnType_INT, Hydrate: branchProtectionRuleHydrateRequiredApprovingReviewCount, Transform: transform.FromValue(), Description: "Number of approving reviews required to update matching branches."},
			{Name: "requires_conversation_resolution", Type: proto.ColumnType_BOOL, Hydrate: branchProtectionRuleHydrateRequiresConversationResolution, Transform: transform.FromValue(), Description: "If true, requires all comments on the pull request to be resolved before it can be merged to a protected branch."},
			{Name: "requires_code_owner_reviews", Type: proto.ColumnType_BOOL, Hydrate: branchProtectionRuleHydrateRequiresCodeOwnerReviews, Transform: transform.FromValue(), Description: "If true, reviews from code owners are required to update matching branches."},
			{Name: "requires_commit_signatures", Type: proto.ColumnType_BOOL, Hydrate: branchProtectionRuleHydrateRequiresCommitSignatures, Transform: transform.FromValue(), Description: "If true, commits are required to be signed by verified signatures."},
			{Name: "requires_deployments", Type: proto.ColumnType_BOOL, Hydrate: branchProtectionRuleHydrateRequiresDeployments, Transform: transform.FromValue(), Description: "If true, matching branches require deployment to specific environments before merging."},
			{Name: "required_deployment_environments", Type: proto.ColumnType_JSON, Hydrate: branchProtectionRuleHydrateRequiredDeploymentEnvironments, Transform: transform.FromValue(), Description: "List of required deployment environments that must be deployed successfully to update matching branches."},
			{Name: "requires_linear_history", Type: proto.ColumnType_BOOL, Hydrate: branchProtectionRuleHydrateRequiresLinearHistory, Transform: transform.FromValue(), Description: "If true, prevent merge commits from being pushed to matching branches."},
			{Name: "requires_status_checks", Type: proto.ColumnType_BOOL, Hydrate: branchProtectionRuleHydrateRequiresStatusChecks, Transform: transform.FromValue(), Description: "If true, status checks are required to update matching branches."},
			{Name: "required_status_checks", Type: proto.ColumnType_JSON, Hydrate: branchProtectionRuleHydrateRequiredStatusChecks, Transform: transform.FromValue(), Description: "Status checks that must pass before a branch can be merged into branches matching this rule."},
			{Name: "requires_strict_status_checks", Type: proto.ColumnType_BOOL, Hydrate: branchProtectionRuleHydrateRequiresStrictStatusChecks, Transform: transform.FromValue(), Description: "If true, branches required to be up to date before merging."},
			{Name: "restricts_review_dismissals", Type: proto.ColumnType_BOOL, Hydrate: branchProtectionRuleHydrateRestrictsReviewDismissals, Transform: transform.FromValue(), Description: "If true, review dismissals are restricted."},
			{Name: "restricts_pushes", Type: proto.ColumnType_BOOL, Hydrate: branchProtectionRuleHydrateRestrictsPushes, Transform: transform.FromValue(), Description: "If true, pushing to matching branches is restricted."},
			{Name: "push_allowance_apps", Type: proto.ColumnType_JSON, Description: "Applications can push to the branch only if in this list."},
			{Name: "push_allowance_teams", Type: proto.ColumnType_JSON, Description: "Teams can push to the branch only if in this list."},
			{Name: "push_allowance_users", Type: proto.ColumnType_JSON, Description: "Users can push to the branch only if in this list."},
			{Name: "bypass_force_push_allowance_apps", Type: proto.ColumnType_JSON, Description: "Applications can force push to the branch only if in this list."},
			{Name: "bypass_force_push_allowance_teams", Type: proto.ColumnType_JSON, Description: "Teams can force push to the branch only if in this list."},
			{Name: "bypass_force_push_allowance_users", Type: proto.ColumnType_JSON, Description: "Users can force push to the branch only if in this list."},
			{Name: "bypass_pull_request_allowance_apps", Type: proto.ColumnType_JSON, Description: "Applications can bypass pull requests to the branch only if in this list."},
			{Name: "bypass_pull_request_allowance_teams", Type: proto.ColumnType_JSON, Description: "Teams can bypass pull requests to the branch only if in this list."},
			{Name: "bypass_pull_request_allowance_users", Type: proto.ColumnType_JSON, Description: "Users can bypass pull requests to the branch only if in this list."},
		}),
	}
}

// branchProtectionRow is used to flatten nested pageable items into separate columns by type
type branchProtectionRow struct {
	ID                              int
	NodeID                          string
	MatchingBranches                int
	IsAdminEnforced                 bool
	AllowsDeletions                 bool
	AllowsForcePushes               bool
	BlocksCreations                 bool
	CreatorLogin                    string
	DismissesStaleReviews           bool
	LockAllowsFetchAndMerge         bool
	LockBranch                      bool
	Pattern                         string
	RequireLastPushApproval         bool
	RequiredApprovingReviewCount    int
	RequiredDeploymentEnvironments  []string
	RequiredStatusChecks            []string
	RequiresApprovingReviews        bool
	RequiresConversationResolution  bool
	RequiresCodeOwnerReviews        bool
	RequiresCommitSignatures        bool
	RequiresDeployments             bool
	RequiresLinearHistory           bool
	RequiresStatusChecks            bool
	RequiresStrictStatusChecks      bool
	RestrictsPushes                 bool
	RestrictsReviewDismissals       bool
	PushAllowanceApps               []models.NameSlug
	PushAllowanceTeams              []models.NameSlug
	PushAllowanceUsers              []models.NameLogin
	BypassForcePushAllowanceApps    []models.NameSlug
	BypassForcePushAllowanceTeams   []models.NameSlug
	BypassForcePushAllowanceUsers   []models.NameLogin
	BypassPullRequestAllowanceApps  []models.NameSlug
	BypassPullRequestAllowanceTeams []models.NameSlug
	BypassPullRequestAllowanceUsers []models.NameLogin
}
