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



func resourceArmImage() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmImageCreateUpdate,
        Read: resourceArmImageRead,
        Update: resourceArmImageCreateUpdate,
        Delete: resourceArmImageDelete,

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

            "image_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "source_virtual_machine": {
                Type: schema.TypeList,
                Optional: true,
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

            "storage_profile": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "os_disk": {
                            Type: schema.TypeList,
                            Required: true,
                            MaxItems: 1,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "os_state": {
                                        Type: schema.TypeString,
                                        Required: true,
                                        ValidateFunc: validation.StringInSlice([]string{
                                            string(compute.Generalized),
                                            string(compute.Specialized),
                                        }, false),
                                    },
                                    "os_type": {
                                        Type: schema.TypeString,
                                        Required: true,
                                        ValidateFunc: validation.StringInSlice([]string{
                                            string(compute.Windows),
                                            string(compute.Linux),
                                        }, false),
                                    },
                                    "blob_uri": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                    "caching": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                        ValidateFunc: validation.StringInSlice([]string{
                                            string(compute.None),
                                            string(compute.ReadOnly),
                                            string(compute.ReadWrite),
                                        }, false),
                                        Default: string(compute.None),
                                    },
                                    "disk_size_gb": {
                                        Type: schema.TypeInt,
                                        Optional: true,
                                    },
                                    "managed_disk": {
                                        Type: schema.TypeList,
                                        Optional: true,
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
                                    "snapshot": {
                                        Type: schema.TypeList,
                                        Optional: true,
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
                                    "storage_account_type": {
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
                        "data_disks": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "lun": {
                                        Type: schema.TypeInt,
                                        Required: true,
                                    },
                                    "blob_uri": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                    "caching": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                        ValidateFunc: validation.StringInSlice([]string{
                                            string(compute.None),
                                            string(compute.ReadOnly),
                                            string(compute.ReadWrite),
                                        }, false),
                                        Default: string(compute.None),
                                    },
                                    "disk_size_gb": {
                                        Type: schema.TypeInt,
                                        Optional: true,
                                    },
                                    "managed_disk": {
                                        Type: schema.TypeList,
                                        Optional: true,
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
                                    "snapshot": {
                                        Type: schema.TypeList,
                                        Optional: true,
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
                                    "storage_account_type": {
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
                    },
                },
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

func resourceArmImageCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).imagesClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    imageName := d.Get("image_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, imageName)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Image (Image Name %q / Resource Group %q): %+v", imageName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_image", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    sourceVirtualMachine := d.Get("source_virtual_machine").([]interface{})
    storageProfile := d.Get("storage_profile").([]interface{})
    t := d.Get("tags").(map[string]interface{})

    parameters := compute.Image{
        Location: utils.String(location),
        ImageProperties: &compute.ImageProperties{
            SourceVirtualMachine: expandArmImageSubResource(sourceVirtualMachine),
            StorageProfile: expandArmImageImageStorageProfile(storageProfile),
        },
        Tags: tags.Expand(t),
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, imageName, parameters)
    if err != nil {
        return fmt.Errorf("Error creating Image (Image Name %q / Resource Group %q): %+v", imageName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Image (Image Name %q / Resource Group %q): %+v", imageName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, imageName)
    if err != nil {
        return fmt.Errorf("Error retrieving Image (Image Name %q / Resource Group %q): %+v", imageName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Image (Image Name %q / Resource Group %q) ID", imageName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmImageRead(d, meta)
}

func resourceArmImageRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).imagesClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    imageName := id.Path["images"]

    resp, err := client.Get(ctx, resourceGroup, imageName)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Image %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Image (Image Name %q / Resource Group %q): %+v", imageName, resourceGroup, err)
    }


    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    d.Set("image_name", imageName)
    if imageProperties := resp.ImageProperties; imageProperties != nil {
        d.Set("provisioning_state", imageProperties.ProvisioningState)
        if err := d.Set("source_virtual_machine", flattenArmImageSubResource(imageProperties.SourceVirtualMachine)); err != nil {
            return fmt.Errorf("Error setting `source_virtual_machine`: %+v", err)
        }
        if err := d.Set("storage_profile", flattenArmImageImageStorageProfile(imageProperties.StorageProfile)); err != nil {
            return fmt.Errorf("Error setting `storage_profile`: %+v", err)
        }
    }
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}


func resourceArmImageDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).imagesClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    imageName := id.Path["images"]

    future, err := client.Delete(ctx, resourceGroup, imageName)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Image (Image Name %q / Resource Group %q): %+v", imageName, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Image (Image Name %q / Resource Group %q): %+v", imageName, resourceGroup, err)
        }
    }

    return nil
}

