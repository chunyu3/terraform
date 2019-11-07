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


func testCheckAzureRMP2sVpnServerConfigurationExists(resourceName string) resource.TestCheckFunc {
    return func(s *terraform.State) error {
        rs, ok := s.RootModule().Resources[resourceName]
        if !ok {
            return fmt.Errorf("P2s Vpn Server Configuration not found: %s", resourceName)
        }

        name := rs.Primary.Attributes["name"]
        resourceGroup := rs.Primary.Attributes["resource_group"]
        virtualWanName := rs.Primary.Attributes["virtual_wan_name"]

        client := testAccProvider.Meta().(*ArmClient).p2sVpnServerConfigurationsClient
        ctx := testAccProvider.Meta().(*ArmClient).StopContext

        if resp, err := client.Get(ctx, resourceGroup, virtualWanName, name); err != nil {
            if utils.ResponseWasNotFound(resp.Response) {
                return fmt.Errorf("Bad: P2s Vpn Server Configuration %q (Virtual Wan Name %q / Resource Group %q) does not exist", name, virtualWanName, resourceGroup)
            }
            return fmt.Errorf("Bad: Get on p2sVpnServerConfigurationsClient: %+v", err)
        }

        return nil
    }
}

func testCheckAzureRMP2sVpnServerConfigurationDestroy(s *terraform.State) error {
    client := testAccProvider.Meta().(*ArmClient).p2sVpnServerConfigurationsClient
    ctx := testAccProvider.Meta().(*ArmClient).StopContext

    for _, rs := range s.RootModule().Resources {
        if rs.Type != "azurerm_p2s_vpn_server_configuration" {
            continue
        }

        name := rs.Primary.Attributes["name"]
        resourceGroup := rs.Primary.Attributes["resource_group"]
        virtualWanName := rs.Primary.Attributes["virtual_wan_name"]

        if resp, err := client.Get(ctx, resourceGroup, virtualWanName, name); err != nil {
            if !utils.ResponseWasNotFound(resp.Response) {
                return fmt.Errorf("Bad: Get on p2sVpnServerConfigurationsClient: %+v", err)
            }
        }

        return nil
    }

    return nil
}