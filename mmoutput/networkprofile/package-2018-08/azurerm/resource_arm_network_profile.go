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



func resourceArmNetworkProfile() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmNetworkProfileCreateUpdate,
        Read: resourceArmNetworkProfileRead,
        Update: resourceArmNetworkProfileCreateUpdate,
        Delete: resourceArmNetworkProfileDelete,

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

            "network_profile_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "container_network_interface_configurations": {
                Type: schema.TypeList,
                Optional: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "container_network_interfaces": {
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
                    },
                },
            },

            "container_network_interfaces": {
                Type: schema.TypeList,
                Optional: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "container": {
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
                        "container_network_interface_configuration": {
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

            "etag": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
            },

            "provisioning_state": {
                Type: schema.TypeString,
                Computed: true,
            },

            "resource_guid": {
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

func resourceArmNetworkProfileCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).networkProfilesClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    networkProfileName := d.Get("network_profile_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, networkProfileName)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Network Profile (Network Profile Name %q / Resource Group %q): %+v", networkProfileName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_network_profile", *existing.ID)
        }
    }

    id := d.Get("id").(string)
    location := azure.NormalizeLocation(d.Get("location").(string))
    containerNetworkInterfaceConfigurations := d.Get("container_network_interface_configurations").([]interface{})
    containerNetworkInterfaces := d.Get("container_network_interfaces").([]interface{})
    etag := d.Get("etag").(string)
    t := d.Get("tags").(map[string]interface{})

    parameters := network.Profile{
        Etag: utils.String(etag),
        ID: utils.String(id),
        Location: utils.String(location),
        ProfilePropertiesFormat: &network.ProfilePropertiesFormat{
            ContainerNetworkInterfaceConfigurations: expandArmNetworkProfileContainerNetworkInterfaceConfiguration(containerNetworkInterfaceConfigurations),
            ContainerNetworkInterfaces: expandArmNetworkProfileContainerNetworkInterface(containerNetworkInterfaces),
        },
        Tags: tags.Expand(t),
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroup, networkProfileName, parameters); err != nil {
        return fmt.Errorf("Error creating Network Profile (Network Profile Name %q / Resource Group %q): %+v", networkProfileName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, networkProfileName)
    if err != nil {
        return fmt.Errorf("Error retrieving Network Profile (Network Profile Name %q / Resource Group %q): %+v", networkProfileName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Network Profile (Network Profile Name %q / Resource Group %q) ID", networkProfileName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmNetworkProfileRead(d, meta)
}

func resourceArmNetworkProfileRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).networkProfilesClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    networkProfileName := id.Path["networkProfiles"]

    resp, err := client.Get(ctx, resourceGroup, networkProfileName)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Network Profile %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Network Profile (Network Profile Name %q / Resource Group %q): %+v", networkProfileName, resourceGroup, err)
    }


    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if profilePropertiesFormat := resp.ProfilePropertiesFormat; profilePropertiesFormat != nil {
        if err := d.Set("container_network_interface_configurations", flattenArmNetworkProfileContainerNetworkInterfaceConfiguration(profilePropertiesFormat.ContainerNetworkInterfaceConfigurations)); err != nil {
            return fmt.Errorf("Error setting `container_network_interface_configurations`: %+v", err)
        }
        if err := d.Set("container_network_interfaces", flattenArmNetworkProfileContainerNetworkInterface(profilePropertiesFormat.ContainerNetworkInterfaces)); err != nil {
            return fmt.Errorf("Error setting `container_network_interfaces`: %+v", err)
        }
        d.Set("provisioning_state", profilePropertiesFormat.ProvisioningState)
        d.Set("resource_guid", profilePropertiesFormat.ResourceGuid)
    }
    d.Set("etag", resp.Etag)
    d.Set("network_profile_name", networkProfileName)
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}


func resourceArmNetworkProfileDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).networkProfilesClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    networkProfileName := id.Path["networkProfiles"]

    future, err := client.Delete(ctx, resourceGroup, networkProfileName)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Network Profile (Network Profile Name %q / Resource Group %q): %+v", networkProfileName, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Network Profile (Network Profile Name %q / Resource Group %q): %+v", networkProfileName, resourceGroup, err)
        }
    }

    return nil
}

func expandArmNetworkProfileContainerNetworkInterfaceConfiguration(input []interface{}) *[]network.ContainerNetworkInterfaceConfiguration {
    results := make([]network.ContainerNetworkInterfaceConfiguration, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        id := v["id"].(string)
        ipConfigurations := v["ip_configurations"].([]interface{})
        containerNetworkInterfaces := v["container_network_interfaces"].([]interface{})
        name := v["name"].(string)
        etag := v["etag"].(string)

        result := network.ContainerNetworkInterfaceConfiguration{
            Etag: utils.String(etag),
            ID: utils.String(id),
            Name: utils.String(name),
            ContainerNetworkInterfaceConfigurationPropertiesFormat: &network.ContainerNetworkInterfaceConfigurationPropertiesFormat{
                ContainerNetworkInterfaces: expandArmNetworkProfileContainerNetworkInterface(containerNetworkInterfaces),
                IpConfigurations: expandArmNetworkProfileIPConfigurationProfile(ipConfigurations),
            },
        }

        results = append(results, result)
    }
    return &results
}

