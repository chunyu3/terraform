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



func resourceArmDiskEncryptionSet() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmDiskEncryptionSetCreate,
        Read: resourceArmDiskEncryptionSetRead,
        Update: resourceArmDiskEncryptionSetUpdate,
        Delete: resourceArmDiskEncryptionSetDelete,

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

            "active_key": {
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
                                string(compute.SystemAssigned),
                            }, false),
                            Default: string(compute.SystemAssigned),
                        },
                    },
                },
            },

            "previous_keys": {
                Type: schema.TypeList,
                Computed: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "key_url": {
                            Type: schema.TypeString,
                            Computed: true,
                        },
                        "source_vault": {
                            Type: schema.TypeList,
                            Computed: true,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "id": {
                                        Type: schema.TypeString,
                                        Computed: true,
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

func resourceArmDiskEncryptionSetCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).diskEncryptionSetsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Disk Encryption Set %q (Resource Group %q): %+v", name, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_disk_encryption_set", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    activeKey := d.Get("active_key").([]interface{})
    identity := d.Get("identity").([]interface{})
    t := d.Get("tags").(map[string]interface{})

    diskEncryptionSet := compute.DiskEncryptionSetUpdate{
        Identity: expandArmDiskEncryptionSetEncryptionSetIdentity(identity),
        Location: utils.String(location),
        DiskEncryptionSetUpdateProperties: &compute.DiskEncryptionSetUpdateProperties{
            ActiveKey: expandArmDiskEncryptionSetKeyVaultAndKeyReference(activeKey),
        },
        Tags: tags.Expand(t),
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, name, diskEncryptionSet)
    if err != nil {
        return fmt.Errorf("Error creating Disk Encryption Set %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Disk Encryption Set %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Disk Encryption Set %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Disk Encryption Set %q (Resource Group %q) ID", name, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmDiskEncryptionSetRead(d, meta)
}

func resourceArmDiskEncryptionSetRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).diskEncryptionSetsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["diskEncryptionSets"]

    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Disk Encryption Set %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Disk Encryption Set %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if diskEncryptionSetUpdateProperties := resp.DiskEncryptionSetUpdateProperties; diskEncryptionSetUpdateProperties != nil {
        if err := d.Set("active_key", flattenArmDiskEncryptionSetKeyVaultAndKeyReference(diskEncryptionSetUpdateProperties.ActiveKey)); err != nil {
            return fmt.Errorf("Error setting `active_key`: %+v", err)
        }
        if err := d.Set("previous_keys", flattenArmDiskEncryptionSetKeyVaultAndKeyReference(diskEncryptionSetUpdateProperties.PreviousKeys)); err != nil {
            return fmt.Errorf("Error setting `previous_keys`: %+v", err)
        }
        d.Set("provisioning_state", diskEncryptionSetUpdateProperties.ProvisioningState)
    }
    if err := d.Set("identity", flattenArmDiskEncryptionSetEncryptionSetIdentity(resp.Identity)); err != nil {
        return fmt.Errorf("Error setting `identity`: %+v", err)
    }
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}

func resourceArmDiskEncryptionSetUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).diskEncryptionSetsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    activeKey := d.Get("active_key").([]interface{})
    identity := d.Get("identity").([]interface{})
    t := d.Get("tags").(map[string]interface{})

    diskEncryptionSet := compute.DiskEncryptionSetUpdate{
        Identity: expandArmDiskEncryptionSetEncryptionSetIdentity(identity),
        DiskEncryptionSetUpdateProperties: &compute.DiskEncryptionSetUpdateProperties{
            ActiveKey: expandArmDiskEncryptionSetKeyVaultAndKeyReference(activeKey),
        },
        Tags: tags.Expand(t),
    }


    future, err := client.Update(ctx, resourceGroup, name, diskEncryptionSet)
    if err != nil {
        return fmt.Errorf("Error updating Disk Encryption Set %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for update of Disk Encryption Set %q (Resource Group %q): %+v", name, resourceGroup, err)
    }

    return resourceArmDiskEncryptionSetRead(d, meta)
}

func resourceArmDiskEncryptionSetDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).diskEncryptionSetsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["diskEncryptionSets"]

    future, err := client.Delete(ctx, resourceGroup, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Disk Encryption Set %q (Resource Group %q): %+v", name, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Disk Encryption Set %q (Resource Group %q): %+v", name, resourceGroup, err)
        }
    }

    return nil
}

func expandArmDiskEncryptionSetEncryptionSetIdentity(input []interface{}) *compute.EncryptionSetIdentity {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    type := v["type"].(string)

    result := compute.EncryptionSetIdentity{
        Type: compute.DiskEncryptionSetIdentityType(type),
    }
    return &result
}

func expandArmDiskEncryptionSetKeyVaultAndKeyReference(input []interface{}) *compute.KeyVaultAndKeyReference {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    sourceVault := v["source_vault"].([]interface{})
    keyUrl := v["key_url"].(string)

    result := compute.KeyVaultAndKeyReference{
        KeyURL: utils.String(keyUrl),
        SourceVault: expandArmDiskEncryptionSetSourceVault(sourceVault),
    }
    return &result
}

func expandArmDiskEncryptionSetSourceVault(input []interface{}) *compute.SourceVault {
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


func flattenArmDiskEncryptionSetKeyVaultAndKeyReference(input *compute.KeyVaultAndKeyReference) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if keyUrl := input.KeyURL; keyUrl != nil {
        result["key_url"] = *keyUrl
    }
    result["source_vault"] = flattenArmDiskEncryptionSetSourceVault(input.SourceVault)

    return []interface{}{result}
}

func flattenArmDiskEncryptionSetKeyVaultAndKeyReference(input *[]compute.KeyVaultAndKeyReference) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        if keyUrl := item.KeyURL; keyUrl != nil {
            v["key_url"] = *keyUrl
        }
        v["source_vault"] = flattenArmDiskEncryptionSetSourceVault(item.SourceVault)

        results = append(results, v)
    }

    return results
}

func flattenArmDiskEncryptionSetEncryptionSetIdentity(input *compute.EncryptionSetIdentity) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    result["type"] = string(input.Type)

    return []interface{}{result}
}

func flattenArmDiskEncryptionSetSourceVault(input *compute.SourceVault) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if id := input.ID; id != nil {
        result["id"] = *id
    }

    return []interface{}{result}
}
