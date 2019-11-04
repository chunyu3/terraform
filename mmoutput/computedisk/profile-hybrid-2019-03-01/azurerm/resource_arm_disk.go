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



func resourceArmDisk() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmDiskCreate,
        Read: resourceArmDiskRead,
        Update: resourceArmDiskUpdate,
        Delete: resourceArmDiskDelete,

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

            "disk_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "disk_size_gb": {
                Type: schema.TypeInt,
                Optional: true,
            },

            "encryption_settings": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
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
                                    "source_vault": {
                                        Type: schema.TypeList,
                                        Required: true,
                                        MaxItems: 1,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "id": {
                                                    Type: schema.TypeString,
                                                    Optional: true,
                                                },
                                            },
                                        },
                                    },
                                },
                            },
                        },
                        "enabled": {
                            Type: schema.TypeBool,
                            Optional: true,
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
                                    "source_vault": {
                                        Type: schema.TypeList,
                                        Required: true,
                                        MaxItems: 1,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "id": {
                                                    Type: schema.TypeString,
                                                    Optional: true,
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
                            }, false),
                            Default: string(compute.Standard_LRS),
                        },
                    },
                },
            },

            "zones": {
                Type: schema.TypeList,
                Optional: true,
                ForceNew: true,
                Elem: &schema.Schema{
                    Type: schema.TypeString,
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

func resourceArmDiskCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).disksClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    diskName := d.Get("disk_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, diskName)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Disk (Disk Name %q / Resource Group %q): %+v", diskName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_disk", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    creationData := d.Get("creation_data").([]interface{})
    diskSizeGb := d.Get("disk_size_gb").(int)
    encryptionSettings := d.Get("encryption_settings").([]interface{})
    osType := d.Get("os_type").(string)
    sku := d.Get("sku").([]interface{})
    zones := d.Get("zones").([]interface{})
    t := d.Get("tags").(map[string]interface{})

    disk := compute.Disk{
        Location: utils.String(location),
        DiskProperties: &compute.DiskProperties{
            CreationData: expandArmDiskCreationData(creationData),
            DiskSizeGb: utils.Int32(int32(diskSizeGb)),
            EncryptionSettings: expandArmDiskEncryptionSettings(encryptionSettings),
            OsType: compute.OperatingSystemTypes(osType),
        },
        Sku: expandArmDiskDiskSku(sku),
        Tags: tags.Expand(t),
        Zones: utils.ExpandStringSlice(zones),
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, diskName, disk)
    if err != nil {
        return fmt.Errorf("Error creating Disk (Disk Name %q / Resource Group %q): %+v", diskName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Disk (Disk Name %q / Resource Group %q): %+v", diskName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, diskName)
    if err != nil {
        return fmt.Errorf("Error retrieving Disk (Disk Name %q / Resource Group %q): %+v", diskName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Disk (Disk Name %q / Resource Group %q) ID", diskName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmDiskRead(d, meta)
}

func resourceArmDiskRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).disksClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    diskName := id.Path["disks"]

    resp, err := client.Get(ctx, resourceGroup, diskName)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Disk %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Disk (Disk Name %q / Resource Group %q): %+v", diskName, resourceGroup, err)
    }


    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if diskProperties := resp.DiskProperties; diskProperties != nil {
        if err := d.Set("creation_data", flattenArmDiskCreationData(diskProperties.CreationData)); err != nil {
            return fmt.Errorf("Error setting `creation_data`: %+v", err)
        }
        d.Set("disk_size_gb", int(*diskProperties.DiskSizeGb))
        if err := d.Set("encryption_settings", flattenArmDiskEncryptionSettings(diskProperties.EncryptionSettings)); err != nil {
            return fmt.Errorf("Error setting `encryption_settings`: %+v", err)
        }
        d.Set("os_type", string(diskProperties.OsType))
        d.Set("provisioning_state", diskProperties.ProvisioningState)
        d.Set("time_created", (diskProperties.TimeCreated).String())
    }
    d.Set("disk_name", diskName)
    d.Set("managed_by", resp.ManagedBy)
    if err := d.Set("sku", flattenArmDiskDiskSku(resp.Sku)); err != nil {
        return fmt.Errorf("Error setting `sku`: %+v", err)
    }
    d.Set("type", resp.Type)
    d.Set("zones", utils.FlattenStringSlice(resp.Zones))

    return tags.FlattenAndSet(d, resp.Tags)
}

func resourceArmDiskUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).disksClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    creationData := d.Get("creation_data").([]interface{})
    diskName := d.Get("disk_name").(string)
    diskSizeGb := d.Get("disk_size_gb").(int)
    encryptionSettings := d.Get("encryption_settings").([]interface{})
    osType := d.Get("os_type").(string)
    sku := d.Get("sku").([]interface{})
    zones := d.Get("zones").([]interface{})
    t := d.Get("tags").(map[string]interface{})

    disk := compute.Disk{
        Location: utils.String(location),
        DiskProperties: &compute.DiskProperties{
            CreationData: expandArmDiskCreationData(creationData),
            DiskSizeGb: utils.Int32(int32(diskSizeGb)),
            EncryptionSettings: expandArmDiskEncryptionSettings(encryptionSettings),
            OsType: compute.OperatingSystemTypes(osType),
        },
        Sku: expandArmDiskDiskSku(sku),
        Tags: tags.Expand(t),
        Zones: utils.ExpandStringSlice(zones),
    }


    future, err := client.Update(ctx, resourceGroup, diskName, disk)
    if err != nil {
        return fmt.Errorf("Error updating Disk (Disk Name %q / Resource Group %q): %+v", diskName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for update of Disk (Disk Name %q / Resource Group %q): %+v", diskName, resourceGroup, err)
    }

    return resourceArmDiskRead(d, meta)
}

func resourceArmDiskDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).disksClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    diskName := id.Path["disks"]

    future, err := client.Delete(ctx, resourceGroup, diskName)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Disk (Disk Name %q / Resource Group %q): %+v", diskName, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Disk (Disk Name %q / Resource Group %q): %+v", diskName, resourceGroup, err)
        }
    }

    return nil
}

