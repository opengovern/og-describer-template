package describer

import (
	"context"
	"github.com/google/go-github/v55/github"
	"github.com/opengovern/og-describer-github/pkg/sdk/models"
	"github.com/opengovern/og-describer-github/provider/model"
)

func GetAllSecrets(ctx context.Context, githubClient GitHubClient, stream *models.StreamSender) ([]models.Resource, error) {
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
		repoValues, err := GetRepositorySecrets(ctx, githubClient, stream, owner, repo.GetName())
		if err != nil {
			return nil, err
		}
		values = append(values, repoValues...)
	}
	return values, nil
}

func GetRepositorySecrets(ctx context.Context, githubClient GitHubClient, stream *models.StreamSender, owner, repo string) ([]models.Resource, error) {
	client := githubClient.RestClient
	opts := &github.ListOptions{PerPage: maxPagesCount}
	repoFullName := formRepositoryFullName(owner, repo)
	var values []models.Resource
	for {
		secrets, resp, err := client.Actions.ListRepoSecrets(ctx, owner, repo, opts)
		if err != nil {
			return nil, err
		}
		for _, secret := range secrets.Secrets {
			value := models.Resource{
				ID:   secret.Name,
				Name: secret.Name,
				Description: JSONAllFieldsMarshaller{
					Value: model.SecretDescription{
						Secret:       *secret,
						RepoFullName: repoFullName,
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

func GetSecret(ctx context.Context, client *github.Client, repo, secretName string) (*models.Resource, error) {
	owner, err := getOwnerName(ctx, client)
	if err != nil {
		return nil, nil
	}
	if secretName == "" || repo == "" {
		return nil, nil
	}
	type GetResponse struct {
		secret *github.Secret
		resp   *github.Response
	}
	secret, _, err := client.Actions.GetRepoSecret(ctx, owner, repo, secretName)
	if err != nil {
		return nil, err
	}
	repoFullName := formRepositoryFullName(owner, repo)
	value := models.Resource{
		ID:   secret.Name,
		Name: secret.Name,
		Description: JSONAllFieldsMarshaller{
			Value: model.SecretDescription{
				Secret:       *secret,
				RepoFullName: repoFullName,
			},
		},
	}
	return &value, nil
}
