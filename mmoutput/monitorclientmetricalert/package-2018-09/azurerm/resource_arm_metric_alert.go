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



func resourceArmMetricAlert() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmMetricAlertCreate,
        Read: resourceArmMetricAlertRead,
        Update: resourceArmMetricAlertUpdate,
        Delete: resourceArmMetricAlertDelete,

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

            "criteria": {
                Type: schema.TypeList,
                Required: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "additional_properties": {
                            Type: schema.TypeMap,
                            Optional: true,
                            Elem: &schema.Schema{Type: schema.TypeString},
                        },
                    },
                },
            },

            "description": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "enabled": {
                Type: schema.TypeBool,
                Required: true,
            },

            "evaluation_frequency": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validate.NoEmptyStrings,
                ValidateFunc: validateIso8601Duration(),
            },

            "severity": {
                Type: schema.TypeInt,
                Required: true,
            },

            "window_size": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validate.NoEmptyStrings,
                ValidateFunc: validateIso8601Duration(),
            },

            "actions": {
                Type: schema.TypeList,
                Optional: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "action_group_id": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "webhook_properties": {
                            Type: schema.TypeMap,
                            Optional: true,
                            Elem: &schema.Schema{Type: schema.TypeString},
                        },
                    },
                },
            },

            "auto_mitigate": {
                Type: schema.TypeBool,
                Optional: true,
            },

            "scopes": {
                Type: schema.TypeList,
                Optional: true,
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
            },

            "target_resource_region": {
                Type: schema.TypeString,
                Optional: true,
            },

            "target_resource_type": {
                Type: schema.TypeString,
                Optional: true,
            },

            "last_updated_time": {
                Type: schema.TypeString,
                Computed: true,
            },

            "type": {
                Type: schema.TypeString,
                Computed: true,
            },

            "tags": tags.Schema(),
        },
    }
}

func resourceArmMetricAlertCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).metricAlertsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Metric Alert %q (Resource Group %q): %+v", name, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_metric_alert", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    actions := d.Get("actions").([]interface{})
    autoMitigate := d.Get("auto_mitigate").(bool)
    criteria := d.Get("criteria").([]interface{})
    description := d.Get("description").(string)
    enabled := d.Get("enabled").(bool)
    evaluationFrequency := d.Get("evaluation_frequency").(string)
    scopes := d.Get("scopes").([]interface{})
    severity := d.Get("severity").(int)
    targetResourceRegion := d.Get("target_resource_region").(string)
    targetResourceType := d.Get("target_resource_type").(string)
    windowSize := d.Get("window_size").(string)
    t := d.Get("tags").(map[string]interface{})

    parameters := monitorclient.MetricAlertResourcePatch{
        Location: utils.String(location),
        MetricAlertProperties: &monitorclient.MetricAlertProperties{
            Actions: expandArmMetricAlertMetricAlertAction(actions),
            AutoMitigate: utils.Bool(autoMitigate),
            Criteria: expandArmMetricAlertMetricAlertCriteria(criteria),
            Description: utils.String(description),
            Enabled: utils.Bool(enabled),
            EvaluationFrequency: utils.String(evaluationFrequency),
            Scopes: utils.ExpandStringSlice(scopes),
            Severity: utils.Int(severity),
            TargetResourceRegion: utils.String(targetResourceRegion),
            TargetResourceType: utils.String(targetResourceType),
            WindowSize: utils.String(windowSize),
        },
        Tags: tags.Expand(t),
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroup, name, parameters); err != nil {
        return fmt.Errorf("Error creating Metric Alert %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Metric Alert %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Metric Alert %q (Resource Group %q) ID", name, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmMetricAlertRead(d, meta)
}

func resourceArmMetricAlertRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).metricAlertsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["metricAlerts"]

    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Metric Alert %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Metric Alert %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if metricAlertProperties := resp.MetricAlertProperties; metricAlertProperties != nil {
        if err := d.Set("actions", flattenArmMetricAlertMetricAlertAction(metricAlertProperties.Actions)); err != nil {
            return fmt.Errorf("Error setting `actions`: %+v", err)
        }
        d.Set("auto_mitigate", metricAlertProperties.AutoMitigate)
        if err := d.Set("criteria", flattenArmMetricAlertMetricAlertCriteria(metricAlertProperties.Criteria)); err != nil {
            return fmt.Errorf("Error setting `criteria`: %+v", err)
        }
        d.Set("description", metricAlertProperties.Description)
        d.Set("enabled", metricAlertProperties.Enabled)
        d.Set("evaluation_frequency", metricAlertProperties.EvaluationFrequency)
        d.Set("last_updated_time", (metricAlertProperties.LastUpdatedTime).String())
        d.Set("scopes", utils.FlattenStringSlice(metricAlertProperties.Scopes))
        d.Set("severity", metricAlertProperties.Severity)
        d.Set("target_resource_region", metricAlertProperties.TargetResourceRegion)
        d.Set("target_resource_type", metricAlertProperties.TargetResourceType)
        d.Set("window_size", metricAlertProperties.WindowSize)
    }
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}

func resourceArmMetricAlertUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).metricAlertsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    actions := d.Get("actions").([]interface{})
    autoMitigate := d.Get("auto_mitigate").(bool)
    criteria := d.Get("criteria").([]interface{})
    description := d.Get("description").(string)
    enabled := d.Get("enabled").(bool)
    evaluationFrequency := d.Get("evaluation_frequency").(string)
    scopes := d.Get("scopes").([]interface{})
    severity := d.Get("severity").(int)
    targetResourceRegion := d.Get("target_resource_region").(string)
    targetResourceType := d.Get("target_resource_type").(string)
    windowSize := d.Get("window_size").(string)
    t := d.Get("tags").(map[string]interface{})

    parameters := monitorclient.MetricAlertResourcePatch{
        MetricAlertProperties: &monitorclient.MetricAlertProperties{
            Actions: expandArmMetricAlertMetricAlertAction(actions),
            AutoMitigate: utils.Bool(autoMitigate),
            Criteria: expandArmMetricAlertMetricAlertCriteria(criteria),
            Description: utils.String(description),
            Enabled: utils.Bool(enabled),
            EvaluationFrequency: utils.String(evaluationFrequency),
            Scopes: utils.ExpandStringSlice(scopes),
            Severity: utils.Int(severity),
            TargetResourceRegion: utils.String(targetResourceRegion),
            TargetResourceType: utils.String(targetResourceType),
            WindowSize: utils.String(windowSize),
        },
        Tags: tags.Expand(t),
    }


    if _, err := client.Update(ctx, resourceGroup, name, parameters); err != nil {
        return fmt.Errorf("Error updating Metric Alert %q (Resource Group %q): %+v", name, resourceGroup, err)
    }

    return resourceArmMetricAlertRead(d, meta)
}

func resourceArmMetricAlertDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).metricAlertsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["metricAlerts"]

    if _, err := client.Delete(ctx, resourceGroup, name); err != nil {
        return fmt.Errorf("Error deleting Metric Alert %q (Resource Group %q): %+v", name, resourceGroup, err)
    }

    return nil
}

func expandArmMetricAlertMetricAlertAction(input []interface{}) *[]monitorclient.MetricAlertAction {
    results := make([]monitorclient.MetricAlertAction, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        actionGroupId := v["action_group_id"].(string)
        webhookProperties := v["webhook_properties"].(map[string]interface{})

        result := monitorclient.MetricAlertAction{
            ActionGroupID: utils.String(actionGroupId),
            WebhookProperties: utils.ExpandKeyValuePairs(webhookProperties),
        }

        results = append(results, result)
    }
    return &results
}

func expandArmMetricAlertMetricAlertCriteria(input []interface{}) *monitorclient.MetricAlertCriteria {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    additionalProperties := v["additional_properties"].(map[string]interface{})

    result := monitorclient.MetricAlertCriteria{
        AdditionalProperties: utils.ExpandKeyValuePairs(additionalProperties),
    }
    return &result
}


func flattenArmMetricAlertMetricAlertAction(input *[]monitorclient.MetricAlertAction) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        if actionGroupId := item.ActionGroupID; actionGroupId != nil {
            v["action_group_id"] = *actionGroupId
        }
        v["webhook_properties"] = utils.FlattenKeyValuePairs(item.WebhookProperties)

        results = append(results, v)
    }

    return results
}

func flattenArmMetricAlertMetricAlertCriteria(input *monitorclient.MetricAlertCriteria) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    result["additional_properties"] = utils.FlattenKeyValuePairs(input.AdditionalProperties)

    return []interface{}{result}
}
