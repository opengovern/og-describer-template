// artifact_dockerfile.go
package describer

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/moby/buildkit/frontend/dockerfile/parser" // BuildKit Dockerfile parser

	"github.com/opengovern/og-describer-github/pkg/sdk/models"
	"github.com/opengovern/og-describer-github/provider/model"
	resilientbridge "github.com/opengovern/resilient-bridge"
	"github.com/opengovern/resilient-bridge/adapters"
)

// MAX_RESULTS is the maximum number of Dockerfiles to collect/stream.
const MAX_RESULTS = 500

// MAX_DOCKERFILE_LEN is the maximum allowed number of lines in a Dockerfile.
const MAX_DOCKERFILE_LEN = 200

// ListArtifactDockerFiles searches for all Dockerfiles in the specified organization.
// If a stream is provided, results are streamed AND appended to the final return slice.
func ListArtifactDockerFiles(
	ctx context.Context,
	githubClient GitHubClient,
	organizationName string,
	stream *models.StreamSender,
) ([]models.Resource, error) {
	sdk := resilientbridge.NewResilientBridge()
	sdk.SetDebug(false)
	sdk.RegisterProvider("github", adapters.NewGitHubAdapter(githubClient.Token), &resilientbridge.ProviderConfig{
		UseProviderLimits: true,
		MaxRetries:        3,
		BaseBackoff:       time.Second,
	})

	// Check context for overrides
	if orgVal := ctx.Value("organization"); orgVal != nil {
		if orgName, ok := orgVal.(string); ok && orgName != "" {
			organizationName = orgName
		}
	}

	// If a specific repository is set, only fetch Dockerfiles for that repo.
	if repoVal := ctx.Value("repository"); repoVal != nil {
		if repoName, ok := repoVal.(string); ok && repoName != "" {
			repoFullName := fmt.Sprintf("%s/%s", organizationName, repoName)
			return fetchRepositoryDockerfiles(ctx, sdk, githubClient, organizationName, repoFullName, stream)
		}
	}

	// Otherwise, fetch all repos in the org.
	repositories, err := getRepositories(ctx, githubClient.RestClient, organizationName)
	if err != nil {
		return nil, fmt.Errorf("error fetching repositories for org %s: %w", organizationName, err)
	}

	var allValues []models.Resource
	totalCollected := 0
	perPage := 100

	// For each repository, search for Dockerfiles
	for _, repo := range repositories {
		if totalCollected >= MAX_RESULTS {
			break
		}
		repoFullName := repo.GetFullName()

		queryParts := []string{
			fmt.Sprintf("repo:%s", repoFullName),
			"filename:Dockerfile",
		}
		finalQuery := strings.Join(queryParts, " ")

		page := 1
		for totalCollected < MAX_RESULTS {
			q := url.QueryEscape(finalQuery)
			searchEndpoint := fmt.Sprintf("/search/code?q=%s&per_page=%d&page=%d", q, perPage, page)

			searchReq := &resilientbridge.NormalizedRequest{
				Method:   "GET",
				Endpoint: searchEndpoint,
				Headers:  map[string]string{"Accept": "application/vnd.github+json"},
			}

			searchResp, err := sdk.Request("github", searchReq)
			if err != nil {
				log.Printf("Error searching code in %s: %v\n", repoFullName, err)
				break
			}
			if searchResp.StatusCode >= 400 {
				log.Printf("HTTP error %d searching code in %s: %s\n",
					searchResp.StatusCode, repoFullName, string(searchResp.Data))
				break
			}

			var result model.CodeSearchResult
			if err := json.Unmarshal(searchResp.Data, &result); err != nil {
				log.Printf("Error parsing code search response for %s: %v\n", repoFullName, err)
				break
			}

			// If no items, no more results
			if len(result.Items) == 0 {
				break
			}

			for _, item := range result.Items {
				resource, err := GetDockerfile(ctx, githubClient, organizationName, item.Repository.FullName, item.Path, stream)
				if err != nil {
					log.Printf("Skipping %s/%s: %v\n", item.Repository.FullName, item.Path, err)
					continue
				}
				if resource == nil {
					continue
				}

				// Always add to our slice
				allValues = append(allValues, *resource)
				totalCollected++

				// If a stream is provided, also stream the resource
				if stream != nil {
					if err := (*stream)(*resource); err != nil {
						return allValues, fmt.Errorf("error streaming resource: %w", err)
					}
				}

				if totalCollected >= MAX_RESULTS {
					break
				}
			}

			if len(result.Items) < perPage {
				break
			}
			page++
		}
	}

	// Return everything, regardless of streaming
	return allValues, nil
}

