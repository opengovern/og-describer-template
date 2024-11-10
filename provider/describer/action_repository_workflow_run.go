package describer

import (
	"context"
	"github.com/google/go-github/v55/github"
	"github.com/opengovern/og-describer-template/pkg/sdk/models"
	"github.com/opengovern/og-describer-template/provider"
	"github.com/opengovern/og-describer-template/provider/model"
	"strconv"
)

func GetAllWorkflowRuns(ctx context.Context, githubClient provider.GitHubClient, stream *models.StreamSender) ([]models.Resource, error) {
	client := githubClient.RestClient
	owner, err := getOwnerName(ctx, client)
	if err != nil {
		return nil, nil
	}
	repositories, err := getRepositoriesName(ctx, client, owner)
	if err != nil {
		return nil, nil
	}
	var values []models.Resource
	for _, repo := range repositories {
		repoValues, err := GetRepositoryWorkflowRuns(ctx, githubClient, stream, owner, repo)
		if err != nil {
			return nil, err
		}
		values = append(values, repoValues...)
	}
	return values, nil
}

func GetRepositoryWorkflowRuns(ctx context.Context, githubClient provider.GitHubClient, stream *models.StreamSender, owner, repo string) ([]models.Resource, error) {
	client := githubClient.RestClient
	opts := &github.ListWorkflowRunsOptions{
		ListOptions: github.ListOptions{PerPage: maxPagesCount},
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
						WorkflowRun:  *workflowRun,
						RepoFullName: repo,
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
				WorkflowRun:  *workflowRun,
				RepoFullName: repo,
			},
		},
	}
	return &value, nil
}
