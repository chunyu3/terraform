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



func resourceArmDatabase() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmDatabaseCreate,
        Read: resourceArmDatabaseRead,
        Update: resourceArmDatabaseUpdate,
        Delete: resourceArmDatabaseDelete,

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

            "server_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "auto_pause_delay": {
                Type: schema.TypeInt,
                Optional: true,
            },

            "catalog_collation": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(sql.DATABASE_DEFAULT),
                    string(sql.SQL_Latin1_General_CP1_CI_AS),
                }, false),
                Default: string(sql.DATABASE_DEFAULT),
            },

            "collation": {
                Type: schema.TypeString,
                Optional: true,
            },

            "create_mode": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(sql.Default),
                    string(sql.Copy),
                    string(sql.Secondary),
                    string(sql.PointInTimeRestore),
                    string(sql.Restore),
                    string(sql.Recovery),
                    string(sql.RestoreExternalBackup),
                    string(sql.RestoreExternalBackupSecondary),
                    string(sql.RestoreLongTermRetentionBackup),
                    string(sql.OnlineSecondary),
                }, false),
                Default: string(sql.Default),
            },

            "elastic_pool_id": {
                Type: schema.TypeString,
                Optional: true,
            },

            "license_type": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(sql.LicenseIncluded),
                    string(sql.BasePrice),
                }, false),
                Default: string(sql.LicenseIncluded),
            },

            "long_term_retention_backup_resource_id": {
                Type: schema.TypeString,
                Optional: true,
            },

            "max_size_bytes": {
                Type: schema.TypeInt,
                Optional: true,
            },

            "min_capacity": {
                Type: schema.TypeFloat,
                Optional: true,
            },

            "read_replica_count": {
                Type: schema.TypeInt,
                Optional: true,
            },

            "read_scale": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(sql.Enabled),
                    string(sql.Disabled),
                }, false),
                Default: string(sql.Enabled),
            },

            "recoverable_database_id": {
                Type: schema.TypeString,
                Optional: true,
            },

            "recovery_services_recovery_point_id": {
                Type: schema.TypeString,
                Optional: true,
            },

            "restorable_dropped_database_id": {
                Type: schema.TypeString,
                Optional: true,
            },

            "restore_point_in_time": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validateRFC3339Date,
            },

            "sample_name": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(sql.AdventureWorksLT),
                    string(sql.WideWorldImportersStd),
                    string(sql.WideWorldImportersFull),
                }, false),
                Default: string(sql.AdventureWorksLT),
            },

            "sku": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "name": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "capacity": {
                            Type: schema.TypeInt,
                            Optional: true,
                        },
                        "family": {
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
                        },
                    },
                },
            },

            "source_database_deletion_date": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validateRFC3339Date,
            },

            "source_database_id": {
                Type: schema.TypeString,
                Optional: true,
            },

            "zone_redundant": {
                Type: schema.TypeBool,
                Optional: true,
            },

            "creation_date": {
                Type: schema.TypeString,
                Computed: true,
            },

            "current_service_objective_name": {
                Type: schema.TypeString,
                Computed: true,
            },

            "current_sku": {
                Type: schema.TypeList,
                Computed: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "name": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "capacity": {
                            Type: schema.TypeInt,
                            Optional: true,
                        },
                        "family": {
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
                        },
                    },
                },
            },

            "database_id": {
                Type: schema.TypeString,
                Computed: true,
            },

            "default_secondary_location": {
                Type: schema.TypeString,
                Computed: true,
            },

            "earliest_restore_date": {
                Type: schema.TypeString,
                Computed: true,
            },

            "failover_group_id": {
                Type: schema.TypeString,
                Computed: true,
            },

            "kind": {
                Type: schema.TypeString,
                Computed: true,
            },

            "managed_by": {
                Type: schema.TypeString,
                Computed: true,
            },

            "max_log_size_bytes": {
                Type: schema.TypeInt,
                Computed: true,
            },

            "paused_date": {
                Type: schema.TypeString,
                Computed: true,
            },

            "requested_service_objective_name": {
                Type: schema.TypeString,
                Computed: true,
            },

            "resumed_date": {
                Type: schema.TypeString,
                Computed: true,
            },

            "status": {
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

func resourceArmDatabaseCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).databasesClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    serverName := d.Get("server_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, serverName, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Database %q (Server Name %q / Resource Group %q): %+v", name, serverName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_database", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    autoPauseDelay := d.Get("auto_pause_delay").(int)
    catalogCollation := d.Get("catalog_collation").(string)
    collation := d.Get("collation").(string)
    createMode := d.Get("create_mode").(string)
    elasticPoolId := d.Get("elastic_pool_id").(string)
    licenseType := d.Get("license_type").(string)
    longTermRetentionBackupResourceId := d.Get("long_term_retention_backup_resource_id").(string)
    maxSizeBytes := d.Get("max_size_bytes").(int)
    minCapacity := d.Get("min_capacity").(float64)
    readReplicaCount := d.Get("read_replica_count").(int)
    readScale := d.Get("read_scale").(string)
    recoverableDatabaseId := d.Get("recoverable_database_id").(string)
    recoveryServicesRecoveryPointId := d.Get("recovery_services_recovery_point_id").(string)
    restorableDroppedDatabaseId := d.Get("restorable_dropped_database_id").(string)
    restorePointInTime := d.Get("restore_point_in_time").(string)
    sampleName := d.Get("sample_name").(string)
    sku := d.Get("sku").([]interface{})
    sourceDatabaseDeletionDate := d.Get("source_database_deletion_date").(string)
    sourceDatabaseId := d.Get("source_database_id").(string)
    zoneRedundant := d.Get("zone_redundant").(bool)
    t := d.Get("tags").(map[string]interface{})

    parameters := sql.Database{
        Location: utils.String(location),
        DatabaseProperties: &sql.DatabaseProperties{
            AutoPauseDelay: utils.Int32(int32(autoPauseDelay)),
            CatalogCollation: sql.CatalogCollationType(catalogCollation),
            Collation: utils.String(collation),
            CreateMode: sql.CreateMode(createMode),
            ElasticPoolID: utils.String(elasticPoolId),
            LicenseType: sql.DatabaseLicenseType(licenseType),
            LongTermRetentionBackupResourceID: utils.String(longTermRetentionBackupResourceId),
            MaxSizeBytes: utils.Int64(int64(maxSizeBytes)),
            MinCapacity: utils.Float(minCapacity),
            ReadReplicaCount: utils.Int32(int32(readReplicaCount)),
            ReadScale: sql.DatabaseReadScale(readScale),
            RecoverableDatabaseID: utils.String(recoverableDatabaseId),
            RecoveryServicesRecoveryPointID: utils.String(recoveryServicesRecoveryPointId),
            RestorableDroppedDatabaseID: utils.String(restorableDroppedDatabaseId),
            RestorePointInTime: convertStringToDate(restorePointInTime),
            SampleName: sql.SampleName(sampleName),
            SourceDatabaseDeletionDate: convertStringToDate(sourceDatabaseDeletionDate),
            SourceDatabaseID: utils.String(sourceDatabaseId),
            ZoneRedundant: utils.Bool(zoneRedundant),
        },
        Sku: expandArmDatabaseSku(sku),
        Tags: tags.Expand(t),
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, serverName, name, parameters)
    if err != nil {
        return fmt.Errorf("Error creating Database %q (Server Name %q / Resource Group %q): %+v", name, serverName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Database %q (Server Name %q / Resource Group %q): %+v", name, serverName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, serverName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Database %q (Server Name %q / Resource Group %q): %+v", name, serverName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Database %q (Server Name %q / Resource Group %q) ID", name, serverName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmDatabaseRead(d, meta)
}

func resourceArmDatabaseRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).databasesClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    serverName := id.Path["servers"]
    name := id.Path["databases"]

    resp, err := client.Get(ctx, resourceGroup, serverName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Database %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Database %q (Server Name %q / Resource Group %q): %+v", name, serverName, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if databaseProperties := resp.DatabaseProperties; databaseProperties != nil {
        d.Set("auto_pause_delay", int(*databaseProperties.AutoPauseDelay))
        d.Set("catalog_collation", string(databaseProperties.CatalogCollation))
        d.Set("collation", databaseProperties.Collation)
        d.Set("create_mode", string(databaseProperties.CreateMode))
        d.Set("creation_date", (databaseProperties.CreationDate).String())
        d.Set("current_service_objective_name", databaseProperties.CurrentServiceObjectiveName)
        if err := d.Set("current_sku", flattenArmDatabaseSku(databaseProperties.CurrentSku)); err != nil {
            return fmt.Errorf("Error setting `current_sku`: %+v", err)
        }
        d.Set("database_id", databaseProperties.DatabaseID)
        d.Set("default_secondary_location", databaseProperties.DefaultSecondaryLocation)
        d.Set("earliest_restore_date", (databaseProperties.EarliestRestoreDate).String())
        d.Set("elastic_pool_id", databaseProperties.ElasticPoolID)
        d.Set("failover_group_id", databaseProperties.FailoverGroupID)
        d.Set("license_type", string(databaseProperties.LicenseType))
        d.Set("long_term_retention_backup_resource_id", databaseProperties.LongTermRetentionBackupResourceID)
        d.Set("max_log_size_bytes", int(*databaseProperties.MaxLogSizeBytes))
        d.Set("max_size_bytes", int(*databaseProperties.MaxSizeBytes))
        d.Set("min_capacity", databaseProperties.MinCapacity)
        d.Set("paused_date", (databaseProperties.PausedDate).String())
        d.Set("read_replica_count", int(*databaseProperties.ReadReplicaCount))
        d.Set("read_scale", string(databaseProperties.ReadScale))
        d.Set("recoverable_database_id", databaseProperties.RecoverableDatabaseID)
        d.Set("recovery_services_recovery_point_id", databaseProperties.RecoveryServicesRecoveryPointID)
        d.Set("requested_service_objective_name", databaseProperties.RequestedServiceObjectiveName)
        d.Set("restorable_dropped_database_id", databaseProperties.RestorableDroppedDatabaseID)
        d.Set("restore_point_in_time", (databaseProperties.RestorePointInTime).String())
        d.Set("resumed_date", (databaseProperties.ResumedDate).String())
        d.Set("sample_name", string(databaseProperties.SampleName))
        d.Set("source_database_deletion_date", (databaseProperties.SourceDatabaseDeletionDate).String())
        d.Set("source_database_id", databaseProperties.SourceDatabaseID)
        d.Set("status", string(databaseProperties.Status))
        d.Set("zone_redundant", databaseProperties.ZoneRedundant)
    }
    d.Set("kind", resp.Kind)
    d.Set("managed_by", resp.ManagedBy)
    d.Set("server_name", serverName)
    if err := d.Set("sku", flattenArmDatabaseSku(resp.Sku)); err != nil {
        return fmt.Errorf("Error setting `sku`: %+v", err)
    }
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}

func resourceArmDatabaseUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).databasesClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    autoPauseDelay := d.Get("auto_pause_delay").(int)
    catalogCollation := d.Get("catalog_collation").(string)
    collation := d.Get("collation").(string)
    createMode := d.Get("create_mode").(string)
    elasticPoolId := d.Get("elastic_pool_id").(string)
    licenseType := d.Get("license_type").(string)
    longTermRetentionBackupResourceId := d.Get("long_term_retention_backup_resource_id").(string)
    maxSizeBytes := d.Get("max_size_bytes").(int)
    minCapacity := d.Get("min_capacity").(float64)
    readReplicaCount := d.Get("read_replica_count").(int)
    readScale := d.Get("read_scale").(string)
    recoverableDatabaseId := d.Get("recoverable_database_id").(string)
    recoveryServicesRecoveryPointId := d.Get("recovery_services_recovery_point_id").(string)
    restorableDroppedDatabaseId := d.Get("restorable_dropped_database_id").(string)
    restorePointInTime := d.Get("restore_point_in_time").(string)
    sampleName := d.Get("sample_name").(string)
    serverName := d.Get("server_name").(string)
    sku := d.Get("sku").([]interface{})
    sourceDatabaseDeletionDate := d.Get("source_database_deletion_date").(string)
    sourceDatabaseId := d.Get("source_database_id").(string)
    zoneRedundant := d.Get("zone_redundant").(bool)
    t := d.Get("tags").(map[string]interface{})

    parameters := sql.Database{
        Location: utils.String(location),
        DatabaseProperties: &sql.DatabaseProperties{
            AutoPauseDelay: utils.Int32(int32(autoPauseDelay)),
            CatalogCollation: sql.CatalogCollationType(catalogCollation),
            Collation: utils.String(collation),
            CreateMode: sql.CreateMode(createMode),
            ElasticPoolID: utils.String(elasticPoolId),
            LicenseType: sql.DatabaseLicenseType(licenseType),
            LongTermRetentionBackupResourceID: utils.String(longTermRetentionBackupResourceId),
            MaxSizeBytes: utils.Int64(int64(maxSizeBytes)),
            MinCapacity: utils.Float(minCapacity),
            ReadReplicaCount: utils.Int32(int32(readReplicaCount)),
            ReadScale: sql.DatabaseReadScale(readScale),
            RecoverableDatabaseID: utils.String(recoverableDatabaseId),
            RecoveryServicesRecoveryPointID: utils.String(recoveryServicesRecoveryPointId),
            RestorableDroppedDatabaseID: utils.String(restorableDroppedDatabaseId),
            RestorePointInTime: convertStringToDate(restorePointInTime),
            SampleName: sql.SampleName(sampleName),
            SourceDatabaseDeletionDate: convertStringToDate(sourceDatabaseDeletionDate),
            SourceDatabaseID: utils.String(sourceDatabaseId),
            ZoneRedundant: utils.Bool(zoneRedundant),
        },
        Sku: expandArmDatabaseSku(sku),
        Tags: tags.Expand(t),
    }


    future, err := client.Update(ctx, resourceGroup, serverName, name, parameters)
    if err != nil {
        return fmt.Errorf("Error updating Database %q (Server Name %q / Resource Group %q): %+v", name, serverName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for update of Database %q (Server Name %q / Resource Group %q): %+v", name, serverName, resourceGroup, err)
    }

    return resourceArmDatabaseRead(d, meta)
}

func resourceArmDatabaseDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).databasesClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    serverName := id.Path["servers"]
    name := id.Path["databases"]

    future, err := client.Delete(ctx, resourceGroup, serverName, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Database %q (Server Name %q / Resource Group %q): %+v", name, serverName, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Database %q (Server Name %q / Resource Group %q): %+v", name, serverName, resourceGroup, err)
        }
    }

    return nil
}

func convertStringToDate(input interface{}) *date.Time {
  v := input.(string)

  dateTime, err := date.ParseTime(time.RFC3339, v)
  if err != nil {
      log.Printf("[ERROR] Cannot convert an invalid string to RFC3339 date %q: %+v", v, err)
      return nil
  }

  result := date.Time{
      Time: dateTime,
  }
  return &result
}

func expandArmDatabaseSku(input []interface{}) *sql.Sku {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    name := v["name"].(string)
    tier := v["tier"].(string)
    size := v["size"].(string)
    family := v["family"].(string)
    capacity := v["capacity"].(int)

    result := sql.Sku{
        Capacity: utils.Int32(int32(capacity)),
        Family: utils.String(family),
        Name: utils.String(name),
        Size: utils.String(size),
        Tier: utils.String(tier),
    }
    return &result
}


func flattenArmDatabaseSku(input *sql.Sku) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})


    return []interface{}{result}
}
