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


func testCheckAzureRMExpressRouteConnectionExists(resourceName string) resource.TestCheckFunc {
    return func(s *terraform.State) error {
        rs, ok := s.RootModule().Resources[resourceName]
        if !ok {
            return fmt.Errorf("Express Route Connection not found: %s", resourceName)
        }

        resourceGroup := rs.Primary.Attributes["resource_group"]
        connectionName := rs.Primary.Attributes["connection_name"]
        expressRouteGatewayName := rs.Primary.Attributes["express_route_gateway_name"]

        client := testAccProvider.Meta().(*ArmClient).expressRouteConnectionsClient
        ctx := testAccProvider.Meta().(*ArmClient).StopContext

        if resp, err := client.Get(ctx, resourceGroup, expressRouteGatewayName, connectionName); err != nil {
            if utils.ResponseWasNotFound(resp.Response) {
                return fmt.Errorf("Bad: Express Route Connection (Connection Name %q / Express Route Gateway Name %q / Resource Group %q) does not exist", connectionName, expressRouteGatewayName, resourceGroup)
            }
            return fmt.Errorf("Bad: Get on expressRouteConnectionsClient: %+v", err)
        }

        return nil
    }
}

func testCheckAzureRMExpressRouteConnectionDestroy(s *terraform.State) error {
    client := testAccProvider.Meta().(*ArmClient).expressRouteConnectionsClient
    ctx := testAccProvider.Meta().(*ArmClient).StopContext

    for _, rs := range s.RootModule().Resources {
        if rs.Type != "azurerm_express_route_connection" {
            continue
        }

        resourceGroup := rs.Primary.Attributes["resource_group"]
        connectionName := rs.Primary.Attributes["connection_name"]
        expressRouteGatewayName := rs.Primary.Attributes["express_route_gateway_name"]

        if resp, err := client.Get(ctx, resourceGroup, expressRouteGatewayName, connectionName); err != nil {
            if !utils.ResponseWasNotFound(resp.Response) {
                return fmt.Errorf("Bad: Get on expressRouteConnectionsClient: %+v", err)
            }
        }

        return nil
    }

    return nil
}