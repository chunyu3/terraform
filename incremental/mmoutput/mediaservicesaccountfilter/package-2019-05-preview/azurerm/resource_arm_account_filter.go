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



func resourceArmAccountFilter() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmAccountFilterCreate,
        Read: resourceArmAccountFilterRead,
        Update: resourceArmAccountFilterUpdate,
        Delete: resourceArmAccountFilterDelete,

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

            "first_quality": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "bitrate": {
                            Type: schema.TypeInt,
                            Required: true,
                        },
                    },
                },
            },

            "presentation_time_range": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "end_timestamp": {
                            Type: schema.TypeInt,
                            Optional: true,
                        },
                        "force_end_timestamp": {
                            Type: schema.TypeBool,
                            Optional: true,
                        },
                        "live_backoff_duration": {
                            Type: schema.TypeInt,
                            Optional: true,
                        },
                        "presentation_window_duration": {
                            Type: schema.TypeInt,
                            Optional: true,
                        },
                        "start_timestamp": {
                            Type: schema.TypeInt,
                            Optional: true,
                        },
                        "timescale": {
                            Type: schema.TypeInt,
                            Optional: true,
                        },
                    },
                },
            },

            "tracks": {
                Type: schema.TypeList,
                Optional: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "track_selections": {
                            Type: schema.TypeList,
                            Required: true,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "operation": {
                                        Type: schema.TypeString,
                                        Required: true,
                                        ValidateFunc: validation.StringInSlice([]string{
                                            string(mediaservices.Equal),
                                            string(mediaservices.NotEqual),
                                        }, false),
                                    },
                                    "property": {
                                        Type: schema.TypeString,
                                        Required: true,
                                        ValidateFunc: validation.StringInSlice([]string{
                                            string(mediaservices.Unknown),
                                            string(mediaservices.Type),
                                            string(mediaservices.Name),
                                            string(mediaservices.Language),
                                            string(mediaservices.FourCC),
                                            string(mediaservices.Bitrate),
                                        }, false),
                                    },
                                    "value": {
                                        Type: schema.TypeString,
                                        Required: true,
                                        ValidateFunc: validate.NoEmptyStrings,
                                    },
                                },
                            },
                        },
                    },
                },
            },

            "type": {
                Type: schema.TypeString,
                Computed: true,
            },
        },
    }
}

func resourceArmAccountFilterCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).accountFiltersClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    accountName := d.Get("account_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, accountName, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Account Filter %q (Account Name %q / Resource Group %q): %+v", name, accountName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_account_filter", *existing.ID)
        }
    }

    firstQuality := d.Get("first_quality").([]interface{})
    presentationTimeRange := d.Get("presentation_time_range").([]interface{})
    tracks := d.Get("tracks").([]interface{})

    parameters := mediaservices.AccountFilter{
        MediaFilterProperties: &mediaservices.MediaFilterProperties{
            FirstQuality: expandArmAccountFilterFirstQuality(firstQuality),
            PresentationTimeRange: expandArmAccountFilterPresentationTimeRange(presentationTimeRange),
            Tracks: expandArmAccountFilterFilterTrackSelection(tracks),
        },
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroup, accountName, name, parameters); err != nil {
        return fmt.Errorf("Error creating Account Filter %q (Account Name %q / Resource Group %q): %+v", name, accountName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, accountName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Account Filter %q (Account Name %q / Resource Group %q): %+v", name, accountName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Account Filter %q (Account Name %q / Resource Group %q) ID", name, accountName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmAccountFilterRead(d, meta)
}

func resourceArmAccountFilterRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).accountFiltersClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    accountName := id.Path["mediaServices"]
    name := id.Path["accountFilters"]

    resp, err := client.Get(ctx, resourceGroup, accountName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Account Filter %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Account Filter %q (Account Name %q / Resource Group %q): %+v", name, accountName, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    d.Set("account_name", accountName)
    d.Set("type", resp.Type)

    return nil
}

func resourceArmAccountFilterUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).accountFiltersClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    accountName := d.Get("account_name").(string)
    firstQuality := d.Get("first_quality").([]interface{})
    presentationTimeRange := d.Get("presentation_time_range").([]interface{})
    tracks := d.Get("tracks").([]interface{})

    parameters := mediaservices.AccountFilter{
        MediaFilterProperties: &mediaservices.MediaFilterProperties{
            FirstQuality: expandArmAccountFilterFirstQuality(firstQuality),
            PresentationTimeRange: expandArmAccountFilterPresentationTimeRange(presentationTimeRange),
            Tracks: expandArmAccountFilterFilterTrackSelection(tracks),
        },
    }


    if _, err := client.Update(ctx, resourceGroup, accountName, name, parameters); err != nil {
        return fmt.Errorf("Error updating Account Filter %q (Account Name %q / Resource Group %q): %+v", name, accountName, resourceGroup, err)
    }

    return resourceArmAccountFilterRead(d, meta)
}

func resourceArmAccountFilterDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).accountFiltersClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    accountName := id.Path["mediaServices"]
    name := id.Path["accountFilters"]

    if _, err := client.Delete(ctx, resourceGroup, accountName, name); err != nil {
        return fmt.Errorf("Error deleting Account Filter %q (Account Name %q / Resource Group %q): %+v", name, accountName, resourceGroup, err)
    }

    return nil
}

func expandArmAccountFilterFirstQuality(input []interface{}) *mediaservices.FirstQuality {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    bitrate := v["bitrate"].(int)

    result := mediaservices.FirstQuality{
        Bitrate: utils.Int32(int32(bitrate)),
    }
    return &result
}

func expandArmAccountFilterPresentationTimeRange(input []interface{}) *mediaservices.PresentationTimeRange {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    startTimestamp := v["start_timestamp"].(int)
    endTimestamp := v["end_timestamp"].(int)
    presentationWindowDuration := v["presentation_window_duration"].(int)
    liveBackoffDuration := v["live_backoff_duration"].(int)
    timescale := v["timescale"].(int)
    forceEndTimestamp := v["force_end_timestamp"].(bool)

    result := mediaservices.PresentationTimeRange{
        EndTimestamp: utils.Int64(int64(endTimestamp)),
        ForceEndTimestamp: utils.Bool(forceEndTimestamp),
        LiveBackoffDuration: utils.Int64(int64(liveBackoffDuration)),
        PresentationWindowDuration: utils.Int64(int64(presentationWindowDuration)),
        StartTimestamp: utils.Int64(int64(startTimestamp)),
        Timescale: utils.Int64(int64(timescale)),
    }
    return &result
}

func expandArmAccountFilterFilterTrackSelection(input []interface{}) *[]mediaservices.FilterTrackSelection {
    results := make([]mediaservices.FilterTrackSelection, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        trackSelections := v["track_selections"].([]interface{})

        result := mediaservices.FilterTrackSelection{
            TrackSelections: expandArmAccountFilterFilterTrackPropertyCondition(trackSelections),
        }

        results = append(results, result)
    }
    return &results
}

func expandArmAccountFilterFilterTrackPropertyCondition(input []interface{}) *[]mediaservices.FilterTrackPropertyCondition {
    results := make([]mediaservices.FilterTrackPropertyCondition, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        property := v["property"].(string)
        value := v["value"].(string)
        operation := v["operation"].(string)

        result := mediaservices.FilterTrackPropertyCondition{
            Operation: mediaservices.FilterTrackPropertyCompareOperation(operation),
            Property: mediaservices.FilterTrackPropertyType(property),
            Value: utils.String(value),
        }

        results = append(results, result)
    }
    return &results
}