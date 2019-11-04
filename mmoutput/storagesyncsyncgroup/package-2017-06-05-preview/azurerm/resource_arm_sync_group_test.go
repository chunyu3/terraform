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


func testCheckAzureRMSyncGroupExists(resourceName string) resource.TestCheckFunc {
    return func(s *terraform.State) error {
        rs, ok := s.RootModule().Resources[resourceName]
        if !ok {
            return fmt.Errorf("Sync Group not found: %s", resourceName)
        }

        resourceGroup := rs.Primary.Attributes["resource_group"]
        storageSyncServiceName := rs.Primary.Attributes["storage_sync_service_name"]
        syncGroupName := rs.Primary.Attributes["sync_group_name"]

        client := testAccProvider.Meta().(*ArmClient).syncGroupsClient
        ctx := testAccProvider.Meta().(*ArmClient).StopContext

        if resp, err := client.Get(ctx, resourceGroup, storageSyncServiceName, syncGroupName); err != nil {
            if utils.ResponseWasNotFound(resp.Response) {
                return fmt.Errorf("Bad: Sync Group (Sync Group Name %q / Storage Sync Service Name %q / Resource Group %q) does not exist", syncGroupName, storageSyncServiceName, resourceGroup)
            }
            return fmt.Errorf("Bad: Get on syncGroupsClient: %+v", err)
        }

        return nil
    }
}

func testCheckAzureRMSyncGroupDestroy(s *terraform.State) error {
    client := testAccProvider.Meta().(*ArmClient).syncGroupsClient
    ctx := testAccProvider.Meta().(*ArmClient).StopContext

    for _, rs := range s.RootModule().Resources {
        if rs.Type != "azurerm_sync_group" {
            continue
        }

        resourceGroup := rs.Primary.Attributes["resource_group"]
        storageSyncServiceName := rs.Primary.Attributes["storage_sync_service_name"]
        syncGroupName := rs.Primary.Attributes["sync_group_name"]

        if resp, err := client.Get(ctx, resourceGroup, storageSyncServiceName, syncGroupName); err != nil {
            if !utils.ResponseWasNotFound(resp.Response) {
                return fmt.Errorf("Bad: Get on syncGroupsClient: %+v", err)
            }
        }

        return nil
    }

    return nil
}
