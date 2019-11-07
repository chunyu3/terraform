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


func testCheckAzureRMReplicationProtectionContainerMappingExists(resourceName string) resource.TestCheckFunc {
    return func(s *terraform.State) error {
        rs, ok := s.RootModule().Resources[resourceName]
        if !ok {
            return fmt.Errorf("Replication Protection Container Mapping not found: %s", resourceName)
        }

        name := rs.Primary.Attributes["name"]
        resourceGroup := rs.Primary.Attributes["resource_group"]
        fabricName := rs.Primary.Attributes["fabric_name"]
        protectionContainerName := rs.Primary.Attributes["protection_container_name"]
        resourceName := rs.Primary.Attributes["resource_name"]

        client := testAccProvider.Meta().(*ArmClient).replicationProtectionContainerMappingsClient
        ctx := testAccProvider.Meta().(*ArmClient).StopContext

        if resp, err := client.Get(ctx, resourceGroup, resourceName, fabricName, protectionContainerName, name); err != nil {
            if utils.ResponseWasNotFound(resp.Response) {
                return fmt.Errorf("Bad: Replication Protection Container Mapping %q (Protection Container Name %q / Fabric Name %q / Resource Group %q / Resource Name %q) does not exist", name, protectionContainerName, fabricName, resourceGroup, resourceName)
            }
            return fmt.Errorf("Bad: Get on replicationProtectionContainerMappingsClient: %+v", err)
        }

        return nil
    }
}

func testCheckAzureRMReplicationProtectionContainerMappingDestroy(s *terraform.State) error {
    client := testAccProvider.Meta().(*ArmClient).replicationProtectionContainerMappingsClient
    ctx := testAccProvider.Meta().(*ArmClient).StopContext

    for _, rs := range s.RootModule().Resources {
        if rs.Type != "azurerm_replication_protection_container_mapping" {
            continue
        }

        name := rs.Primary.Attributes["name"]
        resourceGroup := rs.Primary.Attributes["resource_group"]
        fabricName := rs.Primary.Attributes["fabric_name"]
        protectionContainerName := rs.Primary.Attributes["protection_container_name"]
        resourceName := rs.Primary.Attributes["resource_name"]

        if resp, err := client.Get(ctx, resourceGroup, resourceName, fabricName, protectionContainerName, name); err != nil {
            if !utils.ResponseWasNotFound(resp.Response) {
                return fmt.Errorf("Bad: Get on replicationProtectionContainerMappingsClient: %+v", err)
            }
        }

        return nil
    }

    return nil
}