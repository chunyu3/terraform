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
page_title: "Azure Resource Manager: azurerm_virtual_network_gateway"
sidebar_current: "docs-azurerm-resource-virtual-network-gateway"
description: |-
  Manage Azure VirtualNetworkGateway instance.
---

# azurerm_virtual_network_gateway

Manage Azure VirtualNetworkGateway instance.


## Argument Reference

The following arguments are supported:

* `resource_group` - (Required) The name of the resource group. Changing this forces a new resource to be created.

* `virtual_network_gateway_name` - (Required) The name of the virtual network gateway. Changing this forces a new resource to be created.

* `id` - (Optional) Resource ID. Changing this forces a new resource to be created.

* `location` - (Optional) Resource location. Changing this forces a new resource to be created.

* `active_active` - (Optional) ActiveActive flag

* `bgp_settings` - (Optional) One `bgp_setting` block defined below.

* `enable_bgp` - (Optional) Whether BGP is enabled for this virtual network gateway or not.

* `etag` - (Optional) Gets a unique read-only string that changes whenever the resource is updated. Changing this forces a new resource to be created.

* `gateway_default_site` - (Optional) One `gateway_default_site` block defined below.

* `gateway_type` - (Optional) The type of this virtual network gateway. Possible values are: 'Vpn' and 'ExpressRoute'. Defaults to `Vpn`.

* `ip_configurations` - (Optional) One or more `ip_configuration` block defined below.

* `resource_guid` - (Optional) The resource GUID property of the VirtualNetworkGateway resource.

* `sku` - (Optional) One `sku` block defined below.

* `vpn_client_configuration` - (Optional) One `vpn_client_configuration` block defined below.

* `vpn_type` - (Optional) The type of this virtual network gateway. Possible values are: 'PolicyBased' and 'RouteBased'. Defaults to `PolicyBased`.

* `tags` - (Optional) Resource tags. Changing this forces a new resource to be created.

---

The `bgp_setting` block supports the following:

* `asn` - (Optional) The BGP speaker's ASN.

* `bgp_peering_address` - (Optional) The BGP peering address and BGP identifier of this BGP speaker.

* `peer_weight` - (Optional) The weight added to routes learned from this BGP speaker.

---

The `gateway_default_site` block supports the following:

* `id` - (Optional) Resource ID.

---

The `ip_configuration` block supports the following:

* `id` - (Optional) Resource ID.

* `private_ipallocation_method` - (Optional) The private IP allocation method. Possible values are: 'Static' and 'Dynamic'. Defaults to `Static`.

* `subnet` - (Optional) One `subnet` block defined below.

* `public_ip_address` - (Optional) One `public_ip_address` block defined below.

* `name` - (Optional) The name of the resource that is unique within a resource group. This name can be used to access the resource.

* `etag` - (Optional) A unique read-only string that changes whenever the resource is updated.


---

The `subnet` block supports the following:

* `id` - (Optional) Resource ID.

---

The `public_ip_address` block supports the following:

* `id` - (Optional) Resource ID.

---

The `sku` block supports the following:

* `name` - (Optional) Gateway SKU name. Defaults to `Basic`.

* `tier` - (Optional) Gateway SKU tier. Defaults to `Basic`.

* `capacity` - (Optional) The capacity.

---

The `vpn_client_configuration` block supports the following:

* `vpn_client_address_pool` - (Optional) One `vpn_client_address_pool` block defined below.

* `vpn_client_root_certificates` - (Optional) One or more `vpn_client_root_certificate` block defined below.

* `vpn_client_revoked_certificates` - (Optional) One or more `vpn_client_revoked_certificate` block defined below.

* `vpn_client_protocols` - (Optional) VpnClientProtocols for Virtual network gateway.

* `vpn_client_ipsec_policies` - (Optional) One or more `vpn_client_ipsec_policy` block defined below.

* `radius_server_address` - (Optional) The radius server address property of the VirtualNetworkGateway resource for vpn client connection.

* `radius_server_secret` - (Optional) The radius secret property of the VirtualNetworkGateway resource for vpn client connection.


---

The `vpn_client_address_pool` block supports the following:

* `address_prefixes` - (Optional) A list of address blocks reserved for this virtual network in CIDR notation.

---

The `vpn_client_root_certificate` block supports the following:

* `id` - (Optional) Resource ID.

* `public_cert_data` - (Required) The certificate public data.

* `name` - (Optional) The name of the resource that is unique within a resource group. This name can be used to access the resource.

* `etag` - (Optional) A unique read-only string that changes whenever the resource is updated.

---

The `vpn_client_revoked_certificate` block supports the following:

* `id` - (Optional) Resource ID.

* `thumbprint` - (Optional) The revoked VPN client certificate thumbprint.

* `name` - (Optional) The name of the resource that is unique within a resource group. This name can be used to access the resource.

* `etag` - (Optional) A unique read-only string that changes whenever the resource is updated.

---

The `vpn_client_ipsec_policy` block supports the following:

* `sa_life_time_seconds` - (Required) The IPSec Security Association (also called Quick Mode or Phase 2 SA) lifetime in seconds for a site to site VPN tunnel.

* `sa_data_size_kilobytes` - (Required) The IPSec Security Association (also called Quick Mode or Phase 2 SA) payload size in KB for a site to site VPN tunnel.

* `ipsec_encryption` - (Required) The IPSec encryption algorithm (IKE phase 1).

* `ipsec_integrity` - (Required) The IPSec integrity algorithm (IKE phase 1).

* `ike_encryption` - (Required) The IKE encryption algorithm (IKE phase 2).

* `ike_integrity` - (Required) The IKE integrity algorithm (IKE phase 2).

* `dh_group` - (Required) The DH Groups used in IKE Phase 1 for initial SA.

* `pfs_group` - (Required) The Pfs Groups used in IKE Phase 2 for new child SA.

## Attributes Reference

The following attributes are exported:

* `provisioning_state` - The provisioning state of the VirtualNetworkGateway resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.

* `name` - Resource name.

* `type` - Resource type.
