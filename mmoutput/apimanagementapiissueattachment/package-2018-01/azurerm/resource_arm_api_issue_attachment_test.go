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


func testCheckAzureRMApiIssueAttachmentExists(resourceName string) resource.TestCheckFunc {
    return func(s *terraform.State) error {
        rs, ok := s.RootModule().Resources[resourceName]
        if !ok {
            return fmt.Errorf("Api Issue Attachment not found: %s", resourceName)
        }

        resourceGroup := rs.Primary.Attributes["resource_group"]
        apiID := rs.Primary.Attributes["api_id"]
        attachmentID := rs.Primary.Attributes["attachment_id"]
        issueID := rs.Primary.Attributes["issue_id"]
        serviceName := rs.Primary.Attributes["service_name"]

        client := testAccProvider.Meta().(*ArmClient).apiIssueAttachmentClient
        ctx := testAccProvider.Meta().(*ArmClient).StopContext

        if resp, err := client.Get(ctx, resourceGroup, serviceName, apiID, issueID, attachmentID); err != nil {
            if utils.ResponseWasNotFound(resp.Response) {
                return fmt.Errorf("Bad: Api Issue Attachment (Attachment %q / Issue %q / Api %q / Service Name %q / Resource Group %q) does not exist", attachmentID, issueID, apiID, serviceName, resourceGroup)
            }
            return fmt.Errorf("Bad: Get on apiIssueAttachmentClient: %+v", err)
        }

        return nil
    }
}

func testCheckAzureRMApiIssueAttachmentDestroy(s *terraform.State) error {
    client := testAccProvider.Meta().(*ArmClient).apiIssueAttachmentClient
    ctx := testAccProvider.Meta().(*ArmClient).StopContext

    for _, rs := range s.RootModule().Resources {
        if rs.Type != "azurerm_api_issue_attachment" {
            continue
        }

        resourceGroup := rs.Primary.Attributes["resource_group"]
        apiID := rs.Primary.Attributes["api_id"]
        attachmentID := rs.Primary.Attributes["attachment_id"]
        issueID := rs.Primary.Attributes["issue_id"]
        serviceName := rs.Primary.Attributes["service_name"]

        if resp, err := client.Get(ctx, resourceGroup, serviceName, apiID, issueID, attachmentID); err != nil {
            if !utils.ResponseWasNotFound(resp.Response) {
                return fmt.Errorf("Bad: Get on apiIssueAttachmentClient: %+v", err)
            }
        }

        return nil
    }

    return nil
}