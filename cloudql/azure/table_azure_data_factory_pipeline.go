package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAzureDataFactoryPipeline(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_data_factory_pipeline",
		Description: "Azure Data Factory Pipeline",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "resource_group", "factory_name"}),
			Hydrate:    opengovernance.GetDataFactoryPipeline,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "ResourceGroupNotFound", "404"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListDataFactoryPipeline,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The resource name.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Pipeline.Name")},
			{
				Name:        "id",
				Description: "The resource identifier.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Pipeline.ID"),
			},
			{
				Name:        "factory_name",
				Description: "Name of the factory the pipeline belongs.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Factory.Name")},
			{
				Name:        "description",
				Description: "The description of the pipeline.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Pipeline.Properties.Description")},
			{
				Name:        "etag",
				Description: "An unique read-only string that changes whenever the resource is updated.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Pipeline.Etag")},
			{
				Name:        "type",
				Description: "The resource type.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Pipeline.Type")},
			{
				Name:        "concurrency",
				Description: "The max number of concurrent runs for the pipeline.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Pipeline.Properties.Concurrency")},
			{
				Name:        "pipeline_folder",
				Description: "The folder that this Pipeline is in. If not specified, Pipeline will appear at the root level.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Pipeline.Properties.Folder.Name")},
			{
				Name:        "activities",
				Description: "A list of activities in pipeline.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Pipeline.Properties.Activities")},
			{
				Name:        "annotations",
				Description: "A list of tags that can be used for describing the Pipeline.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Pipeline.Properties.Annotations")},
			{
				Name:        "parameters",
				Description: "A list of parameters for pipeline.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Pipeline.Properties.Parameters")},
			{
				Name:        "pipeline_policy",
				Description: "Pipeline ElapsedTime Metric Policy.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Pipeline.Properties.Policy")},
			{
				Name:        "variables",
				Description: "A list of variables for pipeline.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Pipeline.Properties.Variables")},
			{
				Name:        "run_dimensions",
				Description: "Dimensions emitted by Pipeline.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Pipeline.Properties.RunDimensions")},
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Pipeline.Name")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.Pipeline.ID").Transform(idToAkas),
			},
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ResourceGroup")},
		}),
	}
}
