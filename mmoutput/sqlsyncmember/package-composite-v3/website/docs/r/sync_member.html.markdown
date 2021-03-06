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
page_title: "Azure Resource Manager: azurerm_sync_member"
sidebar_current: "docs-azurerm-resource-sync-member"
description: |-
  Manage Azure SyncMember instance.
---

# azurerm_sync_member

Manage Azure SyncMember instance.


## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the sync member. Changing this forces a new resource to be created.

* `resource_group` - (Required) The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal. Changing this forces a new resource to be created.

* `database_name` - (Required) The name of the database on which the sync group is hosted. Changing this forces a new resource to be created.

* `server_name` - (Required) The name of the server. Changing this forces a new resource to be created.

* `sync_group_name` - (Required) The name of the sync group on which the sync member is hosted. Changing this forces a new resource to be created.

* `database_name` - (Optional) Database name of the member database in the sync member.

* `database_type` - (Optional) Database type of the sync member. Defaults to `AzureSqlDatabase`.

* `password` - (Optional) Password of the member database in the sync member.

* `server_name` - (Optional) Server name of the member database in the sync member

* `sql_server_database_id` - (Optional) SQL Server database id of the sync member.

* `sync_agent_id` - (Optional) ARM resource id of the sync agent in the sync member.

* `sync_direction` - (Optional) Sync direction of the sync member. Defaults to `Bidirectional`.

* `user_name` - (Optional) User name of the member database in the sync member.

## Attributes Reference

The following attributes are exported:

* `id` - Resource ID.

* `name` - Resource name.

* `type` - Resource type.
