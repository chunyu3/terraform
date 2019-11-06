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



func resourceArmMediaService() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmMediaServiceCreate,
        Read: resourceArmMediaServiceRead,
        Update: resourceArmMediaServiceUpdate,
        Delete: resourceArmMediaServiceDelete,

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

            "storage_accounts": {
                Type: schema.TypeList,
                Optional: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "id": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "is_primary": {
                            Type: schema.TypeBool,
                            Required: true,
                        },
                    },
                },
            },

            "api_endpoints": {
                Type: schema.TypeList,
                Computed: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "endpoint": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "major_version": {
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

func resourceArmMediaServiceCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).mediaServiceClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Media Service %q (Resource Group %q): %+v", name, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_media_service", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    storageAccounts := d.Get("storage_accounts").([]interface{})
    t := d.Get("tags").(map[string]interface{})

    parameters := mediaservices.MediaService{
        Location: utils.String(location),
        MediaServiceProperties: &mediaservices.MediaServiceProperties{
            StorageAccounts: expandArmMediaServiceStorageAccount(storageAccounts),
        },
        Tags: tags.Expand(t),
    }


    if _, err := client.Create(ctx, resourceGroup, name, parameters); err != nil {
        return fmt.Errorf("Error creating Media Service %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Media Service %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Media Service %q (Resource Group %q) ID", name, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmMediaServiceRead(d, meta)
}

func resourceArmMediaServiceRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).mediaServiceClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["mediaservices"]

    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Media Service %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Media Service %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if mediaServiceProperties := resp.MediaServiceProperties; mediaServiceProperties != nil {
        if err := d.Set("api_endpoints", flattenArmMediaServiceApiEndpoint(mediaServiceProperties.ApiEndpoints)); err != nil {
            return fmt.Errorf("Error setting `api_endpoints`: %+v", err)
        }
        if err := d.Set("storage_accounts", flattenArmMediaServiceStorageAccount(mediaServiceProperties.StorageAccounts)); err != nil {
            return fmt.Errorf("Error setting `storage_accounts`: %+v", err)
        }
    }
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}

func resourceArmMediaServiceUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).mediaServiceClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    storageAccounts := d.Get("storage_accounts").([]interface{})
    t := d.Get("tags").(map[string]interface{})

    parameters := mediaservices.MediaService{
        Location: utils.String(location),
        MediaServiceProperties: &mediaservices.MediaServiceProperties{
            StorageAccounts: expandArmMediaServiceStorageAccount(storageAccounts),
        },
        Tags: tags.Expand(t),
    }


    if _, err := client.Update(ctx, resourceGroup, name, parameters); err != nil {
        return fmt.Errorf("Error updating Media Service %q (Resource Group %q): %+v", name, resourceGroup, err)
    }

    return resourceArmMediaServiceRead(d, meta)
}

func resourceArmMediaServiceDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).mediaServiceClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["mediaservices"]

    if _, err := client.Delete(ctx, resourceGroup, name); err != nil {
        return fmt.Errorf("Error deleting Media Service %q (Resource Group %q): %+v", name, resourceGroup, err)
    }

    return nil
}

func expandArmMediaServiceStorageAccount(input []interface{}) *[]mediaservices.StorageAccount {
    results := make([]mediaservices.StorageAccount, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        id := v["id"].(string)
        isPrimary := v["is_primary"].(bool)

        result := mediaservices.StorageAccount{
            ID: utils.String(id),
            IsPrimary: utils.Bool(isPrimary),
        }

        results = append(results, result)
    }
    return &results
}


func flattenArmMediaServiceApiEndpoint(input *[]mediaservices.ApiEndpoint) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})


        results = append(results, v)
    }

    return results
}

func flattenArmMediaServiceStorageAccount(input *[]mediaservices.StorageAccount) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        if id := item.ID; id != nil {
            v["id"] = *id
        }
        if isPrimary := item.IsPrimary; isPrimary != nil {
            v["is_primary"] = *isPrimary
        }

        results = append(results, v)
    }

    return results
}