// fetchRepositoryDockerfiles is the same logic as above but scoped to a single repo.
func fetchRepositoryDockerfiles(
	ctx context.Context,
	sdk *resilientbridge.ResilientBridge,
	githubClient GitHubClient,
	organizationName, repoFullName string,
	stream *models.StreamSender,
) ([]models.Resource, error) {

	var allValues []models.Resource
	totalCollected := 0
	perPage := 100

	queryParts := []string{
		fmt.Sprintf("repo:%s", repoFullName),
		"filename:Dockerfile",
	}
	finalQuery := strings.Join(queryParts, " ")

	page := 1
	for totalCollected < MAX_RESULTS {
		q := url.QueryEscape(finalQuery)
		searchEndpoint := fmt.Sprintf("/search/code?q=%s&per_page=%d&page=%d", q, perPage, page)

		searchReq := &resilientbridge.NormalizedRequest{
			Method:   "GET",
			Endpoint: searchEndpoint,
			Headers:  map[string]string{"Accept": "application/vnd.github+json"},
		}

		searchResp, err := sdk.Request("github", searchReq)
		if err != nil {
			log.Printf("Error searching code in %s: %v\n", repoFullName, err)
			break
		}
		if searchResp.StatusCode >= 400 {
			log.Printf("HTTP error %d searching code in %s: %s\n",
				searchResp.StatusCode, repoFullName, string(searchResp.Data))
			break
		}

		var result model.CodeSearchResult
		if err := json.Unmarshal(searchResp.Data, &result); err != nil {
			log.Printf("Error parsing code search response for %s: %v\n", repoFullName, err)
			break
		}

		if len(result.Items) == 0 {
			break
		}

		for _, item := range result.Items {
			resource, err := GetDockerfile(ctx, githubClient, organizationName, item.Repository.FullName, item.Path, stream)
			if err != nil {
				log.Printf("Skipping %s/%s: %v\n", item.Repository.FullName, item.Path, err)
				continue
			}
			if resource == nil {
				continue
			}

			allValues = append(allValues, *resource)
			totalCollected++

			if stream != nil {
				if err := (*stream)(*resource); err != nil {
					return allValues, fmt.Errorf("error streaming resource: %w", err)
				}
			}

			if totalCollected >= MAX_RESULTS {
				break
			}
		}

		if len(result.Items) < perPage {
			break
		}
		page++
	}

	return allValues, nil
}

