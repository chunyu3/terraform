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



func resourceArmStorageAccount() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmStorageAccountCreateUpdate,
        Read: resourceArmStorageAccountRead,
        Update: resourceArmStorageAccountCreateUpdate,
        Delete: resourceArmStorageAccountDelete,

        Importer: &schema.ResourceImporter{
            State: schema.ImportStatePassthrough,
        },


        Schema: map[string]*schema.Schema{
            "device_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "storage_account_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "data_policy": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(databoxedge.Cloud),
                    string(databoxedge.Local),
                }, false),
                Default: string(databoxedge.Cloud),
            },

            "description": {
                Type: schema.TypeString,
                Optional: true,
            },

            "storage_account_credential_id": {
                Type: schema.TypeString,
                Optional: true,
            },

            "storage_account_status": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(databoxedge.OK),
                    string(databoxedge.Offline),
                    string(databoxedge.Unknown),
                    string(databoxedge.Updating),
                    string(databoxedge.NeedsAttention),
                }, false),
                Default: string(databoxedge.OK),
            },

            "blob_endpoint": {
                Type: schema.TypeString,
                Computed: true,
            },

            "container_count": {
                Type: schema.TypeInt,
                Computed: true,
            },

            "id": {
                Type: schema.TypeString,
                Computed: true,
            },

            "name": {
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

func resourceArmStorageAccountCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).storageAccountsClient
    ctx, cancel := timeouts.ForCreateUpdate(meta.(*ArmClient).StopContext, d)
    defer cancel()

    resourceGroupName := d.Get("resource_group").(string)
    deviceName := d.Get("device_name").(string)
    name := d.Get("storage_account_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroupName, deviceName, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Storage Account (Storage Account Name %q / Resource Group %q / Device Name %q): %+v", name, resourceGroupName, deviceName, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_storage_account", *existing.ID)
        }
    }

    dataPolicy := d.Get("data_policy").(string)
    description := d.Get("description").(string)
    storageAccountCredentialID := d.Get("storage_account_credential_id").(string)
    storageAccountStatus := d.Get("storage_account_status").(string)

    storageAccount := databoxedge.StorageAccount{
        StorageAccountProperties: &databoxedge.StorageAccountProperties{
            DataPolicy: databoxedge.DataPolicy(dataPolicy),
            Description: utils.String(description),
            StorageAccountCredentialID: utils.String(storageAccountCredentialID),
            StorageAccountStatus: databoxedge.StorageAccountStatus(storageAccountStatus),
        },
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroupName, deviceName, name, storageAccount)
    if err != nil {
        return fmt.Errorf("Error creating Storage Account (Storage Account Name %q / Resource Group %q / Device Name %q): %+v", name, resourceGroupName, deviceName, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Storage Account (Storage Account Name %q / Resource Group %q / Device Name %q): %+v", name, resourceGroupName, deviceName, err)
    }


    resp, err := client.Get(ctx, resourceGroupName, deviceName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Storage Account (Storage Account Name %q / Resource Group %q / Device Name %q): %+v", name, resourceGroupName, deviceName, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Storage Account (Storage Account Name %q / Resource Group %q / Device Name %q) ID", name, resourceGroupName, deviceName)
    }
    d.SetId(*resp.ID)

    return resourceArmStorageAccountRead(d, meta)
}

func resourceArmStorageAccountRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).storageAccountsClient
    ctx, cancel := timeouts.ForRead(meta.(*ArmClient).StopContext, d)
    defer cancel()

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroupName := id.ResourceGroup
    deviceName := id.Path["dataBoxEdgeDevices"]
    name := id.Path["storageAccounts"]

    resp, err := client.Get(ctx, resourceGroupName, deviceName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Storage Account %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Storage Account (Storage Account Name %q / Resource Group %q / Device Name %q): %+v", name, resourceGroupName, deviceName, err)
    }


    d.Set("resource_group", resourceGroupName)
    if storageAccountProperties := resp.StorageAccountProperties; storageAccountProperties != nil {
        d.Set("blob_endpoint", storageAccountProperties.BlobEndpoint)
        d.Set("container_count", int(*storageAccountProperties.ContainerCount))
        d.Set("data_policy", string(storageAccountProperties.DataPolicy))
        d.Set("description", storageAccountProperties.Description)
        d.Set("storage_account_credential_id", storageAccountProperties.StorageAccountCredentialID)
        d.Set("storage_account_status", string(storageAccountProperties.StorageAccountStatus))
    }
    d.Set("device_name", deviceName)
    d.Set("id", resp.ID)
    d.Set("name", resp.Name)
    d.Set("storage_account_name", name)
    d.Set("type", resp.Type)

    return nil
}


func resourceArmStorageAccountDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).storageAccountsClient
    ctx, cancel := timeouts.ForDelete(meta.(*ArmClient).StopContext, d)
    defer cancel()


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroupName := id.ResourceGroup
    deviceName := id.Path["dataBoxEdgeDevices"]
    name := id.Path["storageAccounts"]

    future, err := client.Delete(ctx, resourceGroupName, deviceName, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Storage Account (Storage Account Name %q / Resource Group %q / Device Name %q): %+v", name, resourceGroupName, deviceName, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Storage Account (Storage Account Name %q / Resource Group %q / Device Name %q): %+v", name, resourceGroupName, deviceName, err)
        }
    }

    return nil
}