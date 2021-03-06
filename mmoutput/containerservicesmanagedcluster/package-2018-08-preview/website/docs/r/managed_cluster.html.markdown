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
page_title: "Azure Resource Manager: azurerm_managed_cluster"
sidebar_current: "docs-azurerm-resource-managed-cluster"
description: |-
  Manage Azure ManagedCluster instance.
---

# azurerm_managed_cluster

Manage Azure ManagedCluster instance.


## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the managed cluster resource. Changing this forces a new resource to be created.

* `resource_group` - (Required) The name of the resource group. Changing this forces a new resource to be created.

* `location` - (Required) Resource location Changing this forces a new resource to be created.

* `aad_profile` - (Optional) One `aad_profile` block defined below.

* `addon_profiles` - (Optional) Profile of managed cluster add-on.

* `agent_pool_profiles` - (Optional) One or more `agent_pool_profile` block defined below.

* `dns_prefix` - (Optional) DNS prefix specified when creating the managed cluster.

* `enable_rbac` - (Optional) Whether to enable Kubernetes Role-Based Access Control.

* `kubernetes_version` - (Optional) Version of Kubernetes specified when creating the managed cluster.

* `linux_profile` - (Optional) One `linux_profile` block defined below.

* `network_profile` - (Optional) One `network_profile` block defined below.

* `service_principal_profile` - (Optional) One `service_principal_profile` block defined below.

* `tags` - (Optional) Resource tags Changing this forces a new resource to be created.

---

The `aad_profile` block supports the following:

* `client_app_id` - (Required) The client AAD application ID.

* `server_app_id` - (Required) The server AAD application ID.

* `server_app_secret` - (Optional) The server AAD application secret.

* `tenant_id` - (Optional) The AAD tenant ID to use for authentication. If not specified, will use the tenant of the deployment subscription.

---

The `agent_pool_profile` block supports the following:

* `name` - (Required) Unique name of the agent pool profile in the context of the subscription and resource group.

* `count` - (Required) Number of agents (VMs) to host docker containers. Allowed values must be in the range of 1 to 100 (inclusive). The default value is 1.

* `vm_size` - (Required) Size of agent VMs.

* `os_disk_size_gb` - (Optional) OS Disk Size in GB to be used to specify the disk size for every machine in this master/agent pool. If you specify 0, it will apply the default osDisk size according to the vmSize specified.

* `vnet_subnet_id` - (Optional) VNet SubnetID specifies the VNet's subnet identifier.

* `max_pods` - (Optional) Maximum number of pods that can run on a node.

* `os_type` - (Optional) OsType to be used to specify os type. Choose from Linux and Windows. Default to Linux. Defaults to `Linux`.

* `max_count` - (Optional) Maximum number of nodes for auto-scaling

* `min_count` - (Optional) Minimum number of nodes for auto-scaling

* `enable_auto_scaling` - (Optional) Whether to enable auto-scaler

* `type` - (Optional) AgentPoolType represents types of an agent pool Defaults to `VirtualMachineScaleSets`.

---

The `linux_profile` block supports the following:

* `admin_username` - (Required) The administrator username to use for Linux VMs.

* `ssh` - (Required) One `ssh` block defined below.


---

The `ssh` block supports the following:

* `public_keys` - (Required) One or more `public_key` block defined below.


---

The `public_key` block supports the following:

* `key_data` - (Required) Certificate public key used to authenticate with VMs through SSH. The certificate must be in PEM format with or without headers.

---

The `network_profile` block supports the following:

* `network_plugin` - (Optional) Network plugin used for building Kubernetes network. Defaults to `azure`.

* `network_policy` - (Optional) Network policy used for building Kubernetes network. Defaults to `calico`.

* `pod_cidr` - (Optional) A CIDR notation IP range from which to assign pod IPs when kubenet is used.

* `service_cidr` - (Optional) A CIDR notation IP range from which to assign service cluster IPs. It must not overlap with any Subnet IP ranges.

* `dns_service_ip` - (Optional) An IP address assigned to the Kubernetes DNS service. It must be within the Kubernetes service address range specified in serviceCidr.

* `docker_bridge_cidr` - (Optional) A CIDR notation IP range assigned to the Docker bridge network. It must not overlap with any Subnet IP ranges or the Kubernetes service address range.

---

The `service_principal_profile` block supports the following:

* `client_id` - (Required) The ID for the service principal.

* `secret` - (Optional) The secret password associated with the service principal in plain text.

## Attributes Reference

The following attributes are exported:

* `provisioning_state` - The current deployment or provisioning state, which only appears in the response.

* `fqdn` - FQDN for the master pool.

* `node_resource_group` - Name of the resource group containing agent pool nodes.

* `id` - Resource Id

* `name` - Resource name

* `type` - Resource type
