package github

import (
	opengovernance "github.com/opengovern/og-describer-github/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func gitHubOrganizationTokenColumns() []*plugin.Column {
	tableCols := []*plugin.Column{
		{
			Name:        "title",
			Type:        proto.ColumnType_STRING,
			Description: "login",
			Transform:   transform.FromField("Description.Title")},
		{
			Name:        "login",
			Type:        proto.ColumnType_STRING,
			Description: "login",
			Transform:   transform.FromField("Description.Login")},
		{
			Name:        "principle_type",
			Type:        proto.ColumnType_STRING,
			Description: "login",
			Transform:   transform.FromField("Description.PrincipleType")},
		{
			Name:        "organization_id",
			Type:        proto.ColumnType_STRING,
			Description: "login",
			Transform:   transform.FromField("Description.OrganizationID")},
		{
			Name:        "principal_id",
			Type:        proto.ColumnType_STRING,
			Description: "login",
			Transform:   transform.FromField("Description.PrincipalID")},
		{
			Name:        "credential_id",
			Type:        proto.ColumnType_INT,
			Description: "login",
			Transform:   transform.FromField("Description.CredentialId")},
		{
			Name:        "credential_type",
			Type:        proto.ColumnType_STRING,
			Description: "login",
			Transform:   transform.FromField("Description.CredentialType")},
		{
			Name:        "token_last_eight",
			Type:        proto.ColumnType_STRING,
			Description: "login",
			Transform:   transform.FromField("Description.TokenLastEight")},
		{
			Name:        "credential_authorized_at",
			Type:        proto.ColumnType_TIMESTAMP,
			Description: "login",
			Transform:   transform.FromField("Description.CredentialAuthorizedAt")},
		{
			Name:        "scopes",
			Type:        proto.ColumnType_JSON,
			Description: "login",
			Transform:   transform.FromField("Description.Scopes")},
		{
			Name:        "fingerprint",
			Type:        proto.ColumnType_STRING,
			Description: "login",
			Transform:   transform.FromField("Description.Fingerprint")},
		{
			Name:        "credential_accessed_at",
			Type:        proto.ColumnType_TIMESTAMP,
			Description: "login",
			Transform:   transform.FromField("Description.CredentialAccessedAt")},
		{
			Name:        "authorized_credential_id",
			Type:        proto.ColumnType_INT,
			Description: "login",
			Transform:   transform.FromField("Description.AuthorizedCredentialId")},
		{
			Name:        "authorized_credential_expires_at",
			Type:        proto.ColumnType_TIMESTAMP,
			Description: "login",
			Transform:   transform.FromField("Description.AuthorizedCredentialExpiresAt")},
	}

	return tableCols
}

func tableGitHubOrganizationToken() *plugin.Table {
	return &plugin.Table{
		Name:        "github_organization_token",
		Description: "GitHub tokens for a given organization. GitHub Users are user accounts in GitHub.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListOrganizationToken,
		},
		Columns: commonColumns(gitHubOrganizationTokenColumns()),
	}
}
