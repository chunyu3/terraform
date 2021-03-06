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
page_title: "Azure Resource Manager: azurerm_volume"
sidebar_current: "docs-azurerm-resource-volume"
description: |-
  Manage Azure Volume instance.
---

# azurerm_volume

Manage Azure Volume instance.


## Argument Reference

The following arguments are supported:

* `name` - (Required) The volume name. Changing this forces a new resource to be created.

* `resource_group` - (Required) The resource group name Changing this forces a new resource to be created.

* `access_control_record_ids` - (Required) The IDs of the access control records, associated with the volume.

* `device_name` - (Required) The device name Changing this forces a new resource to be created.

* `manager_name` - (Required) The manager name Changing this forces a new resource to be created.

* `monitoring_status` - (Required) The monitoring status of the volume.

* `size_in_bytes` - (Required) The size of the volume in bytes.

* `volume_container_name` - (Required) The volume container name. Changing this forces a new resource to be created.

* `volume_status` - (Required) The volume status.

* `volume_type` - (Required) The type of the volume.

* `kind` - (Optional) The Kind of the object. Currently only Series8000 is supported Defaults to `Series8000`. Changing this forces a new resource to be created.

## Attributes Reference

The following attributes are exported:

* `volume_container_id` - The ID of the volume container, in which this volume is created.

* `operation_status` - The operation status on the volume.

* `backup_status` - The backup status of the volume.

* `backup_policy_ids` - The IDs of the backup policies, in which this volume is part of.

* `id` - The path ID that uniquely identifies the object.

* `name` - The name of the object.

* `type` - The hierarchical type of the object.
