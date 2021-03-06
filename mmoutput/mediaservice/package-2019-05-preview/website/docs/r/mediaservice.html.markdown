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
page_title: "Azure Resource Manager: azurerm_mediaservice"
sidebar_current: "docs-azurerm-resource-mediaservice"
description: |-
  Manage Azure Mediaservice instance.
---

# azurerm_mediaservice

Manage Azure Mediaservice instance.


## Argument Reference

The following arguments are supported:

* `name` - (Required) The Media Services account name. Changing this forces a new resource to be created.

* `resource_group` - (Required) The name of the resource group within the Azure subscription. Changing this forces a new resource to be created.

* `id` - (Optional) The ID of the storage account resource. Changing this forces a new resource to be created.

* `location` - (Optional) The Azure Region of the resource. Changing this forces a new resource to be created.

* `storage_accounts` - (Optional) One or more `storage_account` block defined below.

* `tags` - (Optional) Resource tags. Changing this forces a new resource to be created.

---

The `storage_account` block supports the following:

* `id` - (Optional) The ID of the storage account resource. Media Services relies on tables and queues as well as blobs, so the primary storage account must be a Standard Storage account (either Microsoft.ClassicStorage or Microsoft.Storage). Blob only storage accounts can be added as secondary storage accounts.

* `type` - (Required) The type of the storage account.

## Attributes Reference

The following attributes are exported:

* `media_service_id` - The Media Services account ID.

* `name` - The name of the resource.

* `type` - The type of the resource.
