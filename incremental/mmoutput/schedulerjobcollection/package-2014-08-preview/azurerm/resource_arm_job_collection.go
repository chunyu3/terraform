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



func resourceArmJobCollection() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmJobCollectionCreateUpdate,
        Read: resourceArmJobCollectionRead,
        Update: resourceArmJobCollectionCreateUpdate,
        Delete: resourceArmJobCollectionDelete,

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
                Optional: true,
                ForceNew: true,
            },

            "location": azure.SchemaLocation(),

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "quota": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "max_job_count": {
                            Type: schema.TypeInt,
                            Optional: true,
                        },
                        "max_job_occurrence": {
                            Type: schema.TypeInt,
                            Optional: true,
                        },
                        "max_recurrence": {
                            Type: schema.TypeList,
                            Optional: true,
                            MaxItems: 1,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "frequency": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                        ValidateFunc: validation.StringInSlice([]string{
                                            string(scheduler.Minute),
                                            string(scheduler.Hour),
                                            string(scheduler.Day),
                                            string(scheduler.Week),
                                            string(scheduler.Month),
                                        }, false),
                                        Default: string(scheduler.Minute),
                                    },
                                    "interval": {
                                        Type: schema.TypeInt,
                                        Optional: true,
                                    },
                                },
                            },
                        },
                    },
                },
            },

            "sku": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "name": {
                            Type: schema.TypeString,
                            Optional: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(scheduler.Standard),
                                string(scheduler.Free),
                                string(scheduler.Premium),
                            }, false),
                            Default: string(scheduler.Standard),
                        },
                    },
                },
            },

            "state": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(scheduler.Enabled),
                    string(scheduler.Disabled),
                    string(scheduler.Suspended),
                    string(scheduler.Deleted),
                }, false),
                Default: string(scheduler.Enabled),
            },

            "type": {
                Type: schema.TypeString,
                Computed: true,
            },

            "tags": tags.Schema(),
        },
    }
}

func resourceArmJobCollectionCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).jobCollectionsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Job Collection %q (Resource Group %q): %+v", name, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_job_collection", *existing.ID)
        }
    }

    name := d.Get("name").(string)
    location := azure.NormalizeLocation(d.Get("location").(string))
    quota := d.Get("quota").([]interface{})
    sku := d.Get("sku").([]interface{})
    state := d.Get("state").(string)
    t := d.Get("tags").(map[string]interface{})

    jobCollection := scheduler.JobCollectionDefinition{
        Location: utils.String(location),
        Name: utils.String(name),
        JobCollectionProperties: &scheduler.JobCollectionProperties{
            Quota: expandArmJobCollectionJobCollectionQuota(quota),
            Sku: expandArmJobCollectionSku(sku),
            State: scheduler.JobCollectionState(state),
        },
        Tags: tags.Expand(t),
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroup, name, jobCollection); err != nil {
        return fmt.Errorf("Error creating Job Collection %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Job Collection %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Job Collection %q (Resource Group %q) ID", name, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmJobCollectionRead(d, meta)
}

func resourceArmJobCollectionRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).jobCollectionsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["jobCollections"]

    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Job Collection %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Job Collection %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if jobCollectionProperties := resp.JobCollectionProperties; jobCollectionProperties != nil {
        if err := d.Set("quota", flattenArmJobCollectionJobCollectionQuota(jobCollectionProperties.Quota)); err != nil {
            return fmt.Errorf("Error setting `quota`: %+v", err)
        }
        if err := d.Set("sku", flattenArmJobCollectionSku(jobCollectionProperties.Sku)); err != nil {
            return fmt.Errorf("Error setting `sku`: %+v", err)
        }
        d.Set("state", string(jobCollectionProperties.State))
    }
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}


func resourceArmJobCollectionDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).jobCollectionsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["jobCollections"]

    if _, err := client.Delete(ctx, resourceGroup, name); err != nil {
        return fmt.Errorf("Error deleting Job Collection %q (Resource Group %q): %+v", name, resourceGroup, err)
    }

    return nil
}

func expandArmJobCollectionJobCollectionQuota(input []interface{}) *scheduler.JobCollectionQuota {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    maxJobCount := v["max_job_count"].(int)
    maxJobOccurrence := v["max_job_occurrence"].(int)
    maxRecurrence := v["max_recurrence"].([]interface{})

    result := scheduler.JobCollectionQuota{
        MaxJobCount: utils.Int(maxJobCount),
        MaxJobOccurrence: utils.Int(maxJobOccurrence),
        MaxRecurrence: expandArmJobCollectionJobMaxRecurrence(maxRecurrence),
    }
    return &result
}

func expandArmJobCollectionSku(input []interface{}) *scheduler.Sku {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    name := v["name"].(string)

    result := scheduler.Sku{
        Name: scheduler.SkuDefinition(name),
    }
    return &result
}

func expandArmJobCollectionJobMaxRecurrence(input []interface{}) *scheduler.JobMaxRecurrence {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    frequency := v["frequency"].(string)
    interval := v["interval"].(int)

    result := scheduler.JobMaxRecurrence{
        Frequency: scheduler.RecurrenceFrequency(frequency),
        Interval: utils.Int(interval),
    }
    return &result
}


func flattenArmJobCollectionJobCollectionQuota(input *scheduler.JobCollectionQuota) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if maxJobCount := input.MaxJobCount; maxJobCount != nil {
        result["max_job_count"] = *maxJobCount
    }
    if maxJobOccurrence := input.MaxJobOccurrence; maxJobOccurrence != nil {
        result["max_job_occurrence"] = *maxJobOccurrence
    }
    result["max_recurrence"] = flattenArmJobCollectionJobMaxRecurrence(input.MaxRecurrence)

    return []interface{}{result}
}

func flattenArmJobCollectionSku(input *scheduler.Sku) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    result["name"] = string(input.Name)

    return []interface{}{result}
}

func flattenArmJobCollectionJobMaxRecurrence(input *scheduler.JobMaxRecurrence) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    result["frequency"] = string(input.Frequency)
    if interval := input.Interval; interval != nil {
        result["interval"] = *interval
    }

    return []interface{}{result}
}
