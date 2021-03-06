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
page_title: "Azure Resource Manager: azurerm_instance_pool"
sidebar_current: "docs-azurerm-resource-instance-pool"
description: |-
  Manage Azure InstancePool instance.
---

# azurerm_instance_pool

Manage Azure InstancePool instance.


## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the instance pool to be updated. Changing this forces a new resource to be created.

* `resource_group` - (Required) The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal. Changing this forces a new resource to be created.

* `location` - (Required) Resource location. Changing this forces a new resource to be created.

* `license_type` - (Required) The license type. Possible values are 'LicenseIncluded' (price for SQL license is included) and 'BasePrice' (without SQL license price).

* `subnet_id` - (Required) Resource ID of the subnet to place this instance pool in.

* `v_cores` - (Required) Count of vCores belonging to this instance pool.

* `sku` - (Optional) One `sku` block defined below.

* `tags` - (Optional) Resource tags. Changing this forces a new resource to be created.

---

The `sku` block supports the following:

* `name` - (Required) The name of the SKU, typically, a letter + Number code, e.g. P3.

* `tier` - (Optional) The tier or edition of the particular SKU, e.g. Basic, Premium.

* `size` - (Optional) Size of the particular SKU

* `family` - (Optional) If the service has different generations of hardware, for the same SKU, then that can be captured here.

* `capacity` - (Optional) Capacity of the particular SKU.

## Attributes Reference

The following attributes are exported:

* `id` - Resource ID.

* `name` - Resource name.

* `type` - Resource type.
