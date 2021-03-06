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
page_title: "Azure Resource Manager: azurerm_local_network_gateway"
sidebar_current: "docs-azurerm-resource-local-network-gateway"
description: |-
  Manage Azure LocalNetworkGateway instance.
---

# azurerm_local_network_gateway

Manage Azure LocalNetworkGateway instance.


## Argument Reference

The following arguments are supported:

* `resource_group` - (Required) The name of the resource group. Changing this forces a new resource to be created.

* `local_network_gateway_name` - (Required) The name of the local network gateway. Changing this forces a new resource to be created.

* `id` - (Optional) Resource Id Changing this forces a new resource to be created.

* `location` - (Optional) Resource location Changing this forces a new resource to be created.

* `bgp_settings` - (Optional) One `bgp_setting` block defined below.

* `etag` - (Optional) Gets a unique read-only string that changes whenever the resource is updated Changing this forces a new resource to be created.

* `gateway_ip_address` - (Optional) IP address of local network gateway.

* `local_network_address_space` - (Optional) One `local_network_address_space` block defined below.

* `resource_guid` - (Optional) Gets or sets resource guid property of the LocalNetworkGateway resource

* `tags` - (Optional) Resource tags Changing this forces a new resource to be created.

---

The `bgp_setting` block supports the following:

* `asn` - (Optional) Gets or sets this BGP speaker's ASN

* `bgp_peering_address` - (Optional) Gets or sets the BGP peering address and BGP identifier of this BGP speaker

* `peer_weight` - (Optional) Gets or sets the weight added to routes learned from this BGP speaker

---

The `local_network_address_space` block supports the following:

* `address_prefixes` - (Optional) Gets or sets list of address blocks reserved for this virtual network in CIDR notation

## Attributes Reference

The following attributes are exported:

* `provisioning_state` - Gets provisioning state of the LocalNetworkGateway resource Updating/Deleting/Failed

* `name` - Resource name

* `type` - Resource type
