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



func resourceArmProfile() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmProfileCreate,
        Read: resourceArmProfileRead,
        Update: resourceArmProfileUpdate,
        Delete: resourceArmProfileDelete,

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

            "location": azure.SchemaLocation(),

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "dns_config": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "relative_name": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "ttl": {
                            Type: schema.TypeInt,
                            Optional: true,
                        },
                    },
                },
            },

            "endpoints": {
                Type: schema.TypeList,
                Optional: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "custom_headers": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "name": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                    "value": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                },
                            },
                        },
                        "endpoint_location": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "endpoint_monitor_status": {
                            Type: schema.TypeString,
                            Optional: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(trafficmanager.CheckingEndpoint),
                                string(trafficmanager.Online),
                                string(trafficmanager.Degraded),
                                string(trafficmanager.Disabled),
                                string(trafficmanager.Inactive),
                                string(trafficmanager.Stopped),
                            }, false),
                            Default: string(trafficmanager.CheckingEndpoint),
                        },
                        "endpoint_status": {
                            Type: schema.TypeString,
                            Optional: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(trafficmanager.Enabled),
                                string(trafficmanager.Disabled),
                            }, false),
                            Default: string(trafficmanager.Enabled),
                        },
                        "geo_mapping": {
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
                        "min_child_endpoints": {
                            Type: schema.TypeInt,
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
                        "subnets": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "first": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                    "last": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                    "scope": {
                                        Type: schema.TypeInt,
                                        Optional: true,
                                    },
                                },
                            },
                        },
                        "target": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "target_resource_id": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "type": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "weight": {
                            Type: schema.TypeInt,
                            Optional: true,
                        },
                    },
                },
            },

            "max_return": {
                Type: schema.TypeInt,
                Optional: true,
            },

            "monitor_config": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "custom_headers": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "name": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                    "value": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                },
                            },
                        },
                        "expected_status_code_ranges": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "max": {
                                        Type: schema.TypeInt,
                                        Optional: true,
                                    },
                                    "min": {
                                        Type: schema.TypeInt,
                                        Optional: true,
                                    },
                                },
                            },
                        },
                        "interval_in_seconds": {
                            Type: schema.TypeInt,
                            Optional: true,
                        },
                        "path": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "port": {
                            Type: schema.TypeInt,
                            Optional: true,
                        },
                        "profile_monitor_status": {
                            Type: schema.TypeString,
                            Optional: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(trafficmanager.CheckingEndpoints),
                                string(trafficmanager.Online),
                                string(trafficmanager.Degraded),
                                string(trafficmanager.Disabled),
                                string(trafficmanager.Inactive),
                            }, false),
                            Default: string(trafficmanager.CheckingEndpoints),
                        },
                        "protocol": {
                            Type: schema.TypeString,
                            Optional: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(trafficmanager.HTTP),
                                string(trafficmanager.HTTPS),
                                string(trafficmanager.TCP),
                            }, false),
                            Default: string(trafficmanager.HTTP),
                        },
                        "timeout_in_seconds": {
                            Type: schema.TypeInt,
                            Optional: true,
                        },
                        "tolerated_number_of_failures": {
                            Type: schema.TypeInt,
                            Optional: true,
                        },
                    },
                },
            },

            "profile_status": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(trafficmanager.Enabled),
                    string(trafficmanager.Disabled),
                }, false),
                Default: string(trafficmanager.Enabled),
            },

            "traffic_routing_method": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(trafficmanager.Performance),
                    string(trafficmanager.Priority),
                    string(trafficmanager.Weighted),
                    string(trafficmanager.Geographic),
                    string(trafficmanager.MultiValue),
                    string(trafficmanager.Subnet),
                }, false),
                Default: string(trafficmanager.Performance),
            },

            "traffic_view_enrollment_status": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(trafficmanager.Enabled),
                    string(trafficmanager.Disabled),
                }, false),
                Default: string(trafficmanager.Enabled),
            },

            "type": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
            },

            "tags": tags.Schema(),
        },
    }
}

func resourceArmProfileCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).profilesClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Profile %q (Resource Group %q): %+v", name, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_profile", *existing.ID)
        }
    }

    id := d.Get("id").(string)
    name := d.Get("name").(string)
    location := azure.NormalizeLocation(d.Get("location").(string))
    dnsConfig := d.Get("dns_config").([]interface{})
    endpoints := d.Get("endpoints").([]interface{})
    maxReturn := d.Get("max_return").(int)
    monitorConfig := d.Get("monitor_config").([]interface{})
    profileStatus := d.Get("profile_status").(string)
    trafficRoutingMethod := d.Get("traffic_routing_method").(string)
    trafficViewEnrollmentStatus := d.Get("traffic_view_enrollment_status").(string)
    type := d.Get("type").(string)
    t := d.Get("tags").(map[string]interface{})

    parameters := trafficmanager.Profile{
        ID: utils.String(id),
        Location: utils.String(location),
        Name: utils.String(name),
        ProfileProperties: &trafficmanager.ProfileProperties{
            DnsConfig: expandArmProfileDnsConfig(dnsConfig),
            Endpoints: expandArmProfileEndpoint(endpoints),
            MaxReturn: utils.Int64(int64(maxReturn)),
            MonitorConfig: expandArmProfileMonitorConfig(monitorConfig),
            ProfileStatus: trafficmanager.ProfileStatus(profileStatus),
            TrafficRoutingMethod: trafficmanager.TrafficRoutingMethod(trafficRoutingMethod),
            TrafficViewEnrollmentStatus: trafficmanager.TrafficViewEnrollmentStatus(trafficViewEnrollmentStatus),
        },
        Tags: tags.Expand(t),
        Type: utils.String(type),
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroup, name, parameters); err != nil {
        return fmt.Errorf("Error creating Profile %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Profile %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Profile %q (Resource Group %q) ID", name, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmProfileRead(d, meta)
}

func resourceArmProfileRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).profilesClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["trafficmanagerprofiles"]

    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Profile %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Profile %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if profileProperties := resp.ProfileProperties; profileProperties != nil {
        if err := d.Set("dns_config", flattenArmProfileDnsConfig(profileProperties.DnsConfig)); err != nil {
            return fmt.Errorf("Error setting `dns_config`: %+v", err)
        }
        if err := d.Set("endpoints", flattenArmProfileEndpoint(profileProperties.Endpoints)); err != nil {
            return fmt.Errorf("Error setting `endpoints`: %+v", err)
        }
        d.Set("max_return", int(*profileProperties.MaxReturn))
        if err := d.Set("monitor_config", flattenArmProfileMonitorConfig(profileProperties.MonitorConfig)); err != nil {
            return fmt.Errorf("Error setting `monitor_config`: %+v", err)
        }
        d.Set("profile_status", string(profileProperties.ProfileStatus))
        d.Set("traffic_routing_method", string(profileProperties.TrafficRoutingMethod))
        d.Set("traffic_view_enrollment_status", string(profileProperties.TrafficViewEnrollmentStatus))
    }
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}

func resourceArmProfileUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).profilesClient
    ctx := meta.(*ArmClient).StopContext

    id := d.Get("id").(string)
    name := d.Get("name").(string)
    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    dnsConfig := d.Get("dns_config").([]interface{})
    endpoints := d.Get("endpoints").([]interface{})
    maxReturn := d.Get("max_return").(int)
    monitorConfig := d.Get("monitor_config").([]interface{})
    profileStatus := d.Get("profile_status").(string)
    trafficRoutingMethod := d.Get("traffic_routing_method").(string)
    trafficViewEnrollmentStatus := d.Get("traffic_view_enrollment_status").(string)
    type := d.Get("type").(string)
    t := d.Get("tags").(map[string]interface{})

    parameters := trafficmanager.Profile{
        ID: utils.String(id),
        Location: utils.String(location),
        Name: utils.String(name),
        ProfileProperties: &trafficmanager.ProfileProperties{
            DnsConfig: expandArmProfileDnsConfig(dnsConfig),
            Endpoints: expandArmProfileEndpoint(endpoints),
            MaxReturn: utils.Int64(int64(maxReturn)),
            MonitorConfig: expandArmProfileMonitorConfig(monitorConfig),
            ProfileStatus: trafficmanager.ProfileStatus(profileStatus),
            TrafficRoutingMethod: trafficmanager.TrafficRoutingMethod(trafficRoutingMethod),
            TrafficViewEnrollmentStatus: trafficmanager.TrafficViewEnrollmentStatus(trafficViewEnrollmentStatus),
        },
        Tags: tags.Expand(t),
        Type: utils.String(type),
    }


    if _, err := client.Update(ctx, resourceGroup, name, parameters); err != nil {
        return fmt.Errorf("Error updating Profile %q (Resource Group %q): %+v", name, resourceGroup, err)
    }

    return resourceArmProfileRead(d, meta)
}

func resourceArmProfileDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).profilesClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["trafficmanagerprofiles"]

    if _, err := client.Delete(ctx, resourceGroup, name); err != nil {
        return fmt.Errorf("Error deleting Profile %q (Resource Group %q): %+v", name, resourceGroup, err)
    }

    return nil
}

func expandArmProfileDnsConfig(input []interface{}) *trafficmanager.DnsConfig {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    relativeName := v["relative_name"].(string)
    ttl := v["ttl"].(int)

    result := trafficmanager.DnsConfig{
        RelativeName: utils.String(relativeName),
        Ttl: utils.Int64(int64(ttl)),
    }
    return &result
}

func expandArmProfileEndpoint(input []interface{}) *[]trafficmanager.Endpoint {
    results := make([]trafficmanager.Endpoint, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        id := v["id"].(string)
        name := v["name"].(string)
        type := v["type"].(string)
        targetResourceId := v["target_resource_id"].(string)
        target := v["target"].(string)
        endpointStatus := v["endpoint_status"].(string)
        weight := v["weight"].(int)
        priority := v["priority"].(int)
        endpointLocation := v["endpoint_location"].(string)
        endpointMonitorStatus := v["endpoint_monitor_status"].(string)
        minChildEndpoints := v["min_child_endpoints"].(int)
        geoMapping := v["geo_mapping"].([]interface{})
        subnets := v["subnets"].([]interface{})
        customHeaders := v["custom_headers"].([]interface{})

        result := trafficmanager.Endpoint{
            ID: utils.String(id),
            Name: utils.String(name),
            EndpointProperties: &trafficmanager.EndpointProperties{
                CustomHeaders: expandArmProfileEndpointProperties_customHeadersItem(customHeaders),
                EndpointLocation: utils.String(endpointLocation),
                EndpointMonitorStatus: trafficmanager.EndpointMonitorStatus(endpointMonitorStatus),
                EndpointStatus: trafficmanager.EndpointStatus(endpointStatus),
                GeoMapping: utils.ExpandStringSlice(geoMapping),
                MinChildEndpoints: utils.Int64(int64(minChildEndpoints)),
                Priority: utils.Int64(int64(priority)),
                Subnets: expandArmProfileEndpointProperties_subnetsItem(subnets),
                Target: utils.String(target),
                TargetResourceID: utils.String(targetResourceId),
                Weight: utils.Int64(int64(weight)),
            },
            Type: utils.String(type),
        }

        results = append(results, result)
    }
    return &results
}

