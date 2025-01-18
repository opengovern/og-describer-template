package describers

import (
	"context"
	"fmt"
	"strings"

	"github.com/opengovern/og-describer-azure/discovery/pkg/models"
	model "github.com/opengovern/og-describer-azure/discovery/provider"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managedservices/armmanagedservices"
	
)

func LighthouseDefinition(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armmanagedservices.NewClientFactory(cred, nil)
	if err != nil {
		return nil, err
	}

	client := clientFactory.NewRegistrationDefinitionsClient()
	scope := fmt.Sprintf("subscriptions/%s", subscription)
	pager := client.NewListPager(scope, nil)

	var resources []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, definition := range page.Value {
			resource := getLighthouseDefinition(ctx, definition, scope)
			if stream != nil {
				if err := (*stream)(*resource); err != nil {
					return nil, err
				}
			} else {
				resources = append(resources, *resource)
			}
		}
	}
	return resources, nil

}

func getLighthouseDefinition(_ context.Context, lighthouseDefinition *armmanagedservices.RegistrationDefinition, scope string) *models.Resource {
	resourceGroup := strings.Split(*lighthouseDefinition.ID, "/")[4]

	resource := models.Resource{
		ID:   *lighthouseDefinition.ID,
		Name: *lighthouseDefinition.Name,
		Type: *lighthouseDefinition.Type,
		Description: model.LighthouseDefinitionDescription{
			LighthouseDefinition: *lighthouseDefinition,
			Scope:                scope,
			ResourceGroup:        resourceGroup,
		},
	}
	return &resource
}

func LighthouseAssignments(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armmanagedservices.NewClientFactory(cred, nil)
	if err != nil {
		return nil, err
	}

	client := clientFactory.NewRegistrationAssignmentsClient()
	scope := fmt.Sprintf("subscriptions/%s", subscription)
	pager := client.NewListPager(scope, nil)

	var resources []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, assignment := range page.Value {
			resource := getLighthouseAssignment(ctx, assignment, scope)
			if stream != nil {
				if err := (*stream)(*resource); err != nil {
					return nil, err
				}
			} else {
				resources = append(resources, *resource)
			}
		}
	}
	return resources, nil

}

func getLighthouseAssignment(_ context.Context, lighthouseAssignment *armmanagedservices.RegistrationAssignment, scope string) *models.Resource {

	resourceGroup := strings.Split(*lighthouseAssignment.ID, "/")[4]

	resource := models.Resource{
		ID:   *lighthouseAssignment.ID,
		Name: *lighthouseAssignment.Name,
		Type: *lighthouseAssignment.Type,
		Description: model.LighthouseAssignmentDescription{
			LighthouseAssignment: *lighthouseAssignment,
			Scope:                scope,
			ResourceGroup:        resourceGroup,
		},
	}
	return &resource

}
