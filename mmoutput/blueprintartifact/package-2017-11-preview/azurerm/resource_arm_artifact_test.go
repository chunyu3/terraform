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


func testCheckAzureRMArtifactExists(resourceName string) resource.TestCheckFunc {
    return func(s *terraform.State) error {
        rs, ok := s.RootModule().Resources[resourceName]
        if !ok {
            return fmt.Errorf("Artifact not found: %s", resourceName)
        }

        name := rs.Primary.Attributes["name"]
        blueprintName := rs.Primary.Attributes["blueprint_name"]
        managementGroupName := rs.Primary.Attributes["management_group_name"]

        client := testAccProvider.Meta().(*ArmClient).artifactsClient
        ctx := testAccProvider.Meta().(*ArmClient).StopContext

        if resp, err := client.Get(ctx, managementGroupName, blueprintName, name); err != nil {
            if utils.ResponseWasNotFound(resp.Response) {
                return fmt.Errorf("Bad: Artifact %q (Blueprint Name %q / Management Group Name %q) does not exist", name, blueprintName, managementGroupName)
            }
            return fmt.Errorf("Bad: Get on artifactsClient: %+v", err)
        }

        return nil
    }
}

func testCheckAzureRMArtifactDestroy(s *terraform.State) error {
    client := testAccProvider.Meta().(*ArmClient).artifactsClient
    ctx := testAccProvider.Meta().(*ArmClient).StopContext

    for _, rs := range s.RootModule().Resources {
        if rs.Type != "azurerm_artifact" {
            continue
        }

        name := rs.Primary.Attributes["name"]
        blueprintName := rs.Primary.Attributes["blueprint_name"]
        managementGroupName := rs.Primary.Attributes["management_group_name"]

        if resp, err := client.Get(ctx, managementGroupName, blueprintName, name); err != nil {
            if !utils.ResponseWasNotFound(resp.Response) {
                return fmt.Errorf("Bad: Get on artifactsClient: %+v", err)
            }
        }

        return nil
    }

    return nil
}
