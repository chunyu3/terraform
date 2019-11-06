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



func resourceArmNetworkSecurityGroup() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmNetworkSecurityGroupCreateUpdate,
        Read: resourceArmNetworkSecurityGroupRead,
        Update: resourceArmNetworkSecurityGroupCreateUpdate,
        Delete: resourceArmNetworkSecurityGroupDelete,

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

            "default_security_rules": {
                Type: schema.TypeList,
                Optional: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "access": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(network.Allow),
                                string(network.Deny),
                            }, false),
                        },
                        "destination_address_prefix": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "direction": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(network.Inbound),
                                string(network.Outbound),
                            }, false),
                        },
                        "protocol": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(network.Tcp),
                                string(network.Udp),
                                string(network.*),
                            }, false),
                        },
                        "source_address_prefix": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "description": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "destination_port_range": {
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
                        "priority": {
                            Type: schema.TypeInt,
                            Optional: true,
                        },
                        "source_port_range": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },

            "etag": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
            },

            "network_interfaces": {
                Type: schema.TypeList,
                Optional: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
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
                        "location": azure.SchemaLocation(),
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
                        "primary": {
                            Type: schema.TypeBool,
                            Optional: true,
                        },
                        "resource_guid": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "tags": tags.Schema(),
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
                    },
                },
            },

            "resource_guid": {
                Type: schema.TypeString,
                Optional: true,
            },

            "security_rules": {
                Type: schema.TypeList,
                Optional: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "access": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(network.Allow),
                                string(network.Deny),
                            }, false),
                        },
                        "destination_address_prefix": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "direction": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(network.Inbound),
                                string(network.Outbound),
                            }, false),
                        },
                        "protocol": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(network.Tcp),
                                string(network.Udp),
                                string(network.*),
                            }, false),
                        },
                        "source_address_prefix": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "description": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "destination_port_range": {
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
                        "priority": {
                            Type: schema.TypeInt,
                            Optional: true,
                        },
                        "source_port_range": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
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

func resourceArmNetworkSecurityGroupCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).networkSecurityGroupsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Network Security Group %q (Resource Group %q): %+v", name, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_network_security_group", *existing.ID)
        }
    }

    id := d.Get("id").(string)
    location := azure.NormalizeLocation(d.Get("location").(string))
    defaultSecurityRules := d.Get("default_security_rules").([]interface{})
    etag := d.Get("etag").(string)
    networkInterfaces := d.Get("network_interfaces").([]interface{})
    resourceGuid := d.Get("resource_guid").(string)
    securityRules := d.Get("security_rules").([]interface{})
    subnets := d.Get("subnets").([]interface{})
    t := d.Get("tags").(map[string]interface{})

    parameters := network.SecurityGroup{
        Etag: utils.String(etag),
        ID: utils.String(id),
        Location: utils.String(location),
        SecurityGroupPropertiesFormat: &network.SecurityGroupPropertiesFormat{
            DefaultSecurityRules: expandArmNetworkSecurityGroupSecurityRule(defaultSecurityRules),
            NetworkInterfaces: expandArmNetworkSecurityGroupInterface(networkInterfaces),
            ResourceGuid: utils.String(resourceGuid),
            SecurityRules: expandArmNetworkSecurityGroupSecurityRule(securityRules),
            Subnets: expandArmNetworkSecurityGroupSubnet(subnets),
        },
        Tags: tags.Expand(t),
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, name, parameters)
    if err != nil {
        return fmt.Errorf("Error creating Network Security Group %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Network Security Group %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Network Security Group %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Network Security Group %q (Resource Group %q) ID", name, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmNetworkSecurityGroupRead(d, meta)
}

func resourceArmNetworkSecurityGroupRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).networkSecurityGroupsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["networkSecurityGroups"]

    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Network Security Group %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Network Security Group %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if securityGroupPropertiesFormat := resp.SecurityGroupPropertiesFormat; securityGroupPropertiesFormat != nil {
        if err := d.Set("default_security_rules", flattenArmNetworkSecurityGroupSecurityRule(securityGroupPropertiesFormat.DefaultSecurityRules)); err != nil {
            return fmt.Errorf("Error setting `default_security_rules`: %+v", err)
        }
        if err := d.Set("network_interfaces", flattenArmNetworkSecurityGroupInterface(securityGroupPropertiesFormat.NetworkInterfaces)); err != nil {
            return fmt.Errorf("Error setting `network_interfaces`: %+v", err)
        }
        d.Set("provisioning_state", securityGroupPropertiesFormat.ProvisioningState)
        d.Set("resource_guid", securityGroupPropertiesFormat.ResourceGuid)
        if err := d.Set("security_rules", flattenArmNetworkSecurityGroupSecurityRule(securityGroupPropertiesFormat.SecurityRules)); err != nil {
            return fmt.Errorf("Error setting `security_rules`: %+v", err)
        }
        if err := d.Set("subnets", flattenArmNetworkSecurityGroupSubnet(securityGroupPropertiesFormat.Subnets)); err != nil {
            return fmt.Errorf("Error setting `subnets`: %+v", err)
        }
    }
    d.Set("etag", resp.Etag)
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}


func resourceArmNetworkSecurityGroupDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).networkSecurityGroupsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["networkSecurityGroups"]

    future, err := client.Delete(ctx, resourceGroup, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Network Security Group %q (Resource Group %q): %+v", name, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Network Security Group %q (Resource Group %q): %+v", name, resourceGroup, err)
        }
    }

    return nil
}

func expandArmNetworkSecurityGroupSecurityRule(input []interface{}) *[]network.SecurityRule {
    results := make([]network.SecurityRule, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        id := v["id"].(string)
        description := v["description"].(string)
        protocol := v["protocol"].(string)
        sourcePortRange := v["source_port_range"].(string)
        destinationPortRange := v["destination_port_range"].(string)
        sourceAddressPrefix := v["source_address_prefix"].(string)
        destinationAddressPrefix := v["destination_address_prefix"].(string)
        access := v["access"].(string)
        priority := v["priority"].(int)
        direction := v["direction"].(string)
        name := v["name"].(string)
        etag := v["etag"].(string)

        result := network.SecurityRule{
            Etag: utils.String(etag),
            ID: utils.String(id),
            Name: utils.String(name),
            SecurityRulePropertiesFormat: &network.SecurityRulePropertiesFormat{
                Access: network.SecurityRuleAccess(access),
                Description: utils.String(description),
                DestinationAddressPrefix: utils.String(destinationAddressPrefix),
                DestinationPortRange: utils.String(destinationPortRange),
                Direction: network.SecurityRuleDirection(direction),
                Priority: utils.Int32(int32(priority)),
                Protocol: network.SecurityRuleProtocol(protocol),
                SourceAddressPrefix: utils.String(sourceAddressPrefix),
                SourcePortRange: utils.String(sourcePortRange),
            },
        }

        results = append(results, result)
    }
    return &results
}

func expandArmNetworkSecurityGroupInterface(input []interface{}) *[]network.Interface {
    results := make([]network.Interface, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        id := v["id"].(string)
        location := azure.NormalizeLocation(v["location"].(string))
        t := v["tags"].(map[string]interface{})
        virtualMachine := v["virtual_machine"].([]interface{})
        networkSecurityGroup := v["network_security_group"].([]interface{})
        ipConfigurations := v["ip_configurations"].([]interface{})
        dnsSettings := v["dns_settings"].([]interface{})
        macAddress := v["mac_address"].(string)
        primary := v["primary"].(bool)
        enableIpforwarding := v["enable_ipforwarding"].(bool)
        resourceGuid := v["resource_guid"].(string)
        etag := v["etag"].(string)

        result := network.Interface{
            Etag: utils.String(etag),
            ID: utils.String(id),
            Location: utils.String(location),
            InterfacePropertiesFormat: &network.InterfacePropertiesFormat{
                DnsSettings: expandArmNetworkSecurityGroupInterfaceDnsSettings(dnsSettings),
                EnableIpforwarding: utils.Bool(enableIpforwarding),
                IpConfigurations: expandArmNetworkSecurityGroupInterfaceIPConfiguration(ipConfigurations),
                MacAddress: utils.String(macAddress),
                NetworkSecurityGroup: expandArmNetworkSecurityGroupSecurityGroup(networkSecurityGroup),
                Primary: utils.Bool(primary),
                ResourceGuid: utils.String(resourceGuid),
                VirtualMachine: expandArmNetworkSecurityGroupSubResource(virtualMachine),
            },
            Tags: tags.Expand(t),
        }

        results = append(results, result)
    }
    return &results
}

