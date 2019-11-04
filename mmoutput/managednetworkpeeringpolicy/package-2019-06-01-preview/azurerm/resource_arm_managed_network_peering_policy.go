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



func resourceArmManagedNetworkPeeringPolicy() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmManagedNetworkPeeringPolicyCreateUpdate,
        Read: resourceArmManagedNetworkPeeringPolicyRead,
        Update: resourceArmManagedNetworkPeeringPolicyCreateUpdate,
        Delete: resourceArmManagedNetworkPeeringPolicyDelete,

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

            "managed_network_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "managed_network_peering_policy_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "properties": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "properties_type": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(managednetwork.HubAndSpokeTopology),
                                string(managednetwork.MeshTopology),
                            }, false),
                        },
                        "properties_hub": {
                            Type: schema.TypeList,
                            Optional: true,
                            MaxItems: 1,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "hub_id": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                },
                            },
                        },
                        "properties_mesh": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "mesh_id": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                },
                            },
                        },
                        "properties_spokes": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "spokes_id": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                },
                            },
                        },
                    },
                },
            },

            "type": {
                Type: schema.TypeString,
                Computed: true,
            },
        },
    }
}

func resourceArmManagedNetworkPeeringPolicyCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).managedNetworkPeeringPoliciesClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    managedNetworkName := d.Get("managed_network_name").(string)
    managedNetworkPeeringPolicyName := d.Get("managed_network_peering_policy_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, managedNetworkName, managedNetworkPeeringPolicyName)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Managed Network Peering Policy (Managed Network Peering Policy Name %q / Managed Network Name %q / Resource Group %q): %+v", managedNetworkPeeringPolicyName, managedNetworkName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_managed_network_peering_policy", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    properties := d.Get("properties").([]interface{})

    managedNetworkPolicy := managednetwork.PeeringPolicy{
        Location: utils.String(location),
        PeeringPolicyProperties: expandArmManagedNetworkPeeringPolicyPeeringPolicyProperties(properties),
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, managedNetworkName, managedNetworkPeeringPolicyName, managedNetworkPolicy)
    if err != nil {
        return fmt.Errorf("Error creating Managed Network Peering Policy (Managed Network Peering Policy Name %q / Managed Network Name %q / Resource Group %q): %+v", managedNetworkPeeringPolicyName, managedNetworkName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Managed Network Peering Policy (Managed Network Peering Policy Name %q / Managed Network Name %q / Resource Group %q): %+v", managedNetworkPeeringPolicyName, managedNetworkName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, managedNetworkName, managedNetworkPeeringPolicyName)
    if err != nil {
        return fmt.Errorf("Error retrieving Managed Network Peering Policy (Managed Network Peering Policy Name %q / Managed Network Name %q / Resource Group %q): %+v", managedNetworkPeeringPolicyName, managedNetworkName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Managed Network Peering Policy (Managed Network Peering Policy Name %q / Managed Network Name %q / Resource Group %q) ID", managedNetworkPeeringPolicyName, managedNetworkName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmManagedNetworkPeeringPolicyRead(d, meta)
}

func resourceArmManagedNetworkPeeringPolicyRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).managedNetworkPeeringPoliciesClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    managedNetworkName := id.Path["managedNetworks"]
    managedNetworkPeeringPolicyName := id.Path["managedNetworkPeeringPolicies"]

    resp, err := client.Get(ctx, resourceGroup, managedNetworkName, managedNetworkPeeringPolicyName)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Managed Network Peering Policy %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Managed Network Peering Policy (Managed Network Peering Policy Name %q / Managed Network Name %q / Resource Group %q): %+v", managedNetworkPeeringPolicyName, managedNetworkName, resourceGroup, err)
    }


    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    d.Set("managed_network_name", managedNetworkName)
    d.Set("managed_network_peering_policy_name", managedNetworkPeeringPolicyName)
    if err := d.Set("properties", flattenArmManagedNetworkPeeringPolicyPeeringPolicyProperties(resp.PeeringPolicyProperties)); err != nil {
        return fmt.Errorf("Error setting `properties`: %+v", err)
    }
    d.Set("type", resp.Type)

    return nil
}


func resourceArmManagedNetworkPeeringPolicyDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).managedNetworkPeeringPoliciesClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    managedNetworkName := id.Path["managedNetworks"]
    managedNetworkPeeringPolicyName := id.Path["managedNetworkPeeringPolicies"]

    future, err := client.Delete(ctx, resourceGroup, managedNetworkName, managedNetworkPeeringPolicyName)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Managed Network Peering Policy (Managed Network Peering Policy Name %q / Managed Network Name %q / Resource Group %q): %+v", managedNetworkPeeringPolicyName, managedNetworkName, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Managed Network Peering Policy (Managed Network Peering Policy Name %q / Managed Network Name %q / Resource Group %q): %+v", managedNetworkPeeringPolicyName, managedNetworkName, resourceGroup, err)
        }
    }

    return nil
}

