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



func resourceArmVirtualNetworkGateway() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmVirtualNetworkGatewayCreateUpdate,
        Read: resourceArmVirtualNetworkGatewayRead,
        Update: resourceArmVirtualNetworkGatewayCreateUpdate,
        Delete: resourceArmVirtualNetworkGatewayDelete,

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

            "active_active": {
                Type: schema.TypeBool,
                Optional: true,
            },

            "bgp_settings": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "asn": {
                            Type: schema.TypeInt,
                            Optional: true,
                        },
                        "bgp_peering_address": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "peer_weight": {
                            Type: schema.TypeInt,
                            Optional: true,
                        },
                    },
                },
            },

            "enable_bgp": {
                Type: schema.TypeBool,
                Optional: true,
            },

            "etag": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
            },

            "gateway_default_site": {
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

            "gateway_type": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(network.Vpn),
                    string(network.ExpressRoute),
                }, false),
                Default: string(network.Vpn),
            },

            "ip_configurations": {
                Type: schema.TypeList,
                Optional: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "etag": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "id": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "name": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "private_ipallocation_method": {
                            Type: schema.TypeString,
                            Optional: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(network.Static),
                                string(network.Dynamic),
                            }, false),
                            Default: string(network.Static),
                        },
                        "public_ip_address": {
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
                        "subnet": {
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
                    },
                },
            },

            "resource_guid": {
                Type: schema.TypeString,
                Optional: true,
            },

            "sku": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "capacity": {
                            Type: schema.TypeInt,
                            Optional: true,
                        },
                        "name": {
                            Type: schema.TypeString,
                            Optional: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(network.Basic),
                                string(network.HighPerformance),
                                string(network.Standard),
                                string(network.UltraPerformance),
                                string(network.VpnGw1),
                                string(network.VpnGw2),
                                string(network.VpnGw3),
                            }, false),
                            Default: string(network.Basic),
                        },
                        "tier": {
                            Type: schema.TypeString,
                            Optional: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(network.Basic),
                                string(network.HighPerformance),
                                string(network.Standard),
                                string(network.UltraPerformance),
                                string(network.VpnGw1),
                                string(network.VpnGw2),
                                string(network.VpnGw3),
                            }, false),
                            Default: string(network.Basic),
                        },
                    },
                },
            },

            "vpn_client_configuration": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
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
                        "vpn_client_revoked_certificates": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "etag": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                    "id": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                    "name": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                    "thumbprint": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                },
                            },
                        },
                        "vpn_client_root_certificates": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "public_cert_data": {
                                        Type: schema.TypeString,
                                        Required: true,
                                        ValidateFunc: validate.NoEmptyStrings,
                                    },
                                    "etag": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                    "id": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                    "name": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                },
                            },
                        },
                    },
                },
            },

            "vpn_type": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(network.PolicyBased),
                    string(network.RouteBased),
                }, false),
                Default: string(network.PolicyBased),
            },

            "provisioning_state": {
                Type: schema.TypeString,
                Computed: true,
            },

            "type": {
                Type: schema.TypeString,
                Computed: true,
            },

            "tags": tags.Schema(),
        },
    }
}

func resourceArmVirtualNetworkGatewayCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).virtualNetworkGatewaysClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Virtual Network Gateway %q (Resource Group %q): %+v", name, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_virtual_network_gateway", *existing.ID)
        }
    }

    id := d.Get("id").(string)
    location := azure.NormalizeLocation(d.Get("location").(string))
    activeActive := d.Get("active_active").(bool)
    bgpSettings := d.Get("bgp_settings").([]interface{})
    enableBgp := d.Get("enable_bgp").(bool)
    etag := d.Get("etag").(string)
    gatewayDefaultSite := d.Get("gateway_default_site").([]interface{})
    gatewayType := d.Get("gateway_type").(string)
    ipConfigurations := d.Get("ip_configurations").([]interface{})
    resourceGuid := d.Get("resource_guid").(string)
    sku := d.Get("sku").([]interface{})
    vpnClientConfiguration := d.Get("vpn_client_configuration").([]interface{})
    vpnType := d.Get("vpn_type").(string)
    t := d.Get("tags").(map[string]interface{})

    parameters := network.VirtualNetworkGateway{
        Etag: utils.String(etag),
        ID: utils.String(id),
        Location: utils.String(location),
        VirtualNetworkGatewayPropertiesFormat: &network.VirtualNetworkGatewayPropertiesFormat{
            ActiveActive: utils.Bool(activeActive),
            BgpSettings: expandArmVirtualNetworkGatewayBgpSettings(bgpSettings),
            EnableBgp: utils.Bool(enableBgp),
            GatewayDefaultSite: expandArmVirtualNetworkGatewaySubResource(gatewayDefaultSite),
            GatewayType: network.VirtualNetworkGatewayType(gatewayType),
            IpConfigurations: expandArmVirtualNetworkGatewayVirtualNetworkGatewayIPConfiguration(ipConfigurations),
            ResourceGuid: utils.String(resourceGuid),
            Sku: expandArmVirtualNetworkGatewayVirtualNetworkGatewaySku(sku),
            VpnClientConfiguration: expandArmVirtualNetworkGatewayVpnClientConfiguration(vpnClientConfiguration),
            VpnType: network.VpnType(vpnType),
        },
        Tags: tags.Expand(t),
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, name, parameters)
    if err != nil {
        return fmt.Errorf("Error creating Virtual Network Gateway %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Virtual Network Gateway %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Virtual Network Gateway %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Virtual Network Gateway %q (Resource Group %q) ID", name, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmVirtualNetworkGatewayRead(d, meta)
}

func resourceArmVirtualNetworkGatewayRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).virtualNetworkGatewaysClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["virtualNetworkGateways"]

    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Virtual Network Gateway %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Virtual Network Gateway %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if virtualNetworkGatewayPropertiesFormat := resp.VirtualNetworkGatewayPropertiesFormat; virtualNetworkGatewayPropertiesFormat != nil {
        d.Set("active_active", virtualNetworkGatewayPropertiesFormat.ActiveActive)
        if err := d.Set("bgp_settings", flattenArmVirtualNetworkGatewayBgpSettings(virtualNetworkGatewayPropertiesFormat.BgpSettings)); err != nil {
            return fmt.Errorf("Error setting `bgp_settings`: %+v", err)
        }
        d.Set("enable_bgp", virtualNetworkGatewayPropertiesFormat.EnableBgp)
        if err := d.Set("gateway_default_site", flattenArmVirtualNetworkGatewaySubResource(virtualNetworkGatewayPropertiesFormat.GatewayDefaultSite)); err != nil {
            return fmt.Errorf("Error setting `gateway_default_site`: %+v", err)
        }
        d.Set("gateway_type", string(virtualNetworkGatewayPropertiesFormat.GatewayType))
        if err := d.Set("ip_configurations", flattenArmVirtualNetworkGatewayVirtualNetworkGatewayIPConfiguration(virtualNetworkGatewayPropertiesFormat.IpConfigurations)); err != nil {
            return fmt.Errorf("Error setting `ip_configurations`: %+v", err)
        }
        d.Set("provisioning_state", virtualNetworkGatewayPropertiesFormat.ProvisioningState)
        d.Set("resource_guid", virtualNetworkGatewayPropertiesFormat.ResourceGuid)
        if err := d.Set("sku", flattenArmVirtualNetworkGatewayVirtualNetworkGatewaySku(virtualNetworkGatewayPropertiesFormat.Sku)); err != nil {
            return fmt.Errorf("Error setting `sku`: %+v", err)
        }
        if err := d.Set("vpn_client_configuration", flattenArmVirtualNetworkGatewayVpnClientConfiguration(virtualNetworkGatewayPropertiesFormat.VpnClientConfiguration)); err != nil {
            return fmt.Errorf("Error setting `vpn_client_configuration`: %+v", err)
        }
        d.Set("vpn_type", string(virtualNetworkGatewayPropertiesFormat.VpnType))
    }
    d.Set("etag", resp.Etag)
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}


func resourceArmVirtualNetworkGatewayDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).virtualNetworkGatewaysClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["virtualNetworkGateways"]

    future, err := client.Delete(ctx, resourceGroup, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Virtual Network Gateway %q (Resource Group %q): %+v", name, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Virtual Network Gateway %q (Resource Group %q): %+v", name, resourceGroup, err)
        }
    }

    return nil
}

func expandArmVirtualNetworkGatewayBgpSettings(input []interface{}) *network.BgpSettings {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    asn := v["asn"].(int)
    bgpPeeringAddress := v["bgp_peering_address"].(string)
    peerWeight := v["peer_weight"].(int)

    result := network.BgpSettings{
        Asn: utils.Int64(int64(asn)),
        BgpPeeringAddress: utils.String(bgpPeeringAddress),
        PeerWeight: utils.Int32(int32(peerWeight)),
    }
    return &result
}

