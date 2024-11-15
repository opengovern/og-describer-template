package describer

import (
	"context"
	"github.com/opengovern/og-describer-github/pkg/sdk/models"
	"github.com/opengovern/og-describer-github/provider/model"
)

func GetAllBlobs(ctx context.Context, githubClient GitHubClient, stream *models.StreamSender) ([]models.Resource, error) {
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
		repoValues, err := GetRepositoryBlobs(ctx, githubClient, stream, owner, repo.GetName())
		if err != nil {
			return nil, err
		}
		values = append(values, repoValues...)
	}
	return values, nil
}

func GetRepositoryBlobs(ctx context.Context, githubClient GitHubClient, stream *models.StreamSender, owner, repo string) ([]models.Resource, error) {
	client := githubClient.RestClient
	fileSHAs, err := getFileSHAs(client, owner, repo)
	if err != nil {
		return nil, nil
	}
	var values []models.Resource
	for _, sha := range fileSHAs {
		blobValue, err := GetBlob(ctx, githubClient, owner, repo, sha)
		if err != nil {
			return nil, err
		}
		if stream != nil {
			if err := (*stream)(*blobValue); err != nil {
				return nil, err
			}
		} else {
			values = append(values, *blobValue)
		}
	}
	return values, nil
}

func GetBlob(ctx context.Context, githubClient GitHubClient, owner, repo, sha string) (*models.Resource, error) {
	client := githubClient.RestClient
	blob, _, err := client.Git.GetBlob(ctx, owner, repo, sha)
	if err != nil {
		return nil, err
	}
	repoFullName := formRepositoryFullName(owner, repo)
	var value models.Resource
	if blob != nil {
		value = models.Resource{
			ID:   *blob.SHA,
			Name: *blob.SHA,
			Description: JSONAllFieldsMarshaller{
				Value: model.BlobDescription{
					Blob:         *blob,
					RepoFullName: repoFullName,
				},
			},
		}
	}
	return &value, nil
}
