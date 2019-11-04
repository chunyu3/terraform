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



func resourceArmPrediction() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmPredictionCreateUpdate,
        Read: resourceArmPredictionRead,
        Update: resourceArmPredictionCreateUpdate,
        Delete: resourceArmPredictionDelete,

        Importer: &schema.ResourceImporter{
            State: schema.ImportStatePassthrough,
        },


        Schema: map[string]*schema.Schema{
            "name": {
                Type: schema.TypeString,
                Computed: true,
            },

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "auto_analyze": {
                Type: schema.TypeBool,
                Required: true,
            },

            "hub_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "mappings": {
                Type: schema.TypeList,
                Required: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "grade": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "reason": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "score": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                    },
                },
            },

            "negative_outcome_expression": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "positive_outcome_expression": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "prediction_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "primary_profile_type": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "scope_expression": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "score_label": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "description": {
                Type: schema.TypeMap,
                Optional: true,
                Elem: &schema.Schema{Type: schema.TypeString},
            },

            "display_name": {
                Type: schema.TypeMap,
                Optional: true,
                Elem: &schema.Schema{Type: schema.TypeString},
            },

            "grades": {
                Type: schema.TypeList,
                Optional: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "grade_name": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "max_score_threshold": {
                            Type: schema.TypeInt,
                            Optional: true,
                        },
                        "min_score_threshold": {
                            Type: schema.TypeInt,
                            Optional: true,
                        },
                    },
                },
            },

            "involved_interaction_types": {
                Type: schema.TypeList,
                Optional: true,
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
            },

            "involved_kpi_types": {
                Type: schema.TypeList,
                Optional: true,
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
            },

            "involved_relationships": {
                Type: schema.TypeList,
                Optional: true,
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
            },

            "prediction_name": {
                Type: schema.TypeString,
                Optional: true,
            },

            "provisioning_state": {
                Type: schema.TypeString,
                Computed: true,
            },

            "system_generated_entities": {
                Type: schema.TypeList,
                Computed: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "generated_interaction_types": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                        "generated_kpis": {
                            Type: schema.TypeMap,
                            Optional: true,
                            Elem: &schema.Schema{Type: schema.TypeString},
                        },
                        "generated_links": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                    },
                },
            },

            "tenant_id": {
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

func resourceArmPredictionCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).predictionsClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    hubName := d.Get("hub_name").(string)
    predictionName := d.Get("prediction_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, hubName, predictionName)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Prediction (Prediction Name %q / Hub Name %q / Resource Group %q): %+v", predictionName, hubName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_prediction", *existing.ID)
        }
    }

    autoAnalyze := d.Get("auto_analyze").(bool)
    description := d.Get("description").(map[string]interface{})
    displayName := d.Get("display_name").(map[string]interface{})
    grades := d.Get("grades").([]interface{})
    involvedInteractionTypes := d.Get("involved_interaction_types").([]interface{})
    involvedKpiTypes := d.Get("involved_kpi_types").([]interface{})
    involvedRelationships := d.Get("involved_relationships").([]interface{})
    mappings := d.Get("mappings").([]interface{})
    negativeOutcomeExpression := d.Get("negative_outcome_expression").(string)
    positiveOutcomeExpression := d.Get("positive_outcome_expression").(string)
    predictionName := d.Get("prediction_name").(string)
    primaryProfileType := d.Get("primary_profile_type").(string)
    scopeExpression := d.Get("scope_expression").(string)
    scoreLabel := d.Get("score_label").(string)

    parameters := customerinsights.PredictionResourceFormat{
        Prediction: &customerinsights.Prediction{
            AutoAnalyze: utils.Bool(autoAnalyze),
            Description: utils.ExpandKeyValuePairs(description),
            DisplayName: utils.ExpandKeyValuePairs(displayName),
            Grades: expandArmPredictionPrediction_gradesItem(grades),
            InvolvedInteractionTypes: utils.ExpandStringSlice(involvedInteractionTypes),
            InvolvedKpiTypes: utils.ExpandStringSlice(involvedKpiTypes),
            InvolvedRelationships: utils.ExpandStringSlice(involvedRelationships),
            Mappings: expandArmPredictionPrediction_mappings(mappings),
            NegativeOutcomeExpression: utils.String(negativeOutcomeExpression),
            PositiveOutcomeExpression: utils.String(positiveOutcomeExpression),
            PredictionName: utils.String(predictionName),
            PrimaryProfileType: utils.String(primaryProfileType),
            ScopeExpression: utils.String(scopeExpression),
            ScoreLabel: utils.String(scoreLabel),
        },
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, hubName, predictionName, parameters)
    if err != nil {
        return fmt.Errorf("Error creating Prediction (Prediction Name %q / Hub Name %q / Resource Group %q): %+v", predictionName, hubName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Prediction (Prediction Name %q / Hub Name %q / Resource Group %q): %+v", predictionName, hubName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, hubName, predictionName)
    if err != nil {
        return fmt.Errorf("Error retrieving Prediction (Prediction Name %q / Hub Name %q / Resource Group %q): %+v", predictionName, hubName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Prediction (Prediction Name %q / Hub Name %q / Resource Group %q) ID", predictionName, hubName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmPredictionRead(d, meta)
}

func resourceArmPredictionRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).predictionsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    hubName := id.Path["hubs"]
    predictionName := id.Path["predictions"]

    resp, err := client.Get(ctx, resourceGroup, hubName, predictionName)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Prediction %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Prediction (Prediction Name %q / Hub Name %q / Resource Group %q): %+v", predictionName, hubName, resourceGroup, err)
    }


    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if prediction := resp.Prediction; prediction != nil {
        d.Set("auto_analyze", prediction.AutoAnalyze)
        d.Set("description", utils.FlattenKeyValuePairs(prediction.Description))
        d.Set("display_name", utils.FlattenKeyValuePairs(prediction.DisplayName))
        if err := d.Set("grades", flattenArmPredictionPrediction_gradesItem(prediction.Grades)); err != nil {
            return fmt.Errorf("Error setting `grades`: %+v", err)
        }
        d.Set("involved_interaction_types", utils.FlattenStringSlice(prediction.InvolvedInteractionTypes))
        d.Set("involved_kpi_types", utils.FlattenStringSlice(prediction.InvolvedKpiTypes))
        d.Set("involved_relationships", utils.FlattenStringSlice(prediction.InvolvedRelationships))
        if err := d.Set("mappings", flattenArmPredictionPrediction_mappings(prediction.Mappings)); err != nil {
            return fmt.Errorf("Error setting `mappings`: %+v", err)
        }
        d.Set("negative_outcome_expression", prediction.NegativeOutcomeExpression)
        d.Set("positive_outcome_expression", prediction.PositiveOutcomeExpression)
        d.Set("prediction_name", prediction.PredictionName)
        d.Set("primary_profile_type", prediction.PrimaryProfileType)
        d.Set("provisioning_state", string(prediction.ProvisioningState))
        d.Set("scope_expression", prediction.ScopeExpression)
        d.Set("score_label", prediction.ScoreLabel)
        if err := d.Set("system_generated_entities", flattenArmPredictionPrediction_systemGeneratedEntities(prediction.SystemGeneratedEntities)); err != nil {
            return fmt.Errorf("Error setting `system_generated_entities`: %+v", err)
        }
        d.Set("tenant_id", prediction.TenantID)
    }
    d.Set("hub_name", hubName)
    d.Set("prediction_name", predictionName)
    d.Set("type", resp.Type)

    return nil
}


func resourceArmPredictionDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).predictionsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    hubName := id.Path["hubs"]
    predictionName := id.Path["predictions"]

    future, err := client.Delete(ctx, resourceGroup, hubName, predictionName)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Prediction (Prediction Name %q / Hub Name %q / Resource Group %q): %+v", predictionName, hubName, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Prediction (Prediction Name %q / Hub Name %q / Resource Group %q): %+v", predictionName, hubName, resourceGroup, err)
        }
    }

    return nil
}

func expandArmPredictionPrediction_gradesItem(input []interface{}) *[]customerinsights.Prediction_gradesItem {
    results := make([]customerinsights.Prediction_gradesItem, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        gradeName := v["grade_name"].(string)
        minScoreThreshold := v["min_score_threshold"].(int)
        maxScoreThreshold := v["max_score_threshold"].(int)

        result := customerinsights.Prediction_gradesItem{
            GradeName: utils.String(gradeName),
            MaxScoreThreshold: utils.Int(maxScoreThreshold),
            MinScoreThreshold: utils.Int(minScoreThreshold),
        }

        results = append(results, result)
    }
    return &results
}

func expandArmPredictionPrediction_mappings(input []interface{}) *customerinsights.Prediction_mappings {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    score := v["score"].(string)
    grade := v["grade"].(string)
    reason := v["reason"].(string)

    result := customerinsights.Prediction_mappings{
        Grade: utils.String(grade),
        Reason: utils.String(reason),
        Score: utils.String(score),
    }
    return &result
}


func flattenArmPredictionPrediction_gradesItem(input *[]customerinsights.Prediction_gradesItem) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        if gradeName := item.GradeName; gradeName != nil {
            v["grade_name"] = *gradeName
        }
        if maxScoreThreshold := item.MaxScoreThreshold; maxScoreThreshold != nil {
            v["max_score_threshold"] = *maxScoreThreshold
        }
        if minScoreThreshold := item.MinScoreThreshold; minScoreThreshold != nil {
            v["min_score_threshold"] = *minScoreThreshold
        }

        results = append(results, v)
    }

    return results
}

func flattenArmPredictionPrediction_mappings(input *customerinsights.Prediction_mappings) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if grade := input.Grade; grade != nil {
        result["grade"] = *grade
    }
    if reason := input.Reason; reason != nil {
        result["reason"] = *reason
    }
    if score := input.Score; score != nil {
        result["score"] = *score
    }

    return []interface{}{result}
}

func flattenArmPredictionPrediction_systemGeneratedEntities(input *customerinsights.Prediction_systemGeneratedEntities) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})


    return []interface{}{result}
}
