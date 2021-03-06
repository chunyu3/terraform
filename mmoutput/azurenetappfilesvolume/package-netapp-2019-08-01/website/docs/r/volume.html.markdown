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

* `name` - (Required) The name of the volume Changing this forces a new resource to be created.

* `resource_group` - (Required) The name of the resource group. Changing this forces a new resource to be created.

* `account_name` - (Required) The name of the NetApp account Changing this forces a new resource to be created.

* `pool_name` - (Required) The name of the capacity pool Changing this forces a new resource to be created.

* `location` - (Optional) Resource location Changing this forces a new resource to be created.

* `export_policy` - (Optional) One `export_policy` block defined below.

* `service_level` - (Optional) The service level of the file system Defaults to `Standard`.

* `usage_threshold` - (Optional) Maximum storage quota allowed for a file system in bytes. This is a soft quota used for alerting only. Minimum size is 100 GiB. Upper limit is 100TiB. Specified in bytes.

* `tags` - (Optional) Resource tags Changing this forces a new resource to be created.

---

The `export_policy` block supports the following:

* `rules` - (Optional) One or more `rule` block defined below.


---

The `rule` block supports the following:

* `rule_index` - (Optional) Order index

* `unix_read_only` - (Optional) Read only access

* `unix_read_write` - (Optional) Read and write access

* `cifs` - (Optional) Allows CIFS protocol

* `nfsv3` - (Optional) Allows NFSv3 protocol

* `nfsv41` - (Optional) Allows NFSv4.1 protocol

* `allowed_clients` - (Optional) Client ingress specification as comma separated string with IPv4 CIDRs, IPv4 host addresses and host names

## Attributes Reference

The following attributes are exported:

* `id` - Resource Id

* `name` - Resource name

* `type` - Resource type
