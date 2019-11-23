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

            "vpn_type": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(network.PolicyBased),
                    string(network.RouteBased),
                }, false),
                Default: string(network.PolicyBased),
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

    location := azure.NormalizeLocation(d.Get("location").(string))
    enableBgp := d.Get("enable_bgp").(bool)
    etag := d.Get("etag").(string)
    gatewayDefaultSite := d.Get("gateway_default_site").([]interface{})
    gatewayType := d.Get("gateway_type").(string)
    ipConfigurations := d.Get("ip_configurations").([]interface{})
    resourceGuid := d.Get("resource_guid").(string)
    vpnType := d.Get("vpn_type").(string)
    t := d.Get("tags").(map[string]interface{})

    parameters := network.VirtualNetworkGateway{
        Etag: utils.String(etag),
        Location: utils.String(location),
        VirtualNetworkGatewayPropertiesFormat: &network.VirtualNetworkGatewayPropertiesFormat{
            EnableBgp: utils.Bool(enableBgp),
            GatewayDefaultSite: expandArmVirtualNetworkGatewaySubResource(gatewayDefaultSite),
            GatewayType: network.VirtualNetworkGatewayType(gatewayType),
            IPConfigurations: expandArmVirtualNetworkGatewayVirtualNetworkGatewayIpConfiguration(ipConfigurations),
            ResourceGUID: utils.String(resourceGuid),
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
    name := id.Path["virtualnetworkgateways"]

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
    d.Set("type", resp.Type)

    return nil
}


func resourceArmVirtualNetworkGatewayDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).virtualNetworkGatewaysClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["virtualnetworkgateways"]

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

func expandArmVirtualNetworkGatewayVirtualNetworkGatewayIpConfiguration(input []interface{}) *[]network.VirtualNetworkGatewayIpConfiguration {
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
                PublicIPAddress: expandArmVirtualNetworkGatewaySubResource(publicIpAddress),
                Subnet: expandArmVirtualNetworkGatewaySubResource(subnet),
            },
        }

        results = append(results, result)
    }
    return &results
}