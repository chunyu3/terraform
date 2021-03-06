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



func resourceArmAzureFirewall() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmAzureFirewallCreateUpdate,
        Read: resourceArmAzureFirewallRead,
        Update: resourceArmAzureFirewallCreateUpdate,
        Delete: resourceArmAzureFirewallDelete,

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

            "azure_firewall_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "application_rule_collections": {
                Type: schema.TypeList,
                Optional: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "action": {
                            Type: schema.TypeList,
                            Optional: true,
                            MaxItems: 1,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "type": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                        ValidateFunc: validation.StringInSlice([]string{
                                            string(network.Allow),
                                            string(network.Deny),
                                        }, false),
                                        Default: string(network.Allow),
                                    },
                                },
                            },
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
                        "rules": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "description": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                    "name": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                    "source_addresses": {
                                        Type: schema.TypeList,
                                        Optional: true,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                        },
                                    },
                                    "target_urls": {
                                        Type: schema.TypeList,
                                        Optional: true,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                        },
                                    },
                                },
                            },
                        },
                    },
                },
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
                        "internal_public_ip_address": {
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
                        "name": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "private_ip_address": {
                            Type: schema.TypeString,
                            Optional: true,
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

            "network_rule_collections": {
                Type: schema.TypeList,
                Optional: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "action": {
                            Type: schema.TypeList,
                            Optional: true,
                            MaxItems: 1,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "type": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                        ValidateFunc: validation.StringInSlice([]string{
                                            string(network.Allow),
                                            string(network.Deny),
                                        }, false),
                                        Default: string(network.Allow),
                                    },
                                },
                            },
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
                        "rules": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "description": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                    "destination_addresses": {
                                        Type: schema.TypeList,
                                        Optional: true,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                        },
                                    },
                                    "destination_ports": {
                                        Type: schema.TypeList,
                                        Optional: true,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                        },
                                    },
                                    "name": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                    "protocols": {
                                        Type: schema.TypeList,
                                        Optional: true,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                            ValidateFunc: validation.StringInSlice([]string{
                                                string(network.TCP),
                                                string(network.UDP),
                                                string(network.Any),
                                                string(network.ICMP),
                                           }, false),
                                        },
                                    },
                                    "source_addresses": {
                                        Type: schema.TypeList,
                                        Optional: true,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                        },
                                    },
                                },
                            },
                        },
                    },
                },
            },

            "etag": {
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

            "tags": tags.Schema(),
        },
    }
}

func resourceArmAzureFirewallCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).azureFirewallsClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    azureFirewallName := d.Get("azure_firewall_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, azureFirewallName)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Azure Firewall (Azure Firewall Name %q / Resource Group %q): %+v", azureFirewallName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_azure_firewall", *existing.ID)
        }
    }

    id := d.Get("id").(string)
    location := azure.NormalizeLocation(d.Get("location").(string))
    applicationRuleCollections := d.Get("application_rule_collections").([]interface{})
    ipConfigurations := d.Get("ip_configurations").([]interface{})
    networkRuleCollections := d.Get("network_rule_collections").([]interface{})
    t := d.Get("tags").(map[string]interface{})

    parameters := network.AzureFirewall{
        ID: utils.String(id),
        Location: utils.String(location),
        AzureFirewallPropertiesFormat: &network.AzureFirewallPropertiesFormat{
            ApplicationRuleCollections: expandArmAzureFirewallAzureFirewallApplicationRuleCollection(applicationRuleCollections),
            IpConfigurations: expandArmAzureFirewallAzureFirewallIPConfiguration(ipConfigurations),
            NetworkRuleCollections: expandArmAzureFirewallAzureFirewallNetworkRuleCollection(networkRuleCollections),
        },
        Tags: tags.Expand(t),
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, azureFirewallName, parameters)
    if err != nil {
        return fmt.Errorf("Error creating Azure Firewall (Azure Firewall Name %q / Resource Group %q): %+v", azureFirewallName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Azure Firewall (Azure Firewall Name %q / Resource Group %q): %+v", azureFirewallName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, azureFirewallName)
    if err != nil {
        return fmt.Errorf("Error retrieving Azure Firewall (Azure Firewall Name %q / Resource Group %q): %+v", azureFirewallName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Azure Firewall (Azure Firewall Name %q / Resource Group %q) ID", azureFirewallName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmAzureFirewallRead(d, meta)
}

func resourceArmAzureFirewallRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).azureFirewallsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    azureFirewallName := id.Path["azureFirewalls"]

    resp, err := client.Get(ctx, resourceGroup, azureFirewallName)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Azure Firewall %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Azure Firewall (Azure Firewall Name %q / Resource Group %q): %+v", azureFirewallName, resourceGroup, err)
    }


    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if azureFirewallPropertiesFormat := resp.AzureFirewallPropertiesFormat; azureFirewallPropertiesFormat != nil {
        if err := d.Set("application_rule_collections", flattenArmAzureFirewallAzureFirewallApplicationRuleCollection(azureFirewallPropertiesFormat.ApplicationRuleCollections)); err != nil {
            return fmt.Errorf("Error setting `application_rule_collections`: %+v", err)
        }
        if err := d.Set("ip_configurations", flattenArmAzureFirewallAzureFirewallIPConfiguration(azureFirewallPropertiesFormat.IpConfigurations)); err != nil {
            return fmt.Errorf("Error setting `ip_configurations`: %+v", err)
        }
        if err := d.Set("network_rule_collections", flattenArmAzureFirewallAzureFirewallNetworkRuleCollection(azureFirewallPropertiesFormat.NetworkRuleCollections)); err != nil {
            return fmt.Errorf("Error setting `network_rule_collections`: %+v", err)
        }
        d.Set("provisioning_state", string(azureFirewallPropertiesFormat.ProvisioningState))
    }
    d.Set("azure_firewall_name", azureFirewallName)
    d.Set("etag", resp.Etag)
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}


func resourceArmAzureFirewallDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).azureFirewallsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    azureFirewallName := id.Path["azureFirewalls"]

    future, err := client.Delete(ctx, resourceGroup, azureFirewallName)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Azure Firewall (Azure Firewall Name %q / Resource Group %q): %+v", azureFirewallName, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Azure Firewall (Azure Firewall Name %q / Resource Group %q): %+v", azureFirewallName, resourceGroup, err)
        }
    }

    return nil
}

func expandArmAzureFirewallAzureFirewallApplicationRuleCollection(input []interface{}) *[]network.AzureFirewallApplicationRuleCollection {
    results := make([]network.AzureFirewallApplicationRuleCollection, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        id := v["id"].(string)
        priority := v["priority"].(int)
        action := v["action"].([]interface{})
        rules := v["rules"].([]interface{})
        name := v["name"].(string)

        result := network.AzureFirewallApplicationRuleCollection{
            ID: utils.String(id),
            Name: utils.String(name),
            AzureFirewallApplicationRuleCollectionPropertiesFormat: &network.AzureFirewallApplicationRuleCollectionPropertiesFormat{
                Action: expandArmAzureFirewallAzureFirewallRCAction(action),
                Priority: utils.Int32(int32(priority)),
                Rules: expandArmAzureFirewallAzureFirewallApplicationRule(rules),
            },
        }

        results = append(results, result)
    }
    return &results
}

func expandArmAzureFirewallAzureFirewallIPConfiguration(input []interface{}) *[]network.AzureFirewallIPConfiguration {
    results := make([]network.AzureFirewallIPConfiguration, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        id := v["id"].(string)
        privateIpAddress := v["private_ip_address"].(string)
        subnet := v["subnet"].([]interface{})
        internalPublicIpAddress := v["internal_public_ip_address"].([]interface{})
        publicIpAddress := v["public_ip_address"].([]interface{})
        name := v["name"].(string)
        etag := v["etag"].(string)

        result := network.AzureFirewallIPConfiguration{
            Etag: utils.String(etag),
            ID: utils.String(id),
            Name: utils.String(name),
            AzureFirewallIPConfigurationPropertiesFormat: &network.AzureFirewallIPConfigurationPropertiesFormat{
                InternalPublicIpAddress: expandArmAzureFirewallSubResource(internalPublicIpAddress),
                PrivateIpAddress: utils.String(privateIpAddress),
                PublicIpAddress: expandArmAzureFirewallSubResource(publicIpAddress),
                Subnet: expandArmAzureFirewallSubResource(subnet),
            },
        }

        results = append(results, result)
    }
    return &results
}

