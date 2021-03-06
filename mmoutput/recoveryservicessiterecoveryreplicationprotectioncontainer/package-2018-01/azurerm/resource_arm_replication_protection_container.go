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



func resourceArmReplicationProtectionContainer() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmReplicationProtectionContainerCreateUpdate,
        Read: resourceArmReplicationProtectionContainerRead,
        Update: resourceArmReplicationProtectionContainerCreateUpdate,
        Delete: resourceArmReplicationProtectionContainerDelete,

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

            "resource_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "replication_protected_item_name": {
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

func resourceArmReplicationProtectionContainerCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).replicationProtectionContainersClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    fabricName := d.Get("fabric_name").(string)
    resourceName := d.Get("resource_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, resourceName, fabricName, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Replication Protection Container %q (Fabric Name %q / Resource Group %q / Resource Name %q): %+v", name, fabricName, resourceGroup, resourceName, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_replication_protection_container", *existing.ID)
        }
    }

    replicationProtectedItemName := d.Get("replication_protected_item_name").(string)

    creationInput := recoveryservicessiterecovery.CreateProtectionContainerInput{
        SwitchProtectionInputProperties: &recoveryservicessiterecovery.SwitchProtectionInputProperties{
            ReplicationProtectedItemName: utils.String(replicationProtectedItemName),
        },
    }


    future, err := client.Create(ctx, resourceGroup, resourceName, fabricName, name, creationInput)
    if err != nil {
        return fmt.Errorf("Error creating Replication Protection Container %q (Fabric Name %q / Resource Group %q / Resource Name %q): %+v", name, fabricName, resourceGroup, resourceName, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Replication Protection Container %q (Fabric Name %q / Resource Group %q / Resource Name %q): %+v", name, fabricName, resourceGroup, resourceName, err)
    }


    resp, err := client.Get(ctx, resourceGroup, resourceName, fabricName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Replication Protection Container %q (Fabric Name %q / Resource Group %q / Resource Name %q): %+v", name, fabricName, resourceGroup, resourceName, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Replication Protection Container %q (Fabric Name %q / Resource Group %q / Resource Name %q) ID", name, fabricName, resourceGroup, resourceName)
    }
    d.SetId(*resp.ID)

    return resourceArmReplicationProtectionContainerRead(d, meta)
}

func resourceArmReplicationProtectionContainerRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).replicationProtectionContainersClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    resourceName := id.Path["vaults"]
    fabricName := id.Path["replicationFabrics"]
    name := id.Path["replicationProtectionContainers"]

    resp, err := client.Get(ctx, resourceGroup, resourceName, fabricName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Replication Protection Container %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Replication Protection Container %q (Fabric Name %q / Resource Group %q / Resource Name %q): %+v", name, fabricName, resourceGroup, resourceName, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    d.Set("fabric_name", fabricName)
    d.Set("resource_name", resourceName)
    d.Set("type", resp.Type)

    return nil
}


func resourceArmReplicationProtectionContainerDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).replicationProtectionContainersClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    resourceName := id.Path["vaults"]
    fabricName := id.Path["replicationFabrics"]
    name := id.Path["replicationProtectionContainers"]

    future, err := client.Delete(ctx, resourceGroup, resourceName, fabricName, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Replication Protection Container %q (Fabric Name %q / Resource Group %q / Resource Name %q): %+v", name, fabricName, resourceGroup, resourceName, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Replication Protection Container %q (Fabric Name %q / Resource Group %q / Resource Name %q): %+v", name, fabricName, resourceGroup, resourceName, err)
        }
    }

    return nil
}
