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
page_title: "Azure Resource Manager: azurerm_link"
sidebar_current: "docs-azurerm-resource-link"
description: |-
  Manage Azure Link instance.
---

# azurerm_link

Manage Azure Link instance.


## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the link. Changing this forces a new resource to be created.

* `resource_group` - (Required) The name of the resource group. Changing this forces a new resource to be created.

* `hub_name` - (Required) The name of the hub. Changing this forces a new resource to be created.

* `participant_property_references` - (Required) One or more `participant_property_reference` block defined below.

* `source_entity_type` - (Required) Type of source entity.

* `source_entity_type_name` - (Required) Name of the source Entity Type.

* `target_entity_type` - (Required) Type of target entity.

* `target_entity_type_name` - (Required) Name of the target Entity Type.

---

The `participant_property_reference` block supports the following:

* `source_property_name` - (Required) The source property that maps to the target property.

* `target_property_name` - (Required) The target property that maps to the source property.

* `description` - (Optional) Localized descriptions for the Link.

* `display_name` - (Optional) Localized display name for the Link.

* `mappings` - (Optional) One or more `mapping` block defined below.

* `operation_type` - (Optional) Determines whether this link is supposed to create or delete instances if Link is NOT Reference Only. Defaults to `Upsert`.

* `reference_only` - (Optional) Indicating whether the link is reference only link. This flag is ignored if the Mappings are defined. If the mappings are not defined and it is set to true, links processing will not create or update profiles.

---

The `mapping` block supports the following:

* `source_property_name` - (Required) Property name on the source Entity Type.

* `target_property_name` - (Required) Property name on the target Entity Type.

* `link_type` - (Optional) Link type. Defaults to `UpdateAlways`.

## Attributes Reference

The following attributes are exported:

* `tenant_id` - The hub name.

* `link_name` - The link name.

* `provisioning_state` - Provisioning state.

* `id` - Resource ID.

* `name` - Resource name.

* `type` - Resource type.
