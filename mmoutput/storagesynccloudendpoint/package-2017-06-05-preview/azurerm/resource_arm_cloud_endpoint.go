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

            "friendly_name": {
                Type: schema.TypeString,
                Optional: true,
            },

            "last_workflow_id": {
                Type: schema.TypeString,
                Optional: true,
            },

            "partnership_id": {
                Type: schema.TypeString,
                Optional: true,
            },

            "storage_account": {
                Type: schema.TypeString,
                Optional: true,
            },

            "storage_account_key": {
                Type: schema.TypeString,
                Optional: true,
            },

            "storage_account_resource_id": {
                Type: schema.TypeString,
                Optional: true,
            },

            "storage_account_share_name": {
                Type: schema.TypeString,
                Optional: true,
            },

            "storage_account_tenant_id": {
                Type: schema.TypeString,
                Optional: true,
            },

            "backup_enabled": {
                Type: schema.TypeBool,
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

    friendlyName := d.Get("friendly_name").(string)
    lastWorkflowId := d.Get("last_workflow_id").(string)
    partnershipId := d.Get("partnership_id").(string)
    storageAccount := d.Get("storage_account").(string)
    storageAccountKey := d.Get("storage_account_key").(string)
    storageAccountResourceId := d.Get("storage_account_resource_id").(string)
    storageAccountShareName := d.Get("storage_account_share_name").(string)
    storageAccountTenantId := d.Get("storage_account_tenant_id").(string)

    parameters := storagesync.CloudEndpoint{
        CloudEndpointProperties: &storagesync.CloudEndpointProperties{
            FriendlyName: utils.String(friendlyName),
            LastWorkflowID: utils.String(lastWorkflowId),
            PartnershipID: utils.String(partnershipId),
            StorageAccount: utils.String(storageAccount),
            StorageAccountKey: utils.String(storageAccountKey),
            StorageAccountResourceID: utils.String(storageAccountResourceId),
            StorageAccountShareName: utils.String(storageAccountShareName),
            StorageAccountTenantID: utils.String(storageAccountTenantId),
        },
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
    if cloudEndpointProperties := resp.CloudEndpointProperties; cloudEndpointProperties != nil {
        d.Set("backup_enabled", cloudEndpointProperties.BackupEnabled)
        d.Set("friendly_name", cloudEndpointProperties.FriendlyName)
        d.Set("last_workflow_id", cloudEndpointProperties.LastWorkflowID)
        d.Set("partnership_id", cloudEndpointProperties.PartnershipID)
        d.Set("provisioning_state", cloudEndpointProperties.ProvisioningState)
        d.Set("storage_account", cloudEndpointProperties.StorageAccount)
        d.Set("storage_account_key", cloudEndpointProperties.StorageAccountKey)
        d.Set("storage_account_resource_id", cloudEndpointProperties.StorageAccountResourceID)
        d.Set("storage_account_share_name", cloudEndpointProperties.StorageAccountShareName)
        d.Set("storage_account_tenant_id", cloudEndpointProperties.StorageAccountTenantID)
    }
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
