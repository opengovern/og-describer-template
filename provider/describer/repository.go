package describer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/go-github/v55/github"
	"github.com/opengovern/og-describer-github/pkg/sdk/models"
	"github.com/opengovern/og-describer-github/provider/model"
	resilientbridge "github.com/opengovern/resilient-bridge"
	"github.com/opengovern/resilient-bridge/adapters"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func GetRepositoryList(ctx context.Context, githubClient GitHubClient, organizationName string, stream *models.StreamSender) ([]models.Resource, error) {
	maxResults := 100

	sdk := resilientbridge.NewResilientBridge()
	sdk.RegisterProvider("github", adapters.NewGitHubAdapter(githubClient.Token), &resilientbridge.ProviderConfig{
		UseProviderLimits: true,
		MaxRetries:        3,
		BaseBackoff:       0,
	})

	allRepos, err := fetchOrgRepos(sdk, organizationName, maxResults)
	if err != nil {
		log.Fatalf("Error fetching organization repositories: %v", err)
	}

	var values []models.Resource
	for _, r := range allRepos {
		value := getRepositoriesDetail(ctx, sdk, organizationName, r.Name, stream)
		values = append(values, *value)
	}

	return values, nil
}

func getRepositoriesDetail(ctx context.Context, sdk *resilientbridge.ResilientBridge, organizationName, repo string, stream *models.StreamSender) *models.Resource {
	repoDetail, err := fetchRepoDetails(sdk, organizationName, repo)
	if err != nil {
		log.Printf("Error fetching details for %s/%s: %v", organizationName, repo, err)
		return nil
	}

	finalDetail := transformToFinalRepoDetail(repoDetail)
	// Fetch languages
	langs, err := fetchLanguages(sdk, organizationName, repo)
	if err == nil {
		finalDetail.Languages = langs
	}

	// Enrich metrics
	err = enrichRepoMetrics(sdk, organizationName, repo, finalDetail)
	if err != nil {
		log.Printf("Error enriching repo metrics for %s/%s: %v", organizationName, repo, err)
	}

	// **New addition: Fetch private vulnerability reporting status**
	pvrEnabled, err := fetchPrivateVulnerabilityReporting(sdk, organizationName, repo)
	if err != nil {
		log.Printf("Error fetching private vulnerability reporting status for %s/%s: %v", organizationName, repo, err)
	} else {
		finalDetail.SecuritySettings.PrivateVulnerabilityReportingEnabled = pvrEnabled
	}

	value := models.Resource{
		ID:   strconv.Itoa(finalDetail.GitHubRepoID),
		Name: finalDetail.Name,
		Description: JSONAllFieldsMarshaller{
			Value: model.RepositoryDescription{
				GitHubRepoID:            finalDetail.GitHubRepoID,
				NodeID:                  finalDetail.NodeID,
				Name:                    finalDetail.Name,
				NameWithOwner:           finalDetail.NameWithOwner,
				Description:             finalDetail.Description,
				CreatedAt:               finalDetail.CreatedAt,
				UpdatedAt:               finalDetail.UpdatedAt,
				PushedAt:                finalDetail.PushedAt,
				IsActive:                finalDetail.IsActive,
				IsEmpty:                 finalDetail.IsEmpty,
				IsFork:                  finalDetail.IsFork,
				IsSecurityPolicyEnabled: finalDetail.IsSecurityPolicyEnabled,
				Owner:                   finalDetail.Owner,
				HomepageURL:             finalDetail.HomepageURL,
				LicenseInfo:             finalDetail.LicenseInfo,
				Topics:                  finalDetail.Topics,
				Visibility:              finalDetail.Visibility,
				DefaultBranchRef:        finalDetail.DefaultBranchRef,
				Permissions:             finalDetail.Permissions,
				Organization:            finalDetail.Organization,
				Parent:                  finalDetail.Parent,
				Source:                  finalDetail.Source,
				Languages:               finalDetail.Languages,
				RepositorySettings:      finalDetail.RepositorySettings,
				SecuritySettings:        finalDetail.SecuritySettings,
				RepoURLs:                finalDetail.RepoURLs,
				Metrics:                 finalDetail.Metrics,
			},
		},
	}
	if stream != nil {
		if err := (*stream)(value); err != nil {
			return nil
		}
	}
	return &value
}

