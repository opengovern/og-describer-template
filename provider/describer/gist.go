package describer

import (
	"context"
	"github.com/google/go-github/v55/github"
	"github.com/opengovern/og-describer-github/pkg/sdk/models"
	"github.com/opengovern/og-describer-github/provider/model"
)

func GetGistList(ctx context.Context, githubClient GitHubClient, stream *models.StreamSender) ([]models.Resource, error) {
	client := githubClient.RestClient
	opt := &github.GistListOptions{ListOptions: github.ListOptions{PerPage: maxPagesCount}}
	var values []models.Resource
	for {
		gists, resp, err := client.Gists.List(ctx, "", opt)
		if err != nil {
			return nil, err
		}
		for _, gist := range gists {
			value := models.Resource{
				ID:   *gist.ID,
				Name: *gist.ID,
				Description: JSONAllFieldsMarshaller{
					Value: model.GistDescription{
						Gist:       *gist,
						OwnerID:    int(*gist.Owner.ID),
						OwnerLogin: *gist.Owner.Login,
						OwnerType:  *gist.Owner.Type,
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
