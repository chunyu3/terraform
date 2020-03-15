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



func resourceArmServerKey() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmServerKeyCreateUpdate,
        Read: resourceArmServerKeyRead,
        Update: resourceArmServerKeyCreateUpdate,
        Delete: resourceArmServerKeyDelete,

        Importer: &schema.ResourceImporter{
            State: schema.ImportStatePassthrough,
        },


        Schema: map[string]*schema.Schema{
            "key_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "server_key_type": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "server_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "uri": {
                Type: schema.TypeString,
                Optional: true,
            },

            "creation_date": {
                Type: schema.TypeString,
                Computed: true,
            },

            "id": {
                Type: schema.TypeString,
                Computed: true,
            },

            "kind": {
                Type: schema.TypeString,
                Computed: true,
            },

            "name": {
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

func resourceArmServerKeyCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).serverKeysClient
    ctx, cancel := timeouts.ForCreateUpdate(meta.(*ArmClient).StopContext, d)
    defer cancel()

    resourceGroupName := d.Get("resource_group").(string)
    name := d.Get("key_name").(string)
    serverName := d.Get("server_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroupName, serverName, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Server Key (Key Name %q / Server Name %q / Resource Group %q): %+v", name, serverName, resourceGroupName, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_server_key", *existing.ID)
        }
    }

    serverKeyType := d.Get("server_key_type").(string)
    uRI := d.Get("uri").(string)

    parameters := mysql.ServerKey{
        ServerKeyProperties: &mysql.ServerKeyProperties{
            ServerKeyType: utils.String(serverKeyType),
            URI: utils.String(uRI),
        },
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroupName, serverName, name, parameters)
    if err != nil {
        return fmt.Errorf("Error creating Server Key (Key Name %q / Server Name %q / Resource Group %q): %+v", name, serverName, resourceGroupName, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Server Key (Key Name %q / Server Name %q / Resource Group %q): %+v", name, serverName, resourceGroupName, err)
    }


    resp, err := client.Get(ctx, resourceGroupName, serverName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Server Key (Key Name %q / Server Name %q / Resource Group %q): %+v", name, serverName, resourceGroupName, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Server Key (Key Name %q / Server Name %q / Resource Group %q) ID", name, serverName, resourceGroupName)
    }
    d.SetId(*resp.ID)

    return resourceArmServerKeyRead(d, meta)
}

func resourceArmServerKeyRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).serverKeysClient
    ctx, cancel := timeouts.ForRead(meta.(*ArmClient).StopContext, d)
    defer cancel()

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroupName := id.ResourceGroup
    serverName := id.Path["servers"]
    name := id.Path["keys"]

    resp, err := client.Get(ctx, resourceGroupName, serverName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Server Key %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Server Key (Key Name %q / Server Name %q / Resource Group %q): %+v", name, serverName, resourceGroupName, err)
    }


    d.Set("resource_group", resourceGroupName)
    if serverKeyProperties := resp.ServerKeyProperties; serverKeyProperties != nil {
        d.Set("creation_date", (serverKeyProperties.CreationDate).String())
        d.Set("server_key_type", serverKeyProperties.ServerKeyType)
        d.Set("uri", serverKeyProperties.URI)
    }
    d.Set("id", resp.ID)
    d.Set("key_name", name)
    d.Set("kind", resp.Kind)
    d.Set("name", resp.Name)
    d.Set("server_name", serverName)
    d.Set("type", resp.Type)

    return nil
}


func resourceArmServerKeyDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).serverKeysClient
    ctx, cancel := timeouts.ForDelete(meta.(*ArmClient).StopContext, d)
    defer cancel()


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroupName := id.ResourceGroup
    serverName := id.Path["servers"]
    name := id.Path["keys"]

    future, err := client.Delete(ctx, resourceGroupName, serverName, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Server Key (Key Name %q / Server Name %q / Resource Group %q): %+v", name, serverName, resourceGroupName, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Server Key (Key Name %q / Server Name %q / Resource Group %q): %+v", name, serverName, resourceGroupName, err)
        }
    }

    return nil
}
