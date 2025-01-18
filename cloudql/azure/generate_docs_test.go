package azure

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"html"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGenerateDocs(t *testing.T) {
	plg := Plugin(context.Background())

	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	pathPrefix := "./docs/table_def"
	if strings.HasSuffix(currentDir, "azure") {
		pathPrefix = "." + pathPrefix
	}

	for _, table := range plg.TableMap {
		doc := `# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
`
		for _, column := range table.Columns {
			desc := column.Description
			desc = html.EscapeString(desc)
			doc += fmt.Sprintf(`	<tr><td>%s</td><td>%s</td></tr>
`, column.Name, desc)
		}

		doc += "</table>"

		err := os.WriteFile(pathPrefix+"/"+table.Name+".md", []byte(doc), os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
}

func TestGenerateTableList(t *testing.T) {
	var tablesFiles []string
	err := filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		if strings.HasPrefix(info.Name(), "table_") {
			name := info.Name()
			name = name[6:]
			name = strings.ReplaceAll(name, ".go", "")
			if name == "aws_api_gateway_api_authorizer" {
				name = "aws_api_gateway_authorizer"
			}
			tablesFiles = append(tablesFiles, name)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	plg := Plugin(context.Background())

	awsResourceType, err := os.ReadFile("../../../opengovernance-deploy/opengovernance/inventory-data/azure-resource-types.json")
	if err != nil {
		panic(err)
	}

	var rts []ResourceType
	err = json.Unmarshal(awsResourceType, &rts)
	if err != nil {
		panic(err)
	}

	for _, rt := range rts {
		exists := false
		for idx, _ := range plg.TableMap {
			if rt.SteampipeTable == idx {
				exists = true
			}
		}
		if strings.HasPrefix(rt.SteampipeTable, "azuread") {
			continue
		}
		if !exists {
			panic("rt " + rt.SteampipeTable + " does not exists")
		}
	}

	for idx, _ := range plg.TableMap {
		exists := false
		for _, rt := range rts {
			if rt.SteampipeTable == idx {
				exists = true
			}
		}
		if !exists {
			fmt.Println(idx + " not supported")
		}
	}

	var cv [][]string
	for _, t := range tablesToPopulate {
		resourceName := ""
		status := ""
		for _, rt := range rts {
			if rt.SteampipeTable == t {
				resourceName = rt.ResourceName
				status = rt.Discovery
			}
		}
		cv = append(cv, []string{t, resourceName, status})
	}
	csvfiler, err := os.Create("tableInformation-part2.csv")
	if err != nil {
		panic(err)
	}
	defer csvfiler.Close()

	csvWriter := csv.NewWriter(csvfiler)
	err = csvWriter.WriteAll(cv)
	if err != nil {
		panic(err)
	}
	csvWriter.Flush()

}

type ResourceType struct {
	ResourceName         string
	ResourceLabel        string
	Category             []string
	Tags                 map[string][]string
	TagsString           string `json:"-"`
	ServiceName          string
	ListDescriber        string
	GetDescriber         string
	TerraformName        []string
	TerraformNameString  string `json:"-"`
	TerraformServiceName string
	Discovery            string
	IgnoreSummarize      bool
	SteampipeTable       string
	Model                string
}

var tablesToPopulate = []string{
	"azure_alert_management",
	"azure_api_management",
	"azure_app_configuration",
	"azure_app_service_environment",
	"azure_app_service_web_app",
	"azure_app_service_plan",
	"azure_app_service_web_app_slot",
	"azure_application_security_group",
	"azure_application_gateway",
	"azure_automation_account",
	"azure_automation_variable",
	"azure_bastion_host",
	"azure_batch_account",
	"azure_cognitive_account",
	"azure_compute_availability_set",
	"azure_compute_disk_encryption_set",
	"azure_compute_image",
	"azure_compute_snapshot",
	"azure_compute_ssh_key",
	"azure_compute_virtual_machine",
	"azure_compute_virtual_machine_scale_set",
	"azure_compute_virtual_machine_scale_set_network_interface",
	"azure_compute_virtual_machine_scale_set_vm",
	"azure_app_service_function_app",
	"azure_container_group",
	"azure_container_registry",
	"azure_cosmosdb_account",
	"azure_cosmosdb_mongo_collection",
	"azure_cosmosdb_mongo_database",
	"azure_cosmosdb_restorable_database_account",
	"azure_cosmosdb_sql_database",
	"azure_data_factory",
	"azure_data_factory_dataset",
	"azure_data_factory_pipeline",
	"azure_data_lake_analytics_account",
	"azure_data_lake_store",
	"azure_databox_edge_device",
	"azure_databricks_workspace",
	"azure_dns_zone",
	"azure_eventgrid_domain",
	"azure_eventgrid_topic",
	"azure_eventhub_namespace",
	"azure_express_route_circuit",
	"azure_firewall",
	"azure_firewall_policy",
	"azure_frontdoor",
	"azure_hdinsight_cluster",
	"azure_healthcare_service",
	"azure_hpc_cache",
	"azure_hybrid_compute_machine",
	"azure_hybrid_kubernetes_connected_cluster",
	"azure_iothub",
	"azure_iothub_dps",
	"azure_key_vault",
	"azure_key_vault_key",
	"azure_key_vault_managed_hardware_security_module",
	"azure_key_vault_secret",
	"azure_kubernetes_cluster",
	"azure_kusto_cluster",
	"azure_lb",
	"azure_lb_backend_address_pool",
	"azure_lb_nat_rule",
	"azure_lb_outbound_rule",
	"azure_lb_probe",
	"azure_lb_rule",
	"azure_location",
	"azure_log_alert",
	"azure_log_profile",
	"azure_logic_app_workflow",
	"azure_machine_learning_workspace",
	"azure_management_group",
	"azure_management_lock",
	"azure_mariadb_server",
	"azure_monitor_activity_log_event",
	"azure_mssql_elasticpool",
	"azure_mssql_managed_instance",
	"azure_mssql_virtual_machine",
	"azure_mysql_flexible_server",
	"azure_mysql_server",
	"azure_nat_gateway",
	"azure_network_interface",
	"azure_network_security_group",
	"azure_network_watcher",
	"azure_network_watcher_flow_log",
	"azure_policy_assignment",
	"azure_policy_definition",
	"azure_postgresql_flexible_server",
	"azure_postgresql_server",
	"azure_private_dns_zone",
	"azure_provider",
	"azure_public_ip",
	"azure_recovery_services_backup_job",
	"azure_recovery_services_vault",
	"azure_redis_cache",
	"azure_resource_group",
	"azure_role_assignment",
	"azure_role_definition",
	"azure_route_table",
	"azure_search_service",
	"azure_security_center_auto_provisioning",
	"azure_security_center_automation",
	"azure_security_center_contact",
	"azure_security_center_jit_network_access_policy",
	"azure_security_center_setting",
	"azure_security_center_sub_assessment",
	"azure_security_center_subscription_pricing",
	"azure_service_fabric_cluster",
	"azure_storage_blob",
	"azure_storage_blob_service",
	"azure_storage_container",
	"azure_storage_queue",
	"azure_storage_share_file",
	"azure_storage_table",
	"azure_stream_analytics_job",
	"azure_subnet",
	"azure_subscription",
	"azure_synapse_workspace",
	"azure_tenant",
	"azure_virtual_network",
	"azure_virtual_network_gateway",
	"azure_compute_disk",
	"azure_servicebus_namespace",
	"azure_signalr_service",
	"azure_spring_cloud_service",
	"azure_sql_database",
	"azure_sql_server",
	"azure_storage_account"}
