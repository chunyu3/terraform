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


func testCheckAzureRMCloudConnectorExists(resourceName string) resource.TestCheckFunc {
    return func(s *terraform.State) error {
        rs, ok := s.RootModule().Resources[resourceName]
        if !ok {
            return fmt.Errorf("Cloud Connector not found: %s", resourceName)
        }

        connectorName := rs.Primary.Attributes["connector_name"]

        client := testAccProvider.Meta().(*ArmClient).cloudConnectorClient
        ctx := testAccProvider.Meta().(*ArmClient).StopContext

        if resp, err := client.Get(ctx, connectorName); err != nil {
            if utils.ResponseWasNotFound(resp.Response) {
                return fmt.Errorf("Bad: Cloud Connector (Connector Name %q) does not exist", connectorName)
            }
            return fmt.Errorf("Bad: Get on cloudConnectorClient: %+v", err)
        }

        return nil
    }
}

func testCheckAzureRMCloudConnectorDestroy(s *terraform.State) error {
    client := testAccProvider.Meta().(*ArmClient).cloudConnectorClient
    ctx := testAccProvider.Meta().(*ArmClient).StopContext

    for _, rs := range s.RootModule().Resources {
        if rs.Type != "azurerm_cloud_connector" {
            continue
        }

        connectorName := rs.Primary.Attributes["connector_name"]

        if resp, err := client.Get(ctx, connectorName); err != nil {
            if !utils.ResponseWasNotFound(resp.Response) {
                return fmt.Errorf("Bad: Get on cloudConnectorClient: %+v", err)
            }
        }

        return nil
    }

    return nil
}
