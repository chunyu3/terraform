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



func resourceArmSubnet() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmSubnetCreateUpdate,
        Read: resourceArmSubnetRead,
        Update: resourceArmSubnetCreateUpdate,
        Delete: resourceArmSubnetDelete,

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

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "address_prefix": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "virtual_network_name": {
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

            "ip_configurations": {
                Type: schema.TypeList,
                Optional: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "id": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },

            "network_security_group": {
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

            "route_table": {
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

            "provisioning_state": {
                Type: schema.TypeString,
                Computed: true,
            },
        },
    }
}

func resourceArmSubnetCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).subnetsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    virtualNetworkName := d.Get("virtual_network_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, virtualNetworkName, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Subnet %q (Virtual Network Name %q / Resource Group %q): %+v", name, virtualNetworkName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_subnet", *existing.ID)
        }
    }

    id := d.Get("id").(string)
    name := d.Get("name").(string)
    addressPrefix := d.Get("address_prefix").(string)
    etag := d.Get("etag").(string)
    ipConfigurations := d.Get("ip_configurations").([]interface{})
    networkSecurityGroup := d.Get("network_security_group").([]interface{})
    routeTable := d.Get("route_table").([]interface{})

    subnetParameters := network.Subnet{
        Etag: utils.String(etag),
        ID: utils.String(id),
        Name: utils.String(name),
        SubnetPropertiesFormat: &network.SubnetPropertiesFormat{
            AddressPrefix: utils.String(addressPrefix),
            IPConfigurations: expandArmSubnetSubResource(ipConfigurations),
            NetworkSecurityGroup: expandArmSubnetSubResource(networkSecurityGroup),
            RouteTable: expandArmSubnetSubResource(routeTable),
        },
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, virtualNetworkName, name, subnetParameters)
    if err != nil {
        return fmt.Errorf("Error creating Subnet %q (Virtual Network Name %q / Resource Group %q): %+v", name, virtualNetworkName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Subnet %q (Virtual Network Name %q / Resource Group %q): %+v", name, virtualNetworkName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, virtualNetworkName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Subnet %q (Virtual Network Name %q / Resource Group %q): %+v", name, virtualNetworkName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Subnet %q (Virtual Network Name %q / Resource Group %q) ID", name, virtualNetworkName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmSubnetRead(d, meta)
}

func resourceArmSubnetRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).subnetsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    virtualNetworkName := id.Path["virtualnetworks"]
    name := id.Path["subnets"]

    resp, err := client.Get(ctx, resourceGroup, virtualNetworkName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Subnet %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Subnet %q (Virtual Network Name %q / Resource Group %q): %+v", name, virtualNetworkName, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if subnetPropertiesFormat := resp.SubnetPropertiesFormat; subnetPropertiesFormat != nil {
        d.Set("address_prefix", subnetPropertiesFormat.AddressPrefix)
        if err := d.Set("ip_configurations", flattenArmSubnetSubResource(subnetPropertiesFormat.IPConfigurations)); err != nil {
            return fmt.Errorf("Error setting `ip_configurations`: %+v", err)
        }
        if err := d.Set("network_security_group", flattenArmSubnetSubResource(subnetPropertiesFormat.NetworkSecurityGroup)); err != nil {
            return fmt.Errorf("Error setting `network_security_group`: %+v", err)
        }
        d.Set("provisioning_state", subnetPropertiesFormat.ProvisioningState)
        if err := d.Set("route_table", flattenArmSubnetSubResource(subnetPropertiesFormat.RouteTable)); err != nil {
            return fmt.Errorf("Error setting `route_table`: %+v", err)
        }
    }
    d.Set("etag", resp.Etag)
    d.Set("virtual_network_name", virtualNetworkName)

    return nil
}


func resourceArmSubnetDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).subnetsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    virtualNetworkName := id.Path["virtualnetworks"]
    name := id.Path["subnets"]

    future, err := client.Delete(ctx, resourceGroup, virtualNetworkName, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Subnet %q (Virtual Network Name %q / Resource Group %q): %+v", name, virtualNetworkName, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Subnet %q (Virtual Network Name %q / Resource Group %q): %+v", name, virtualNetworkName, resourceGroup, err)
        }
    }

    return nil
}

func expandArmSubnetSubResource(input []interface{}) *[]network.SubResource {
    results := make([]network.SubResource, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        id := v["id"].(string)

        result := network.SubResource{
            ID: utils.String(id),
        }

        results = append(results, result)
    }
    return &results
}

func expandArmSubnetSubResource(input []interface{}) *network.SubResource {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    id := v["id"].(string)

    result := network.SubResource{
        ID: utils.String(id),
    }
    return &result
}


func flattenArmSubnetSubResource(input *[]network.SubResource) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        if id := item.ID; id != nil {
            v["id"] = *id
        }

        results = append(results, v)
    }

    return results
}

func flattenArmSubnetSubResource(input *network.SubResource) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if id := input.ID; id != nil {
        result["id"] = *id
    }

    return []interface{}{result}
}
