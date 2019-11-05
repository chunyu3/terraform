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



func resourceArmNetworkInterface() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmNetworkInterfaceCreateUpdate,
        Read: resourceArmNetworkInterfaceRead,
        Update: resourceArmNetworkInterfaceCreateUpdate,
        Delete: resourceArmNetworkInterfaceDelete,

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

            "network_interface_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "dns_settings": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "applied_dns_servers": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                        "dns_servers": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                        "internal_dns_name_label": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "internal_fqdn": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },

            "enable_ipforwarding": {
                Type: schema.TypeBool,
                Optional: true,
            },

            "etag": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
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
                        "load_balancer_backend_address_pools": {
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
                        "load_balancer_inbound_nat_rules": {
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

            "mac_address": {
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

            "primary": {
                Type: schema.TypeBool,
                Optional: true,
            },

            "resource_guid": {
                Type: schema.TypeString,
                Optional: true,
            },

            "virtual_machine": {
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

func resourceArmNetworkInterfaceCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).networkInterfacesClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    networkInterfaceName := d.Get("network_interface_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, networkInterfaceName)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Network Interface (Network Interface Name %q / Resource Group %q): %+v", networkInterfaceName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_network_interface", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    dnsSettings := d.Get("dns_settings").([]interface{})
    enableIpforwarding := d.Get("enable_ipforwarding").(bool)
    etag := d.Get("etag").(string)
    ipConfigurations := d.Get("ip_configurations").([]interface{})
    macAddress := d.Get("mac_address").(string)
    networkSecurityGroup := d.Get("network_security_group").([]interface{})
    primary := d.Get("primary").(bool)
    resourceGuid := d.Get("resource_guid").(string)
    virtualMachine := d.Get("virtual_machine").([]interface{})
    t := d.Get("tags").(map[string]interface{})

    parameters := network.Interface{
        Etag: utils.String(etag),
        Location: utils.String(location),
        InterfacePropertiesFormat: &network.InterfacePropertiesFormat{
            DnsSettings: expandArmNetworkInterfaceInterfaceDnsSettings(dnsSettings),
            EnableIpforwarding: utils.Bool(enableIpforwarding),
            IpConfigurations: expandArmNetworkInterfaceInterfaceIpConfiguration(ipConfigurations),
            MacAddress: utils.String(macAddress),
            NetworkSecurityGroup: expandArmNetworkInterfaceSubResource(networkSecurityGroup),
            Primary: utils.Bool(primary),
            ResourceGuid: utils.String(resourceGuid),
            VirtualMachine: expandArmNetworkInterfaceSubResource(virtualMachine),
        },
        Tags: tags.Expand(t),
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, networkInterfaceName, parameters)
    if err != nil {
        return fmt.Errorf("Error creating Network Interface (Network Interface Name %q / Resource Group %q): %+v", networkInterfaceName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Network Interface (Network Interface Name %q / Resource Group %q): %+v", networkInterfaceName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, networkInterfaceName)
    if err != nil {
        return fmt.Errorf("Error retrieving Network Interface (Network Interface Name %q / Resource Group %q): %+v", networkInterfaceName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Network Interface (Network Interface Name %q / Resource Group %q) ID", networkInterfaceName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmNetworkInterfaceRead(d, meta)
}

func resourceArmNetworkInterfaceRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).networkInterfacesClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    networkInterfaceName := id.Path["networkInterfaces"]

    resp, err := client.Get(ctx, resourceGroup, networkInterfaceName)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Network Interface %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Network Interface (Network Interface Name %q / Resource Group %q): %+v", networkInterfaceName, resourceGroup, err)
    }


    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if interfacePropertiesFormat := resp.InterfacePropertiesFormat; interfacePropertiesFormat != nil {
        if err := d.Set("dns_settings", flattenArmNetworkInterfaceInterfaceDnsSettings(interfacePropertiesFormat.DnsSettings)); err != nil {
            return fmt.Errorf("Error setting `dns_settings`: %+v", err)
        }
        d.Set("enable_ipforwarding", interfacePropertiesFormat.EnableIpforwarding)
        if err := d.Set("ip_configurations", flattenArmNetworkInterfaceInterfaceIpConfiguration(interfacePropertiesFormat.IpConfigurations)); err != nil {
            return fmt.Errorf("Error setting `ip_configurations`: %+v", err)
        }
        d.Set("mac_address", interfacePropertiesFormat.MacAddress)
        if err := d.Set("network_security_group", flattenArmNetworkInterfaceSubResource(interfacePropertiesFormat.NetworkSecurityGroup)); err != nil {
            return fmt.Errorf("Error setting `network_security_group`: %+v", err)
        }
        d.Set("primary", interfacePropertiesFormat.Primary)
        d.Set("provisioning_state", interfacePropertiesFormat.ProvisioningState)
        d.Set("resource_guid", interfacePropertiesFormat.ResourceGuid)
        if err := d.Set("virtual_machine", flattenArmNetworkInterfaceSubResource(interfacePropertiesFormat.VirtualMachine)); err != nil {
            return fmt.Errorf("Error setting `virtual_machine`: %+v", err)
        }
    }
    d.Set("etag", resp.Etag)
    d.Set("network_interface_name", networkInterfaceName)
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}


func resourceArmNetworkInterfaceDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).networkInterfacesClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    networkInterfaceName := id.Path["networkInterfaces"]

    future, err := client.Delete(ctx, resourceGroup, networkInterfaceName)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Network Interface (Network Interface Name %q / Resource Group %q): %+v", networkInterfaceName, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Network Interface (Network Interface Name %q / Resource Group %q): %+v", networkInterfaceName, resourceGroup, err)
        }
    }

    return nil
}

