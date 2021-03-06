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
page_title: "Azure Resource Manager: azurerm_workspace"
sidebar_current: "docs-azurerm-resource-workspace"
description: |-
  Manage Azure Workspace instance.
---

# azurerm_workspace

Manage Azure Workspace instance.


## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the workspace. Workspace names can only contain a combination of alphanumeric characters along with dash (-) and underscore (_). The name must be from 1 through 64 characters long. Changing this forces a new resource to be created.

* `resource_group` - (Required) Name of the resource group to which the resource belongs. Changing this forces a new resource to be created.

* `location` - (Required) The region in which to create the Workspace. Changing this forces a new resource to be created.

* `tags` - (Optional) The user specified tags associated with the Workspace. Changing this forces a new resource to be created.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the resource

* `name` - The name of the resource

* `type` - The type of the resource

* `creation_time` - Time when the Workspace was created.

* `provisioning_state` - The provisioned state of the Workspace

* `provisioning_state_transition_time` - The time at which the workspace entered its current provisioning state.
