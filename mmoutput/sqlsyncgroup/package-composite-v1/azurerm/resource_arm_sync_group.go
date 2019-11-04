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



func resourceArmSyncGroup() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmSyncGroupCreate,
        Read: resourceArmSyncGroupRead,
        Update: resourceArmSyncGroupUpdate,
        Delete: resourceArmSyncGroupDelete,

        Importer: &schema.ResourceImporter{
            State: schema.ImportStatePassthrough,
        },


        Schema: map[string]*schema.Schema{
            "name": {
                Type: schema.TypeString,
                Computed: true,
            },

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "database_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "server_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "sync_group_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "conflict_resolution_policy": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(sql.HubWin),
                    string(sql.MemberWin),
                }, false),
                Default: string(sql.HubWin),
            },

            "hub_database_password": {
                Type: schema.TypeString,
                Optional: true,
            },

            "hub_database_user_name": {
                Type: schema.TypeString,
                Optional: true,
            },

            "interval": {
                Type: schema.TypeInt,
                Optional: true,
            },

            "schema": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "master_sync_member_name": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "tables": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "columns": {
                                        Type: schema.TypeList,
                                        Optional: true,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "data_size": {
                                                    Type: schema.TypeString,
                                                    Optional: true,
                                                },
                                                "data_type": {
                                                    Type: schema.TypeString,
                                                    Optional: true,
                                                },
                                                "quoted_name": {
                                                    Type: schema.TypeString,
                                                    Optional: true,
                                                },
                                            },
                                        },
                                    },
                                    "quoted_name": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                },
                            },
                        },
                    },
                },
            },

            "sync_database_id": {
                Type: schema.TypeString,
                Optional: true,
            },

            "last_sync_time": {
                Type: schema.TypeString,
                Computed: true,
            },

            "sync_state": {
                Type: schema.TypeString,
                Computed: true,
            },

            "type": {
                Type: schema.TypeString,
                Computed: true,
            },
        },
    }
}

func resourceArmSyncGroupCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).syncGroupsClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    databaseName := d.Get("database_name").(string)
    serverName := d.Get("server_name").(string)
    syncGroupName := d.Get("sync_group_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, serverName, databaseName, syncGroupName)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Sync Group (Sync Group Name %q / Database Name %q / Server Name %q / Resource Group %q): %+v", syncGroupName, databaseName, serverName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_sync_group", *existing.ID)
        }
    }

    conflictResolutionPolicy := d.Get("conflict_resolution_policy").(string)
    hubDatabasePassword := d.Get("hub_database_password").(string)
    hubDatabaseUserName := d.Get("hub_database_user_name").(string)
    interval := d.Get("interval").(int)
    schema := d.Get("schema").([]interface{})
    syncDatabaseId := d.Get("sync_database_id").(string)

    parameters := sql.SyncGroup{
        SyncGroupProperties: &sql.SyncGroupProperties{
            ConflictResolutionPolicy: sql.SyncConflictResolutionPolicy(conflictResolutionPolicy),
            HubDatabasePassword: utils.String(hubDatabasePassword),
            HubDatabaseUserName: utils.String(hubDatabaseUserName),
            Interval: utils.Int32(int32(interval)),
            Schema: expandArmSyncGroupSyncGroupSchema(schema),
            SyncDatabaseID: utils.String(syncDatabaseId),
        },
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, serverName, databaseName, syncGroupName, parameters)
    if err != nil {
        return fmt.Errorf("Error creating Sync Group (Sync Group Name %q / Database Name %q / Server Name %q / Resource Group %q): %+v", syncGroupName, databaseName, serverName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Sync Group (Sync Group Name %q / Database Name %q / Server Name %q / Resource Group %q): %+v", syncGroupName, databaseName, serverName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, serverName, databaseName, syncGroupName)
    if err != nil {
        return fmt.Errorf("Error retrieving Sync Group (Sync Group Name %q / Database Name %q / Server Name %q / Resource Group %q): %+v", syncGroupName, databaseName, serverName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Sync Group (Sync Group Name %q / Database Name %q / Server Name %q / Resource Group %q) ID", syncGroupName, databaseName, serverName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmSyncGroupRead(d, meta)
}

func resourceArmSyncGroupRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).syncGroupsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    serverName := id.Path["servers"]
    databaseName := id.Path["databases"]
    syncGroupName := id.Path["syncGroups"]

    resp, err := client.Get(ctx, resourceGroup, serverName, databaseName, syncGroupName)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Sync Group %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Sync Group (Sync Group Name %q / Database Name %q / Server Name %q / Resource Group %q): %+v", syncGroupName, databaseName, serverName, resourceGroup, err)
    }


    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if syncGroupProperties := resp.SyncGroupProperties; syncGroupProperties != nil {
        d.Set("conflict_resolution_policy", string(syncGroupProperties.ConflictResolutionPolicy))
        d.Set("hub_database_password", syncGroupProperties.HubDatabasePassword)
        d.Set("hub_database_user_name", syncGroupProperties.HubDatabaseUserName)
        d.Set("interval", int(*syncGroupProperties.Interval))
        d.Set("last_sync_time", (syncGroupProperties.LastSyncTime).String())
        if err := d.Set("schema", flattenArmSyncGroupSyncGroupSchema(syncGroupProperties.Schema)); err != nil {
            return fmt.Errorf("Error setting `schema`: %+v", err)
        }
        d.Set("sync_database_id", syncGroupProperties.SyncDatabaseID)
        d.Set("sync_state", string(syncGroupProperties.SyncState))
    }
    d.Set("database_name", databaseName)
    d.Set("server_name", serverName)
    d.Set("sync_group_name", syncGroupName)
    d.Set("type", resp.Type)

    return nil
}

func resourceArmSyncGroupUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).syncGroupsClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    conflictResolutionPolicy := d.Get("conflict_resolution_policy").(string)
    databaseName := d.Get("database_name").(string)
    hubDatabasePassword := d.Get("hub_database_password").(string)
    hubDatabaseUserName := d.Get("hub_database_user_name").(string)
    interval := d.Get("interval").(int)
    schema := d.Get("schema").([]interface{})
    serverName := d.Get("server_name").(string)
    syncDatabaseId := d.Get("sync_database_id").(string)
    syncGroupName := d.Get("sync_group_name").(string)

    parameters := sql.SyncGroup{
        SyncGroupProperties: &sql.SyncGroupProperties{
            ConflictResolutionPolicy: sql.SyncConflictResolutionPolicy(conflictResolutionPolicy),
            HubDatabasePassword: utils.String(hubDatabasePassword),
            HubDatabaseUserName: utils.String(hubDatabaseUserName),
            Interval: utils.Int32(int32(interval)),
            Schema: expandArmSyncGroupSyncGroupSchema(schema),
            SyncDatabaseID: utils.String(syncDatabaseId),
        },
    }


    future, err := client.Update(ctx, resourceGroup, serverName, databaseName, syncGroupName, parameters)
    if err != nil {
        return fmt.Errorf("Error updating Sync Group (Sync Group Name %q / Database Name %q / Server Name %q / Resource Group %q): %+v", syncGroupName, databaseName, serverName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for update of Sync Group (Sync Group Name %q / Database Name %q / Server Name %q / Resource Group %q): %+v", syncGroupName, databaseName, serverName, resourceGroup, err)
    }

    return resourceArmSyncGroupRead(d, meta)
}

func resourceArmSyncGroupDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).syncGroupsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    serverName := id.Path["servers"]
    databaseName := id.Path["databases"]
    syncGroupName := id.Path["syncGroups"]

    future, err := client.Delete(ctx, resourceGroup, serverName, databaseName, syncGroupName)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Sync Group (Sync Group Name %q / Database Name %q / Server Name %q / Resource Group %q): %+v", syncGroupName, databaseName, serverName, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Sync Group (Sync Group Name %q / Database Name %q / Server Name %q / Resource Group %q): %+v", syncGroupName, databaseName, serverName, resourceGroup, err)
        }
    }

    return nil
}

func expandArmSyncGroupSyncGroupSchema(input []interface{}) *sql.SyncGroupSchema {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    tables := v["tables"].([]interface{})
    masterSyncMemberName := v["master_sync_member_name"].(string)

    result := sql.SyncGroupSchema{
        MasterSyncMemberName: utils.String(masterSyncMemberName),
        Tables: expandArmSyncGroupSyncGroupSchemaTable(tables),
    }
    return &result
}

func expandArmSyncGroupSyncGroupSchemaTable(input []interface{}) *[]sql.SyncGroupSchemaTable {
    results := make([]sql.SyncGroupSchemaTable, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        columns := v["columns"].([]interface{})
        quotedName := v["quoted_name"].(string)

        result := sql.SyncGroupSchemaTable{
            Columns: expandArmSyncGroupSyncGroupSchemaTableColumn(columns),
            QuotedName: utils.String(quotedName),
        }

        results = append(results, result)
    }
    return &results
}

func expandArmSyncGroupSyncGroupSchemaTableColumn(input []interface{}) *[]sql.SyncGroupSchemaTableColumn {
    results := make([]sql.SyncGroupSchemaTableColumn, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        quotedName := v["quoted_name"].(string)
        dataSize := v["data_size"].(string)
        dataType := v["data_type"].(string)

        result := sql.SyncGroupSchemaTableColumn{
            DataSize: utils.String(dataSize),
            DataType: utils.String(dataType),
            QuotedName: utils.String(quotedName),
        }

        results = append(results, result)
    }
    return &results
}


func flattenArmSyncGroupSyncGroupSchema(input *sql.SyncGroupSchema) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if masterSyncMemberName := input.MasterSyncMemberName; masterSyncMemberName != nil {
        result["master_sync_member_name"] = *masterSyncMemberName
    }
    result["tables"] = flattenArmSyncGroupSyncGroupSchemaTable(input.Tables)

    return []interface{}{result}
}

func flattenArmSyncGroupSyncGroupSchemaTable(input *[]sql.SyncGroupSchemaTable) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        v["columns"] = flattenArmSyncGroupSyncGroupSchemaTableColumn(item.Columns)
        if quotedName := item.QuotedName; quotedName != nil {
            v["quoted_name"] = *quotedName
        }

        results = append(results, v)
    }

    return results
}

func flattenArmSyncGroupSyncGroupSchemaTableColumn(input *[]sql.SyncGroupSchemaTableColumn) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        if dataSize := item.DataSize; dataSize != nil {
            v["data_size"] = *dataSize
        }
        if dataType := item.DataType; dataType != nil {
            v["data_type"] = *dataType
        }
        if quotedName := item.QuotedName; quotedName != nil {
            v["quoted_name"] = *quotedName
        }

        results = append(results, v)
    }

    return results
}