func expandArmVirtualNetworkGatewaySubResource(input []interface{}) *network.SubResource {
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

func expandArmVirtualNetworkGatewayVirtualNetworkGatewayIPConfiguration(input []interface{}) *[]network.VirtualNetworkGatewayIPConfiguration {
    results := make([]network.VirtualNetworkGatewayIPConfiguration, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        id := v["id"].(string)
        privateIpallocationMethod := v["private_ipallocation_method"].(string)
        subnet := v["subnet"].([]interface{})
        publicIpAddress := v["public_ip_address"].([]interface{})
        name := v["name"].(string)
        etag := v["etag"].(string)

        result := network.VirtualNetworkGatewayIPConfiguration{
            Etag: utils.String(etag),
            ID: utils.String(id),
            Name: utils.String(name),
            VirtualNetworkGatewayIPConfigurationPropertiesFormat: &network.VirtualNetworkGatewayIPConfigurationPropertiesFormat{
                PrivateIpallocationMethod: network.IPAllocationMethod(privateIpallocationMethod),
                PublicIpAddress: expandArmVirtualNetworkGatewaySubResource(publicIpAddress),
                Subnet: expandArmVirtualNetworkGatewaySubResource(subnet),
            },
        }

        results = append(results, result)
    }
    return &results
}

func expandArmVirtualNetworkGatewayVirtualNetworkGatewaySku(input []interface{}) *network.VirtualNetworkGatewaySku {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    name := v["name"].(string)
    tier := v["tier"].(string)
    capacity := v["capacity"].(int)

    result := network.VirtualNetworkGatewaySku{
        Capacity: utils.Int32(int32(capacity)),
        Name: network.VirtualNetworkGatewaySkuName(name),
        Tier: network.VirtualNetworkGatewaySkuTier(tier),
    }
    return &result
}

func expandArmVirtualNetworkGatewayVpnClientConfiguration(input []interface{}) *network.VpnClientConfiguration {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    vpnClientAddressPool := v["vpn_client_address_pool"].([]interface{})
    vpnClientRootCertificates := v["vpn_client_root_certificates"].([]interface{})
    vpnClientRevokedCertificates := v["vpn_client_revoked_certificates"].([]interface{})

    result := network.VpnClientConfiguration{
        VpnClientAddressPool: expandArmVirtualNetworkGatewayAddressSpace(vpnClientAddressPool),
        VpnClientRevokedCertificates: expandArmVirtualNetworkGatewayVpnClientRevokedCertificate(vpnClientRevokedCertificates),
        VpnClientRootCertificates: expandArmVirtualNetworkGatewayVpnClientRootCertificate(vpnClientRootCertificates),
    }
    return &result
}

func expandArmVirtualNetworkGatewayAddressSpace(input []interface{}) *network.AddressSpace {
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

func expandArmVirtualNetworkGatewayVpnClientRevokedCertificate(input []interface{}) *[]network.VpnClientRevokedCertificate {
    results := make([]network.VpnClientRevokedCertificate, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        id := v["id"].(string)
        thumbprint := v["thumbprint"].(string)
        name := v["name"].(string)
        etag := v["etag"].(string)

        result := network.VpnClientRevokedCertificate{
            Etag: utils.String(etag),
            ID: utils.String(id),
            Name: utils.String(name),
            VpnClientRevokedCertificatePropertiesFormat: &network.VpnClientRevokedCertificatePropertiesFormat{
                Thumbprint: utils.String(thumbprint),
            },
        }

        results = append(results, result)
    }
    return &results
}

func expandArmVirtualNetworkGatewayVpnClientRootCertificate(input []interface{}) *[]network.VpnClientRootCertificate {
    results := make([]network.VpnClientRootCertificate, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        id := v["id"].(string)
        publicCertData := v["public_cert_data"].(string)
        name := v["name"].(string)
        etag := v["etag"].(string)

        result := network.VpnClientRootCertificate{
            Etag: utils.String(etag),
            ID: utils.String(id),
            Name: utils.String(name),
            VpnClientRootCertificatePropertiesFormat: &network.VpnClientRootCertificatePropertiesFormat{
                PublicCertData: utils.String(publicCertData),
            },
        }

        results = append(results, result)
    }
    return &results
}


func flattenArmVirtualNetworkGatewayBgpSettings(input *network.BgpSettings) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if asn := input.Asn; asn != nil {
        result["asn"] = int(*asn)
    }
    if bgpPeeringAddress := input.BgpPeeringAddress; bgpPeeringAddress != nil {
        result["bgp_peering_address"] = *bgpPeeringAddress
    }
    if peerWeight := input.PeerWeight; peerWeight != nil {
        result["peer_weight"] = int(*peerWeight)
    }

    return []interface{}{result}
}

func flattenArmVirtualNetworkGatewaySubResource(input *network.SubResource) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if id := input.ID; id != nil {
        result["id"] = *id
    }

    return []interface{}{result}
}

