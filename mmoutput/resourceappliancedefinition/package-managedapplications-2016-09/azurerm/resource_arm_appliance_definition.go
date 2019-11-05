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



func resourceArmApplianceDefinition() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmApplianceDefinitionCreateUpdate,
        Read: resourceArmApplianceDefinitionRead,
        Update: resourceArmApplianceDefinitionCreateUpdate,
        Delete: resourceArmApplianceDefinitionDelete,

        Importer: &schema.ResourceImporter{
            State: schema.ImportStatePassthrough,
        },


        Schema: map[string]*schema.Schema{
            "name": {
                Type: schema.TypeString,
                Computed: true,
            },

            "location": azure.SchemaLocation(),

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "appliance_definition_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "authorizations": {
                Type: schema.TypeList,
                Required: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "principal_id": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "role_definition_id": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                    },
                },
            },

            "lock_level": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(resource.CanNotDelete),
                    string(resource.ReadOnly),
                    string(resource.None),
                }, false),
            },

            "package_file_uri": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "artifacts": {
                Type: schema.TypeList,
                Optional: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "name": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "type": {
                            Type: schema.TypeString,
                            Optional: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(resource.Template),
                                string(resource.Custom),
                            }, false),
                            Default: string(resource.Template),
                        },
                        "uri": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },

            "description": {
                Type: schema.TypeString,
                Optional: true,
            },

            "display_name": {
                Type: schema.TypeString,
                Optional: true,
            },

            "identity": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "type": {
                            Type: schema.TypeString,
                            Optional: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(resource.SystemAssigned),
                            }, false),
                            Default: string(resource.SystemAssigned),
                        },
                    },
                },
            },

            "managed_by": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
            },

            "sku": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "name": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "capacity": {
                            Type: schema.TypeInt,
                            Optional: true,
                        },
                        "family": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "model": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "size": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "tier": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },

            "type": {
                Type: schema.TypeString,
                Computed: true,
            },

            "tags": tags.Schema(),
        },
    }
}

func resourceArmApplianceDefinitionCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).applianceDefinitionsClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    applianceDefinitionName := d.Get("appliance_definition_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, applianceDefinitionName)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Appliance Definition (Appliance Definition Name %q / Resource Group %q): %+v", applianceDefinitionName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_appliance_definition", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    artifacts := d.Get("artifacts").([]interface{})
    authorizations := d.Get("authorizations").([]interface{})
    description := d.Get("description").(string)
    displayName := d.Get("display_name").(string)
    identity := d.Get("identity").([]interface{})
    lockLevel := d.Get("lock_level").(string)
    managedBy := d.Get("managed_by").(string)
    packageFileUri := d.Get("package_file_uri").(string)
    sku := d.Get("sku").([]interface{})
    t := d.Get("tags").(map[string]interface{})

    parameters := resource.ApplianceDefinition{
        Identity: expandArmApplianceDefinitionIdentity(identity),
        Location: utils.String(location),
        ManagedBy: utils.String(managedBy),
        ApplianceDefinitionProperties: &resource.ApplianceDefinitionProperties{
            Artifacts: expandArmApplianceDefinitionApplianceArtifact(artifacts),
            Authorizations: expandArmApplianceDefinitionApplianceProviderAuthorization(authorizations),
            Description: utils.String(description),
            DisplayName: utils.String(displayName),
            LockLevel: resource.ApplianceLockLevel(lockLevel),
            PackageFileUri: utils.String(packageFileUri),
        },
        Sku: expandArmApplianceDefinitionSku(sku),
        Tags: tags.Expand(t),
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, applianceDefinitionName, parameters)
    if err != nil {
        return fmt.Errorf("Error creating Appliance Definition (Appliance Definition Name %q / Resource Group %q): %+v", applianceDefinitionName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Appliance Definition (Appliance Definition Name %q / Resource Group %q): %+v", applianceDefinitionName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, applianceDefinitionName)
    if err != nil {
        return fmt.Errorf("Error retrieving Appliance Definition (Appliance Definition Name %q / Resource Group %q): %+v", applianceDefinitionName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Appliance Definition (Appliance Definition Name %q / Resource Group %q) ID", applianceDefinitionName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmApplianceDefinitionRead(d, meta)
}

func resourceArmApplianceDefinitionRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).applianceDefinitionsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    applianceDefinitionName := id.Path["applianceDefinitions"]

    resp, err := client.Get(ctx, resourceGroup, applianceDefinitionName)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Appliance Definition %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Appliance Definition (Appliance Definition Name %q / Resource Group %q): %+v", applianceDefinitionName, resourceGroup, err)
    }


    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    d.Set("appliance_definition_name", applianceDefinitionName)
    if applianceDefinitionProperties := resp.ApplianceDefinitionProperties; applianceDefinitionProperties != nil {
        if err := d.Set("artifacts", flattenArmApplianceDefinitionApplianceArtifact(applianceDefinitionProperties.Artifacts)); err != nil {
            return fmt.Errorf("Error setting `artifacts`: %+v", err)
        }
        if err := d.Set("authorizations", flattenArmApplianceDefinitionApplianceProviderAuthorization(applianceDefinitionProperties.Authorizations)); err != nil {
            return fmt.Errorf("Error setting `authorizations`: %+v", err)
        }
        d.Set("description", applianceDefinitionProperties.Description)
        d.Set("display_name", applianceDefinitionProperties.DisplayName)
        d.Set("lock_level", string(applianceDefinitionProperties.LockLevel))
        d.Set("package_file_uri", applianceDefinitionProperties.PackageFileUri)
    }
    if err := d.Set("identity", flattenArmApplianceDefinitionIdentity(resp.Identity)); err != nil {
        return fmt.Errorf("Error setting `identity`: %+v", err)
    }
    d.Set("managed_by", resp.ManagedBy)
    if err := d.Set("sku", flattenArmApplianceDefinitionSku(resp.Sku)); err != nil {
        return fmt.Errorf("Error setting `sku`: %+v", err)
    }
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}


func resourceArmApplianceDefinitionDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).applianceDefinitionsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    applianceDefinitionName := id.Path["applianceDefinitions"]

    future, err := client.Delete(ctx, resourceGroup, applianceDefinitionName)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Appliance Definition (Appliance Definition Name %q / Resource Group %q): %+v", applianceDefinitionName, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Appliance Definition (Appliance Definition Name %q / Resource Group %q): %+v", applianceDefinitionName, resourceGroup, err)
        }
    }

    return nil
}

