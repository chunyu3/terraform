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
            "lab_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

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

            "external_subnets": {
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

            "location": azure.SchemaLocation(),

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

            "tags": tags.Schema(),

            "unique_identifier": {
                Type: schema.TypeString,
                Optional: true,
            },

            "created_date": {
                Type: schema.TypeString,
                Computed: true,
            },

            "id": {
                Type: schema.TypeString,
                Computed: true,
            },

            "name": {
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
        },
    }
}

func resourceArmVirtualNetworkCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).virtualNetworksClient
    ctx, cancel := timeouts.ForCreate(meta.(*ArmClient).StopContext, d)
    defer cancel()

    resourceGroupName := d.Get("resource_group").(string)
    name := d.Get("lab_name").(string)
    name := d.Get("name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroupName, name, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Virtual Network (Name %q / Lab Name %q / Resource Group %q): %+v", name, name, resourceGroupName, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_virtual_network", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    allowedSubnets := d.Get("allowed_subnets").([]interface{})
    description := d.Get("description").(string)
    externalProviderResourceID := d.Get("external_provider_resource_id").(string)
    externalSubnets := d.Get("external_subnets").([]interface{})
    subnetOverrides := d.Get("subnet_overrides").([]interface{})
    uniqueIdentifier := d.Get("unique_identifier").(string)
    tags := d.Get("tags").(map[string]interface{})

    virtualNetwork := devtestlab.VirtualNetworkFragment{
        Location: utils.String(location),
        VirtualNetworkPropertiesFragment: &devtestlab.VirtualNetworkPropertiesFragment{
            AllowedSubnets: expandArmVirtualNetworkSubnetFragment(allowedSubnets),
            Description: utils.String(description),
            ExternalProviderResourceID: utils.String(externalProviderResourceID),
            ExternalSubnets: expandArmVirtualNetworkExternalSubnetFragment(externalSubnets),
            SubnetOverrides: expandArmVirtualNetworkSubnetOverrideFragment(subnetOverrides),
            UniqueIdentifier: utils.String(uniqueIdentifier),
        },
        Tags: tags.Expand(tags),
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroupName, name, name, virtualNetwork)
    if err != nil {
        return fmt.Errorf("Error creating Virtual Network (Name %q / Lab Name %q / Resource Group %q): %+v", name, name, resourceGroupName, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Virtual Network (Name %q / Lab Name %q / Resource Group %q): %+v", name, name, resourceGroupName, err)
    }


    resp, err := client.Get(ctx, resourceGroupName, name, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Virtual Network (Name %q / Lab Name %q / Resource Group %q): %+v", name, name, resourceGroupName, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Virtual Network (Name %q / Lab Name %q / Resource Group %q) ID", name, name, resourceGroupName)
    }
    d.SetId(*resp.ID)

    return resourceArmVirtualNetworkRead(d, meta)
}

func resourceArmVirtualNetworkRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).virtualNetworksClient
    ctx, cancel := timeouts.ForRead(meta.(*ArmClient).StopContext, d)
    defer cancel()

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroupName := id.ResourceGroup
    name := id.Path["labs"]
    name := id.Path["virtualnetworks"]

    resp, err := client.Get(ctx, resourceGroupName, name, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Virtual Network %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Virtual Network (Name %q / Lab Name %q / Resource Group %q): %+v", name, name, resourceGroupName, err)
    }


    d.Set("resource_group", resourceGroupName)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if virtualNetworkPropertiesFragment := resp.VirtualNetworkPropertiesFragment; virtualNetworkPropertiesFragment != nil {
        if err := d.Set("allowed_subnets", flattenArmVirtualNetworkSubnetFragment(virtualNetworkPropertiesFragment.AllowedSubnets)); err != nil {
            return fmt.Errorf("Error setting `allowed_subnets`: %+v", err)
        }
        d.Set("created_date", (virtualNetworkPropertiesFragment.CreatedDate).String())
        d.Set("description", virtualNetworkPropertiesFragment.Description)
        d.Set("external_provider_resource_id", virtualNetworkPropertiesFragment.ExternalProviderResourceID)
        if err := d.Set("external_subnets", flattenArmVirtualNetworkExternalSubnetFragment(virtualNetworkPropertiesFragment.ExternalSubnets)); err != nil {
            return fmt.Errorf("Error setting `external_subnets`: %+v", err)
        }
        d.Set("provisioning_state", virtualNetworkPropertiesFragment.ProvisioningState)
        if err := d.Set("subnet_overrides", flattenArmVirtualNetworkSubnetOverrideFragment(virtualNetworkPropertiesFragment.SubnetOverrides)); err != nil {
            return fmt.Errorf("Error setting `subnet_overrides`: %+v", err)
        }
        d.Set("unique_identifier", virtualNetworkPropertiesFragment.UniqueIdentifier)
    }
    d.Set("id", resp.ID)
    d.Set("lab_name", name)
    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}

func resourceArmVirtualNetworkUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).virtualNetworksClient
    ctx, cancel := timeouts.ForUpdate(meta.(*ArmClient).StopContext, d)
    defer cancel()

      resourceGroupName := d.Get("resource_group").(string)
    location := azure.NormalizeLocation(d.Get("location").(string))
    allowedSubnets := d.Get("allowed_subnets").([]interface{})
    description := d.Get("description").(string)
    externalProviderResourceID := d.Get("external_provider_resource_id").(string)
    externalSubnets := d.Get("external_subnets").([]interface{})
    name := d.Get("lab_name").(string)
    name := d.Get("name").(string)
    subnetOverrides := d.Get("subnet_overrides").([]interface{})
    uniqueIdentifier := d.Get("unique_identifier").(string)
    tags := d.Get("tags").(map[string]interface{})

    virtualNetwork := devtestlab.VirtualNetworkFragment{
        Location: utils.String(location),
        VirtualNetworkPropertiesFragment: &devtestlab.VirtualNetworkPropertiesFragment{
            AllowedSubnets: expandArmVirtualNetworkSubnetFragment(allowedSubnets),
            Description: utils.String(description),
            ExternalProviderResourceID: utils.String(externalProviderResourceID),
            ExternalSubnets: expandArmVirtualNetworkExternalSubnetFragment(externalSubnets),
            SubnetOverrides: expandArmVirtualNetworkSubnetOverrideFragment(subnetOverrides),
            UniqueIdentifier: utils.String(uniqueIdentifier),
        },
        Tags: tags.Expand(tags),
    }


    if _, err := client.Update(ctx, resourceGroupName, name, name, virtualNetwork); err != nil {
        return fmt.Errorf("Error updating Virtual Network (Name %q / Lab Name %q / Resource Group %q): %+v", name, name, resourceGroupName, err)
    }

    return resourceArmVirtualNetworkRead(d, meta)
}

func resourceArmVirtualNetworkDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).virtualNetworksClient
    ctx, cancel := timeouts.ForDelete(meta.(*ArmClient).StopContext, d)
    defer cancel()


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroupName := id.ResourceGroup
    name := id.Path["labs"]
    name := id.Path["virtualnetworks"]

    future, err := client.Delete(ctx, resourceGroupName, name, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Virtual Network (Name %q / Lab Name %q / Resource Group %q): %+v", name, name, resourceGroupName, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Virtual Network (Name %q / Lab Name %q / Resource Group %q): %+v", name, name, resourceGroupName, err)
        }
    }

    return nil
}

func expandArmVirtualNetworkSubnetFragment(input []interface{}) *[]devtestlab.SubnetFragment {
    results := make([]devtestlab.SubnetFragment, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        resourceID := v["resource_id"].(string)
        labSubnetName := v["lab_subnet_name"].(string)
        allowPublicIP := v["allow_public_ip"].(string)

        result := devtestlab.SubnetFragment{
            AllowPublicIP: devtestlab.UsagePermissionType(allowPublicIP),
            LabSubnetName: utils.String(labSubnetName),
            ResourceID: utils.String(resourceID),
        }

        results = append(results, result)
    }
    return &results
}

