package describer

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"github.com/opengovern/og-describer-github/pkg/sdk/models"
	"github.com/opengovern/og-describer-github/provider/model"
	resilientbridge "github.com/opengovern/resilient-bridge"
	"github.com/opengovern/resilient-bridge/adapters"
)

// GetRepositoryList returns a list of all active (non-archived, non-disabled) repos in the organization.
// By default, no excludes are applied, so this returns only active repositories.
func GetRepositoryList(ctx context.Context, githubClient GitHubClient, organizationName string, stream *models.StreamSender) ([]models.Resource, error) {
	// Call the helper with default options (no excludes)
	return GetRepositoryListWithOptions(ctx, githubClient, organizationName, stream, false, false)
}

// GetRepositoryListWithOptions returns a list of all active repos in the organization with options to exclude archived or disabled.
func GetRepositoryListWithOptions(ctx context.Context, githubClient GitHubClient, organizationName string, stream *models.StreamSender, excludeArchived bool, excludeDisabled bool) ([]models.Resource, error) {
	maxResults := 100

	sdk := resilientbridge.NewResilientBridge()
	sdk.RegisterProvider("github", adapters.NewGitHubAdapter(githubClient.Token), &resilientbridge.ProviderConfig{
		UseProviderLimits: true,
		MaxRetries:        3,
		BaseBackoff:       0,
	})

	allRepos, err := _fetchOrgRepos(sdk, organizationName, maxResults)
	if err != nil {
		return nil, fmt.Errorf("error fetching organization repositories: %w", err)
	}

	// Filter repositories based on excludeArchived and excludeDisabled
	var filteredRepos []model.MinimalRepoInfo
	for _, r := range allRepos {
		if excludeArchived && r.Archived {
			continue
		}
		if excludeDisabled && r.Disabled {
			continue
		}
		filteredRepos = append(filteredRepos, r)
	}

	// Multi-threading (5 workers) for fetching repository details
	concurrency := 5
	results := make([]models.Resource, len(filteredRepos))

	type job struct {
		index int
		repo  string
	}

	jobCh := make(chan job)
	wg := sync.WaitGroup{}

	worker := func() {
		defer wg.Done()
		for j := range jobCh {
			value := _getRepositoriesDetail(ctx, sdk, organizationName, j.repo, stream)
			if value != nil {
				results[j.index] = *value
			}
		}
	}

	// Start workers
	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go worker()
	}

	// Send jobs
	for i, r := range filteredRepos {
		jobCh <- job{index: i, repo: r.Name}
	}
	close(jobCh)

	// Wait for all workers to finish
	wg.Wait()

	// Filter out empty results in case some fetches failed
	var finalResults []models.Resource
	for _, res := range results {
		if res.ID != "" {
			finalResults = append(finalResults, res)
		}
	}

	return finalResults, nil
}

// GetRepositoryDetails returns details for a given repo
func GetRepositoryDetails(ctx context.Context, githubClient GitHubClient, organizationName, repositoryName string) (*models.Resource, error) {
	sdk := resilientbridge.NewResilientBridge()
	sdk.RegisterProvider("github", adapters.NewGitHubAdapter(githubClient.Token), &resilientbridge.ProviderConfig{
		UseProviderLimits: true,
		MaxRetries:        3,
		BaseBackoff:       0,
	})

	repoDetail, err := _fetchRepoDetails(sdk, organizationName, repositoryName)
	if err != nil {
		return nil, fmt.Errorf("error fetching repository details for %s/%s: %w", organizationName, repositoryName, err)
	}

	finalDetail := _transformToFinalRepoDetail(repoDetail)

	langs, err := _fetchLanguages(sdk, organizationName, repositoryName)
	if err == nil {
		finalDetail.Languages = langs
	}

	err = _enrichRepoMetrics(sdk, organizationName, repositoryName, finalDetail)
	if err != nil {
		log.Printf("Error enriching repo metrics for %s/%s: %v", organizationName, repositoryName, err)
	}

	value := models.Resource{
		ID:   strconv.Itoa(finalDetail.GitHubRepoID),
		Name: finalDetail.Name,
		Description: JSONAllFieldsMarshaller{
			Value: finalDetail,
		},
	}
	return &value, nil
}

// Utility/helper functions (prefixed with underscore)

// _fetchLanguages fetches repository languages.
func _fetchLanguages(sdk *resilientbridge.ResilientBridge, owner, repo string) (map[string]int, error) {
	req := &resilientbridge.NormalizedRequest{
		Method:   "GET",
		Endpoint: fmt.Sprintf("/repos/%s/%s/languages", owner, repo),
		Headers:  map[string]string{"Accept": "application/vnd.github+json"},
	}

	resp, err := sdk.Request("github", req)
	if err != nil {
		return nil, fmt.Errorf("error fetching languages: %w", err)
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("HTTP error %d: %s", resp.StatusCode, string(resp.Data))
	}

	var langs map[string]int
	if err := json.Unmarshal(resp.Data, &langs); err != nil {
		return nil, fmt.Errorf("error decoding languages: %w", err)
	}
	return langs, nil
}