func expandArmApplianceDefinitionIdentity(input []interface{}) *resource.Identity {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    type := v["type"].(string)

    result := resource.Identity{
        Type: resource.IdentityType(type),
    }
    return &result
}

func expandArmApplianceDefinitionApplianceArtifact(input []interface{}) *[]resource.ApplianceArtifact {
    results := make([]resource.ApplianceArtifact, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        name := v["name"].(string)
        uri := v["uri"].(string)
        type := v["type"].(string)

        result := resource.ApplianceArtifact{
            Name: utils.String(name),
            Type: resource.ApplianceArtifactType(type),
            Uri: utils.String(uri),
        }

        results = append(results, result)
    }
    return &results
}

func expandArmApplianceDefinitionApplianceProviderAuthorization(input []interface{}) *[]resource.ApplianceProviderAuthorization {
    results := make([]resource.ApplianceProviderAuthorization, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        principalId := v["principal_id"].(string)
        roleDefinitionId := v["role_definition_id"].(string)

        result := resource.ApplianceProviderAuthorization{
            PrincipalID: utils.String(principalId),
            RoleDefinitionID: utils.String(roleDefinitionId),
        }

        results = append(results, result)
    }
    return &results
}

func expandArmApplianceDefinitionSku(input []interface{}) *resource.Sku {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    name := v["name"].(string)
    tier := v["tier"].(string)
    size := v["size"].(string)
    family := v["family"].(string)
    model := v["model"].(string)
    capacity := v["capacity"].(int)

    result := resource.Sku{
        Capacity: utils.Int32(int32(capacity)),
        Family: utils.String(family),
        Model: utils.String(model),
        Name: utils.String(name),
        Size: utils.String(size),
        Tier: utils.String(tier),
    }
    return &result
}


func flattenArmApplianceDefinitionApplianceArtifact(input *[]resource.ApplianceArtifact) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        if name := item.Name; name != nil {
            v["name"] = *name
        }
        v["type"] = string(item.Type)
        if uri := item.Uri; uri != nil {
            v["uri"] = *uri
        }

        results = append(results, v)
    }

    return results
}

func flattenArmApplianceDefinitionApplianceProviderAuthorization(input *[]resource.ApplianceProviderAuthorization) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        if principalId := item.PrincipalID; principalId != nil {
            v["principal_id"] = *principalId
        }
        if roleDefinitionId := item.RoleDefinitionID; roleDefinitionId != nil {
            v["role_definition_id"] = *roleDefinitionId
        }

        results = append(results, v)
    }

    return results
}

func flattenArmApplianceDefinitionIdentity(input *resource.Identity) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    result["type"] = string(input.Type)

    return []interface{}{result}
}

func flattenArmApplianceDefinitionSku(input *resource.Sku) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if name := input.Name; name != nil {
        result["name"] = *name
    }
    if capacity := input.Capacity; capacity != nil {
        result["capacity"] = int(*capacity)
    }
    if family := input.Family; family != nil {
        result["family"] = *family
    }
    if model := input.Model; model != nil {
        result["model"] = *model
    }
    if size := input.Size; size != nil {
        result["size"] = *size
    }
    if tier := input.Tier; tier != nil {
        result["tier"] = *tier
    }

    return []interface{}{result}
}
