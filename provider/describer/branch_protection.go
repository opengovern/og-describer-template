package describer

import (
	"context"
	"github.com/opengovern/og-describer-github/pkg/sdk/models"
	"github.com/opengovern/og-describer-github/provider/model"
	"github.com/shurcooL/githubv4"
	steampipemodels "github.com/turbot/steampipe-plugin-github/github/models"
	"strconv"
)

func GetAllBranchProtections(ctx context.Context, githubClient GitHubClient, stream *models.StreamSender) ([]models.Resource, error) {
	client := githubClient.RestClient
	owner, err := getOwnerName(ctx, client)
	if err != nil {
		return nil, nil
	}
	repositories, err := getRepositories(ctx, client, owner)
	if err != nil {
		return nil, nil
	}
	var values []models.Resource
	for _, repo := range repositories {
		repoValues, err := GetRepositoryBranchProtections(ctx, githubClient, stream, owner, repo.GetName())
		if err != nil {
			return nil, err
		}
		values = append(values, repoValues...)
	}
	return values, nil
}

func GetRepositoryBranchProtections(ctx context.Context, githubClient GitHubClient, stream *models.StreamSender, owner, repo string) ([]models.Resource, error) {
	client := githubClient.GraphQLClient
	var query struct {
		RateLimit  steampipemodels.RateLimit
		Repository struct {
			BranchProtectionRules struct {
				TotalCount int
				PageInfo   steampipemodels.PageInfo
				Nodes      []steampipemodels.BranchProtectionRuleWithFirstPageEmbeddedItems
			} `graphql:"branchProtectionRules(first: $pageSize, after: $cursor)"`
		} `graphql:"repository(owner: $owner, name: $repo)"`
	}
	variables := map[string]interface{}{
		"owner":    githubv4.String(owner),
		"repo":     githubv4.String(repo),
		"pageSize": githubv4.Int(pageSize),
		"cursor":   (*githubv4.String)(nil),
	}
	appendBranchProtectionRuleColumnIncludes(&variables, branchProtectionCols())
	repoFullName := formRepositoryFullName(owner, repo)
	var values []models.Resource
	var pushAllowanceApps []model.BranchApp
	var pushAllowanceTeams []model.BranchTeam
	var pushAllowanceUsers []model.BranchUser
	var bypassForcePushAllowanceApps []model.BranchApp
	var bypassForcePushAllowanceTeams []model.BranchTeam
	var bypassForcePushAllowanceUsers []model.BranchUser
	var bypassPullRequestAllowanceApps []model.BranchApp
	var bypassPullRequestAllowanceTeams []model.BranchTeam
	var bypassPullRequestAllowanceUsers []model.BranchUser
	for {
		err := client.Query(ctx, &query, variables)
		if err != nil {
			return nil, err
		}
		for _, rule := range query.Repository.BranchProtectionRules.Nodes {
			row := mapBranchProtectionRule(&rule)
			if rule.PushAllowances.PageInfo.HasNextPage {
				err := branchProtectionGetPushAllowances(ctx, client, &row, rule.PushAllowances.PageInfo.EndCursor)
				if err != nil {
					return nil, err
				}
			}
			if rule.BypassForcePushAllowances.PageInfo.HasNextPage {
				err := branchProtectionGetBypassForcePushAllowances(ctx, client, &row, rule.BypassForcePushAllowances.PageInfo.EndCursor)
				if err != nil {
					return nil, err
				}
			}
			if rule.BypassPullRequestAllowances.PageInfo.HasNextPage {
				err := branchProtectionGetBypassPullRequestAllowances(ctx, client, &row, rule.BypassPullRequestAllowances.PageInfo.EndCursor)
				if err != nil {
					return nil, err
				}
			}
			for _, node := range rule.PushAllowances.Nodes {
				pushAllowanceApps = append(pushAllowanceApps, node.Actor.App)
				pushAllowanceTeams = append(pushAllowanceTeams, node.Actor.Team)
				pushAllowanceUsers = append(pushAllowanceUsers, node.Actor.User)
			}
			for _, node := range rule.BypassForcePushAllowances.Nodes {
				bypassForcePushAllowanceApps = append(bypassForcePushAllowanceApps, node.Actor.App)
				bypassForcePushAllowanceTeams = append(bypassForcePushAllowanceTeams, node.Actor.Team)
				bypassForcePushAllowanceUsers = append(bypassForcePushAllowanceUsers, node.Actor.User)
			}
			for _, node := range rule.BypassForcePushAllowances.Nodes {
				bypassPullRequestAllowanceApps = append(bypassPullRequestAllowanceApps, node.Actor.App)
				bypassPullRequestAllowanceTeams = append(bypassPullRequestAllowanceTeams, node.Actor.Team)
				bypassPullRequestAllowanceUsers = append(bypassPullRequestAllowanceUsers, node.Actor.User)
			}
			value := models.Resource{
				ID:   strconv.Itoa(rule.Id),
				Name: strconv.Itoa(rule.Id),
				Description: JSONAllFieldsMarshaller{
					Value: model.BranchProtectionDescription{
						BranchProtectionRule:            rule.BranchProtectionRule,
						RepoFullName:                    repoFullName,
						CreatorLogin:                    rule.Creator.Login,
						MatchingBranches:                rule.MatchingBranches.TotalCount,
						PushAllowanceApps:               pushAllowanceApps,
						PushAllowanceTeams:              pushAllowanceTeams,
						PushAllowanceUsers:              pushAllowanceUsers,
						BypassForcePushAllowanceApps:    bypassForcePushAllowanceApps,
						BypassForcePushAllowanceTeams:   bypassForcePushAllowanceTeams,
						BypassForcePushAllowanceUsers:   bypassForcePushAllowanceUsers,
						BypassPullRequestAllowanceApps:  bypassPullRequestAllowanceApps,
						BypassPullRequestAllowanceTeams: bypassPullRequestAllowanceTeams,
						BypassPullRequestAllowanceUsers: bypassPullRequestAllowanceUsers,
					},
				},
			}
			if stream != nil {
				if err := (*stream)(value); err != nil {
					return nil, err
				}
			} else {
				values = append(values, value)
			}
		}
		if !query.Repository.BranchProtectionRules.PageInfo.HasNextPage {
			break
		}
		variables["cursor"] = githubv4.NewString(query.Repository.BranchProtectionRules.PageInfo.EndCursor)
	}
	return values, nil
}

