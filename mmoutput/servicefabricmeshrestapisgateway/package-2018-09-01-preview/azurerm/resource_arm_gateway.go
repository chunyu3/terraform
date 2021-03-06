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



func resourceArmGateway() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmGatewayCreateUpdate,
        Read: resourceArmGatewayRead,
        Update: resourceArmGatewayCreateUpdate,
        Delete: resourceArmGatewayDelete,

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

            "destination_network": {
                Type: schema.TypeList,
                Required: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "endpoint_refs": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
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
                    },
                },
            },

            "source_network": {
                Type: schema.TypeList,
                Required: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "endpoint_refs": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
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
                    },
                },
            },

            "description": {
                Type: schema.TypeString,
                Optional: true,
            },

            "http": {
                Type: schema.TypeList,
                Optional: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "hosts": {
                            Type: schema.TypeList,
                            Required: true,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "name": {
                                        Type: schema.TypeString,
                                        Required: true,
                                        ValidateFunc: validate.NoEmptyStrings,
                                    },
                                    "routes": {
                                        Type: schema.TypeList,
                                        Required: true,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "destination": {
                                                    Type: schema.TypeList,
                                                    Required: true,
                                                    MaxItems: 1,
                                                    Elem: &schema.Resource{
                                                        Schema: map[string]*schema.Schema{
                                                            "application_name": {
                                                                Type: schema.TypeString,
                                                                Required: true,
                                                                ValidateFunc: validate.NoEmptyStrings,
                                                            },
                                                            "endpoint_name": {
                                                                Type: schema.TypeString,
                                                                Required: true,
                                                                ValidateFunc: validate.NoEmptyStrings,
                                                            },
                                                            "service_name": {
                                                                Type: schema.TypeString,
                                                                Required: true,
                                                                ValidateFunc: validate.NoEmptyStrings,
                                                            },
                                                        },
                                                    },
                                                },
                                                "match": {
                                                    Type: schema.TypeList,
                                                    Required: true,
                                                    MaxItems: 1,
                                                    Elem: &schema.Resource{
                                                        Schema: map[string]*schema.Schema{
                                                            "path": {
                                                                Type: schema.TypeList,
                                                                Required: true,
                                                                MaxItems: 1,
                                                                Elem: &schema.Resource{
                                                                    Schema: map[string]*schema.Schema{
                                                                        "type": {
                                                                            Type: schema.TypeString,
                                                                            Required: true,
                                                                            ValidateFunc: validate.NoEmptyStrings,
                                                                        },
                                                                        "value": {
                                                                            Type: schema.TypeString,
                                                                            Required: true,
                                                                            ValidateFunc: validate.NoEmptyStrings,
                                                                        },
                                                                        "rewrite": {
                                                                            Type: schema.TypeString,
                                                                            Optional: true,
                                                                        },
                                                                    },
                                                                },
                                                            },
                                                            "headers": {
                                                                Type: schema.TypeList,
                                                                Optional: true,
                                                                Elem: &schema.Resource{
                                                                    Schema: map[string]*schema.Schema{
                                                                        "name": {
                                                                            Type: schema.TypeString,
                                                                            Required: true,
                                                                            ValidateFunc: validate.NoEmptyStrings,
                                                                        },
                                                                        "type": {
                                                                            Type: schema.TypeString,
                                                                            Optional: true,
                                                                            ValidateFunc: validation.StringInSlice([]string{
                                                                                string(servicefabricmeshrestapis.exact),
                                                                            }, false),
                                                                            Default: string(servicefabricmeshrestapis.exact),
                                                                        },
                                                                        "value": {
                                                                            Type: schema.TypeString,
                                                                            Optional: true,
                                                                        },
                                                                    },
                                                                },
                                                            },
                                                        },
                                                    },
                                                },
                                                "name": {
                                                    Type: schema.TypeString,
                                                    Required: true,
                                                    ValidateFunc: validate.NoEmptyStrings,
                                                },
                                            },
                                        },
                                    },
                                },
                            },
                        },
                        "name": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "port": {
                            Type: schema.TypeInt,
                            Required: true,
                        },
                    },
                },
            },

            "tcp": {
                Type: schema.TypeList,
                Optional: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "destination": {
                            Type: schema.TypeList,
                            Required: true,
                            MaxItems: 1,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "application_name": {
                                        Type: schema.TypeString,
                                        Required: true,
                                        ValidateFunc: validate.NoEmptyStrings,
                                    },
                                    "endpoint_name": {
                                        Type: schema.TypeString,
                                        Required: true,
                                        ValidateFunc: validate.NoEmptyStrings,
                                    },
                                    "service_name": {
                                        Type: schema.TypeString,
                                        Required: true,
                                        ValidateFunc: validate.NoEmptyStrings,
                                    },
                                },
                            },
                        },
                        "name": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "port": {
                            Type: schema.TypeInt,
                            Required: true,
                        },
                    },
                },
            },

            "ip_address": {
                Type: schema.TypeString,
                Computed: true,
            },

            "provisioning_state": {
                Type: schema.TypeString,
                Computed: true,
            },

            "status": {
                Type: schema.TypeString,
                Computed: true,
            },

            "status_details": {
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

func resourceArmGatewayCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).gatewayClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Gateway %q (Resource Group %q): %+v", name, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_gateway", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    description := d.Get("description").(string)
    destinationNetwork := d.Get("destination_network").([]interface{})
    http := d.Get("http").([]interface{})
    sourceNetwork := d.Get("source_network").([]interface{})
    tcp := d.Get("tcp").([]interface{})
    t := d.Get("tags").(map[string]interface{})

    gatewayResourceDescription := servicefabricmeshrestapis.GatewayResourceDescription{
        Location: utils.String(location),
        GatewayResourceProperties: &servicefabricmeshrestapis.GatewayResourceProperties{
            Description: utils.String(description),
            DestinationNetwork: expandArmGatewayNetworkRef(destinationNetwork),
            HTTP: expandArmGatewayHttpConfig(http),
            SourceNetwork: expandArmGatewayNetworkRef(sourceNetwork),
            TCP: expandArmGatewayTcpConfig(tcp),
        },
        Tags: tags.Expand(t),
    }


    if _, err := client.Create(ctx, resourceGroup, name, gatewayResourceDescription); err != nil {
        return fmt.Errorf("Error creating Gateway %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Gateway %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Gateway %q (Resource Group %q) ID", name, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmGatewayRead(d, meta)
}

func resourceArmGatewayRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).gatewayClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["gateways"]

    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Gateway %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Gateway %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if gatewayResourceProperties := resp.GatewayResourceProperties; gatewayResourceProperties != nil {
        d.Set("description", gatewayResourceProperties.Description)
        if err := d.Set("destination_network", flattenArmGatewayNetworkRef(gatewayResourceProperties.DestinationNetwork)); err != nil {
            return fmt.Errorf("Error setting `destination_network`: %+v", err)
        }
        if err := d.Set("http", flattenArmGatewayHttpConfig(gatewayResourceProperties.HTTP)); err != nil {
            return fmt.Errorf("Error setting `http`: %+v", err)
        }
        d.Set("ip_address", gatewayResourceProperties.IPAddress)
        d.Set("provisioning_state", gatewayResourceProperties.ProvisioningState)
        if err := d.Set("source_network", flattenArmGatewayNetworkRef(gatewayResourceProperties.SourceNetwork)); err != nil {
            return fmt.Errorf("Error setting `source_network`: %+v", err)
        }
        d.Set("status", string(gatewayResourceProperties.Status))
        d.Set("status_details", gatewayResourceProperties.StatusDetails)
        if err := d.Set("tcp", flattenArmGatewayTcpConfig(gatewayResourceProperties.TCP)); err != nil {
            return fmt.Errorf("Error setting `tcp`: %+v", err)
        }
    }
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}


func resourceArmGatewayDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).gatewayClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["gateways"]

    if _, err := client.Delete(ctx, resourceGroup, name); err != nil {
        return fmt.Errorf("Error deleting Gateway %q (Resource Group %q): %+v", name, resourceGroup, err)
    }

    return nil
}

func expandArmGatewayNetworkRef(input []interface{}) *servicefabricmeshrestapis.NetworkRef {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    name := v["name"].(string)
    endpointRefs := v["endpoint_refs"].([]interface{})

    result := servicefabricmeshrestapis.NetworkRef{
        EndpointRefs: expandArmGatewayEndpointRef(endpointRefs),
        Name: utils.String(name),
    }
    return &result
}

func expandArmGatewayHttpConfig(input []interface{}) *[]servicefabricmeshrestapis.HttpConfig {
    results := make([]servicefabricmeshrestapis.HttpConfig, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        name := v["name"].(string)
        port := v["port"].(int)
        hosts := v["hosts"].([]interface{})

        result := servicefabricmeshrestapis.HttpConfig{
            Hosts: expandArmGatewayHttpHostConfig(hosts),
            Name: utils.String(name),
            Port: utils.Int(port),
        }

        results = append(results, result)
    }
    return &results
}

func expandArmGatewayTcpConfig(input []interface{}) *[]servicefabricmeshrestapis.TcpConfig {
    results := make([]servicefabricmeshrestapis.TcpConfig, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        name := v["name"].(string)
        port := v["port"].(int)
        destination := v["destination"].([]interface{})

        result := servicefabricmeshrestapis.TcpConfig{
            Destination: expandArmGatewayGatewayDestination(destination),
            Name: utils.String(name),
            Port: utils.Int(port),
        }

        results = append(results, result)
    }
    return &results
}

