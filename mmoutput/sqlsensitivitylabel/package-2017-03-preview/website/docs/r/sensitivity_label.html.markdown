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
page_title: "Azure Resource Manager: azurerm_sensitivity_label"
sidebar_current: "docs-azurerm-resource-sensitivity-label"
description: |-
  Manage Azure SensitivityLabel instance.
---

# azurerm_sensitivity_label

Manage Azure SensitivityLabel instance.


## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the column. Changing this forces a new resource to be created.

* `resource_group` - (Required) The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal. Changing this forces a new resource to be created.

* `database_name` - (Required) The name of the database. Changing this forces a new resource to be created.

* `schema_name` - (Required) The name of the schema. Changing this forces a new resource to be created.

* `sensitivity_label_source` - (Required) The source of the sensitivity label. Changing this forces a new resource to be created.

* `server_name` - (Required) The name of the server. Changing this forces a new resource to be created.

* `table_name` - (Required) The name of the table. Changing this forces a new resource to be created.

* `information_type` - (Optional) The information type.

* `information_type_id` - (Optional) The information type ID.

* `label_id` - (Optional) The label ID.

* `label_name` - (Optional) The label name.

## Attributes Reference

The following attributes are exported:

* `is_disabled` - Is sensitivity recommendation disabled. Applicable for recommended sensitivity label only. Specifies whether the sensitivity recommendation on this column is disabled (dismissed) or not.

* `id` - Resource ID.

* `name` - Resource name.

* `type` - Resource type.
