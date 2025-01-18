package describers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/opengovern/og-describer-render/discovery/pkg/models"
	"github.com/opengovern/og-describer-render/discovery/provider"
	"net/http"
	"net/url"
	"sync"
)

func ListProjects(ctx context.Context, handler *provider.RenderAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	renderChan := make(chan models.Resource)
	errorChan := make(chan error, 1) // Buffered channel to capture errors

	go func() {
		defer close(renderChan)
		defer close(errorChan)
		if err := processProjects(ctx, handler, renderChan, &wg); err != nil {
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

func GetProject(ctx context.Context, handler *provider.RenderAPIHandler, resourceID string) (*models.Resource, error) {
	project, err := processProject(ctx, handler, resourceID)
	if err != nil {
		return nil, err
	}
	owner := provider.Owner{
		ID:                   project.Owner.ID,
		Name:                 project.Owner.Name,
		Email:                project.Owner.Email,
		TwoFactorAuthEnabled: project.Owner.TwoFactorAuthEnabled,
		Type:                 project.Owner.Type,
	}
	value := models.Resource{
		ID:   project.ID,
		Name: project.Name,
		Description: JSONAllFieldsMarshaller{
			Value: provider.ProjectDescription{
				ID:             project.ID,
				CreatedAt:      project.CreatedAt,
				UpdatedAt:      project.UpdatedAt,
				Name:           project.Name,
				Owner:          owner,
				EnvironmentIDs: project.EnvironmentIDs,
			},
		},
	}
	return &value, nil
}

func processProjects(ctx context.Context, handler *provider.RenderAPIHandler, renderChan chan<- models.Resource, wg *sync.WaitGroup) error {
	var projects []provider.ProjectJSON
	var projectListResponse []provider.ProjectResponse
	var resp *http.Response
	baseURL := "https://api.render.com/v1/projects"
	cursor := ""

	for {
		params := url.Values{}
		params.Set("limit", provider.Limit)
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

			if e = json.NewDecoder(resp.Body).Decode(&projectListResponse); e != nil {
				return nil, fmt.Errorf("failed to decode response: %w", e)
			}
			for i, projectResp := range projectListResponse {
				projects = append(projects, projectResp.Project)
				if i == len(projectListResponse)-1 {
					cursor = projectResp.Cursor
				}
			}
			return resp, nil
		}
		err = handler.DoRequest(ctx, req, requestFunc)
		if err != nil {
			return fmt.Errorf("error during request handling: %w", err)
		}

		if len(projectListResponse) < 100 {
			break
		}
	}
	for _, project := range projects {
		wg.Add(1)
		go func(project provider.ProjectJSON) {
			defer wg.Done()
			owner := provider.Owner{
				ID:                   project.Owner.ID,
				Name:                 project.Owner.Name,
				Email:                project.Owner.Email,
				TwoFactorAuthEnabled: project.Owner.TwoFactorAuthEnabled,
				Type:                 project.Owner.Type,
			}
			value := models.Resource{
				ID:   project.ID,
				Name: project.Name,
				Description: JSONAllFieldsMarshaller{
					Value: provider.ProjectDescription{
						ID:             project.ID,
						CreatedAt:      project.CreatedAt,
						UpdatedAt:      project.UpdatedAt,
						Name:           project.Name,
						Owner:          owner,
						EnvironmentIDs: project.EnvironmentIDs,
					},
				},
			}
			renderChan <- value
		}(project)
	}
	return nil
}

func processProject(ctx context.Context, handler *provider.RenderAPIHandler, resourceID string) (*provider.ProjectJSON, error) {
	var project provider.ProjectJSON
	var resp *http.Response
	baseURL := "https://api.render.com/v1/projects/"

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

		if e = json.NewDecoder(resp.Body).Decode(&project); e != nil {
			return nil, fmt.Errorf("failed to decode response: %w", e)
		}
		return resp, e
	}

	err = handler.DoRequest(ctx, req, requestFunc)
	if err != nil {
		return nil, fmt.Errorf("error during request handling: %w", err)
	}
	return &project, nil
}
