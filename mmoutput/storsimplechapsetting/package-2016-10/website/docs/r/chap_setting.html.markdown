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
page_title: "Azure Resource Manager: azurerm_chap_setting"
sidebar_current: "docs-azurerm-resource-chap-setting"
description: |-
  Manage Azure ChapSetting instance.
---

# azurerm_chap_setting

Manage Azure ChapSetting instance.


## Argument Reference

The following arguments are supported:

* `name` - (Required) The chap user name. Changing this forces a new resource to be created.

* `resource_group` - (Required) The resource group name Changing this forces a new resource to be created.

* `device_name` - (Required) The device name. Changing this forces a new resource to be created.

* `manager_name` - (Required) The manager name Changing this forces a new resource to be created.

* `password` - (Required) One `password` block defined below.

---

The `password` block supports the following:

* `value` - (Required) The value of the secret itself. If the secret is in plaintext then EncryptionAlgorithm will be none and EncryptionCertThumbprint will be null.

* `encryption_certificate_thumbprint` - (Optional) Thumbprint certificate that was used to encrypt "Value"

* `encryption_algorithm` - (Required) Algorithm used to encrypt "Value"

## Attributes Reference

The following attributes are exported:

* `id` - The identifier.

* `name` - The name.

* `type` - The type.