//func GetRepositoryList(ctx context.Context, githubClient GitHubClient, organizationName string, stream *models.StreamSender) ([]models.Resource, error) {
//	client := githubClient.GraphQLClient
//	query := struct {
//		RateLimit    steampipemodels.RateLimit
//		Organization struct {
//			Repositories struct {
//				PageInfo   steampipemodels.PageInfo
//				TotalCount int
//				Nodes      []steampipemodels.Repository
//			} `graphql:"repositories(first: $pageSize, after: $cursor)"`
//		} `graphql:"organization(login: $owner)"` // <-- $owner used here
//	}{}
//	variables := map[string]interface{}{
//		"owner":    githubv4.String(organizationName),
//		"pageSize": githubv4.Int(repoPageSize),
//		"cursor":   (*githubv4.String)(nil),
//	}
//	columnNames := repositoryCols()
//	appendRepoColumnIncludes(&variables, columnNames)
//	var values []models.Resource
//	for {
//		err := client.Query(ctx, &query, variables)
//		if err != nil {
//			return nil, err
//		}
//		for _, repo := range query.Organization.Repositories.Nodes {
//			hooks, err := GetRepositoryHooks(ctx, githubClient.RestClient, organizationName, repo.Name)
//			if err != nil {
//				return nil, err
//			}
//			additionalRepoInfo, err := GetRepositoryAdditionalData(ctx, githubClient.RestClient, organizationName, repo.Name)
//			value := models.Resource{
//				ID:   strconv.Itoa(repo.Id),
//				Name: repo.Name,
//				Description: JSONAllFieldsMarshaller{
//					Value: model.RepositoryDescription{
//						ID:                            repo.Id,
//						NodeID:                        repo.NodeId,
//						Name:                          repo.Name,
//						AllowUpdateBranch:             repo.AllowUpdateBranch,
//						ArchivedAt:                    repo.ArchivedAt,
//						AutoMergeAllowed:              repo.AutoMergeAllowed,
//						CodeOfConduct:                 repo.CodeOfConduct,
//						ContactLinks:                  repo.ContactLinks,
//						CreatedAt:                     repo.CreatedAt,
//						DefaultBranchRef:              repo.DefaultBranchRef,
//						DeleteBranchOnMerge:           repo.DeleteBranchOnMerge,
//						Description:                   repo.Description,
//						DiskUsage:                     repo.DiskUsage,
//						ForkCount:                     repo.ForkCount,
//						ForkingAllowed:                repo.ForkingAllowed,
//						FundingLinks:                  repo.FundingLinks,
//						HasDiscussionsEnabled:         repo.HasDiscussionsEnabled,
//						HasIssuesEnabled:              repo.HasIssuesEnabled,
//						HasProjectsEnabled:            repo.HasProjectsEnabled,
//						HasVulnerabilityAlertsEnabled: repo.HasVulnerabilityAlertsEnabled,
//						HasWikiEnabled:                repo.HasWikiEnabled,
//						HomepageURL:                   repo.HomepageUrl,
//						InteractionAbility:            repo.InteractionAbility,
//						IsArchived:                    repo.IsArchived,
//						IsBlankIssuesEnabled:          repo.IsBlankIssuesEnabled,
//						IsDisabled:                    repo.IsDisabled,
//						IsEmpty:                       repo.IsEmpty,
//						IsFork:                        repo.IsFork,
//						IsInOrganization:              repo.IsInOrganization,
//						IsLocked:                      repo.IsLocked,
//						IsMirror:                      repo.IsMirror,
//						IsPrivate:                     repo.IsPrivate,
//						IsSecurityPolicyEnabled:       repo.IsSecurityPolicyEnabled,
//						IsTemplate:                    repo.IsTemplate,
//						IsUserConfigurationRepository: repo.IsUserConfigurationRepository,
//						IssueTemplates:                repo.IssueTemplates,
//						LicenseInfo:                   repo.LicenseInfo,
//						LockReason:                    repo.LockReason,
//						MergeCommitAllowed:            repo.MergeCommitAllowed,
//						MergeCommitMessage:            repo.MergeCommitMessage,
//						MergeCommitTitle:              repo.MergeCommitTitle,
//						MirrorURL:                     repo.MirrorUrl,
//						NameWithOwner:                 repo.NameWithOwner,
//						OpenGraphImageURL:             repo.OpenGraphImageUrl,
//						OwnerLogin:                    repo.Owner.Login,
//						PrimaryLanguage:               repo.PrimaryLanguage,
//						ProjectsURL:                   repo.ProjectsUrl,
//						PullRequestTemplates:          repo.PullRequestTemplates,
//						PushedAt:                      repo.PushedAt,
//						RebaseMergeAllowed:            repo.RebaseMergeAllowed,
//						SecurityPolicyURL:             repo.SecurityPolicyUrl,
//						SquashMergeAllowed:            repo.SquashMergeAllowed,
//						SquashMergeCommitMessage:      repo.SquashMergeCommitMessage,
//						SquashMergeCommitTitle:        repo.SquashMergeCommitTitle,
//						SSHURL:                        repo.SshUrl,
//						StargazerCount:                repo.StargazerCount,
//						UpdatedAt:                     repo.UpdatedAt,
//						URL:                           repo.Url,
//						// UsesCustomOpenGraphImage:      repo.UsesCustomOpenGraphImage,
//						// CanAdminister:                 repo.CanAdminister,
//						// CanCreateProjects:             repo.CanCreateProjects,
//						// CanSubscribe:                  repo.CanSubscribe,
//						// CanUpdateTopics:               repo.CanUpdateTopics,
//						// HasStarred:                    repo.HasStarred,
//						PossibleCommitEmails: repo.PossibleCommitEmails,
//						// Subscription:                  repo.Subscription,
//						Visibility: repo.Visibility,
//						// YourPermission:                repo.YourPermission,
//						WebCommitSignOffRequired:   repo.WebCommitSignoffRequired,
//						RepositoryTopicsTotalCount: repo.RepositoryTopics.TotalCount,
//						OpenIssuesTotalCount:       repo.OpenIssues.TotalCount,
//						WatchersTotalCount:         repo.Watchers.TotalCount,
//						Hooks:                      hooks,
//						Topics:                     additionalRepoInfo.Topics,
//						SubscribersCount:           additionalRepoInfo.GetSubscribersCount(),
//						HasDownloads:               additionalRepoInfo.GetHasDownloads(),
//						HasPages:                   additionalRepoInfo.GetHasPages(),
//						NetworkCount:               additionalRepoInfo.GetNetworkCount(),
//					},
//				},
//			}
//			if stream != nil {
//				if err := (*stream)(value); err != nil {
//					return nil, err
//				}
//			} else {
//				values = append(values, value)
//			}
//		}
//		if !query.Organization.Repositories.PageInfo.HasNextPage {
//			break
//		}
//		variables["cursor"] = githubv4.NewString(query.Organization.Repositories.PageInfo.EndCursor)
//	}
//	return values, nil
//}

