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

            "container": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "files": {
                Type: schema.TypeList,
                Required: true,
                ForceNew: true,
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
            },

            "key_type": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(redis.Primary),
                    string(redis.Secondary),
                }, false),
            },

            "prefix": {
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

            "enable_non_ssl_port": {
                Type: schema.TypeBool,
                Optional: true,
            },

            "format": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
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

            "shard_id": {
                Type: schema.TypeInt,
                Optional: true,
                ForceNew: true,
            },

            "sku": {
                Type: schema.TypeList,
                Optional: true,
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
    container := d.Get("container").(string)
    enableNonSslPort := d.Get("enable_non_ssl_port").(bool)
    files := d.Get("files").([]interface{})
    format := d.Get("format").(string)
    keyType := d.Get("key_type").(string)
    prefix := d.Get("prefix").(string)
    rebootType := d.Get("reboot_type").(string)
    redisConfiguration := d.Get("redis_configuration").(map[string]interface{})
    shardCount := d.Get("shard_count").(int)
    shardId := d.Get("shard_id").(int)
    sku := d.Get("sku").([]interface{})
    staticIp := d.Get("static_ip").(string)
    subnetId := d.Get("subnet_id").(string)
    tenantSettings := d.Get("tenant_settings").(map[string]interface{})
    t := d.Get("tags").(map[string]interface{})

    parameters := redis.ExportRDBParameters{
        Container: utils.String(container),
        Files: utils.ExpandStringSlice(files),
        Format: utils.String(format),
        KeyType: redis.KeyType(keyType),
        Location: utils.String(location),
        Prefix: utils.String(prefix),
        UpdateProperties: &redis.UpdateProperties{
            EnableNonSslPort: utils.Bool(enableNonSslPort),
            RedisConfiguration: utils.ExpandKeyValuePairs(redisConfiguration),
            ShardCount: utils.Int32(int32(shardCount)),
            Sku: expandArmRediSku(sku),
            StaticIP: utils.String(staticIp),
            SubnetID: utils.String(subnetId),
            TenantSettings: utils.ExpandKeyValuePairs(tenantSettings),
        },
        RebootType: redis.RebootType(rebootType),
        ShardID: utils.Int32(int32(shardId)),
        Tags: tags.Expand(t),
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
    d.Set("type", resp.Type)

    return nil
}

func resourceArmRediUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).redisClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    container := d.Get("container").(string)
    enableNonSslPort := d.Get("enable_non_ssl_port").(bool)
    files := d.Get("files").([]interface{})
    format := d.Get("format").(string)
    keyType := d.Get("key_type").(string)
    prefix := d.Get("prefix").(string)
    rebootType := d.Get("reboot_type").(string)
    redisConfiguration := d.Get("redis_configuration").(map[string]interface{})
    shardCount := d.Get("shard_count").(int)
    shardId := d.Get("shard_id").(int)
    sku := d.Get("sku").([]interface{})
    staticIp := d.Get("static_ip").(string)
    subnetId := d.Get("subnet_id").(string)
    tenantSettings := d.Get("tenant_settings").(map[string]interface{})
    t := d.Get("tags").(map[string]interface{})

    parameters := redis.ExportRDBParameters{
        Container: utils.String(container),
        Files: utils.ExpandStringSlice(files),
        Format: utils.String(format),
        KeyType: redis.KeyType(keyType),
        Prefix: utils.String(prefix),
        UpdateProperties: &redis.UpdateProperties{
            EnableNonSslPort: utils.Bool(enableNonSslPort),
            RedisConfiguration: utils.ExpandKeyValuePairs(redisConfiguration),
            ShardCount: utils.Int32(int32(shardCount)),
            Sku: expandArmRediSku(sku),
            StaticIP: utils.String(staticIp),
            SubnetID: utils.String(subnetId),
            TenantSettings: utils.ExpandKeyValuePairs(tenantSettings),
        },
        RebootType: redis.RebootType(rebootType),
        ShardID: utils.Int32(int32(shardId)),
        Tags: tags.Expand(t),
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
