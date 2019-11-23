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

            "administrator_login_password": {
                Type: schema.TypeString,
                Optional: true,
            },

            "replication_role": {
                Type: schema.TypeString,
                Optional: true,
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
                                string(postgresql.Basic),
                                string(postgresql.GeneralPurpose),
                                string(postgresql.MemoryOptimized),
                            }, false),
                            Default: string(postgresql.Basic),
                        },
                    },
                },
            },

            "ssl_enforcement": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(postgresql.Enabled),
                    string(postgresql.Disabled),
                }, false),
                Default: string(postgresql.Enabled),
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
                                string(postgresql.Enabled),
                                string(postgresql.Disabled),
                            }, false),
                            Default: string(postgresql.Enabled),
                        },
                        "storage_autogrow": {
                            Type: schema.TypeString,
                            Optional: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(postgresql.Enabled),
                                string(postgresql.Disabled),
                            }, false),
                            Default: string(postgresql.Enabled),
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
                    string(postgresql.9.5),
                    string(postgresql.9.6),
                    string(postgresql.10),
                    string(postgresql.10.0),
                    string(postgresql.10.2),
                    string(postgresql.11),
                }, false),
                Default: string(postgresql.9.5),
            },

            "type": {
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
    administratorLoginPassword := d.Get("administrator_login_password").(string)
    replicationRole := d.Get("replication_role").(string)
    sku := d.Get("sku").([]interface{})
    sslEnforcement := d.Get("ssl_enforcement").(string)
    storageProfile := d.Get("storage_profile").([]interface{})
    version := d.Get("version").(string)
    t := d.Get("tags").(map[string]interface{})

    parameters := postgresql.ServerUpdateParameters{
        Location: utils.String(location),
        ServerUpdateParameters_properties: &postgresql.ServerUpdateParameters_properties{
            AdministratorLoginPassword: utils.String(administratorLoginPassword),
            ReplicationRole: utils.String(replicationRole),
            SslEnforcement: postgresql.SslEnforcementEnum(sslEnforcement),
            StorageProfile: expandArmServerStorageProfile(storageProfile),
            Version: postgresql.ServerVersion(version),
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
    d.Set("type", resp.Type)

    return nil
}

func resourceArmServerUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).serversClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    administratorLoginPassword := d.Get("administrator_login_password").(string)
    replicationRole := d.Get("replication_role").(string)
    sku := d.Get("sku").([]interface{})
    sslEnforcement := d.Get("ssl_enforcement").(string)
    storageProfile := d.Get("storage_profile").([]interface{})
    version := d.Get("version").(string)
    t := d.Get("tags").(map[string]interface{})

    parameters := postgresql.ServerUpdateParameters{
        ServerUpdateParameters_properties: &postgresql.ServerUpdateParameters_properties{
            AdministratorLoginPassword: utils.String(administratorLoginPassword),
            ReplicationRole: utils.String(replicationRole),
            SslEnforcement: postgresql.SslEnforcementEnum(sslEnforcement),
            StorageProfile: expandArmServerStorageProfile(storageProfile),
            Version: postgresql.ServerVersion(version),
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

func expandArmServerStorageProfile(input []interface{}) *postgresql.StorageProfile {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    backupRetentionDays := v["backup_retention_days"].(int)
    geoRedundantBackup := v["geo_redundant_backup"].(string)
    storageMb := v["storage_mb"].(int)
    storageAutogrow := v["storage_autogrow"].(string)

    result := postgresql.StorageProfile{
        BackupRetentionDays: utils.Int(backupRetentionDays),
        GeoRedundantBackup: postgresql.GeoRedundantBackup(geoRedundantBackup),
        StorageAutogrow: postgresql.StorageAutogrow(storageAutogrow),
        StorageMB: utils.Int32(int32(storageMb)),
    }
    return &result
}

func expandArmServerSku(input []interface{}) *postgresql.Sku {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    name := v["name"].(string)
    tier := v["tier"].(string)
    capacity := v["capacity"].(int)
    size := v["size"].(string)
    family := v["family"].(string)

    result := postgresql.Sku{
        Capacity: utils.Int32(int32(capacity)),
        Family: utils.String(family),
        Name: utils.String(name),
        Size: utils.String(size),
        Tier: postgresql.SkuTier(tier),
    }
    return &result
}