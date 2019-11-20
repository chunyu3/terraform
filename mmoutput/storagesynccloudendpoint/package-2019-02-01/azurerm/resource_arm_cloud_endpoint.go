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



func resourceArmCloudEndpoint() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmCloudEndpointCreateUpdate,
        Read: resourceArmCloudEndpointRead,
        Update: resourceArmCloudEndpointCreateUpdate,
        Delete: resourceArmCloudEndpointDelete,

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

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "storage_sync_service_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "sync_group_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "azure_file_share": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
            },

            "azure_file_share_name": {
                Type: schema.TypeString,
                Optional: true,
            },

            "azure_file_share_uri": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
            },

            "backup_metadata_property_bag": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
            },

            "failed_file_list": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
            },

            "partition": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
            },

            "pause_wait_for_sync_drain_time_period_in_seconds": {
                Type: schema.TypeInt,
                Optional: true,
                ForceNew: true,
            },

            "replica_group": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
            },

            "request_id": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
            },

            "restore_file_spec": {
                Type: schema.TypeList,
                Optional: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "path": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },

            "source_azure_file_share_uri": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
            },

            "status": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
            },

            "storage_account_resource_id": {
                Type: schema.TypeString,
                Optional: true,
            },

            "storage_account_tenant_id": {
                Type: schema.TypeString,
                Optional: true,
            },

            "type": {
                Type: schema.TypeString,
                Computed: true,
            },
        },
    }
}

func resourceArmCloudEndpointCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).cloudEndpointsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    storageSyncServiceName := d.Get("storage_sync_service_name").(string)
    syncGroupName := d.Get("sync_group_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, storageSyncServiceName, syncGroupName, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Cloud Endpoint %q (Sync Group Name %q / Storage Sync Service Name %q / Resource Group %q): %+v", name, syncGroupName, storageSyncServiceName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_cloud_endpoint", *existing.ID)
        }
    }

    azureFileShare := d.Get("azure_file_share").(string)
    azureFileShareName := d.Get("azure_file_share_name").(string)
    azureFileShareUri := d.Get("azure_file_share_uri").(string)
    backupMetadataPropertyBag := d.Get("backup_metadata_property_bag").(string)
    failedFileList := d.Get("failed_file_list").(string)
    partition := d.Get("partition").(string)
    pauseWaitForSyncDrainTimePeriodInSeconds := d.Get("pause_wait_for_sync_drain_time_period_in_seconds").(int)
    replicaGroup := d.Get("replica_group").(string)
    requestId := d.Get("request_id").(string)
    restoreFileSpec := d.Get("restore_file_spec").([]interface{})
    sourceAzureFileShareUri := d.Get("source_azure_file_share_uri").(string)
    status := d.Get("status").(string)
    storageAccountResourceId := d.Get("storage_account_resource_id").(string)
    storageAccountTenantId := d.Get("storage_account_tenant_id").(string)

    parameters := storagesync.PostRestoreRequest{
        AzureFileShare: utils.String(azureFileShare),
        AzureFileShareURI: utils.String(azureFileShareUri),
        BackupMetadataPropertyBag: utils.String(backupMetadataPropertyBag),
        FailedFileList: utils.String(failedFileList),
        Partition: utils.String(partition),
        PauseWaitForSyncDrainTimePeriodInSeconds: utils.Int(pauseWaitForSyncDrainTimePeriodInSeconds),
        CloudEndpointCreateParametersProperties: &storagesync.CloudEndpointCreateParametersProperties{
            AzureFileShareName: utils.String(azureFileShareName),
            StorageAccountResourceID: utils.String(storageAccountResourceId),
            StorageAccountTenantID: utils.String(storageAccountTenantId),
        },
        ReplicaGroup: utils.String(replicaGroup),
        RequestID: utils.String(requestId),
        RestoreFileSpec: expandArmCloudEndpointRestoreFileSpec(restoreFileSpec),
        SourceAzureFileShareURI: utils.String(sourceAzureFileShareUri),
        Status: utils.String(status),
    }


    future, err := client.Create(ctx, resourceGroup, storageSyncServiceName, syncGroupName, name, parameters)
    if err != nil {
        return fmt.Errorf("Error creating Cloud Endpoint %q (Sync Group Name %q / Storage Sync Service Name %q / Resource Group %q): %+v", name, syncGroupName, storageSyncServiceName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Cloud Endpoint %q (Sync Group Name %q / Storage Sync Service Name %q / Resource Group %q): %+v", name, syncGroupName, storageSyncServiceName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, storageSyncServiceName, syncGroupName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Cloud Endpoint %q (Sync Group Name %q / Storage Sync Service Name %q / Resource Group %q): %+v", name, syncGroupName, storageSyncServiceName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Cloud Endpoint %q (Sync Group Name %q / Storage Sync Service Name %q / Resource Group %q) ID", name, syncGroupName, storageSyncServiceName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmCloudEndpointRead(d, meta)
}

func resourceArmCloudEndpointRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).cloudEndpointsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    storageSyncServiceName := id.Path["storageSyncServices"]
    syncGroupName := id.Path["syncGroups"]
    name := id.Path["cloudEndpoints"]

    resp, err := client.Get(ctx, resourceGroup, storageSyncServiceName, syncGroupName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Cloud Endpoint %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Cloud Endpoint %q (Sync Group Name %q / Storage Sync Service Name %q / Resource Group %q): %+v", name, syncGroupName, storageSyncServiceName, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    d.Set("storage_sync_service_name", storageSyncServiceName)
    d.Set("sync_group_name", syncGroupName)
    d.Set("type", resp.Type)

    return nil
}


func resourceArmCloudEndpointDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).cloudEndpointsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    storageSyncServiceName := id.Path["storageSyncServices"]
    syncGroupName := id.Path["syncGroups"]
    name := id.Path["cloudEndpoints"]

    future, err := client.Delete(ctx, resourceGroup, storageSyncServiceName, syncGroupName, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Cloud Endpoint %q (Sync Group Name %q / Storage Sync Service Name %q / Resource Group %q): %+v", name, syncGroupName, storageSyncServiceName, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Cloud Endpoint %q (Sync Group Name %q / Storage Sync Service Name %q / Resource Group %q): %+v", name, syncGroupName, storageSyncServiceName, resourceGroup, err)
        }
    }

    return nil
}

func expandArmCloudEndpointRestoreFileSpec(input []interface{}) *[]storagesync.RestoreFileSpec {
    results := make([]storagesync.RestoreFileSpec, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        path := v["path"].(string)

        result := storagesync.RestoreFileSpec{
            Path: utils.String(path),
        }

        results = append(results, result)
    }
    return &results
}
