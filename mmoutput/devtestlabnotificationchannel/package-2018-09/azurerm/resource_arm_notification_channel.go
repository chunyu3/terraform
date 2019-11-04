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



func resourceArmNotificationChannel() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmNotificationChannelCreate,
        Read: resourceArmNotificationChannelRead,
        Update: resourceArmNotificationChannelUpdate,
        Delete: resourceArmNotificationChannelDelete,

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

            "location": azure.SchemaLocation(),

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "lab_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "description": {
                Type: schema.TypeString,
                Optional: true,
            },

            "email_recipient": {
                Type: schema.TypeString,
                Optional: true,
            },

            "events": {
                Type: schema.TypeList,
                Optional: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "event_name": {
                            Type: schema.TypeString,
                            Optional: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(devtestlab.AutoShutdown),
                                string(devtestlab.Cost),
                            }, false),
                            Default: string(devtestlab.AutoShutdown),
                        },
                    },
                },
            },

            "notification_locale": {
                Type: schema.TypeString,
                Optional: true,
            },

            "web_hook_url": {
                Type: schema.TypeString,
                Optional: true,
            },

            "created_date": {
                Type: schema.TypeString,
                Computed: true,
            },

            "provisioning_state": {
                Type: schema.TypeString,
                Computed: true,
            },

            "type": {
                Type: schema.TypeString,
                Computed: true,
            },

            "unique_identifier": {
                Type: schema.TypeString,
                Computed: true,
            },

            "tags": tags.Schema(),
        },
    }
}

func resourceArmNotificationChannelCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).notificationChannelsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    labName := d.Get("lab_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, labName, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Notification Channel %q (Lab Name %q / Resource Group %q): %+v", name, labName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_notification_channel", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    description := d.Get("description").(string)
    emailRecipient := d.Get("email_recipient").(string)
    events := d.Get("events").([]interface{})
    notificationLocale := d.Get("notification_locale").(string)
    webHookUrl := d.Get("web_hook_url").(string)
    t := d.Get("tags").(map[string]interface{})

    notificationChannel := devtestlab.NotificationChannel{
        Location: utils.String(location),
        NotificationChannelProperties: &devtestlab.NotificationChannelProperties{
            Description: utils.String(description),
            EmailRecipient: utils.String(emailRecipient),
            Events: expandArmNotificationChannelEvent(events),
            NotificationLocale: utils.String(notificationLocale),
            WebHookURL: utils.String(webHookUrl),
        },
        Tags: tags.Expand(t),
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroup, labName, name, notificationChannel); err != nil {
        return fmt.Errorf("Error creating Notification Channel %q (Lab Name %q / Resource Group %q): %+v", name, labName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, labName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Notification Channel %q (Lab Name %q / Resource Group %q): %+v", name, labName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Notification Channel %q (Lab Name %q / Resource Group %q) ID", name, labName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmNotificationChannelRead(d, meta)
}

func resourceArmNotificationChannelRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).notificationChannelsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    labName := id.Path["labs"]
    name := id.Path["notificationchannels"]

    resp, err := client.Get(ctx, resourceGroup, labName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Notification Channel %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Notification Channel %q (Lab Name %q / Resource Group %q): %+v", name, labName, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if notificationChannelProperties := resp.NotificationChannelProperties; notificationChannelProperties != nil {
        d.Set("created_date", (notificationChannelProperties.CreatedDate).String())
        d.Set("description", notificationChannelProperties.Description)
        d.Set("email_recipient", notificationChannelProperties.EmailRecipient)
        if err := d.Set("events", flattenArmNotificationChannelEvent(notificationChannelProperties.Events)); err != nil {
            return fmt.Errorf("Error setting `events`: %+v", err)
        }
        d.Set("notification_locale", notificationChannelProperties.NotificationLocale)
        d.Set("provisioning_state", notificationChannelProperties.ProvisioningState)
        d.Set("unique_identifier", notificationChannelProperties.UniqueIdentifier)
        d.Set("web_hook_url", notificationChannelProperties.WebHookURL)
    }
    d.Set("lab_name", labName)
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}

func resourceArmNotificationChannelUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).notificationChannelsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    description := d.Get("description").(string)
    emailRecipient := d.Get("email_recipient").(string)
    events := d.Get("events").([]interface{})
    labName := d.Get("lab_name").(string)
    notificationLocale := d.Get("notification_locale").(string)
    webHookUrl := d.Get("web_hook_url").(string)
    t := d.Get("tags").(map[string]interface{})

    notificationChannel := devtestlab.NotificationChannel{
        Location: utils.String(location),
        NotificationChannelProperties: &devtestlab.NotificationChannelProperties{
            Description: utils.String(description),
            EmailRecipient: utils.String(emailRecipient),
            Events: expandArmNotificationChannelEvent(events),
            NotificationLocale: utils.String(notificationLocale),
            WebHookURL: utils.String(webHookUrl),
        },
        Tags: tags.Expand(t),
    }


    if _, err := client.Update(ctx, resourceGroup, labName, name, notificationChannel); err != nil {
        return fmt.Errorf("Error updating Notification Channel %q (Lab Name %q / Resource Group %q): %+v", name, labName, resourceGroup, err)
    }

    return resourceArmNotificationChannelRead(d, meta)
}

func resourceArmNotificationChannelDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).notificationChannelsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    labName := id.Path["labs"]
    name := id.Path["notificationchannels"]

    if _, err := client.Delete(ctx, resourceGroup, labName, name); err != nil {
        return fmt.Errorf("Error deleting Notification Channel %q (Lab Name %q / Resource Group %q): %+v", name, labName, resourceGroup, err)
    }

    return nil
}

func expandArmNotificationChannelEvent(input []interface{}) *[]devtestlab.Event {
    results := make([]devtestlab.Event, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        eventName := v["event_name"].(string)

        result := devtestlab.Event{
            EventName: devtestlab.NotificationChannelEventType(eventName),
        }

        results = append(results, result)
    }
    return &results
}


func flattenArmNotificationChannelEvent(input *[]devtestlab.Event) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        v["event_name"] = string(item.EventName)

        results = append(results, v)
    }

    return results
}