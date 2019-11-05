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


func testCheckAzureRMEnvironmentExists(resourceName string) resource.TestCheckFunc {
    return func(s *terraform.State) error {
        rs, ok := s.RootModule().Resources[resourceName]
        if !ok {
            return fmt.Errorf("Environment not found: %s", resourceName)
        }

        resourceGroup := rs.Primary.Attributes["resource_group"]
        environmentName := rs.Primary.Attributes["environment_name"]
        environmentSettingName := rs.Primary.Attributes["environment_setting_name"]
        labAccountName := rs.Primary.Attributes["lab_account_name"]
        labName := rs.Primary.Attributes["lab_name"]

        client := testAccProvider.Meta().(*ArmClient).environmentsClient
        ctx := testAccProvider.Meta().(*ArmClient).StopContext

        if resp, err := client.Get(ctx, resourceGroup, labAccountName, labName, environmentSettingName, environmentName); err != nil {
            if utils.ResponseWasNotFound(resp.Response) {
                return fmt.Errorf("Bad: Environment (Environment Name %q / Environment Setting Name %q / Lab Name %q / Lab Account Name %q / Resource Group %q) does not exist", environmentName, environmentSettingName, labName, labAccountName, resourceGroup)
            }
            return fmt.Errorf("Bad: Get on environmentsClient: %+v", err)
        }

        return nil
    }
}

func testCheckAzureRMEnvironmentDestroy(s *terraform.State) error {
    client := testAccProvider.Meta().(*ArmClient).environmentsClient
    ctx := testAccProvider.Meta().(*ArmClient).StopContext

    for _, rs := range s.RootModule().Resources {
        if rs.Type != "azurerm_environment" {
            continue
        }

        resourceGroup := rs.Primary.Attributes["resource_group"]
        environmentName := rs.Primary.Attributes["environment_name"]
        environmentSettingName := rs.Primary.Attributes["environment_setting_name"]
        labAccountName := rs.Primary.Attributes["lab_account_name"]
        labName := rs.Primary.Attributes["lab_name"]

        if resp, err := client.Get(ctx, resourceGroup, labAccountName, labName, environmentSettingName, environmentName); err != nil {
            if !utils.ResponseWasNotFound(resp.Response) {
                return fmt.Errorf("Bad: Get on environmentsClient: %+v", err)
            }
        }

        return nil
    }

    return nil
}
