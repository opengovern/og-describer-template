package github

import (
	opengovernance "github.com/opengovern/og-describer-github/pkg/sdk/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableGitHubActionsRepositoryWorkflowRun() *plugin.Table {
	return &plugin.Table{
		Name:        "github_actions_repository_workflow_run",
		Description: "WorkflowRun represents a repository action workflow run",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListWorkflowRun,
		},
		Get: &plugin.GetConfig{
			KeyColumns:        plugin.AllColumns([]string{"repository_full_name", "id"}),
			ShouldIgnoreError: isNotFoundError([]string{"404"}),
			Hydrate:           opengovernance.GetWorkflowRun,
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{
				Name:        "repository_full_name",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromQual("Description.RepoFullName"),
				Description: "Full name of the repository that specifies the workflow run.",
			},
			{
				Name:        "id",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromQual("Description.ID"),
				Description: "The unque identifier of the workflow run.",
			},
			{
				Name:        "event",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromQual("Description.Event"),
				Description: "The event for which workflow triggered off.",
			},
			{
				Name:        "workflow_id",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromQual("Description.WorkflowID"),
				Description: "The workflow id of the worflow run.",
			},
			{
				Name:        "node_id",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromQual("Description.NodeID"),
				Description: "The node id of the worflow run.",
			},
			{
				Name:        "conclusion",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromQual("Description.Conclusion"),
				Description: "The conclusion for workflow run."},
			{
				Name:        "status",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromQual("Description.Status"),
				Description: "The status of the worflow run."},
			{
				Name:        "run_number",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromQual("Description.RunNumber"),
				Description: "The number of time workflow has run."},
			{
				Name:        "artifacts_url",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromQual("Description.ArtifactsURL"),
				Description: "The address for artifact GitHub web page."},
			{
				Name:        "cancel_url",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromQual("Description.CancelURL"),
				Description: "The address for workflow run cancel GitHub web page."},
			{
				Name:        "check_suite_url",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromQual("Description.CheckSuiteURL"),
				Description: "The address for the workflow check suite GitHub web page."},
			{
				Name:        "head_branch",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromQual("Description.HeadBranch"),
				Description: "The head branch of the workflow run branch."},
			{
				Name:        "head_sha",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromQual("Description.HeadSHA"),
				Description: "The head sha of the workflow run."},
			{
				Name:        "html_url",
				Type:        proto.ColumnType_STRING,
				Description: "The address for the organization's GitHub web page.",
				Transform:   transform.FromQual("Description.HTMLURL"),
			},
			{
				Name:        "jobs_url",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromQual("Description.JobsURL"),
				Description: "The address for the workflow job GitHub web page."},
			{
				Name:        "logs_url",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromQual("Description.LogsURL"),
				Description: "The address for the workflow logs GitHub web page."},
			{
				Name:        "rerun_url",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromQual("Description.RerunURL"),
				Description: "The address for workflow rerun GitHub web page."},
			{
				Name:        "url",
				Type:        proto.ColumnType_STRING,
				Description: "The address for the workflow run GitHub web page.",
				Transform:   transform.FromQual("Description.URL"),
			},
			{
				Name:        "workflow_url",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromQual("Description.WorkflowURL"),
				Description: "The address for workflow GitHub web page."},

			// Other columns
			{
				Name:        "created_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.CreatedAt"),
				Description: "Time when the workflow run was created."},
			{
				Name:        "head_commit",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromQual("Description.HeadCommit"),
				Description: "The head commit details for workflow run."},
			{
				Name:        "head_repository",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromQual("Description.HeadRepository"),
				Description: "The head repository info for the workflow run."},
			{
				Name:        "pull_requests",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromQual("Description.PullRequests"),
				Description: "The pull request details for the workflow run."},
			{
				Name:        "repository",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromQual("Description.Repository"),
				Description: "The repository info for the workflow run."},
			{
				Name:        "run_started_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.RunStartedAt"),
				Description: "Time when the workflow run was started."},
			{
				Name:        "updated_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.UpdatedAt"),
				Description: "Time when the workflow run was updated."},
			{
				Name:        "actor",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Actor"),
				Description: "The user whom initiated the first instance of this workflow run.",
			},
			{
				Name:        "actor_login",
				Type:        proto.ColumnType_STRING,
				Description: "The login of the user whom initiated the first instance of the workflow run.",
				Transform:   transform.FromField("Description.Actor.Login"),
			},
			{
				Name:        "triggering_actor",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.TriggeringActor"),
				Description: "The user whom initiated the latest instance of this workflow run.",
			},
			{
				Name:        "triggering_actor_login",
				Type:        proto.ColumnType_STRING,
				Description: "The login of the user whom initiated the latest instance of this workflow run.",
				Transform:   transform.FromField("Description.TriggeringActor.Login"),
			},
		}),
	}
}
