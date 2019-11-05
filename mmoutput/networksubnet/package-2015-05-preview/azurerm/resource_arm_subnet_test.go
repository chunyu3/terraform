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


func testCheckAzureRMSubnetExists(resourceName string) resource.TestCheckFunc {
    return func(s *terraform.State) error {
        rs, ok := s.RootModule().Resources[resourceName]
        if !ok {
            return fmt.Errorf("Subnet not found: %s", resourceName)
        }

        resourceGroup := rs.Primary.Attributes["resource_group"]
        subnetName := rs.Primary.Attributes["subnet_name"]
        virtualNetworkName := rs.Primary.Attributes["virtual_network_name"]

        client := testAccProvider.Meta().(*ArmClient).subnetsClient
        ctx := testAccProvider.Meta().(*ArmClient).StopContext

        if resp, err := client.Get(ctx, resourceGroup, virtualNetworkName, subnetName); err != nil {
            if utils.ResponseWasNotFound(resp.Response) {
                return fmt.Errorf("Bad: Subnet (Subnet Name %q / Virtual Network Name %q / Resource Group %q) does not exist", subnetName, virtualNetworkName, resourceGroup)
            }
            return fmt.Errorf("Bad: Get on subnetsClient: %+v", err)
        }

        return nil
    }
}

func testCheckAzureRMSubnetDestroy(s *terraform.State) error {
    client := testAccProvider.Meta().(*ArmClient).subnetsClient
    ctx := testAccProvider.Meta().(*ArmClient).StopContext

    for _, rs := range s.RootModule().Resources {
        if rs.Type != "azurerm_subnet" {
            continue
        }

        resourceGroup := rs.Primary.Attributes["resource_group"]
        subnetName := rs.Primary.Attributes["subnet_name"]
        virtualNetworkName := rs.Primary.Attributes["virtual_network_name"]

        if resp, err := client.Get(ctx, resourceGroup, virtualNetworkName, subnetName); err != nil {
            if !utils.ResponseWasNotFound(resp.Response) {
                return fmt.Errorf("Bad: Get on subnetsClient: %+v", err)
            }
        }

        return nil
    }

    return nil
}