func expandArmNetworkSecurityGroupSubnet(input []interface{}) *[]network.Subnet {
    results := make([]network.Subnet, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        id := v["id"].(string)
        addressPrefix := v["address_prefix"].(string)
        networkSecurityGroup := v["network_security_group"].([]interface{})
        routeTable := v["route_table"].([]interface{})
        ipConfigurations := v["ip_configurations"].([]interface{})
        name := v["name"].(string)
        etag := v["etag"].(string)

        result := network.Subnet{
            Etag: utils.String(etag),
            ID: utils.String(id),
            Name: utils.String(name),
            SubnetPropertiesFormat: &network.SubnetPropertiesFormat{
                AddressPrefix: utils.String(addressPrefix),
                IpConfigurations: expandArmNetworkSecurityGroupIPConfiguration(ipConfigurations),
                NetworkSecurityGroup: expandArmNetworkSecurityGroupSecurityGroup(networkSecurityGroup),
                RouteTable: expandArmNetworkSecurityGroupRouteTable(routeTable),
            },
        }

        results = append(results, result)
    }
    return &results
}

func expandArmNetworkSecurityGroupInterfaceDnsSettings(input []interface{}) *network.InterfaceDnsSettings {
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

func expandArmNetworkSecurityGroupInterfaceIPConfiguration(input []interface{}) *[]network.InterfaceIPConfiguration {
    results := make([]network.InterfaceIPConfiguration, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        id := v["id"].(string)
        name := v["name"].(string)
        etag := v["etag"].(string)

        result := network.InterfaceIPConfiguration{
            Etag: utils.String(etag),
            ID: utils.String(id),
            Name: utils.String(name),
        }

        results = append(results, result)
    }
    return &results
}

func expandArmNetworkSecurityGroupSecurityGroup(input []interface{}) *network.SecurityGroup {
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

func expandArmNetworkSecurityGroupSubResource(input []interface{}) *network.SubResource {
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

func expandArmNetworkSecurityGroupIPConfiguration(input []interface{}) *[]network.IPConfiguration {
    results := make([]network.IPConfiguration, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        id := v["id"].(string)
        name := v["name"].(string)
        etag := v["etag"].(string)

        result := network.IPConfiguration{
            Etag: utils.String(etag),
            ID: utils.String(id),
            Name: utils.String(name),
        }

        results = append(results, result)
    }
    return &results
}

func expandArmNetworkSecurityGroupRouteTable(input []interface{}) *network.RouteTable {
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


func flattenArmNetworkSecurityGroupSecurityRule(input *[]network.SecurityRule) []interface{} {
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
        if securityRulePropertiesFormat := item.SecurityRulePropertiesFormat; securityRulePropertiesFormat != nil {
            v["access"] = string(securityRulePropertiesFormat.Access)
            if description := securityRulePropertiesFormat.Description; description != nil {
                v["description"] = *description
            }
            if destinationAddressPrefix := securityRulePropertiesFormat.DestinationAddressPrefix; destinationAddressPrefix != nil {
                v["destination_address_prefix"] = *destinationAddressPrefix
            }
            if destinationPortRange := securityRulePropertiesFormat.DestinationPortRange; destinationPortRange != nil {
                v["destination_port_range"] = *destinationPortRange
            }
            v["direction"] = string(securityRulePropertiesFormat.Direction)
            if priority := securityRulePropertiesFormat.Priority; priority != nil {
                v["priority"] = int(*priority)
            }
            v["protocol"] = string(securityRulePropertiesFormat.Protocol)
            if sourceAddressPrefix := securityRulePropertiesFormat.SourceAddressPrefix; sourceAddressPrefix != nil {
                v["source_address_prefix"] = *sourceAddressPrefix
            }
            if sourcePortRange := securityRulePropertiesFormat.SourcePortRange; sourcePortRange != nil {
                v["source_port_range"] = *sourcePortRange
            }
        }
        if etag := item.Etag; etag != nil {
            v["etag"] = *etag
        }

        results = append(results, v)
    }

    return results
}

func flattenArmNetworkSecurityGroupInterface(input *[]network.Interface) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        if id := item.ID; id != nil {
            v["id"] = *id
        }
        if location := item.Location; location != nil {
            v["location"] = azure.NormalizeLocation(*location)
        }
        if interfacePropertiesFormat := item.InterfacePropertiesFormat; interfacePropertiesFormat != nil {
            v["dns_settings"] = flattenArmNetworkSecurityGroupInterfaceDnsSettings(interfacePropertiesFormat.DnsSettings)
            if enableIpforwarding := interfacePropertiesFormat.EnableIpforwarding; enableIpforwarding != nil {
                v["enable_ipforwarding"] = *enableIpforwarding
            }
            v["ip_configurations"] = flattenArmNetworkSecurityGroupInterfaceIPConfiguration(interfacePropertiesFormat.IpConfigurations)
            if macAddress := interfacePropertiesFormat.MacAddress; macAddress != nil {
                v["mac_address"] = *macAddress
            }
            v["network_security_group"] = flattenArmNetworkSecurityGroupSecurityGroup(interfacePropertiesFormat.NetworkSecurityGroup)
            if primary := interfacePropertiesFormat.Primary; primary != nil {
                v["primary"] = *primary
            }
            if resourceGuid := interfacePropertiesFormat.ResourceGuid; resourceGuid != nil {
                v["resource_guid"] = *resourceGuid
            }
            v["virtual_machine"] = flattenArmNetworkSecurityGroupSubResource(interfacePropertiesFormat.VirtualMachine)
        }
        if etag := item.Etag; etag != nil {
            v["etag"] = *etag
        }

        results = append(results, v)
    }

    return results
}

func flattenArmNetworkSecurityGroupSubnet(input *[]network.Subnet) []interface{} {
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
            v["ip_configurations"] = flattenArmNetworkSecurityGroupIPConfiguration(subnetPropertiesFormat.IpConfigurations)
            v["network_security_group"] = flattenArmNetworkSecurityGroupSecurityGroup(subnetPropertiesFormat.NetworkSecurityGroup)
            v["route_table"] = flattenArmNetworkSecurityGroupRouteTable(subnetPropertiesFormat.RouteTable)
        }
        if etag := item.Etag; etag != nil {
            v["etag"] = *etag
        }

        results = append(results, v)
    }

    return results
}

func flattenArmNetworkSecurityGroupInterfaceDnsSettings(input *network.InterfaceDnsSettings) []interface{} {
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

func flattenArmNetworkSecurityGroupInterfaceIPConfiguration(input *[]network.InterfaceIPConfiguration) []interface{} {
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

func flattenArmNetworkSecurityGroupSecurityGroup(input *network.SecurityGroup) []interface{} {
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

func flattenArmNetworkSecurityGroupSubResource(input *network.SubResource) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if id := input.ID; id != nil {
        result["id"] = *id
    }

    return []interface{}{result}
}

func flattenArmNetworkSecurityGroupIPConfiguration(input *[]network.IPConfiguration) []interface{} {
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

func flattenArmNetworkSecurityGroupRouteTable(input *network.RouteTable) []interface{} {
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
