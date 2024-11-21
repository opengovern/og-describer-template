package entraid

import (
	"context"
	"github.com/opengovern/og-describer-entraid/pkg/sdk/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableEntraIdSignInReport(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "entraid_sign_in_report",
		Description: "Represents an Azure Active Directory (Azure AD) sign-in report.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListAdSignInReport,
		},

		Columns: azureOGColumns([]*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Unique ID representing the sign-in activity.", Transform: transform.FromField("Description.Id")},
			{Name: "created_date_time", Type: proto.ColumnType_TIMESTAMP, Description: "Date and time (UTC) the sign-in was initiated.", Transform: transform.FromField("Description.CreatedDateTime")},
			{Name: "user_display_name", Type: proto.ColumnType_STRING, Description: "Display name of the user that initiated the sign-in.", Transform: transform.FromField("Description.UserDisplayName")},
			{Name: "user_principal_name", Type: proto.ColumnType_STRING, Description: "User principal name of the user that initiated the sign-in.", Transform: transform.FromField("Description.UserPrincipalName")},
			{Name: "user_id", Type: proto.ColumnType_STRING, Description: "ID of the user that initiated the sign-in.", Transform: transform.FromField("Description.UserId")},
			{Name: "app_id", Type: proto.ColumnType_STRING, Description: "Unique GUID representing the app ID in the Azure Active Directory.", Transform: transform.FromField("Description.AppId")},
			{Name: "app_display_name", Type: proto.ColumnType_STRING, Description: "App name displayed in the Azure Portal.", Transform: transform.FromField("Description.AppDisplayName")},
			{Name: "ip_address", Type: proto.ColumnType_STRING, Description: "IP address of the client used to sign in.", Transform: transform.FromField("Description.IpAddress")},
			{Name: "client_app_used", Type: proto.ColumnType_STRING, Description: "Identifies the legacy client used for sign-in activity.", Transform: transform.FromField("Description.ClientAppUsed")},
			{Name: "correlation_id", Type: proto.ColumnType_STRING, Description: "The request ID sent from the client when the sign-in is initiated; used to troubleshoot sign-in activity.", Transform: transform.FromField("Description.CorrelationId")},
			{Name: "conditional_access_status", Type: proto.ColumnType_STRING, Description: "Reports status of an activated conditional access policy. Possible values are: success, failure, notApplied, and unknownFutureValue.", Transform: transform.FromField("Description.ConditionalAccessStatus")},
			{Name: "is_interactive", Type: proto.ColumnType_BOOL, Description: "Indicates if a sign-in is interactive or not.", Transform: transform.FromField("Description.IsInteractive")},
			{Name: "risk_detail", Type: proto.ColumnType_STRING, Description: "Provides the 'reason' behind a specific state of a risky user, sign-in or a risk event. The possible values are: none, adminGeneratedTemporaryPassword, userPerformedSecuredPasswordChange, userPerformedSecuredPasswordReset, adminConfirmedSigninSafe, aiConfirmedSigninSafe, userPassedMFADrivenByRiskBasedPolicy, adminDismissedAllRiskForUser, adminConfirmedSigninCompromised, unknownFutureValue.", Transform: transform.FromField("Description.RiskDetail")},
			{Name: "risk_level_aggregated", Type: proto.ColumnType_STRING, Description: "Aggregated risk level. The possible values are: none, low, medium, high, hidden, and unknownFutureValue.", Transform: transform.FromField("Description.RiskLevelAggregated")},
			{Name: "risk_level_during_sign_in", Type: proto.ColumnType_STRING, Description: "Risk level during sign-in. The possible values are: none, low, medium, high, hidden, and unknownFutureValue.", Transform: transform.FromField("Description.RiskLevelDuringSignIn")},
			{Name: "risk_state", Type: proto.ColumnType_STRING, Description: "Reports status of the risky user, sign-in, or a risk event. The possible values are: none, confirmedSafe, remediated, dismissed, atRisk, confirmedCompromised, unknownFutureValue.", Transform: transform.FromField("Description.RiskState")},
			{Name: "resource_display_name", Type: proto.ColumnType_STRING, Description: "Name of the resource the user signed into.", Transform: transform.FromField("Description.ResourceDisplayName")},
			{Name: "resource_id", Type: proto.ColumnType_STRING, Description: "ID of the resource that the user signed into.", Transform: transform.FromField("Description.ResourceId")},

			// JSON fields
			{Name: "risk_event_types", Type: proto.ColumnType_JSON, Description: "Risk event types associated with the sign-in. The possible values are: unlikelyTravel, anonymizedIPAddress, maliciousIPAddress, unfamiliarFeatures, malwareInfectedIPAddress, suspiciousIPAddress, leakedCredentials, investigationsThreatIntelligence, generic, and unknownFutureValue.", Transform: transform.FromField("Description.RiskEventTypes").Transform(formatSignInReportRiskEventTypes)},
			{Name: "status", Type: proto.ColumnType_JSON, Description: "Sign-in status. Includes the error code and description of the error (in case of a sign-in failure).", Transform: transform.FromField("Description.Status")},
			{Name: "device_detail", Type: proto.ColumnType_JSON, Description: "Device information from where the sign-in occurred; includes device ID, operating system, and browser.", Transform: transform.FromField("Description.DeviceDetail")},
			{Name: "location", Type: proto.ColumnType_JSON, Description: "Provides the city, state, and country code where the sign-in originated.", Transform: transform.FromField("Description.Location")},
			{Name: "applied_conditional_access_policies", Type: proto.ColumnType_JSON, Description: "Provides a list of conditional access policies that are triggered by the corresponding sign-in activity.", Transform: transform.FromField("Description.AppliedConditionalAccessPolicies")},

			// Standard columns
			{Name: "title", Type: proto.ColumnType_STRING, Description: ColumnDescriptionTitle, Transform: transform.FromField("Description.Id")},
			{Name: "tenant_id", Type: proto.ColumnType_STRING, Description: ColumnDescriptionTenant, Transform: transform.FromField("Description.TenantID")},
		}),
	}
}

func formatSignInReportRiskEventTypes(_ context.Context, d *transform.TransformData) (interface{}, error) {
	data := d.HydrateItem.(opengovernance.AdSignInReport).Description
	riskEventTypes := data.RiskEventTypes
	if len(riskEventTypes) == 0 {
		return nil, nil
	}

	return riskEventTypes, nil
}
