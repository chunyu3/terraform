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


func testCheckAzureRMSyncMemberExists(resourceName string) resource.TestCheckFunc {
    return func(s *terraform.State) error {
        rs, ok := s.RootModule().Resources[resourceName]
        if !ok {
            return fmt.Errorf("Sync Member not found: %s", resourceName)
        }

        name := rs.Primary.Attributes["name"]
        resourceGroup := rs.Primary.Attributes["resource_group"]
        databaseName := rs.Primary.Attributes["database_name"]
        serverName := rs.Primary.Attributes["server_name"]
        syncGroupName := rs.Primary.Attributes["sync_group_name"]

        client := testAccProvider.Meta().(*ArmClient).syncMembersClient
        ctx := testAccProvider.Meta().(*ArmClient).StopContext

        if resp, err := client.Get(ctx, resourceGroup, serverName, databaseName, syncGroupName, name); err != nil {
            if utils.ResponseWasNotFound(resp.Response) {
                return fmt.Errorf("Bad: Sync Member %q (Sync Group Name %q / Database Name %q / Server Name %q / Resource Group %q) does not exist", name, syncGroupName, databaseName, serverName, resourceGroup)
            }
            return fmt.Errorf("Bad: Get on syncMembersClient: %+v", err)
        }

        return nil
    }
}

func testCheckAzureRMSyncMemberDestroy(s *terraform.State) error {
    client := testAccProvider.Meta().(*ArmClient).syncMembersClient
    ctx := testAccProvider.Meta().(*ArmClient).StopContext

    for _, rs := range s.RootModule().Resources {
        if rs.Type != "azurerm_sync_member" {
            continue
        }

        name := rs.Primary.Attributes["name"]
        resourceGroup := rs.Primary.Attributes["resource_group"]
        databaseName := rs.Primary.Attributes["database_name"]
        serverName := rs.Primary.Attributes["server_name"]
        syncGroupName := rs.Primary.Attributes["sync_group_name"]

        if resp, err := client.Get(ctx, resourceGroup, serverName, databaseName, syncGroupName, name); err != nil {
            if !utils.ResponseWasNotFound(resp.Response) {
                return fmt.Errorf("Bad: Get on syncMembersClient: %+v", err)
            }
        }

        return nil
    }

    return nil
}