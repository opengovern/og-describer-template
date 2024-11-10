//go:generate go run ../../SDK/runnable/models/main.go --file $GOFILE --output ../../SDK/generated/resources_clients.go --type $PROVIDER

// Implement types for each resource

package model

import (
	"github.com/google/go-github/v55/github"
	steampipemodels "github.com/opengovern/og-describer-template/steampipe-plugin-github/github/models"
)

type Artifact struct {
	ArtifactInfo github.Artifact
	RepoFullName string
}

type Runner struct {
	RunnerInfo   github.Runner
	RepoFullName string
}

type Secret struct {
	SecretInfo   github.Secret
	RepoFullName string
}

type WorkflowRun struct {
	WorkflowRunInfo github.WorkflowRun
	RepoFullName    string
}

type Repository struct {
	RepositoryInfo steampipemodels.Repository
}
