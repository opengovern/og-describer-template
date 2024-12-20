package describer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/opengovern/og-describer-github/pkg/sdk/models"
	"github.com/opengovern/og-describer-github/provider/model"
	resilientbridge "github.com/opengovern/resilient-bridge"
	"github.com/opengovern/resilient-bridge/adapters"
	"log"
	"net/url"
	"strconv"
	"strings"
)

func GetAllWorkflowRuns(ctx context.Context, githubClient GitHubClient, organizationName string, stream *models.StreamSender) ([]models.Resource, error) {
	client := githubClient.RestClient
	owner := organizationName
	repositories, err := getRepositories(ctx, client, owner)
	if err != nil {
		return nil, nil
	}

	sdk := resilientbridge.NewResilientBridge()
	sdk.RegisterProvider("github", adapters.NewGitHubAdapter(githubClient.Token), &resilientbridge.ProviderConfig{
		UseProviderLimits: true,
		MaxRetries:        3,
		BaseBackoff:       0,
	})

	var values []models.Resource
	for _, repo := range repositories {
		active, err := checkRepositoryActive(sdk, owner, repo.GetName())
		if err != nil {
			log.Fatalf("Error checking repository: %v", err)
		}

		if !active {
			// Repository is archived or disabled, return 0 workflow runs
			// No output needed, just exit gracefully.
			continue
		}
		repoValues, err := GetRepositoryWorkflowRuns(ctx, sdk, stream, owner, repo.GetName())
		if err != nil {
			return nil, err
		}
		values = append(values, repoValues...)
	}
	return values, nil
}

