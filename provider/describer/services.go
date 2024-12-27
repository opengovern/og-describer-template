package describer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/opengovern/og-describer-render/pkg/sdk/models"
	"github.com/opengovern/og-describer-render/provider/model"
	"net/http"
	"net/url"
	"sync"
)

func ListServices(ctx context.Context, handler *RenderAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	renderChan := make(chan models.Resource)
	errorChan := make(chan error, 1) // Buffered channel to capture errors

	go func() {
		defer close(renderChan)
		defer close(errorChan)
		if err := processServices(ctx, handler, renderChan, &wg); err != nil {
			errorChan <- err // Send error to the error channel
		}
		wg.Wait()
	}()

	var values []models.Resource
	for {
		select {
		case value, ok := <-renderChan:
			if !ok {
				return values, nil
			}
			if stream != nil {
				if err := (*stream)(value); err != nil {
					return nil, err
				}
			} else {
				values = append(values, value)
			}
		case err := <-errorChan:
			return nil, err
		}
	}
}

func GetService(ctx context.Context, handler *RenderAPIHandler, resourceID string) (*models.Resource, error) {
	service, err := processService(ctx, handler, resourceID)
	if err != nil {
		return nil, err
	}
	buildFilter := model.BuildFilter{
		Paths:        service.BuildFilter.Paths,
		IgnoredPaths: service.BuildFilter.IgnoredPaths,
	}
	registryCredential := model.RegistryCredential{
		ID:   service.RegistryCredential.ID,
		Name: service.RegistryCredential.Name,
	}
	parentServer := model.ParentServer{
		ID:   service.ServiceDetails.ParentServer.ID,
		Name: service.ServiceDetails.ParentServer.Name,
	}
	previews := model.Previews{
		Generation: service.ServiceDetails.Previews.Generation,
	}
	serviceDetails := model.ServiceDetails{
		BuildCommand: service.ServiceDetails.BuildCommand,
		ParentServer: parentServer,
		PublishPath:  service.ServiceDetails.PublishPath,
		Previews:     previews,
		URL:          service.ServiceDetails.URL,
		BuildPlan:    service.ServiceDetails.BuildPlan,
	}
	value := models.Resource{
		ID:   service.ID,
		Name: service.Name,
		Description: JSONAllFieldsMarshaller{
			Value: model.ServiceDescription{
				ID:                 service.ID,
				AutoDeploy:         service.AutoDeploy,
				Branch:             service.Branch,
				BuildFilter:        buildFilter,
				CreatedAt:          service.CreatedAt,
				DashboardURL:       service.DashboardURL,
				EnvironmentID:      service.EnvironmentID,
				ImagePath:          service.ImagePath,
				Name:               service.Name,
				NotifyOnFail:       service.NotifyOnFail,
				OwnerID:            service.OwnerID,
				RegistryCredential: registryCredential,
				Repo:               service.Repo,
				RootDir:            service.RootDir,
				Slug:               service.Slug,
				Suspended:          service.Suspended,
				Suspenders:         service.Suspenders,
				Type:               service.Type,
				UpdatedAt:          service.UpdatedAt,
				ServiceDetails:     serviceDetails,
			},
		},
	}
	return &value, nil
}

