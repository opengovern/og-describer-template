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

func ListDeploys(ctx context.Context, handler *provider.RenderAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	renderChan := make(chan models.Resource)
	errorChan := make(chan error, 1) // Buffered channel to capture errors
	services, err := provider.GetServices(ctx, handler)
	if err != nil {
		return nil, err
	}

	go func() {
		defer close(renderChan)
		defer close(errorChan)
		for _, service := range services {
			if err := processDeploys(ctx, handler, service.ID, renderChan, &wg); err != nil {
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

func processDeploys(ctx context.Context, handler *provider.RenderAPIHandler, serviceID string, renderChan chan<- models.Resource, wg *sync.WaitGroup) error {
	var deploys []provider.DeployJSON
	var deployListResponse []provider.DeployResponse
	var resp *http.Response
	baseURL := "https://api.render.com/v1/services/"
	cursor := ""

	for {
		params := url.Values{}
		params.Set("limit", provider.Limit)
		if cursor != "" {
			params.Set("cursor", cursor)
		}
		finalURL := fmt.Sprintf("%s%s/deploys?%s", baseURL, serviceID, params.Encode())
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

			if e = json.NewDecoder(resp.Body).Decode(&deployListResponse); e != nil {
				return nil, fmt.Errorf("failed to decode response: %w", e)
			}
			for i, deployResp := range deployListResponse {
				deploys = append(deploys, deployResp.Deploy)
				if i == len(deployListResponse)-1 {
					cursor = deployResp.Cursor
				}
			}
			return resp, nil
		}
		err = handler.DoRequest(ctx, req, requestFunc)
		if err != nil {
			return fmt.Errorf("error during request handling: %w", err)
		}

		if len(deployListResponse) < 100 {
			break
		}
	}
	for _, deploy := range deploys {
		wg.Add(1)
		go func(deploy provider.DeployJSON) {
			defer wg.Done()
			commit := provider.Commit{
				ID:        deploy.Commit.ID,
				Message:   deploy.Commit.Message,
				CreatedAt: deploy.Commit.CreatedAt,
			}
			image := provider.Image{
				Ref:                deploy.Image.Ref,
				SHA:                deploy.Image.SHA,
				RegistryCredential: deploy.Image.RegistryCredential,
			}
			value := models.Resource{
				ID:   deploy.ID,
				Name: deploy.Status,
				Description: JSONAllFieldsMarshaller{
					Value: provider.DeployDescription{
						ID:         deploy.ID,
						Commit:     commit,
						Image:      image,
						Status:     deploy.Status,
						Trigger:    deploy.Trigger,
						FinishedAt: deploy.FinishedAt,
						CreatedAt:  deploy.CreatedAt,
						UpdatedAt:  deploy.UpdatedAt,
					},
				},
			}
			renderChan <- value
		}(deploy)
	}
	return nil
}
