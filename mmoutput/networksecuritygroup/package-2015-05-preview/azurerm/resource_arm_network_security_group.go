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
                Computed: true,
            },

            "location": azure.SchemaLocation(),

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "network_security_group_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

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

func resourceArmNetworkSecurityGroupCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).networkSecurityGroupsClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    networkSecurityGroupName := d.Get("network_security_group_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, networkSecurityGroupName)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Network Security Group (Network Security Group Name %q / Resource Group %q): %+v", networkSecurityGroupName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_network_security_group", *existing.ID)
        }
    }

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
        Location: utils.String(location),
        SecurityGroupPropertiesFormat: &network.SecurityGroupPropertiesFormat{
            DefaultSecurityRules: expandArmNetworkSecurityGroupSecurityRule(defaultSecurityRules),
            NetworkInterfaces: expandArmNetworkSecurityGroupSubResource(networkInterfaces),
            ResourceGuid: utils.String(resourceGuid),
            SecurityRules: expandArmNetworkSecurityGroupSecurityRule(securityRules),
            Subnets: expandArmNetworkSecurityGroupSubResource(subnets),
        },
        Tags: tags.Expand(t),
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, networkSecurityGroupName, parameters)
    if err != nil {
        return fmt.Errorf("Error creating Network Security Group (Network Security Group Name %q / Resource Group %q): %+v", networkSecurityGroupName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Network Security Group (Network Security Group Name %q / Resource Group %q): %+v", networkSecurityGroupName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, networkSecurityGroupName)
    if err != nil {
        return fmt.Errorf("Error retrieving Network Security Group (Network Security Group Name %q / Resource Group %q): %+v", networkSecurityGroupName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Network Security Group (Network Security Group Name %q / Resource Group %q) ID", networkSecurityGroupName, resourceGroup)
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
    networkSecurityGroupName := id.Path["networkSecurityGroups"]

    resp, err := client.Get(ctx, resourceGroup, networkSecurityGroupName)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Network Security Group %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Network Security Group (Network Security Group Name %q / Resource Group %q): %+v", networkSecurityGroupName, resourceGroup, err)
    }


    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if securityGroupPropertiesFormat := resp.SecurityGroupPropertiesFormat; securityGroupPropertiesFormat != nil {
        if err := d.Set("default_security_rules", flattenArmNetworkSecurityGroupSecurityRule(securityGroupPropertiesFormat.DefaultSecurityRules)); err != nil {
            return fmt.Errorf("Error setting `default_security_rules`: %+v", err)
        }
        if err := d.Set("network_interfaces", flattenArmNetworkSecurityGroupSubResource(securityGroupPropertiesFormat.NetworkInterfaces)); err != nil {
            return fmt.Errorf("Error setting `network_interfaces`: %+v", err)
        }
        d.Set("provisioning_state", securityGroupPropertiesFormat.ProvisioningState)
        d.Set("resource_guid", securityGroupPropertiesFormat.ResourceGuid)
        if err := d.Set("security_rules", flattenArmNetworkSecurityGroupSecurityRule(securityGroupPropertiesFormat.SecurityRules)); err != nil {
            return fmt.Errorf("Error setting `security_rules`: %+v", err)
        }
        if err := d.Set("subnets", flattenArmNetworkSecurityGroupSubResource(securityGroupPropertiesFormat.Subnets)); err != nil {
            return fmt.Errorf("Error setting `subnets`: %+v", err)
        }
    }
    d.Set("etag", resp.Etag)
    d.Set("network_security_group_name", networkSecurityGroupName)
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
    networkSecurityGroupName := id.Path["networkSecurityGroups"]

    future, err := client.Delete(ctx, resourceGroup, networkSecurityGroupName)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Network Security Group (Network Security Group Name %q / Resource Group %q): %+v", networkSecurityGroupName, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Network Security Group (Network Security Group Name %q / Resource Group %q): %+v", networkSecurityGroupName, resourceGroup, err)
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

func expandArmNetworkSecurityGroupSubResource(input []interface{}) *[]network.SubResource {
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

func flattenArmNetworkSecurityGroupSubResource(input *[]network.SubResource) []interface{} {
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
