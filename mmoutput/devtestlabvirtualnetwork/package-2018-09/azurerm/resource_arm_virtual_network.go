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



func resourceArmVirtualNetwork() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmVirtualNetworkCreate,
        Read: resourceArmVirtualNetworkRead,
        Update: resourceArmVirtualNetworkUpdate,
        Delete: resourceArmVirtualNetworkDelete,

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

            "location": azure.SchemaLocation(),

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "lab_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "allowed_subnets": {
                Type: schema.TypeList,
                Optional: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "allow_public_ip": {
                            Type: schema.TypeString,
                            Optional: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(devtestlab.Default),
                                string(devtestlab.Deny),
                                string(devtestlab.Allow),
                            }, false),
                            Default: string(devtestlab.Default),
                        },
                        "lab_subnet_name": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "resource_id": {
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

            "external_provider_resource_id": {
                Type: schema.TypeString,
                Optional: true,
            },

            "subnet_overrides": {
                Type: schema.TypeList,
                Optional: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "lab_subnet_name": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "resource_id": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "shared_public_ip_address_configuration": {
                            Type: schema.TypeList,
                            Optional: true,
                            MaxItems: 1,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "allowed_ports": {
                                        Type: schema.TypeList,
                                        Optional: true,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "backend_port": {
                                                    Type: schema.TypeInt,
                                                    Optional: true,
                                                },
                                                "transport_protocol": {
                                                    Type: schema.TypeString,
                                                    Optional: true,
                                                    ValidateFunc: validation.StringInSlice([]string{
                                                        string(devtestlab.Tcp),
                                                        string(devtestlab.Udp),
                                                    }, false),
                                                    Default: string(devtestlab.Tcp),
                                                },
                                            },
                                        },
                                    },
                                },
                            },
                        },
                        "use_in_vm_creation_permission": {
                            Type: schema.TypeString,
                            Optional: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(devtestlab.Default),
                                string(devtestlab.Deny),
                                string(devtestlab.Allow),
                            }, false),
                            Default: string(devtestlab.Default),
                        },
                        "use_public_ip_address_permission": {
                            Type: schema.TypeString,
                            Optional: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(devtestlab.Default),
                                string(devtestlab.Deny),
                                string(devtestlab.Allow),
                            }, false),
                            Default: string(devtestlab.Default),
                        },
                        "virtual_network_pool_name": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },

            "created_date": {
                Type: schema.TypeString,
                Computed: true,
            },

            "external_subnets": {
                Type: schema.TypeList,
                Computed: true,
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

            "provisioning_state": {
                Type: schema.TypeString,
                Computed: true,
            },

            "type": {
                Type: schema.TypeString,
                Computed: true,
            },

            "unique_identifier": {
                Type: schema.TypeString,
                Computed: true,
            },

            "tags": tags.Schema(),
        },
    }
}

func resourceArmVirtualNetworkCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).virtualNetworksClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    labName := d.Get("lab_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, labName, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Virtual Network %q (Lab Name %q / Resource Group %q): %+v", name, labName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_virtual_network", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    allowedSubnets := d.Get("allowed_subnets").([]interface{})
    description := d.Get("description").(string)
    externalProviderResourceId := d.Get("external_provider_resource_id").(string)
    subnetOverrides := d.Get("subnet_overrides").([]interface{})
    t := d.Get("tags").(map[string]interface{})

    virtualNetwork := devtestlab.VirtualNetwork{
        Location: utils.String(location),
        VirtualNetworkProperties: &devtestlab.VirtualNetworkProperties{
            AllowedSubnets: expandArmVirtualNetworkSubnet(allowedSubnets),
            Description: utils.String(description),
            ExternalProviderResourceID: utils.String(externalProviderResourceId),
            SubnetOverrides: expandArmVirtualNetworkSubnetOverride(subnetOverrides),
        },
        Tags: tags.Expand(t),
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, labName, name, virtualNetwork)
    if err != nil {
        return fmt.Errorf("Error creating Virtual Network %q (Lab Name %q / Resource Group %q): %+v", name, labName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Virtual Network %q (Lab Name %q / Resource Group %q): %+v", name, labName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, labName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Virtual Network %q (Lab Name %q / Resource Group %q): %+v", name, labName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Virtual Network %q (Lab Name %q / Resource Group %q) ID", name, labName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmVirtualNetworkRead(d, meta)
}

func resourceArmVirtualNetworkRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).virtualNetworksClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    labName := id.Path["labs"]
    name := id.Path["virtualnetworks"]

    resp, err := client.Get(ctx, resourceGroup, labName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Virtual Network %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Virtual Network %q (Lab Name %q / Resource Group %q): %+v", name, labName, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if virtualNetworkProperties := resp.VirtualNetworkProperties; virtualNetworkProperties != nil {
        if err := d.Set("allowed_subnets", flattenArmVirtualNetworkSubnet(virtualNetworkProperties.AllowedSubnets)); err != nil {
            return fmt.Errorf("Error setting `allowed_subnets`: %+v", err)
        }
        d.Set("created_date", (virtualNetworkProperties.CreatedDate).String())
        d.Set("description", virtualNetworkProperties.Description)
        d.Set("external_provider_resource_id", virtualNetworkProperties.ExternalProviderResourceID)
        if err := d.Set("external_subnets", flattenArmVirtualNetworkExternalSubnet(virtualNetworkProperties.ExternalSubnets)); err != nil {
            return fmt.Errorf("Error setting `external_subnets`: %+v", err)
        }
        d.Set("provisioning_state", virtualNetworkProperties.ProvisioningState)
        if err := d.Set("subnet_overrides", flattenArmVirtualNetworkSubnetOverride(virtualNetworkProperties.SubnetOverrides)); err != nil {
            return fmt.Errorf("Error setting `subnet_overrides`: %+v", err)
        }
        d.Set("unique_identifier", virtualNetworkProperties.UniqueIdentifier)
    }
    d.Set("lab_name", labName)
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}

func resourceArmVirtualNetworkUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).virtualNetworksClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    allowedSubnets := d.Get("allowed_subnets").([]interface{})
    description := d.Get("description").(string)
    externalProviderResourceId := d.Get("external_provider_resource_id").(string)
    labName := d.Get("lab_name").(string)
    subnetOverrides := d.Get("subnet_overrides").([]interface{})
    t := d.Get("tags").(map[string]interface{})

    virtualNetwork := devtestlab.VirtualNetwork{
        Location: utils.String(location),
        VirtualNetworkProperties: &devtestlab.VirtualNetworkProperties{
            AllowedSubnets: expandArmVirtualNetworkSubnet(allowedSubnets),
            Description: utils.String(description),
            ExternalProviderResourceID: utils.String(externalProviderResourceId),
            SubnetOverrides: expandArmVirtualNetworkSubnetOverride(subnetOverrides),
        },
        Tags: tags.Expand(t),
    }


    if _, err := client.Update(ctx, resourceGroup, labName, name, virtualNetwork); err != nil {
        return fmt.Errorf("Error updating Virtual Network %q (Lab Name %q / Resource Group %q): %+v", name, labName, resourceGroup, err)
    }

    return resourceArmVirtualNetworkRead(d, meta)
}

func resourceArmVirtualNetworkDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).virtualNetworksClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    labName := id.Path["labs"]
    name := id.Path["virtualnetworks"]

    future, err := client.Delete(ctx, resourceGroup, labName, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Virtual Network %q (Lab Name %q / Resource Group %q): %+v", name, labName, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Virtual Network %q (Lab Name %q / Resource Group %q): %+v", name, labName, resourceGroup, err)
        }
    }

    return nil
}

func expandArmVirtualNetworkSubnet(input []interface{}) *[]devtestlab.Subnet {
    results := make([]devtestlab.Subnet, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        resourceId := v["resource_id"].(string)
        labSubnetName := v["lab_subnet_name"].(string)
        allowPublicIp := v["allow_public_ip"].(string)

        result := devtestlab.Subnet{
            AllowPublicIp: devtestlab.UsagePermissionType(allowPublicIp),
            LabSubnetName: utils.String(labSubnetName),
            ResourceID: utils.String(resourceId),
        }

        results = append(results, result)
    }
    return &results
}

