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



func resourceArmApplication() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmApplicationCreate,
        Read: resourceArmApplicationRead,
        Update: resourceArmApplicationUpdate,
        Delete: resourceArmApplicationDelete,

        Importer: &schema.ResourceImporter{
            State: schema.ImportStatePassthrough,
        },


        Schema: map[string]*schema.Schema{
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

            "application_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "allow_updates": {
                Type: schema.TypeBool,
                Optional: true,
            },

            "default_version": {
                Type: schema.TypeString,
                Optional: true,
            },

            "display_name": {
                Type: schema.TypeString,
                Optional: true,
            },

            "etag": {
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

func resourceArmApplicationCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).applicationClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    accountName := d.Get("account_name").(string)
    applicationName := d.Get("application_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, accountName, applicationName)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Application (Application Name %q / Account Name %q / Resource Group %q): %+v", applicationName, accountName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_application", *existing.ID)
        }
    }

    allowUpdates := d.Get("allow_updates").(bool)
    defaultVersion := d.Get("default_version").(string)
    displayName := d.Get("display_name").(string)

    parameters := batch.Application{
        ApplicationProperties: &batch.ApplicationProperties{
            AllowUpdates: utils.Bool(allowUpdates),
            DefaultVersion: utils.String(defaultVersion),
            DisplayName: utils.String(displayName),
        },
    }


    if _, err := client.Create(ctx, resourceGroup, accountName, applicationName, parameters); err != nil {
        return fmt.Errorf("Error creating Application (Application Name %q / Account Name %q / Resource Group %q): %+v", applicationName, accountName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, accountName, applicationName)
    if err != nil {
        return fmt.Errorf("Error retrieving Application (Application Name %q / Account Name %q / Resource Group %q): %+v", applicationName, accountName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Application (Application Name %q / Account Name %q / Resource Group %q) ID", applicationName, accountName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmApplicationRead(d, meta)
}

func resourceArmApplicationRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).applicationClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    accountName := id.Path["batchAccounts"]
    applicationName := id.Path["applications"]

    resp, err := client.Get(ctx, resourceGroup, accountName, applicationName)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Application %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Application (Application Name %q / Account Name %q / Resource Group %q): %+v", applicationName, accountName, resourceGroup, err)
    }


    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    d.Set("account_name", accountName)
    if applicationProperties := resp.ApplicationProperties; applicationProperties != nil {
        d.Set("allow_updates", applicationProperties.AllowUpdates)
        d.Set("default_version", applicationProperties.DefaultVersion)
        d.Set("display_name", applicationProperties.DisplayName)
    }
    d.Set("application_name", applicationName)
    d.Set("etag", resp.Etag)
    d.Set("type", resp.Type)

    return nil
}

func resourceArmApplicationUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).applicationClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    accountName := d.Get("account_name").(string)
    allowUpdates := d.Get("allow_updates").(bool)
    applicationName := d.Get("application_name").(string)
    defaultVersion := d.Get("default_version").(string)
    displayName := d.Get("display_name").(string)

    parameters := batch.Application{
        ApplicationProperties: &batch.ApplicationProperties{
            AllowUpdates: utils.Bool(allowUpdates),
            DefaultVersion: utils.String(defaultVersion),
            DisplayName: utils.String(displayName),
        },
    }


    if _, err := client.Update(ctx, resourceGroup, accountName, applicationName, parameters); err != nil {
        return fmt.Errorf("Error updating Application (Application Name %q / Account Name %q / Resource Group %q): %+v", applicationName, accountName, resourceGroup, err)
    }

    return resourceArmApplicationRead(d, meta)
}

func resourceArmApplicationDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).applicationClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    accountName := id.Path["batchAccounts"]
    applicationName := id.Path["applications"]

    if _, err := client.Delete(ctx, resourceGroup, accountName, applicationName); err != nil {
        return fmt.Errorf("Error deleting Application (Application Name %q / Account Name %q / Resource Group %q): %+v", applicationName, accountName, resourceGroup, err)
    }

    return nil
}
