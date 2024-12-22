// artifact_dockerfile.go
package describer

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strings"
	"time"

	"encoding/base64"

	"github.com/opengovern/og-describer-github/pkg/sdk/models"
	"github.com/opengovern/og-describer-github/provider/model"
	resilientbridge "github.com/opengovern/resilient-bridge"
	"github.com/opengovern/resilient-bridge/adapters"
)

// MAX_RESULTS is the maximum number of Dockerfiles to collect or stream.
const MAX_RESULTS = 500

// ListArtifactDockerFiles searches for all Dockerfiles in an org’s repos,
// then for each file found, it calls GetDockerFile(...) to fetch/stream
// the *final* Dockerfile resource instead of returning a partial reference.
func ListArtifactDockerFiles(
	ctx context.Context,
	githubClient GitHubClient,
	organizationName string,
	stream *models.StreamSender,
) ([]models.Resource, error) {

	// 1) Enumerate all repositories.
	repositories, err := getRepositories(ctx, githubClient.RestClient, organizationName)
	if err != nil {
		return nil, fmt.Errorf("error fetching repositories for org %s: %w", organizationName, err)
	}

	// 2) Set up the resilient-bridge client
	sdk := resilientbridge.NewResilientBridge()
	sdk.SetDebug(false)
	sdk.RegisterProvider("github", adapters.NewGitHubAdapter(githubClient.Token), &resilientbridge.ProviderConfig{
		UseProviderLimits: true,
		MaxRetries:        3,
		BaseBackoff:       time.Second,
	})

	// We'll accumulate final Dockerfile resources here (only if no stream).
	var allValues []models.Resource

	totalCollected := 0
	perPage := 100

	// 3) For each repository, search for Dockerfiles by code search
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
				log.Printf("HTTP error %d searching code in %s: %s\n", searchResp.StatusCode, repoFullName, string(searchResp.Data))
				break
			}

			var result model.CodeSearchResult
			if err := json.Unmarshal(searchResp.Data, &result); err != nil {
				log.Printf("Error parsing code search response for %s: %v\n", repoFullName, err)
				break
			}

			// If no items returned, no more results
			if len(result.Items) == 0 {
				break
			}

			// 4) For each search result referencing a Dockerfile,
			//    call GetDockerFile(...) to get the final resource.
			for _, item := range result.Items {
				if totalCollected >= MAX_RESULTS {
					break
				}

				finalResource, err := GetDockerfile(
					ctx,
					githubClient,
					organizationName,
					item.Repository.FullName, // e.g. "orgName/repoName"
					item.Path,                // path to Dockerfile
					stream,
				)
				if err != nil {
					// e.g. if file was >200 lines, or some other fetch error
					log.Printf("Skipping %s/%s: %v\n", item.Repository.FullName, item.Path, err)
					continue
				}
				if finalResource == nil {
					// Means it didn't stream anything or some unknown reason
					continue
				}

				// If we have a stream, we've *already* streamed the final resource in GetDockerfile.
				// If we do *not* have a stream, we accumulate to return later.
				if stream == nil {
					allValues = append(allValues, *finalResource)
				}

				totalCollected++
			}

			// 5) If fewer than perPage results, we're done with this repo
			if len(result.Items) < perPage {
				break
			}
			page++
		}
	}

	// 6) If we streamed everything, return empty slice
	if stream != nil {
		return []models.Resource{}, nil
	}

	return allValues, nil
}

// GetDockerfile fetches the actual Dockerfile content, line-count checks, last-updated info, etc.
// (Unmodified from your existing code; only repeated here for clarity.)
func GetDockerfile(ctx context.Context, githubClient GitHubClient, organizationName, repoFullName, filePath string, stream *models.StreamSender) (*models.Resource, error) {
	sdk := resilientbridge.NewResilientBridge()
	sdk.SetDebug(false)

	sdk.RegisterProvider("github", adapters.NewGitHubAdapter(githubClient.Token), &resilientbridge.ProviderConfig{
		UseProviderLimits: true,
		MaxRetries:        3,
		BaseBackoff:       time.Second,
	})

	// Fetch file content
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

	var fileContent string
	if contentData.Encoding == "base64" {
		decoded, err := base64.StdEncoding.DecodeString(contentData.Content)
		if err != nil {
			return nil, fmt.Errorf("error decoding base64 content for %s/%s: %w", repoFullName, filePath, err)
		}
		fileContent = string(decoded)
	} else {
		fileContent = contentData.Content
	}

	lines := strings.Split(fileContent, "\n")
	if len(lines) > 200 {
		// Example “business rule” => skip if >200 lines
		return nil, fmt.Errorf("skipping %s/%s: more than 200 lines (%d lines)",
			repoFullName, filePath, len(lines))
	}

	// Grab last_updated_at
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

	// Final output describing this Dockerfile in detail
	output := model.ArtifactDockerFileDescription{
		Sha:                     contentData.Sha,
		Name:                    contentData.Name,
		Path:                    contentData.Path,
		LastUpdatedAt:           lastUpdatedAt,
		GitURL:                  contentData.GitURL,
		HTMLURL:                 contentData.HTMLURL,
		URI:                     contentData.HTMLURL,
		DockerfileContent:       fileContent,
		DockerfileContentBase64: contentData.Content,
		Repository:              repoObj,
	}

	value := models.Resource{
		ID:   output.URI,
		Name: output.Name,
		Description: JSONAllFieldsMarshaller{
			Value: output,
		},
	}

	// Stream if provided
	if stream != nil {
		if err := (*stream)(value); err != nil {
			return nil, err
		}
	}

	// Return the final Dockerfile resource
	return &value, nil
}
