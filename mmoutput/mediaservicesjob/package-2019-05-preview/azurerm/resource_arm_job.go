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
        Create: resourceArmJobCreate,
        Read: resourceArmJobRead,
        Update: resourceArmJobUpdate,
        Delete: resourceArmJobDelete,

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

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "account_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "outputs": {
                Type: schema.TypeList,
                Required: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "label": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },

            "transform_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "correlation_data": {
                Type: schema.TypeMap,
                Optional: true,
                Elem: &schema.Schema{Type: schema.TypeString},
            },

            "description": {
                Type: schema.TypeString,
                Optional: true,
            },

            "priority": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(mediaservices.Low),
                    string(mediaservices.Normal),
                    string(mediaservices.High),
                }, false),
                Default: string(mediaservices.Low),
            },

            "created": {
                Type: schema.TypeString,
                Computed: true,
            },

            "last_modified": {
                Type: schema.TypeString,
                Computed: true,
            },

            "state": {
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

func resourceArmJobCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).jobsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    accountName := d.Get("account_name").(string)
    transformName := d.Get("transform_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, accountName, transformName, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Job %q (Transform Name %q / Account Name %q / Resource Group %q): %+v", name, transformName, accountName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_job", *existing.ID)
        }
    }

    correlationData := d.Get("correlation_data").(map[string]interface{})
    description := d.Get("description").(string)
    outputs := d.Get("outputs").([]interface{})
    priority := d.Get("priority").(string)

    parameters := mediaservices.Job{
        JobProperties: &mediaservices.JobProperties{
            CorrelationData: utils.ExpandKeyValuePairs(correlationData),
            Description: utils.String(description),
            Outputs: expandArmJobJobOutput(outputs),
            Priority: mediaservices.Priority(priority),
        },
    }


    if _, err := client.Create(ctx, resourceGroup, accountName, transformName, name, parameters); err != nil {
        return fmt.Errorf("Error creating Job %q (Transform Name %q / Account Name %q / Resource Group %q): %+v", name, transformName, accountName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, accountName, transformName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Job %q (Transform Name %q / Account Name %q / Resource Group %q): %+v", name, transformName, accountName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Job %q (Transform Name %q / Account Name %q / Resource Group %q) ID", name, transformName, accountName, resourceGroup)
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
    accountName := id.Path["mediaServices"]
    transformName := id.Path["transforms"]
    name := id.Path["jobs"]

    resp, err := client.Get(ctx, resourceGroup, accountName, transformName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Job %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Job %q (Transform Name %q / Account Name %q / Resource Group %q): %+v", name, transformName, accountName, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    d.Set("account_name", accountName)
    if jobProperties := resp.JobProperties; jobProperties != nil {
        d.Set("correlation_data", utils.FlattenKeyValuePairs(jobProperties.CorrelationData))
        d.Set("created", (jobProperties.Created).String())
        d.Set("description", jobProperties.Description)
        d.Set("last_modified", (jobProperties.LastModified).String())
        if err := d.Set("outputs", flattenArmJobJobOutput(jobProperties.Outputs)); err != nil {
            return fmt.Errorf("Error setting `outputs`: %+v", err)
        }
        d.Set("priority", string(jobProperties.Priority))
        d.Set("state", string(jobProperties.State))
    }
    d.Set("transform_name", transformName)
    d.Set("type", resp.Type)

    return nil
}

func resourceArmJobUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).jobsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    accountName := d.Get("account_name").(string)
    correlationData := d.Get("correlation_data").(map[string]interface{})
    description := d.Get("description").(string)
    outputs := d.Get("outputs").([]interface{})
    priority := d.Get("priority").(string)
    transformName := d.Get("transform_name").(string)

    parameters := mediaservices.Job{
        JobProperties: &mediaservices.JobProperties{
            CorrelationData: utils.ExpandKeyValuePairs(correlationData),
            Description: utils.String(description),
            Outputs: expandArmJobJobOutput(outputs),
            Priority: mediaservices.Priority(priority),
        },
    }


    if _, err := client.Update(ctx, resourceGroup, accountName, transformName, name, parameters); err != nil {
        return fmt.Errorf("Error updating Job %q (Transform Name %q / Account Name %q / Resource Group %q): %+v", name, transformName, accountName, resourceGroup, err)
    }

    return resourceArmJobRead(d, meta)
}

func resourceArmJobDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).jobsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    accountName := id.Path["mediaServices"]
    transformName := id.Path["transforms"]
    name := id.Path["jobs"]

    if _, err := client.Delete(ctx, resourceGroup, accountName, transformName, name); err != nil {
        return fmt.Errorf("Error deleting Job %q (Transform Name %q / Account Name %q / Resource Group %q): %+v", name, transformName, accountName, resourceGroup, err)
    }

    return nil
}

func expandArmJobJobOutput(input []interface{}) *[]mediaservices.JobOutput {
    results := make([]mediaservices.JobOutput, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        label := v["label"].(string)

        result := mediaservices.JobOutput{
            Label: utils.String(label),
        }

        results = append(results, result)
    }
    return &results
}


func flattenArmJobJobOutput(input *[]mediaservices.JobOutput) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        if label := item.Label; label != nil {
            v["label"] = *label
        }

        results = append(results, v)
    }

    return results
}
