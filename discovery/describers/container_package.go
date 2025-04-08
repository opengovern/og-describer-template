package describers

import (
	"context"
	"encoding/json"
	"errors" // Import errors package
	"fmt"
	"log"
	"net/url"
	"strconv"
	"strings"
	//"sync" // No longer needed here

	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/name"
	"github.com/google/go-containerregistry/pkg/v1/remote"
	"github.com/opengovern/og-describer-github/discovery/pkg/models"
	model "github.com/opengovern/og-describer-github/discovery/provider"
	resilientbridge "github.com/opengovern/resilient-bridge"
	"github.com/opengovern/resilient-bridge/adapters"
)

// -----------------------------------------------------------------------------
// 1. GetContainerPackageList
// -----------------------------------------------------------------------------
func GetContainerPackageList(
	ctx context.Context, // Use the passed-in context
	githubClient model.GitHubClient,
	organizationName string,
	stream *models.StreamSender,
) ([]models.Resource, error) {
	sdk := resilientbridge.NewResilientBridge()
	sdk.RegisterProvider("github", adapters.NewGitHubAdapter(githubClient.Token), &resilientbridge.ProviderConfig{
		UseProviderLimits: true,
		MaxRetries:        3,
		BaseBackoff:       0,
	})

	org := organizationName

	organization := ctx.Value("organization")
	if organization != nil {
		orgVal, ok := organization.(string)
		if ok && orgVal != "" {
			org = orgVal
		}
	}

	packages := fetchPackages(sdk, org, "container")

	maxVersions := 1 // Still limiting to 1 version per package for now
	var allValues []models.Resource

	fmt.Println("packages:", len(packages)) // Removed packages list from log

	// Loop through each package and version
	for _, p := range packages {
		packageName := p.Name

		versions := fetchVersions(sdk, org, "container", packageName)

		if len(versions) > maxVersions {
			versions = versions[:maxVersions]
		}

		for _, v := range versions {

			// Pass the context down
			packageValues, err := getVersionOutput(ctx, githubClient.Token, org, packageName, v)

			if err != nil {
				if errors.Is(err, context.DeadlineExceeded) {
					log.Printf("Timeout processing version %s/%s: %v", packageName, v.Name, err)
				} else if errors.Is(err, context.Canceled) {
					log.Printf("Cancelled processing version %s/%s: %v", packageName, v.Name, err)
					return allValues, err // If context is cancelled, stop processing
				} else {
					log.Printf("Error getting version output for %s/%s: %v", packageName, v.Name, err)
				}
				continue
			}

			if packageValues.ID == "" && packageValues.Name == "" {
				log.Printf("Skipping empty result for %s/%s", packageName, v.Name)
				continue
			}

			if stream != nil {
				if e := (*stream)(packageValues); e != nil {
					log.Printf("Error streaming result for %s/%s: %v", packageName, v.Name, e)
					return allValues, e
				}
			}

			allValues = append(allValues, packageValues)
		}
	}
	return allValues, nil
}

// -----------------------------------------------------------------------------
// 2. getVersionOutput - Simplified, removed broken concurrency, added context
// -----------------------------------------------------------------------------
func getVersionOutput(
	ctx context.Context,
	apiToken, org, packageName string,
	version model.PackageVersion,
) (models.Resource, error) {

	var emptyResult models.Resource

	normalizedPackageName := strings.ToLower(packageName)
	tags := version.Metadata.Container.Tags
	if len(tags) == 0 {
		log.Printf("No tags found for version %s of package %s, skipping", version.Name, packageName)
		return emptyResult, nil
	}

	image_uri := make([]string, 0, len(tags))
	for _, tag := range tags {
		normalizedTag := strings.ToLower(tag)
		image_uri = append(image_uri, fmt.Sprintf("ghcr.io/%s/%s:%s", org, normalizedPackageName, normalizedTag))
	}

	parts := strings.Split(version.Name, ":")
	if len(parts) != 2 || parts[0] != "sha256" {
		log.Printf("Warning: Unexpected format for version name '%s' for package %s. Expected 'sha256:<digest>'.", version.Name, packageName)
	}
	digestPart := ""
	if len(parts) > 1 {
		digestPart = parts[1]
	} else {
		log.Printf("Warning: Version name '%s' for package %s did not contain ':'. Attempting to use directly.", version.Name, packageName)
		if !strings.Contains(version.Name, ":") {
			log.Printf("Cannot reliably determine digest prefix for version name '%s'. Skipping.", version.Name)
			return emptyResult, fmt.Errorf("cannot determine digest from version name '%s'", version.Name)
		}
	}

	if digestPart == "" {
		log.Printf("Error: Could not extract digest from version name '%s' for package %s.", version.Name, packageName)
		return emptyResult, fmt.Errorf("could not extract digest from version name '%s'", version.Name)
	}

	imageRef := fmt.Sprintf("ghcr.io/%s/%s@sha256:%s", org, normalizedPackageName, digestPart)

	ov, err := fetchAndAssembleOutput(ctx, apiToken, org, normalizedPackageName, version, imageRef, image_uri)
	if err != nil {
		return emptyResult, fmt.Errorf("failed in fetchAndAssembleOutput for %s: %w", imageRef, err)
	}

	value := models.Resource{
		ID:          strconv.Itoa(ov.ID),
		Name:        ov.Name,
		Description: ov,
	}
	return value, nil
}

