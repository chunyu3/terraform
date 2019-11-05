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



func resourceArmExpressRouteCircuitPeering() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmExpressRouteCircuitPeeringCreateUpdate,
        Read: resourceArmExpressRouteCircuitPeeringRead,
        Update: resourceArmExpressRouteCircuitPeeringCreateUpdate,
        Delete: resourceArmExpressRouteCircuitPeeringDelete,

        Importer: &schema.ResourceImporter{
            State: schema.ImportStatePassthrough,
        },


        Schema: map[string]*schema.Schema{
            "name": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
            },

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "circuit_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "peering_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "azure_asn": {
                Type: schema.TypeInt,
                Optional: true,
            },

            "gateway_manager_etag": {
                Type: schema.TypeString,
                Optional: true,
            },

            "ipv6peering_config": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
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
                        "peerings": {
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
                        "rules": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
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
                                },
                            },
                        },
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

            "etag": {
                Type: schema.TypeString,
                Computed: true,
            },

            "provisioning_state": {
                Type: schema.TypeString,
                Computed: true,
            },
        },
    }
}

func resourceArmExpressRouteCircuitPeeringCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).expressRouteCircuitPeeringsClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    circuitName := d.Get("circuit_name").(string)
    peeringName := d.Get("peering_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, circuitName, peeringName)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Express Route Circuit Peering (Peering Name %q / Circuit Name %q / Resource Group %q): %+v", peeringName, circuitName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_express_route_circuit_peering", *existing.ID)
        }
    }

    id := d.Get("id").(string)
    name := d.Get("name").(string)
    azureAsn := d.Get("azure_asn").(int)
    gatewayManagerEtag := d.Get("gateway_manager_etag").(string)
    ipv6peeringConfig := d.Get("ipv6peering_config").([]interface{})
    lastModifiedBy := d.Get("last_modified_by").(string)
    microsoftPeeringConfig := d.Get("microsoft_peering_config").([]interface{})
    peerAsn := d.Get("peer_asn").(int)
    peeringType := d.Get("peering_type").(string)
    primaryAzurePort := d.Get("primary_azure_port").(string)
    primaryPeerAddressPrefix := d.Get("primary_peer_address_prefix").(string)
    routeFilter := d.Get("route_filter").([]interface{})
    secondaryAzurePort := d.Get("secondary_azure_port").(string)
    secondaryPeerAddressPrefix := d.Get("secondary_peer_address_prefix").(string)
    sharedKey := d.Get("shared_key").(string)
    state := d.Get("state").(string)
    stats := d.Get("stats").([]interface{})
    vlanId := d.Get("vlan_id").(int)

    peeringParameters := network.ExpressRouteCircuitPeering{
        ID: utils.String(id),
        Name: utils.String(name),
        ExpressRouteCircuitPeeringPropertiesFormat: &network.ExpressRouteCircuitPeeringPropertiesFormat{
            AzureAsn: utils.Int32(int32(azureAsn)),
            GatewayManagerEtag: utils.String(gatewayManagerEtag),
            Ipv6peeringConfig: expandArmExpressRouteCircuitPeeringIpv6ExpressRouteCircuitPeeringConfig(ipv6peeringConfig),
            LastModifiedBy: utils.String(lastModifiedBy),
            MicrosoftPeeringConfig: expandArmExpressRouteCircuitPeeringExpressRouteCircuitPeeringConfig(microsoftPeeringConfig),
            PeerAsn: utils.Int32(int32(peerAsn)),
            PeeringType: network.ExpressRouteCircuitPeeringType(peeringType),
            PrimaryAzurePort: utils.String(primaryAzurePort),
            PrimaryPeerAddressPrefix: utils.String(primaryPeerAddressPrefix),
            RouteFilter: expandArmExpressRouteCircuitPeeringRouteFilter(routeFilter),
            SecondaryAzurePort: utils.String(secondaryAzurePort),
            SecondaryPeerAddressPrefix: utils.String(secondaryPeerAddressPrefix),
            SharedKey: utils.String(sharedKey),
            State: network.ExpressRouteCircuitPeeringState(state),
            Stats: expandArmExpressRouteCircuitPeeringExpressRouteCircuitStats(stats),
            VlanID: utils.Int32(int32(vlanId)),
        },
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, circuitName, peeringName, peeringParameters)
    if err != nil {
        return fmt.Errorf("Error creating Express Route Circuit Peering (Peering Name %q / Circuit Name %q / Resource Group %q): %+v", peeringName, circuitName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Express Route Circuit Peering (Peering Name %q / Circuit Name %q / Resource Group %q): %+v", peeringName, circuitName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, circuitName, peeringName)
    if err != nil {
        return fmt.Errorf("Error retrieving Express Route Circuit Peering (Peering Name %q / Circuit Name %q / Resource Group %q): %+v", peeringName, circuitName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Express Route Circuit Peering (Peering Name %q / Circuit Name %q / Resource Group %q) ID", peeringName, circuitName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmExpressRouteCircuitPeeringRead(d, meta)
}

func resourceArmExpressRouteCircuitPeeringRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).expressRouteCircuitPeeringsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    circuitName := id.Path["expressRouteCircuits"]
    peeringName := id.Path["peerings"]

    resp, err := client.Get(ctx, resourceGroup, circuitName, peeringName)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Express Route Circuit Peering %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Express Route Circuit Peering (Peering Name %q / Circuit Name %q / Resource Group %q): %+v", peeringName, circuitName, resourceGroup, err)
    }


    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if expressRouteCircuitPeeringPropertiesFormat := resp.ExpressRouteCircuitPeeringPropertiesFormat; expressRouteCircuitPeeringPropertiesFormat != nil {
        d.Set("azure_asn", int(*expressRouteCircuitPeeringPropertiesFormat.AzureAsn))
        d.Set("gateway_manager_etag", expressRouteCircuitPeeringPropertiesFormat.GatewayManagerEtag)
        if err := d.Set("ipv6peering_config", flattenArmExpressRouteCircuitPeeringIpv6ExpressRouteCircuitPeeringConfig(expressRouteCircuitPeeringPropertiesFormat.Ipv6peeringConfig)); err != nil {
            return fmt.Errorf("Error setting `ipv6peering_config`: %+v", err)
        }
        d.Set("last_modified_by", expressRouteCircuitPeeringPropertiesFormat.LastModifiedBy)
        if err := d.Set("microsoft_peering_config", flattenArmExpressRouteCircuitPeeringExpressRouteCircuitPeeringConfig(expressRouteCircuitPeeringPropertiesFormat.MicrosoftPeeringConfig)); err != nil {
            return fmt.Errorf("Error setting `microsoft_peering_config`: %+v", err)
        }
        d.Set("peer_asn", int(*expressRouteCircuitPeeringPropertiesFormat.PeerAsn))
        d.Set("peering_type", string(expressRouteCircuitPeeringPropertiesFormat.PeeringType))
        d.Set("primary_azure_port", expressRouteCircuitPeeringPropertiesFormat.PrimaryAzurePort)
        d.Set("primary_peer_address_prefix", expressRouteCircuitPeeringPropertiesFormat.PrimaryPeerAddressPrefix)
        d.Set("provisioning_state", expressRouteCircuitPeeringPropertiesFormat.ProvisioningState)
        if err := d.Set("route_filter", flattenArmExpressRouteCircuitPeeringRouteFilter(expressRouteCircuitPeeringPropertiesFormat.RouteFilter)); err != nil {
            return fmt.Errorf("Error setting `route_filter`: %+v", err)
        }
        d.Set("secondary_azure_port", expressRouteCircuitPeeringPropertiesFormat.SecondaryAzurePort)
        d.Set("secondary_peer_address_prefix", expressRouteCircuitPeeringPropertiesFormat.SecondaryPeerAddressPrefix)
        d.Set("shared_key", expressRouteCircuitPeeringPropertiesFormat.SharedKey)
        d.Set("state", string(expressRouteCircuitPeeringPropertiesFormat.State))
        if err := d.Set("stats", flattenArmExpressRouteCircuitPeeringExpressRouteCircuitStats(expressRouteCircuitPeeringPropertiesFormat.Stats)); err != nil {
            return fmt.Errorf("Error setting `stats`: %+v", err)
        }
        d.Set("vlan_id", int(*expressRouteCircuitPeeringPropertiesFormat.VlanID))
    }
    d.Set("circuit_name", circuitName)
    d.Set("etag", resp.Etag)
    d.Set("peering_name", peeringName)

    return nil
}


func resourceArmExpressRouteCircuitPeeringDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).expressRouteCircuitPeeringsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    circuitName := id.Path["expressRouteCircuits"]
    peeringName := id.Path["peerings"]

    future, err := client.Delete(ctx, resourceGroup, circuitName, peeringName)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Express Route Circuit Peering (Peering Name %q / Circuit Name %q / Resource Group %q): %+v", peeringName, circuitName, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Express Route Circuit Peering (Peering Name %q / Circuit Name %q / Resource Group %q): %+v", peeringName, circuitName, resourceGroup, err)
        }
    }

    return nil
}

func expandArmExpressRouteCircuitPeeringIpv6ExpressRouteCircuitPeeringConfig(input []interface{}) *network.Ipv6ExpressRouteCircuitPeeringConfig {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    primaryPeerAddressPrefix := v["primary_peer_address_prefix"].(string)
    secondaryPeerAddressPrefix := v["secondary_peer_address_prefix"].(string)
    microsoftPeeringConfig := v["microsoft_peering_config"].([]interface{})
    routeFilter := v["route_filter"].([]interface{})
    state := v["state"].(string)

    result := network.Ipv6ExpressRouteCircuitPeeringConfig{
        MicrosoftPeeringConfig: expandArmExpressRouteCircuitPeeringExpressRouteCircuitPeeringConfig(microsoftPeeringConfig),
        PrimaryPeerAddressPrefix: utils.String(primaryPeerAddressPrefix),
        RouteFilter: expandArmExpressRouteCircuitPeeringRouteFilter(routeFilter),
        SecondaryPeerAddressPrefix: utils.String(secondaryPeerAddressPrefix),
        State: network.ExpressRouteCircuitPeeringState(state),
    }
    return &result
}

