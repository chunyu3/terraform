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



func resourceArmBackupScheduleGroup() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmBackupScheduleGroupCreateUpdate,
        Read: resourceArmBackupScheduleGroupRead,
        Update: resourceArmBackupScheduleGroupCreateUpdate,
        Delete: resourceArmBackupScheduleGroupDelete,

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

            "device_name": {
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

            "start_time": {
                Type: schema.TypeList,
                Required: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "hour": {
                            Type: schema.TypeInt,
                            Required: true,
                        },
                        "minute": {
                            Type: schema.TypeInt,
                            Required: true,
                        },
                    },
                },
            },

            "type": {
                Type: schema.TypeString,
                Computed: true,
            },
        },
    }
}

func resourceArmBackupScheduleGroupCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).backupScheduleGroupsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    deviceName := d.Get("device_name").(string)
    managerName := d.Get("manager_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, managerName, deviceName, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Backup Schedule Group %q (Manager Name %q / Resource Group %q / Device Name %q): %+v", name, managerName, resourceGroup, deviceName, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_backup_schedule_group", *existing.ID)
        }
    }

    startTime := d.Get("start_time").([]interface{})

    scheduleGroup := storsimple.BackupScheduleGroup{
        BackupScheduleGroupProperties: &storsimple.BackupScheduleGroupProperties{
            StartTime: expandArmBackupScheduleGroupTime(startTime),
        },
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, managerName, deviceName, name, scheduleGroup)
    if err != nil {
        return fmt.Errorf("Error creating Backup Schedule Group %q (Manager Name %q / Resource Group %q / Device Name %q): %+v", name, managerName, resourceGroup, deviceName, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Backup Schedule Group %q (Manager Name %q / Resource Group %q / Device Name %q): %+v", name, managerName, resourceGroup, deviceName, err)
    }


    resp, err := client.Get(ctx, resourceGroup, managerName, deviceName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Backup Schedule Group %q (Manager Name %q / Resource Group %q / Device Name %q): %+v", name, managerName, resourceGroup, deviceName, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Backup Schedule Group %q (Manager Name %q / Resource Group %q / Device Name %q) ID", name, managerName, resourceGroup, deviceName)
    }
    d.SetId(*resp.ID)

    return resourceArmBackupScheduleGroupRead(d, meta)
}

func resourceArmBackupScheduleGroupRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).backupScheduleGroupsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    managerName := id.Path["managers"]
    deviceName := id.Path["devices"]
    name := id.Path["backupScheduleGroups"]

    resp, err := client.Get(ctx, resourceGroup, managerName, deviceName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Backup Schedule Group %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Backup Schedule Group %q (Manager Name %q / Resource Group %q / Device Name %q): %+v", name, managerName, resourceGroup, deviceName, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    d.Set("device_name", deviceName)
    d.Set("manager_name", managerName)
    d.Set("type", resp.Type)

    return nil
}


func resourceArmBackupScheduleGroupDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).backupScheduleGroupsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    managerName := id.Path["managers"]
    deviceName := id.Path["devices"]
    name := id.Path["backupScheduleGroups"]

    future, err := client.Delete(ctx, resourceGroup, managerName, deviceName, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Backup Schedule Group %q (Manager Name %q / Resource Group %q / Device Name %q): %+v", name, managerName, resourceGroup, deviceName, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Backup Schedule Group %q (Manager Name %q / Resource Group %q / Device Name %q): %+v", name, managerName, resourceGroup, deviceName, err)
        }
    }

    return nil
}

func expandArmBackupScheduleGroupTime(input []interface{}) *storsimple.Time {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    hour := v["hour"].(int)
    minute := v["minute"].(int)

    result := storsimple.Time{
        Hour: utils.Int32(int32(hour)),
        Minute: utils.Int32(int32(minute)),
    }
    return &result
}
