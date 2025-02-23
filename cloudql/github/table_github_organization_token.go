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
			Name:        "authorized_credential_id",
			Type:        proto.ColumnType_INT,
			Description: "The organization the member is associated with.",
			Transform:   transform.FromField("Description.AuthorizedCredentialId")},
		{
			Name:        "authorized_credential_title",
			Type:        proto.ColumnType_STRING,
			Description: "The role this user has in the organization. Returns null if information is not available to viewer.",
			Transform:   transform.FromField("Description.AuthorizedCredentialTitle")},
		{
			Name:      "authorized_credential_note",
			Type:      proto.ColumnType_STRING,
			Transform: transform.FromField("Description.AuthorizedCredentialNote")},
		{
			Name:        "authorized_credential_expires_at",
			Type:        proto.ColumnType_TIMESTAMP,
			Description: "permissions",
			Transform:   transform.FromField("Description.AuthorizedCredentialExpiresAt")},
		{
			Name:        "login",
			Type:        proto.ColumnType_STRING,
			Description: "login",
			Transform:   transform.FromField("Description.Login")},
		{
			Name:        "scopes",
			Description: "permissions",
			Type:        proto.ColumnType_JSON,
			Transform:   transform.FromField("Description.Scopes")},
		{
			Name:        "credential_id",
			Type:        proto.ColumnType_INT,
			Description: "permissions",
			Transform:   transform.FromField("Description.CredentialId")},
		{
			Name:        "credential_type",
			Type:        proto.ColumnType_STRING,
			Description: "permissions",
			Transform:   transform.FromField("Description.CredentialType")},
		{
			Name:        "credential_accessed_at",
			Type:        proto.ColumnType_TIMESTAMP,
			Description: "permissions",
			Transform:   transform.FromField("Description.CredentialAccessedAt")},
		{
			Name:        "credential_authorized_at",
			Type:        proto.ColumnType_TIMESTAMP,
			Description: "permissions",
			Transform:   transform.FromField("Description.CredentialAuthorizedAt")},
		{
			Name:        "token_last_eight",
			Type:        proto.ColumnType_STRING,
			Description: "permissions",
			Transform:   transform.FromField("Description.TokenLastEight")},
		{
			Name:        "fingerprint",
			Type:        proto.ColumnType_STRING,
			Description: "permissions",
			Transform:   transform.FromField("Description.Fingerprint")},
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
