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



func resourceArmAvailabilitySet() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmAvailabilitySetCreate,
        Read: resourceArmAvailabilitySetRead,
        Update: resourceArmAvailabilitySetUpdate,
        Delete: resourceArmAvailabilitySetDelete,

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
                Computed: true,
            },

            "location": azure.SchemaLocation(),

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "platform_fault_domain_count": {
                Type: schema.TypeInt,
                Optional: true,
            },

            "platform_update_domain_count": {
                Type: schema.TypeInt,
                Optional: true,
            },

            "sku": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "capacity": {
                            Type: schema.TypeInt,
                            Optional: true,
                        },
                        "name": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "tier": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },

            "virtual_machines": {
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

            "statuses": {
                Type: schema.TypeList,
                Computed: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "code": {
                            Type: schema.TypeString,
                            Computed: true,
                        },
                        "display_status": {
                            Type: schema.TypeString,
                            Computed: true,
                        },
                        "level": {
                            Type: schema.TypeString,
                            Computed: true,
                        },
                        "message": {
                            Type: schema.TypeString,
                            Computed: true,
                        },
                        "time": {
                            Type: schema.TypeString,
                            Computed: true,
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

func resourceArmAvailabilitySetCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).availabilitySetsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Availability Set %q (Resource Group %q): %+v", name, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_availability_set", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    platformFaultDomainCount := d.Get("platform_fault_domain_count").(int)
    platformUpdateDomainCount := d.Get("platform_update_domain_count").(int)
    sku := d.Get("sku").([]interface{})
    virtualMachines := d.Get("virtual_machines").([]interface{})
    t := d.Get("tags").(map[string]interface{})

    parameters := compute.AvailabilitySetUpdate{
        Location: utils.String(location),
        AvailabilitySetProperties: &compute.AvailabilitySetProperties{
            PlatformFaultDomainCount: utils.Int32(int32(platformFaultDomainCount)),
            PlatformUpdateDomainCount: utils.Int32(int32(platformUpdateDomainCount)),
            VirtualMachines: expandArmAvailabilitySetSubResource(virtualMachines),
        },
        Sku: expandArmAvailabilitySetSku(sku),
        Tags: tags.Expand(t),
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroup, name, parameters); err != nil {
        return fmt.Errorf("Error creating Availability Set %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Availability Set %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Availability Set %q (Resource Group %q) ID", name, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmAvailabilitySetRead(d, meta)
}

func resourceArmAvailabilitySetRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).availabilitySetsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["availabilitySets"]

    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Availability Set %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Availability Set %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if availabilitySetProperties := resp.AvailabilitySetProperties; availabilitySetProperties != nil {
        d.Set("platform_fault_domain_count", int(*availabilitySetProperties.PlatformFaultDomainCount))
        d.Set("platform_update_domain_count", int(*availabilitySetProperties.PlatformUpdateDomainCount))
        if err := d.Set("statuses", flattenArmAvailabilitySetInstanceViewStatus(availabilitySetProperties.Statuses)); err != nil {
            return fmt.Errorf("Error setting `statuses`: %+v", err)
        }
        if err := d.Set("virtual_machines", flattenArmAvailabilitySetSubResource(availabilitySetProperties.VirtualMachines)); err != nil {
            return fmt.Errorf("Error setting `virtual_machines`: %+v", err)
        }
    }
    if err := d.Set("sku", flattenArmAvailabilitySetSku(resp.Sku)); err != nil {
        return fmt.Errorf("Error setting `sku`: %+v", err)
    }
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}

func resourceArmAvailabilitySetUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).availabilitySetsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    platformFaultDomainCount := d.Get("platform_fault_domain_count").(int)
    platformUpdateDomainCount := d.Get("platform_update_domain_count").(int)
    sku := d.Get("sku").([]interface{})
    virtualMachines := d.Get("virtual_machines").([]interface{})
    t := d.Get("tags").(map[string]interface{})

    parameters := compute.AvailabilitySetUpdate{
        AvailabilitySetProperties: &compute.AvailabilitySetProperties{
            PlatformFaultDomainCount: utils.Int32(int32(platformFaultDomainCount)),
            PlatformUpdateDomainCount: utils.Int32(int32(platformUpdateDomainCount)),
            VirtualMachines: expandArmAvailabilitySetSubResource(virtualMachines),
        },
        Sku: expandArmAvailabilitySetSku(sku),
        Tags: tags.Expand(t),
    }


    if _, err := client.Update(ctx, resourceGroup, name, parameters); err != nil {
        return fmt.Errorf("Error updating Availability Set %q (Resource Group %q): %+v", name, resourceGroup, err)
    }

    return resourceArmAvailabilitySetRead(d, meta)
}

func resourceArmAvailabilitySetDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).availabilitySetsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["availabilitySets"]

    if _, err := client.Delete(ctx, resourceGroup, name); err != nil {
        return fmt.Errorf("Error deleting Availability Set %q (Resource Group %q): %+v", name, resourceGroup, err)
    }

    return nil
}

func expandArmAvailabilitySetSubResource(input []interface{}) *[]compute.SubResource {
    results := make([]compute.SubResource, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        id := v["id"].(string)

        result := compute.SubResource{
            ID: utils.String(id),
        }

        results = append(results, result)
    }
    return &results
}

func expandArmAvailabilitySetSku(input []interface{}) *compute.Sku {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    name := v["name"].(string)
    tier := v["tier"].(string)
    capacity := v["capacity"].(int)

    result := compute.Sku{
        Capacity: utils.Int64(int64(capacity)),
        Name: utils.String(name),
        Tier: utils.String(tier),
    }
    return &result
}


func flattenArmAvailabilitySetInstanceViewStatus(input *[]compute.InstanceViewStatus) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        if code := item.Code; code != nil {
            v["code"] = *code
        }
        if displayStatus := item.DisplayStatus; displayStatus != nil {
            v["display_status"] = *displayStatus
        }
        v["level"] = string(item.Level)
        if message := item.Message; message != nil {
            v["message"] = *message
        }
        if time := item.Time; time != nil {
            v["time"] = (*time).String()
        }

        results = append(results, v)
    }

    return results
}

func flattenArmAvailabilitySetSubResource(input *[]compute.SubResource) []interface{} {
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

func flattenArmAvailabilitySetSku(input *compute.Sku) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if name := input.Name; name != nil {
        result["name"] = *name
    }
    if capacity := input.Capacity; capacity != nil {
        result["capacity"] = int(*capacity)
    }
    if tier := input.Tier; tier != nil {
        result["tier"] = *tier
    }

    return []interface{}{result}
}
