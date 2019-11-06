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



func resourceArmSnapshot() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmSnapshotCreate,
        Read: resourceArmSnapshotRead,
        Update: resourceArmSnapshotUpdate,
        Delete: resourceArmSnapshotDelete,

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

            "creation_data": {
                Type: schema.TypeList,
                Required: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "create_option": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(compute.Empty),
                                string(compute.Attach),
                                string(compute.FromImage),
                                string(compute.Import),
                                string(compute.Copy),
                                string(compute.Restore),
                                string(compute.Upload),
                            }, false),
                        },
                        "image_reference": {
                            Type: schema.TypeList,
                            Optional: true,
                            MaxItems: 1,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "id": {
                                        Type: schema.TypeString,
                                        Required: true,
                                        ValidateFunc: validate.NoEmptyStrings,
                                    },
                                    "lun": {
                                        Type: schema.TypeInt,
                                        Optional: true,
                                    },
                                },
                            },
                        },
                        "source_resource_id": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "source_uri": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "storage_account_id": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },

            "disk_size_gb": {
                Type: schema.TypeInt,
                Optional: true,
            },

            "encryption_settings_collection": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "enabled": {
                            Type: schema.TypeBool,
                            Required: true,
                        },
                        "encryption_settings": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "disk_encryption_key": {
                                        Type: schema.TypeList,
                                        Optional: true,
                                        MaxItems: 1,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "secret_url": {
                                                    Type: schema.TypeString,
                                                    Required: true,
                                                    ValidateFunc: validate.NoEmptyStrings,
                                                },
                                            },
                                        },
                                    },
                                    "key_encryption_key": {
                                        Type: schema.TypeList,
                                        Optional: true,
                                        MaxItems: 1,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "key_url": {
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
                    },
                },
            },

            "hyper_vgeneration": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(compute.V1),
                    string(compute.V2),
                }, false),
                Default: string(compute.V1),
            },

            "os_type": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(compute.Windows),
                    string(compute.Linux),
                }, false),
                Default: string(compute.Windows),
            },

            "sku": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "name": {
                            Type: schema.TypeString,
                            Optional: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(compute.Standard_LRS),
                                string(compute.Premium_LRS),
                                string(compute.Standard_ZRS),
                            }, false),
                            Default: string(compute.Standard_LRS),
                        },
                    },
                },
            },

            "managed_by": {
                Type: schema.TypeString,
                Computed: true,
            },

            "provisioning_state": {
                Type: schema.TypeString,
                Computed: true,
            },

            "time_created": {
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

func resourceArmSnapshotCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).snapshotsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Snapshot %q (Resource Group %q): %+v", name, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_snapshot", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    creationData := d.Get("creation_data").([]interface{})
    diskSizeGb := d.Get("disk_size_gb").(int)
    encryptionSettingsCollection := d.Get("encryption_settings_collection").([]interface{})
    hyperVgeneration := d.Get("hyper_vgeneration").(string)
    osType := d.Get("os_type").(string)
    sku := d.Get("sku").([]interface{})
    t := d.Get("tags").(map[string]interface{})

    snapshot := compute.Snapshot{
        Location: utils.String(location),
        SnapshotProperties: &compute.SnapshotProperties{
            CreationData: expandArmSnapshotCreationData(creationData),
            DiskSizeGb: utils.Int32(int32(diskSizeGb)),
            EncryptionSettingsCollection: expandArmSnapshotEncryptionSettingsCollection(encryptionSettingsCollection),
            HyperVgeneration: compute.HyperVGeneration(hyperVgeneration),
            OsType: compute.OperatingSystemTypes(osType),
        },
        Sku: expandArmSnapshotSnapshotSku(sku),
        Tags: tags.Expand(t),
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, name, snapshot)
    if err != nil {
        return fmt.Errorf("Error creating Snapshot %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Snapshot %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Snapshot %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Snapshot %q (Resource Group %q) ID", name, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmSnapshotRead(d, meta)
}

func resourceArmSnapshotRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).snapshotsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["snapshots"]

    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Snapshot %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Snapshot %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if snapshotProperties := resp.SnapshotProperties; snapshotProperties != nil {
        if err := d.Set("creation_data", flattenArmSnapshotCreationData(snapshotProperties.CreationData)); err != nil {
            return fmt.Errorf("Error setting `creation_data`: %+v", err)
        }
        d.Set("disk_size_gb", int(*snapshotProperties.DiskSizeGb))
        if err := d.Set("encryption_settings_collection", flattenArmSnapshotEncryptionSettingsCollection(snapshotProperties.EncryptionSettingsCollection)); err != nil {
            return fmt.Errorf("Error setting `encryption_settings_collection`: %+v", err)
        }
        d.Set("hyper_vgeneration", string(snapshotProperties.HyperVgeneration))
        d.Set("os_type", string(snapshotProperties.OsType))
        d.Set("provisioning_state", snapshotProperties.ProvisioningState)
        d.Set("time_created", (snapshotProperties.TimeCreated).String())
    }
    d.Set("managed_by", resp.ManagedBy)
    if err := d.Set("sku", flattenArmSnapshotSnapshotSku(resp.Sku)); err != nil {
        return fmt.Errorf("Error setting `sku`: %+v", err)
    }
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}

func resourceArmSnapshotUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).snapshotsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    creationData := d.Get("creation_data").([]interface{})
    diskSizeGb := d.Get("disk_size_gb").(int)
    encryptionSettingsCollection := d.Get("encryption_settings_collection").([]interface{})
    hyperVgeneration := d.Get("hyper_vgeneration").(string)
    osType := d.Get("os_type").(string)
    sku := d.Get("sku").([]interface{})
    t := d.Get("tags").(map[string]interface{})

    snapshot := compute.Snapshot{
        Location: utils.String(location),
        SnapshotProperties: &compute.SnapshotProperties{
            CreationData: expandArmSnapshotCreationData(creationData),
            DiskSizeGb: utils.Int32(int32(diskSizeGb)),
            EncryptionSettingsCollection: expandArmSnapshotEncryptionSettingsCollection(encryptionSettingsCollection),
            HyperVgeneration: compute.HyperVGeneration(hyperVgeneration),
            OsType: compute.OperatingSystemTypes(osType),
        },
        Sku: expandArmSnapshotSnapshotSku(sku),
        Tags: tags.Expand(t),
    }


    future, err := client.Update(ctx, resourceGroup, name, snapshot)
    if err != nil {
        return fmt.Errorf("Error updating Snapshot %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for update of Snapshot %q (Resource Group %q): %+v", name, resourceGroup, err)
    }

    return resourceArmSnapshotRead(d, meta)
}

func resourceArmSnapshotDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).snapshotsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["snapshots"]

    future, err := client.Delete(ctx, resourceGroup, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Snapshot %q (Resource Group %q): %+v", name, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Snapshot %q (Resource Group %q): %+v", name, resourceGroup, err)
        }
    }

    return nil
}

func expandArmSnapshotCreationData(input []interface{}) *compute.CreationData {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    createOption := v["create_option"].(string)
    storageAccountId := v["storage_account_id"].(string)
    imageReference := v["image_reference"].([]interface{})
    sourceUri := v["source_uri"].(string)
    sourceResourceId := v["source_resource_id"].(string)

    result := compute.CreationData{
        CreateOption: compute.DiskCreateOption(createOption),
        ImageReference: expandArmSnapshotImageDiskReference(imageReference),
        SourceResourceID: utils.String(sourceResourceId),
        SourceUri: utils.String(sourceUri),
        StorageAccountID: utils.String(storageAccountId),
    }
    return &result
}

func expandArmSnapshotEncryptionSettingsCollection(input []interface{}) *compute.EncryptionSettingsCollection {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    enabled := v["enabled"].(bool)
    encryptionSettings := v["encryption_settings"].([]interface{})

    result := compute.EncryptionSettingsCollection{
        Enabled: utils.Bool(enabled),
        EncryptionSettings: expandArmSnapshotEncryptionSettingsElement(encryptionSettings),
    }
    return &result
}

