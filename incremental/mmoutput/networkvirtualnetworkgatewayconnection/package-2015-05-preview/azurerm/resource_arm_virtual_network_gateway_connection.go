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



func resourceArmVirtualNetworkGatewayConnection() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmVirtualNetworkGatewayConnectionCreateUpdate,
        Read: resourceArmVirtualNetworkGatewayConnectionRead,
        Update: resourceArmVirtualNetworkGatewayConnectionCreateUpdate,
        Delete: resourceArmVirtualNetworkGatewayConnectionDelete,

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

            "connection_status": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(network.Unknown),
                    string(network.Connecting),
                    string(network.Connected),
                    string(network.NotConnected),
                }, false),
                Default: string(network.Unknown),
            },

            "connection_type": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(network.IPsec),
                    string(network.Vnet2Vnet),
                    string(network.ExpressRoute),
                    string(network.VPNClient),
                }, false),
                Default: string(network.IPsec),
            },

            "egress_bytes_transferred": {
                Type: schema.TypeInt,
                Optional: true,
            },

            "etag": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
            },

            "ingress_bytes_transferred": {
                Type: schema.TypeInt,
                Optional: true,
            },

            "key_length": {
                Type: schema.TypeInt,
                Optional: true,
                ForceNew: true,
            },

            "local_network_gateway2": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "location": azure.SchemaLocation(),
                        "etag": {
                            Type: schema.TypeString,
                            Optional: true,
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
                    },
                },
            },

            "peer": {
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

            "resource_guid": {
                Type: schema.TypeString,
                Optional: true,
            },

            "routing_weight": {
                Type: schema.TypeInt,
                Optional: true,
            },

            "shared_key": {
                Type: schema.TypeString,
                Optional: true,
            },

            "value": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
            },

            "virtual_network_gateway1": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "location": azure.SchemaLocation(),
                        "enable_bgp": {
                            Type: schema.TypeBool,
                            Optional: true,
                        },
                        "etag": {
                            Type: schema.TypeString,
                            Optional: true,
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
                                    "private_ip_address": {
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
                        "tags": tags.Schema(),
                        "vpn_type": {
                            Type: schema.TypeString,
                            Optional: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(network.PolicyBased),
                                string(network.RouteBased),
                            }, false),
                            Default: string(network.PolicyBased),
                        },
                    },
                },
            },

            "virtual_network_gateway2": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "location": azure.SchemaLocation(),
                        "enable_bgp": {
                            Type: schema.TypeBool,
                            Optional: true,
                        },
                        "etag": {
                            Type: schema.TypeString,
                            Optional: true,
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
                                    "private_ip_address": {
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
                        "tags": tags.Schema(),
                        "vpn_type": {
                            Type: schema.TypeString,
                            Optional: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(network.PolicyBased),
                                string(network.RouteBased),
                            }, false),
                            Default: string(network.PolicyBased),
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

func resourceArmVirtualNetworkGatewayConnectionCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).virtualNetworkGatewayConnectionsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Virtual Network Gateway Connection %q (Resource Group %q): %+v", name, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_virtual_network_gateway_connection", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    connectionStatus := d.Get("connection_status").(string)
    connectionType := d.Get("connection_type").(string)
    egressBytesTransferred := d.Get("egress_bytes_transferred").(int)
    etag := d.Get("etag").(string)
    ingressBytesTransferred := d.Get("ingress_bytes_transferred").(int)
    keyLength := d.Get("key_length").(int)
    localNetworkGateway2 := d.Get("local_network_gateway2").([]interface{})
    peer := d.Get("peer").([]interface{})
    resourceGuid := d.Get("resource_guid").(string)
    routingWeight := d.Get("routing_weight").(int)
    sharedKey := d.Get("shared_key").(string)
    value := d.Get("value").(string)
    virtualNetworkGateway1 := d.Get("virtual_network_gateway1").([]interface{})
    virtualNetworkGateway2 := d.Get("virtual_network_gateway2").([]interface{})
    t := d.Get("tags").(map[string]interface{})

    parameters := network.ConnectionResetSharedKey{
        Etag: utils.String(etag),
        KeyLength: utils.Int64(int64(keyLength)),
        Location: utils.String(location),
        VirtualNetworkGatewayConnectionPropertiesFormat: &network.VirtualNetworkGatewayConnectionPropertiesFormat{
            ConnectionStatus: network.VirtualNetworkGatewayConnectionStatus(connectionStatus),
            ConnectionType: network.VirtualNetworkGatewayConnectionType(connectionType),
            EgressBytesTransferred: utils.Int64(int64(egressBytesTransferred)),
            IngressBytesTransferred: utils.Int64(int64(ingressBytesTransferred)),
            LocalNetworkGateway2: expandArmVirtualNetworkGatewayConnectionLocalNetworkGateway(localNetworkGateway2),
            Peer: expandArmVirtualNetworkGatewayConnectionSubResource(peer),
            ResourceGUID: utils.String(resourceGuid),
            RoutingWeight: utils.Int32(int32(routingWeight)),
            SharedKey: utils.String(sharedKey),
            VirtualNetworkGateway1: expandArmVirtualNetworkGatewayConnectionVirtualNetworkGateway(virtualNetworkGateway1),
            VirtualNetworkGateway2: expandArmVirtualNetworkGatewayConnectionVirtualNetworkGateway(virtualNetworkGateway2),
        },
        Tags: tags.Expand(t),
        Value: utils.String(value),
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, name, parameters)
    if err != nil {
        return fmt.Errorf("Error creating Virtual Network Gateway Connection %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Virtual Network Gateway Connection %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Virtual Network Gateway Connection %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Virtual Network Gateway Connection %q (Resource Group %q) ID", name, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmVirtualNetworkGatewayConnectionRead(d, meta)
}

func resourceArmVirtualNetworkGatewayConnectionRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).virtualNetworkGatewayConnectionsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["connections"]

    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Virtual Network Gateway Connection %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Virtual Network Gateway Connection %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    d.Set("type", resp.Type)

    return nil
}


