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



func resourceArmInterfaceEndpoint() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmInterfaceEndpointCreateUpdate,
        Read: resourceArmInterfaceEndpointRead,
        Update: resourceArmInterfaceEndpointCreateUpdate,
        Delete: resourceArmInterfaceEndpointDelete,

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

            "interface_endpoint_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "endpoint_service": {
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

            "etag": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
            },

            "fqdn": {
                Type: schema.TypeString,
                Optional: true,
            },

            "subnet": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "address_prefix": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "address_prefixes": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                        "delegations": {
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
                                    "tags": tags.Schema(),
                                },
                            },
                        },
                        "resource_navigation_links": {
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
                                    "tags": tags.Schema(),
                                },
                            },
                        },
                        "service_association_links": {
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
                        "service_endpoint_policies": {
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
                                    "location": azure.SchemaLocation(),
                                    "tags": tags.Schema(),
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
                        "hosted_workloads": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                        "id": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "interface_endpoint": {
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
                        "tap_configurations": {
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
                                    "type": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                },
                            },
                        },
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

            "owner": {
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

func resourceArmInterfaceEndpointCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).interfaceEndpointsClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    interfaceEndpointName := d.Get("interface_endpoint_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, interfaceEndpointName)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Interface Endpoint (Interface Endpoint Name %q / Resource Group %q): %+v", interfaceEndpointName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_interface_endpoint", *existing.ID)
        }
    }

    id := d.Get("id").(string)
    location := azure.NormalizeLocation(d.Get("location").(string))
    endpointService := d.Get("endpoint_service").([]interface{})
    etag := d.Get("etag").(string)
    fqdn := d.Get("fqdn").(string)
    subnet := d.Get("subnet").([]interface{})
    t := d.Get("tags").(map[string]interface{})

    parameters := network.InterfaceEndpoint{
        Etag: utils.String(etag),
        ID: utils.String(id),
        Location: utils.String(location),
        InterfaceEndpointProperties: &network.InterfaceEndpointProperties{
            EndpointService: expandArmInterfaceEndpointEndpointService(endpointService),
            Fqdn: utils.String(fqdn),
            Subnet: expandArmInterfaceEndpointSubnet(subnet),
        },
        Tags: tags.Expand(t),
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, interfaceEndpointName, parameters)
    if err != nil {
        return fmt.Errorf("Error creating Interface Endpoint (Interface Endpoint Name %q / Resource Group %q): %+v", interfaceEndpointName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Interface Endpoint (Interface Endpoint Name %q / Resource Group %q): %+v", interfaceEndpointName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, interfaceEndpointName)
    if err != nil {
        return fmt.Errorf("Error retrieving Interface Endpoint (Interface Endpoint Name %q / Resource Group %q): %+v", interfaceEndpointName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Interface Endpoint (Interface Endpoint Name %q / Resource Group %q) ID", interfaceEndpointName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmInterfaceEndpointRead(d, meta)
}

func resourceArmInterfaceEndpointRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).interfaceEndpointsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    interfaceEndpointName := id.Path["interfaceEndpoints"]

    resp, err := client.Get(ctx, resourceGroup, interfaceEndpointName)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Interface Endpoint %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Interface Endpoint (Interface Endpoint Name %q / Resource Group %q): %+v", interfaceEndpointName, resourceGroup, err)
    }


    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if interfaceEndpointProperties := resp.InterfaceEndpointProperties; interfaceEndpointProperties != nil {
        if err := d.Set("endpoint_service", flattenArmInterfaceEndpointEndpointService(interfaceEndpointProperties.EndpointService)); err != nil {
            return fmt.Errorf("Error setting `endpoint_service`: %+v", err)
        }
        d.Set("fqdn", interfaceEndpointProperties.Fqdn)
        if err := d.Set("network_interfaces", flattenArmInterfaceEndpointInterface(interfaceEndpointProperties.NetworkInterfaces)); err != nil {
            return fmt.Errorf("Error setting `network_interfaces`: %+v", err)
        }
        d.Set("owner", interfaceEndpointProperties.Owner)
        d.Set("provisioning_state", interfaceEndpointProperties.ProvisioningState)
        if err := d.Set("subnet", flattenArmInterfaceEndpointSubnet(interfaceEndpointProperties.Subnet)); err != nil {
            return fmt.Errorf("Error setting `subnet`: %+v", err)
        }
    }
    d.Set("etag", resp.Etag)
    d.Set("interface_endpoint_name", interfaceEndpointName)
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}


func resourceArmInterfaceEndpointDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).interfaceEndpointsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    interfaceEndpointName := id.Path["interfaceEndpoints"]

    future, err := client.Delete(ctx, resourceGroup, interfaceEndpointName)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Interface Endpoint (Interface Endpoint Name %q / Resource Group %q): %+v", interfaceEndpointName, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Interface Endpoint (Interface Endpoint Name %q / Resource Group %q): %+v", interfaceEndpointName, resourceGroup, err)
        }
    }

    return nil
}

func expandArmInterfaceEndpointEndpointService(input []interface{}) *network.EndpointService {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    id := v["id"].(string)

    result := network.EndpointService{
        ID: utils.String(id),
    }
    return &result
}

func expandArmInterfaceEndpointSubnet(input []interface{}) *network.Subnet {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    id := v["id"].(string)
    addressPrefix := v["address_prefix"].(string)
    addressPrefixes := v["address_prefixes"].([]interface{})
    networkSecurityGroup := v["network_security_group"].([]interface{})
    routeTable := v["route_table"].([]interface{})
    serviceEndpoints := v["service_endpoints"].([]interface{})
    serviceEndpointPolicies := v["service_endpoint_policies"].([]interface{})
    resourceNavigationLinks := v["resource_navigation_links"].([]interface{})
    serviceAssociationLinks := v["service_association_links"].([]interface{})
    delegations := v["delegations"].([]interface{})
    name := v["name"].(string)
    etag := v["etag"].(string)

    result := network.Subnet{
        Etag: utils.String(etag),
        ID: utils.String(id),
        Name: utils.String(name),
        SubnetPropertiesFormat: &network.SubnetPropertiesFormat{
            AddressPrefix: utils.String(addressPrefix),
            AddressPrefixes: utils.ExpandStringSlice(addressPrefixes),
            Delegations: expandArmInterfaceEndpointDelegation(delegations),
            NetworkSecurityGroup: expandArmInterfaceEndpointSecurityGroup(networkSecurityGroup),
            ResourceNavigationLinks: expandArmInterfaceEndpointResourceNavigationLink(resourceNavigationLinks),
            RouteTable: expandArmInterfaceEndpointRouteTable(routeTable),
            ServiceAssociationLinks: expandArmInterfaceEndpointServiceAssociationLink(serviceAssociationLinks),
            ServiceEndpointPolicies: expandArmInterfaceEndpointServiceEndpointPolicy(serviceEndpointPolicies),
            ServiceEndpoints: expandArmInterfaceEndpointServiceEndpointPropertiesFormat(serviceEndpoints),
        },
    }
    return &result
}

func expandArmInterfaceEndpointDelegation(input []interface{}) *[]network.Delegation {
    results := make([]network.Delegation, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        id := v["id"].(string)
        name := v["name"].(string)
        etag := v["etag"].(string)

        result := network.Delegation{
            Etag: utils.String(etag),
            ID: utils.String(id),
            Name: utils.String(name),
        }

        results = append(results, result)
    }
    return &results
}

func expandArmInterfaceEndpointSecurityGroup(input []interface{}) *network.SecurityGroup {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    id := v["id"].(string)
    location := azure.NormalizeLocation(v["location"].(string))
    t := v["tags"].(map[string]interface{})
    etag := v["etag"].(string)

    result := network.SecurityGroup{
        Etag: utils.String(etag),
        ID: utils.String(id),
        Location: utils.String(location),
        Tags: tags.Expand(t),
    }
    return &result
}

func expandArmInterfaceEndpointResourceNavigationLink(input []interface{}) *[]network.ResourceNavigationLink {
    results := make([]network.ResourceNavigationLink, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        id := v["id"].(string)
        name := v["name"].(string)

        result := network.ResourceNavigationLink{
            ID: utils.String(id),
            Name: utils.String(name),
        }

        results = append(results, result)
    }
    return &results
}

func expandArmInterfaceEndpointRouteTable(input []interface{}) *network.RouteTable {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    id := v["id"].(string)
    location := azure.NormalizeLocation(v["location"].(string))
    t := v["tags"].(map[string]interface{})
    etag := v["etag"].(string)

    result := network.RouteTable{
        Etag: utils.String(etag),
        ID: utils.String(id),
        Location: utils.String(location),
        Tags: tags.Expand(t),
    }
    return &result
}

func expandArmInterfaceEndpointServiceAssociationLink(input []interface{}) *[]network.ServiceAssociationLink {
    results := make([]network.ServiceAssociationLink, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        id := v["id"].(string)
        name := v["name"].(string)

        result := network.ServiceAssociationLink{
            ID: utils.String(id),
            Name: utils.String(name),
        }

        results = append(results, result)
    }
    return &results
}

