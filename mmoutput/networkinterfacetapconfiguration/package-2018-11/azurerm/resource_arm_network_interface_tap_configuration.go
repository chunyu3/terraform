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



func resourceArmNetworkInterfaceTapConfiguration() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmNetworkInterfaceTapConfigurationCreateUpdate,
        Read: resourceArmNetworkInterfaceTapConfigurationRead,
        Update: resourceArmNetworkInterfaceTapConfigurationCreateUpdate,
        Delete: resourceArmNetworkInterfaceTapConfigurationDelete,

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

            "network_interface_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "tap_configuration_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "etag": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
            },

            "virtual_network_tap": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "destination_load_balancer_front_end_ipconfiguration": {
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
                                    "name": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                    "zones": {
                                        Type: schema.TypeList,
                                        Optional: true,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                        },
                                    },
                                },
                            },
                        },
                        "destination_network_interface_ipconfiguration": {
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
                                    "name": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                },
                            },
                        },
                        "destination_port": {
                            Type: schema.TypeInt,
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
                        "location": azure.SchemaLocation(),
                        "tags": tags.Schema(),
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
        },
    }
}

func resourceArmNetworkInterfaceTapConfigurationCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).networkInterfaceTapConfigurationsClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    networkInterfaceName := d.Get("network_interface_name").(string)
    tapConfigurationName := d.Get("tap_configuration_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, networkInterfaceName, tapConfigurationName)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Network Interface Tap Configuration (Tap Configuration Name %q / Network Interface Name %q / Resource Group %q): %+v", tapConfigurationName, networkInterfaceName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_network_interface_tap_configuration", *existing.ID)
        }
    }

    id := d.Get("id").(string)
    name := d.Get("name").(string)
    etag := d.Get("etag").(string)
    virtualNetworkTap := d.Get("virtual_network_tap").([]interface{})

    tapConfigurationParameters := network.InterfaceTapConfiguration{
        Etag: utils.String(etag),
        ID: utils.String(id),
        Name: utils.String(name),
        InterfaceTapConfigurationPropertiesFormat: &network.InterfaceTapConfigurationPropertiesFormat{
            VirtualNetworkTap: expandArmNetworkInterfaceTapConfigurationVirtualNetworkTap(virtualNetworkTap),
        },
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, networkInterfaceName, tapConfigurationName, tapConfigurationParameters)
    if err != nil {
        return fmt.Errorf("Error creating Network Interface Tap Configuration (Tap Configuration Name %q / Network Interface Name %q / Resource Group %q): %+v", tapConfigurationName, networkInterfaceName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Network Interface Tap Configuration (Tap Configuration Name %q / Network Interface Name %q / Resource Group %q): %+v", tapConfigurationName, networkInterfaceName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, networkInterfaceName, tapConfigurationName)
    if err != nil {
        return fmt.Errorf("Error retrieving Network Interface Tap Configuration (Tap Configuration Name %q / Network Interface Name %q / Resource Group %q): %+v", tapConfigurationName, networkInterfaceName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Network Interface Tap Configuration (Tap Configuration Name %q / Network Interface Name %q / Resource Group %q) ID", tapConfigurationName, networkInterfaceName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmNetworkInterfaceTapConfigurationRead(d, meta)
}

func resourceArmNetworkInterfaceTapConfigurationRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).networkInterfaceTapConfigurationsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    networkInterfaceName := id.Path["networkInterfaces"]
    tapConfigurationName := id.Path["tapConfigurations"]

    resp, err := client.Get(ctx, resourceGroup, networkInterfaceName, tapConfigurationName)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Network Interface Tap Configuration %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Network Interface Tap Configuration (Tap Configuration Name %q / Network Interface Name %q / Resource Group %q): %+v", tapConfigurationName, networkInterfaceName, resourceGroup, err)
    }


    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    d.Set("etag", resp.Etag)
    d.Set("network_interface_name", networkInterfaceName)
    if interfaceTapConfigurationPropertiesFormat := resp.InterfaceTapConfigurationPropertiesFormat; interfaceTapConfigurationPropertiesFormat != nil {
        d.Set("provisioning_state", interfaceTapConfigurationPropertiesFormat.ProvisioningState)
        if err := d.Set("virtual_network_tap", flattenArmNetworkInterfaceTapConfigurationVirtualNetworkTap(interfaceTapConfigurationPropertiesFormat.VirtualNetworkTap)); err != nil {
            return fmt.Errorf("Error setting `virtual_network_tap`: %+v", err)
        }
    }
    d.Set("tap_configuration_name", tapConfigurationName)
    d.Set("type", resp.Type)

    return nil
}


