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



func resourceArmServiceFabric() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmServiceFabricCreate,
        Read: resourceArmServiceFabricRead,
        Update: resourceArmServiceFabricUpdate,
        Delete: resourceArmServiceFabricDelete,

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

            "name": {
                Type: schema.TypeString,
                Computed: true,
            },

            "location": azure.SchemaLocation(),

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "lab_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "environment_id": {
                Type: schema.TypeString,
                Optional: true,
            },

            "external_service_fabric_id": {
                Type: schema.TypeString,
                Optional: true,
            },

            "applicable_schedule": {
                Type: schema.TypeList,
                Computed: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "id": {
                            Type: schema.TypeString,
                            Computed: true,
                        },
                        "lab_vms_shutdown": {
                            Type: schema.TypeList,
                            Computed: true,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "created_date": {
                                        Type: schema.TypeString,
                                        Computed: true,
                                    },
                                    "daily_recurrence": {
                                        Type: schema.TypeList,
                                        Computed: true,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "time": {
                                                    Type: schema.TypeString,
                                                    Computed: true,
                                                },
                                            },
                                        },
                                    },
                                    "hourly_recurrence": {
                                        Type: schema.TypeList,
                                        Computed: true,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "minute": {
                                                    Type: schema.TypeInt,
                                                    Computed: true,
                                                },
                                            },
                                        },
                                    },
                                    "id": {
                                        Type: schema.TypeString,
                                        Computed: true,
                                    },
                                    "location": azure.SchemaLocation(),
                                    "name": {
                                        Type: schema.TypeString,
                                        Computed: true,
                                    },
                                    "notification_settings": {
                                        Type: schema.TypeList,
                                        Computed: true,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "email_recipient": {
                                                    Type: schema.TypeString,
                                                    Computed: true,
                                                },
                                                "notification_locale": {
                                                    Type: schema.TypeString,
                                                    Computed: true,
                                                },
                                                "status": {
                                                    Type: schema.TypeString,
                                                    Computed: true,
                                                },
                                                "time_in_minutes": {
                                                    Type: schema.TypeInt,
                                                    Computed: true,
                                                },
                                                "webhook_url": {
                                                    Type: schema.TypeString,
                                                    Computed: true,
                                                },
                                            },
                                        },
                                    },
                                    "provisioning_state": {
                                        Type: schema.TypeString,
                                        Computed: true,
                                    },
                                    "status": {
                                        Type: schema.TypeString,
                                        Computed: true,
                                    },
                                    "tags": tags.Schema(),
                                    "target_resource_id": {
                                        Type: schema.TypeString,
                                        Computed: true,
                                    },
                                    "task_type": {
                                        Type: schema.TypeString,
                                        Computed: true,
                                    },
                                    "time_zone_id": {
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
                                    "weekly_recurrence": {
                                        Type: schema.TypeList,
                                        Computed: true,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "time": {
                                                    Type: schema.TypeString,
                                                    Computed: true,
                                                },
                                                "weekdays": {
                                                    Type: schema.TypeList,
                                                    Computed: true,
                                                    Elem: &schema.Schema{
                                                        Type: schema.TypeString,
                                                    },
                                                },
                                            },
                                        },
                                    },
                                },
                            },
                        },
                        "lab_vms_startup": {
                            Type: schema.TypeList,
                            Computed: true,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "created_date": {
                                        Type: schema.TypeString,
                                        Computed: true,
                                    },
                                    "daily_recurrence": {
                                        Type: schema.TypeList,
                                        Computed: true,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "time": {
                                                    Type: schema.TypeString,
                                                    Computed: true,
                                                },
                                            },
                                        },
                                    },
                                    "hourly_recurrence": {
                                        Type: schema.TypeList,
                                        Computed: true,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "minute": {
                                                    Type: schema.TypeInt,
                                                    Computed: true,
                                                },
                                            },
                                        },
                                    },
                                    "id": {
                                        Type: schema.TypeString,
                                        Computed: true,
                                    },
                                    "location": azure.SchemaLocation(),
                                    "name": {
                                        Type: schema.TypeString,
                                        Computed: true,
                                    },
                                    "notification_settings": {
                                        Type: schema.TypeList,
                                        Computed: true,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "email_recipient": {
                                                    Type: schema.TypeString,
                                                    Computed: true,
                                                },
                                                "notification_locale": {
                                                    Type: schema.TypeString,
                                                    Computed: true,
                                                },
                                                "status": {
                                                    Type: schema.TypeString,
                                                    Computed: true,
                                                },
                                                "time_in_minutes": {
                                                    Type: schema.TypeInt,
                                                    Computed: true,
                                                },
                                                "webhook_url": {
                                                    Type: schema.TypeString,
                                                    Computed: true,
                                                },
                                            },
                                        },
                                    },
                                    "provisioning_state": {
                                        Type: schema.TypeString,
                                        Computed: true,
                                    },
                                    "status": {
                                        Type: schema.TypeString,
                                        Computed: true,
                                    },
                                    "tags": tags.Schema(),
                                    "target_resource_id": {
                                        Type: schema.TypeString,
                                        Computed: true,
                                    },
                                    "task_type": {
                                        Type: schema.TypeString,
                                        Computed: true,
                                    },
                                    "time_zone_id": {
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
                                    "weekly_recurrence": {
                                        Type: schema.TypeList,
                                        Computed: true,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "time": {
                                                    Type: schema.TypeString,
                                                    Computed: true,
                                                },
                                                "weekdays": {
                                                    Type: schema.TypeList,
                                                    Computed: true,
                                                    Elem: &schema.Schema{
                                                        Type: schema.TypeString,
                                                    },
                                                },
                                            },
                                        },
                                    },
                                },
                            },
                        },
                        "location": azure.SchemaLocation(),
                        "name": {
                            Type: schema.TypeString,
                            Computed: true,
                        },
                        "tags": tags.Schema(),
                        "type": {
                            Type: schema.TypeString,
                            Computed: true,
                        },
                    },
                },
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

func resourceArmServiceFabricCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).serviceFabricsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    labName := d.Get("lab_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, labName, name, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Service Fabric %q (Lab Name %q / Resource Group %q): %+v", name, labName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_service_fabric", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    environmentId := d.Get("environment_id").(string)
    externalServiceFabricId := d.Get("external_service_fabric_id").(string)
    t := d.Get("tags").(map[string]interface{})

    serviceFabric := devtestlab.ServiceFabricFragment{
        Location: utils.String(location),
        ServiceFabricPropertiesFragment: &devtestlab.ServiceFabricPropertiesFragment{
            EnvironmentID: utils.String(environmentId),
            ExternalServiceFabricID: utils.String(externalServiceFabricId),
        },
        Tags: tags.Expand(t),
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, labName, name, name, serviceFabric)
    if err != nil {
        return fmt.Errorf("Error creating Service Fabric %q (Lab Name %q / Resource Group %q): %+v", name, labName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Service Fabric %q (Lab Name %q / Resource Group %q): %+v", name, labName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, labName, name, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Service Fabric %q (Lab Name %q / Resource Group %q): %+v", name, labName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Service Fabric %q (Lab Name %q / Resource Group %q) ID", name, labName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmServiceFabricRead(d, meta)
}

func resourceArmServiceFabricRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).serviceFabricsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    labName := id.Path["labs"]
    name := id.Path["users"]
    name := id.Path["servicefabrics"]

    resp, err := client.Get(ctx, resourceGroup, labName, name, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Service Fabric %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Service Fabric %q (Lab Name %q / Resource Group %q): %+v", name, labName, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if serviceFabricPropertiesFragment := resp.ServiceFabricPropertiesFragment; serviceFabricPropertiesFragment != nil {
        if err := d.Set("applicable_schedule", flattenArmServiceFabricApplicableSchedule(serviceFabricPropertiesFragment.ApplicableSchedule)); err != nil {
            return fmt.Errorf("Error setting `applicable_schedule`: %+v", err)
        }
        d.Set("environment_id", serviceFabricPropertiesFragment.EnvironmentID)
        d.Set("external_service_fabric_id", serviceFabricPropertiesFragment.ExternalServiceFabricID)
        d.Set("provisioning_state", serviceFabricPropertiesFragment.ProvisioningState)
        d.Set("unique_identifier", serviceFabricPropertiesFragment.UniqueIdentifier)
    }
    d.Set("lab_name", labName)
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}

func resourceArmServiceFabricUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).serviceFabricsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    environmentId := d.Get("environment_id").(string)
    externalServiceFabricId := d.Get("external_service_fabric_id").(string)
    labName := d.Get("lab_name").(string)
    t := d.Get("tags").(map[string]interface{})

    serviceFabric := devtestlab.ServiceFabricFragment{
        ServiceFabricPropertiesFragment: &devtestlab.ServiceFabricPropertiesFragment{
            EnvironmentID: utils.String(environmentId),
            ExternalServiceFabricID: utils.String(externalServiceFabricId),
        },
        Tags: tags.Expand(t),
    }


    if _, err := client.Update(ctx, resourceGroup, labName, name, name, serviceFabric); err != nil {
        return fmt.Errorf("Error updating Service Fabric %q (Lab Name %q / Resource Group %q): %+v", name, labName, resourceGroup, err)
    }

    return resourceArmServiceFabricRead(d, meta)
}

func resourceArmServiceFabricDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).serviceFabricsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    labName := id.Path["labs"]
    name := id.Path["users"]
    name := id.Path["servicefabrics"]

    future, err := client.Delete(ctx, resourceGroup, labName, name, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Service Fabric %q (Lab Name %q / Resource Group %q): %+v", name, labName, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Service Fabric %q (Lab Name %q / Resource Group %q): %+v", name, labName, resourceGroup, err)
        }
    }

    return nil
}


func flattenArmServiceFabricApplicableSchedule(input *devtestlab.ApplicableSchedule) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if id := input.ID; id != nil {
        result["id"] = *id
    }
    if name := input.Name; name != nil {
        result["name"] = *name
    }
    if location := input.Location; location != nil {
        result["location"] = azure.NormalizeLocation(*location)
    }
    if applicableScheduleProperties := input.ApplicableScheduleProperties; applicableScheduleProperties != nil {
        result["lab_vms_shutdown"] = flattenArmServiceFabricSchedule(applicableScheduleProperties.LabVmsShutdown)
        result["lab_vms_startup"] = flattenArmServiceFabricSchedule(applicableScheduleProperties.LabVmsStartup)
    }
    if type := input.Type; type != nil {
        result["type"] = *type
    }

    return []interface{}{result}
}

func flattenArmServiceFabricSchedule(input *devtestlab.Schedule) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if id := input.ID; id != nil {
        result["id"] = *id
    }
    if name := input.Name; name != nil {
        result["name"] = *name
    }
    if location := input.Location; location != nil {
        result["location"] = azure.NormalizeLocation(*location)
    }
    if scheduleProperties := input.ScheduleProperties; scheduleProperties != nil {
        if createdDate := scheduleProperties.CreatedDate; createdDate != nil {
            result["created_date"] = (*createdDate).String()
        }
        result["daily_recurrence"] = flattenArmServiceFabricDayDetails(scheduleProperties.DailyRecurrence)
        result["hourly_recurrence"] = flattenArmServiceFabricHourDetails(scheduleProperties.HourlyRecurrence)
        result["notification_settings"] = flattenArmServiceFabricNotificationSettings(scheduleProperties.NotificationSettings)
        if provisioningState := scheduleProperties.ProvisioningState; provisioningState != nil {
            result["provisioning_state"] = *provisioningState
        }
        result["status"] = string(scheduleProperties.Status)
        if targetResourceId := scheduleProperties.TargetResourceID; targetResourceId != nil {
            result["target_resource_id"] = *targetResourceId
        }
        if taskType := scheduleProperties.TaskType; taskType != nil {
            result["task_type"] = *taskType
        }
        if timeZoneId := scheduleProperties.TimeZoneID; timeZoneId != nil {
            result["time_zone_id"] = *timeZoneId
        }
        if uniqueIdentifier := scheduleProperties.UniqueIdentifier; uniqueIdentifier != nil {
            result["unique_identifier"] = *uniqueIdentifier
        }
        result["weekly_recurrence"] = flattenArmServiceFabricWeekDetails(scheduleProperties.WeeklyRecurrence)
    }
    if type := input.Type; type != nil {
        result["type"] = *type
    }

    return []interface{}{result}
}

func flattenArmServiceFabricDayDetails(input *devtestlab.DayDetails) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if time := input.Time; time != nil {
        result["time"] = *time
    }

    return []interface{}{result}
}

func flattenArmServiceFabricHourDetails(input *devtestlab.HourDetails) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if minute := input.Minute; minute != nil {
        result["minute"] = int(*minute)
    }

    return []interface{}{result}
}

func flattenArmServiceFabricNotificationSettings(input *devtestlab.NotificationSettings) []interface{} {
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

func flattenArmServiceFabricWeekDetails(input *devtestlab.WeekDetails) []interface{} {
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
