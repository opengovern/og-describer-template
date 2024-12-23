// artifact_dockerfile.go
package describer

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strings"
	"time"

	"github.com/opengovern/og-describer-github/pkg/sdk/models"
	"github.com/opengovern/og-describer-github/provider/model"
	resilientbridge "github.com/opengovern/resilient-bridge"
	"github.com/opengovern/resilient-bridge/adapters"
	"github.com/opengovern/resilient-bridge/utils" // <--- calls ExtractExternalBaseImagesFromBase64
)

// MAX_RESULTS is the maximum number of Dockerfiles to collect or stream.
const MAX_RESULTS = 500

// MAX_DOCKERFILE_LEN is the maximum allowed number of lines in a Dockerfile.
const MAX_DOCKERFILE_LEN = 200

// ListArtifactDockerFiles searches for all Dockerfiles in the specified organization,
// streaming each Dockerfile as it’s processed, and returning the final slice.
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

	// Override org/repo if set in context
	if orgVal := ctx.Value("organization"); orgVal != nil {
		if orgName, ok := orgVal.(string); ok && orgName != "" {
			organizationName = orgName
		}
	}
	if repoVal := ctx.Value("repository"); repoVal != nil {
		if repoName, ok := repoVal.(string); ok && repoName != "" {
			repoFullName := fmt.Sprintf("%s/%s", organizationName, repoName)
			return fetchRepositoryDockerfiles(ctx, sdk, githubClient, organizationName, repoFullName, stream)
		}
	}

	// Otherwise list all repos in the org
	repositories, err := getRepositories(ctx, githubClient.RestClient, organizationName)
	if err != nil {
		return nil, fmt.Errorf("error fetching repositories for org %s: %w", organizationName, err)
	}

	var allValues []models.Resource
	totalCollected := 0
	perPage := 100

	// For each repository, find Dockerfiles
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

			// Process each search result
			for _, item := range result.Items {
				resource, err := GetDockerfile(ctx, githubClient, organizationName, item.Repository.FullName, item.Path, stream)
				if err != nil {
					log.Printf("Skipping %s/%s: %v\n", item.Repository.FullName, item.Path, err)
					continue
				}
				if resource == nil {
					continue
				}

				// 1) Add to local slice
				allValues = append(allValues, *resource)
				totalCollected++

				// 2) Stream the resource *now*, just like repository.go
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
				// No more pages
				break
			}
			page++
		}
	}

	// Return everything, even if we streamed
	return allValues, nil
}

// fetchRepositoryDockerfiles is the same logic but limited to a single repo.
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

		// Process each Dockerfile search result
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

			// Stream each Dockerfile now
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

// GetDockerfile fetches a single Dockerfile from GitHub, decodes the base64 content,
// checks line count, uses utils.ExtractExternalBaseImagesFromBase64 to parse images.
// If parsing fails, Images remains empty.
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

	// 1) Fetch file content from GitHub
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

	// 2) We rely on base64 content (contentData.Content).
	dockerfileB64 := contentData.Content
	if dockerfileB64 == "" {
		// No content => skip
		return nil, fmt.Errorf("no base64 content found for %s/%s", repoFullName, filePath)
	}

	// 3) Decode just to do line count check
	decoded, err := base64.StdEncoding.DecodeString(dockerfileB64)
	if err != nil {
		return nil, fmt.Errorf("error decoding base64 for %s/%s: %w", repoFullName, filePath, err)
	}
	lines := strings.Split(string(decoded), "\n")
	if len(lines) > MAX_DOCKERFILE_LEN {
		return nil, fmt.Errorf("skipping %s/%s: Dockerfile has %d lines (> %d)",
			repoFullName, filePath, len(lines), MAX_DOCKERFILE_LEN)
	}

	// 4) Parse to find external base images
	images, parseErr := utils.ExtractExternalBaseImagesFromBase64(dockerfileB64)
	if parseErr != nil {
		log.Printf("Parsing error for Dockerfile at %s/%s: %v\n", repoFullName, filePath, parseErr)
		images = []string{} // fail-safe
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

	// 6) Build the final struct
	repoObj := map[string]interface{}{
		"full_name": repoFullName,
	}

	output := model.ArtifactDockerFileDescription{
		Sha:                     contentData.Sha,
		Name:                    contentData.Name,
		Path:                    contentData.Path,
		LastUpdatedAt:           lastUpdatedAt,
		GitURL:                  contentData.GitURL,
		HTMLURL:                 contentData.HTMLURL,
		URI:                     contentData.HTMLURL,
		DockerfileContent:       string(decoded),
		DockerfileContentBase64: dockerfileB64,
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
