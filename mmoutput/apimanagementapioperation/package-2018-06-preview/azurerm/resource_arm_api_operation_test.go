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


func testCheckAzureRMApiOperationExists(resourceName string) resource.TestCheckFunc {
    return func(s *terraform.State) error {
        rs, ok := s.RootModule().Resources[resourceName]
        if !ok {
            return fmt.Errorf("Api Operation not found: %s", resourceName)
        }

        name := rs.Primary.Attributes["name"]
        resourceGroup := rs.Primary.Attributes["resource_group"]
        apiID := rs.Primary.Attributes["api_id"]
        operationID := rs.Primary.Attributes["operation_id"]

        client := testAccProvider.Meta().(*ArmClient).apiOperationClient
        ctx := testAccProvider.Meta().(*ArmClient).StopContext

        if resp, err := client.Get(ctx, resourceGroup, name, apiID, operationID); err != nil {
            if utils.ResponseWasNotFound(resp.Response) {
                return fmt.Errorf("Bad: Api Operation %q (Operation %q / Api %q / Resource Group %q) does not exist", name, operationID, apiID, resourceGroup)
            }
            return fmt.Errorf("Bad: Get on apiOperationClient: %+v", err)
        }

        return nil
    }
}

func testCheckAzureRMApiOperationDestroy(s *terraform.State) error {
    client := testAccProvider.Meta().(*ArmClient).apiOperationClient
    ctx := testAccProvider.Meta().(*ArmClient).StopContext

    for _, rs := range s.RootModule().Resources {
        if rs.Type != "azurerm_api_operation" {
            continue
        }

        name := rs.Primary.Attributes["name"]
        resourceGroup := rs.Primary.Attributes["resource_group"]
        apiID := rs.Primary.Attributes["api_id"]
        operationID := rs.Primary.Attributes["operation_id"]

        if resp, err := client.Get(ctx, resourceGroup, name, apiID, operationID); err != nil {
            if !utils.ResponseWasNotFound(resp.Response) {
                return fmt.Errorf("Bad: Get on apiOperationClient: %+v", err)
            }
        }

        return nil
    }

    return nil
}
