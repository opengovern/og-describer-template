package describer

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/opengovern/og-describer-github/pkg/sdk/models"
	"github.com/opengovern/og-describer-github/provider/model"
	resilientbridge "github.com/opengovern/resilient-bridge"
	"github.com/opengovern/resilient-bridge/adapters"
	"log"
	"net/url"
	"strings"
	"time"
)

func ListArtifactDockerFiles(ctx context.Context, githubClient GitHubClient, organizationName string, stream *models.StreamSender) ([]models.Resource, error) {
	repositories, err := getRepositories(ctx, githubClient.RestClient, organizationName)
	if err != nil {
		return nil, err
	}

	var values []models.Resource
	for _, repo := range repositories {
		repoValues := fetchArtifactDockerFiles(ctx, githubClient, stream, repo.GetFullName())
		values = append(values, repoValues...)
	}

	return values, nil
}

func fetchArtifactDockerFiles(ctx context.Context, githubClient GitHubClient, stream *models.StreamSender, repoName string) []models.Resource {
	sdk := resilientbridge.NewResilientBridge()
	sdk.SetDebug(false) // Disable debug

	sdk.RegisterProvider("github", adapters.NewGitHubAdapter(githubClient.Token), &resilientbridge.ProviderConfig{
		UseProviderLimits: true,
		MaxRetries:        3,
		BaseBackoff:       time.Second,
	})

	// Construct the query
	queryParts := []string{
		fmt.Sprintf("repo:%s", repoName),
		"filename:Dockerfile",
	}

	finalQuery := strings.Join(queryParts, " ")

	perPage := 100
	page := 1
	totalCollected := 0

	var allItems []model.CodeSearchHit

	for {
		if totalCollected >= 500 {
			break
		}

		q := url.QueryEscape(finalQuery)
		searchEndpoint := fmt.Sprintf("/search/code?q=%s&per_page=%d&page=%d", q, perPage, page)

		searchReq := &resilientbridge.NormalizedRequest{
			Method:   "GET",
			Endpoint: searchEndpoint,
			Headers:  map[string]string{"Accept": "application/vnd.github+json"},
		}

		searchResp, err := sdk.Request("github", searchReq)
		if err != nil {
			log.Printf("Error searching code: %v\n", err)
			break
		}

		if searchResp.StatusCode >= 400 {
			log.Printf("HTTP error %d searching code: %s\n", searchResp.StatusCode, string(searchResp.Data))
			break
		}

		var result model.CodeSearchResult
		if err := json.Unmarshal(searchResp.Data, &result); err != nil {
			log.Printf("Error parsing code search response: %v\n", err)
			break
		}

		if len(result.Items) == 0 {
			// No more results
			break
		}

		for _, item := range result.Items {
			allItems = append(allItems, item)
			totalCollected++
			if totalCollected >= 500 {
				break
			}
		}

		if totalCollected >= 500 || len(result.Items) < perPage {
			break
		}
		page++
	}

	var values []models.Resource

	for _, item := range allItems {
		contentReq := &resilientbridge.NormalizedRequest{
			Method:   "GET",
			Endpoint: strings.TrimPrefix(item.URL, "https://api.github.com"),
			Headers:  map[string]string{"Accept": "application/vnd.github+json"},
		}

		contentResp, err := sdk.Request("github", contentReq)
		if err != nil {
			log.Printf("Error fetching content for %s/%s: %v\n", item.Repository.FullName, item.Path, err)
			continue
		}

		if contentResp.StatusCode >= 400 {
			log.Printf("HTTP error %d fetching content for %s/%s: %s\n", contentResp.StatusCode, item.Repository.FullName, item.Path, string(contentResp.Data))
			continue
		}

		var contentData model.ContentResponse
		if err := json.Unmarshal(contentResp.Data, &contentData); err != nil {
			log.Printf("Error parsing content response for %s/%s: %v\n", item.Repository.FullName, item.Path, err)
			continue
		}

		var fileContent string
		if contentData.Encoding == "base64" {
			decoded, err := base64.StdEncoding.DecodeString(contentData.Content)
			if err != nil {
				log.Printf("Error decoding base64 content for %s/%s: %v\n", item.Repository.FullName, item.Path, err)
				continue
			}
			fileContent = string(decoded)
		} else {
			fileContent = contentData.Content
		}

		lines := strings.Split(fileContent, "\n")
		if len(lines) > 200 {
			// Skip if more than 200 lines
			log.Printf("Skipping %s/%s: more than 200 lines (%d lines)\n", item.Repository.FullName, item.Path, len(lines))
			continue
		}

		// Fetch last_updated_at via commits API
		lastUpdatedAt := ""
		if item.Repository.FullName != "" && item.Path != "" {
			commitsEndpoint := fmt.Sprintf("/repos/%s/commits?path=%s&per_page=1", item.Repository.FullName, url.QueryEscape(item.Path))
			commitReq := &resilientbridge.NormalizedRequest{
				Method:   "GET",
				Endpoint: commitsEndpoint,
				Headers:  map[string]string{"Accept": "application/vnd.github+json"},
			}

			commitResp, err := sdk.Request("github", commitReq)
			if err != nil {
				log.Printf("Error fetching commits for %s/%s: %v\n", item.Repository.FullName, item.Path, err)
			} else if commitResp.StatusCode < 400 {
				var commits []model.CommitResponse
				if err := json.Unmarshal(commitResp.Data, &commits); err == nil && len(commits) > 0 {
					if commits[0].Commit.Author.Date != "" {
						lastUpdatedAt = commits[0].Commit.Author.Date
					} else if commits[0].Commit.Committer.Date != "" {
						lastUpdatedAt = commits[0].Commit.Committer.Date
					}
				}
			}
		}

		repoObj := map[string]interface{}{
			"id":        item.Repository.ID,
			"node_id":   item.Repository.NodeID,
			"name":      item.Repository.Name,
			"full_name": item.Repository.FullName,
			"private":   item.Repository.Private,
			"public":    !item.Repository.Private,
			"owner": map[string]interface{}{
				"login":    item.Repository.Owner.Login,
				"id":       item.Repository.Owner.ID,
				"node_id":  item.Repository.Owner.NodeID,
				"html_url": item.Repository.Owner.HTMLURL,
				"type":     item.Repository.Owner.Type,
			},
			"html_url":    item.Repository.HTMLURL,
			"description": item.Repository.Description,
			"fork":        item.Repository.Fork,
		}

		output := model.ArtifactDockerFileDescription{
			Sha:                     item.Sha,
			Name:                    item.Name,
			Path:                    item.Path,
			LastUpdatedAt:           lastUpdatedAt,
			GitURL:                  item.GitURL,
			HTMLURL:                 item.HTMLURL,
			URI:                     item.HTMLURL, // Unique identifier
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
		if stream != nil {
			if err := (*stream)(value); err != nil {
				return nil
			}
		} else {
			values = append(values, value)
		}
	}

	return values
}
