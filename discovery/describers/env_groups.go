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

func ListEnvGroups(ctx context.Context, handler *provider.RenderAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	renderChan := make(chan models.Resource)
	errorChan := make(chan error, 1) // Buffered channel to capture errors

	go func() {
		defer close(renderChan)
		defer close(errorChan)
		if err := processEnvGroups(ctx, handler, renderChan, &wg); err != nil {
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

func GetEnvGroup(ctx context.Context, handler *provider.RenderAPIHandler, resourceID string) (*models.Resource, error) {
	envGroup, err := processEnvGroup(ctx, handler, resourceID)
	if err != nil {
		return nil, err
	}
	var serviceLinks []provider.ServiceLink
	for _, serviceLink := range envGroup.ServiceLinks {
		serviceLinks = append(serviceLinks, provider.ServiceLink{
			ID:   serviceLink.ID,
			Name: serviceLink.Name,
			Type: serviceLink.Type,
		})
	}
	value := models.Resource{
		ID:   envGroup.ID,
		Name: envGroup.Name,
		Description: JSONAllFieldsMarshaller{
			Value: provider.EnvGroupDescription{
				ID:            envGroup.ID,
				Name:          envGroup.Name,
				OwnerID:       envGroup.OwnerID,
				CreatedAt:     envGroup.CreatedAt,
				UpdatedAt:     envGroup.UpdatedAt,
				ServiceLinks:  serviceLinks,
				EnvironmentID: envGroup.EnvironmentID,
			},
		},
	}
	return &value, nil
}

func processEnvGroups(ctx context.Context, handler *provider.RenderAPIHandler, renderChan chan<- models.Resource, wg *sync.WaitGroup) error {
	var envGroups []provider.EnvGroupJSON
	var envGroupResp []provider.EnvGroupResponse
	var resp *http.Response
	baseURL := "https://api.render.com/v1/env-groups"
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

			if e = json.NewDecoder(resp.Body).Decode(&envGroupResp); e != nil {
				return nil, fmt.Errorf("failed to decode response: %w", e)
			}
			for i, envGroup := range envGroupResp {
				envGroups = append(envGroups, envGroup.EnvGroup)
				if i == len(envGroupResp)-1 {
					cursor = envGroup.Cursor
				}
			}
			return resp, nil
		}
		err = handler.DoRequest(ctx, req, requestFunc)
		if err != nil {
			return fmt.Errorf("error during request handling: %w", err)
		}

		if len(envGroupResp) < 100 {
			break
		}
	}
	for _, envGroup := range envGroups {
		wg.Add(1)
		go func(envGroup provider.EnvGroupJSON) {
			defer wg.Done()
			var serviceLinks []provider.ServiceLink
			for _, serviceLink := range envGroup.ServiceLinks {
				serviceLinks = append(serviceLinks, provider.ServiceLink{
					ID:   serviceLink.ID,
					Name: serviceLink.Name,
					Type: serviceLink.Type,
				})
			}
			value := models.Resource{
				ID:   envGroup.ID,
				Name: envGroup.Name,
				Description: JSONAllFieldsMarshaller{
					Value: provider.EnvGroupDescription{
						ID:            envGroup.ID,
						Name:          envGroup.Name,
						OwnerID:       envGroup.OwnerID,
						CreatedAt:     envGroup.CreatedAt,
						UpdatedAt:     envGroup.UpdatedAt,
						ServiceLinks:  serviceLinks,
						EnvironmentID: envGroup.EnvironmentID,
					},
				},
			}
			renderChan <- value
		}(envGroup)
	}
	return nil
}

func processEnvGroup(ctx context.Context, handler *provider.RenderAPIHandler, resourceID string) (*provider.EnvGroupJSON, error) {
	var envGroup provider.EnvGroupJSON
	var resp *http.Response
	baseURL := "https://api.render.com/v1/env-groups/"

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

		if e = json.NewDecoder(resp.Body).Decode(&envGroup); e != nil {
			return nil, fmt.Errorf("failed to decode response: %w", e)
		}
		return resp, e
	}

	err = handler.DoRequest(ctx, req, requestFunc)
	if err != nil {
		return nil, fmt.Errorf("error during request handling: %w", err)
	}
	return &envGroup, nil
}
