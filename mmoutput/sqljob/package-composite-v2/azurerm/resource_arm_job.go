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



func resourceArmJob() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmJobCreateUpdate,
        Read: resourceArmJobRead,
        Update: resourceArmJobCreateUpdate,
        Delete: resourceArmJobDelete,

        Importer: &schema.ResourceImporter{
            State: schema.ImportStatePassthrough,
        },


        Schema: map[string]*schema.Schema{
            "name": {
                Type: schema.TypeString,
                Computed: true,
            },

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "job_agent_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "job_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "server_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "description": {
                Type: schema.TypeString,
                Optional: true,
            },

            "schedule": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "enabled": {
                            Type: schema.TypeBool,
                            Optional: true,
                        },
                        "end_time": {
                            Type: schema.TypeString,
                            Optional: true,
                            ValidateFunc: validateRFC3339Date,
                        },
                        "interval": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "start_time": {
                            Type: schema.TypeString,
                            Optional: true,
                            ValidateFunc: validateRFC3339Date,
                        },
                        "type": {
                            Type: schema.TypeString,
                            Optional: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(sql.Once),
                                string(sql.Recurring),
                            }, false),
                            Default: string(sql.Once),
                        },
                    },
                },
            },

            "type": {
                Type: schema.TypeString,
                Computed: true,
            },

            "version": {
                Type: schema.TypeInt,
                Computed: true,
            },
        },
    }
}

func resourceArmJobCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).jobsClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    jobAgentName := d.Get("job_agent_name").(string)
    jobName := d.Get("job_name").(string)
    serverName := d.Get("server_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, serverName, jobAgentName, jobName)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Job (Job Name %q / Job Agent Name %q / Server Name %q / Resource Group %q): %+v", jobName, jobAgentName, serverName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_job", *existing.ID)
        }
    }

    description := d.Get("description").(string)
    schedule := d.Get("schedule").([]interface{})

    parameters := sql.Job{
        JobProperties: &sql.JobProperties{
            Description: utils.String(description),
            Schedule: expandArmJobJobSchedule(schedule),
        },
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroup, serverName, jobAgentName, jobName, parameters); err != nil {
        return fmt.Errorf("Error creating Job (Job Name %q / Job Agent Name %q / Server Name %q / Resource Group %q): %+v", jobName, jobAgentName, serverName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, serverName, jobAgentName, jobName)
    if err != nil {
        return fmt.Errorf("Error retrieving Job (Job Name %q / Job Agent Name %q / Server Name %q / Resource Group %q): %+v", jobName, jobAgentName, serverName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Job (Job Name %q / Job Agent Name %q / Server Name %q / Resource Group %q) ID", jobName, jobAgentName, serverName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmJobRead(d, meta)
}

func resourceArmJobRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).jobsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    serverName := id.Path["servers"]
    jobAgentName := id.Path["jobAgents"]
    jobName := id.Path["jobs"]

    resp, err := client.Get(ctx, resourceGroup, serverName, jobAgentName, jobName)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Job %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Job (Job Name %q / Job Agent Name %q / Server Name %q / Resource Group %q): %+v", jobName, jobAgentName, serverName, resourceGroup, err)
    }


    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if jobProperties := resp.JobProperties; jobProperties != nil {
        d.Set("description", jobProperties.Description)
        if err := d.Set("schedule", flattenArmJobJobSchedule(jobProperties.Schedule)); err != nil {
            return fmt.Errorf("Error setting `schedule`: %+v", err)
        }
        d.Set("version", int(*jobProperties.Version))
    }
    d.Set("job_agent_name", jobAgentName)
    d.Set("job_name", jobName)
    d.Set("server_name", serverName)
    d.Set("type", resp.Type)

    return nil
}


func resourceArmJobDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).jobsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    serverName := id.Path["servers"]
    jobAgentName := id.Path["jobAgents"]
    jobName := id.Path["jobs"]

    if _, err := client.Delete(ctx, resourceGroup, serverName, jobAgentName, jobName); err != nil {
        return fmt.Errorf("Error deleting Job (Job Name %q / Job Agent Name %q / Server Name %q / Resource Group %q): %+v", jobName, jobAgentName, serverName, resourceGroup, err)
    }

    return nil
}

func expandArmJobJobSchedule(input []interface{}) *sql.JobSchedule {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    startTime := v["start_time"].(string)
    endTime := v["end_time"].(string)
    type := v["type"].(string)
    enabled := v["enabled"].(bool)
    interval := v["interval"].(string)

    result := sql.JobSchedule{
        Enabled: utils.Bool(enabled),
        EndTime: convertStringToDate(endTime),
        Interval: utils.String(interval),
        StartTime: convertStringToDate(startTime),
        Type: sql.JobScheduleType(type),
    }
    return &result
}

func convertStringToDate(input interface{}) *date.Time {
  v := input.(string)

  dateTime, err := date.ParseTime(time.RFC3339, v)
  if err != nil {
      log.Printf("[ERROR] Cannot convert an invalid string to RFC3339 date %q: %+v", v, err)
      return nil
  }

  result := date.Time{
      Time: dateTime,
  }
  return &result
}


func flattenArmJobJobSchedule(input *sql.JobSchedule) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if enabled := input.Enabled; enabled != nil {
        result["enabled"] = *enabled
    }
    if endTime := input.EndTime; endTime != nil {
        result["end_time"] = (*endTime).String()
    }
    if interval := input.Interval; interval != nil {
        result["interval"] = *interval
    }
    if startTime := input.StartTime; startTime != nil {
        result["start_time"] = (*startTime).String()
    }
    result["type"] = string(input.Type)

    return []interface{}{result}
}
