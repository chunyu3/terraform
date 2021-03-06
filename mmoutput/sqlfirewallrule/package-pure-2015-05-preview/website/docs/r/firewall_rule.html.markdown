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
page_title: "Azure Resource Manager: azurerm_firewall_rule"
sidebar_current: "docs-azurerm-resource-firewall-rule"
description: |-
  Manage Azure FirewallRule instance.
---

# azurerm_firewall_rule

Manage Azure FirewallRule instance.


## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the firewall rule. Changing this forces a new resource to be created.

* `resource_group` - (Required) The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal. Changing this forces a new resource to be created.

* `server_name` - (Required) The name of the server. Changing this forces a new resource to be created.

* `name` - (Optional) Resource name. Changing this forces a new resource to be created.

* `end_ip_address` - (Optional) The end IP address of the firewall rule. Must be IPv4 format. Must be greater than or equal to startIpAddress. Use value '0.0.0.0' for all Azure-internal IP addresses.

* `start_ip_address` - (Optional) The start IP address of the firewall rule. Must be IPv4 format. Use value '0.0.0.0' for all Azure-internal IP addresses.

* `values` - (Optional) One or more `value` block defined below.

---

The `value` block supports the following:

* `name` - (Optional) Resource name.

* `start_ip_address` - (Optional) The start IP address of the firewall rule. Must be IPv4 format. Use value '0.0.0.0' for all Azure-internal IP addresses.

* `end_ip_address` - (Optional) The end IP address of the firewall rule. Must be IPv4 format. Must be greater than or equal to startIpAddress. Use value '0.0.0.0' for all Azure-internal IP addresses.

## Attributes Reference

The following attributes are exported:

* `id` - Resource ID.

* `type` - Resource type.