//func GetRepository(ctx context.Context, githubClient GitHubClient, organizationName string, repositoryName string, resourceID string, stream *models.StreamSender) (*models.Resource, error) {
//	client := githubClient.GraphQLClient
//	query := struct {
//		RateLimit    steampipemodels.RateLimit
//		Organization struct {
//			Repository steampipemodels.Repository `graphql:"repository(name: $repoName)"`
//		} `graphql:"organization(login: $owner)"` // <-- $owner used here
//	}{}
//
//	variables := map[string]interface{}{
//		"owner":    githubv4.String(organizationName),
//		"repoName": githubv4.String(repositoryName),
//	}
//
//	columnNames := repositoryCols()
//	appendRepoColumnIncludes(&variables, columnNames)
//	err := client.Query(ctx, &query, variables)
//	if err != nil {
//		return nil, err
//	}
//	repo := query.Organization.Repository
//	hooks, err := GetRepositoryHooks(ctx, githubClient.RestClient, organizationName, repo.Name)
//	if err != nil {
//		return nil, err
//	}
//	additionalRepoInfo, err := GetRepositoryAdditionalData(ctx, githubClient.RestClient, organizationName, repo.Name)
//	value := models.Resource{
//		ID:   strconv.Itoa(repo.Id),
//		Name: repo.Name,
//		Description: JSONAllFieldsMarshaller{
//			Value: model.RepositoryDescription{
//				ID:                            repo.Id,
//				NodeID:                        repo.NodeId,
//				Name:                          repo.Name,
//				AllowUpdateBranch:             repo.AllowUpdateBranch,
//				ArchivedAt:                    repo.ArchivedAt,
//				AutoMergeAllowed:              repo.AutoMergeAllowed,
//				CodeOfConduct:                 repo.CodeOfConduct,
//				ContactLinks:                  repo.ContactLinks,
//				CreatedAt:                     repo.CreatedAt,
//				DefaultBranchRef:              repo.DefaultBranchRef,
//				DeleteBranchOnMerge:           repo.DeleteBranchOnMerge,
//				Description:                   repo.Description,
//				DiskUsage:                     repo.DiskUsage,
//				ForkCount:                     repo.ForkCount,
//				ForkingAllowed:                repo.ForkingAllowed,
//				FundingLinks:                  repo.FundingLinks,
//				HasDiscussionsEnabled:         repo.HasDiscussionsEnabled,
//				HasIssuesEnabled:              repo.HasIssuesEnabled,
//				HasProjectsEnabled:            repo.HasProjectsEnabled,
//				HasVulnerabilityAlertsEnabled: repo.HasVulnerabilityAlertsEnabled,
//				HasWikiEnabled:                repo.HasWikiEnabled,
//				HomepageURL:                   repo.HomepageUrl,
//				InteractionAbility:            repo.InteractionAbility,
//				IsArchived:                    repo.IsArchived,
//				IsBlankIssuesEnabled:          repo.IsBlankIssuesEnabled,
//				IsDisabled:                    repo.IsDisabled,
//				IsEmpty:                       repo.IsEmpty,
//				IsFork:                        repo.IsFork,
//				IsInOrganization:              repo.IsInOrganization,
//				IsLocked:                      repo.IsLocked,
//				IsMirror:                      repo.IsMirror,
//				IsPrivate:                     repo.IsPrivate,
//				IsSecurityPolicyEnabled:       repo.IsSecurityPolicyEnabled,
//				IsTemplate:                    repo.IsTemplate,
//				IsUserConfigurationRepository: repo.IsUserConfigurationRepository,
//				IssueTemplates:                repo.IssueTemplates,
//				LicenseInfo:                   repo.LicenseInfo,
//				LockReason:                    repo.LockReason,
//				MergeCommitAllowed:            repo.MergeCommitAllowed,
//				MergeCommitMessage:            repo.MergeCommitMessage,
//				MergeCommitTitle:              repo.MergeCommitTitle,
//				MirrorURL:                     repo.MirrorUrl,
//				NameWithOwner:                 repo.NameWithOwner,
//				OpenGraphImageURL:             repo.OpenGraphImageUrl,
//				OwnerLogin:                    repo.Owner.Login,
//				PrimaryLanguage:               repo.PrimaryLanguage,
//				ProjectsURL:                   repo.ProjectsUrl,
//				PullRequestTemplates:          repo.PullRequestTemplates,
//				PushedAt:                      repo.PushedAt,
//				RebaseMergeAllowed:            repo.RebaseMergeAllowed,
//				SecurityPolicyURL:             repo.SecurityPolicyUrl,
//				SquashMergeAllowed:            repo.SquashMergeAllowed,
//				SquashMergeCommitMessage:      repo.SquashMergeCommitMessage,
//				SquashMergeCommitTitle:        repo.SquashMergeCommitTitle,
//				SSHURL:                        repo.SshUrl,
//				StargazerCount:                repo.StargazerCount,
//				UpdatedAt:                     repo.UpdatedAt,
//				URL:                           repo.Url,
//				// UsesCustomOpenGraphImage:      repo.UsesCustomOpenGraphImage,
//				// CanAdminister:                 repo.CanAdminister,
//				// CanCreateProjects:             repo.CanCreateProjects,
//				// CanSubscribe:                  repo.CanSubscribe,
//				// CanUpdateTopics:               repo.CanUpdateTopics,
//				// HasStarred:                    repo.HasStarred,
//				PossibleCommitEmails: repo.PossibleCommitEmails,
//				// Subscription:                  repo.Subscription,
//				Visibility: repo.Visibility,
//				// YourPermission:                repo.YourPermission,
//				WebCommitSignOffRequired:   repo.WebCommitSignoffRequired,
//				RepositoryTopicsTotalCount: repo.RepositoryTopics.TotalCount,
//				OpenIssuesTotalCount:       repo.OpenIssues.TotalCount,
//				WatchersTotalCount:         repo.Watchers.TotalCount,
//				Hooks:                      hooks,
//				Topics:                     additionalRepoInfo.Topics,
//				SubscribersCount:           additionalRepoInfo.GetSubscribersCount(),
//				HasDownloads:               additionalRepoInfo.GetHasDownloads(),
//				HasPages:                   additionalRepoInfo.GetHasPages(),
//				NetworkCount:               additionalRepoInfo.GetNetworkCount(),
//			},
//		},
//	}
//	if stream != nil {
//		if err := (*stream)(value); err != nil {
//			return nil, err
//		}
//	}
//
//	return &value, nil
//}

