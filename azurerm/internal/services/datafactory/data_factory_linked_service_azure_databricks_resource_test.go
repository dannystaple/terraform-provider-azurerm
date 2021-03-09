package datafactory_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/azure"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/acceptance"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/acceptance/check"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/clients"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/utils"
)

type LinkedServiceDatabricksResource struct {
}

func TestAccDataFactoryLinkedServiceDatabricks_authViaMSI(t *testing.T) {
	// Build random test data.
	data := acceptance.BuildTestData(t, "azurerm_data_factory_linked_service_azure_databricks", "test")

	//Create an instance of this class so we can reference the functions.
	r := LinkedServiceDatabricksResource{}

	//Execute a test case
	data.ResourceTest(t, r, []resource.TestStep{
		{
			//Define the configuration to use the resource definition returned by the "basic" function (below)
			Config: r.authentication_msi(data),
			// Create a composition of validations to perform post-creation
			Check: resource.ComposeTestCheckFunc(
				// This verifies that the resource now exists in Azure (i.e. Creation was successful)
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		// ?? Is this just validating that the service_principal_key exists on the resource. This should be true because this was not created with a managed identity.
		// If it was created with a managed identity, then this shouldn't exist?
		data.ImportStep(""),
	})
}

func TestAccDataFactoryLinkedServiceDatabricks_authViaAccessToken(t *testing.T) {
	// Build random test data.
	data := acceptance.BuildTestData(t, "azurerm_data_factory_linked_service_azure_databricks", "test")

	//Create an instance of this class so we can reference the functions.
	r := LinkedServiceDatabricksResource{}

	//Execute a test case
	data.ResourceTest(t, r, []resource.TestStep{
		{
			//Define the configuration to use the resource definition returned by the "basic" function (below)
			Config: r.authentication_access_token(data),
			// Create a composition of validations to perform post-creation
			Check: resource.ComposeTestCheckFunc(
				// This verifies that the resource now exists in Azure (i.e. Creation was successful)
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		// ?? Is this just validating that the service_principal_key exists on the resource. This should be true because this was not created with a managed identity.
		// If it was created with a managed identity, then this shouldn't exist?
		data.ImportStep(""),
	})
}

func TestAccDataFactoryLinkedServiceDatabricks_authViaKeyVault(t *testing.T) {
	// Build random test data.
	data := acceptance.BuildTestData(t, "azurerm_data_factory_linked_service_azure_databricks", "test")

	//Create an instance of this class so we can reference the functions.
	r := LinkedServiceDatabricksResource{}

	//Execute a test case
	data.ResourceTest(t, r, []resource.TestStep{
		{
			//Define the configuration to use the resource definition returned by the "basic" function (below)
			Config: r.authentication_key_vault(data),
			// Create a composition of validations to perform post-creation
			Check: resource.ComposeTestCheckFunc(
				// This verifies that the resource now exists in Azure (i.e. Creation was successful)
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		// ?? Is this just validating that the service_principal_key exists on the resource. This should be true because this was not created with a managed identity.
		// If it was created with a managed identity, then this shouldn't exist?
		data.ImportStep(""),
	})
}

func TestAccDataFactoryLinkedServiceDatabricks_newClusterConfig(t *testing.T) {
	// Build random test data.
	data := acceptance.BuildTestData(t, "azurerm_data_factory_linked_service_azure_databricks", "test")

	//Create an instance of this class so we can reference the functions.
	r := LinkedServiceDatabricksResource{}

	//Execute a test case
	data.ResourceTest(t, r, []resource.TestStep{
		{
			//Define the configuration to use the resource definition returned by the "basic" function (below)
			Config: r.newClusterConfig(data),
			// Create a composition of validations to perform post-creation
			Check: resource.ComposeTestCheckFunc(
				// This verifies that the resource now exists in Azure (i.e. Creation was successful)
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("new_cluster_config.0.cluster_version").HasValue("5.5.x-gpu-scala2.11"),
				check.That(data.ResourceName).Key("new_cluster_config.0.node_type").HasValue("Standard_NC12"),
				check.That(data.ResourceName).Key("new_cluster_config.0.number_of_workers").HasValue("5"),
				check.That(data.ResourceName).Key("new_cluster_config.0.driver_node_type").HasValue("Standard_NC13"),
				check.That(data.ResourceName).Key("new_cluster_config.0.log_destination").HasValue("dbfs:/logs"),
				check.That(data.ResourceName).Key("new_cluster_config.0.init_scripts.#").HasValue("2"),
				check.That(data.ResourceName).Key("new_cluster_config.0.custom_tags.sct1").HasValue("sct_value_1"),
				check.That(data.ResourceName).Key("new_cluster_config.0.custom_tags.sct2").HasValue("sct_value_2"),
				check.That(data.ResourceName).Key("new_cluster_config.0.spark_config.sc1").HasValue("sc_value_1"),
				check.That(data.ResourceName).Key("new_cluster_config.0.spark_config.sc2").HasValue("sc_value_2"),
				check.That(data.ResourceName).Key("new_cluster_config.0.spark_environment_variables.sev1").HasValue("sev_value_1"),
				check.That(data.ResourceName).Key("new_cluster_config.0.spark_environment_variables.sev2").HasValue("sev_value_2"),
			),
		},
		// ?? Is this just validating that the service_principal_key exists on the resource. This should be true because this was not created with a managed identity.
		// If it was created with a managed identity, then this shouldn't exist?
		data.ImportStep(""),
	})
}
func TestAccDataFactoryLinkedServiceDatabricks_update(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_data_factory_linked_service_azure_databricks", "test")
	r := LinkedServiceDatabricksResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.update1(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("parameters.%").HasValue("2"),
				check.That(data.ResourceName).Key("parameters.key1").HasValue("u1k1"),
				check.That(data.ResourceName).Key("parameters.key2").HasValue("u1k2"),
				check.That(data.ResourceName).Key("annotations.#").HasValue("2"),
				check.That(data.ResourceName).Key("annotations.0").HasValue("a1"),
				check.That(data.ResourceName).Key("annotations.1").HasValue("a2"),
				check.That(data.ResourceName).Key("description").HasValue("Initial Description"),
				check.That(data.ResourceName).Key("new_cluster_config.0.cluster_version").HasValue("5.5.x-gpu-scala2.11"),
				check.That(data.ResourceName).Key("new_cluster_config.0.node_type").HasValue("Standard_NC12"),
				check.That(data.ResourceName).Key("new_cluster_config.0.number_of_workers").HasValue("1:10"),
				check.That(data.ResourceName).Key("new_cluster_config.0.driver_node_type").HasValue("Standard_NC12"),
				check.That(data.ResourceName).Key("new_cluster_config.0.log_destination").HasValue("dbfs:/logs"),
				check.That(data.ResourceName).Key("new_cluster_config.0.init_scripts.#").HasValue("2"),
				check.That(data.ResourceName).Key("new_cluster_config.0.custom_tags.sct1").HasValue("sct_value_1"),
				check.That(data.ResourceName).Key("new_cluster_config.0.custom_tags.sct2").HasValue("sct_value_2"),
				check.That(data.ResourceName).Key("new_cluster_config.0.spark_config.sc1").HasValue("sc_value_1"),
				check.That(data.ResourceName).Key("new_cluster_config.0.spark_config.sc2").HasValue("sc_value_2"),
				check.That(data.ResourceName).Key("new_cluster_config.0.spark_environment_variables.sev1").HasValue("sev_value_1"),
				check.That(data.ResourceName).Key("new_cluster_config.0.spark_environment_variables.sev2").HasValue("sev_value_2"),
			),
		},
		{
			Config: r.update2(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("parameters.%").HasValue("2"),
				check.That(data.ResourceName).Key("parameters.key1").HasValue("u2k1"),
				check.That(data.ResourceName).Key("parameters.key2").HasValue("u2k2"),
				check.That(data.ResourceName).Key("annotations.#").HasValue("2"),
				check.That(data.ResourceName).Key("annotations.0").HasValue("b1"),
				check.That(data.ResourceName).Key("annotations.1").HasValue("b2"),
				check.That(data.ResourceName).Key("description").HasValue("Updated Description"),
				check.That(data.ResourceName).Key("new_cluster_config.0.cluster_version").HasValue("6.5.x-gpu-scala2.11"),
				check.That(data.ResourceName).Key("new_cluster_config.0.node_type").HasValue("Standard_NC20"),
				check.That(data.ResourceName).Key("new_cluster_config.0.number_of_workers").HasValue("5"),
				check.That(data.ResourceName).Key("new_cluster_config.0.driver_node_type").HasValue("Standard_NC13"),
				check.That(data.ResourceName).Key("new_cluster_config.0.log_destination").HasValue("dbfs:/logs_updated"),
				check.That(data.ResourceName).Key("new_cluster_config.0.init_scripts.#").HasValue("3"),
				check.That(data.ResourceName).Key("new_cluster_config.0.custom_tags.sct1").HasValue("updated_sct_value_1"),
				check.That(data.ResourceName).Key("new_cluster_config.0.custom_tags.sct2").HasValue("updated_sct_value_2"),
				check.That(data.ResourceName).Key("new_cluster_config.0.spark_config.sc1").HasValue("updated_sc_value_1"),
				check.That(data.ResourceName).Key("new_cluster_config.0.spark_config.sc2").HasValue("updated_sc_value_2"),
				check.That(data.ResourceName).Key("new_cluster_config.0.spark_environment_variables.sev1").HasValue("updated_sev_value_1"),
				check.That(data.ResourceName).Key("new_cluster_config.0.spark_environment_variables.sev2").HasValue("updated_sev_value_2"),
			),
		},
		data.ImportStep(""),
	})
}

func (t LinkedServiceDatabricksResource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	id, err := azure.ParseAzureResourceID(state.ID)
	if err != nil {
		return nil, err
	}
	resourceGroup := id.ResourceGroup
	dataFactoryName := id.Path["factories"]
	name := id.Path["linkedservices"]

	resp, err := clients.DataFactory.LinkedServiceClient.Get(ctx, resourceGroup, dataFactoryName, name, "")
	if err != nil {
		return nil, fmt.Errorf("reading Data Factory LinkedServiceDatabricksResource (%s): %+v", id, err)
	}

	return utils.Bool(resp.ID != nil), nil
}

func (LinkedServiceDatabricksResource) authentication_msi(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "test" {
  name     = "acctestRG-df-%d"
  location = "%s"
}

resource "azurerm_data_factory" "test" {
  name                = "acctestdf%d"
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name
}

data "azurerm_client_config" "current" {
}

resource "azurerm_data_factory_linked_service_azure_databricks" "test" {
	name                  = "acctestDatabricksLinkedService%d"
	resource_group_name   = azurerm_resource_group.test.name
	data_factory_name     = azurerm_data_factory.test.name
	authentication_msi = {
		workspaceResourceId="/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/test/providers/Microsoft.Databricks/workspaces/testworkspace"
	}
	description				= "Initial description"
	annotations				= ["test1", "test2"]
	existing_cluster_id		= "test"
	adb_domain = "https://adb-111111111.11.azuredatabricks.net"
	
}
`, data.RandomInteger, data.Locations.Primary, data.RandomInteger, data.RandomInteger)
}

func (LinkedServiceDatabricksResource) authentication_access_token(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "test" {
  name     = "acctestRG-df-%d"
  location = "%s"
}

resource "azurerm_data_factory" "test" {
  name                = "acctestdf%d"
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name
}

data "azurerm_client_config" "current" {
}

resource "azurerm_data_factory_linked_service_azure_databricks" "test" {
	name                  = "acctestDatabricksLinkedService%d"
	resource_group_name   = azurerm_resource_group.test.name
	data_factory_name     = azurerm_data_factory.test.name
	authentication_access_token = {
		access_token	= "SomeFakeAccessToken"
	}
	description				= "Initial description"
	annotations				= ["test1", "test2"]

	adb_domain = "https://adb-111111111.11.azuredatabricks.net"
	existing_cluster_id = "1234"

}
`, data.RandomInteger, data.Locations.Primary, data.RandomInteger, data.RandomInteger)
}

func (LinkedServiceDatabricksResource) authentication_key_vault(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

data "azurerm_client_config" "current" {}
// Create the RG
resource "azurerm_resource_group" "test" {
  name     = "acctestRG-df-%d"
  location = "%s"
}

resource "azurerm_data_factory" "test" {
  name                = "acctestdf%d"
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name
}


// Create a key vault so we can setup a KV linked service
resource "azurerm_key_vault" "test" {
	name                = "acctkv%d"
	location            = azurerm_resource_group.test.location
	resource_group_name = azurerm_resource_group.test.name
	tenant_id           = data.azurerm_client_config.current.tenant_id
	sku_name            = "standard"
  }

// Create the KV linked service so we can test out integration the Databricks linked service
resource "azurerm_data_factory_linked_service_key_vault" "test" {
	name                = "linkkv"
	resource_group_name = azurerm_resource_group.test.name
	data_factory_name   = azurerm_data_factory.test.name
	key_vault_id        = azurerm_key_vault.test.id
  }

//   Create a databricks linked service that leveraged the KV linked service for password management
resource "azurerm_data_factory_linked_service_azure_databricks" "test" {
	name                  = "acctestDatabricksLinkedService%d"
	resource_group_name   = azurerm_resource_group.test.name
	data_factory_name     = azurerm_data_factory.test.name
	authentication_key_vault_password {
		linked_service_name = azurerm_data_factory_linked_service_key_vault.test.name
		secret_name         = "secret"
	  }
	description				= "Initial description"
	annotations				= ["test1", "test2"]
	adb_domain = "https://adb-111111111.11.azuredatabricks.net"
	instance_pool = {
		instance_pool_id = "0308-201055-safes631-pool-EHfwukQo",
		number_of_workers = "1",
		cluster_version = "5.5.x-gpu-scala2.11"
	  }
	
}
`, data.RandomInteger, data.Locations.Primary, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (LinkedServiceDatabricksResource) newClusterConfig(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "test" {
  name     = "acctestRG-df-%d"
  location = "%s"
}

resource "azurerm_data_factory" "test" {
  name                = "acctestdf%d"
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name
}

data "azurerm_client_config" "current" {
}

resource "azurerm_data_factory_linked_service_azure_databricks" "test" {
	name                  = "acctestDatabricksLinkedService%d"
	resource_group_name   = azurerm_resource_group.test.name
	data_factory_name     = azurerm_data_factory.test.name
	authentication_msi = {
		workspaceResourceId="/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/test/providers/Microsoft.Databricks/workspaces/testworkspace"
	}
	description				= "Initial description"
	annotations				= ["test1", "test2"]
	adb_domain = "https://adb-111111111.11.azuredatabricks.net"
	new_cluster_config  {
		cluster_version = "5.5.x-gpu-scala2.11"
		number_of_workers = "5"
		node_type = "Standard_NC12"
		driver_node_type = "Standard_NC13"
		log_destination = "dbfs:/logs"
		
		custom_tags = {
			sct1 = "sct_value_1"
			sct2 = "sct_value_2"
		}
		spark_config = {
			 sc1 = "sc_value_1"
			 sc2 = "sc_value_2"
		}
		spark_environment_variables = {
			sev1 = "sev_value_1"
			sev2 = "sev_value_2"
		}

		init_scripts = ["init.sh", "init2.sh"]
	}
	
}
`, data.RandomInteger, data.Locations.Primary, data.RandomInteger, data.RandomInteger)
}

func (LinkedServiceDatabricksResource) update1(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "test" {
  name     = "acctestRG-df-%d"
  location = "%s"
}

resource "azurerm_data_factory" "test" {
  name                = "acctestdf%d"
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name
}


resource "azurerm_data_factory_linked_service_azure_databricks" "test" {
	name                  = "acctestDatabricksLinkedService%d"
	resource_group_name = azurerm_resource_group.test.name
	data_factory_name     = azurerm_data_factory.test.name
    description = "Initial Description"
	annotations = ["a1", "a2"]
	parameters = {
		key1 = "u1k1"
		key2 = "u1k2"
	}
    authentication_msi = {
      workspaceResourceId = "/subscriptions/d111111-1111-1111-1111-111111111111/resourceGroups/Testing-rg-creation/providers/Microsoft.Databricks/workspaces/databricks-test"
    }
    adb_domain = "https://someFakeDomain"

    new_cluster_config {
      node_type = "Standard_NC12"
      cluster_version = "5.5.x-gpu-scala2.11"
      number_of_workers = "1:10"
     
      driver_node_type = "Standard_NC12"
      log_destination = "dbfs:/logs"
      
      custom_tags = {
        sct1 = "sct_value_1"
        sct2 = "sct_value_2"
      }
      spark_config = {
        sc1 = "sc_value_1"
        sc2 = "sc_value_2"
      }
      spark_environment_variables = {
        sev1 = "sev_value_1"
        sev2 = "sev_value_2"
      }

      init_scripts = ["init.sh", "init2.sh"]
    }
}


`, data.RandomInteger, data.Locations.Primary, data.RandomInteger, data.RandomInteger)
}

func (LinkedServiceDatabricksResource) update2(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "test" {
  name     = "acctestRG-df-%d"
  location = "%s"
}

resource "azurerm_data_factory" "test" {
  name                = "acctestdf%d"
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name
}


resource "azurerm_data_factory_linked_service_azure_databricks" "test" {
	name                  = "acctestDatabricksLinkedService%d"
	resource_group_name = azurerm_resource_group.test.name
	data_factory_name     = azurerm_data_factory.test.name
	description = "Updated Description"
	annotations = ["b1", "b2"]
	parameters = {
		key1 = "u2k1"
		key2 = "u2k2"
	}
    authentication_msi = {
      workspaceResourceId = "/subscriptions/d111111-1111-1111-1111-111111111111/resourceGroups/Testing-rg-creation/providers/Microsoft.Databricks/workspaces/databricks-test"
    }
    adb_domain = "https://someFakedomain"

    new_cluster_config {
      node_type = "Standard_NC20"
      cluster_version = "6.5.x-gpu-scala2.11"
      number_of_workers = "5"
     
      driver_node_type = "Standard_NC13"
      log_destination = "dbfs:/logs_updated"
      
      custom_tags = {
        sct1 = "updated_sct_value_1"
        sct2 = "updated_sct_value_2"
      }
      spark_config = {
        sc1 = "updated_sc_value_1"
        sc2 = "updated_sc_value_2"
      }
      spark_environment_variables = {
        sev1 = "updated_sev_value_1"
        sev2 = "updated_sev_value_2"
      }

      init_scripts = ["updated_init.sh", "updated_init2.sh", "updated_init3.sh"]
    }
}


`, data.RandomInteger, data.Locations.Primary, data.RandomInteger, data.RandomInteger)
}
