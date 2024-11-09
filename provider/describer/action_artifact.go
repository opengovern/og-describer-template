package describer

import (
	"context"
	"github.com/google/go-github/v55/github"
	"github.com/opengovern/og-describer-template/pkg/sdk/models"
	"github.com/opengovern/og-describer-template/provider/model"
	"strconv"
)

func GetArtifactList(ctx context.Context, client *github.Client, repo string) ([]models.Resource, error) {
	owner, _, err := client.Users.Get(ctx, "")
	if err != nil {
		return nil, err
	}
	ownerName := *owner.Name
	opts := &github.ListOptions{PerPage: 100}
	var values []models.Resource
	for {
		artifacts, resp, err := client.Actions.ListArtifacts(ctx, ownerName, repo, opts)
		if err != nil {
			return nil, err
		}
		for _, artifact := range artifacts.Artifacts {
			value := models.Resource{
				ID:   strconv.Itoa(int(*artifact.ID)),
				Name: *artifact.Name,
				Description: JSONAllFieldsMarshaller{
					Value: model.Artifact{
						ArtifactInfo: *artifact,
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

func GetArtifact(ctx context.Context, client *github.Client, repo string, artifactID int64) (*models.Resource, error) {
	owner, _, err := client.Users.Get(ctx, "")
	if err != nil {
		return nil, err
	}
	ownerName := *owner.Name
	if artifactID == 0 || repo == "" {
		return nil, nil
	}
	artifact, _, err := client.Actions.GetArtifact(ctx, ownerName, repo, artifactID)
	if err != nil {
		return nil, err
	}
	value := models.Resource{
		ID:   strconv.Itoa(int(*artifact.ID)),
		Name: *artifact.Name,
		Description: JSONAllFieldsMarshaller{
			Value: model.Artifact{
				ArtifactInfo: *artifact,
			},
		},
	}
	return &value, nil
}
