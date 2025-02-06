package describers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/opengovern/og-describer-render/discovery/pkg/models"
	"github.com/opengovern/og-describer-render/discovery/provider"
	"net/http"
	"sync"
)

func ListPostgresqlBackups(ctx context.Context, handler *provider.RenderAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	renderChan := make(chan models.Resource)
	errorChan := make(chan error, 1) // Buffered channel to capture errors
	postgresInstances, err := provider.GetPostgresqlInstances(ctx, handler)
	if err != nil {
		return nil, err
	}

	go func() {
		defer close(renderChan)
		defer close(errorChan)
		for _, postgres := range postgresInstances {
			if err := processPostgresqlBackups(ctx, handler, postgres.ID, renderChan, &wg); err != nil {
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

func processPostgresqlBackups(ctx context.Context, handler *provider.RenderAPIHandler, postgresID string, renderChan chan<- models.Resource, wg *sync.WaitGroup) error {
	var postgresqlBackups []provider.PostgresqlBackupJSON
	var resp *http.Response
	baseURL := "https://api.render.com/v1/postgres/"

	finalURL := fmt.Sprintf("%s%s/backup", baseURL, postgresID)

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

		if e = json.NewDecoder(resp.Body).Decode(&postgresqlBackups); e != nil {
			return nil, fmt.Errorf("failed to decode response: %w", e)
		}
		return resp, nil
	}
	err = handler.DoRequest(ctx, req, requestFunc)
	if err != nil {
		return fmt.Errorf("error during request handling: %w", err)
	}

	for _, postgresqlBackup := range postgresqlBackups {
		wg.Add(1)
		go func(postgresqlBackup provider.PostgresqlBackupJSON) {
			defer wg.Done()
			value := models.Resource{
				ID:   postgresqlBackup.ID,
				Name: postgresqlBackup.URL,
				Description: JSONAllFieldsMarshaller{
					Value: provider.PostgresqlBackupDescription{
						ID:        postgresqlBackup.ID,
						CreatedAt: postgresqlBackup.CreatedAt,
						URL:       postgresqlBackup.URL,
					},
				},
			}
			renderChan <- value
		}(postgresqlBackup)
	}
	return nil
}
