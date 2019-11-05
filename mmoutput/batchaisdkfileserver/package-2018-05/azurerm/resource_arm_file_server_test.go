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


func testCheckAzureRMFileServerExists(resourceName string) resource.TestCheckFunc {
    return func(s *terraform.State) error {
        rs, ok := s.RootModule().Resources[resourceName]
        if !ok {
            return fmt.Errorf("File Server not found: %s", resourceName)
        }

        resourceGroup := rs.Primary.Attributes["resource_group"]
        fileServerName := rs.Primary.Attributes["file_server_name"]
        workspaceName := rs.Primary.Attributes["workspace_name"]

        client := testAccProvider.Meta().(*ArmClient).fileServersClient
        ctx := testAccProvider.Meta().(*ArmClient).StopContext

        if resp, err := client.Get(ctx, resourceGroup, workspaceName, fileServerName); err != nil {
            if utils.ResponseWasNotFound(resp.Response) {
                return fmt.Errorf("Bad: File Server (File Server Name %q / Workspace Name %q / Resource Group %q) does not exist", fileServerName, workspaceName, resourceGroup)
            }
            return fmt.Errorf("Bad: Get on fileServersClient: %+v", err)
        }

        return nil
    }
}

func testCheckAzureRMFileServerDestroy(s *terraform.State) error {
    client := testAccProvider.Meta().(*ArmClient).fileServersClient
    ctx := testAccProvider.Meta().(*ArmClient).StopContext

    for _, rs := range s.RootModule().Resources {
        if rs.Type != "azurerm_file_server" {
            continue
        }

        resourceGroup := rs.Primary.Attributes["resource_group"]
        fileServerName := rs.Primary.Attributes["file_server_name"]
        workspaceName := rs.Primary.Attributes["workspace_name"]

        if resp, err := client.Get(ctx, resourceGroup, workspaceName, fileServerName); err != nil {
            if !utils.ResponseWasNotFound(resp.Response) {
                return fmt.Errorf("Bad: Get on fileServersClient: %+v", err)
            }
        }

        return nil
    }

    return nil
}