// GetDockerfile fetches details of a single Dockerfile, including content and metadata,
// then **parses the Dockerfile from base64** to populate Images. If parsing fails, Images stays empty.
func GetDockerfile(
	ctx context.Context,
	githubClient GitHubClient,
	organizationName, repoFullName, filePath string,
	stream *models.StreamSender,
) (*models.Resource, error) {

	sdk := resilientbridge.NewResilientBridge()
	sdk.SetDebug(false)
	sdk.RegisterProvider("github", adapters.NewGitHubAdapter(githubClient.Token), &resilientbridge.ProviderConfig{
		UseProviderLimits: true,
		MaxRetries:        3,
		BaseBackoff:       time.Second,
	})

	// 1) Fetch file content (including base64-encoded content)
	contentEndpoint := fmt.Sprintf("/repos/%s/contents/%s", repoFullName, url.PathEscape(filePath))
	contentReq := &resilientbridge.NormalizedRequest{
		Method:   "GET",
		Endpoint: contentEndpoint,
		Headers:  map[string]string{"Accept": "application/vnd.github+json"},
	}

	contentResp, err := sdk.Request("github", contentReq)
	if err != nil {
		return nil, fmt.Errorf("error fetching content for %s/%s: %w", repoFullName, filePath, err)
	}
	if contentResp.StatusCode >= 400 {
		return nil, fmt.Errorf("HTTP error %d fetching content for %s/%s: %s",
			contentResp.StatusCode, repoFullName, filePath, string(contentResp.Data))
	}

	var contentData model.ContentResponse
	if err := json.Unmarshal(contentResp.Data, &contentData); err != nil {
		return nil, fmt.Errorf("error parsing content response for %s/%s: %w", repoFullName, filePath, err)
	}

	// 2) Extract the "base64" Dockerfile content from the JSON
	//    We skip raw fileContent except for line-count check, if you still want that.
	dockerB64 := contentData.Content
	if dockerB64 == "" {
		// Fallback if no content
		return nil, fmt.Errorf("no base64 content found for %s/%s", repoFullName, filePath)
	}

	// 3) (Optional) decode just to do line count check (skip if huge)
	decodedContent, err := base64.StdEncoding.DecodeString(dockerB64)
	if err != nil {
		return nil, fmt.Errorf("error decoding base64 content for %s/%s: %w", repoFullName, filePath, err)
	}
	lines := strings.Split(string(decodedContent), "\n")
	if len(lines) > MAX_DOCKERFILE_LEN {
		return nil, fmt.Errorf(
			"skipping %s/%s: Dockerfile has %d lines, exceeds MAX_DOCKERFILE_LEN of %d",
			repoFullName, filePath, len(lines), MAX_DOCKERFILE_LEN,
		)
	}

	// 4) Attempt to parse from base64 and collect external images
	images, parseErr := extractExternalBaseImagesFromBase64(dockerB64)
	if parseErr != nil {
		log.Printf("Failed to parse Dockerfile for %s/%s: %v", repoFullName, filePath, parseErr)
		images = []string{} // fail safe
	}

	// 5) Last updated time
	var lastUpdatedAt string
	commitsEndpoint := fmt.Sprintf("/repos/%s/commits?path=%s&per_page=1", repoFullName, url.QueryEscape(filePath))
	commitReq := &resilientbridge.NormalizedRequest{
		Method:   "GET",
		Endpoint: commitsEndpoint,
		Headers:  map[string]string{"Accept": "application/vnd.github+json"},
	}
	commitResp, err := sdk.Request("github", commitReq)
	if err == nil && commitResp.StatusCode < 400 {
		var commits []model.CommitResponse
		if json.Unmarshal(commitResp.Data, &commits) == nil && len(commits) > 0 {
			if commits[0].Commit.Author.Date != "" {
				lastUpdatedAt = commits[0].Commit.Author.Date
			} else if commits[0].Commit.Committer.Date != "" {
				lastUpdatedAt = commits[0].Commit.Committer.Date
			}
		}
	}

	repoObj := map[string]interface{}{
		"full_name": repoFullName,
	}

	// 6) Build the final output struct
	output := model.ArtifactDockerFileDescription{
		Sha:                     contentData.Sha,
		Name:                    contentData.Name,
		Path:                    contentData.Path,
		LastUpdatedAt:           lastUpdatedAt,
		GitURL:                  contentData.GitURL,
		HTMLURL:                 contentData.HTMLURL,
		URI:                     contentData.HTMLURL,
		DockerfileContent:       string(decodedContent), // if you want the raw decoded content
		DockerfileContentBase64: dockerB64,
		Repository:              repoObj,
		Images:                  images,
	}

	value := models.Resource{
		ID:   output.URI,
		Name: output.Name,
		Description: JSONAllFieldsMarshaller{
			Value: output,
		},
	}
	return &value, nil
}

/*  ---------------------------------------------------------
    Use the same logic as your working snippet, but generalized
    to parse the Dockerfile from base64 content.
    --------------------------------------------------------- */

// fromInfo tracks a single FROM instruction (base image + optional alias).
type fromInfo struct {
	baseImage string
	alias     string
}

