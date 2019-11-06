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



func resourceArmServer() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmServerCreate,
        Read: resourceArmServerRead,
        Update: resourceArmServerUpdate,
        Delete: resourceArmServerDelete,

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

            "sku": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "capacity": {
                            Type: schema.TypeInt,
                            Optional: true,
                        },
                        "family": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "name": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "size": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "tier": {
                            Type: schema.TypeString,
                            Optional: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(mysql.Basic),
                                string(mysql.GeneralPurpose),
                                string(mysql.MemoryOptimized),
                            }, false),
                            Default: string(mysql.Basic),
                        },
                    },
                },
            },

            "ssl_enforcement": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(mysql.Enabled),
                    string(mysql.Disabled),
                }, false),
                Default: string(mysql.Enabled),
            },

            "storage_profile": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "backup_retention_days": {
                            Type: schema.TypeInt,
                            Optional: true,
                        },
                        "geo_redundant_backup": {
                            Type: schema.TypeString,
                            Optional: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(mysql.Enabled),
                                string(mysql.Disabled),
                            }, false),
                            Default: string(mysql.Enabled),
                        },
                        "storage_autogrow": {
                            Type: schema.TypeString,
                            Optional: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(mysql.Enabled),
                                string(mysql.Disabled),
                            }, false),
                            Default: string(mysql.Enabled),
                        },
                        "storage_mb": {
                            Type: schema.TypeInt,
                            Optional: true,
                        },
                    },
                },
            },

            "version": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(mysql.5.6),
                    string(mysql.5.7),
                }, false),
                Default: string(mysql.5.6),
            },

            "administrator_login": {
                Type: schema.TypeString,
                Computed: true,
            },

            "earliest_restore_date": {
                Type: schema.TypeString,
                Computed: true,
            },

            "fully_qualified_domain_name": {
                Type: schema.TypeString,
                Computed: true,
            },

            "master_server_id": {
                Type: schema.TypeString,
                Computed: true,
            },

            "replica_capacity": {
                Type: schema.TypeInt,
                Computed: true,
            },

            "replication_role": {
                Type: schema.TypeString,
                Computed: true,
            },

            "type": {
                Type: schema.TypeString,
                Computed: true,
            },

            "user_visible_state": {
                Type: schema.TypeString,
                Computed: true,
            },

            "tags": tags.Schema(),
        },
    }
}

func resourceArmServerCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).serversClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Server %q (Resource Group %q): %+v", name, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_server", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    sku := d.Get("sku").([]interface{})
    sslEnforcement := d.Get("ssl_enforcement").(string)
    storageProfile := d.Get("storage_profile").([]interface{})
    version := d.Get("version").(string)
    t := d.Get("tags").(map[string]interface{})

    parameters := mysql.ServerForCreate{
        Location: utils.String(location),
        ServerPropertiesForCreate: &mysql.ServerPropertiesForCreate{
            SslEnforcement: mysql.SslEnforcementEnum(sslEnforcement),
            StorageProfile: expandArmServerStorageProfile(storageProfile),
            Version: mysql.ServerVersion(version),
        },
        Sku: expandArmServerSku(sku),
        Tags: tags.Expand(t),
    }


    future, err := client.Create(ctx, resourceGroup, name, parameters)
    if err != nil {
        return fmt.Errorf("Error creating Server %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Server %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Server %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Server %q (Resource Group %q) ID", name, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmServerRead(d, meta)
}

func resourceArmServerRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).serversClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["servers"]

    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Server %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Server %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if serverPropertiesForCreate := resp.ServerPropertiesForCreate; serverPropertiesForCreate != nil {
        d.Set("administrator_login", serverPropertiesForCreate.AdministratorLogin)
        d.Set("earliest_restore_date", (serverPropertiesForCreate.EarliestRestoreDate).String())
        d.Set("fully_qualified_domain_name", serverPropertiesForCreate.FullyQualifiedDomainName)
        d.Set("master_server_id", serverPropertiesForCreate.MasterServerID)
        d.Set("replica_capacity", int(*serverPropertiesForCreate.ReplicaCapacity))
        d.Set("replication_role", serverPropertiesForCreate.ReplicationRole)
        d.Set("ssl_enforcement", string(serverPropertiesForCreate.SslEnforcement))
        if err := d.Set("storage_profile", flattenArmServerStorageProfile(serverPropertiesForCreate.StorageProfile)); err != nil {
            return fmt.Errorf("Error setting `storage_profile`: %+v", err)
        }
        d.Set("user_visible_state", string(serverPropertiesForCreate.UserVisibleState))
        d.Set("version", string(serverPropertiesForCreate.Version))
    }
    if err := d.Set("sku", flattenArmServerSku(resp.Sku)); err != nil {
        return fmt.Errorf("Error setting `sku`: %+v", err)
    }
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}

func resourceArmServerUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).serversClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    sku := d.Get("sku").([]interface{})
    sslEnforcement := d.Get("ssl_enforcement").(string)
    storageProfile := d.Get("storage_profile").([]interface{})
    version := d.Get("version").(string)
    t := d.Get("tags").(map[string]interface{})

    parameters := mysql.ServerForCreate{
        Location: utils.String(location),
        ServerPropertiesForCreate: &mysql.ServerPropertiesForCreate{
            SslEnforcement: mysql.SslEnforcementEnum(sslEnforcement),
            StorageProfile: expandArmServerStorageProfile(storageProfile),
            Version: mysql.ServerVersion(version),
        },
        Sku: expandArmServerSku(sku),
        Tags: tags.Expand(t),
    }


    future, err := client.Update(ctx, resourceGroup, name, parameters)
    if err != nil {
        return fmt.Errorf("Error updating Server %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for update of Server %q (Resource Group %q): %+v", name, resourceGroup, err)
    }

    return resourceArmServerRead(d, meta)
}

func resourceArmServerDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).serversClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["servers"]

    future, err := client.Delete(ctx, resourceGroup, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Server %q (Resource Group %q): %+v", name, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Server %q (Resource Group %q): %+v", name, resourceGroup, err)
        }
    }

    return nil
}

func expandArmServerStorageProfile(input []interface{}) *mysql.StorageProfile {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    backupRetentionDays := v["backup_retention_days"].(int)
    geoRedundantBackup := v["geo_redundant_backup"].(string)
    storageMb := v["storage_mb"].(int)
    storageAutogrow := v["storage_autogrow"].(string)

    result := mysql.StorageProfile{
        BackupRetentionDays: utils.Int(backupRetentionDays),
        GeoRedundantBackup: mysql.GeoRedundantBackup(geoRedundantBackup),
        StorageAutogrow: mysql.StorageAutogrow(storageAutogrow),
        StorageMb: utils.Int32(int32(storageMb)),
    }
    return &result
}

func expandArmServerSku(input []interface{}) *mysql.Sku {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    name := v["name"].(string)
    tier := v["tier"].(string)
    capacity := v["capacity"].(int)
    size := v["size"].(string)
    family := v["family"].(string)

    result := mysql.Sku{
        Capacity: utils.Int32(int32(capacity)),
        Family: utils.String(family),
        Name: utils.String(name),
        Size: utils.String(size),
        Tier: mysql.SkuTier(tier),
    }
    return &result
}


func flattenArmServerStorageProfile(input *mysql.StorageProfile) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if backupRetentionDays := input.BackupRetentionDays; backupRetentionDays != nil {
        result["backup_retention_days"] = *backupRetentionDays
    }
    result["geo_redundant_backup"] = string(input.GeoRedundantBackup)
    result["storage_autogrow"] = string(input.StorageAutogrow)
    if storageMb := input.StorageMb; storageMb != nil {
        result["storage_mb"] = int(*storageMb)
    }

    return []interface{}{result}
}

func flattenArmServerSku(input *mysql.Sku) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if name := input.Name; name != nil {
        result["name"] = *name
    }
    if capacity := input.Capacity; capacity != nil {
        result["capacity"] = int(*capacity)
    }
    if family := input.Family; family != nil {
        result["family"] = *family
    }
    if size := input.Size; size != nil {
        result["size"] = *size
    }
    result["tier"] = string(input.Tier)

    return []interface{}{result}
}
