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


func testCheckAzureRMDiagnosticSettingExists(resourceName string) resource.TestCheckFunc {
    return func(s *terraform.State) error {
        rs, ok := s.RootModule().Resources[resourceName]
        if !ok {
            return fmt.Errorf("Diagnostic Setting not found: %s", resourceName)
        }

        name := rs.Primary.Attributes["name"]
        resourceURI := rs.Primary.Attributes["resource_uri"]

        client := testAccProvider.Meta().(*ArmClient).diagnosticSettingsClient
        ctx := testAccProvider.Meta().(*ArmClient).StopContext

        if resp, err := client.Get(ctx, resourceURI, name); err != nil {
            if utils.ResponseWasNotFound(resp.Response) {
                return fmt.Errorf("Bad: Diagnostic Setting %q (Resource Uri %q) does not exist", name, resourceURI)
            }
            return fmt.Errorf("Bad: Get on diagnosticSettingsClient: %+v", err)
        }

        return nil
    }
}

func testCheckAzureRMDiagnosticSettingDestroy(s *terraform.State) error {
    client := testAccProvider.Meta().(*ArmClient).diagnosticSettingsClient
    ctx := testAccProvider.Meta().(*ArmClient).StopContext

    for _, rs := range s.RootModule().Resources {
        if rs.Type != "azurerm_diagnostic_setting" {
            continue
        }

        name := rs.Primary.Attributes["name"]
        resourceURI := rs.Primary.Attributes["resource_uri"]

        if resp, err := client.Get(ctx, resourceURI, name); err != nil {
            if !utils.ResponseWasNotFound(resp.Response) {
                return fmt.Errorf("Bad: Get on diagnosticSettingsClient: %+v", err)
            }
        }

        return nil
    }

    return nil
}