// extractExternalBaseImagesFromBase64 decodes the Dockerfile from base64,
// parses with BuildKit, collects base images, expands ARG references, and
// filters out internal aliases.
func extractExternalBaseImagesFromBase64(encodedDockerfile string) ([]string, error) {
	// --- A: Decode base64 Dockerfile content ---
	decoded, err := base64.StdEncoding.DecodeString(encodedDockerfile)
	if err != nil {
		return nil, fmt.Errorf("failed to base64-decode Dockerfile: %w", err)
	}

	// --- B: Parse with Docker BuildKit parser ---
	res, err := parser.Parse(strings.NewReader(string(decoded)))
	if err != nil {
		return nil, fmt.Errorf("BuildKit parser error: %w", err)
	}
	if res == nil || res.AST == nil {
		return nil, nil
	}

	// --- C: Collect top-level ARG instructions for naive variable expansion ---
	argsMap := collectArgs(res.AST)

	// --- D: Gather FROM instructions, expand them
	var fromLines []fromInfo
	stageAliases := make(map[string]bool)
	for _, stmt := range res.AST.Children {
		if strings.EqualFold(stmt.Value, "from") {
			tokens := collectStatementTokens(stmt)
			base, alias := parseFromLine(tokens, argsMap)
			if alias != "" {
				stageAliases[alias] = true
			}
			fromLines = append(fromLines, fromInfo{baseImage: base, alias: alias})
		}
	}

	// --- E: Filter out references to internal aliases (FROM builder, etc.)
	var external []string
	for _, f := range fromLines {
		// If the baseImage is itself a known alias, skip it
		if stageAliases[f.baseImage] {
			continue
		}
		external = append(external, f.baseImage)
	}
	return external, nil
}

// collectArgs gathers top-level ARG instructions from the AST
func collectArgs(ast *parser.Node) map[string]string {
	argsMap := make(map[string]string)
	for _, stmt := range ast.Children {
		if strings.EqualFold(stmt.Value, "arg") {
			tokens := collectStatementTokens(stmt)
			for _, t := range tokens {
				k, v := parseArgKeyValue(t)
				if k != "" && v != "" {
					argsMap[k] = v
				}
			}
		}
	}
	return argsMap
}

// collectStatementTokens flattens tokens for a single statement, stopping if we
// hit another Dockerfile instruction. This matches your original snippet.
func collectStatementTokens(stmt *parser.Node) []string {
	var tokens []string
	cur := stmt.Next
	for cur != nil {
		if isInstructionKeyword(cur.Value) {
			break
		}
		tokens = append(tokens, cur.Value)
		cur = cur.Next
	}
	return tokens
}

// parseFromLine processes tokens from a FROM statement, e.g. ["--platform=${PLATFORM}", "${GO_IMAGE}", "AS", "builder"]
// returns (baseImage, alias).
func parseFromLine(tokens []string, argsMap map[string]string) (string, string) {
	var base, alias string
	for i := 0; i < len(tokens); i++ {
		t := tokens[i]
		if strings.HasPrefix(t, "--") {
			// e.g. --platform=...
			continue
		}
		if strings.EqualFold(t, "AS") && i+1 < len(tokens) {
			alias = tokens[i+1]
			break
		}
		base = expandArgs(t, argsMap)
	}
	return base, alias
}

// parseArgKeyValue splits "KEY=VALUE". If no '=', we get (KEY, "").
func parseArgKeyValue(argToken string) (string, string) {
	parts := strings.SplitN(argToken, "=", 2)
	if len(parts) == 1 {
		return parts[0], ""
	}
	return parts[0], parts[1]
}

// expandArgs does a naive expansion of $VAR / ${VAR} with known defaults in argsMap.
// If no default is found, we leave it as ${VAR}.
func expandArgs(input string, argsMap map[string]string) string {
	return os.Expand(input, func(key string) string {
		if val, ok := argsMap[key]; ok {
			return val
		}
		return fmt.Sprintf("${%s}", key)
	})
}

// isInstructionKeyword checks if a token is recognized as a Dockerfile instruction.
func isInstructionKeyword(s string) bool {
	switch strings.ToUpper(s) {
	case "ADD",
		"ARG",
		"CMD",
		"COPY",
		"ENTRYPOINT",
		"ENV",
		"EXPOSE",
		"FROM",
		"HEALTHCHECK",
		"LABEL",
		"MAINTAINER",
		"ONBUILD",
		"RUN",
		"SHELL",
		"STOPSIGNAL",
		"USER",
		"VOLUME",
		"WORKDIR":
		return true
	}
	return false
}
