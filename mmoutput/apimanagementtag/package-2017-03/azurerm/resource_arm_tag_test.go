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


func testCheckAzureRMTagExists(resourceName string) resource.TestCheckFunc {
    return func(s *terraform.State) error {
        rs, ok := s.RootModule().Resources[resourceName]
        if !ok {
            return fmt.Errorf("Tag not found: %s", resourceName)
        }

        resourceGroup := rs.Primary.Attributes["resource_group"]
        serviceName := rs.Primary.Attributes["service_name"]
        tagID := rs.Primary.Attributes["tag_id"]

        client := testAccProvider.Meta().(*ArmClient).tagClient
        ctx := testAccProvider.Meta().(*ArmClient).StopContext

        if resp, err := client.Get(ctx, resourceGroup, serviceName, tagID); err != nil {
            if utils.ResponseWasNotFound(resp.Response) {
                return fmt.Errorf("Bad: Tag (Tag %q / Service Name %q / Resource Group %q) does not exist", tagID, serviceName, resourceGroup)
            }
            return fmt.Errorf("Bad: Get on tagClient: %+v", err)
        }

        return nil
    }
}

func testCheckAzureRMTagDestroy(s *terraform.State) error {
    client := testAccProvider.Meta().(*ArmClient).tagClient
    ctx := testAccProvider.Meta().(*ArmClient).StopContext

    for _, rs := range s.RootModule().Resources {
        if rs.Type != "azurerm_tag" {
            continue
        }

        resourceGroup := rs.Primary.Attributes["resource_group"]
        serviceName := rs.Primary.Attributes["service_name"]
        tagID := rs.Primary.Attributes["tag_id"]

        if resp, err := client.Get(ctx, resourceGroup, serviceName, tagID); err != nil {
            if !utils.ResponseWasNotFound(resp.Response) {
                return fmt.Errorf("Bad: Get on tagClient: %+v", err)
            }
        }

        return nil
    }

    return nil
}