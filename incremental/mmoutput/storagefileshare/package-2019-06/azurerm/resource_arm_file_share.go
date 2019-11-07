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
        Create: resourceArmFileShareCreate,
        Read: resourceArmFileShareRead,
        Update: resourceArmFileShareUpdate,
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

            "account_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "metadata": {
                Type: schema.TypeMap,
                Optional: true,
                Elem: &schema.Schema{Type: schema.TypeString},
            },

            "share_quota": {
                Type: schema.TypeInt,
                Optional: true,
            },

            "etag": {
                Type: schema.TypeString,
                Computed: true,
            },

            "last_modified_time": {
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

func resourceArmFileShareCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).fileSharesClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    accountName := d.Get("account_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, accountName, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing File Share %q (Account Name %q / Resource Group %q): %+v", name, accountName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_file_share", *existing.ID)
        }
    }

    metadata := d.Get("metadata").(map[string]interface{})
    shareQuota := d.Get("share_quota").(int)

    fileShare := storage.FileShare{
        FileShareProperties: &storage.FileShareProperties{
            Metadata: utils.ExpandKeyValuePairs(metadata),
            ShareQuota: utils.Int(shareQuota),
        },
    }


    if _, err := client.Create(ctx, resourceGroup, accountName, name, fileShare); err != nil {
        return fmt.Errorf("Error creating File Share %q (Account Name %q / Resource Group %q): %+v", name, accountName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, accountName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving File Share %q (Account Name %q / Resource Group %q): %+v", name, accountName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read File Share %q (Account Name %q / Resource Group %q) ID", name, accountName, resourceGroup)
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
    accountName := id.Path["storageAccounts"]
    name := id.Path["shares"]

    resp, err := client.Get(ctx, resourceGroup, accountName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] File Share %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading File Share %q (Account Name %q / Resource Group %q): %+v", name, accountName, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    d.Set("account_name", accountName)
    d.Set("etag", resp.Etag)
    if fileShareProperties := resp.FileShareProperties; fileShareProperties != nil {
        d.Set("last_modified_time", (fileShareProperties.LastModifiedTime).String())
        d.Set("metadata", utils.FlattenKeyValuePairs(fileShareProperties.Metadata))
        d.Set("share_quota", fileShareProperties.ShareQuota)
    }
    d.Set("type", resp.Type)

    return nil
}

func resourceArmFileShareUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).fileSharesClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    accountName := d.Get("account_name").(string)
    metadata := d.Get("metadata").(map[string]interface{})
    shareQuota := d.Get("share_quota").(int)

    fileShare := storage.FileShare{
        FileShareProperties: &storage.FileShareProperties{
            Metadata: utils.ExpandKeyValuePairs(metadata),
            ShareQuota: utils.Int(shareQuota),
        },
    }


    if _, err := client.Update(ctx, resourceGroup, accountName, name, fileShare); err != nil {
        return fmt.Errorf("Error updating File Share %q (Account Name %q / Resource Group %q): %+v", name, accountName, resourceGroup, err)
    }

    return resourceArmFileShareRead(d, meta)
}

func resourceArmFileShareDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).fileSharesClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    accountName := id.Path["storageAccounts"]
    name := id.Path["shares"]

    if _, err := client.Delete(ctx, resourceGroup, accountName, name); err != nil {
        return fmt.Errorf("Error deleting File Share %q (Account Name %q / Resource Group %q): %+v", name, accountName, resourceGroup, err)
    }

    return nil
}