func processServices(ctx context.Context, handler *RenderAPIHandler, renderChan chan<- models.Resource, wg *sync.WaitGroup) error {
	var services []model.ServiceJSON
	var serviceListResponse []model.ServiceResponse
	var resp *http.Response
	baseURL := "https://api.render.com/v1/services"
	cursor := ""

	for {
		params := url.Values{}
		params.Set("limit", limit)
		params.Set("includePreviews", includePreviews)
		if cursor != "" {
			params.Set("cursor", cursor)
		}
		finalURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

		req, err := http.NewRequest("GET", finalURL, nil)
		if err != nil {
			return fmt.Errorf("failed to create request: %w", err)
		}

		requestFunc := func(req *http.Request) (*http.Response, error) {
			var e error
			resp, e = handler.Client.Do(req)
			if e != nil {
				return nil, fmt.Errorf("request execution failed: %w", e)
			}
			defer resp.Body.Close()

			if e = json.NewDecoder(resp.Body).Decode(&serviceListResponse); e != nil {
				return nil, fmt.Errorf("failed to decode response: %w", e)
			}
			for i, serviceResp := range serviceListResponse {
				services = append(services, serviceResp.Service)
				if i == len(serviceListResponse)-1 {
					cursor = serviceResp.Cursor
				}
			}
			return resp, nil
		}
		err = handler.DoRequest(ctx, req, requestFunc)
		if err != nil {
			return fmt.Errorf("error during request handling: %w", err)
		}

		if len(serviceListResponse) < 100 {
			break
		}
	}
	for _, service := range services {
		wg.Add(1)
		go func(service model.ServiceJSON) {
			defer wg.Done()
			buildFilter := model.BuildFilter{
				Paths:        service.BuildFilter.Paths,
				IgnoredPaths: service.BuildFilter.IgnoredPaths,
			}
			registryCredential := model.RegistryCredential{
				ID:   service.RegistryCredential.ID,
				Name: service.RegistryCredential.Name,
			}
			parentServer := model.ParentServer{
				ID:   service.ServiceDetails.ParentServer.ID,
				Name: service.ServiceDetails.ParentServer.Name,
			}
			previews := model.Previews{
				Generation: service.ServiceDetails.Previews.Generation,
			}
			serviceDetails := model.ServiceDetails{
				BuildCommand: service.ServiceDetails.BuildCommand,
				ParentServer: parentServer,
				PublishPath:  service.ServiceDetails.PublishPath,
				Previews:     previews,
				URL:          service.ServiceDetails.URL,
				BuildPlan:    service.ServiceDetails.BuildPlan,
			}
			value := models.Resource{
				ID:   service.ID,
				Name: service.Name,
				Description: JSONAllFieldsMarshaller{
					Value: model.ServiceDescription{
						ID:                 service.ID,
						AutoDeploy:         service.AutoDeploy,
						Branch:             service.Branch,
						BuildFilter:        buildFilter,
						CreatedAt:          service.CreatedAt,
						DashboardURL:       service.DashboardURL,
						EnvironmentID:      service.EnvironmentID,
						ImagePath:          service.ImagePath,
						Name:               service.Name,
						NotifyOnFail:       service.NotifyOnFail,
						OwnerID:            service.OwnerID,
						RegistryCredential: registryCredential,
						Repo:               service.Repo,
						RootDir:            service.RootDir,
						Slug:               service.Slug,
						Suspended:          service.Suspended,
						Suspenders:         service.Suspenders,
						Type:               service.Type,
						UpdatedAt:          service.UpdatedAt,
						ServiceDetails:     serviceDetails,
					},
				},
			}
			renderChan <- value
		}(service)
	}
	return nil
}

func processService(ctx context.Context, handler *RenderAPIHandler, resourceID string) (*model.ServiceJSON, error) {
	var service model.ServiceJSON
	var resp *http.Response
	baseURL := "https://api.render.com/v1/services/"

	finalURL := fmt.Sprintf("%s%s", baseURL, resourceID)
	req, err := http.NewRequest("GET", finalURL, nil)
	if err != nil {
		return nil, err
	}

	requestFunc := func(req *http.Request) (*http.Response, error) {
		var e error
		resp, e = handler.Client.Do(req)
		if e != nil {
			return nil, fmt.Errorf("request execution failed: %w", e)
		}
		defer resp.Body.Close()

		if e = json.NewDecoder(resp.Body).Decode(&service); e != nil {
			return nil, fmt.Errorf("failed to decode response: %w", e)
		}
		return resp, e
	}

	err = handler.DoRequest(ctx, req, requestFunc)
	if err != nil {
		return nil, fmt.Errorf("error during request handling: %w", err)
	}
	return &service, nil
}
