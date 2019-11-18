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



func resourceArmCluster() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmClusterCreate,
        Read: resourceArmClusterRead,
        Update: resourceArmClusterUpdate,
        Delete: resourceArmClusterDelete,

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
                                string(kusto.Standard_DS13_v2+1TB_PS),
                                string(kusto.Standard_DS13_v2+2TB_PS),
                                string(kusto.Standard_DS14_v2+3TB_PS),
                                string(kusto.Standard_DS14_v2+4TB_PS),
                                string(kusto.Standard_D13_v2),
                                string(kusto.Standard_D14_v2),
                                string(kusto.Standard_L8s),
                                string(kusto.Standard_L16s),
                                string(kusto.Standard_D11_v2),
                                string(kusto.Standard_D12_v2),
                                string(kusto.Standard_L4s),
                                string(kusto.Dev(No SLA)_Standard_D11_v2),
                            }, false),
                        },
                        "tier": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(kusto.Basic),
                                string(kusto.Standard),
                            }, false),
                        },
                        "capacity": {
                            Type: schema.TypeInt,
                            Optional: true,
                        },
                    },
                },
            },

            "trusted_external_tenants": {
                Type: schema.TypeList,
                Optional: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "value": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },

            "data_ingestion_uri": {
                Type: schema.TypeString,
                Computed: true,
            },

            "provisioning_state": {
                Type: schema.TypeString,
                Computed: true,
            },

            "state": {
                Type: schema.TypeString,
                Computed: true,
            },

            "type": {
                Type: schema.TypeString,
                Computed: true,
            },

            "uri": {
                Type: schema.TypeString,
                Computed: true,
            },

            "tags": tags.Schema(),
        },
    }
}

func resourceArmClusterCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).clustersClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Cluster %q (Resource Group %q): %+v", name, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_cluster", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    sku := d.Get("sku").([]interface{})
    trustedExternalTenants := d.Get("trusted_external_tenants").([]interface{})
    t := d.Get("tags").(map[string]interface{})

    parameters := kusto.ClusterUpdate{
        Location: utils.String(location),
        ClusterProperties: &kusto.ClusterProperties{
            TrustedExternalTenants: expandArmClusterTrustedExternalTenant(trustedExternalTenants),
        },
        Sku: expandArmClusterAzureSku(sku),
        Tags: tags.Expand(t),
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, name, parameters)
    if err != nil {
        return fmt.Errorf("Error creating Cluster %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Cluster %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Cluster %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Cluster %q (Resource Group %q) ID", name, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmClusterRead(d, meta)
}

func resourceArmClusterRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).clustersClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["clusters"]

    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Cluster %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Cluster %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if clusterProperties := resp.ClusterProperties; clusterProperties != nil {
        d.Set("data_ingestion_uri", clusterProperties.DataIngestionUri)
        d.Set("provisioning_state", string(clusterProperties.ProvisioningState))
        d.Set("state", string(clusterProperties.State))
        if err := d.Set("trusted_external_tenants", flattenArmClusterTrustedExternalTenant(clusterProperties.TrustedExternalTenants)); err != nil {
            return fmt.Errorf("Error setting `trusted_external_tenants`: %+v", err)
        }
        d.Set("uri", clusterProperties.Uri)
    }
    if err := d.Set("sku", flattenArmClusterAzureSku(resp.Sku)); err != nil {
        return fmt.Errorf("Error setting `sku`: %+v", err)
    }
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}

func resourceArmClusterUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).clustersClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    sku := d.Get("sku").([]interface{})
    trustedExternalTenants := d.Get("trusted_external_tenants").([]interface{})
    t := d.Get("tags").(map[string]interface{})

    parameters := kusto.ClusterUpdate{
        Location: utils.String(location),
        ClusterProperties: &kusto.ClusterProperties{
            TrustedExternalTenants: expandArmClusterTrustedExternalTenant(trustedExternalTenants),
        },
        Sku: expandArmClusterAzureSku(sku),
        Tags: tags.Expand(t),
    }


    future, err := client.Update(ctx, resourceGroup, name, parameters)
    if err != nil {
        return fmt.Errorf("Error updating Cluster %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for update of Cluster %q (Resource Group %q): %+v", name, resourceGroup, err)
    }

    return resourceArmClusterRead(d, meta)
}

func resourceArmClusterDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).clustersClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["clusters"]

    future, err := client.Delete(ctx, resourceGroup, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Cluster %q (Resource Group %q): %+v", name, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Cluster %q (Resource Group %q): %+v", name, resourceGroup, err)
        }
    }

    return nil
}

func expandArmClusterTrustedExternalTenant(input []interface{}) *[]kusto.TrustedExternalTenant {
    results := make([]kusto.TrustedExternalTenant, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        value := v["value"].(string)

        result := kusto.TrustedExternalTenant{
            Value: utils.String(value),
        }

        results = append(results, result)
    }
    return &results
}

func expandArmClusterAzureSku(input []interface{}) *kusto.AzureSku {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    name := v["name"].(string)
    capacity := v["capacity"].(int)
    tier := v["tier"].(string)

    result := kusto.AzureSku{
        Capacity: utils.Int(capacity),
        Name: kusto.AzureSkuName(name),
        Tier: kusto.AzureSkuTier(tier),
    }
    return &result
}


func flattenArmClusterTrustedExternalTenant(input *[]kusto.TrustedExternalTenant) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        if value := item.Value; value != nil {
            v["value"] = *value
        }

        results = append(results, v)
    }

    return results
}

func flattenArmClusterAzureSku(input *kusto.AzureSku) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    result["name"] = string(input.Name)
    if capacity := input.Capacity; capacity != nil {
        result["capacity"] = *capacity
    }
    result["tier"] = string(input.Tier)

    return []interface{}{result}
}
