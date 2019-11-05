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
                Computed: true,
            },

            "location": azure.SchemaLocation(),

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "server_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

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
                                string(mariadb.Basic),
                                string(mariadb.GeneralPurpose),
                                string(mariadb.MemoryOptimized),
                            }, false),
                            Default: string(mariadb.Basic),
                        },
                    },
                },
            },

            "ssl_enforcement": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(mariadb.Enabled),
                    string(mariadb.Disabled),
                }, false),
                Default: string(mariadb.Enabled),
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
                                string(mariadb.Enabled),
                                string(mariadb.Disabled),
                            }, false),
                            Default: string(mariadb.Enabled),
                        },
                        "storage_autogrow": {
                            Type: schema.TypeString,
                            Optional: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(mariadb.Enabled),
                                string(mariadb.Disabled),
                            }, false),
                            Default: string(mariadb.Enabled),
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
                    string(mariadb.5.6),
                    string(mariadb.5.7),
                }, false),
                Default: string(mariadb.5.6),
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

    resourceGroup := d.Get("resource_group").(string)
    serverName := d.Get("server_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, serverName)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Server (Server Name %q / Resource Group %q): %+v", serverName, resourceGroup, err)
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

    parameters := mariadb.ServerForCreate{
        Location: utils.String(location),
        ServerPropertiesForCreate: &mariadb.ServerPropertiesForCreate{
            SslEnforcement: mariadb.SslEnforcementEnum(sslEnforcement),
            StorageProfile: expandArmServerStorageProfile(storageProfile),
            Version: mariadb.ServerVersion(version),
        },
        Sku: expandArmServerSku(sku),
        Tags: tags.Expand(t),
    }


    future, err := client.Create(ctx, resourceGroup, serverName, parameters)
    if err != nil {
        return fmt.Errorf("Error creating Server (Server Name %q / Resource Group %q): %+v", serverName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Server (Server Name %q / Resource Group %q): %+v", serverName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, serverName)
    if err != nil {
        return fmt.Errorf("Error retrieving Server (Server Name %q / Resource Group %q): %+v", serverName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Server (Server Name %q / Resource Group %q) ID", serverName, resourceGroup)
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
    serverName := id.Path["servers"]

    resp, err := client.Get(ctx, resourceGroup, serverName)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Server %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Server (Server Name %q / Resource Group %q): %+v", serverName, resourceGroup, err)
    }


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
    d.Set("server_name", serverName)
    if err := d.Set("sku", flattenArmServerSku(resp.Sku)); err != nil {
        return fmt.Errorf("Error setting `sku`: %+v", err)
    }
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}

func resourceArmServerUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).serversClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    serverName := d.Get("server_name").(string)
    sku := d.Get("sku").([]interface{})
    sslEnforcement := d.Get("ssl_enforcement").(string)
    storageProfile := d.Get("storage_profile").([]interface{})
    version := d.Get("version").(string)
    t := d.Get("tags").(map[string]interface{})

    parameters := mariadb.ServerForCreate{
        Location: utils.String(location),
        ServerPropertiesForCreate: &mariadb.ServerPropertiesForCreate{
            SslEnforcement: mariadb.SslEnforcementEnum(sslEnforcement),
            StorageProfile: expandArmServerStorageProfile(storageProfile),
            Version: mariadb.ServerVersion(version),
        },
        Sku: expandArmServerSku(sku),
        Tags: tags.Expand(t),
    }


    future, err := client.Update(ctx, resourceGroup, serverName, parameters)
    if err != nil {
        return fmt.Errorf("Error updating Server (Server Name %q / Resource Group %q): %+v", serverName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for update of Server (Server Name %q / Resource Group %q): %+v", serverName, resourceGroup, err)
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
    serverName := id.Path["servers"]

    future, err := client.Delete(ctx, resourceGroup, serverName)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Server (Server Name %q / Resource Group %q): %+v", serverName, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Server (Server Name %q / Resource Group %q): %+v", serverName, resourceGroup, err)
        }
    }

    return nil
}

func expandArmServerStorageProfile(input []interface{}) *mariadb.StorageProfile {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    backupRetentionDays := v["backup_retention_days"].(int)
    geoRedundantBackup := v["geo_redundant_backup"].(string)
    storageMb := v["storage_mb"].(int)
    storageAutogrow := v["storage_autogrow"].(string)

    result := mariadb.StorageProfile{
        BackupRetentionDays: utils.Int(backupRetentionDays),
        GeoRedundantBackup: mariadb.GeoRedundantBackup(geoRedundantBackup),
        StorageAutogrow: mariadb.StorageAutogrow(storageAutogrow),
        StorageMb: utils.Int32(int32(storageMb)),
    }
    return &result
}

func expandArmServerSku(input []interface{}) *mariadb.Sku {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    name := v["name"].(string)
    tier := v["tier"].(string)
    capacity := v["capacity"].(int)
    size := v["size"].(string)
    family := v["family"].(string)

    result := mariadb.Sku{
        Capacity: utils.Int32(int32(capacity)),
        Family: utils.String(family),
        Name: utils.String(name),
        Size: utils.String(size),
        Tier: mariadb.SkuTier(tier),
    }
    return &result
}


func flattenArmServerStorageProfile(input *mariadb.StorageProfile) []interface{} {
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

func flattenArmServerSku(input *mariadb.Sku) []interface{} {
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