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



func resourceArmServiceFabricSchedule() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmServiceFabricScheduleCreate,
        Read: resourceArmServiceFabricScheduleRead,
        Update: resourceArmServiceFabricScheduleUpdate,
        Delete: resourceArmServiceFabricScheduleDelete,

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

            "location": azure.SchemaLocation(),

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "lab_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "service_fabric_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "user_name": {
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

func resourceArmServiceFabricScheduleCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).serviceFabricSchedulesClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    labName := d.Get("lab_name").(string)
    serviceFabricName := d.Get("service_fabric_name").(string)
    userName := d.Get("user_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, labName, userName, serviceFabricName, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Service Fabric Schedule %q (Service Fabric Name %q / User Name %q / Lab Name %q / Resource Group %q): %+v", name, serviceFabricName, userName, labName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_service_fabric_schedule", *existing.ID)
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

    schedule := devtestlab.Schedule{
        Location: utils.String(location),
        ScheduleProperties: &devtestlab.ScheduleProperties{
            DailyRecurrence: expandArmServiceFabricScheduleDayDetails(dailyRecurrence),
            HourlyRecurrence: expandArmServiceFabricScheduleHourDetails(hourlyRecurrence),
            NotificationSettings: expandArmServiceFabricScheduleNotificationSettings(notificationSettings),
            Status: devtestlab.EnableStatus(status),
            TargetResourceID: utils.String(targetResourceId),
            TaskType: utils.String(taskType),
            TimeZoneID: utils.String(timeZoneId),
            WeeklyRecurrence: expandArmServiceFabricScheduleWeekDetails(weeklyRecurrence),
        },
        Tags: tags.Expand(t),
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroup, labName, userName, serviceFabricName, name, schedule); err != nil {
        return fmt.Errorf("Error creating Service Fabric Schedule %q (Service Fabric Name %q / User Name %q / Lab Name %q / Resource Group %q): %+v", name, serviceFabricName, userName, labName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, labName, userName, serviceFabricName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Service Fabric Schedule %q (Service Fabric Name %q / User Name %q / Lab Name %q / Resource Group %q): %+v", name, serviceFabricName, userName, labName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Service Fabric Schedule %q (Service Fabric Name %q / User Name %q / Lab Name %q / Resource Group %q) ID", name, serviceFabricName, userName, labName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmServiceFabricScheduleRead(d, meta)
}

func resourceArmServiceFabricScheduleRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).serviceFabricSchedulesClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    labName := id.Path["labs"]
    userName := id.Path["users"]
    serviceFabricName := id.Path["servicefabrics"]
    name := id.Path["schedules"]

    resp, err := client.Get(ctx, resourceGroup, labName, userName, serviceFabricName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Service Fabric Schedule %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Service Fabric Schedule %q (Service Fabric Name %q / User Name %q / Lab Name %q / Resource Group %q): %+v", name, serviceFabricName, userName, labName, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if scheduleProperties := resp.ScheduleProperties; scheduleProperties != nil {
        d.Set("created_date", (scheduleProperties.CreatedDate).String())
        if err := d.Set("daily_recurrence", flattenArmServiceFabricScheduleDayDetails(scheduleProperties.DailyRecurrence)); err != nil {
            return fmt.Errorf("Error setting `daily_recurrence`: %+v", err)
        }
        if err := d.Set("hourly_recurrence", flattenArmServiceFabricScheduleHourDetails(scheduleProperties.HourlyRecurrence)); err != nil {
            return fmt.Errorf("Error setting `hourly_recurrence`: %+v", err)
        }
        if err := d.Set("notification_settings", flattenArmServiceFabricScheduleNotificationSettings(scheduleProperties.NotificationSettings)); err != nil {
            return fmt.Errorf("Error setting `notification_settings`: %+v", err)
        }
        d.Set("provisioning_state", scheduleProperties.ProvisioningState)
        d.Set("status", string(scheduleProperties.Status))
        d.Set("target_resource_id", scheduleProperties.TargetResourceID)
        d.Set("task_type", scheduleProperties.TaskType)
        d.Set("time_zone_id", scheduleProperties.TimeZoneID)
        d.Set("unique_identifier", scheduleProperties.UniqueIdentifier)
        if err := d.Set("weekly_recurrence", flattenArmServiceFabricScheduleWeekDetails(scheduleProperties.WeeklyRecurrence)); err != nil {
            return fmt.Errorf("Error setting `weekly_recurrence`: %+v", err)
        }
    }
    d.Set("lab_name", labName)
    d.Set("service_fabric_name", serviceFabricName)
    d.Set("type", resp.Type)
    d.Set("user_name", userName)

    return tags.FlattenAndSet(d, resp.Tags)
}

func resourceArmServiceFabricScheduleUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).serviceFabricSchedulesClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    dailyRecurrence := d.Get("daily_recurrence").([]interface{})
    hourlyRecurrence := d.Get("hourly_recurrence").([]interface{})
    labName := d.Get("lab_name").(string)
    notificationSettings := d.Get("notification_settings").([]interface{})
    serviceFabricName := d.Get("service_fabric_name").(string)
    status := d.Get("status").(string)
    targetResourceId := d.Get("target_resource_id").(string)
    taskType := d.Get("task_type").(string)
    timeZoneId := d.Get("time_zone_id").(string)
    userName := d.Get("user_name").(string)
    weeklyRecurrence := d.Get("weekly_recurrence").([]interface{})
    t := d.Get("tags").(map[string]interface{})

    schedule := devtestlab.Schedule{
        Location: utils.String(location),
        ScheduleProperties: &devtestlab.ScheduleProperties{
            DailyRecurrence: expandArmServiceFabricScheduleDayDetails(dailyRecurrence),
            HourlyRecurrence: expandArmServiceFabricScheduleHourDetails(hourlyRecurrence),
            NotificationSettings: expandArmServiceFabricScheduleNotificationSettings(notificationSettings),
            Status: devtestlab.EnableStatus(status),
            TargetResourceID: utils.String(targetResourceId),
            TaskType: utils.String(taskType),
            TimeZoneID: utils.String(timeZoneId),
            WeeklyRecurrence: expandArmServiceFabricScheduleWeekDetails(weeklyRecurrence),
        },
        Tags: tags.Expand(t),
    }


    if _, err := client.Update(ctx, resourceGroup, labName, userName, serviceFabricName, name, schedule); err != nil {
        return fmt.Errorf("Error updating Service Fabric Schedule %q (Service Fabric Name %q / User Name %q / Lab Name %q / Resource Group %q): %+v", name, serviceFabricName, userName, labName, resourceGroup, err)
    }

    return resourceArmServiceFabricScheduleRead(d, meta)
}

func resourceArmServiceFabricScheduleDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).serviceFabricSchedulesClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    labName := id.Path["labs"]
    userName := id.Path["users"]
    serviceFabricName := id.Path["servicefabrics"]
    name := id.Path["schedules"]

    if _, err := client.Delete(ctx, resourceGroup, labName, userName, serviceFabricName, name); err != nil {
        return fmt.Errorf("Error deleting Service Fabric Schedule %q (Service Fabric Name %q / User Name %q / Lab Name %q / Resource Group %q): %+v", name, serviceFabricName, userName, labName, resourceGroup, err)
    }

    return nil
}

func expandArmServiceFabricScheduleDayDetails(input []interface{}) *devtestlab.DayDetails {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    time := v["time"].(string)

    result := devtestlab.DayDetails{
        Time: utils.String(time),
    }
    return &result
}

func expandArmServiceFabricScheduleHourDetails(input []interface{}) *devtestlab.HourDetails {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    minute := v["minute"].(int)

    result := devtestlab.HourDetails{
        Minute: utils.Int32(int32(minute)),
    }
    return &result
}

func expandArmServiceFabricScheduleNotificationSettings(input []interface{}) *devtestlab.NotificationSettings {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    status := v["status"].(string)
    timeInMinutes := v["time_in_minutes"].(int)
    webhookUrl := v["webhook_url"].(string)
    emailRecipient := v["email_recipient"].(string)
    notificationLocale := v["notification_locale"].(string)

    result := devtestlab.NotificationSettings{
        EmailRecipient: utils.String(emailRecipient),
        NotificationLocale: utils.String(notificationLocale),
        Status: devtestlab.EnableStatus(status),
        TimeInMinutes: utils.Int32(int32(timeInMinutes)),
        WebhookURL: utils.String(webhookUrl),
    }
    return &result
}

func expandArmServiceFabricScheduleWeekDetails(input []interface{}) *devtestlab.WeekDetails {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    weekdays := v["weekdays"].([]interface{})
    time := v["time"].(string)

    result := devtestlab.WeekDetails{
        Time: utils.String(time),
        Weekdays: utils.ExpandStringSlice(weekdays),
    }
    return &result
}


func flattenArmServiceFabricScheduleDayDetails(input *devtestlab.DayDetails) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if time := input.Time; time != nil {
        result["time"] = *time
    }

    return []interface{}{result}
}

func flattenArmServiceFabricScheduleHourDetails(input *devtestlab.HourDetails) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if minute := input.Minute; minute != nil {
        result["minute"] = int(*minute)
    }

    return []interface{}{result}
}

func flattenArmServiceFabricScheduleNotificationSettings(input *devtestlab.NotificationSettings) []interface{} {
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

func flattenArmServiceFabricScheduleWeekDetails(input *devtestlab.WeekDetails) []interface{} {
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