func resourceArmNetworkInterfaceTapConfigurationDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).networkInterfaceTapConfigurationsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    networkInterfaceName := id.Path["networkInterfaces"]
    tapConfigurationName := id.Path["tapConfigurations"]

    future, err := client.Delete(ctx, resourceGroup, networkInterfaceName, tapConfigurationName)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Network Interface Tap Configuration (Tap Configuration Name %q / Network Interface Name %q / Resource Group %q): %+v", tapConfigurationName, networkInterfaceName, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Network Interface Tap Configuration (Tap Configuration Name %q / Network Interface Name %q / Resource Group %q): %+v", tapConfigurationName, networkInterfaceName, resourceGroup, err)
        }
    }

    return nil
}

func expandArmNetworkInterfaceTapConfigurationVirtualNetworkTap(input []interface{}) *network.VirtualNetworkTap {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    id := v["id"].(string)
    location := azure.NormalizeLocation(v["location"].(string))
    t := v["tags"].(map[string]interface{})
    destinationNetworkInterfaceIpconfiguration := v["destination_network_interface_ipconfiguration"].([]interface{})
    destinationLoadBalancerFrontEndIpconfiguration := v["destination_load_balancer_front_end_ipconfiguration"].([]interface{})
    destinationPort := v["destination_port"].(int)
    etag := v["etag"].(string)

    result := network.VirtualNetworkTap{
        Etag: utils.String(etag),
        ID: utils.String(id),
        Location: utils.String(location),
        VirtualNetworkTapPropertiesFormat: &network.VirtualNetworkTapPropertiesFormat{
            DestinationLoadBalancerFrontEndIpconfiguration: expandArmNetworkInterfaceTapConfigurationFrontendIPConfiguration(destinationLoadBalancerFrontEndIpconfiguration),
            DestinationNetworkInterfaceIpconfiguration: expandArmNetworkInterfaceTapConfigurationInterfaceIPConfiguration(destinationNetworkInterfaceIpconfiguration),
            DestinationPort: utils.Int(destinationPort),
        },
        Tags: tags.Expand(t),
    }
    return &result
}

func expandArmNetworkInterfaceTapConfigurationFrontendIPConfiguration(input []interface{}) *network.FrontendIPConfiguration {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    id := v["id"].(string)
    name := v["name"].(string)
    etag := v["etag"].(string)
    zones := v["zones"].([]interface{})

    result := network.FrontendIPConfiguration{
        Etag: utils.String(etag),
        ID: utils.String(id),
        Name: utils.String(name),
        Zones: utils.ExpandStringSlice(zones),
    }
    return &result
}

func expandArmNetworkInterfaceTapConfigurationInterfaceIPConfiguration(input []interface{}) *network.InterfaceIPConfiguration {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    id := v["id"].(string)
    name := v["name"].(string)
    etag := v["etag"].(string)

    result := network.InterfaceIPConfiguration{
        Etag: utils.String(etag),
        ID: utils.String(id),
        Name: utils.String(name),
    }
    return &result
}


func flattenArmNetworkInterfaceTapConfigurationVirtualNetworkTap(input *network.VirtualNetworkTap) []interface{} {
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
    if virtualNetworkTapPropertiesFormat := input.VirtualNetworkTapPropertiesFormat; virtualNetworkTapPropertiesFormat != nil {
        result["destination_load_balancer_front_end_ipconfiguration"] = flattenArmNetworkInterfaceTapConfigurationFrontendIPConfiguration(virtualNetworkTapPropertiesFormat.DestinationLoadBalancerFrontEndIpconfiguration)
        result["destination_network_interface_ipconfiguration"] = flattenArmNetworkInterfaceTapConfigurationInterfaceIPConfiguration(virtualNetworkTapPropertiesFormat.DestinationNetworkInterfaceIpconfiguration)
        if destinationPort := virtualNetworkTapPropertiesFormat.DestinationPort; destinationPort != nil {
            result["destination_port"] = *destinationPort
        }
    }
    if etag := input.Etag; etag != nil {
        result["etag"] = *etag
    }

    return []interface{}{result}
}

func flattenArmNetworkInterfaceTapConfigurationFrontendIPConfiguration(input *network.FrontendIPConfiguration) []interface{} {
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
    if etag := input.Etag; etag != nil {
        result["etag"] = *etag
    }
    result["zones"] = utils.FlattenStringSlice(input.Zones)

    return []interface{}{result}
}

func flattenArmNetworkInterfaceTapConfigurationInterfaceIPConfiguration(input *network.InterfaceIPConfiguration) []interface{} {
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
    if etag := input.Etag; etag != nil {
        result["etag"] = *etag
    }

    return []interface{}{result}
}