func flattenArmVirtualNetworkGatewayVirtualNetworkGatewayIPConfiguration(input *[]network.VirtualNetworkGatewayIPConfiguration) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        if id := item.ID; id != nil {
            v["id"] = *id
        }
        if name := item.Name; name != nil {
            v["name"] = *name
        }
        if etag := item.Etag; etag != nil {
            v["etag"] = *etag
        }
        if virtualNetworkGatewayIPConfigurationPropertiesFormat := item.VirtualNetworkGatewayIPConfigurationPropertiesFormat; virtualNetworkGatewayIPConfigurationPropertiesFormat != nil {
            v["private_ipallocation_method"] = string(virtualNetworkGatewayIPConfigurationPropertiesFormat.PrivateIpallocationMethod)
            v["public_ip_address"] = flattenArmVirtualNetworkGatewaySubResource(virtualNetworkGatewayIPConfigurationPropertiesFormat.PublicIpAddress)
            v["subnet"] = flattenArmVirtualNetworkGatewaySubResource(virtualNetworkGatewayIPConfigurationPropertiesFormat.Subnet)
        }

        results = append(results, v)
    }

    return results
}

func flattenArmVirtualNetworkGatewayVirtualNetworkGatewaySku(input *network.VirtualNetworkGatewaySku) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    result["name"] = string(input.Name)
    if capacity := input.Capacity; capacity != nil {
        result["capacity"] = int(*capacity)
    }
    result["tier"] = string(input.Tier)

    return []interface{}{result}
}

func flattenArmVirtualNetworkGatewayVpnClientConfiguration(input *network.VpnClientConfiguration) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    result["vpn_client_address_pool"] = flattenArmVirtualNetworkGatewayAddressSpace(input.VpnClientAddressPool)
    result["vpn_client_revoked_certificates"] = flattenArmVirtualNetworkGatewayVpnClientRevokedCertificate(input.VpnClientRevokedCertificates)
    result["vpn_client_root_certificates"] = flattenArmVirtualNetworkGatewayVpnClientRootCertificate(input.VpnClientRootCertificates)

    return []interface{}{result}
}

func flattenArmVirtualNetworkGatewayAddressSpace(input *network.AddressSpace) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    result["address_prefixes"] = utils.FlattenStringSlice(input.AddressPrefixes)

    return []interface{}{result}
}

func flattenArmVirtualNetworkGatewayVpnClientRevokedCertificate(input *[]network.VpnClientRevokedCertificate) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        if id := item.ID; id != nil {
            v["id"] = *id
        }
        if name := item.Name; name != nil {
            v["name"] = *name
        }
        if etag := item.Etag; etag != nil {
            v["etag"] = *etag
        }
        if vpnClientRevokedCertificatePropertiesFormat := item.VpnClientRevokedCertificatePropertiesFormat; vpnClientRevokedCertificatePropertiesFormat != nil {
            if thumbprint := vpnClientRevokedCertificatePropertiesFormat.Thumbprint; thumbprint != nil {
                v["thumbprint"] = *thumbprint
            }
        }

        results = append(results, v)
    }

    return results
}

func flattenArmVirtualNetworkGatewayVpnClientRootCertificate(input *[]network.VpnClientRootCertificate) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        if id := item.ID; id != nil {
            v["id"] = *id
        }
        if name := item.Name; name != nil {
            v["name"] = *name
        }
        if etag := item.Etag; etag != nil {
            v["etag"] = *etag
        }
        if vpnClientRootCertificatePropertiesFormat := item.VpnClientRootCertificatePropertiesFormat; vpnClientRootCertificatePropertiesFormat != nil {
            if publicCertData := vpnClientRootCertificatePropertiesFormat.PublicCertData; publicCertData != nil {
                v["public_cert_data"] = *publicCertData
            }
        }

        results = append(results, v)
    }

    return results
}
