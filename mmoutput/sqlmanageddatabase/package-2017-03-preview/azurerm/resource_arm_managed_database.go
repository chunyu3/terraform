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



func resourceArmManagedDatabase() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmManagedDatabaseCreate,
        Read: resourceArmManagedDatabaseRead,
        Update: resourceArmManagedDatabaseUpdate,
        Delete: resourceArmManagedDatabaseDelete,

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

            "database_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "managed_instance_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
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
                    string(sql.RestoreExternalBackup),
                    string(sql.PointInTimeRestore),
                    string(sql.Recovery),
                }, false),
                Default: string(sql.Default),
            },

            "recoverable_database_id": {
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

            "source_database_id": {
                Type: schema.TypeString,
                Optional: true,
            },

            "storage_container_sas_token": {
                Type: schema.TypeString,
                Optional: true,
            },

            "storage_container_uri": {
                Type: schema.TypeString,
                Optional: true,
            },

            "creation_date": {
                Type: schema.TypeString,
                Computed: true,
            },

            "default_secondary_location": {
                Type: schema.TypeString,
                Computed: true,
            },

            "earliest_restore_point": {
                Type: schema.TypeString,
                Computed: true,
            },

            "failover_group_id": {
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

func resourceArmManagedDatabaseCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).managedDatabasesClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    databaseName := d.Get("database_name").(string)
    managedInstanceName := d.Get("managed_instance_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, managedInstanceName, databaseName)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Managed Database (Database Name %q / Managed Instance Name %q / Resource Group %q): %+v", databaseName, managedInstanceName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_managed_database", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    catalogCollation := d.Get("catalog_collation").(string)
    collation := d.Get("collation").(string)
    createMode := d.Get("create_mode").(string)
    recoverableDatabaseId := d.Get("recoverable_database_id").(string)
    restorableDroppedDatabaseId := d.Get("restorable_dropped_database_id").(string)
    restorePointInTime := d.Get("restore_point_in_time").(string)
    sourceDatabaseId := d.Get("source_database_id").(string)
    storageContainerSasToken := d.Get("storage_container_sas_token").(string)
    storageContainerUri := d.Get("storage_container_uri").(string)
    t := d.Get("tags").(map[string]interface{})

    parameters := sql.ManagedDatabase{
        Location: utils.String(location),
        ManagedDatabaseProperties: &sql.ManagedDatabaseProperties{
            CatalogCollation: sql.CatalogCollationType(catalogCollation),
            Collation: utils.String(collation),
            CreateMode: sql.ManagedDatabaseCreateMode(createMode),
            RecoverableDatabaseID: utils.String(recoverableDatabaseId),
            RestorableDroppedDatabaseID: utils.String(restorableDroppedDatabaseId),
            RestorePointInTime: convertStringToDate(restorePointInTime),
            SourceDatabaseID: utils.String(sourceDatabaseId),
            StorageContainerSasToken: utils.String(storageContainerSasToken),
            StorageContainerUri: utils.String(storageContainerUri),
        },
        Tags: tags.Expand(t),
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, managedInstanceName, databaseName, parameters)
    if err != nil {
        return fmt.Errorf("Error creating Managed Database (Database Name %q / Managed Instance Name %q / Resource Group %q): %+v", databaseName, managedInstanceName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Managed Database (Database Name %q / Managed Instance Name %q / Resource Group %q): %+v", databaseName, managedInstanceName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, managedInstanceName, databaseName)
    if err != nil {
        return fmt.Errorf("Error retrieving Managed Database (Database Name %q / Managed Instance Name %q / Resource Group %q): %+v", databaseName, managedInstanceName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Managed Database (Database Name %q / Managed Instance Name %q / Resource Group %q) ID", databaseName, managedInstanceName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmManagedDatabaseRead(d, meta)
}

func resourceArmManagedDatabaseRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).managedDatabasesClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    managedInstanceName := id.Path["managedInstances"]
    databaseName := id.Path["databases"]

    resp, err := client.Get(ctx, resourceGroup, managedInstanceName, databaseName)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Managed Database %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Managed Database (Database Name %q / Managed Instance Name %q / Resource Group %q): %+v", databaseName, managedInstanceName, resourceGroup, err)
    }


    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if managedDatabaseProperties := resp.ManagedDatabaseProperties; managedDatabaseProperties != nil {
        d.Set("catalog_collation", string(managedDatabaseProperties.CatalogCollation))
        d.Set("collation", managedDatabaseProperties.Collation)
        d.Set("create_mode", string(managedDatabaseProperties.CreateMode))
        d.Set("creation_date", (managedDatabaseProperties.CreationDate).String())
        d.Set("default_secondary_location", managedDatabaseProperties.DefaultSecondaryLocation)
        d.Set("earliest_restore_point", (managedDatabaseProperties.EarliestRestorePoint).String())
        d.Set("failover_group_id", managedDatabaseProperties.FailoverGroupID)
        d.Set("recoverable_database_id", managedDatabaseProperties.RecoverableDatabaseID)
        d.Set("restorable_dropped_database_id", managedDatabaseProperties.RestorableDroppedDatabaseID)
        d.Set("restore_point_in_time", (managedDatabaseProperties.RestorePointInTime).String())
        d.Set("source_database_id", managedDatabaseProperties.SourceDatabaseID)
        d.Set("status", string(managedDatabaseProperties.Status))
        d.Set("storage_container_sas_token", managedDatabaseProperties.StorageContainerSasToken)
        d.Set("storage_container_uri", managedDatabaseProperties.StorageContainerUri)
    }
    d.Set("database_name", databaseName)
    d.Set("managed_instance_name", managedInstanceName)
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}

func resourceArmManagedDatabaseUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).managedDatabasesClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    catalogCollation := d.Get("catalog_collation").(string)
    collation := d.Get("collation").(string)
    createMode := d.Get("create_mode").(string)
    databaseName := d.Get("database_name").(string)
    managedInstanceName := d.Get("managed_instance_name").(string)
    recoverableDatabaseId := d.Get("recoverable_database_id").(string)
    restorableDroppedDatabaseId := d.Get("restorable_dropped_database_id").(string)
    restorePointInTime := d.Get("restore_point_in_time").(string)
    sourceDatabaseId := d.Get("source_database_id").(string)
    storageContainerSasToken := d.Get("storage_container_sas_token").(string)
    storageContainerUri := d.Get("storage_container_uri").(string)
    t := d.Get("tags").(map[string]interface{})

    parameters := sql.ManagedDatabase{
        Location: utils.String(location),
        ManagedDatabaseProperties: &sql.ManagedDatabaseProperties{
            CatalogCollation: sql.CatalogCollationType(catalogCollation),
            Collation: utils.String(collation),
            CreateMode: sql.ManagedDatabaseCreateMode(createMode),
            RecoverableDatabaseID: utils.String(recoverableDatabaseId),
            RestorableDroppedDatabaseID: utils.String(restorableDroppedDatabaseId),
            RestorePointInTime: convertStringToDate(restorePointInTime),
            SourceDatabaseID: utils.String(sourceDatabaseId),
            StorageContainerSasToken: utils.String(storageContainerSasToken),
            StorageContainerUri: utils.String(storageContainerUri),
        },
        Tags: tags.Expand(t),
    }


    future, err := client.Update(ctx, resourceGroup, managedInstanceName, databaseName, parameters)
    if err != nil {
        return fmt.Errorf("Error updating Managed Database (Database Name %q / Managed Instance Name %q / Resource Group %q): %+v", databaseName, managedInstanceName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for update of Managed Database (Database Name %q / Managed Instance Name %q / Resource Group %q): %+v", databaseName, managedInstanceName, resourceGroup, err)
    }

    return resourceArmManagedDatabaseRead(d, meta)
}

func resourceArmManagedDatabaseDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).managedDatabasesClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    managedInstanceName := id.Path["managedInstances"]
    databaseName := id.Path["databases"]

    future, err := client.Delete(ctx, resourceGroup, managedInstanceName, databaseName)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Managed Database (Database Name %q / Managed Instance Name %q / Resource Group %q): %+v", databaseName, managedInstanceName, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Managed Database (Database Name %q / Managed Instance Name %q / Resource Group %q): %+v", databaseName, managedInstanceName, resourceGroup, err)
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
