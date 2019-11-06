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



func resourceArmRedi() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmRediCreate,
        Read: resourceArmRediRead,
        Update: resourceArmRediUpdate,
        Delete: resourceArmRediDelete,

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

            "location": azure.SchemaLocation(),

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "sku": {
                Type: schema.TypeList,
                Required: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "capacity": {
                            Type: schema.TypeInt,
                            Required: true,
                        },
                        "family": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(redis.C),
                                string(redis.P),
                            }, false),
                        },
                        "name": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(redis.Basic),
                                string(redis.Standard),
                                string(redis.Premium),
                            }, false),
                        },
                    },
                },
            },

            "enable_non_ssl_port": {
                Type: schema.TypeBool,
                Optional: true,
            },

            "redis_configuration": {
                Type: schema.TypeMap,
                Optional: true,
                Elem: &schema.Schema{Type: schema.TypeString},
            },

            "shard_count": {
                Type: schema.TypeInt,
                Optional: true,
            },

            "static_ip": {
                Type: schema.TypeString,
                Optional: true,
            },

            "subnet_id": {
                Type: schema.TypeString,
                Optional: true,
            },

            "tenant_settings": {
                Type: schema.TypeMap,
                Optional: true,
                Elem: &schema.Schema{Type: schema.TypeString},
            },

            "zones": {
                Type: schema.TypeList,
                Optional: true,
                ForceNew: true,
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
            },

            "access_keys": {
                Type: schema.TypeList,
                Computed: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "primary_key": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "secondary_key": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },

            "host_name": {
                Type: schema.TypeString,
                Computed: true,
            },

            "linked_servers": {
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

            "port": {
                Type: schema.TypeInt,
                Computed: true,
            },

            "provisioning_state": {
                Type: schema.TypeString,
                Computed: true,
            },

            "redis_version": {
                Type: schema.TypeString,
                Computed: true,
            },

            "ssl_port": {
                Type: schema.TypeInt,
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

func resourceArmRediCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).redisClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Redi %q (Resource Group %q): %+v", name, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_redi", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    enableNonSslPort := d.Get("enable_non_ssl_port").(bool)
    redisConfiguration := d.Get("redis_configuration").(map[string]interface{})
    shardCount := d.Get("shard_count").(int)
    sku := d.Get("sku").([]interface{})
    staticIp := d.Get("static_ip").(string)
    subnetId := d.Get("subnet_id").(string)
    tenantSettings := d.Get("tenant_settings").(map[string]interface{})
    zones := d.Get("zones").([]interface{})
    t := d.Get("tags").(map[string]interface{})

    parameters := redis.CreateParameters{
        Location: utils.String(location),
        CreateProperties: &redis.CreateProperties{
            EnableNonSslPort: utils.Bool(enableNonSslPort),
            RedisConfiguration: utils.ExpandKeyValuePairs(redisConfiguration),
            ShardCount: utils.Int32(int32(shardCount)),
            Sku: expandArmRediSku(sku),
            StaticIp: utils.String(staticIp),
            SubnetID: utils.String(subnetId),
            TenantSettings: utils.ExpandKeyValuePairs(tenantSettings),
        },
        Tags: tags.Expand(t),
        Zones: utils.ExpandStringSlice(zones),
    }


    future, err := client.Create(ctx, resourceGroup, name, parameters)
    if err != nil {
        return fmt.Errorf("Error creating Redi %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Redi %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Redi %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Redi %q (Resource Group %q) ID", name, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmRediRead(d, meta)
}

func resourceArmRediRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).redisClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["Redis"]

    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Redi %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Redi %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if createProperties := resp.CreateProperties; createProperties != nil {
        if err := d.Set("access_keys", flattenArmRediAccessKeys(createProperties.AccessKeys)); err != nil {
            return fmt.Errorf("Error setting `access_keys`: %+v", err)
        }
        d.Set("enable_non_ssl_port", createProperties.EnableNonSslPort)
        d.Set("host_name", createProperties.HostName)
        if err := d.Set("linked_servers", flattenArmRediLinkedServer(createProperties.LinkedServers)); err != nil {
            return fmt.Errorf("Error setting `linked_servers`: %+v", err)
        }
        d.Set("port", int(*createProperties.Port))
        d.Set("provisioning_state", createProperties.ProvisioningState)
        d.Set("redis_configuration", utils.FlattenKeyValuePairs(createProperties.RedisConfiguration))
        d.Set("redis_version", createProperties.RedisVersion)
        d.Set("shard_count", int(*createProperties.ShardCount))
        if err := d.Set("sku", flattenArmRediSku(createProperties.Sku)); err != nil {
            return fmt.Errorf("Error setting `sku`: %+v", err)
        }
        d.Set("ssl_port", int(*createProperties.SslPort))
        d.Set("static_ip", createProperties.StaticIp)
        d.Set("subnet_id", createProperties.SubnetID)
        d.Set("tenant_settings", utils.FlattenKeyValuePairs(createProperties.TenantSettings))
    }
    d.Set("type", resp.Type)
    d.Set("zones", utils.FlattenStringSlice(resp.Zones))

    return tags.FlattenAndSet(d, resp.Tags)
}

func resourceArmRediUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).redisClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    enableNonSslPort := d.Get("enable_non_ssl_port").(bool)
    redisConfiguration := d.Get("redis_configuration").(map[string]interface{})
    shardCount := d.Get("shard_count").(int)
    sku := d.Get("sku").([]interface{})
    staticIp := d.Get("static_ip").(string)
    subnetId := d.Get("subnet_id").(string)
    tenantSettings := d.Get("tenant_settings").(map[string]interface{})
    zones := d.Get("zones").([]interface{})
    t := d.Get("tags").(map[string]interface{})

    parameters := redis.CreateParameters{
        Location: utils.String(location),
        CreateProperties: &redis.CreateProperties{
            EnableNonSslPort: utils.Bool(enableNonSslPort),
            RedisConfiguration: utils.ExpandKeyValuePairs(redisConfiguration),
            ShardCount: utils.Int32(int32(shardCount)),
            Sku: expandArmRediSku(sku),
            StaticIp: utils.String(staticIp),
            SubnetID: utils.String(subnetId),
            TenantSettings: utils.ExpandKeyValuePairs(tenantSettings),
        },
        Tags: tags.Expand(t),
        Zones: utils.ExpandStringSlice(zones),
    }


    if _, err := client.Update(ctx, resourceGroup, name, parameters); err != nil {
        return fmt.Errorf("Error updating Redi %q (Resource Group %q): %+v", name, resourceGroup, err)
    }

    return resourceArmRediRead(d, meta)
}

func resourceArmRediDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).redisClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["Redis"]

    future, err := client.Delete(ctx, resourceGroup, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Redi %q (Resource Group %q): %+v", name, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Redi %q (Resource Group %q): %+v", name, resourceGroup, err)
        }
    }

    return nil
}

func expandArmRediSku(input []interface{}) *redis.Sku {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    name := v["name"].(string)
    family := v["family"].(string)
    capacity := v["capacity"].(int)

    result := redis.Sku{
        Capacity: utils.Int32(int32(capacity)),
        Family: redis.SkuFamily(family),
        Name: redis.SkuName(name),
    }
    return &result
}


func flattenArmRediAccessKeys(input *redis.AccessKeys) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})


    return []interface{}{result}
}

func flattenArmRediLinkedServer(input *[]redis.LinkedServer) []interface{} {
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

func flattenArmRediSku(input *redis.Sku) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    result["name"] = string(input.Name)
    if capacity := input.Capacity; capacity != nil {
        result["capacity"] = int(*capacity)
    }
    result["family"] = string(input.Family)

    return []interface{}{result}
}
