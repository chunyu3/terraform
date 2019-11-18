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



func resourceArmLab() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmLabCreate,
        Read: resourceArmLabRead,
        Update: resourceArmLabUpdate,
        Delete: resourceArmLabDelete,

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

            "location": azure.SchemaLocation(),

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "lab_storage_type": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(devtestlab.Standard),
                    string(devtestlab.Premium),
                }, false),
                Default: string(devtestlab.Standard),
            },

            "premium_data_disks": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(devtestlab.Disabled),
                    string(devtestlab.Enabled),
                }, false),
                Default: string(devtestlab.Disabled),
            },

            "unique_identifier": {
                Type: schema.TypeString,
                Optional: true,
            },

            "artifacts_storage_account": {
                Type: schema.TypeString,
                Computed: true,
            },

            "created_date": {
                Type: schema.TypeString,
                Computed: true,
            },

            "default_premium_storage_account": {
                Type: schema.TypeString,
                Computed: true,
            },

            "default_storage_account": {
                Type: schema.TypeString,
                Computed: true,
            },

            "premium_data_disk_storage_account": {
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

            "vault_name": {
                Type: schema.TypeString,
                Computed: true,
            },

            "tags": tags.Schema(),
        },
    }
}

func resourceArmLabCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).labsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Lab %q (Resource Group %q): %+v", name, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_lab", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    labStorageType := d.Get("lab_storage_type").(string)
    premiumDataDisks := d.Get("premium_data_disks").(string)
    uniqueIdentifier := d.Get("unique_identifier").(string)
    t := d.Get("tags").(map[string]interface{})

    lab := devtestlab.LabFragment{
        Location: utils.String(location),
        LabPropertiesFragment: &devtestlab.LabPropertiesFragment{
            LabStorageType: devtestlab.StorageType(labStorageType),
            PremiumDataDisks: devtestlab.PremiumDataDisk(premiumDataDisks),
            UniqueIdentifier: utils.String(uniqueIdentifier),
        },
        Tags: tags.Expand(t),
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, name, lab)
    if err != nil {
        return fmt.Errorf("Error creating Lab %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Lab %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Lab %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Lab %q (Resource Group %q) ID", name, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmLabRead(d, meta)
}

func resourceArmLabRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).labsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["labs"]

    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Lab %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Lab %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if labPropertiesFragment := resp.LabPropertiesFragment; labPropertiesFragment != nil {
        d.Set("artifacts_storage_account", labPropertiesFragment.ArtifactsStorageAccount)
        d.Set("created_date", (labPropertiesFragment.CreatedDate).String())
        d.Set("default_premium_storage_account", labPropertiesFragment.DefaultPremiumStorageAccount)
        d.Set("default_storage_account", labPropertiesFragment.DefaultStorageAccount)
        d.Set("lab_storage_type", string(labPropertiesFragment.LabStorageType))
        d.Set("premium_data_disk_storage_account", labPropertiesFragment.PremiumDataDiskStorageAccount)
        d.Set("premium_data_disks", string(labPropertiesFragment.PremiumDataDisks))
        d.Set("provisioning_state", labPropertiesFragment.ProvisioningState)
        d.Set("unique_identifier", labPropertiesFragment.UniqueIdentifier)
        d.Set("vault_name", labPropertiesFragment.VaultName)
    }
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}

func resourceArmLabUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).labsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    labStorageType := d.Get("lab_storage_type").(string)
    premiumDataDisks := d.Get("premium_data_disks").(string)
    uniqueIdentifier := d.Get("unique_identifier").(string)
    t := d.Get("tags").(map[string]interface{})

    lab := devtestlab.LabFragment{
        Location: utils.String(location),
        LabPropertiesFragment: &devtestlab.LabPropertiesFragment{
            LabStorageType: devtestlab.StorageType(labStorageType),
            PremiumDataDisks: devtestlab.PremiumDataDisk(premiumDataDisks),
            UniqueIdentifier: utils.String(uniqueIdentifier),
        },
        Tags: tags.Expand(t),
    }


    if _, err := client.Update(ctx, resourceGroup, name, lab); err != nil {
        return fmt.Errorf("Error updating Lab %q (Resource Group %q): %+v", name, resourceGroup, err)
    }

    return resourceArmLabRead(d, meta)
}

func resourceArmLabDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).labsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["labs"]

    future, err := client.Delete(ctx, resourceGroup, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Lab %q (Resource Group %q): %+v", name, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Lab %q (Resource Group %q): %+v", name, resourceGroup, err)
        }
    }

    return nil
}
