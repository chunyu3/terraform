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



func resourceArmSyncMember() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmSyncMemberCreate,
        Read: resourceArmSyncMemberRead,
        Update: resourceArmSyncMemberUpdate,
        Delete: resourceArmSyncMemberDelete,

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

            "sync_member_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "database_name": {
                Type: schema.TypeString,
                Optional: true,
            },

            "database_type": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(sql.AzureSqlDatabase),
                    string(sql.SqlServerDatabase),
                }, false),
                Default: string(sql.AzureSqlDatabase),
            },

            "password": {
                Type: schema.TypeString,
                Optional: true,
            },

            "server_name": {
                Type: schema.TypeString,
                Optional: true,
            },

            "sql_server_database_id": {
                Type: schema.TypeString,
                Optional: true,
            },

            "sync_agent_id": {
                Type: schema.TypeString,
                Optional: true,
            },

            "sync_direction": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(sql.Bidirectional),
                    string(sql.OneWayMemberToHub),
                    string(sql.OneWayHubToMember),
                }, false),
                Default: string(sql.Bidirectional),
            },

            "user_name": {
                Type: schema.TypeString,
                Optional: true,
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

func resourceArmSyncMemberCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).syncMembersClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    databaseName := d.Get("database_name").(string)
    serverName := d.Get("server_name").(string)
    syncGroupName := d.Get("sync_group_name").(string)
    syncMemberName := d.Get("sync_member_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, serverName, databaseName, syncGroupName, syncMemberName)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Sync Member (Sync Member Name %q / Sync Group Name %q / Database Name %q / Server Name %q / Resource Group %q): %+v", syncMemberName, syncGroupName, databaseName, serverName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_sync_member", *existing.ID)
        }
    }

    databaseName := d.Get("database_name").(string)
    databaseType := d.Get("database_type").(string)
    password := d.Get("password").(string)
    serverName := d.Get("server_name").(string)
    sqlServerDatabaseId := d.Get("sql_server_database_id").(string)
    syncAgentId := d.Get("sync_agent_id").(string)
    syncDirection := d.Get("sync_direction").(string)
    userName := d.Get("user_name").(string)

    parameters := sql.SyncMember{
        SyncMemberProperties: &sql.SyncMemberProperties{
            DatabaseName: utils.String(databaseName),
            DatabaseType: sql.SyncMemberDbType(databaseType),
            Password: utils.String(password),
            ServerName: utils.String(serverName),
            SqlServerDatabaseID: utils.String(sqlServerDatabaseId),
            SyncAgentID: utils.String(syncAgentId),
            SyncDirection: sql.SyncDirection(syncDirection),
            UserName: utils.String(userName),
        },
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, serverName, databaseName, syncGroupName, syncMemberName, parameters)
    if err != nil {
        return fmt.Errorf("Error creating Sync Member (Sync Member Name %q / Sync Group Name %q / Database Name %q / Server Name %q / Resource Group %q): %+v", syncMemberName, syncGroupName, databaseName, serverName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Sync Member (Sync Member Name %q / Sync Group Name %q / Database Name %q / Server Name %q / Resource Group %q): %+v", syncMemberName, syncGroupName, databaseName, serverName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, serverName, databaseName, syncGroupName, syncMemberName)
    if err != nil {
        return fmt.Errorf("Error retrieving Sync Member (Sync Member Name %q / Sync Group Name %q / Database Name %q / Server Name %q / Resource Group %q): %+v", syncMemberName, syncGroupName, databaseName, serverName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Sync Member (Sync Member Name %q / Sync Group Name %q / Database Name %q / Server Name %q / Resource Group %q) ID", syncMemberName, syncGroupName, databaseName, serverName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmSyncMemberRead(d, meta)
}

func resourceArmSyncMemberRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).syncMembersClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    serverName := id.Path["servers"]
    databaseName := id.Path["databases"]
    syncGroupName := id.Path["syncGroups"]
    syncMemberName := id.Path["syncMembers"]

    resp, err := client.Get(ctx, resourceGroup, serverName, databaseName, syncGroupName, syncMemberName)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Sync Member %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Sync Member (Sync Member Name %q / Sync Group Name %q / Database Name %q / Server Name %q / Resource Group %q): %+v", syncMemberName, syncGroupName, databaseName, serverName, resourceGroup, err)
    }


    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    d.Set("database_name", databaseName)
    if syncMemberProperties := resp.SyncMemberProperties; syncMemberProperties != nil {
        d.Set("database_name", syncMemberProperties.DatabaseName)
        d.Set("database_type", string(syncMemberProperties.DatabaseType))
        d.Set("password", syncMemberProperties.Password)
        d.Set("server_name", syncMemberProperties.ServerName)
        d.Set("sql_server_database_id", syncMemberProperties.SqlServerDatabaseID)
        d.Set("sync_agent_id", syncMemberProperties.SyncAgentID)
        d.Set("sync_direction", string(syncMemberProperties.SyncDirection))
        d.Set("sync_state", string(syncMemberProperties.SyncState))
        d.Set("user_name", syncMemberProperties.UserName)
    }
    d.Set("server_name", serverName)
    d.Set("sync_group_name", syncGroupName)
    d.Set("sync_member_name", syncMemberName)
    d.Set("type", resp.Type)

    return nil
}

func resourceArmSyncMemberUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).syncMembersClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    databaseName := d.Get("database_name").(string)
    databaseName := d.Get("database_name").(string)
    databaseType := d.Get("database_type").(string)
    password := d.Get("password").(string)
    serverName := d.Get("server_name").(string)
    serverName := d.Get("server_name").(string)
    sqlServerDatabaseId := d.Get("sql_server_database_id").(string)
    syncAgentId := d.Get("sync_agent_id").(string)
    syncDirection := d.Get("sync_direction").(string)
    syncGroupName := d.Get("sync_group_name").(string)
    syncMemberName := d.Get("sync_member_name").(string)
    userName := d.Get("user_name").(string)

    parameters := sql.SyncMember{
        SyncMemberProperties: &sql.SyncMemberProperties{
            DatabaseName: utils.String(databaseName),
            DatabaseType: sql.SyncMemberDbType(databaseType),
            Password: utils.String(password),
            ServerName: utils.String(serverName),
            SqlServerDatabaseID: utils.String(sqlServerDatabaseId),
            SyncAgentID: utils.String(syncAgentId),
            SyncDirection: sql.SyncDirection(syncDirection),
            UserName: utils.String(userName),
        },
    }


    future, err := client.Update(ctx, resourceGroup, serverName, databaseName, syncGroupName, syncMemberName, parameters)
    if err != nil {
        return fmt.Errorf("Error updating Sync Member (Sync Member Name %q / Sync Group Name %q / Database Name %q / Server Name %q / Resource Group %q): %+v", syncMemberName, syncGroupName, databaseName, serverName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for update of Sync Member (Sync Member Name %q / Sync Group Name %q / Database Name %q / Server Name %q / Resource Group %q): %+v", syncMemberName, syncGroupName, databaseName, serverName, resourceGroup, err)
    }

    return resourceArmSyncMemberRead(d, meta)
}

func resourceArmSyncMemberDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).syncMembersClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    serverName := id.Path["servers"]
    databaseName := id.Path["databases"]
    syncGroupName := id.Path["syncGroups"]
    syncMemberName := id.Path["syncMembers"]

    future, err := client.Delete(ctx, resourceGroup, serverName, databaseName, syncGroupName, syncMemberName)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Sync Member (Sync Member Name %q / Sync Group Name %q / Database Name %q / Server Name %q / Resource Group %q): %+v", syncMemberName, syncGroupName, databaseName, serverName, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Sync Member (Sync Member Name %q / Sync Group Name %q / Database Name %q / Server Name %q / Resource Group %q): %+v", syncMemberName, syncGroupName, databaseName, serverName, resourceGroup, err)
        }
    }

    return nil
}
