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



func resourceArmIscsiServer() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmIscsiServerCreateUpdate,
        Read: resourceArmIscsiServerRead,
        Update: resourceArmIscsiServerCreateUpdate,
        Delete: resourceArmIscsiServerDelete,

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

            "iscsi_server_name": {
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

            "chap_id": {
                Type: schema.TypeString,
                Optional: true,
            },

            "description": {
                Type: schema.TypeString,
                Optional: true,
            },

            "reverse_chap_id": {
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

func resourceArmIscsiServerCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).iscsiServersClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    deviceName := d.Get("device_name").(string)
    iscsiServerName := d.Get("iscsi_server_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, name, deviceName, iscsiServerName)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Iscsi Server %q (Resource Group %q / Iscsi Server Name %q / Device Name %q): %+v", name, resourceGroup, iscsiServerName, deviceName, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_iscsi_server", *existing.ID)
        }
    }

    backupScheduleGroupId := d.Get("backup_schedule_group_id").(string)
    chapId := d.Get("chap_id").(string)
    description := d.Get("description").(string)
    reverseChapId := d.Get("reverse_chap_id").(string)
    storageDomainId := d.Get("storage_domain_id").(string)

    iscsiServer := storsimple.ISCSIServer{
        ISCSIServerProperties: &storsimple.ISCSIServerProperties{
            BackupScheduleGroupID: utils.String(backupScheduleGroupId),
            ChapID: utils.String(chapId),
            Description: utils.String(description),
            ReverseChapID: utils.String(reverseChapId),
            StorageDomainID: utils.String(storageDomainId),
        },
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, name, deviceName, iscsiServerName, iscsiServer)
    if err != nil {
        return fmt.Errorf("Error creating Iscsi Server %q (Resource Group %q / Iscsi Server Name %q / Device Name %q): %+v", name, resourceGroup, iscsiServerName, deviceName, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Iscsi Server %q (Resource Group %q / Iscsi Server Name %q / Device Name %q): %+v", name, resourceGroup, iscsiServerName, deviceName, err)
    }


    resp, err := client.Get(ctx, resourceGroup, name, deviceName, iscsiServerName)
    if err != nil {
        return fmt.Errorf("Error retrieving Iscsi Server %q (Resource Group %q / Iscsi Server Name %q / Device Name %q): %+v", name, resourceGroup, iscsiServerName, deviceName, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Iscsi Server %q (Resource Group %q / Iscsi Server Name %q / Device Name %q) ID", name, resourceGroup, iscsiServerName, deviceName)
    }
    d.SetId(*resp.ID)

    return resourceArmIscsiServerRead(d, meta)
}

func resourceArmIscsiServerRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).iscsiServersClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["managers"]
    deviceName := id.Path["devices"]
    iscsiServerName := id.Path["iscsiservers"]

    resp, err := client.Get(ctx, resourceGroup, name, deviceName, iscsiServerName)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Iscsi Server %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Iscsi Server %q (Resource Group %q / Iscsi Server Name %q / Device Name %q): %+v", name, resourceGroup, iscsiServerName, deviceName, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if iSCSIServerProperties := resp.ISCSIServerProperties; iSCSIServerProperties != nil {
        d.Set("backup_schedule_group_id", iSCSIServerProperties.BackupScheduleGroupID)
        d.Set("chap_id", iSCSIServerProperties.ChapID)
        d.Set("description", iSCSIServerProperties.Description)
        d.Set("reverse_chap_id", iSCSIServerProperties.ReverseChapID)
        d.Set("storage_domain_id", iSCSIServerProperties.StorageDomainID)
    }
    d.Set("device_name", deviceName)
    d.Set("iscsi_server_name", iscsiServerName)
    d.Set("type", resp.Type)

    return nil
}


func resourceArmIscsiServerDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).iscsiServersClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["managers"]
    deviceName := id.Path["devices"]
    iscsiServerName := id.Path["iscsiservers"]

    future, err := client.Delete(ctx, resourceGroup, name, deviceName, iscsiServerName)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Iscsi Server %q (Resource Group %q / Iscsi Server Name %q / Device Name %q): %+v", name, resourceGroup, iscsiServerName, deviceName, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Iscsi Server %q (Resource Group %q / Iscsi Server Name %q / Device Name %q): %+v", name, resourceGroup, iscsiServerName, deviceName, err)
        }
    }

    return nil
}