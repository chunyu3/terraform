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



func resourceArmDedicatedCloudNode() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmDedicatedCloudNodeCreate,
        Read: resourceArmDedicatedCloudNodeRead,
        Update: resourceArmDedicatedCloudNodeUpdate,
        Delete: resourceArmDedicatedCloudNodeDelete,

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

            "availability_zone_id": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "nodes_count": {
                Type: schema.TypeInt,
                Required: true,
            },

            "placement_group_id": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "purchase_id": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "referer": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "sku": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "name": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "capacity": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "description": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "family": {
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

            "sku_description": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "id": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "name": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
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

func resourceArmDedicatedCloudNodeCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).dedicatedCloudNodesClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    referer := d.Get("referer").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Dedicated Cloud Node %q (Resource Group %q): %+v", name, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_dedicated_cloud_node", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    availabilityZoneId := d.Get("availability_zone_id").(string)
    nodesCount := d.Get("nodes_count").(int)
    placementGroupId := d.Get("placement_group_id").(string)
    purchaseId := d.Get("purchase_id").(string)
    sku := d.Get("sku").([]interface{})
    skuDescription := d.Get("sku_description").([]interface{})
    t := d.Get("tags").(map[string]interface{})

    dedicatedCloudNodeRequest := vmwarecloudsimple.PatchPayload{
        Location: utils.String(location),
        DedicatedCloudNodeProperties: &vmwarecloudsimple.DedicatedCloudNodeProperties{
            AvailabilityZoneID: utils.String(availabilityZoneId),
            NodesCount: utils.Int(nodesCount),
            PlacementGroupID: utils.String(placementGroupId),
            PurchaseID: utils.String(purchaseId),
            SkuDescription: expandArmDedicatedCloudNodeSkuDescription(skuDescription),
        },
        Sku: expandArmDedicatedCloudNodeSku(sku),
        Tags: tags.Expand(t),
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, name, dedicatedCloudNodeRequest, referer)
    if err != nil {
        return fmt.Errorf("Error creating Dedicated Cloud Node %q (Referer %q / Resource Group %q): %+v", name, referer, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Dedicated Cloud Node %q (Referer %q / Resource Group %q): %+v", name, referer, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Dedicated Cloud Node %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Dedicated Cloud Node %q (Resource Group %q) ID", name, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmDedicatedCloudNodeRead(d, meta)
}

func resourceArmDedicatedCloudNodeRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).dedicatedCloudNodesClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["dedicatedCloudNodes"]

    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Dedicated Cloud Node %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Dedicated Cloud Node %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    d.Set("type", resp.Type)

    return nil
}

func resourceArmDedicatedCloudNodeUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).dedicatedCloudNodesClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    availabilityZoneId := d.Get("availability_zone_id").(string)
    nodesCount := d.Get("nodes_count").(int)
    placementGroupId := d.Get("placement_group_id").(string)
    purchaseId := d.Get("purchase_id").(string)
    sku := d.Get("sku").([]interface{})
    skuDescription := d.Get("sku_description").([]interface{})
    t := d.Get("tags").(map[string]interface{})

    dedicatedCloudNodeRequest := vmwarecloudsimple.PatchPayload{
        DedicatedCloudNodeProperties: &vmwarecloudsimple.DedicatedCloudNodeProperties{
            AvailabilityZoneID: utils.String(availabilityZoneId),
            NodesCount: utils.Int(nodesCount),
            PlacementGroupID: utils.String(placementGroupId),
            PurchaseID: utils.String(purchaseId),
            SkuDescription: expandArmDedicatedCloudNodeSkuDescription(skuDescription),
        },
        Sku: expandArmDedicatedCloudNodeSku(sku),
        Tags: tags.Expand(t),
    }


    if _, err := client.Update(ctx, resourceGroup, name, dedicatedCloudNodeRequest); err != nil {
        return fmt.Errorf("Error updating Dedicated Cloud Node %q (Resource Group %q): %+v", name, resourceGroup, err)
    }

    return resourceArmDedicatedCloudNodeRead(d, meta)
}

func resourceArmDedicatedCloudNodeDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).dedicatedCloudNodesClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["dedicatedCloudNodes"]

    if _, err := client.Delete(ctx, resourceGroup, name); err != nil {
        return fmt.Errorf("Error deleting Dedicated Cloud Node %q (Resource Group %q): %+v", name, resourceGroup, err)
    }

    return nil
}

func expandArmDedicatedCloudNodeSkuDescription(input []interface{}) *vmwarecloudsimple.SkuDescription {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    id := v["id"].(string)
    name := v["name"].(string)

    result := vmwarecloudsimple.SkuDescription{
        ID: utils.String(id),
        Name: utils.String(name),
    }
    return &result
}

func expandArmDedicatedCloudNodeSku(input []interface{}) *vmwarecloudsimple.Sku {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    capacity := v["capacity"].(string)
    description := v["description"].(string)
    family := v["family"].(string)
    name := v["name"].(string)
    tier := v["tier"].(string)

    result := vmwarecloudsimple.Sku{
        Capacity: utils.String(capacity),
        Description: utils.String(description),
        Family: utils.String(family),
        Name: utils.String(name),
        Tier: utils.String(tier),
    }
    return &result
}
