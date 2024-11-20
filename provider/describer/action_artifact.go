package describer

import (
	"context"
	"github.com/google/go-github/v55/github"
	"github.com/opengovern/og-describer-github/pkg/sdk/models"
	"github.com/opengovern/og-describer-github/provider/model"
	"strconv"
)

func GetAllArtifacts(ctx context.Context, githubClient GitHubClient, organizationName string, stream *models.StreamSender) ([]models.Resource, error) {
	client := githubClient.RestClient
	owner := organizationName
	repositories, err := getRepositories(ctx, client, owner)
	if err != nil {
		return nil, nil
	}
	var values []models.Resource
	for _, repo := range repositories {
		repoValues, err := GetRepositoryArtifacts(ctx, githubClient, stream, owner, repo.GetName())
		if err != nil {
			return nil, err
		}
		values = append(values, repoValues...)
	}
	return values, nil
}

func GetRepositoryArtifacts(ctx context.Context, githubClient GitHubClient, stream *models.StreamSender, owner, repo string) ([]models.Resource, error) {
	client := githubClient.RestClient
	opts := &github.ListOptions{PerPage: maxPagesCount}
	repoFullName := formRepositoryFullName(owner, repo)
	var values []models.Resource
	for {
		artifacts, resp, err := client.Actions.ListArtifacts(ctx, owner, repo, opts)
		if err != nil {
			return nil, err
		}
		for _, artifact := range artifacts.Artifacts {
			value := models.Resource{
				ID:   strconv.Itoa(int(artifact.GetID())),
				Name: artifact.GetName(),
				Description: JSONAllFieldsMarshaller{
					Value: model.ArtifactDescription{
						ID:                 artifact.GetID(),
						NodeID:             artifact.GetNodeID(),
						Name:               artifact.GetName(),
						SizeInBytes:        artifact.GetSizeInBytes(),
						ArchiveDownloadURL: artifact.GetArchiveDownloadURL(),
						Expired:            artifact.GetExpired(),
						CreatedAt:          artifact.GetCreatedAt(),
						ExpiresAt:          artifact.GetExpiresAt(),
						RepoFullName:       repoFullName,
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
		opts.Page = resp.NextPage
	}
	return values, nil
}

func GetArtifact(ctx context.Context, githubClient GitHubClient, organizationName string, repositoryName string, resourceID string, stream *models.StreamSender) (*models.Resource, error) {
	client := githubClient.RestClient
	artifactID, err := strconv.ParseInt(resourceID, 10, 64)
	if err != nil {
		return nil, err
	}
	repoFullName := formRepositoryFullName(organizationName, repositoryName)
	artifact, _, err := client.Actions.GetArtifact(ctx, organizationName, repositoryName, artifactID)
	if err != nil {
		return nil, err
	}
	value := models.Resource{
		ID:   strconv.Itoa(int(artifact.GetID())),
		Name: artifact.GetName(),
		Description: JSONAllFieldsMarshaller{
			Value: model.ArtifactDescription{
				ID:                 artifact.GetID(),
				NodeID:             artifact.GetNodeID(),
				Name:               artifact.GetName(),
				SizeInBytes:        artifact.GetSizeInBytes(),
				ArchiveDownloadURL: artifact.GetArchiveDownloadURL(),
				Expired:            artifact.GetExpired(),
				CreatedAt:          artifact.GetCreatedAt(),
				ExpiresAt:          artifact.GetExpiresAt(),
				RepoFullName:       repoFullName,
			},
		},
	}
	if stream != nil {
		if err := (*stream)(value); err != nil {
			return nil, err
		}
	}

	return &value, nil
}
