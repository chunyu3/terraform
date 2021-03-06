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



func resourceArmServerCommunicationLink() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmServerCommunicationLinkCreateUpdate,
        Read: resourceArmServerCommunicationLinkRead,
        Update: resourceArmServerCommunicationLinkCreateUpdate,
        Delete: resourceArmServerCommunicationLinkDelete,

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

            "partner_server": {
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

            "kind": {
                Type: schema.TypeString,
                Computed: true,
            },

            "state": {
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

func resourceArmServerCommunicationLinkCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).serverCommunicationLinksClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    serverName := d.Get("server_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, serverName, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Server Communication Link %q (Server Name %q / Resource Group %q): %+v", name, serverName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_server_communication_link", *existing.ID)
        }
    }

    partnerServer := d.Get("partner_server").(string)

    parameters := sql.ServerCommunicationLink{
        ServerCommunicationLinkProperties: &sql.ServerCommunicationLinkProperties{
            PartnerServer: utils.String(partnerServer),
        },
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, serverName, name, parameters)
    if err != nil {
        return fmt.Errorf("Error creating Server Communication Link %q (Server Name %q / Resource Group %q): %+v", name, serverName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Server Communication Link %q (Server Name %q / Resource Group %q): %+v", name, serverName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, serverName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Server Communication Link %q (Server Name %q / Resource Group %q): %+v", name, serverName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Server Communication Link %q (Server Name %q / Resource Group %q) ID", name, serverName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmServerCommunicationLinkRead(d, meta)
}

func resourceArmServerCommunicationLinkRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).serverCommunicationLinksClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    serverName := id.Path["servers"]
    name := id.Path["communicationLinks"]

    resp, err := client.Get(ctx, resourceGroup, serverName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Server Communication Link %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Server Communication Link %q (Server Name %q / Resource Group %q): %+v", name, serverName, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    d.Set("kind", resp.Kind)
    if serverCommunicationLinkProperties := resp.ServerCommunicationLinkProperties; serverCommunicationLinkProperties != nil {
        d.Set("partner_server", serverCommunicationLinkProperties.PartnerServer)
        d.Set("state", serverCommunicationLinkProperties.State)
    }
    d.Set("server_name", serverName)
    d.Set("type", resp.Type)

    return nil
}


func resourceArmServerCommunicationLinkDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).serverCommunicationLinksClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    serverName := id.Path["servers"]
    name := id.Path["communicationLinks"]

    if _, err := client.Delete(ctx, resourceGroup, serverName, name); err != nil {
        return fmt.Errorf("Error deleting Server Communication Link %q (Server Name %q / Resource Group %q): %+v", name, serverName, resourceGroup, err)
    }

    return nil
}
