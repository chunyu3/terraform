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
            "location": azure.SchemaLocation(),

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "server_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "administrator_login_password": {
                Type: schema.TypeString,
                Optional: true,
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
                                string(mysql.SystemAssigned),
                            }, false),
                            Default: string(mysql.SystemAssigned),
                        },
                    },
                },
            },

            "minimal_tls_version": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(mysql.TLS1_0),
                    string(mysql.TLS1_1),
                    string(mysql.TLS1_2),
                    string(mysql.TLSEnforcementDisabled),
                }, false),
                Default: string(mysql.TLS1_0),
            },

            "public_network_access": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(mysql.Enabled),
                    string(mysql.Disabled),
                }, false),
                Default: string(mysql.Enabled),
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

            "tags": tags.Schema(),

            "version": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(mysql.5.6),
                    string(mysql.5.7),
                    string(mysql.8.0),
                }, false),
                Default: string(mysql.5.6),
            },

            "administrator_login": {
                Type: schema.TypeString,
                Computed: true,
            },

            "byok_enforcement": {
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

            "id": {
                Type: schema.TypeString,
                Computed: true,
            },

            "infrastructure_encryption": {
                Type: schema.TypeString,
                Computed: true,
            },

            "master_server_id": {
                Type: schema.TypeString,
                Computed: true,
            },

            "name": {
                Type: schema.TypeString,
                Computed: true,
            },

            "private_endpoint_connections": {
                Type: schema.TypeList,
                Computed: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "id": {
                            Type: schema.TypeString,
                            Computed: true,
                        },
                        "private_endpoint": {
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
                        "private_link_service_connection_state": {
                            Type: schema.TypeList,
                            Computed: true,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "actions_required": {
                                        Type: schema.TypeString,
                                        Computed: true,
                                    },
                                    "description": {
                                        Type: schema.TypeString,
                                        Computed: true,
                                    },
                                    "status": {
                                        Type: schema.TypeString,
                                        Computed: true,
                                    },
                                },
                            },
                        },
                        "provisioning_state": {
                            Type: schema.TypeString,
                            Computed: true,
                        },
                    },
                },
            },

            "replica_capacity": {
                Type: schema.TypeInt,
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
        },
    }
}

func resourceArmServerCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).serversClient
    ctx, cancel := timeouts.ForCreate(meta.(*ArmClient).StopContext, d)
    defer cancel()

    resourceGroupName := d.Get("resource_group").(string)
    name := d.Get("server_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroupName, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Server (Server Name %q / Resource Group %q): %+v", name, resourceGroupName, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_server", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    administratorLoginPassword := d.Get("administrator_login_password").(string)
    identity := d.Get("identity").([]interface{})
    minimalTLSVersion := d.Get("minimal_tls_version").(string)
    publicNetworkAccess := d.Get("public_network_access").(string)
    replicationRole := d.Get("replication_role").(string)
    sku := d.Get("sku").([]interface{})
    sslEnforcement := d.Get("ssl_enforcement").(string)
    storageProfile := d.Get("storage_profile").([]interface{})
    version := d.Get("version").(string)
    tags := d.Get("tags").(map[string]interface{})

    parameters := mysql.ServerUpdateParameters{
        Identity: expandArmServerResourceIdentity(identity),
        Location: utils.String(location),
        ServerUpdateParameters_properties: &mysql.ServerUpdateParameters_properties{
            AdministratorLoginPassword: utils.String(administratorLoginPassword),
            MinimalTLSVersion: mysql.MinimalTlsVersionEnum(minimalTLSVersion),
            PublicNetworkAccess: mysql.PublicNetworkAccessEnum(publicNetworkAccess),
            ReplicationRole: utils.String(replicationRole),
            SslEnforcement: mysql.SslEnforcementEnum(sslEnforcement),
            StorageProfile: expandArmServerStorageProfile(storageProfile),
            Version: mysql.ServerVersion(version),
        },
        Sku: expandArmServerSku(sku),
        Tags: tags.Expand(tags),
    }


    future, err := client.Create(ctx, resourceGroupName, name, parameters)
    if err != nil {
        return fmt.Errorf("Error creating Server (Server Name %q / Resource Group %q): %+v", name, resourceGroupName, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Server (Server Name %q / Resource Group %q): %+v", name, resourceGroupName, err)
    }


    resp, err := client.Get(ctx, resourceGroupName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Server (Server Name %q / Resource Group %q): %+v", name, resourceGroupName, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Server (Server Name %q / Resource Group %q) ID", name, resourceGroupName)
    }
    d.SetId(*resp.ID)

    return resourceArmServerRead(d, meta)
}

func resourceArmServerRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).serversClient
    ctx, cancel := timeouts.ForRead(meta.(*ArmClient).StopContext, d)
    defer cancel()

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroupName := id.ResourceGroup
    name := id.Path["servers"]

    resp, err := client.Get(ctx, resourceGroupName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Server %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Server (Server Name %q / Resource Group %q): %+v", name, resourceGroupName, err)
    }


    d.Set("resource_group", resourceGroupName)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if serverUpdateParametersProperties := resp.ServerUpdateParameters_properties; serverUpdateParametersProperties != nil {
        d.Set("administrator_login", serverUpdateParametersProperties.AdministratorLogin)
        d.Set("byok_enforcement", serverUpdateParametersProperties.ByokEnforcement)
        d.Set("earliest_restore_date", (serverUpdateParametersProperties.EarliestRestoreDate).String())
        d.Set("fully_qualified_domain_name", serverUpdateParametersProperties.FullyQualifiedDomainName)
        d.Set("infrastructure_encryption", string(serverUpdateParametersProperties.InfrastructureEncryption))
        d.Set("master_server_id", serverUpdateParametersProperties.MasterServerID)
        d.Set("minimal_tls_version", string(serverUpdateParametersProperties.MinimalTLSVersion))
        if err := d.Set("private_endpoint_connections", flattenArmServerServerPrivateEndpointConnection(serverUpdateParametersProperties.PrivateEndpointConnections)); err != nil {
            return fmt.Errorf("Error setting `private_endpoint_connections`: %+v", err)
        }
        d.Set("public_network_access", string(serverUpdateParametersProperties.PublicNetworkAccess))
        d.Set("replica_capacity", int(*serverUpdateParametersProperties.ReplicaCapacity))
        d.Set("replication_role", serverUpdateParametersProperties.ReplicationRole)
        d.Set("ssl_enforcement", string(serverUpdateParametersProperties.SslEnforcement))
        if err := d.Set("storage_profile", flattenArmServerStorageProfile(serverUpdateParametersProperties.StorageProfile)); err != nil {
            return fmt.Errorf("Error setting `storage_profile`: %+v", err)
        }
        d.Set("user_visible_state", string(serverUpdateParametersProperties.UserVisibleState))
        d.Set("version", string(serverUpdateParametersProperties.Version))
    }
    d.Set("id", resp.ID)
    if err := d.Set("identity", flattenArmServerResourceIdentity(resp.Identity)); err != nil {
        return fmt.Errorf("Error setting `identity`: %+v", err)
    }
    d.Set("name", resp.Name)
    d.Set("server_name", name)
    if err := d.Set("sku", flattenArmServerSku(resp.Sku)); err != nil {
        return fmt.Errorf("Error setting `sku`: %+v", err)
    }
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}

func resourceArmServerUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).serversClient
    ctx, cancel := timeouts.ForUpdate(meta.(*ArmClient).StopContext, d)
    defer cancel()

      resourceGroupName := d.Get("resource_group").(string)
    location := azure.NormalizeLocation(d.Get("location").(string))
    administratorLoginPassword := d.Get("administrator_login_password").(string)
    identity := d.Get("identity").([]interface{})
    minimalTLSVersion := d.Get("minimal_tls_version").(string)
    publicNetworkAccess := d.Get("public_network_access").(string)
    replicationRole := d.Get("replication_role").(string)
    name := d.Get("server_name").(string)
    sku := d.Get("sku").([]interface{})
    sslEnforcement := d.Get("ssl_enforcement").(string)
    storageProfile := d.Get("storage_profile").([]interface{})
    version := d.Get("version").(string)
    tags := d.Get("tags").(map[string]interface{})

    parameters := mysql.ServerUpdateParameters{
        Identity: expandArmServerResourceIdentity(identity),
        Location: utils.String(location),
        ServerUpdateParameters_properties: &mysql.ServerUpdateParameters_properties{
            AdministratorLoginPassword: utils.String(administratorLoginPassword),
            MinimalTLSVersion: mysql.MinimalTlsVersionEnum(minimalTLSVersion),
            PublicNetworkAccess: mysql.PublicNetworkAccessEnum(publicNetworkAccess),
            ReplicationRole: utils.String(replicationRole),
            SslEnforcement: mysql.SslEnforcementEnum(sslEnforcement),
            StorageProfile: expandArmServerStorageProfile(storageProfile),
            Version: mysql.ServerVersion(version),
        },
        Sku: expandArmServerSku(sku),
        Tags: tags.Expand(tags),
    }


    future, err := client.Update(ctx, resourceGroupName, name, parameters)
    if err != nil {
        return fmt.Errorf("Error updating Server (Server Name %q / Resource Group %q): %+v", name, resourceGroupName, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for update of Server (Server Name %q / Resource Group %q): %+v", name, resourceGroupName, err)
    }

    return resourceArmServerRead(d, meta)
}

func resourceArmServerDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).serversClient
    ctx, cancel := timeouts.ForDelete(meta.(*ArmClient).StopContext, d)
    defer cancel()


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroupName := id.ResourceGroup
    name := id.Path["servers"]

    future, err := client.Delete(ctx, resourceGroupName, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Server (Server Name %q / Resource Group %q): %+v", name, resourceGroupName, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Server (Server Name %q / Resource Group %q): %+v", name, resourceGroupName, err)
        }
    }

    return nil
}

func expandArmServerResourceIdentity(input []interface{}) *mysql.ResourceIdentity {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    type := v["type"].(string)

    result := mysql.ResourceIdentity{
        Type: mysql.IdentityType(type),
    }
    return &result
}

func expandArmServerStorageProfile(input []interface{}) *mysql.StorageProfile {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    backupRetentionDays := v["backup_retention_days"].(int)
    geoRedundantBackup := v["geo_redundant_backup"].(string)
    storageMB := v["storage_mb"].(int)
    storageAutogrow := v["storage_autogrow"].(string)

    result := mysql.StorageProfile{
        BackupRetentionDays: utils.Int(backupRetentionDays),
        GeoRedundantBackup: mysql.GeoRedundantBackup(geoRedundantBackup),
        StorageAutogrow: mysql.StorageAutogrow(storageAutogrow),
        StorageMB: utils.Int32(int32(storageMB)),
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


func flattenArmServerServerPrivateEndpointConnection(input *[]mysql.ServerPrivateEndpointConnection) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        if id := item.ID; id != nil {
            v["id"] = *id
        }
        if serverPrivateEndpointConnectionProperties := item.ServerPrivateEndpointConnectionProperties; serverPrivateEndpointConnectionProperties != nil {
            v["private_endpoint"] = flattenArmServerPrivateEndpointProperty(serverPrivateEndpointConnectionProperties.PrivateEndpoint)
            v["private_link_service_connection_state"] = flattenArmServerServerPrivateLinkServiceConnectionStateProperty(serverPrivateEndpointConnectionProperties.PrivateLinkServiceConnectionState)
            v["provisioning_state"] = string(serverPrivateEndpointConnectionProperties.ProvisioningState)
        }

        results = append(results, v)
    }

    return results
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
    if storageMb := input.StorageMB; storageMb != nil {
        result["storage_mb"] = int(*storageMb)
    }

    return []interface{}{result}
}

func flattenArmServerResourceIdentity(input *mysql.ResourceIdentity) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    result["type"] = string(input.Type)

    return []interface{}{result}
}

func flattenArmServerSku(input *mysql.Sku) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if capacity := input.Capacity; capacity != nil {
        result["capacity"] = int(*capacity)
    }
    if family := input.Family; family != nil {
        result["family"] = *family
    }
    if name := input.Name; name != nil {
        result["name"] = *name
    }
    if size := input.Size; size != nil {
        result["size"] = *size
    }
    result["tier"] = string(input.Tier)

    return []interface{}{result}
}

func flattenArmServerPrivateEndpointProperty(input *mysql.PrivateEndpointProperty) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if id := input.ID; id != nil {
        result["id"] = *id
    }

    return []interface{}{result}
}

func flattenArmServerServerPrivateLinkServiceConnectionStateProperty(input *mysql.ServerPrivateLinkServiceConnectionStateProperty) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    result["actions_required"] = string(input.ActionsRequired)
    if description := input.Description; description != nil {
        result["description"] = *description
    }
    result["status"] = string(input.Status)

    return []interface{}{result}
}
