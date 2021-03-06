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
page_title: "Azure Resource Manager: azurerm_registration_assignment"
sidebar_current: "docs-azurerm-resource-registration-assignment"
description: |-
  Manage Azure RegistrationAssignment instance.
---

# azurerm_registration_assignment

Manage Azure RegistrationAssignment instance.


## Argument Reference

The following arguments are supported:

* `registration_assignment_id` - (Required) Guid of the registration assignment. Changing this forces a new resource to be created.

* `registration_definition_id` - (Required) Fully qualified path of the registration definition.

* `scope` - (Required) Scope of the resource. Changing this forces a new resource to be created.

## Attributes Reference

The following attributes are exported:

* `provisioning_state` - Current state of the registration assignment.

* `registration_definition` - One `registration_definition` block defined below.

* `id` - The fully qualified path of the registration assignment.

* `type` - Type of the resource.

* `name` - Name of the registration assignment.


---

The `registration_definition` block contains the following:

* `description` - Description of the registration definition.

* `authorizations` - One or more `authorization` block defined below.

* `registration_definition_name` - Name of the registration definition.

* `provisioning_state` - Current state of the registration definition.

* `managee_tenant_id` - Id of the home tenant.

* `managee_tenant_name` - Name of the home tenant.

* `managed_by_tenant_id` - Id of the managedBy tenant.

* `managed_by_tenant_name` - Name of the managedBy tenant.

* `plan` - One `plan` block defined below.

* `id` - Fully qualified path of the registration definition.

* `type` - Type of the resource (Microsoft.ManagedServices/registrationDefinitions).

* `name` - Name of the registration definition.


---

The `authorization` block contains the following:

* `principal_id` - Principal Id of the security group/service principal/user that would be assigned permissions to the projected subscription

* `role_definition_id` - The role definition identifier. This role will define all the permissions that the security group/service principal/user must have on the projected subscription. This role cannot be an owner role.

---

The `plan` block contains the following:

* `name` - The plan name.

* `publisher` - The publisher ID.

* `product` - The product code.

* `version` - The plan's version.