func resourceArmVirtualNetworkGatewayConnectionDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).virtualNetworkGatewayConnectionsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["connections"]

    future, err := client.Delete(ctx, resourceGroup, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Virtual Network Gateway Connection %q (Resource Group %q): %+v", name, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Virtual Network Gateway Connection %q (Resource Group %q): %+v", name, resourceGroup, err)
        }
    }

    return nil
}

func expandArmVirtualNetworkGatewayConnectionLocalNetworkGateway(input []interface{}) *network.LocalNetworkGateway {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    location := azure.NormalizeLocation(v["location"].(string))
    t := v["tags"].(map[string]interface{})
    localNetworkAddressSpace := v["local_network_address_space"].([]interface{})
    gatewayIpAddress := v["gateway_ip_address"].(string)
    resourceGuid := v["resource_guid"].(string)
    etag := v["etag"].(string)

    result := network.LocalNetworkGateway{
        Etag: utils.String(etag),
        Location: utils.String(location),
        LocalNetworkGatewayPropertiesFormat: &network.LocalNetworkGatewayPropertiesFormat{
            GatewayIPAddress: utils.String(gatewayIpAddress),
            LocalNetworkAddressSpace: expandArmVirtualNetworkGatewayConnectionAddressSpace(localNetworkAddressSpace),
            ResourceGUID: utils.String(resourceGuid),
        },
        Tags: tags.Expand(t),
    }
    return &result
}

func expandArmVirtualNetworkGatewayConnectionSubResource(input []interface{}) *network.SubResource {
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

func expandArmVirtualNetworkGatewayConnectionVirtualNetworkGateway(input []interface{}) *network.VirtualNetworkGateway {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    location := azure.NormalizeLocation(v["location"].(string))
    t := v["tags"].(map[string]interface{})
    ipConfigurations := v["ip_configurations"].([]interface{})
    gatewayType := v["gateway_type"].(string)
    vpnType := v["vpn_type"].(string)
    enableBgp := v["enable_bgp"].(bool)
    gatewayDefaultSite := v["gateway_default_site"].([]interface{})
    resourceGuid := v["resource_guid"].(string)
    etag := v["etag"].(string)

    result := network.VirtualNetworkGateway{
        Etag: utils.String(etag),
        Location: utils.String(location),
        VirtualNetworkGatewayPropertiesFormat: &network.VirtualNetworkGatewayPropertiesFormat{
            EnableBgp: utils.Bool(enableBgp),
            GatewayDefaultSite: expandArmVirtualNetworkGatewayConnectionSubResource(gatewayDefaultSite),
            GatewayType: network.VirtualNetworkGatewayType(gatewayType),
            IPConfigurations: expandArmVirtualNetworkGatewayConnectionVirtualNetworkGatewayIpConfiguration(ipConfigurations),
            ResourceGUID: utils.String(resourceGuid),
            VpnType: network.VpnType(vpnType),
        },
        Tags: tags.Expand(t),
    }
    return &result
}

func expandArmVirtualNetworkGatewayConnectionAddressSpace(input []interface{}) *network.AddressSpace {
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

func expandArmVirtualNetworkGatewayConnectionVirtualNetworkGatewayIpConfiguration(input []interface{}) *[]network.VirtualNetworkGatewayIpConfiguration {
    results := make([]network.VirtualNetworkGatewayIpConfiguration, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        id := v["id"].(string)
        privateIpAddress := v["private_ip_address"].(string)
        privateIpallocationMethod := v["private_ipallocation_method"].(string)
        subnet := v["subnet"].([]interface{})
        publicIpAddress := v["public_ip_address"].([]interface{})
        name := v["name"].(string)
        etag := v["etag"].(string)

        result := network.VirtualNetworkGatewayIpConfiguration{
            Etag: utils.String(etag),
            ID: utils.String(id),
            Name: utils.String(name),
            VirtualNetworkGatewayIpConfigurationPropertiesFormat: &network.VirtualNetworkGatewayIpConfigurationPropertiesFormat{
                PrivateIPAddress: utils.String(privateIpAddress),
                PrivateIPAllocationMethod: network.IpAllocationMethod(privateIpallocationMethod),
                PublicIPAddress: expandArmVirtualNetworkGatewayConnectionSubResource(publicIpAddress),
                Subnet: expandArmVirtualNetworkGatewayConnectionSubResource(subnet),
            },
        }

        results = append(results, result)
    }
    return &results
}