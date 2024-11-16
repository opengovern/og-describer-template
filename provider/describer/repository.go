package describer

import (
	"context"
	"github.com/google/go-github/v55/github"
	"github.com/opengovern/og-describer-github/pkg/sdk/models"
	"github.com/opengovern/og-describer-github/provider/model"
	"github.com/shurcooL/githubv4"
	steampipemodels "github.com/turbot/steampipe-plugin-github/github/models"
	"strconv"
	"strings"
)

func GetRepositoryList(ctx context.Context, githubClient GitHubClient, stream *models.StreamSender) ([]models.Resource, error) {
	client := githubClient.GraphQLClient
	var query struct {
		RateLimit steampipemodels.RateLimit
		Viewer    struct {
			Repositories struct {
				PageInfo   steampipemodels.PageInfo
				TotalCount int
				Nodes      []steampipemodels.Repository
			} `graphql:"repositories(first: $pageSize, after: $cursor, affiliations: [COLLABORATOR, OWNER, ORGANIZATION_MEMBER], ownerAffiliations: [COLLABORATOR, OWNER, ORGANIZATION_MEMBER])"`
		}
	}
	variables := map[string]interface{}{
		"pageSize": githubv4.Int(repoPageSize),
		"cursor":   (*githubv4.String)(nil),
	}
	columnNames := repositoryCols()
	appendRepoColumnIncludes(&variables, columnNames)
	var values []models.Resource
	for {
		err := client.Query(ctx, &query, variables)
		if err != nil {
			return nil, err
		}
		for _, repo := range query.Viewer.Repositories.Nodes {
			hooks, err := GetRepositoryHooks(ctx, githubClient.RestClient, repo.Name)
			if err != nil {
				return nil, err
			}
			additionalRepoInfo, err := GetRepositoryAdditionalData(ctx, githubClient.RestClient, repo.Name)
			value := models.Resource{
				ID:   strconv.Itoa(repo.Id),
				Name: repo.Name,
				Description: JSONAllFieldsMarshaller{
					Value: model.RepositoryDescription{
						ID:                            repo.Id,
						NodeID:                        repo.NodeId,
						Name:                          repo.Name,
						AllowUpdateBranch:             repo.AllowUpdateBranch,
						ArchivedAt:                    repo.ArchivedAt,
						AutoMergeAllowed:              repo.AutoMergeAllowed,
						CodeOfConduct:                 repo.CodeOfConduct,
						ContactLinks:                  repo.ContactLinks,
						CreatedAt:                     repo.CreatedAt,
						DefaultBranchRef:              repo.DefaultBranchRef,
						DeleteBranchOnMerge:           repo.DeleteBranchOnMerge,
						Description:                   repo.Description,
						DiskUsage:                     repo.DiskUsage,
						ForkCount:                     repo.ForkCount,
						ForkingAllowed:                repo.ForkingAllowed,
						FundingLinks:                  repo.FundingLinks,
						HasDiscussionsEnabled:         repo.HasDiscussionsEnabled,
						HasIssuesEnabled:              repo.HasIssuesEnabled,
						HasProjectsEnabled:            repo.HasProjectsEnabled,
						HasVulnerabilityAlertsEnabled: repo.HasVulnerabilityAlertsEnabled,
						HasWikiEnabled:                repo.HasWikiEnabled,
						HomepageURL:                   repo.HomepageUrl,
						InteractionAbility:            repo.InteractionAbility,
						IsArchived:                    repo.IsArchived,
						IsBlankIssuesEnabled:          repo.IsBlankIssuesEnabled,
						IsDisabled:                    repo.IsDisabled,
						IsEmpty:                       repo.IsEmpty,
						IsFork:                        repo.IsFork,
						IsInOrganization:              repo.IsInOrganization,
						IsLocked:                      repo.IsLocked,
						IsMirror:                      repo.IsMirror,
						IsPrivate:                     repo.IsPrivate,
						IsSecurityPolicyEnabled:       repo.IsSecurityPolicyEnabled,
						IsTemplate:                    repo.IsTemplate,
						IsUserConfigurationRepository: repo.IsUserConfigurationRepository,
						IssueTemplates:                repo.IssueTemplates,
						LicenseInfo:                   repo.LicenseInfo,
						LockReason:                    repo.LockReason,
						MergeCommitAllowed:            repo.MergeCommitAllowed,
						MergeCommitMessage:            repo.MergeCommitMessage,
						MergeCommitTitle:              repo.MergeCommitTitle,
						MirrorURL:                     repo.MirrorUrl,
						NameWithOwner:                 repo.NameWithOwner,
						OpenGraphImageURL:             repo.OpenGraphImageUrl,
						OwnerLogin:                    repo.Owner.Login,
						PrimaryLanguage:               repo.PrimaryLanguage,
						ProjectsURL:                   repo.ProjectsUrl,
						PullRequestTemplates:          repo.PullRequestTemplates,
						PushedAt:                      repo.PushedAt,
						RebaseMergeAllowed:            repo.RebaseMergeAllowed,
						SecurityPolicyURL:             repo.SecurityPolicyUrl,
						SquashMergeAllowed:            repo.SquashMergeAllowed,
						SquashMergeCommitMessage:      repo.SquashMergeCommitMessage,
						SquashMergeCommitTitle:        repo.SquashMergeCommitTitle,
						SSHURL:                        repo.SshUrl,
						StargazerCount:                repo.StargazerCount,
						UpdatedAt:                     repo.UpdatedAt,
						URL:                           repo.Url,
						UsesCustomOpenGraphImage:      repo.UsesCustomOpenGraphImage,
						CanAdminister:                 repo.CanAdminister,
						CanCreateProjects:             repo.CanCreateProjects,
						CanSubscribe:                  repo.CanSubscribe,
						CanUpdateTopics:               repo.CanUpdateTopics,
						HasStarred:                    repo.HasStarred,
						PossibleCommitEmails:          repo.PossibleCommitEmails,
						Subscription:                  repo.Subscription,
						Visibility:                    repo.Visibility,
						YourPermission:                repo.YourPermission,
						WebCommitSignOffRequired:      repo.WebCommitSignoffRequired,
						RepositoryTopicsTotalCount:    repo.RepositoryTopics.TotalCount,
						OpenIssuesTotalCount:          repo.OpenIssues.TotalCount,
						WatchersTotalCount:            repo.Watchers.TotalCount,
						Hooks:                         hooks,
						Topics:                        additionalRepoInfo.Topics,
						SubscribersCount:              *additionalRepoInfo.SubscribersCount,
						HasDownloads:                  *additionalRepoInfo.HasDownloads,
						HasPages:                      *additionalRepoInfo.HasPages,
						NetworkCount:                  *additionalRepoInfo.NetworkCount,
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
		if !query.Viewer.Repositories.PageInfo.HasNextPage {
			break
		}
		variables["cursor"] = githubv4.NewString(query.Viewer.Repositories.PageInfo.EndCursor)
	}
	return values, nil
}

func GetRepositoryAdditionalData(ctx context.Context, client *github.Client, repo string) (*github.Repository, error) {
	owner, err := getOwnerName(ctx, client)
	if err != nil {
		return nil, nil
	}
	repository, _, err := client.Repositories.Get(ctx, owner, repo)
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			return nil, nil
		}
		return nil, nil
	}
	if repository == nil {
		return nil, nil
	}
	return repository, nil
}

func GetRepositoryHooks(ctx context.Context, client *github.Client, repo string) ([]*github.Hook, error) {
	owner, err := getOwnerName(ctx, client)
	if err != nil {
		return nil, nil
	}
	var repositoryHooks []*github.Hook
	opt := &github.ListOptions{PerPage: pageSize}
	for {
		hooks, resp, err := client.Repositories.ListHooks(ctx, owner, repo, opt)
		if err != nil && strings.Contains(err.Error(), "Not Found") {
			return nil, nil
		} else if err != nil {
			return nil, err
		}
		repositoryHooks = append(repositoryHooks, hooks...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}
	return repositoryHooks, nil
}
