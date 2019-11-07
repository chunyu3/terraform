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



func resourceArmFailoverGroup() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmFailoverGroupCreate,
        Read: resourceArmFailoverGroupRead,
        Update: resourceArmFailoverGroupUpdate,
        Delete: resourceArmFailoverGroupDelete,

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

            "partner_servers": {
                Type: schema.TypeList,
                Required: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "id": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                    },
                },
            },

            "read_write_endpoint": {
                Type: schema.TypeList,
                Required: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "failover_policy": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(sql.Manual),
                                string(sql.Automatic),
                            }, false),
                        },
                        "failover_with_data_loss_grace_period_minutes": {
                            Type: schema.TypeInt,
                            Optional: true,
                        },
                    },
                },
            },

            "server_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "databases": {
                Type: schema.TypeList,
                Optional: true,
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
            },

            "read_only_endpoint": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "failover_policy": {
                            Type: schema.TypeString,
                            Optional: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(sql.Disabled),
                                string(sql.Enabled),
                            }, false),
                            Default: string(sql.Disabled),
                        },
                    },
                },
            },

            "replication_role": {
                Type: schema.TypeString,
                Computed: true,
            },

            "replication_state": {
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

func resourceArmFailoverGroupCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).failoverGroupsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    serverName := d.Get("server_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, serverName, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Failover Group %q (Server Name %q / Resource Group %q): %+v", name, serverName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_failover_group", *existing.ID)
        }
    }

    databases := d.Get("databases").([]interface{})
    partnerServers := d.Get("partner_servers").([]interface{})
    readOnlyEndpoint := d.Get("read_only_endpoint").([]interface{})
    readWriteEndpoint := d.Get("read_write_endpoint").([]interface{})
    t := d.Get("tags").(map[string]interface{})

    parameters := sql.FailoverGroup{
        FailoverGroupProperties: &sql.FailoverGroupProperties{
            Databases: utils.ExpandStringSlice(databases),
            PartnerServers: expandArmFailoverGroupPartnerInfo(partnerServers),
            ReadOnlyEndpoint: expandArmFailoverGroupFailoverGroupReadOnlyEndpoint(readOnlyEndpoint),
            ReadWriteEndpoint: expandArmFailoverGroupFailoverGroupReadWriteEndpoint(readWriteEndpoint),
        },
        Tags: tags.Expand(t),
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, serverName, name, parameters)
    if err != nil {
        return fmt.Errorf("Error creating Failover Group %q (Server Name %q / Resource Group %q): %+v", name, serverName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Failover Group %q (Server Name %q / Resource Group %q): %+v", name, serverName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, serverName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Failover Group %q (Server Name %q / Resource Group %q): %+v", name, serverName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Failover Group %q (Server Name %q / Resource Group %q) ID", name, serverName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmFailoverGroupRead(d, meta)
}

func resourceArmFailoverGroupRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).failoverGroupsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    serverName := id.Path["servers"]
    name := id.Path["failoverGroups"]

    resp, err := client.Get(ctx, resourceGroup, serverName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Failover Group %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Failover Group %q (Server Name %q / Resource Group %q): %+v", name, serverName, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if failoverGroupProperties := resp.FailoverGroupProperties; failoverGroupProperties != nil {
        d.Set("databases", utils.FlattenStringSlice(failoverGroupProperties.Databases))
        if err := d.Set("partner_servers", flattenArmFailoverGroupPartnerInfo(failoverGroupProperties.PartnerServers)); err != nil {
            return fmt.Errorf("Error setting `partner_servers`: %+v", err)
        }
        if err := d.Set("read_only_endpoint", flattenArmFailoverGroupFailoverGroupReadOnlyEndpoint(failoverGroupProperties.ReadOnlyEndpoint)); err != nil {
            return fmt.Errorf("Error setting `read_only_endpoint`: %+v", err)
        }
        if err := d.Set("read_write_endpoint", flattenArmFailoverGroupFailoverGroupReadWriteEndpoint(failoverGroupProperties.ReadWriteEndpoint)); err != nil {
            return fmt.Errorf("Error setting `read_write_endpoint`: %+v", err)
        }
        d.Set("replication_role", string(failoverGroupProperties.ReplicationRole))
        d.Set("replication_state", failoverGroupProperties.ReplicationState)
    }
    d.Set("server_name", serverName)
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}

func resourceArmFailoverGroupUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).failoverGroupsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    databases := d.Get("databases").([]interface{})
    partnerServers := d.Get("partner_servers").([]interface{})
    readOnlyEndpoint := d.Get("read_only_endpoint").([]interface{})
    readWriteEndpoint := d.Get("read_write_endpoint").([]interface{})
    serverName := d.Get("server_name").(string)
    t := d.Get("tags").(map[string]interface{})

    parameters := sql.FailoverGroup{
        FailoverGroupProperties: &sql.FailoverGroupProperties{
            Databases: utils.ExpandStringSlice(databases),
            PartnerServers: expandArmFailoverGroupPartnerInfo(partnerServers),
            ReadOnlyEndpoint: expandArmFailoverGroupFailoverGroupReadOnlyEndpoint(readOnlyEndpoint),
            ReadWriteEndpoint: expandArmFailoverGroupFailoverGroupReadWriteEndpoint(readWriteEndpoint),
        },
        Tags: tags.Expand(t),
    }


    future, err := client.Update(ctx, resourceGroup, serverName, name, parameters)
    if err != nil {
        return fmt.Errorf("Error updating Failover Group %q (Server Name %q / Resource Group %q): %+v", name, serverName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for update of Failover Group %q (Server Name %q / Resource Group %q): %+v", name, serverName, resourceGroup, err)
    }

    return resourceArmFailoverGroupRead(d, meta)
}

func resourceArmFailoverGroupDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).failoverGroupsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    serverName := id.Path["servers"]
    name := id.Path["failoverGroups"]

    future, err := client.Delete(ctx, resourceGroup, serverName, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Failover Group %q (Server Name %q / Resource Group %q): %+v", name, serverName, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Failover Group %q (Server Name %q / Resource Group %q): %+v", name, serverName, resourceGroup, err)
        }
    }

    return nil
}

func expandArmFailoverGroupPartnerInfo(input []interface{}) *[]sql.PartnerInfo {
    results := make([]sql.PartnerInfo, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        id := v["id"].(string)

        result := sql.PartnerInfo{
            ID: utils.String(id),
        }

        results = append(results, result)
    }
    return &results
}

func expandArmFailoverGroupFailoverGroupReadOnlyEndpoint(input []interface{}) *sql.FailoverGroupReadOnlyEndpoint {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    failoverPolicy := v["failover_policy"].(string)

    result := sql.FailoverGroupReadOnlyEndpoint{
        FailoverPolicy: sql.ReadOnlyEndpointFailoverPolicy(failoverPolicy),
    }
    return &result
}

func expandArmFailoverGroupFailoverGroupReadWriteEndpoint(input []interface{}) *sql.FailoverGroupReadWriteEndpoint {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    failoverPolicy := v["failover_policy"].(string)
    failoverWithDataLossGracePeriodMinutes := v["failover_with_data_loss_grace_period_minutes"].(int)

    result := sql.FailoverGroupReadWriteEndpoint{
        FailoverPolicy: sql.ReadWriteEndpointFailoverPolicy(failoverPolicy),
        FailoverWithDataLossGracePeriodMinutes: utils.Int32(int32(failoverWithDataLossGracePeriodMinutes)),
    }
    return &result
}


func flattenArmFailoverGroupPartnerInfo(input *[]sql.PartnerInfo) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        if id := item.ID; id != nil {
            v["id"] = *id
        }

        results = append(results, v)
    }

    return results
}

func flattenArmFailoverGroupFailoverGroupReadOnlyEndpoint(input *sql.FailoverGroupReadOnlyEndpoint) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    result["failover_policy"] = string(input.FailoverPolicy)

    return []interface{}{result}
}

func flattenArmFailoverGroupFailoverGroupReadWriteEndpoint(input *sql.FailoverGroupReadWriteEndpoint) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    result["failover_policy"] = string(input.FailoverPolicy)
    if failoverWithDataLossGracePeriodMinutes := input.FailoverWithDataLossGracePeriodMinutes; failoverWithDataLossGracePeriodMinutes != nil {
        result["failover_with_data_loss_grace_period_minutes"] = int(*failoverWithDataLossGracePeriodMinutes)
    }

    return []interface{}{result}
}