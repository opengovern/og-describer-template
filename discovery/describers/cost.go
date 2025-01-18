package describers

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement"

	"github.com/opengovern/og-util/pkg/describe/enums"

	"github.com/opengovern/og-describer-azure/discovery/pkg/models"
	model "github.com/opengovern/og-describer-azure/discovery/provider"
)

const serviceNameDimension = "ServiceName"
const publisherTypeDimension = "PublisherType"
const subscriptionDimension = "SubscriptionId"

func cost(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, from time.Time, to time.Time, dimension string) ([]model.CostManagementQueryRow, *string, error) {
	var err error
	clientFactory, err := armcostmanagement.NewClientFactory(cred, nil)
	if err != nil {
		return nil, nil, err
	}
	client := clientFactory.NewQueryClient()

	scope := fmt.Sprintf("subscriptions/%s", subscription)
	groupingType := armcostmanagement.QueryColumnTypeDimension
	groupings := []*armcostmanagement.QueryGrouping{
		{
			Type: &groupingType,
			Name: &dimension,
		},
	}
	if dimension == serviceNameDimension {
		d := publisherTypeDimension
		groupings = append(groupings, &armcostmanagement.QueryGrouping{
			Type: &groupingType,
			Name: &d,
		})
	}

	costAggregationString := "Cost"
	var costs armcostmanagement.QueryResult
	queryFunction := armcostmanagement.FunctionTypeSum
	queryGranularity := armcostmanagement.GranularityTypeDaily
	queryTimeFrame := armcostmanagement.TimeframeTypeCustom
	queryType := armcostmanagement.ExportTypeAmortizedCost
	queryDefinition := armcostmanagement.QueryDefinition{
		Dataset: &armcostmanagement.QueryDataset{
			Aggregation: map[string]*armcostmanagement.QueryAggregation{
				"Cost": {
					Name:     &costAggregationString,
					Function: &queryFunction,
				},
			},
			Granularity: &queryGranularity,
			Grouping:    groupings,
		},
		Timeframe: &queryTimeFrame,
		Type:      &queryType,
		TimePeriod: &armcostmanagement.QueryTimePeriod{
			From: &from,
			To:   &to,
		},
	}
	costsResponse, err := client.Usage(ctx, scope, queryDefinition, nil)
	if err != nil {
		return nil, nil, err
	}
	costs = costsResponse.QueryResult

	mapResult := make([]map[string]any, 0)
	for _, row := range costs.Properties.Rows {
		rowMap := make(map[string]any)
		for i, column := range costs.Properties.Columns {
			rowMap[*column.Name] = row[i]
		}
		mapResult = append(mapResult, rowMap)
	}
	jsonMapResult, err := json.Marshal(mapResult)
	if err != nil {
		return nil, nil, err
	}

	result := make([]model.CostManagementQueryRow, 0, len(mapResult))
	err = json.Unmarshal(jsonMapResult, &result)
	if err != nil {
		return nil, nil, err
	}

	return result, costs.Location, nil
}

func DailyCostByResourceType(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	triggerType := GetTriggerTypeFromContext(ctx)
	from := time.Now().AddDate(0, 0, -7)
	if time.Now().Day() == 6 {
		y, m, _ := time.Now().Date()
		from = time.Date(y, m, 1, 0, 0, 0, 0, time.UTC).AddDate(0, -1, 0)
	}
	if triggerType == enums.DescribeTriggerTypeInitialDiscovery {
		from = time.Now().AddDate(0, -3, -7)
	} else if triggerType == enums.DescribeTriggerTypeCostFullDiscovery {
		from = time.Date(2022, 12, 1, 0, 0, 0, 0, time.UTC)
	}
	to := time.Now()

	var costResult []model.CostManagementQueryRow
	var locationPtr *string
	pageFrom := from
	pageTo := to
	for {
		if pageFrom.Add(4 * 30 * 24 * time.Hour).Before(to) {
			pageTo = pageFrom.Add(4 * 30 * 24 * time.Hour)
		} else {
			pageTo = to
		}

		pageCostResult, pageLocationPtr, err := cost(ctx, cred, subscription, pageFrom, pageTo, serviceNameDimension)
		if err != nil {
			return nil, err
		}

		if pageLocationPtr != nil {
			locationPtr = pageLocationPtr
		}
		costResult = append(costResult, pageCostResult...)

		pageFrom = pageTo

		if pageFrom == to {
			break
		}
	}

	location := "global"
	if locationPtr != nil {
		location = *locationPtr
	}
	var values []models.Resource
	for _, row := range costResult {
		usageDateStr := strconv.FormatInt(int64(row.UsageDate), 10)
		year, month, day := usageDateStr[:4], usageDateStr[4:6], usageDateStr[6:]
		costDate, err := time.Parse("2006-01-02", fmt.Sprintf("%s-%s-%s", year, month, day))
		if err != nil {
			return nil, err
		}

		resource := models.Resource{
			ID:       fmt.Sprintf("resource-cost-%s/%s-%d", subscription, *row.ServiceName, row.UsageDate),
			Location: location,
			Description: model.CostManagementCostByResourceTypeDescription{
				CostManagementCostByResourceType: row,
				CostDateMillis:                   costDate.UnixMilli(),
			},
		}
		if stream != nil {
			if err := (*stream)(resource); err != nil {
				return nil, err
			}
		} else {
			values = append(values, resource)
		}
	}

	return values, nil
}

func DailyCostBySubscription(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	triggerType := GetTriggerTypeFromContext(ctx)
	from := time.Now().AddDate(0, 0, -7)
	if triggerType == enums.DescribeTriggerTypeInitialDiscovery {
		from = time.Now().AddDate(0, -3, -7)
	}
	to := time.Now()

	costResult, locationPtr, err := cost(ctx, cred, subscription, from, to, subscriptionDimension)
	if err != nil {
		return nil, err
	}
	location := "global"
	if locationPtr != nil {
		location = *locationPtr
	}
	var values []models.Resource
	for _, row := range costResult {
		resource := models.Resource{
			ID:       fmt.Sprintf("resource-cost-%s/%d", subscription, row.UsageDate),
			Location: location,
			Description: model.CostManagementCostBySubscriptionDescription{
				CostManagementCostBySubscription: row,
			},
		}
		if stream != nil {
			if err := (*stream)(resource); err != nil {
				return nil, err
			}
		} else {
			values = append(values, resource)
		}
	}

	return values, nil
}
