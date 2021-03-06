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
page_title: "Azure Resource Manager: azurerm_content_key_policy"
sidebar_current: "docs-azurerm-resource-content-key-policy"
description: |-
  Manage Azure ContentKeyPolicy instance.
---

# azurerm_content_key_policy

Manage Azure ContentKeyPolicy instance.


## Argument Reference

The following arguments are supported:

* `name` - (Required) The Content Key Policy name. Changing this forces a new resource to be created.

* `resource_group` - (Required) The name of the resource group within the Azure subscription. Changing this forces a new resource to be created.

* `account_name` - (Required) The Media Services account name. Changing this forces a new resource to be created.

* `options` - (Required) One or more `option` block defined below.

---

The `option` block supports the following:

* `name` - (Optional) The Policy Option description.

* `description` - (Optional) A description for the Policy.

## Attributes Reference

The following attributes are exported:

* `policy_id` - The legacy Policy ID.

* `created` - The creation date of the Policy

* `last_modified` - The last modified date of the Policy

* `id` - Fully qualified resource ID for the resource.

* `name` - The name of the resource.

* `type` - The type of the resource.
