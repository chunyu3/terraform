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
                Computed: true,
            },

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "cloud_endpoint_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

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

            "azure_file_share_name": {
                Type: schema.TypeString,
                Optional: true,
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

            "friendly_name": {
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
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    cloudEndpointName := d.Get("cloud_endpoint_name").(string)
    storageSyncServiceName := d.Get("storage_sync_service_name").(string)
    syncGroupName := d.Get("sync_group_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, storageSyncServiceName, syncGroupName, cloudEndpointName)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Cloud Endpoint (Cloud Endpoint Name %q / Sync Group Name %q / Storage Sync Service Name %q / Resource Group %q): %+v", cloudEndpointName, syncGroupName, storageSyncServiceName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_cloud_endpoint", *existing.ID)
        }
    }

    azureFileShareName := d.Get("azure_file_share_name").(string)
    storageAccountResourceId := d.Get("storage_account_resource_id").(string)
    storageAccountTenantId := d.Get("storage_account_tenant_id").(string)

    parameters := storagesync.CloudEndpointCreateParameters{
        CloudEndpointCreateParametersProperties: &storagesync.CloudEndpointCreateParametersProperties{
            AzureFileShareName: utils.String(azureFileShareName),
            StorageAccountResourceID: utils.String(storageAccountResourceId),
            StorageAccountTenantID: utils.String(storageAccountTenantId),
        },
    }


    future, err := client.Create(ctx, resourceGroup, storageSyncServiceName, syncGroupName, cloudEndpointName, parameters)
    if err != nil {
        return fmt.Errorf("Error creating Cloud Endpoint (Cloud Endpoint Name %q / Sync Group Name %q / Storage Sync Service Name %q / Resource Group %q): %+v", cloudEndpointName, syncGroupName, storageSyncServiceName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Cloud Endpoint (Cloud Endpoint Name %q / Sync Group Name %q / Storage Sync Service Name %q / Resource Group %q): %+v", cloudEndpointName, syncGroupName, storageSyncServiceName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, storageSyncServiceName, syncGroupName, cloudEndpointName)
    if err != nil {
        return fmt.Errorf("Error retrieving Cloud Endpoint (Cloud Endpoint Name %q / Sync Group Name %q / Storage Sync Service Name %q / Resource Group %q): %+v", cloudEndpointName, syncGroupName, storageSyncServiceName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Cloud Endpoint (Cloud Endpoint Name %q / Sync Group Name %q / Storage Sync Service Name %q / Resource Group %q) ID", cloudEndpointName, syncGroupName, storageSyncServiceName, resourceGroup)
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
    cloudEndpointName := id.Path["cloudEndpoints"]

    resp, err := client.Get(ctx, resourceGroup, storageSyncServiceName, syncGroupName, cloudEndpointName)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Cloud Endpoint %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Cloud Endpoint (Cloud Endpoint Name %q / Sync Group Name %q / Storage Sync Service Name %q / Resource Group %q): %+v", cloudEndpointName, syncGroupName, storageSyncServiceName, resourceGroup, err)
    }


    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
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
    d.Set("cloud_endpoint_name", cloudEndpointName)
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
    cloudEndpointName := id.Path["cloudEndpoints"]

    future, err := client.Delete(ctx, resourceGroup, storageSyncServiceName, syncGroupName, cloudEndpointName)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Cloud Endpoint (Cloud Endpoint Name %q / Sync Group Name %q / Storage Sync Service Name %q / Resource Group %q): %+v", cloudEndpointName, syncGroupName, storageSyncServiceName, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Cloud Endpoint (Cloud Endpoint Name %q / Sync Group Name %q / Storage Sync Service Name %q / Resource Group %q): %+v", cloudEndpointName, syncGroupName, storageSyncServiceName, resourceGroup, err)
        }
    }

    return nil
}