func expandArmInterfaceEndpointServiceEndpointPolicy(input []interface{}) *[]network.ServiceEndpointPolicy {
    results := make([]network.ServiceEndpointPolicy, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        id := v["id"].(string)
        location := azure.NormalizeLocation(v["location"].(string))
        t := v["tags"].(map[string]interface{})
        etag := v["etag"].(string)

        result := network.ServiceEndpointPolicy{
            Etag: utils.String(etag),
            ID: utils.String(id),
            Location: utils.String(location),
            Tags: tags.Expand(t),
        }

        results = append(results, result)
    }
    return &results
}

func expandArmInterfaceEndpointServiceEndpointPropertiesFormat(input []interface{}) *[]network.ServiceEndpointPropertiesFormat {
    results := make([]network.ServiceEndpointPropertiesFormat, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        service := v["service"].(string)
        locations := v["locations"].([]interface{})

        result := network.ServiceEndpointPropertiesFormat{
            Locations: utils.ExpandStringSlice(locations),
            Service: utils.String(service),
        }

        results = append(results, result)
    }
    return &results
}


func flattenArmInterfaceEndpointEndpointService(input *network.EndpointService) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if id := input.ID; id != nil {
        result["id"] = *id
    }

    return []interface{}{result}
}

func flattenArmInterfaceEndpointInterface(input *[]network.Interface) []interface{} {
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

func flattenArmInterfaceEndpointSubnet(input *network.Subnet) []interface{} {
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
    if subnetPropertiesFormat := input.SubnetPropertiesFormat; subnetPropertiesFormat != nil {
        if addressPrefix := subnetPropertiesFormat.AddressPrefix; addressPrefix != nil {
            result["address_prefix"] = *addressPrefix
        }
        result["address_prefixes"] = utils.FlattenStringSlice(subnetPropertiesFormat.AddressPrefixes)
        result["delegations"] = flattenArmInterfaceEndpointDelegation(subnetPropertiesFormat.Delegations)
        result["network_security_group"] = flattenArmInterfaceEndpointSecurityGroup(subnetPropertiesFormat.NetworkSecurityGroup)
        result["resource_navigation_links"] = flattenArmInterfaceEndpointResourceNavigationLink(subnetPropertiesFormat.ResourceNavigationLinks)
        result["route_table"] = flattenArmInterfaceEndpointRouteTable(subnetPropertiesFormat.RouteTable)
        result["service_association_links"] = flattenArmInterfaceEndpointServiceAssociationLink(subnetPropertiesFormat.ServiceAssociationLinks)
        result["service_endpoint_policies"] = flattenArmInterfaceEndpointServiceEndpointPolicy(subnetPropertiesFormat.ServiceEndpointPolicies)
        result["service_endpoints"] = flattenArmInterfaceEndpointServiceEndpointPropertiesFormat(subnetPropertiesFormat.ServiceEndpoints)
    }
    if etag := input.Etag; etag != nil {
        result["etag"] = *etag
    }

    return []interface{}{result}
}

func flattenArmInterfaceEndpointDelegation(input *[]network.Delegation) []interface{} {
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

        results = append(results, v)
    }

    return results
}

func flattenArmInterfaceEndpointSecurityGroup(input *network.SecurityGroup) []interface{} {
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

func flattenArmInterfaceEndpointResourceNavigationLink(input *[]network.ResourceNavigationLink) []interface{} {
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

func flattenArmInterfaceEndpointRouteTable(input *network.RouteTable) []interface{} {
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

func flattenArmInterfaceEndpointServiceAssociationLink(input *[]network.ServiceAssociationLink) []interface{} {
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

func flattenArmInterfaceEndpointServiceEndpointPolicy(input *[]network.ServiceEndpointPolicy) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        if id := item.ID; id != nil {
            v["id"] = *id
        }
        if location := item.Location; location != nil {
            v["location"] = azure.NormalizeLocation(*location)
        }
        if etag := item.Etag; etag != nil {
            v["etag"] = *etag
        }

        results = append(results, v)
    }

    return results
}

func flattenArmInterfaceEndpointServiceEndpointPropertiesFormat(input *[]network.ServiceEndpointPropertiesFormat) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        v["locations"] = utils.FlattenStringSlice(item.Locations)
        if service := item.Service; service != nil {
            v["service"] = *service
        }

        results = append(results, v)
    }

    return results
}
