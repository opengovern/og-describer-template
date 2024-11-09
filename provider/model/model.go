//go:generate go run ../../SDK/runnable/models/main.go --file $GOFILE --output ../../SDK/generated/resources_clients.go --type $PROVIDER

// Implement types for each resource

package model

import steampipemodels "github.com/turbot/steampipe-plugin-github/github/models"

type Repository struct {
	RepositoryInfo steampipemodels.Repository
}
