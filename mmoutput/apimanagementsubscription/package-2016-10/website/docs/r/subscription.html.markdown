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
page_title: "Azure Resource Manager: azurerm_subscription"
sidebar_current: "docs-azurerm-resource-subscription"
description: |-
  Manage Azure Subscription instance.
---

# azurerm_subscription

Manage Azure Subscription instance.


## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the API Management service. Changing this forces a new resource to be created.

* `resource_group` - (Required) The name of the resource group. Changing this forces a new resource to be created.

* `sid` - (Required) Subscription entity Identifier. The entity represents the association between a user and a product in API Management. Changing this forces a new resource to be created.

* `name` - (Optional) Subscription name. Changing this forces a new resource to be created.

* `expiration_date` - (Optional) New subscription expiration date. Changing this forces a new resource to be created.

* `primary_key` - (Optional) Primary subscription key. Changing this forces a new resource to be created.

* `product_id` - (Optional) Product identifier path: /products/{productId} Changing this forces a new resource to be created.

* `secondary_key` - (Optional) Secondary subscription key. Changing this forces a new resource to be created.

* `state` - (Optional) Subscription state. Possible states are * active – the subscription is active, * suspended – the subscription is blocked, and the subscriber cannot call any APIs of the product, * submitted – the subscription request has been made by the developer, but has not yet been approved or rejected, * rejected – the subscription request has been denied by an administrator, * cancelled – the subscription has been cancelled by the developer or administrator, * expired – the subscription reached its expiration date and was deactivated. Defaults to `Suspended`. Changing this forces a new resource to be created.

* `state_comment` - (Optional) Comments describing subscription state change by the administrator. Changing this forces a new resource to be created.

* `user_id` - (Optional) User identifier path: /users/{uid} Changing this forces a new resource to be created.

## Attributes Reference

The following attributes are exported:

* `id` - Uniquely identifies the subscription within the current API Management service instance. The value is a valid relative URL in the format of /subscriptions/{sid} where {sid} is a subscription identifier.

* `created_date` - Subscription creation date. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.<br>

* `start_date` - Subscription activation date. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.<br>

* `end_date` - Date when subscription was cancelled or expired. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.<br>

* `notification_date` - Upcoming subscription expiration notification date. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.<br>
