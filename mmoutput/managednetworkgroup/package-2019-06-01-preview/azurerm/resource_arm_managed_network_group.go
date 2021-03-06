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



func resourceArmManagedNetworkGroup() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmManagedNetworkGroupCreateUpdate,
        Read: resourceArmManagedNetworkGroupRead,
        Update: resourceArmManagedNetworkGroupCreateUpdate,
        Delete: resourceArmManagedNetworkGroupDelete,

        Importer: &schema.ResourceImporter{
            State: schema.ImportStatePassthrough,
        },


        Schema: map[string]*schema.Schema{
            "managed_network_group_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "managed_network_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "management_groups": {
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

            "subnets": {
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

            "subscriptions": {
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

            "virtual_networks": {
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
        },
    }
}

func resourceArmManagedNetworkGroupCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).managedNetworkGroupsClient
    ctx, cancel := timeouts.ForCreateUpdate(meta.(*ArmClient).StopContext, d)
    defer cancel()

    resourceGroupName := d.Get("resource_group").(string)
    name := d.Get("managed_network_group_name").(string)
    managedNetworkName := d.Get("managed_network_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroupName, managedNetworkName, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Managed Network Group (Managed Network Group Name %q / Managed Network Name %q / Resource Group %q): %+v", name, managedNetworkName, resourceGroupName, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_managed_network_group", *existing.ID)
        }
    }

    managementGroups := d.Get("management_groups").([]interface{})
    subnets := d.Get("subnets").([]interface{})
    subscriptions := d.Get("subscriptions").([]interface{})
    virtualNetworks := d.Get("virtual_networks").([]interface{})

    managedNetworkGroup := managednetwork.Group{
        GroupProperties: &managednetwork.GroupProperties{
            ManagementGroups: expandArmManagedNetworkGroupResourceId(managementGroups),
            Subnets: expandArmManagedNetworkGroupResourceId(subnets),
            Subscriptions: expandArmManagedNetworkGroupResourceId(subscriptions),
            VirtualNetworks: expandArmManagedNetworkGroupResourceId(virtualNetworks),
        },
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroupName, managedNetworkName, name, managedNetworkGroup)
    if err != nil {
        return fmt.Errorf("Error creating Managed Network Group (Managed Network Group Name %q / Managed Network Name %q / Resource Group %q): %+v", name, managedNetworkName, resourceGroupName, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Managed Network Group (Managed Network Group Name %q / Managed Network Name %q / Resource Group %q): %+v", name, managedNetworkName, resourceGroupName, err)
    }


    resp, err := client.Get(ctx, resourceGroupName, managedNetworkName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Managed Network Group (Managed Network Group Name %q / Managed Network Name %q / Resource Group %q): %+v", name, managedNetworkName, resourceGroupName, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Managed Network Group (Managed Network Group Name %q / Managed Network Name %q / Resource Group %q) ID", name, managedNetworkName, resourceGroupName)
    }
    d.SetId(*resp.ID)

    return resourceArmManagedNetworkGroupRead(d, meta)
}

func resourceArmManagedNetworkGroupRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).managedNetworkGroupsClient
    ctx, cancel := timeouts.ForRead(meta.(*ArmClient).StopContext, d)
    defer cancel()

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroupName := id.ResourceGroup
    managedNetworkName := id.Path["managedNetworks"]
    name := id.Path["managedNetworkGroups"]

    resp, err := client.Get(ctx, resourceGroupName, managedNetworkName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Managed Network Group %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Managed Network Group (Managed Network Group Name %q / Managed Network Name %q / Resource Group %q): %+v", name, managedNetworkName, resourceGroupName, err)
    }


    d.Set("resource_group", resourceGroupName)
    d.Set("managed_network_group_name", name)
    d.Set("managed_network_name", managedNetworkName)
    if groupProperties := resp.GroupProperties; groupProperties != nil {
        if err := d.Set("management_groups", flattenArmManagedNetworkGroupResourceId(groupProperties.ManagementGroups)); err != nil {
            return fmt.Errorf("Error setting `management_groups`: %+v", err)
        }
        if err := d.Set("subnets", flattenArmManagedNetworkGroupResourceId(groupProperties.Subnets)); err != nil {
            return fmt.Errorf("Error setting `subnets`: %+v", err)
        }
        if err := d.Set("subscriptions", flattenArmManagedNetworkGroupResourceId(groupProperties.Subscriptions)); err != nil {
            return fmt.Errorf("Error setting `subscriptions`: %+v", err)
        }
        if err := d.Set("virtual_networks", flattenArmManagedNetworkGroupResourceId(groupProperties.VirtualNetworks)); err != nil {
            return fmt.Errorf("Error setting `virtual_networks`: %+v", err)
        }
    }

    return nil
}


func resourceArmManagedNetworkGroupDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).managedNetworkGroupsClient
    ctx, cancel := timeouts.ForDelete(meta.(*ArmClient).StopContext, d)
    defer cancel()


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroupName := id.ResourceGroup
    managedNetworkName := id.Path["managedNetworks"]
    name := id.Path["managedNetworkGroups"]

    future, err := client.Delete(ctx, resourceGroupName, managedNetworkName, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Managed Network Group (Managed Network Group Name %q / Managed Network Name %q / Resource Group %q): %+v", name, managedNetworkName, resourceGroupName, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Managed Network Group (Managed Network Group Name %q / Managed Network Name %q / Resource Group %q): %+v", name, managedNetworkName, resourceGroupName, err)
        }
    }

    return nil
}

func expandArmManagedNetworkGroupResourceId(input []interface{}) *[]managednetwork.ResourceId {
    results := make([]managednetwork.ResourceId, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        iD := v["id"].(string)

        result := managednetwork.ResourceId{
            ID: utils.String(iD),
        }

        results = append(results, result)
    }
    return &results
}


func flattenArmManagedNetworkGroupResourceId(input *[]managednetwork.ResourceId) []interface{} {
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