func expandArmAzureFirewallAzureFirewallNetworkRuleCollection(input []interface{}) *[]network.AzureFirewallNetworkRuleCollection {
    results := make([]network.AzureFirewallNetworkRuleCollection, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        id := v["id"].(string)
        priority := v["priority"].(int)
        action := v["action"].([]interface{})
        rules := v["rules"].([]interface{})
        name := v["name"].(string)

        result := network.AzureFirewallNetworkRuleCollection{
            ID: utils.String(id),
            Name: utils.String(name),
            AzureFirewallNetworkRuleCollectionPropertiesFormat: &network.AzureFirewallNetworkRuleCollectionPropertiesFormat{
                Action: expandArmAzureFirewallAzureFirewallRCAction(action),
                Priority: utils.Int32(int32(priority)),
                Rules: expandArmAzureFirewallAzureFirewallNetworkRule(rules),
            },
        }

        results = append(results, result)
    }
    return &results
}

func expandArmAzureFirewallAzureFirewallRCAction(input []interface{}) *network.AzureFirewallRCAction {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    type := v["type"].(string)

    result := network.AzureFirewallRCAction{
        Type: network.AzureFirewallRCActionType(type),
    }
    return &result
}

func expandArmAzureFirewallAzureFirewallApplicationRule(input []interface{}) *[]network.AzureFirewallApplicationRule {
    results := make([]network.AzureFirewallApplicationRule, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        name := v["name"].(string)
        description := v["description"].(string)
        sourceAddresses := v["source_addresses"].([]interface{})
        targetUrls := v["target_urls"].([]interface{})

        result := network.AzureFirewallApplicationRule{
            Description: utils.String(description),
            Name: utils.String(name),
            SourceAddresses: utils.ExpandStringSlice(sourceAddresses),
            TargetUrls: utils.ExpandStringSlice(targetUrls),
        }

        results = append(results, result)
    }
    return &results
}

func expandArmAzureFirewallSubResource(input []interface{}) *network.SubResource {
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

func expandArmAzureFirewallAzureFirewallNetworkRule(input []interface{}) *[]network.AzureFirewallNetworkRule {
    results := make([]network.AzureFirewallNetworkRule, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        name := v["name"].(string)
        description := v["description"].(string)
        protocols := v["protocols"].([]interface{})
        sourceAddresses := v["source_addresses"].([]interface{})
        destinationAddresses := v["destination_addresses"].([]interface{})
        destinationPorts := v["destination_ports"].([]interface{})

        result := network.AzureFirewallNetworkRule{
            Description: utils.String(description),
            DestinationAddresses: utils.ExpandStringSlice(destinationAddresses),
            DestinationPorts: utils.ExpandStringSlice(destinationPorts),
            Name: utils.String(name),
            Protocols: expandArmAzureFirewall(protocols),
            SourceAddresses: utils.ExpandStringSlice(sourceAddresses),
        }

        results = append(results, result)
    }
    return &results
}

func expandArmAzureFirewall(input []interface{}) *[]network. {
    results := make([]network., 0)
    for _, item := range input {
        v := item.(string)
        result := network.(v)
        results = append(results, result)
    }
    return &results
}


func flattenArmAzureFirewallAzureFirewallApplicationRuleCollection(input *[]network.AzureFirewallApplicationRuleCollection) []interface{} {
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
        if azureFirewallApplicationRuleCollectionPropertiesFormat := item.AzureFirewallApplicationRuleCollectionPropertiesFormat; azureFirewallApplicationRuleCollectionPropertiesFormat != nil {
            v["action"] = flattenArmAzureFirewallAzureFirewallRCAction(azureFirewallApplicationRuleCollectionPropertiesFormat.Action)
            if priority := azureFirewallApplicationRuleCollectionPropertiesFormat.Priority; priority != nil {
                v["priority"] = int(*priority)
            }
            v["rules"] = flattenArmAzureFirewallAzureFirewallApplicationRule(azureFirewallApplicationRuleCollectionPropertiesFormat.Rules)
        }

        results = append(results, v)
    }

    return results
}