// _enrichRepoMetrics enriches repo metrics such as commits, issues, branches, etc.
func _enrichRepoMetrics(sdk *resilientbridge.ResilientBridge, owner, repoName string, finalDetail *model.RepositoryDescription) error {
	var dbObj map[string]string
	if finalDetail.DefaultBranchRef != nil {
		if err := json.Unmarshal(finalDetail.DefaultBranchRef, &dbObj); err != nil {
			return err
		}
	}
	defaultBranch := dbObj["name"]
	if defaultBranch == "" {
		defaultBranch = "main"
	}

	commitsCount, err := _countCommits(sdk, owner, repoName, defaultBranch)
	if err != nil {
		return fmt.Errorf("counting commits: %w", err)
	}
	finalDetail.Metrics.Commits = commitsCount

	issuesCount, err := _countIssues(sdk, owner, repoName)
	if err != nil {
		return fmt.Errorf("counting issues: %w", err)
	}
	finalDetail.Metrics.Issues = issuesCount

	branchesCount, err := _countBranches(sdk, owner, repoName)
	if err != nil {
		return fmt.Errorf("counting branches: %w", err)
	}
	finalDetail.Metrics.Branches = branchesCount

	prCount, err := _countPullRequests(sdk, owner, repoName)
	if err != nil {
		return fmt.Errorf("counting PRs: %w", err)
	}
	finalDetail.Metrics.PullRequests = prCount

	releasesCount, err := _countReleases(sdk, owner, repoName)
	if err != nil {
		return fmt.Errorf("counting releases: %w", err)
	}
	finalDetail.Metrics.Releases = releasesCount

	tagsCount, err := _countTags(sdk, owner, repoName)
	if err != nil {
		return fmt.Errorf("counting tags: %w", err)
	}
	finalDetail.Metrics.Tags = tagsCount

	return nil
}

func _countTags(sdk *resilientbridge.ResilientBridge, owner, repoName string) (int, error) {
	endpoint := fmt.Sprintf("/repos/%s/%s/tags?per_page=1", owner, repoName)
	return _countItemsFromEndpoint(sdk, endpoint)
}

func _countCommits(sdk *resilientbridge.ResilientBridge, owner, repoName, defaultBranch string) (int, error) {
	endpoint := fmt.Sprintf("/repos/%s/%s/commits?sha=%s&per_page=1", owner, repoName, defaultBranch)
	return _countItemsFromEndpoint(sdk, endpoint)
}

func _countIssues(sdk *resilientbridge.ResilientBridge, owner, repoName string) (int, error) {
	endpoint := fmt.Sprintf("/repos/%s/%s/issues?state=all&per_page=1", owner, repoName)
	return _countItemsFromEndpoint(sdk, endpoint)
}

func _countBranches(sdk *resilientbridge.ResilientBridge, owner, repoName string) (int, error) {
	endpoint := fmt.Sprintf("/repos/%s/%s/branches?per_page=1", owner, repoName)
	return _countItemsFromEndpoint(sdk, endpoint)
}

func _countPullRequests(sdk *resilientbridge.ResilientBridge, owner, repoName string) (int, error) {
	endpoint := fmt.Sprintf("/repos/%s/%s/pulls?state=all&per_page=1", owner, repoName)
	return _countItemsFromEndpoint(sdk, endpoint)
}

func _countReleases(sdk *resilientbridge.ResilientBridge, owner, repoName string) (int, error) {
	endpoint := fmt.Sprintf("/repos/%s/%s/releases?per_page=1", owner, repoName)
	return _countItemsFromEndpoint(sdk, endpoint)
}

func _countItemsFromEndpoint(sdk *resilientbridge.ResilientBridge, endpoint string) (int, error) {
	req := &resilientbridge.NormalizedRequest{
		Method:   "GET",
		Endpoint: endpoint,
		Headers:  map[string]string{"Accept": "application/vnd.github+json"},
	}

	resp, err := sdk.Request("github", req)
	if err != nil {
		return 0, fmt.Errorf("error fetching data: %w", err)
	}

	if resp.StatusCode == 409 {
		return 0, nil
	}

	if resp.StatusCode >= 400 {
		return 0, fmt.Errorf("HTTP error %d: %s", resp.StatusCode, string(resp.Data))
	}

	var linkHeader string
	for k, v := range resp.Headers {
		if strings.ToLower(k) == "link" {
			linkHeader = v
			break
		}
	}

	if linkHeader == "" {
		if len(resp.Data) > 2 {
			var items []interface{}
			if err := json.Unmarshal(resp.Data, &items); err != nil {
				return 1, nil
			}
			return len(items), nil
		}
		return 0, nil
	}

	lastPage, err := _parseLastPage(linkHeader)
	if err != nil {
		return 0, fmt.Errorf("could not parse last page: %w", err)
	}

	return lastPage, nil
}

