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



func resourceArmRemoteRenderingAccount() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmRemoteRenderingAccountCreate,
        Read: resourceArmRemoteRenderingAccountRead,
        Update: resourceArmRemoteRenderingAccountUpdate,
        Delete: resourceArmRemoteRenderingAccountDelete,

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

            "location": azure.SchemaLocation(),

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

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

func resourceArmRemoteRenderingAccountCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).remoteRenderingAccountsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Remote Rendering Account %q (Resource Group %q): %+v", name, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_remote_rendering_account", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    t := d.Get("tags").(map[string]interface{})

    remoteRenderingAccount := mixedreality.RemoteRenderingAccount{
        Location: utils.String(location),
        Tags: tags.Expand(t),
    }


    if _, err := client.Create(ctx, resourceGroup, name, remoteRenderingAccount); err != nil {
        return fmt.Errorf("Error creating Remote Rendering Account %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Remote Rendering Account %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Remote Rendering Account %q (Resource Group %q) ID", name, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmRemoteRenderingAccountRead(d, meta)
}

func resourceArmRemoteRenderingAccountRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).remoteRenderingAccountsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["remoteRenderingAccounts"]

    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Remote Rendering Account %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Remote Rendering Account %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if accountProperties := resp.AccountProperties; accountProperties != nil {
        d.Set("account_domain", accountProperties.AccountDomain)
        d.Set("account_id", accountProperties.AccountID)
    }
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}

func resourceArmRemoteRenderingAccountUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).remoteRenderingAccountsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    t := d.Get("tags").(map[string]interface{})

    remoteRenderingAccount := mixedreality.RemoteRenderingAccount{
        Location: utils.String(location),
        Tags: tags.Expand(t),
    }


    if _, err := client.Update(ctx, resourceGroup, name, remoteRenderingAccount); err != nil {
        return fmt.Errorf("Error updating Remote Rendering Account %q (Resource Group %q): %+v", name, resourceGroup, err)
    }

    return resourceArmRemoteRenderingAccountRead(d, meta)
}

func resourceArmRemoteRenderingAccountDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).remoteRenderingAccountsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["remoteRenderingAccounts"]

    if _, err := client.Delete(ctx, resourceGroup, name); err != nil {
        return fmt.Errorf("Error deleting Remote Rendering Account %q (Resource Group %q): %+v", name, resourceGroup, err)
    }

    return nil
}