func expandArmExpressRouteCircuitPeeringExpressRouteCircuitPeeringConfig(input []interface{}) *network.ExpressRouteCircuitPeeringConfig {
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

func expandArmExpressRouteCircuitPeeringRouteFilter(input []interface{}) *network.RouteFilter {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    id := v["id"].(string)
    location := azure.NormalizeLocation(v["location"].(string))
    t := v["tags"].(map[string]interface{})
    rules := v["rules"].([]interface{})
    peerings := v["peerings"].([]interface{})

    result := network.RouteFilter{
        ID: utils.String(id),
        Location: utils.String(location),
        RouteFilterPropertiesFormat: &network.RouteFilterPropertiesFormat{
            Peerings: expandArmExpressRouteCircuitPeeringExpressRouteCircuitPeering(peerings),
            Rules: expandArmExpressRouteCircuitPeeringRouteFilterRule(rules),
        },
        Tags: tags.Expand(t),
    }
    return &result
}

func expandArmExpressRouteCircuitPeeringExpressRouteCircuitStats(input []interface{}) *network.ExpressRouteCircuitStats {
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

func expandArmExpressRouteCircuitPeeringRouteFilter(input []interface{}) *network.RouteFilter {
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

func expandArmExpressRouteCircuitPeeringExpressRouteCircuitPeering(input []interface{}) *[]network.ExpressRouteCircuitPeering {
    results := make([]network.ExpressRouteCircuitPeering, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        id := v["id"].(string)
        name := v["name"].(string)

        result := network.ExpressRouteCircuitPeering{
            ID: utils.String(id),
            Name: utils.String(name),
        }

        results = append(results, result)
    }
    return &results
}

func expandArmExpressRouteCircuitPeeringRouteFilterRule(input []interface{}) *[]network.RouteFilterRule {
    results := make([]network.RouteFilterRule, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        id := v["id"].(string)
        name := v["name"].(string)
        location := azure.NormalizeLocation(v["location"].(string))
        t := v["tags"].(map[string]interface{})

        result := network.RouteFilterRule{
            ID: utils.String(id),
            Location: utils.String(location),
            Name: utils.String(name),
            Tags: tags.Expand(t),
        }

        results = append(results, result)
    }
    return &results
}


func flattenArmExpressRouteCircuitPeeringIpv6ExpressRouteCircuitPeeringConfig(input *network.Ipv6ExpressRouteCircuitPeeringConfig) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    result["microsoft_peering_config"] = flattenArmExpressRouteCircuitPeeringExpressRouteCircuitPeeringConfig(input.MicrosoftPeeringConfig)
    if primaryPeerAddressPrefix := input.PrimaryPeerAddressPrefix; primaryPeerAddressPrefix != nil {
        result["primary_peer_address_prefix"] = *primaryPeerAddressPrefix
    }
    result["route_filter"] = flattenArmExpressRouteCircuitPeeringRouteFilter(input.RouteFilter)
    if secondaryPeerAddressPrefix := input.SecondaryPeerAddressPrefix; secondaryPeerAddressPrefix != nil {
        result["secondary_peer_address_prefix"] = *secondaryPeerAddressPrefix
    }
    result["state"] = string(input.State)

    return []interface{}{result}
}

func flattenArmExpressRouteCircuitPeeringExpressRouteCircuitPeeringConfig(input *network.ExpressRouteCircuitPeeringConfig) []interface{} {
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

func flattenArmExpressRouteCircuitPeeringRouteFilter(input *network.RouteFilter) []interface{} {
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
    if routeFilterPropertiesFormat := input.RouteFilterPropertiesFormat; routeFilterPropertiesFormat != nil {
        result["peerings"] = flattenArmExpressRouteCircuitPeeringExpressRouteCircuitPeering(routeFilterPropertiesFormat.Peerings)
        result["rules"] = flattenArmExpressRouteCircuitPeeringRouteFilterRule(routeFilterPropertiesFormat.Rules)
    }

    return []interface{}{result}
}

func flattenArmExpressRouteCircuitPeeringExpressRouteCircuitStats(input *network.ExpressRouteCircuitStats) []interface{} {
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

func flattenArmExpressRouteCircuitPeeringRouteFilter(input *network.RouteFilter) []interface{} {
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

func flattenArmExpressRouteCircuitPeeringExpressRouteCircuitPeering(input *[]network.ExpressRouteCircuitPeering) []interface{} {
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

func flattenArmExpressRouteCircuitPeeringRouteFilterRule(input *[]network.RouteFilterRule) []interface{} {
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

        results = append(results, v)
    }

    return results
}
