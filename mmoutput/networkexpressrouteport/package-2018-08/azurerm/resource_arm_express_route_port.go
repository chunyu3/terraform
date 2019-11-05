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



func resourceArmExpressRoutePort() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmExpressRoutePortCreateUpdate,
        Read: resourceArmExpressRoutePortRead,
        Update: resourceArmExpressRoutePortCreateUpdate,
        Delete: resourceArmExpressRoutePortDelete,

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

            "express_route_port_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "bandwidth_in_gbps": {
                Type: schema.TypeInt,
                Optional: true,
            },

            "encapsulation": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(network.Dot1Q),
                    string(network.QinQ),
                }, false),
                Default: string(network.Dot1Q),
            },

            "links": {
                Type: schema.TypeList,
                Optional: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "admin_state": {
                            Type: schema.TypeString,
                            Optional: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(network.Enabled),
                                string(network.Disabled),
                            }, false),
                            Default: string(network.Enabled),
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

            "peering_location": {
                Type: schema.TypeString,
                Optional: true,
            },

            "resource_guid": {
                Type: schema.TypeString,
                Optional: true,
            },

            "allocation_date": {
                Type: schema.TypeString,
                Computed: true,
            },

            "circuits": {
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

            "etag": {
                Type: schema.TypeString,
                Computed: true,
            },

            "ether_type": {
                Type: schema.TypeString,
                Computed: true,
            },

            "mtu": {
                Type: schema.TypeString,
                Computed: true,
            },

            "provisioned_bandwidth_in_gbps": {
                Type: schema.TypeInt,
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

func resourceArmExpressRoutePortCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).expressRoutePortsClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    expressRoutePortName := d.Get("express_route_port_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, expressRoutePortName)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Express Route Port (Express Route Port Name %q / Resource Group %q): %+v", expressRoutePortName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_express_route_port", *existing.ID)
        }
    }

    id := d.Get("id").(string)
    location := azure.NormalizeLocation(d.Get("location").(string))
    bandwidthInGbps := d.Get("bandwidth_in_gbps").(int)
    encapsulation := d.Get("encapsulation").(string)
    links := d.Get("links").([]interface{})
    peeringLocation := d.Get("peering_location").(string)
    resourceGuid := d.Get("resource_guid").(string)
    t := d.Get("tags").(map[string]interface{})

    parameters := network.ExpressRoutePort{
        ID: utils.String(id),
        Location: utils.String(location),
        ExpressRoutePortPropertiesFormat: &network.ExpressRoutePortPropertiesFormat{
            BandwidthInGbps: utils.Int(bandwidthInGbps),
            Encapsulation: network.ExpressRoutePortsEncapsulation(encapsulation),
            Links: expandArmExpressRoutePortExpressRouteLink(links),
            PeeringLocation: utils.String(peeringLocation),
            ResourceGuid: utils.String(resourceGuid),
        },
        Tags: tags.Expand(t),
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, expressRoutePortName, parameters)
    if err != nil {
        return fmt.Errorf("Error creating Express Route Port (Express Route Port Name %q / Resource Group %q): %+v", expressRoutePortName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Express Route Port (Express Route Port Name %q / Resource Group %q): %+v", expressRoutePortName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, expressRoutePortName)
    if err != nil {
        return fmt.Errorf("Error retrieving Express Route Port (Express Route Port Name %q / Resource Group %q): %+v", expressRoutePortName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Express Route Port (Express Route Port Name %q / Resource Group %q) ID", expressRoutePortName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmExpressRoutePortRead(d, meta)
}

func resourceArmExpressRoutePortRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).expressRoutePortsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    expressRoutePortName := id.Path["ExpressRoutePorts"]

    resp, err := client.Get(ctx, resourceGroup, expressRoutePortName)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Express Route Port %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Express Route Port (Express Route Port Name %q / Resource Group %q): %+v", expressRoutePortName, resourceGroup, err)
    }


    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if expressRoutePortPropertiesFormat := resp.ExpressRoutePortPropertiesFormat; expressRoutePortPropertiesFormat != nil {
        d.Set("allocation_date", expressRoutePortPropertiesFormat.AllocationDate)
        d.Set("bandwidth_in_gbps", expressRoutePortPropertiesFormat.BandwidthInGbps)
        if err := d.Set("circuits", flattenArmExpressRoutePortSubResource(expressRoutePortPropertiesFormat.Circuits)); err != nil {
            return fmt.Errorf("Error setting `circuits`: %+v", err)
        }
        d.Set("encapsulation", string(expressRoutePortPropertiesFormat.Encapsulation))
        d.Set("ether_type", expressRoutePortPropertiesFormat.EtherType)
        if err := d.Set("links", flattenArmExpressRoutePortExpressRouteLink(expressRoutePortPropertiesFormat.Links)); err != nil {
            return fmt.Errorf("Error setting `links`: %+v", err)
        }
        d.Set("mtu", expressRoutePortPropertiesFormat.Mtu)
        d.Set("peering_location", expressRoutePortPropertiesFormat.PeeringLocation)
        d.Set("provisioned_bandwidth_in_gbps", expressRoutePortPropertiesFormat.ProvisionedBandwidthInGbps)
        d.Set("provisioning_state", expressRoutePortPropertiesFormat.ProvisioningState)
        d.Set("resource_guid", expressRoutePortPropertiesFormat.ResourceGuid)
    }
    d.Set("etag", resp.Etag)
    d.Set("express_route_port_name", expressRoutePortName)
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}


func resourceArmExpressRoutePortDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).expressRoutePortsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    expressRoutePortName := id.Path["ExpressRoutePorts"]

    future, err := client.Delete(ctx, resourceGroup, expressRoutePortName)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Express Route Port (Express Route Port Name %q / Resource Group %q): %+v", expressRoutePortName, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Express Route Port (Express Route Port Name %q / Resource Group %q): %+v", expressRoutePortName, resourceGroup, err)
        }
    }

    return nil
}

func expandArmExpressRoutePortExpressRouteLink(input []interface{}) *[]network.ExpressRouteLink {
    results := make([]network.ExpressRouteLink, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        id := v["id"].(string)
        adminState := v["admin_state"].(string)
        name := v["name"].(string)

        result := network.ExpressRouteLink{
            ID: utils.String(id),
            Name: utils.String(name),
            ExpressRouteLinkPropertiesFormat: &network.ExpressRouteLinkPropertiesFormat{
                AdminState: network.ExpressRouteLinkAdminState(adminState),
            },
        }

        results = append(results, result)
    }
    return &results
}


func flattenArmExpressRoutePortSubResource(input *[]network.SubResource) []interface{} {
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

func flattenArmExpressRoutePortExpressRouteLink(input *[]network.ExpressRouteLink) []interface{} {
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
        if expressRouteLinkPropertiesFormat := item.ExpressRouteLinkPropertiesFormat; expressRouteLinkPropertiesFormat != nil {
            v["admin_state"] = string(expressRouteLinkPropertiesFormat.AdminState)
        }

        results = append(results, v)
    }

    return results
}