func expandArmProfileMonitorConfig(input []interface{}) *trafficmanager.MonitorConfig {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    profileMonitorStatus := v["profile_monitor_status"].(string)
    protocol := v["protocol"].(string)
    port := v["port"].(int)
    path := v["path"].(string)
    intervalInSeconds := v["interval_in_seconds"].(int)
    timeoutInSeconds := v["timeout_in_seconds"].(int)
    toleratedNumberOfFailures := v["tolerated_number_of_failures"].(int)
    customHeaders := v["custom_headers"].([]interface{})
    expectedStatusCodeRanges := v["expected_status_code_ranges"].([]interface{})

    result := trafficmanager.MonitorConfig{
        CustomHeaders: expandArmProfileMonitorConfig_customHeadersItem(customHeaders),
        ExpectedStatusCodeRanges: expandArmProfileMonitorConfig_expectedStatusCodeRangesItem(expectedStatusCodeRanges),
        IntervalInSeconds: utils.Int64(int64(intervalInSeconds)),
        Path: utils.String(path),
        Port: utils.Int64(int64(port)),
        ProfileMonitorStatus: trafficmanager.ProfileMonitorStatus(profileMonitorStatus),
        Protocol: trafficmanager.MonitorProtocol(protocol),
        TimeoutInSeconds: utils.Int64(int64(timeoutInSeconds)),
        ToleratedNumberOfFailures: utils.Int64(int64(toleratedNumberOfFailures)),
    }
    return &result
}

func expandArmProfileEndpointProperties_customHeadersItem(input []interface{}) *[]trafficmanager.EndpointProperties_customHeadersItem {
    results := make([]trafficmanager.EndpointProperties_customHeadersItem, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        name := v["name"].(string)
        value := v["value"].(string)

        result := trafficmanager.EndpointProperties_customHeadersItem{
            Name: utils.String(name),
            Value: utils.String(value),
        }

        results = append(results, result)
    }
    return &results
}

func expandArmProfileEndpointProperties_subnetsItem(input []interface{}) *[]trafficmanager.EndpointProperties_subnetsItem {
    results := make([]trafficmanager.EndpointProperties_subnetsItem, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        first := v["first"].(string)
        last := v["last"].(string)
        scope := v["scope"].(int)

        result := trafficmanager.EndpointProperties_subnetsItem{
            First: utils.String(first),
            Last: utils.String(last),
            Scope: utils.Int(scope),
        }

        results = append(results, result)
    }
    return &results
}

func expandArmProfileMonitorConfig_customHeadersItem(input []interface{}) *[]trafficmanager.MonitorConfig_customHeadersItem {
    results := make([]trafficmanager.MonitorConfig_customHeadersItem, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        name := v["name"].(string)
        value := v["value"].(string)

        result := trafficmanager.MonitorConfig_customHeadersItem{
            Name: utils.String(name),
            Value: utils.String(value),
        }

        results = append(results, result)
    }
    return &results
}

func expandArmProfileMonitorConfig_expectedStatusCodeRangesItem(input []interface{}) *[]trafficmanager.MonitorConfig_expectedStatusCodeRangesItem {
    results := make([]trafficmanager.MonitorConfig_expectedStatusCodeRangesItem, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        min := v["min"].(int)
        max := v["max"].(int)

        result := trafficmanager.MonitorConfig_expectedStatusCodeRangesItem{
            Max: utils.Int(max),
            Min: utils.Int(min),
        }

        results = append(results, result)
    }
    return &results
}


func flattenArmProfileDnsConfig(input *trafficmanager.DnsConfig) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if relativeName := input.RelativeName; relativeName != nil {
        result["relative_name"] = *relativeName
    }
    if ttl := input.Ttl; ttl != nil {
        result["ttl"] = int(*ttl)
    }

    return []interface{}{result}
}

