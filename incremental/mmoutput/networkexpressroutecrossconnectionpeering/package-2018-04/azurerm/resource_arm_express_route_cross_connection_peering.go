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



func resourceArmExpressRouteCrossConnectionPeering() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmExpressRouteCrossConnectionPeeringCreateUpdate,
        Read: resourceArmExpressRouteCrossConnectionPeeringRead,
        Update: resourceArmExpressRouteCrossConnectionPeeringCreateUpdate,
        Delete: resourceArmExpressRouteCrossConnectionPeeringDelete,

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
                Optional: true,
                ForceNew: true,
            },

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "cross_connection_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
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

            "primary_peer_address_prefix": {
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

            "vlan_id": {
                Type: schema.TypeInt,
                Optional: true,
            },

            "azure_asn": {
                Type: schema.TypeInt,
                Computed: true,
            },

            "etag": {
                Type: schema.TypeString,
                Computed: true,
            },

            "primary_azure_port": {
                Type: schema.TypeString,
                Computed: true,
            },

            "provisioning_state": {
                Type: schema.TypeString,
                Computed: true,
            },

            "secondary_azure_port": {
                Type: schema.TypeString,
                Computed: true,
            },
        },
    }
}

func resourceArmExpressRouteCrossConnectionPeeringCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).expressRouteCrossConnectionPeeringsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    crossConnectionName := d.Get("cross_connection_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, crossConnectionName, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Express Route Cross Connection Peering %q (Cross Connection Name %q / Resource Group %q): %+v", name, crossConnectionName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_express_route_cross_connection_peering", *existing.ID)
        }
    }

    id := d.Get("id").(string)
    name := d.Get("name").(string)
    gatewayManagerEtag := d.Get("gateway_manager_etag").(string)
    ipv6peeringConfig := d.Get("ipv6peering_config").([]interface{})
    lastModifiedBy := d.Get("last_modified_by").(string)
    microsoftPeeringConfig := d.Get("microsoft_peering_config").([]interface{})
    peerAsn := d.Get("peer_asn").(int)
    peeringType := d.Get("peering_type").(string)
    primaryPeerAddressPrefix := d.Get("primary_peer_address_prefix").(string)
    secondaryPeerAddressPrefix := d.Get("secondary_peer_address_prefix").(string)
    sharedKey := d.Get("shared_key").(string)
    state := d.Get("state").(string)
    vlanId := d.Get("vlan_id").(int)

    peeringParameters := network.ExpressRouteCrossConnectionPeering{
        ID: utils.String(id),
        Name: utils.String(name),
        ExpressRouteCrossConnectionPeeringProperties: &network.ExpressRouteCrossConnectionPeeringProperties{
            GatewayManagerEtag: utils.String(gatewayManagerEtag),
            Ipv6peeringConfig: expandArmExpressRouteCrossConnectionPeeringIpv6ExpressRouteCircuitPeeringConfig(ipv6peeringConfig),
            LastModifiedBy: utils.String(lastModifiedBy),
            MicrosoftPeeringConfig: expandArmExpressRouteCrossConnectionPeeringExpressRouteCircuitPeeringConfig(microsoftPeeringConfig),
            PeerAsn: utils.Int64(int64(peerAsn)),
            PeeringType: network.ExpressRoutePeeringType(peeringType),
            PrimaryPeerAddressPrefix: utils.String(primaryPeerAddressPrefix),
            SecondaryPeerAddressPrefix: utils.String(secondaryPeerAddressPrefix),
            SharedKey: utils.String(sharedKey),
            State: network.ExpressRoutePeeringState(state),
            VlanID: utils.Int32(int32(vlanId)),
        },
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, crossConnectionName, name, peeringParameters)
    if err != nil {
        return fmt.Errorf("Error creating Express Route Cross Connection Peering %q (Cross Connection Name %q / Resource Group %q): %+v", name, crossConnectionName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Express Route Cross Connection Peering %q (Cross Connection Name %q / Resource Group %q): %+v", name, crossConnectionName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, crossConnectionName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Express Route Cross Connection Peering %q (Cross Connection Name %q / Resource Group %q): %+v", name, crossConnectionName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Express Route Cross Connection Peering %q (Cross Connection Name %q / Resource Group %q) ID", name, crossConnectionName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmExpressRouteCrossConnectionPeeringRead(d, meta)
}

func resourceArmExpressRouteCrossConnectionPeeringRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).expressRouteCrossConnectionPeeringsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    crossConnectionName := id.Path["expressRouteCrossConnections"]
    name := id.Path["peerings"]

    resp, err := client.Get(ctx, resourceGroup, crossConnectionName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Express Route Cross Connection Peering %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Express Route Cross Connection Peering %q (Cross Connection Name %q / Resource Group %q): %+v", name, crossConnectionName, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if expressRouteCrossConnectionPeeringProperties := resp.ExpressRouteCrossConnectionPeeringProperties; expressRouteCrossConnectionPeeringProperties != nil {
        d.Set("azure_asn", int(*expressRouteCrossConnectionPeeringProperties.AzureAsn))
        d.Set("gateway_manager_etag", expressRouteCrossConnectionPeeringProperties.GatewayManagerEtag)
        if err := d.Set("ipv6peering_config", flattenArmExpressRouteCrossConnectionPeeringIpv6ExpressRouteCircuitPeeringConfig(expressRouteCrossConnectionPeeringProperties.Ipv6peeringConfig)); err != nil {
            return fmt.Errorf("Error setting `ipv6peering_config`: %+v", err)
        }
        d.Set("last_modified_by", expressRouteCrossConnectionPeeringProperties.LastModifiedBy)
        if err := d.Set("microsoft_peering_config", flattenArmExpressRouteCrossConnectionPeeringExpressRouteCircuitPeeringConfig(expressRouteCrossConnectionPeeringProperties.MicrosoftPeeringConfig)); err != nil {
            return fmt.Errorf("Error setting `microsoft_peering_config`: %+v", err)
        }
        d.Set("peer_asn", int(*expressRouteCrossConnectionPeeringProperties.PeerAsn))
        d.Set("peering_type", string(expressRouteCrossConnectionPeeringProperties.PeeringType))
        d.Set("primary_azure_port", expressRouteCrossConnectionPeeringProperties.PrimaryAzurePort)
        d.Set("primary_peer_address_prefix", expressRouteCrossConnectionPeeringProperties.PrimaryPeerAddressPrefix)
        d.Set("provisioning_state", expressRouteCrossConnectionPeeringProperties.ProvisioningState)
        d.Set("secondary_azure_port", expressRouteCrossConnectionPeeringProperties.SecondaryAzurePort)
        d.Set("secondary_peer_address_prefix", expressRouteCrossConnectionPeeringProperties.SecondaryPeerAddressPrefix)
        d.Set("shared_key", expressRouteCrossConnectionPeeringProperties.SharedKey)
        d.Set("state", string(expressRouteCrossConnectionPeeringProperties.State))
        d.Set("vlan_id", int(*expressRouteCrossConnectionPeeringProperties.VlanID))
    }
    d.Set("cross_connection_name", crossConnectionName)
    d.Set("etag", resp.Etag)

    return nil
}


