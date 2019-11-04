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


func testCheckAzureRMBandwidthScheduleExists(resourceName string) resource.TestCheckFunc {
    return func(s *terraform.State) error {
        rs, ok := s.RootModule().Resources[resourceName]
        if !ok {
            return fmt.Errorf("Bandwidth Schedule not found: %s", resourceName)
        }

        name := rs.Primary.Attributes["name"]
        resourceGroup := rs.Primary.Attributes["resource_group"]
        deviceName := rs.Primary.Attributes["device_name"]

        client := testAccProvider.Meta().(*ArmClient).bandwidthSchedulesClient
        ctx := testAccProvider.Meta().(*ArmClient).StopContext

        if resp, err := client.Get(ctx, deviceName, name, resourceGroup); err != nil {
            if utils.ResponseWasNotFound(resp.Response) {
                return fmt.Errorf("Bad: Bandwidth Schedule %q (Resource Group %q / Device Name %q) does not exist", name, resourceGroup, deviceName)
            }
            return fmt.Errorf("Bad: Get on bandwidthSchedulesClient: %+v", err)
        }

        return nil
    }
}

func testCheckAzureRMBandwidthScheduleDestroy(s *terraform.State) error {
    client := testAccProvider.Meta().(*ArmClient).bandwidthSchedulesClient
    ctx := testAccProvider.Meta().(*ArmClient).StopContext

    for _, rs := range s.RootModule().Resources {
        if rs.Type != "azurerm_bandwidth_schedule" {
            continue
        }

        name := rs.Primary.Attributes["name"]
        resourceGroup := rs.Primary.Attributes["resource_group"]
        deviceName := rs.Primary.Attributes["device_name"]

        if resp, err := client.Get(ctx, deviceName, name, resourceGroup); err != nil {
            if !utils.ResponseWasNotFound(resp.Response) {
                return fmt.Errorf("Bad: Get on bandwidthSchedulesClient: %+v", err)
            }
        }

        return nil
    }

    return nil
}