func branchProtectionGetPushAllowances(ctx context.Context, client *githubv4.Client, row *branchProtectionRow, initialCursor githubv4.String) error {
	var query struct {
		RateLimit steampipemodels.RateLimit
		Node      struct {
			BranchProtectionRule steampipemodels.BranchProtectionRuleWithPushAllowances `graphql:"... on BranchProtectionRule"`
		} `graphql:"node(id: $nodeId)"`
	}
	vars := map[string]interface{}{
		"nodeId":   githubv4.ID(row.NodeID),
		"pageSize": githubv4.Int(100),
		"cursor":   githubv4.NewString(initialCursor),
	}
	for {
		err := client.Query(ctx, &query, vars)
		if err != nil {
			return err
		}
		a, t, u := query.Node.BranchProtectionRule.PushAllowances.Explode()
		row.PushAllowanceApps = append(row.PushAllowanceApps, a...)
		row.PushAllowanceTeams = append(row.PushAllowanceTeams, t...)
		row.PushAllowanceUsers = append(row.PushAllowanceUsers, u...)
		if !query.Node.BranchProtectionRule.PushAllowances.PageInfo.HasNextPage {
			break
		}
		vars["cursor"] = githubv4.NewString(query.Node.BranchProtectionRule.PushAllowances.PageInfo.EndCursor)
	}
	return nil
}

func branchProtectionGetBypassForcePushAllowances(ctx context.Context, client *githubv4.Client, row *branchProtectionRow, initialCursor githubv4.String) error {
	var query struct {
		RateLimit steampipemodels.RateLimit
		Node      struct {
			BranchProtectionRule steampipemodels.BranchProtectionRuleWithBypassForcePushAllowances `graphql:"... on BranchProtectionRule"`
		} `graphql:"node(id: $nodeId)"`
	}
	vars := map[string]interface{}{
		"nodeId":   githubv4.ID(row.NodeID),
		"pageSize": githubv4.Int(100),
		"cursor":   githubv4.NewString(initialCursor),
	}
	for {
		err := client.Query(ctx, &query, vars)
		if err != nil {
			return err
		}
		a, t, u := query.Node.BranchProtectionRule.BypassForcePushAllowances.Explode()
		row.BypassForcePushAllowanceApps = append(row.BypassForcePushAllowanceApps, a...)
		row.BypassForcePushAllowanceTeams = append(row.BypassForcePushAllowanceTeams, t...)
		row.BypassForcePushAllowanceUsers = append(row.BypassForcePushAllowanceUsers, u...)
		if !query.Node.BranchProtectionRule.BypassForcePushAllowances.PageInfo.HasNextPage {
			break
		}
		vars["cursor"] = githubv4.NewString(query.Node.BranchProtectionRule.BypassForcePushAllowances.PageInfo.EndCursor)
	}
	return nil
}

