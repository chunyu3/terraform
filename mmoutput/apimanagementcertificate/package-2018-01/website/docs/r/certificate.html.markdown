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
page_title: "Azure Resource Manager: azurerm_certificate"
sidebar_current: "docs-azurerm-resource-certificate"
description: |-
  Manage Azure Certificate instance.
---

# azurerm_certificate

Manage Azure Certificate instance.


## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the API Management service. Changing this forces a new resource to be created.

* `resource_group` - (Required) The name of the resource group. Changing this forces a new resource to be created.

* `certificate_id` - (Required) Identifier of the certificate entity. Must be unique in the current API Management service instance. Changing this forces a new resource to be created.

* `data` - (Required) Base 64 encoded certificate using the application/x-pkcs12 representation.

* `password` - (Required) Password for the Certificate

## Attributes Reference

The following attributes are exported:

* `subject` - Subject attribute of the certificate.

* `thumbprint` - Thumbprint of the certificate.

* `expiration_date` - Expiration date of the certificate. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.<br>

* `id` - Resource ID.

* `name` - Resource name.

* `type` - Resource type for API Management resource.
