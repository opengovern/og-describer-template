package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAzureStreamAnalyticsJob(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_stream_analytics_job",
		Description: "Azure Stream Analytics Job",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "resource_group"}),
			Hydrate:    opengovernance.GetStreamAnalyticsJob,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "ResourceGroupNotFound", "Invalid input"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListStreamAnalyticsJob,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The resource name.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.StreamingJob.Name")},
			{
				Name:        "id",
				Description: "The resource identifier.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.StreamingJob.ID")},
			{
				Name:        "job_id",
				Description: "A GUID uniquely identifying the streaming job.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.StreamingJob.Properties.JobID")},
			{
				Name:        "job_state",
				Description: "Describes the state of the streaming job.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.StreamingJob.Properties.JobState")},
			{
				Name:        "provisioning_state",
				Description: "Describes the provisioning status of the streaming job.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.StreamingJob.Properties.ProvisioningState")},
			{
				Name:        "type",
				Description: "The resource type.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.StreamingJob.Type")},
			{
				Name:        "compatibility_level",
				Description: "Controls certain runtime behaviors of the streaming job.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.StreamingJob.Properties.CompatibilityLevel")},
			{
				Name:        "created_date",
				Description: "Specifies the time when the stream analytics job was created.",
				Type:        proto.ColumnType_TIMESTAMP,

				Transform: transform.FromField("Description.StreamingJob.Properties.CreatedDate").Transform(convertDateToTime),
			},
			{
				Name:        "data_locale",
				Description: "The data locale of the stream analytics job.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.StreamingJob.Properties.DataLocale")},
			{
				Name:        "etag",
				Description: "An unique read-only string that changes whenever the resource is updated.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.StreamingJob.Properties.Etag")},
			{
				Name:        "events_late_arrival_max_delay_in_seconds",
				Description: "The maximum tolerable delay in seconds where events arriving late could be included.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.StreamingJob.Properties.EventsLateArrivalMaxDelayInSeconds")},
			{
				Name:        "events_out_of_order_max_delay_in_seconds",
				Description: "The maximum tolerable delay in seconds where out-of-order events can be adjusted to be back in order.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.StreamingJob.Properties.EventsOutOfOrderMaxDelayInSeconds")},
			{
				Name:        "events_out_of_order_policy",
				Description: "Indicates the policy to apply to events that arrive out of order in the input event stream.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.StreamingJob.Properties.EventsOutOfOrderPolicy")},
			{
				Name:        "last_output_event_time",
				Description: "Indicating the last output event time of the streaming job or null indicating that output has not yet been produced.",
				Type:        proto.ColumnType_TIMESTAMP,

				Transform: transform.FromField("Description.StreamingJob.Properties.LastOutputEventTime").Transform(convertDateToTime),
			},
			{
				Name:        "output_error_policy",
				Description: "Indicates the policy to apply to events that arrive at the output and cannot be written to the external storage due to being malformed (missing column values, column values of wrong type or size).",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.StreamingJob.Properties.OutputErrorPolicy")},
			{
				Name:        "output_start_mode",
				Description: "This property should only be utilized when it is desired that the job be started immediately upon creation.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.StreamingJob.Properties.OutputStartMode")},
			{
				Name:        "output_start_time",
				Description: "Indicates the starting point of the output event stream, or null to indicate that the output event stream will start whenever the streaming job is started.",
				Type:        proto.ColumnType_TIMESTAMP,

				Transform: transform.FromField("Description.StreamingJob.Properties.OutputStartTime").Transform(convertDateToTime),
			},
			{
				Name:        "sku_name",
				Description: "Describes the sku name of the streaming job.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.StreamingJob.Properties.SKU.Name")},
			{
				Name:        "diagnostic_settings",
				Description: "A list of active diagnostic settings for the streaming job.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DiagnosticSettingsResources")},
			{
				Name:        "functions",
				Description: "A list of one or more functions for the streaming job.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.StreamingJob.Properties.Functions")},
			{
				Name:        "inputs",
				Description: "A list of one or more inputs to the streaming job.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.StreamingJob.Properties.Inputs")},
			{
				Name:        "outputs",
				Description: "A list of one or more outputs for the streaming job.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.StreamingJob.Properties.Outputs")},
			{
				Name:        "transformation",
				Description: "Indicates the query and the number of streaming units to use for the streaming job.",
				Type:        proto.ColumnType_JSON,

				// Steampipe standard columns
				Transform: transform.FromField("Description.StreamingJob.Properties.Transformation")},

			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.StreamingJob.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.StreamingJob.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				// Azure standard columns

				Transform: transform.FromField("Description.StreamingJob.ID").Transform(idToAkas),
			},

			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.StreamingJob.Location").Transform(formatRegion).Transform(toLower),
			},
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,

				//// LIST FUNCTION

				Transform: transform.

					// Check if context has been cancelled or if the limit has been hit (if specified)
					// if there is a limit, it will return the number of rows required to reach this limit
					FromField("Description.ResourceGroup")},
		}),
	}
}

// Check if context has been cancelled or if the limit has been hit (if specified)
// if there is a limit, it will return the number of rows required to reach this limit

//// HYDRATE FUNCTIONS

// Create session

// Return nil, if no input provide

// Create session

// If we return the API response directly, the output only gives
// the contents of DiagnosticSettings
