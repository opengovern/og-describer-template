package describer

import (
	"context"
	"github.com/google/go-github/v55/github"
	"github.com/opengovern/og-describer-template/pkg/sdk/models"
	"github.com/opengovern/og-describer-template/provider/model"
)

const maxSecretCount = 100

func GetSecretList(ctx context.Context, client *github.Client, repo string) ([]models.Resource, error) {
	owner, err := getOwnerName(ctx, client)
	if err != nil {
		return nil, nil
	}
	opts := &github.ListOptions{PerPage: maxSecretCount}
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
					Value: model.Secret{
						SecretInfo: *secret,
					},
				},
			}
			values = append(values, value)
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
	value := models.Resource{
		ID:   secret.Name,
		Name: secret.Name,
		Description: JSONAllFieldsMarshaller{
			Value: model.Secret{
				SecretInfo: *secret,
			},
		},
	}
	return &value, nil
}
