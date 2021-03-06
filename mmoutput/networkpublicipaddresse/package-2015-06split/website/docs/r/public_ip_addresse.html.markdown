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
page_title: "Azure Resource Manager: azurerm_public_ip_addresse"
sidebar_current: "docs-azurerm-resource-public-ip-addresse"
description: |-
  Manage Azure PublicIPAddresse instance.
---

# azurerm_public_ip_addresse

Manage Azure PublicIPAddresse instance.


## Argument Reference

The following arguments are supported:

* `resource_group` - (Required) The name of the resource group. Changing this forces a new resource to be created.

* `public_ip_address_name` - (Required) The name of the public IP address. Changing this forces a new resource to be created.

* `id` - (Optional) Resource Identifier. Changing this forces a new resource to be created.

* `location` - (Optional) Resource location. Changing this forces a new resource to be created.

* `dns_settings` - (Optional) One `dns_setting` block defined below.

* `etag` - (Optional) A unique read-only string that changes whenever the resource is updated. Changing this forces a new resource to be created.

* `idle_timeout_in_minutes` - (Optional) The idle timeout of the public IP address.

* `ip_address` - (Optional) 

* `ip_configuration` - (Optional) One `ip_configuration` block defined below.

* `public_ipallocation_method` - (Optional) The public IP allocation method. Possible values are: 'Static' and 'Dynamic'. Defaults to `Static`.

* `resource_guid` - (Optional) The resource GUID property of the public IP resource.

* `tags` - (Optional) Resource tags. Changing this forces a new resource to be created.

---

The `dns_setting` block supports the following:

* `domain_name_label` - (Optional) Gets or sets the Domain name label.The concatenation of the domain name label and the regionalized DNS zone make up the fully qualified domain name associated with the public IP address. If a domain name label is specified, an A DNS record is created for the public IP in the Microsoft Azure DNS system.

* `fqdn` - (Optional) Gets the FQDN, Fully qualified domain name of the A DNS record associated with the public IP. This is the concatenation of the domainNameLabel and the regionalized DNS zone.

* `reverse_fqdn` - (Optional) Gets or Sets the Reverse FQDN. A user-visible, fully qualified domain name that resolves to this public IP address. If the reverseFqdn is specified, then a PTR DNS record is created pointing from the IP address in the in-addr.arpa domain to the reverse FQDN.

---

The `ip_configuration` block supports the following:

* `id` - (Optional) Resource Identifier.

* `private_ip_address` - (Optional) The private IP address of the IP configuration.

* `private_ipallocation_method` - (Optional) The private IP allocation method. Possible values are 'Static' and 'Dynamic'. Defaults to `Static`.

* `subnet` - (Optional) One `subnet` block defined below.

* `public_ip_address` - (Optional) One `public_ip_address` block defined below.

* `name` - (Optional) The name of the resource that is unique within a resource group. This name can be used to access the resource.

* `etag` - (Optional) A unique read-only string that changes whenever the resource is updated.


---

The `subnet` block supports the following:

* `id` - (Optional) Resource Identifier.

* `name` - (Optional) The name of the resource that is unique within a resource group. This name can be used to access the resource.

* `etag` - (Optional) A unique read-only string that changes whenever the resource is updated.

---

The `public_ip_address` block supports the following:

* `id` - (Optional) Resource Identifier.

* `location` - (Optional) Resource location. Changing this forces a new resource to be created.

* `tags` - (Optional) Resource tags.

* `etag` - (Optional) A unique read-only string that changes whenever the resource is updated.

## Attributes Reference

The following attributes are exported:

* `provisioning_state` - The provisioning state of the PublicIP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.

* `name` - Resource name.

* `type` - Resource type.
