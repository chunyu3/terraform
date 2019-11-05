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



func resourceArmDdosProtectionPlan() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmDdosProtectionPlanCreateUpdate,
        Read: resourceArmDdosProtectionPlanRead,
        Update: resourceArmDdosProtectionPlanCreateUpdate,
        Delete: resourceArmDdosProtectionPlanDelete,

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

            "ddos_protection_plan_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "etag": {
                Type: schema.TypeString,
                Computed: true,
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

            "virtual_networks": {
                Type: schema.TypeList,
                Computed: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "id": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },

            "tags": tags.Schema(),
        },
    }
}

func resourceArmDdosProtectionPlanCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).ddosProtectionPlansClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    ddosProtectionPlanName := d.Get("ddos_protection_plan_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, ddosProtectionPlanName)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Ddos Protection Plan (Ddos Protection Plan Name %q / Resource Group %q): %+v", ddosProtectionPlanName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_ddos_protection_plan", *existing.ID)
        }
    }

    id := d.Get("id").(string)
    location := azure.NormalizeLocation(d.Get("location").(string))
    t := d.Get("tags").(map[string]interface{})

    parameters := network.DdosProtectionPlan{
        ID: utils.String(id),
        Location: utils.String(location),
        Tags: tags.Expand(t),
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, ddosProtectionPlanName, parameters)
    if err != nil {
        return fmt.Errorf("Error creating Ddos Protection Plan (Ddos Protection Plan Name %q / Resource Group %q): %+v", ddosProtectionPlanName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Ddos Protection Plan (Ddos Protection Plan Name %q / Resource Group %q): %+v", ddosProtectionPlanName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, ddosProtectionPlanName)
    if err != nil {
        return fmt.Errorf("Error retrieving Ddos Protection Plan (Ddos Protection Plan Name %q / Resource Group %q): %+v", ddosProtectionPlanName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Ddos Protection Plan (Ddos Protection Plan Name %q / Resource Group %q) ID", ddosProtectionPlanName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmDdosProtectionPlanRead(d, meta)
}

func resourceArmDdosProtectionPlanRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).ddosProtectionPlansClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    ddosProtectionPlanName := id.Path["ddosProtectionPlans"]

    resp, err := client.Get(ctx, resourceGroup, ddosProtectionPlanName)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Ddos Protection Plan %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Ddos Protection Plan (Ddos Protection Plan Name %q / Resource Group %q): %+v", ddosProtectionPlanName, resourceGroup, err)
    }


    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    d.Set("ddos_protection_plan_name", ddosProtectionPlanName)
    d.Set("etag", resp.Etag)
    if ddosProtectionPlanPropertiesFormat := resp.DdosProtectionPlanPropertiesFormat; ddosProtectionPlanPropertiesFormat != nil {
        d.Set("provisioning_state", ddosProtectionPlanPropertiesFormat.ProvisioningState)
        d.Set("resource_guid", ddosProtectionPlanPropertiesFormat.ResourceGuid)
        if err := d.Set("virtual_networks", flattenArmDdosProtectionPlanSubResource(ddosProtectionPlanPropertiesFormat.VirtualNetworks)); err != nil {
            return fmt.Errorf("Error setting `virtual_networks`: %+v", err)
        }
    }
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}


func resourceArmDdosProtectionPlanDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).ddosProtectionPlansClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    ddosProtectionPlanName := id.Path["ddosProtectionPlans"]

    future, err := client.Delete(ctx, resourceGroup, ddosProtectionPlanName)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Ddos Protection Plan (Ddos Protection Plan Name %q / Resource Group %q): %+v", ddosProtectionPlanName, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Ddos Protection Plan (Ddos Protection Plan Name %q / Resource Group %q): %+v", ddosProtectionPlanName, resourceGroup, err)
        }
    }

    return nil
}


func flattenArmDdosProtectionPlanSubResource(input *[]network.SubResource) []interface{} {
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
