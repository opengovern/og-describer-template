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
)

// MAX_RESULTS is the maximum number of Dockerfiles to collect or stream.
const MAX_RESULTS = 500

// ListDockerFile lists references to Dockerfiles in all repositories for the given organization.
// If a stream is provided, results are streamed. If not, a slice of resources is returned.
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

	org := ctx.Value("organization")
	if org != nil {
		orgName := org.(string)
		if orgName != "" {
			organizationName = orgName
		}
	}

	repo := ctx.Value("repository")
	if repo != nil {
		repoName := repo.(string)
		if repoName != "" {
			repoFullName := fmt.Sprintf("%s/%s", organizationName, repoName)
			return fetchRepositoryDockerfiles(ctx, sdk, githubClient, organizationName, repoFullName, stream)
		}
	}

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

			for _, item := range result.Items {
				resource, err := GetDockerfile(ctx, githubClient, organizationName, item.Repository.FullName, item.Path, stream)
				if err != nil {
					log.Printf("Skipping %s/%s: %v\n", item.Repository.FullName, item.Path, err)
					continue
				}
				if resource == nil {
					// Means it didn't stream anything or some unknown reason
					continue
				}

				totalCollected++
				if stream != nil {
					// Stream the resource
					if err := (*stream)(*resource); err != nil {
						return nil, fmt.Errorf("error streaming resource: %w", err)
					}
				} else {
					// Accumulate to return later
					allValues = append(allValues, *resource)
				}

				if totalCollected >= MAX_RESULTS {
					break
				}
			}

			if len(result.Items) < perPage {
				// Fewer than perPage results means no more pages
				break
			}
			page++
		}
	}

	// If we streamed, return an empty slice since results are already sent via stream
	if stream != nil {
		return []models.Resource{}, nil
	}

	return allValues, nil
}

func fetchRepositoryDockerfiles(ctx context.Context, sdk *resilientbridge.ResilientBridge, githubClient GitHubClient, organizationName, repoFullName string, stream *models.StreamSender) ([]models.Resource, error) {
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

		for _, item := range result.Items {
			resource, err := GetDockerfile(ctx, githubClient, organizationName, item.Repository.FullName, item.Path, stream)
			if err != nil {
				log.Printf("Skipping %s/%s: %v\n", item.Repository.FullName, item.Path, err)
				continue
			}
			if resource == nil {
				// Means it didn't stream anything or some unknown reason
				continue
			}

			totalCollected++
			if stream != nil {
				// Stream the resource
				if err := (*stream)(*resource); err != nil {
					return nil, fmt.Errorf("error streaming resource: %w", err)
				}
			} else {
				// Accumulate to return later
				allValues = append(allValues, *resource)
			}

			if totalCollected >= MAX_RESULTS {
				break
			}
		}

		if len(result.Items) < perPage {
			// Fewer than perPage results means no more pages
			break
		}
		page++
	}

	return allValues, nil
}

// GetDockerfile fetches the details and content of a single Dockerfile given the repo and file path.
// It returns a fully populated resource with Dockerfile content, line count checks, and last updated at info.
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
		return nil, fmt.Errorf("HTTP error %d fetching content for %s/%s: %s", contentResp.StatusCode, repoFullName, filePath, string(contentResp.Data))
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
		return nil, fmt.Errorf("skipping %s/%s: more than 200 lines (%d lines)", repoFullName, filePath, len(lines))
	}

	// Fetch last_updated_at via commits API
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

	return &value, nil
}
