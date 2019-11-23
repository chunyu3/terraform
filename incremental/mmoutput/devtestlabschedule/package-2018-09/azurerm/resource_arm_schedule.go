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



func resourceArmSchedule() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmScheduleCreate,
        Read: resourceArmScheduleRead,
        Update: resourceArmScheduleUpdate,
        Delete: resourceArmScheduleDelete,

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

            "type": {
                Type: schema.TypeString,
                Computed: true,
            },

            "tags": tags.Schema(),
        },
    }
}

func resourceArmScheduleCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).schedulesClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, name, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Schedule %q (Resource Group %q): %+v", name, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_schedule", *existing.ID)
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
            DailyRecurrence: expandArmScheduleDayDetailsFragment(dailyRecurrence),
            HourlyRecurrence: expandArmScheduleHourDetailsFragment(hourlyRecurrence),
            NotificationSettings: expandArmScheduleNotificationSettingsFragment(notificationSettings),
            Status: devtestlab.EnableStatus(status),
            TargetResourceID: utils.String(targetResourceId),
            TaskType: utils.String(taskType),
            TimeZoneID: utils.String(timeZoneId),
            WeeklyRecurrence: expandArmScheduleWeekDetailsFragment(weeklyRecurrence),
        },
        Tags: tags.Expand(t),
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroup, name, name, schedule); err != nil {
        return fmt.Errorf("Error creating Schedule %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, name, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Schedule %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Schedule %q (Resource Group %q) ID", name, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmScheduleRead(d, meta)
}

func resourceArmScheduleRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).schedulesClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["labs"]
    name := id.Path["schedules"]

    resp, err := client.Get(ctx, resourceGroup, name, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Schedule %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Schedule %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", name)
    d.Set("resource_group", resourceGroup)
    d.Set("type", resp.Type)

    return nil
}

func resourceArmScheduleUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).schedulesClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
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
        SchedulePropertiesFragment: &devtestlab.SchedulePropertiesFragment{
            DailyRecurrence: expandArmScheduleDayDetailsFragment(dailyRecurrence),
            HourlyRecurrence: expandArmScheduleHourDetailsFragment(hourlyRecurrence),
            NotificationSettings: expandArmScheduleNotificationSettingsFragment(notificationSettings),
            Status: devtestlab.EnableStatus(status),
            TargetResourceID: utils.String(targetResourceId),
            TaskType: utils.String(taskType),
            TimeZoneID: utils.String(timeZoneId),
            WeeklyRecurrence: expandArmScheduleWeekDetailsFragment(weeklyRecurrence),
        },
        Tags: tags.Expand(t),
    }


    if _, err := client.Update(ctx, resourceGroup, name, name, schedule); err != nil {
        return fmt.Errorf("Error updating Schedule %q (Resource Group %q): %+v", name, resourceGroup, err)
    }

    return resourceArmScheduleRead(d, meta)
}

func resourceArmScheduleDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).schedulesClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["labs"]
    name := id.Path["schedules"]

    if _, err := client.Delete(ctx, resourceGroup, name, name); err != nil {
        return fmt.Errorf("Error deleting Schedule %q (Resource Group %q): %+v", name, resourceGroup, err)
    }

    return nil
}

func expandArmScheduleDayDetailsFragment(input []interface{}) *devtestlab.DayDetailsFragment {
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

func expandArmScheduleHourDetailsFragment(input []interface{}) *devtestlab.HourDetailsFragment {
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

func expandArmScheduleNotificationSettingsFragment(input []interface{}) *devtestlab.NotificationSettingsFragment {
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

func expandArmScheduleWeekDetailsFragment(input []interface{}) *devtestlab.WeekDetailsFragment {
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