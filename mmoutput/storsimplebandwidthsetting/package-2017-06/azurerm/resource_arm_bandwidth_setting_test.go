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


func testCheckAzureRMBandwidthSettingExists(resourceName string) resource.TestCheckFunc {
    return func(s *terraform.State) error {
        rs, ok := s.RootModule().Resources[resourceName]
        if !ok {
            return fmt.Errorf("Bandwidth Setting not found: %s", resourceName)
        }

        name := rs.Primary.Attributes["name"]
        resourceGroup := rs.Primary.Attributes["resource_group"]
        bandwidthSettingName := rs.Primary.Attributes["bandwidth_setting_name"]

        client := testAccProvider.Meta().(*ArmClient).bandwidthSettingsClient
        ctx := testAccProvider.Meta().(*ArmClient).StopContext

        if resp, err := client.Get(ctx, resourceGroup, name, bandwidthSettingName); err != nil {
            if utils.ResponseWasNotFound(resp.Response) {
                return fmt.Errorf("Bad: Bandwidth Setting %q (Resource Group %q / Bandwidth Setting Name %q) does not exist", name, resourceGroup, bandwidthSettingName)
            }
            return fmt.Errorf("Bad: Get on bandwidthSettingsClient: %+v", err)
        }

        return nil
    }
}

func testCheckAzureRMBandwidthSettingDestroy(s *terraform.State) error {
    client := testAccProvider.Meta().(*ArmClient).bandwidthSettingsClient
    ctx := testAccProvider.Meta().(*ArmClient).StopContext

    for _, rs := range s.RootModule().Resources {
        if rs.Type != "azurerm_bandwidth_setting" {
            continue
        }

        name := rs.Primary.Attributes["name"]
        resourceGroup := rs.Primary.Attributes["resource_group"]
        bandwidthSettingName := rs.Primary.Attributes["bandwidth_setting_name"]

        if resp, err := client.Get(ctx, resourceGroup, name, bandwidthSettingName); err != nil {
            if !utils.ResponseWasNotFound(resp.Response) {
                return fmt.Errorf("Bad: Get on bandwidthSettingsClient: %+v", err)
            }
        }

        return nil
    }

    return nil
}