func flattenArmAzureFirewallAzureFirewallIPConfiguration(input *[]network.AzureFirewallIPConfiguration) []interface{} {
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
        if azureFirewallIPConfigurationPropertiesFormat := item.AzureFirewallIPConfigurationPropertiesFormat; azureFirewallIPConfigurationPropertiesFormat != nil {
            v["internal_public_ip_address"] = flattenArmAzureFirewallSubResource(azureFirewallIPConfigurationPropertiesFormat.InternalPublicIpAddress)
            if privateIpAddress := azureFirewallIPConfigurationPropertiesFormat.PrivateIpAddress; privateIpAddress != nil {
                v["private_ip_address"] = *privateIpAddress
            }
            v["public_ip_address"] = flattenArmAzureFirewallSubResource(azureFirewallIPConfigurationPropertiesFormat.PublicIpAddress)
            v["subnet"] = flattenArmAzureFirewallSubResource(azureFirewallIPConfigurationPropertiesFormat.Subnet)
        }

        results = append(results, v)
    }

    return results
}

func flattenArmAzureFirewallAzureFirewallNetworkRuleCollection(input *[]network.AzureFirewallNetworkRuleCollection) []interface{} {
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
        if azureFirewallNetworkRuleCollectionPropertiesFormat := item.AzureFirewallNetworkRuleCollectionPropertiesFormat; azureFirewallNetworkRuleCollectionPropertiesFormat != nil {
            v["action"] = flattenArmAzureFirewallAzureFirewallRCAction(azureFirewallNetworkRuleCollectionPropertiesFormat.Action)
            if priority := azureFirewallNetworkRuleCollectionPropertiesFormat.Priority; priority != nil {
                v["priority"] = int(*priority)
            }
            v["rules"] = flattenArmAzureFirewallAzureFirewallNetworkRule(azureFirewallNetworkRuleCollectionPropertiesFormat.Rules)
        }

        results = append(results, v)
    }

    return results
}

func flattenArmAzureFirewallAzureFirewallRCAction(input *network.AzureFirewallRCAction) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    result["type"] = string(input.Type)

    return []interface{}{result}
}

func flattenArmAzureFirewallAzureFirewallApplicationRule(input *[]network.AzureFirewallApplicationRule) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        if name := item.Name; name != nil {
            v["name"] = *name
        }
        if description := item.Description; description != nil {
            v["description"] = *description
        }
        v["source_addresses"] = utils.FlattenStringSlice(item.SourceAddresses)
        v["target_urls"] = utils.FlattenStringSlice(item.TargetUrls)

        results = append(results, v)
    }

    return results
}

func flattenArmAzureFirewallSubResource(input *network.SubResource) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if id := input.ID; id != nil {
        result["id"] = *id
    }

    return []interface{}{result}
}

func flattenArmAzureFirewallAzureFirewallNetworkRule(input *[]network.AzureFirewallNetworkRule) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        if name := item.Name; name != nil {
            v["name"] = *name
        }
        if description := item.Description; description != nil {
            v["description"] = *description
        }
        v["destination_addresses"] = utils.FlattenStringSlice(item.DestinationAddresses)
        v["destination_ports"] = utils.FlattenStringSlice(item.DestinationPorts)
        v["protocols"] = flattenArmAzureFirewall(string(item.Protocols))
        v["source_addresses"] = utils.FlattenStringSlice(item.SourceAddresses)

        results = append(results, v)
    }

    return results
}

func flattenArmAzureFirewall(input *[]network.) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        result := string(item)
        results = append(results, result)
    }

    return results
}
