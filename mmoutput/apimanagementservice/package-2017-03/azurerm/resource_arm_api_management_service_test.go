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


func testCheckAzureRMApiManagementServiceExists(resourceName string) resource.TestCheckFunc {
    return func(s *terraform.State) error {
        rs, ok := s.RootModule().Resources[resourceName]
        if !ok {
            return fmt.Errorf("Api Management Service not found: %s", resourceName)
        }

        resourceGroup := rs.Primary.Attributes["resource_group"]
        serviceName := rs.Primary.Attributes["service_name"]

        client := testAccProvider.Meta().(*ArmClient).apiManagementServiceClient
        ctx := testAccProvider.Meta().(*ArmClient).StopContext

        if resp, err := client.Get(ctx, resourceGroup, serviceName); err != nil {
            if utils.ResponseWasNotFound(resp.Response) {
                return fmt.Errorf("Bad: Api Management Service (Service Name %q / Resource Group %q) does not exist", serviceName, resourceGroup)
            }
            return fmt.Errorf("Bad: Get on apiManagementServiceClient: %+v", err)
        }

        return nil
    }
}

func testCheckAzureRMApiManagementServiceDestroy(s *terraform.State) error {
    client := testAccProvider.Meta().(*ArmClient).apiManagementServiceClient
    ctx := testAccProvider.Meta().(*ArmClient).StopContext

    for _, rs := range s.RootModule().Resources {
        if rs.Type != "azurerm_api_management_service" {
            continue
        }

        resourceGroup := rs.Primary.Attributes["resource_group"]
        serviceName := rs.Primary.Attributes["service_name"]

        if resp, err := client.Get(ctx, resourceGroup, serviceName); err != nil {
            if !utils.ResponseWasNotFound(resp.Response) {
                return fmt.Errorf("Bad: Get on apiManagementServiceClient: %+v", err)
            }
        }

        return nil
    }

    return nil
}