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

            "account_type": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(compute.Standard_LRS),
                    string(compute.Premium_LRS),
                }, false),
                Default: string(compute.Standard_LRS),
            },

            "creation_data": {
                Type: schema.TypeList,
                Optional: true,
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
    accountType := d.Get("account_type").(string)
    creationData := d.Get("creation_data").([]interface{})
    diskSizeGb := d.Get("disk_size_gb").(int)
    durationInSeconds := d.Get("duration_in_seconds").(int)
    encryptionSettings := d.Get("encryption_settings").([]interface{})
    osType := d.Get("os_type").(string)
    t := d.Get("tags").(map[string]interface{})

    disk := compute.DiskUpdate{
        Access: compute.AccessLevel(access),
        DurationInSeconds: utils.Int32(int32(durationInSeconds)),
        Location: utils.String(location),
        DiskUpdateProperties: &compute.DiskUpdateProperties{
            AccountType: compute.StorageAccountTypes(accountType),
            CreationData: expandArmDiskCreationData(creationData),
            DiskSizeGB: utils.Int32(int32(diskSizeGb)),
            EncryptionSettings: expandArmDiskEncryptionSettings(encryptionSettings),
            OsType: compute.OperatingSystemTypes(osType),
        },
        Tags: tags.Expand(t),
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
    d.Set("type", resp.Type)

    return nil
}

func resourceArmDiskUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).disksClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    access := d.Get("access").(string)
    accountType := d.Get("account_type").(string)
    creationData := d.Get("creation_data").([]interface{})
    diskSizeGb := d.Get("disk_size_gb").(int)
    durationInSeconds := d.Get("duration_in_seconds").(int)
    encryptionSettings := d.Get("encryption_settings").([]interface{})
    osType := d.Get("os_type").(string)
    t := d.Get("tags").(map[string]interface{})

    disk := compute.DiskUpdate{
        Access: compute.AccessLevel(access),
        DurationInSeconds: utils.Int32(int32(durationInSeconds)),
        DiskUpdateProperties: &compute.DiskUpdateProperties{
            AccountType: compute.StorageAccountTypes(accountType),
            CreationData: expandArmDiskCreationData(creationData),
            DiskSizeGB: utils.Int32(int32(diskSizeGb)),
            EncryptionSettings: expandArmDiskEncryptionSettings(encryptionSettings),
            OsType: compute.OperatingSystemTypes(osType),
        },
        Tags: tags.Expand(t),
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
        SourceURI: utils.String(sourceUri),
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
