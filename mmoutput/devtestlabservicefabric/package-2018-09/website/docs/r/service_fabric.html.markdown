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
page_title: "Azure Resource Manager: azurerm_service_fabric"
sidebar_current: "docs-azurerm-resource-service-fabric"
description: |-
  Manage Azure ServiceFabric instance.
---

# azurerm_service_fabric

Manage Azure ServiceFabric instance.


## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the user profile. Changing this forces a new resource to be created.

* `name` - (Required) The name of the service fabric. Changing this forces a new resource to be created.

* `resource_group` - (Required) The name of the resource group. Changing this forces a new resource to be created.

* `lab_name` - (Required) The name of the lab. Changing this forces a new resource to be created.

* `location` - (Optional) The location of the resource. Changing this forces a new resource to be created.

* `environment_id` - (Optional) The resource id of the environment under which the service fabric resource is present

* `external_service_fabric_id` - (Optional) The backing service fabric resource's id

* `tags` - (Optional) The tags of the resource. Changing this forces a new resource to be created.

## Attributes Reference

The following attributes are exported:

* `applicable_schedule` - One `applicable_schedule` block defined below.

* `provisioning_state` - The provisioning status of the resource.

* `unique_identifier` - The unique immutable identifier of a resource (Guid).

* `id` - The identifier of the resource.

* `name` - The name of the resource.

* `type` - The type of the resource.


---

The `applicable_schedule` block contains the following:

* `id` - The identifier of the resource.

* `name` - The name of the resource.

* `type` - The type of the resource.

* `location` - The location of the resource.

* `tags` - The tags of the resource.

* `lab_vms_shutdown` - One `lab_vms_shutdown` block defined below.

* `lab_vms_startup` - One `lab_vms_startup` block defined below.


---

The `lab_vms_shutdown` block contains the following:

* `id` - The identifier of the resource.

* `name` - The name of the resource.

* `type` - The type of the resource.

* `location` - The location of the resource.

* `tags` - The tags of the resource.

* `status` - The status of the schedule (i.e. Enabled, Disabled)

* `task_type` - The task type of the schedule (e.g. LabVmsShutdownTask, LabVmAutoStart).

* `weekly_recurrence` - One `weekly_recurrence` block defined below.

* `daily_recurrence` - One `daily_recurrence` block defined below.

* `hourly_recurrence` - One `hourly_recurrence` block defined below.

* `time_zone_id` - The time zone ID (e.g. Pacific Standard time).

* `notification_settings` - One `notification_setting` block defined below.

* `created_date` - The creation date of the schedule.

* `target_resource_id` - The resource ID to which the schedule belongs

* `provisioning_state` - The provisioning status of the resource.

* `unique_identifier` - The unique immutable identifier of a resource (Guid).


---

The `weekly_recurrence` block contains the following:

* `weekdays` - The days of the week for which the schedule is set (e.g. Sunday, Monday, Tuesday, etc.).

* `time` - The time of the day the schedule will occur.

---

The `daily_recurrence` block contains the following:

* `time` - The time of day the schedule will occur.

---

The `hourly_recurrence` block contains the following:

* `minute` - Minutes of the hour the schedule will run.

---

The `notification_setting` block contains the following:

* `status` - If notifications are enabled for this schedule (i.e. Enabled, Disabled).

* `time_in_minutes` - Time in minutes before event at which notification will be sent.

* `webhook_url` - The webhook URL to which the notification will be sent.

* `email_recipient` - The email recipient to send notifications to (can be a list of semi-colon separated email addresses).

* `notification_locale` - The locale to use when sending a notification (fallback for unsupported languages is EN).

---

The `lab_vms_startup` block contains the following:

* `id` - The identifier of the resource.

* `name` - The name of the resource.

* `type` - The type of the resource.

* `location` - The location of the resource.

* `tags` - The tags of the resource.

* `status` - The status of the schedule (i.e. Enabled, Disabled)

* `task_type` - The task type of the schedule (e.g. LabVmsShutdownTask, LabVmAutoStart).

* `weekly_recurrence` - One `weekly_recurrence` block defined below.

* `daily_recurrence` - One `daily_recurrence` block defined below.

* `hourly_recurrence` - One `hourly_recurrence` block defined below.

* `time_zone_id` - The time zone ID (e.g. Pacific Standard time).

* `notification_settings` - One `notification_setting` block defined below.

* `created_date` - The creation date of the schedule.

* `target_resource_id` - The resource ID to which the schedule belongs

* `provisioning_state` - The provisioning status of the resource.

* `unique_identifier` - The unique immutable identifier of a resource (Guid).


---

The `weekly_recurrence` block contains the following:

* `weekdays` - The days of the week for which the schedule is set (e.g. Sunday, Monday, Tuesday, etc.).

* `time` - The time of the day the schedule will occur.

---

The `daily_recurrence` block contains the following:

* `time` - The time of day the schedule will occur.

---

The `hourly_recurrence` block contains the following:

* `minute` - Minutes of the hour the schedule will run.

---

The `notification_setting` block contains the following:

* `status` - If notifications are enabled for this schedule (i.e. Enabled, Disabled).

* `time_in_minutes` - Time in minutes before event at which notification will be sent.

* `webhook_url` - The webhook URL to which the notification will be sent.

* `email_recipient` - The email recipient to send notifications to (can be a list of semi-colon separated email addresses).

* `notification_locale` - The locale to use when sending a notification (fallback for unsupported languages is EN).