func expandArmManagedNetworkPeeringPolicyPeeringPolicyProperties(input []interface{}) *managednetwork.PeeringPolicyProperties {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    propertiesType := v["properties_type"].(string)
    propertiesHub := v["properties_hub"].([]interface{})
    propertiesSpokes := v["properties_spokes"].([]interface{})
    propertiesMesh := v["properties_mesh"].([]interface{})

    result := managednetwork.PeeringPolicyProperties{
        PropertiesHub: expandArmManagedNetworkPeeringPolicyResourceId(propertiesHub),
        PropertiesMesh: expandArmManagedNetworkPeeringPolicyResourceId(propertiesMesh),
        PropertiesSpokes: expandArmManagedNetworkPeeringPolicyResourceId(propertiesSpokes),
        PropertiesType: managednetwork.Type(propertiesType),
    }
    return &result
}

func expandArmManagedNetworkPeeringPolicyResourceId(input []interface{}) *managednetwork.ResourceId {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    hubId := v["hub_id"].(string)

    result := managednetwork.ResourceId{
        HubID: utils.String(hubId),
    }
    return &result
}

func expandArmManagedNetworkPeeringPolicyResourceId(input []interface{}) *[]managednetwork.ResourceId {
    results := make([]managednetwork.ResourceId, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        meshId := v["mesh_id"].(string)

        result := managednetwork.ResourceId{
            MeshID: utils.String(meshId),
        }

        results = append(results, result)
    }
    return &results
}

func expandArmManagedNetworkPeeringPolicyResourceId(input []interface{}) *[]managednetwork.ResourceId {
    results := make([]managednetwork.ResourceId, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        spokesId := v["spokes_id"].(string)

        result := managednetwork.ResourceId{
            SpokesID: utils.String(spokesId),
        }

        results = append(results, result)
    }
    return &results
}


func flattenArmManagedNetworkPeeringPolicyPeeringPolicyProperties(input *managednetwork.PeeringPolicyProperties) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    result["properties_hub"] = flattenArmManagedNetworkPeeringPolicyResourceId(input.PropertiesHub)
    result["properties_mesh"] = flattenArmManagedNetworkPeeringPolicyResourceId(input.PropertiesMesh)
    result["properties_spokes"] = flattenArmManagedNetworkPeeringPolicyResourceId(input.PropertiesSpokes)
    result["properties_type"] = string(input.PropertiesType)

    return []interface{}{result}
}

func flattenArmManagedNetworkPeeringPolicyResourceId(input *managednetwork.ResourceId) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if hubId := input.HubID; hubId != nil {
        result["hub_id"] = *hubId
    }

    return []interface{}{result}
}

func flattenArmManagedNetworkPeeringPolicyResourceId(input *[]managednetwork.ResourceId) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        if meshId := item.MeshID; meshId != nil {
            v["mesh_id"] = *meshId
        }

        results = append(results, v)
    }

    return results
}

func flattenArmManagedNetworkPeeringPolicyResourceId(input *[]managednetwork.ResourceId) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        if spokesId := item.SpokesID; spokesId != nil {
            v["spokes_id"] = *spokesId
        }

        results = append(results, v)
    }

    return results
}
