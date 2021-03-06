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
page_title: "Azure Resource Manager: azurerm_iot_hub_resource"
sidebar_current: "docs-azurerm-resource-iot-hub-resource"
description: |-
  Manage Azure IotHubResource instance.
---

# azurerm_iot_hub_resource

Manage Azure IotHubResource instance.


## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the IoT hub to create or update. Changing this forces a new resource to be created.

* `resource_group` - (Required) The name of the resource group that contains the IoT hub. Changing this forces a new resource to be created.

* `location` - (Required) The resource location. Changing this forces a new resource to be created.

* `resourcegroup` - (Required) The name of the resource group that contains the IoT hub. A resource group name uniquely identifies the resource group within the subscription. Changing this forces a new resource to be created.

* `sku` - (Required) One `sku` block defined below.

* `subscriptionid` - (Required) The subscription identifier. Changing this forces a new resource to be created.

---

The `sku` block supports the following:

* `name` - (Required) The name of the SKU.

* `capacity` - (Required) The number of provisioned IoT Hub units. See: https://docs.microsoft.com/azure/azure-subscription-service-limits#iot-hub-limits.

* `authorization_policies` - (Optional) One or more `authorization_policy` block defined below.

* `cloud_to_device` - (Optional) One `cloud_to_device` block defined below.

* `comments` - (Optional) Comments.

* `enable_file_upload_notifications` - (Optional) If True, file upload notifications are enabled.

* `etag` - (Optional) The Etag field is *not* required. If it is provided in the response body, it must also be provided as a header per the normal ETag convention. Changing this forces a new resource to be created.

* `event_hub_endpoints` - (Optional) The Event Hub-compatible endpoint properties. The possible keys to this dictionary are events and operationsMonitoringEvents. Both of these keys have to be present in the dictionary while making create or update calls for the IoT hub.

* `features` - (Optional) The capabilities and features enabled for the IoT hub. Defaults to `None`.

* `ip_filter_rules` - (Optional) One or more `ip_filter_rule` block defined below.

* `messaging_endpoints` - (Optional) The messaging endpoint properties for the file upload notification queue.

* `operations_monitoring_properties` - (Optional) One `operations_monitoring_property` block defined below.

* `storage_endpoints` - (Optional) The list of Azure Storage endpoints where you can upload files. Currently you can configure only one Azure Storage account and that MUST have its key as $default. Specifying more than one storage account causes an error to be thrown. Not specifying a value for this property when the enableFileUploadNotifications property is set to True, causes an error to be thrown.

* `tags` - (Optional) The resource tags. Changing this forces a new resource to be created.

---

The `authorization_policy` block supports the following:

* `key_name` - (Required) The name of the shared access policy.

* `primary_key` - (Optional) The primary key.

* `secondary_key` - (Optional) The secondary key.

* `rights` - (Required) The permissions assigned to the shared access policy.

---

The `cloud_to_device` block supports the following:

* `max_delivery_count` - (Optional) The max delivery count for cloud-to-device messages in the device queue. See: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-messaging#cloud-to-device-messages.

* `default_ttl_as_iso8601` - (Optional) The default time to live for cloud-to-device messages in the device queue. See: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-messaging#cloud-to-device-messages.

* `feedback` - (Optional) One `feedback` block defined below.


---

The `feedback` block supports the following:

* `lock_duration_as_iso8601` - (Optional) The lock duration for the feedback queue. See: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-messaging#cloud-to-device-messages.

* `ttl_as_iso8601` - (Optional) The period of time for which a message is available to consume before it is expired by the IoT hub. See: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-messaging#cloud-to-device-messages.

* `max_delivery_count` - (Optional) The number of times the IoT hub attempts to deliver a message on the feedback queue. See: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-messaging#cloud-to-device-messages.

---

The `ip_filter_rule` block supports the following:

* `filter_name` - (Required) The name of the IP filter rule.

* `action` - (Required) The desired action for requests captured by this rule.

* `ip_mask` - (Required) A string that contains the IP address range in CIDR notation for the rule.

---

The `operations_monitoring_property` block supports the following:

* `events` - (Optional) 

## Attributes Reference

The following attributes are exported:

* `provisioning_state` - The provisioning state.

* `host_name` - The name of the host.

* `id` - The resource identifier.

* `name` - The resource name.

* `type` - The resource type.
