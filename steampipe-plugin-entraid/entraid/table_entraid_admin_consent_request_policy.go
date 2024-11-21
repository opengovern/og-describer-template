package entraid

import (
	"context"
	"github.com/opengovern/og-describer-entraid/pkg/sdk/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableEntraIdAdminConsentRequestPolicy(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "entraid_admin_consent_request_policy",
		Description: "Represents the policy for enabling or disabling the Azure AD admin consent workflow.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListAdAdminConsentRequestPolicy,
		},

		Columns: azureKaytuColumns([]*plugin.Column{
			{
				Name:        "is_enabled",
				Type:        proto.ColumnType_BOOL,
				Description: "Specifies whether the admin consent request feature is enabled or disabled.",
				Transform:   transform.FromField("Description.IsEnabled"),
			},

			// Other fields
			{
				Name:        "notify_reviewers",
				Type:        proto.ColumnType_BOOL,
				Description: "Specifies whether reviewers will receive notifications.",
				Transform:   transform.FromField("Description.NotifyReviewers"),
			},
			{
				Name:        "reminders_enabled",
				Type:        proto.ColumnType_BOOL,
				Description: "Specifies whether reviewers will receive reminder emails.",
				Transform:   transform.FromField("Description.RemindersEnabled"),
			},
			{
				Name:        "request_duration_in_days",
				Type:        proto.ColumnType_INT,
				Description: "Specifies the duration the request is active before it automatically expires if no decision is applied.",
				Transform:   transform.FromField("Description.RequestDurationInDays"),
			},
			{
				Name:        "version",
				Type:        proto.ColumnType_INT,
				Description: "Specifies the version of this policy. When the policy is updated, this version is updated.",
				Transform:   transform.FromField("Description.Version"),
			},

			// JSON fields
			{
				Name:        "reviewers",
				Type:        proto.ColumnType_JSON,
				Description: "The list of reviewers for the admin consent.",
				Transform:   transform.FromField("Description.Reviewers"),
			},

			// Standard columns
			{
				Name:        "title",
				Type:        proto.ColumnType_STRING,
				Description: ColumnDescriptionTitle,
				Transform:   transform.FromField("Description.Id"),
			},
			{
				Name:        "tenant_id",
				Type:        proto.ColumnType_STRING,
				Description: ColumnDescriptionTenant,
				Transform:   transform.FromField("Description.TenantID"),
			},
		}),
	}
}
