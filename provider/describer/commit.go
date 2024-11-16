package describer

import (
	"context"
	"github.com/opengovern/og-describer-github/pkg/sdk/models"
	"github.com/opengovern/og-describer-github/provider/model"
	"github.com/shurcooL/githubv4"
	steampipemodels "github.com/turbot/steampipe-plugin-github/github/models"
)

func GetAllCommits(ctx context.Context, githubClient GitHubClient, stream *models.StreamSender) ([]models.Resource, error) {
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
		repoValues, err := GetRepositoryCommits(ctx, githubClient, stream, owner, repo.GetName())
		if err != nil {
			return nil, err
		}
		values = append(values, repoValues...)
	}
	return values, nil
}

func GetRepositoryCommits(ctx context.Context, githubClient GitHubClient, stream *models.StreamSender, owner, repo string) ([]models.Resource, error) {
	client := githubClient.GraphQLClient
	var query struct {
		RateLimit  steampipemodels.RateLimit
		Repository struct {
			DefaultBranchRef struct {
				Target struct {
					Commit struct {
						History struct {
							TotalCount int
							PageInfo   steampipemodels.PageInfo
							Nodes      []steampipemodels.Commit
						} `graphql:"history(first: $pageSize, after: $cursor, since: $since, until: $until)"`
					} `graphql:"... on Commit"`
				}
			}
		} `graphql:"repository(owner: $owner, name: $name)"`
	}
	variables := map[string]interface{}{
		"owner":    githubv4.String(owner),
		"name":     githubv4.String(repo),
		"pageSize": githubv4.Int(pageSize),
		"cursor":   (*githubv4.String)(nil),
		"since":    (*githubv4.GitTimestamp)(nil),
		"until":    (*githubv4.GitTimestamp)(nil),
	}
	appendCommitColumnIncludes(&variables, commitCols())
	repoFullName := formRepositoryFullName(owner, repo)
	var values []models.Resource
	for {
		err := client.Query(ctx, &query, variables)
		if err != nil {
			return nil, err
		}
		for _, commit := range query.Repository.DefaultBranchRef.Target.Commit.History.Nodes {
			value := models.Resource{
				ID:   commit.Sha,
				Name: commit.Sha,
				Description: JSONAllFieldsMarshaller{
					Value: model.CommitDescription{
						Commit:         commit,
						RepoFullName:   repoFullName,
						AuthorLogin:    commit.Author.User.Login,
						CommitterLogin: commit.Committer.User.Login,
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
		if !query.Repository.DefaultBranchRef.Target.Commit.History.PageInfo.HasNextPage {
			break
		}
		variables["cursor"] = githubv4.NewString(query.Repository.DefaultBranchRef.Target.Commit.History.PageInfo.EndCursor)
	}
	return values, nil
}
