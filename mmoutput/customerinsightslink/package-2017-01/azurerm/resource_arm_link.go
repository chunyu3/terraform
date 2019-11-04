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



func resourceArmLink() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmLinkCreateUpdate,
        Read: resourceArmLinkRead,
        Update: resourceArmLinkCreateUpdate,
        Delete: resourceArmLinkDelete,

        Importer: &schema.ResourceImporter{
            State: schema.ImportStatePassthrough,
        },


        Schema: map[string]*schema.Schema{
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

            "link_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "participant_property_references": {
                Type: schema.TypeList,
                Required: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "interaction_property_name": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "profile_property_name": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                    },
                },
            },

            "source_interaction_type": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "target_profile_type": {
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

            "mappings": {
                Type: schema.TypeList,
                Optional: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "interaction_type_property_name": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "profile_type_property_name": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "is_profile_type_id": {
                            Type: schema.TypeBool,
                            Optional: true,
                        },
                        "link_type": {
                            Type: schema.TypeString,
                            Optional: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(customerinsights.UpdateAlways),
                                string(customerinsights.CopyIfNull),
                            }, false),
                            Default: string(customerinsights.UpdateAlways),
                        },
                    },
                },
            },

            "operation_type": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(customerinsights.Upsert),
                    string(customerinsights.Delete),
                }, false),
                Default: string(customerinsights.Upsert),
            },

            "reference_only": {
                Type: schema.TypeBool,
                Optional: true,
            },

            "link_name": {
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

func resourceArmLinkCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).linksClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    hubName := d.Get("hub_name").(string)
    linkName := d.Get("link_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, hubName, linkName)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Link (Link Name %q / Hub Name %q / Resource Group %q): %+v", linkName, hubName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_link", *existing.ID)
        }
    }

    description := d.Get("description").(map[string]interface{})
    displayName := d.Get("display_name").(map[string]interface{})
    mappings := d.Get("mappings").([]interface{})
    operationType := d.Get("operation_type").(string)
    participantPropertyReferences := d.Get("participant_property_references").([]interface{})
    referenceOnly := d.Get("reference_only").(bool)
    sourceInteractionType := d.Get("source_interaction_type").(string)
    targetProfileType := d.Get("target_profile_type").(string)

    parameters := customerinsights.LinkResourceFormat{
        LinkDefinition: &customerinsights.LinkDefinition{
            Description: utils.ExpandKeyValuePairs(description),
            DisplayName: utils.ExpandKeyValuePairs(displayName),
            Mappings: expandArmLinkTypePropertiesMapping(mappings),
            OperationType: customerinsights.InstanceOperationType(operationType),
            ParticipantPropertyReferences: expandArmLinkParticipantPropertyReference(participantPropertyReferences),
            ReferenceOnly: utils.Bool(referenceOnly),
            SourceInteractionType: utils.String(sourceInteractionType),
            TargetProfileType: utils.String(targetProfileType),
        },
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, hubName, linkName, parameters)
    if err != nil {
        return fmt.Errorf("Error creating Link (Link Name %q / Hub Name %q / Resource Group %q): %+v", linkName, hubName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Link (Link Name %q / Hub Name %q / Resource Group %q): %+v", linkName, hubName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, hubName, linkName)
    if err != nil {
        return fmt.Errorf("Error retrieving Link (Link Name %q / Hub Name %q / Resource Group %q): %+v", linkName, hubName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Link (Link Name %q / Hub Name %q / Resource Group %q) ID", linkName, hubName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmLinkRead(d, meta)
}

func resourceArmLinkRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).linksClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    hubName := id.Path["hubs"]
    linkName := id.Path["links"]

    resp, err := client.Get(ctx, resourceGroup, hubName, linkName)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Link %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Link (Link Name %q / Hub Name %q / Resource Group %q): %+v", linkName, hubName, resourceGroup, err)
    }


    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if linkDefinition := resp.LinkDefinition; linkDefinition != nil {
        d.Set("description", utils.FlattenKeyValuePairs(linkDefinition.Description))
        d.Set("display_name", utils.FlattenKeyValuePairs(linkDefinition.DisplayName))
        d.Set("link_name", linkDefinition.LinkName)
        if err := d.Set("mappings", flattenArmLinkTypePropertiesMapping(linkDefinition.Mappings)); err != nil {
            return fmt.Errorf("Error setting `mappings`: %+v", err)
        }
        d.Set("operation_type", string(linkDefinition.OperationType))
        if err := d.Set("participant_property_references", flattenArmLinkParticipantPropertyReference(linkDefinition.ParticipantPropertyReferences)); err != nil {
            return fmt.Errorf("Error setting `participant_property_references`: %+v", err)
        }
        d.Set("provisioning_state", string(linkDefinition.ProvisioningState))
        d.Set("reference_only", linkDefinition.ReferenceOnly)
        d.Set("source_interaction_type", linkDefinition.SourceInteractionType)
        d.Set("target_profile_type", linkDefinition.TargetProfileType)
        d.Set("tenant_id", linkDefinition.TenantID)
    }
    d.Set("hub_name", hubName)
    d.Set("link_name", linkName)
    d.Set("type", resp.Type)

    return nil
}


func resourceArmLinkDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).linksClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    hubName := id.Path["hubs"]
    linkName := id.Path["links"]

    if _, err := client.Delete(ctx, resourceGroup, hubName, linkName); err != nil {
        return fmt.Errorf("Error deleting Link (Link Name %q / Hub Name %q / Resource Group %q): %+v", linkName, hubName, resourceGroup, err)
    }

    return nil
}

func expandArmLinkTypePropertiesMapping(input []interface{}) *[]customerinsights.TypePropertiesMapping {
    results := make([]customerinsights.TypePropertiesMapping, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        interactionTypePropertyName := v["interaction_type_property_name"].(string)
        profileTypePropertyName := v["profile_type_property_name"].(string)
        isProfileTypeId := v["is_profile_type_id"].(bool)
        linkType := v["link_type"].(string)

        result := customerinsights.TypePropertiesMapping{
            InteractionTypePropertyName: utils.String(interactionTypePropertyName),
            IsProfileTypeID: utils.Bool(isProfileTypeId),
            LinkType: customerinsights.LinkTypes(linkType),
            ProfileTypePropertyName: utils.String(profileTypePropertyName),
        }

        results = append(results, result)
    }
    return &results
}

func expandArmLinkParticipantPropertyReference(input []interface{}) *[]customerinsights.ParticipantPropertyReference {
    results := make([]customerinsights.ParticipantPropertyReference, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        interactionPropertyName := v["interaction_property_name"].(string)
        profilePropertyName := v["profile_property_name"].(string)

        result := customerinsights.ParticipantPropertyReference{
            InteractionPropertyName: utils.String(interactionPropertyName),
            ProfilePropertyName: utils.String(profilePropertyName),
        }

        results = append(results, result)
    }
    return &results
}


func flattenArmLinkTypePropertiesMapping(input *[]customerinsights.TypePropertiesMapping) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        if interactionTypePropertyName := item.InteractionTypePropertyName; interactionTypePropertyName != nil {
            v["interaction_type_property_name"] = *interactionTypePropertyName
        }
        if isProfileTypeId := item.IsProfileTypeID; isProfileTypeId != nil {
            v["is_profile_type_id"] = *isProfileTypeId
        }
        v["link_type"] = string(item.LinkType)
        if profileTypePropertyName := item.ProfileTypePropertyName; profileTypePropertyName != nil {
            v["profile_type_property_name"] = *profileTypePropertyName
        }

        results = append(results, v)
    }

    return results
}

func flattenArmLinkParticipantPropertyReference(input *[]customerinsights.ParticipantPropertyReference) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        if interactionPropertyName := item.InteractionPropertyName; interactionPropertyName != nil {
            v["interaction_property_name"] = *interactionPropertyName
        }
        if profilePropertyName := item.ProfilePropertyName; profilePropertyName != nil {
            v["profile_property_name"] = *profilePropertyName
        }

        results = append(results, v)
    }

    return results
}