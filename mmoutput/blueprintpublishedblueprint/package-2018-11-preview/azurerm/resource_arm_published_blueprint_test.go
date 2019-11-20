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


func testCheckAzureRMPublishedBlueprintExists(resourceName string) resource.TestCheckFunc {
    return func(s *terraform.State) error {
        rs, ok := s.RootModule().Resources[resourceName]
        if !ok {
            return fmt.Errorf("Published Blueprint not found: %s", resourceName)
        }

        name := rs.Primary.Attributes["name"]
        scope := rs.Primary.Attributes["scope"]
        versionID := rs.Primary.Attributes["version_id"]

        client := testAccProvider.Meta().(*ArmClient).publishedBlueprintsClient
        ctx := testAccProvider.Meta().(*ArmClient).StopContext

        if resp, err := client.Get(ctx, scope, name, versionID); err != nil {
            if utils.ResponseWasNotFound(resp.Response) {
                return fmt.Errorf("Bad: Published Blueprint %q (Version %q / Scope %q) does not exist", name, versionID, scope)
            }
            return fmt.Errorf("Bad: Get on publishedBlueprintsClient: %+v", err)
        }

        return nil
    }
}

func testCheckAzureRMPublishedBlueprintDestroy(s *terraform.State) error {
    client := testAccProvider.Meta().(*ArmClient).publishedBlueprintsClient
    ctx := testAccProvider.Meta().(*ArmClient).StopContext

    for _, rs := range s.RootModule().Resources {
        if rs.Type != "azurerm_published_blueprint" {
            continue
        }

        name := rs.Primary.Attributes["name"]
        scope := rs.Primary.Attributes["scope"]
        versionID := rs.Primary.Attributes["version_id"]

        if resp, err := client.Get(ctx, scope, name, versionID); err != nil {
            if !utils.ResponseWasNotFound(resp.Response) {
                return fmt.Errorf("Bad: Get on publishedBlueprintsClient: %+v", err)
            }
        }

        return nil
    }

    return nil
}
