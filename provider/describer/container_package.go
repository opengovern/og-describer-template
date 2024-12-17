package describer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/name"
	"github.com/google/go-containerregistry/pkg/v1/remote"
	"github.com/opengovern/og-describer-github/pkg/sdk/models"
	"github.com/opengovern/og-describer-github/provider/model"
	resilientbridge "github.com/opengovern/resilient-bridge"
	"github.com/opengovern/resilient-bridge/adapters"
	"strconv"
)

func GetContainerPackageList(ctx context.Context, githubClient GitHubClient, organizationName string, stream *models.StreamSender) ([]models.Resource, error) {
	sdk := resilientbridge.NewResilientBridge()
	sdk.RegisterProvider("github", adapters.NewGitHubAdapter(githubClient.Token), &resilientbridge.ProviderConfig{
		UseProviderLimits: true,
		MaxRetries:        3,
		BaseBackoff:       0,
	})
	packages, err := fetchPackages(sdk, organizationName, "container")
	if err != nil {
		return nil, err
	}
	var allResults []model.PackageOutputVersionDescription
	for _, p := range packages {
		packageName := p.Name
		versions, err := fetchVersions(sdk, organizationName, "container", packageName)
		if err != nil {
			return nil, err
		}
		for _, v := range versions {
			results, err := getVersionOutput(githubClient.Token, organizationName, packageName, v)
			if err != nil {
				return nil, err
			}
			allResults = append(allResults, results...)
		}
	}
	var values []models.Resource
	for _, result := range allResults {
		value := models.Resource{
			ID:   strconv.Itoa(result.ID),
			Name: result.Name,
			Description: JSONAllFieldsMarshaller{
				Value: result,
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

func GetContainerPackage(ctx context.Context, githubClient GitHubClient, organizationName string, repositoryName string, resourceID string, stream *models.StreamSender) (*models.Resource, error) {
	client := githubClient.RestClient
	packageType := "container"
	respPackages, _, err := client.Organizations.GetPackage(ctx, organizationName, packageType, resourceID)
	if err != nil {
		return nil, err
	}
	value := models.Resource{
		ID:   strconv.Itoa(int(respPackages.GetID())),
		Name: respPackages.GetName(),
		Description: JSONAllFieldsMarshaller{
			Value: model.PackageDescription{
				ID:         strconv.Itoa(int(respPackages.GetID())),
				RegistryID: respPackages.Registry.GetURL(),
				Name:       respPackages.GetName(),
				URL:        respPackages.GetURL(),
				CreatedAt:  respPackages.GetCreatedAt(),
				UpdatedAt:  respPackages.GetUpdatedAt(),
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

func getVersionOutput(apiToken, org, packageName string, version model.PackageVersion) ([]model.PackageOutputVersionDescription, error) {
	// Each version can have multiple tags. We'll produce one output object per tag.
	var results []model.PackageOutputVersionDescription
	for _, tag := range version.Metadata.Container.Tags {
		imageRef := fmt.Sprintf("ghcr.io/%s/%s:%s", org, packageName, tag)
		ov, err := fetchAndAssembleOutput(apiToken, version, imageRef)
		if err != nil {
			return nil, err
		}
		results = append(results, *ov)
	}
	return results, nil
}

func fetchAndAssembleOutput(apiToken string, version model.PackageVersion, imageRef string) (*model.PackageOutputVersionDescription, error) {
	authOption := remote.WithAuth(&authn.Basic{
		Username: "github",
		Password: apiToken,
	})

	ref, err := name.ParseReference(imageRef)
	if err != nil {
		return nil, err
	}

	desc, err := remote.Get(ref, authOption)
	if err != nil {
		return nil, err
	}

	var manifestStruct struct {
		SchemaVersion int    `json:"schemaVersion"`
		MediaType     string `json:"mediaType"`
		Config        struct {
			Size      int64  `json:"size"`
			Digest    string `json:"digest"`
			MediaType string `json:"mediaType"`
		} `json:"config"`
		Layers []struct {
			Size      int64  `json:"size"`
			Digest    string `json:"digest"`
			MediaType string `json:"mediaType"`
		} `json:"layers"`
	}
	if err := json.Unmarshal(desc.Manifest, &manifestStruct); err != nil {
		return nil, err
	}

	totalSize := manifestStruct.Config.Size
	for _, layer := range manifestStruct.Layers {
		totalSize += layer.Size
	}

	var manifestInterface interface{}
	if err := json.Unmarshal(desc.Manifest, &manifestInterface); err != nil {
		return nil, err
	}

	return &model.PackageOutputVersionDescription{
		ID:             version.ID,
		Digest:         version.Name, // version digest from "name"
		URL:            version.URL,
		PackageURI:     imageRef,
		PackageHTMLURL: version.PackageHTMLURL,
		CreatedAt:      version.CreatedAt,
		UpdatedAt:      version.UpdatedAt,
		HTMLURL:        version.HTMLURL,
		Name:           imageRef,
		MediaType:      string(desc.Descriptor.MediaType),
		TotalSize:      totalSize,
		Metadata:       version.Metadata,
		Manifest:       manifestInterface,
	}, nil
}

func fetchPackages(sdk *resilientbridge.ResilientBridge, org, packageType string) ([]model.Package, error) {
	listReq := &resilientbridge.NormalizedRequest{
		Method:   "GET",
		Endpoint: fmt.Sprintf("/orgs/%s/packages?package_type=%s", org, packageType),
		Headers:  map[string]string{"Accept": "application/vnd.github+json"},
	}
	listResp, err := sdk.Request("github", listReq)
	if err != nil {
		return nil, err
	}
	if listResp.StatusCode >= 400 {
		return nil, fmt.Errorf("HTTP error %d: %s", listResp.StatusCode, string(listResp.Data))
	}
	var packages []model.Package
	if err := json.Unmarshal(listResp.Data, &packages); err != nil {
		return nil, fmt.Errorf("error parsing packages list response: %v", err)
	}
	return packages, nil
}

func fetchVersions(sdk *resilientbridge.ResilientBridge, org, packageType, packageName string) ([]model.PackageVersion, error) {
	versionsReq := &resilientbridge.NormalizedRequest{
		Method:   "GET",
		Endpoint: fmt.Sprintf("/orgs/%s/packages/%s/%s/versions", org, packageType, packageName),
		Headers:  map[string]string{"Accept": "application/vnd.github+json"},
	}

	versionsResp, err := sdk.Request("github", versionsReq)
	if err != nil {
		return nil, err
	}
	if versionsResp.StatusCode >= 400 {
		return nil, fmt.Errorf("HTTP error %d: %s", versionsResp.StatusCode, string(versionsResp.Data))
	}

	var versions []model.PackageVersion
	if err := json.Unmarshal(versionsResp.Data, &versions); err != nil {
		return nil, fmt.Errorf("error parsing package versions response: %v", err)
	}
	return versions, nil
}