func expandArmNetworkProfileContainerNetworkInterface(input []interface{}) *[]network.ContainerNetworkInterface {
    results := make([]network.ContainerNetworkInterface, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        id := v["id"].(string)
        containerNetworkInterfaceConfiguration := v["container_network_interface_configuration"].([]interface{})
        container := v["container"].([]interface{})
        ipConfigurations := v["ip_configurations"].([]interface{})
        name := v["name"].(string)
        etag := v["etag"].(string)

        result := network.ContainerNetworkInterface{
            Etag: utils.String(etag),
            ID: utils.String(id),
            Name: utils.String(name),
            ContainerNetworkInterfacePropertiesFormat: &network.ContainerNetworkInterfacePropertiesFormat{
                Container: expandArmNetworkProfileContainer(container),
                ContainerNetworkInterfaceConfiguration: expandArmNetworkProfileContainerNetworkInterfaceConfiguration(containerNetworkInterfaceConfiguration),
                IpConfigurations: expandArmNetworkProfileContainerNetworkInterfaceIpConfiguration(ipConfigurations),
            },
        }

        results = append(results, result)
    }
    return &results
}

func expandArmNetworkProfileContainerNetworkInterface(input []interface{}) *[]network.ContainerNetworkInterface {
    results := make([]network.ContainerNetworkInterface, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        id := v["id"].(string)
        name := v["name"].(string)
        etag := v["etag"].(string)

        result := network.ContainerNetworkInterface{
            Etag: utils.String(etag),
            ID: utils.String(id),
            Name: utils.String(name),
        }

        results = append(results, result)
    }
    return &results
}

func expandArmNetworkProfileIPConfigurationProfile(input []interface{}) *[]network.IPConfigurationProfile {
    results := make([]network.IPConfigurationProfile, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        id := v["id"].(string)
        name := v["name"].(string)
        etag := v["etag"].(string)

        result := network.IPConfigurationProfile{
            Etag: utils.String(etag),
            ID: utils.String(id),
            Name: utils.String(name),
        }

        results = append(results, result)
    }
    return &results
}

func expandArmNetworkProfileContainer(input []interface{}) *network.Container {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    id := v["id"].(string)

    result := network.Container{
        ID: utils.String(id),
    }
    return &result
}

func expandArmNetworkProfileContainerNetworkInterfaceConfiguration(input []interface{}) *network.ContainerNetworkInterfaceConfiguration {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    id := v["id"].(string)
    name := v["name"].(string)
    etag := v["etag"].(string)

    result := network.ContainerNetworkInterfaceConfiguration{
        Etag: utils.String(etag),
        ID: utils.String(id),
        Name: utils.String(name),
    }
    return &result
}

func expandArmNetworkProfileContainerNetworkInterfaceIpConfiguration(input []interface{}) *[]network.ContainerNetworkInterfaceIpConfiguration {
    results := make([]network.ContainerNetworkInterfaceIpConfiguration, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        name := v["name"].(string)
        etag := v["etag"].(string)

        result := network.ContainerNetworkInterfaceIpConfiguration{
            Etag: utils.String(etag),
            Name: utils.String(name),
        }

        results = append(results, result)
    }
    return &results
}


func flattenArmNetworkProfileContainerNetworkInterfaceConfiguration(input *[]network.ContainerNetworkInterfaceConfiguration) []interface{} {
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
        if containerNetworkInterfaceConfigurationPropertiesFormat := item.ContainerNetworkInterfaceConfigurationPropertiesFormat; containerNetworkInterfaceConfigurationPropertiesFormat != nil {
            v["container_network_interfaces"] = flattenArmNetworkProfileContainerNetworkInterface(containerNetworkInterfaceConfigurationPropertiesFormat.ContainerNetworkInterfaces)
            v["ip_configurations"] = flattenArmNetworkProfileIPConfigurationProfile(containerNetworkInterfaceConfigurationPropertiesFormat.IpConfigurations)
        }
        if etag := item.Etag; etag != nil {
            v["etag"] = *etag
        }

        results = append(results, v)
    }

    return results
}

func flattenArmNetworkProfileContainerNetworkInterface(input *[]network.ContainerNetworkInterface) []interface{} {
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
        if containerNetworkInterfacePropertiesFormat := item.ContainerNetworkInterfacePropertiesFormat; containerNetworkInterfacePropertiesFormat != nil {
            v["container"] = flattenArmNetworkProfileContainer(containerNetworkInterfacePropertiesFormat.Container)
            v["container_network_interface_configuration"] = flattenArmNetworkProfileContainerNetworkInterfaceConfiguration(containerNetworkInterfacePropertiesFormat.ContainerNetworkInterfaceConfiguration)
            v["ip_configurations"] = flattenArmNetworkProfileContainerNetworkInterfaceIpConfiguration(containerNetworkInterfacePropertiesFormat.IpConfigurations)
        }
        if etag := item.Etag; etag != nil {
            v["etag"] = *etag
        }

        results = append(results, v)
    }

    return results
}

func flattenArmNetworkProfileContainerNetworkInterface(input *[]network.ContainerNetworkInterface) []interface{} {
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

func flattenArmNetworkProfileIPConfigurationProfile(input *[]network.IPConfigurationProfile) []interface{} {
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

func flattenArmNetworkProfileContainer(input *network.Container) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if id := input.ID; id != nil {
        result["id"] = *id
    }

    return []interface{}{result}
}

func flattenArmNetworkProfileContainerNetworkInterfaceConfiguration(input *network.ContainerNetworkInterfaceConfiguration) []interface{} {
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

func flattenArmNetworkProfileContainerNetworkInterfaceIpConfiguration(input *[]network.ContainerNetworkInterfaceIpConfiguration) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

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
