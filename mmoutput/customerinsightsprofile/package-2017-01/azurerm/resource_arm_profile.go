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



func resourceArmProfile() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmProfileCreateUpdate,
        Read: resourceArmProfileRead,
        Update: resourceArmProfileCreateUpdate,
        Delete: resourceArmProfileDelete,

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

            "hub_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "api_entity_set_name": {
                Type: schema.TypeString,
                Optional: true,
            },

            "attributes": {
                Type: schema.TypeMap,
                Optional: true,
                Elem: &schema.Schema{Type: schema.TypeString},
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

            "entity_type": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(customerinsights.None),
                    string(customerinsights.Profile),
                    string(customerinsights.Interaction),
                    string(customerinsights.Relationship),
                }, false),
                Default: string(customerinsights.None),
            },

            "fields": {
                Type: schema.TypeList,
                Optional: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "field_name": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "field_type": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "array_value_separator": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "enum_valid_values": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "localized_value_names": {
                                        Type: schema.TypeMap,
                                        Optional: true,
                                        Elem: &schema.Schema{Type: schema.TypeString},
                                    },
                                    "value": {
                                        Type: schema.TypeInt,
                                        Optional: true,
                                    },
                                },
                            },
                        },
                        "is_array": {
                            Type: schema.TypeBool,
                            Optional: true,
                        },
                        "is_available_in_graph": {
                            Type: schema.TypeBool,
                            Optional: true,
                        },
                        "is_enum": {
                            Type: schema.TypeBool,
                            Optional: true,
                        },
                        "is_flag_enum": {
                            Type: schema.TypeBool,
                            Optional: true,
                        },
                        "is_image": {
                            Type: schema.TypeBool,
                            Optional: true,
                        },
                        "is_localized_string": {
                            Type: schema.TypeBool,
                            Optional: true,
                        },
                        "is_name": {
                            Type: schema.TypeBool,
                            Optional: true,
                        },
                        "is_required": {
                            Type: schema.TypeBool,
                            Optional: true,
                        },
                        "max_length": {
                            Type: schema.TypeInt,
                            Optional: true,
                        },
                        "property_id": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "schema_item_prop_link": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },

            "instances_count": {
                Type: schema.TypeInt,
                Optional: true,
            },

            "large_image": {
                Type: schema.TypeString,
                Optional: true,
            },

            "localized_attributes": {
                Type: schema.TypeMap,
                Optional: true,
                Elem: &schema.Schema{Type: schema.TypeString},
            },

            "medium_image": {
                Type: schema.TypeString,
                Optional: true,
            },

            "schema_item_type_link": {
                Type: schema.TypeString,
                Optional: true,
            },

            "small_image": {
                Type: schema.TypeString,
                Optional: true,
            },

            "strong_ids": {
                Type: schema.TypeList,
                Optional: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "key_property_names": {
                            Type: schema.TypeList,
                            Required: true,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                        "strong_id_name": {
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
                    },
                },
            },

            "timestamp_field_name": {
                Type: schema.TypeString,
                Optional: true,
            },

            "type_name": {
                Type: schema.TypeString,
                Optional: true,
            },

            "last_changed_utc": {
                Type: schema.TypeString,
                Computed: true,
            },

            "provisioning_state": {
                Type: schema.TypeString,
                Computed: true,
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

func resourceArmProfileCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).profilesClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    hubName := d.Get("hub_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, hubName, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Profile %q (Hub Name %q / Resource Group %q): %+v", name, hubName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_profile", *existing.ID)
        }
    }

    apiEntitySetName := d.Get("api_entity_set_name").(string)
    attributes := d.Get("attributes").(map[string]interface{})
    description := d.Get("description").(map[string]interface{})
    displayName := d.Get("display_name").(map[string]interface{})
    entityType := d.Get("entity_type").(string)
    fields := d.Get("fields").([]interface{})
    instancesCount := d.Get("instances_count").(int)
    largeImage := d.Get("large_image").(string)
    localizedAttributes := d.Get("localized_attributes").(map[string]interface{})
    mediumImage := d.Get("medium_image").(string)
    schemaItemTypeLink := d.Get("schema_item_type_link").(string)
    smallImage := d.Get("small_image").(string)
    strongIds := d.Get("strong_ids").([]interface{})
    timestampFieldName := d.Get("timestamp_field_name").(string)
    typeName := d.Get("type_name").(string)

    parameters := customerinsights.ProfileResourceFormat{
        ProfileTypeDefinition: &customerinsights.ProfileTypeDefinition{
            ApiEntitySetName: utils.String(apiEntitySetName),
            Attributes: utils.ExpandKeyValuePairs(attributes),
            Description: utils.ExpandKeyValuePairs(description),
            DisplayName: utils.ExpandKeyValuePairs(displayName),
            EntityType: customerinsights.EntityTypes(entityType),
            Fields: expandArmProfilePropertyDefinition(fields),
            InstancesCount: utils.Int(instancesCount),
            LargeImage: utils.String(largeImage),
            LocalizedAttributes: utils.ExpandKeyValuePairs(localizedAttributes),
            MediumImage: utils.String(mediumImage),
            SchemaItemTypeLink: utils.String(schemaItemTypeLink),
            SmallImage: utils.String(smallImage),
            StrongIds: expandArmProfileStrongId(strongIds),
            TimestampFieldName: utils.String(timestampFieldName),
            TypeName: utils.String(typeName),
        },
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, hubName, name, parameters)
    if err != nil {
        return fmt.Errorf("Error creating Profile %q (Hub Name %q / Resource Group %q): %+v", name, hubName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Profile %q (Hub Name %q / Resource Group %q): %+v", name, hubName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, hubName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Profile %q (Hub Name %q / Resource Group %q): %+v", name, hubName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Profile %q (Hub Name %q / Resource Group %q) ID", name, hubName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmProfileRead(d, meta)
}

func resourceArmProfileRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).profilesClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    hubName := id.Path["hubs"]
    name := id.Path["profiles"]

    resp, err := client.Get(ctx, resourceGroup, hubName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Profile %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Profile %q (Hub Name %q / Resource Group %q): %+v", name, hubName, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if profileTypeDefinition := resp.ProfileTypeDefinition; profileTypeDefinition != nil {
        d.Set("api_entity_set_name", profileTypeDefinition.ApiEntitySetName)
        d.Set("attributes", utils.FlattenKeyValuePairs(profileTypeDefinition.Attributes))
        d.Set("description", utils.FlattenKeyValuePairs(profileTypeDefinition.Description))
        d.Set("display_name", utils.FlattenKeyValuePairs(profileTypeDefinition.DisplayName))
        d.Set("entity_type", string(profileTypeDefinition.EntityType))
        if err := d.Set("fields", flattenArmProfilePropertyDefinition(profileTypeDefinition.Fields)); err != nil {
            return fmt.Errorf("Error setting `fields`: %+v", err)
        }
        d.Set("instances_count", profileTypeDefinition.InstancesCount)
        d.Set("large_image", profileTypeDefinition.LargeImage)
        d.Set("last_changed_utc", (profileTypeDefinition.LastChangedUtc).String())
        d.Set("localized_attributes", utils.FlattenKeyValuePairs(profileTypeDefinition.LocalizedAttributes))
        d.Set("medium_image", profileTypeDefinition.MediumImage)
        d.Set("provisioning_state", string(profileTypeDefinition.ProvisioningState))
        d.Set("schema_item_type_link", profileTypeDefinition.SchemaItemTypeLink)
        d.Set("small_image", profileTypeDefinition.SmallImage)
        if err := d.Set("strong_ids", flattenArmProfileStrongId(profileTypeDefinition.StrongIds)); err != nil {
            return fmt.Errorf("Error setting `strong_ids`: %+v", err)
        }
        d.Set("tenant_id", profileTypeDefinition.TenantID)
        d.Set("timestamp_field_name", profileTypeDefinition.TimestampFieldName)
        d.Set("type_name", profileTypeDefinition.TypeName)
    }
    d.Set("hub_name", hubName)
    d.Set("type", resp.Type)

    return nil
}


func resourceArmProfileDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).profilesClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    hubName := id.Path["hubs"]
    name := id.Path["profiles"]

    future, err := client.Delete(ctx, resourceGroup, hubName, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Profile %q (Hub Name %q / Resource Group %q): %+v", name, hubName, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Profile %q (Hub Name %q / Resource Group %q): %+v", name, hubName, resourceGroup, err)
        }
    }

    return nil
}

func expandArmProfilePropertyDefinition(input []interface{}) *[]customerinsights.PropertyDefinition {
    results := make([]customerinsights.PropertyDefinition, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        arrayValueSeparator := v["array_value_separator"].(string)
        enumValidValues := v["enum_valid_values"].([]interface{})
        fieldName := v["field_name"].(string)
        fieldType := v["field_type"].(string)
        isArray := v["is_array"].(bool)
        isEnum := v["is_enum"].(bool)
        isFlagEnum := v["is_flag_enum"].(bool)
        isImage := v["is_image"].(bool)
        isLocalizedString := v["is_localized_string"].(bool)
        isName := v["is_name"].(bool)
        isRequired := v["is_required"].(bool)
        propertyId := v["property_id"].(string)
        schemaItemPropLink := v["schema_item_prop_link"].(string)
        maxLength := v["max_length"].(int)
        isAvailableInGraph := v["is_available_in_graph"].(bool)

        result := customerinsights.PropertyDefinition{
            ArrayValueSeparator: utils.String(arrayValueSeparator),
            EnumValidValues: expandArmProfileProfileEnumValidValuesFormat(enumValidValues),
            FieldName: utils.String(fieldName),
            FieldType: utils.String(fieldType),
            IsArray: utils.Bool(isArray),
            IsAvailableInGraph: utils.Bool(isAvailableInGraph),
            IsEnum: utils.Bool(isEnum),
            IsFlagEnum: utils.Bool(isFlagEnum),
            IsImage: utils.Bool(isImage),
            IsLocalizedString: utils.Bool(isLocalizedString),
            IsName: utils.Bool(isName),
            IsRequired: utils.Bool(isRequired),
            MaxLength: utils.Int(maxLength),
            PropertyID: utils.String(propertyId),
            SchemaItemPropLink: utils.String(schemaItemPropLink),
        }

        results = append(results, result)
    }
    return &results
}

