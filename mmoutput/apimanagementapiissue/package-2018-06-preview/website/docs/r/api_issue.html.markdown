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
page_title: "Azure Resource Manager: azurerm_api_issue"
sidebar_current: "docs-azurerm-resource-api-issue"
description: |-
  Manage Azure ApiIssue instance.
---

# azurerm_api_issue

Manage Azure ApiIssue instance.


## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the API Management service. Changing this forces a new resource to be created.

* `resource_group` - (Required) The name of the resource group. Changing this forces a new resource to be created.

* `api_id` - (Required) API identifier. Must be unique in the current API Management service instance. Changing this forces a new resource to be created.

* `description` - (Required) Text describing the issue.

* `issue_id` - (Required) Issue identifier. Must be unique in the current API Management service instance. Changing this forces a new resource to be created.

* `title` - (Required) The issue title.

* `user_id` - (Required) A resource identifier for the user created the issue.

* `api_id` - (Optional) A resource identifier for the API the issue was created for.

* `created_date` - (Optional) Date and time when the issue was created.

* `state` - (Optional) Status of the issue. Defaults to `proposed`.

## Attributes Reference

The following attributes are exported:

* `id` - Resource ID.

* `name` - Resource name.

* `type` - Resource type for API Management resource.
