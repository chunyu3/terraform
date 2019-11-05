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


func testCheckAzureRMRedisLinkedServerExists(resourceName string) resource.TestCheckFunc {
    return func(s *terraform.State) error {
        rs, ok := s.RootModule().Resources[resourceName]
        if !ok {
            return fmt.Errorf("Redis Linked Server not found: %s", resourceName)
        }

        name := rs.Primary.Attributes["name"]
        resourceGroup := rs.Primary.Attributes["resource_group"]
        linkedServerName := rs.Primary.Attributes["linked_server_name"]

        client := testAccProvider.Meta().(*ArmClient).redisLinkedServerClient
        ctx := testAccProvider.Meta().(*ArmClient).StopContext

        if resp, err := client.Get(ctx, resourceGroup, name, linkedServerName); err != nil {
            if utils.ResponseWasNotFound(resp.Response) {
                return fmt.Errorf("Bad: Redis Linked Server %q (Linked Server Name %q / Resource Group %q) does not exist", name, linkedServerName, resourceGroup)
            }
            return fmt.Errorf("Bad: Get on redisLinkedServerClient: %+v", err)
        }

        return nil
    }
}

func testCheckAzureRMRedisLinkedServerDestroy(s *terraform.State) error {
    client := testAccProvider.Meta().(*ArmClient).redisLinkedServerClient
    ctx := testAccProvider.Meta().(*ArmClient).StopContext

    for _, rs := range s.RootModule().Resources {
        if rs.Type != "azurerm_redis_linked_server" {
            continue
        }

        name := rs.Primary.Attributes["name"]
        resourceGroup := rs.Primary.Attributes["resource_group"]
        linkedServerName := rs.Primary.Attributes["linked_server_name"]

        if resp, err := client.Get(ctx, resourceGroup, name, linkedServerName); err != nil {
            if !utils.ResponseWasNotFound(resp.Response) {
                return fmt.Errorf("Bad: Get on redisLinkedServerClient: %+v", err)
            }
        }

        return nil
    }

    return nil
}
