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



func resourceArmExpressRouteCircuit() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmExpressRouteCircuitCreateUpdate,
        Read: resourceArmExpressRouteCircuitRead,
        Update: resourceArmExpressRouteCircuitCreateUpdate,
        Delete: resourceArmExpressRouteCircuitDelete,

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

            "circuit_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "allow_classic_operations": {
                Type: schema.TypeBool,
                Optional: true,
            },

            "authorizations": {
                Type: schema.TypeList,
                Optional: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "authorization_key": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "authorization_use_status": {
                            Type: schema.TypeString,
                            Optional: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(network.Available),
                                string(network.InUse),
                            }, false),
                            Default: string(network.Available),
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

            "circuit_provisioning_state": {
                Type: schema.TypeString,
                Optional: true,
            },

            "gateway_manager_etag": {
                Type: schema.TypeString,
                Optional: true,
            },

            "peerings": {
                Type: schema.TypeList,
                Optional: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "azure_asn": {
                            Type: schema.TypeInt,
                            Optional: true,
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

            "service_key": {
                Type: schema.TypeString,
                Optional: true,
            },

            "service_provider_notes": {
                Type: schema.TypeString,
                Optional: true,
            },

            "service_provider_properties": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "bandwidth_in_mbps": {
                            Type: schema.TypeInt,
                            Optional: true,
                        },
                        "peering_location": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "service_provider_name": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },

            "service_provider_provisioning_state": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(network.NotProvisioned),
                    string(network.Provisioning),
                    string(network.Provisioned),
                    string(network.Deprovisioning),
                }, false),
                Default: string(network.NotProvisioned),
            },

            "sku": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "family": {
                            Type: schema.TypeString,
                            Optional: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(network.UnlimitedData),
                                string(network.MeteredData),
                            }, false),
                            Default: string(network.UnlimitedData),
                        },
                        "name": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "tier": {
                            Type: schema.TypeString,
                            Optional: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(network.Standard),
                                string(network.Premium),
                            }, false),
                            Default: string(network.Standard),
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

func resourceArmExpressRouteCircuitCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).expressRouteCircuitsClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    circuitName := d.Get("circuit_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, circuitName)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Express Route Circuit (Circuit Name %q / Resource Group %q): %+v", circuitName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_express_route_circuit", *existing.ID)
        }
    }

    id := d.Get("id").(string)
    location := azure.NormalizeLocation(d.Get("location").(string))
    allowClassicOperations := d.Get("allow_classic_operations").(bool)
    authorizations := d.Get("authorizations").([]interface{})
    circuitProvisioningState := d.Get("circuit_provisioning_state").(string)
    gatewayManagerEtag := d.Get("gateway_manager_etag").(string)
    peerings := d.Get("peerings").([]interface{})
    serviceKey := d.Get("service_key").(string)
    serviceProviderNotes := d.Get("service_provider_notes").(string)
    serviceProviderProperties := d.Get("service_provider_properties").([]interface{})
    serviceProviderProvisioningState := d.Get("service_provider_provisioning_state").(string)
    sku := d.Get("sku").([]interface{})
    t := d.Get("tags").(map[string]interface{})

    parameters := network.ExpressRouteCircuit{
        ID: utils.String(id),
        Location: utils.String(location),
        ExpressRouteCircuitPropertiesFormat: &network.ExpressRouteCircuitPropertiesFormat{
            AllowClassicOperations: utils.Bool(allowClassicOperations),
            Authorizations: expandArmExpressRouteCircuitExpressRouteCircuitAuthorization(authorizations),
            CircuitProvisioningState: utils.String(circuitProvisioningState),
            GatewayManagerEtag: utils.String(gatewayManagerEtag),
            Peerings: expandArmExpressRouteCircuitExpressRouteCircuitPeering(peerings),
            ServiceKey: utils.String(serviceKey),
            ServiceProviderNotes: utils.String(serviceProviderNotes),
            ServiceProviderProperties: expandArmExpressRouteCircuitExpressRouteCircuitServiceProviderProperties(serviceProviderProperties),
            ServiceProviderProvisioningState: network.ServiceProviderProvisioningState(serviceProviderProvisioningState),
        },
        Sku: expandArmExpressRouteCircuitExpressRouteCircuitSku(sku),
        Tags: tags.Expand(t),
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, circuitName, parameters)
    if err != nil {
        return fmt.Errorf("Error creating Express Route Circuit (Circuit Name %q / Resource Group %q): %+v", circuitName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Express Route Circuit (Circuit Name %q / Resource Group %q): %+v", circuitName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, circuitName)
    if err != nil {
        return fmt.Errorf("Error retrieving Express Route Circuit (Circuit Name %q / Resource Group %q): %+v", circuitName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Express Route Circuit (Circuit Name %q / Resource Group %q) ID", circuitName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmExpressRouteCircuitRead(d, meta)
}

func resourceArmExpressRouteCircuitRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).expressRouteCircuitsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    circuitName := id.Path["expressRouteCircuits"]

    resp, err := client.Get(ctx, resourceGroup, circuitName)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Express Route Circuit %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Express Route Circuit (Circuit Name %q / Resource Group %q): %+v", circuitName, resourceGroup, err)
    }


    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if expressRouteCircuitPropertiesFormat := resp.ExpressRouteCircuitPropertiesFormat; expressRouteCircuitPropertiesFormat != nil {
        d.Set("allow_classic_operations", expressRouteCircuitPropertiesFormat.AllowClassicOperations)
        if err := d.Set("authorizations", flattenArmExpressRouteCircuitExpressRouteCircuitAuthorization(expressRouteCircuitPropertiesFormat.Authorizations)); err != nil {
            return fmt.Errorf("Error setting `authorizations`: %+v", err)
        }
        d.Set("circuit_provisioning_state", expressRouteCircuitPropertiesFormat.CircuitProvisioningState)
        d.Set("gateway_manager_etag", expressRouteCircuitPropertiesFormat.GatewayManagerEtag)
        if err := d.Set("peerings", flattenArmExpressRouteCircuitExpressRouteCircuitPeering(expressRouteCircuitPropertiesFormat.Peerings)); err != nil {
            return fmt.Errorf("Error setting `peerings`: %+v", err)
        }
        d.Set("provisioning_state", expressRouteCircuitPropertiesFormat.ProvisioningState)
        d.Set("service_key", expressRouteCircuitPropertiesFormat.ServiceKey)
        d.Set("service_provider_notes", expressRouteCircuitPropertiesFormat.ServiceProviderNotes)
        if err := d.Set("service_provider_properties", flattenArmExpressRouteCircuitExpressRouteCircuitServiceProviderProperties(expressRouteCircuitPropertiesFormat.ServiceProviderProperties)); err != nil {
            return fmt.Errorf("Error setting `service_provider_properties`: %+v", err)
        }
        d.Set("service_provider_provisioning_state", string(expressRouteCircuitPropertiesFormat.ServiceProviderProvisioningState))
    }
    d.Set("circuit_name", circuitName)
    d.Set("etag", resp.Etag)
    if err := d.Set("sku", flattenArmExpressRouteCircuitExpressRouteCircuitSku(resp.Sku)); err != nil {
        return fmt.Errorf("Error setting `sku`: %+v", err)
    }
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}


func resourceArmExpressRouteCircuitDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).expressRouteCircuitsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    circuitName := id.Path["expressRouteCircuits"]

    future, err := client.Delete(ctx, resourceGroup, circuitName)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Express Route Circuit (Circuit Name %q / Resource Group %q): %+v", circuitName, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Express Route Circuit (Circuit Name %q / Resource Group %q): %+v", circuitName, resourceGroup, err)
        }
    }

    return nil
}