func expandArmVirtualNetworkSubnetOverride(input []interface{}) *[]devtestlab.SubnetOverride {
    results := make([]devtestlab.SubnetOverride, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        resourceId := v["resource_id"].(string)
        labSubnetName := v["lab_subnet_name"].(string)
        useInVmCreationPermission := v["use_in_vm_creation_permission"].(string)
        usePublicIpAddressPermission := v["use_public_ip_address_permission"].(string)
        sharedPublicIpAddressConfiguration := v["shared_public_ip_address_configuration"].([]interface{})
        virtualNetworkPoolName := v["virtual_network_pool_name"].(string)

        result := devtestlab.SubnetOverride{
            LabSubnetName: utils.String(labSubnetName),
            ResourceID: utils.String(resourceId),
            SharedPublicIpAddressConfiguration: expandArmVirtualNetworkSubnetSharedPublicIpAddressConfiguration(sharedPublicIpAddressConfiguration),
            UseInVmCreationPermission: devtestlab.UsagePermissionType(useInVmCreationPermission),
            UsePublicIpAddressPermission: devtestlab.UsagePermissionType(usePublicIpAddressPermission),
            VirtualNetworkPoolName: utils.String(virtualNetworkPoolName),
        }

        results = append(results, result)
    }
    return &results
}

func expandArmVirtualNetworkSubnetSharedPublicIpAddressConfiguration(input []interface{}) *devtestlab.SubnetSharedPublicIpAddressConfiguration {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    allowedPorts := v["allowed_ports"].([]interface{})

    result := devtestlab.SubnetSharedPublicIpAddressConfiguration{
        AllowedPorts: expandArmVirtualNetworkPort(allowedPorts),
    }
    return &result
}

func expandArmVirtualNetworkPort(input []interface{}) *[]devtestlab.Port {
    results := make([]devtestlab.Port, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        transportProtocol := v["transport_protocol"].(string)
        backendPort := v["backend_port"].(int)

        result := devtestlab.Port{
            BackendPort: utils.Int32(int32(backendPort)),
            TransportProtocol: devtestlab.TransportProtocol(transportProtocol),
        }

        results = append(results, result)
    }
    return &results
}


func flattenArmVirtualNetworkSubnet(input *[]devtestlab.Subnet) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        v["allow_public_ip"] = string(item.AllowPublicIp)
        if labSubnetName := item.LabSubnetName; labSubnetName != nil {
            v["lab_subnet_name"] = *labSubnetName
        }
        if resourceId := item.ResourceID; resourceId != nil {
            v["resource_id"] = *resourceId
        }

        results = append(results, v)
    }

    return results
}

func flattenArmVirtualNetworkExternalSubnet(input *[]devtestlab.ExternalSubnet) []interface{} {
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

func flattenArmVirtualNetworkSubnetOverride(input *[]devtestlab.SubnetOverride) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        if labSubnetName := item.LabSubnetName; labSubnetName != nil {
            v["lab_subnet_name"] = *labSubnetName
        }
        if resourceId := item.ResourceID; resourceId != nil {
            v["resource_id"] = *resourceId
        }
        v["shared_public_ip_address_configuration"] = flattenArmVirtualNetworkSubnetSharedPublicIpAddressConfiguration(item.SharedPublicIpAddressConfiguration)
        v["use_in_vm_creation_permission"] = string(item.UseInVmCreationPermission)
        v["use_public_ip_address_permission"] = string(item.UsePublicIpAddressPermission)
        if virtualNetworkPoolName := item.VirtualNetworkPoolName; virtualNetworkPoolName != nil {
            v["virtual_network_pool_name"] = *virtualNetworkPoolName
        }

        results = append(results, v)
    }

    return results
}

func flattenArmVirtualNetworkSubnetSharedPublicIpAddressConfiguration(input *devtestlab.SubnetSharedPublicIpAddressConfiguration) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    result["allowed_ports"] = flattenArmVirtualNetworkPort(input.AllowedPorts)

    return []interface{}{result}
}

func flattenArmVirtualNetworkPort(input *[]devtestlab.Port) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        if backendPort := item.BackendPort; backendPort != nil {
            v["backend_port"] = int(*backendPort)
        }
        v["transport_protocol"] = string(item.TransportProtocol)

        results = append(results, v)
    }

    return results
}
