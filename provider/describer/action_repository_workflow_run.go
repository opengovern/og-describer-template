package describer

import (
	"context"
	"github.com/google/go-github/v55/github"
	"github.com/opengovern/og-describer-github/pkg/sdk/models"
	"github.com/opengovern/og-describer-github/provider/model"
	"strconv"
)

func GetAllWorkflowRuns(ctx context.Context, githubClient GitHubClient, organizationName string, stream *models.StreamSender) ([]models.Resource, error) {
	client := githubClient.RestClient
	owner := organizationName
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
				ID:   strconv.Itoa(int(workflowRun.GetID())),
				Name: workflowRun.GetName(),
				Description: JSONAllFieldsMarshaller{
					Value: model.WorkflowRunDescription{
						ID:                 workflowRun.GetID(),
						Name:               workflowRun.GetName(),
						NodeID:             workflowRun.GetNodeID(),
						HeadBranch:         workflowRun.GetHeadBranch(),
						HeadSHA:            workflowRun.GetHeadSHA(),
						RunNumber:          workflowRun.GetRunNumber(),
						RunAttempt:         workflowRun.GetRunAttempt(),
						Event:              workflowRun.GetEvent(),
						DisplayTitle:       workflowRun.GetDisplayTitle(),
						Status:             workflowRun.GetStatus(),
						Conclusion:         workflowRun.GetConclusion(),
						WorkflowID:         workflowRun.GetWorkflowID(),
						CheckSuiteID:       workflowRun.GetCheckSuiteID(),
						CheckSuiteNodeID:   workflowRun.GetCheckSuiteNodeID(),
						URL:                workflowRun.GetURL(),
						HTMLURL:            workflowRun.GetHTMLURL(),
						PullRequests:       workflowRun.PullRequests,
						CreatedAt:          workflowRun.GetCreatedAt(),
						UpdatedAt:          workflowRun.GetUpdatedAt(),
						RunStartedAt:       workflowRun.GetRunStartedAt(),
						JobsURL:            workflowRun.GetJobsURL(),
						LogsURL:            workflowRun.GetLogsURL(),
						CheckSuiteURL:      workflowRun.GetCheckSuiteURL(),
						ArtifactsURL:       workflowRun.GetArtifactsURL(),
						CancelURL:          workflowRun.GetCancelURL(),
						RerunURL:           workflowRun.GetRerunURL(),
						PreviousAttemptURL: workflowRun.GetPreviousAttemptURL(),
						HeadCommit:         workflowRun.GetHeadCommit(),
						WorkflowURL:        workflowRun.GetWorkflowURL(),
						Repository:         workflowRun.GetRepository(),
						HeadRepository:     workflowRun.GetHeadRepository(),
						Actor:              workflowRun.GetActor(),
						TriggeringActor:    workflowRun.GetTriggeringActor(),
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

func GetRepoWorkflowRun(ctx context.Context, client *github.Client, organizationName string, repo string, workflowRunID int64, stream *models.StreamSender) (*models.Resource, error) {
	owner := organizationName
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
				ID:                 workflowRun.GetID(),
				Name:               workflowRun.GetName(),
				NodeID:             workflowRun.GetNodeID(),
				HeadBranch:         workflowRun.GetHeadBranch(),
				HeadSHA:            workflowRun.GetHeadSHA(),
				RunNumber:          workflowRun.GetRunNumber(),
				RunAttempt:         workflowRun.GetRunAttempt(),
				Event:              workflowRun.GetEvent(),
				DisplayTitle:       workflowRun.GetDisplayTitle(),
				Status:             workflowRun.GetStatus(),
				Conclusion:         workflowRun.GetConclusion(),
				WorkflowID:         workflowRun.GetWorkflowID(),
				CheckSuiteID:       workflowRun.GetCheckSuiteID(),
				CheckSuiteNodeID:   workflowRun.GetCheckSuiteNodeID(),
				URL:                workflowRun.GetURL(),
				HTMLURL:            workflowRun.GetHTMLURL(),
				PullRequests:       workflowRun.PullRequests,
				CreatedAt:          workflowRun.GetCreatedAt(),
				UpdatedAt:          workflowRun.GetUpdatedAt(),
				RunStartedAt:       workflowRun.GetRunStartedAt(),
				JobsURL:            workflowRun.GetJobsURL(),
				LogsURL:            workflowRun.GetLogsURL(),
				CheckSuiteURL:      workflowRun.GetCheckSuiteURL(),
				ArtifactsURL:       workflowRun.GetArtifactsURL(),
				CancelURL:          workflowRun.GetCancelURL(),
				RerunURL:           workflowRun.GetRerunURL(),
				PreviousAttemptURL: workflowRun.GetPreviousAttemptURL(),
				HeadCommit:         workflowRun.GetHeadCommit(),
				WorkflowURL:        workflowRun.GetWorkflowURL(),
				Repository:         workflowRun.GetRepository(),
				HeadRepository:     workflowRun.GetHeadRepository(),
				Actor:              workflowRun.GetActor(),
				TriggeringActor:    workflowRun.GetTriggeringActor(),
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
