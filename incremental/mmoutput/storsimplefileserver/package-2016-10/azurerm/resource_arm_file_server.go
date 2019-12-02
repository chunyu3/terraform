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



func resourceArmFileServer() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmFileServerCreateUpdate,
        Read: resourceArmFileServerRead,
        Update: resourceArmFileServerCreateUpdate,
        Delete: resourceArmFileServerDelete,

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

            "backup_schedule_group_id": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "device_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "domain_name": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "manager_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "storage_domain_id": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validate.NoEmptyStrings,
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

func resourceArmFileServerCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).fileServersClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    deviceName := d.Get("device_name").(string)
    managerName := d.Get("manager_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, managerName, deviceName, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing File Server %q (Manager Name %q / Resource Group %q / Device Name %q): %+v", name, managerName, resourceGroup, deviceName, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_file_server", *existing.ID)
        }
    }

    backupScheduleGroupId := d.Get("backup_schedule_group_id").(string)
    description := d.Get("description").(string)
    domainName := d.Get("domain_name").(string)
    storageDomainId := d.Get("storage_domain_id").(string)

    fileServer := storsimple.FileServer{
        FileServerProperties: &storsimple.FileServerProperties{
            BackupScheduleGroupID: utils.String(backupScheduleGroupId),
            Description: utils.String(description),
            DomainName: utils.String(domainName),
            StorageDomainID: utils.String(storageDomainId),
        },
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, managerName, deviceName, name, fileServer)
    if err != nil {
        return fmt.Errorf("Error creating File Server %q (Manager Name %q / Resource Group %q / Device Name %q): %+v", name, managerName, resourceGroup, deviceName, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of File Server %q (Manager Name %q / Resource Group %q / Device Name %q): %+v", name, managerName, resourceGroup, deviceName, err)
    }


    resp, err := client.Get(ctx, resourceGroup, managerName, deviceName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving File Server %q (Manager Name %q / Resource Group %q / Device Name %q): %+v", name, managerName, resourceGroup, deviceName, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read File Server %q (Manager Name %q / Resource Group %q / Device Name %q) ID", name, managerName, resourceGroup, deviceName)
    }
    d.SetId(*resp.ID)

    return resourceArmFileServerRead(d, meta)
}

func resourceArmFileServerRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).fileServersClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    managerName := id.Path["managers"]
    deviceName := id.Path["devices"]
    name := id.Path["fileservers"]

    resp, err := client.Get(ctx, resourceGroup, managerName, deviceName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] File Server %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading File Server %q (Manager Name %q / Resource Group %q / Device Name %q): %+v", name, managerName, resourceGroup, deviceName, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if fileServerProperties := resp.FileServerProperties; fileServerProperties != nil {
        d.Set("backup_schedule_group_id", fileServerProperties.BackupScheduleGroupID)
        d.Set("description", fileServerProperties.Description)
        d.Set("domain_name", fileServerProperties.DomainName)
        d.Set("storage_domain_id", fileServerProperties.StorageDomainID)
    }
    d.Set("device_name", deviceName)
    d.Set("manager_name", managerName)
    d.Set("type", resp.Type)

    return nil
}


func resourceArmFileServerDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).fileServersClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    managerName := id.Path["managers"]
    deviceName := id.Path["devices"]
    name := id.Path["fileservers"]

    future, err := client.Delete(ctx, resourceGroup, managerName, deviceName, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting File Server %q (Manager Name %q / Resource Group %q / Device Name %q): %+v", name, managerName, resourceGroup, deviceName, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting File Server %q (Manager Name %q / Resource Group %q / Device Name %q): %+v", name, managerName, resourceGroup, deviceName, err)
        }
    }

    return nil
}