func expandArmNetworkInterfaceInterfaceDnsSettings(input []interface{}) *network.InterfaceDnsSettings {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    dnsServers := v["dns_servers"].([]interface{})
    appliedDnsServers := v["applied_dns_servers"].([]interface{})
    internalDnsNameLabel := v["internal_dns_name_label"].(string)
    internalFqdn := v["internal_fqdn"].(string)

    result := network.InterfaceDnsSettings{
        AppliedDnsServers: utils.ExpandStringSlice(appliedDnsServers),
        DnsServers: utils.ExpandStringSlice(dnsServers),
        InternalDnsNameLabel: utils.String(internalDnsNameLabel),
        InternalFqdn: utils.String(internalFqdn),
    }
    return &result
}

func expandArmNetworkInterfaceInterfaceIpConfiguration(input []interface{}) *[]network.InterfaceIpConfiguration {
    results := make([]network.InterfaceIpConfiguration, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        id := v["id"].(string)
        privateIpAddress := v["private_ip_address"].(string)
        privateIpallocationMethod := v["private_ipallocation_method"].(string)
        subnet := v["subnet"].([]interface{})
        publicIpAddress := v["public_ip_address"].([]interface{})
        loadBalancerBackendAddressPools := v["load_balancer_backend_address_pools"].([]interface{})
        loadBalancerInboundNatRules := v["load_balancer_inbound_nat_rules"].([]interface{})
        name := v["name"].(string)
        etag := v["etag"].(string)

        result := network.InterfaceIpConfiguration{
            Etag: utils.String(etag),
            ID: utils.String(id),
            Name: utils.String(name),
            InterfaceIpConfigurationPropertiesFormat: &network.InterfaceIpConfigurationPropertiesFormat{
                LoadBalancerBackendAddressPools: expandArmNetworkInterfaceSubResource(loadBalancerBackendAddressPools),
                LoadBalancerInboundNatRules: expandArmNetworkInterfaceSubResource(loadBalancerInboundNatRules),
                PrivateIpAddress: utils.String(privateIpAddress),
                PrivateIpallocationMethod: network.IpAllocationMethod(privateIpallocationMethod),
                PublicIpAddress: expandArmNetworkInterfaceSubResource(publicIpAddress),
                Subnet: expandArmNetworkInterfaceSubResource(subnet),
            },
        }

        results = append(results, result)
    }
    return &results
}

func expandArmNetworkInterfaceSubResource(input []interface{}) *network.SubResource {
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

func expandArmNetworkInterfaceSubResource(input []interface{}) *[]network.SubResource {
    results := make([]network.SubResource, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        id := v["id"].(string)

        result := network.SubResource{
            ID: utils.String(id),
        }

        results = append(results, result)
    }
    return &results
}


func flattenArmNetworkInterfaceInterfaceDnsSettings(input *network.InterfaceDnsSettings) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    result["applied_dns_servers"] = utils.FlattenStringSlice(input.AppliedDnsServers)
    result["dns_servers"] = utils.FlattenStringSlice(input.DnsServers)
    if internalDnsNameLabel := input.InternalDnsNameLabel; internalDnsNameLabel != nil {
        result["internal_dns_name_label"] = *internalDnsNameLabel
    }
    if internalFqdn := input.InternalFqdn; internalFqdn != nil {
        result["internal_fqdn"] = *internalFqdn
    }

    return []interface{}{result}
}

func flattenArmNetworkInterfaceInterfaceIpConfiguration(input *[]network.InterfaceIpConfiguration) []interface{} {
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
        if interfaceIpConfigurationPropertiesFormat := item.InterfaceIpConfigurationPropertiesFormat; interfaceIpConfigurationPropertiesFormat != nil {
            v["load_balancer_backend_address_pools"] = flattenArmNetworkInterfaceSubResource(interfaceIpConfigurationPropertiesFormat.LoadBalancerBackendAddressPools)
            v["load_balancer_inbound_nat_rules"] = flattenArmNetworkInterfaceSubResource(interfaceIpConfigurationPropertiesFormat.LoadBalancerInboundNatRules)
            if privateIpAddress := interfaceIpConfigurationPropertiesFormat.PrivateIpAddress; privateIpAddress != nil {
                v["private_ip_address"] = *privateIpAddress
            }
            v["private_ipallocation_method"] = string(interfaceIpConfigurationPropertiesFormat.PrivateIpallocationMethod)
            v["public_ip_address"] = flattenArmNetworkInterfaceSubResource(interfaceIpConfigurationPropertiesFormat.PublicIpAddress)
            v["subnet"] = flattenArmNetworkInterfaceSubResource(interfaceIpConfigurationPropertiesFormat.Subnet)
        }

        results = append(results, v)
    }

    return results
}

func flattenArmNetworkInterfaceSubResource(input *network.SubResource) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if id := input.ID; id != nil {
        result["id"] = *id
    }

    return []interface{}{result}
}

func flattenArmNetworkInterfaceSubResource(input *[]network.SubResource) []interface{} {
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
