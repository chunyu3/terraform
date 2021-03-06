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
        Create: resourceArmRediCreateUpdate,
        Read: resourceArmRediRead,
        Update: resourceArmRediCreateUpdate,
        Delete: resourceArmRediDelete,

        Importer: &schema.ResourceImporter{
            State: schema.ImportStatePassthrough,
        },


        Schema: map[string]*schema.Schema{
            "key_type": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(redis.Primary),
                    string(redis.Secondary),
                }, false),
            },

            "location": azure.SchemaLocation(),

            "name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "reboot_type": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(redis.PrimaryNode),
                    string(redis.SecondaryNode),
                    string(redis.AllNodes),
                }, false),
            },

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

            "redis_version": {
                Type: schema.TypeString,
                Optional: true,
            },

            "shard_count": {
                Type: schema.TypeInt,
                Optional: true,
            },

            "shard_id": {
                Type: schema.TypeInt,
                Optional: true,
                ForceNew: true,
            },

            "static_ip": {
                Type: schema.TypeString,
                Optional: true,
            },

            "subnet": {
                Type: schema.TypeString,
                Optional: true,
            },

            "tags": tags.Schema(),

            "tenant_settings": {
                Type: schema.TypeMap,
                Optional: true,
                Elem: &schema.Schema{Type: schema.TypeString},
            },

            "virtual_network": {
                Type: schema.TypeString,
                Optional: true,
            },

            "host_name": {
                Type: schema.TypeString,
                Computed: true,
            },

            "id": {
                Type: schema.TypeString,
                Computed: true,
            },

            "name": {
                Type: schema.TypeString,
                Computed: true,
            },

            "port": {
                Type: schema.TypeInt,
                Computed: true,
            },

            "provisioning_state": {
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
        },
    }
}

func resourceArmRediCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).redisClient
    ctx, cancel := timeouts.ForCreateUpdate(meta.(*ArmClient).StopContext, d)
    defer cancel()

    resourceGroupName := d.Get("resource_group").(string)
    name := d.Get("name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroupName, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Redi (Name %q / Resource Group %q): %+v", name, resourceGroupName, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_redi", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    enableNonSslPort := d.Get("enable_non_ssl_port").(bool)
    keyType := d.Get("key_type").(string)
    rebootType := d.Get("reboot_type").(string)
    redisConfiguration := d.Get("redis_configuration").(map[string]interface{})
    redisVersion := d.Get("redis_version").(string)
    shardCount := d.Get("shard_count").(int)
    shardID := d.Get("shard_id").(int)
    sku := d.Get("sku").([]interface{})
    staticIP := d.Get("static_ip").(string)
    subnet := d.Get("subnet").(string)
    tenantSettings := d.Get("tenant_settings").(map[string]interface{})
    virtualNetwork := d.Get("virtual_network").(string)
    tags := d.Get("tags").(map[string]interface{})

    parameters := redis.RebootParameters{
        KeyType: redis.KeyType(keyType),
        Location: utils.String(location),
        Properties: &redis.Properties{
            EnableNonSslPort: utils.Bool(enableNonSslPort),
            RedisConfiguration: utils.ExpandKeyValuePairs(redisConfiguration),
            RedisVersion: utils.String(redisVersion),
            ShardCount: utils.Int32(int32(shardCount)),
            Sku: expandArmRediSku(sku),
            StaticIP: utils.String(staticIP),
            Subnet: utils.String(subnet),
            TenantSettings: utils.ExpandKeyValuePairs(tenantSettings),
            VirtualNetwork: utils.String(virtualNetwork),
        },
        RebootType: redis.RebootType(rebootType),
        ShardID: utils.Int32(int32(shardID)),
        Tags: tags.Expand(tags),
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroupName, name, parameters); err != nil {
        return fmt.Errorf("Error creating Redi (Name %q / Resource Group %q): %+v", name, resourceGroupName, err)
    }


    resp, err := client.Get(ctx, resourceGroupName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Redi (Name %q / Resource Group %q): %+v", name, resourceGroupName, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Redi (Name %q / Resource Group %q) ID", name, resourceGroupName)
    }
    d.SetId(*resp.ID)

    return resourceArmRediRead(d, meta)
}

func resourceArmRediRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).redisClient
    ctx, cancel := timeouts.ForRead(meta.(*ArmClient).StopContext, d)
    defer cancel()

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroupName := id.ResourceGroup
    name := id.Path["Redis"]

    resp, err := client.Get(ctx, resourceGroupName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Redi %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Redi (Name %q / Resource Group %q): %+v", name, resourceGroupName, err)
    }


    d.Set("resource_group", resourceGroupName)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if properties := resp.Properties; properties != nil {
        d.Set("enable_non_ssl_port", properties.EnableNonSslPort)
        d.Set("host_name", properties.HostName)
        d.Set("port", int(*properties.Port))
        d.Set("provisioning_state", properties.ProvisioningState)
        d.Set("redis_configuration", utils.FlattenKeyValuePairs(properties.RedisConfiguration))
        d.Set("redis_version", properties.RedisVersion)
        d.Set("shard_count", int(*properties.ShardCount))
        if err := d.Set("sku", flattenArmRediSku(properties.Sku)); err != nil {
            return fmt.Errorf("Error setting `sku`: %+v", err)
        }
        d.Set("ssl_port", int(*properties.SslPort))
        d.Set("static_ip", properties.StaticIP)
        d.Set("subnet", properties.Subnet)
        d.Set("tenant_settings", utils.FlattenKeyValuePairs(properties.TenantSettings))
        d.Set("virtual_network", properties.VirtualNetwork)
    }
    d.Set("id", resp.ID)
    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}


func resourceArmRediDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).redisClient
    ctx, cancel := timeouts.ForDelete(meta.(*ArmClient).StopContext, d)
    defer cancel()


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroupName := id.ResourceGroup
    name := id.Path["Redis"]

    if _, err := client.Delete(ctx, resourceGroupName, name); err != nil {
        return fmt.Errorf("Error deleting Redi (Name %q / Resource Group %q): %+v", name, resourceGroupName, err)
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


func flattenArmRediSku(input *redis.Sku) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if capacity := input.Capacity; capacity != nil {
        result["capacity"] = int(*capacity)
    }
    result["family"] = string(input.Family)
    result["name"] = string(input.Name)

    return []interface{}{result}
}
