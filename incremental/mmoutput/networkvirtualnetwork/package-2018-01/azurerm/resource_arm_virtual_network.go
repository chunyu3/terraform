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



func resourceArmVirtualNetwork() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmVirtualNetworkCreateUpdate,
        Read: resourceArmVirtualNetworkRead,
        Update: resourceArmVirtualNetworkCreateUpdate,
        Delete: resourceArmVirtualNetworkDelete,

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

            "address_space": {
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

            "dhcp_options": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "dns_servers": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                    },
                },
            },

            "enable_ddos_protection": {
                Type: schema.TypeBool,
                Optional: true,
            },

            "enable_vm_protection": {
                Type: schema.TypeBool,
                Optional: true,
            },

            "etag": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
            },

            "resource_guid": {
                Type: schema.TypeString,
                Optional: true,
            },

            "subnets": {
                Type: schema.TypeList,
                Optional: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "address_prefix": {
                            Type: schema.TypeString,
                            Optional: true,
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
                        "network_security_group": {
                            Type: schema.TypeList,
                            Optional: true,
                            MaxItems: 1,
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
                                    "location": azure.SchemaLocation(),
                                    "tags": tags.Schema(),
                                },
                            },
                        },
                        "resource_navigation_links": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
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
                        "route_table": {
                            Type: schema.TypeList,
                            Optional: true,
                            MaxItems: 1,
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
                                    "location": azure.SchemaLocation(),
                                    "tags": tags.Schema(),
                                },
                            },
                        },
                        "service_endpoints": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "locations": {
                                        Type: schema.TypeList,
                                        Optional: true,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                        },
                                    },
                                    "service": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                },
                            },
                        },
                    },
                },
            },

            "virtual_network_peerings": {
                Type: schema.TypeList,
                Optional: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "allow_forwarded_traffic": {
                            Type: schema.TypeBool,
                            Optional: true,
                        },
                        "allow_gateway_transit": {
                            Type: schema.TypeBool,
                            Optional: true,
                        },
                        "allow_virtual_network_access": {
                            Type: schema.TypeBool,
                            Optional: true,
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
                        "peering_state": {
                            Type: schema.TypeString,
                            Optional: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(network.Initiated),
                                string(network.Connected),
                                string(network.Disconnected),
                            }, false),
                            Default: string(network.Initiated),
                        },
                        "remote_address_space": {
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
                        "remote_virtual_network": {
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
                        "use_remote_gateways": {
                            Type: schema.TypeBool,
                            Optional: true,
                        },
                    },
                },
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

func resourceArmVirtualNetworkCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).virtualNetworksClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Virtual Network %q (Resource Group %q): %+v", name, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_virtual_network", *existing.ID)
        }
    }

    id := d.Get("id").(string)
    location := azure.NormalizeLocation(d.Get("location").(string))
    addressSpace := d.Get("address_space").([]interface{})
    dhcpOptions := d.Get("dhcp_options").([]interface{})
    enableDdosProtection := d.Get("enable_ddos_protection").(bool)
    enableVmProtection := d.Get("enable_vm_protection").(bool)
    etag := d.Get("etag").(string)
    resourceGuid := d.Get("resource_guid").(string)
    subnets := d.Get("subnets").([]interface{})
    virtualNetworkPeerings := d.Get("virtual_network_peerings").([]interface{})
    t := d.Get("tags").(map[string]interface{})

    parameters := network.VirtualNetwork{
        Etag: utils.String(etag),
        ID: utils.String(id),
        Location: utils.String(location),
        VirtualNetworkPropertiesFormat: &network.VirtualNetworkPropertiesFormat{
            AddressSpace: expandArmVirtualNetworkAddressSpace(addressSpace),
            DhcpOptions: expandArmVirtualNetworkDhcpOptions(dhcpOptions),
            EnableDdosProtection: utils.Bool(enableDdosProtection),
            EnableVmProtection: utils.Bool(enableVmProtection),
            ResourceGuid: utils.String(resourceGuid),
            Subnets: expandArmVirtualNetworkSubnet(subnets),
            VirtualNetworkPeerings: expandArmVirtualNetworkVirtualNetworkPeering(virtualNetworkPeerings),
        },
        Tags: tags.Expand(t),
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, name, parameters)
    if err != nil {
        return fmt.Errorf("Error creating Virtual Network %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Virtual Network %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Virtual Network %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Virtual Network %q (Resource Group %q) ID", name, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmVirtualNetworkRead(d, meta)
}

func resourceArmVirtualNetworkRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).virtualNetworksClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["virtualNetworks"]

    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Virtual Network %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Virtual Network %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if virtualNetworkPropertiesFormat := resp.VirtualNetworkPropertiesFormat; virtualNetworkPropertiesFormat != nil {
        if err := d.Set("address_space", flattenArmVirtualNetworkAddressSpace(virtualNetworkPropertiesFormat.AddressSpace)); err != nil {
            return fmt.Errorf("Error setting `address_space`: %+v", err)
        }
        if err := d.Set("dhcp_options", flattenArmVirtualNetworkDhcpOptions(virtualNetworkPropertiesFormat.DhcpOptions)); err != nil {
            return fmt.Errorf("Error setting `dhcp_options`: %+v", err)
        }
        d.Set("enable_ddos_protection", virtualNetworkPropertiesFormat.EnableDdosProtection)
        d.Set("enable_vm_protection", virtualNetworkPropertiesFormat.EnableVmProtection)
        d.Set("provisioning_state", virtualNetworkPropertiesFormat.ProvisioningState)
        d.Set("resource_guid", virtualNetworkPropertiesFormat.ResourceGuid)
        if err := d.Set("subnets", flattenArmVirtualNetworkSubnet(virtualNetworkPropertiesFormat.Subnets)); err != nil {
            return fmt.Errorf("Error setting `subnets`: %+v", err)
        }
        if err := d.Set("virtual_network_peerings", flattenArmVirtualNetworkVirtualNetworkPeering(virtualNetworkPropertiesFormat.VirtualNetworkPeerings)); err != nil {
            return fmt.Errorf("Error setting `virtual_network_peerings`: %+v", err)
        }
    }
    d.Set("etag", resp.Etag)
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}


func resourceArmVirtualNetworkDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).virtualNetworksClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["virtualNetworks"]

    future, err := client.Delete(ctx, resourceGroup, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Virtual Network %q (Resource Group %q): %+v", name, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Virtual Network %q (Resource Group %q): %+v", name, resourceGroup, err)
        }
    }

    return nil
}

