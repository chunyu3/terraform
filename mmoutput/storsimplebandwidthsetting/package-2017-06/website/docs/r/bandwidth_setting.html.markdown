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
page_title: "Azure Resource Manager: azurerm_bandwidth_setting"
sidebar_current: "docs-azurerm-resource-bandwidth-setting"
description: |-
  Manage Azure BandwidthSetting instance.
---

# azurerm_bandwidth_setting

Manage Azure BandwidthSetting instance.


## Argument Reference

The following arguments are supported:

* `name` - (Required) The manager name Changing this forces a new resource to be created.

* `resource_group` - (Required) The resource group name Changing this forces a new resource to be created.

* `bandwidth_setting_name` - (Required) The bandwidth setting name. Changing this forces a new resource to be created.

* `schedules` - (Required) One or more `schedule` block defined below.

---

The `schedule` block supports the following:

* `start` - (Required) One `start` block defined below.

* `stop` - (Required) One `stop` block defined below.

* `rate_in_mbps` - (Required) The rate in Mbps.

* `days` - (Required) The days of the week when this schedule is applicable.


---

The `start` block supports the following:

* `hours` - (Required) The hour.

* `minutes` - (Required) The minute.

* `seconds` - (Required) The second.

---

The `stop` block supports the following:

* `hours` - (Required) The hour.

* `minutes` - (Required) The minute.

* `seconds` - (Required) The second.

* `kind` - (Optional) The Kind of the object. Currently only Series8000 is supported Defaults to `Series8000`. Changing this forces a new resource to be created.

## Attributes Reference

The following attributes are exported:

* `volume_count` - The number of volumes that uses the bandwidth setting.

* `id` - The path ID that uniquely identifies the object.

* `name` - The name of the object.

* `type` - The hierarchical type of the object.
