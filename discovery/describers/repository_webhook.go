package describers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/opengovern/og-describer-github/discovery/pkg/models"
	model "github.com/opengovern/og-describer-github/discovery/provider"
	resilientbridge "github.com/opengovern/resilient-bridge"
	"github.com/opengovern/resilient-bridge/adapters"
	"net/url"
	"strconv"
	"time"
)

func ListRepositoryWebhooks(ctx context.Context, githubClient model.GitHubClient, organizationName string, stream *models.StreamSender) ([]models.Resource, error) {
	handler := resilientbridge.NewResilientBridge()
	handler.SetDebug(false)
	handler.RegisterProvider("github", adapters.NewGitHubAdapter(githubClient.Token), &resilientbridge.ProviderConfig{
		UseProviderLimits: true,
		MaxRetries:        3,
		BaseBackoff:       time.Second,
	})

	repositories, err := getRepositories(ctx, githubClient.RestClient, organizationName)
	if err != nil {
		return nil, err
	}

	var values []models.Resource
	for _, repo := range repositories {
		webhooks, err := processRepositoryWebhooks(ctx, handler, organizationName, repo.GetName())
		if err != nil {
			return nil, err
		}
		for _, webhook := range webhooks {
			config := model.WebhookConfig{
				URL:         webhook.Config.URL,
				ContentType: webhook.Config.ContentType,
				Secret:      webhook.Config.Secret,
				InsecureSSL: webhook.Config.InsecureSSL,
			}
			lastResponse := model.HookResponse{
				Code:    webhook.LastResponse.Code,
				Status:  webhook.LastResponse.Status,
				Message: webhook.LastResponse.Message,
			}
			value := models.Resource{
				ID:   strconv.Itoa(int(webhook.ID)),
				Name: webhook.Name,
				Description: model.WebhookDescription{
					Type:          webhook.Type,
					ID:            webhook.ID,
					Name:          webhook.Name,
					Active:        webhook.Active,
					Events:        webhook.Events,
					Config:        config,
					UpdatedAt:     webhook.UpdatedAt,
					CreatedAt:     webhook.CreatedAt,
					URL:           webhook.URL,
					TestURL:       webhook.TestURL,
					PingURL:       webhook.PingURL,
					DeliveriesURL: webhook.DeliveriesURL,
					LastResponse:  lastResponse,
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
	}

	return values, nil
}

func GetRepositoryWebhook(ctx context.Context, githubClient model.GitHubClient, organizationName string, repositoryName string, resourceID string, stream *models.StreamSender) (*models.Resource, error) {
	handler := resilientbridge.NewResilientBridge()
	handler.SetDebug(false)
	handler.RegisterProvider("github", adapters.NewGitHubAdapter(githubClient.Token), &resilientbridge.ProviderConfig{
		UseProviderLimits: true,
		MaxRetries:        3,
		BaseBackoff:       time.Second,
	})

	webhook, err := processRepositoryWebhook(ctx, handler, organizationName, repositoryName, resourceID)
	if err != nil {
		return nil, err
	}
	config := model.WebhookConfig{
		URL:         webhook.Config.URL,
		ContentType: webhook.Config.ContentType,
		Secret:      webhook.Config.Secret,
		InsecureSSL: webhook.Config.InsecureSSL,
	}
	lastResponse := model.HookResponse{
		Code:    webhook.LastResponse.Code,
		Status:  webhook.LastResponse.Status,
		Message: webhook.LastResponse.Message,
	}
	value := models.Resource{
		ID:   strconv.Itoa(int(webhook.ID)),
		Name: webhook.Name,
		Description: model.WebhookDescription{
			Type:          webhook.Type,
			ID:            webhook.ID,
			Name:          webhook.Name,
			Active:        webhook.Active,
			Events:        webhook.Events,
			Config:        config,
			UpdatedAt:     webhook.UpdatedAt,
			CreatedAt:     webhook.CreatedAt,
			URL:           webhook.URL,
			TestURL:       webhook.TestURL,
			PingURL:       webhook.PingURL,
			DeliveriesURL: webhook.DeliveriesURL,
			LastResponse:  lastResponse,
		},
	}

	return &value, nil
}

func processRepositoryWebhooks(ctx context.Context, handler *resilientbridge.ResilientBridge, organization, repo string) ([]model.WebhookJSON, error) {
	var webhooks []model.WebhookJSON
	var responseWebhooks []model.WebhookJSON
	baseURL := "/repos/"
	page := 1

	for {
		params := url.Values{}
		params.Set("per_page", "100")
		params.Set("page", strconv.Itoa(page))
		finalURL := fmt.Sprintf("%s%s/%s/hooks?%s", baseURL, organization, repo, params.Encode())

		req := &resilientbridge.NormalizedRequest{
			Method:   "GET",
			Endpoint: finalURL,
			Headers:  map[string]string{"accept": "application/vnd.github+json"},
		}

		resp, err := handler.Request("github", req)
		if err != nil {
			return nil, fmt.Errorf("request execution failed: %w", err)
		}

		if resp.StatusCode >= 400 {
			return nil, fmt.Errorf("error %d: %s", resp.StatusCode, string(resp.Data))
		}

		if err = json.Unmarshal(resp.Data, &responseWebhooks); err != nil {
			return nil, fmt.Errorf("error parsing response: %w", err)
		}

		webhooks = append(webhooks, responseWebhooks...)
		if len(responseWebhooks) < 100 {
			break
		}
		page++
	}

	return webhooks, nil
}

func processRepositoryWebhook(ctx context.Context, handler *resilientbridge.ResilientBridge, organization, repo, resourceID string) (*model.WebhookJSON, error) {
	var webhook model.WebhookJSON
	baseURL := "/repos/"

	finalURL := fmt.Sprintf("%s%s/%s/hooks/%s", baseURL, organization, repo, resourceID)

	req := &resilientbridge.NormalizedRequest{
		Method:   "GET",
		Endpoint: finalURL,
		Headers:  map[string]string{"accept": "application/vnd.github+json"},
	}

	resp, err := handler.Request("github", req)
	if err != nil {
		return nil, fmt.Errorf("request execution failed: %w", err)
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("error %d: %s", resp.StatusCode, string(resp.Data))
	}

	if err = json.Unmarshal(resp.Data, &webhook); err != nil {
		return nil, fmt.Errorf("error parsing response: %w", err)
	}

	return &webhook, nil
}
