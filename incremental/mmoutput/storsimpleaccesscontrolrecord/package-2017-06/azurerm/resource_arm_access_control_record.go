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



func resourceArmAccessControlRecord() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmAccessControlRecordCreateUpdate,
        Read: resourceArmAccessControlRecordRead,
        Update: resourceArmAccessControlRecordCreateUpdate,
        Delete: resourceArmAccessControlRecordDelete,

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

            "access_control_record_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "initiator_name": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validate.NoEmptyStrings,
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

            "type": {
                Type: schema.TypeString,
                Computed: true,
            },

            "volume_count": {
                Type: schema.TypeInt,
                Computed: true,
            },
        },
    }
}

func resourceArmAccessControlRecordCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).accessControlRecordsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    accessControlRecordName := d.Get("access_control_record_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, name, accessControlRecordName)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Access Control Record %q (Resource Group %q / Access Control Record Name %q): %+v", name, resourceGroup, accessControlRecordName, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_access_control_record", *existing.ID)
        }
    }

    initiatorName := d.Get("initiator_name").(string)
    kind := d.Get("kind").(string)

    parameters := storsimple.AccessControlRecord{
        Kind: storsimple.Kind(kind),
        AccessControlRecordProperties: &storsimple.AccessControlRecordProperties{
            InitiatorName: utils.String(initiatorName),
        },
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, name, accessControlRecordName, parameters)
    if err != nil {
        return fmt.Errorf("Error creating Access Control Record %q (Resource Group %q / Access Control Record Name %q): %+v", name, resourceGroup, accessControlRecordName, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Access Control Record %q (Resource Group %q / Access Control Record Name %q): %+v", name, resourceGroup, accessControlRecordName, err)
    }


    resp, err := client.Get(ctx, resourceGroup, name, accessControlRecordName)
    if err != nil {
        return fmt.Errorf("Error retrieving Access Control Record %q (Resource Group %q / Access Control Record Name %q): %+v", name, resourceGroup, accessControlRecordName, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Access Control Record %q (Resource Group %q / Access Control Record Name %q) ID", name, resourceGroup, accessControlRecordName)
    }
    d.SetId(*resp.ID)

    return resourceArmAccessControlRecordRead(d, meta)
}

func resourceArmAccessControlRecordRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).accessControlRecordsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["managers"]
    accessControlRecordName := id.Path["accessControlRecords"]

    resp, err := client.Get(ctx, resourceGroup, name, accessControlRecordName)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Access Control Record %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Access Control Record %q (Resource Group %q / Access Control Record Name %q): %+v", name, resourceGroup, accessControlRecordName, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    d.Set("access_control_record_name", accessControlRecordName)
    if accessControlRecordProperties := resp.AccessControlRecordProperties; accessControlRecordProperties != nil {
        d.Set("initiator_name", accessControlRecordProperties.InitiatorName)
        d.Set("volume_count", int(*accessControlRecordProperties.VolumeCount))
    }
    d.Set("kind", string(resp.Kind))
    d.Set("type", resp.Type)

    return nil
}


func resourceArmAccessControlRecordDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).accessControlRecordsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["managers"]
    accessControlRecordName := id.Path["accessControlRecords"]

    future, err := client.Delete(ctx, resourceGroup, name, accessControlRecordName)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Access Control Record %q (Resource Group %q / Access Control Record Name %q): %+v", name, resourceGroup, accessControlRecordName, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Access Control Record %q (Resource Group %q / Access Control Record Name %q): %+v", name, resourceGroup, accessControlRecordName, err)
        }
    }

    return nil
}