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



func resourceArmVolume() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmVolumeCreateUpdate,
        Read: resourceArmVolumeRead,
        Update: resourceArmVolumeCreateUpdate,
        Delete: resourceArmVolumeDelete,

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

            "provider": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "azure_file_parameters": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "account_name": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "share_name": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "account_key": {
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

            "provisioning_state": {
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

func resourceArmVolumeCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).volumeClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Volume %q (Resource Group %q): %+v", name, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_volume", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    azureFileParameters := d.Get("azure_file_parameters").([]interface{})
    description := d.Get("description").(string)
    provider := d.Get("provider").(string)
    t := d.Get("tags").(map[string]interface{})

    volumeResourceDescription := servicefabricmeshrestapis.VolumeResourceDescription{
        Location: utils.String(location),
        VolumeResourceProperties: &servicefabricmeshrestapis.VolumeResourceProperties{
            AzureFileParameters: expandArmVolumeVolumeProviderParametersAzureFile(azureFileParameters),
            Description: utils.String(description),
            Provider: utils.String(provider),
        },
        Tags: tags.Expand(t),
    }


    if _, err := client.Create(ctx, resourceGroup, name, volumeResourceDescription); err != nil {
        return fmt.Errorf("Error creating Volume %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Volume %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Volume %q (Resource Group %q) ID", name, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmVolumeRead(d, meta)
}

func resourceArmVolumeRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).volumeClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["volumes"]

    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Volume %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Volume %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if volumeResourceProperties := resp.VolumeResourceProperties; volumeResourceProperties != nil {
        if err := d.Set("azure_file_parameters", flattenArmVolumeVolumeProviderParametersAzureFile(volumeResourceProperties.AzureFileParameters)); err != nil {
            return fmt.Errorf("Error setting `azure_file_parameters`: %+v", err)
        }
        d.Set("description", volumeResourceProperties.Description)
        d.Set("provider", volumeResourceProperties.Provider)
        d.Set("provisioning_state", volumeResourceProperties.ProvisioningState)
    }
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}


func resourceArmVolumeDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).volumeClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["volumes"]

    if _, err := client.Delete(ctx, resourceGroup, name); err != nil {
        return fmt.Errorf("Error deleting Volume %q (Resource Group %q): %+v", name, resourceGroup, err)
    }

    return nil
}

func expandArmVolumeVolumeProviderParametersAzureFile(input []interface{}) *servicefabricmeshrestapis.VolumeProviderParametersAzureFile {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    accountName := v["account_name"].(string)
    accountKey := v["account_key"].(string)
    shareName := v["share_name"].(string)

    result := servicefabricmeshrestapis.VolumeProviderParametersAzureFile{
        AccountKey: utils.String(accountKey),
        AccountName: utils.String(accountName),
        ShareName: utils.String(shareName),
    }
    return &result
}


func flattenArmVolumeVolumeProviderParametersAzureFile(input *servicefabricmeshrestapis.VolumeProviderParametersAzureFile) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if accountKey := input.AccountKey; accountKey != nil {
        result["account_key"] = *accountKey
    }
    if accountName := input.AccountName; accountName != nil {
        result["account_name"] = *accountName
    }
    if shareName := input.ShareName; shareName != nil {
        result["share_name"] = *shareName
    }

    return []interface{}{result}
}
