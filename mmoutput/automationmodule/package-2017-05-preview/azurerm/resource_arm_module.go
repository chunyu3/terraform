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



func resourceArmModule() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmModuleCreate,
        Read: resourceArmModuleRead,
        Update: resourceArmModuleUpdate,
        Delete: resourceArmModuleDelete,

        Importer: &schema.ResourceImporter{
            State: schema.ImportStatePassthrough,
        },


        Schema: map[string]*schema.Schema{
            "automation_account_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "module_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "content_link": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "content_hash": {
                            Type: schema.TypeList,
                            Optional: true,
                            MaxItems: 1,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "algorithm": {
                                        Type: schema.TypeString,
                                        Required: true,
                                        ValidateFunc: validate.NoEmptyStrings,
                                    },
                                    "value": {
                                        Type: schema.TypeString,
                                        Required: true,
                                        ValidateFunc: validate.NoEmptyStrings,
                                    },
                                },
                            },
                        },
                        "uri": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "version": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },

            "location": azure.SchemaLocation(),

            "name": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
            },

            "tags": tags.Schema(),

            "activity_count": {
                Type: schema.TypeInt,
                Computed: true,
            },

            "creation_time": {
                Type: schema.TypeString,
                Computed: true,
            },

            "description": {
                Type: schema.TypeString,
                Computed: true,
            },

            "error": {
                Type: schema.TypeList,
                Computed: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "code": {
                            Type: schema.TypeString,
                            Computed: true,
                        },
                        "message": {
                            Type: schema.TypeString,
                            Computed: true,
                        },
                    },
                },
            },

            "etag": {
                Type: schema.TypeString,
                Computed: true,
            },

            "id": {
                Type: schema.TypeString,
                Computed: true,
            },

            "is_composite": {
                Type: schema.TypeBool,
                Computed: true,
            },

            "is_global": {
                Type: schema.TypeBool,
                Computed: true,
            },

            "last_modified_time": {
                Type: schema.TypeString,
                Computed: true,
            },

            "provisioning_state": {
                Type: schema.TypeString,
                Computed: true,
            },

            "size_in_bytes": {
                Type: schema.TypeInt,
                Computed: true,
            },

            "type": {
                Type: schema.TypeString,
                Computed: true,
            },

            "version": {
                Type: schema.TypeString,
                Computed: true,
            },
        },
    }
}

func resourceArmModuleCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).moduleClient
    ctx, cancel := timeouts.ForCreate(meta.(*ArmClient).StopContext, d)
    defer cancel()

    resourceGroupName := d.Get("resource_group").(string)
    automationAccountName := d.Get("automation_account_name").(string)
    name := d.Get("module_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroupName, automationAccountName, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Module (Module Name %q / Automation Account Name %q / Resource Group %q): %+v", name, automationAccountName, resourceGroupName, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_module", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    contentLink := d.Get("content_link").([]interface{})
    name := d.Get("name").(string)
    tags := d.Get("tags").(map[string]interface{})

    parameters := automation.ModuleUpdateParameters{
        Location: utils.String(location),
        Name: utils.String(name),
        ModuleUpdateProperties: &automation.ModuleUpdateProperties{
            ContentLink: expandArmModuleContentLink(contentLink),
        },
        Tags: tags.Expand(tags),
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroupName, automationAccountName, name, parameters); err != nil {
        return fmt.Errorf("Error creating Module (Module Name %q / Automation Account Name %q / Resource Group %q): %+v", name, automationAccountName, resourceGroupName, err)
    }


    resp, err := client.Get(ctx, resourceGroupName, automationAccountName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Module (Module Name %q / Automation Account Name %q / Resource Group %q): %+v", name, automationAccountName, resourceGroupName, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Module (Module Name %q / Automation Account Name %q / Resource Group %q) ID", name, automationAccountName, resourceGroupName)
    }
    d.SetId(*resp.ID)

    return resourceArmModuleRead(d, meta)
}

func resourceArmModuleRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).moduleClient
    ctx, cancel := timeouts.ForRead(meta.(*ArmClient).StopContext, d)
    defer cancel()

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroupName := id.ResourceGroup
    automationAccountName := id.Path["automationAccounts"]
    name := id.Path["modules"]

    resp, err := client.Get(ctx, resourceGroupName, automationAccountName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Module %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Module (Module Name %q / Automation Account Name %q / Resource Group %q): %+v", name, automationAccountName, resourceGroupName, err)
    }


    d.Set("resource_group", resourceGroupName)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if moduleUpdateProperties := resp.ModuleUpdateProperties; moduleUpdateProperties != nil {
        d.Set("activity_count", int(*moduleUpdateProperties.ActivityCount))
        if err := d.Set("content_link", flattenArmModuleContentLink(moduleUpdateProperties.ContentLink)); err != nil {
            return fmt.Errorf("Error setting `content_link`: %+v", err)
        }
        d.Set("creation_time", (moduleUpdateProperties.CreationTime).String())
        d.Set("description", moduleUpdateProperties.Description)
        if err := d.Set("error", flattenArmModuleModuleErrorInfo(moduleUpdateProperties.Error)); err != nil {
            return fmt.Errorf("Error setting `error`: %+v", err)
        }
        d.Set("is_composite", moduleUpdateProperties.IsComposite)
        d.Set("is_global", moduleUpdateProperties.IsGlobal)
        d.Set("last_modified_time", (moduleUpdateProperties.LastModifiedTime).String())
        d.Set("provisioning_state", string(moduleUpdateProperties.ProvisioningState))
        d.Set("size_in_bytes", int(*moduleUpdateProperties.SizeInBytes))
        d.Set("version", moduleUpdateProperties.Version)
    }
    d.Set("automation_account_name", automationAccountName)
    d.Set("etag", resp.Etag)
    d.Set("id", resp.ID)
    d.Set("module_name", name)
    d.Set("name", resp.Name)
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}

func resourceArmModuleUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).moduleClient
    ctx, cancel := timeouts.ForUpdate(meta.(*ArmClient).StopContext, d)
    defer cancel()

      resourceGroupName := d.Get("resource_group").(string)
    location := azure.NormalizeLocation(d.Get("location").(string))
    automationAccountName := d.Get("automation_account_name").(string)
    contentLink := d.Get("content_link").([]interface{})
    name := d.Get("module_name").(string)
    name := d.Get("name").(string)
    tags := d.Get("tags").(map[string]interface{})

    parameters := automation.ModuleUpdateParameters{
        Location: utils.String(location),
        Name: utils.String(name),
        ModuleUpdateProperties: &automation.ModuleUpdateProperties{
            ContentLink: expandArmModuleContentLink(contentLink),
        },
        Tags: tags.Expand(tags),
    }


    if _, err := client.Update(ctx, resourceGroupName, automationAccountName, name, parameters); err != nil {
        return fmt.Errorf("Error updating Module (Module Name %q / Automation Account Name %q / Resource Group %q): %+v", name, automationAccountName, resourceGroupName, err)
    }

    return resourceArmModuleRead(d, meta)
}

func resourceArmModuleDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).moduleClient
    ctx, cancel := timeouts.ForDelete(meta.(*ArmClient).StopContext, d)
    defer cancel()


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroupName := id.ResourceGroup
    automationAccountName := id.Path["automationAccounts"]
    name := id.Path["modules"]

    if _, err := client.Delete(ctx, resourceGroupName, automationAccountName, name); err != nil {
        return fmt.Errorf("Error deleting Module (Module Name %q / Automation Account Name %q / Resource Group %q): %+v", name, automationAccountName, resourceGroupName, err)
    }

    return nil
}

func expandArmModuleContentLink(input []interface{}) *automation.ContentLink {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    uRI := v["uri"].(string)
    contentHash := v["content_hash"].([]interface{})
    version := v["version"].(string)

    result := automation.ContentLink{
        ContentHash: expandArmModuleContentHash(contentHash),
        URI: utils.String(uRI),
        Version: utils.String(version),
    }
    return &result
}

func expandArmModuleContentHash(input []interface{}) *automation.ContentHash {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    algorithm := v["algorithm"].(string)
    value := v["value"].(string)

    result := automation.ContentHash{
        Algorithm: utils.String(algorithm),
        Value: utils.String(value),
    }
    return &result
}


func flattenArmModuleContentLink(input *automation.ContentLink) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    result["content_hash"] = flattenArmModuleContentHash(input.ContentHash)
    if uri := input.URI; uri != nil {
        result["uri"] = *uri
    }
    if version := input.Version; version != nil {
        result["version"] = *version
    }

    return []interface{}{result}
}

func flattenArmModuleModuleErrorInfo(input *automation.ModuleErrorInfo) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if code := input.Code; code != nil {
        result["code"] = *code
    }
    if message := input.Message; message != nil {
        result["message"] = *message
    }

    return []interface{}{result}
}

func flattenArmModuleContentHash(input *automation.ContentHash) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if algorithm := input.Algorithm; algorithm != nil {
        result["algorithm"] = *algorithm
    }
    if value := input.Value; value != nil {
        result["value"] = *value
    }

    return []interface{}{result}
}