func expandArmProfileStrongId(input []interface{}) *[]customerinsights.StrongId {
    results := make([]customerinsights.StrongId, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        keyPropertyNames := v["key_property_names"].([]interface{})
        strongIdName := v["strong_id_name"].(string)
        displayName := v["display_name"].(map[string]interface{})
        description := v["description"].(map[string]interface{})

        result := customerinsights.StrongId{
            Description: utils.ExpandKeyValuePairs(description),
            DisplayName: utils.ExpandKeyValuePairs(displayName),
            KeyPropertyNames: utils.ExpandStringSlice(keyPropertyNames),
            StrongIDName: utils.String(strongIdName),
        }

        results = append(results, result)
    }
    return &results
}

func expandArmProfileProfileEnumValidValuesFormat(input []interface{}) *[]customerinsights.ProfileEnumValidValuesFormat {
    results := make([]customerinsights.ProfileEnumValidValuesFormat, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        value := v["value"].(int)
        localizedValueNames := v["localized_value_names"].(map[string]interface{})

        result := customerinsights.ProfileEnumValidValuesFormat{
            LocalizedValueNames: utils.ExpandKeyValuePairs(localizedValueNames),
            Value: utils.Int(value),
        }

        results = append(results, result)
    }
    return &results
}


func flattenArmProfilePropertyDefinition(input *[]customerinsights.PropertyDefinition) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        if arrayValueSeparator := item.ArrayValueSeparator; arrayValueSeparator != nil {
            v["array_value_separator"] = *arrayValueSeparator
        }
        v["enum_valid_values"] = flattenArmProfileProfileEnumValidValuesFormat(item.EnumValidValues)
        if fieldName := item.FieldName; fieldName != nil {
            v["field_name"] = *fieldName
        }
        if fieldType := item.FieldType; fieldType != nil {
            v["field_type"] = *fieldType
        }
        if isArray := item.IsArray; isArray != nil {
            v["is_array"] = *isArray
        }
        if isAvailableInGraph := item.IsAvailableInGraph; isAvailableInGraph != nil {
            v["is_available_in_graph"] = *isAvailableInGraph
        }
        if isEnum := item.IsEnum; isEnum != nil {
            v["is_enum"] = *isEnum
        }
        if isFlagEnum := item.IsFlagEnum; isFlagEnum != nil {
            v["is_flag_enum"] = *isFlagEnum
        }
        if isImage := item.IsImage; isImage != nil {
            v["is_image"] = *isImage
        }
        if isLocalizedString := item.IsLocalizedString; isLocalizedString != nil {
            v["is_localized_string"] = *isLocalizedString
        }
        if isName := item.IsName; isName != nil {
            v["is_name"] = *isName
        }
        if isRequired := item.IsRequired; isRequired != nil {
            v["is_required"] = *isRequired
        }
        if maxLength := item.MaxLength; maxLength != nil {
            v["max_length"] = *maxLength
        }
        if propertyId := item.PropertyID; propertyId != nil {
            v["property_id"] = *propertyId
        }
        if schemaItemPropLink := item.SchemaItemPropLink; schemaItemPropLink != nil {
            v["schema_item_prop_link"] = *schemaItemPropLink
        }

        results = append(results, v)
    }

    return results
}

func flattenArmProfileStrongId(input *[]customerinsights.StrongId) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        v["description"] = utils.FlattenKeyValuePairs(item.Description)
        v["display_name"] = utils.FlattenKeyValuePairs(item.DisplayName)
        v["key_property_names"] = utils.FlattenStringSlice(item.KeyPropertyNames)
        if strongIdName := item.StrongIDName; strongIdName != nil {
            v["strong_id_name"] = *strongIdName
        }

        results = append(results, v)
    }

    return results
}

func flattenArmProfileProfileEnumValidValuesFormat(input *[]customerinsights.ProfileEnumValidValuesFormat) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        v["localized_value_names"] = utils.FlattenKeyValuePairs(item.LocalizedValueNames)
        if value := item.Value; value != nil {
            v["value"] = *value
        }

        results = append(results, v)
    }

    return results
}
