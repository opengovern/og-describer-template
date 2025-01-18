package main

import (
	"errors"
	"fmt"
	"net/http"
)

// ConfigHealthCheck represents the JSON input configuration
type ConfigHealthCheck struct {
	APIKey string `json:"api_key"`
}

// IsHealthy checks if the JWT has read access to all required resources
func IsHealthy(apiKey string) error {
	url := "https://api.render.com/v1/users"

	client := http.DefaultClient

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("request execution failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to retrieve user information correctly. status code: %v", resp.StatusCode)
	}

	return nil
}

func RenderIntegrationHealthcheck(cfg ConfigHealthCheck) (bool, error) {
	// Check for the api key
	if cfg.APIKey == "" {
		return false, errors.New("api key must be configured")
	}

	err := IsHealthy(cfg.APIKey)
	if err != nil {
		return false, err
	}

	return true, nil
}