func expandArmVirtualNetworkAddressSpace(input []interface{}) *network.AddressSpace {
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

func expandArmVirtualNetworkDhcpOptions(input []interface{}) *network.DhcpOptions {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    dnsServers := v["dns_servers"].([]interface{})

    result := network.DhcpOptions{
        DnsServers: utils.ExpandStringSlice(dnsServers),
    }
    return &result
}

func expandArmVirtualNetworkSubnet(input []interface{}) *[]network.Subnet {
    results := make([]network.Subnet, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        id := v["id"].(string)
        addressPrefix := v["address_prefix"].(string)
        networkSecurityGroup := v["network_security_group"].([]interface{})
        routeTable := v["route_table"].([]interface{})
        serviceEndpoints := v["service_endpoints"].([]interface{})
        resourceNavigationLinks := v["resource_navigation_links"].([]interface{})
        name := v["name"].(string)
        etag := v["etag"].(string)

        result := network.Subnet{
            Etag: utils.String(etag),
            ID: utils.String(id),
            Name: utils.String(name),
            SubnetPropertiesFormat: &network.SubnetPropertiesFormat{
                AddressPrefix: utils.String(addressPrefix),
                NetworkSecurityGroup: expandArmVirtualNetworkSecurityGroup(networkSecurityGroup),
                ResourceNavigationLinks: expandArmVirtualNetworkResourceNavigationLink(resourceNavigationLinks),
                RouteTable: expandArmVirtualNetworkRouteTable(routeTable),
                ServiceEndpoints: expandArmVirtualNetworkServiceEndpointPropertiesFormat(serviceEndpoints),
            },
        }

        results = append(results, result)
    }
    return &results
}

func expandArmVirtualNetworkVirtualNetworkPeering(input []interface{}) *[]network.VirtualNetworkPeering {
    results := make([]network.VirtualNetworkPeering, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        id := v["id"].(string)
        allowVirtualNetworkAccess := v["allow_virtual_network_access"].(bool)
        allowForwardedTraffic := v["allow_forwarded_traffic"].(bool)
        allowGatewayTransit := v["allow_gateway_transit"].(bool)
        useRemoteGateways := v["use_remote_gateways"].(bool)
        remoteVirtualNetwork := v["remote_virtual_network"].([]interface{})
        remoteAddressSpace := v["remote_address_space"].([]interface{})
        peeringState := v["peering_state"].(string)
        name := v["name"].(string)
        etag := v["etag"].(string)

        result := network.VirtualNetworkPeering{
            Etag: utils.String(etag),
            ID: utils.String(id),
            Name: utils.String(name),
            VirtualNetworkPeeringPropertiesFormat: &network.VirtualNetworkPeeringPropertiesFormat{
                AllowForwardedTraffic: utils.Bool(allowForwardedTraffic),
                AllowGatewayTransit: utils.Bool(allowGatewayTransit),
                AllowVirtualNetworkAccess: utils.Bool(allowVirtualNetworkAccess),
                PeeringState: network.VirtualNetworkPeeringState(peeringState),
                RemoteAddressSpace: expandArmVirtualNetworkAddressSpace(remoteAddressSpace),
                RemoteVirtualNetwork: expandArmVirtualNetworkSubResource(remoteVirtualNetwork),
                UseRemoteGateways: utils.Bool(useRemoteGateways),
            },
        }

        results = append(results, result)
    }
    return &results
}

func expandArmVirtualNetworkSecurityGroup(input []interface{}) *network.SecurityGroup {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    id := v["id"].(string)
    location := azure.NormalizeLocation(v["location"].(string))
    t := v["tags"].(map[string]interface{})
    etag := v["etag"].(string)

    result := network.SecurityGroup{
        Etag: utils.String(etag),
        ID: utils.String(id),
        Location: utils.String(location),
        Tags: tags.Expand(t),
    }
    return &result
}

func expandArmVirtualNetworkResourceNavigationLink(input []interface{}) *[]network.ResourceNavigationLink {
    results := make([]network.ResourceNavigationLink, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        id := v["id"].(string)
        name := v["name"].(string)

        result := network.ResourceNavigationLink{
            ID: utils.String(id),
            Name: utils.String(name),
        }

        results = append(results, result)
    }
    return &results
}

func expandArmVirtualNetworkRouteTable(input []interface{}) *network.RouteTable {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    id := v["id"].(string)
    location := azure.NormalizeLocation(v["location"].(string))
    t := v["tags"].(map[string]interface{})
    etag := v["etag"].(string)

    result := network.RouteTable{
        Etag: utils.String(etag),
        ID: utils.String(id),
        Location: utils.String(location),
        Tags: tags.Expand(t),
    }
    return &result
}

func expandArmVirtualNetworkServiceEndpointPropertiesFormat(input []interface{}) *[]network.ServiceEndpointPropertiesFormat {
    results := make([]network.ServiceEndpointPropertiesFormat, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        service := v["service"].(string)
        locations := v["locations"].([]interface{})

        result := network.ServiceEndpointPropertiesFormat{
            Locations: utils.ExpandStringSlice(locations),
            Service: utils.String(service),
        }

        results = append(results, result)
    }
    return &results
}

func expandArmVirtualNetworkSubResource(input []interface{}) *network.SubResource {
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


func flattenArmVirtualNetworkAddressSpace(input *network.AddressSpace) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    result["address_prefixes"] = utils.FlattenStringSlice(input.AddressPrefixes)

    return []interface{}{result}
}

func flattenArmVirtualNetworkDhcpOptions(input *network.DhcpOptions) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    result["dns_servers"] = utils.FlattenStringSlice(input.DnsServers)

    return []interface{}{result}
}

