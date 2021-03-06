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
page_title: "Azure Resource Manager: azurerm_express_route_circuit_authorization"
sidebar_current: "docs-azurerm-resource-express-route-circuit-authorization"
description: |-
  Manage Azure ExpressRouteCircuitAuthorization instance.
---

# azurerm_express_route_circuit_authorization

Manage Azure ExpressRouteCircuitAuthorization instance.


## Argument Reference

The following arguments are supported:

* `resource_group` - (Required) The name of the resource group. Changing this forces a new resource to be created.

* `authorization_name` - (Required) The name of the authorization. Changing this forces a new resource to be created.

* `circuit_name` - (Required) The name of the express route circuit. Changing this forces a new resource to be created.

* `id` - (Optional) Resource Id Changing this forces a new resource to be created.

* `name` - (Optional) Gets name of the resource that is unique within a resource group. This name can be used to access the resource Changing this forces a new resource to be created.

* `authorization_key` - (Optional) Gets or sets the authorization key

* `authorization_use_status` - (Optional) Gets or sets AuthorizationUseStatus Defaults to `Available`.

* `etag` - (Optional) A unique read-only string that changes whenever the resource is updated Changing this forces a new resource to be created.

## Attributes Reference

The following attributes are exported:

* `provisioning_state` - Gets provisioning state of the PublicIP resource Updating/Deleting/Failed
