// describer.go
package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/opengovern/og-describer-fly/global"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	model "github.com/opengovern/og-describer-fly/discovery/pkg/models"
	"github.com/opengovern/og-describer-fly/discovery/pkg/orchestrator"
	"github.com/opengovern/og-describer-fly/discovery/provider"
	"github.com/opengovern/og-util/pkg/describe"
	"github.com/opengovern/og-util/pkg/es"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var (
	resourceType string
	outputFile   string
)

// describerCmd represents the describer command
var describerCmd = &cobra.Command{
	Use:   "describer",
	Short: "A brief description of your command",
	RunE: func(cmd *cobra.Command, args []string) error {

		// Open the output file
		file, err := os.Create(outputFile)
		if err != nil {
			return fmt.Errorf("failed to create file: %w", err)
		}
		defer file.Close() // Ensure the file is closed at the end

		job := describe.DescribeJob{
			JobID:           uint(uuid.New().ID()),
			ResourceType:    resourceType,
			IntegrationID:   "",
			ProviderID:      "",
			DescribedAt:     time.Now().UnixMilli(),
			IntegrationType: global.IntegrationTypeLower,
			CipherText:      "",
			IntegrationLabels: map[string]string{
				"AppName": "AppName",
			},
			IntegrationAnnotations: nil,
		}

		ctx := context.Background()
		logger, _ := zap.NewProduction()

		creds, err := provider.AccountCredentialsFromMap(map[string]any{
			"token": "FlyV1 fm2_lJPECAAAAAAACBDXxBD7ANdkUElUGPZn1cfDsGSEwrVodHRwczovL2FwaS5mbHkuaW8vdjGUAJLOAA42Kh8Lk7lodHRwczovL2FwaS5mbHkuaW8vYWFhL3YxxDyBGaQX+lIsNtJCDfYhXNuQ0aheXzmSvd9zLdiYBGjTDHtuize88AdnHa0oDChKF/p4ZGZtORrgG1uHkDXETgmBIP42/CEBs+lwwk+JCAARG9iV1ybL1XE6nD+fB+7Hf5s3rBcYDPpMcQK2LcRXhtEaSGQ4c+54rlAARurRxv6Veok0I9iUyY/h0j//YsQgngnw3Qq8VM9fv/+YwSeLTuJNn9z2YYET69PxP7tiRKY=,fm2_lJPETgmBIP42/CEBs+lwwk+JCAARG9iV1ybL1XE6nD+fB+7Hf5s3rBcYDPpMcQK2LcRXhtEaSGQ4c+54rlAARurRxv6Veok0I9iUyY/h0j//YsQQ2TgFVddqJpqyIC/vPEbqlsO5aHR0cHM6Ly9hcGkuZmx5LmlvL2FhYS92MZgEks5nisp/zmeUPV0XzgANvzsKkc4ADb87DMQQUBIxjBKv1wDjYm5L+ZmW38QgW5IvR7f6BbHLUKDQ++EUGgYCuikbnE7P5cTBRRfT1CE=",
		})
		if err != nil {
			return fmt.Errorf(" account credentials: %w", err)
		}

		additionalParameters, err := provider.GetAdditionalParameters(job)
		if err != nil {
			return err
		}
		plg := global.Plugin()

		f := func(resource model.Resource) error {
			if resource.Description == nil {
				return nil
			}
			descriptionJSON, err := json.Marshal(resource.Description)
			if err != nil {
				return fmt.Errorf("failed to marshal description: %w", err)
			}
			descriptionJSON, err = trimJsonFromEmptyObjects(descriptionJSON)
			if err != nil {
				return fmt.Errorf("failed to trim json: %w", err)
			}

			metadata, err := provider.GetResourceMetadata(job, resource)
			if err != nil {
				return fmt.Errorf("failed to get resource metadata")
			}
			err = provider.AdjustResource(job, &resource)
			if err != nil {
				return fmt.Errorf("failed to adjust resource metadata")
			}

			desc := resource.Description
			err = json.Unmarshal(descriptionJSON, &desc)
			if err != nil {
				return fmt.Errorf("unmarshal description: %v", err.Error())
			}

			if plg != nil {
				_, _, err = global.ExtractTagsAndNames(logger, plg, job.ResourceType, resource)
				if err != nil {
					logger.Error("failed to build tags for service", zap.Error(err), zap.String("resourceType", job.ResourceType), zap.Any("resource", resource))
				}
			}

			var description any
			err = json.Unmarshal([]byte(descriptionJSON), &description)
			if err != nil {
				logger.Error("failed to parse resource description json", zap.Error(err))
				return fmt.Errorf("failed to parse resource description json")
			}

			res := es.Resource{
				PlatformID:      fmt.Sprintf("%s:::%s:::%s", job.IntegrationID, job.ResourceType, resource.UniqueID()),
				ResourceID:      resource.UniqueID(),
				ResourceName:    resource.Name,
				Description:     description,
				IntegrationType: global.IntegrationName,
				ResourceType:    strings.ToLower(job.ResourceType),
				IntegrationID:   job.IntegrationID,
				Metadata:        metadata,
				DescribedAt:     job.DescribedAt,
				DescribedBy:     strconv.FormatUint(uint64(job.JobID), 10),
			}

			// Write the resource JSON to the file
			resJSON, err := json.Marshal(res)
			if err != nil {
				return fmt.Errorf("failed to marshal resource JSON: %w", err)
			}
			_, err = file.Write(resJSON)
			if err != nil {
				return fmt.Errorf("failed to write to file: %w", err)
			}
			_, err = file.Write([]byte(",\n")) // Add a newline for readability
			if err != nil {
				return fmt.Errorf("failed to write newline to file: %w", err)
			}

			return nil
		}
		clientStream := (*model.StreamSender)(&f)

		err = orchestrator.GetResources(
			ctx,
			logger,
			job.ResourceType,
			job.TriggerType,
			creds,
			additionalParameters,
			clientStream,
		)
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	describerCmd.Flags().StringVar(&resourceType, "resourceType", "", "Resource type")
	describerCmd.Flags().StringVar(&outputFile, "outputFile", "output.json", "File to write JSON outputs")
}

func trimJsonFromEmptyObjects(input []byte) ([]byte, error) {
	unknownData := map[string]any{}
	err := json.Unmarshal(input, &unknownData)
	if err != nil {
		return nil, err
	}
	trimEmptyMaps(unknownData)
	return json.Marshal(unknownData)
}

func trimEmptyMaps(input map[string]any) {
	for key, value := range input {
		switch value.(type) {
		case map[string]any:
			if len(value.(map[string]any)) != 0 {
				trimEmptyMaps(value.(map[string]any))
			}
			if len(value.(map[string]any)) == 0 {
				delete(input, key)
			}
		}
	}
}
