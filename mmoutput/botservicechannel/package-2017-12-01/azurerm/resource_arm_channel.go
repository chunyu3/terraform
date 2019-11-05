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



func resourceArmChannel() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmChannelCreate,
        Read: resourceArmChannelRead,
        Update: resourceArmChannelUpdate,
        Delete: resourceArmChannelDelete,

        Importer: &schema.ResourceImporter{
            State: schema.ImportStatePassthrough,
        },


        Schema: map[string]*schema.Schema{
            "name": {
                Type: schema.TypeString,
                Computed: true,
            },

            "location": azure.SchemaLocation(),

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "channel_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "resource_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "etag": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
            },

            "kind": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(botservice.sdk),
                    string(botservice.designer),
                    string(botservice.bot),
                    string(botservice.function),
                }, false),
                Default: string(botservice.sdk),
            },

            "sku": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "name": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(botservice.F0),
                                string(botservice.S1),
                            }, false),
                        },
                    },
                },
            },

            "type": {
                Type: schema.TypeString,
                Computed: true,
            },

            "tags": tags.Schema(),
        },
    }
}

func resourceArmChannelCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).channelsClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    channelName := d.Get("channel_name").(string)
    resourceName := d.Get("resource_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, resourceName, channelName)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Channel (Channel Name %q / Resource Name %q / Resource Group %q): %+v", channelName, resourceName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_channel", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    etag := d.Get("etag").(string)
    kind := d.Get("kind").(string)
    sku := d.Get("sku").([]interface{})
    t := d.Get("tags").(map[string]interface{})

    parameters := botservice.BotChannel{
        Etag: utils.String(etag),
        Kind: botservice.Kind(kind),
        Location: utils.String(location),
        Sku: expandArmChannelSku(sku),
        Tags: tags.Expand(t),
    }


    if _, err := client.Create(ctx, resourceGroup, resourceName, channelName, parameters); err != nil {
        return fmt.Errorf("Error creating Channel (Channel Name %q / Resource Name %q / Resource Group %q): %+v", channelName, resourceName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, resourceName, channelName)
    if err != nil {
        return fmt.Errorf("Error retrieving Channel (Channel Name %q / Resource Name %q / Resource Group %q): %+v", channelName, resourceName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Channel (Channel Name %q / Resource Name %q / Resource Group %q) ID", channelName, resourceName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmChannelRead(d, meta)
}

func resourceArmChannelRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).channelsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    resourceName := id.Path["botServices"]
    channelName := id.Path["channels"]

    resp, err := client.Get(ctx, resourceGroup, resourceName, channelName)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Channel %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Channel (Channel Name %q / Resource Name %q / Resource Group %q): %+v", channelName, resourceName, resourceGroup, err)
    }


    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    d.Set("channel_name", channelName)
    d.Set("etag", resp.Etag)
    d.Set("kind", string(resp.Kind))
    d.Set("resource_name", resourceName)
    if err := d.Set("sku", flattenArmChannelSku(resp.Sku)); err != nil {
        return fmt.Errorf("Error setting `sku`: %+v", err)
    }
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}

func resourceArmChannelUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).channelsClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    channelName := d.Get("channel_name").(string)
    etag := d.Get("etag").(string)
    kind := d.Get("kind").(string)
    resourceName := d.Get("resource_name").(string)
    sku := d.Get("sku").([]interface{})
    t := d.Get("tags").(map[string]interface{})

    parameters := botservice.BotChannel{
        Etag: utils.String(etag),
        Kind: botservice.Kind(kind),
        Location: utils.String(location),
        Sku: expandArmChannelSku(sku),
        Tags: tags.Expand(t),
    }


    if _, err := client.Update(ctx, resourceGroup, resourceName, channelName, parameters); err != nil {
        return fmt.Errorf("Error updating Channel (Channel Name %q / Resource Name %q / Resource Group %q): %+v", channelName, resourceName, resourceGroup, err)
    }

    return resourceArmChannelRead(d, meta)
}

func resourceArmChannelDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).channelsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    resourceName := id.Path["botServices"]
    channelName := id.Path["channels"]

    if _, err := client.Delete(ctx, resourceGroup, resourceName, channelName); err != nil {
        return fmt.Errorf("Error deleting Channel (Channel Name %q / Resource Name %q / Resource Group %q): %+v", channelName, resourceName, resourceGroup, err)
    }

    return nil
}

func expandArmChannelSku(input []interface{}) *botservice.Sku {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    name := v["name"].(string)

    result := botservice.Sku{
        Name: botservice.SkuName(name),
    }
    return &result
}


func flattenArmChannelSku(input *botservice.Sku) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    result["name"] = string(input.Name)

    return []interface{}{result}
}