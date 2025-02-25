package github

import (
	opengovernance "github.com/opengovern/og-describer-github/discovery/pkg/es"

	"github.com/opengovern/og-describer-github/cloudql/github/models"
	"github.com/shurcooL/githubv4"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func gitHubRepositoryCollaboratorColumns() []*plugin.Column {
	return []*plugin.Column{
		{Name: "repository_name", Type: proto.ColumnType_STRING, Description: "The full name of the repository, including the owner and repo name.",
			Transform: transform.FromField("Description.RepositoryName")},
		{Name: "repository_full_name", Type: proto.ColumnType_STRING, Description: "The full name of the repository, including the owner and repo name.",
			Transform: transform.FromField("Description.RepoFullName")},
		{Name: "collaborator_id", Type: proto.ColumnType_STRING, Description: "The full name of the repository, including the owner and repo name.",
			Transform: transform.FromField("Description.CollaboratorID")},
		{Name: "collaborator_type", Type: proto.ColumnType_STRING, Description: "The full name of the repository, including the owner and repo name.",
			Transform: transform.FromField("Description.CollaboratorType")},
		{Name: "permission", Type: proto.ColumnType_STRING, Description: "The permission the collaborator has on the repository.",
			Transform: transform.FromField("Description.Permission")},
		{
			Name:        "organization_id",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("Description.OrganizationID"),
			Description: "repository name",
		},
	}
}

func tableGitHubRepositoryCollaborator() *plugin.Table {
	return &plugin.Table{
		Name:        "github_repository_collaborator",
		Description: "Collaborators are users that have contributed to the repository.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListRepoCollaborators,
		},
		Columns: commonColumns(gitHubRepositoryCollaboratorColumns()),
	}
}

type RepositoryCollaborator struct {
	Permission githubv4.RepositoryPermission
	Node       models.BasicUser
}
