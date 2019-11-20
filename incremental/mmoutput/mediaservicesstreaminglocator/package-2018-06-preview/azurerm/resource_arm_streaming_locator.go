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



func resourceArmStreamingLocator() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmStreamingLocatorCreateUpdate,
        Read: resourceArmStreamingLocatorRead,
        Update: resourceArmStreamingLocatorCreateUpdate,
        Delete: resourceArmStreamingLocatorDelete,

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

            "asset_name": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "streaming_policy_name": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "alternative_media_id": {
                Type: schema.TypeString,
                Optional: true,
            },

            "content_keys": {
                Type: schema.TypeList,
                Optional: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "id": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "label": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "tracks": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "track_selections": {
                                        Type: schema.TypeList,
                                        Optional: true,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "operation": {
                                                    Type: schema.TypeString,
                                                    Required: true,
                                                    ValidateFunc: validation.StringInSlice([]string{
                                                        string(mediaservices.Unknown),
                                                        string(mediaservices.Equal),
                                                    }, false),
                                                },
                                                "property": {
                                                    Type: schema.TypeString,
                                                    Required: true,
                                                    ValidateFunc: validation.StringInSlice([]string{
                                                        string(mediaservices.Unknown),
                                                        string(mediaservices.FourCC),
                                                    }, false),
                                                },
                                                "value": {
                                                    Type: schema.TypeString,
                                                    Optional: true,
                                                },
                                            },
                                        },
                                    },
                                },
                            },
                        },
                        "value": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },

            "default_content_key_policy_name": {
                Type: schema.TypeString,
                Optional: true,
            },

            "end_time": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validateRFC3339Date,
            },

            "start_time": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validateRFC3339Date,
            },

            "streaming_locator_id": {
                Type: schema.TypeString,
                Optional: true,
            },

            "type": {
                Type: schema.TypeString,
                Computed: true,
            },
        },
    }
}

func resourceArmStreamingLocatorCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).streamingLocatorsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    accountName := d.Get("account_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, accountName, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Streaming Locator %q (Account Name %q / Resource Group %q): %+v", name, accountName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_streaming_locator", *existing.ID)
        }
    }

    alternativeMediaId := d.Get("alternative_media_id").(string)
    assetName := d.Get("asset_name").(string)
    contentKeys := d.Get("content_keys").([]interface{})
    defaultContentKeyPolicyName := d.Get("default_content_key_policy_name").(string)
    endTime := d.Get("end_time").(string)
    startTime := d.Get("start_time").(string)
    streamingLocatorId := d.Get("streaming_locator_id").(string)
    streamingPolicyName := d.Get("streaming_policy_name").(string)

    parameters := mediaservices.StreamingLocator{
        StreamingLocatorProperties: &mediaservices.StreamingLocatorProperties{
            AlternativeMediaID: utils.String(alternativeMediaId),
            AssetName: utils.String(assetName),
            ContentKeys: expandArmStreamingLocatorStreamingLocatorContentKey(contentKeys),
            DefaultContentKeyPolicyName: utils.String(defaultContentKeyPolicyName),
            EndTime: convertStringToDate(endTime),
            StartTime: convertStringToDate(startTime),
            StreamingLocatorID: utils.String(streamingLocatorId),
            StreamingPolicyName: utils.String(streamingPolicyName),
        },
    }


    if _, err := client.Create(ctx, resourceGroup, accountName, name, parameters); err != nil {
        return fmt.Errorf("Error creating Streaming Locator %q (Account Name %q / Resource Group %q): %+v", name, accountName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, accountName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Streaming Locator %q (Account Name %q / Resource Group %q): %+v", name, accountName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Streaming Locator %q (Account Name %q / Resource Group %q) ID", name, accountName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmStreamingLocatorRead(d, meta)
}

func resourceArmStreamingLocatorRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).streamingLocatorsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    accountName := id.Path["mediaServices"]
    name := id.Path["streamingLocators"]

    resp, err := client.Get(ctx, resourceGroup, accountName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Streaming Locator %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Streaming Locator %q (Account Name %q / Resource Group %q): %+v", name, accountName, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    d.Set("account_name", accountName)
    d.Set("type", resp.Type)

    return nil
}


func resourceArmStreamingLocatorDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).streamingLocatorsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    accountName := id.Path["mediaServices"]
    name := id.Path["streamingLocators"]

    if _, err := client.Delete(ctx, resourceGroup, accountName, name); err != nil {
        return fmt.Errorf("Error deleting Streaming Locator %q (Account Name %q / Resource Group %q): %+v", name, accountName, resourceGroup, err)
    }

    return nil
}

func expandArmStreamingLocatorStreamingLocatorContentKey(input []interface{}) *[]mediaservices.StreamingLocatorContentKey {
    results := make([]mediaservices.StreamingLocatorContentKey, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        id := v["id"].(string)
        label := v["label"].(string)
        value := v["value"].(string)
        tracks := v["tracks"].([]interface{})

        result := mediaservices.StreamingLocatorContentKey{
            ID: utils.String(id),
            Label: utils.String(label),
            Tracks: expandArmStreamingLocatorTrackSelection(tracks),
            Value: utils.String(value),
        }

        results = append(results, result)
    }
    return &results
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

func expandArmStreamingLocatorTrackSelection(input []interface{}) *[]mediaservices.TrackSelection {
    results := make([]mediaservices.TrackSelection, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        trackSelections := v["track_selections"].([]interface{})

        result := mediaservices.TrackSelection{
            TrackSelections: expandArmStreamingLocatorTrackPropertyCondition(trackSelections),
        }

        results = append(results, result)
    }
    return &results
}

func expandArmStreamingLocatorTrackPropertyCondition(input []interface{}) *[]mediaservices.TrackPropertyCondition {
    results := make([]mediaservices.TrackPropertyCondition, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        property := v["property"].(string)
        operation := v["operation"].(string)
        value := v["value"].(string)

        result := mediaservices.TrackPropertyCondition{
            Operation: mediaservices.TrackPropertyCompareOperation(operation),
            Property: mediaservices.TrackPropertyType(property),
            Value: utils.String(value),
        }

        results = append(results, result)
    }
    return &results
}
