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
page_title: "Azure Resource Manager: azurerm_load_balancer"
sidebar_current: "docs-azurerm-resource-load-balancer"
description: |-
  Manage Azure LoadBalancer instance.
---

# azurerm_load_balancer

Manage Azure LoadBalancer instance.


## Argument Reference

The following arguments are supported:

* `resource_group` - (Required) The name of the resource group. Changing this forces a new resource to be created.

* `load_balancer_name` - (Required) The name of the load balancer. Changing this forces a new resource to be created.

* `id` - (Optional) Resource ID. Changing this forces a new resource to be created.

* `location` - (Optional) Resource location. Changing this forces a new resource to be created.

* `backend_address_pools` - (Optional) One or more `backend_address_pool` block defined below.

* `etag` - (Optional) A unique read-only string that changes whenever the resource is updated. Changing this forces a new resource to be created.

* `frontend_ipconfigurations` - (Optional) One or more `frontend_ipconfiguration` block defined below.

* `inbound_nat_pools` - (Optional) One or more `inbound_nat_pool` block defined below.

* `inbound_nat_rules` - (Optional) One or more `inbound_nat_rule` block defined below.

* `load_balancing_rules` - (Optional) One or more `load_balancing_rule` block defined below.

* `outbound_nat_rules` - (Optional) One or more `outbound_nat_rule` block defined below.

* `probes` - (Optional) One or more `probe` block defined below.

* `resource_guid` - (Optional) The resource GUID property of the load balancer resource.

* `tags` - (Optional) Resource tags. Changing this forces a new resource to be created.

---

The `backend_address_pool` block supports the following:

* `id` - (Optional) Resource ID.

* `name` - (Optional) Gets name of the resource that is unique within a resource group. This name can be used to access the resource.

* `etag` - (Optional) A unique read-only string that changes whenever the resource is updated.

---

The `frontend_ipconfiguration` block supports the following:

* `id` - (Optional) Resource ID.

* `private_ip_address` - (Optional) The private IP address of the IP configuration.

* `private_ipallocation_method` - (Optional) The Private IP allocation method. Possible values are: 'Static' and 'Dynamic'. Defaults to `Static`.

* `subnet` - (Optional) One `subnet` block defined below.

* `public_ip_address` - (Optional) One `public_ip_address` block defined below.

* `name` - (Optional) The name of the resource that is unique within a resource group. This name can be used to access the resource.

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

---

The `inbound_nat_pool` block supports the following:

* `id` - (Optional) Resource ID.

* `frontend_ipconfiguration` - (Optional) One `frontend_ipconfiguration` block defined below.

* `protocol` - (Required) The transport protocol for the endpoint. Possible values are: 'Udp' or 'Tcp'.

* `frontend_port_range_start` - (Required) The first port number in the range of external ports that will be used to provide Inbound Nat to NICs associated with a load balancer. Acceptable values range between 1 and 65534.

* `frontend_port_range_end` - (Required) The last port number in the range of external ports that will be used to provide Inbound Nat to NICs associated with a load balancer. Acceptable values range between 1 and 65535.

* `backend_port` - (Required) The port used for internal connections on the endpoint. Acceptable values are between 1 and 65535.

* `name` - (Optional) The name of the resource that is unique within a resource group. This name can be used to access the resource.

* `etag` - (Optional) A unique read-only string that changes whenever the resource is updated.


---

The `frontend_ipconfiguration` block supports the following:

* `id` - (Optional) Resource ID.

---

The `inbound_nat_rule` block supports the following:

* `id` - (Optional) Resource ID.

* `frontend_ipconfiguration` - (Optional) One `frontend_ipconfiguration` block defined below.

* `protocol` - (Optional) The transport protocol for the endpoint. Possible values are: 'Udp' or 'Tcp' Defaults to `Udp`.

* `frontend_port` - (Optional) The port for the external endpoint. Port numbers for each rule must be unique within the Load Balancer. Acceptable values range from 1 to 65534.

* `backend_port` - (Optional) The port used for the internal endpoint. Acceptable values range from 1 to 65535.

* `idle_timeout_in_minutes` - (Optional) The timeout for the TCP idle connection. The value can be set between 4 and 30 minutes. The default value is 4 minutes. This element is only used when the protocol is set to TCP.

