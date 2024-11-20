package describer

import (
	"context"
	"github.com/opengovern/og-describer-github/pkg/sdk/models"
	"github.com/opengovern/og-describer-github/provider/model"
)

func GetAllTrees(ctx context.Context, githubClient GitHubClient, organizationName string, stream *models.StreamSender) ([]models.Resource, error) {
	client := githubClient.RestClient

	repositories, err := getRepositories(ctx, client, organizationName)
	if err != nil {
		return nil, nil
	}
	var values []models.Resource
	for _, repo := range repositories {
		repoValues, err := GetRepositoryTrees(ctx, githubClient, stream, organizationName, repo.GetName())
		if err != nil {
			return nil, err
		}
		values = append(values, repoValues...)
	}
	return values, nil
}

func GetRepositoryTrees(ctx context.Context, githubClient GitHubClient, stream *models.StreamSender, owner, repo string) ([]models.Resource, error) {
	client := githubClient.RestClient
	repository, _, err := client.Repositories.Get(ctx, owner, repo)
	if err != nil {
		return nil, err
	}
	branch, _, err := client.Repositories.GetBranch(ctx, owner, repo, repository.GetDefaultBranch(), false)
	if err != nil {
		return nil, err
	}
	sha := branch.Commit.GetSHA()
	tree, _, err := client.Git.GetTree(ctx, owner, repo, sha, true)
	if err != nil {
		return nil, err
	}
	entries := tree.Entries
	var values []models.Resource
	repoFullName := formRepositoryFullName(owner, repo)
	for _, entry := range entries {
		value := models.Resource{
			ID:   entry.GetSHA(),
			Name: entry.GetSHA(),
			Description: JSONAllFieldsMarshaller{
				Value: model.TreeDescription{
					TreeSHA:            sha,
					RepositoryFullName: repoFullName,
					Recursive:          true,
					Truncated:          tree.GetTruncated(),
					SHA:                entry.GetSHA(),
					Path:               entry.GetPath(),
					Mode:               entry.GetMode(),
					Type:               entry.GetType(),
					Size:               entry.GetSize(),
					URL:                entry.GetURL(),
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
	return values, nil
}