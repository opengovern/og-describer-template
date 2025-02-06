package provider

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/time/rate"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const (
	Limit           = "100"
	IncludeReplicas = "true"
	IncludePreviews = "true"
)

type RenderAPIHandler struct {
	Client       *http.Client
	APIKey       string
	RateLimiter  *rate.Limiter
	Semaphore    chan struct{}
	MaxRetries   int
	RetryBackoff time.Duration
}

func NewRenderAPIHandler(apiKey string, rateLimit rate.Limit, burst int, maxConcurrency int, maxRetries int, retryBackoff time.Duration) *RenderAPIHandler {
	return &RenderAPIHandler{
		Client:       http.DefaultClient,
		APIKey:       apiKey,
		RateLimiter:  rate.NewLimiter(rateLimit, burst),
		Semaphore:    make(chan struct{}, maxConcurrency),
		MaxRetries:   maxRetries,
		RetryBackoff: retryBackoff,
	}
}

// DoRequest executes the render API request with rate limiting, retries, and concurrency control.
func (h *RenderAPIHandler) DoRequest(ctx context.Context, req *http.Request, requestFunc func(req *http.Request) (*http.Response, error)) error {
	h.Semaphore <- struct{}{}
	defer func() { <-h.Semaphore }()
	var resp *http.Response
	var err error
	for attempt := 0; attempt <= h.MaxRetries; attempt++ {
		// Wait based on rate limiter
		if err = h.RateLimiter.Wait(ctx); err != nil {
			return err
		}
		// Set request headers
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", h.APIKey))
		// Execute the request function
		resp, err = requestFunc(req)
		if err == nil {
			return nil
		}
		// Set rate limiter new value
		var remainTime int
		if resp != nil {
			remainTimeStr := resp.Header.Get("Ratelimit-Reset")
			if remainTimeStr != "" {
				remainTime, _ = strconv.Atoi(remainTimeStr)
			}
			var remainRequests int
			remainRequestsStr := resp.Header.Get("Ratelimit-Remaining")
			if remainRequestsStr != "" {
				remainRequests, err = strconv.Atoi(remainRequestsStr)
				if err == nil && remainTime > 0 {
					h.RateLimiter = rate.NewLimiter(rate.Every(time.Duration(remainTime)/time.Duration(remainRequests)), 1)
				}
			}
		}
		// Handle rate limit errors
		if resp != nil && resp.StatusCode == http.StatusTooManyRequests {
			if remainTime > 0 {
				time.Sleep(time.Duration(remainTime))
				continue
			}
			// Exponential backoff if headers are missing
			backoff := h.RetryBackoff * (1 << attempt)
			time.Sleep(backoff)
			continue
		}
		// Handle temporary network errors
		if IsTemporary(err) {
			backoff := h.RetryBackoff * (1 << attempt)
			time.Sleep(backoff)
			continue
		}
		break
	}
	return err
}

// IsTemporary checks if an error is temporary.
func IsTemporary(err error) bool {
	if err == nil {
		return false
	}
	var netErr interface{ Temporary() bool }
	if errors.As(err, &netErr) {
		return netErr.Temporary()
	}
	return false
}

func GetServices(ctx context.Context, handler *RenderAPIHandler) ([]ServiceJSON, error) {
	var services []ServiceJSON
	var serviceListResponse []ServiceResponse
	var resp *http.Response
	baseURL := "https://api.render.com/v1/services"
	cursor := ""

	for {
		params := url.Values{}
		params.Set("limit", Limit)
		params.Set("includePreviews", IncludePreviews)
		if cursor != "" {
			params.Set("cursor", cursor)
		}
		finalURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

		req, err := http.NewRequest("GET", finalURL, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create request: %w", err)
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
			return nil, fmt.Errorf("error during request handling: %w", err)
		}

		if len(serviceListResponse) < 100 {
			break
		}
	}
	return services, nil
}

func GetProjects(ctx context.Context, handler *RenderAPIHandler) ([]ProjectJSON, error) {
	var projects []ProjectJSON
	var projectListResponse []ProjectResponse
	var resp *http.Response
	baseURL := "https://api.render.com/v1/projects"
	cursor := ""

	for {
		params := url.Values{}
		params.Set("limit", Limit)
		if cursor != "" {
			params.Set("cursor", cursor)
		}
		finalURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

		req, err := http.NewRequest("GET", finalURL, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create request: %w", err)
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
			return nil, fmt.Errorf("error during request handling: %w", err)
		}

		if len(projectListResponse) < 100 {
			break
		}
	}
	return projects, nil
}

func GetPostgresqlInstances(ctx context.Context, handler *RenderAPIHandler) ([]PostgresJSON, error) {
	var postgresInstances []PostgresJSON
	var postgresListResponse []PostgresResponse
	var resp *http.Response
	baseURL := "https://api.render.com/v1/postgres"
	cursor := ""

	for {
		params := url.Values{}
		params.Set("limit", Limit)
		params.Set("includeReplicas", IncludeReplicas)
		if cursor != "" {
			params.Set("cursor", cursor)
		}
		finalURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

		req, err := http.NewRequest("GET", finalURL, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create request: %w", err)
		}

		requestFunc := func(req *http.Request) (*http.Response, error) {
			var e error
			resp, e = handler.Client.Do(req)
			if e != nil {
				return nil, fmt.Errorf("request execution failed: %w", e)
			}
			defer resp.Body.Close()

			if e = json.NewDecoder(resp.Body).Decode(&postgresListResponse); e != nil {
				return nil, fmt.Errorf("failed to decode response: %w", e)
			}
			for i, postgresResp := range postgresListResponse {
				postgresInstances = append(postgresInstances, postgresResp.Postgres)
				if i == len(postgresListResponse)-1 {
					cursor = postgresResp.Cursor
				}
			}
			return resp, nil
		}
		err = handler.DoRequest(ctx, req, requestFunc)
		if err != nil {
			return nil, fmt.Errorf("error during request handling: %w", err)
		}

		if len(postgresListResponse) < 100 {
			break
		}
	}

	return postgresInstances, nil
}