func expandArmExpressRouteCircuitExpressRouteCircuitAuthorization(input []interface{}) *[]network.ExpressRouteCircuitAuthorization {
    results := make([]network.ExpressRouteCircuitAuthorization, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        id := v["id"].(string)
        authorizationKey := v["authorization_key"].(string)
        authorizationUseStatus := v["authorization_use_status"].(string)
        name := v["name"].(string)

        result := network.ExpressRouteCircuitAuthorization{
            ID: utils.String(id),
            Name: utils.String(name),
            AuthorizationPropertiesFormat: &network.AuthorizationPropertiesFormat{
                AuthorizationKey: utils.String(authorizationKey),
                AuthorizationUseStatus: network.AuthorizationUseStatus(authorizationUseStatus),
            },
        }

        results = append(results, result)
    }
    return &results
}

func expandArmExpressRouteCircuitExpressRouteCircuitPeering(input []interface{}) *[]network.ExpressRouteCircuitPeering {
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
        name := v["name"].(string)

        result := network.ExpressRouteCircuitPeering{
            ID: utils.String(id),
            Name: utils.String(name),
            ExpressRouteCircuitPeeringPropertiesFormat: &network.ExpressRouteCircuitPeeringPropertiesFormat{
                AzureAsn: utils.Int32(int32(azureAsn)),
                GatewayManagerEtag: utils.String(gatewayManagerEtag),
                Ipv6peeringConfig: expandArmExpressRouteCircuitIpv6ExpressRouteCircuitPeeringConfig(ipv6peeringConfig),
                LastModifiedBy: utils.String(lastModifiedBy),
                MicrosoftPeeringConfig: expandArmExpressRouteCircuitExpressRouteCircuitPeeringConfig(microsoftPeeringConfig),
                PeerAsn: utils.Int64(int64(peerAsn)),
                PeeringType: network.ExpressRouteCircuitPeeringType(peeringType),
                PrimaryAzurePort: utils.String(primaryAzurePort),
                PrimaryPeerAddressPrefix: utils.String(primaryPeerAddressPrefix),
                RouteFilter: expandArmExpressRouteCircuitRouteFilter(routeFilter),
                SecondaryAzurePort: utils.String(secondaryAzurePort),
                SecondaryPeerAddressPrefix: utils.String(secondaryPeerAddressPrefix),
                SharedKey: utils.String(sharedKey),
                State: network.ExpressRouteCircuitPeeringState(state),
                Stats: expandArmExpressRouteCircuitExpressRouteCircuitStats(stats),
                VlanID: utils.Int32(int32(vlanId)),
            },
        }

        results = append(results, result)
    }
    return &results
}

func expandArmExpressRouteCircuitExpressRouteCircuitServiceProviderProperties(input []interface{}) *network.ExpressRouteCircuitServiceProviderProperties {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    serviceProviderName := v["service_provider_name"].(string)
    peeringLocation := v["peering_location"].(string)
    bandwidthInMbps := v["bandwidth_in_mbps"].(int)

    result := network.ExpressRouteCircuitServiceProviderProperties{
        BandwidthInMbps: utils.Int32(int32(bandwidthInMbps)),
        PeeringLocation: utils.String(peeringLocation),
        ServiceProviderName: utils.String(serviceProviderName),
    }
    return &result
}

func expandArmExpressRouteCircuitExpressRouteCircuitSku(input []interface{}) *network.ExpressRouteCircuitSku {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    name := v["name"].(string)
    tier := v["tier"].(string)
    family := v["family"].(string)

    result := network.ExpressRouteCircuitSku{
        Family: network.ExpressRouteCircuitSkuFamily(family),
        Name: utils.String(name),
        Tier: network.ExpressRouteCircuitSkuTier(tier),
    }
    return &result
}

