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



func resourceArmGlobalSchedule() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmGlobalScheduleCreate,
        Read: resourceArmGlobalScheduleRead,
        Update: resourceArmGlobalScheduleUpdate,
        Delete: resourceArmGlobalScheduleDelete,

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
                Computed: true,
            },

            "location": azure.SchemaLocation(),

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "current_resource_id": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
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

            "target_resource_id": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
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

func resourceArmGlobalScheduleCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).globalSchedulesClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Global Schedule %q (Resource Group %q): %+v", name, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_global_schedule", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    currentResourceId := d.Get("current_resource_id").(string)
    dailyRecurrence := d.Get("daily_recurrence").([]interface{})
    hourlyRecurrence := d.Get("hourly_recurrence").([]interface{})
    notificationSettings := d.Get("notification_settings").([]interface{})
    status := d.Get("status").(string)
    targetResourceId := d.Get("target_resource_id").(string)
    targetResourceId := d.Get("target_resource_id").(string)
    taskType := d.Get("task_type").(string)
    timeZoneId := d.Get("time_zone_id").(string)
    weeklyRecurrence := d.Get("weekly_recurrence").([]interface{})
    t := d.Get("tags").(map[string]interface{})

    schedule := devtestlab.ScheduleFragment{
        CurrentResourceID: utils.String(currentResourceId),
        Location: utils.String(location),
        SchedulePropertiesFragment: &devtestlab.SchedulePropertiesFragment{
            DailyRecurrence: expandArmGlobalScheduleDayDetailsFragment(dailyRecurrence),
            HourlyRecurrence: expandArmGlobalScheduleHourDetailsFragment(hourlyRecurrence),
            NotificationSettings: expandArmGlobalScheduleNotificationSettingsFragment(notificationSettings),
            Status: devtestlab.EnableStatus(status),
            TargetResourceID: utils.String(targetResourceId),
            TaskType: utils.String(taskType),
            TimeZoneID: utils.String(timeZoneId),
            WeeklyRecurrence: expandArmGlobalScheduleWeekDetailsFragment(weeklyRecurrence),
        },
        Tags: tags.Expand(t),
        TargetResourceID: utils.String(targetResourceId),
        TargetResourceID: utils.String(targetResourceId),
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroup, name, schedule); err != nil {
        return fmt.Errorf("Error creating Global Schedule %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Global Schedule %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Global Schedule %q (Resource Group %q) ID", name, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmGlobalScheduleRead(d, meta)
}

func resourceArmGlobalScheduleRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).globalSchedulesClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["schedules"]

    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Global Schedule %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Global Schedule %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if schedulePropertiesFragment := resp.SchedulePropertiesFragment; schedulePropertiesFragment != nil {
        d.Set("created_date", (schedulePropertiesFragment.CreatedDate).String())
        if err := d.Set("daily_recurrence", flattenArmGlobalScheduleDayDetailsFragment(schedulePropertiesFragment.DailyRecurrence)); err != nil {
            return fmt.Errorf("Error setting `daily_recurrence`: %+v", err)
        }
        if err := d.Set("hourly_recurrence", flattenArmGlobalScheduleHourDetailsFragment(schedulePropertiesFragment.HourlyRecurrence)); err != nil {
            return fmt.Errorf("Error setting `hourly_recurrence`: %+v", err)
        }
        if err := d.Set("notification_settings", flattenArmGlobalScheduleNotificationSettingsFragment(schedulePropertiesFragment.NotificationSettings)); err != nil {
            return fmt.Errorf("Error setting `notification_settings`: %+v", err)
        }
        d.Set("provisioning_state", schedulePropertiesFragment.ProvisioningState)
        d.Set("status", string(schedulePropertiesFragment.Status))
        d.Set("target_resource_id", schedulePropertiesFragment.TargetResourceID)
        d.Set("task_type", schedulePropertiesFragment.TaskType)
        d.Set("time_zone_id", schedulePropertiesFragment.TimeZoneID)
        d.Set("unique_identifier", schedulePropertiesFragment.UniqueIdentifier)
        if err := d.Set("weekly_recurrence", flattenArmGlobalScheduleWeekDetailsFragment(schedulePropertiesFragment.WeeklyRecurrence)); err != nil {
            return fmt.Errorf("Error setting `weekly_recurrence`: %+v", err)
        }
    }
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}

func resourceArmGlobalScheduleUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).globalSchedulesClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    currentResourceId := d.Get("current_resource_id").(string)
    dailyRecurrence := d.Get("daily_recurrence").([]interface{})
    hourlyRecurrence := d.Get("hourly_recurrence").([]interface{})
    notificationSettings := d.Get("notification_settings").([]interface{})
    status := d.Get("status").(string)
    targetResourceId := d.Get("target_resource_id").(string)
    targetResourceId := d.Get("target_resource_id").(string)
    taskType := d.Get("task_type").(string)
    timeZoneId := d.Get("time_zone_id").(string)
    weeklyRecurrence := d.Get("weekly_recurrence").([]interface{})
    t := d.Get("tags").(map[string]interface{})

    schedule := devtestlab.ScheduleFragment{
        CurrentResourceID: utils.String(currentResourceId),
        SchedulePropertiesFragment: &devtestlab.SchedulePropertiesFragment{
            DailyRecurrence: expandArmGlobalScheduleDayDetailsFragment(dailyRecurrence),
            HourlyRecurrence: expandArmGlobalScheduleHourDetailsFragment(hourlyRecurrence),
            NotificationSettings: expandArmGlobalScheduleNotificationSettingsFragment(notificationSettings),
            Status: devtestlab.EnableStatus(status),
            TargetResourceID: utils.String(targetResourceId),
            TaskType: utils.String(taskType),
            TimeZoneID: utils.String(timeZoneId),
            WeeklyRecurrence: expandArmGlobalScheduleWeekDetailsFragment(weeklyRecurrence),
        },
        Tags: tags.Expand(t),
        TargetResourceID: utils.String(targetResourceId),
        TargetResourceID: utils.String(targetResourceId),
    }


    if _, err := client.Update(ctx, resourceGroup, name, schedule); err != nil {
        return fmt.Errorf("Error updating Global Schedule %q (Resource Group %q): %+v", name, resourceGroup, err)
    }

    return resourceArmGlobalScheduleRead(d, meta)
}

func resourceArmGlobalScheduleDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).globalSchedulesClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["schedules"]

    if _, err := client.Delete(ctx, resourceGroup, name); err != nil {
        return fmt.Errorf("Error deleting Global Schedule %q (Resource Group %q): %+v", name, resourceGroup, err)
    }

    return nil
}

func expandArmGlobalScheduleDayDetailsFragment(input []interface{}) *devtestlab.DayDetailsFragment {
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

func expandArmGlobalScheduleHourDetailsFragment(input []interface{}) *devtestlab.HourDetailsFragment {
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

func expandArmGlobalScheduleNotificationSettingsFragment(input []interface{}) *devtestlab.NotificationSettingsFragment {
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

func expandArmGlobalScheduleWeekDetailsFragment(input []interface{}) *devtestlab.WeekDetailsFragment {
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


func flattenArmGlobalScheduleDayDetailsFragment(input *devtestlab.DayDetailsFragment) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if time := input.Time; time != nil {
        result["time"] = *time
    }

    return []interface{}{result}
}

func flattenArmGlobalScheduleHourDetailsFragment(input *devtestlab.HourDetailsFragment) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if minute := input.Minute; minute != nil {
        result["minute"] = int(*minute)
    }

    return []interface{}{result}
}

func flattenArmGlobalScheduleNotificationSettingsFragment(input *devtestlab.NotificationSettingsFragment) []interface{} {
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

func flattenArmGlobalScheduleWeekDetailsFragment(input *devtestlab.WeekDetailsFragment) []interface{} {
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
