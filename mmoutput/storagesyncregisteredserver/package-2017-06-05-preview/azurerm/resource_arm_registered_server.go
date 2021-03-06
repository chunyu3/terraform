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



func resourceArmRegisteredServer() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmRegisteredServerCreateUpdate,
        Read: resourceArmRegisteredServerRead,
        Update: resourceArmRegisteredServerCreateUpdate,
        Delete: resourceArmRegisteredServerDelete,

        Importer: &schema.ResourceImporter{
            State: schema.ImportStatePassthrough,
        },


        Schema: map[string]*schema.Schema{
            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "server_id": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "storage_sync_service_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "agent_version": {
                Type: schema.TypeString,
                Optional: true,
            },

            "cluster_id": {
                Type: schema.TypeString,
                Optional: true,
            },

            "cluster_name": {
                Type: schema.TypeString,
                Optional: true,
            },

            "last_heart_beat": {
                Type: schema.TypeString,
                Optional: true,
            },

            "last_workflow_id": {
                Type: schema.TypeString,
                Optional: true,
            },

            "server_certificate": {
                Type: schema.TypeString,
                Optional: true,
            },

            "server_id": {
                Type: schema.TypeString,
                Optional: true,
            },

            "server_managementt_error_code": {
                Type: schema.TypeInt,
                Optional: true,
            },

            "server_os_version": {
                Type: schema.TypeString,
                Optional: true,
            },

            "server_role": {
                Type: schema.TypeString,
                Optional: true,
            },

            "storage_sync_service_uid": {
                Type: schema.TypeString,
                Optional: true,
            },

            "id": {
                Type: schema.TypeString,
                Computed: true,
            },

            "name": {
                Type: schema.TypeString,
                Computed: true,
            },

            "provisioning_state": {
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

func resourceArmRegisteredServerCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).registeredServersClient
    ctx, cancel := timeouts.ForCreateUpdate(meta.(*ArmClient).StopContext, d)
    defer cancel()

    resourceGroupName := d.Get("resource_group").(string)
    serverID := d.Get("server_id").(string)
    name := d.Get("storage_sync_service_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroupName, name, serverID)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Registered Server (Server %q / Storage Sync Service Name %q / Resource Group %q): %+v", serverID, name, resourceGroupName, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_registered_server", *existing.ID)
        }
    }

    agentVersion := d.Get("agent_version").(string)
    clusterID := d.Get("cluster_id").(string)
    clusterName := d.Get("cluster_name").(string)
    lastHeartBeat := d.Get("last_heart_beat").(string)
    lastWorkflowID := d.Get("last_workflow_id").(string)
    serverCertificate := d.Get("server_certificate").(string)
    serverID := d.Get("server_id").(string)
    serverManagementtErrorCode := d.Get("server_managementt_error_code").(int)
    serverOSVersion := d.Get("server_os_version").(string)
    serverRole := d.Get("server_role").(string)
    storageSyncServiceUID := d.Get("storage_sync_service_uid").(string)

    parameters := storagesync.RegisteredServer{
        RegisteredServerProperties: &storagesync.RegisteredServerProperties{
            AgentVersion: utils.String(agentVersion),
            ClusterID: utils.String(clusterID),
            ClusterName: utils.String(clusterName),
            LastHeartBeat: utils.String(lastHeartBeat),
            LastWorkflowID: utils.String(lastWorkflowID),
            ServerCertificate: utils.String(serverCertificate),
            ServerID: utils.String(serverID),
            ServerManagementtErrorCode: utils.Int(serverManagementtErrorCode),
            ServerOSVersion: utils.String(serverOSVersion),
            ServerRole: utils.String(serverRole),
            StorageSyncServiceUID: utils.String(storageSyncServiceUID),
        },
    }


    future, err := client.Create(ctx, resourceGroupName, name, serverID, parameters)
    if err != nil {
        return fmt.Errorf("Error creating Registered Server (Server %q / Storage Sync Service Name %q / Resource Group %q): %+v", serverID, name, resourceGroupName, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Registered Server (Server %q / Storage Sync Service Name %q / Resource Group %q): %+v", serverID, name, resourceGroupName, err)
    }


    resp, err := client.Get(ctx, resourceGroupName, name, serverID)
    if err != nil {
        return fmt.Errorf("Error retrieving Registered Server (Server %q / Storage Sync Service Name %q / Resource Group %q): %+v", serverID, name, resourceGroupName, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Registered Server (Server %q / Storage Sync Service Name %q / Resource Group %q) ID", serverID, name, resourceGroupName)
    }
    d.SetId(*resp.ID)

    return resourceArmRegisteredServerRead(d, meta)
}

func resourceArmRegisteredServerRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).registeredServersClient
    ctx, cancel := timeouts.ForRead(meta.(*ArmClient).StopContext, d)
    defer cancel()

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroupName := id.ResourceGroup
    name := id.Path["storageSyncServices"]
    serverID := id.Path["registeredServers"]

    resp, err := client.Get(ctx, resourceGroupName, name, serverID)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Registered Server %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Registered Server (Server %q / Storage Sync Service Name %q / Resource Group %q): %+v", serverID, name, resourceGroupName, err)
    }


    d.Set("resource_group", resourceGroupName)
    if registeredServerProperties := resp.RegisteredServerProperties; registeredServerProperties != nil {
        d.Set("agent_version", registeredServerProperties.AgentVersion)
        d.Set("cluster_id", registeredServerProperties.ClusterID)
        d.Set("cluster_name", registeredServerProperties.ClusterName)
        d.Set("last_heart_beat", registeredServerProperties.LastHeartBeat)
        d.Set("last_workflow_id", registeredServerProperties.LastWorkflowID)
        d.Set("provisioning_state", registeredServerProperties.ProvisioningState)
        d.Set("server_certificate", registeredServerProperties.ServerCertificate)
        d.Set("server_id", registeredServerProperties.ServerID)
        d.Set("server_managementt_error_code", registeredServerProperties.ServerManagementtErrorCode)
        d.Set("server_os_version", registeredServerProperties.ServerOSVersion)
        d.Set("server_role", registeredServerProperties.ServerRole)
        d.Set("storage_sync_service_uid", registeredServerProperties.StorageSyncServiceUID)
    }
    d.Set("id", resp.ID)
    d.Set("name", resp.Name)
    d.Set("server_id", serverID)
    d.Set("storage_sync_service_name", name)
    d.Set("type", resp.Type)

    return nil
}


func resourceArmRegisteredServerDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).registeredServersClient
    ctx, cancel := timeouts.ForDelete(meta.(*ArmClient).StopContext, d)
    defer cancel()


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroupName := id.ResourceGroup
    name := id.Path["storageSyncServices"]
    serverID := id.Path["registeredServers"]

    future, err := client.Delete(ctx, resourceGroupName, name, serverID)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Registered Server (Server %q / Storage Sync Service Name %q / Resource Group %q): %+v", serverID, name, resourceGroupName, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Registered Server (Server %q / Storage Sync Service Name %q / Resource Group %q): %+v", serverID, name, resourceGroupName, err)
        }
    }

    return nil
}
