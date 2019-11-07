// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file at
//     https://github.com/Azure/magic-module-specs
//
// ----------------------------------------------------------------------------

package azurerm

import (
    "fmt"
    "strings"
    "testing"

    "github.com/hashicorp/terraform/helper/acctest"
    "github.com/hashicorp/terraform/helper/resource"
    "github.com/hashicorp/terraform/terraform"
    "github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/tf"
    "github.com/terraform-providers/terraform-provider-azurerm/azurerm/utils"
)


func testCheckAzureRMSensitivityLabelExists(resourceName string) resource.TestCheckFunc {
    return func(s *terraform.State) error {
        rs, ok := s.RootModule().Resources[resourceName]
        if !ok {
            return fmt.Errorf("Sensitivity Label not found: %s", resourceName)
        }

        name := rs.Primary.Attributes["name"]
        resourceGroup := rs.Primary.Attributes["resource_group"]
        databaseName := rs.Primary.Attributes["database_name"]
        schemaName := rs.Primary.Attributes["schema_name"]
        sensitivityLabelSource := rs.Primary.Attributes["sensitivity_label_source"]
        serverName := rs.Primary.Attributes["server_name"]
        tableName := rs.Primary.Attributes["table_name"]

        client := testAccProvider.Meta().(*ArmClient).sensitivityLabelsClient
        ctx := testAccProvider.Meta().(*ArmClient).StopContext

        if resp, err := client.Get(ctx, resourceGroup, serverName, databaseName, schemaName, tableName, name, sensitivityLabelSource); err != nil {
            if utils.ResponseWasNotFound(resp.Response) {
                return fmt.Errorf("Bad: Sensitivity Label %q (Sensitivity Label Source %q / Table Name %q / Schema Name %q / Database Name %q / Server Name %q / Resource Group %q) does not exist", name, sensitivityLabelSource, tableName, schemaName, databaseName, serverName, resourceGroup)
            }
            return fmt.Errorf("Bad: Get on sensitivityLabelsClient: %+v", err)
        }

        return nil
    }
}

func testCheckAzureRMSensitivityLabelDestroy(s *terraform.State) error {
    client := testAccProvider.Meta().(*ArmClient).sensitivityLabelsClient
    ctx := testAccProvider.Meta().(*ArmClient).StopContext

    for _, rs := range s.RootModule().Resources {
        if rs.Type != "azurerm_sensitivity_label" {
            continue
        }

        name := rs.Primary.Attributes["name"]
        resourceGroup := rs.Primary.Attributes["resource_group"]
        databaseName := rs.Primary.Attributes["database_name"]
        schemaName := rs.Primary.Attributes["schema_name"]
        sensitivityLabelSource := rs.Primary.Attributes["sensitivity_label_source"]
        serverName := rs.Primary.Attributes["server_name"]
        tableName := rs.Primary.Attributes["table_name"]

        if resp, err := client.Get(ctx, resourceGroup, serverName, databaseName, schemaName, tableName, name, sensitivityLabelSource); err != nil {
            if !utils.ResponseWasNotFound(resp.Response) {
                return fmt.Errorf("Bad: Get on sensitivityLabelsClient: %+v", err)
            }
        }

        return nil
    }

    return nil
}
