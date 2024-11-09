package describer

import (
	"context"
	"github.com/google/go-github/v55/github"
	"github.com/opengovern/og-describer-template/pkg/sdk/models"
	"github.com/opengovern/og-describer-template/provider/model"
	"strconv"
)

const maxWorkflowRunCount = 100

func GetRepoWorkflowRunList(ctx context.Context, client *github.Client, repo string) ([]models.Resource, error) {
	owner, err := getOwnerName(ctx, client)
	if err != nil {
		return nil, nil
	}
	opts := &github.ListWorkflowRunsOptions{
		ListOptions: github.ListOptions{PerPage: maxWorkflowRunCount},
	}
	var values []models.Resource
	for {
		workflowRuns, resp, err := client.Actions.ListRepositoryWorkflowRuns(ctx, owner, repo, opts)
		if err != nil {
			return nil, err
		}
		for _, workflowRun := range workflowRuns.WorkflowRuns {
			value := models.Resource{
				ID:   strconv.Itoa(int(*workflowRun.ID)),
				Name: *workflowRun.Name,
				Description: JSONAllFieldsMarshaller{
					Value: model.WorkflowRun{
						WorkflowRunInfo: *workflowRun,
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

func GetRepoWorkflowRun(ctx context.Context, client *github.Client, repo string, workflowRunID int64) (*models.Resource, error) {
	owner, err := getOwnerName(ctx, client)
	if err != nil {
		return nil, nil
	}
	if workflowRunID == 0 || repo == "" {
		return nil, nil
	}
	workflowRun, _, err := client.Actions.GetWorkflowRunByID(ctx, owner, repo, workflowRunID)
	if err != nil {
		return nil, err
	}
	value := models.Resource{
		ID:   strconv.Itoa(int(*workflowRun.ID)),
		Name: *workflowRun.Name,
		Description: JSONAllFieldsMarshaller{
			Value: model.WorkflowRun{
				WorkflowRunInfo: *workflowRun,
			},
		},
	}
	return &value, nil
}
