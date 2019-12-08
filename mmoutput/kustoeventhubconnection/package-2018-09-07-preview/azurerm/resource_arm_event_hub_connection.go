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



func resourceArmEventHubConnection() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmEventHubConnectionCreate,
        Read: resourceArmEventHubConnectionRead,
        Update: resourceArmEventHubConnectionUpdate,
        Delete: resourceArmEventHubConnectionDelete,

        Importer: &schema.ResourceImporter{
            State: schema.ImportStatePassthrough,
        },


        Schema: map[string]*schema.Schema{
            "cluster_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "database_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "event_hub_connection_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),
        },
    }
}

func resourceArmEventHubConnectionCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).eventHubConnectionsClient
    ctx, cancel := timeouts.ForCreate(meta.(*ArmClient).StopContext, d)
    defer cancel()

    resourceGroupName := d.Get("resource_group").(string)
    clusterName := d.Get("cluster_name").(string)
    databaseName := d.Get("database_name").(string)
    name := d.Get("event_hub_connection_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroupName, clusterName, databaseName, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Event Hub Connection (Event Hub Connection Name %q / Database Name %q / Cluster Name %q / Resource Group %q): %+v", name, databaseName, clusterName, resourceGroupName, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_event_hub_connection", *existing.ID)
        }
    }


    parameters := kusto.EventHubConnectionUpdate{
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroupName, clusterName, databaseName, name, parameters)
    if err != nil {
        return fmt.Errorf("Error creating Event Hub Connection (Event Hub Connection Name %q / Database Name %q / Cluster Name %q / Resource Group %q): %+v", name, databaseName, clusterName, resourceGroupName, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Event Hub Connection (Event Hub Connection Name %q / Database Name %q / Cluster Name %q / Resource Group %q): %+v", name, databaseName, clusterName, resourceGroupName, err)
    }


    resp, err := client.Get(ctx, resourceGroupName, clusterName, databaseName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Event Hub Connection (Event Hub Connection Name %q / Database Name %q / Cluster Name %q / Resource Group %q): %+v", name, databaseName, clusterName, resourceGroupName, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Event Hub Connection (Event Hub Connection Name %q / Database Name %q / Cluster Name %q / Resource Group %q) ID", name, databaseName, clusterName, resourceGroupName)
    }
    d.SetId(*resp.ID)

    return resourceArmEventHubConnectionRead(d, meta)
}

func resourceArmEventHubConnectionRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).eventHubConnectionsClient
    ctx, cancel := timeouts.ForRead(meta.(*ArmClient).StopContext, d)
    defer cancel()

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroupName := id.ResourceGroup
    clusterName := id.Path["clusters"]
    databaseName := id.Path["databases"]
    name := id.Path["eventhubconnections"]

    resp, err := client.Get(ctx, resourceGroupName, clusterName, databaseName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Event Hub Connection %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Event Hub Connection (Event Hub Connection Name %q / Database Name %q / Cluster Name %q / Resource Group %q): %+v", name, databaseName, clusterName, resourceGroupName, err)
    }


    d.Set("resource_group", resourceGroupName)
    d.Set("cluster_name", clusterName)
    d.Set("database_name", databaseName)
    d.Set("event_hub_connection_name", name)

    return nil
}

func resourceArmEventHubConnectionUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).eventHubConnectionsClient
    ctx, cancel := timeouts.ForUpdate(meta.(*ArmClient).StopContext, d)
    defer cancel()

      resourceGroupName := d.Get("resource_group").(string)
    clusterName := d.Get("cluster_name").(string)
    databaseName := d.Get("database_name").(string)
    name := d.Get("event_hub_connection_name").(string)

    parameters := kusto.EventHubConnectionUpdate{
    }


    future, err := client.Update(ctx, resourceGroupName, clusterName, databaseName, name, parameters)
    if err != nil {
        return fmt.Errorf("Error updating Event Hub Connection (Event Hub Connection Name %q / Database Name %q / Cluster Name %q / Resource Group %q): %+v", name, databaseName, clusterName, resourceGroupName, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for update of Event Hub Connection (Event Hub Connection Name %q / Database Name %q / Cluster Name %q / Resource Group %q): %+v", name, databaseName, clusterName, resourceGroupName, err)
    }

    return resourceArmEventHubConnectionRead(d, meta)
}

func resourceArmEventHubConnectionDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).eventHubConnectionsClient
    ctx, cancel := timeouts.ForDelete(meta.(*ArmClient).StopContext, d)
    defer cancel()


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroupName := id.ResourceGroup
    clusterName := id.Path["clusters"]
    databaseName := id.Path["databases"]
    name := id.Path["eventhubconnections"]

    future, err := client.Delete(ctx, resourceGroupName, clusterName, databaseName, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Event Hub Connection (Event Hub Connection Name %q / Database Name %q / Cluster Name %q / Resource Group %q): %+v", name, databaseName, clusterName, resourceGroupName, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Event Hub Connection (Event Hub Connection Name %q / Database Name %q / Cluster Name %q / Resource Group %q): %+v", name, databaseName, clusterName, resourceGroupName, err)
        }
    }

    return nil
}
