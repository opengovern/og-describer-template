package github

import (
	opengovernance "github.com/opengovern/og-describer-github/pkg/sdk/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func sharedCommentsColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "repository_full_name",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("Description.RepoFullName"),
			Description: "The full name of the repository (login/repo-name)."},
		{
			Name:        "number",
			Type:        proto.ColumnType_INT,
			Transform:   transform.FromField("Description.Number"),
			Description: "The issue/pr number."},
		{
			Name:        "id",
			Type:        proto.ColumnType_INT,
			Transform:   transform.FromField("Description.Id"),
			Description: "The ID of the comment."},
		{
			Name:        "node_id",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("Description.NodeId"),
			Description: "The node ID of the comment."},
		{
			Name:        "author",
			Type:        proto.ColumnType_JSON,
			Transform:   transform.FromField("Description.Author"),
			Description: "The actor who authored the comment."},
		{
			Name:        "author_login",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("Description.AuthorLogin"),
			Description: "The login of the comment author."},
		{
			Name:        "author_association",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("Description.AuthorAssociation"),
			Description: "Author's association with the subject of the issue/pr the comment was raised on."},
		{
			Name:        "body",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("Description.Body"),
			Description: "The contents of the comment as markdown."},
		{
			Name:        "body_text",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("Description.BodyText"),
			Description: "The contents of the comment as text."},
		{
			Name:        "created_at",
			Type:        proto.ColumnType_TIMESTAMP,
			Transform:   transform.FromField("Description.CreatedAt").Transform(convertTimestamp),
			Description: "Timestamp when comment was created."},
		{
			Name:      "created_via_email",
			Type:      proto.ColumnType_BOOL,
			Transform: transform.FromField("Description.CreatedViaEmail"),

			Description: "If true, comment was created via email."},
		{
			Name:      "editor",
			Type:      proto.ColumnType_JSON,
			Transform: transform.FromField("Description.Editor"),

			Description: "The actor who edited the comment."},
		{
			Name:      "editor_login",
			Type:      proto.ColumnType_STRING,
			Transform: transform.FromField("Description.EditorLogin"),

			Description: "The login of the comment editor."},
		{
			Name:      "includes_created_edit",
			Type:      proto.ColumnType_BOOL,
			Transform: transform.FromField("Description.IncludesCreatedEdit"),

			Description: "If true, comment was edited and includes an edit with the creation data."},
		{
			Name:      "is_minimized",
			Type:      proto.ColumnType_BOOL,
			Transform: transform.FromField("Description.IsMinimized"),

			Description: "If true, comment has been minimized."},
		{
			Name:      "minimized_reason",
			Type:      proto.ColumnType_STRING,
			Transform: transform.FromField("Description.MinimizedReason"),

			Description: "The reason for comment being minimized."},
		{
			Name:      "last_edited_at",
			Type:      proto.ColumnType_TIMESTAMP,
			Transform: transform.FromField("Description.LastEditedAt").Transform(convertTimestamp),

			Description: "Timestamp when comment was last edited."},
		{
			Name:      "published_at",
			Type:      proto.ColumnType_TIMESTAMP,
			Transform: transform.FromField("Description.PublishedAt").Transform(convertTimestamp),

			Description: "Timestamp when comment was published."},
		{
			Name:      "updated_at",
			Type:      proto.ColumnType_TIMESTAMP,
			Transform: transform.FromField("Description.UpdatedAt").Transform(convertTimestamp),

			Description: "Timestamp when comment was last updated."},
		{
			Name:      "url",
			Type:      proto.ColumnType_STRING,
			Transform: transform.FromField("Description.Url"),

			Description: "URL for the comment."},
		{
			Name:      "can_delete",
			Type:      proto.ColumnType_BOOL,
			Transform: transform.FromField("Description.CanDelete"),

			Description: "If true, user can delete the comment."},
		{
			Name:      "can_minimize",
			Type:      proto.ColumnType_BOOL,
			Transform: transform.FromField("Description.CanMinimize"),

			Description: "If true, user can minimize the comment."},
		{
			Name:      "can_react",
			Type:      proto.ColumnType_BOOL,
			Transform: transform.FromField("Description.CanReact"),

			Description: "If true, user can react to the comment."},
		{
			Name:      "can_update",
			Type:      proto.ColumnType_BOOL,
			Transform: transform.FromField("Description.CanUpdate"),

			Description: "If true, user can update the comment."},
		{
			Name:      "cannot_update_reasons",
			Type:      proto.ColumnType_JSON,
			Transform: transform.FromField("Description.CannotUpdateReasons"),

			Description: "A list of reasons why user cannot update the comment."},
		{
			Name:      "did_author",
			Type:      proto.ColumnType_BOOL,
			Transform: transform.FromField("Description.DidAuthor"),

			Description: "If true, user authored the comment."},
	}
}

func tableGitHubIssueComment() *plugin.Table {
	return &plugin.Table{
		Name:        "github_issue_comment",
		Description: "GitHub Issue Comments are the responses/comments on GitHub Issues.",
		List: &plugin.ListConfig{
			KeyColumns:        plugin.AllColumns([]string{"repository_full_name", "number"}),
			ShouldIgnoreError: isNotFoundError([]string{"404"}),
			Hydrate:           opengovernance.ListIssueComment,
		},
		Columns: commonColumns(sharedCommentsColumns()),
	}
}
