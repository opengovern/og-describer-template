package describers

import (
	"context"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
	"github.com/opengovern/og-describer-azure/discovery/pkg/models"
	model "github.com/opengovern/og-describer-azure/discovery/provider"
)

func SynapseWorkspace(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armsynapse.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	synapseClient := clientFactory.NewWorkspaceManagedSQLServerVulnerabilityAssessmentsClient()
	client := clientFactory.NewWorkspacesClient()

	monitorClientFactory, err := armmonitor.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	diagnosticClient := monitorClientFactory.NewDiagnosticSettingsClient()

	var values []models.Resource
	pager := client.NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource, err := GetSynapseWorkspace(ctx, synapseClient, diagnosticClient, v)
			if err != nil {
				return nil, err
			}
			if stream != nil {
				if err := (*stream)(*resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, *resource)
			}
		}
	}
	return values, nil
}

func GetSynapseWorkspace(ctx context.Context, synapseClient *armsynapse.WorkspaceManagedSQLServerVulnerabilityAssessmentsClient, diagnosticClient *armmonitor.DiagnosticSettingsClient, config *armsynapse.Workspace) (*models.Resource, error) {
	resourceGroup := strings.Split(*config.ID, "/")[4]

	ignoreAssesment := false
	var synapseListResult []*armsynapse.ServerVulnerabilityAssessment
	pager1 := synapseClient.NewListPager(resourceGroup, *config.Name, nil)
	for pager1.More() {
		pager1, err := pager1.NextPage(ctx)
		if err != nil {
			if strings.Contains(err.Error(), "UnsupportedOperation") {
				ignoreAssesment = true
			} else {
				return nil, err
			}
		}
		synapseListResult = append(synapseListResult, pager1.Value...)
	}

	var serverVulnerabilityAssessments []*armsynapse.ServerVulnerabilityAssessment
	if !ignoreAssesment {
		serverVulnerabilityAssessments = append(serverVulnerabilityAssessments, synapseListResult...)
	}

	var synapseListOp []*armmonitor.DiagnosticSettingsResource
	pager2 := diagnosticClient.NewListPager(*config.ID, nil)
	for pager2.More() {
		page2, err := pager2.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		synapseListOp = append(synapseListOp, page2.Value...)
	}

	resource := models.Resource{
		ID:       *config.ID,
		Name:     *config.Name,
		Location: *config.Location,
		Description: model.SynapseWorkspaceDescription{
			Workspace:                      *config,
			ServerVulnerabilityAssessments: serverVulnerabilityAssessments,
			DiagnosticSettingsResources:    synapseListOp,
			ResourceGroup:                  resourceGroup,
		},
	}
	return &resource, nil
}

func SynapseWorkspaceBigdataPools(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armsynapse.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	bigDataPoolsClient := clientFactory.NewBigDataPoolsClient()
	client := clientFactory.NewWorkspacesClient()

	var values []models.Resource
	pager := client.NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resources, err := ListSynapseWorkspaceBigdataPools(ctx, bigDataPoolsClient, v)
			if err != nil {
				return nil, err
			}
			values = append(values, resources...)
		}
	}
	return values, err
}

func ListSynapseWorkspaceBigdataPools(ctx context.Context, bigDataPoolsClient *armsynapse.BigDataPoolsClient, v *armsynapse.Workspace) ([]models.Resource, error) {
	resourceGroup := strings.Split(*v.ID, "/")[4]

	var values []models.Resource
	pager := bigDataPoolsClient.NewListByWorkspacePager(resourceGroup, *v.Name, nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, bp := range page.Value {
			resource := GetSynapseWorkspaceBigdataPools(ctx, resourceGroup, bp, v)
			values = append(values, *resource)
		}
	}
	return values, nil
}

func GetSynapseWorkspaceBigdataPools(ctx context.Context, resourceGroup string, bp *armsynapse.BigDataPoolResourceInfo, v *armsynapse.Workspace) *models.Resource {
	resource := models.Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: *v.Location,
		Description: model.SynapseWorkspaceBigdatapoolsDescription{
			Workspace:     *v,
			BigDataPool:   *bp,
			ResourceGroup: resourceGroup,
		},
	}
	return &resource
}

func SynapseWorkspaceSqlpools(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armsynapse.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewWorkspacesClient()
	bpClient := clientFactory.NewSQLPoolsClient()

	var values []models.Resource
	pager := client.NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resources, err := ListSynapseWorkspaceSqlpools(ctx, bpClient, v)
			if err != nil {
				return nil, err
			}
			for _, resource := range resources {
				if stream != nil {
					if err := (*stream)(resource); err != nil {
						return nil, err
					}
				} else {
					values = append(values, resource)
				}
			}
		}
	}
	return values, nil
}

func ListSynapseWorkspaceSqlpools(ctx context.Context, bpClient *armsynapse.SQLPoolsClient, v *armsynapse.Workspace) ([]models.Resource, error) {
	resourceGroup := strings.Split(*v.ID, "/")[4]

	var values []models.Resource
	pager := bpClient.NewListByWorkspacePager(resourceGroup, *v.Name, nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			if strings.Contains(err.Error(), "UnsupportedOperation") {
				continue
			}
			return nil, err
		}
		for _, bp := range page.Value {
			resource := GetSynapseWorkspaceSqlpools(ctx, v, bp)
			values = append(values, *resource)
		}
	}
	return values, nil
}

func GetSynapseWorkspaceSqlpools(ctx context.Context, v *armsynapse.Workspace, bp *armsynapse.SQLPool) *models.Resource {
	resourceGroup := strings.Split(*v.ID, "/")[4]

	resource := models.Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: *v.Location,
		Description: model.SynapseWorkspaceSqlpoolsDescription{
			Workspace:     *v,
			SqlPool:       *bp,
			ResourceGroup: resourceGroup,
		},
	}
	return &resource
}
