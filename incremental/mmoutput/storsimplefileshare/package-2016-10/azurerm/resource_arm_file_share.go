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



func resourceArmFileShare() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmFileShareCreateUpdate,
        Read: resourceArmFileShareRead,
        Update: resourceArmFileShareCreateUpdate,
        Delete: resourceArmFileShareDelete,

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

            "admin_user": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "data_policy": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(storsimple.Invalid),
                    string(storsimple.Local),
                    string(storsimple.Tiered),
                    string(storsimple.Cloud),
                }, false),
            },

            "device_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "file_server_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "manager_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "monitoring_status": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(storsimple.Enabled),
                    string(storsimple.Disabled),
                }, false),
            },

            "provisioned_capacity_in_bytes": {
                Type: schema.TypeInt,
                Required: true,
            },

            "share_status": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(storsimple.Online),
                    string(storsimple.Offline),
                }, false),
            },

            "description": {
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

func resourceArmFileShareCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).fileSharesClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    deviceName := d.Get("device_name").(string)
    fileServerName := d.Get("file_server_name").(string)
    managerName := d.Get("manager_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, managerName, deviceName, fileServerName, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing File Share %q (Manager Name %q / Resource Group %q / File Server Name %q / Device Name %q): %+v", name, managerName, resourceGroup, fileServerName, deviceName, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_file_share", *existing.ID)
        }
    }

    adminUser := d.Get("admin_user").(string)
    dataPolicy := d.Get("data_policy").(string)
    description := d.Get("description").(string)
    monitoringStatus := d.Get("monitoring_status").(string)
    provisionedCapacityInBytes := d.Get("provisioned_capacity_in_bytes").(int)
    shareStatus := d.Get("share_status").(string)

    fileShare := storsimple.FileShare{
        FileShareProperties: &storsimple.FileShareProperties{
            AdminUser: utils.String(adminUser),
            DataPolicy: storsimple.DataPolicy(dataPolicy),
            Description: utils.String(description),
            MonitoringStatus: storsimple.MonitoringStatus(monitoringStatus),
            ProvisionedCapacityInBytes: utils.Int64(int64(provisionedCapacityInBytes)),
            ShareStatus: storsimple.ShareStatus(shareStatus),
        },
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, managerName, deviceName, fileServerName, name, fileShare)
    if err != nil {
        return fmt.Errorf("Error creating File Share %q (Manager Name %q / Resource Group %q / File Server Name %q / Device Name %q): %+v", name, managerName, resourceGroup, fileServerName, deviceName, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of File Share %q (Manager Name %q / Resource Group %q / File Server Name %q / Device Name %q): %+v", name, managerName, resourceGroup, fileServerName, deviceName, err)
    }


    resp, err := client.Get(ctx, resourceGroup, managerName, deviceName, fileServerName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving File Share %q (Manager Name %q / Resource Group %q / File Server Name %q / Device Name %q): %+v", name, managerName, resourceGroup, fileServerName, deviceName, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read File Share %q (Manager Name %q / Resource Group %q / File Server Name %q / Device Name %q) ID", name, managerName, resourceGroup, fileServerName, deviceName)
    }
    d.SetId(*resp.ID)

    return resourceArmFileShareRead(d, meta)
}

func resourceArmFileShareRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).fileSharesClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    managerName := id.Path["managers"]
    deviceName := id.Path["devices"]
    fileServerName := id.Path["fileservers"]
    name := id.Path["shares"]

    resp, err := client.Get(ctx, resourceGroup, managerName, deviceName, fileServerName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] File Share %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading File Share %q (Manager Name %q / Resource Group %q / File Server Name %q / Device Name %q): %+v", name, managerName, resourceGroup, fileServerName, deviceName, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    d.Set("device_name", deviceName)
    d.Set("file_server_name", fileServerName)
    d.Set("manager_name", managerName)
    d.Set("type", resp.Type)

    return nil
}


func resourceArmFileShareDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).fileSharesClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    managerName := id.Path["managers"]
    deviceName := id.Path["devices"]
    fileServerName := id.Path["fileservers"]
    name := id.Path["shares"]

    future, err := client.Delete(ctx, resourceGroup, managerName, deviceName, fileServerName, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting File Share %q (Manager Name %q / Resource Group %q / File Server Name %q / Device Name %q): %+v", name, managerName, resourceGroup, fileServerName, deviceName, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting File Share %q (Manager Name %q / Resource Group %q / File Server Name %q / Device Name %q): %+v", name, managerName, resourceGroup, fileServerName, deviceName, err)
        }
    }

    return nil
}