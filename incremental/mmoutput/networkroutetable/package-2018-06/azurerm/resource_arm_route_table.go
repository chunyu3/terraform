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



func resourceArmRouteTable() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmRouteTableCreateUpdate,
        Read: resourceArmRouteTableRead,
        Update: resourceArmRouteTableCreateUpdate,
        Delete: resourceArmRouteTableDelete,

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

            "disable_bgp_route_propagation": {
                Type: schema.TypeBool,
                Optional: true,
            },

            "etag": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
            },

            "routes": {
                Type: schema.TypeList,
                Optional: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "next_hop_type": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(network.VirtualNetworkGateway),
                                string(network.VnetLocal),
                                string(network.Internet),
                                string(network.VirtualAppliance),
                                string(network.None),
                            }, false),
                        },
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
                        "name": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "next_hop_ip_address": {
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

func resourceArmRouteTableCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).routeTablesClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Route Table %q (Resource Group %q): %+v", name, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_route_table", *existing.ID)
        }
    }

    id := d.Get("id").(string)
    location := azure.NormalizeLocation(d.Get("location").(string))
    disableBgpRoutePropagation := d.Get("disable_bgp_route_propagation").(bool)
    etag := d.Get("etag").(string)
    routes := d.Get("routes").([]interface{})
    t := d.Get("tags").(map[string]interface{})

    parameters := network.RouteTable{
        Etag: utils.String(etag),
        ID: utils.String(id),
        Location: utils.String(location),
        RouteTablePropertiesFormat: &network.RouteTablePropertiesFormat{
            DisableBgpRoutePropagation: utils.Bool(disableBgpRoutePropagation),
            Routes: expandArmRouteTableRoute(routes),
        },
        Tags: tags.Expand(t),
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, name, parameters)
    if err != nil {
        return fmt.Errorf("Error creating Route Table %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Route Table %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Route Table %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Route Table %q (Resource Group %q) ID", name, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmRouteTableRead(d, meta)
}

func resourceArmRouteTableRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).routeTablesClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["routeTables"]

    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Route Table %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Route Table %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if routeTablePropertiesFormat := resp.RouteTablePropertiesFormat; routeTablePropertiesFormat != nil {
        d.Set("disable_bgp_route_propagation", routeTablePropertiesFormat.DisableBgpRoutePropagation)
        d.Set("provisioning_state", routeTablePropertiesFormat.ProvisioningState)
        if err := d.Set("routes", flattenArmRouteTableRoute(routeTablePropertiesFormat.Routes)); err != nil {
            return fmt.Errorf("Error setting `routes`: %+v", err)
        }
        if err := d.Set("subnets", flattenArmRouteTableSubnet(routeTablePropertiesFormat.Subnets)); err != nil {
            return fmt.Errorf("Error setting `subnets`: %+v", err)
        }
    }
    d.Set("etag", resp.Etag)
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}


func resourceArmRouteTableDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).routeTablesClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["routeTables"]

    future, err := client.Delete(ctx, resourceGroup, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Route Table %q (Resource Group %q): %+v", name, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Route Table %q (Resource Group %q): %+v", name, resourceGroup, err)
        }
    }

    return nil
}

func expandArmRouteTableRoute(input []interface{}) *[]network.Route {
    results := make([]network.Route, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        id := v["id"].(string)
        addressPrefix := v["address_prefix"].(string)
        nextHopType := v["next_hop_type"].(string)
        nextHopIpAddress := v["next_hop_ip_address"].(string)
        name := v["name"].(string)
        etag := v["etag"].(string)

        result := network.Route{
            Etag: utils.String(etag),
            ID: utils.String(id),
            Name: utils.String(name),
            RoutePropertiesFormat: &network.RoutePropertiesFormat{
                AddressPrefix: utils.String(addressPrefix),
                NextHopIpAddress: utils.String(nextHopIpAddress),
                NextHopType: network.RouteNextHopType(nextHopType),
            },
        }

        results = append(results, result)
    }
    return &results
}


func flattenArmRouteTableRoute(input *[]network.Route) []interface{} {
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
        if routePropertiesFormat := item.RoutePropertiesFormat; routePropertiesFormat != nil {
            if addressPrefix := routePropertiesFormat.AddressPrefix; addressPrefix != nil {
                v["address_prefix"] = *addressPrefix
            }
            if nextHopIpAddress := routePropertiesFormat.NextHopIpAddress; nextHopIpAddress != nil {
                v["next_hop_ip_address"] = *nextHopIpAddress
            }
            v["next_hop_type"] = string(routePropertiesFormat.NextHopType)
        }
        if etag := item.Etag; etag != nil {
            v["etag"] = *etag
        }

        results = append(results, v)
    }

    return results
}

func flattenArmRouteTableSubnet(input *[]network.Subnet) []interface{} {
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
