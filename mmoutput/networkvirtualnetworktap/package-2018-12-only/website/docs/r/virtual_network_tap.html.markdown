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
page_title: "Azure Resource Manager: azurerm_virtual_network_tap"
sidebar_current: "docs-azurerm-resource-virtual-network-tap"
description: |-
  Manage Azure VirtualNetworkTap instance.
---

# azurerm_virtual_network_tap

Manage Azure VirtualNetworkTap instance.


## Argument Reference

The following arguments are supported:

* `resource_group` - (Required) The name of the resource group. Changing this forces a new resource to be created.

* `tap_name` - (Required) The name of the virtual network tap. Changing this forces a new resource to be created.

* `id` - (Optional) Resource ID. Changing this forces a new resource to be created.

* `location` - (Optional) Resource location. Changing this forces a new resource to be created.

* `destination_load_balancer_front_end_ipconfiguration` - (Optional) One `destination_load_balancer_front_end_ipconfiguration` block defined below.

* `destination_network_interface_ipconfiguration` - (Optional) One `destination_network_interface_ipconfiguration` block defined below.

* `destination_port` - (Optional) The VXLAN destination port that will receive the tapped traffic.

* `etag` - (Optional) Gets a unique read-only string that changes whenever the resource is updated. Changing this forces a new resource to be created.

* `tags` - (Optional) Resource tags. Changing this forces a new resource to be created.

---

The `destination_load_balancer_front_end_ipconfiguration` block supports the following:

* `id` - (Optional) Resource ID.

* `private_ip_address` - (Optional) The private IP address of the IP configuration.

* `private_ipallocation_method` - (Optional) The Private IP allocation method. Possible values are: 'Static' and 'Dynamic'. Defaults to `Static`.

* `subnet` - (Optional) One `subnet` block defined below.

* `public_ip_address` - (Optional) One `public_ip_address` block defined below.

* `public_ipprefix` - (Optional) One `public_ipprefix` block defined below.

* `name` - (Optional) The name of the resource that is unique within a resource group. This name can be used to access the resource.

* `etag` - (Optional) A unique read-only string that changes whenever the resource is updated.

* `zones` - (Optional) A list of availability zones denoting the IP allocated for the resource needs to come from.


---

The `subnet` block supports the following:

* `id` - (Optional) Resource ID.

* `name` - (Optional) The name of the resource that is unique within a resource group. This name can be used to access the resource.

* `etag` - (Optional) A unique read-only string that changes whenever the resource is updated.

---

The `public_ip_address` block supports the following:

* `id` - (Optional) Resource ID.

* `location` - (Optional) Resource location. Changing this forces a new resource to be created.

* `tags` - (Optional) Resource tags.

* `etag` - (Optional) A unique read-only string that changes whenever the resource is updated.

* `zones` - (Optional) A list of availability zones denoting the IP allocated for the resource needs to come from.

---

The `public_ipprefix` block supports the following:

* `id` - (Optional) Resource ID.

---

The `destination_network_interface_ipconfiguration` block supports the following:

* `id` - (Optional) Resource ID.

* `virtual_network_taps` - (Optional) One or more `virtual_network_tap` block defined below.

* `application_gateway_backend_address_pools` - (Optional) One or more `application_gateway_backend_address_pool` block defined below.

* `load_balancer_backend_address_pools` - (Optional) One or more `load_balancer_backend_address_pool` block defined below.

* `load_balancer_inbound_nat_rules` - (Optional) One or more `load_balancer_inbound_nat_rule` block defined below.

* `private_ip_address` - (Optional) Private IP address of the IP configuration.

* `private_ipallocation_method` - (Optional) Defines how a private IP address is assigned. Possible values are: 'Static' and 'Dynamic'. Defaults to `Static`.

* `private_ip_address_version` - (Optional) Available from Api-Version 2016-03-30 onwards, it represents whether the specific ipconfiguration is IPv4 or IPv6. Default is taken as IPv4.  Possible values are: 'IPv4' and 'IPv6'. Defaults to `IPv4`.

