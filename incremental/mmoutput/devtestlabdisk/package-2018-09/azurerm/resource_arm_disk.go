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
        Create: resourceArmDiskCreate,
        Read: resourceArmDiskRead,
        Update: resourceArmDiskUpdate,
        Delete: resourceArmDiskDelete,

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
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "location": azure.SchemaLocation(),

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "lab_name": {
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
                    string(devtestlab.StandardSSD),
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

            "leased_by_lab_vm_id": {
                Type: schema.TypeString,
                Optional: true,
            },

            "managed_disk_id": {
                Type: schema.TypeString,
                Optional: true,
            },

            "created_date": {
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

            "unique_identifier": {
                Type: schema.TypeString,
                Computed: true,
            },

            "tags": tags.Schema(),
        },
    }
}

func resourceArmDiskCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).disksClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    labName := d.Get("lab_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, labName, name, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Disk %q (Lab Name %q / Resource Group %q): %+v", name, labName, resourceGroup, err)
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
    diskUri := d.Get("disk_uri").(string)
    hostCaching := d.Get("host_caching").(string)
    leasedByLabVmId := d.Get("leased_by_lab_vm_id").(string)
    managedDiskId := d.Get("managed_disk_id").(string)
    t := d.Get("tags").(map[string]interface{})

    disk := devtestlab.Disk{
        Location: utils.String(location),
        DiskProperties: &devtestlab.DiskProperties{
            DiskBlobName: utils.String(diskBlobName),
            DiskSizeGiB: utils.Int32(int32(diskSizeGiB)),
            DiskType: devtestlab.StorageType(diskType),
            DiskUri: utils.String(diskUri),
            HostCaching: utils.String(hostCaching),
            LeasedByLabVmID: utils.String(leasedByLabVmId),
            ManagedDiskID: utils.String(managedDiskId),
        },
        Tags: tags.Expand(t),
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, labName, name, name, disk)
    if err != nil {
        return fmt.Errorf("Error creating Disk %q (Lab Name %q / Resource Group %q): %+v", name, labName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Disk %q (Lab Name %q / Resource Group %q): %+v", name, labName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, labName, name, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Disk %q (Lab Name %q / Resource Group %q): %+v", name, labName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Disk %q (Lab Name %q / Resource Group %q) ID", name, labName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmDiskRead(d, meta)
}

func resourceArmDiskRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).disksClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    labName := id.Path["labs"]
    name := id.Path["users"]
    name := id.Path["disks"]

    resp, err := client.Get(ctx, resourceGroup, labName, name, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Disk %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Disk %q (Lab Name %q / Resource Group %q): %+v", name, labName, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if diskProperties := resp.DiskProperties; diskProperties != nil {
        d.Set("created_date", (diskProperties.CreatedDate).String())
        d.Set("disk_blob_name", diskProperties.DiskBlobName)
        d.Set("disk_size_gi_b", int(*diskProperties.DiskSizeGiB))
        d.Set("disk_type", string(diskProperties.DiskType))
        d.Set("disk_uri", diskProperties.DiskUri)
        d.Set("host_caching", diskProperties.HostCaching)
        d.Set("leased_by_lab_vm_id", diskProperties.LeasedByLabVmID)
        d.Set("managed_disk_id", diskProperties.ManagedDiskID)
        d.Set("provisioning_state", diskProperties.ProvisioningState)
        d.Set("unique_identifier", diskProperties.UniqueIdentifier)
    }
    d.Set("lab_name", labName)
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}

func resourceArmDiskUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).disksClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    diskBlobName := d.Get("disk_blob_name").(string)
    diskSizeGiB := d.Get("disk_size_gi_b").(int)
    diskType := d.Get("disk_type").(string)
    diskUri := d.Get("disk_uri").(string)
    hostCaching := d.Get("host_caching").(string)
    labName := d.Get("lab_name").(string)
    leasedByLabVmId := d.Get("leased_by_lab_vm_id").(string)
    managedDiskId := d.Get("managed_disk_id").(string)
    t := d.Get("tags").(map[string]interface{})

    disk := devtestlab.Disk{
        Location: utils.String(location),
        DiskProperties: &devtestlab.DiskProperties{
            DiskBlobName: utils.String(diskBlobName),
            DiskSizeGiB: utils.Int32(int32(diskSizeGiB)),
            DiskType: devtestlab.StorageType(diskType),
            DiskUri: utils.String(diskUri),
            HostCaching: utils.String(hostCaching),
            LeasedByLabVmID: utils.String(leasedByLabVmId),
            ManagedDiskID: utils.String(managedDiskId),
        },
        Tags: tags.Expand(t),
    }


    if _, err := client.Update(ctx, resourceGroup, labName, name, name, disk); err != nil {
        return fmt.Errorf("Error updating Disk %q (Lab Name %q / Resource Group %q): %+v", name, labName, resourceGroup, err)
    }

    return resourceArmDiskRead(d, meta)
}

func resourceArmDiskDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).disksClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    labName := id.Path["labs"]
    name := id.Path["users"]
    name := id.Path["disks"]

    future, err := client.Delete(ctx, resourceGroup, labName, name, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Disk %q (Lab Name %q / Resource Group %q): %+v", name, labName, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Disk %q (Lab Name %q / Resource Group %q): %+v", name, labName, resourceGroup, err)
        }
    }

    return nil
}
