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
                Computed: true,
            },

            "location": azure.SchemaLocation(),

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "profile_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "dns_config": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "fqdn": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
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
                        "endpoint_location": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "endpoint_monitor_status": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "endpoint_status": {
                            Type: schema.TypeString,
                            Optional: true,
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

            "monitor_config": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
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
                        },
                        "protocol": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },

            "profile_status": {
                Type: schema.TypeString,
                Optional: true,
            },

            "traffic_routing_method": {
                Type: schema.TypeString,
                Optional: true,
            },

            "type": {
                Type: schema.TypeString,
                Computed: true,
            },

            "tags": tags.Schema(),
        },
    }
}

func resourceArmProfileCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).profilesClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    profileName := d.Get("profile_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, profileName)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Profile (Profile Name %q / Resource Group %q): %+v", profileName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_profile", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    dnsConfig := d.Get("dns_config").([]interface{})
    endpoints := d.Get("endpoints").([]interface{})
    monitorConfig := d.Get("monitor_config").([]interface{})
    profileStatus := d.Get("profile_status").(string)
    trafficRoutingMethod := d.Get("traffic_routing_method").(string)
    t := d.Get("tags").(map[string]interface{})

    parameters := trafficmanager.Profile{
        Location: utils.String(location),
        ProfileProperties: &trafficmanager.ProfileProperties{
            DnsConfig: expandArmProfileDnsConfig(dnsConfig),
            Endpoints: expandArmProfileEndpoint(endpoints),
            MonitorConfig: expandArmProfileMonitorConfig(monitorConfig),
            ProfileStatus: utils.String(profileStatus),
            TrafficRoutingMethod: utils.String(trafficRoutingMethod),
        },
        Tags: tags.Expand(t),
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroup, profileName, parameters); err != nil {
        return fmt.Errorf("Error creating Profile (Profile Name %q / Resource Group %q): %+v", profileName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, profileName)
    if err != nil {
        return fmt.Errorf("Error retrieving Profile (Profile Name %q / Resource Group %q): %+v", profileName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Profile (Profile Name %q / Resource Group %q) ID", profileName, resourceGroup)
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
    profileName := id.Path["trafficmanagerprofiles"]

    resp, err := client.Get(ctx, resourceGroup, profileName)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Profile %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Profile (Profile Name %q / Resource Group %q): %+v", profileName, resourceGroup, err)
    }


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
        if err := d.Set("monitor_config", flattenArmProfileMonitorConfig(profileProperties.MonitorConfig)); err != nil {
            return fmt.Errorf("Error setting `monitor_config`: %+v", err)
        }
        d.Set("profile_status", profileProperties.ProfileStatus)
        d.Set("traffic_routing_method", profileProperties.TrafficRoutingMethod)
    }
    d.Set("profile_name", profileName)
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}

func resourceArmProfileUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).profilesClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    dnsConfig := d.Get("dns_config").([]interface{})
    endpoints := d.Get("endpoints").([]interface{})
    monitorConfig := d.Get("monitor_config").([]interface{})
    profileName := d.Get("profile_name").(string)
    profileStatus := d.Get("profile_status").(string)
    trafficRoutingMethod := d.Get("traffic_routing_method").(string)
    t := d.Get("tags").(map[string]interface{})

    parameters := trafficmanager.Profile{
        Location: utils.String(location),
        ProfileProperties: &trafficmanager.ProfileProperties{
            DnsConfig: expandArmProfileDnsConfig(dnsConfig),
            Endpoints: expandArmProfileEndpoint(endpoints),
            MonitorConfig: expandArmProfileMonitorConfig(monitorConfig),
            ProfileStatus: utils.String(profileStatus),
            TrafficRoutingMethod: utils.String(trafficRoutingMethod),
        },
        Tags: tags.Expand(t),
    }


    if _, err := client.Update(ctx, resourceGroup, profileName, parameters); err != nil {
        return fmt.Errorf("Error updating Profile (Profile Name %q / Resource Group %q): %+v", profileName, resourceGroup, err)
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
    profileName := id.Path["trafficmanagerprofiles"]

    if _, err := client.Delete(ctx, resourceGroup, profileName); err != nil {
        return fmt.Errorf("Error deleting Profile (Profile Name %q / Resource Group %q): %+v", profileName, resourceGroup, err)
    }

    return nil
}

func expandArmProfileDnsConfig(input []interface{}) *trafficmanager.DnsConfig {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    relativeName := v["relative_name"].(string)
    fqdn := v["fqdn"].(string)
    ttl := v["ttl"].(int)

    result := trafficmanager.DnsConfig{
        Fqdn: utils.String(fqdn),
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

        result := trafficmanager.Endpoint{
            ID: utils.String(id),
            Name: utils.String(name),
            EndpointProperties: &trafficmanager.EndpointProperties{
                EndpointLocation: utils.String(endpointLocation),
                EndpointMonitorStatus: utils.String(endpointMonitorStatus),
                EndpointStatus: utils.String(endpointStatus),
                GeoMapping: utils.ExpandStringSlice(geoMapping),
                MinChildEndpoints: utils.Int64(int64(minChildEndpoints)),
                Priority: utils.Int64(int64(priority)),
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

    result := trafficmanager.MonitorConfig{
        Path: utils.String(path),
        Port: utils.Int64(int64(port)),
        ProfileMonitorStatus: utils.String(profileMonitorStatus),
        Protocol: utils.String(protocol),
    }
    return &result
}


func flattenArmProfileDnsConfig(input *trafficmanager.DnsConfig) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if fqdn := input.Fqdn; fqdn != nil {
        result["fqdn"] = *fqdn
    }
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
            if endpointLocation := endpointProperties.EndpointLocation; endpointLocation != nil {
                v["endpoint_location"] = *endpointLocation
            }
            if endpointMonitorStatus := endpointProperties.EndpointMonitorStatus; endpointMonitorStatus != nil {
                v["endpoint_monitor_status"] = *endpointMonitorStatus
            }
            if endpointStatus := endpointProperties.EndpointStatus; endpointStatus != nil {
                v["endpoint_status"] = *endpointStatus
            }
            v["geo_mapping"] = utils.FlattenStringSlice(endpointProperties.GeoMapping)
            if minChildEndpoints := endpointProperties.MinChildEndpoints; minChildEndpoints != nil {
                v["min_child_endpoints"] = int(*minChildEndpoints)
            }
            if priority := endpointProperties.Priority; priority != nil {
                v["priority"] = int(*priority)
            }
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

    if path := input.Path; path != nil {
        result["path"] = *path
    }
    if port := input.Port; port != nil {
        result["port"] = int(*port)
    }
    if profileMonitorStatus := input.ProfileMonitorStatus; profileMonitorStatus != nil {
        result["profile_monitor_status"] = *profileMonitorStatus
    }
    if protocol := input.Protocol; protocol != nil {
        result["protocol"] = *protocol
    }

    return []interface{}{result}
}