func expandArmDiskCreationData(input []interface{}) *compute.CreationData {
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
        ImageReference: expandArmDiskImageDiskReference(imageReference),
        SourceResourceID: utils.String(sourceResourceId),
        SourceUri: utils.String(sourceUri),
        StorageAccountID: utils.String(storageAccountId),
    }
    return &result
}

func expandArmDiskEncryptionSettings(input []interface{}) *compute.EncryptionSettings {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    enabled := v["enabled"].(bool)
    diskEncryptionKey := v["disk_encryption_key"].([]interface{})
    keyEncryptionKey := v["key_encryption_key"].([]interface{})

    result := compute.EncryptionSettings{
        DiskEncryptionKey: expandArmDiskKeyVaultAndSecretReference(diskEncryptionKey),
        Enabled: utils.Bool(enabled),
        KeyEncryptionKey: expandArmDiskKeyVaultAndKeyReference(keyEncryptionKey),
    }
    return &result
}

func expandArmDiskDiskSku(input []interface{}) *compute.DiskSku {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    name := v["name"].(string)

    result := compute.DiskSku{
        Name: compute.StorageAccountTypes(name),
    }
    return &result
}

func expandArmDiskImageDiskReference(input []interface{}) *compute.ImageDiskReference {
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

func expandArmDiskKeyVaultAndSecretReference(input []interface{}) *compute.KeyVaultAndSecretReference {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    sourceVault := v["source_vault"].([]interface{})
    secretUrl := v["secret_url"].(string)

    result := compute.KeyVaultAndSecretReference{
        SecretURL: utils.String(secretUrl),
        SourceVault: expandArmDiskSourceVault(sourceVault),
    }
    return &result
}

func expandArmDiskKeyVaultAndKeyReference(input []interface{}) *compute.KeyVaultAndKeyReference {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    sourceVault := v["source_vault"].([]interface{})
    keyUrl := v["key_url"].(string)

    result := compute.KeyVaultAndKeyReference{
        KeyURL: utils.String(keyUrl),
        SourceVault: expandArmDiskSourceVault(sourceVault),
    }
    return &result
}

func expandArmDiskSourceVault(input []interface{}) *compute.SourceVault {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    id := v["id"].(string)

    result := compute.SourceVault{
        ID: utils.String(id),
    }
    return &result
}


func flattenArmDiskCreationData(input *compute.CreationData) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    result["create_option"] = string(input.CreateOption)
    result["image_reference"] = flattenArmDiskImageDiskReference(input.ImageReference)
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

func flattenArmDiskEncryptionSettings(input *compute.EncryptionSettings) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    result["disk_encryption_key"] = flattenArmDiskKeyVaultAndSecretReference(input.DiskEncryptionKey)
    if enabled := input.Enabled; enabled != nil {
        result["enabled"] = *enabled
    }
    result["key_encryption_key"] = flattenArmDiskKeyVaultAndKeyReference(input.KeyEncryptionKey)

    return []interface{}{result}
}

func flattenArmDiskDiskSku(input *compute.DiskSku) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    result["name"] = string(input.Name)

    return []interface{}{result}
}

func flattenArmDiskImageDiskReference(input *compute.ImageDiskReference) []interface{} {
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

func flattenArmDiskKeyVaultAndSecretReference(input *compute.KeyVaultAndSecretReference) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if secretUrl := input.SecretURL; secretUrl != nil {
        result["secret_url"] = *secretUrl
    }
    result["source_vault"] = flattenArmDiskSourceVault(input.SourceVault)

    return []interface{}{result}
}

func flattenArmDiskKeyVaultAndKeyReference(input *compute.KeyVaultAndKeyReference) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if keyUrl := input.KeyURL; keyUrl != nil {
        result["key_url"] = *keyUrl
    }
    result["source_vault"] = flattenArmDiskSourceVault(input.SourceVault)

    return []interface{}{result}
}

func flattenArmDiskSourceVault(input *compute.SourceVault) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if id := input.ID; id != nil {
        result["id"] = *id
    }

    return []interface{}{result}
}
