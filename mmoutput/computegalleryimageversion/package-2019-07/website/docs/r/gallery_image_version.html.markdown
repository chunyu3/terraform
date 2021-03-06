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
page_title: "Azure Resource Manager: azurerm_gallery_image_version"
sidebar_current: "docs-azurerm-resource-gallery-image-version"
description: |-
  Manage Azure GalleryImageVersion instance.
---

# azurerm_gallery_image_version

Manage Azure GalleryImageVersion instance.


## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the gallery Image Version to be deleted. Changing this forces a new resource to be created.

* `resource_group` - (Required) The name of the resource group. Changing this forces a new resource to be created.

* `location` - (Required) Resource location Changing this forces a new resource to be created.

* `gallery_image_name` - (Required) The name of the Shared Image Gallery Image Definition from which the Image Versions are to be listed. Changing this forces a new resource to be created.

* `gallery_name` - (Required) The name of the Shared Image Gallery in which the Image Definition resides. Changing this forces a new resource to be created.

* `storage_profile` - (Required) One `storage_profile` block defined below.

---

The `storage_profile` block supports the following:

* `source` - (Optional) One `source` block defined below.

* `os_disk_image` - (Optional) One `os_disk_image` block defined below.

* `data_disk_images` - (Optional) One or more `data_disk_image` block defined below.


---

The `source` block supports the following:

* `id` - (Required) The id of the gallery artifact version source. Can specify a disk uri, snapshot uri, or user image.

---

The `os_disk_image` block supports the following:

* `host_caching` - (Optional) The host caching of the disk. Valid values are 'None', 'ReadOnly', and 'ReadWrite' Defaults to `None`.

* `source` - (Optional) One `source` block defined below.


---

The `source` block supports the following:

* `id` - (Required) The id of the gallery artifact version source. Can specify a disk uri, snapshot uri, or user image.

---

The `data_disk_image` block supports the following:

* `host_caching` - (Optional) The host caching of the disk. Valid values are 'None', 'ReadOnly', and 'ReadWrite' Defaults to `None`.

* `source` - (Optional) One `source` block defined below.

* `lun` - (Required) This property specifies the logical unit number of the data disk. This value is used to identify data disks within the Virtual Machine and therefore must be unique for each data disk attached to the Virtual Machine.


---

The `source` block supports the following:

* `id` - (Required) The id of the gallery artifact version source. Can specify a disk uri, snapshot uri, or user image.

* `publishing_profile` - (Optional) One `publishing_profile` block defined below.

* `tags` - (Optional) Resource tags Changing this forces a new resource to be created.

---

The `publishing_profile` block supports the following:

* `target_regions` - (Optional) One or more `target_region` block defined below.

* `replica_count` - (Optional) The number of replicas of the Image Version to be created per region. This property would take effect for a region when regionalReplicaCount is not specified. This property is updatable.

* `exclude_from_latest` - (Optional) If set to true, Virtual Machines deployed from the latest version of the Image Definition won't use this Image Version.

* `end_of_life_date` - (Optional) The end of life date of the gallery Image Version. This property can be used for decommissioning purposes. This property is updatable.

* `storage_account_type` - (Optional) Specifies the storage account type to be used to store the image. This property is not updatable. Defaults to `Standard_LRS`.


---

The `target_region` block supports the following:

* `name` - (Required) The name of the region.

* `regional_replica_count` - (Optional) The number of replicas of the Image Version to be created per region. This property is updatable.

* `storage_account_type` - (Optional) Specifies the storage account type to be used to store the image. This property is not updatable. Defaults to `Standard_LRS`.

## Attributes Reference

The following attributes are exported:

* `provisioning_state` - The provisioning state, which only appears in the response.

* `replication_status` - One `replication_status` block defined below.

* `id` - Resource Id

* `name` - Resource name

* `type` - Resource type


---

The `replication_status` block contains the following:

* `aggregated_state` - This is the aggregated replication status based on all the regional replication status flags.

* `summary` - One or more `summary` block defined below.


---

The `summary` block contains the following:

* `region` - The region to which the gallery Image Version is being replicated to.

* `state` - This is the regional replication state.

* `details` - The details of the replication status.

* `progress` - It indicates progress of the replication job.
