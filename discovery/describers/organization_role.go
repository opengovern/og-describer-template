package describers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/opengovern/og-describer-github/discovery/pkg/models"
	model "github.com/opengovern/og-describer-github/discovery/provider"
	resilientbridge "github.com/opengovern/resilient-bridge"
	"github.com/opengovern/resilient-bridge/adapters"
	"strconv"
	"time"
)

func ListOrganizationRoles(ctx context.Context,
	githubClient model.GitHubClient,
	organizationName string,
	stream *models.StreamSender) ([]models.Resource, error) {
	sdk := resilientbridge.NewResilientBridge()
	sdk.RegisterProvider("github", adapters.NewGitHubAdapter(githubClient.Token), &resilientbridge.ProviderConfig{
		UseProviderLimits: true,
		MaxRetries:        3,
		BaseBackoff:       0,
	})

	var values []models.Resource

	endpoint := fmt.Sprintf("/orgs/%s/organization-roles", organizationName)
	req := &resilientbridge.NormalizedRequest{
		Method:   "GET",
		Endpoint: endpoint,
		Headers:  map[string]string{"Accept": "application/vnd.github+json"},
	}

	resp, err := sdk.Request("github", req)
	if err != nil {
		return nil, fmt.Errorf("error fetching repos: %w", err)
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP error %d: %s", resp.StatusCode, string(resp.Data))
	}

	// Decode into a slice of generic maps. We'll only parse name, archived, disabled, etc.
	var rolesResponse OrganizationRolesResponse
	if err := json.Unmarshal(resp.Data, &rolesResponse); err != nil {
		return nil, fmt.Errorf("error decoding repos list: %w", err)
	}

	for _, r := range rolesResponse.Roles {
		value := models.Resource{
			ID:   strconv.Itoa(r.ID),
			Name: r.Name,
			Description: model.OrganizationRoleDescription{
				Organization: organizationName,
				Name:         r.Name,
				ID:           r.ID,
				Source:       r.Source,
				BaseRole:     r.BaseRole,
				Permissions:  r.Permissions,
				Description:  r.Description,
				CreatedAt:    r.CreatedAt,
				UpdatedAt:    r.UpdatedAt,
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

	return values, nil
}

type OrganizationRole struct {
	ID           int         `json:"id"`
	Name         string      `json:"name"`
	Description  string      `json:"description"`
	Organization interface{} `json:"organization"`
	Permissions  []string    `json:"permissions"`
	BaseRole     string      `json:"base_role"`
	Source       string      `json:"source"`
	CreatedAt    time.Time   `json:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at"`
}

type OrganizationRolesResponse struct {
	Roles []OrganizationRole `json:"roles"`
}