* `subnet` - (Optional) One `subnet` block defined below.

* `primary` - (Optional) Gets whether this is a primary customer address on the network interface.

* `public_ip_address` - (Optional) One `public_ip_address` block defined below.

* `application_security_groups` - (Optional) One or more `application_security_group` block defined below.

* `name` - (Optional) The name of the resource that is unique within a resource group. This name can be used to access the resource.

* `etag` - (Optional) A unique read-only string that changes whenever the resource is updated.


---

The `virtual_network_tap` block supports the following:

* `id` - (Optional) Resource ID.

* `location` - (Optional) Resource location. Changing this forces a new resource to be created.

* `tags` - (Optional) Resource tags.

* `etag` - (Optional) Gets a unique read-only string that changes whenever the resource is updated.

---

The `application_gateway_backend_address_pool` block supports the following:

* `id` - (Optional) Resource ID.

* `name` - (Optional) Name of the backend address pool that is unique within an Application Gateway.

* `etag` - (Optional) A unique read-only string that changes whenever the resource is updated.

* `type` - (Optional) Type of the resource.

---

The `load_balancer_backend_address_pool` block supports the following:

* `id` - (Optional) Resource ID.

* `name` - (Optional) Gets name of the resource that is unique within a resource group. This name can be used to access the resource.

* `etag` - (Optional) A unique read-only string that changes whenever the resource is updated.

---

The `load_balancer_inbound_nat_rule` block supports the following:

* `id` - (Optional) Resource ID.

* `name` - (Optional) Gets name of the resource that is unique within a resource group. This name can be used to access the resource.

* `etag` - (Optional) A unique read-only string that changes whenever the resource is updated.

---

The `subnet` block supports the following:

* `id` - (Optional) Resource ID.

* `name` - (Optional) The name of the resource that is unique within a resource group. This name can be used to access the resource.

* `etag` - (Optional) A unique read-only string that changes whenever the resource is updated.

---

The `public_ip_address` block supports the following:

* `id` - (Optional) Resource ID.

* `location` - (Optional) Resource location. Changing this forces a new resource to be created.

* `tags` - (Optional) Resource tags.

* `etag` - (Optional) A unique read-only string that changes whenever the resource is updated.

* `zones` - (Optional) A list of availability zones denoting the IP allocated for the resource needs to come from.

---

The `application_security_group` block supports the following:

* `id` - (Optional) Resource ID.

* `location` - (Optional) Resource location. Changing this forces a new resource to be created.

* `tags` - (Optional) Resource tags.

## Attributes Reference

The following attributes are exported:

* `network_interface_tap_configurations` - One or more `network_interface_tap_configuration` block defined below.

* `resource_guid` - The resourceGuid property of the virtual network tap.

* `provisioning_state` - The provisioning state of the virtual network tap. Possible values are: 'Updating', 'Deleting', and 'Failed'.

* `name` - Resource name.

* `type` - Resource type.


---

The `network_interface_tap_configuration` block contains the following:

* `id` - (Optional) Resource ID.

* `virtual_network_tap` - (Optional) One `virtual_network_tap` block defined below.

* `provisioning_state` - (Optional) The provisioning state of the network interface tap configuration. Possible values are: 'Updating', 'Deleting', and 'Failed'.

* `name` - (Optional) The name of the resource that is unique within a resource group. This name can be used to access the resource.

* `etag` - (Optional) A unique read-only string that changes whenever the resource is updated.

* `type` - (Optional) Sub Resource type.


---

The `virtual_network_tap` block supports the following:

* `id` - (Optional) Resource ID.

* `name` - (Optional) Resource name.

* `type` - (Optional) Resource type.

* `location` - (Optional) Resource location. Changing this forces a new resource to be created.

* `tags` - (Optional) Resource tags.

* `etag` - (Optional) Gets a unique read-only string that changes whenever the resource is updated.
