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



func resourceArmRouteFilter() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmRouteFilterCreate,
        Read: resourceArmRouteFilterRead,
        Update: resourceArmRouteFilterUpdate,
        Delete: resourceArmRouteFilterDelete,

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

            "peerings": {
                Type: schema.TypeList,
                Optional: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "azure_asn": {
                            Type: schema.TypeInt,
                            Optional: true,
                        },
                        "connections": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
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
                        "gateway_manager_etag": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "id": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "ipv6peering_config": {
                            Type: schema.TypeList,
                            Optional: true,
                            MaxItems: 1,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "primary_peer_address_prefix": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                    "secondary_peer_address_prefix": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                    "state": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                        ValidateFunc: validation.StringInSlice([]string{
                                            string(network.Disabled),
                                            string(network.Enabled),
                                        }, false),
                                        Default: string(network.Disabled),
                                    },
                                },
                            },
                        },
                        "last_modified_by": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "microsoft_peering_config": {
                            Type: schema.TypeList,
                            Optional: true,
                            MaxItems: 1,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "advertised_communities": {
                                        Type: schema.TypeList,
                                        Optional: true,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                        },
                                    },
                                    "advertised_public_prefixes": {
                                        Type: schema.TypeList,
                                        Optional: true,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                        },
                                    },
                                    "advertised_public_prefixes_state": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                        ValidateFunc: validation.StringInSlice([]string{
                                            string(network.NotConfigured),
                                            string(network.Configuring),
                                            string(network.Configured),
                                            string(network.ValidationNeeded),
                                        }, false),
                                        Default: string(network.NotConfigured),
                                    },
                                    "customer_asn": {
                                        Type: schema.TypeInt,
                                        Optional: true,
                                    },
                                    "legacy_mode": {
                                        Type: schema.TypeInt,
                                        Optional: true,
                                    },
                                    "routing_registry_name": {
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
                        "peer_asn": {
                            Type: schema.TypeInt,
                            Optional: true,
                        },
                        "peering_type": {
                            Type: schema.TypeString,
                            Optional: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(network.AzurePublicPeering),
                                string(network.AzurePrivatePeering),
                                string(network.MicrosoftPeering),
                            }, false),
                            Default: string(network.AzurePublicPeering),
                        },
                        "primary_azure_port": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "primary_peer_address_prefix": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "route_filter": {
                            Type: schema.TypeList,
                            Optional: true,
                            MaxItems: 1,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "id": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                    "location": azure.SchemaLocation(),
                                    "tags": tags.Schema(),
                                },
                            },
                        },
                        "secondary_azure_port": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "secondary_peer_address_prefix": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "shared_key": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "state": {
                            Type: schema.TypeString,
                            Optional: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(network.Disabled),
                                string(network.Enabled),
                            }, false),
                            Default: string(network.Disabled),
                        },
                        "stats": {
                            Type: schema.TypeList,
                            Optional: true,
                            MaxItems: 1,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "primarybytes_in": {
                                        Type: schema.TypeInt,
                                        Optional: true,
                                    },
                                    "primarybytes_out": {
                                        Type: schema.TypeInt,
                                        Optional: true,
                                    },
                                    "secondarybytes_in": {
                                        Type: schema.TypeInt,
                                        Optional: true,
                                    },
                                    "secondarybytes_out": {
                                        Type: schema.TypeInt,
                                        Optional: true,
                                    },
                                },
                            },
                        },
                        "vlan_id": {
                            Type: schema.TypeInt,
                            Optional: true,
                        },
                    },
                },
            },

            "rules": {
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
                        "communities": {
                            Type: schema.TypeList,
                            Required: true,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                        "route_filter_rule_type": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
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

func resourceArmRouteFilterCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).routeFiltersClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Route Filter %q (Resource Group %q): %+v", name, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_route_filter", *existing.ID)
        }
    }

    id := d.Get("id").(string)
    location := azure.NormalizeLocation(d.Get("location").(string))
    peerings := d.Get("peerings").([]interface{})
    rules := d.Get("rules").([]interface{})
    t := d.Get("tags").(map[string]interface{})

    routeFilterParameters := network.RouteFilter{
        ID: utils.String(id),
        Location: utils.String(location),
        RouteFilterPropertiesFormat: &network.RouteFilterPropertiesFormat{
            Peerings: expandArmRouteFilterExpressRouteCircuitPeering(peerings),
            Rules: expandArmRouteFilterRouteFilterRule(rules),
        },
        Tags: tags.Expand(t),
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, name, routeFilterParameters)
    if err != nil {
        return fmt.Errorf("Error creating Route Filter %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Route Filter %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Route Filter %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Route Filter %q (Resource Group %q) ID", name, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmRouteFilterRead(d, meta)
}

func resourceArmRouteFilterRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).routeFiltersClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["routeFilters"]

    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Route Filter %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Route Filter %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    d.Set("etag", resp.Etag)
    if routeFilterPropertiesFormat := resp.RouteFilterPropertiesFormat; routeFilterPropertiesFormat != nil {
        if err := d.Set("peerings", flattenArmRouteFilterExpressRouteCircuitPeering(routeFilterPropertiesFormat.Peerings)); err != nil {
            return fmt.Errorf("Error setting `peerings`: %+v", err)
        }
        d.Set("provisioning_state", routeFilterPropertiesFormat.ProvisioningState)
        if err := d.Set("rules", flattenArmRouteFilterRouteFilterRule(routeFilterPropertiesFormat.Rules)); err != nil {
            return fmt.Errorf("Error setting `rules`: %+v", err)
        }
    }
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}

func resourceArmRouteFilterUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).routeFiltersClient
    ctx := meta.(*ArmClient).StopContext

    id := d.Get("id").(string)
    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    peerings := d.Get("peerings").([]interface{})
    rules := d.Get("rules").([]interface{})
    t := d.Get("tags").(map[string]interface{})

    routeFilterParameters := network.RouteFilter{
        ID: utils.String(id),
        Location: utils.String(location),
        RouteFilterPropertiesFormat: &network.RouteFilterPropertiesFormat{
            Peerings: expandArmRouteFilterExpressRouteCircuitPeering(peerings),
            Rules: expandArmRouteFilterRouteFilterRule(rules),
        },
        Tags: tags.Expand(t),
    }


    future, err := client.Update(ctx, resourceGroup, name, routeFilterParameters)
    if err != nil {
        return fmt.Errorf("Error updating Route Filter %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for update of Route Filter %q (Resource Group %q): %+v", name, resourceGroup, err)
    }

    return resourceArmRouteFilterRead(d, meta)
}

func resourceArmRouteFilterDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).routeFiltersClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["routeFilters"]

    future, err := client.Delete(ctx, resourceGroup, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Route Filter %q (Resource Group %q): %+v", name, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Route Filter %q (Resource Group %q): %+v", name, resourceGroup, err)
        }
    }

    return nil
}

func expandArmRouteFilterExpressRouteCircuitPeering(input []interface{}) *[]network.ExpressRouteCircuitPeering {
    results := make([]network.ExpressRouteCircuitPeering, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        id := v["id"].(string)
        peeringType := v["peering_type"].(string)
        state := v["state"].(string)
        azureAsn := v["azure_asn"].(int)
        peerAsn := v["peer_asn"].(int)
        primaryPeerAddressPrefix := v["primary_peer_address_prefix"].(string)
        secondaryPeerAddressPrefix := v["secondary_peer_address_prefix"].(string)
        primaryAzurePort := v["primary_azure_port"].(string)
        secondaryAzurePort := v["secondary_azure_port"].(string)
        sharedKey := v["shared_key"].(string)
        vlanId := v["vlan_id"].(int)
        microsoftPeeringConfig := v["microsoft_peering_config"].([]interface{})
        stats := v["stats"].([]interface{})
        gatewayManagerEtag := v["gateway_manager_etag"].(string)
        lastModifiedBy := v["last_modified_by"].(string)
        routeFilter := v["route_filter"].([]interface{})
        ipv6peeringConfig := v["ipv6peering_config"].([]interface{})
        connections := v["connections"].([]interface{})
        name := v["name"].(string)

        result := network.ExpressRouteCircuitPeering{
            ID: utils.String(id),
            Name: utils.String(name),
            ExpressRouteCircuitPeeringPropertiesFormat: &network.ExpressRouteCircuitPeeringPropertiesFormat{
                AzureAsn: utils.Int32(int32(azureAsn)),
                Connections: expandArmRouteFilterExpressRouteCircuitConnection(connections),
                GatewayManagerEtag: utils.String(gatewayManagerEtag),
                Ipv6peeringConfig: expandArmRouteFilterIpv6ExpressRouteCircuitPeeringConfig(ipv6peeringConfig),
                LastModifiedBy: utils.String(lastModifiedBy),
                MicrosoftPeeringConfig: expandArmRouteFilterExpressRouteCircuitPeeringConfig(microsoftPeeringConfig),
                PeerAsn: utils.Int64(int64(peerAsn)),
                PeeringType: network.ExpressRoutePeeringType(peeringType),
                PrimaryAzurePort: utils.String(primaryAzurePort),
                PrimaryPeerAddressPrefix: utils.String(primaryPeerAddressPrefix),
                RouteFilter: expandArmRouteFilterRouteFilter(routeFilter),
                SecondaryAzurePort: utils.String(secondaryAzurePort),
                SecondaryPeerAddressPrefix: utils.String(secondaryPeerAddressPrefix),
                SharedKey: utils.String(sharedKey),
                State: network.ExpressRoutePeeringState(state),
                Stats: expandArmRouteFilterExpressRouteCircuitStats(stats),
                VlanID: utils.Int32(int32(vlanId)),
            },
        }

        results = append(results, result)
    }
    return &results
}

