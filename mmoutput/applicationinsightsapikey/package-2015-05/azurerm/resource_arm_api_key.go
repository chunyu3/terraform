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



func resourceArmAPIKey() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmAPIKeyCreateUpdate,
        Read: resourceArmAPIKeyRead,
        Update: resourceArmAPIKeyCreateUpdate,
        Delete: resourceArmAPIKeyDelete,

        Importer: &schema.ResourceImporter{
            State: schema.ImportStatePassthrough,
        },


        Schema: map[string]*schema.Schema{
            "name": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
            },

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "resource_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "linked_read_properties": {
                Type: schema.TypeList,
                Optional: true,
                ForceNew: true,
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
            },

            "linked_write_properties": {
                Type: schema.TypeList,
                Optional: true,
                ForceNew: true,
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
            },

            "api_key": {
                Type: schema.TypeString,
                Computed: true,
            },

            "created_date": {
                Type: schema.TypeString,
                Computed: true,
            },
        },
    }
}

func resourceArmAPIKeyCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).aPIKeysClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    resourceName := d.Get("resource_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, resourceName)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Api Key (Resource Name %q / Resource Group %q): %+v", resourceName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_api_key", *existing.ID)
        }
    }

    name := d.Get("name").(string)
    linkedReadProperties := d.Get("linked_read_properties").([]interface{})
    linkedWriteProperties := d.Get("linked_write_properties").([]interface{})

    apikeyProperties := applicationinsights.APIKeyRequest{
        LinkedReadProperties: utils.ExpandStringSlice(linkedReadProperties),
        LinkedWriteProperties: utils.ExpandStringSlice(linkedWriteProperties),
        Name: utils.String(name),
    }


    if _, err := client.Create(ctx, resourceGroup, resourceName, apikeyProperties); err != nil {
        return fmt.Errorf("Error creating Api Key (Resource Name %q / Resource Group %q): %+v", resourceName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, resourceName)
    if err != nil {
        return fmt.Errorf("Error retrieving Api Key (Resource Name %q / Resource Group %q): %+v", resourceName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Api Key (Resource Name %q / Resource Group %q) ID", resourceName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmAPIKeyRead(d, meta)
}

func resourceArmAPIKeyRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).aPIKeysClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    resourceName := id.Path["components"]

    resp, err := client.Get(ctx, resourceGroup, resourceName)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Api Key %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Api Key (Resource Name %q / Resource Group %q): %+v", resourceName, resourceGroup, err)
    }


    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    d.Set("api_key", resp.ApiKey)
    d.Set("created_date", resp.CreatedDate)
    d.Set("linked_read_properties", utils.FlattenStringSlice(resp.LinkedReadProperties))
    d.Set("linked_write_properties", utils.FlattenStringSlice(resp.LinkedWriteProperties))
    d.Set("resource_name", resourceName)

    return nil
}


func resourceArmAPIKeyDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).aPIKeysClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    resourceName := id.Path["components"]

    if _, err := client.Delete(ctx, resourceGroup, resourceName); err != nil {
        return fmt.Errorf("Error deleting Api Key (Resource Name %q / Resource Group %q): %+v", resourceName, resourceGroup, err)
    }

    return nil
}