func expandArmSnapshotSnapshotSku(input []interface{}) *compute.SnapshotSku {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    name := v["name"].(string)

    result := compute.SnapshotSku{
        Name: compute.SnapshotStorageAccountTypes(name),
    }
    return &result
}

func expandArmSnapshotImageDiskReference(input []interface{}) *compute.ImageDiskReference {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    id := v["id"].(string)
    lun := v["lun"].(int)

    result := compute.ImageDiskReference{
        ID: utils.String(id),
        Lun: utils.Int32(int32(lun)),
    }
    return &result
}

func expandArmSnapshotEncryptionSettingsElement(input []interface{}) *[]compute.EncryptionSettingsElement {
    results := make([]compute.EncryptionSettingsElement, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        diskEncryptionKey := v["disk_encryption_key"].([]interface{})
        keyEncryptionKey := v["key_encryption_key"].([]interface{})

        result := compute.EncryptionSettingsElement{
            DiskEncryptionKey: expandArmSnapshotKeyVaultAndSecretReference(diskEncryptionKey),
            KeyEncryptionKey: expandArmSnapshotKeyVaultAndKeyReference(keyEncryptionKey),
        }

        results = append(results, result)
    }
    return &results
}

func expandArmSnapshotKeyVaultAndSecretReference(input []interface{}) *compute.KeyVaultAndSecretReference {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    secretUrl := v["secret_url"].(string)

    result := compute.KeyVaultAndSecretReference{
        SecretURL: utils.String(secretUrl),
    }
    return &result
}

func expandArmSnapshotKeyVaultAndKeyReference(input []interface{}) *compute.KeyVaultAndKeyReference {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    keyUrl := v["key_url"].(string)

    result := compute.KeyVaultAndKeyReference{
        KeyURL: utils.String(keyUrl),
    }
    return &result
}


func flattenArmSnapshotCreationData(input *compute.CreationData) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    result["create_option"] = string(input.CreateOption)
    result["image_reference"] = flattenArmSnapshotImageDiskReference(input.ImageReference)
    if sourceResourceId := input.SourceResourceID; sourceResourceId != nil {
        result["source_resource_id"] = *sourceResourceId
    }
    if sourceUri := input.SourceUri; sourceUri != nil {
        result["source_uri"] = *sourceUri
    }
    if storageAccountId := input.StorageAccountID; storageAccountId != nil {
        result["storage_account_id"] = *storageAccountId
    }

    return []interface{}{result}
}

func flattenArmSnapshotEncryptionSettingsCollection(input *compute.EncryptionSettingsCollection) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if enabled := input.Enabled; enabled != nil {
        result["enabled"] = *enabled
    }
    result["encryption_settings"] = flattenArmSnapshotEncryptionSettingsElement(input.EncryptionSettings)

    return []interface{}{result}
}

func flattenArmSnapshotSnapshotSku(input *compute.SnapshotSku) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    result["name"] = string(input.Name)

    return []interface{}{result}
}

func flattenArmSnapshotImageDiskReference(input *compute.ImageDiskReference) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if id := input.ID; id != nil {
        result["id"] = *id
    }
    if lun := input.Lun; lun != nil {
        result["lun"] = int(*lun)
    }

    return []interface{}{result}
}

func flattenArmSnapshotEncryptionSettingsElement(input *[]compute.EncryptionSettingsElement) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        v["disk_encryption_key"] = flattenArmSnapshotKeyVaultAndSecretReference(item.DiskEncryptionKey)
        v["key_encryption_key"] = flattenArmSnapshotKeyVaultAndKeyReference(item.KeyEncryptionKey)

        results = append(results, v)
    }

    return results
}

func flattenArmSnapshotKeyVaultAndSecretReference(input *compute.KeyVaultAndSecretReference) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if secretUrl := input.SecretURL; secretUrl != nil {
        result["secret_url"] = *secretUrl
    }

    return []interface{}{result}
}

func flattenArmSnapshotKeyVaultAndKeyReference(input *compute.KeyVaultAndKeyReference) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if keyUrl := input.KeyURL; keyUrl != nil {
        result["key_url"] = *keyUrl
    }

    return []interface{}{result}
}
