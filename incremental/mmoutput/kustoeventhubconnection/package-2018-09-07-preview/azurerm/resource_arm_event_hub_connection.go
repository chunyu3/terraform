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

            "cluster_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "consumer_group": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "database_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "event_hub_resource_id": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "data_format": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(kusto.MULTIJSON),
                    string(kusto.JSON),
                    string(kusto.CSV),
                }, false),
                Default: string(kusto.MULTIJSON),
            },

            "mapping_rule_name": {
                Type: schema.TypeString,
                Optional: true,
            },

            "table_name": {
                Type: schema.TypeString,
                Optional: true,
            },

            "type": {
                Type: schema.TypeString,
                Computed: true,
            },
        },
    }
}

func resourceArmEventHubConnectionCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).eventHubConnectionsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    clusterName := d.Get("cluster_name").(string)
    databaseName := d.Get("database_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, clusterName, databaseName, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Event Hub Connection %q (Database Name %q / Cluster Name %q / Resource Group %q): %+v", name, databaseName, clusterName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_event_hub_connection", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    consumerGroup := d.Get("consumer_group").(string)
    dataFormat := d.Get("data_format").(string)
    eventHubResourceId := d.Get("event_hub_resource_id").(string)
    mappingRuleName := d.Get("mapping_rule_name").(string)
    tableName := d.Get("table_name").(string)

    parameters := kusto.EventHubConnectionUpdate{
        Location: utils.String(location),
        EventHubConnectionProperties: &kusto.EventHubConnectionProperties{
            ConsumerGroup: utils.String(consumerGroup),
            DataFormat: kusto.DataFormat(dataFormat),
            EventHubResourceID: utils.String(eventHubResourceId),
            MappingRuleName: utils.String(mappingRuleName),
            TableName: utils.String(tableName),
        },
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, clusterName, databaseName, name, parameters)
    if err != nil {
        return fmt.Errorf("Error creating Event Hub Connection %q (Database Name %q / Cluster Name %q / Resource Group %q): %+v", name, databaseName, clusterName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Event Hub Connection %q (Database Name %q / Cluster Name %q / Resource Group %q): %+v", name, databaseName, clusterName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, clusterName, databaseName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Event Hub Connection %q (Database Name %q / Cluster Name %q / Resource Group %q): %+v", name, databaseName, clusterName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Event Hub Connection %q (Database Name %q / Cluster Name %q / Resource Group %q) ID", name, databaseName, clusterName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmEventHubConnectionRead(d, meta)
}

func resourceArmEventHubConnectionRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).eventHubConnectionsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    clusterName := id.Path["clusters"]
    databaseName := id.Path["databases"]
    name := id.Path["eventhubconnections"]

    resp, err := client.Get(ctx, resourceGroup, clusterName, databaseName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Event Hub Connection %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Event Hub Connection %q (Database Name %q / Cluster Name %q / Resource Group %q): %+v", name, databaseName, clusterName, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    d.Set("cluster_name", clusterName)
    if eventHubConnectionProperties := resp.EventHubConnectionProperties; eventHubConnectionProperties != nil {
        d.Set("consumer_group", eventHubConnectionProperties.ConsumerGroup)
        d.Set("data_format", string(eventHubConnectionProperties.DataFormat))
        d.Set("event_hub_resource_id", eventHubConnectionProperties.EventHubResourceID)
        d.Set("mapping_rule_name", eventHubConnectionProperties.MappingRuleName)
        d.Set("table_name", eventHubConnectionProperties.TableName)
    }
    d.Set("database_name", databaseName)
    d.Set("type", resp.Type)

    return nil
}

func resourceArmEventHubConnectionUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).eventHubConnectionsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    clusterName := d.Get("cluster_name").(string)
    consumerGroup := d.Get("consumer_group").(string)
    dataFormat := d.Get("data_format").(string)
    databaseName := d.Get("database_name").(string)
    eventHubResourceId := d.Get("event_hub_resource_id").(string)
    mappingRuleName := d.Get("mapping_rule_name").(string)
    tableName := d.Get("table_name").(string)

    parameters := kusto.EventHubConnectionUpdate{
        Location: utils.String(location),
        EventHubConnectionProperties: &kusto.EventHubConnectionProperties{
            ConsumerGroup: utils.String(consumerGroup),
            DataFormat: kusto.DataFormat(dataFormat),
            EventHubResourceID: utils.String(eventHubResourceId),
            MappingRuleName: utils.String(mappingRuleName),
            TableName: utils.String(tableName),
        },
    }


    future, err := client.Update(ctx, resourceGroup, clusterName, databaseName, name, parameters)
    if err != nil {
        return fmt.Errorf("Error updating Event Hub Connection %q (Database Name %q / Cluster Name %q / Resource Group %q): %+v", name, databaseName, clusterName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for update of Event Hub Connection %q (Database Name %q / Cluster Name %q / Resource Group %q): %+v", name, databaseName, clusterName, resourceGroup, err)
    }

    return resourceArmEventHubConnectionRead(d, meta)
}

func resourceArmEventHubConnectionDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).eventHubConnectionsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    clusterName := id.Path["clusters"]
    databaseName := id.Path["databases"]
    name := id.Path["eventhubconnections"]

    future, err := client.Delete(ctx, resourceGroup, clusterName, databaseName, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Event Hub Connection %q (Database Name %q / Cluster Name %q / Resource Group %q): %+v", name, databaseName, clusterName, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Event Hub Connection %q (Database Name %q / Cluster Name %q / Resource Group %q): %+v", name, databaseName, clusterName, resourceGroup, err)
        }
    }

    return nil
}
