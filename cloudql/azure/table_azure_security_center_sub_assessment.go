package azure

import (
	"context"
	"encoding/json"
	"regexp"
	"strings"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAzureSecurityCenterSubAssessment(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_security_center_sub_assessment",
		Description: "Azure Security Center Sub Assessment",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListSecurityCenterSubAssessment,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "The resource id.",
				Transform:   transform.FromField("Description.SubAssessment.ID"),
			},
			{
				Name:        "name",
				Description: "The resource name.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SubAssessment.Name")},
			{
				Name:        "type",
				Description: "The resource type.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SubAssessment.Type")},
			{
				Name:        "assessment_name",
				Description: "The assessment name.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.From(extractAssessmentName),
			},
			{
				Name:        "category",
				Description: "Category of the sub-assessment.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SubAssessment.Properties.Category")},
			{
				Name:        "description",
				Description: "Human readable description of the assessment status.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SubAssessment.Properties.Description")},
			{
				Name:        "display_name",
				Description: "User friendly display name of the sub-assessment.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SubAssessment.Properties.DisplayName")},
			{
				Name:        "impact",
				Description: "Description of the impact of this sub-assessment.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SubAssessment.Properties.Impact")},
			{
				Name:        "remediation",
				Description: "Information on how to remediate this sub-assessment.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SubAssessment.Properties.Remediation")},
			{
				Name:        "time_generated",
				Description: "The date and time the sub-assessment was generated.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SubAssessment.Properties.TimeGenerated")},
			{
				Name:        "assessed_resource_type",
				Description: "Details of the sub-assessment.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.From(extractAssessedResourceType),
			},
			{
				Name:        "container_registry_vulnerability_properties",
				Description: "ContainerRegistryVulnerabilityProperties details of the resource that was assessed.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(extractContainerRegistryVulnerabilityProperties),
			},
			{
				Name:        "resource_details",
				Description: "Details of the resource that was assessed.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(extractResourceDetails),
			},
			{
				Name:        "status",
				Description: "The status of the sub-assessment.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(extractSubAssessmentStatus),
			},
			{
				Name:        "server_vulnerability_properties",
				Description: "ServerVulnerabilityProperties details of the resource that was assessed.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(extractServerVulnerabilityProperties),
			},
			{
				Name:        "sql_server_vulnerability_properties",
				Description: "SQLServerVulnerabilityProperties details of the resource that was assessed.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(extractSQLServerVulnerabilityProperties),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SubAssessment.Name")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.SubAssessment.ID").Transform(idToAkas),
			},
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,

				Transform: transform.

					//// TRANSFORM FUNCTIONS
					FromField("Description.ResourceGroup")},
		}),
	}
}

func extractAssessmentName(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	subAssessment := d.HydrateItem.(opengovernance.SecurityCenterSubAssessment).Description.SubAssessment
	if subAssessment.ID == nil {
		return nil, nil
	}
	r := regexp.MustCompile(`\bassessments\b`)
	splitStr := r.Split(*subAssessment.ID, len(*subAssessment.ID))[1]
	assessmentName := strings.Split(splitStr, "/")[1]
	return assessmentName, nil
}

func extractSQLServerVulnerabilityProperties(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	subAssessment := d.HydrateItem.(opengovernance.SecurityCenterSubAssessment).Description.SubAssessment
	additionalData := subAssessment.Properties.AdditionalData
	if additionalData == nil {
		return nil, nil
	}
	objectMap := additionalData.GetAdditionalData()

	jsonStr, err := json.Marshal(objectMap)
	if err != nil {
		return nil, nil
	}
	return string(jsonStr), nil

}

func extractContainerRegistryVulnerabilityProperties(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	subAssessment := d.HydrateItem.(opengovernance.SecurityCenterSubAssessment).Description.SubAssessment
	additionalData := subAssessment.Properties.AdditionalData
	if additionalData == nil {
		return nil, nil
	}
	objectMap := additionalData.GetAdditionalData()

	jsonStr, err := json.Marshal(objectMap)
	if err != nil {
		return nil, nil
	}
	return string(jsonStr), nil

}

func extractServerVulnerabilityProperties(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	subAssessment := d.HydrateItem.(opengovernance.SecurityCenterSubAssessment).Description.SubAssessment
	additionalData := subAssessment.Properties.AdditionalData
	if additionalData == nil {
		return nil, nil
	}
	objectMap := additionalData.GetAdditionalData()

	jsonStr, err := json.Marshal(objectMap)
	if err != nil {
		return nil, nil
	}
	return string(jsonStr), nil

}

func extractAssessedResourceType(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	subAssessment := d.HydrateItem.(opengovernance.SecurityCenterSubAssessment).Description.SubAssessment
	additional := subAssessment.Properties.AdditionalData
	if additional == nil {
		return nil, nil
	}
	additionalData := additional.GetAdditionalData()
	if additionalData == nil {
		return nil, nil
	}
	return additionalData.AssessedResourceType, nil
}

func extractResourceDetails(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	subAssessment := d.HydrateItem.(opengovernance.SecurityCenterSubAssessment).Description.SubAssessment
	resourceDetails := subAssessment.Properties.ResourceDetails
	if resourceDetails == nil {
		return nil, nil
	}
	azureResourceDetails := resourceDetails.GetResourceDetails()
	return azureResourceDetails, nil
}

func extractSubAssessmentStatus(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	subAssessment := d.HydrateItem.(opengovernance.SecurityCenterSubAssessment).Description.SubAssessment
	subAssessmentStatus := subAssessment.Properties.Status
	objectMap := make(map[string]interface{})
	if subAssessmentStatus.Cause != nil {
		objectMap["Cause"] = *subAssessmentStatus.Cause
	}
	if subAssessmentStatus.Code != nil {
		if *subAssessmentStatus.Code != "" {
			objectMap["Code"] = subAssessmentStatus.Code
		}
	}
	if subAssessmentStatus.Description != nil {
		objectMap["Description"] = *subAssessmentStatus.Description
	}
	if subAssessmentStatus.Severity != nil {
		if *subAssessmentStatus.Severity != "" {
			objectMap["Severity"] = subAssessmentStatus.Severity
		}
	}
	jsonStr, err := json.Marshal(objectMap)
	if err != nil {
		return nil, err
	}
	return string(jsonStr), nil
}