// -----------------------------------------------------------------------------
// 4. fetchAndAssembleOutput - Added context parameter and usage
// -----------------------------------------------------------------------------
func fetchAndAssembleOutput(
	ctx context.Context, // Added context parameter
	apiToken, org, packageName string,
	version model.PackageVersion,
	imageRef string,
	imageURIs []string,
) (model.ContainerPackageDescription, error) {

	var emptyResult model.ContainerPackageDescription // Return this on error

	authOption := remote.WithAuth(&authn.Basic{
		Username: "github", // Consider making username configurable if not always github
		Password: apiToken,
	})
	imageRef = strings.ToLower(imageRef)

	// Add context option for the remote call
	ctxOption := remote.WithContext(ctx)

	ref, err := name.ParseReference(imageRef, name.WeakValidation)
	if err != nil {
		return emptyResult,
			fmt.Errorf("error parsing reference %s: %w", imageRef, err)
	}

	// Use context with remote.Get
	desc, err := remote.Get(ref, authOption, ctxOption) // Added ctxOption
	if err != nil {
		// Check context errors specifically
		if errors.Is(err, context.DeadlineExceeded) {
			return emptyResult, fmt.Errorf("timeout fetching manifest for %s: %w", imageRef, err)
		}
		if errors.Is(err, context.Canceled) {
			// Don't log cancellation as an error usually, just return it
			return emptyResult, fmt.Errorf("cancelled fetching manifest for %s: %w", imageRef, err)
		}
		// Log other errors before returning
		log.Printf("Error response from remote.Get for %s: %v", imageRef, err)
		return emptyResult, fmt.Errorf("error fetching manifest for %s: %w", imageRef, err)
	}

	actualDigest := desc.Descriptor.Digest.String()

	// Unmarshal manifest to calculate size
	var manifestStruct struct {
		SchemaVersion int    `json:"schemaVersion"`
		MediaType     string `json:"mediaType"`
		Config        struct {
			Size   int64  `json:"size"`
			Digest string `json:"digest"`
			//MediaType string `json:"mediaType"` // Keep if needed, commented out if unused
		} `json:"config"`
		Layers []struct {
			Size   int64  `json:"size"`
			Digest string `json:"digest"`
			//MediaType string `json:"mediaType"` // Keep if needed
		} `json:"layers"`
	}
	// Use RawManifest if Manifest is deprecated or causes issues
	manifestBytes := desc.Manifest // Or desc.RawManifest if applicable
	if manifestBytes == nil {
		return emptyResult, fmt.Errorf("manifest content is nil for %s", imageRef)
	}

	if err := json.Unmarshal(manifestBytes, &manifestStruct); err != nil {
		// Log the manifest content if unmarshaling fails for debugging
		log.Printf("Manifest content for %s: %s", imageRef, string(manifestBytes))
		return emptyResult, fmt.Errorf("error unmarshaling manifest JSON for size calculation for %s: %w", imageRef, err)
	}

	totalSize := manifestStruct.Config.Size
	for _, layer := range manifestStruct.Layers {
		totalSize += layer.Size
	}

	// Unmarshal manifest again for storing as interface{}
	var manifestInterface interface{}
	if err := json.Unmarshal(manifestBytes, &manifestInterface); err != nil {
		// This shouldn't fail if the previous one succeeded, but check anyway
		return emptyResult, fmt.Errorf("error unmarshaling manifest for output interface for %s: %w", imageRef, err)
	}

	ov := model.ContainerPackageDescription{
		ID:                    version.ID,
		Digest:                actualDigest, // Use the actual digest from registry
		AdditionalPackageURIs: []string{},   // Initialize, will be populated by deduplication if used
		CreatedAt:             version.CreatedAt,
		UpdatedAt:             version.UpdatedAt,
		PackageURL:            version.HTMLURL, // This is the package version URL, might differ from specific tag URLs
		Name:                  packageName,     // The package name
		ImageUri:              imageURIs,       // List of tag-based URIs generated earlier
		ImageRef:              imageRef,        // The specific digest-based ref used for fetching

		MediaType:    string(desc.Descriptor.MediaType), // Use MediaType from descriptor
		TotalSize:    totalSize,
		Metadata:     version.Metadata, // Contains original tags
		Manifest:     manifestInterface,
		Organization: org,
		// GHVersionName: version.Name, // Add this field to model.ContainerPackageDescription if you want to store the original version name separately
	}

	//fmt.Println("ov:", ov) // Optional: uncomment for debugging fetch result
	return ov, nil
}