func expandArmImageSubResource(input []interface{}) *compute.SubResource {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    id := v["id"].(string)

    result := compute.SubResource{
        ID: utils.String(id),
    }
    return &result
}

func expandArmImageImageStorageProfile(input []interface{}) *compute.ImageStorageProfile {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    osDisk := v["os_disk"].([]interface{})
    dataDisks := v["data_disks"].([]interface{})

    result := compute.ImageStorageProfile{
        DataDisks: expandArmImageImageDataDisk(dataDisks),
        OsDisk: expandArmImageImageOSDisk(osDisk),
    }
    return &result
}

func expandArmImageImageDataDisk(input []interface{}) *[]compute.ImageDataDisk {
    results := make([]compute.ImageDataDisk, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        lun := v["lun"].(int)
        snapshot := v["snapshot"].([]interface{})
        managedDisk := v["managed_disk"].([]interface{})
        blobUri := v["blob_uri"].(string)
        caching := v["caching"].(string)
        diskSizeGb := v["disk_size_gb"].(int)
        storageAccountType := v["storage_account_type"].(string)

        result := compute.ImageDataDisk{
            BlobUri: utils.String(blobUri),
            Caching: compute.CachingTypes(caching),
            DiskSizeGb: utils.Int32(int32(diskSizeGb)),
            Lun: utils.Int32(int32(lun)),
            ManagedDisk: expandArmImageSubResource(managedDisk),
            Snapshot: expandArmImageSubResource(snapshot),
            StorageAccountType: compute.StorageAccountTypes(storageAccountType),
        }

        results = append(results, result)
    }
    return &results
}

func expandArmImageImageOSDisk(input []interface{}) *compute.ImageOSDisk {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    osType := v["os_type"].(string)
    osState := v["os_state"].(string)
    snapshot := v["snapshot"].([]interface{})
    managedDisk := v["managed_disk"].([]interface{})
    blobUri := v["blob_uri"].(string)
    caching := v["caching"].(string)
    diskSizeGb := v["disk_size_gb"].(int)
    storageAccountType := v["storage_account_type"].(string)

    result := compute.ImageOSDisk{
        BlobUri: utils.String(blobUri),
        Caching: compute.CachingTypes(caching),
        DiskSizeGb: utils.Int32(int32(diskSizeGb)),
        ManagedDisk: expandArmImageSubResource(managedDisk),
        OsState: compute.OperatingSystemStateTypes(osState),
        OsType: compute.OperatingSystemTypes(osType),
        Snapshot: expandArmImageSubResource(snapshot),
        StorageAccountType: compute.StorageAccountTypes(storageAccountType),
    }
    return &result
}


func flattenArmImageSubResource(input *compute.SubResource) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if id := input.ID; id != nil {
        result["id"] = *id
    }

    return []interface{}{result}
}

func flattenArmImageImageStorageProfile(input *compute.ImageStorageProfile) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    result["data_disks"] = flattenArmImageImageDataDisk(input.DataDisks)
    result["os_disk"] = flattenArmImageImageOSDisk(input.OsDisk)

    return []interface{}{result}
}

func flattenArmImageImageDataDisk(input *[]compute.ImageDataDisk) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        if blobUri := item.BlobUri; blobUri != nil {
            v["blob_uri"] = *blobUri
        }
        v["caching"] = string(item.Caching)
        if diskSizeGb := item.DiskSizeGb; diskSizeGb != nil {
            v["disk_size_gb"] = int(*diskSizeGb)
        }
        if lun := item.Lun; lun != nil {
            v["lun"] = int(*lun)
        }
        v["managed_disk"] = flattenArmImageSubResource(item.ManagedDisk)
        v["snapshot"] = flattenArmImageSubResource(item.Snapshot)
        v["storage_account_type"] = string(item.StorageAccountType)

        results = append(results, v)
    }

    return results
}

func flattenArmImageImageOSDisk(input *compute.ImageOSDisk) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if blobUri := input.BlobUri; blobUri != nil {
        result["blob_uri"] = *blobUri
    }
    result["caching"] = string(input.Caching)
    if diskSizeGb := input.DiskSizeGb; diskSizeGb != nil {
        result["disk_size_gb"] = int(*diskSizeGb)
    }
    result["managed_disk"] = flattenArmImageSubResource(input.ManagedDisk)
    result["os_state"] = string(input.OsState)
    result["os_type"] = string(input.OsType)
    result["snapshot"] = flattenArmImageSubResource(input.Snapshot)
    result["storage_account_type"] = string(input.StorageAccountType)

    return []interface{}{result}
}
