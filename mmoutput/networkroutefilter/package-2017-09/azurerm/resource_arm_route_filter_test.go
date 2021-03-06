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


func testCheckAzureRMRouteFilterExists(resourceName string) resource.TestCheckFunc {
    return func(s *terraform.State) error {
        rs, ok := s.RootModule().Resources[resourceName]
        if !ok {
            return fmt.Errorf("Route Filter not found: %s", resourceName)
        }

        resourceGroup := rs.Primary.Attributes["resource_group"]
        routeFilterName := rs.Primary.Attributes["route_filter_name"]

        client := testAccProvider.Meta().(*ArmClient).routeFiltersClient
        ctx := testAccProvider.Meta().(*ArmClient).StopContext

        if resp, err := client.Get(ctx, resourceGroup, routeFilterName); err != nil {
            if utils.ResponseWasNotFound(resp.Response) {
                return fmt.Errorf("Bad: Route Filter (Route Filter Name %q / Resource Group %q) does not exist", routeFilterName, resourceGroup)
            }
            return fmt.Errorf("Bad: Get on routeFiltersClient: %+v", err)
        }

        return nil
    }
}

func testCheckAzureRMRouteFilterDestroy(s *terraform.State) error {
    client := testAccProvider.Meta().(*ArmClient).routeFiltersClient
    ctx := testAccProvider.Meta().(*ArmClient).StopContext

    for _, rs := range s.RootModule().Resources {
        if rs.Type != "azurerm_route_filter" {
            continue
        }

        resourceGroup := rs.Primary.Attributes["resource_group"]
        routeFilterName := rs.Primary.Attributes["route_filter_name"]

        if resp, err := client.Get(ctx, resourceGroup, routeFilterName); err != nil {
            if !utils.ResponseWasNotFound(resp.Response) {
                return fmt.Errorf("Bad: Get on routeFiltersClient: %+v", err)
            }
        }

        return nil
    }

    return nil
}
