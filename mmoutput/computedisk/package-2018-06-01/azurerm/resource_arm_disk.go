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

            "access": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(compute.None),
                    string(compute.Read),
                }, false),
            },

            "duration_in_seconds": {
                Type: schema.TypeInt,
                Required: true,
                ForceNew: true,
            },

            "disk_iopsread_write": {
                Type: schema.TypeInt,
                Optional: true,
            },

            "disk_mbps_read_write": {
                Type: schema.TypeInt,
                Optional: true,
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
                                string(compute.StandardSSD_LRS),
                                string(compute.UltraSSD_LRS),
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

            "creation_data": {
                Type: schema.TypeList,
                Computed: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "create_option": {
                            Type: schema.TypeString,
                            Computed: true,
                        },
                        "image_reference": {
                            Type: schema.TypeList,
                            Computed: true,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "id": {
                                        Type: schema.TypeString,
                                        Computed: true,
                                    },
                                    "lun": {
                                        Type: schema.TypeInt,
                                        Computed: true,
                                    },
                                },
                            },
                        },
                        "source_resource_id": {
                            Type: schema.TypeString,
                            Computed: true,
                        },
                        "source_uri": {
                            Type: schema.TypeString,
                            Computed: true,
                        },
                        "storage_account_id": {
                            Type: schema.TypeString,
                            Computed: true,
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

func resourceArmDiskCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).disksClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Disk %q (Resource Group %q): %+v", name, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_disk", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    access := d.Get("access").(string)
    diskIopsreadWrite := d.Get("disk_iopsread_write").(int)
    diskMbpsReadWrite := d.Get("disk_mbps_read_write").(int)
    diskSizeGb := d.Get("disk_size_gb").(int)
    durationInSeconds := d.Get("duration_in_seconds").(int)
    encryptionSettings := d.Get("encryption_settings").([]interface{})
    osType := d.Get("os_type").(string)
    sku := d.Get("sku").([]interface{})
    zones := d.Get("zones").([]interface{})
    t := d.Get("tags").(map[string]interface{})

    disk := compute.DiskUpdate{
        Access: compute.AccessLevel(access),
        DurationInSeconds: utils.Int32(int32(durationInSeconds)),
        Location: utils.String(location),
        DiskUpdateProperties: &compute.DiskUpdateProperties{
            DiskIOPSReadWrite: utils.Int64(int64(diskIopsreadWrite)),
            DiskMBpsReadWrite: utils.Int32(int32(diskMbpsReadWrite)),
            DiskSizeGB: utils.Int32(int32(diskSizeGb)),
            EncryptionSettings: expandArmDiskEncryptionSettings(encryptionSettings),
            OsType: compute.OperatingSystemTypes(osType),
        },
        Sku: expandArmDiskDiskSku(sku),
        Tags: tags.Expand(t),
        Zones: utils.ExpandStringSlice(zones),
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, name, disk)
    if err != nil {
        return fmt.Errorf("Error creating Disk %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Disk %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Disk %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Disk %q (Resource Group %q) ID", name, resourceGroup)
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
    name := id.Path["disks"]

    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Disk %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Disk %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if diskUpdateProperties := resp.DiskUpdateProperties; diskUpdateProperties != nil {
        if err := d.Set("creation_data", flattenArmDiskCreationData(diskUpdateProperties.CreationData)); err != nil {
            return fmt.Errorf("Error setting `creation_data`: %+v", err)
        }
        d.Set("disk_iopsread_write", int(*diskUpdateProperties.DiskIOPSReadWrite))
        d.Set("disk_mbps_read_write", int(*diskUpdateProperties.DiskMBpsReadWrite))
        d.Set("disk_size_gb", int(*diskUpdateProperties.DiskSizeGB))
        if err := d.Set("encryption_settings", flattenArmDiskEncryptionSettings(diskUpdateProperties.EncryptionSettings)); err != nil {
            return fmt.Errorf("Error setting `encryption_settings`: %+v", err)
        }
        d.Set("os_type", string(diskUpdateProperties.OsType))
        d.Set("provisioning_state", diskUpdateProperties.ProvisioningState)
        d.Set("time_created", (diskUpdateProperties.TimeCreated).String())
    }
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

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    access := d.Get("access").(string)
    diskIopsreadWrite := d.Get("disk_iopsread_write").(int)
    diskMbpsReadWrite := d.Get("disk_mbps_read_write").(int)
    diskSizeGb := d.Get("disk_size_gb").(int)
    durationInSeconds := d.Get("duration_in_seconds").(int)
    encryptionSettings := d.Get("encryption_settings").([]interface{})
    osType := d.Get("os_type").(string)
    sku := d.Get("sku").([]interface{})
    zones := d.Get("zones").([]interface{})
    t := d.Get("tags").(map[string]interface{})

    disk := compute.DiskUpdate{
        Access: compute.AccessLevel(access),
        DurationInSeconds: utils.Int32(int32(durationInSeconds)),
        DiskUpdateProperties: &compute.DiskUpdateProperties{
            DiskIOPSReadWrite: utils.Int64(int64(diskIopsreadWrite)),
            DiskMBpsReadWrite: utils.Int32(int32(diskMbpsReadWrite)),
            DiskSizeGB: utils.Int32(int32(diskSizeGb)),
            EncryptionSettings: expandArmDiskEncryptionSettings(encryptionSettings),
            OsType: compute.OperatingSystemTypes(osType),
        },
        Sku: expandArmDiskDiskSku(sku),
        Tags: tags.Expand(t),
        Zones: utils.ExpandStringSlice(zones),
    }


    future, err := client.Update(ctx, resourceGroup, name, disk)
    if err != nil {
        return fmt.Errorf("Error updating Disk %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for update of Disk %q (Resource Group %q): %+v", name, resourceGroup, err)
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
    name := id.Path["disks"]

    future, err := client.Delete(ctx, resourceGroup, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Disk %q (Resource Group %q): %+v", name, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Disk %q (Resource Group %q): %+v", name, resourceGroup, err)
        }
    }

    return nil
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
        Name: compute.DiskStorageAccountTypes(name),
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
    if sourceUri := input.SourceURI; sourceUri != nil {
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
