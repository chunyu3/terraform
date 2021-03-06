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



func resourceArmLocalNetworkGateway() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmLocalNetworkGatewayCreateUpdate,
        Read: resourceArmLocalNetworkGatewayRead,
        Update: resourceArmLocalNetworkGatewayCreateUpdate,
        Delete: resourceArmLocalNetworkGatewayDelete,

        Importer: &schema.ResourceImporter{
            State: schema.ImportStatePassthrough,
        },


        Schema: map[string]*schema.Schema{
            "local_network_gateway_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "location": azure.SchemaLocation(),

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "etag": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
            },

            "gateway_ip_address": {
                Type: schema.TypeString,
                Optional: true,
            },

            "local_network_address_space": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "address_prefixes": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                    },
                },
            },

            "resource_guid": {
                Type: schema.TypeString,
                Optional: true,
            },

            "tags": tags.Schema(),

            "id": {
                Type: schema.TypeString,
                Computed: true,
            },

            "name": {
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
        },
    }
}

func resourceArmLocalNetworkGatewayCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).localNetworkGatewaysClient
    ctx, cancel := timeouts.ForCreateUpdate(meta.(*ArmClient).StopContext, d)
    defer cancel()

    resourceGroupName := d.Get("resource_group").(string)
    name := d.Get("local_network_gateway_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroupName, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Local Network Gateway (Local Network Gateway Name %q / Resource Group %q): %+v", name, resourceGroupName, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_local_network_gateway", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    etag := d.Get("etag").(string)
    gatewayIPAddress := d.Get("gateway_ip_address").(string)
    localNetworkAddressSpace := d.Get("local_network_address_space").([]interface{})
    resourceGUID := d.Get("resource_guid").(string)
    tags := d.Get("tags").(map[string]interface{})

    parameters := network.LocalNetworkGateway{
        Etag: utils.String(etag),
        Location: utils.String(location),
        LocalNetworkGatewayPropertiesFormat: &network.LocalNetworkGatewayPropertiesFormat{
            GatewayIPAddress: utils.String(gatewayIPAddress),
            LocalNetworkAddressSpace: expandArmLocalNetworkGatewayAddressSpace(localNetworkAddressSpace),
            ResourceGUID: utils.String(resourceGUID),
        },
        Tags: tags.Expand(tags),
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroupName, name, parameters)
    if err != nil {
        return fmt.Errorf("Error creating Local Network Gateway (Local Network Gateway Name %q / Resource Group %q): %+v", name, resourceGroupName, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Local Network Gateway (Local Network Gateway Name %q / Resource Group %q): %+v", name, resourceGroupName, err)
    }


    resp, err := client.Get(ctx, resourceGroupName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Local Network Gateway (Local Network Gateway Name %q / Resource Group %q): %+v", name, resourceGroupName, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Local Network Gateway (Local Network Gateway Name %q / Resource Group %q) ID", name, resourceGroupName)
    }
    d.SetId(*resp.ID)

    return resourceArmLocalNetworkGatewayRead(d, meta)
}

func resourceArmLocalNetworkGatewayRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).localNetworkGatewaysClient
    ctx, cancel := timeouts.ForRead(meta.(*ArmClient).StopContext, d)
    defer cancel()

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroupName := id.ResourceGroup
    name := id.Path["localNetworkGateways"]

    resp, err := client.Get(ctx, resourceGroupName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Local Network Gateway %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Local Network Gateway (Local Network Gateway Name %q / Resource Group %q): %+v", name, resourceGroupName, err)
    }


    d.Set("resource_group", resourceGroupName)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    d.Set("etag", resp.Etag)
    if localNetworkGatewayPropertiesFormat := resp.LocalNetworkGatewayPropertiesFormat; localNetworkGatewayPropertiesFormat != nil {
        d.Set("gateway_ip_address", localNetworkGatewayPropertiesFormat.GatewayIPAddress)
        if err := d.Set("local_network_address_space", flattenArmLocalNetworkGatewayAddressSpace(localNetworkGatewayPropertiesFormat.LocalNetworkAddressSpace)); err != nil {
            return fmt.Errorf("Error setting `local_network_address_space`: %+v", err)
        }
        d.Set("provisioning_state", localNetworkGatewayPropertiesFormat.ProvisioningState)
        d.Set("resource_guid", localNetworkGatewayPropertiesFormat.ResourceGUID)
    }
    d.Set("id", resp.ID)
    d.Set("local_network_gateway_name", name)
    d.Set("name", resp.Name)
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}


func resourceArmLocalNetworkGatewayDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).localNetworkGatewaysClient
    ctx, cancel := timeouts.ForDelete(meta.(*ArmClient).StopContext, d)
    defer cancel()


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroupName := id.ResourceGroup
    name := id.Path["localNetworkGateways"]

    future, err := client.Delete(ctx, resourceGroupName, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Local Network Gateway (Local Network Gateway Name %q / Resource Group %q): %+v", name, resourceGroupName, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Local Network Gateway (Local Network Gateway Name %q / Resource Group %q): %+v", name, resourceGroupName, err)
        }
    }

    return nil
}

func expandArmLocalNetworkGatewayAddressSpace(input []interface{}) *network.AddressSpace {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    addressPrefixes := v["address_prefixes"].([]interface{})

    result := network.AddressSpace{
        AddressPrefixes: utils.ExpandStringSlice(addressPrefixes),
    }
    return &result
}


func flattenArmLocalNetworkGatewayAddressSpace(input *network.AddressSpace) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    result["address_prefixes"] = utils.FlattenStringSlice(input.AddressPrefixes)

    return []interface{}{result}
}
