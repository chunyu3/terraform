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



func resourceArmP2sVpnGateway() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmP2sVpnGatewayCreateUpdate,
        Read: resourceArmP2sVpnGatewayRead,
        Update: resourceArmP2sVpnGatewayCreateUpdate,
        Delete: resourceArmP2sVpnGatewayDelete,

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

            "gateway_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "p2svpn_server_configuration": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "id": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },

            "virtual_hub": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "id": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },

            "vpn_client_address_pool": {
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

            "vpn_gateway_scale_unit": {
                Type: schema.TypeInt,
                Optional: true,
            },

            "etag": {
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

            "vpn_client_connection_health": {
                Type: schema.TypeList,
                Computed: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "allocated_ip_addresses": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                        "total_egress_bytes_transferred": {
                            Type: schema.TypeInt,
                            Optional: true,
                        },
                        "total_ingress_bytes_transferred": {
                            Type: schema.TypeInt,
                            Optional: true,
                        },
                        "vpn_client_connections_count": {
                            Type: schema.TypeInt,
                            Optional: true,
                        },
                    },
                },
            },

            "tags": tags.Schema(),
        },
    }
}

func resourceArmP2sVpnGatewayCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).p2sVpnGatewaysClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    gatewayName := d.Get("gateway_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, gatewayName)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing P2s Vpn Gateway (Gateway Name %q / Resource Group %q): %+v", gatewayName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_p2s_vpn_gateway", *existing.ID)
        }
    }

    id := d.Get("id").(string)
    location := azure.NormalizeLocation(d.Get("location").(string))
    p2svpnServerConfiguration := d.Get("p2svpn_server_configuration").([]interface{})
    virtualHub := d.Get("virtual_hub").([]interface{})
    vpnClientAddressPool := d.Get("vpn_client_address_pool").([]interface{})
    vpnGatewayScaleUnit := d.Get("vpn_gateway_scale_unit").(int)
    t := d.Get("tags").(map[string]interface{})

    p2svpnGatewayParameters := network.P2SVpnGateway{
        ID: utils.String(id),
        Location: utils.String(location),
        P2SVpnGatewayProperties: &network.P2SVpnGatewayProperties{
            P2svpnServerConfiguration: expandArmP2sVpnGatewaySubResource(p2svpnServerConfiguration),
            VirtualHub: expandArmP2sVpnGatewaySubResource(virtualHub),
            VpnClientAddressPool: expandArmP2sVpnGatewayAddressSpace(vpnClientAddressPool),
            VpnGatewayScaleUnit: utils.Int32(int32(vpnGatewayScaleUnit)),
        },
        Tags: tags.Expand(t),
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, gatewayName, p2svpnGatewayParameters)
    if err != nil {
        return fmt.Errorf("Error creating P2s Vpn Gateway (Gateway Name %q / Resource Group %q): %+v", gatewayName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of P2s Vpn Gateway (Gateway Name %q / Resource Group %q): %+v", gatewayName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, gatewayName)
    if err != nil {
        return fmt.Errorf("Error retrieving P2s Vpn Gateway (Gateway Name %q / Resource Group %q): %+v", gatewayName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read P2s Vpn Gateway (Gateway Name %q / Resource Group %q) ID", gatewayName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmP2sVpnGatewayRead(d, meta)
}

func resourceArmP2sVpnGatewayRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).p2sVpnGatewaysClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    gatewayName := id.Path["p2svpnGateways"]

    resp, err := client.Get(ctx, resourceGroup, gatewayName)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] P2s Vpn Gateway %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading P2s Vpn Gateway (Gateway Name %q / Resource Group %q): %+v", gatewayName, resourceGroup, err)
    }


    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    d.Set("etag", resp.Etag)
    d.Set("gateway_name", gatewayName)
    if p2SVpnGatewayProperties := resp.P2SVpnGatewayProperties; p2SVpnGatewayProperties != nil {
        if err := d.Set("p2svpn_server_configuration", flattenArmP2sVpnGatewaySubResource(p2SVpnGatewayProperties.P2svpnServerConfiguration)); err != nil {
            return fmt.Errorf("Error setting `p2svpn_server_configuration`: %+v", err)
        }
        d.Set("provisioning_state", string(p2SVpnGatewayProperties.ProvisioningState))
        if err := d.Set("virtual_hub", flattenArmP2sVpnGatewaySubResource(p2SVpnGatewayProperties.VirtualHub)); err != nil {
            return fmt.Errorf("Error setting `virtual_hub`: %+v", err)
        }
        if err := d.Set("vpn_client_address_pool", flattenArmP2sVpnGatewayAddressSpace(p2SVpnGatewayProperties.VpnClientAddressPool)); err != nil {
            return fmt.Errorf("Error setting `vpn_client_address_pool`: %+v", err)
        }
        if err := d.Set("vpn_client_connection_health", flattenArmP2sVpnGatewayVpnClientConnectionHealth(p2SVpnGatewayProperties.VpnClientConnectionHealth)); err != nil {
            return fmt.Errorf("Error setting `vpn_client_connection_health`: %+v", err)
        }
        d.Set("vpn_gateway_scale_unit", int(*p2SVpnGatewayProperties.VpnGatewayScaleUnit))
    }
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}


func resourceArmP2sVpnGatewayDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).p2sVpnGatewaysClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    gatewayName := id.Path["p2svpnGateways"]

    future, err := client.Delete(ctx, resourceGroup, gatewayName)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting P2s Vpn Gateway (Gateway Name %q / Resource Group %q): %+v", gatewayName, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting P2s Vpn Gateway (Gateway Name %q / Resource Group %q): %+v", gatewayName, resourceGroup, err)
        }
    }

    return nil
}

func expandArmP2sVpnGatewaySubResource(input []interface{}) *network.SubResource {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    id := v["id"].(string)

    result := network.SubResource{
        ID: utils.String(id),
    }
    return &result
}

func expandArmP2sVpnGatewayAddressSpace(input []interface{}) *network.AddressSpace {
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


func flattenArmP2sVpnGatewaySubResource(input *network.SubResource) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if id := input.ID; id != nil {
        result["id"] = *id
    }

    return []interface{}{result}
}

func flattenArmP2sVpnGatewayAddressSpace(input *network.AddressSpace) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    result["address_prefixes"] = utils.FlattenStringSlice(input.AddressPrefixes)

    return []interface{}{result}
}

func flattenArmP2sVpnGatewayVpnClientConnectionHealth(input *network.VpnClientConnectionHealth) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})


    return []interface{}{result}
}
