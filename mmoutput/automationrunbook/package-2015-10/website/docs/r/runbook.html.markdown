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
page_title: "Azure Resource Manager: azurerm_runbook"
sidebar_current: "docs-azurerm-resource-runbook"
description: |-
  Manage Azure Runbook instance.
---

# azurerm_runbook

Manage Azure Runbook instance.


## Argument Reference

The following arguments are supported:

* `name` - (Required) The runbook name. Changing this forces a new resource to be created.

* `resource_group` - (Required) Name of an Azure Resource group. Changing this forces a new resource to be created.

* `automation_account_name` - (Required) The name of the automation account. Changing this forces a new resource to be created.

* `name` - (Optional) Gets or sets the name of the resource. Changing this forces a new resource to be created.

* `location` - (Optional) Gets or sets the location of the resource. Changing this forces a new resource to be created.

* `description` - (Optional) Gets or sets the description of the runbook.

* `log_activity_trace` - (Optional) Gets or sets the activity-level tracing options of the runbook.

* `log_progress` - (Optional) Gets or sets progress log option.

* `log_verbose` - (Optional) Gets or sets verbose log option.

* `tags` - (Optional) Gets or sets the tags attached to the resource. Changing this forces a new resource to be created.

## Attributes Reference

The following attributes are exported:

* `id` - Fully qualified resource Id for the resource

* `type` - The type of the resource.
