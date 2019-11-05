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



func resourceArmPublicIPAddresse() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmPublicIPAddresseCreateUpdate,
        Read: resourceArmPublicIPAddresseRead,
        Update: resourceArmPublicIPAddresseCreateUpdate,
        Delete: resourceArmPublicIPAddresseDelete,

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

            "public_ip_address_name": {
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
                        "domain_name_label": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "fqdn": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "reverse_fqdn": {
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

            "idle_timeout_in_minutes": {
                Type: schema.TypeInt,
                Optional: true,
            },

            "ip_address": {
                Type: schema.TypeString,
                Optional: true,
            },

            "ip_configuration": {
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

            "public_ipallocation_method": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(network.Static),
                    string(network.Dynamic),
                }, false),
                Default: string(network.Static),
            },

            "resource_guid": {
                Type: schema.TypeString,
                Optional: true,
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

func resourceArmPublicIPAddresseCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).publicIPAddressesClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    publicIpAddressName := d.Get("public_ip_address_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, publicIpAddressName)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Public Ip Addresse (Public Ip Address Name %q / Resource Group %q): %+v", publicIpAddressName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_public_ip_addresse", *existing.ID)
        }
    }

    id := d.Get("id").(string)
    location := azure.NormalizeLocation(d.Get("location").(string))
    dnsSettings := d.Get("dns_settings").([]interface{})
    etag := d.Get("etag").(string)
    idleTimeoutInMinutes := d.Get("idle_timeout_in_minutes").(int)
    ipAddress := d.Get("ip_address").(string)
    ipConfiguration := d.Get("ip_configuration").([]interface{})
    publicIpallocationMethod := d.Get("public_ipallocation_method").(string)
    resourceGuid := d.Get("resource_guid").(string)
    t := d.Get("tags").(map[string]interface{})

    parameters := network.PublicIPAddress{
        Etag: utils.String(etag),
        ID: utils.String(id),
        Location: utils.String(location),
        PublicIPAddressPropertiesFormat: &network.PublicIPAddressPropertiesFormat{
            DnsSettings: expandArmPublicIPAddressePublicIPAddressDnsSettings(dnsSettings),
            IdleTimeoutInMinutes: utils.Int32(int32(idleTimeoutInMinutes)),
            IpAddress: utils.String(ipAddress),
            IpConfiguration: expandArmPublicIPAddresseIPConfiguration(ipConfiguration),
            PublicIpallocationMethod: network.IPAllocationMethod(publicIpallocationMethod),
            ResourceGuid: utils.String(resourceGuid),
        },
        Tags: tags.Expand(t),
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, publicIpAddressName, parameters)
    if err != nil {
        return fmt.Errorf("Error creating Public Ip Addresse (Public Ip Address Name %q / Resource Group %q): %+v", publicIpAddressName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Public Ip Addresse (Public Ip Address Name %q / Resource Group %q): %+v", publicIpAddressName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, publicIpAddressName)
    if err != nil {
        return fmt.Errorf("Error retrieving Public Ip Addresse (Public Ip Address Name %q / Resource Group %q): %+v", publicIpAddressName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Public Ip Addresse (Public Ip Address Name %q / Resource Group %q) ID", publicIpAddressName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmPublicIPAddresseRead(d, meta)
}

func resourceArmPublicIPAddresseRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).publicIPAddressesClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    publicIpAddressName := id.Path["publicIPAddresses"]

    resp, err := client.Get(ctx, resourceGroup, publicIpAddressName)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Public Ip Addresse %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Public Ip Addresse (Public Ip Address Name %q / Resource Group %q): %+v", publicIpAddressName, resourceGroup, err)
    }


    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if publicIPAddressPropertiesFormat := resp.PublicIPAddressPropertiesFormat; publicIPAddressPropertiesFormat != nil {
        if err := d.Set("dns_settings", flattenArmPublicIPAddressePublicIPAddressDnsSettings(publicIPAddressPropertiesFormat.DnsSettings)); err != nil {
            return fmt.Errorf("Error setting `dns_settings`: %+v", err)
        }
        d.Set("idle_timeout_in_minutes", int(*publicIPAddressPropertiesFormat.IdleTimeoutInMinutes))
        d.Set("ip_address", publicIPAddressPropertiesFormat.IpAddress)
        if err := d.Set("ip_configuration", flattenArmPublicIPAddresseIPConfiguration(publicIPAddressPropertiesFormat.IpConfiguration)); err != nil {
            return fmt.Errorf("Error setting `ip_configuration`: %+v", err)
        }
        d.Set("provisioning_state", publicIPAddressPropertiesFormat.ProvisioningState)
        d.Set("public_ipallocation_method", string(publicIPAddressPropertiesFormat.PublicIpallocationMethod))
        d.Set("resource_guid", publicIPAddressPropertiesFormat.ResourceGuid)
    }
    d.Set("etag", resp.Etag)
    d.Set("public_ip_address_name", publicIpAddressName)
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}


func resourceArmPublicIPAddresseDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).publicIPAddressesClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    publicIpAddressName := id.Path["publicIPAddresses"]

    future, err := client.Delete(ctx, resourceGroup, publicIpAddressName)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Public Ip Addresse (Public Ip Address Name %q / Resource Group %q): %+v", publicIpAddressName, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Public Ip Addresse (Public Ip Address Name %q / Resource Group %q): %+v", publicIpAddressName, resourceGroup, err)
        }
    }

    return nil
}

func expandArmPublicIPAddressePublicIPAddressDnsSettings(input []interface{}) *network.PublicIPAddressDnsSettings {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    domainNameLabel := v["domain_name_label"].(string)
    fqdn := v["fqdn"].(string)
    reverseFqdn := v["reverse_fqdn"].(string)

    result := network.PublicIPAddressDnsSettings{
        DomainNameLabel: utils.String(domainNameLabel),
        Fqdn: utils.String(fqdn),
        ReverseFqdn: utils.String(reverseFqdn),
    }
    return &result
}

func expandArmPublicIPAddresseIPConfiguration(input []interface{}) *network.IPConfiguration {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    id := v["id"].(string)
    privateIpAddress := v["private_ip_address"].(string)
    privateIpallocationMethod := v["private_ipallocation_method"].(string)
    subnet := v["subnet"].([]interface{})
    publicIpAddress := v["public_ip_address"].([]interface{})
    name := v["name"].(string)
    etag := v["etag"].(string)

    result := network.IPConfiguration{
        Etag: utils.String(etag),
        ID: utils.String(id),
        Name: utils.String(name),
        IPConfigurationPropertiesFormat: &network.IPConfigurationPropertiesFormat{
            PrivateIpAddress: utils.String(privateIpAddress),
            PrivateIpallocationMethod: network.IPAllocationMethod(privateIpallocationMethod),
            PublicIpAddress: expandArmPublicIPAddressePublicIPAddress(publicIpAddress),
            Subnet: expandArmPublicIPAddresseSubnet(subnet),
        },
    }
    return &result
}

func expandArmPublicIPAddressePublicIPAddress(input []interface{}) *network.PublicIPAddress {
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

func expandArmPublicIPAddresseSubnet(input []interface{}) *network.Subnet {
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


func flattenArmPublicIPAddressePublicIPAddressDnsSettings(input *network.PublicIPAddressDnsSettings) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if domainNameLabel := input.DomainNameLabel; domainNameLabel != nil {
        result["domain_name_label"] = *domainNameLabel
    }
    if fqdn := input.Fqdn; fqdn != nil {
        result["fqdn"] = *fqdn
    }
    if reverseFqdn := input.ReverseFqdn; reverseFqdn != nil {
        result["reverse_fqdn"] = *reverseFqdn
    }

    return []interface{}{result}
}

func flattenArmPublicIPAddresseIPConfiguration(input *network.IPConfiguration) []interface{} {
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
    if iPConfigurationPropertiesFormat := input.IPConfigurationPropertiesFormat; iPConfigurationPropertiesFormat != nil {
        if privateIpAddress := iPConfigurationPropertiesFormat.PrivateIpAddress; privateIpAddress != nil {
            result["private_ip_address"] = *privateIpAddress
        }
        result["private_ipallocation_method"] = string(iPConfigurationPropertiesFormat.PrivateIpallocationMethod)
        result["public_ip_address"] = flattenArmPublicIPAddressePublicIPAddress(iPConfigurationPropertiesFormat.PublicIpAddress)
        result["subnet"] = flattenArmPublicIPAddresseSubnet(iPConfigurationPropertiesFormat.Subnet)
    }

    return []interface{}{result}
}

func flattenArmPublicIPAddressePublicIPAddress(input *network.PublicIPAddress) []interface{} {
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

func flattenArmPublicIPAddresseSubnet(input *network.Subnet) []interface{} {
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
