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


func testCheckAzureRMSavedSearcheExists(resourceName string) resource.TestCheckFunc {
    return func(s *terraform.State) error {
        rs, ok := s.RootModule().Resources[resourceName]
        if !ok {
            return fmt.Errorf("Saved Searche not found: %s", resourceName)
        }

        name := rs.Primary.Attributes["name"]
        resourceGroup := rs.Primary.Attributes["resource_group"]
        savedSearchID := rs.Primary.Attributes["saved_search_id"]

        client := testAccProvider.Meta().(*ArmClient).savedSearchesClient
        ctx := testAccProvider.Meta().(*ArmClient).StopContext

        if resp, err := client.Get(ctx, resourceGroup, name, savedSearchID); err != nil {
            if utils.ResponseWasNotFound(resp.Response) {
                return fmt.Errorf("Bad: Saved Searche %q (Saved Search %q / Resource Group %q) does not exist", name, savedSearchID, resourceGroup)
            }
            return fmt.Errorf("Bad: Get on savedSearchesClient: %+v", err)
        }

        return nil
    }
}

func testCheckAzureRMSavedSearcheDestroy(s *terraform.State) error {
    client := testAccProvider.Meta().(*ArmClient).savedSearchesClient
    ctx := testAccProvider.Meta().(*ArmClient).StopContext

    for _, rs := range s.RootModule().Resources {
        if rs.Type != "azurerm_saved_searche" {
            continue
        }

        name := rs.Primary.Attributes["name"]
        resourceGroup := rs.Primary.Attributes["resource_group"]
        savedSearchID := rs.Primary.Attributes["saved_search_id"]

        if resp, err := client.Get(ctx, resourceGroup, name, savedSearchID); err != nil {
            if !utils.ResponseWasNotFound(resp.Response) {
                return fmt.Errorf("Bad: Get on savedSearchesClient: %+v", err)
            }
        }

        return nil
    }

    return nil
}
