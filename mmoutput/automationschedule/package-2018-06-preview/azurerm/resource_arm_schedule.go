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
            "automation_account_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "schedule_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "description": {
                Type: schema.TypeString,
                Optional: true,
            },

            "is_enabled": {
                Type: schema.TypeBool,
                Optional: true,
            },

            "name": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
            },

            "advanced_schedule": {
                Type: schema.TypeList,
                Computed: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "month_days": {
                            Type: schema.TypeList,
                            Computed: true,
                            Elem: &schema.Schema{
                                Type: schema.TypeInt,
                            },
                        },
                        "monthly_occurrences": {
                            Type: schema.TypeList,
                            Computed: true,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "day": {
                                        Type: schema.TypeString,
                                        Computed: true,
                                    },
                                    "occurrence": {
                                        Type: schema.TypeInt,
                                        Computed: true,
                                    },
                                },
                            },
                        },
                        "week_days": {
                            Type: schema.TypeList,
                            Computed: true,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                    },
                },
            },

            "creation_time": {
                Type: schema.TypeString,
                Computed: true,
            },

            "expiry_time": {
                Type: schema.TypeString,
                Computed: true,
            },

            "expiry_time_offset_minutes": {
                Type: schema.TypeFloat,
                Computed: true,
            },

            "frequency": {
                Type: schema.TypeString,
                Computed: true,
            },

            "id": {
                Type: schema.TypeString,
                Computed: true,
            },

            "interval": {
                Type: schema.TypeInt,
                Computed: true,
            },

            "last_modified_time": {
                Type: schema.TypeString,
                Computed: true,
            },

            "next_run": {
                Type: schema.TypeString,
                Computed: true,
            },

            "next_run_offset_minutes": {
                Type: schema.TypeFloat,
                Computed: true,
            },

            "start_time": {
                Type: schema.TypeString,
                Computed: true,
            },

            "start_time_offset_minutes": {
                Type: schema.TypeFloat,
                Computed: true,
            },

            "time_zone": {
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

func resourceArmScheduleCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).scheduleClient
    ctx, cancel := timeouts.ForCreate(meta.(*ArmClient).StopContext, d)
    defer cancel()

    resourceGroupName := d.Get("resource_group").(string)
    automationAccountName := d.Get("automation_account_name").(string)
    name := d.Get("schedule_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroupName, automationAccountName, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Schedule (Schedule Name %q / Automation Account Name %q / Resource Group %q): %+v", name, automationAccountName, resourceGroupName, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_schedule", *existing.ID)
        }
    }

    description := d.Get("description").(string)
    isEnabled := d.Get("is_enabled").(bool)
    name := d.Get("name").(string)

    parameters := automation.ScheduleUpdateParameters{
        Name: utils.String(name),
        ScheduleUpdateProperties: &automation.ScheduleUpdateProperties{
            Description: utils.String(description),
            IsEnabled: utils.Bool(isEnabled),
        },
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroupName, automationAccountName, name, parameters); err != nil {
        return fmt.Errorf("Error creating Schedule (Schedule Name %q / Automation Account Name %q / Resource Group %q): %+v", name, automationAccountName, resourceGroupName, err)
    }


    resp, err := client.Get(ctx, resourceGroupName, automationAccountName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Schedule (Schedule Name %q / Automation Account Name %q / Resource Group %q): %+v", name, automationAccountName, resourceGroupName, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Schedule (Schedule Name %q / Automation Account Name %q / Resource Group %q) ID", name, automationAccountName, resourceGroupName)
    }
    d.SetId(*resp.ID)

    return resourceArmScheduleRead(d, meta)
}

func resourceArmScheduleRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).scheduleClient
    ctx, cancel := timeouts.ForRead(meta.(*ArmClient).StopContext, d)
    defer cancel()

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroupName := id.ResourceGroup
    automationAccountName := id.Path["automationAccounts"]
    name := id.Path["schedules"]

    resp, err := client.Get(ctx, resourceGroupName, automationAccountName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Schedule %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Schedule (Schedule Name %q / Automation Account Name %q / Resource Group %q): %+v", name, automationAccountName, resourceGroupName, err)
    }


    d.Set("resource_group", resourceGroupName)
    if scheduleUpdateProperties := resp.ScheduleUpdateProperties; scheduleUpdateProperties != nil {
        if err := d.Set("advanced_schedule", flattenArmScheduleAdvancedSchedule(scheduleUpdateProperties.AdvancedSchedule)); err != nil {
            return fmt.Errorf("Error setting `advanced_schedule`: %+v", err)
        }
        d.Set("creation_time", (scheduleUpdateProperties.CreationTime).String())
        d.Set("description", scheduleUpdateProperties.Description)
        d.Set("expiry_time", (scheduleUpdateProperties.ExpiryTime).String())
        d.Set("expiry_time_offset_minutes", scheduleUpdateProperties.ExpiryTimeOffsetMinutes)
        d.Set("frequency", string(scheduleUpdateProperties.Frequency))
        d.Set("interval", scheduleUpdateProperties.Interval)
        d.Set("is_enabled", scheduleUpdateProperties.IsEnabled)
        d.Set("last_modified_time", (scheduleUpdateProperties.LastModifiedTime).String())
        d.Set("next_run", (scheduleUpdateProperties.NextRun).String())
        d.Set("next_run_offset_minutes", scheduleUpdateProperties.NextRunOffsetMinutes)
        d.Set("start_time", (scheduleUpdateProperties.StartTime).String())
        d.Set("start_time_offset_minutes", scheduleUpdateProperties.StartTimeOffsetMinutes)
        d.Set("time_zone", scheduleUpdateProperties.TimeZone)
    }
    d.Set("automation_account_name", automationAccountName)
    d.Set("id", resp.ID)
    d.Set("name", resp.Name)
    d.Set("schedule_name", name)
    d.Set("type", resp.Type)

    return nil
}

func resourceArmScheduleUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).scheduleClient
    ctx, cancel := timeouts.ForUpdate(meta.(*ArmClient).StopContext, d)
    defer cancel()

      resourceGroupName := d.Get("resource_group").(string)
    automationAccountName := d.Get("automation_account_name").(string)
    description := d.Get("description").(string)
    isEnabled := d.Get("is_enabled").(bool)
    name := d.Get("name").(string)
    name := d.Get("schedule_name").(string)

    parameters := automation.ScheduleUpdateParameters{
        Name: utils.String(name),
        ScheduleUpdateProperties: &automation.ScheduleUpdateProperties{
            Description: utils.String(description),
            IsEnabled: utils.Bool(isEnabled),
        },
    }


    if _, err := client.Update(ctx, resourceGroupName, automationAccountName, name, parameters); err != nil {
        return fmt.Errorf("Error updating Schedule (Schedule Name %q / Automation Account Name %q / Resource Group %q): %+v", name, automationAccountName, resourceGroupName, err)
    }

    return resourceArmScheduleRead(d, meta)
}

func resourceArmScheduleDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).scheduleClient
    ctx, cancel := timeouts.ForDelete(meta.(*ArmClient).StopContext, d)
    defer cancel()


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroupName := id.ResourceGroup
    automationAccountName := id.Path["automationAccounts"]
    name := id.Path["schedules"]

    if _, err := client.Delete(ctx, resourceGroupName, automationAccountName, name); err != nil {
        return fmt.Errorf("Error deleting Schedule (Schedule Name %q / Automation Account Name %q / Resource Group %q): %+v", name, automationAccountName, resourceGroupName, err)
    }

    return nil
}


func flattenArmScheduleAdvancedSchedule(input *automation.AdvancedSchedule) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    result["month_days"] = utils.FlattenInteger32Slice(input.MonthDays)
    result["monthly_occurrences"] = flattenArmScheduleAdvancedScheduleMonthlyOccurrence(input.MonthlyOccurrences)
    result["week_days"] = utils.FlattenStringSlice(input.WeekDays)

    return []interface{}{result}
}

func flattenArmScheduleAdvancedScheduleMonthlyOccurrence(input *[]automation.AdvancedScheduleMonthlyOccurrence) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        v["day"] = string(item.Day)
        if occurrence := item.Occurrence; occurrence != nil {
            v["occurrence"] = int(*occurrence)
        }

        results = append(results, v)
    }

    return results
}