// -----------------------------------------------------------------------------
// 5. fetchPackages - No changes needed here
// -----------------------------------------------------------------------------
func fetchPackages(sdk *resilientbridge.ResilientBridge, org, packageType string) []model.Package {
	var allPackages []model.Package
	page := 1
	perPage := 100 // Max allowed by GitHub API

	for {
		req := &resilientbridge.NormalizedRequest{
			Method: "GET",
			Endpoint: fmt.Sprintf("/orgs/%s/packages?package_type=%s&per_page=%d&page=%d",
				org, packageType, perPage, page),
			Headers: map[string]string{"Accept": "application/vnd.github+json"},
		}

		resp, err := sdk.Request("github", req)
		if err != nil {
			// Log fatal stops execution, maybe just log error and return partial/empty list?
			log.Printf("Error requesting packages page %d: %v", page, err)
			break // Stop pagination on SDK error
		}
		// Check for HTTP errors specifically
		if resp.StatusCode == 404 {
			log.Printf("Organization '%s' not found or packages not enabled/visible.", org)
			break // Org not found, no point retrying/continuing
		} else if resp.StatusCode == 403 {
			log.Printf("Forbidden access to packages for org '%s'. Check token permissions (needs read:packages).", org)
			break // Permissions issue
		} else if resp.StatusCode >= 400 {
			log.Printf("HTTP error %d fetching packages page %d: %s", resp.StatusCode, page, string(resp.Data))
			break // Stop pagination on other HTTP error
		}

		var packages []model.Package
		if err := json.Unmarshal(resp.Data, &packages); err != nil {
			log.Printf("Error parsing packages list response page %d: %v", page, err)
			break // Stop pagination on parsing error
		}
		if len(packages) == 0 {
			// No more packages found on this page
			break
		}

		allPackages = append(allPackages, packages...)

		// Check if the number of packages received is less than requested perPage
		if len(packages) < perPage {
			// This was the last page
			break
		}
		page++
	}
	log.Printf("Fetched %d total container packages for org %s", len(allPackages), org)
	return allPackages
}

// -----------------------------------------------------------------------------
// 6. fetchVersions - No changes needed here
// -----------------------------------------------------------------------------
func fetchVersions(
	sdk *resilientbridge.ResilientBridge,
	org, packageType, packageName string,
) []model.PackageVersion {

	packageNameEncoded := url.PathEscape(packageName)
	var allVersions []model.PackageVersion
	page := 1
	perPage := 100 // Max allowed by GitHub API

	for {
		req := &resilientbridge.NormalizedRequest{
			Method: "GET",
			Endpoint: fmt.Sprintf(
				"/orgs/%s/packages/%s/%s/versions?per_page=%d&page=%d",
				org, packageType, packageNameEncoded, perPage, page,
			),
			Headers: map[string]string{"Accept": "application/vnd.github+json"},
		}

		resp, err := sdk.Request("github", req)
		if err != nil {
			log.Printf("Error requesting versions for package '%s' page %d: %v", packageName, page, err)
			break // Stop pagination on SDK error
		}

		// Check for HTTP errors specifically
		if resp.StatusCode == 404 {
			log.Printf("Package '%s' not found or no versions available in org '%s'.", packageName, org)
			break // Package not found
		} else if resp.StatusCode == 403 {
			log.Printf("Forbidden access to versions for package '%s' in org '%s'. Check token permissions.", packageName, org)
			break // Permissions issue
		} else if resp.StatusCode >= 400 {
			log.Printf("HTTP error %d fetching versions for package '%s' page %d: %s", resp.StatusCode, packageName, page, string(resp.Data))
			break // Stop pagination on other HTTP error
		}

		var versions []model.PackageVersion
		if err := json.Unmarshal(resp.Data, &versions); err != nil {
			log.Printf("Error parsing package versions response for package '%s' page %d: %v", packageName, page, err)
			break // Stop pagination on parsing error
		}

		if len(versions) == 0 {
			// No more versions found on this page
			break
		}

		allVersions = append(allVersions, versions...)

		// Check if the number of versions received is less than requested perPage
		if len(versions) < perPage {
			// This was the last page
			break
		}
		page++
	}
	log.Printf("Fetched %d total versions for container package %s/%s", len(allVersions), org, packageName)
	return allVersions
}
