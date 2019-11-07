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



func resourceArmVolume() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmVolumeCreateUpdate,
        Read: resourceArmVolumeRead,
        Update: resourceArmVolumeCreateUpdate,
        Delete: resourceArmVolumeDelete,

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

            "access_control_record_ids": {
                Type: schema.TypeList,
                Required: true,
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
            },

            "device_name": {
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

            "size_in_bytes": {
                Type: schema.TypeInt,
                Required: true,
            },

            "volume_container_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "volume_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "volume_status": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(storsimple.Online),
                    string(storsimple.Offline),
                }, false),
            },

            "volume_type": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(storsimple.Tiered),
                    string(storsimple.Archival),
                    string(storsimple.LocallyPinned),
                }, false),
            },

            "kind": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(storsimple.Series8000),
                }, false),
                Default: string(storsimple.Series8000),
            },

            "backup_policy_ids": {
                Type: schema.TypeList,
                Computed: true,
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
            },

            "backup_status": {
                Type: schema.TypeString,
                Computed: true,
            },

            "operation_status": {
                Type: schema.TypeString,
                Computed: true,
            },

            "type": {
                Type: schema.TypeString,
                Computed: true,
            },

            "volume_container_id": {
                Type: schema.TypeString,
                Computed: true,
            },
        },
    }
}

func resourceArmVolumeCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).volumesClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    deviceName := d.Get("device_name").(string)
    volumeContainerName := d.Get("volume_container_name").(string)
    volumeName := d.Get("volume_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, name, deviceName, volumeContainerName, volumeName)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Volume %q (Resource Group %q / Volume Name %q / Volume Container Name %q / Device Name %q): %+v", name, resourceGroup, volumeName, volumeContainerName, deviceName, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_volume", *existing.ID)
        }
    }

    accessControlRecordIds := d.Get("access_control_record_ids").([]interface{})
    kind := d.Get("kind").(string)
    monitoringStatus := d.Get("monitoring_status").(string)
    sizeInBytes := d.Get("size_in_bytes").(int)
    volumeStatus := d.Get("volume_status").(string)
    volumeType := d.Get("volume_type").(string)

    parameters := storsimple.Volume{
        Kind: storsimple.Kind(kind),
        VolumeProperties: &storsimple.VolumeProperties{
            AccessControlRecordIds: utils.ExpandStringSlice(accessControlRecordIds),
            MonitoringStatus: storsimple.MonitoringStatus(monitoringStatus),
            SizeInBytes: utils.Int64(int64(sizeInBytes)),
            VolumeStatus: storsimple.VolumeStatus(volumeStatus),
            VolumeType: storsimple.VolumeType(volumeType),
        },
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, name, deviceName, volumeContainerName, volumeName, parameters)
    if err != nil {
        return fmt.Errorf("Error creating Volume %q (Resource Group %q / Volume Name %q / Volume Container Name %q / Device Name %q): %+v", name, resourceGroup, volumeName, volumeContainerName, deviceName, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Volume %q (Resource Group %q / Volume Name %q / Volume Container Name %q / Device Name %q): %+v", name, resourceGroup, volumeName, volumeContainerName, deviceName, err)
    }


    resp, err := client.Get(ctx, resourceGroup, name, deviceName, volumeContainerName, volumeName)
    if err != nil {
        return fmt.Errorf("Error retrieving Volume %q (Resource Group %q / Volume Name %q / Volume Container Name %q / Device Name %q): %+v", name, resourceGroup, volumeName, volumeContainerName, deviceName, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Volume %q (Resource Group %q / Volume Name %q / Volume Container Name %q / Device Name %q) ID", name, resourceGroup, volumeName, volumeContainerName, deviceName)
    }
    d.SetId(*resp.ID)

    return resourceArmVolumeRead(d, meta)
}

func resourceArmVolumeRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).volumesClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["managers"]
    deviceName := id.Path["devices"]
    volumeContainerName := id.Path["volumeContainers"]
    volumeName := id.Path["volumes"]

    resp, err := client.Get(ctx, resourceGroup, name, deviceName, volumeContainerName, volumeName)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Volume %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Volume %q (Resource Group %q / Volume Name %q / Volume Container Name %q / Device Name %q): %+v", name, resourceGroup, volumeName, volumeContainerName, deviceName, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if volumeProperties := resp.VolumeProperties; volumeProperties != nil {
        d.Set("access_control_record_ids", utils.FlattenStringSlice(volumeProperties.AccessControlRecordIds))
        d.Set("backup_policy_ids", utils.FlattenStringSlice(volumeProperties.BackupPolicyIds))
        d.Set("backup_status", string(volumeProperties.BackupStatus))
        d.Set("monitoring_status", string(volumeProperties.MonitoringStatus))
        d.Set("operation_status", string(volumeProperties.OperationStatus))
        d.Set("size_in_bytes", int(*volumeProperties.SizeInBytes))
        d.Set("volume_container_id", volumeProperties.VolumeContainerID)
        d.Set("volume_status", string(volumeProperties.VolumeStatus))
        d.Set("volume_type", string(volumeProperties.VolumeType))
    }
    d.Set("device_name", deviceName)
    d.Set("kind", string(resp.Kind))
    d.Set("type", resp.Type)
    d.Set("volume_container_name", volumeContainerName)
    d.Set("volume_name", volumeName)

    return nil
}


func resourceArmVolumeDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).volumesClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["managers"]
    deviceName := id.Path["devices"]
    volumeContainerName := id.Path["volumeContainers"]
    volumeName := id.Path["volumes"]

    future, err := client.Delete(ctx, resourceGroup, name, deviceName, volumeContainerName, volumeName)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Volume %q (Resource Group %q / Volume Name %q / Volume Container Name %q / Device Name %q): %+v", name, resourceGroup, volumeName, volumeContainerName, deviceName, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Volume %q (Resource Group %q / Volume Name %q / Volume Container Name %q / Device Name %q): %+v", name, resourceGroup, volumeName, volumeContainerName, deviceName, err)
        }
    }

    return nil
}