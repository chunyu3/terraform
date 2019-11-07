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



func resourceArmServerEndpoint() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmServerEndpointCreate,
        Read: resourceArmServerEndpointRead,
        Update: resourceArmServerEndpointUpdate,
        Delete: resourceArmServerEndpointDelete,

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

            "byte_progress": {
                Type: schema.TypeInt,
                Optional: true,
            },

            "byte_total": {
                Type: schema.TypeInt,
                Optional: true,
            },

            "cloud_tiering": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(storagesync.on),
                    string(storagesync.off),
                }, false),
                Default: string(storagesync.on),
            },

            "current_progress_type": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(storagesync.none),
                    string(storagesync.initialize),
                    string(storagesync.download),
                    string(storagesync.upload),
                    string(storagesync.recall),
                }, false),
                Default: string(storagesync.none),
            },

            "friendly_name": {
                Type: schema.TypeString,
                Optional: true,
            },

            "item_download_error_count": {
                Type: schema.TypeInt,
                Optional: true,
            },

            "item_progress_count": {
                Type: schema.TypeInt,
                Optional: true,
            },

            "item_total_count": {
                Type: schema.TypeInt,
                Optional: true,
            },

            "item_upload_error_count": {
                Type: schema.TypeInt,
                Optional: true,
            },

            "last_sync_success": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validateRFC3339Date,
            },

            "last_workflow_id": {
                Type: schema.TypeString,
                Optional: true,
            },

            "server_local_path": {
                Type: schema.TypeString,
                Optional: true,
            },

            "server_resource_id": {
                Type: schema.TypeString,
                Optional: true,
            },

            "sync_error_context": {
                Type: schema.TypeString,
                Optional: true,
            },

            "sync_error_direction": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(storagesync.none),
                    string(storagesync.initialize),
                    string(storagesync.download),
                    string(storagesync.upload),
                    string(storagesync.recall),
                }, false),
                Default: string(storagesync.none),
            },

            "sync_error_state": {
                Type: schema.TypeString,
                Optional: true,
            },

            "sync_error_state_timestamp": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validateRFC3339Date,
            },

            "total_progress": {
                Type: schema.TypeInt,
                Optional: true,
            },

            "volume_free_space_percent": {
                Type: schema.TypeInt,
                Optional: true,
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

func resourceArmServerEndpointCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).serverEndpointsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    storageSyncServiceName := d.Get("storage_sync_service_name").(string)
    syncGroupName := d.Get("sync_group_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, storageSyncServiceName, syncGroupName, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Server Endpoint %q (Sync Group Name %q / Storage Sync Service Name %q / Resource Group %q): %+v", name, syncGroupName, storageSyncServiceName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_server_endpoint", *existing.ID)
        }
    }

    byteProgress := d.Get("byte_progress").(int)
    byteTotal := d.Get("byte_total").(int)
    cloudTiering := d.Get("cloud_tiering").(string)
    currentProgressType := d.Get("current_progress_type").(string)
    friendlyName := d.Get("friendly_name").(string)
    itemDownloadErrorCount := d.Get("item_download_error_count").(int)
    itemProgressCount := d.Get("item_progress_count").(int)
    itemTotalCount := d.Get("item_total_count").(int)
    itemUploadErrorCount := d.Get("item_upload_error_count").(int)
    lastSyncSuccess := d.Get("last_sync_success").(string)
    lastWorkflowId := d.Get("last_workflow_id").(string)
    serverLocalPath := d.Get("server_local_path").(string)
    serverResourceId := d.Get("server_resource_id").(string)
    syncErrorContext := d.Get("sync_error_context").(string)
    syncErrorDirection := d.Get("sync_error_direction").(string)
    syncErrorState := d.Get("sync_error_state").(string)
    syncErrorStateTimestamp := d.Get("sync_error_state_timestamp").(string)
    totalProgress := d.Get("total_progress").(int)
    volumeFreeSpacePercent := d.Get("volume_free_space_percent").(int)

    parameters := storagesync.ServerEndpoint{
        ServerEndpointProperties: &storagesync.ServerEndpointProperties{
            ByteProgress: utils.Int(byteProgress),
            ByteTotal: utils.Int(byteTotal),
            CloudTiering: storagesync.(cloudTiering),
            CurrentProgressType: storagesync.(currentProgressType),
            FriendlyName: utils.String(friendlyName),
            ItemDownloadErrorCount: utils.Int(itemDownloadErrorCount),
            ItemProgressCount: utils.Int(itemProgressCount),
            ItemTotalCount: utils.Int(itemTotalCount),
            ItemUploadErrorCount: utils.Int(itemUploadErrorCount),
            LastSyncSuccess: convertStringToDate(lastSyncSuccess),
            LastWorkflowID: utils.String(lastWorkflowId),
            ServerLocalPath: utils.String(serverLocalPath),
            ServerResourceID: utils.String(serverResourceId),
            SyncErrorContext: utils.String(syncErrorContext),
            SyncErrorDirection: storagesync.(syncErrorDirection),
            SyncErrorState: utils.String(syncErrorState),
            SyncErrorStateTimestamp: convertStringToDate(syncErrorStateTimestamp),
            TotalProgress: utils.Int(totalProgress),
            VolumeFreeSpacePercent: utils.Int(volumeFreeSpacePercent),
        },
    }


    future, err := client.Create(ctx, resourceGroup, storageSyncServiceName, syncGroupName, name, parameters)
    if err != nil {
        return fmt.Errorf("Error creating Server Endpoint %q (Sync Group Name %q / Storage Sync Service Name %q / Resource Group %q): %+v", name, syncGroupName, storageSyncServiceName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Server Endpoint %q (Sync Group Name %q / Storage Sync Service Name %q / Resource Group %q): %+v", name, syncGroupName, storageSyncServiceName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, storageSyncServiceName, syncGroupName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Server Endpoint %q (Sync Group Name %q / Storage Sync Service Name %q / Resource Group %q): %+v", name, syncGroupName, storageSyncServiceName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Server Endpoint %q (Sync Group Name %q / Storage Sync Service Name %q / Resource Group %q) ID", name, syncGroupName, storageSyncServiceName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmServerEndpointRead(d, meta)
}

func resourceArmServerEndpointRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).serverEndpointsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    storageSyncServiceName := id.Path["storageSyncServices"]
    syncGroupName := id.Path["syncGroups"]
    name := id.Path["serverEndpoints"]

    resp, err := client.Get(ctx, resourceGroup, storageSyncServiceName, syncGroupName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Server Endpoint %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Server Endpoint %q (Sync Group Name %q / Storage Sync Service Name %q / Resource Group %q): %+v", name, syncGroupName, storageSyncServiceName, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if serverEndpointProperties := resp.ServerEndpointProperties; serverEndpointProperties != nil {
        d.Set("byte_progress", serverEndpointProperties.ByteProgress)
        d.Set("byte_total", serverEndpointProperties.ByteTotal)
        d.Set("cloud_tiering", string(serverEndpointProperties.CloudTiering))
        d.Set("current_progress_type", string(serverEndpointProperties.CurrentProgressType))
        d.Set("friendly_name", serverEndpointProperties.FriendlyName)
        d.Set("item_download_error_count", serverEndpointProperties.ItemDownloadErrorCount)
        d.Set("item_progress_count", serverEndpointProperties.ItemProgressCount)
        d.Set("item_total_count", serverEndpointProperties.ItemTotalCount)
        d.Set("item_upload_error_count", serverEndpointProperties.ItemUploadErrorCount)
        d.Set("last_sync_success", (serverEndpointProperties.LastSyncSuccess).String())
        d.Set("last_workflow_id", serverEndpointProperties.LastWorkflowID)
        d.Set("provisioning_state", serverEndpointProperties.ProvisioningState)
        d.Set("server_local_path", serverEndpointProperties.ServerLocalPath)
        d.Set("server_resource_id", serverEndpointProperties.ServerResourceID)
        d.Set("sync_error_context", serverEndpointProperties.SyncErrorContext)
        d.Set("sync_error_direction", string(serverEndpointProperties.SyncErrorDirection))
        d.Set("sync_error_state", serverEndpointProperties.SyncErrorState)
        d.Set("sync_error_state_timestamp", (serverEndpointProperties.SyncErrorStateTimestamp).String())
        d.Set("total_progress", serverEndpointProperties.TotalProgress)
        d.Set("volume_free_space_percent", serverEndpointProperties.VolumeFreeSpacePercent)
    }
    d.Set("storage_sync_service_name", storageSyncServiceName)
    d.Set("sync_group_name", syncGroupName)
    d.Set("type", resp.Type)

    return nil
}

func resourceArmServerEndpointUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).serverEndpointsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    byteProgress := d.Get("byte_progress").(int)
    byteTotal := d.Get("byte_total").(int)
    cloudTiering := d.Get("cloud_tiering").(string)
    currentProgressType := d.Get("current_progress_type").(string)
    friendlyName := d.Get("friendly_name").(string)
    itemDownloadErrorCount := d.Get("item_download_error_count").(int)
    itemProgressCount := d.Get("item_progress_count").(int)
    itemTotalCount := d.Get("item_total_count").(int)
    itemUploadErrorCount := d.Get("item_upload_error_count").(int)
    lastSyncSuccess := d.Get("last_sync_success").(string)
    lastWorkflowId := d.Get("last_workflow_id").(string)
    serverLocalPath := d.Get("server_local_path").(string)
    serverResourceId := d.Get("server_resource_id").(string)
    storageSyncServiceName := d.Get("storage_sync_service_name").(string)
    syncErrorContext := d.Get("sync_error_context").(string)
    syncErrorDirection := d.Get("sync_error_direction").(string)
    syncErrorState := d.Get("sync_error_state").(string)
    syncErrorStateTimestamp := d.Get("sync_error_state_timestamp").(string)
    syncGroupName := d.Get("sync_group_name").(string)
    totalProgress := d.Get("total_progress").(int)
    volumeFreeSpacePercent := d.Get("volume_free_space_percent").(int)

    parameters := storagesync.ServerEndpoint{
        ServerEndpointProperties: &storagesync.ServerEndpointProperties{
            ByteProgress: utils.Int(byteProgress),
            ByteTotal: utils.Int(byteTotal),
            CloudTiering: storagesync.(cloudTiering),
            CurrentProgressType: storagesync.(currentProgressType),
            FriendlyName: utils.String(friendlyName),
            ItemDownloadErrorCount: utils.Int(itemDownloadErrorCount),
            ItemProgressCount: utils.Int(itemProgressCount),
            ItemTotalCount: utils.Int(itemTotalCount),
            ItemUploadErrorCount: utils.Int(itemUploadErrorCount),
            LastSyncSuccess: convertStringToDate(lastSyncSuccess),
            LastWorkflowID: utils.String(lastWorkflowId),
            ServerLocalPath: utils.String(serverLocalPath),
            ServerResourceID: utils.String(serverResourceId),
            SyncErrorContext: utils.String(syncErrorContext),
            SyncErrorDirection: storagesync.(syncErrorDirection),
            SyncErrorState: utils.String(syncErrorState),
            SyncErrorStateTimestamp: convertStringToDate(syncErrorStateTimestamp),
            TotalProgress: utils.Int(totalProgress),
            VolumeFreeSpacePercent: utils.Int(volumeFreeSpacePercent),
        },
    }


    future, err := client.Update(ctx, resourceGroup, storageSyncServiceName, syncGroupName, name, parameters)
    if err != nil {
        return fmt.Errorf("Error updating Server Endpoint %q (Sync Group Name %q / Storage Sync Service Name %q / Resource Group %q): %+v", name, syncGroupName, storageSyncServiceName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for update of Server Endpoint %q (Sync Group Name %q / Storage Sync Service Name %q / Resource Group %q): %+v", name, syncGroupName, storageSyncServiceName, resourceGroup, err)
    }

    return resourceArmServerEndpointRead(d, meta)
}

func resourceArmServerEndpointDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).serverEndpointsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    storageSyncServiceName := id.Path["storageSyncServices"]
    syncGroupName := id.Path["syncGroups"]
    name := id.Path["serverEndpoints"]

    future, err := client.Delete(ctx, resourceGroup, storageSyncServiceName, syncGroupName, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Server Endpoint %q (Sync Group Name %q / Storage Sync Service Name %q / Resource Group %q): %+v", name, syncGroupName, storageSyncServiceName, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Server Endpoint %q (Sync Group Name %q / Storage Sync Service Name %q / Resource Group %q): %+v", name, syncGroupName, storageSyncServiceName, resourceGroup, err)
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
