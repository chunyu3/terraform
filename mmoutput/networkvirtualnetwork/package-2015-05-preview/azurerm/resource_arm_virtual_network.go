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
            "location": azure.SchemaLocation(),

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "virtual_network_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

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
                        "ip_configurations": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "id": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                },
                            },
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
                                    "id": {
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

func resourceArmVirtualNetworkCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).virtualNetworksClient
    ctx, cancel := timeouts.ForCreateUpdate(meta.(*ArmClient).StopContext, d)
    defer cancel()

    resourceGroupName := d.Get("resource_group").(string)
    name := d.Get("virtual_network_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroupName, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Virtual Network (Virtual Network Name %q / Resource Group %q): %+v", name, resourceGroupName, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_virtual_network", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    addressSpace := d.Get("address_space").([]interface{})
    dhcpOptions := d.Get("dhcp_options").([]interface{})
    etag := d.Get("etag").(string)
    resourceGUID := d.Get("resource_guid").(string)
    subnets := d.Get("subnets").([]interface{})
    tags := d.Get("tags").(map[string]interface{})

    parameters := network.VirtualNetwork{
        Etag: utils.String(etag),
        Location: utils.String(location),
        VirtualNetworkPropertiesFormat: &network.VirtualNetworkPropertiesFormat{
            AddressSpace: expandArmVirtualNetworkAddressSpace(addressSpace),
            DhcpOptions: expandArmVirtualNetworkDhcpOptions(dhcpOptions),
            ResourceGUID: utils.String(resourceGUID),
            Subnets: expandArmVirtualNetworkSubnet(subnets),
        },
        Tags: tags.Expand(tags),
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroupName, name, parameters)
    if err != nil {
        return fmt.Errorf("Error creating Virtual Network (Virtual Network Name %q / Resource Group %q): %+v", name, resourceGroupName, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Virtual Network (Virtual Network Name %q / Resource Group %q): %+v", name, resourceGroupName, err)
    }


    resp, err := client.Get(ctx, resourceGroupName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Virtual Network (Virtual Network Name %q / Resource Group %q): %+v", name, resourceGroupName, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Virtual Network (Virtual Network Name %q / Resource Group %q) ID", name, resourceGroupName)
    }
    d.SetId(*resp.ID)

    return resourceArmVirtualNetworkRead(d, meta)
}

func resourceArmVirtualNetworkRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).virtualNetworksClient
    ctx, cancel := timeouts.ForRead(meta.(*ArmClient).StopContext, d)
    defer cancel()

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroupName := id.ResourceGroup
    name := id.Path["virtualnetworks"]

    resp, err := client.Get(ctx, resourceGroupName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Virtual Network %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Virtual Network (Virtual Network Name %q / Resource Group %q): %+v", name, resourceGroupName, err)
    }


    d.Set("resource_group", resourceGroupName)
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
        d.Set("provisioning_state", virtualNetworkPropertiesFormat.ProvisioningState)
        d.Set("resource_guid", virtualNetworkPropertiesFormat.ResourceGUID)
        if err := d.Set("subnets", flattenArmVirtualNetworkSubnet(virtualNetworkPropertiesFormat.Subnets)); err != nil {
            return fmt.Errorf("Error setting `subnets`: %+v", err)
        }
    }
    d.Set("etag", resp.Etag)
    d.Set("id", resp.ID)
    d.Set("name", resp.Name)
    d.Set("type", resp.Type)
    d.Set("virtual_network_name", name)

    return tags.FlattenAndSet(d, resp.Tags)
}


func resourceArmVirtualNetworkDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).virtualNetworksClient
    ctx, cancel := timeouts.ForDelete(meta.(*ArmClient).StopContext, d)
    defer cancel()


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroupName := id.ResourceGroup
    name := id.Path["virtualnetworks"]

    future, err := client.Delete(ctx, resourceGroupName, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Virtual Network (Virtual Network Name %q / Resource Group %q): %+v", name, resourceGroupName, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Virtual Network (Virtual Network Name %q / Resource Group %q): %+v", name, resourceGroupName, err)
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

    dNSServers := v["dns_servers"].([]interface{})

    result := network.DhcpOptions{
        DNSServers: utils.ExpandStringSlice(dNSServers),
    }
    return &result
}

func expandArmVirtualNetworkSubnet(input []interface{}) *[]network.Subnet {
    results := make([]network.Subnet, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        iD := v["id"].(string)
        addressPrefix := v["address_prefix"].(string)
        networkSecurityGroup := v["network_security_group"].([]interface{})
        routeTable := v["route_table"].([]interface{})
        iPConfigurations := v["ip_configurations"].([]interface{})
        name := v["name"].(string)
        etag := v["etag"].(string)

        result := network.Subnet{
            Etag: utils.String(etag),
            ID: utils.String(iD),
            Name: utils.String(name),
            SubnetPropertiesFormat: &network.SubnetPropertiesFormat{
                AddressPrefix: utils.String(addressPrefix),
                IPConfigurations: expandArmVirtualNetworkSubResource(iPConfigurations),
                NetworkSecurityGroup: expandArmVirtualNetworkSubResource(networkSecurityGroup),
                RouteTable: expandArmVirtualNetworkSubResource(routeTable),
            },
        }

        results = append(results, result)
    }
    return &results
}

func expandArmVirtualNetworkSubResource(input []interface{}) *[]network.SubResource {
    results := make([]network.SubResource, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        iD := v["id"].(string)

        result := network.SubResource{
            ID: utils.String(iD),
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

    iD := v["id"].(string)

    result := network.SubResource{
        ID: utils.String(iD),
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

    result["dns_servers"] = utils.FlattenStringSlice(input.DNSServers)

    return []interface{}{result}
}

func flattenArmVirtualNetworkSubnet(input *[]network.Subnet) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        if subnetPropertiesFormat := item.SubnetPropertiesFormat; subnetPropertiesFormat != nil {
            if addressPrefix := subnetPropertiesFormat.AddressPrefix; addressPrefix != nil {
                v["address_prefix"] = *addressPrefix
            }
            v["ip_configurations"] = flattenArmVirtualNetworkSubResource(subnetPropertiesFormat.IPConfigurations)
            v["network_security_group"] = flattenArmVirtualNetworkSubResource(subnetPropertiesFormat.NetworkSecurityGroup)
            v["route_table"] = flattenArmVirtualNetworkSubResource(subnetPropertiesFormat.RouteTable)
        }
        if etag := item.Etag; etag != nil {
            v["etag"] = *etag
        }
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

func flattenArmVirtualNetworkSubResource(input *[]network.SubResource) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        if id := item.ID; id != nil {
            v["id"] = *id
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