func GetRepositoryAdditionalData(ctx context.Context, client *github.Client, organizationName string, repo string) (*github.Repository, error) {
	repository, _, err := client.Repositories.Get(ctx, organizationName, repo)
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			return nil, nil
		}
		return nil, nil
	}
	if repository == nil {
		return nil, nil
	}
	return repository, nil
}

func GetRepositoryHooks(ctx context.Context, client *github.Client, organizationName string, repo string) ([]*github.Hook, error) {
	var repositoryHooks []*github.Hook
	opt := &github.ListOptions{PerPage: pageSize}
	for {
		hooks, resp, err := client.Repositories.ListHooks(ctx, organizationName, repo, opt)
		if err != nil && strings.Contains(err.Error(), "Not Found") {
			return nil, nil
		} else if err != nil {
			return nil, err
		}
		repositoryHooks = append(repositoryHooks, hooks...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}
	return repositoryHooks, nil
}

func enrichRepoMetrics(sdk *resilientbridge.ResilientBridge, owner, repoName string, finalDetail *model.RepositoryDescription) error {
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

	commitsCount, err := countCommits(sdk, owner, repoName, defaultBranch)
	if err != nil {
		return fmt.Errorf("counting commits: %w", err)
	}
	finalDetail.Metrics.Commits = commitsCount

	issuesCount, err := countIssues(sdk, owner, repoName)
	if err != nil {
		return fmt.Errorf("counting issues: %w", err)
	}
	finalDetail.Metrics.Issues = issuesCount

	branchesCount, err := countBranches(sdk, owner, repoName)
	if err != nil {
		return fmt.Errorf("counting branches: %w", err)
	}
	finalDetail.Metrics.Branches = branchesCount

	prCount, err := countPullRequests(sdk, owner, repoName)
	if err != nil {
		return fmt.Errorf("counting PRs: %w", err)
	}
	finalDetail.Metrics.PullRequests = prCount

	releasesCount, err := countReleases(sdk, owner, repoName)
	if err != nil {
		return fmt.Errorf("counting releases: %w", err)
	}
	finalDetail.Metrics.Releases = releasesCount

	// New: Count tags
	tagsCount, err := countTags(sdk, owner, repoName)
	if err != nil {
		return fmt.Errorf("counting tags: %w", err)
	}
	// Add "TotalTags" field to RepoMetrics struct and assign here
	// You need to add the `TotalTags int `json:"total_tags"` field to RepoMetrics beforehand
	finalDetail.Metrics.Tags = tagsCount

	return nil
}

// New function countTags
func countTags(sdk *resilientbridge.ResilientBridge, owner, repoName string) (int, error) {
	endpoint := fmt.Sprintf("/repos/%s/%s/tags?per_page=1", owner, repoName)
	return countItemsFromEndpoint(sdk, endpoint)
}

func fetchLanguages(sdk *resilientbridge.ResilientBridge, owner, repo string) (map[string]int, error) {
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

// The rest of the functions (parseScopeURL, fetchOrgRepos, fetchRepoDetails, transformToFinalRepoDetail,
// enrichRepoMetrics, countCommits, countIssues, countBranches, countPullRequests, countReleases,
// countItemsFromEndpoint, and parseLastPage) remain unchanged.

func parseScopeURL(repoURL string) (owner, repo string, err error) {
	if !strings.HasPrefix(repoURL, "https://github.com/") {
		return "", "", fmt.Errorf("URL must start with https://github.com/")
	}
	parts := strings.Split(strings.TrimPrefix(repoURL, "https://github.com/"), "/")
	if len(parts) == 0 || parts[0] == "" {
		return "", "", fmt.Errorf("invalid URL format")
	}
	owner = parts[0]
	if len(parts) > 1 {
		repo = parts[1]
	}
	return owner, repo, nil
}

func fetchPrivateVulnerabilityReporting(sdk *resilientbridge.ResilientBridge, owner, repoName string) (bool, error) {
	req := &resilientbridge.NormalizedRequest{
		Method:   "GET",
		Endpoint: fmt.Sprintf("/repos/%s/%s/private-vulnerability-reporting", owner, repoName),
		Headers:  map[string]string{"Accept": "application/vnd.github+json"},
	}

	resp, err := sdk.Request("github", req)
	if err != nil {
		return false, fmt.Errorf("error fetching private vulnerability reporting: %w", err)
	}

	if resp.StatusCode == 404 {
		// Endpoint returns 404 if private vulnerability reporting is not enabled
		// or the resource is not found. Default to false.
		return false, nil
	}

	if resp.StatusCode >= 400 {
		return false, fmt.Errorf("HTTP error %d: %s", resp.StatusCode, string(resp.Data))
	}

	var result struct {
		Enabled bool `json:"enabled"`
	}
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return false, fmt.Errorf("error decoding private vulnerability reporting status: %w", err)
	}

	return result.Enabled, nil
}

func fetchOrgRepos(sdk *resilientbridge.ResilientBridge, org string, maxResults int) ([]model.MinimalRepoInfo, error) {
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

func fetchRepoDetails(sdk *resilientbridge.ResilientBridge, owner, repo string) (*model.RepoDetail, error) {
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

func transformToFinalRepoDetail(detail *model.RepoDetail) *model.RepositoryDescription {
	var parent *model.RepositoryDescription
	if detail.Parent != nil {
		parent = transformToFinalRepoDetail(detail.Parent)
	}
	var source *model.RepositoryDescription
	if detail.Source != nil {
		source = transformToFinalRepoDetail(detail.Source)
	}

	// Owner: no user_view_type, no site_admin
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

	// Organization: includes user_view_type, site_admin
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

	// Prepare default_branch_ref as before
	dbObj := map[string]string{"name": detail.DefaultBranch}
	dbBytes, _ := json.Marshal(dbObj)

	// Determine is_active: true if not archived and not disabled
	isActive := !(detail.Archived || detail.Disabled)
	//isInOrganization := (detail.Organization != nil && detail.Organization.Type == "Organization")
	//isMirror := (detail.MirrorURL != nil)
	isEmpty := (detail.Size == 0)

	var licenseJSON json.RawMessage
	if detail.License != nil {
		lj, _ := json.Marshal(detail.License)
		licenseJSON = lj
	}

	finalDetail := &model.RepositoryDescription{
		GitHubRepoID:  detail.ID,
		NodeID:        detail.NodeID,
		Name:          detail.Name,
		NameWithOwner: detail.FullName,
		Description:   detail.Description,
		CreatedAt:     detail.CreatedAt,
		UpdatedAt:     detail.UpdatedAt,
		PushedAt:      detail.PushedAt,
		IsActive:      isActive,
		IsEmpty:       isEmpty,
		IsFork:        detail.Fork,
		//IsInOrganization:        isInOrganization,
		//IsMirror:                isMirror,
		IsSecurityPolicyEnabled: false, // as before
		//IsTemplate:              detail.IsTemplate,
		Owner:            finalOwner,
		HomepageURL:      detail.Homepage,
		LicenseInfo:      licenseJSON,
		Topics:           detail.Topics,
		Visibility:       detail.Visibility,
		DefaultBranchRef: dbBytes,
		Permissions:      detail.Permissions,
		Organization:     finalOrg,
		Parent:           parent,
		Source:           source,
		Languages:        nil, // set after fetchLanguages
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
			PrivateVulnerabilityReportingEnabled:     false,
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
			// The rest (tags, commits, issues, open_issues, branches, pull_requests, releases)
			// will be set after calling enrichRepoMetrics and assigning open issues from detail.
		},
	}

	// Set open_issues before enrichRepoMetrics if needed:
	finalDetail.Metrics.OpenIssues = detail.OpenIssuesCount

	return finalDetail
}

func countCommits(sdk *resilientbridge.ResilientBridge, owner, repoName, defaultBranch string) (int, error) {
	endpoint := fmt.Sprintf("/repos/%s/%s/commits?sha=%s&per_page=1", owner, repoName, defaultBranch)
	return countItemsFromEndpoint(sdk, endpoint)
}

func countIssues(sdk *resilientbridge.ResilientBridge, owner, repoName string) (int, error) {
	endpoint := fmt.Sprintf("/repos/%s/%s/issues?state=all&per_page=1", owner, repoName)
	return countItemsFromEndpoint(sdk, endpoint)
}

func countBranches(sdk *resilientbridge.ResilientBridge, owner, repoName string) (int, error) {
	endpoint := fmt.Sprintf("/repos/%s/%s/branches?per_page=1", owner, repoName)
	return countItemsFromEndpoint(sdk, endpoint)
}

func countPullRequests(sdk *resilientbridge.ResilientBridge, owner, repoName string) (int, error) {
	endpoint := fmt.Sprintf("/repos/%s/%s/pulls?state=all&per_page=1", owner, repoName)
	return countItemsFromEndpoint(sdk, endpoint)
}

func countReleases(sdk *resilientbridge.ResilientBridge, owner, repoName string) (int, error) {
	endpoint := fmt.Sprintf("/repos/%s/%s/releases?per_page=1", owner, repoName)
	return countItemsFromEndpoint(sdk, endpoint)
}

func countItemsFromEndpoint(sdk *resilientbridge.ResilientBridge, endpoint string) (int, error) {
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

	lastPage, err := parseLastPage(linkHeader)
	if err != nil {
		return 0, fmt.Errorf("could not parse last page: %w", err)
	}

	return lastPage, nil
}

func parseLastPage(linkHeader string) (int, error) {
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
