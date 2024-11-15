package describer

import (
	"context"
	"github.com/google/go-github/v55/github"
	"github.com/opengovern/og-describer-github/pkg/sdk/models"
	"github.com/opengovern/og-describer-github/provider/model"
	"strconv"
)

func GetAllWorkflowRuns(ctx context.Context, githubClient GitHubClient, stream *models.StreamSender) ([]models.Resource, error) {
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
		repoValues, err := GetRepositoryWorkflowRuns(ctx, githubClient, stream, owner, repo.GetName())
		if err != nil {
			return nil, err
		}
		values = append(values, repoValues...)
	}
	return values, nil
}

func GetRepositoryWorkflowRuns(ctx context.Context, githubClient GitHubClient, stream *models.StreamSender, owner, repo string) ([]models.Resource, error) {
	client := githubClient.RestClient
	opts := &github.ListWorkflowRunsOptions{
		ListOptions: github.ListOptions{PerPage: maxPagesCount},
	}
	repoFullName := formRepositoryFullName(owner, repo)
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
					Value: model.WorkflowRunDescription{
						ID:                 workflowRun.ID,
						Name:               workflowRun.Name,
						NodeID:             workflowRun.NodeID,
						HeadBranch:         workflowRun.HeadBranch,
						HeadSHA:            workflowRun.HeadSHA,
						RunNumber:          workflowRun.RunNumber,
						RunAttempt:         workflowRun.RunAttempt,
						Event:              workflowRun.Event,
						DisplayTitle:       workflowRun.DisplayTitle,
						Status:             workflowRun.Status,
						Conclusion:         workflowRun.Conclusion,
						WorkflowID:         workflowRun.WorkflowID,
						CheckSuiteID:       workflowRun.CheckSuiteID,
						CheckSuiteNodeID:   workflowRun.CheckSuiteNodeID,
						URL:                workflowRun.URL,
						HTMLURL:            workflowRun.HTMLURL,
						PullRequests:       workflowRun.PullRequests,
						CreatedAt:          workflowRun.CreatedAt,
						UpdatedAt:          workflowRun.UpdatedAt,
						RunStartedAt:       workflowRun.RunStartedAt,
						JobsURL:            workflowRun.JobsURL,
						LogsURL:            workflowRun.LogsURL,
						CheckSuiteURL:      workflowRun.CheckSuiteURL,
						ArtifactsURL:       workflowRun.ArtifactsURL,
						CancelURL:          workflowRun.CancelURL,
						RerunURL:           workflowRun.RerunURL,
						PreviousAttemptURL: workflowRun.PreviousAttemptURL,
						HeadCommit:         workflowRun.HeadCommit,
						WorkflowURL:        workflowRun.WorkflowURL,
						Repository:         workflowRun.Repository,
						HeadRepository:     workflowRun.HeadRepository,
						Actor:              workflowRun.Actor,
						TriggeringActor:    workflowRun.TriggeringActor,
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
	repoFullName := formRepositoryFullName(owner, repo)
	value := models.Resource{
		ID:   strconv.Itoa(int(*workflowRun.ID)),
		Name: *workflowRun.Name,
		Description: JSONAllFieldsMarshaller{
			Value: model.WorkflowRunDescription{
				ID:                 workflowRun.ID,
				Name:               workflowRun.Name,
				NodeID:             workflowRun.NodeID,
				HeadBranch:         workflowRun.HeadBranch,
				HeadSHA:            workflowRun.HeadSHA,
				RunNumber:          workflowRun.RunNumber,
				RunAttempt:         workflowRun.RunAttempt,
				Event:              workflowRun.Event,
				DisplayTitle:       workflowRun.DisplayTitle,
				Status:             workflowRun.Status,
				Conclusion:         workflowRun.Conclusion,
				WorkflowID:         workflowRun.WorkflowID,
				CheckSuiteID:       workflowRun.CheckSuiteID,
				CheckSuiteNodeID:   workflowRun.CheckSuiteNodeID,
				URL:                workflowRun.URL,
				HTMLURL:            workflowRun.HTMLURL,
				PullRequests:       workflowRun.PullRequests,
				CreatedAt:          workflowRun.CreatedAt,
				UpdatedAt:          workflowRun.UpdatedAt,
				RunStartedAt:       workflowRun.RunStartedAt,
				JobsURL:            workflowRun.JobsURL,
				LogsURL:            workflowRun.LogsURL,
				CheckSuiteURL:      workflowRun.CheckSuiteURL,
				ArtifactsURL:       workflowRun.ArtifactsURL,
				CancelURL:          workflowRun.CancelURL,
				RerunURL:           workflowRun.RerunURL,
				PreviousAttemptURL: workflowRun.PreviousAttemptURL,
				HeadCommit:         workflowRun.HeadCommit,
				WorkflowURL:        workflowRun.WorkflowURL,
				Repository:         workflowRun.Repository,
				HeadRepository:     workflowRun.HeadRepository,
				Actor:              workflowRun.Actor,
				TriggeringActor:    workflowRun.TriggeringActor,
				RepoFullName:       repoFullName,
			},
		},
	}
	return &value, nil
}
