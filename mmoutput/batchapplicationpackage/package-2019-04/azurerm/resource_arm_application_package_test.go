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


func testCheckAzureRMApplicationPackageExists(resourceName string) resource.TestCheckFunc {
    return func(s *terraform.State) error {
        rs, ok := s.RootModule().Resources[resourceName]
        if !ok {
            return fmt.Errorf("Application Package not found: %s", resourceName)
        }

        resourceGroup := rs.Primary.Attributes["resource_group"]
        accountName := rs.Primary.Attributes["account_name"]
        applicationName := rs.Primary.Attributes["application_name"]
        versionName := rs.Primary.Attributes["version_name"]

        client := testAccProvider.Meta().(*ArmClient).applicationPackageClient
        ctx := testAccProvider.Meta().(*ArmClient).StopContext

        if resp, err := client.Get(ctx, resourceGroup, accountName, applicationName, versionName); err != nil {
            if utils.ResponseWasNotFound(resp.Response) {
                return fmt.Errorf("Bad: Application Package (Version Name %q / Application Name %q / Account Name %q / Resource Group %q) does not exist", versionName, applicationName, accountName, resourceGroup)
            }
            return fmt.Errorf("Bad: Get on applicationPackageClient: %+v", err)
        }

        return nil
    }
}

func testCheckAzureRMApplicationPackageDestroy(s *terraform.State) error {
    client := testAccProvider.Meta().(*ArmClient).applicationPackageClient
    ctx := testAccProvider.Meta().(*ArmClient).StopContext

    for _, rs := range s.RootModule().Resources {
        if rs.Type != "azurerm_application_package" {
            continue
        }

        resourceGroup := rs.Primary.Attributes["resource_group"]
        accountName := rs.Primary.Attributes["account_name"]
        applicationName := rs.Primary.Attributes["application_name"]
        versionName := rs.Primary.Attributes["version_name"]

        if resp, err := client.Get(ctx, resourceGroup, accountName, applicationName, versionName); err != nil {
            if !utils.ResponseWasNotFound(resp.Response) {
                return fmt.Errorf("Bad: Get on applicationPackageClient: %+v", err)
            }
        }

        return nil
    }

    return nil
}
