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


func testCheckAzureRMProtectionPolicyExists(resourceName string) resource.TestCheckFunc {
    return func(s *terraform.State) error {
        rs, ok := s.RootModule().Resources[resourceName]
        if !ok {
            return fmt.Errorf("Protection Policy not found: %s", resourceName)
        }

        resourceGroup := rs.Primary.Attributes["resource_group"]
        policyName := rs.Primary.Attributes["policy_name"]
        vaultName := rs.Primary.Attributes["vault_name"]

        client := testAccProvider.Meta().(*ArmClient).protectionPoliciesClient
        ctx := testAccProvider.Meta().(*ArmClient).StopContext

        if resp, err := client.Get(ctx, vaultName, resourceGroup, policyName); err != nil {
            if utils.ResponseWasNotFound(resp.Response) {
                return fmt.Errorf("Bad: Protection Policy (Policy Name %q / Resource Group %q / Vault Name %q) does not exist", policyName, resourceGroup, vaultName)
            }
            return fmt.Errorf("Bad: Get on protectionPoliciesClient: %+v", err)
        }

        return nil
    }
}

func testCheckAzureRMProtectionPolicyDestroy(s *terraform.State) error {
    client := testAccProvider.Meta().(*ArmClient).protectionPoliciesClient
    ctx := testAccProvider.Meta().(*ArmClient).StopContext

    for _, rs := range s.RootModule().Resources {
        if rs.Type != "azurerm_protection_policy" {
            continue
        }

        resourceGroup := rs.Primary.Attributes["resource_group"]
        policyName := rs.Primary.Attributes["policy_name"]
        vaultName := rs.Primary.Attributes["vault_name"]

        if resp, err := client.Get(ctx, vaultName, resourceGroup, policyName); err != nil {
            if !utils.ResponseWasNotFound(resp.Response) {
                return fmt.Errorf("Bad: Get on protectionPoliciesClient: %+v", err)
            }
        }

        return nil
    }

    return nil
}