* `enable_floating_ip` - (Optional) Configures a virtual machine's endpoint for the floating IP capability required to configure a SQL AlwaysOn Availability Group. This setting is required when using the SQL AlwaysOn Availability Groups in SQL server. This setting can't be changed after you create the endpoint.

* `name` - (Optional) Gets name of the resource that is unique within a resource group. This name can be used to access the resource.

* `etag` - (Optional) A unique read-only string that changes whenever the resource is updated.


---

The `frontend_ipconfiguration` block supports the following:

* `id` - (Optional) Resource ID.

---

The `load_balancing_rule` block supports the following:

* `id` - (Optional) Resource ID.

* `frontend_ipconfiguration` - (Optional) One `frontend_ipconfiguration` block defined below.

* `backend_address_pool` - (Optional) One `backend_address_pool` block defined below.

* `probe` - (Optional) One `probe` block defined below.

* `protocol` - (Required) The transport protocol for the external endpoint. Possible values are 'Udp' or 'Tcp'

* `load_distribution` - (Optional) The load distribution policy for this rule. Possible values are 'Default', 'SourceIP', and 'SourceIPProtocol'. Defaults to `Default`.

* `frontend_port` - (Required) The port for the external endpoint. Port numbers for each rule must be unique within the Load Balancer. Acceptable values are between 1 and 65534.

* `backend_port` - (Optional) The port used for internal connections on the endpoint. Acceptable values are between 1 and 65535.

* `idle_timeout_in_minutes` - (Optional) The timeout for the TCP idle connection. The value can be set between 4 and 30 minutes. The default value is 4 minutes. This element is only used when the protocol is set to TCP.

* `enable_floating_ip` - (Optional) Configures a virtual machine's endpoint for the floating IP capability required to configure a SQL AlwaysOn Availability Group. This setting is required when using the SQL AlwaysOn Availability Groups in SQL server. This setting can't be changed after you create the endpoint.

* `name` - (Optional) The name of the resource that is unique within a resource group. This name can be used to access the resource.

* `etag` - (Optional) A unique read-only string that changes whenever the resource is updated.


---

The `frontend_ipconfiguration` block supports the following:

* `id` - (Optional) Resource ID.

---

The `backend_address_pool` block supports the following:

* `id` - (Optional) Resource ID.

---

The `probe` block supports the following:

* `id` - (Optional) Resource ID.

---

The `outbound_nat_rule` block supports the following:

* `id` - (Optional) Resource ID.

* `allocated_outbound_ports` - (Optional) The number of outbound ports to be used for NAT.

* `frontend_ipconfigurations` - (Optional) One or more `frontend_ipconfiguration` block defined below.

* `backend_address_pool` - (Required) One `backend_address_pool` block defined below.

* `name` - (Optional) The name of the resource that is unique within a resource group. This name can be used to access the resource.

* `etag` - (Optional) A unique read-only string that changes whenever the resource is updated.


---

The `frontend_ipconfiguration` block supports the following:

* `id` - (Optional) Resource ID.

---

The `backend_address_pool` block supports the following:

* `id` - (Optional) Resource ID.

---

The `probe` block supports the following:

* `id` - (Optional) Resource ID.

* `protocol` - (Required) The protocol of the end point. Possible values are: 'Http' or 'Tcp'. If 'Tcp' is specified, a received ACK is required for the probe to be successful. If 'Http' is specified, a 200 OK response from the specifies URI is required for the probe to be successful.

* `port` - (Required) The port for communicating the probe. Possible values range from 1 to 65535, inclusive.

* `interval_in_seconds` - (Optional) The interval, in seconds, for how frequently to probe the endpoint for health status. Typically, the interval is slightly less than half the allocated timeout period (in seconds) which allows two full probes before taking the instance out of rotation. The default value is 15, the minimum value is 5.

* `number_of_probes` - (Optional) The number of probes where if no response, will result in stopping further traffic from being delivered to the endpoint. This values allows endpoints to be taken out of rotation faster or slower than the typical times used in Azure.

* `request_path` - (Optional) The URI used for requesting health status from the VM. Path is required if a protocol is set to http. Otherwise, it is not allowed. There is no default value.

* `name` - (Optional) Gets name of the resource that is unique within a resource group. This name can be used to access the resource.

* `etag` - (Optional) A unique read-only string that changes whenever the resource is updated.

## Attributes Reference

The following attributes are exported:

* `provisioning_state` - Gets the provisioning state of the PublicIP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.

* `name` - Resource name.

* `type` - Resource type.
