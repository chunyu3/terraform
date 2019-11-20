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



func resourceArmVolume() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmVolumeCreate,
        Read: resourceArmVolumeRead,
        Update: resourceArmVolumeUpdate,
        Delete: resourceArmVolumeDelete,

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

            "account_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "pool_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "export_policy": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "rules": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "allowed_clients": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                    "cifs": {
                                        Type: schema.TypeBool,
                                        Optional: true,
                                    },
                                    "nfsv3": {
                                        Type: schema.TypeBool,
                                        Optional: true,
                                    },
                                    "nfsv41": {
                                        Type: schema.TypeBool,
                                        Optional: true,
                                    },
                                    "rule_index": {
                                        Type: schema.TypeInt,
                                        Optional: true,
                                    },
                                    "unix_read_only": {
                                        Type: schema.TypeBool,
                                        Optional: true,
                                    },
                                    "unix_read_write": {
                                        Type: schema.TypeBool,
                                        Optional: true,
                                    },
                                },
                            },
                        },
                    },
                },
            },

            "service_level": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(azurenetappfiles.Standard),
                    string(azurenetappfiles.Premium),
                    string(azurenetappfiles.Ultra),
                }, false),
                Default: string(azurenetappfiles.Standard),
            },

            "usage_threshold": {
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

func resourceArmVolumeCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).volumesClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    accountName := d.Get("account_name").(string)
    poolName := d.Get("pool_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, accountName, poolName, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Volume %q (Pool Name %q / Account Name %q / Resource Group %q): %+v", name, poolName, accountName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_volume", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    exportPolicy := d.Get("export_policy").([]interface{})
    serviceLevel := d.Get("service_level").(string)
    usageThreshold := d.Get("usage_threshold").(int)
    t := d.Get("tags").(map[string]interface{})

    body := azurenetappfiles.VolumePatch{
        Location: utils.String(location),
        VolumePatchProperties: &azurenetappfiles.VolumePatchProperties{
            ExportPolicy: expandArmVolumeVolumePatchProperties_exportPolicy(exportPolicy),
            ServiceLevel: azurenetappfiles.ServiceLevel(serviceLevel),
            UsageThreshold: utils.Int64(int64(usageThreshold)),
        },
        Tags: tags.Expand(t),
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, accountName, poolName, name, body)
    if err != nil {
        return fmt.Errorf("Error creating Volume %q (Pool Name %q / Account Name %q / Resource Group %q): %+v", name, poolName, accountName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Volume %q (Pool Name %q / Account Name %q / Resource Group %q): %+v", name, poolName, accountName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, accountName, poolName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Volume %q (Pool Name %q / Account Name %q / Resource Group %q): %+v", name, poolName, accountName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Volume %q (Pool Name %q / Account Name %q / Resource Group %q) ID", name, poolName, accountName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmVolumeRead(d, meta)
}

func resourceArmVolumeRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).volumesClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    accountName := id.Path["netAppAccounts"]
    poolName := id.Path["capacityPools"]
    name := id.Path["volumes"]

    resp, err := client.Get(ctx, resourceGroup, accountName, poolName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Volume %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Volume %q (Pool Name %q / Account Name %q / Resource Group %q): %+v", name, poolName, accountName, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    d.Set("account_name", accountName)
    d.Set("pool_name", poolName)
    d.Set("type", resp.Type)

    return nil
}

func resourceArmVolumeUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).volumesClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    accountName := d.Get("account_name").(string)
    exportPolicy := d.Get("export_policy").([]interface{})
    poolName := d.Get("pool_name").(string)
    serviceLevel := d.Get("service_level").(string)
    usageThreshold := d.Get("usage_threshold").(int)
    t := d.Get("tags").(map[string]interface{})

    body := azurenetappfiles.VolumePatch{
        VolumePatchProperties: &azurenetappfiles.VolumePatchProperties{
            ExportPolicy: expandArmVolumeVolumePatchProperties_exportPolicy(exportPolicy),
            ServiceLevel: azurenetappfiles.ServiceLevel(serviceLevel),
            UsageThreshold: utils.Int64(int64(usageThreshold)),
        },
        Tags: tags.Expand(t),
    }


    if _, err := client.Update(ctx, resourceGroup, accountName, poolName, name, body); err != nil {
        return fmt.Errorf("Error updating Volume %q (Pool Name %q / Account Name %q / Resource Group %q): %+v", name, poolName, accountName, resourceGroup, err)
    }

    return resourceArmVolumeRead(d, meta)
}

func resourceArmVolumeDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).volumesClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    accountName := id.Path["netAppAccounts"]
    poolName := id.Path["capacityPools"]
    name := id.Path["volumes"]

    future, err := client.Delete(ctx, resourceGroup, accountName, poolName, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Volume %q (Pool Name %q / Account Name %q / Resource Group %q): %+v", name, poolName, accountName, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Volume %q (Pool Name %q / Account Name %q / Resource Group %q): %+v", name, poolName, accountName, resourceGroup, err)
        }
    }

    return nil
}

func expandArmVolumeVolumePatchProperties_exportPolicy(input []interface{}) *azurenetappfiles.VolumePatchProperties_exportPolicy {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    rules := v["rules"].([]interface{})

    result := azurenetappfiles.VolumePatchProperties_exportPolicy{
        Rules: expandArmVolumeExportPolicyRule(rules),
    }
    return &result
}

func expandArmVolumeExportPolicyRule(input []interface{}) *[]azurenetappfiles.ExportPolicyRule {
    results := make([]azurenetappfiles.ExportPolicyRule, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        ruleIndex := v["rule_index"].(int)
        unixReadOnly := v["unix_read_only"].(bool)
        unixReadWrite := v["unix_read_write"].(bool)
        cifs := v["cifs"].(bool)
        nfsv3 := v["nfsv3"].(bool)
        nfsv41 := v["nfsv41"].(bool)
        allowedClients := v["allowed_clients"].(string)

        result := azurenetappfiles.ExportPolicyRule{
            AllowedClients: utils.String(allowedClients),
            Cifs: utils.Bool(cifs),
            Nfsv3: utils.Bool(nfsv3),
            Nfsv41: utils.Bool(nfsv41),
            RuleIndex: utils.Int(ruleIndex),
            UnixReadOnly: utils.Bool(unixReadOnly),
            UnixReadWrite: utils.Bool(unixReadWrite),
        }

        results = append(results, result)
    }
    return &results
}