func expandArmGatewayEndpointRef(input []interface{}) *[]servicefabricmeshrestapis.EndpointRef {
    results := make([]servicefabricmeshrestapis.EndpointRef, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        name := v["name"].(string)

        result := servicefabricmeshrestapis.EndpointRef{
            Name: utils.String(name),
        }

        results = append(results, result)
    }
    return &results
}

func expandArmGatewayHttpHostConfig(input []interface{}) *[]servicefabricmeshrestapis.HttpHostConfig {
    results := make([]servicefabricmeshrestapis.HttpHostConfig, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        name := v["name"].(string)
        routes := v["routes"].([]interface{})

        result := servicefabricmeshrestapis.HttpHostConfig{
            Name: utils.String(name),
            Routes: expandArmGatewayHttpRouteConfig(routes),
        }

        results = append(results, result)
    }
    return &results
}

func expandArmGatewayGatewayDestination(input []interface{}) *servicefabricmeshrestapis.GatewayDestination {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    applicationName := v["application_name"].(string)
    serviceName := v["service_name"].(string)
    endpointName := v["endpoint_name"].(string)

    result := servicefabricmeshrestapis.GatewayDestination{
        ApplicationName: utils.String(applicationName),
        EndpointName: utils.String(endpointName),
        ServiceName: utils.String(serviceName),
    }
    return &result
}

func expandArmGatewayHttpRouteConfig(input []interface{}) *[]servicefabricmeshrestapis.HttpRouteConfig {
    results := make([]servicefabricmeshrestapis.HttpRouteConfig, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        name := v["name"].(string)
        match := v["match"].([]interface{})
        destination := v["destination"].([]interface{})

        result := servicefabricmeshrestapis.HttpRouteConfig{
            Destination: expandArmGatewayGatewayDestination(destination),
            Match: expandArmGatewayHttpRouteMatchRule(match),
            Name: utils.String(name),
        }

        results = append(results, result)
    }
    return &results
}

func expandArmGatewayHttpRouteMatchRule(input []interface{}) *servicefabricmeshrestapis.HttpRouteMatchRule {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    path := v["path"].([]interface{})
    headers := v["headers"].([]interface{})

    result := servicefabricmeshrestapis.HttpRouteMatchRule{
        Headers: expandArmGatewayHttpRouteMatchHeader(headers),
        Path: expandArmGatewayHttpRouteMatchPath(path),
    }
    return &result
}

func expandArmGatewayHttpRouteMatchHeader(input []interface{}) *[]servicefabricmeshrestapis.HttpRouteMatchHeader {
    results := make([]servicefabricmeshrestapis.HttpRouteMatchHeader, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        name := v["name"].(string)
        value := v["value"].(string)
        type := v["type"].(string)

        result := servicefabricmeshrestapis.HttpRouteMatchHeader{
            Name: utils.String(name),
            Type: servicefabricmeshrestapis.HeaderMatchType(type),
            Value: utils.String(value),
        }

        results = append(results, result)
    }
    return &results
}

