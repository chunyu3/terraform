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



func resourceArmInstancePool() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmInstancePoolCreate,
        Read: resourceArmInstancePoolRead,
        Update: resourceArmInstancePoolUpdate,
        Delete: resourceArmInstancePoolDelete,

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

            "instance_pool_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "license_type": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(sql.LicenseIncluded),
                    string(sql.BasePrice),
                }, false),
            },

            "subnet_id": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "v_cores": {
                Type: schema.TypeInt,
                Required: true,
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
                            Type: schema.TypeInt,
                            Optional: true,
                        },
                        "family": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "size": {
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

            "type": {
                Type: schema.TypeString,
                Computed: true,
            },

            "tags": tags.Schema(),
        },
    }
}

func resourceArmInstancePoolCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).instancePoolsClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    instancePoolName := d.Get("instance_pool_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, instancePoolName)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Instance Pool (Instance Pool Name %q / Resource Group %q): %+v", instancePoolName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_instance_pool", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    licenseType := d.Get("license_type").(string)
    sku := d.Get("sku").([]interface{})
    subnetId := d.Get("subnet_id").(string)
    vCores := d.Get("v_cores").(int)
    t := d.Get("tags").(map[string]interface{})

    parameters := sql.InstancePool{
        Location: utils.String(location),
        InstancePoolProperties: &sql.InstancePoolProperties{
            LicenseType: sql.InstancePoolLicenseType(licenseType),
            SubnetID: utils.String(subnetId),
            VCores: utils.Int32(int32(vCores)),
        },
        Sku: expandArmInstancePoolSku(sku),
        Tags: tags.Expand(t),
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, instancePoolName, parameters)
    if err != nil {
        return fmt.Errorf("Error creating Instance Pool (Instance Pool Name %q / Resource Group %q): %+v", instancePoolName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Instance Pool (Instance Pool Name %q / Resource Group %q): %+v", instancePoolName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, instancePoolName)
    if err != nil {
        return fmt.Errorf("Error retrieving Instance Pool (Instance Pool Name %q / Resource Group %q): %+v", instancePoolName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Instance Pool (Instance Pool Name %q / Resource Group %q) ID", instancePoolName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmInstancePoolRead(d, meta)
}

func resourceArmInstancePoolRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).instancePoolsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    instancePoolName := id.Path["instancePools"]

    resp, err := client.Get(ctx, resourceGroup, instancePoolName)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Instance Pool %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Instance Pool (Instance Pool Name %q / Resource Group %q): %+v", instancePoolName, resourceGroup, err)
    }


    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    d.Set("instance_pool_name", instancePoolName)
    if instancePoolProperties := resp.InstancePoolProperties; instancePoolProperties != nil {
        d.Set("license_type", string(instancePoolProperties.LicenseType))
        d.Set("subnet_id", instancePoolProperties.SubnetID)
        d.Set("v_cores", int(*instancePoolProperties.VCores))
    }
    if err := d.Set("sku", flattenArmInstancePoolSku(resp.Sku)); err != nil {
        return fmt.Errorf("Error setting `sku`: %+v", err)
    }
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}

func resourceArmInstancePoolUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).instancePoolsClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    instancePoolName := d.Get("instance_pool_name").(string)
    licenseType := d.Get("license_type").(string)
    sku := d.Get("sku").([]interface{})
    subnetId := d.Get("subnet_id").(string)
    vCores := d.Get("v_cores").(int)
    t := d.Get("tags").(map[string]interface{})

    parameters := sql.InstancePool{
        Location: utils.String(location),
        InstancePoolProperties: &sql.InstancePoolProperties{
            LicenseType: sql.InstancePoolLicenseType(licenseType),
            SubnetID: utils.String(subnetId),
            VCores: utils.Int32(int32(vCores)),
        },
        Sku: expandArmInstancePoolSku(sku),
        Tags: tags.Expand(t),
    }


    future, err := client.Update(ctx, resourceGroup, instancePoolName, parameters)
    if err != nil {
        return fmt.Errorf("Error updating Instance Pool (Instance Pool Name %q / Resource Group %q): %+v", instancePoolName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for update of Instance Pool (Instance Pool Name %q / Resource Group %q): %+v", instancePoolName, resourceGroup, err)
    }

    return resourceArmInstancePoolRead(d, meta)
}

func resourceArmInstancePoolDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).instancePoolsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    instancePoolName := id.Path["instancePools"]

    future, err := client.Delete(ctx, resourceGroup, instancePoolName)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Instance Pool (Instance Pool Name %q / Resource Group %q): %+v", instancePoolName, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Instance Pool (Instance Pool Name %q / Resource Group %q): %+v", instancePoolName, resourceGroup, err)
        }
    }

    return nil
}

func expandArmInstancePoolSku(input []interface{}) *sql.Sku {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    name := v["name"].(string)
    tier := v["tier"].(string)
    size := v["size"].(string)
    family := v["family"].(string)
    capacity := v["capacity"].(int)

    result := sql.Sku{
        Capacity: utils.Int32(int32(capacity)),
        Family: utils.String(family),
        Name: utils.String(name),
        Size: utils.String(size),
        Tier: utils.String(tier),
    }
    return &result
}


func flattenArmInstancePoolSku(input *sql.Sku) []interface{} {
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
    if family := input.Family; family != nil {
        result["family"] = *family
    }
    if size := input.Size; size != nil {
        result["size"] = *size
    }
    if tier := input.Tier; tier != nil {
        result["tier"] = *tier
    }

    return []interface{}{result}
}