func flattenArmProfileEndpoint(input *[]trafficmanager.Endpoint) []interface{} {
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
        if endpointProperties := item.EndpointProperties; endpointProperties != nil {
            v["custom_headers"] = flattenArmProfileEndpointProperties_customHeadersItem(endpointProperties.CustomHeaders)
            if endpointLocation := endpointProperties.EndpointLocation; endpointLocation != nil {
                v["endpoint_location"] = *endpointLocation
            }
            v["endpoint_monitor_status"] = string(endpointProperties.EndpointMonitorStatus)
            v["endpoint_status"] = string(endpointProperties.EndpointStatus)
            v["geo_mapping"] = utils.FlattenStringSlice(endpointProperties.GeoMapping)
            if minChildEndpoints := endpointProperties.MinChildEndpoints; minChildEndpoints != nil {
                v["min_child_endpoints"] = int(*minChildEndpoints)
            }
            if priority := endpointProperties.Priority; priority != nil {
                v["priority"] = int(*priority)
            }
            v["subnets"] = flattenArmProfileEndpointProperties_subnetsItem(endpointProperties.Subnets)
            if target := endpointProperties.Target; target != nil {
                v["target"] = *target
            }
            if targetResourceId := endpointProperties.TargetResourceID; targetResourceId != nil {
                v["target_resource_id"] = *targetResourceId
            }
            if weight := endpointProperties.Weight; weight != nil {
                v["weight"] = int(*weight)
            }
        }
        if type := item.Type; type != nil {
            v["type"] = *type
        }

        results = append(results, v)
    }

    return results
}

func flattenArmProfileMonitorConfig(input *trafficmanager.MonitorConfig) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    result["custom_headers"] = flattenArmProfileMonitorConfig_customHeadersItem(input.CustomHeaders)
    result["expected_status_code_ranges"] = flattenArmProfileMonitorConfig_expectedStatusCodeRangesItem(input.ExpectedStatusCodeRanges)
    if intervalInSeconds := input.IntervalInSeconds; intervalInSeconds != nil {
        result["interval_in_seconds"] = int(*intervalInSeconds)
    }
    if path := input.Path; path != nil {
        result["path"] = *path
    }
    if port := input.Port; port != nil {
        result["port"] = int(*port)
    }
    result["profile_monitor_status"] = string(input.ProfileMonitorStatus)
    result["protocol"] = string(input.Protocol)
    if timeoutInSeconds := input.TimeoutInSeconds; timeoutInSeconds != nil {
        result["timeout_in_seconds"] = int(*timeoutInSeconds)
    }
    if toleratedNumberOfFailures := input.ToleratedNumberOfFailures; toleratedNumberOfFailures != nil {
        result["tolerated_number_of_failures"] = int(*toleratedNumberOfFailures)
    }

    return []interface{}{result}
}

func flattenArmProfileEndpointProperties_customHeadersItem(input *[]trafficmanager.EndpointProperties_customHeadersItem) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        if name := item.Name; name != nil {
            v["name"] = *name
        }
        if value := item.Value; value != nil {
            v["value"] = *value
        }

        results = append(results, v)
    }

    return results
}

func flattenArmProfileEndpointProperties_subnetsItem(input *[]trafficmanager.EndpointProperties_subnetsItem) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        if first := item.First; first != nil {
            v["first"] = *first
        }
        if last := item.Last; last != nil {
            v["last"] = *last
        }
        if scope := item.Scope; scope != nil {
            v["scope"] = *scope
        }

        results = append(results, v)
    }

    return results
}

func flattenArmProfileMonitorConfig_customHeadersItem(input *[]trafficmanager.MonitorConfig_customHeadersItem) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        if name := item.Name; name != nil {
            v["name"] = *name
        }
        if value := item.Value; value != nil {
            v["value"] = *value
        }

        results = append(results, v)
    }

    return results
}

func flattenArmProfileMonitorConfig_expectedStatusCodeRangesItem(input *[]trafficmanager.MonitorConfig_expectedStatusCodeRangesItem) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        if max := item.Max; max != nil {
            v["max"] = *max
        }
        if min := item.Min; min != nil {
            v["min"] = *min
        }

        results = append(results, v)
    }

    return results
}