func expandArmExpressRouteCircuitIpv6ExpressRouteCircuitPeeringConfig(input []interface{}) *network.Ipv6ExpressRouteCircuitPeeringConfig {
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

func expandArmExpressRouteCircuitExpressRouteCircuitPeeringConfig(input []interface{}) *network.ExpressRouteCircuitPeeringConfig {
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

func expandArmExpressRouteCircuitRouteFilter(input []interface{}) *network.RouteFilter {
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

func expandArmExpressRouteCircuitExpressRouteCircuitStats(input []interface{}) *network.ExpressRouteCircuitStats {
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


func flattenArmExpressRouteCircuitExpressRouteCircuitAuthorization(input *[]network.ExpressRouteCircuitAuthorization) []interface{} {
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
        if authorizationPropertiesFormat := item.AuthorizationPropertiesFormat; authorizationPropertiesFormat != nil {
            if authorizationKey := authorizationPropertiesFormat.AuthorizationKey; authorizationKey != nil {
                v["authorization_key"] = *authorizationKey
            }
            v["authorization_use_status"] = string(authorizationPropertiesFormat.AuthorizationUseStatus)
        }

        results = append(results, v)
    }

    return results
}

func flattenArmExpressRouteCircuitExpressRouteCircuitPeering(input *[]network.ExpressRouteCircuitPeering) []interface{} {
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
            if gatewayManagerEtag := expressRouteCircuitPeeringPropertiesFormat.GatewayManagerEtag; gatewayManagerEtag != nil {
                v["gateway_manager_etag"] = *gatewayManagerEtag
            }
            v["ipv6peering_config"] = flattenArmExpressRouteCircuitIpv6ExpressRouteCircuitPeeringConfig(expressRouteCircuitPeeringPropertiesFormat.Ipv6peeringConfig)
            if lastModifiedBy := expressRouteCircuitPeeringPropertiesFormat.LastModifiedBy; lastModifiedBy != nil {
                v["last_modified_by"] = *lastModifiedBy
            }
            v["microsoft_peering_config"] = flattenArmExpressRouteCircuitExpressRouteCircuitPeeringConfig(expressRouteCircuitPeeringPropertiesFormat.MicrosoftPeeringConfig)
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
            v["route_filter"] = flattenArmExpressRouteCircuitRouteFilter(expressRouteCircuitPeeringPropertiesFormat.RouteFilter)
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
            v["stats"] = flattenArmExpressRouteCircuitExpressRouteCircuitStats(expressRouteCircuitPeeringPropertiesFormat.Stats)
            if vlanId := expressRouteCircuitPeeringPropertiesFormat.VlanID; vlanId != nil {
                v["vlan_id"] = int(*vlanId)
            }
        }

        results = append(results, v)
    }

    return results
}

func flattenArmExpressRouteCircuitExpressRouteCircuitServiceProviderProperties(input *network.ExpressRouteCircuitServiceProviderProperties) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if bandwidthInMbps := input.BandwidthInMbps; bandwidthInMbps != nil {
        result["bandwidth_in_mbps"] = int(*bandwidthInMbps)
    }
    if peeringLocation := input.PeeringLocation; peeringLocation != nil {
        result["peering_location"] = *peeringLocation
    }
    if serviceProviderName := input.ServiceProviderName; serviceProviderName != nil {
        result["service_provider_name"] = *serviceProviderName
    }

    return []interface{}{result}
}

func flattenArmExpressRouteCircuitExpressRouteCircuitSku(input *network.ExpressRouteCircuitSku) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if name := input.Name; name != nil {
        result["name"] = *name
    }
    result["family"] = string(input.Family)
    result["tier"] = string(input.Tier)

    return []interface{}{result}
}

func flattenArmExpressRouteCircuitIpv6ExpressRouteCircuitPeeringConfig(input *network.Ipv6ExpressRouteCircuitPeeringConfig) []interface{} {
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

func flattenArmExpressRouteCircuitExpressRouteCircuitPeeringConfig(input *network.ExpressRouteCircuitPeeringConfig) []interface{} {
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

func flattenArmExpressRouteCircuitRouteFilter(input *network.RouteFilter) []interface{} {
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

func flattenArmExpressRouteCircuitExpressRouteCircuitStats(input *network.ExpressRouteCircuitStats) []interface{} {
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
