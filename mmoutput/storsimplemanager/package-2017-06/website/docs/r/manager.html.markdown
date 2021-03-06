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
page_title: "Azure Resource Manager: azurerm_manager"
sidebar_current: "docs-azurerm-resource-manager"
description: |-
  Manage Azure Manager instance.
---

# azurerm_manager

Manage Azure Manager instance.


## Argument Reference

The following arguments are supported:

* `name` - (Required) The manager name Changing this forces a new resource to be created.

* `resource_group` - (Required) The resource group name Changing this forces a new resource to be created.

* `location` - (Required) The geo location of the resource. Changing this forces a new resource to be created.

* `cis_intrinsic_settings` - (Optional) One `cis_intrinsic_setting` block defined below.

* `etag` - (Optional) The etag of the manager. Changing this forces a new resource to be created.

* `sku` - (Optional) One `sku` block defined below.

* `tags` - (Optional) The tags attached to the Manager. Changing this forces a new resource to be created.

---

The `cis_intrinsic_setting` block supports the following:

* `type` - (Required) The type of StorSimple Manager.

---

The `sku` block supports the following:

* `name` - (Required) Refers to the sku name which should be "Standard"

## Attributes Reference

The following attributes are exported:

* `provisioning_state` - Specifies the state of the resource as it is getting provisioned. Value of "Succeeded" means the Manager was successfully created.

* `id` - The resource ID.

* `name` - The resource name.

* `type` - The resource type.
