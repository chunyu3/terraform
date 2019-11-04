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


func testCheckAzureRMSynchronizationSettingExists(resourceName string) resource.TestCheckFunc {
    return func(s *terraform.State) error {
        rs, ok := s.RootModule().Resources[resourceName]
        if !ok {
            return fmt.Errorf("Synchronization Setting not found: %s", resourceName)
        }

        resourceGroup := rs.Primary.Attributes["resource_group"]
        accountName := rs.Primary.Attributes["account_name"]
        shareName := rs.Primary.Attributes["share_name"]
        synchronizationSettingName := rs.Primary.Attributes["synchronization_setting_name"]

        client := testAccProvider.Meta().(*ArmClient).synchronizationSettingsClient
        ctx := testAccProvider.Meta().(*ArmClient).StopContext

        if resp, err := client.Get(ctx, resourceGroup, accountName, shareName, synchronizationSettingName); err != nil {
            if utils.ResponseWasNotFound(resp.Response) {
                return fmt.Errorf("Bad: Synchronization Setting (Synchronization Setting Name %q / Share Name %q / Account Name %q / Resource Group %q) does not exist", synchronizationSettingName, shareName, accountName, resourceGroup)
            }
            return fmt.Errorf("Bad: Get on synchronizationSettingsClient: %+v", err)
        }

        return nil
    }
}

func testCheckAzureRMSynchronizationSettingDestroy(s *terraform.State) error {
    client := testAccProvider.Meta().(*ArmClient).synchronizationSettingsClient
    ctx := testAccProvider.Meta().(*ArmClient).StopContext

    for _, rs := range s.RootModule().Resources {
        if rs.Type != "azurerm_synchronization_setting" {
            continue
        }

        resourceGroup := rs.Primary.Attributes["resource_group"]
        accountName := rs.Primary.Attributes["account_name"]
        shareName := rs.Primary.Attributes["share_name"]
        synchronizationSettingName := rs.Primary.Attributes["synchronization_setting_name"]

        if resp, err := client.Get(ctx, resourceGroup, accountName, shareName, synchronizationSettingName); err != nil {
            if !utils.ResponseWasNotFound(resp.Response) {
                return fmt.Errorf("Bad: Get on synchronizationSettingsClient: %+v", err)
            }
        }

        return nil
    }

    return nil
}