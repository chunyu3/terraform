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
        Create: resourceArmChannelCreateUpdate,
        Read: resourceArmChannelRead,
        Update: resourceArmChannelCreateUpdate,
        Delete: resourceArmChannelDelete,

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

            "account_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "channel_type": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "channel_functions": {
                Type: schema.TypeList,
                Optional: true,
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
            },

            "credentials": {
                Type: schema.TypeMap,
                Optional: true,
                Elem: &schema.Schema{Type: schema.TypeString},
            },

            "type": {
                Type: schema.TypeString,
                Computed: true,
            },
        },
    }
}

func resourceArmChannelCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).channelsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    accountName := d.Get("account_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, accountName, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Channel %q (Account Name %q / Resource Group %q): %+v", name, accountName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_channel", *existing.ID)
        }
    }

    channelFunctions := d.Get("channel_functions").([]interface{})
    channelType := d.Get("channel_type").(string)
    credentials := d.Get("credentials").(map[string]interface{})

    channel := engagementfabric.Channel{
        ChannelProperties: &engagementfabric.ChannelProperties{
            ChannelFunctions: utils.ExpandStringSlice(channelFunctions),
            ChannelType: utils.String(channelType),
            Credentials: utils.ExpandKeyValuePairs(credentials),
        },
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroup, accountName, name, channel); err != nil {
        return fmt.Errorf("Error creating Channel %q (Account Name %q / Resource Group %q): %+v", name, accountName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, accountName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Channel %q (Account Name %q / Resource Group %q): %+v", name, accountName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Channel %q (Account Name %q / Resource Group %q) ID", name, accountName, resourceGroup)
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
    accountName := id.Path["Accounts"]
    name := id.Path["Channels"]

    resp, err := client.Get(ctx, resourceGroup, accountName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Channel %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Channel %q (Account Name %q / Resource Group %q): %+v", name, accountName, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    d.Set("account_name", accountName)
    if channelProperties := resp.ChannelProperties; channelProperties != nil {
        d.Set("channel_functions", utils.FlattenStringSlice(channelProperties.ChannelFunctions))
        d.Set("channel_type", channelProperties.ChannelType)
        d.Set("credentials", utils.FlattenKeyValuePairs(channelProperties.Credentials))
    }
    d.Set("type", resp.Type)

    return nil
}


func resourceArmChannelDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).channelsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    accountName := id.Path["Accounts"]
    name := id.Path["Channels"]

    if _, err := client.Delete(ctx, resourceGroup, accountName, name); err != nil {
        return fmt.Errorf("Error deleting Channel %q (Account Name %q / Resource Group %q): %+v", name, accountName, resourceGroup, err)
    }

    return nil
}