func branchProtectionGetBypassPullRequestAllowances(ctx context.Context, client *githubv4.Client, row *branchProtectionRow, initialCursor githubv4.String) error {
	var query struct {
		RateLimit steampipemodels.RateLimit
		Node      struct {
			BranchProtectionRule steampipemodels.BranchProtectionRuleWithBypassPullRequestAllowances `graphql:"... on BranchProtectionRule"`
		} `graphql:"node(id: $nodeId)"`
	}
	vars := map[string]interface{}{
		"nodeId":   githubv4.ID(row.NodeID),
		"pageSize": githubv4.Int(100),
		"cursor":   githubv4.NewString(initialCursor),
	}
	for {
		err := client.Query(ctx, &query, vars)
		if err != nil {
			return err
		}
		a, t, u := query.Node.BranchProtectionRule.BypassPullRequestAllowances.Explode()
		row.BypassPullRequestAllowanceApps = append(row.BypassPullRequestAllowanceApps, a...)
		row.BypassPullRequestAllowanceTeams = append(row.BypassPullRequestAllowanceTeams, t...)
		row.BypassPullRequestAllowanceUsers = append(row.BypassPullRequestAllowanceUsers, u...)
		if !query.Node.BranchProtectionRule.BypassPullRequestAllowances.PageInfo.HasNextPage {
			break
		}
		vars["cursor"] = githubv4.NewString(query.Node.BranchProtectionRule.BypassPullRequestAllowances.PageInfo.EndCursor)
	}
	return nil
}

func mapBranchProtectionRule(rule *steampipemodels.BranchProtectionRuleWithFirstPageEmbeddedItems) branchProtectionRow {
	row := branchProtectionRow{
		ID:                             rule.Id,
		NodeID:                         rule.NodeId,
		MatchingBranches:               rule.MatchingBranches.TotalCount,
		IsAdminEnforced:                rule.IsAdminEnforced,
		AllowsDeletions:                rule.AllowsDeletions,
		AllowsForcePushes:              rule.AllowsForcePushes,
		BlocksCreations:                rule.BlocksCreations,
		CreatorLogin:                   rule.Creator.Login,
		DismissesStaleReviews:          rule.DismissesStaleReviews,
		LockAllowsFetchAndMerge:        rule.LockAllowsFetchAndMerge,
		LockBranch:                     rule.LockBranch,
		Pattern:                        rule.Pattern,
		RequireLastPushApproval:        rule.RequireLastPushApproval,
		RequiredApprovingReviewCount:   rule.RequiredApprovingReviewCount,
		RequiredDeploymentEnvironments: rule.RequiredDeploymentEnvironments,
		RequiredStatusChecks:           rule.RequiredStatusChecks,
		RequiresApprovingReviews:       rule.RequiresApprovingReviews,
		RequiresConversationResolution: rule.RequiresConversationResolution,
		RequiresCodeOwnerReviews:       rule.RequiresCodeOwnerReviews,
		RequiresCommitSignatures:       rule.RequiresCommitSignatures,
		RequiresDeployments:            rule.RequiresDeployments,
		RequiresLinearHistory:          rule.RequiresLinearHistory,
		RequiresStatusChecks:           rule.RequiresStatusChecks,
		RequiresStrictStatusChecks:     rule.RequiresStrictStatusChecks,
		RestrictsPushes:                rule.RestrictsPushes,
		RestrictsReviewDismissals:      rule.RestrictsReviewDismissals,
	}
	row.PushAllowanceApps, row.PushAllowanceTeams, row.PushAllowanceUsers = rule.PushAllowances.Explode()
	row.BypassForcePushAllowanceApps, row.BypassForcePushAllowanceTeams, row.BypassForcePushAllowanceUsers = rule.BypassForcePushAllowances.Explode()
	row.BypassPullRequestAllowanceApps, row.BypassPullRequestAllowanceTeams, row.BypassPullRequestAllowanceUsers = rule.BypassPullRequestAllowances.Explode()
	return row
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
	PushAllowanceApps               []steampipemodels.NameSlug
	PushAllowanceTeams              []steampipemodels.NameSlug
	PushAllowanceUsers              []steampipemodels.NameLogin
	BypassForcePushAllowanceApps    []steampipemodels.NameSlug
	BypassForcePushAllowanceTeams   []steampipemodels.NameSlug
	BypassForcePushAllowanceUsers   []steampipemodels.NameLogin
	BypassPullRequestAllowanceApps  []steampipemodels.NameSlug
	BypassPullRequestAllowanceTeams []steampipemodels.NameSlug
	BypassPullRequestAllowanceUsers []steampipemodels.NameLogin
}
