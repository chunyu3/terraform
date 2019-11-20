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

            "managed_by": {
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
    access := d.Get("access").(string)
    diskSizeGb := d.Get("disk_size_gb").(int)
    durationInSeconds := d.Get("duration_in_seconds").(int)
    encryptionSettings := d.Get("encryption_settings").([]interface{})
    osType := d.Get("os_type").(string)
    sku := d.Get("sku").([]interface{})
    t := d.Get("tags").(map[string]interface{})

    snapshot := compute.SnapshotUpdate{
        Access: compute.AccessLevel(access),
        DurationInSeconds: utils.Int32(int32(durationInSeconds)),
        Location: utils.String(location),
        DiskUpdateProperties: &compute.DiskUpdateProperties{
            DiskSizeGB: utils.Int32(int32(diskSizeGb)),
            EncryptionSettings: expandArmSnapshotEncryptionSettings(encryptionSettings),
            OsType: compute.OperatingSystemTypes(osType),
        },
        Sku: expandArmSnapshotDiskSku(sku),
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
    d.Set("managed_by", resp.ManagedBy)
    d.Set("type", resp.Type)

    return nil
}

func resourceArmSnapshotUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).snapshotsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    access := d.Get("access").(string)
    diskSizeGb := d.Get("disk_size_gb").(int)
    durationInSeconds := d.Get("duration_in_seconds").(int)
    encryptionSettings := d.Get("encryption_settings").([]interface{})
    osType := d.Get("os_type").(string)
    sku := d.Get("sku").([]interface{})
    t := d.Get("tags").(map[string]interface{})

    snapshot := compute.SnapshotUpdate{
        Access: compute.AccessLevel(access),
        DurationInSeconds: utils.Int32(int32(durationInSeconds)),
        DiskUpdateProperties: &compute.DiskUpdateProperties{
            DiskSizeGB: utils.Int32(int32(diskSizeGb)),
            EncryptionSettings: expandArmSnapshotEncryptionSettings(encryptionSettings),
            OsType: compute.OperatingSystemTypes(osType),
        },
        Sku: expandArmSnapshotDiskSku(sku),
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

func expandArmSnapshotEncryptionSettings(input []interface{}) *compute.EncryptionSettings {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    enabled := v["enabled"].(bool)
    diskEncryptionKey := v["disk_encryption_key"].([]interface{})
    keyEncryptionKey := v["key_encryption_key"].([]interface{})

    result := compute.EncryptionSettings{
        DiskEncryptionKey: expandArmSnapshotKeyVaultAndSecretReference(diskEncryptionKey),
        Enabled: utils.Bool(enabled),
        KeyEncryptionKey: expandArmSnapshotKeyVaultAndKeyReference(keyEncryptionKey),
    }
    return &result
}

func expandArmSnapshotDiskSku(input []interface{}) *compute.DiskSku {
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

func expandArmSnapshotKeyVaultAndSecretReference(input []interface{}) *compute.KeyVaultAndSecretReference {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    sourceVault := v["source_vault"].([]interface{})
    secretUrl := v["secret_url"].(string)

    result := compute.KeyVaultAndSecretReference{
        SecretURL: utils.String(secretUrl),
        SourceVault: expandArmSnapshotSourceVault(sourceVault),
    }
    return &result
}

func expandArmSnapshotKeyVaultAndKeyReference(input []interface{}) *compute.KeyVaultAndKeyReference {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    sourceVault := v["source_vault"].([]interface{})
    keyUrl := v["key_url"].(string)

    result := compute.KeyVaultAndKeyReference{
        KeyURL: utils.String(keyUrl),
        SourceVault: expandArmSnapshotSourceVault(sourceVault),
    }
    return &result
}

func expandArmSnapshotSourceVault(input []interface{}) *compute.SourceVault {
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
