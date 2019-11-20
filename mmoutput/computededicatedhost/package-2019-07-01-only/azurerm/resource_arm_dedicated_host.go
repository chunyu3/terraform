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



func resourceArmDedicatedHost() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmDedicatedHostCreate,
        Read: resourceArmDedicatedHostRead,
        Update: resourceArmDedicatedHostUpdate,
        Delete: resourceArmDedicatedHostDelete,

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

            "host_group_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "sku": {
                Type: schema.TypeList,
                Required: true,
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

            "auto_replace_on_failure": {
                Type: schema.TypeBool,
                Optional: true,
            },

            "license_type": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(compute.None),
                    string(compute.Windows_Server_Hybrid),
                    string(compute.Windows_Server_Perpetual),
                }, false),
                Default: string(compute.None),
            },

            "platform_fault_domain": {
                Type: schema.TypeInt,
                Optional: true,
            },

            "type": {
                Type: schema.TypeString,
                Computed: true,
            },

            "tags": tags.Schema(),
        },
    }
}

func resourceArmDedicatedHostCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).dedicatedHostsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    hostGroupName := d.Get("host_group_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, hostGroupName, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Dedicated Host %q (Host Group Name %q / Resource Group %q): %+v", name, hostGroupName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_dedicated_host", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    autoReplaceOnFailure := d.Get("auto_replace_on_failure").(bool)
    licenseType := d.Get("license_type").(string)
    platformFaultDomain := d.Get("platform_fault_domain").(int)
    sku := d.Get("sku").([]interface{})
    t := d.Get("tags").(map[string]interface{})

    parameters := compute.DedicatedHostUpdate{
        Location: utils.String(location),
        DedicatedHostProperties: &compute.DedicatedHostProperties{
            AutoReplaceOnFailure: utils.Bool(autoReplaceOnFailure),
            LicenseType: compute.DedicatedHostLicenseTypes(licenseType),
            PlatformFaultDomain: utils.Int32(int32(platformFaultDomain)),
        },
        Sku: expandArmDedicatedHostSku(sku),
        Tags: tags.Expand(t),
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, hostGroupName, name, parameters)
    if err != nil {
        return fmt.Errorf("Error creating Dedicated Host %q (Host Group Name %q / Resource Group %q): %+v", name, hostGroupName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Dedicated Host %q (Host Group Name %q / Resource Group %q): %+v", name, hostGroupName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, hostGroupName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Dedicated Host %q (Host Group Name %q / Resource Group %q): %+v", name, hostGroupName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Dedicated Host %q (Host Group Name %q / Resource Group %q) ID", name, hostGroupName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmDedicatedHostRead(d, meta)
}

func resourceArmDedicatedHostRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).dedicatedHostsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    hostGroupName := id.Path["hostGroups"]
    name := id.Path["hosts"]

    resp, err := client.Get(ctx, resourceGroup, hostGroupName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Dedicated Host %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Dedicated Host %q (Host Group Name %q / Resource Group %q): %+v", name, hostGroupName, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    d.Set("host_group_name", hostGroupName)
    d.Set("type", resp.Type)

    return nil
}

func resourceArmDedicatedHostUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).dedicatedHostsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    autoReplaceOnFailure := d.Get("auto_replace_on_failure").(bool)
    hostGroupName := d.Get("host_group_name").(string)
    licenseType := d.Get("license_type").(string)
    platformFaultDomain := d.Get("platform_fault_domain").(int)
    sku := d.Get("sku").([]interface{})
    t := d.Get("tags").(map[string]interface{})

    parameters := compute.DedicatedHostUpdate{
        DedicatedHostProperties: &compute.DedicatedHostProperties{
            AutoReplaceOnFailure: utils.Bool(autoReplaceOnFailure),
            LicenseType: compute.DedicatedHostLicenseTypes(licenseType),
            PlatformFaultDomain: utils.Int32(int32(platformFaultDomain)),
        },
        Sku: expandArmDedicatedHostSku(sku),
        Tags: tags.Expand(t),
    }


    future, err := client.Update(ctx, resourceGroup, hostGroupName, name, parameters)
    if err != nil {
        return fmt.Errorf("Error updating Dedicated Host %q (Host Group Name %q / Resource Group %q): %+v", name, hostGroupName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for update of Dedicated Host %q (Host Group Name %q / Resource Group %q): %+v", name, hostGroupName, resourceGroup, err)
    }

    return resourceArmDedicatedHostRead(d, meta)
}

func resourceArmDedicatedHostDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).dedicatedHostsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    hostGroupName := id.Path["hostGroups"]
    name := id.Path["hosts"]

    future, err := client.Delete(ctx, resourceGroup, hostGroupName, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Dedicated Host %q (Host Group Name %q / Resource Group %q): %+v", name, hostGroupName, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Dedicated Host %q (Host Group Name %q / Resource Group %q): %+v", name, hostGroupName, resourceGroup, err)
        }
    }

    return nil
}

func expandArmDedicatedHostSku(input []interface{}) *compute.Sku {
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
