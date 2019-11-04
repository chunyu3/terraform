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


func testCheckAzureRMLinkExists(resourceName string) resource.TestCheckFunc {
    return func(s *terraform.State) error {
        rs, ok := s.RootModule().Resources[resourceName]
        if !ok {
            return fmt.Errorf("Link not found: %s", resourceName)
        }

        resourceGroup := rs.Primary.Attributes["resource_group"]
        hubName := rs.Primary.Attributes["hub_name"]
        linkName := rs.Primary.Attributes["link_name"]

        client := testAccProvider.Meta().(*ArmClient).linksClient
        ctx := testAccProvider.Meta().(*ArmClient).StopContext

        if resp, err := client.Get(ctx, resourceGroup, hubName, linkName); err != nil {
            if utils.ResponseWasNotFound(resp.Response) {
                return fmt.Errorf("Bad: Link (Link Name %q / Hub Name %q / Resource Group %q) does not exist", linkName, hubName, resourceGroup)
            }
            return fmt.Errorf("Bad: Get on linksClient: %+v", err)
        }

        return nil
    }
}

func testCheckAzureRMLinkDestroy(s *terraform.State) error {
    client := testAccProvider.Meta().(*ArmClient).linksClient
    ctx := testAccProvider.Meta().(*ArmClient).StopContext

    for _, rs := range s.RootModule().Resources {
        if rs.Type != "azurerm_link" {
            continue
        }

        resourceGroup := rs.Primary.Attributes["resource_group"]
        hubName := rs.Primary.Attributes["hub_name"]
        linkName := rs.Primary.Attributes["link_name"]

        if resp, err := client.Get(ctx, resourceGroup, hubName, linkName); err != nil {
            if !utils.ResponseWasNotFound(resp.Response) {
                return fmt.Errorf("Bad: Get on linksClient: %+v", err)
            }
        }

        return nil
    }

    return nil
}