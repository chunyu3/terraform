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



func resourceArmStorageSyncService() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmStorageSyncServiceCreate,
        Read: resourceArmStorageSyncServiceRead,
        Update: resourceArmStorageSyncServiceUpdate,
        Delete: resourceArmStorageSyncServiceDelete,

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

            "storage_sync_service_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "storage_sync_service_status": {
                Type: schema.TypeInt,
                Optional: true,
            },

            "storage_sync_service_uid": {
                Type: schema.TypeString,
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

func resourceArmStorageSyncServiceCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).storageSyncServicesClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    storageSyncServiceName := d.Get("storage_sync_service_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, storageSyncServiceName)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Storage Sync Service (Storage Sync Service Name %q / Resource Group %q): %+v", storageSyncServiceName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_storage_sync_service", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    storageSyncServiceStatus := d.Get("storage_sync_service_status").(int)
    storageSyncServiceUid := d.Get("storage_sync_service_uid").(string)
    t := d.Get("tags").(map[string]interface{})

    parameters := storagesync.ServiceCreateParameters{
        Location: utils.String(location),
        // TODO: SDK Reference /properties is not supported
        Tags: tags.Expand(t),
    }


    if _, err := client.Create(ctx, resourceGroup, storageSyncServiceName, parameters); err != nil {
        return fmt.Errorf("Error creating Storage Sync Service (Storage Sync Service Name %q / Resource Group %q): %+v", storageSyncServiceName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, storageSyncServiceName)
    if err != nil {
        return fmt.Errorf("Error retrieving Storage Sync Service (Storage Sync Service Name %q / Resource Group %q): %+v", storageSyncServiceName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Storage Sync Service (Storage Sync Service Name %q / Resource Group %q) ID", storageSyncServiceName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmStorageSyncServiceRead(d, meta)
}

func resourceArmStorageSyncServiceRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).storageSyncServicesClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    storageSyncServiceName := id.Path["storageSyncServices"]

    resp, err := client.Get(ctx, resourceGroup, storageSyncServiceName)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Storage Sync Service %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Storage Sync Service (Storage Sync Service Name %q / Resource Group %q): %+v", storageSyncServiceName, resourceGroup, err)
    }


    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    d.Set("storage_sync_service_name", storageSyncServiceName)
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}

func resourceArmStorageSyncServiceUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).storageSyncServicesClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    storageSyncServiceName := d.Get("storage_sync_service_name").(string)
    storageSyncServiceStatus := d.Get("storage_sync_service_status").(int)
    storageSyncServiceUid := d.Get("storage_sync_service_uid").(string)
    t := d.Get("tags").(map[string]interface{})

    parameters := storagesync.ServiceCreateParameters{
        Location: utils.String(location),
        // TODO: SDK Reference /properties is not supported
        Tags: tags.Expand(t),
    }


    if _, err := client.Update(ctx, resourceGroup, storageSyncServiceName, parameters); err != nil {
        return fmt.Errorf("Error updating Storage Sync Service (Storage Sync Service Name %q / Resource Group %q): %+v", storageSyncServiceName, resourceGroup, err)
    }

    return resourceArmStorageSyncServiceRead(d, meta)
}

func resourceArmStorageSyncServiceDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).storageSyncServicesClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    storageSyncServiceName := id.Path["storageSyncServices"]

    if _, err := client.Delete(ctx, resourceGroup, storageSyncServiceName); err != nil {
        return fmt.Errorf("Error deleting Storage Sync Service (Storage Sync Service Name %q / Resource Group %q): %+v", storageSyncServiceName, resourceGroup, err)
    }

    return nil
}
