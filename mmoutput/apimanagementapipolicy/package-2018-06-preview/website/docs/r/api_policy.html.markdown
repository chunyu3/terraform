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
page_title: "Azure Resource Manager: azurerm_api_policy"
sidebar_current: "docs-azurerm-resource-api-policy"
description: |-
  Manage Azure ApiPolicy instance.
---

# azurerm_api_policy

Manage Azure ApiPolicy instance.


## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the API Management service. Changing this forces a new resource to be created.

* `resource_group` - (Required) The name of the resource group. Changing this forces a new resource to be created.

* `api_id` - (Required) API revision identifier. Must be unique in the current API Management service instance. Non-current revision has ;rev=n as a suffix where n is the revision number. Changing this forces a new resource to be created.

* `policy_content` - (Required) Json escaped Xml Encoded contents of the Policy.

* `policy_id` - (Required) The identifier of the Policy. Changing this forces a new resource to be created.

* `content_format` - (Optional) Format of the policyContent. Defaults to `xml`.

## Attributes Reference

The following attributes are exported:

* `id` - Resource ID.

* `name` - Resource name.

* `type` - Resource type for API Management resource.
