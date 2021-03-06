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
page_title: "Azure Resource Manager: azurerm_event_hub"
sidebar_current: "docs-azurerm-resource-event-hub"
description: |-
  Manage Azure EventHub instance.
---

# azurerm_event_hub

Manage Azure EventHub instance.


## Argument Reference

The following arguments are supported:

* `name` - (Required) The Event Hub name Changing this forces a new resource to be created.

* `resource_group` - (Required) Name of the resource group within the azure subscription. Changing this forces a new resource to be created.

* `location` - (Required) Location of the resource. Changing this forces a new resource to be created.

* `namespace_name` - (Required) The Namespace name Changing this forces a new resource to be created.

* `name` - (Optional) Name of the Event Hub. Changing this forces a new resource to be created.

* `message_retention_in_days` - (Optional) Number of days to retain the events for this Event Hub.

* `partition_count` - (Optional) Number of partitions created for the Event Hub.

* `status` - (Optional) Enumerates the possible values for the status of the Event Hub. Defaults to `Active`.

* `type` - (Optional) ARM type of the Namespace. Changing this forces a new resource to be created.

## Attributes Reference

The following attributes are exported:

* `created_at` - Exact time the Event Hub was created.

* `partition_ids` - Current number of shards on the Event Hub.

* `updated_at` - The exact time the message was updated.

* `id` - Resource Id
