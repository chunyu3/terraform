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
page_title: "Azure Resource Manager: azurerm_security_rule"
sidebar_current: "docs-azurerm-resource-security-rule"
description: |-
  Manage Azure SecurityRule instance.
---

# azurerm_security_rule

Manage Azure SecurityRule instance.


## Argument Reference

The following arguments are supported:

* `resource_group` - (Required) The name of the resource group. Changing this forces a new resource to be created.

* `access` - (Required) Gets or sets network traffic is allowed or denied. Possible values are 'Allow' and 'Deny'

* `destination_address_prefix` - (Required) Gets or sets destination address prefix. CIDR or source IP range. Asterisk '*' can also be used to match all source IPs. Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used.

* `direction` - (Required) Gets or sets the direction of the rule.InBound or Outbound. The direction specifies if rule will be evaluated on incoming or outgoing traffic.

* `network_security_group_name` - (Required) The name of the network security group. Changing this forces a new resource to be created.

* `protocol` - (Required) Gets or sets Network protocol this rule applies to. Can be Tcp, Udp or All(*).

* `security_rule_name` - (Required) The name of the security rule. Changing this forces a new resource to be created.

* `source_address_prefix` - (Required) Gets or sets source address prefix. CIDR or source IP range. Asterisk '*' can also be used to match all source IPs. Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used. If this is an ingress rule, specifies where network traffic originates from.

* `id` - (Optional) Resource Id Changing this forces a new resource to be created.

* `name` - (Optional) Gets name of the resource that is unique within a resource group. This name can be used to access the resource Changing this forces a new resource to be created.

* `description` - (Optional) Gets or sets a description for this rule. Restricted to 140 chars.

* `destination_port_range` - (Optional) Gets or sets Destination Port or Range. Integer or range between 0 and 65535. Asterisk '*' can also be used to match all ports.

* `etag` - (Optional) A unique read-only string that changes whenever the resource is updated Changing this forces a new resource to be created.

* `priority` - (Optional) Gets or sets the priority of the rule. The value can be between 100 and 4096. The priority number must be unique for each rule in the collection. The lower the priority number, the higher the priority of the rule.

* `source_port_range` - (Optional) Gets or sets Source Port or Range. Integer or range between 0 and 65535. Asterisk '*' can also be used to match all ports.

## Attributes Reference

The following attributes are exported:

* `provisioning_state` - Gets or sets Provisioning state of the PublicIP resource Updating/Deleting/Failed
