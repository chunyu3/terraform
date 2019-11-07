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



func resourceArmElasticPool() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmElasticPoolCreate,
        Read: resourceArmElasticPoolRead,
        Update: resourceArmElasticPoolUpdate,
        Delete: resourceArmElasticPoolDelete,

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

            "server_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "license_type": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(sql.LicenseIncluded),
                    string(sql.BasePrice),
                }, false),
                Default: string(sql.LicenseIncluded),
            },

            "max_size_bytes": {
                Type: schema.TypeInt,
                Optional: true,
            },

            "per_database_settings": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "max_capacity": {
                            Type: schema.TypeFloat,
                            Optional: true,
                        },
                        "min_capacity": {
                            Type: schema.TypeFloat,
                            Optional: true,
                        },
                    },
                },
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

            "zone_redundant": {
                Type: schema.TypeBool,
                Optional: true,
            },

            "creation_date": {
                Type: schema.TypeString,
                Computed: true,
            },

            "kind": {
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

            "tags": tags.Schema(),
        },
    }
}

func resourceArmElasticPoolCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).elasticPoolsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    serverName := d.Get("server_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, serverName, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Elastic Pool %q (Server Name %q / Resource Group %q): %+v", name, serverName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_elastic_pool", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    licenseType := d.Get("license_type").(string)
    maxSizeBytes := d.Get("max_size_bytes").(int)
    perDatabaseSettings := d.Get("per_database_settings").([]interface{})
    sku := d.Get("sku").([]interface{})
    zoneRedundant := d.Get("zone_redundant").(bool)
    t := d.Get("tags").(map[string]interface{})

    parameters := sql.ElasticPool{
        Location: utils.String(location),
        ElasticPoolProperties: &sql.ElasticPoolProperties{
            LicenseType: sql.ElasticPoolLicenseType(licenseType),
            MaxSizeBytes: utils.Int64(int64(maxSizeBytes)),
            PerDatabaseSettings: expandArmElasticPoolElasticPoolPerDatabaseSettings(perDatabaseSettings),
            ZoneRedundant: utils.Bool(zoneRedundant),
        },
        Sku: expandArmElasticPoolSku(sku),
        Tags: tags.Expand(t),
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, serverName, name, parameters)
    if err != nil {
        return fmt.Errorf("Error creating Elastic Pool %q (Server Name %q / Resource Group %q): %+v", name, serverName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Elastic Pool %q (Server Name %q / Resource Group %q): %+v", name, serverName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, serverName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Elastic Pool %q (Server Name %q / Resource Group %q): %+v", name, serverName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Elastic Pool %q (Server Name %q / Resource Group %q) ID", name, serverName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmElasticPoolRead(d, meta)
}

func resourceArmElasticPoolRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).elasticPoolsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    serverName := id.Path["servers"]
    name := id.Path["elasticPools"]

    resp, err := client.Get(ctx, resourceGroup, serverName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Elastic Pool %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Elastic Pool %q (Server Name %q / Resource Group %q): %+v", name, serverName, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if elasticPoolProperties := resp.ElasticPoolProperties; elasticPoolProperties != nil {
        d.Set("creation_date", (elasticPoolProperties.CreationDate).String())
        d.Set("license_type", string(elasticPoolProperties.LicenseType))
        d.Set("max_size_bytes", int(*elasticPoolProperties.MaxSizeBytes))
        if err := d.Set("per_database_settings", flattenArmElasticPoolElasticPoolPerDatabaseSettings(elasticPoolProperties.PerDatabaseSettings)); err != nil {
            return fmt.Errorf("Error setting `per_database_settings`: %+v", err)
        }
        d.Set("state", string(elasticPoolProperties.State))
        d.Set("zone_redundant", elasticPoolProperties.ZoneRedundant)
    }
    d.Set("kind", resp.Kind)
    d.Set("server_name", serverName)
    if err := d.Set("sku", flattenArmElasticPoolSku(resp.Sku)); err != nil {
        return fmt.Errorf("Error setting `sku`: %+v", err)
    }
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}

func resourceArmElasticPoolUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).elasticPoolsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    licenseType := d.Get("license_type").(string)
    maxSizeBytes := d.Get("max_size_bytes").(int)
    perDatabaseSettings := d.Get("per_database_settings").([]interface{})
    serverName := d.Get("server_name").(string)
    sku := d.Get("sku").([]interface{})
    zoneRedundant := d.Get("zone_redundant").(bool)
    t := d.Get("tags").(map[string]interface{})

    parameters := sql.ElasticPool{
        Location: utils.String(location),
        ElasticPoolProperties: &sql.ElasticPoolProperties{
            LicenseType: sql.ElasticPoolLicenseType(licenseType),
            MaxSizeBytes: utils.Int64(int64(maxSizeBytes)),
            PerDatabaseSettings: expandArmElasticPoolElasticPoolPerDatabaseSettings(perDatabaseSettings),
            ZoneRedundant: utils.Bool(zoneRedundant),
        },
        Sku: expandArmElasticPoolSku(sku),
        Tags: tags.Expand(t),
    }


    future, err := client.Update(ctx, resourceGroup, serverName, name, parameters)
    if err != nil {
        return fmt.Errorf("Error updating Elastic Pool %q (Server Name %q / Resource Group %q): %+v", name, serverName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for update of Elastic Pool %q (Server Name %q / Resource Group %q): %+v", name, serverName, resourceGroup, err)
    }

    return resourceArmElasticPoolRead(d, meta)
}

func resourceArmElasticPoolDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).elasticPoolsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    serverName := id.Path["servers"]
    name := id.Path["elasticPools"]

    future, err := client.Delete(ctx, resourceGroup, serverName, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Elastic Pool %q (Server Name %q / Resource Group %q): %+v", name, serverName, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Elastic Pool %q (Server Name %q / Resource Group %q): %+v", name, serverName, resourceGroup, err)
        }
    }

    return nil
}

func expandArmElasticPoolElasticPoolPerDatabaseSettings(input []interface{}) *sql.ElasticPoolPerDatabaseSettings {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    minCapacity := v["min_capacity"].(float64)
    maxCapacity := v["max_capacity"].(float64)

    result := sql.ElasticPoolPerDatabaseSettings{
        MaxCapacity: utils.Float(maxCapacity),
        MinCapacity: utils.Float(minCapacity),
    }
    return &result
}

func expandArmElasticPoolSku(input []interface{}) *sql.Sku {
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


func flattenArmElasticPoolElasticPoolPerDatabaseSettings(input *sql.ElasticPoolPerDatabaseSettings) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if maxCapacity := input.MaxCapacity; maxCapacity != nil {
        result["max_capacity"] = *maxCapacity
    }
    if minCapacity := input.MinCapacity; minCapacity != nil {
        result["min_capacity"] = *minCapacity
    }

    return []interface{}{result}
}

func flattenArmElasticPoolSku(input *sql.Sku) []interface{} {
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
