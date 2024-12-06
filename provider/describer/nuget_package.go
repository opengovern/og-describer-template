package describer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/opengovern/og-describer-github/pkg/sdk/models"
	"github.com/opengovern/og-describer-github/provider/model"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func GetNugetPackageList(ctx context.Context, githubClient GitHubClient, organizationName string, stream *models.StreamSender) ([]models.Resource, error) {
	var packages []model.PackageResponse
	var packagesResp model.PackageListResponse
	var pagesRemaining = true
	page := 1
	baseURL := "https://api.github.com/orgs/"
	client := http.DefaultClient
	var values []models.Resource
	for pagesRemaining {
		params := url.Values{}
		params.Set("package_type", "nuget")
		params.Set("page", strconv.Itoa(page))
		params.Set("per_page", "100")
		finalURL := fmt.Sprintf("%s%s/packages?%s", baseURL, organizationName, params.Encode())
		req, err := http.NewRequest("GET", finalURL, nil)
		if err != nil {
			return nil, err
		}
		req.Header.Set("Accept", "application/vnd.github+json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", githubClient.Token))
		req.Header.Set("X-GitHub-Api-Version", "2022-11-28")
		resp, err := client.Do(req)
		if err = json.NewDecoder(resp.Body).Decode(&packagesResp); err != nil {
			return nil, err
		}
		packages = append(packages, packagesResp.Items...)
		linkHeader := resp.Header.Get("Link")
		pagesRemaining = strings.Contains(linkHeader, `rel="next"`)
		if pagesRemaining {
			page++
			break
		}
	}
	for _, packageData := range packages {
		value := models.Resource{
			ID:   strconv.Itoa(packageData.ID),
			Name: packageData.Name,
			Description: JSONAllFieldsMarshaller{
				Value: model.PackageDescription{
					ID:        strconv.Itoa(packageData.ID),
					Name:      packageData.Name,
					URL:       packageData.URL,
					CreatedAt: packageData.CreatedAt,
					UpdatedAt: packageData.UpdatedAt,
				},
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

func GetNugetPackage(ctx context.Context, githubClient GitHubClient, organizationName string, repositoryName string, resourceID string, stream *models.StreamSender) (*models.Resource, error) {
	var packageData model.PackageResponse
	baseURL := "https://api.github.com/orgs/"
	client := http.DefaultClient
	finalURL := fmt.Sprintf("%s%s/packages/%s/%s", baseURL, organizationName, "nuget", resourceID)
	req, err := http.NewRequest("GET", finalURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", githubClient.Token))
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")
	resp, err := client.Do(req)
	if err = json.NewDecoder(resp.Body).Decode(&packageData); err != nil {
		return nil, err
	}
	value := models.Resource{
		ID:   strconv.Itoa(packageData.ID),
		Name: packageData.Name,
		Description: JSONAllFieldsMarshaller{
			Value: model.PackageDescription{
				ID:        strconv.Itoa(packageData.ID),
				Name:      packageData.Name,
				URL:       packageData.URL,
				CreatedAt: packageData.CreatedAt,
				UpdatedAt: packageData.UpdatedAt,
			},
		},
	}
	if stream != nil {
		if err := (*stream)(value); err != nil {
			return nil, err
		}
	}

	return &value, nil
}
