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
page_title: "Azure Resource Manager: azurerm_group"
sidebar_current: "docs-azurerm-resource-group"
description: |-
  Manage Azure Group instance.
---

# azurerm_group

Manage Azure Group instance.


## Argument Reference

The following arguments are supported:

* `name` - (Required) Unique name of a group within a project. Changing this forces a new resource to be created.

* `resource_group` - (Required) Name of the Azure Resource Group that project is part of. Changing this forces a new resource to be created.

* `project_name` - (Required) Name of the Azure Migrate project. Changing this forces a new resource to be created.

* `e_tag` - (Optional) For optimistic concurrency control. Changing this forces a new resource to be created.

* `machines` - (Optional) List of machine names that are part of this group.

* `operation_type` - (Optional) Whether to add or remove the machines. Defaults to `Add`.

## Attributes Reference

The following attributes are exported:

* `group_status` - Whether the group has been created and is valid.

* `machine_count` - Number of machines part of this group.

* `assessments` - List of References to Assessments created on this group.

* `are_assessments_running` - If the assessments are in running state.

* `created_timestamp` - Time when this group was created. Date-Time represented in ISO-8601 format.

* `updated_timestamp` - Time when this group was last updated. Date-Time represented in ISO-8601 format.

* `id` - Path reference to this group. /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}/groups/{groupName}

* `name` - Name of the group.

* `type` - Type of the object = [Microsoft.Migrate/assessmentProjects/groups].