func _parseLastPage(linkHeader string) (int, error) {
	re := regexp.MustCompile(`page=(\d+)>; rel="last"`)
	matches := re.FindStringSubmatch(linkHeader)
	if len(matches) < 2 {
		return 1, nil
	}
	var lastPage int
	_, err := fmt.Sscanf(matches[1], "%d", &lastPage)
	if err != nil {
		return 0, err
	}
	return lastPage, nil
}

func _fetchOrgRepos(sdk *resilientbridge.ResilientBridge, org string, maxResults int) ([]model.MinimalRepoInfo, error) {
	var allRepos []model.MinimalRepoInfo
	perPage := 100
	page := 1

	for len(allRepos) < maxResults {
		remaining := maxResults - len(allRepos)
		if remaining < perPage {
			perPage = remaining
		}

		endpoint := fmt.Sprintf("/orgs/%s/repos?per_page=%d&page=%d", org, perPage, page)
		listReq := &resilientbridge.NormalizedRequest{
			Method:   "GET",
			Endpoint: endpoint,
			Headers:  map[string]string{"Accept": "application/vnd.github+json"},
		}

		listResp, err := sdk.Request("github", listReq)
		if err != nil {
			return nil, fmt.Errorf("error fetching repos: %w", err)
		}

		if listResp.StatusCode >= 400 {
			return nil, fmt.Errorf("HTTP error %d: %s", listResp.StatusCode, string(listResp.Data))
		}

		var repos []model.MinimalRepoInfo
		if err := json.Unmarshal(listResp.Data, &repos); err != nil {
			return nil, fmt.Errorf("error decoding repos list: %w", err)
		}

		if len(repos) == 0 {
			break
		}

		allRepos = append(allRepos, repos...)
		if len(allRepos) >= maxResults {
			break
		}
		page++
	}
	if len(allRepos) > maxResults {
		allRepos = allRepos[:maxResults]
	}
	return allRepos, nil
}

func _fetchRepoDetails(sdk *resilientbridge.ResilientBridge, owner, repo string) (*model.RepoDetail, error) {
	req := &resilientbridge.NormalizedRequest{
		Method:   "GET",
		Endpoint: fmt.Sprintf("/repos/%s/%s", owner, repo),
		Headers:  map[string]string{"Accept": "application/vnd.github+json"},
	}
	resp, err := sdk.Request("github", req)
	if err != nil {
		return nil, fmt.Errorf("error fetching repo details: %w", err)
	}
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("HTTP error %d: %s", resp.StatusCode, string(resp.Data))
	}

	var detail model.RepoDetail
	if err := json.Unmarshal(resp.Data, &detail); err != nil {
		return nil, fmt.Errorf("error decoding repo details: %w", err)
	}
	return &detail, nil
}

