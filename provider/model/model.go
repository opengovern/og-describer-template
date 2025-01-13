//go:generate go run ../../pkg/sdk/runable/steampipe_es_client_generator/main.go -pluginPath ../../steampipe-plugin-REPLACEME/REPLACEME -file $GOFILE -output ../../pkg/sdk/es/resources_clients.go -resourceTypesFile ../resource_types/resource-types.json

// Implement types for each resource

package model

type Metadata struct{}

type AppJSON struct {
	ID           string `json:"id"`
	MachineCount int    `json:"machine_count,omitempty"`
	Name         string `json:"name"`
	Network      any    `json:"network,omitempty"`
}

type AppDescription struct {
	ID           string
	MachineCount int
	Name         string
	Network      any
}

type AppsResponse struct {
	Apps []AppJSON `json:"apps"`
}
