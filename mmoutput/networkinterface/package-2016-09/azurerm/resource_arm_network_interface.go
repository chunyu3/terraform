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
                        "internal_domain_name_suffix": {
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

            "enable_accelerated_networking": {
                Type: schema.TypeBool,
                Optional: true,
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
                        "application_gateway_backend_address_pools": {
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
                                },
                            },
                        },
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
                        "load_balancer_inbound_nat_rules": {
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
                                },
                            },
                        },
                        "name": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "primary": {
                            Type: schema.TypeBool,
                            Optional: true,
                        },
                        "private_ip_address": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "private_ip_address_version": {
                            Type: schema.TypeString,
                            Optional: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(network.IPv4),
                                string(network.IPv6),
                            }, false),
                            Default: string(network.IPv4),
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
                        "subnet": {
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
                        "default_security_rules": {
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
                                },
                            },
                        },
                        "etag": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "id": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "location": azure.SchemaLocation(),
                        "resource_guid": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "security_rules": {
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
                                },
                            },
                        },
                        "tags": tags.Schema(),
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

    id := d.Get("id").(string)
    location := azure.NormalizeLocation(d.Get("location").(string))
    dnsSettings := d.Get("dns_settings").([]interface{})
    enableAcceleratedNetworking := d.Get("enable_accelerated_networking").(bool)
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
        ID: utils.String(id),
        Location: utils.String(location),
        InterfacePropertiesFormat: &network.InterfacePropertiesFormat{
            DnsSettings: expandArmNetworkInterfaceInterfaceDnsSettings(dnsSettings),
            EnableAcceleratedNetworking: utils.Bool(enableAcceleratedNetworking),
            EnableIpforwarding: utils.Bool(enableIpforwarding),
            IpConfigurations: expandArmNetworkInterfaceInterfaceIPConfiguration(ipConfigurations),
            MacAddress: utils.String(macAddress),
            NetworkSecurityGroup: expandArmNetworkInterfaceSecurityGroup(networkSecurityGroup),
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
        d.Set("enable_accelerated_networking", interfacePropertiesFormat.EnableAcceleratedNetworking)
        d.Set("enable_ipforwarding", interfacePropertiesFormat.EnableIpforwarding)
        if err := d.Set("ip_configurations", flattenArmNetworkInterfaceInterfaceIPConfiguration(interfacePropertiesFormat.IpConfigurations)); err != nil {
            return fmt.Errorf("Error setting `ip_configurations`: %+v", err)
        }
        d.Set("mac_address", interfacePropertiesFormat.MacAddress)
        if err := d.Set("network_security_group", flattenArmNetworkInterfaceSecurityGroup(interfacePropertiesFormat.NetworkSecurityGroup)); err != nil {
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
    internalDomainNameSuffix := v["internal_domain_name_suffix"].(string)

    result := network.InterfaceDnsSettings{
        AppliedDnsServers: utils.ExpandStringSlice(appliedDnsServers),
        DnsServers: utils.ExpandStringSlice(dnsServers),
        InternalDnsNameLabel: utils.String(internalDnsNameLabel),
        InternalDomainNameSuffix: utils.String(internalDomainNameSuffix),
        InternalFqdn: utils.String(internalFqdn),
    }
    return &result
}

func expandArmNetworkInterfaceInterfaceIPConfiguration(input []interface{}) *[]network.InterfaceIPConfiguration {
    results := make([]network.InterfaceIPConfiguration, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        id := v["id"].(string)
        applicationGatewayBackendAddressPools := v["application_gateway_backend_address_pools"].([]interface{})
        loadBalancerBackendAddressPools := v["load_balancer_backend_address_pools"].([]interface{})
        loadBalancerInboundNatRules := v["load_balancer_inbound_nat_rules"].([]interface{})
        privateIpAddress := v["private_ip_address"].(string)
        privateIpallocationMethod := v["private_ipallocation_method"].(string)
        privateIpAddressVersion := v["private_ip_address_version"].(string)
        subnet := v["subnet"].([]interface{})
        primary := v["primary"].(bool)
        publicIpAddress := v["public_ip_address"].([]interface{})
        name := v["name"].(string)
        etag := v["etag"].(string)

        result := network.InterfaceIPConfiguration{
            Etag: utils.String(etag),
            ID: utils.String(id),
            Name: utils.String(name),
            InterfaceIPConfigurationPropertiesFormat: &network.InterfaceIPConfigurationPropertiesFormat{
                ApplicationGatewayBackendAddressPools: expandArmNetworkInterfaceApplicationGatewayBackendAddressPool(applicationGatewayBackendAddressPools),
                LoadBalancerBackendAddressPools: expandArmNetworkInterfaceBackendAddressPool(loadBalancerBackendAddressPools),
                LoadBalancerInboundNatRules: expandArmNetworkInterfaceInboundNatRule(loadBalancerInboundNatRules),
                Primary: utils.Bool(primary),
                PrivateIpAddress: utils.String(privateIpAddress),
                PrivateIpAddressVersion: network.IPVersion(privateIpAddressVersion),
                PrivateIpallocationMethod: network.IPAllocationMethod(privateIpallocationMethod),
                PublicIpAddress: expandArmNetworkInterfacePublicIPAddress(publicIpAddress),
                Subnet: expandArmNetworkInterfaceSubnet(subnet),
            },
        }

        results = append(results, result)
    }
    return &results
}

func expandArmNetworkInterfaceSecurityGroup(input []interface{}) *network.SecurityGroup {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    id := v["id"].(string)
    location := azure.NormalizeLocation(v["location"].(string))
    t := v["tags"].(map[string]interface{})
    securityRules := v["security_rules"].([]interface{})
    defaultSecurityRules := v["default_security_rules"].([]interface{})
    resourceGuid := v["resource_guid"].(string)
    etag := v["etag"].(string)

    result := network.SecurityGroup{
        Etag: utils.String(etag),
        ID: utils.String(id),
        Location: utils.String(location),
        SecurityGroupPropertiesFormat: &network.SecurityGroupPropertiesFormat{
            DefaultSecurityRules: expandArmNetworkInterfaceSecurityRule(defaultSecurityRules),
            ResourceGuid: utils.String(resourceGuid),
            SecurityRules: expandArmNetworkInterfaceSecurityRule(securityRules),
        },
        Tags: tags.Expand(t),
    }
    return &result
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

func expandArmNetworkInterfaceApplicationGatewayBackendAddressPool(input []interface{}) *[]network.ApplicationGatewayBackendAddressPool {
    results := make([]network.ApplicationGatewayBackendAddressPool, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        id := v["id"].(string)
        name := v["name"].(string)
        etag := v["etag"].(string)

        result := network.ApplicationGatewayBackendAddressPool{
            Etag: utils.String(etag),
            ID: utils.String(id),
            Name: utils.String(name),
        }

        results = append(results, result)
    }
    return &results
}

func expandArmNetworkInterfaceBackendAddressPool(input []interface{}) *[]network.BackendAddressPool {
    results := make([]network.BackendAddressPool, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        id := v["id"].(string)
        name := v["name"].(string)
        etag := v["etag"].(string)

        result := network.BackendAddressPool{
            Etag: utils.String(etag),
            ID: utils.String(id),
            Name: utils.String(name),
        }

        results = append(results, result)
    }
    return &results
}

func expandArmNetworkInterfaceInboundNatRule(input []interface{}) *[]network.InboundNatRule {
    results := make([]network.InboundNatRule, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        id := v["id"].(string)
        name := v["name"].(string)
        etag := v["etag"].(string)

        result := network.InboundNatRule{
            Etag: utils.String(etag),
            ID: utils.String(id),
            Name: utils.String(name),
        }

        results = append(results, result)
    }
    return &results
}

func expandArmNetworkInterfacePublicIPAddress(input []interface{}) *network.PublicIPAddress {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    id := v["id"].(string)
    location := azure.NormalizeLocation(v["location"].(string))
    t := v["tags"].(map[string]interface{})
    etag := v["etag"].(string)

    result := network.PublicIPAddress{
        Etag: utils.String(etag),
        ID: utils.String(id),
        Location: utils.String(location),
        Tags: tags.Expand(t),
    }
    return &result
}

func expandArmNetworkInterfaceSubnet(input []interface{}) *network.Subnet {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    id := v["id"].(string)
    name := v["name"].(string)
    etag := v["etag"].(string)

    result := network.Subnet{
        Etag: utils.String(etag),
        ID: utils.String(id),
        Name: utils.String(name),
    }
    return &result
}

func expandArmNetworkInterfaceSecurityRule(input []interface{}) *[]network.SecurityRule {
    results := make([]network.SecurityRule, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        id := v["id"].(string)
        name := v["name"].(string)
        etag := v["etag"].(string)

        result := network.SecurityRule{
            Etag: utils.String(etag),
            ID: utils.String(id),
            Name: utils.String(name),
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
    if internalDomainNameSuffix := input.InternalDomainNameSuffix; internalDomainNameSuffix != nil {
        result["internal_domain_name_suffix"] = *internalDomainNameSuffix
    }
    if internalFqdn := input.InternalFqdn; internalFqdn != nil {
        result["internal_fqdn"] = *internalFqdn
    }

    return []interface{}{result}
}

func flattenArmNetworkInterfaceInterfaceIPConfiguration(input *[]network.InterfaceIPConfiguration) []interface{} {
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
        if interfaceIPConfigurationPropertiesFormat := item.InterfaceIPConfigurationPropertiesFormat; interfaceIPConfigurationPropertiesFormat != nil {
            v["application_gateway_backend_address_pools"] = flattenArmNetworkInterfaceApplicationGatewayBackendAddressPool(interfaceIPConfigurationPropertiesFormat.ApplicationGatewayBackendAddressPools)
            v["load_balancer_backend_address_pools"] = flattenArmNetworkInterfaceBackendAddressPool(interfaceIPConfigurationPropertiesFormat.LoadBalancerBackendAddressPools)
            v["load_balancer_inbound_nat_rules"] = flattenArmNetworkInterfaceInboundNatRule(interfaceIPConfigurationPropertiesFormat.LoadBalancerInboundNatRules)
            if primary := interfaceIPConfigurationPropertiesFormat.Primary; primary != nil {
                v["primary"] = *primary
            }
            if privateIpAddress := interfaceIPConfigurationPropertiesFormat.PrivateIpAddress; privateIpAddress != nil {
                v["private_ip_address"] = *privateIpAddress
            }
            v["private_ip_address_version"] = string(interfaceIPConfigurationPropertiesFormat.PrivateIpAddressVersion)
            v["private_ipallocation_method"] = string(interfaceIPConfigurationPropertiesFormat.PrivateIpallocationMethod)
            v["public_ip_address"] = flattenArmNetworkInterfacePublicIPAddress(interfaceIPConfigurationPropertiesFormat.PublicIpAddress)
            v["subnet"] = flattenArmNetworkInterfaceSubnet(interfaceIPConfigurationPropertiesFormat.Subnet)
        }
        if etag := item.Etag; etag != nil {
            v["etag"] = *etag
        }

        results = append(results, v)
    }

    return results
}

func flattenArmNetworkInterfaceSecurityGroup(input *network.SecurityGroup) []interface{} {
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
    if securityGroupPropertiesFormat := input.SecurityGroupPropertiesFormat; securityGroupPropertiesFormat != nil {
        result["default_security_rules"] = flattenArmNetworkInterfaceSecurityRule(securityGroupPropertiesFormat.DefaultSecurityRules)
        if resourceGuid := securityGroupPropertiesFormat.ResourceGuid; resourceGuid != nil {
            result["resource_guid"] = *resourceGuid
        }
        result["security_rules"] = flattenArmNetworkInterfaceSecurityRule(securityGroupPropertiesFormat.SecurityRules)
    }
    if etag := input.Etag; etag != nil {
        result["etag"] = *etag
    }

    return []interface{}{result}
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

func flattenArmNetworkInterfaceApplicationGatewayBackendAddressPool(input *[]network.ApplicationGatewayBackendAddressPool) []interface{} {
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

        results = append(results, v)
    }

    return results
}

func flattenArmNetworkInterfaceBackendAddressPool(input *[]network.BackendAddressPool) []interface{} {
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

        results = append(results, v)
    }

    return results
}

func flattenArmNetworkInterfaceInboundNatRule(input *[]network.InboundNatRule) []interface{} {
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

        results = append(results, v)
    }

    return results
}

func flattenArmNetworkInterfacePublicIPAddress(input *network.PublicIPAddress) []interface{} {
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

func flattenArmNetworkInterfaceSubnet(input *network.Subnet) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if id := input.ID; id != nil {
        result["id"] = *id
    }
    if name := input.Name; name != nil {
        result["name"] = *name
    }
    if etag := input.Etag; etag != nil {
        result["etag"] = *etag
    }

    return []interface{}{result}
}

func flattenArmNetworkInterfaceSecurityRule(input *[]network.SecurityRule) []interface{} {
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

        results = append(results, v)
    }

    return results
}
