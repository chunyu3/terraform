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
                        "description": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "destination_address_prefix": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "destination_address_prefixes": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                        "destination_port_range": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "destination_port_ranges": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
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
                        "name": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "priority": {
                            Type: schema.TypeInt,
                            Optional: true,
                        },
                        "source_address_prefix": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "source_address_prefixes": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                        "source_port_range": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "source_port_ranges": {
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
                        "description": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "destination_address_prefix": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "destination_address_prefixes": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                        "destination_port_range": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "destination_port_ranges": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
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
                        "name": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "priority": {
                            Type: schema.TypeInt,
                            Optional: true,
                        },
                        "source_address_prefix": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "source_address_prefixes": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                        "source_port_range": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "source_port_ranges": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                    },
                },
            },

            "network_interfaces": {
                Type: schema.TypeList,
                Computed: true,
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
                                    "name": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                    "tags": tags.Schema(),
                                    "type": {
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
                        "provisioning_state": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "resource_guid": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "tags": tags.Schema(),
                        "type": {
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
                    },
                },
            },

            "provisioning_state": {
                Type: schema.TypeString,
                Computed: true,
            },

            "subnets": {
                Type: schema.TypeList,
                Computed: true,
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
                                    "name": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                    "tags": tags.Schema(),
                                    "type": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                },
                            },
                        },
                        "provisioning_state": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "resource_navigation_links": {
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
                                    "name": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                    "tags": tags.Schema(),
                                    "type": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
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
                                    "provisioning_state": {
                                        Type: schema.TypeString,
                                        Optional: true,
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
    resourceGuid := d.Get("resource_guid").(string)
    securityRules := d.Get("security_rules").([]interface{})
    t := d.Get("tags").(map[string]interface{})

    parameters := network.SecurityGroup{
        Etag: utils.String(etag),
        ID: utils.String(id),
        Location: utils.String(location),
        SecurityGroupPropertiesFormat: &network.SecurityGroupPropertiesFormat{
            DefaultSecurityRules: expandArmNetworkSecurityGroupSecurityRule(defaultSecurityRules),
            ResourceGuid: utils.String(resourceGuid),
            SecurityRules: expandArmNetworkSecurityGroupSecurityRule(securityRules),
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
        sourceAddressPrefixes := v["source_address_prefixes"].([]interface{})
        destinationAddressPrefix := v["destination_address_prefix"].(string)
        destinationAddressPrefixes := v["destination_address_prefixes"].([]interface{})
        sourcePortRanges := v["source_port_ranges"].([]interface{})
        destinationPortRanges := v["destination_port_ranges"].([]interface{})
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
                DestinationAddressPrefixes: utils.ExpandStringSlice(destinationAddressPrefixes),
                DestinationPortRange: utils.String(destinationPortRange),
                DestinationPortRanges: utils.ExpandStringSlice(destinationPortRanges),
                Direction: network.SecurityRuleDirection(direction),
                Priority: utils.Int32(int32(priority)),
                Protocol: network.SecurityRuleProtocol(protocol),
                SourceAddressPrefix: utils.String(sourceAddressPrefix),
                SourceAddressPrefixes: utils.ExpandStringSlice(sourceAddressPrefixes),
                SourcePortRange: utils.String(sourcePortRange),
                SourcePortRanges: utils.ExpandStringSlice(sourcePortRanges),
            },
        }

        results = append(results, result)
    }
    return &results
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
            v["destination_address_prefixes"] = utils.FlattenStringSlice(securityRulePropertiesFormat.DestinationAddressPrefixes)
            if destinationPortRange := securityRulePropertiesFormat.DestinationPortRange; destinationPortRange != nil {
                v["destination_port_range"] = *destinationPortRange
            }
            v["destination_port_ranges"] = utils.FlattenStringSlice(securityRulePropertiesFormat.DestinationPortRanges)
            v["direction"] = string(securityRulePropertiesFormat.Direction)
            if priority := securityRulePropertiesFormat.Priority; priority != nil {
                v["priority"] = int(*priority)
            }
            v["protocol"] = string(securityRulePropertiesFormat.Protocol)
            if sourceAddressPrefix := securityRulePropertiesFormat.SourceAddressPrefix; sourceAddressPrefix != nil {
                v["source_address_prefix"] = *sourceAddressPrefix
            }
            v["source_address_prefixes"] = utils.FlattenStringSlice(securityRulePropertiesFormat.SourceAddressPrefixes)
            if sourcePortRange := securityRulePropertiesFormat.SourcePortRange; sourcePortRange != nil {
                v["source_port_range"] = *sourcePortRange
            }
            v["source_port_ranges"] = utils.FlattenStringSlice(securityRulePropertiesFormat.SourcePortRanges)
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


        results = append(results, v)
    }

    return results
}
