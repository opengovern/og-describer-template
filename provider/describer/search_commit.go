package describer

import (
	"context"
	"fmt"
	"github.com/google/go-github/v55/github"
	"github.com/opengovern/og-describer-github/pkg/sdk/models"
	"github.com/opengovern/og-describer-github/provider/model"
)

func GetAllSearchCommits(ctx context.Context, githubClient GitHubClient, stream *models.StreamSender) ([]models.Resource, error) {
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
		repoValues, err := GetSearchCommits(ctx, githubClient, stream, owner, repo.GetName())
		if err != nil {
			return nil, err
		}
		values = append(values, repoValues...)
	}
	return values, nil
}

func GetSearchCommits(ctx context.Context, githubClient GitHubClient, stream *models.StreamSender, owner, repo string) ([]models.Resource, error) {
	client := githubClient.RestClient
	repoFullName := formRepositoryFullName(owner, repo)
	query := fmt.Sprintf("repo:%s", repoFullName)
	opt := &github.SearchOptions{
		ListOptions: github.ListOptions{PerPage: 100},
		TextMatch:   true,
	}
	var values []models.Resource
	for {
		result, resp, err := client.Search.Commits(ctx, query, opt)
		if err != nil {
			return nil, err
		}
		codeResults := result.Commits
		for _, codeResult := range codeResults {
			value := models.Resource{
				ID:   *codeResult.SHA,
				Name: *codeResult.Commit.Message,
				Description: JSONAllFieldsMarshaller{
					Value: model.SearchCommitDescription{
						CommitResult: *codeResult,
						RepoFullName: repoFullName,
						Query:        query,
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
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}
	return values, nil
}