func GetRepositoryWorkflowRuns(ctx context.Context, sdk *resilientbridge.ResilientBridge, stream *models.StreamSender, owner, repo string) ([]models.Resource, error) {
	maxRuns := 50

	runs, err := fetchWorkflowRuns(sdk, owner, repo, "", maxRuns)
	if err != nil {
		log.Fatalf("Error fetching workflow runs: %v", err)
	}

	var values []models.Resource
	for _, runBasic := range runs {
		runDetail, err := fetchRunDetails(sdk, owner, repo, runBasic.ID)
		if err != nil {
			log.Printf("Error fetching details for run %d: %v", runBasic.ID, err)
			continue
		}

		artifactCount, artifacts, err := fetchArtifactsForRun(sdk, owner, repo, runBasic.ID)
		if err != nil {
			log.Printf("Error fetching artifacts for run %d: %v", runBasic.ID, err)
			continue
		}
		runDetail.ArtifactCount = artifactCount
		runDetail.Artifacts = artifacts

		value := models.Resource{
			ID:   strconv.Itoa(runDetail.ID),
			Name: runDetail.Name,
			Description: JSONAllFieldsMarshaller{
				Value: runDetail,
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
	return values, nil
}

//func GetRepoWorkflowRun(ctx context.Context, client GitHubClient, organizationName string, repo string, resourceID string, stream *models.StreamSender) (*models.Resource, error) {
//	owner := organizationName
//	workflowRunID, err := strconv.ParseInt(resourceID, 10, 64)
//	if err != nil {
//		return nil, err
//	}
//	if workflowRunID == 0 || repo == "" {
//		return nil, nil
//	}
//	workflowRun, _, err := client.RestClient.Actions.GetWorkflowRunByID(ctx, owner, repo, workflowRunID)
//	if err != nil {
//		return nil, err
//	}
//	repoFullName := formRepositoryFullName(owner, repo)
//	value := models.Resource{
//		ID:   strconv.Itoa(int(*workflowRun.ID)),
//		Name: *workflowRun.Name,
//		Description: JSONAllFieldsMarshaller{
//			Value: model.WorkflowRunDescription{
//				ID:                 workflowRun.GetID(),
//				Name:               workflowRun.GetName(),
//				NodeID:             workflowRun.GetNodeID(),
//				HeadBranch:         workflowRun.GetHeadBranch(),
//				HeadSHA:            workflowRun.GetHeadSHA(),
//				RunNumber:          workflowRun.GetRunNumber(),
//				RunAttempt:         workflowRun.GetRunAttempt(),
//				Event:              workflowRun.GetEvent(),
//				DisplayTitle:       workflowRun.GetDisplayTitle(),
//				Status:             workflowRun.GetStatus(),
//				Conclusion:         workflowRun.GetConclusion(),
//				WorkflowID:         workflowRun.GetWorkflowID(),
//				CheckSuiteID:       workflowRun.GetCheckSuiteID(),
//				CheckSuiteNodeID:   workflowRun.GetCheckSuiteNodeID(),
//				URL:                workflowRun.GetURL(),
//				HTMLURL:            workflowRun.GetHTMLURL(),
//				PullRequests:       workflowRun.PullRequests,
//				CreatedAt:          workflowRun.GetCreatedAt(),
//				UpdatedAt:          workflowRun.GetUpdatedAt(),
//				RunStartedAt:       workflowRun.GetRunStartedAt(),
//				JobsURL:            workflowRun.GetJobsURL(),
//				LogsURL:            workflowRun.GetLogsURL(),
//				CheckSuiteURL:      workflowRun.GetCheckSuiteURL(),
//				ArtifactsURL:       workflowRun.GetArtifactsURL(),
//				CancelURL:          workflowRun.GetCancelURL(),
//				RerunURL:           workflowRun.GetRerunURL(),
//				PreviousAttemptURL: workflowRun.GetPreviousAttemptURL(),
//				HeadCommit:         workflowRun.GetHeadCommit(),
//				WorkflowURL:        workflowRun.GetWorkflowURL(),
//				Repository:         workflowRun.GetRepository(),
//				HeadRepository:     workflowRun.GetHeadRepository(),
//				Actor:              workflowRun.GetActor(),
//				TriggeringActor:    workflowRun.GetTriggeringActor(),
//				RepoFullName:       repoFullName,
//			},
//		},
//	}
//	if stream != nil {
//		if err := (*stream)(value); err != nil {
//			return nil, err
//		}
//	}
//	return &value, nil
//}

// parseRunNumberFlag parses the run_number flag.
// It handles:
// - Single run number: "23"
// - Comma-separated: "23,25"
// - Range: "23-56"
//
// The result is returned as a slice of runNumberCriterion, which can represent either single values or ranges.
type runNumberCriterion struct {
	From int
	To   int
}

func parseRunNumberFlag(flagVal string) []runNumberCriterion {
	if strings.TrimSpace(flagVal) == "" {
		return nil
	}

	parts := strings.Split(flagVal, ",")
	var criteria []runNumberCriterion
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		if strings.Contains(p, "-") {
			rangeParts := strings.SplitN(p, "-", 2)
			if len(rangeParts) == 2 {
				startStr := strings.TrimSpace(rangeParts[0])
				endStr := strings.TrimSpace(rangeParts[1])
				start, err1 := strconv.Atoi(startStr)
				end, err2 := strconv.Atoi(endStr)
				if err1 == nil && err2 == nil && start <= end {
					criteria = append(criteria, runNumberCriterion{From: start, To: end})
				}
			}
		} else {
			// Single number
			n, err := strconv.Atoi(p)
			if err == nil {
				criteria = append(criteria, runNumberCriterion{From: n, To: n})
			}
		}
	}
	return criteria
}

// filterRunsByNumber filters the given runs to include only those that match the runNumberCriterion(s).
func filterRunsByNumber(runs []model.WorkflowRunDescription, criteria []runNumberCriterion) []model.WorkflowRunDescription {
	var filtered []model.WorkflowRunDescription

	for _, run := range runs {
		if runNumberMatches(run.RunNumber, criteria) {
			filtered = append(filtered, run)
		}
	}
	return filtered
}

func runNumberMatches(runNum int, criteria []runNumberCriterion) bool {
	for _, c := range criteria {
		if runNum >= c.From && runNum <= c.To {
			return true
		}
	}
	return false
}

// fetchWorkflowRuns returns up to maxRuns workflow runs. If branch is specified, filter by that branch, otherwise fetch all.
func fetchWorkflowRuns(sdk *resilientbridge.ResilientBridge, owner, repo, branch string, maxRuns int) ([]model.WorkflowRunDescription, error) {
	var allRuns []model.WorkflowRunDescription
	perPage := 100
	page := 1

	for len(allRuns) < maxRuns {
		remaining := maxRuns - len(allRuns)
		if remaining < perPage {
			perPage = remaining
		}

		params := url.Values{}
		params.Set("per_page", fmt.Sprintf("%d", perPage))
		params.Set("page", fmt.Sprintf("%d", page))

		// If branch is provided, add it to the query params
		if branch != "" {
			params.Set("branch", branch)
		}

		endpoint := fmt.Sprintf("/repos/%s/%s/actions/runs?%s", owner, repo, params.Encode())

		req := &resilientbridge.NormalizedRequest{
			Method:   "GET",
			Endpoint: endpoint,
			Headers:  map[string]string{"Accept": "application/vnd.github+json"},
		}

		resp, err := sdk.Request("github", req)
		if err != nil {
			return nil, fmt.Errorf("error fetching workflow runs: %w", err)
		}

		if resp.StatusCode >= 400 {
			return nil, fmt.Errorf("HTTP error %d: %s", resp.StatusCode, string(resp.Data))
		}

		var runsResp model.WorkflowRunsResponse
		if err := json.Unmarshal(resp.Data, &runsResp); err != nil {
			return nil, fmt.Errorf("error decoding workflow runs: %w", err)
		}

		if len(runsResp.WorkflowRuns) == 0 {
			// No more runs
			break
		}

		allRuns = append(allRuns, runsResp.WorkflowRuns...)
		if len(allRuns) >= maxRuns {
			break
		}
		page++
	}

	if len(allRuns) > maxRuns {
		allRuns = allRuns[:maxRuns]
	}

	return allRuns, nil
}

// fetchRunDetails fetches the full details for a specific run, including actor, repository info, and referenced_workflows.
func fetchRunDetails(sdk *resilientbridge.ResilientBridge, owner, repo string, runID int) (model.WorkflowRunDescription, error) {
	endpoint := fmt.Sprintf("/repos/%s/%s/actions/runs/%d", owner, repo, runID)
	req := &resilientbridge.NormalizedRequest{
		Method:   "GET",
		Endpoint: endpoint,
		Headers:  map[string]string{"Accept": "application/vnd.github+json"},
	}

	resp, err := sdk.Request("github", req)
	if err != nil {
		return model.WorkflowRunDescription{}, fmt.Errorf("error fetching run details: %w", err)
	}

	if resp.StatusCode >= 400 {
		return model.WorkflowRunDescription{}, fmt.Errorf("HTTP error %d: %s", resp.StatusCode, string(resp.Data))
	}

	var fullDetail struct {
		ID                  int                      `json:"id"`
		Name                string                   `json:"name"`
		HeadBranch          string                   `json:"head_branch"`
		HeadSHA             string                   `json:"head_sha"`
		Status              string                   `json:"status"`
		Conclusion          string                   `json:"conclusion"`
		HTMLURL             string                   `json:"html_url"`
		WorkflowID          int                      `json:"workflow_id"`
		RunNumber           int                      `json:"run_number"`
		Event               string                   `json:"event"`
		CreatedAt           string                   `json:"created_at"`
		UpdatedAt           string                   `json:"updated_at"`
		RunAttempt          int                      `json:"run_attempt"`
		RunStartedAt        string                   `json:"run_started_at"`
		Actor               *model.SimpleActor       `json:"actor"`
		HeadCommit          *model.CommitRefWorkflow `json:"head_commit"`
		Repository          *model.SimpleRepo        `json:"repository"`
		HeadRepository      *model.SimpleRepo        `json:"head_repository"`
		ReferencedWorkflows []interface{}            `json:"referenced_workflows"`
	}

	if err := json.Unmarshal(resp.Data, &fullDetail); err != nil {
		return model.WorkflowRunDescription{}, fmt.Errorf("error decoding run details: %w", err)
	}

	return model.WorkflowRunDescription{
		ID:                  fullDetail.ID,
		Name:                fullDetail.Name,
		HeadBranch:          fullDetail.HeadBranch,
		HeadSHA:             fullDetail.HeadSHA,
		Status:              fullDetail.Status,
		Conclusion:          fullDetail.Conclusion,
		HTMLURL:             fullDetail.HTMLURL,
		WorkflowID:          fullDetail.WorkflowID,
		RunNumber:           fullDetail.RunNumber,
		Event:               fullDetail.Event,
		CreatedAt:           fullDetail.CreatedAt,
		UpdatedAt:           fullDetail.UpdatedAt,
		RunAttempt:          fullDetail.RunAttempt,
		RunStartedAt:        fullDetail.RunStartedAt,
		Actor:               fullDetail.Actor,
		HeadCommit:          fullDetail.HeadCommit,
		Repository:          fullDetail.Repository,
		HeadRepository:      fullDetail.HeadRepository,
		ReferencedWorkflows: fullDetail.ReferencedWorkflows,
	}, nil
}

func fetchArtifactsForRun(sdk *resilientbridge.ResilientBridge, owner, repo string, runID int) (int, []model.WorkflowArtifact, error) {
	endpoint := fmt.Sprintf("/repos/%s/%s/actions/runs/%d/artifacts", owner, repo, runID)
	req := &resilientbridge.NormalizedRequest{
		Method:   "GET",
		Endpoint: endpoint,
		Headers:  map[string]string{"Accept": "application/vnd.github+json"},
	}

	resp, err := sdk.Request("github", req)
	if err != nil {
		return 0, nil, fmt.Errorf("error fetching artifacts: %w", err)
	}

	if resp.StatusCode >= 400 {
		return 0, nil, fmt.Errorf("HTTP error %d: %s", resp.StatusCode, string(resp.Data))
	}

	var artResp model.ArtifactsResponse
	if err := json.Unmarshal(resp.Data, &artResp); err != nil {
		return 0, nil, fmt.Errorf("error decoding artifacts response: %w", err)
	}

	return artResp.TotalCount, artResp.Artifacts, nil
}
