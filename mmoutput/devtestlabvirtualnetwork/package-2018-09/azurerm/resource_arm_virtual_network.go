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

            "name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "location": azure.SchemaLocation(),

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

            "type": {
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
    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, name, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Virtual Network %q (Resource Group %q): %+v", name, resourceGroup, err)
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

    virtualNetwork := devtestlab.VirtualNetworkFragment{
        Location: utils.String(location),
        VirtualNetworkPropertiesFragment: &devtestlab.VirtualNetworkPropertiesFragment{
            AllowedSubnets: expandArmVirtualNetworkSubnetFragment(allowedSubnets),
            Description: utils.String(description),
            ExternalProviderResourceID: utils.String(externalProviderResourceId),
            SubnetOverrides: expandArmVirtualNetworkSubnetOverrideFragment(subnetOverrides),
        },
        Tags: tags.Expand(t),
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, name, name, virtualNetwork)
    if err != nil {
        return fmt.Errorf("Error creating Virtual Network %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Virtual Network %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, name, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Virtual Network %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Virtual Network %q (Resource Group %q) ID", name, resourceGroup)
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
    name := id.Path["labs"]
    name := id.Path["virtualnetworks"]

    resp, err := client.Get(ctx, resourceGroup, name, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Virtual Network %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Virtual Network %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", name)
    d.Set("resource_group", resourceGroup)
    d.Set("type", resp.Type)

    return nil
}

func resourceArmVirtualNetworkUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).virtualNetworksClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    allowedSubnets := d.Get("allowed_subnets").([]interface{})
    description := d.Get("description").(string)
    externalProviderResourceId := d.Get("external_provider_resource_id").(string)
    subnetOverrides := d.Get("subnet_overrides").([]interface{})
    t := d.Get("tags").(map[string]interface{})

    virtualNetwork := devtestlab.VirtualNetworkFragment{
        VirtualNetworkPropertiesFragment: &devtestlab.VirtualNetworkPropertiesFragment{
            AllowedSubnets: expandArmVirtualNetworkSubnetFragment(allowedSubnets),
            Description: utils.String(description),
            ExternalProviderResourceID: utils.String(externalProviderResourceId),
            SubnetOverrides: expandArmVirtualNetworkSubnetOverrideFragment(subnetOverrides),
        },
        Tags: tags.Expand(t),
    }


    if _, err := client.Update(ctx, resourceGroup, name, name, virtualNetwork); err != nil {
        return fmt.Errorf("Error updating Virtual Network %q (Resource Group %q): %+v", name, resourceGroup, err)
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
    name := id.Path["labs"]
    name := id.Path["virtualnetworks"]

    future, err := client.Delete(ctx, resourceGroup, name, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Virtual Network %q (Resource Group %q): %+v", name, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Virtual Network %q (Resource Group %q): %+v", name, resourceGroup, err)
        }
    }

    return nil
}

func expandArmVirtualNetworkSubnetFragment(input []interface{}) *[]devtestlab.SubnetFragment {
    results := make([]devtestlab.SubnetFragment, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        resourceId := v["resource_id"].(string)
        labSubnetName := v["lab_subnet_name"].(string)
        allowPublicIp := v["allow_public_ip"].(string)

        result := devtestlab.SubnetFragment{
            AllowPublicIP: devtestlab.UsagePermissionType(allowPublicIp),
            LabSubnetName: utils.String(labSubnetName),
            ResourceID: utils.String(resourceId),
        }

        results = append(results, result)
    }
    return &results
}

func expandArmVirtualNetworkSubnetOverrideFragment(input []interface{}) *[]devtestlab.SubnetOverrideFragment {
    results := make([]devtestlab.SubnetOverrideFragment, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        resourceId := v["resource_id"].(string)
        labSubnetName := v["lab_subnet_name"].(string)
        useInVmCreationPermission := v["use_in_vm_creation_permission"].(string)
        usePublicIpAddressPermission := v["use_public_ip_address_permission"].(string)
        sharedPublicIpAddressConfiguration := v["shared_public_ip_address_configuration"].([]interface{})
        virtualNetworkPoolName := v["virtual_network_pool_name"].(string)

        result := devtestlab.SubnetOverrideFragment{
            LabSubnetName: utils.String(labSubnetName),
            ResourceID: utils.String(resourceId),
            SharedPublicIPAddressConfiguration: expandArmVirtualNetworkSubnetSharedPublicIpAddressConfigurationFragment(sharedPublicIpAddressConfiguration),
            UseInVMCreationPermission: devtestlab.UsagePermissionType(useInVmCreationPermission),
            UsePublicIPAddressPermission: devtestlab.UsagePermissionType(usePublicIpAddressPermission),
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
