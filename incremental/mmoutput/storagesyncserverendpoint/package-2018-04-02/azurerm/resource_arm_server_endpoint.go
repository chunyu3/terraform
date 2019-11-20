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

            "location": azure.SchemaLocation(),

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

            "cloud_tiering": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(storagesync.on),
                    string(storagesync.off),
                }, false),
                Default: string(storagesync.on),
            },

            "pattern": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
            },

            "recall_path": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
            },

            "volume_free_space_percent": {
                Type: schema.TypeInt,
                Optional: true,
            },

            "type": {
                Type: schema.TypeString,
                Computed: true,
            },

            "tags": tags.Schema(),
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

    location := azure.NormalizeLocation(d.Get("location").(string))
    cloudTiering := d.Get("cloud_tiering").(string)
    pattern := d.Get("pattern").(string)
    recallPath := d.Get("recall_path").(string)
    volumeFreeSpacePercent := d.Get("volume_free_space_percent").(int)
    t := d.Get("tags").(map[string]interface{})

    parameters := storagesync.RecallActionParameters{
        Location: utils.String(location),
        Pattern: utils.String(pattern),
        ServerEndpointUpdateProperties: &storagesync.ServerEndpointUpdateProperties{
            CloudTiering: storagesync.(cloudTiering),
            VolumeFreeSpacePercent: utils.Int(volumeFreeSpacePercent),
        },
        RecallPath: utils.String(recallPath),
        Tags: tags.Expand(t),
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
    cloudTiering := d.Get("cloud_tiering").(string)
    pattern := d.Get("pattern").(string)
    recallPath := d.Get("recall_path").(string)
    storageSyncServiceName := d.Get("storage_sync_service_name").(string)
    syncGroupName := d.Get("sync_group_name").(string)
    volumeFreeSpacePercent := d.Get("volume_free_space_percent").(int)
    t := d.Get("tags").(map[string]interface{})

    parameters := storagesync.RecallActionParameters{
        Pattern: utils.String(pattern),
        ServerEndpointUpdateProperties: &storagesync.ServerEndpointUpdateProperties{
            CloudTiering: storagesync.(cloudTiering),
            VolumeFreeSpacePercent: utils.Int(volumeFreeSpacePercent),
        },
        RecallPath: utils.String(recallPath),
        Tags: tags.Expand(t),
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
