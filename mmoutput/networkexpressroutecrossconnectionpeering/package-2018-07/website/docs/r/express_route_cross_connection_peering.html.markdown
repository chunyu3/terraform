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
page_title: "Azure Resource Manager: azurerm_express_route_cross_connection_peering"
sidebar_current: "docs-azurerm-resource-express-route-cross-connection-peering"
description: |-
  Manage Azure ExpressRouteCrossConnectionPeering instance.
---

# azurerm_express_route_cross_connection_peering

Manage Azure ExpressRouteCrossConnectionPeering instance.


## Argument Reference

The following arguments are supported:

* `resource_group` - (Required) The name of the resource group. Changing this forces a new resource to be created.

* `cross_connection_name` - (Required) The name of the ExpressRouteCrossConnection. Changing this forces a new resource to be created.

* `peering_name` - (Required) The name of the peering. Changing this forces a new resource to be created.

* `id` - (Optional) Resource ID. Changing this forces a new resource to be created.

* `name` - (Optional) Gets name of the resource that is unique within a resource group. This name can be used to access the resource. Changing this forces a new resource to be created.

* `gateway_manager_etag` - (Optional) The GatewayManager Etag.

* `ipv6peering_config` - (Optional) One `ipv6peering_config` block defined below.

* `last_modified_by` - (Optional) Gets whether the provider or the customer last modified the peering.

* `microsoft_peering_config` - (Optional) One `microsoft_peering_config` block defined below.

* `peer_asn` - (Optional) The peer ASN.

* `peering_type` - (Optional) The peering type. Defaults to `AzurePublicPeering`.

* `primary_peer_address_prefix` - (Optional) The primary address prefix.

* `secondary_peer_address_prefix` - (Optional) The secondary address prefix.

* `shared_key` - (Optional) The shared key.

* `state` - (Optional) The peering state. Defaults to `Disabled`.

* `vlan_id` - (Optional) The VLAN ID.

---

The `ipv6peering_config` block supports the following:

* `primary_peer_address_prefix` - (Optional) The primary address prefix.

* `secondary_peer_address_prefix` - (Optional) The secondary address prefix.

* `microsoft_peering_config` - (Optional) One `microsoft_peering_config` block defined below.

* `route_filter` - (Optional) One `route_filter` block defined below.

* `state` - (Optional) The state of peering. Possible values are: 'Disabled' and 'Enabled' Defaults to `Disabled`.


---

The `microsoft_peering_config` block supports the following:

* `advertised_public_prefixes` - (Optional) The reference of AdvertisedPublicPrefixes.

* `advertised_communities` - (Optional) The communities of bgp peering. Specified for microsoft peering

* `advertised_public_prefixes_state` - (Optional) AdvertisedPublicPrefixState of the Peering resource. Possible values are 'NotConfigured', 'Configuring', 'Configured', and 'ValidationNeeded'. Defaults to `NotConfigured`.

* `legacy_mode` - (Optional) The legacy mode of the peering.

* `customer_asn` - (Optional) The CustomerASN of the peering.

* `routing_registry_name` - (Optional) The RoutingRegistryName of the configuration.

---

The `route_filter` block supports the following:

* `id` - (Optional) Resource ID.

* `location` - (Optional) Resource location. Changing this forces a new resource to be created.

* `tags` - (Optional) Resource tags.

---

The `microsoft_peering_config` block supports the following:

* `advertised_public_prefixes` - (Optional) The reference of AdvertisedPublicPrefixes.

* `advertised_communities` - (Optional) The communities of bgp peering. Specified for microsoft peering

* `advertised_public_prefixes_state` - (Optional) AdvertisedPublicPrefixState of the Peering resource. Possible values are 'NotConfigured', 'Configuring', 'Configured', and 'ValidationNeeded'. Defaults to `NotConfigured`.

* `legacy_mode` - (Optional) The legacy mode of the peering.

* `customer_asn` - (Optional) The CustomerASN of the peering.

* `routing_registry_name` - (Optional) The RoutingRegistryName of the configuration.

## Attributes Reference

The following attributes are exported:

* `azure_asn` - The Azure ASN.

* `primary_azure_port` - The primary port.

* `secondary_azure_port` - The secondary port.

* `provisioning_state` - Gets the provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.

* `etag` - A unique read-only string that changes whenever the resource is updated.
