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


func testCheckAzureRMBackupScheduleExists(resourceName string) resource.TestCheckFunc {
    return func(s *terraform.State) error {
        rs, ok := s.RootModule().Resources[resourceName]
        if !ok {
            return fmt.Errorf("Backup Schedule not found: %s", resourceName)
        }

        name := rs.Primary.Attributes["name"]
        resourceGroup := rs.Primary.Attributes["resource_group"]
        backupPolicyName := rs.Primary.Attributes["backup_policy_name"]
        backupScheduleName := rs.Primary.Attributes["backup_schedule_name"]
        deviceName := rs.Primary.Attributes["device_name"]

        client := testAccProvider.Meta().(*ArmClient).backupSchedulesClient
        ctx := testAccProvider.Meta().(*ArmClient).StopContext

        if resp, err := client.Get(ctx, resourceGroup, name, deviceName, backupPolicyName, backupScheduleName); err != nil {
            if utils.ResponseWasNotFound(resp.Response) {
                return fmt.Errorf("Bad: Backup Schedule %q (Resource Group %q / Backup Schedule Name %q / Backup Policy Name %q / Device Name %q) does not exist", name, resourceGroup, backupScheduleName, backupPolicyName, deviceName)
            }
            return fmt.Errorf("Bad: Get on backupSchedulesClient: %+v", err)
        }

        return nil
    }
}

func testCheckAzureRMBackupScheduleDestroy(s *terraform.State) error {
    client := testAccProvider.Meta().(*ArmClient).backupSchedulesClient
    ctx := testAccProvider.Meta().(*ArmClient).StopContext

    for _, rs := range s.RootModule().Resources {
        if rs.Type != "azurerm_backup_schedule" {
            continue
        }

        name := rs.Primary.Attributes["name"]
        resourceGroup := rs.Primary.Attributes["resource_group"]
        backupPolicyName := rs.Primary.Attributes["backup_policy_name"]
        backupScheduleName := rs.Primary.Attributes["backup_schedule_name"]
        deviceName := rs.Primary.Attributes["device_name"]

        if resp, err := client.Get(ctx, resourceGroup, name, deviceName, backupPolicyName, backupScheduleName); err != nil {
            if !utils.ResponseWasNotFound(resp.Response) {
                return fmt.Errorf("Bad: Get on backupSchedulesClient: %+v", err)
            }
        }

        return nil
    }

    return nil
}
