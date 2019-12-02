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



func resourceArmSubscription() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmSubscriptionCreateUpdate,
        Read: resourceArmSubscriptionRead,
        Update: resourceArmSubscriptionCreateUpdate,
        Delete: resourceArmSubscriptionDelete,

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

            "namespace_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "topic_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "auto_delete_on_idle": {
                Type: schema.TypeString,
                Optional: true,
            },

            "dead_lettering_on_filter_evaluation_exceptions": {
                Type: schema.TypeBool,
                Optional: true,
            },

            "dead_lettering_on_message_expiration": {
                Type: schema.TypeBool,
                Optional: true,
            },

            "default_message_time_to_live": {
                Type: schema.TypeString,
                Optional: true,
            },

            "enable_batched_operations": {
                Type: schema.TypeBool,
                Optional: true,
            },

            "entity_availability_status": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(servicebus.Available),
                    string(servicebus.Limited),
                    string(servicebus.Renaming),
                    string(servicebus.Restoring),
                    string(servicebus.Unknown),
                }, false),
                Default: string(servicebus.Available),
            },

            "is_read_only": {
                Type: schema.TypeBool,
                Optional: true,
            },

            "lock_duration": {
                Type: schema.TypeString,
                Optional: true,
            },

            "max_delivery_count": {
                Type: schema.TypeInt,
                Optional: true,
            },

            "requires_session": {
                Type: schema.TypeBool,
                Optional: true,
            },

            "status": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(servicebus.Active),
                    string(servicebus.Creating),
                    string(servicebus.Deleting),
                    string(servicebus.Disabled),
                    string(servicebus.ReceiveDisabled),
                    string(servicebus.Renaming),
                    string(servicebus.Restoring),
                    string(servicebus.SendDisabled),
                    string(servicebus.Unknown),
                }, false),
                Default: string(servicebus.Active),
            },

            "type": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
            },

            "accessed_at": {
                Type: schema.TypeString,
                Computed: true,
            },

            "count_details": {
                Type: schema.TypeList,
                Computed: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "active_message_count": {
                            Type: schema.TypeInt,
                            Computed: true,
                        },
                        "dead_letter_message_count": {
                            Type: schema.TypeInt,
                            Computed: true,
                        },
                        "scheduled_message_count": {
                            Type: schema.TypeInt,
                            Computed: true,
                        },
                        "transfer_dead_letter_message_count": {
                            Type: schema.TypeInt,
                            Computed: true,
                        },
                        "transfer_message_count": {
                            Type: schema.TypeInt,
                            Computed: true,
                        },
                    },
                },
            },

            "created_at": {
                Type: schema.TypeString,
                Computed: true,
            },

            "message_count": {
                Type: schema.TypeInt,
                Computed: true,
            },

            "updated_at": {
                Type: schema.TypeString,
                Computed: true,
            },
        },
    }
}

func resourceArmSubscriptionCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).subscriptionsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    namespaceName := d.Get("namespace_name").(string)
    topicName := d.Get("topic_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, namespaceName, topicName, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Subscription %q (Topic Name %q / Namespace Name %q / Resource Group %q): %+v", name, topicName, namespaceName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_subscription", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    autoDeleteOnIdle := d.Get("auto_delete_on_idle").(string)
    deadLetteringOnFilterEvaluationExceptions := d.Get("dead_lettering_on_filter_evaluation_exceptions").(bool)
    deadLetteringOnMessageExpiration := d.Get("dead_lettering_on_message_expiration").(bool)
    defaultMessageTimeToLive := d.Get("default_message_time_to_live").(string)
    enableBatchedOperations := d.Get("enable_batched_operations").(bool)
    entityAvailabilityStatus := d.Get("entity_availability_status").(string)
    isReadOnly := d.Get("is_read_only").(bool)
    lockDuration := d.Get("lock_duration").(string)
    maxDeliveryCount := d.Get("max_delivery_count").(int)
    requiresSession := d.Get("requires_session").(bool)
    status := d.Get("status").(string)
    type := d.Get("type").(string)

    parameters := servicebus.SubscriptionCreateOrUpdateParameters{
        Location: utils.String(location),
        SubscriptionProperties: &servicebus.SubscriptionProperties{
            AutoDeleteOnIdle: utils.String(autoDeleteOnIdle),
            DeadLetteringOnFilterEvaluationExceptions: utils.Bool(deadLetteringOnFilterEvaluationExceptions),
            DeadLetteringOnMessageExpiration: utils.Bool(deadLetteringOnMessageExpiration),
            DefaultMessageTimeToLive: utils.String(defaultMessageTimeToLive),
            EnableBatchedOperations: utils.Bool(enableBatchedOperations),
            EntityAvailabilityStatus: servicebus.EntityAvailabilityStatus(entityAvailabilityStatus),
            IsReadOnly: utils.Bool(isReadOnly),
            LockDuration: utils.String(lockDuration),
            MaxDeliveryCount: utils.Int32(int32(maxDeliveryCount)),
            RequiresSession: utils.Bool(requiresSession),
            Status: servicebus.EntityStatus(status),
        },
        Type: utils.String(type),
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroup, namespaceName, topicName, name, parameters); err != nil {
        return fmt.Errorf("Error creating Subscription %q (Topic Name %q / Namespace Name %q / Resource Group %q): %+v", name, topicName, namespaceName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, namespaceName, topicName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Subscription %q (Topic Name %q / Namespace Name %q / Resource Group %q): %+v", name, topicName, namespaceName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Subscription %q (Topic Name %q / Namespace Name %q / Resource Group %q) ID", name, topicName, namespaceName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmSubscriptionRead(d, meta)
}

func resourceArmSubscriptionRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).subscriptionsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    namespaceName := id.Path["namespaces"]
    topicName := id.Path["topics"]
    name := id.Path["subscriptions"]

    resp, err := client.Get(ctx, resourceGroup, namespaceName, topicName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Subscription %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Subscription %q (Topic Name %q / Namespace Name %q / Resource Group %q): %+v", name, topicName, namespaceName, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if subscriptionProperties := resp.SubscriptionProperties; subscriptionProperties != nil {
        d.Set("accessed_at", (subscriptionProperties.AccessedAt).String())
        d.Set("auto_delete_on_idle", subscriptionProperties.AutoDeleteOnIdle)
        if err := d.Set("count_details", flattenArmSubscriptionMessageCountDetails(subscriptionProperties.CountDetails)); err != nil {
            return fmt.Errorf("Error setting `count_details`: %+v", err)
        }
        d.Set("created_at", (subscriptionProperties.CreatedAt).String())
        d.Set("dead_lettering_on_filter_evaluation_exceptions", subscriptionProperties.DeadLetteringOnFilterEvaluationExceptions)
        d.Set("dead_lettering_on_message_expiration", subscriptionProperties.DeadLetteringOnMessageExpiration)
        d.Set("default_message_time_to_live", subscriptionProperties.DefaultMessageTimeToLive)
        d.Set("enable_batched_operations", subscriptionProperties.EnableBatchedOperations)
        d.Set("entity_availability_status", string(subscriptionProperties.EntityAvailabilityStatus))
        d.Set("is_read_only", subscriptionProperties.IsReadOnly)
        d.Set("lock_duration", subscriptionProperties.LockDuration)
        d.Set("max_delivery_count", int(*subscriptionProperties.MaxDeliveryCount))
        d.Set("message_count", int(*subscriptionProperties.MessageCount))
        d.Set("requires_session", subscriptionProperties.RequiresSession)
        d.Set("status", string(subscriptionProperties.Status))
        d.Set("updated_at", (subscriptionProperties.UpdatedAt).String())
    }
    d.Set("namespace_name", namespaceName)
    d.Set("topic_name", topicName)
    d.Set("type", resp.Type)

    return nil
}


func resourceArmSubscriptionDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).subscriptionsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    namespaceName := id.Path["namespaces"]
    topicName := id.Path["topics"]
    name := id.Path["subscriptions"]

    if _, err := client.Delete(ctx, resourceGroup, namespaceName, topicName, name); err != nil {
        return fmt.Errorf("Error deleting Subscription %q (Topic Name %q / Namespace Name %q / Resource Group %q): %+v", name, topicName, namespaceName, resourceGroup, err)
    }

    return nil
}


func flattenArmSubscriptionMessageCountDetails(input *servicebus.MessageCountDetails) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if activeMessageCount := input.ActiveMessageCount; activeMessageCount != nil {
        result["active_message_count"] = int(*activeMessageCount)
    }
    if deadLetterMessageCount := input.DeadLetterMessageCount; deadLetterMessageCount != nil {
        result["dead_letter_message_count"] = int(*deadLetterMessageCount)
    }
    if scheduledMessageCount := input.ScheduledMessageCount; scheduledMessageCount != nil {
        result["scheduled_message_count"] = int(*scheduledMessageCount)
    }
    if transferDeadLetterMessageCount := input.TransferDeadLetterMessageCount; transferDeadLetterMessageCount != nil {
        result["transfer_dead_letter_message_count"] = int(*transferDeadLetterMessageCount)
    }
    if transferMessageCount := input.TransferMessageCount; transferMessageCount != nil {
        result["transfer_message_count"] = int(*transferMessageCount)
    }

    return []interface{}{result}
}
