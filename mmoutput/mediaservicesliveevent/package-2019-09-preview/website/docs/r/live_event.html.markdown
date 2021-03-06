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
page_title: "Azure Resource Manager: azurerm_live_event"
sidebar_current: "docs-azurerm-resource-live-event"
description: |-
  Manage Azure LiveEvent instance.
---

# azurerm_live_event

Manage Azure LiveEvent instance.


## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the Live Event. Changing this forces a new resource to be created.

* `resource_group` - (Required) The name of the resource group within the Azure subscription. Changing this forces a new resource to be created.

* `account_name` - (Required) The Media Services account name. Changing this forces a new resource to be created.

* `input` - (Required) One `input` block defined below.

---

The `input` block supports the following:

* `streaming_protocol` - (Required) The streaming protocol for the Live Event.  This is specified at creation time and cannot be updated.

* `access_control` - (Optional) One `access_control` block defined below.

* `key_frame_interval_duration` - (Optional) ISO 8601 timespan duration of the key frame interval duration.

* `access_token` - (Optional) A unique identifier for a stream.  This can be specified at creation time but cannot be updated.  If omitted, the service will generate a unique value.

* `endpoints` - (Optional) One or more `endpoint` block defined below.


---

The `access_control` block supports the following:

* `ip` - (Optional) One `ip` block defined below.


---

The `ip` block supports the following:

* `allow` - (Optional) One or more `allow` block defined below.


---

The `allow` block supports the following:

* `name` - (Optional) The friendly name for the IP address range.

* `address` - (Optional) The IP address.

* `subnet_prefix_length` - (Optional) The subnet mask prefix length (see CIDR notation).

---

The `endpoint` block supports the following:

* `protocol` - (Optional) The endpoint protocol.

* `url` - (Optional) The endpoint URL.

* `location` - (Optional) The Azure Region of the resource. Changing this forces a new resource to be created.

* `auto_start` - (Optional) The flag indicates if the resource should be automatically started on creation. Changing this forces a new resource to be created.

* `cross_site_access_policies` - (Optional) One `cross_site_access_policy` block defined below.

* `description` - (Optional) The Live Event description.

* `encoding` - (Optional) One `encoding` block defined below.

* `preview` - (Optional) One `preview` block defined below.

* `stream_options` - (Optional) The options to use for the LiveEvent.  This value is specified at creation time and cannot be updated.

* `vanity_url` - (Optional) Specifies whether to use a vanity url with the Live Event.  This value is specified at creation time and cannot be updated.

* `tags` - (Optional) Resource tags. Changing this forces a new resource to be created.

---

The `cross_site_access_policy` block supports the following:

* `client_access_policy` - (Optional) The content of clientaccesspolicy.xml used by Silverlight.

* `cross_domain_policy` - (Optional) The content of crossdomain.xml used by Silverlight.

---

The `encoding` block supports the following:

* `encoding_type` - (Optional) The encoding type for Live Event.  This value is specified at creation time and cannot be updated. Defaults to `None`.

* `preset_name` - (Optional) The encoding preset name.  This value is specified at creation time and cannot be updated.

---

The `preview` block supports the following:

* `endpoints` - (Optional) One or more `endpoint` block defined below.

* `access_control` - (Optional) One `access_control` block defined below.

* `preview_locator` - (Optional) The identifier of the preview locator in Guid format.  Specifying this at creation time allows the caller to know the preview locator url before the event is created.  If omitted, the service will generate a random identifier.  This value cannot be updated once the live event is created.

* `streaming_policy_name` - (Optional) The name of streaming policy used for the LiveEvent preview.  This value is specified at creation time and cannot be updated.

* `alternative_media_id` - (Optional) An Alternative Media Identifier associated with the StreamingLocator created for the preview.  This value is specified at creation time and cannot be updated.  The identifier can be used in the CustomLicenseAcquisitionUrlTemplate or the CustomKeyAcquisitionUrlTemplate of the StreamingPolicy specified in the StreamingPolicyName field.


---

The `endpoint` block supports the following:

* `protocol` - (Optional) The endpoint protocol.

* `url` - (Optional) The endpoint URL.

---

The `access_control` block supports the following:

* `ip` - (Optional) One `ip` block defined below.


---

The `ip` block supports the following:

* `allow` - (Optional) One or more `allow` block defined below.


---

The `allow` block supports the following:

* `name` - (Optional) The friendly name for the IP address range.

* `address` - (Optional) The IP address.

* `subnet_prefix_length` - (Optional) The subnet mask prefix length (see CIDR notation).

## Attributes Reference

The following attributes are exported:

* `provisioning_state` - The provisioning state of the Live Event.

* `resource_state` - The resource state of the Live Event.

* `created` - The exact time the Live Event was created.

* `last_modified` - The exact time the Live Event was last modified.

* `id` - Fully qualified resource ID for the resource.

* `name` - The name of the resource.

* `type` - The type of the resource.
