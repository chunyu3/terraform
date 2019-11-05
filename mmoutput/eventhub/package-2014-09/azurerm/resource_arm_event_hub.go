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



func resourceArmEventHub() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmEventHubCreateUpdate,
        Read: resourceArmEventHubRead,
        Update: resourceArmEventHubCreateUpdate,
        Delete: resourceArmEventHubDelete,

        Importer: &schema.ResourceImporter{
            State: schema.ImportStatePassthrough,
        },


        Schema: map[string]*schema.Schema{
            "name": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
            },

            "location": azure.SchemaLocation(),

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "event_hub_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "namespace_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "message_retention_in_days": {
                Type: schema.TypeInt,
                Optional: true,
            },

            "partition_count": {
                Type: schema.TypeInt,
                Optional: true,
            },

            "status": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(eventhub.Active),
                    string(eventhub.Disabled),
                    string(eventhub.Restoring),
                    string(eventhub.SendDisabled),
                    string(eventhub.ReceiveDisabled),
                    string(eventhub.Creating),
                    string(eventhub.Deleting),
                    string(eventhub.Renaming),
                    string(eventhub.Unknown),
                }, false),
                Default: string(eventhub.Active),
            },

            "type": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
            },

            "created_at": {
                Type: schema.TypeString,
                Computed: true,
            },

            "partition_ids": {
                Type: schema.TypeList,
                Computed: true,
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
            },

            "updated_at": {
                Type: schema.TypeString,
                Computed: true,
            },
        },
    }
}

func resourceArmEventHubCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).eventHubsClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    eventHubName := d.Get("event_hub_name").(string)
    namespaceName := d.Get("namespace_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, namespaceName, eventHubName)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Event Hub (Event Hub Name %q / Namespace Name %q / Resource Group %q): %+v", eventHubName, namespaceName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_event_hub", *existing.ID)
        }
    }

    name := d.Get("name").(string)
    location := azure.NormalizeLocation(d.Get("location").(string))
    messageRetentionInDays := d.Get("message_retention_in_days").(int)
    partitionCount := d.Get("partition_count").(int)
    status := d.Get("status").(string)
    type := d.Get("type").(string)

    parameters := eventhub.CreateOrUpdateParameters{
        Location: utils.String(location),
        Name: utils.String(name),
        Properties: &eventhub.Properties{
            MessageRetentionInDays: utils.Int64(int64(messageRetentionInDays)),
            PartitionCount: utils.Int64(int64(partitionCount)),
            Status: eventhub.EntityStatus(status),
        },
        Type: utils.String(type),
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroup, namespaceName, eventHubName, parameters); err != nil {
        return fmt.Errorf("Error creating Event Hub (Event Hub Name %q / Namespace Name %q / Resource Group %q): %+v", eventHubName, namespaceName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, namespaceName, eventHubName)
    if err != nil {
        return fmt.Errorf("Error retrieving Event Hub (Event Hub Name %q / Namespace Name %q / Resource Group %q): %+v", eventHubName, namespaceName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Event Hub (Event Hub Name %q / Namespace Name %q / Resource Group %q) ID", eventHubName, namespaceName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmEventHubRead(d, meta)
}

func resourceArmEventHubRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).eventHubsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    namespaceName := id.Path["namespaces"]
    eventHubName := id.Path["eventhubs"]

    resp, err := client.Get(ctx, resourceGroup, namespaceName, eventHubName)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Event Hub %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Event Hub (Event Hub Name %q / Namespace Name %q / Resource Group %q): %+v", eventHubName, namespaceName, resourceGroup, err)
    }


    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if properties := resp.Properties; properties != nil {
        d.Set("created_at", (properties.CreatedAt).String())
        d.Set("message_retention_in_days", int(*properties.MessageRetentionInDays))
        d.Set("partition_count", int(*properties.PartitionCount))
        d.Set("partition_ids", utils.FlattenStringSlice(properties.PartitionIds))
        d.Set("status", string(properties.Status))
        d.Set("updated_at", (properties.UpdatedAt).String())
    }
    d.Set("event_hub_name", eventHubName)
    d.Set("namespace_name", namespaceName)
    d.Set("type", resp.Type)

    return nil
}


func resourceArmEventHubDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).eventHubsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    namespaceName := id.Path["namespaces"]
    eventHubName := id.Path["eventhubs"]

    if _, err := client.Delete(ctx, resourceGroup, namespaceName, eventHubName); err != nil {
        return fmt.Errorf("Error deleting Event Hub (Event Hub Name %q / Namespace Name %q / Resource Group %q): %+v", eventHubName, namespaceName, resourceGroup, err)
    }

    return nil
}