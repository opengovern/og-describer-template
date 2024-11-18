package github

import (
	opengovernance "github.com/opengovern/og-describer-github/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableGitHubCommit() *plugin.Table {
	return &plugin.Table{
		Name:        "github_commit",
		Description: "GitHub Commits bundle project files for download by users.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListCommit,
		},
		Get: &plugin.GetConfig{
			KeyColumns:        plugin.AllColumns([]string{"repository_full_name", "sha"}),
			ShouldIgnoreError: isNotFoundError([]string{"404"}),
			Hydrate:           opengovernance.GetCommit,
		},
		Columns: commonColumns([]*plugin.Column{
			{
				Name:        "repository_full_name",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RepoFullName"),
				Description: "Full name of the repository that contains the commit.",
			},
			{
				Name:        "sha",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Sha"),
				Description: "SHA of the commit."},
			{
				Name:        "short_sha",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ShortSha"),
				Description: "Short SHA of the commit."},
			{
				Name:        "message",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Message"),
				Description: "Commit message."},
			{
				Name:        "author_login",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AuthorLogin"),
				Description: "The login name of the author of the commit."},
			{
				Name:        "authored_date",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Describer.AuthoredDate").NullIfZero().Transform(convertTimestamp),
				Description: "Timestamp when the author made this commit."},
			{
				Name:        "author",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Author"),
				Description: "The commit author."},
			{
				Name:        "committer_login",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.CommitterLogin"),
				Description: "The login name of the committer."},
			{
				Name:        "committed_date",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.CommittedDate").NullIfZero().Transform(convertTimestamp),
				Description: "Timestamp when commit was committed."},
			{
				Name:        "committer",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Committer").NullIfZero(),
				Description: "The committer."},
			{
				Name:        "additions",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Additions"),
				Description: "Number of additions in the commit."},
			{
				Name:        "authored_by_committer",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.AuthoredByCommitter"),
				Description: "Check if the committer and the author match."},
			{
				Name:        "deletions",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Deletions"),
				Description: "Number of deletions in the commit."},
			{
				Name:        "changed_files",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.ChangedFiles"),
				Description: "Count of files changed in the commit."},
			{
				Name:        "committed_via_web",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.CommittedViaWeb"),
				Description: "If true, commit was made via GitHub web ui."},
			{
				Name:        "commit_url",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.CommitUrl"),
				Description: "URL of the commit."},
			{
				Name:        "signature",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Signature"),
				Description: "The signature of commit."},
			{
				Name:        "status",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Status"),
				Description: "Status of the commit."},
			{
				Name:        "tarball_url",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TarballUrl"),
				Description: "URL to download a tar of commit."},
			{
				Name:        "zipball_url",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ZipballUrl"),
				Description: "URL to download a zip of commit."},
			{
				Name:        "tree_url",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TreeUrl"),
				Description: "URL to tree of the commit."},
			{
				Name:        "can_subscribe",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.CanSubscribe"),
				Description: "If true, user can subscribe to this commit."},
			{
				Name:        "subscription",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Subscription"),
				Description: "Users subscription state of the commit."},
			{
				Name:        "url",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Url"),
				Description: "URL of the commit."},
			{
				Name:        "node_id",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.NodeId"),
				Description: "The node ID of the commit."},
			{
				Name:        "message_headline",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.MessageHeadline"),
				Description: "The Git commit message headline."},
		}),
	}
}