func expandArmRouteFilterRouteFilterRule(input []interface{}) *[]network.RouteFilterRule {
    results := make([]network.RouteFilterRule, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        id := v["id"].(string)
        access := v["access"].(string)
        routeFilterRuleType := v["route_filter_rule_type"].(string)
        communities := v["communities"].([]interface{})
        name := v["name"].(string)
        location := azure.NormalizeLocation(v["location"].(string))

        result := network.RouteFilterRule{
            ID: utils.String(id),
            Location: utils.String(location),
            Name: utils.String(name),
            RouteFilterRulePropertiesFormat: &network.RouteFilterRulePropertiesFormat{
                Access: network.Access(access),
                Communities: utils.ExpandStringSlice(communities),
                RouteFilterRuleType: utils.String(routeFilterRuleType),
            },
        }

        results = append(results, result)
    }
    return &results
}

func expandArmRouteFilterExpressRouteCircuitConnection(input []interface{}) *[]network.ExpressRouteCircuitConnection {
    results := make([]network.ExpressRouteCircuitConnection, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        id := v["id"].(string)
        name := v["name"].(string)

        result := network.ExpressRouteCircuitConnection{
            ID: utils.String(id),
            Name: utils.String(name),
        }

        results = append(results, result)
    }
    return &results
}

func expandArmRouteFilterIpv6ExpressRouteCircuitPeeringConfig(input []interface{}) *network.Ipv6ExpressRouteCircuitPeeringConfig {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    primaryPeerAddressPrefix := v["primary_peer_address_prefix"].(string)
    secondaryPeerAddressPrefix := v["secondary_peer_address_prefix"].(string)
    state := v["state"].(string)

    result := network.Ipv6ExpressRouteCircuitPeeringConfig{
        PrimaryPeerAddressPrefix: utils.String(primaryPeerAddressPrefix),
        SecondaryPeerAddressPrefix: utils.String(secondaryPeerAddressPrefix),
        State: network.ExpressRouteCircuitPeeringState(state),
    }
    return &result
}

func expandArmRouteFilterExpressRouteCircuitPeeringConfig(input []interface{}) *network.ExpressRouteCircuitPeeringConfig {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    advertisedPublicPrefixes := v["advertised_public_prefixes"].([]interface{})
    advertisedCommunities := v["advertised_communities"].([]interface{})
    advertisedPublicPrefixesState := v["advertised_public_prefixes_state"].(string)
    legacyMode := v["legacy_mode"].(int)
    customerAsn := v["customer_asn"].(int)
    routingRegistryName := v["routing_registry_name"].(string)

    result := network.ExpressRouteCircuitPeeringConfig{
        AdvertisedCommunities: utils.ExpandStringSlice(advertisedCommunities),
        AdvertisedPublicPrefixes: utils.ExpandStringSlice(advertisedPublicPrefixes),
        AdvertisedPublicPrefixesState: network.ExpressRouteCircuitPeeringAdvertisedPublicPrefixState(advertisedPublicPrefixesState),
        CustomerAsn: utils.Int32(int32(customerAsn)),
        LegacyMode: utils.Int32(int32(legacyMode)),
        RoutingRegistryName: utils.String(routingRegistryName),
    }
    return &result
}

