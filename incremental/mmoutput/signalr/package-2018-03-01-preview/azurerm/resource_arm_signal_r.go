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



func resourceArmSignalR() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmSignalRCreate,
        Read: resourceArmSignalRRead,
        Update: resourceArmSignalRUpdate,
        Delete: resourceArmSignalRDelete,

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

            "host_name_prefix": {
                Type: schema.TypeString,
                Optional: true,
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
                            ValidateFunc: validation.StringInSlice([]string{
                                string(signalr.Free),
                                string(signalr.Basic),
                                string(signalr.Standard),
                                string(signalr.Premium),
                            }, false),
                            Default: string(signalr.Free),
                        },
                    },
                },
            },

            "external_ip": {
                Type: schema.TypeString,
                Computed: true,
            },

            "host_name": {
                Type: schema.TypeString,
                Computed: true,
            },

            "provisioning_state": {
                Type: schema.TypeString,
                Computed: true,
            },

            "public_port": {
                Type: schema.TypeInt,
                Computed: true,
            },

            "server_port": {
                Type: schema.TypeInt,
                Computed: true,
            },

            "type": {
                Type: schema.TypeString,
                Computed: true,
            },

            "version": {
                Type: schema.TypeString,
                Computed: true,
            },

            "tags": tags.Schema(),
        },
    }
}

func resourceArmSignalRCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).signalRClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Signal R %q (Resource Group %q): %+v", name, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_signal_r", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    hostNamePrefix := d.Get("host_name_prefix").(string)
    sku := d.Get("sku").([]interface{})
    t := d.Get("tags").(map[string]interface{})

    parameters := signalr.UpdateParameters{
        Location: utils.String(location),
        CreateOrUpdateProperties: &signalr.CreateOrUpdateProperties{
            HostNamePrefix: utils.String(hostNamePrefix),
        },
        Sku: expandArmSignalRResourceSku(sku),
        Tags: tags.Expand(t),
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, name, parameters)
    if err != nil {
        return fmt.Errorf("Error creating Signal R %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Signal R %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Signal R %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Signal R %q (Resource Group %q) ID", name, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmSignalRRead(d, meta)
}

func resourceArmSignalRRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).signalRClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["SignalR"]

    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Signal R %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Signal R %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if createOrUpdateProperties := resp.CreateOrUpdateProperties; createOrUpdateProperties != nil {
        d.Set("external_ip", createOrUpdateProperties.ExternalIp)
        d.Set("host_name", createOrUpdateProperties.HostName)
        d.Set("host_name_prefix", createOrUpdateProperties.HostNamePrefix)
        d.Set("provisioning_state", string(createOrUpdateProperties.ProvisioningState))
        d.Set("public_port", int(*createOrUpdateProperties.PublicPort))
        d.Set("server_port", int(*createOrUpdateProperties.ServerPort))
        d.Set("version", createOrUpdateProperties.Version)
    }
    if err := d.Set("sku", flattenArmSignalRResourceSku(resp.Sku)); err != nil {
        return fmt.Errorf("Error setting `sku`: %+v", err)
    }
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}

func resourceArmSignalRUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).signalRClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    hostNamePrefix := d.Get("host_name_prefix").(string)
    sku := d.Get("sku").([]interface{})
    t := d.Get("tags").(map[string]interface{})

    parameters := signalr.UpdateParameters{
        Location: utils.String(location),
        CreateOrUpdateProperties: &signalr.CreateOrUpdateProperties{
            HostNamePrefix: utils.String(hostNamePrefix),
        },
        Sku: expandArmSignalRResourceSku(sku),
        Tags: tags.Expand(t),
    }


    future, err := client.Update(ctx, resourceGroup, name, parameters)
    if err != nil {
        return fmt.Errorf("Error updating Signal R %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for update of Signal R %q (Resource Group %q): %+v", name, resourceGroup, err)
    }

    return resourceArmSignalRRead(d, meta)
}

func resourceArmSignalRDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).signalRClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["SignalR"]

    future, err := client.Delete(ctx, resourceGroup, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Signal R %q (Resource Group %q): %+v", name, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Signal R %q (Resource Group %q): %+v", name, resourceGroup, err)
        }
    }

    return nil
}

func expandArmSignalRResourceSku(input []interface{}) *signalr.ResourceSku {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    name := v["name"].(string)
    tier := v["tier"].(string)
    size := v["size"].(string)
    family := v["family"].(string)
    capacity := v["capacity"].(int)

    result := signalr.ResourceSku{
        Capacity: utils.Int32(int32(capacity)),
        Family: utils.String(family),
        Name: utils.String(name),
        Size: utils.String(size),
        Tier: signalr.SkuTier(tier),
    }
    return &result
}


func flattenArmSignalRResourceSku(input *signalr.ResourceSku) []interface{} {
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
    result["tier"] = string(input.Tier)

    return []interface{}{result}
}
