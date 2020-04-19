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
            "cloud_endpoint_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
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

            "change_detection_mode": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(storagesync.Default),
                    string(storagesync.Recursive),
                }, false),
                Default: string(storagesync.Default),
            },

            "directory_path": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
            },

            "failed_file_list": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
            },

            "friendly_name": {
                Type: schema.TypeString,
                Optional: true,
            },

            "partition": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
            },

            "paths": {
                Type: schema.TypeList,
                Optional: true,
                ForceNew: true,
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
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
                        "isdir": {
                            Type: schema.TypeBool,
                            Optional: true,
                        },
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

            "backup_enabled": {
                Type: schema.TypeString,
                Computed: true,
            },

            "id": {
                Type: schema.TypeString,
                Computed: true,
            },

            "last_operation_name": {
                Type: schema.TypeString,
                Computed: true,
            },

            "last_workflow_id": {
                Type: schema.TypeString,
                Computed: true,
            },

            "name": {
                Type: schema.TypeString,
                Computed: true,
            },

            "partnership_id": {
                Type: schema.TypeString,
                Computed: true,
            },

            "provisioning_state": {
                Type: schema.TypeString,
                Computed: true,
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
    ctx, cancel := timeouts.ForCreateUpdate(meta.(*ArmClient).StopContext, d)
    defer cancel()

    resourceGroupName := d.Get("resource_group").(string)
    name := d.Get("cloud_endpoint_name").(string)
    storageSyncServiceName := d.Get("storage_sync_service_name").(string)
    syncGroupName := d.Get("sync_group_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroupName, storageSyncServiceName, syncGroupName, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Cloud Endpoint (Cloud Endpoint Name %q / Sync Group Name %q / Storage Sync Service Name %q / Resource Group %q): %+v", name, syncGroupName, storageSyncServiceName, resourceGroupName, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_cloud_endpoint", *existing.ID)
        }
    }

    azureFileShare := d.Get("azure_file_share").(string)
    azureFileShareName := d.Get("azure_file_share_name").(string)
    azureFileShareURI := d.Get("azure_file_share_uri").(string)
    backupMetadataPropertyBag := d.Get("backup_metadata_property_bag").(string)
    changeDetectionMode := d.Get("change_detection_mode").(string)
    directoryPath := d.Get("directory_path").(string)
    failedFileList := d.Get("failed_file_list").(string)
    friendlyName := d.Get("friendly_name").(string)
    partition := d.Get("partition").(string)
    paths := d.Get("paths").([]interface{})
    pauseWaitForSyncDrainTimePeriodInSeconds := d.Get("pause_wait_for_sync_drain_time_period_in_seconds").(int)
    replicaGroup := d.Get("replica_group").(string)
    requestID := d.Get("request_id").(string)
    restoreFileSpec := d.Get("restore_file_spec").([]interface{})
    sourceAzureFileShareURI := d.Get("source_azure_file_share_uri").(string)
    status := d.Get("status").(string)
    storageAccountResourceID := d.Get("storage_account_resource_id").(string)
    storageAccountTenantID := d.Get("storage_account_tenant_id").(string)

    parameters := storagesync.TriggerChangeDetectionParameters{
        AzureFileShare: utils.String(azureFileShare),
        AzureFileShareURI: utils.String(azureFileShareURI),
        BackupMetadataPropertyBag: utils.String(backupMetadataPropertyBag),
        ChangeDetectionMode: storagesync.ChangeDetectionMode(changeDetectionMode),
        DirectoryPath: utils.String(directoryPath),
        FailedFileList: utils.String(failedFileList),
        Partition: utils.String(partition),
        Paths: utils.ExpandStringSlice(paths),
        PauseWaitForSyncDrainTimePeriodInSeconds: utils.Int(pauseWaitForSyncDrainTimePeriodInSeconds),
        CloudEndpointCreateParametersProperties: &storagesync.CloudEndpointCreateParametersProperties{
            AzureFileShareName: utils.String(azureFileShareName),
            FriendlyName: utils.String(friendlyName),
            StorageAccountResourceID: utils.String(storageAccountResourceID),
            StorageAccountTenantID: utils.String(storageAccountTenantID),
        },
        ReplicaGroup: utils.String(replicaGroup),
        RequestID: utils.String(requestID),
        RestoreFileSpec: expandArmCloudEndpointRestoreFileSpec(restoreFileSpec),
        SourceAzureFileShareURI: utils.String(sourceAzureFileShareURI),
        Status: utils.String(status),
    }


    future, err := client.Create(ctx, resourceGroupName, storageSyncServiceName, syncGroupName, name, parameters)
    if err != nil {
        return fmt.Errorf("Error creating Cloud Endpoint (Cloud Endpoint Name %q / Sync Group Name %q / Storage Sync Service Name %q / Resource Group %q): %+v", name, syncGroupName, storageSyncServiceName, resourceGroupName, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Cloud Endpoint (Cloud Endpoint Name %q / Sync Group Name %q / Storage Sync Service Name %q / Resource Group %q): %+v", name, syncGroupName, storageSyncServiceName, resourceGroupName, err)
    }


    resp, err := client.Get(ctx, resourceGroupName, storageSyncServiceName, syncGroupName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Cloud Endpoint (Cloud Endpoint Name %q / Sync Group Name %q / Storage Sync Service Name %q / Resource Group %q): %+v", name, syncGroupName, storageSyncServiceName, resourceGroupName, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Cloud Endpoint (Cloud Endpoint Name %q / Sync Group Name %q / Storage Sync Service Name %q / Resource Group %q) ID", name, syncGroupName, storageSyncServiceName, resourceGroupName)
    }
    d.SetId(*resp.ID)

    return resourceArmCloudEndpointRead(d, meta)
}

func resourceArmCloudEndpointRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).cloudEndpointsClient
    ctx, cancel := timeouts.ForRead(meta.(*ArmClient).StopContext, d)
    defer cancel()

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroupName := id.ResourceGroup
    storageSyncServiceName := id.Path["storageSyncServices"]
    syncGroupName := id.Path["syncGroups"]
    name := id.Path["cloudEndpoints"]

    resp, err := client.Get(ctx, resourceGroupName, storageSyncServiceName, syncGroupName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Cloud Endpoint %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Cloud Endpoint (Cloud Endpoint Name %q / Sync Group Name %q / Storage Sync Service Name %q / Resource Group %q): %+v", name, syncGroupName, storageSyncServiceName, resourceGroupName, err)
    }


    d.Set("resource_group", resourceGroupName)
    if cloudEndpointCreateParametersProperties := resp.CloudEndpointCreateParametersProperties; cloudEndpointCreateParametersProperties != nil {
        d.Set("azure_file_share_name", cloudEndpointCreateParametersProperties.AzureFileShareName)
        d.Set("backup_enabled", cloudEndpointCreateParametersProperties.BackupEnabled)
        d.Set("friendly_name", cloudEndpointCreateParametersProperties.FriendlyName)
        d.Set("last_operation_name", cloudEndpointCreateParametersProperties.LastOperationName)
        d.Set("last_workflow_id", cloudEndpointCreateParametersProperties.LastWorkflowID)
        d.Set("partnership_id", cloudEndpointCreateParametersProperties.PartnershipID)
        d.Set("provisioning_state", cloudEndpointCreateParametersProperties.ProvisioningState)
        d.Set("storage_account_resource_id", cloudEndpointCreateParametersProperties.StorageAccountResourceID)
        d.Set("storage_account_tenant_id", cloudEndpointCreateParametersProperties.StorageAccountTenantID)
    }
    d.Set("cloud_endpoint_name", name)
    d.Set("id", resp.ID)
    d.Set("name", resp.Name)
    d.Set("storage_sync_service_name", storageSyncServiceName)
    d.Set("sync_group_name", syncGroupName)
    d.Set("type", resp.Type)

    return nil
}


func resourceArmCloudEndpointDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).cloudEndpointsClient
    ctx, cancel := timeouts.ForDelete(meta.(*ArmClient).StopContext, d)
    defer cancel()


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroupName := id.ResourceGroup
    storageSyncServiceName := id.Path["storageSyncServices"]
    syncGroupName := id.Path["syncGroups"]
    name := id.Path["cloudEndpoints"]

    future, err := client.Delete(ctx, resourceGroupName, storageSyncServiceName, syncGroupName, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Cloud Endpoint (Cloud Endpoint Name %q / Sync Group Name %q / Storage Sync Service Name %q / Resource Group %q): %+v", name, syncGroupName, storageSyncServiceName, resourceGroupName, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Cloud Endpoint (Cloud Endpoint Name %q / Sync Group Name %q / Storage Sync Service Name %q / Resource Group %q): %+v", name, syncGroupName, storageSyncServiceName, resourceGroupName, err)
        }
    }

    return nil
}

func expandArmCloudEndpointRestoreFileSpec(input []interface{}) *[]storagesync.RestoreFileSpec {
    results := make([]storagesync.RestoreFileSpec, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        path := v["path"].(string)
        isdir := v["isdir"].(bool)

        result := storagesync.RestoreFileSpec{
            Isdir: utils.Bool(isdir),
            Path: utils.String(path),
        }

        results = append(results, result)
    }
    return &results
}