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



func resourceArmDisk() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmDiskCreateUpdate,
        Read: resourceArmDiskRead,
        Update: resourceArmDiskCreateUpdate,
        Delete: resourceArmDiskDelete,

        Importer: &schema.ResourceImporter{
            State: schema.ImportStatePassthrough,
        },


        Schema: map[string]*schema.Schema{
            "lab_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "user_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "disk_blob_name": {
                Type: schema.TypeString,
                Optional: true,
            },

            "disk_size_gi_b": {
                Type: schema.TypeInt,
                Optional: true,
            },

            "disk_type": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(devtestlab.Standard),
                    string(devtestlab.Premium),
                }, false),
                Default: string(devtestlab.Standard),
            },

            "disk_uri": {
                Type: schema.TypeString,
                Optional: true,
            },

            "host_caching": {
                Type: schema.TypeString,
                Optional: true,
            },

            "leased_by_lab_vmid": {
                Type: schema.TypeString,
                Optional: true,
            },

            "leased_by_lab_vmid": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
            },

            "location": azure.SchemaLocation(),

            "managed_disk_id": {
                Type: schema.TypeString,
                Optional: true,
            },

            "tags": tags.Schema(),

            "unique_identifier": {
                Type: schema.TypeString,
                Optional: true,
            },

            "created_date": {
                Type: schema.TypeString,
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

            "provisioning_state": {
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

func resourceArmDiskCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).disksClient
    ctx, cancel := timeouts.ForCreateUpdate(meta.(*ArmClient).StopContext, d)
    defer cancel()

    resourceGroupName := d.Get("resource_group").(string)
    labName := d.Get("lab_name").(string)
    name := d.Get("name").(string)
    name := d.Get("user_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroupName, labName, name, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Disk (Name %q / User Name %q / Lab Name %q / Resource Group %q): %+v", name, name, labName, resourceGroupName, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_disk", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    diskBlobName := d.Get("disk_blob_name").(string)
    diskSizeGiB := d.Get("disk_size_gi_b").(int)
    diskType := d.Get("disk_type").(string)
    diskURI := d.Get("disk_uri").(string)
    hostCaching := d.Get("host_caching").(string)
    leasedByLabVMID := d.Get("leased_by_lab_vmid").(string)
    leasedByLabVMID := d.Get("leased_by_lab_vmid").(string)
    managedDiskID := d.Get("managed_disk_id").(string)
    uniqueIdentifier := d.Get("unique_identifier").(string)
    tags := d.Get("tags").(map[string]interface{})

    disk := devtestlab.Disk{
        LeasedByLabVMID: utils.String(leasedByLabVMID),
        LeasedByLabVMID: utils.String(leasedByLabVMID),
        Location: utils.String(location),
        DiskProperties: &devtestlab.DiskProperties{
            DiskBlobName: utils.String(diskBlobName),
            DiskSizeGiB: utils.Int32(int32(diskSizeGiB)),
            DiskType: devtestlab.StorageType(diskType),
            DiskURI: utils.String(diskURI),
            HostCaching: utils.String(hostCaching),
            LeasedByLabVMID: utils.String(leasedByLabVMID),
            ManagedDiskID: utils.String(managedDiskID),
            UniqueIdentifier: utils.String(uniqueIdentifier),
        },
        Tags: tags.Expand(tags),
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroupName, labName, name, name, disk)
    if err != nil {
        return fmt.Errorf("Error creating Disk (Name %q / User Name %q / Lab Name %q / Resource Group %q): %+v", name, name, labName, resourceGroupName, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Disk (Name %q / User Name %q / Lab Name %q / Resource Group %q): %+v", name, name, labName, resourceGroupName, err)
    }


    resp, err := client.Get(ctx, resourceGroupName, labName, name, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Disk (Name %q / User Name %q / Lab Name %q / Resource Group %q): %+v", name, name, labName, resourceGroupName, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Disk (Name %q / User Name %q / Lab Name %q / Resource Group %q) ID", name, name, labName, resourceGroupName)
    }
    d.SetId(*resp.ID)

    return resourceArmDiskRead(d, meta)
}

func resourceArmDiskRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).disksClient
    ctx, cancel := timeouts.ForRead(meta.(*ArmClient).StopContext, d)
    defer cancel()

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroupName := id.ResourceGroup
    labName := id.Path["labs"]
    name := id.Path["users"]
    name := id.Path["disks"]

    resp, err := client.Get(ctx, resourceGroupName, labName, name, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Disk %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Disk (Name %q / User Name %q / Lab Name %q / Resource Group %q): %+v", name, name, labName, resourceGroupName, err)
    }


    d.Set("resource_group", resourceGroupName)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if diskProperties := resp.DiskProperties; diskProperties != nil {
        d.Set("created_date", (diskProperties.CreatedDate).String())
        d.Set("disk_blob_name", diskProperties.DiskBlobName)
        d.Set("disk_size_gi_b", int(*diskProperties.DiskSizeGiB))
        d.Set("disk_type", string(diskProperties.DiskType))
        d.Set("disk_uri", diskProperties.DiskURI)
        d.Set("host_caching", diskProperties.HostCaching)
        d.Set("leased_by_lab_vmid", diskProperties.LeasedByLabVMID)
        d.Set("managed_disk_id", diskProperties.ManagedDiskID)
        d.Set("provisioning_state", diskProperties.ProvisioningState)
        d.Set("unique_identifier", diskProperties.UniqueIdentifier)
    }
    d.Set("id", resp.ID)
    d.Set("lab_name", labName)
    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("type", resp.Type)
    d.Set("user_name", name)

    return tags.FlattenAndSet(d, resp.Tags)
}


func resourceArmDiskDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).disksClient
    ctx, cancel := timeouts.ForDelete(meta.(*ArmClient).StopContext, d)
    defer cancel()


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroupName := id.ResourceGroup
    labName := id.Path["labs"]
    name := id.Path["users"]
    name := id.Path["disks"]

    future, err := client.Delete(ctx, resourceGroupName, labName, name, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Disk (Name %q / User Name %q / Lab Name %q / Resource Group %q): %+v", name, name, labName, resourceGroupName, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Disk (Name %q / User Name %q / Lab Name %q / Resource Group %q): %+v", name, name, labName, resourceGroupName, err)
        }
    }

    return nil
}
