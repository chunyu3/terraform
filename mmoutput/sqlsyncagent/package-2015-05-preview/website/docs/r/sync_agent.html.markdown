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
page_title: "Azure Resource Manager: azurerm_sync_agent"
sidebar_current: "docs-azurerm-resource-sync-agent"
description: |-
  Manage Azure SyncAgent instance.
---

# azurerm_sync_agent

Manage Azure SyncAgent instance.


## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the sync agent. Changing this forces a new resource to be created.

* `resource_group` - (Required) The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal. Changing this forces a new resource to be created.

* `server_name` - (Required) The name of the server on which the sync agent is hosted. Changing this forces a new resource to be created.

* `sync_database_id` - (Optional) ARM resource id of the sync database in the sync agent.

## Attributes Reference

The following attributes are exported:

* `name` - Name of the sync agent.

* `last_alive_time` - Last alive time of the sync agent.

* `state` - State of the sync agent.

* `is_up_to_date` - If the sync agent version is up to date.

* `expiry_time` - Expiration time of the sync agent version.

* `version` - Version of the sync agent.

* `id` - Resource ID.

* `name` - Resource name.

* `type` - Resource type.
