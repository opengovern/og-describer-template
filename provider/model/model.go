//go:generate go run ../../SDK/runnable/models/main.go --file $GOFILE --output ../../SDK/generated/resources_clients.go --type $PROVIDER

// Implement types for each resource

package model

import (
	"github.com/google/go-github/v55/github"
	steampipemodels "github.com/turbot/steampipe-plugin-github/github/models"
)

type Artifact struct {
	ArtifactInfo github.Artifact
}

type Runner struct {
	RunnerInfo github.Runner
}

type Secret struct {
	SecretInfo github.Secret
}

type Repository struct {
	RepositoryInfo steampipemodels.Repository
}
