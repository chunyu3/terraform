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



func resourceArmIscsiDisk() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmIscsiDiskCreateUpdate,
        Read: resourceArmIscsiDiskRead,
        Update: resourceArmIscsiDiskCreateUpdate,
        Delete: resourceArmIscsiDiskDelete,

        Importer: &schema.ResourceImporter{
            State: schema.ImportStatePassthrough,
        },


        Schema: map[string]*schema.Schema{
            "name": {
                Type: schema.TypeString,
                Computed: true,
            },

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "access_control_records": {
                Type: schema.TypeList,
                Required: true,
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
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

            "disk_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "disk_status": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(storsimple.Online),
                    string(storsimple.Offline),
                }, false),
            },

            "iscsi_server_name": {
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

            "description": {
                Type: schema.TypeString,
                Optional: true,
            },

            "local_used_capacity_in_bytes": {
                Type: schema.TypeInt,
                Computed: true,
            },

            "type": {
                Type: schema.TypeString,
                Computed: true,
            },

            "used_capacity_in_bytes": {
                Type: schema.TypeInt,
                Computed: true,
            },
        },
    }
}

func resourceArmIscsiDiskCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).iscsiDisksClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    deviceName := d.Get("device_name").(string)
    diskName := d.Get("disk_name").(string)
    iscsiServerName := d.Get("iscsi_server_name").(string)
    managerName := d.Get("manager_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, deviceName, iscsiServerName, diskName, resourceGroup, managerName)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Iscsi Disk (Manager Name %q / Resource Group %q / Disk Name %q / Iscsi Server Name %q / Device Name %q): %+v", managerName, resourceGroup, diskName, iscsiServerName, deviceName, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_iscsi_disk", *existing.ID)
        }
    }

    accessControlRecords := d.Get("access_control_records").([]interface{})
    dataPolicy := d.Get("data_policy").(string)
    description := d.Get("description").(string)
    diskStatus := d.Get("disk_status").(string)
    monitoringStatus := d.Get("monitoring_status").(string)
    provisionedCapacityInBytes := d.Get("provisioned_capacity_in_bytes").(int)

    iscsiDisk := storsimple.ISCSIDisk{
        ISCSIDiskProperties: &storsimple.ISCSIDiskProperties{
            AccessControlRecords: utils.ExpandStringSlice(accessControlRecords),
            DataPolicy: storsimple.DataPolicy(dataPolicy),
            Description: utils.String(description),
            DiskStatus: storsimple.DiskStatus(diskStatus),
            MonitoringStatus: storsimple.MonitoringStatus(monitoringStatus),
            ProvisionedCapacityInBytes: utils.Int64(int64(provisionedCapacityInBytes)),
        },
    }


    future, err := client.CreateOrUpdate(ctx, deviceName, iscsiServerName, diskName, resourceGroup, managerName, iscsiDisk)
    if err != nil {
        return fmt.Errorf("Error creating Iscsi Disk (Manager Name %q / Resource Group %q / Disk Name %q / Iscsi Server Name %q / Device Name %q): %+v", managerName, resourceGroup, diskName, iscsiServerName, deviceName, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Iscsi Disk (Manager Name %q / Resource Group %q / Disk Name %q / Iscsi Server Name %q / Device Name %q): %+v", managerName, resourceGroup, diskName, iscsiServerName, deviceName, err)
    }


    resp, err := client.Get(ctx, deviceName, iscsiServerName, diskName, resourceGroup, managerName)
    if err != nil {
        return fmt.Errorf("Error retrieving Iscsi Disk (Manager Name %q / Resource Group %q / Disk Name %q / Iscsi Server Name %q / Device Name %q): %+v", managerName, resourceGroup, diskName, iscsiServerName, deviceName, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Iscsi Disk (Manager Name %q / Resource Group %q / Disk Name %q / Iscsi Server Name %q / Device Name %q) ID", managerName, resourceGroup, diskName, iscsiServerName, deviceName)
    }
    d.SetId(*resp.ID)

    return resourceArmIscsiDiskRead(d, meta)
}

func resourceArmIscsiDiskRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).iscsiDisksClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    deviceName := id.Path["devices"]
    iscsiServerName := id.Path["iscsiservers"]
    diskName := id.Path["disks"]
    resourceGroup := id.ResourceGroup
    managerName := id.Path["managers"]

    resp, err := client.Get(ctx, deviceName, iscsiServerName, diskName, resourceGroup, managerName)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Iscsi Disk %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Iscsi Disk (Manager Name %q / Resource Group %q / Disk Name %q / Iscsi Server Name %q / Device Name %q): %+v", managerName, resourceGroup, diskName, iscsiServerName, deviceName, err)
    }


    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if iSCSIDiskProperties := resp.ISCSIDiskProperties; iSCSIDiskProperties != nil {
        d.Set("access_control_records", utils.FlattenStringSlice(iSCSIDiskProperties.AccessControlRecords))
        d.Set("data_policy", string(iSCSIDiskProperties.DataPolicy))
        d.Set("description", iSCSIDiskProperties.Description)
        d.Set("disk_status", string(iSCSIDiskProperties.DiskStatus))
        d.Set("local_used_capacity_in_bytes", int(*iSCSIDiskProperties.LocalUsedCapacityInBytes))
        d.Set("monitoring_status", string(iSCSIDiskProperties.MonitoringStatus))
        d.Set("provisioned_capacity_in_bytes", int(*iSCSIDiskProperties.ProvisionedCapacityInBytes))
        d.Set("used_capacity_in_bytes", int(*iSCSIDiskProperties.UsedCapacityInBytes))
    }
    d.Set("device_name", deviceName)
    d.Set("disk_name", diskName)
    d.Set("iscsi_server_name", iscsiServerName)
    d.Set("manager_name", managerName)
    d.Set("type", resp.Type)

    return nil
}


func resourceArmIscsiDiskDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).iscsiDisksClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    deviceName := id.Path["devices"]
    iscsiServerName := id.Path["iscsiservers"]
    diskName := id.Path["disks"]
    resourceGroup := id.ResourceGroup
    managerName := id.Path["managers"]

    future, err := client.Delete(ctx, deviceName, iscsiServerName, diskName, resourceGroup, managerName)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Iscsi Disk (Manager Name %q / Resource Group %q / Disk Name %q / Iscsi Server Name %q / Device Name %q): %+v", managerName, resourceGroup, diskName, iscsiServerName, deviceName, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Iscsi Disk (Manager Name %q / Resource Group %q / Disk Name %q / Iscsi Server Name %q / Device Name %q): %+v", managerName, resourceGroup, diskName, iscsiServerName, deviceName, err)
        }
    }

    return nil
}
