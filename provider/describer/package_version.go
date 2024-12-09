package describer

import (
	"context"
	"fmt"
	"github.com/google/go-github/v55/github"
	"github.com/opengovern/og-describer-github/pkg/sdk/models"
	"github.com/opengovern/og-describer-github/provider/model"
	"strconv"
)

func GetAllPackageVersionList(ctx context.Context, githubClient GitHubClient, organizationName string, stream *models.StreamSender) ([]models.Resource, error) {
	packages, err := getPackages(ctx, githubClient, organizationName)
	if err != nil {
		return nil, err
	}
	var values []models.Resource
	for _, packageData := range packages {
		packageVersionValues, err := GetPackageVersionList(ctx, githubClient, organizationName, packageData.GetPackageType(), packageData.GetName(), int(packageData.GetID()), stream)
		if err != nil {
			return nil, err
		}
		values = append(values, packageVersionValues...)
	}

	return values, nil
}

func GetPackageVersionList(ctx context.Context, githubClient GitHubClient, organizationName, packageType, packageName string, packageID int, stream *models.StreamSender) ([]models.Resource, error) {
	client := githubClient.RestClient
	var values []models.Resource
	page := 1
	for {
		var opts = &github.PackageListOptions{
			PackageType: &packageType,
			ListOptions: github.ListOptions{
				Page:    page,
				PerPage: packagePageSize,
			},
		}
		respPackageVersions, resp, err := client.Organizations.PackageGetAllVersions(ctx, organizationName, packageType, packageName, opts)
		if err != nil {
			return nil, err
		}
		for _, packageVersion := range respPackageVersions {
			value := models.Resource{
				ID:   strconv.Itoa(int(packageVersion.GetID())),
				Name: packageVersion.GetName(),
				Description: JSONAllFieldsMarshaller{
					Value: model.PackageVersionDescription{
						ID:          int(packageVersion.GetID()),
						Name:        fmt.Sprintf("%s:%s", packageName, packageVersion.GetVersion()),
						VersionURI:  fmt.Sprintf("ghcr.io/%s/%s:%s", organizationName, packageName, packageVersion.GetVersion()),
						PackageName: packageName,
						Digest:      packageVersion.Name,
						CreatedAt:   packageVersion.GetCreatedAt(),
						UpdatedAt:   packageVersion.GetUpdatedAt(),
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
		if resp.After == "" {
			break
		}
		opts.ListOptions.Page += 1
	}
	return values, nil
}
