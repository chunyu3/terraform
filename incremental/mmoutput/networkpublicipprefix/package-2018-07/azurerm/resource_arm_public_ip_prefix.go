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



func resourceArmPublicIPPrefix() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmPublicIPPrefixCreateUpdate,
        Read: resourceArmPublicIPPrefixRead,
        Update: resourceArmPublicIPPrefixCreateUpdate,
        Delete: resourceArmPublicIPPrefixDelete,

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

            "etag": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
            },

            "ip_prefix": {
                Type: schema.TypeString,
                Optional: true,
            },

            "ip_tags": {
                Type: schema.TypeList,
                Optional: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "ip_tag_type": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "tag": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },

            "prefix_length": {
                Type: schema.TypeInt,
                Optional: true,
            },

            "public_ip_address_version": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(network.IPv4),
                    string(network.IPv6),
                }, false),
                Default: string(network.IPv4),
            },

            "public_ip_addresses": {
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

            "sku": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "name": {
                            Type: schema.TypeString,
                            Optional: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(network.Standard),
                            }, false),
                            Default: string(network.Standard),
                        },
                    },
                },
            },

            "zones": {
                Type: schema.TypeList,
                Optional: true,
                ForceNew: true,
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
            },

            "load_balancer_frontend_ip_configuration": {
                Type: schema.TypeList,
                Computed: true,
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

func resourceArmPublicIPPrefixCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).publicIPPrefixesClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Public Ip Prefix %q (Resource Group %q): %+v", name, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_public_ip_prefix", *existing.ID)
        }
    }

    id := d.Get("id").(string)
    location := azure.NormalizeLocation(d.Get("location").(string))
    etag := d.Get("etag").(string)
    ipPrefix := d.Get("ip_prefix").(string)
    ipTags := d.Get("ip_tags").([]interface{})
    prefixLength := d.Get("prefix_length").(int)
    publicIpAddressVersion := d.Get("public_ip_address_version").(string)
    publicIpAddresses := d.Get("public_ip_addresses").([]interface{})
    resourceGuid := d.Get("resource_guid").(string)
    sku := d.Get("sku").([]interface{})
    zones := d.Get("zones").([]interface{})
    t := d.Get("tags").(map[string]interface{})

    parameters := network.PublicIPPrefix{
        Etag: utils.String(etag),
        ID: utils.String(id),
        Location: utils.String(location),
        PublicIPPrefixPropertiesFormat: &network.PublicIPPrefixPropertiesFormat{
            IpPrefix: utils.String(ipPrefix),
            IpTags: expandArmPublicIPPrefixIpTag(ipTags),
            PrefixLength: utils.Int32(int32(prefixLength)),
            PublicIpAddressVersion: network.IPVersion(publicIpAddressVersion),
            PublicIpAddresses: expandArmPublicIPPrefixReferencedPublicIpAddress(publicIpAddresses),
            ResourceGuid: utils.String(resourceGuid),
        },
        Sku: expandArmPublicIPPrefixPublicIPPrefixSku(sku),
        Tags: tags.Expand(t),
        Zones: utils.ExpandStringSlice(zones),
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, name, parameters)
    if err != nil {
        return fmt.Errorf("Error creating Public Ip Prefix %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Public Ip Prefix %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Public Ip Prefix %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Public Ip Prefix %q (Resource Group %q) ID", name, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmPublicIPPrefixRead(d, meta)
}

func resourceArmPublicIPPrefixRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).publicIPPrefixesClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["publicIPPrefixes"]

    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Public Ip Prefix %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Public Ip Prefix %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    d.Set("etag", resp.Etag)
    if publicIPPrefixPropertiesFormat := resp.PublicIPPrefixPropertiesFormat; publicIPPrefixPropertiesFormat != nil {
        d.Set("ip_prefix", publicIPPrefixPropertiesFormat.IpPrefix)
        if err := d.Set("ip_tags", flattenArmPublicIPPrefixIpTag(publicIPPrefixPropertiesFormat.IpTags)); err != nil {
            return fmt.Errorf("Error setting `ip_tags`: %+v", err)
        }
        if err := d.Set("load_balancer_frontend_ip_configuration", flattenArmPublicIPPrefixSubResource(publicIPPrefixPropertiesFormat.LoadBalancerFrontendIpConfiguration)); err != nil {
            return fmt.Errorf("Error setting `load_balancer_frontend_ip_configuration`: %+v", err)
        }
        d.Set("prefix_length", int(*publicIPPrefixPropertiesFormat.PrefixLength))
        d.Set("provisioning_state", publicIPPrefixPropertiesFormat.ProvisioningState)
        d.Set("public_ip_address_version", string(publicIPPrefixPropertiesFormat.PublicIpAddressVersion))
        if err := d.Set("public_ip_addresses", flattenArmPublicIPPrefixReferencedPublicIpAddress(publicIPPrefixPropertiesFormat.PublicIpAddresses)); err != nil {
            return fmt.Errorf("Error setting `public_ip_addresses`: %+v", err)
        }
        d.Set("resource_guid", publicIPPrefixPropertiesFormat.ResourceGuid)
    }
    if err := d.Set("sku", flattenArmPublicIPPrefixPublicIPPrefixSku(resp.Sku)); err != nil {
        return fmt.Errorf("Error setting `sku`: %+v", err)
    }
    d.Set("type", resp.Type)
    d.Set("zones", utils.FlattenStringSlice(resp.Zones))

    return tags.FlattenAndSet(d, resp.Tags)
}


func resourceArmPublicIPPrefixDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).publicIPPrefixesClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["publicIPPrefixes"]

    future, err := client.Delete(ctx, resourceGroup, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Public Ip Prefix %q (Resource Group %q): %+v", name, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Public Ip Prefix %q (Resource Group %q): %+v", name, resourceGroup, err)
        }
    }

    return nil
}

func expandArmPublicIPPrefixIpTag(input []interface{}) *[]network.IpTag {
    results := make([]network.IpTag, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        ipTagType := v["ip_tag_type"].(string)
        tag := v["tag"].(string)

        result := network.IpTag{
            IpTagType: utils.String(ipTagType),
            Tag: utils.String(tag),
        }

        results = append(results, result)
    }
    return &results
}

func expandArmPublicIPPrefixReferencedPublicIpAddress(input []interface{}) *[]network.ReferencedPublicIpAddress {
    results := make([]network.ReferencedPublicIpAddress, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        id := v["id"].(string)

        result := network.ReferencedPublicIpAddress{
            ID: utils.String(id),
        }

        results = append(results, result)
    }
    return &results
}

func expandArmPublicIPPrefixPublicIPPrefixSku(input []interface{}) *network.PublicIPPrefixSku {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    name := v["name"].(string)

    result := network.PublicIPPrefixSku{
        Name: network.PublicIPPrefixSkuName(name),
    }
    return &result
}


func flattenArmPublicIPPrefixIpTag(input *[]network.IpTag) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        if ipTagType := item.IpTagType; ipTagType != nil {
            v["ip_tag_type"] = *ipTagType
        }
        if tag := item.Tag; tag != nil {
            v["tag"] = *tag
        }

        results = append(results, v)
    }

    return results
}

func flattenArmPublicIPPrefixSubResource(input *network.SubResource) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})


    return []interface{}{result}
}

func flattenArmPublicIPPrefixReferencedPublicIpAddress(input *[]network.ReferencedPublicIpAddress) []interface{} {
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

func flattenArmPublicIPPrefixPublicIPPrefixSku(input *network.PublicIPPrefixSku) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    result["name"] = string(input.Name)

    return []interface{}{result}
}