func expandArmVirtualNetworkExternalSubnetFragment(input []interface{}) *[]devtestlab.ExternalSubnetFragment {
    results := make([]devtestlab.ExternalSubnetFragment, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        iD := v["id"].(string)
        name := v["name"].(string)

        result := devtestlab.ExternalSubnetFragment{
            ID: utils.String(iD),
            Name: utils.String(name),
        }

        results = append(results, result)
    }
    return &results
}

func expandArmVirtualNetworkSubnetOverrideFragment(input []interface{}) *[]devtestlab.SubnetOverrideFragment {
    results := make([]devtestlab.SubnetOverrideFragment, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        resourceID := v["resource_id"].(string)
        labSubnetName := v["lab_subnet_name"].(string)
        useInVMCreationPermission := v["use_in_vm_creation_permission"].(string)
        usePublicIPAddressPermission := v["use_public_ip_address_permission"].(string)
        sharedPublicIPAddressConfiguration := v["shared_public_ip_address_configuration"].([]interface{})
        virtualNetworkPoolName := v["virtual_network_pool_name"].(string)

        result := devtestlab.SubnetOverrideFragment{
            LabSubnetName: utils.String(labSubnetName),
            ResourceID: utils.String(resourceID),
            SharedPublicIPAddressConfiguration: expandArmVirtualNetworkSubnetSharedPublicIpAddressConfigurationFragment(sharedPublicIPAddressConfiguration),
            UseInVMCreationPermission: devtestlab.UsagePermissionType(useInVMCreationPermission),
            UsePublicIPAddressPermission: devtestlab.UsagePermissionType(usePublicIPAddressPermission),
            VirtualNetworkPoolName: utils.String(virtualNetworkPoolName),
        }

        results = append(results, result)
    }
    return &results
}

func expandArmVirtualNetworkSubnetSharedPublicIpAddressConfigurationFragment(input []interface{}) *devtestlab.SubnetSharedPublicIpAddressConfigurationFragment {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    allowedPorts := v["allowed_ports"].([]interface{})

    result := devtestlab.SubnetSharedPublicIpAddressConfigurationFragment{
        AllowedPorts: expandArmVirtualNetworkPortFragment(allowedPorts),
    }
    return &result
}

func expandArmVirtualNetworkPortFragment(input []interface{}) *[]devtestlab.PortFragment {
    results := make([]devtestlab.PortFragment, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        transportProtocol := v["transport_protocol"].(string)
        backendPort := v["backend_port"].(int)

        result := devtestlab.PortFragment{
            BackendPort: utils.Int32(int32(backendPort)),
            TransportProtocol: devtestlab.TransportProtocol(transportProtocol),
        }

        results = append(results, result)
    }
    return &results
}


func flattenArmVirtualNetworkSubnetFragment(input *[]devtestlab.SubnetFragment) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        v["allow_public_ip"] = string(item.AllowPublicIP)
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

func flattenArmVirtualNetworkExternalSubnetFragment(input *[]devtestlab.ExternalSubnetFragment) []interface{} {
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

func flattenArmVirtualNetworkSubnetOverrideFragment(input *[]devtestlab.SubnetOverrideFragment) []interface{} {
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
        v["shared_public_ip_address_configuration"] = flattenArmVirtualNetworkSubnetSharedPublicIpAddressConfigurationFragment(item.SharedPublicIPAddressConfiguration)
        v["use_in_vm_creation_permission"] = string(item.UseInVMCreationPermission)
        v["use_public_ip_address_permission"] = string(item.UsePublicIPAddressPermission)
        if virtualNetworkPoolName := item.VirtualNetworkPoolName; virtualNetworkPoolName != nil {
            v["virtual_network_pool_name"] = *virtualNetworkPoolName
        }

        results = append(results, v)
    }

    return results
}

func flattenArmVirtualNetworkSubnetSharedPublicIpAddressConfigurationFragment(input *devtestlab.SubnetSharedPublicIpAddressConfigurationFragment) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    result["allowed_ports"] = flattenArmVirtualNetworkPortFragment(input.AllowedPorts)

    return []interface{}{result}
}

func flattenArmVirtualNetworkPortFragment(input *[]devtestlab.PortFragment) []interface{} {
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