func _transformToFinalRepoDetail(detail *model.RepoDetail) *model.RepositoryDescription {
	var parent *model.RepositoryDescription
	if detail.Parent != nil {
		parent = _transformToFinalRepoDetail(detail.Parent)
	}
	var source *model.RepositoryDescription
	if detail.Source != nil {
		source = _transformToFinalRepoDetail(detail.Source)
	}

	var finalOwner *model.Owner
	if detail.Owner != nil {
		finalOwner = &model.Owner{
			Login:   detail.Owner.Login,
			ID:      detail.Owner.ID,
			NodeID:  detail.Owner.NodeID,
			HTMLURL: detail.Owner.HTMLURL,
			Type:    detail.Owner.Type,
		}
	}

	var finalOrg *model.Organization
	if detail.Organization != nil {
		finalOrg = &model.Organization{
			Login:        detail.Organization.Login,
			ID:           detail.Organization.ID,
			NodeID:       detail.Organization.NodeID,
			HTMLURL:      detail.Organization.HTMLURL,
			Type:         detail.Organization.Type,
			UserViewType: detail.Organization.UserViewType,
			SiteAdmin:    detail.Organization.SiteAdmin,
		}
	}

	dbObj := map[string]string{"name": detail.DefaultBranch}
	dbBytes, _ := json.Marshal(dbObj)

	isActive := !(detail.Archived || detail.Disabled)
	isEmpty := (detail.Size == 0)

	var licenseJSON json.RawMessage
	if detail.License != nil {
		lj, _ := json.Marshal(detail.License)
		licenseJSON = lj
	}

	finalDetail := &model.RepositoryDescription{
		GitHubRepoID:            detail.ID,
		NodeID:                  detail.NodeID,
		Name:                    detail.Name,
		NameWithOwner:           detail.FullName,
		Description:             detail.Description,
		CreatedAt:               detail.CreatedAt,
		UpdatedAt:               detail.UpdatedAt,
		PushedAt:                detail.PushedAt,
		IsActive:                isActive,
		IsEmpty:                 isEmpty,
		IsFork:                  detail.Fork,
		IsSecurityPolicyEnabled: false,
		Owner:                   finalOwner,
		HomepageURL:             detail.Homepage,
		LicenseInfo:             licenseJSON,
		Topics:                  detail.Topics,
		Visibility:              detail.Visibility,
		DefaultBranchRef:        dbBytes,
		Permissions:             detail.Permissions,
		Organization:            finalOrg,
		Parent:                  parent,
		Source:                  source,
		Languages:               nil,
		RepositorySettings: model.RepositorySettings{
			HasDiscussionsEnabled:     detail.HasDiscussions,
			HasIssuesEnabled:          detail.HasIssues,
			HasProjectsEnabled:        detail.HasProjects,
			HasWikiEnabled:            detail.HasWiki,
			MergeCommitAllowed:        detail.AllowMergeCommit,
			MergeCommitMessage:        detail.MergeCommitMessage,
			MergeCommitTitle:          detail.MergeCommitTitle,
			SquashMergeAllowed:        detail.AllowSquashMerge,
			SquashMergeCommitMessage:  detail.SquashMergeCommitMessage,
			SquashMergeCommitTitle:    detail.SquashMergeCommitTitle,
			HasDownloads:              detail.HasDownloads,
			HasPages:                  detail.HasPages,
			WebCommitSignoffRequired:  detail.WebCommitSignoffRequired,
			MirrorURL:                 detail.MirrorURL,
			AllowAutoMerge:            detail.AllowAutoMerge,
			DeleteBranchOnMerge:       detail.DeleteBranchOnMerge,
			AllowUpdateBranch:         detail.AllowUpdateBranch,
			UseSquashPRTitleAsDefault: detail.UseSquashPRTitleAsDefault,
			CustomProperties:          detail.CustomProperties,
			ForkingAllowed:            detail.AllowForking,
			IsTemplate:                detail.IsTemplate,
			AllowRebaseMerge:          detail.AllowRebaseMerge,
			Archived:                  detail.Archived,
			Disabled:                  detail.Disabled,
			Locked:                    detail.Locked,
		},
		SecuritySettings: model.SecuritySettings{
			VulnerabilityAlertsEnabled:               false,
			SecretScanningEnabled:                    false,
			SecretScanningPushProtectionEnabled:      false,
			DependabotSecurityUpdatesEnabled:         false,
			SecretScanningNonProviderPatternsEnabled: false,
			SecretScanningValidityChecksEnabled:      false,
		},
		RepoURLs: model.RepoURLs{
			GitURL:   detail.GitURL,
			SSHURL:   detail.SSHURL,
			CloneURL: detail.CloneURL,
			SVNURL:   detail.SVNURL,
			HTMLURL:  detail.HTMLURL,
		},
		Metrics: model.Metrics{
			Stargazers:  detail.StargazersCount,
			Forks:       detail.ForksCount,
			Subscribers: detail.SubscribersCount,
			Size:        detail.Size,
			OpenIssues:  detail.OpenIssuesCount,
		},
	}
	return finalDetail
}

func _getRepositoriesDetail(ctx context.Context, sdk *resilientbridge.ResilientBridge, organizationName, repo string, stream *models.StreamSender) *models.Resource {
	repoDetail, err := _fetchRepoDetails(sdk, organizationName, repo)
	if err != nil {
		log.Printf("Error fetching details for %s/%s: %v", organizationName, repo, err)
		return nil
	}

	finalDetail := _transformToFinalRepoDetail(repoDetail)
	langs, err := _fetchLanguages(sdk, organizationName, repo)
	if err == nil {
		finalDetail.Languages = langs
	}

	err = _enrichRepoMetrics(sdk, organizationName, repo, finalDetail)
	if err != nil {
		log.Printf("Error enriching repo metrics for %s/%s: %v", organizationName, repo, err)
	}

	value := models.Resource{
		ID:   strconv.Itoa(finalDetail.GitHubRepoID),
		Name: finalDetail.Name,
		Description: JSONAllFieldsMarshaller{
			Value: finalDetail,
		},
	}
	if stream != nil {
		if err := (*stream)(value); err != nil {
			return nil
		}
	}
	return &value
}
