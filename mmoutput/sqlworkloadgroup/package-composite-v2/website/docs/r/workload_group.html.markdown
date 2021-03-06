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
page_title: "Azure Resource Manager: azurerm_workload_group"
sidebar_current: "docs-azurerm-resource-workload-group"
description: |-
  Manage Azure WorkloadGroup instance.
---

# azurerm_workload_group

Manage Azure WorkloadGroup instance.


## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the workload group to delete. Changing this forces a new resource to be created.

* `resource_group` - (Required) The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal. Changing this forces a new resource to be created.

* `database_name` - (Required) The name of the database. Changing this forces a new resource to be created.

* `max_resource_percent` - (Required) The workload group cap percentage resource.

* `min_resource_percent` - (Required) The workload group minimum percentage resource.

* `min_resource_percent_per_request` - (Required) The workload group request minimum grant percentage.

* `server_name` - (Required) The name of the server. Changing this forces a new resource to be created.

* `importance` - (Optional) The workload group importance level.

* `max_resource_percent_per_request` - (Optional) The workload group request maximum grant percentage.

* `query_execution_timeout` - (Optional) The workload group query execution timeout.

## Attributes Reference

The following attributes are exported:

* `id` - Resource ID.

* `name` - Resource name.

* `type` - Resource type.