func expandArmRouteFilterRouteFilter(input []interface{}) *network.RouteFilter {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    id := v["id"].(string)
    location := azure.NormalizeLocation(v["location"].(string))
    t := v["tags"].(map[string]interface{})

    result := network.RouteFilter{
        ID: utils.String(id),
        Location: utils.String(location),
        Tags: tags.Expand(t),
    }
    return &result
}

func expandArmRouteFilterExpressRouteCircuitStats(input []interface{}) *network.ExpressRouteCircuitStats {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    primarybytesIn := v["primarybytes_in"].(int)
    primarybytesOut := v["primarybytes_out"].(int)
    secondarybytesIn := v["secondarybytes_in"].(int)
    secondarybytesOut := v["secondarybytes_out"].(int)

    result := network.ExpressRouteCircuitStats{
        PrimarybytesIn: utils.Int64(int64(primarybytesIn)),
        PrimarybytesOut: utils.Int64(int64(primarybytesOut)),
        SecondarybytesIn: utils.Int64(int64(secondarybytesIn)),
        SecondarybytesOut: utils.Int64(int64(secondarybytesOut)),
    }
    return &result
}


func flattenArmRouteFilterExpressRouteCircuitPeering(input *[]network.ExpressRouteCircuitPeering) []interface{} {
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
        if expressRouteCircuitPeeringPropertiesFormat := item.ExpressRouteCircuitPeeringPropertiesFormat; expressRouteCircuitPeeringPropertiesFormat != nil {
            if azureAsn := expressRouteCircuitPeeringPropertiesFormat.AzureAsn; azureAsn != nil {
                v["azure_asn"] = int(*azureAsn)
            }
            v["connections"] = flattenArmRouteFilterExpressRouteCircuitConnection(expressRouteCircuitPeeringPropertiesFormat.Connections)
            if gatewayManagerEtag := expressRouteCircuitPeeringPropertiesFormat.GatewayManagerEtag; gatewayManagerEtag != nil {
                v["gateway_manager_etag"] = *gatewayManagerEtag
            }
            v["ipv6peering_config"] = flattenArmRouteFilterIpv6ExpressRouteCircuitPeeringConfig(expressRouteCircuitPeeringPropertiesFormat.Ipv6peeringConfig)
            if lastModifiedBy := expressRouteCircuitPeeringPropertiesFormat.LastModifiedBy; lastModifiedBy != nil {
                v["last_modified_by"] = *lastModifiedBy
            }
            v["microsoft_peering_config"] = flattenArmRouteFilterExpressRouteCircuitPeeringConfig(expressRouteCircuitPeeringPropertiesFormat.MicrosoftPeeringConfig)
            if peerAsn := expressRouteCircuitPeeringPropertiesFormat.PeerAsn; peerAsn != nil {
                v["peer_asn"] = int(*peerAsn)
            }
            v["peering_type"] = string(expressRouteCircuitPeeringPropertiesFormat.PeeringType)
            if primaryAzurePort := expressRouteCircuitPeeringPropertiesFormat.PrimaryAzurePort; primaryAzurePort != nil {
                v["primary_azure_port"] = *primaryAzurePort
            }
            if primaryPeerAddressPrefix := expressRouteCircuitPeeringPropertiesFormat.PrimaryPeerAddressPrefix; primaryPeerAddressPrefix != nil {
                v["primary_peer_address_prefix"] = *primaryPeerAddressPrefix
            }
            v["route_filter"] = flattenArmRouteFilterRouteFilter(expressRouteCircuitPeeringPropertiesFormat.RouteFilter)
            if secondaryAzurePort := expressRouteCircuitPeeringPropertiesFormat.SecondaryAzurePort; secondaryAzurePort != nil {
                v["secondary_azure_port"] = *secondaryAzurePort
            }
            if secondaryPeerAddressPrefix := expressRouteCircuitPeeringPropertiesFormat.SecondaryPeerAddressPrefix; secondaryPeerAddressPrefix != nil {
                v["secondary_peer_address_prefix"] = *secondaryPeerAddressPrefix
            }
            if sharedKey := expressRouteCircuitPeeringPropertiesFormat.SharedKey; sharedKey != nil {
                v["shared_key"] = *sharedKey
            }
            v["state"] = string(expressRouteCircuitPeeringPropertiesFormat.State)
            v["stats"] = flattenArmRouteFilterExpressRouteCircuitStats(expressRouteCircuitPeeringPropertiesFormat.Stats)
            if vlanId := expressRouteCircuitPeeringPropertiesFormat.VlanID; vlanId != nil {
                v["vlan_id"] = int(*vlanId)
            }
        }

        results = append(results, v)
    }

    return results
}

