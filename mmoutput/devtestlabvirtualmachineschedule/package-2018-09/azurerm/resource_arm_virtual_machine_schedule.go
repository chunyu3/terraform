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
            "name": {
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

            "location": azure.SchemaLocation(),

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "lab_name": {
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

            "notification_settings": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "email_recipient": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "notification_locale": {
                            Type: schema.TypeString,
                            Optional: true,
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

            "provisioning_state": {
                Type: schema.TypeString,
                Computed: true,
            },

            "type": {
                Type: schema.TypeString,
                Computed: true,
            },

            "unique_identifier": {
                Type: schema.TypeString,
                Computed: true,
            },

            "tags": tags.Schema(),
        },
    }
}

func resourceArmVirtualMachineScheduleCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).virtualMachineSchedulesClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    labName := d.Get("lab_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, labName, name, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Virtual Machine Schedule %q (Lab Name %q / Resource Group %q): %+v", name, labName, resourceGroup, err)
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
    targetResourceId := d.Get("target_resource_id").(string)
    taskType := d.Get("task_type").(string)
    timeZoneId := d.Get("time_zone_id").(string)
    weeklyRecurrence := d.Get("weekly_recurrence").([]interface{})
    t := d.Get("tags").(map[string]interface{})

    schedule := devtestlab.ScheduleFragment{
        Location: utils.String(location),
        SchedulePropertiesFragment: &devtestlab.SchedulePropertiesFragment{
            DailyRecurrence: expandArmVirtualMachineScheduleDayDetailsFragment(dailyRecurrence),
            HourlyRecurrence: expandArmVirtualMachineScheduleHourDetailsFragment(hourlyRecurrence),
            NotificationSettings: expandArmVirtualMachineScheduleNotificationSettingsFragment(notificationSettings),
            Status: devtestlab.EnableStatus(status),
            TargetResourceID: utils.String(targetResourceId),
            TaskType: utils.String(taskType),
            TimeZoneID: utils.String(timeZoneId),
            WeeklyRecurrence: expandArmVirtualMachineScheduleWeekDetailsFragment(weeklyRecurrence),
        },
        Tags: tags.Expand(t),
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroup, labName, name, name, schedule); err != nil {
        return fmt.Errorf("Error creating Virtual Machine Schedule %q (Lab Name %q / Resource Group %q): %+v", name, labName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, labName, name, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Virtual Machine Schedule %q (Lab Name %q / Resource Group %q): %+v", name, labName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Virtual Machine Schedule %q (Lab Name %q / Resource Group %q) ID", name, labName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmVirtualMachineScheduleRead(d, meta)
}

func resourceArmVirtualMachineScheduleRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).virtualMachineSchedulesClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    labName := id.Path["labs"]
    name := id.Path["virtualmachines"]
    name := id.Path["schedules"]

    resp, err := client.Get(ctx, resourceGroup, labName, name, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Virtual Machine Schedule %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Virtual Machine Schedule %q (Lab Name %q / Resource Group %q): %+v", name, labName, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", name)
    d.Set("resource_group", resourceGroup)
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
    d.Set("lab_name", labName)
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}

func resourceArmVirtualMachineScheduleUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).virtualMachineSchedulesClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    dailyRecurrence := d.Get("daily_recurrence").([]interface{})
    hourlyRecurrence := d.Get("hourly_recurrence").([]interface{})
    labName := d.Get("lab_name").(string)
    notificationSettings := d.Get("notification_settings").([]interface{})
    status := d.Get("status").(string)
    targetResourceId := d.Get("target_resource_id").(string)
    taskType := d.Get("task_type").(string)
    timeZoneId := d.Get("time_zone_id").(string)
    weeklyRecurrence := d.Get("weekly_recurrence").([]interface{})
    t := d.Get("tags").(map[string]interface{})

    schedule := devtestlab.ScheduleFragment{
        Location: utils.String(location),
        SchedulePropertiesFragment: &devtestlab.SchedulePropertiesFragment{
            DailyRecurrence: expandArmVirtualMachineScheduleDayDetailsFragment(dailyRecurrence),
            HourlyRecurrence: expandArmVirtualMachineScheduleHourDetailsFragment(hourlyRecurrence),
            NotificationSettings: expandArmVirtualMachineScheduleNotificationSettingsFragment(notificationSettings),
            Status: devtestlab.EnableStatus(status),
            TargetResourceID: utils.String(targetResourceId),
            TaskType: utils.String(taskType),
            TimeZoneID: utils.String(timeZoneId),
            WeeklyRecurrence: expandArmVirtualMachineScheduleWeekDetailsFragment(weeklyRecurrence),
        },
        Tags: tags.Expand(t),
    }


    if _, err := client.Update(ctx, resourceGroup, labName, name, name, schedule); err != nil {
        return fmt.Errorf("Error updating Virtual Machine Schedule %q (Lab Name %q / Resource Group %q): %+v", name, labName, resourceGroup, err)
    }

    return resourceArmVirtualMachineScheduleRead(d, meta)
}

func resourceArmVirtualMachineScheduleDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).virtualMachineSchedulesClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    labName := id.Path["labs"]
    name := id.Path["virtualmachines"]
    name := id.Path["schedules"]

    if _, err := client.Delete(ctx, resourceGroup, labName, name, name); err != nil {
        return fmt.Errorf("Error deleting Virtual Machine Schedule %q (Lab Name %q / Resource Group %q): %+v", name, labName, resourceGroup, err)
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
    webhookUrl := v["webhook_url"].(string)
    emailRecipient := v["email_recipient"].(string)
    notificationLocale := v["notification_locale"].(string)

    result := devtestlab.NotificationSettingsFragment{
        EmailRecipient: utils.String(emailRecipient),
        NotificationLocale: utils.String(notificationLocale),
        Status: devtestlab.EnableStatus(status),
        TimeInMinutes: utils.Int32(int32(timeInMinutes)),
        WebhookURL: utils.String(webhookUrl),
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

    if emailRecipient := input.EmailRecipient; emailRecipient != nil {
        result["email_recipient"] = *emailRecipient
    }
    if notificationLocale := input.NotificationLocale; notificationLocale != nil {
        result["notification_locale"] = *notificationLocale
    }
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
