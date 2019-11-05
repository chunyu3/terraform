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


func testCheckAzureRMServiceUnitExists(resourceName string) resource.TestCheckFunc {
    return func(s *terraform.State) error {
        rs, ok := s.RootModule().Resources[resourceName]
        if !ok {
            return fmt.Errorf("Service Unit not found: %s", resourceName)
        }

        resourceGroup := rs.Primary.Attributes["resource_group"]
        serviceName := rs.Primary.Attributes["service_name"]
        serviceTopologyName := rs.Primary.Attributes["service_topology_name"]
        serviceUnitName := rs.Primary.Attributes["service_unit_name"]

        client := testAccProvider.Meta().(*ArmClient).serviceUnitsClient
        ctx := testAccProvider.Meta().(*ArmClient).StopContext

        if resp, err := client.Get(ctx, resourceGroup, serviceTopologyName, serviceName, serviceUnitName); err != nil {
            if utils.ResponseWasNotFound(resp.Response) {
                return fmt.Errorf("Bad: Service Unit (Service Unit Name %q / Service Name %q / Service Topology Name %q / Resource Group %q) does not exist", serviceUnitName, serviceName, serviceTopologyName, resourceGroup)
            }
            return fmt.Errorf("Bad: Get on serviceUnitsClient: %+v", err)
        }

        return nil
    }
}

func testCheckAzureRMServiceUnitDestroy(s *terraform.State) error {
    client := testAccProvider.Meta().(*ArmClient).serviceUnitsClient
    ctx := testAccProvider.Meta().(*ArmClient).StopContext

    for _, rs := range s.RootModule().Resources {
        if rs.Type != "azurerm_service_unit" {
            continue
        }

        resourceGroup := rs.Primary.Attributes["resource_group"]
        serviceName := rs.Primary.Attributes["service_name"]
        serviceTopologyName := rs.Primary.Attributes["service_topology_name"]
        serviceUnitName := rs.Primary.Attributes["service_unit_name"]

        if resp, err := client.Get(ctx, resourceGroup, serviceTopologyName, serviceName, serviceUnitName); err != nil {
            if !utils.ResponseWasNotFound(resp.Response) {
                return fmt.Errorf("Bad: Get on serviceUnitsClient: %+v", err)
            }
        }

        return nil
    }

    return nil
}