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



func resourceArmSpatialAnchorsAccount() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmSpatialAnchorsAccountCreate,
        Read: resourceArmSpatialAnchorsAccountRead,
        Update: resourceArmSpatialAnchorsAccountUpdate,
        Delete: resourceArmSpatialAnchorsAccountDelete,

        Importer: &schema.ResourceImporter{
            State: schema.ImportStatePassthrough,
        },


        Schema: map[string]*schema.Schema{
            "name": {
                Type: schema.TypeString,
                Computed: true,
            },

            "location": azure.SchemaLocation(),

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "account_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "account_domain": {
                Type: schema.TypeString,
                Computed: true,
            },

            "account_id": {
                Type: schema.TypeString,
                Computed: true,
            },

            "type": {
                Type: schema.TypeString,
                Computed: true,
            },

            "tags": tags.Schema(),
        },
    }
}

func resourceArmSpatialAnchorsAccountCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).spatialAnchorsAccountsClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    accountName := d.Get("account_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, accountName)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Spatial Anchors Account (Account Name %q / Resource Group %q): %+v", accountName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_spatial_anchors_account", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    t := d.Get("tags").(map[string]interface{})

    spatialAnchorsAccount := mixedreality.SpatialAnchorsAccount{
        Location: utils.String(location),
        Tags: tags.Expand(t),
    }


    if _, err := client.Create(ctx, resourceGroup, accountName, spatialAnchorsAccount); err != nil {
        return fmt.Errorf("Error creating Spatial Anchors Account (Account Name %q / Resource Group %q): %+v", accountName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, accountName)
    if err != nil {
        return fmt.Errorf("Error retrieving Spatial Anchors Account (Account Name %q / Resource Group %q): %+v", accountName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Spatial Anchors Account (Account Name %q / Resource Group %q) ID", accountName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmSpatialAnchorsAccountRead(d, meta)
}

func resourceArmSpatialAnchorsAccountRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).spatialAnchorsAccountsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    accountName := id.Path["spatialAnchorsAccounts"]

    resp, err := client.Get(ctx, resourceGroup, accountName)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Spatial Anchors Account %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Spatial Anchors Account (Account Name %q / Resource Group %q): %+v", accountName, resourceGroup, err)
    }


    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if accountProperties := resp.AccountProperties; accountProperties != nil {
        d.Set("account_domain", accountProperties.AccountDomain)
        d.Set("account_id", accountProperties.AccountID)
    }
    d.Set("account_name", accountName)
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}

func resourceArmSpatialAnchorsAccountUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).spatialAnchorsAccountsClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    accountName := d.Get("account_name").(string)
    t := d.Get("tags").(map[string]interface{})

    spatialAnchorsAccount := mixedreality.SpatialAnchorsAccount{
        Location: utils.String(location),
        Tags: tags.Expand(t),
    }


    if _, err := client.Update(ctx, resourceGroup, accountName, spatialAnchorsAccount); err != nil {
        return fmt.Errorf("Error updating Spatial Anchors Account (Account Name %q / Resource Group %q): %+v", accountName, resourceGroup, err)
    }

    return resourceArmSpatialAnchorsAccountRead(d, meta)
}

func resourceArmSpatialAnchorsAccountDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).spatialAnchorsAccountsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    accountName := id.Path["spatialAnchorsAccounts"]

    if _, err := client.Delete(ctx, resourceGroup, accountName); err != nil {
        return fmt.Errorf("Error deleting Spatial Anchors Account (Account Name %q / Resource Group %q): %+v", accountName, resourceGroup, err)
    }

    return nil
}
