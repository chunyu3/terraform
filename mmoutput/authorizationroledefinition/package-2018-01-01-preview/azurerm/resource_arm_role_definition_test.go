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


func testCheckAzureRMRoleDefinitionExists(resourceName string) resource.TestCheckFunc {
    return func(s *terraform.State) error {
        rs, ok := s.RootModule().Resources[resourceName]
        if !ok {
            return fmt.Errorf("Role Definition not found: %s", resourceName)
        }

        roleDefinitionID := rs.Primary.Attributes["role_definition_id"]
        scope := rs.Primary.Attributes["scope"]

        client := testAccProvider.Meta().(*ArmClient).roleDefinitionsClient
        ctx := testAccProvider.Meta().(*ArmClient).StopContext

        if resp, err := client.Get(ctx, scope, roleDefinitionID); err != nil {
            if utils.ResponseWasNotFound(resp.Response) {
                return fmt.Errorf("Bad: Role Definition (Role Definition %q / Scope %q) does not exist", roleDefinitionID, scope)
            }
            return fmt.Errorf("Bad: Get on roleDefinitionsClient: %+v", err)
        }

        return nil
    }
}

func testCheckAzureRMRoleDefinitionDestroy(s *terraform.State) error {
    client := testAccProvider.Meta().(*ArmClient).roleDefinitionsClient
    ctx := testAccProvider.Meta().(*ArmClient).StopContext

    for _, rs := range s.RootModule().Resources {
        if rs.Type != "azurerm_role_definition" {
            continue
        }

        roleDefinitionID := rs.Primary.Attributes["role_definition_id"]
        scope := rs.Primary.Attributes["scope"]

        if resp, err := client.Get(ctx, scope, roleDefinitionID); err != nil {
            if !utils.ResponseWasNotFound(resp.Response) {
                return fmt.Errorf("Bad: Get on roleDefinitionsClient: %+v", err)
            }
        }

        return nil
    }

    return nil
}
