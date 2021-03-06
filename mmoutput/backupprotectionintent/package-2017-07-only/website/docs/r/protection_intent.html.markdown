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
page_title: "Azure Resource Manager: azurerm_protection_intent"
sidebar_current: "docs-azurerm-resource-protection-intent"
description: |-
  Manage Azure ProtectionIntent instance.
---

# azurerm_protection_intent

Manage Azure ProtectionIntent instance.


## Argument Reference

The following arguments are supported:

* `resource_group` - (Required) The name of the resource group where the recovery services vault is present. Changing this forces a new resource to be created.

* `fabric_name` - (Required) Fabric name associated with the backup item. Changing this forces a new resource to be created.

* `intent_object_name` - (Required) Intent object name. Changing this forces a new resource to be created.

* `vault_name` - (Required) The name of the recovery services vault. Changing this forces a new resource to be created.

* `location` - (Optional) Resource location. Changing this forces a new resource to be created.

* `backup_management_type` - (Optional) Type of backup management for the backed up item. Defaults to `Invalid`.

* `e_tag` - (Optional) Optional ETag. Changing this forces a new resource to be created.

* `item_id` - (Optional) ID of the item which is getting protected, In case of Azure Vm , it is ProtectedItemId

* `policy_id` - (Optional) ID of the backup policy with which this item is backed up.

* `protection_state` - (Optional) Backup state of this backup item. Defaults to `Invalid`.

* `source_resource_id` - (Optional) ARM ID of the resource to be backed up.

* `tags` - (Optional) Resource tags. Changing this forces a new resource to be created.

## Attributes Reference

The following attributes are exported:

* `id` - Resource Id represents the complete path to the resource.

* `name` - Resource name associated with the resource.

* `type` - Resource type represents the complete path of the form Namespace/ResourceType/ResourceType/...
