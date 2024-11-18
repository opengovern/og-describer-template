package describer

import (
	"context"
	"fmt"
	"github.com/opengovern/og-describer-github/pkg/sdk/models"
	"github.com/opengovern/og-describer-github/provider/model"
	"github.com/shurcooL/githubv4"
	steampipemodels "github.com/turbot/steampipe-plugin-github/github/models"
)

func GetAllBranches(ctx context.Context, githubClient GitHubClient, organizationName string, stream *models.StreamSender) ([]models.Resource, error) {
	client := githubClient.RestClient
	owner := organizationName
	repositories, err := getRepositories(ctx, client, owner)
	if err != nil {
		return nil, nil
	}
	var values []models.Resource
	for _, repo := range repositories {
		repoValues, err := GetRepositoryBranches(ctx, githubClient, stream, owner, repo.GetName())
		if err != nil {
			return nil, err
		}
		values = append(values, repoValues...)
	}
	return values, nil
}

func GetRepositoryBranches(ctx context.Context, githubClient GitHubClient, stream *models.StreamSender, owner, repo string) ([]models.Resource, error) {
	graphQLClient := githubClient.GraphQLClient
	restClient := githubClient.RestClient
	var query struct {
		RateLimit  steampipemodels.RateLimit
		Repository struct {
			Refs struct {
				TotalCount int
				PageInfo   steampipemodels.PageInfo
				Edges      []struct {
					Node steampipemodels.Branch
				}
			} `graphql:"refs(refPrefix: \"refs/heads/\", first: $pageSize, after: $cursor)"`
		} `graphql:"repository(owner: $owner, name: $repo)"`
	}
	variables := map[string]interface{}{
		"owner":    githubv4.String(owner),
		"repo":     githubv4.String(repo),
		"pageSize": githubv4.Int(pageSize),
		"cursor":   (*githubv4.String)(nil),
	}
	appendBranchColumnIncludes(&variables, branchCols())
	repoFullName := formRepositoryFullName(owner, repo)
	var values []models.Resource
	for {
		err := graphQLClient.Query(ctx, &query, variables)
		if err != nil {
			return nil, err
		}
		for _, branch := range query.Repository.Refs.Edges {
			branchInfo, _, err := restClient.Repositories.GetBranch(ctx, owner, repo, branch.Node.Name, true)
			if err != nil {
				return nil, err
			}
			protected := branchInfo.GetProtected()
			id := fmt.Sprintf("%s/%s/%s", owner, repo, branch.Node.Name)
			value := models.Resource{
				ID:   id,
				Name: branch.Node.Name,
				Description: JSONAllFieldsMarshaller{
					Value: model.BranchDescription{
						Name:                 branch.Node.Name,
						Commit:               branch.Node.Target.Commit,
						BranchProtectionRule: branch.Node.BranchProtectionRule,
						RepoFullName:         repoFullName,
						Protected:            protected,
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
		if !query.Repository.Refs.PageInfo.HasNextPage {
			break
		}
		variables["cursor"] = githubv4.NewString(query.Repository.Refs.PageInfo.EndCursor)
	}
	return values, nil
}
