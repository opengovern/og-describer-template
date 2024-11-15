package describer

import (
	"context"
	"github.com/opengovern/og-describer-github/pkg/sdk/models"
	"github.com/opengovern/og-describer-github/provider/model"
	"github.com/shurcooL/githubv4"
	steampipemodels "github.com/turbot/steampipe-plugin-github/github/models"
	"strconv"
	"strings"
)

func GetAllTeamsRepositories(ctx context.Context, githubClient GitHubClient, stream *models.StreamSender) ([]models.Resource, error) {
	client := githubClient.RestClient
	teams, err := getTeams(ctx, client)
	if err != nil {
		return nil, nil
	}
	var values []models.Resource
	for _, team := range teams {
		teamValues, err := GetTeamRepositories(ctx, githubClient, stream, *team.Organization.Name, *team.Slug)
		if err != nil {
			return nil, err
		}
		values = append(values, teamValues...)
	}
	return values, nil
}

func GetTeamRepositories(ctx context.Context, githubClient GitHubClient, stream *models.StreamSender, org, slug string) ([]models.Resource, error) {
	client := githubClient.GraphQLClient
	var query struct {
		RateLimit    steampipemodels.RateLimit
		Organization struct {
			Team struct {
				Repositories struct {
					TotalCount int
					PageInfo   steampipemodels.PageInfo
					Edges      []steampipemodels.TeamRepositoryWithPermission
				} `graphql:"repositories(first: $pageSize, after: $cursor)"`
			} `graphql:"team(slug: $slug)"`
		} `graphql:"organization(login: $login)"`
	}
	variables := map[string]interface{}{
		"login":    githubv4.String(org),
		"slug":     githubv4.String(slug),
		"pageSize": githubv4.Int(teamMembersPageSize),
		"cursor":   (*githubv4.String)(nil),
	}
	appendRepoColumnIncludes(&variables, teamRepositoriesCols())
	var values []models.Resource
	for {
		err := client.Query(ctx, &query, variables)
		if err != nil {
			if strings.Contains(err.Error(), "Could not resolve to an Organization with the login of") {
				return nil, nil
			}
			return nil, err
		}
		for _, repo := range query.Organization.Team.Repositories.Edges {
			hooks, err := GetRepositoryHooks(ctx, githubClient.RestClient, repo.Node.Name)
			if err != nil {
				return nil, err
			}
			additionalRepoInfo, err := GetRepositoryAdditionalData(ctx, githubClient.RestClient, repo.Node.Name)
			value := models.Resource{
				ID:   strconv.Itoa(repo.Node.Id),
				Name: repo.Node.Name,
				Description: JSONAllFieldsMarshaller{
					Value: model.TeamRepositoryDescription{
						RepositoryDescription: model.RepositoryDescription{
							ID:                            repo.Node.Id,
							NodeID:                        repo.Node.NodeId,
							Name:                          repo.Node.Name,
							AllowUpdateBranch:             repo.Node.AllowUpdateBranch,
							ArchivedAt:                    repo.Node.ArchivedAt,
							AutoMergeAllowed:              repo.Node.AutoMergeAllowed,
							CodeOfConduct:                 repo.Node.CodeOfConduct,
							ContactLinks:                  repo.Node.ContactLinks,
							CreatedAt:                     repo.Node.CreatedAt,
							DefaultBranchRef:              repo.Node.DefaultBranchRef,
							DeleteBranchOnMerge:           repo.Node.DeleteBranchOnMerge,
							Description:                   repo.Node.Description,
							DiskUsage:                     repo.Node.DiskUsage,
							ForkCount:                     repo.Node.ForkCount,
							ForkingAllowed:                repo.Node.ForkingAllowed,
							FundingLinks:                  repo.Node.FundingLinks,
							HasDiscussionsEnabled:         repo.Node.HasDiscussionsEnabled,
							HasIssuesEnabled:              repo.Node.HasIssuesEnabled,
							HasProjectsEnabled:            repo.Node.HasProjectsEnabled,
							HasVulnerabilityAlertsEnabled: repo.Node.HasVulnerabilityAlertsEnabled,
							HasWikiEnabled:                repo.Node.HasWikiEnabled,
							HomepageURL:                   repo.Node.HomepageUrl,
							InteractionAbility:            repo.Node.InteractionAbility,
							IsArchived:                    repo.Node.IsArchived,
							IsBlankIssuesEnabled:          repo.Node.IsBlankIssuesEnabled,
							IsDisabled:                    repo.Node.IsDisabled,
							IsEmpty:                       repo.Node.IsEmpty,
							IsFork:                        repo.Node.IsFork,
							IsInOrganization:              repo.Node.IsInOrganization,
							IsLocked:                      repo.Node.IsLocked,
							IsMirror:                      repo.Node.IsMirror,
							IsPrivate:                     repo.Node.IsPrivate,
							IsSecurityPolicyEnabled:       repo.Node.IsSecurityPolicyEnabled,
							IsTemplate:                    repo.Node.IsTemplate,
							IsUserConfigurationRepository: repo.Node.IsUserConfigurationRepository,
							IssueTemplates:                repo.Node.IssueTemplates,
							LicenseInfo:                   repo.Node.LicenseInfo,
							LockReason:                    repo.Node.LockReason,
							MergeCommitAllowed:            repo.Node.MergeCommitAllowed,
							MergeCommitMessage:            repo.Node.MergeCommitMessage,
							MergeCommitTitle:              repo.Node.MergeCommitTitle,
							MirrorURL:                     repo.Node.MirrorUrl,
							NameWithOwner:                 repo.Node.NameWithOwner,
							OpenGraphImageURL:             repo.Node.OpenGraphImageUrl,
							OwnerLogin:                    repo.Node.Owner.Login,
							PrimaryLanguage:               repo.Node.PrimaryLanguage,
							ProjectsURL:                   repo.Node.ProjectsUrl,
							PullRequestTemplates:          repo.Node.PullRequestTemplates,
							PushedAt:                      repo.Node.PushedAt,
							RebaseMergeAllowed:            repo.Node.RebaseMergeAllowed,
							SecurityPolicyURL:             repo.Node.SecurityPolicyUrl,
							SquashMergeAllowed:            repo.Node.SquashMergeAllowed,
							SquashMergeCommitMessage:      repo.Node.SquashMergeCommitMessage,
							SquashMergeCommitTitle:        repo.Node.SquashMergeCommitTitle,
							SSHURL:                        repo.Node.SshUrl,
							StargazerCount:                repo.Node.StargazerCount,
							UpdatedAt:                     repo.Node.UpdatedAt,
							URL:                           repo.Node.Url,
							UsesCustomOpenGraphImage:      repo.Node.UsesCustomOpenGraphImage,
							CanAdminister:                 repo.Node.CanAdminister,
							CanCreateProjects:             repo.Node.CanCreateProjects,
							CanSubscribe:                  repo.Node.CanSubscribe,
							CanUpdateTopics:               repo.Node.CanUpdateTopics,
							HasStarred:                    repo.Node.HasStarred,
							PossibleCommitEmails:          repo.Node.PossibleCommitEmails,
							Subscription:                  repo.Node.Subscription,
							Visibility:                    repo.Node.Visibility,
							YourPermission:                repo.Node.YourPermission,
							WebCommitSignOffRequired:      repo.Node.WebCommitSignoffRequired,
							RepositoryTopicsTotalCount:    repo.Node.RepositoryTopics.TotalCount,
							OpenIssuesTotalCount:          repo.Node.OpenIssues.TotalCount,
							WatchersTotalCount:            repo.Node.Watchers.TotalCount,
							Hooks:                         hooks,
							Topics:                        additionalRepoInfo.Topics,
							SubscribersCount:              *additionalRepoInfo.SubscribersCount,
							HasDownloads:                  *additionalRepoInfo.HasDownloads,
							HasPages:                      *additionalRepoInfo.HasPages,
							NetworkCount:                  *additionalRepoInfo.NetworkCount,
						},
						Organization: org,
						Slug:         slug,
						Permission:   repo.Permission,
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
		if !query.Organization.Team.Repositories.PageInfo.HasNextPage {
			break
		}
		variables["cursor"] = githubv4.NewString(query.Organization.Team.Repositories.PageInfo.EndCursor)
	}
	return values, nil
}
