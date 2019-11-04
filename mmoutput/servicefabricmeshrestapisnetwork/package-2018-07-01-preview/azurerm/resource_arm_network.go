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



func resourceArmNetwork() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmNetworkCreateUpdate,
        Read: resourceArmNetworkRead,
        Update: resourceArmNetworkCreateUpdate,
        Delete: resourceArmNetworkDelete,

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

            "address_prefix": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "network_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "description": {
                Type: schema.TypeString,
                Optional: true,
            },

            "ingress_config": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "layer4": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "application_name": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                    "endpoint_name": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                    "name": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                    "public_port": {
                                        Type: schema.TypeInt,
                                        Optional: true,
                                    },
                                    "service_name": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                },
                            },
                        },
                        "qos_level": {
                            Type: schema.TypeString,
                            Optional: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(servicefabricmeshrestapis.Bronze),
                            }, false),
                            Default: string(servicefabricmeshrestapis.Bronze),
                        },
                    },
                },
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

func resourceArmNetworkCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).networkClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    networkName := d.Get("network_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, networkName)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Network (Network Name %q / Resource Group %q): %+v", networkName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_network", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    addressPrefix := d.Get("address_prefix").(string)
    description := d.Get("description").(string)
    ingressConfig := d.Get("ingress_config").([]interface{})
    t := d.Get("tags").(map[string]interface{})

    networkResourceDescription := servicefabricmeshrestapis.NetworkResourceDescription{
        Location: utils.String(location),
        NetworkResourceProperties: &servicefabricmeshrestapis.NetworkResourceProperties{
            AddressPrefix: utils.String(addressPrefix),
            Description: utils.String(description),
            IngressConfig: expandArmNetworkIngressConfig(ingressConfig),
        },
        Tags: tags.Expand(t),
    }


    if _, err := client.Create(ctx, resourceGroup, networkName, networkResourceDescription); err != nil {
        return fmt.Errorf("Error creating Network (Network Name %q / Resource Group %q): %+v", networkName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, networkName)
    if err != nil {
        return fmt.Errorf("Error retrieving Network (Network Name %q / Resource Group %q): %+v", networkName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Network (Network Name %q / Resource Group %q) ID", networkName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmNetworkRead(d, meta)
}

func resourceArmNetworkRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).networkClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    networkName := id.Path["networks"]

    resp, err := client.Get(ctx, resourceGroup, networkName)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Network %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Network (Network Name %q / Resource Group %q): %+v", networkName, resourceGroup, err)
    }


    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if networkResourceProperties := resp.NetworkResourceProperties; networkResourceProperties != nil {
        d.Set("address_prefix", networkResourceProperties.AddressPrefix)
        d.Set("description", networkResourceProperties.Description)
        if err := d.Set("ingress_config", flattenArmNetworkIngressConfig(networkResourceProperties.IngressConfig)); err != nil {
            return fmt.Errorf("Error setting `ingress_config`: %+v", err)
        }
        d.Set("provisioning_state", networkResourceProperties.ProvisioningState)
    }
    d.Set("network_name", networkName)
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}


func resourceArmNetworkDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).networkClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    networkName := id.Path["networks"]

    if _, err := client.Delete(ctx, resourceGroup, networkName); err != nil {
        return fmt.Errorf("Error deleting Network (Network Name %q / Resource Group %q): %+v", networkName, resourceGroup, err)
    }

    return nil
}

func expandArmNetworkIngressConfig(input []interface{}) *servicefabricmeshrestapis.IngressConfig {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    qosLevel := v["qos_level"].(string)
    layer4 := v["layer4"].([]interface{})

    result := servicefabricmeshrestapis.IngressConfig{
        Layer4: expandArmNetworkLayer4IngressConfig(layer4),
        QosLevel: servicefabricmeshrestapis.IngressQoSLevel(qosLevel),
    }
    return &result
}

func expandArmNetworkLayer4IngressConfig(input []interface{}) *[]servicefabricmeshrestapis.Layer4IngressConfig {
    results := make([]servicefabricmeshrestapis.Layer4IngressConfig, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        name := v["name"].(string)
        publicPort := v["public_port"].(int)
        applicationName := v["application_name"].(string)
        serviceName := v["service_name"].(string)
        endpointName := v["endpoint_name"].(string)

        result := servicefabricmeshrestapis.Layer4IngressConfig{
            ApplicationName: utils.String(applicationName),
            EndpointName: utils.String(endpointName),
            Name: utils.String(name),
            PublicPort: utils.Int(publicPort),
            ServiceName: utils.String(serviceName),
        }

        results = append(results, result)
    }
    return &results
}


func flattenArmNetworkIngressConfig(input *servicefabricmeshrestapis.IngressConfig) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    result["layer4"] = flattenArmNetworkLayer4IngressConfig(input.Layer4)
    result["qos_level"] = string(input.QosLevel)

    return []interface{}{result}
}

func flattenArmNetworkLayer4IngressConfig(input *[]servicefabricmeshrestapis.Layer4IngressConfig) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        if name := item.Name; name != nil {
            v["name"] = *name
        }
        if applicationName := item.ApplicationName; applicationName != nil {
            v["application_name"] = *applicationName
        }
        if endpointName := item.EndpointName; endpointName != nil {
            v["endpoint_name"] = *endpointName
        }
        if publicPort := item.PublicPort; publicPort != nil {
            v["public_port"] = *publicPort
        }
        if serviceName := item.ServiceName; serviceName != nil {
            v["service_name"] = *serviceName
        }

        results = append(results, v)
    }

    return results
}
