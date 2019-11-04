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


func testCheckAzureRMReplicationFabricExists(resourceName string) resource.TestCheckFunc {
    return func(s *terraform.State) error {
        rs, ok := s.RootModule().Resources[resourceName]
        if !ok {
            return fmt.Errorf("Replication Fabric not found: %s", resourceName)
        }

        resourceGroup := rs.Primary.Attributes["resource_group"]
        fabricName := rs.Primary.Attributes["fabric_name"]
        resourceName := rs.Primary.Attributes["resource_name"]

        client := testAccProvider.Meta().(*ArmClient).replicationFabricsClient
        ctx := testAccProvider.Meta().(*ArmClient).StopContext

        if resp, err := client.Get(ctx, resourceName, resourceGroup, fabricName); err != nil {
            if utils.ResponseWasNotFound(resp.Response) {
                return fmt.Errorf("Bad: Replication Fabric (Fabric Name %q / Resource Group %q / Resource Name %q) does not exist", fabricName, resourceGroup, resourceName)
            }
            return fmt.Errorf("Bad: Get on replicationFabricsClient: %+v", err)
        }

        return nil
    }
}

func testCheckAzureRMReplicationFabricDestroy(s *terraform.State) error {
    client := testAccProvider.Meta().(*ArmClient).replicationFabricsClient
    ctx := testAccProvider.Meta().(*ArmClient).StopContext

    for _, rs := range s.RootModule().Resources {
        if rs.Type != "azurerm_replication_fabric" {
            continue
        }

        resourceGroup := rs.Primary.Attributes["resource_group"]
        fabricName := rs.Primary.Attributes["fabric_name"]
        resourceName := rs.Primary.Attributes["resource_name"]

        if resp, err := client.Get(ctx, resourceName, resourceGroup, fabricName); err != nil {
            if !utils.ResponseWasNotFound(resp.Response) {
                return fmt.Errorf("Bad: Get on replicationFabricsClient: %+v", err)
            }
        }

        return nil
    }

    return nil
}
