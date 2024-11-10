//go:generate go run ../../SDK/runnable/models/main.go --file $GOFILE --output ../../SDK/generated/resources_clients.go --type $PROVIDER

// Implement types for each resource

package model

import (
	"github.com/google/go-github/v55/github"
	steampipemodels "github.com/opengovern/og-describer-template/steampipe-plugin-github/github/models"
)

type Artifact struct {
	github.Artifact
	RepoFullName string
}

type Runner struct {
	github.Runner
	RepoFullName string
}

type Secret struct {
	github.Secret
	RepoFullName string
}

type WorkflowRun struct {
	github.WorkflowRun
	RepoFullName string
}

type Branch struct {
	steampipemodels.Branch
	RepoFullName string
	Protected    bool
}

type BranchProtection struct {
	steampipemodels.BranchProtectionRuleWithFirstPageEmbeddedItems
	RepoFullName                    string
	CreatorLogin                    string
	PushAllowanceApps               []App
	PushAllowanceTeams              []Team
	PushAllowanceUsers              []User
	BypassForcePushAllowanceApps    []App
	BypassForcePushAllowanceTeams   []Team
	BypassForcePushAllowanceUsers   []User
	BypassPullRequestAllowanceApps  []App
	BypassPullRequestAllowanceTeams []Team
	BypassPullRequestAllowanceUsers []User
}

type App struct {
	Name string
	Slug string
}

type Team struct {
	Name string
	Slug string
}

type User struct {
	Name  string
	Login string
}

type Commit struct {
	steampipemodels.Commit
	RepoFullName   string
	AuthorLogin    string
	CommitterLogin string
}

type CommunityProfile struct {
	steampipemodels.CommunityProfile
	RepoFullName string
	ReadMe       steampipemodels.Blob
	Contributing steampipemodels.Blob
	Security     steampipemodels.Blob
}

type GitIgnore struct {
	github.Gitignore
}

type Gist struct {
	github.Gist
}

type Organization struct {
	steampipemodels.OrganizationWithCounts
}

type Repository struct {
	steampipemodels.Repository
}

type Star struct {
	RepoFullName string
	StarredAt    steampipemodels.NullableTime
	Url          string
}

type Issue struct {
	steampipemodels.Issue
	RepoFullName string
}

type IssueComment struct {
	steampipemodels.IssueComment
	RepoFullName string
	Number       int
}
