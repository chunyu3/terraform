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


func testCheckAzureRMExpressRouteCircuitConnectionExists(resourceName string) resource.TestCheckFunc {
    return func(s *terraform.State) error {
        rs, ok := s.RootModule().Resources[resourceName]
        if !ok {
            return fmt.Errorf("Express Route Circuit Connection not found: %s", resourceName)
        }

        resourceGroup := rs.Primary.Attributes["resource_group"]
        circuitName := rs.Primary.Attributes["circuit_name"]
        connectionName := rs.Primary.Attributes["connection_name"]
        peeringName := rs.Primary.Attributes["peering_name"]

        client := testAccProvider.Meta().(*ArmClient).expressRouteCircuitConnectionsClient
        ctx := testAccProvider.Meta().(*ArmClient).StopContext

        if resp, err := client.Get(ctx, resourceGroup, circuitName, peeringName, connectionName); err != nil {
            if utils.ResponseWasNotFound(resp.Response) {
                return fmt.Errorf("Bad: Express Route Circuit Connection (Connection Name %q / Peering Name %q / Circuit Name %q / Resource Group %q) does not exist", connectionName, peeringName, circuitName, resourceGroup)
            }
            return fmt.Errorf("Bad: Get on expressRouteCircuitConnectionsClient: %+v", err)
        }

        return nil
    }
}

func testCheckAzureRMExpressRouteCircuitConnectionDestroy(s *terraform.State) error {
    client := testAccProvider.Meta().(*ArmClient).expressRouteCircuitConnectionsClient
    ctx := testAccProvider.Meta().(*ArmClient).StopContext

    for _, rs := range s.RootModule().Resources {
        if rs.Type != "azurerm_express_route_circuit_connection" {
            continue
        }

        resourceGroup := rs.Primary.Attributes["resource_group"]
        circuitName := rs.Primary.Attributes["circuit_name"]
        connectionName := rs.Primary.Attributes["connection_name"]
        peeringName := rs.Primary.Attributes["peering_name"]

        if resp, err := client.Get(ctx, resourceGroup, circuitName, peeringName, connectionName); err != nil {
            if !utils.ResponseWasNotFound(resp.Response) {
                return fmt.Errorf("Bad: Get on expressRouteCircuitConnectionsClient: %+v", err)
            }
        }

        return nil
    }

    return nil
}