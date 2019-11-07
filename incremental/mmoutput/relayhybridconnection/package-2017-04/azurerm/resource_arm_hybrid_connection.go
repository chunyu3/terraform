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



func resourceArmHybridConnection() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmHybridConnectionCreateUpdate,
        Read: resourceArmHybridConnectionRead,
        Update: resourceArmHybridConnectionCreateUpdate,
        Delete: resourceArmHybridConnectionDelete,

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

            "namespace_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "requires_client_authorization": {
                Type: schema.TypeBool,
                Optional: true,
            },

            "user_metadata": {
                Type: schema.TypeString,
                Optional: true,
            },

            "created_at": {
                Type: schema.TypeString,
                Computed: true,
            },

            "listener_count": {
                Type: schema.TypeInt,
                Computed: true,
            },

            "type": {
                Type: schema.TypeString,
                Computed: true,
            },

            "updated_at": {
                Type: schema.TypeString,
                Computed: true,
            },
        },
    }
}

func resourceArmHybridConnectionCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).hybridConnectionsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    namespaceName := d.Get("namespace_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, namespaceName, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Hybrid Connection %q (Namespace Name %q / Resource Group %q): %+v", name, namespaceName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_hybrid_connection", *existing.ID)
        }
    }

    requiresClientAuthorization := d.Get("requires_client_authorization").(bool)
    userMetadata := d.Get("user_metadata").(string)

    parameters := relay.HybridConnection{
        HybridConnection_properties: &relay.HybridConnection_properties{
            RequiresClientAuthorization: utils.Bool(requiresClientAuthorization),
            UserMetadata: utils.String(userMetadata),
        },
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroup, namespaceName, name, parameters); err != nil {
        return fmt.Errorf("Error creating Hybrid Connection %q (Namespace Name %q / Resource Group %q): %+v", name, namespaceName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, namespaceName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Hybrid Connection %q (Namespace Name %q / Resource Group %q): %+v", name, namespaceName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Hybrid Connection %q (Namespace Name %q / Resource Group %q) ID", name, namespaceName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmHybridConnectionRead(d, meta)
}

func resourceArmHybridConnectionRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).hybridConnectionsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    namespaceName := id.Path["namespaces"]
    name := id.Path["hybridConnections"]

    resp, err := client.Get(ctx, resourceGroup, namespaceName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Hybrid Connection %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Hybrid Connection %q (Namespace Name %q / Resource Group %q): %+v", name, namespaceName, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if hybridConnectionProperties := resp.HybridConnection_properties; hybridConnectionProperties != nil {
        d.Set("created_at", (hybridConnectionProperties.CreatedAt).String())
        d.Set("listener_count", int(*hybridConnectionProperties.ListenerCount))
        d.Set("requires_client_authorization", hybridConnectionProperties.RequiresClientAuthorization)
        d.Set("updated_at", (hybridConnectionProperties.UpdatedAt).String())
        d.Set("user_metadata", hybridConnectionProperties.UserMetadata)
    }
    d.Set("namespace_name", namespaceName)
    d.Set("type", resp.Type)

    return nil
}


func resourceArmHybridConnectionDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).hybridConnectionsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    namespaceName := id.Path["namespaces"]
    name := id.Path["hybridConnections"]

    if _, err := client.Delete(ctx, resourceGroup, namespaceName, name); err != nil {
        return fmt.Errorf("Error deleting Hybrid Connection %q (Namespace Name %q / Resource Group %q): %+v", name, namespaceName, resourceGroup, err)
    }

    return nil
}