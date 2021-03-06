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
page_title: "Azure Resource Manager: azurerm_access_control_record"
sidebar_current: "docs-azurerm-resource-access-control-record"
description: |-
  Manage Azure AccessControlRecord instance.
---

# azurerm_access_control_record

Manage Azure AccessControlRecord instance.


## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the access control record to delete. Changing this forces a new resource to be created.

* `resource_group` - (Required) The resource group name Changing this forces a new resource to be created.

* `initiator_name` - (Required) The iSCSI initiator name (IQN).

* `manager_name` - (Required) The manager name Changing this forces a new resource to be created.

* `kind` - (Optional) The Kind of the object. Currently only Series8000 is supported Defaults to `Series8000`. Changing this forces a new resource to be created.

## Attributes Reference

The following attributes are exported:

* `volume_count` - The number of volumes using the access control record.

* `id` - The path ID that uniquely identifies the object.

* `name` - The name of the object.

* `type` - The hierarchical type of the object.
