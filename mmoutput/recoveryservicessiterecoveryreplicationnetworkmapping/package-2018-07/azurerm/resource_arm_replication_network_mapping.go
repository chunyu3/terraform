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



func resourceArmReplicationNetworkMapping() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmReplicationNetworkMappingCreate,
        Read: resourceArmReplicationNetworkMappingRead,
        Update: resourceArmReplicationNetworkMappingUpdate,
        Delete: resourceArmReplicationNetworkMappingDelete,

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

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "fabric_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "network_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "resource_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "recovery_fabric_name": {
                Type: schema.TypeString,
                Optional: true,
            },

            "recovery_network_id": {
                Type: schema.TypeString,
                Optional: true,
            },

            "type": {
                Type: schema.TypeString,
                Computed: true,
            },
        },
    }
}

func resourceArmReplicationNetworkMappingCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).replicationNetworkMappingsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    fabricName := d.Get("fabric_name").(string)
    networkName := d.Get("network_name").(string)
    resourceName := d.Get("resource_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, resourceName, fabricName, networkName, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Replication Network Mapping %q (Network Name %q / Fabric Name %q / Resource Group %q / Resource Name %q): %+v", name, networkName, fabricName, resourceGroup, resourceName, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_replication_network_mapping", *existing.ID)
        }
    }

    recoveryFabricName := d.Get("recovery_fabric_name").(string)
    recoveryNetworkId := d.Get("recovery_network_id").(string)

    input := recoveryservicessiterecovery.UpdateNetworkMappingInput{
        UpdateNetworkMappingInputProperties: &recoveryservicessiterecovery.UpdateNetworkMappingInputProperties{
            RecoveryFabricName: utils.String(recoveryFabricName),
            RecoveryNetworkID: utils.String(recoveryNetworkId),
        },
    }


    future, err := client.Create(ctx, resourceGroup, resourceName, fabricName, networkName, name, input)
    if err != nil {
        return fmt.Errorf("Error creating Replication Network Mapping %q (Network Name %q / Fabric Name %q / Resource Group %q / Resource Name %q): %+v", name, networkName, fabricName, resourceGroup, resourceName, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Replication Network Mapping %q (Network Name %q / Fabric Name %q / Resource Group %q / Resource Name %q): %+v", name, networkName, fabricName, resourceGroup, resourceName, err)
    }


    resp, err := client.Get(ctx, resourceGroup, resourceName, fabricName, networkName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Replication Network Mapping %q (Network Name %q / Fabric Name %q / Resource Group %q / Resource Name %q): %+v", name, networkName, fabricName, resourceGroup, resourceName, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Replication Network Mapping %q (Network Name %q / Fabric Name %q / Resource Group %q / Resource Name %q) ID", name, networkName, fabricName, resourceGroup, resourceName)
    }
    d.SetId(*resp.ID)

    return resourceArmReplicationNetworkMappingRead(d, meta)
}

func resourceArmReplicationNetworkMappingRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).replicationNetworkMappingsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    resourceName := id.Path["vaults"]
    fabricName := id.Path["replicationFabrics"]
    networkName := id.Path["replicationNetworks"]
    name := id.Path["replicationNetworkMappings"]

    resp, err := client.Get(ctx, resourceGroup, resourceName, fabricName, networkName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Replication Network Mapping %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Replication Network Mapping %q (Network Name %q / Fabric Name %q / Resource Group %q / Resource Name %q): %+v", name, networkName, fabricName, resourceGroup, resourceName, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    d.Set("fabric_name", fabricName)
    d.Set("network_name", networkName)
    d.Set("resource_name", resourceName)
    d.Set("type", resp.Type)

    return nil
}

func resourceArmReplicationNetworkMappingUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).replicationNetworkMappingsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    fabricName := d.Get("fabric_name").(string)
    networkName := d.Get("network_name").(string)
    recoveryFabricName := d.Get("recovery_fabric_name").(string)
    recoveryNetworkId := d.Get("recovery_network_id").(string)
    resourceName := d.Get("resource_name").(string)

    input := recoveryservicessiterecovery.UpdateNetworkMappingInput{
        UpdateNetworkMappingInputProperties: &recoveryservicessiterecovery.UpdateNetworkMappingInputProperties{
            RecoveryFabricName: utils.String(recoveryFabricName),
            RecoveryNetworkID: utils.String(recoveryNetworkId),
        },
    }


    future, err := client.Update(ctx, resourceGroup, resourceName, fabricName, networkName, name, input)
    if err != nil {
        return fmt.Errorf("Error updating Replication Network Mapping %q (Network Name %q / Fabric Name %q / Resource Group %q / Resource Name %q): %+v", name, networkName, fabricName, resourceGroup, resourceName, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for update of Replication Network Mapping %q (Network Name %q / Fabric Name %q / Resource Group %q / Resource Name %q): %+v", name, networkName, fabricName, resourceGroup, resourceName, err)
    }

    return resourceArmReplicationNetworkMappingRead(d, meta)
}

func resourceArmReplicationNetworkMappingDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).replicationNetworkMappingsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    resourceName := id.Path["vaults"]
    fabricName := id.Path["replicationFabrics"]
    networkName := id.Path["replicationNetworks"]
    name := id.Path["replicationNetworkMappings"]

    future, err := client.Delete(ctx, resourceGroup, resourceName, fabricName, networkName, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Replication Network Mapping %q (Network Name %q / Fabric Name %q / Resource Group %q / Resource Name %q): %+v", name, networkName, fabricName, resourceGroup, resourceName, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Replication Network Mapping %q (Network Name %q / Fabric Name %q / Resource Group %q / Resource Name %q): %+v", name, networkName, fabricName, resourceGroup, resourceName, err)
        }
    }

    return nil
}
