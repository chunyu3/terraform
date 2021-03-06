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
page_title: "Azure Resource Manager: azurerm_volume_container"
sidebar_current: "docs-azurerm-resource-volume-container"
description: |-
  Manage Azure VolumeContainer instance.
---

# azurerm_volume_container

Manage Azure VolumeContainer instance.


## Argument Reference

The following arguments are supported:

* `name` - (Required) The volume container name. Changing this forces a new resource to be created.

* `resource_group` - (Required) The resource group name Changing this forces a new resource to be created.

* `device_name` - (Required) The device name Changing this forces a new resource to be created.

* `manager_name` - (Required) The manager name Changing this forces a new resource to be created.

* `storage_account_credential_id` - (Required) The path ID of storage account associated with the volume container.

* `band_width_rate_in_mbps` - (Optional) The bandwidth-rate set on the volume container.

* `bandwidth_setting_id` - (Optional) The ID of the bandwidth setting associated with the volume container.

* `encryption_key` - (Optional) One `encryption_key` block defined below.

* `kind` - (Optional) The Kind of the object. Currently only Series8000 is supported Defaults to `Series8000`. Changing this forces a new resource to be created.

---

The `encryption_key` block supports the following:

* `value` - (Required) The value of the secret.

* `encryption_cert_thumbprint` - (Optional) Thumbprint certificate that was used to encrypt "Value". If the value in unencrypted, it will be null.

* `encryption_algorithm` - (Required) The algorithm used to encrypt "Value".

## Attributes Reference

The following attributes are exported:

* `encryption_status` - The flag to denote whether encryption is enabled or not.

* `volume_count` - The number of volumes in the volume Container.

* `owner_ship_status` - The owner ship status of the volume container. Only when the status is "NotOwned", the delete operation on the volume container is permitted.

* `total_cloud_storage_usage_in_bytes` - The total cloud storage for the volume container.

* `id` - The path ID that uniquely identifies the object.

* `name` - The name of the object.

* `type` - The hierarchical type of the object.
