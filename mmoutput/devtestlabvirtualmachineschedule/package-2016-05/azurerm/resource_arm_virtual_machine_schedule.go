// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file at
//     https://github.com/Azure/magic-module-specs
//
// ----------------------------------------------------------------------------

package azurerm



func resourceArmVirtualMachineSchedule() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmVirtualMachineScheduleCreate,
        Read: resourceArmVirtualMachineScheduleRead,
        Update: resourceArmVirtualMachineScheduleUpdate,
        Delete: resourceArmVirtualMachineScheduleDelete,

        Importer: &schema.ResourceImporter{
            State: schema.ImportStatePassthrough,
        },


        Schema: map[string]*schema.Schema{
            "lab_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "virtual_machine_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "daily_recurrence": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "time": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },

            "hourly_recurrence": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "minute": {
                            Type: schema.TypeInt,
                            Optional: true,
                        },
                    },
                },
            },

            "location": azure.SchemaLocation(),

            "notification_settings": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "status": {
                            Type: schema.TypeString,
                            Optional: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(devtestlab.Disabled),
                                string(devtestlab.Enabled),
                            }, false),
                            Default: string(devtestlab.Disabled),
                        },
                        "time_in_minutes": {
                            Type: schema.TypeInt,
                            Optional: true,
                        },
                        "webhook_url": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },

            "status": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(devtestlab.Enabled),
                    string(devtestlab.Disabled),
                }, false),
                Default: string(devtestlab.Enabled),
            },

            "tags": tags.Schema(),

            "target_resource_id": {
                Type: schema.TypeString,
                Optional: true,
            },

            "task_type": {
                Type: schema.TypeString,
                Optional: true,
            },

            "time_zone_id": {
                Type: schema.TypeString,
                Optional: true,
            },

            "unique_identifier": {
                Type: schema.TypeString,
                Optional: true,
            },

            "weekly_recurrence": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "time": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "weekdays": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                    },
                },
            },

            "created_date": {
                Type: schema.TypeString,
                Computed: true,
            },

            "id": {
                Type: schema.TypeString,
                Computed: true,
            },

            "name": {
                Type: schema.TypeString,
                Computed: true,
            },

            "provisioning_state": {
                Type: schema.TypeString,
                Computed: true,
            },

            "type": {
                Type: schema.TypeString,
                Computed: true,
            },
        },
    }
}

func resourceArmVirtualMachineScheduleCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).virtualMachineSchedulesClient
    ctx, cancel := timeouts.ForCreate(meta.(*ArmClient).StopContext, d)
    defer cancel()

    resourceGroupName := d.Get("resource_group").(string)
    labName := d.Get("lab_name").(string)
    name := d.Get("name").(string)
    name := d.Get("virtual_machine_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroupName, labName, name, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Virtual Machine Schedule (Name %q / Virtual Machine Name %q / Lab Name %q / Resource Group %q): %+v", name, name, labName, resourceGroupName, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_virtual_machine_schedule", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    dailyRecurrence := d.Get("daily_recurrence").([]interface{})
    hourlyRecurrence := d.Get("hourly_recurrence").([]interface{})
    notificationSettings := d.Get("notification_settings").([]interface{})
    status := d.Get("status").(string)
    targetResourceID := d.Get("target_resource_id").(string)
    taskType := d.Get("task_type").(string)
    timeZoneID := d.Get("time_zone_id").(string)
    uniqueIdentifier := d.Get("unique_identifier").(string)
    weeklyRecurrence := d.Get("weekly_recurrence").([]interface{})
    tags := d.Get("tags").(map[string]interface{})

    schedule := devtestlab.ScheduleFragment{
        Location: utils.String(location),
        SchedulePropertiesFragment: &devtestlab.SchedulePropertiesFragment{
            DailyRecurrence: expandArmVirtualMachineScheduleDayDetailsFragment(dailyRecurrence),
            HourlyRecurrence: expandArmVirtualMachineScheduleHourDetailsFragment(hourlyRecurrence),
            NotificationSettings: expandArmVirtualMachineScheduleNotificationSettingsFragment(notificationSettings),
            Status: devtestlab.EnableStatus(status),
            TargetResourceID: utils.String(targetResourceID),
            TaskType: utils.String(taskType),
            TimeZoneID: utils.String(timeZoneID),
            UniqueIdentifier: utils.String(uniqueIdentifier),
            WeeklyRecurrence: expandArmVirtualMachineScheduleWeekDetailsFragment(weeklyRecurrence),
        },
        Tags: tags.Expand(tags),
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroupName, labName, name, name, schedule); err != nil {
        return fmt.Errorf("Error creating Virtual Machine Schedule (Name %q / Virtual Machine Name %q / Lab Name %q / Resource Group %q): %+v", name, name, labName, resourceGroupName, err)
    }


    resp, err := client.Get(ctx, resourceGroupName, labName, name, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Virtual Machine Schedule (Name %q / Virtual Machine Name %q / Lab Name %q / Resource Group %q): %+v", name, name, labName, resourceGroupName, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Virtual Machine Schedule (Name %q / Virtual Machine Name %q / Lab Name %q / Resource Group %q) ID", name, name, labName, resourceGroupName)
    }
    d.SetId(*resp.ID)

    return resourceArmVirtualMachineScheduleRead(d, meta)
}

func resourceArmVirtualMachineScheduleRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).virtualMachineSchedulesClient
    ctx, cancel := timeouts.ForRead(meta.(*ArmClient).StopContext, d)
    defer cancel()

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroupName := id.ResourceGroup
    labName := id.Path["labs"]
    name := id.Path["virtualmachines"]
    name := id.Path["schedules"]

    resp, err := client.Get(ctx, resourceGroupName, labName, name, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Virtual Machine Schedule %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Virtual Machine Schedule (Name %q / Virtual Machine Name %q / Lab Name %q / Resource Group %q): %+v", name, name, labName, resourceGroupName, err)
    }


    d.Set("resource_group", resourceGroupName)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if schedulePropertiesFragment := resp.SchedulePropertiesFragment; schedulePropertiesFragment != nil {
        d.Set("created_date", (schedulePropertiesFragment.CreatedDate).String())
        if err := d.Set("daily_recurrence", flattenArmVirtualMachineScheduleDayDetailsFragment(schedulePropertiesFragment.DailyRecurrence)); err != nil {
            return fmt.Errorf("Error setting `daily_recurrence`: %+v", err)
        }
        if err := d.Set("hourly_recurrence", flattenArmVirtualMachineScheduleHourDetailsFragment(schedulePropertiesFragment.HourlyRecurrence)); err != nil {
            return fmt.Errorf("Error setting `hourly_recurrence`: %+v", err)
        }
        if err := d.Set("notification_settings", flattenArmVirtualMachineScheduleNotificationSettingsFragment(schedulePropertiesFragment.NotificationSettings)); err != nil {
            return fmt.Errorf("Error setting `notification_settings`: %+v", err)
        }
        d.Set("provisioning_state", schedulePropertiesFragment.ProvisioningState)
        d.Set("status", string(schedulePropertiesFragment.Status))
        d.Set("target_resource_id", schedulePropertiesFragment.TargetResourceID)
        d.Set("task_type", schedulePropertiesFragment.TaskType)
        d.Set("time_zone_id", schedulePropertiesFragment.TimeZoneID)
        d.Set("unique_identifier", schedulePropertiesFragment.UniqueIdentifier)
        if err := d.Set("weekly_recurrence", flattenArmVirtualMachineScheduleWeekDetailsFragment(schedulePropertiesFragment.WeeklyRecurrence)); err != nil {
            return fmt.Errorf("Error setting `weekly_recurrence`: %+v", err)
        }
    }
    d.Set("id", resp.ID)
    d.Set("lab_name", labName)
    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("type", resp.Type)
    d.Set("virtual_machine_name", name)

    return tags.FlattenAndSet(d, resp.Tags)
}

func resourceArmVirtualMachineScheduleUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).virtualMachineSchedulesClient
    ctx, cancel := timeouts.ForUpdate(meta.(*ArmClient).StopContext, d)
    defer cancel()

      resourceGroupName := d.Get("resource_group").(string)
    location := azure.NormalizeLocation(d.Get("location").(string))
    dailyRecurrence := d.Get("daily_recurrence").([]interface{})
    hourlyRecurrence := d.Get("hourly_recurrence").([]interface{})
    labName := d.Get("lab_name").(string)
    name := d.Get("name").(string)
    notificationSettings := d.Get("notification_settings").([]interface{})
    status := d.Get("status").(string)
    targetResourceID := d.Get("target_resource_id").(string)
    taskType := d.Get("task_type").(string)
    timeZoneID := d.Get("time_zone_id").(string)
    uniqueIdentifier := d.Get("unique_identifier").(string)
    name := d.Get("virtual_machine_name").(string)
    weeklyRecurrence := d.Get("weekly_recurrence").([]interface{})
    tags := d.Get("tags").(map[string]interface{})

    schedule := devtestlab.ScheduleFragment{
        Location: utils.String(location),
        SchedulePropertiesFragment: &devtestlab.SchedulePropertiesFragment{
            DailyRecurrence: expandArmVirtualMachineScheduleDayDetailsFragment(dailyRecurrence),
            HourlyRecurrence: expandArmVirtualMachineScheduleHourDetailsFragment(hourlyRecurrence),
            NotificationSettings: expandArmVirtualMachineScheduleNotificationSettingsFragment(notificationSettings),
            Status: devtestlab.EnableStatus(status),
            TargetResourceID: utils.String(targetResourceID),
            TaskType: utils.String(taskType),
            TimeZoneID: utils.String(timeZoneID),
            UniqueIdentifier: utils.String(uniqueIdentifier),
            WeeklyRecurrence: expandArmVirtualMachineScheduleWeekDetailsFragment(weeklyRecurrence),
        },
        Tags: tags.Expand(tags),
    }


    if _, err := client.Update(ctx, resourceGroupName, labName, name, name, schedule); err != nil {
        return fmt.Errorf("Error updating Virtual Machine Schedule (Name %q / Virtual Machine Name %q / Lab Name %q / Resource Group %q): %+v", name, name, labName, resourceGroupName, err)
    }

    return resourceArmVirtualMachineScheduleRead(d, meta)
}

func resourceArmVirtualMachineScheduleDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).virtualMachineSchedulesClient
    ctx, cancel := timeouts.ForDelete(meta.(*ArmClient).StopContext, d)
    defer cancel()


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroupName := id.ResourceGroup
    labName := id.Path["labs"]
    name := id.Path["virtualmachines"]
    name := id.Path["schedules"]

    if _, err := client.Delete(ctx, resourceGroupName, labName, name, name); err != nil {
        return fmt.Errorf("Error deleting Virtual Machine Schedule (Name %q / Virtual Machine Name %q / Lab Name %q / Resource Group %q): %+v", name, name, labName, resourceGroupName, err)
    }

    return nil
}

func expandArmVirtualMachineScheduleDayDetailsFragment(input []interface{}) *devtestlab.DayDetailsFragment {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    time := v["time"].(string)

    result := devtestlab.DayDetailsFragment{
        Time: utils.String(time),
    }
    return &result
}

func expandArmVirtualMachineScheduleHourDetailsFragment(input []interface{}) *devtestlab.HourDetailsFragment {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    minute := v["minute"].(int)

    result := devtestlab.HourDetailsFragment{
        Minute: utils.Int32(int32(minute)),
    }
    return &result
}

func expandArmVirtualMachineScheduleNotificationSettingsFragment(input []interface{}) *devtestlab.NotificationSettingsFragment {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    status := v["status"].(string)
    timeInMinutes := v["time_in_minutes"].(int)
    webhookURL := v["webhook_url"].(string)

    result := devtestlab.NotificationSettingsFragment{
        Status: devtestlab.NotificationStatus(status),
        TimeInMinutes: utils.Int32(int32(timeInMinutes)),
        WebhookURL: utils.String(webhookURL),
    }
    return &result
}

func expandArmVirtualMachineScheduleWeekDetailsFragment(input []interface{}) *devtestlab.WeekDetailsFragment {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    weekdays := v["weekdays"].([]interface{})
    time := v["time"].(string)

    result := devtestlab.WeekDetailsFragment{
        Time: utils.String(time),
        Weekdays: utils.ExpandStringSlice(weekdays),
    }
    return &result
}


func flattenArmVirtualMachineScheduleDayDetailsFragment(input *devtestlab.DayDetailsFragment) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if time := input.Time; time != nil {
        result["time"] = *time
    }

    return []interface{}{result}
}

func flattenArmVirtualMachineScheduleHourDetailsFragment(input *devtestlab.HourDetailsFragment) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if minute := input.Minute; minute != nil {
        result["minute"] = int(*minute)
    }

    return []interface{}{result}
}

func flattenArmVirtualMachineScheduleNotificationSettingsFragment(input *devtestlab.NotificationSettingsFragment) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    result["status"] = string(input.Status)
    if timeInMinutes := input.TimeInMinutes; timeInMinutes != nil {
        result["time_in_minutes"] = int(*timeInMinutes)
    }
    if webhookUrl := input.WebhookURL; webhookUrl != nil {
        result["webhook_url"] = *webhookUrl
    }

    return []interface{}{result}
}

func flattenArmVirtualMachineScheduleWeekDetailsFragment(input *devtestlab.WeekDetailsFragment) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if time := input.Time; time != nil {
        result["time"] = *time
    }
    result["weekdays"] = utils.FlattenStringSlice(input.Weekdays)

    return []interface{}{result}
}