func flattenArmRouteFilterRouteFilterRule(input *[]network.RouteFilterRule) []interface{} {
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
        if location := item.Location; location != nil {
            v["location"] = azure.NormalizeLocation(*location)
        }
        if routeFilterRulePropertiesFormat := item.RouteFilterRulePropertiesFormat; routeFilterRulePropertiesFormat != nil {
            v["access"] = string(routeFilterRulePropertiesFormat.Access)
            v["communities"] = utils.FlattenStringSlice(routeFilterRulePropertiesFormat.Communities)
            if routeFilterRuleType := routeFilterRulePropertiesFormat.RouteFilterRuleType; routeFilterRuleType != nil {
                v["route_filter_rule_type"] = *routeFilterRuleType
            }
        }

        results = append(results, v)
    }

    return results
}

func flattenArmRouteFilterExpressRouteCircuitConnection(input *[]network.ExpressRouteCircuitConnection) []interface{} {
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

        results = append(results, v)
    }

    return results
}

func flattenArmRouteFilterIpv6ExpressRouteCircuitPeeringConfig(input *network.Ipv6ExpressRouteCircuitPeeringConfig) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if primaryPeerAddressPrefix := input.PrimaryPeerAddressPrefix; primaryPeerAddressPrefix != nil {
        result["primary_peer_address_prefix"] = *primaryPeerAddressPrefix
    }
    if secondaryPeerAddressPrefix := input.SecondaryPeerAddressPrefix; secondaryPeerAddressPrefix != nil {
        result["secondary_peer_address_prefix"] = *secondaryPeerAddressPrefix
    }
    result["state"] = string(input.State)

    return []interface{}{result}
}

func flattenArmRouteFilterExpressRouteCircuitPeeringConfig(input *network.ExpressRouteCircuitPeeringConfig) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    result["advertised_communities"] = utils.FlattenStringSlice(input.AdvertisedCommunities)
    result["advertised_public_prefixes"] = utils.FlattenStringSlice(input.AdvertisedPublicPrefixes)
    result["advertised_public_prefixes_state"] = string(input.AdvertisedPublicPrefixesState)
    if customerAsn := input.CustomerAsn; customerAsn != nil {
        result["customer_asn"] = int(*customerAsn)
    }
    if legacyMode := input.LegacyMode; legacyMode != nil {
        result["legacy_mode"] = int(*legacyMode)
    }
    if routingRegistryName := input.RoutingRegistryName; routingRegistryName != nil {
        result["routing_registry_name"] = *routingRegistryName
    }

    return []interface{}{result}
}

func flattenArmRouteFilterRouteFilter(input *network.RouteFilter) []interface{} {
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

    return []interface{}{result}
}

func flattenArmRouteFilterExpressRouteCircuitStats(input *network.ExpressRouteCircuitStats) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if primarybytesIn := input.PrimarybytesIn; primarybytesIn != nil {
        result["primarybytes_in"] = int(*primarybytesIn)
    }
    if primarybytesOut := input.PrimarybytesOut; primarybytesOut != nil {
        result["primarybytes_out"] = int(*primarybytesOut)
    }
    if secondarybytesIn := input.SecondarybytesIn; secondarybytesIn != nil {
        result["secondarybytes_in"] = int(*secondarybytesIn)
    }
    if secondarybytesOut := input.SecondarybytesOut; secondarybytesOut != nil {
        result["secondarybytes_out"] = int(*secondarybytesOut)
    }

    return []interface{}{result}
}
