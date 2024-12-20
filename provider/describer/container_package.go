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
	"log"
	"net/url"
	"strconv"
	"strings"
)

func GetContainerPackageList(ctx context.Context, githubClient GitHubClient, organizationName string, stream *models.StreamSender) ([]models.Resource, error) {
	sdk := resilientbridge.NewResilientBridge()
	sdk.RegisterProvider("github", adapters.NewGitHubAdapter(githubClient.Token), &resilientbridge.ProviderConfig{
		UseProviderLimits: true,
		MaxRetries:        3,
		BaseBackoff:       0,
	})

	org := organizationName

	packages := fetchPackages(sdk, org, "container")

	maxVersions := 1

	var values []models.Resource

	for _, p := range packages {
		packageName := p.Name
		versions := fetchVersions(sdk, org, "container", packageName)
		if len(versions) > maxVersions {
			versions = versions[:maxVersions]
		}
		for _, v := range versions {
			packageValues := getVersionOutput(githubClient.Token, org, packageName, v, stream)
			values = append(values, packageValues...)
		}
	}

	return values, nil
}

func getVersionOutput(apiToken, org, packageName string, version model.PackageVersion, stream *models.StreamSender) []models.Resource {
	var values []models.Resource
	normalizedPackageName := strings.ToLower(packageName)

	for _, tag := range version.Metadata.Container.Tags {
		normalizedTag := strings.ToLower(tag)
		imageRef := fmt.Sprintf("ghcr.io/%s/%s:%s", org, normalizedPackageName, normalizedTag)
		ov := fetchAndAssembleOutput(apiToken, org, normalizedPackageName, version, imageRef)
		value := models.Resource{
			ID:   strconv.Itoa(ov.ID),
			Name: ov.Name,
			Description: JSONAllFieldsMarshaller{
				Value: ov,
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

func fetchAndAssembleOutput(apiToken, org, packageName string, version model.PackageVersion, imageRef string) model.ContainerPackageDescription {
	authOption := remote.WithAuth(&authn.Basic{
		Username: "github",
		Password: apiToken,
	})

	ref, err := name.ParseReference(imageRef)
	if err != nil {
		log.Fatalf("Error parsing reference %s: %v", imageRef, err)
	}

	desc, err := remote.Get(ref, authOption)
	if err != nil {
		log.Fatalf("Error fetching manifest for %s: %v", imageRef, err)
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
		log.Fatalf("Error unmarshaling manifest JSON: %v", err)
	}

	totalSize := manifestStruct.Config.Size
	for _, layer := range manifestStruct.Layers {
		totalSize += layer.Size
	}

	var manifestInterface interface{}
	if err := json.Unmarshal(desc.Manifest, &manifestInterface); err != nil {
		log.Fatalf("Error unmarshaling manifest for output: %v", err)
	}

	return model.ContainerPackageDescription{
		ID:             version.ID,
		Digest:         version.Name,
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
	}
}

func fetchPackages(sdk *resilientbridge.ResilientBridge, org, packageType string) []model.Package {
	listReq := &resilientbridge.NormalizedRequest{
		Method:   "GET",
		Endpoint: fmt.Sprintf("/orgs/%s/packages?package_type=%s", org, packageType),
		Headers:  map[string]string{"Accept": "application/vnd.github+json"},
	}
	listResp, err := sdk.Request("github", listReq)
	if err != nil {
		log.Fatalf("Error listing packages: %v", err)
	}
	if listResp.StatusCode >= 400 {
		log.Fatalf("HTTP error %d: %s", listResp.StatusCode, string(listResp.Data))
	}
	var packages []model.Package
	if err := json.Unmarshal(listResp.Data, &packages); err != nil {
		log.Fatalf("Error parsing packages list response: %v", err)
	}
	return packages
}

func fetchVersions(sdk *resilientbridge.ResilientBridge, org, packageType, packageName string) []model.PackageVersion {
	packageNameEncoded := url.PathEscape(packageName)
	versionsReq := &resilientbridge.NormalizedRequest{
		Method:   "GET",
		Endpoint: fmt.Sprintf("/orgs/%s/packages/%s/%s/versions", org, packageType, packageNameEncoded),
		Headers:  map[string]string{"Accept": "application/vnd.github+json"},
	}

	versionsResp, err := sdk.Request("github", versionsReq)
	if err != nil {
		log.Fatalf("Error listing package versions: %v", err)
	}
	if versionsResp.StatusCode >= 400 {
		log.Fatalf("HTTP error %d: %s", versionsResp.StatusCode, string(versionsResp.Data))
	}

	var versions []model.PackageVersion
	if err := json.Unmarshal(versionsResp.Data, &versions); err != nil {
		log.Fatalf("Error parsing package versions response: %v", err)
	}
	return versions
}
