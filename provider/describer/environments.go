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

func ListEnvironments(ctx context.Context, handler *RenderAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	renderChan := make(chan models.Resource)
	errorChan := make(chan error, 1) // Buffered channel to capture errors
	projects, err := getProjects(ctx, handler)
	if err != nil {
		return nil, err
	}

	go func() {
		defer close(renderChan)
		defer close(errorChan)
		for _, project := range projects {
			if err := processEnvironments(ctx, handler, project.ID, renderChan, &wg); err != nil {
				errorChan <- err // Send error to the error channel
			}
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

func GetEnvironment(ctx context.Context, handler *RenderAPIHandler, resourceID string) (*models.Resource, error) {
	environment, err := processEnvironment(ctx, handler, resourceID)
	if err != nil {
		return nil, err
	}
	value := models.Resource{
		ID:   environment.ID,
		Name: environment.Name,
		Description: JSONAllFieldsMarshaller{
			Value: model.EnvironmentDescription{
				ID:              environment.ID,
				Name:            environment.Name,
				ProjectID:       resourceID,
				DatabasesIDs:    environment.DatabasesIDs,
				RedisIDs:        environment.RedisIDs,
				ServiceIDs:      environment.ServiceIDs,
				EnvGroupIDs:     environment.EnvGroupIDs,
				ProtectedStatus: environment.ProtectedStatus,
			},
		},
	}
	return &value, nil
}

func processEnvironments(ctx context.Context, handler *RenderAPIHandler, projectID string, renderChan chan<- models.Resource, wg *sync.WaitGroup) error {
	var environments []model.EnvironmentJSON
	var environmentListResponse []model.EnvironmentResponse
	var resp *http.Response
	baseURL := "https://api.render.com/v1/environments"
	cursor := ""

	for {
		params := url.Values{}
		params.Set("projectId", projectID)
		params.Set("limit", limit)
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

			if e = json.NewDecoder(resp.Body).Decode(&environmentListResponse); e != nil {
				return nil, fmt.Errorf("failed to decode response: %w", e)
			}
			for i, environmentResp := range environmentListResponse {
				environments = append(environments, environmentResp.Environment)
				if i == len(environmentListResponse)-1 {
					cursor = environmentResp.Cursor
				}
			}
			return resp, nil
		}
		err = handler.DoRequest(ctx, req, requestFunc)
		if err != nil {
			return fmt.Errorf("error during request handling: %w", err)
		}

		if len(environmentListResponse) < 100 {
			break
		}
	}
	for _, environment := range environments {
		wg.Add(1)
		go func(environment model.EnvironmentJSON) {
			defer wg.Done()
			value := models.Resource{
				ID:   environment.ID,
				Name: environment.Name,
				Description: JSONAllFieldsMarshaller{
					Value: model.EnvironmentDescription{
						ID:              environment.ID,
						Name:            environment.Name,
						ProjectID:       projectID,
						DatabasesIDs:    environment.DatabasesIDs,
						RedisIDs:        environment.RedisIDs,
						ServiceIDs:      environment.ServiceIDs,
						EnvGroupIDs:     environment.EnvGroupIDs,
						ProtectedStatus: environment.ProtectedStatus,
					},
				},
			}
			renderChan <- value
		}(environment)
	}
	return nil
}

func processEnvironment(ctx context.Context, handler *RenderAPIHandler, resourceID string) (*model.EnvironmentJSON, error) {
	var environment model.EnvironmentJSON
	var resp *http.Response
	baseURL := "https://api.render.com/v1/environments/"

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

		if e = json.NewDecoder(resp.Body).Decode(&environment); e != nil {
			return nil, fmt.Errorf("failed to decode response: %w", e)
		}
		return resp, e
	}

	err = handler.DoRequest(ctx, req, requestFunc)
	if err != nil {
		return nil, fmt.Errorf("error during request handling: %w", err)
	}
	return &environment, nil
}
