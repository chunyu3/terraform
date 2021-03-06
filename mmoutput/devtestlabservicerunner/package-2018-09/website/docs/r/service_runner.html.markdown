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
page_title: "Azure Resource Manager: azurerm_service_runner"
sidebar_current: "docs-azurerm-resource-service-runner"
description: |-
  Manage Azure ServiceRunner instance.
---

# azurerm_service_runner

Manage Azure ServiceRunner instance.


## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the lab. Changing this forces a new resource to be created.

* `name` - (Required) The name of the service runner. Changing this forces a new resource to be created.

* `resource_group` - (Required) The name of the resource group. Changing this forces a new resource to be created.

* `location` - (Optional) The location of the resource. Changing this forces a new resource to be created.

* `identity` - (Optional) One `identity` block defined below.

* `tags` - (Optional) The tags of the resource. Changing this forces a new resource to be created.

---

The `identity` block supports the following:

* `type` - (Optional) Managed identity.

* `principal_id` - (Optional) The principal id of resource identity.

* `tenant_id` - (Optional) The tenant identifier of resource.

* `client_secret_url` - (Optional) The client secret URL of the identity.

## Attributes Reference

The following attributes are exported:

* `id` - The identifier of the resource.

* `name` - The name of the resource.

* `type` - The type of the resource.
