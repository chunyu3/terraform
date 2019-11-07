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


func testCheckAzureRMWorkflowAccessKeyExists(resourceName string) resource.TestCheckFunc {
    return func(s *terraform.State) error {
        rs, ok := s.RootModule().Resources[resourceName]
        if !ok {
            return fmt.Errorf("Workflow Access Key not found: %s", resourceName)
        }

        name := rs.Primary.Attributes["name"]
        resourceGroup := rs.Primary.Attributes["resource_group"]
        workflowName := rs.Primary.Attributes["workflow_name"]

        client := testAccProvider.Meta().(*ArmClient).workflowAccessKeysClient
        ctx := testAccProvider.Meta().(*ArmClient).StopContext

        if resp, err := client.Get(ctx, resourceGroup, workflowName, name); err != nil {
            if utils.ResponseWasNotFound(resp.Response) {
                return fmt.Errorf("Bad: Workflow Access Key %q (Workflow Name %q / Resource Group %q) does not exist", name, workflowName, resourceGroup)
            }
            return fmt.Errorf("Bad: Get on workflowAccessKeysClient: %+v", err)
        }

        return nil
    }
}

func testCheckAzureRMWorkflowAccessKeyDestroy(s *terraform.State) error {
    client := testAccProvider.Meta().(*ArmClient).workflowAccessKeysClient
    ctx := testAccProvider.Meta().(*ArmClient).StopContext

    for _, rs := range s.RootModule().Resources {
        if rs.Type != "azurerm_workflow_access_key" {
            continue
        }

        name := rs.Primary.Attributes["name"]
        resourceGroup := rs.Primary.Attributes["resource_group"]
        workflowName := rs.Primary.Attributes["workflow_name"]

        if resp, err := client.Get(ctx, resourceGroup, workflowName, name); err != nil {
            if !utils.ResponseWasNotFound(resp.Response) {
                return fmt.Errorf("Bad: Get on workflowAccessKeysClient: %+v", err)
            }
        }

        return nil
    }

    return nil
}
