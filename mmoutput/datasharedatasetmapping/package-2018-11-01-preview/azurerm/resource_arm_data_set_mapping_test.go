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


func testCheckAzureRMDataSetMappingExists(resourceName string) resource.TestCheckFunc {
    return func(s *terraform.State) error {
        rs, ok := s.RootModule().Resources[resourceName]
        if !ok {
            return fmt.Errorf("Data Set Mapping not found: %s", resourceName)
        }

        name := rs.Primary.Attributes["name"]
        resourceGroup := rs.Primary.Attributes["resource_group"]
        accountName := rs.Primary.Attributes["account_name"]
        shareSubscriptionName := rs.Primary.Attributes["share_subscription_name"]

        client := testAccProvider.Meta().(*ArmClient).dataSetMappingsClient
        ctx := testAccProvider.Meta().(*ArmClient).StopContext

        if resp, err := client.Get(ctx, resourceGroup, accountName, shareSubscriptionName, name); err != nil {
            if utils.ResponseWasNotFound(resp.Response) {
                return fmt.Errorf("Bad: Data Set Mapping %q (Share Subscription Name %q / Account Name %q / Resource Group %q) does not exist", name, shareSubscriptionName, accountName, resourceGroup)
            }
            return fmt.Errorf("Bad: Get on dataSetMappingsClient: %+v", err)
        }

        return nil
    }
}

func testCheckAzureRMDataSetMappingDestroy(s *terraform.State) error {
    client := testAccProvider.Meta().(*ArmClient).dataSetMappingsClient
    ctx := testAccProvider.Meta().(*ArmClient).StopContext

    for _, rs := range s.RootModule().Resources {
        if rs.Type != "azurerm_data_set_mapping" {
            continue
        }

        name := rs.Primary.Attributes["name"]
        resourceGroup := rs.Primary.Attributes["resource_group"]
        accountName := rs.Primary.Attributes["account_name"]
        shareSubscriptionName := rs.Primary.Attributes["share_subscription_name"]

        if resp, err := client.Get(ctx, resourceGroup, accountName, shareSubscriptionName, name); err != nil {
            if !utils.ResponseWasNotFound(resp.Response) {
                return fmt.Errorf("Bad: Get on dataSetMappingsClient: %+v", err)
            }
        }

        return nil
    }

    return nil
}
