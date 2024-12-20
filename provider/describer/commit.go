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
)

func ListCommits(ctx context.Context, githubClient GitHubClient, organizationName string, stream *models.StreamSender) ([]models.Resource, error) {
	repositories, err := getRepositories(ctx, githubClient.RestClient, organizationName)
	if err != nil {
		return nil, err
	}

	sdk := resilientbridge.NewResilientBridge()
	sdk.RegisterProvider("github", adapters.NewGitHubAdapter(githubClient.Token), &resilientbridge.ProviderConfig{
		UseProviderLimits: true,
		MaxRetries:        3,
		BaseBackoff:       0,
	})

	var values []models.Resource
	for _, repo := range repositories {
		active, err := checkRepositoryActive(sdk, organizationName, repo.GetName())
		if err != nil {
			return nil, err
		}

		if !active {
			// Repository is archived or disabled, return 0 commits
			// No output needed, just exit gracefully.
			continue
		}
		repoValues, err := GetRepositoryCommits(ctx, sdk, stream, organizationName, repo.GetName())
		if err != nil {
			return nil, err
		}
		values = append(values, repoValues...)
	}

	return values, nil
}

func GetRepositoryCommits(ctx context.Context, sdk *resilientbridge.ResilientBridge, stream *models.StreamSender, owner, repo string) ([]models.Resource, error) {
	maxCommits := 50
	commits, err := fetchCommitList(sdk, owner, repo, maxCommits)
	if err != nil {
		log.Fatalf("Error fetching commits list: %v", err)
	}

	var values []models.Resource
	for _, c := range commits {
		commitJSON, err := fetchCommitDetails(sdk, owner, repo, c.SHA)
		if err != nil {
			log.Printf("Error fetching commit %s details: %v", c.SHA, err)
			continue
		}

		var commit model.CommitDescription

		err = json.Unmarshal(commitJSON, &commit)
		if err != nil {
			log.Println("Error unmarshaling JSON:", err)
			continue
		}

		value := models.Resource{
			ID:   commit.ID,
			Name: commit.ID,
			Description: JSONAllFieldsMarshaller{
				Value: commit,
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

//func GetAllCommits(ctx context.Context, githubClient GitHubClient, organizationName string, stream *models.StreamSender) ([]models.Resource, error) {
//	client := githubClient.RestClient
//	owner := organizationName
//	repositories, err := getRepositories(ctx, client, owner)
//	if err != nil {
//		return nil, nil
//	}
//	var values []models.Resource
//	for _, repo := range repositories {
//		repoValues, err := GetRepositoryCommits(ctx, githubClient, stream, owner, repo.GetName())
//		if err != nil {
//			return nil, err
//		}
//		values = append(values, repoValues...)
//	}
//	return values, nil
//}
//
//func GetRepositoryCommits(ctx context.Context, githubClient GitHubClient, stream *models.StreamSender, owner, repo string) ([]models.Resource, error) {
//	client := githubClient.GraphQLClient
//	var query struct {
//		RateLimit  steampipemodels.RateLimit
//		Repository struct {
//			DefaultBranchRef struct {
//				Target struct {
//					Commit struct {
//						History struct {
//							TotalCount int
//							PageInfo   steampipemodels.PageInfo
//							Nodes      []steampipemodels.Commit
//						} `graphql:"history(first: $pageSize, after: $cursor, since: $since, until: $until)"`
//					} `graphql:"... on Commit"`
//				}
//			}
//		} `graphql:"repository(owner: $owner, name: $name)"`
//	}
//	variables := map[string]interface{}{
//		"owner":    githubv4.String(owner),
//		"name":     githubv4.String(repo),
//		"pageSize": githubv4.Int(pageSize),
//		"cursor":   (*githubv4.String)(nil),
//		"since":    (*githubv4.GitTimestamp)(nil),
//		"until":    (*githubv4.GitTimestamp)(nil),
//	}
//	appendCommitColumnIncludes(&variables, commitCols())
//	repoFullName := formRepositoryFullName(owner, repo)
//	var values []models.Resource
//	for {
//		err := client.Query(ctx, &query, variables)
//		if err != nil {
//			return nil, err
//		}
//		for _, commit := range query.Repository.DefaultBranchRef.Target.Commit.History.Nodes {
//			value := models.Resource{
//				ID:   commit.Sha,
//				Name: commit.Sha,
//				Description: JSONAllFieldsMarshaller{
//					Value: model.CommitDescription{
//						Commit:         commit,
//						RepoFullName:   repoFullName,
//						AuthorLogin:    commit.Author.User.Login,
//						CommitterLogin: commit.Committer.User.Login,
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
//		if !query.Repository.DefaultBranchRef.Target.Commit.History.PageInfo.HasNextPage {
//			break
//		}
//		variables["cursor"] = githubv4.NewString(query.Repository.DefaultBranchRef.Target.Commit.History.PageInfo.EndCursor)
//	}
//	return values, nil
//}
//
//func GetRepositoryCommit(ctx context.Context, githubClient GitHubClient, organizationName string, repositoryName string, resourceID string, stream *models.StreamSender) (*models.Resource, error) {
//	repoFullName := formRepositoryFullName(organizationName, repositoryName)
//
//	var query struct {
//		RateLimit  steampipemodels.RateLimit
//		Repository struct {
//			Object struct {
//				Commit steampipemodels.Commit `graphql:"... on Commit"`
//			} `graphql:"object(oid: $sha)"`
//		} `graphql:"repository(owner: $owner, name: $name)"`
//	}
//
//	variables := map[string]interface{}{
//		"owner": githubv4.String(organizationName),
//		"name":  githubv4.String(repositoryName),
//		"sha":   githubv4.GitObjectID(resourceID),
//	}
//
//	client := githubClient.GraphQLClient
//	appendCommitColumnIncludes(&variables, commitCols())
//
//	err := client.Query(ctx, &query, variables)
//	if err != nil {
//		return nil, err
//	}
//
//	value := models.Resource{
//		ID:   query.Repository.Object.Commit.Sha,
//		Name: query.Repository.Object.Commit.Sha,
//		Description: JSONAllFieldsMarshaller{
//			Value: model.CommitDescription{
//				Commit:         query.Repository.Object.Commit,
//				RepoFullName:   repoFullName,
//				AuthorLogin:    query.Repository.Object.Commit.Author.User.Login,
//				CommitterLogin: query.Repository.Object.Commit.Committer.User.Login,
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

type commitRef struct {
	SHA string `json:"sha"`
}

// checkRepositoryActive returns false if the repository is archived or disabled, true otherwise.
func checkRepositoryActive(sdk *resilientbridge.ResilientBridge, owner, repo string) (bool, error) {
	req := &resilientbridge.NormalizedRequest{
		Method:   "GET",
		Endpoint: fmt.Sprintf("/repos/%s/%s", owner, repo),
		Headers:  map[string]string{"Accept": "application/vnd.github+json"},
	}

	resp, err := sdk.Request("github", req)
	if err != nil {
		return false, fmt.Errorf("error checking repository: %w", err)
	}
	if resp.StatusCode == 404 {
		// Repo not found, treat as inactive
		return false, nil
	} else if resp.StatusCode >= 400 {
		return false, fmt.Errorf("HTTP error %d: %s", resp.StatusCode, string(resp.Data))
	}

	var repoInfo struct {
		Archived bool `json:"archived"`
		Disabled bool `json:"disabled"`
	}
	if err := json.Unmarshal(resp.Data, &repoInfo); err != nil {
		return false, fmt.Errorf("error decoding repository info: %w", err)
	}

	if repoInfo.Archived || repoInfo.Disabled {
		return false, nil
	}
	return true, nil
}

// fetchCommitList returns up to maxCommits commit references from the repo’s default branch.
func fetchCommitList(sdk *resilientbridge.ResilientBridge, owner, repo string, maxCommits int) ([]commitRef, error) {
	var allCommits []commitRef
	perPage := 100
	page := 1

	for len(allCommits) < maxCommits {
		remaining := maxCommits - len(allCommits)
		if remaining < perPage {
			perPage = remaining
		}

		endpoint := fmt.Sprintf("/repos/%s/%s/commits?per_page=%d&page=%d", owner, repo, perPage, page)
		req := &resilientbridge.NormalizedRequest{
			Method:   "GET",
			Endpoint: endpoint,
			Headers:  map[string]string{"Accept": "application/vnd.github+json"},
		}

		resp, err := sdk.Request("github", req)
		if err != nil {
			return nil, fmt.Errorf("error fetching commits: %w", err)
		}

		if resp.StatusCode >= 400 {
			return nil, fmt.Errorf("HTTP error %d: %s", resp.StatusCode, string(resp.Data))
		}

		var commits []commitRef
		if err := json.Unmarshal(resp.Data, &commits); err != nil {
			return nil, fmt.Errorf("error decoding commit list: %w", err)
		}

		if len(commits) == 0 {
			// No more commits
			break
		}

		allCommits = append(allCommits, commits...)
		if len(allCommits) >= maxCommits {
			break
		}
		page++
	}

	if len(allCommits) > maxCommits {
		allCommits = allCommits[:maxCommits]
	}

	return allCommits, nil
}

func fetchCommitDetails(sdk *resilientbridge.ResilientBridge, owner, repo, sha string) ([]byte, error) {
	// Fetch the commit details
	req := &resilientbridge.NormalizedRequest{
		Method:   "GET",
		Endpoint: fmt.Sprintf("/repos/%s/%s/commits/%s", owner, repo, sha),
		Headers:  map[string]string{"Accept": "application/vnd.github+json"},
	}
	resp, err := sdk.Request("github", req)
	if err != nil {
		return nil, fmt.Errorf("error fetching commit details: %w", err)
	}
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("HTTP error %d: %s", resp.StatusCode, string(resp.Data))
	}

	var commitData map[string]interface{}
	if err := json.Unmarshal(resp.Data, &commitData); err != nil {
		return nil, fmt.Errorf("error unmarshaling commit details: %w", err)
	}

	// Helper functions
	getString := func(m map[string]interface{}, key string) *string {
		if m == nil {
			return nil
		}
		if val, ok := m[key].(string); ok {
			return &val
		}
		return nil
	}
	getFloat := func(m map[string]interface{}, key string) *int {
		if m == nil {
			return nil
		}
		if val, ok := m[key].(float64); ok {
			v := int(val)
			return &v
		}
		return nil
	}
	getBool := func(m map[string]interface{}, key string) *bool {
		if m == nil {
			return nil
		}
		if val, ok := m[key].(bool); ok {
			return &val
		}
		return nil
	}

	commitSha := getString(commitData, "sha")
	htmlURL := getString(commitData, "html_url")
	nodeID := getString(commitData, "node_id")

	commitSection, _ := commitData["commit"].(map[string]interface{})
	message := getString(commitSection, "message")

	var date *string
	var commitAuthor map[string]interface{}
	if commitSection != nil {
		if ca, ok := commitSection["author"].(map[string]interface{}); ok {
			commitAuthor = ca
			date = getString(ca, "date")
		}
	}

	// stats
	var stats map[string]interface{}
	if s, ok := commitData["stats"].(map[string]interface{}); ok {
		stats = s
	}
	additions := getFloat(stats, "additions")
	deletions := getFloat(stats, "deletions")
	total := getFloat(stats, "total")

	// author
	authorObj := map[string]interface{}{
		"email":    nil,
		"html_url": nil,
		"id":       nil,
		"login":    nil,
		"name":     nil,
		"node_id":  nil,
		"type":     nil,
	}
	if commitAuthor != nil {
		if email := getString(commitAuthor, "email"); email != nil {
			authorObj["email"] = *email
		}
		if name := getString(commitAuthor, "name"); name != nil {
			authorObj["name"] = *name
		}
	}

	if topAuthor, ok := commitData["author"].(map[string]interface{}); ok {
		if login := getString(topAuthor, "login"); login != nil {
			authorObj["login"] = *login
		}
		if idVal, ok := topAuthor["id"].(float64); ok {
			authorObj["id"] = int(idVal)
		}
		if n := getString(topAuthor, "node_id"); n != nil {
			authorObj["node_id"] = *n
		}
		if h := getString(topAuthor, "html_url"); h != nil {
			authorObj["html_url"] = *h
		}
		if t := getString(topAuthor, "type"); t != nil {
			authorObj["type"] = *t
		}
	}

	// files
	filesArray := []interface{}{}
	if files, ok := commitData["files"].([]interface{}); ok {
		for _, f := range files {
			if fm, ok := f.(map[string]interface{}); ok {
				newFile := map[string]interface{}{
					"additions": nil,
					"changes":   nil,
					"deletions": nil,
					"filename":  nil,
					"sha":       nil,
					"status":    nil,
				}
				if a := getFloat(fm, "additions"); a != nil {
					newFile["additions"] = *a
				}
				if c := getFloat(fm, "changes"); c != nil {
					newFile["changes"] = *c
				}
				if d := getFloat(fm, "deletions"); d != nil {
					newFile["deletions"] = *d
				}
				if fn := getString(fm, "filename"); fn != nil {
					newFile["filename"] = *fn
				}
				if sh := getString(fm, "sha"); sh != nil {
					newFile["sha"] = *sh
				}
				if st := getString(fm, "status"); st != nil {
					newFile["status"] = *st
				}
				filesArray = append(filesArray, newFile)
			}
		}
	}

	// parents
	parentsArray := []interface{}{}
	if parents, ok := commitData["parents"].([]interface{}); ok {
		for _, p := range parents {
			if pm, ok := p.(map[string]interface{}); ok {
				newParent := map[string]interface{}{
					"sha": nil,
				}
				if ps := getString(pm, "sha"); ps != nil {
					newParent["sha"] = *ps
				}
				parentsArray = append(parentsArray, newParent)
			}
		}
	}

	// comment_count
	var commentCount *int
	if commitSection != nil {
		commentCount = getFloat(commitSection, "comment_count")
	}

	// tree (only sha)
	var treeObj map[string]interface{}
	if commitSection != nil {
		if tree, ok := commitSection["tree"].(map[string]interface{}); ok {
			treeObj = map[string]interface{}{
				"sha": nil,
			}
			if tsha := getString(tree, "sha"); tsha != nil {
				treeObj["sha"] = *tsha
			}
		}
	}

	// verification details
	var isVerified *bool
	var verificationDetails map[string]interface{}
	if commitSection != nil {
		if verification, ok := commitSection["verification"].(map[string]interface{}); ok {
			isVerified = getBool(verification, "verified")
			reason := getString(verification, "reason")
			signature := getString(verification, "signature")
			verifiedAt := getString(verification, "verified_at")

			verificationDetails = map[string]interface{}{
				"reason":      nil,
				"signature":   nil,
				"verified_at": nil,
			}
			if reason != nil {
				verificationDetails["reason"] = *reason
			}
			if signature != nil {
				verificationDetails["signature"] = *signature
			}
			if verifiedAt != nil {
				verificationDetails["verified_at"] = *verifiedAt
			}
		}
	}

	// additional_details
	additionalDetailsObj := map[string]interface{}{
		"node_id":              nil,
		"parents":              parentsArray,
		"tree":                 nil,
		"verification_details": nil,
	}
	if nodeID != nil {
		additionalDetailsObj["node_id"] = *nodeID
	}
	if treeObj != nil {
		additionalDetailsObj["tree"] = treeObj
	}
	if verificationDetails != nil {
		additionalDetailsObj["verification_details"] = verificationDetails
	}

	// Fetch associated pull requests
	prs, err := fetchPullRequestsForCommit(sdk, owner, repo, sha)
	if err != nil {
		prs = []int{}
	}

	// Determine the branch:
	// If we have at least one PR, fetch the first PR details and use its base.ref as the branch.
	var branchName string
	if len(prs) > 0 {
		prBranch, err := fetchFirstPRBranch(sdk, owner, repo, prs[0])
		if err == nil && prBranch != "" {
			branchName = prBranch
		}
	}

	// If no PR-based branch found, fallback to findBranchByCommit
	if branchName == "" {
		bname, berr := findBranchByCommit(sdk, owner, repo, sha)
		if berr == nil && bname != "" {
			branchName = bname
		}
	}

	// target
	targetObj := map[string]interface{}{
		"branch":       nil,
		"organization": owner,
		"repository":   repo,
	}
	if branchName != "" {
		targetObj["branch"] = branchName
	}

	// Convert pointers to interface{}
	finalID := func() interface{} {
		if commitSha == nil {
			return nil
		}
		return *commitSha
	}()
	finalDate := func() interface{} {
		if date == nil {
			return nil
		}
		return *date
	}()
	finalMessage := func() interface{} {
		if message == nil {
			return nil
		}
		return *message
	}()
	finalHtmlURL := func() interface{} {
		if htmlURL == nil {
			return nil
		}
		return *htmlURL
	}()
	finalIsVerified := func() interface{} {
		if isVerified == nil {
			return nil
		}
		return *isVerified
	}()
	finalCommentCount := func() interface{} {
		if commentCount == nil {
			return nil
		}
		return *commentCount
	}()
	finalAdditions := func() interface{} {
		if additions == nil {
			return nil
		}
		return *additions
	}()
	finalDeletions := func() interface{} {
		if deletions == nil {
			return nil
		}
		return *deletions
	}()
	finalTotal := func() interface{} {
		if total == nil {
			return nil
		}
		return *total
	}()

	output := map[string]interface{}{
		"id":          finalID,
		"date":        finalDate,
		"message":     finalMessage,
		"html_url":    finalHtmlURL,
		"target":      targetObj,
		"is_verified": finalIsVerified,
		"author":      authorObj,
		"changes": map[string]interface{}{
			"additions": finalAdditions,
			"deletions": finalDeletions,
			"total":     finalTotal,
		},
		"comment_count":      finalCommentCount,
		"additional_details": additionalDetailsObj,
		"files":              filesArray,
		"pull_requests":      prs,
	}

	modifiedData, err := json.MarshalIndent(output, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("error encoding final commit details: %w", err)
	}

	return modifiedData, nil
}

func fetchPullRequestsForCommit(sdk *resilientbridge.ResilientBridge, owner, repo, sha string) ([]int, error) {
	req := &resilientbridge.NormalizedRequest{
		Method:   "GET",
		Endpoint: fmt.Sprintf("/repos/%s/%s/commits/%s/pulls", owner, repo, sha),
		Headers:  map[string]string{"Accept": "application/vnd.github+json"},
	}

	resp, err := sdk.Request("github", req)
	if err != nil {
		return nil, fmt.Errorf("error fetching pull requests for commit: %w", err)
	}

	if resp.StatusCode == 409 || resp.StatusCode == 404 {
		return []int{}, nil
	}
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("HTTP error %d: %s", resp.StatusCode, string(resp.Data))
	}

	var pulls []map[string]interface{}
	if err := json.Unmarshal(resp.Data, &pulls); err != nil {
		return nil, fmt.Errorf("error decoding pull requests: %w", err)
	}

	var prNumbers []int
	for _, pr := range pulls {
		if num, ok := pr["number"].(float64); ok {
			prNumbers = append(prNumbers, int(num))
		}
	}

	return prNumbers, nil
}

// fetchFirstPRBranch fetches details of a given pull request number and returns the base branch
func fetchFirstPRBranch(sdk *resilientbridge.ResilientBridge, owner, repo string, prNumber int) (string, error) {
	req := &resilientbridge.NormalizedRequest{
		Method:   "GET",
		Endpoint: fmt.Sprintf("/repos/%s/%s/pulls/%d", owner, repo, prNumber),
		Headers:  map[string]string{"Accept": "application/vnd.github+json"},
	}
	resp, err := sdk.Request("github", req)
	if err != nil {
		return "", fmt.Errorf("error fetching pull request details: %w", err)
	}
	if resp.StatusCode >= 400 {
		return "", fmt.Errorf("HTTP error %d: %s", resp.StatusCode, string(resp.Data))
	}

	var prData map[string]interface{}
	if err := json.Unmarshal(resp.Data, &prData); err != nil {
		return "", fmt.Errorf("error decoding pull request details: %w", err)
	}

	base, ok := prData["base"].(map[string]interface{})
	if !ok {
		return "", nil
	}

	ref, ok := base["ref"].(string)
	if !ok {
		return "", nil
	}

	return ref, nil
}

func findBranchByCommit(sdk *resilientbridge.ResilientBridge, owner, repo, sha string) (string, error) {
	req := &resilientbridge.NormalizedRequest{
		Method:   "GET",
		Endpoint: fmt.Sprintf("/repos/%s/%s/branches?sha=%s", owner, repo, sha),
		Headers:  map[string]string{"Accept": "application/vnd.github+json"},
	}
	resp, err := sdk.Request("github", req)
	if err != nil {
		return "", fmt.Errorf("error fetching branches for commit: %w", err)
	}
	if resp.StatusCode >= 400 {
		return "", fmt.Errorf("HTTP error %d: %s", resp.StatusCode, string(resp.Data))
	}

	var branches []map[string]interface{}
	if err := json.Unmarshal(resp.Data, &branches); err != nil {
		return "", fmt.Errorf("error decoding branches: %w", err)
	}

	if len(branches) == 0 {
		return "", nil
	}

	if name, ok := branches[0]["name"].(string); ok && name != "" {
		return name, nil
	}

	return "", nil
}
