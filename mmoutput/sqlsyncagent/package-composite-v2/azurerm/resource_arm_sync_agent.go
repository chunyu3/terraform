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



func resourceArmSyncAgent() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmSyncAgentCreateUpdate,
        Read: resourceArmSyncAgentRead,
        Update: resourceArmSyncAgentCreateUpdate,
        Delete: resourceArmSyncAgentDelete,

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

            "name": {
                Type: schema.TypeString,
                Computed: true,
            },

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "server_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "sync_database_id": {
                Type: schema.TypeString,
                Optional: true,
            },

            "expiry_time": {
                Type: schema.TypeString,
                Computed: true,
            },

            "is_up_to_date": {
                Type: schema.TypeBool,
                Computed: true,
            },

            "last_alive_time": {
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

            "version": {
                Type: schema.TypeString,
                Computed: true,
            },
        },
    }
}

func resourceArmSyncAgentCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).syncAgentsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    serverName := d.Get("server_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, serverName, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Sync Agent %q (Server Name %q / Resource Group %q): %+v", name, serverName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_sync_agent", *existing.ID)
        }
    }

    syncDatabaseId := d.Get("sync_database_id").(string)

    parameters := sql.SyncAgent{
        SyncAgentProperties: &sql.SyncAgentProperties{
            SyncDatabaseID: utils.String(syncDatabaseId),
        },
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, serverName, name, parameters)
    if err != nil {
        return fmt.Errorf("Error creating Sync Agent %q (Server Name %q / Resource Group %q): %+v", name, serverName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Sync Agent %q (Server Name %q / Resource Group %q): %+v", name, serverName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, serverName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Sync Agent %q (Server Name %q / Resource Group %q): %+v", name, serverName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Sync Agent %q (Server Name %q / Resource Group %q) ID", name, serverName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmSyncAgentRead(d, meta)
}

func resourceArmSyncAgentRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).syncAgentsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    serverName := id.Path["servers"]
    name := id.Path["syncAgents"]

    resp, err := client.Get(ctx, resourceGroup, serverName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Sync Agent %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Sync Agent %q (Server Name %q / Resource Group %q): %+v", name, serverName, resourceGroup, err)
    }


    d.Set("name", name)
    if syncAgentProperties := resp.SyncAgentProperties; syncAgentProperties != nil {
        d.Set("name", syncAgentProperties.Name)
        d.Set("expiry_time", (syncAgentProperties.ExpiryTime).String())
        d.Set("is_up_to_date", syncAgentProperties.IsUpToDate)
        d.Set("last_alive_time", (syncAgentProperties.LastAliveTime).String())
        d.Set("state", string(syncAgentProperties.State))
        d.Set("sync_database_id", syncAgentProperties.SyncDatabaseID)
        d.Set("version", syncAgentProperties.Version)
    }
    d.Set("name", resp.Name)
    d.Set("name", resp.Name)
    d.Set("name", resp.Name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    d.Set("server_name", serverName)
    d.Set("type", resp.Type)

    return nil
}


func resourceArmSyncAgentDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).syncAgentsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    serverName := id.Path["servers"]
    name := id.Path["syncAgents"]

    future, err := client.Delete(ctx, resourceGroup, serverName, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Sync Agent %q (Server Name %q / Resource Group %q): %+v", name, serverName, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Sync Agent %q (Server Name %q / Resource Group %q): %+v", name, serverName, resourceGroup, err)
        }
    }

    return nil
}
