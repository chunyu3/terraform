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



func resourceArmNamespace() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmNamespaceCreate,
        Read: resourceArmNamespaceRead,
        Update: resourceArmNamespaceUpdate,
        Delete: resourceArmNamespaceDelete,

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

            "sku": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "name": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(servicebus.Basic),
                                string(servicebus.Standard),
                                string(servicebus.Premium),
                            }, false),
                        },
                        "capacity": {
                            Type: schema.TypeInt,
                            Optional: true,
                        },
                        "tier": {
                            Type: schema.TypeString,
                            Optional: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(servicebus.Basic),
                                string(servicebus.Standard),
                                string(servicebus.Premium),
                            }, false),
                            Default: string(servicebus.Basic),
                        },
                    },
                },
            },

            "zone_redundant": {
                Type: schema.TypeBool,
                Optional: true,
            },

            "created_at": {
                Type: schema.TypeString,
                Computed: true,
            },

            "metric_id": {
                Type: schema.TypeString,
                Computed: true,
            },

            "provisioning_state": {
                Type: schema.TypeString,
                Computed: true,
            },

            "service_bus_endpoint": {
                Type: schema.TypeString,
                Computed: true,
            },

            "type": {
                Type: schema.TypeString,
                Computed: true,
            },

            "updated_at": {
                Type: schema.TypeString,
                Computed: true,
            },

            "tags": tags.Schema(),
        },
    }
}

func resourceArmNamespaceCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).namespacesClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Namespace %q (Resource Group %q): %+v", name, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_namespace", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    sku := d.Get("sku").([]interface{})
    zoneRedundant := d.Get("zone_redundant").(bool)
    t := d.Get("tags").(map[string]interface{})

    parameters := servicebus.SBNamespace{
        Location: utils.String(location),
        SBNamespaceProperties: &servicebus.SBNamespaceProperties{
            ZoneRedundant: utils.Bool(zoneRedundant),
        },
        Sku: expandArmNamespaceSBSku(sku),
        Tags: tags.Expand(t),
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, name, parameters)
    if err != nil {
        return fmt.Errorf("Error creating Namespace %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Namespace %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Namespace %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Namespace %q (Resource Group %q) ID", name, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmNamespaceRead(d, meta)
}

func resourceArmNamespaceRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).namespacesClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["namespaces"]

    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Namespace %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Namespace %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if sBNamespaceProperties := resp.SBNamespaceProperties; sBNamespaceProperties != nil {
        d.Set("created_at", (sBNamespaceProperties.CreatedAt).String())
        d.Set("metric_id", sBNamespaceProperties.MetricID)
        d.Set("provisioning_state", sBNamespaceProperties.ProvisioningState)
        d.Set("service_bus_endpoint", sBNamespaceProperties.ServiceBusEndpoint)
        d.Set("updated_at", (sBNamespaceProperties.UpdatedAt).String())
        d.Set("zone_redundant", sBNamespaceProperties.ZoneRedundant)
    }
    if err := d.Set("sku", flattenArmNamespaceSBSku(resp.Sku)); err != nil {
        return fmt.Errorf("Error setting `sku`: %+v", err)
    }
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}

func resourceArmNamespaceUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).namespacesClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    sku := d.Get("sku").([]interface{})
    zoneRedundant := d.Get("zone_redundant").(bool)
    t := d.Get("tags").(map[string]interface{})

    parameters := servicebus.SBNamespace{
        Location: utils.String(location),
        SBNamespaceProperties: &servicebus.SBNamespaceProperties{
            ZoneRedundant: utils.Bool(zoneRedundant),
        },
        Sku: expandArmNamespaceSBSku(sku),
        Tags: tags.Expand(t),
    }


    if _, err := client.Update(ctx, resourceGroup, name, parameters); err != nil {
        return fmt.Errorf("Error updating Namespace %q (Resource Group %q): %+v", name, resourceGroup, err)
    }

    return resourceArmNamespaceRead(d, meta)
}

func resourceArmNamespaceDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).namespacesClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["namespaces"]

    future, err := client.Delete(ctx, resourceGroup, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Namespace %q (Resource Group %q): %+v", name, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Namespace %q (Resource Group %q): %+v", name, resourceGroup, err)
        }
    }

    return nil
}

func expandArmNamespaceSBSku(input []interface{}) *servicebus.SBSku {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    name := v["name"].(string)
    tier := v["tier"].(string)
    capacity := v["capacity"].(int)

    result := servicebus.SBSku{
        Capacity: utils.Int32(int32(capacity)),
        Name: servicebus.SkuName(name),
        Tier: servicebus.SkuTier(tier),
    }
    return &result
}


func flattenArmNamespaceSBSku(input *servicebus.SBSku) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    result["name"] = string(input.Name)
    if capacity := input.Capacity; capacity != nil {
        result["capacity"] = int(*capacity)
    }
    result["tier"] = string(input.Tier)

    return []interface{}{result}
}