func resourceArmExpressRouteCrossConnectionPeeringDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).expressRouteCrossConnectionPeeringsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    crossConnectionName := id.Path["expressRouteCrossConnections"]
    name := id.Path["peerings"]

    future, err := client.Delete(ctx, resourceGroup, crossConnectionName, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Express Route Cross Connection Peering %q (Cross Connection Name %q / Resource Group %q): %+v", name, crossConnectionName, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Express Route Cross Connection Peering %q (Cross Connection Name %q / Resource Group %q): %+v", name, crossConnectionName, resourceGroup, err)
        }
    }

    return nil
}

func expandArmExpressRouteCrossConnectionPeeringIpv6ExpressRouteCircuitPeeringConfig(input []interface{}) *network.Ipv6ExpressRouteCircuitPeeringConfig {
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
        MicrosoftPeeringConfig: expandArmExpressRouteCrossConnectionPeeringExpressRouteCircuitPeeringConfig(microsoftPeeringConfig),
        PrimaryPeerAddressPrefix: utils.String(primaryPeerAddressPrefix),
        RouteFilter: expandArmExpressRouteCrossConnectionPeeringRouteFilter(routeFilter),
        SecondaryPeerAddressPrefix: utils.String(secondaryPeerAddressPrefix),
        State: network.ExpressRouteCircuitPeeringState(state),
    }
    return &result
}

func expandArmExpressRouteCrossConnectionPeeringExpressRouteCircuitPeeringConfig(input []interface{}) *network.ExpressRouteCircuitPeeringConfig {
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

func expandArmExpressRouteCrossConnectionPeeringRouteFilter(input []interface{}) *network.RouteFilter {
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


func flattenArmExpressRouteCrossConnectionPeeringIpv6ExpressRouteCircuitPeeringConfig(input *network.Ipv6ExpressRouteCircuitPeeringConfig) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    result["microsoft_peering_config"] = flattenArmExpressRouteCrossConnectionPeeringExpressRouteCircuitPeeringConfig(input.MicrosoftPeeringConfig)
    if primaryPeerAddressPrefix := input.PrimaryPeerAddressPrefix; primaryPeerAddressPrefix != nil {
        result["primary_peer_address_prefix"] = *primaryPeerAddressPrefix
    }
    result["route_filter"] = flattenArmExpressRouteCrossConnectionPeeringRouteFilter(input.RouteFilter)
    if secondaryPeerAddressPrefix := input.SecondaryPeerAddressPrefix; secondaryPeerAddressPrefix != nil {
        result["secondary_peer_address_prefix"] = *secondaryPeerAddressPrefix
    }
    result["state"] = string(input.State)

    return []interface{}{result}
}

func flattenArmExpressRouteCrossConnectionPeeringExpressRouteCircuitPeeringConfig(input *network.ExpressRouteCircuitPeeringConfig) []interface{} {
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

func flattenArmExpressRouteCrossConnectionPeeringRouteFilter(input *network.RouteFilter) []interface{} {
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