func expandArmGatewayHttpRouteMatchPath(input []interface{}) *servicefabricmeshrestapis.HttpRouteMatchPath {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    value := v["value"].(string)
    rewrite := v["rewrite"].(string)
    type := v["type"].(string)

    result := servicefabricmeshrestapis.HttpRouteMatchPath{
        Rewrite: utils.String(rewrite),
        Type: utils.String(type),
        Value: utils.String(value),
    }
    return &result
}


func flattenArmGatewayNetworkRef(input *servicefabricmeshrestapis.NetworkRef) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if name := input.Name; name != nil {
        result["name"] = *name
    }
    result["endpoint_refs"] = flattenArmGatewayEndpointRef(input.EndpointRefs)

    return []interface{}{result}
}

func flattenArmGatewayHttpConfig(input *[]servicefabricmeshrestapis.HttpConfig) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        if name := item.Name; name != nil {
            v["name"] = *name
        }
        v["hosts"] = flattenArmGatewayHttpHostConfig(item.Hosts)
        if port := item.Port; port != nil {
            v["port"] = *port
        }

        results = append(results, v)
    }

    return results
}

func flattenArmGatewayTcpConfig(input *[]servicefabricmeshrestapis.TcpConfig) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        if name := item.Name; name != nil {
            v["name"] = *name
        }
        v["destination"] = flattenArmGatewayGatewayDestination(item.Destination)
        if port := item.Port; port != nil {
            v["port"] = *port
        }

        results = append(results, v)
    }

    return results
}

func flattenArmGatewayEndpointRef(input *[]servicefabricmeshrestapis.EndpointRef) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        if name := item.Name; name != nil {
            v["name"] = *name
        }

        results = append(results, v)
    }

    return results
}

func flattenArmGatewayHttpHostConfig(input *[]servicefabricmeshrestapis.HttpHostConfig) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        if name := item.Name; name != nil {
            v["name"] = *name
        }
        v["routes"] = flattenArmGatewayHttpRouteConfig(item.Routes)

        results = append(results, v)
    }

    return results
}

func flattenArmGatewayGatewayDestination(input *servicefabricmeshrestapis.GatewayDestination) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if applicationName := input.ApplicationName; applicationName != nil {
        result["application_name"] = *applicationName
    }
    if endpointName := input.EndpointName; endpointName != nil {
        result["endpoint_name"] = *endpointName
    }
    if serviceName := input.ServiceName; serviceName != nil {
        result["service_name"] = *serviceName
    }

    return []interface{}{result}
}

func flattenArmGatewayHttpRouteConfig(input *[]servicefabricmeshrestapis.HttpRouteConfig) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        if name := item.Name; name != nil {
            v["name"] = *name
        }
        v["destination"] = flattenArmGatewayGatewayDestination(item.Destination)
        v["match"] = flattenArmGatewayHttpRouteMatchRule(item.Match)

        results = append(results, v)
    }

    return results
}

func flattenArmGatewayHttpRouteMatchRule(input *servicefabricmeshrestapis.HttpRouteMatchRule) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    result["headers"] = flattenArmGatewayHttpRouteMatchHeader(input.Headers)
    result["path"] = flattenArmGatewayHttpRouteMatchPath(input.Path)

    return []interface{}{result}
}

func flattenArmGatewayHttpRouteMatchHeader(input *[]servicefabricmeshrestapis.HttpRouteMatchHeader) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        if name := item.Name; name != nil {
            v["name"] = *name
        }
        v["type"] = string(item.Type)
        if value := item.Value; value != nil {
            v["value"] = *value
        }

        results = append(results, v)
    }

    return results
}

func flattenArmGatewayHttpRouteMatchPath(input *servicefabricmeshrestapis.HttpRouteMatchPath) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if rewrite := input.Rewrite; rewrite != nil {
        result["rewrite"] = *rewrite
    }
    if type := input.Type; type != nil {
        result["type"] = *type
    }
    if value := input.Value; value != nil {
        result["value"] = *value
    }

    return []interface{}{result}
}
