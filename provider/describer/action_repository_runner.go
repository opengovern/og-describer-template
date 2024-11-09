package describer

import (
	"context"
	"github.com/google/go-github/v55/github"
	"github.com/opengovern/og-describer-template/pkg/sdk/models"
	"github.com/opengovern/og-describer-template/provider/model"
	"strconv"
)

const maxRunnerCount = 100

func GetRunnerList(ctx context.Context, client *github.Client, repo string) ([]models.Resource, error) {
	owner, err := getOwnerName(ctx, client)
	if err != nil {
		return nil, nil
	}
	opts := &github.ListOptions{PerPage: maxRunnerCount}
	var values []models.Resource
	for {
		runners, resp, err := client.Actions.ListRunners(ctx, owner, repo, opts)
		if err != nil {
			return nil, err
		}
		for _, runner := range runners.Runners {
			value := models.Resource{
				ID:   strconv.Itoa(int(*runner.ID)),
				Name: *runner.Name,
				Description: JSONAllFieldsMarshaller{
					Value: model.Runner{
						RunnerInfo: *runner,
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

func GetRunner(ctx context.Context, client *github.Client, repo string, runnerID int64) (*models.Resource, error) {
	owner, err := getOwnerName(ctx, client)
	if err != nil {
		return nil, nil
	}
	if runnerID == 0 || repo == "" {
		return nil, nil
	}
	runner, _, err := client.Actions.GetRunner(ctx, owner, repo, runnerID)
	if err != nil {
		return nil, err
	}
	value := models.Resource{
		ID:   strconv.Itoa(int(*runner.ID)),
		Name: *runner.Name,
		Description: JSONAllFieldsMarshaller{
			Value: model.Runner{
				RunnerInfo: *runner,
			},
		},
	}
	return &value, nil
}