func flattenArmVirtualNetworkSubnet(input *[]network.Subnet) []interface{} {
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
        if subnetPropertiesFormat := item.SubnetPropertiesFormat; subnetPropertiesFormat != nil {
            if addressPrefix := subnetPropertiesFormat.AddressPrefix; addressPrefix != nil {
                v["address_prefix"] = *addressPrefix
            }
            v["network_security_group"] = flattenArmVirtualNetworkSecurityGroup(subnetPropertiesFormat.NetworkSecurityGroup)
            v["resource_navigation_links"] = flattenArmVirtualNetworkResourceNavigationLink(subnetPropertiesFormat.ResourceNavigationLinks)
            v["route_table"] = flattenArmVirtualNetworkRouteTable(subnetPropertiesFormat.RouteTable)
            v["service_endpoints"] = flattenArmVirtualNetworkServiceEndpointPropertiesFormat(subnetPropertiesFormat.ServiceEndpoints)
        }
        if etag := item.Etag; etag != nil {
            v["etag"] = *etag
        }

        results = append(results, v)
    }

    return results
}

func flattenArmVirtualNetworkVirtualNetworkPeering(input *[]network.VirtualNetworkPeering) []interface{} {
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
        if virtualNetworkPeeringPropertiesFormat := item.VirtualNetworkPeeringPropertiesFormat; virtualNetworkPeeringPropertiesFormat != nil {
            if allowForwardedTraffic := virtualNetworkPeeringPropertiesFormat.AllowForwardedTraffic; allowForwardedTraffic != nil {
                v["allow_forwarded_traffic"] = *allowForwardedTraffic
            }
            if allowGatewayTransit := virtualNetworkPeeringPropertiesFormat.AllowGatewayTransit; allowGatewayTransit != nil {
                v["allow_gateway_transit"] = *allowGatewayTransit
            }
            if allowVirtualNetworkAccess := virtualNetworkPeeringPropertiesFormat.AllowVirtualNetworkAccess; allowVirtualNetworkAccess != nil {
                v["allow_virtual_network_access"] = *allowVirtualNetworkAccess
            }
            v["peering_state"] = string(virtualNetworkPeeringPropertiesFormat.PeeringState)
            v["remote_address_space"] = flattenArmVirtualNetworkAddressSpace(virtualNetworkPeeringPropertiesFormat.RemoteAddressSpace)
            v["remote_virtual_network"] = flattenArmVirtualNetworkSubResource(virtualNetworkPeeringPropertiesFormat.RemoteVirtualNetwork)
            if useRemoteGateways := virtualNetworkPeeringPropertiesFormat.UseRemoteGateways; useRemoteGateways != nil {
                v["use_remote_gateways"] = *useRemoteGateways
            }
        }
        if etag := item.Etag; etag != nil {
            v["etag"] = *etag
        }

        results = append(results, v)
    }

    return results
}

func flattenArmVirtualNetworkSecurityGroup(input *network.SecurityGroup) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if id := input.ID; id != nil {
        result["id"] = *id
    }
    if location := input.Location; location != nil {
        result["location"] = azure.NormalizeLocation(*location)
    }
    if etag := input.Etag; etag != nil {
        result["etag"] = *etag
    }

    return []interface{}{result}
}

func flattenArmVirtualNetworkResourceNavigationLink(input *[]network.ResourceNavigationLink) []interface{} {
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

        results = append(results, v)
    }

    return results
}

func flattenArmVirtualNetworkRouteTable(input *network.RouteTable) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if id := input.ID; id != nil {
        result["id"] = *id
    }
    if location := input.Location; location != nil {
        result["location"] = azure.NormalizeLocation(*location)
    }
    if etag := input.Etag; etag != nil {
        result["etag"] = *etag
    }

    return []interface{}{result}
}

func flattenArmVirtualNetworkServiceEndpointPropertiesFormat(input *[]network.ServiceEndpointPropertiesFormat) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        v["locations"] = utils.FlattenStringSlice(item.Locations)
        if service := item.Service; service != nil {
            v["service"] = *service
        }

        results = append(results, v)
    }

    return results
}

func flattenArmVirtualNetworkSubResource(input *network.SubResource) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if id := input.ID; id != nil {
        result["id"] = *id
    }

    return []interface{}{result}
}
