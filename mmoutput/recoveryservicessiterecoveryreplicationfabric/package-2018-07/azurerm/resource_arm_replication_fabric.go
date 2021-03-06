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



func resourceArmReplicationFabric() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmReplicationFabricCreateUpdate,
        Read: resourceArmReplicationFabricRead,
        Update: resourceArmReplicationFabricCreateUpdate,
        Delete: resourceArmReplicationFabricDelete,

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

            "resource_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "renew_certificate_type": {
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

func resourceArmReplicationFabricCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).replicationFabricsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    resourceName := d.Get("resource_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, resourceName, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Replication Fabric %q (Resource Group %q / Resource Name %q): %+v", name, resourceGroup, resourceName, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_replication_fabric", *existing.ID)
        }
    }

    renewCertificateType := d.Get("renew_certificate_type").(string)

    input := recoveryservicessiterecovery.FabricCreationInput{
        RenewCertificateInputProperties: &recoveryservicessiterecovery.RenewCertificateInputProperties{
            RenewCertificateType: utils.String(renewCertificateType),
        },
    }


    future, err := client.Create(ctx, resourceGroup, resourceName, name, input)
    if err != nil {
        return fmt.Errorf("Error creating Replication Fabric %q (Resource Group %q / Resource Name %q): %+v", name, resourceGroup, resourceName, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Replication Fabric %q (Resource Group %q / Resource Name %q): %+v", name, resourceGroup, resourceName, err)
    }


    resp, err := client.Get(ctx, resourceGroup, resourceName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Replication Fabric %q (Resource Group %q / Resource Name %q): %+v", name, resourceGroup, resourceName, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Replication Fabric %q (Resource Group %q / Resource Name %q) ID", name, resourceGroup, resourceName)
    }
    d.SetId(*resp.ID)

    return resourceArmReplicationFabricRead(d, meta)
}

func resourceArmReplicationFabricRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).replicationFabricsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    resourceName := id.Path["vaults"]
    name := id.Path["replicationFabrics"]

    resp, err := client.Get(ctx, resourceGroup, resourceName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Replication Fabric %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Replication Fabric %q (Resource Group %q / Resource Name %q): %+v", name, resourceGroup, resourceName, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    d.Set("resource_name", resourceName)
    d.Set("type", resp.Type)

    return nil
}


func resourceArmReplicationFabricDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).replicationFabricsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    resourceName := id.Path["vaults"]
    name := id.Path["replicationFabrics"]

    future, err := client.Delete(ctx, resourceGroup, resourceName, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Replication Fabric %q (Resource Group %q / Resource Name %q): %+v", name, resourceGroup, resourceName, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Replication Fabric %q (Resource Group %q / Resource Name %q): %+v", name, resourceGroup, resourceName, err)
        }
    }

    return nil
}
