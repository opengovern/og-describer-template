package describers

import (
	"context"
	"github.com/opengovern/og-describer-github/discovery/pkg/models"
	model "github.com/opengovern/og-describer-github/discovery/provider"
	resilientbridge "github.com/opengovern/resilient-bridge"
	"github.com/opengovern/resilient-bridge/adapters"
	"github.com/opengovern/resilient-bridge/utils"
	"log"
	"math"
	"strings"
)

func ListArtifactAIModels(ctx context.Context, githubClient model.GitHubClient, organizationName string, stream *models.StreamSender) ([]models.Resource, error) {
	sdk := resilientbridge.NewResilientBridge()
	sdk.RegisterProvider("github", adapters.NewGitHubAdapter(githubClient.Token), &resilientbridge.ProviderConfig{
		UseProviderLimits: true,
		MaxRetries:        3,
		BaseBackoff:       0,
	})

	chunkSize := utils.DEFAULT_CHUNK_SIZE

	var allItems []utils.Item
	extChunks := utils.ChunkBySize(utils.FileExtensions, chunkSize)

	for _, chunk := range extChunks {
		var queryParts []string
		for _, ext := range chunk {
			queryParts = append(queryParts, "extension:"+ext)
		}
		query := strings.Join(queryParts, " OR ")

		// Page 1
		firstResult, err := utils.SearchGitHub(sdk, query, 1)
		if err != nil {
			log.Printf("Error searching GitHub: %v", err)
			continue
		}
		allItems = append(allItems, firstResult.Items...)

		// If total_count is huge, cap at ~10 pages
		totalPages := int(math.Ceil(float64(firstResult.TotalCount) / 100.0))
		if totalPages > 10 {
			totalPages = 10
		}

		// Page 2..N
		for page := 2; page <= totalPages; page++ {
			result, err := utils.SearchGitHub(sdk, query, page)
			if err != nil {
				log.Printf("Error searching GitHub (page %d): %v", page, err)
				break
			}
			allItems = append(allItems, result.Items...)
			if len(result.Items) < 100 {
				// no more results to fetch
				break
			}
		}
	}

	if len(allItems) == 0 {
		return nil, nil
	}

	groups := utils.GatherDirectories(allItems, false)

	keptBinaries := utils.SampleAndFilterDirectories(sdk, groups, 5, false)

	repoMap := utils.CreateDetailedRepoExtensionMap(keptBinaries)

	var values []models.Resource
	for key, value := range repoMap {
		resource := models.Resource{
			ID:   key,
			Name: key,
			Description: model.ArtifactAIModelDescription{
				Name:               key,
				RepositoryID:       value.RepositoryID,
				RepositoryName:     value.RepositoryName,
				RepositoryFullName: value.RepositoryFullName,
				Extensions:         value.Extensions,
			},
		}
		if stream != nil {
			if err := (*stream)(resource); err != nil {
				return nil, err
			}
		} else {
			values = append(values, resource)
		}
	}

	return values, nil
}
