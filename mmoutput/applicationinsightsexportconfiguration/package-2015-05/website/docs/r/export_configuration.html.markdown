---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
#
# ----------------------------------------------------------------------------
#
#     This file is automatically generated by Magic Modules and manual
#     changes will be clobbered when the file is regenerated.
#
#     Please read more about how to change this file at
#     https://github.com/Azure/magic-module-specs
#
# ----------------------------------------------------------------------------
layout: "azurerm"
page_title: "Azure Resource Manager: azurerm_export_configuration"
sidebar_current: "docs-azurerm-resource-export-configuration"
description: |-
  Manage Azure ExportConfiguration instance.
---

# azurerm_export_configuration

Manage Azure ExportConfiguration instance.


## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the Application Insights component resource. Changing this forces a new resource to be created.

* `resource_group` - (Required) The name of the resource group. The name is case insensitive. Changing this forces a new resource to be created.

* `destination_account_id` - (Optional) The name of destination storage account. Changing this forces a new resource to be created.

* `destination_address` - (Optional) The SAS URL for the destination storage container. It must grant write permission. Changing this forces a new resource to be created.

* `destination_storage_location_id` - (Optional) The location ID of the destination storage container. Changing this forces a new resource to be created.

* `destination_storage_subscription_id` - (Optional) The subscription ID of the destination storage container. Changing this forces a new resource to be created.

* `destination_type` - (Optional) The Continuous Export destination type. This has to be 'Blob'. Changing this forces a new resource to be created.

* `is_enabled` - (Optional) Set to 'true' to create a Continuous Export configuration as enabled, otherwise set it to 'false'. Changing this forces a new resource to be created.

* `notification_queue_enabled` - (Optional) Deprecated Changing this forces a new resource to be created.

* `notification_queue_uri` - (Optional) Deprecated Changing this forces a new resource to be created.

* `record_types` - (Optional) The document types to be exported, as comma separated values. Allowed values include 'Requests', 'Event', 'Exceptions', 'Metrics', 'PageViews', 'PageViewPerformance', 'Rdd', 'PerformanceCounters', 'Availability', 'Messages'. Changing this forces a new resource to be created.

## Attributes Reference

The following attributes are exported:

* `export_id` - The unique ID of the export configuration inside an Application Insights component. It is auto generated when the Continuous Export configuration is created.

* `instrumentation_key` - The instrumentation key of the Application Insights component.

* `application_name` - The name of the Application Insights component.

* `subscription_id` - The subscription of the Application Insights component.

* `resource_group` - The resource group of the Application Insights component.

* `is_user_enabled` - This will be 'true' if the Continuous Export configuration is enabled, otherwise it will be 'false'.

* `last_user_update` - Last time the Continuous Export configuration was updated.

* `export_status` - This indicates current Continuous Export configuration status. The possible values are 'Preparing', 'Success', 'Failure'.

* `last_success_time` - The last time data was successfully delivered to the destination storage container for this Continuous Export configuration.

* `last_gap_time` - The last time the Continuous Export configuration started failing.

* `permanent_error_reason` - This is the reason the Continuous Export configuration started failing. It can be 'AzureStorageNotFound' or 'AzureStorageAccessDenied'.

* `storage_name` - The name of the destination storage account.

* `container_name` - The name of the destination storage container.
