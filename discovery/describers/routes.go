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

func ListRoutes(ctx context.Context, handler *provider.RenderAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
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
			if service.Type == "static_site" {
				if err := processRoutes(ctx, handler, service.ID, renderChan, &wg); err != nil {
					errorChan <- err // Send error to the error channel
				}
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

func processRoutes(ctx context.Context, handler *provider.RenderAPIHandler, serviceID string, renderChan chan<- models.Resource, wg *sync.WaitGroup) error {
	var routes []provider.RouteJSON
	var routeListResponse []provider.RouteResponse
	var resp *http.Response
	baseURL := "https://api.render.com/v1/services/"
	cursor := ""

	for {
		params := url.Values{}
		params.Set("limit", provider.Limit)
		if cursor != "" {
			params.Set("cursor", cursor)
		}
		finalURL := fmt.Sprintf("%s%s/routes?%s", baseURL, serviceID, params.Encode())

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

			if e = json.NewDecoder(resp.Body).Decode(&routeListResponse); e != nil {
				return nil, fmt.Errorf("failed to decode response: %w", e)
			}
			for i, routeResp := range routeListResponse {
				routes = append(routes, routeResp.Route)
				if i == len(routeListResponse)-1 {
					cursor = routeResp.Cursor
				}
			}
			return resp, nil
		}
		err = handler.DoRequest(ctx, req, requestFunc)
		if err != nil {
			return fmt.Errorf("error during request handling: %w", err)
		}

		if len(routeListResponse) < 100 {
			break
		}
	}
	for _, route := range routes {
		wg.Add(1)
		go func(route provider.RouteJSON) {
			defer wg.Done()
			value := models.Resource{
				ID:   route.ID,
				Name: route.Type,
				Description: JSONAllFieldsMarshaller{
					Value: provider.RouteDescription{
						ID:          route.ID,
						Type:        route.Type,
						Source:      route.Source,
						Destination: route.Destination,
						Priority:    route.Priority,
					},
				},
			}
			renderChan <- value
		}(route)
	}
	return nil
}
