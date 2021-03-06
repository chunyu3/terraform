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
page_title: "Azure Resource Manager: azurerm_share"
sidebar_current: "docs-azurerm-resource-share"
description: |-
  Manage Azure Share instance.
---

# azurerm_share

Manage Azure Share instance.


## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the share. Changing this forces a new resource to be created.

* `resource_group` - (Required) The resource group name. Changing this forces a new resource to be created.

* `account_name` - (Required) The name of the share account. Changing this forces a new resource to be created.

* `consumer_email` - (Optional) Email of the user who created the synchronization Changing this forces a new resource to be created.

* `consumer_name` - (Optional) Name of the user who created the synchronization Changing this forces a new resource to be created.

* `consumer_tenant_name` - (Optional) Tenant name of the consumer who created the synchronization Changing this forces a new resource to be created.

* `description` - (Optional) Share description.

* `duration_ms` - (Optional) synchronization duration Changing this forces a new resource to be created.

* `end_time` - (Optional) End time of synchronization Changing this forces a new resource to be created.

* `message` - (Optional) message of synchronization Changing this forces a new resource to be created.

* `share_kind` - (Optional) Share kind. Defaults to `CopyBased`.

* `start_time` - (Optional) start time of synchronization Changing this forces a new resource to be created.

* `status` - (Optional) Raw Status Changing this forces a new resource to be created.

* `synchronization_id` - (Optional) Synchronization id Changing this forces a new resource to be created.

* `terms` - (Optional) Share terms.

## Attributes Reference

The following attributes are exported:

* `created_at` - Time at which the share was created.

* `provisioning_state` - Gets or sets the provisioning state

* `user_email` - Email of the user who created the resource

* `user_name` - Name of the user who created the resource

* `id` - The resource id of the azure resource

* `name` - Name of the azure resource

* `type` - Type of the azure resource
