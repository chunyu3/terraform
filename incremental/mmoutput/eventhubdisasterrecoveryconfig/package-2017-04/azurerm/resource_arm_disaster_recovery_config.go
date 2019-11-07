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



func resourceArmDisasterRecoveryConfig() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmDisasterRecoveryConfigCreateUpdate,
        Read: resourceArmDisasterRecoveryConfigRead,
        Update: resourceArmDisasterRecoveryConfigCreateUpdate,
        Delete: resourceArmDisasterRecoveryConfigDelete,

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

            "alias": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "alternate_name": {
                Type: schema.TypeString,
                Optional: true,
            },

            "partner_namespace": {
                Type: schema.TypeString,
                Optional: true,
            },

            "pending_replication_operations_count": {
                Type: schema.TypeInt,
                Computed: true,
            },

            "provisioning_state": {
                Type: schema.TypeString,
                Computed: true,
            },

            "role": {
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

func resourceArmDisasterRecoveryConfigCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).disasterRecoveryConfigsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    alias := d.Get("alias").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, name, alias)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Disaster Recovery Config %q (Alias %q / Resource Group %q): %+v", name, alias, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_disaster_recovery_config", *existing.ID)
        }
    }

    alternateName := d.Get("alternate_name").(string)
    partnerNamespace := d.Get("partner_namespace").(string)

    parameters := eventhub.ArmDisasterRecovery{
        ArmDisasterRecovery_properties: &eventhub.ArmDisasterRecovery_properties{
            AlternateName: utils.String(alternateName),
            PartnerNamespace: utils.String(partnerNamespace),
        },
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroup, name, alias, parameters); err != nil {
        return fmt.Errorf("Error creating Disaster Recovery Config %q (Alias %q / Resource Group %q): %+v", name, alias, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, name, alias)
    if err != nil {
        return fmt.Errorf("Error retrieving Disaster Recovery Config %q (Alias %q / Resource Group %q): %+v", name, alias, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Disaster Recovery Config %q (Alias %q / Resource Group %q) ID", name, alias, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmDisasterRecoveryConfigRead(d, meta)
}

func resourceArmDisasterRecoveryConfigRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).disasterRecoveryConfigsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["namespaces"]
    alias := id.Path["disasterRecoveryConfigs"]

    resp, err := client.Get(ctx, resourceGroup, name, alias)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Disaster Recovery Config %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Disaster Recovery Config %q (Alias %q / Resource Group %q): %+v", name, alias, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    d.Set("alias", alias)
    if armDisasterRecoveryProperties := resp.ArmDisasterRecovery_properties; armDisasterRecoveryProperties != nil {
        d.Set("alternate_name", armDisasterRecoveryProperties.AlternateName)
        d.Set("partner_namespace", armDisasterRecoveryProperties.PartnerNamespace)
        d.Set("pending_replication_operations_count", int(*armDisasterRecoveryProperties.PendingReplicationOperationsCount))
        d.Set("provisioning_state", string(armDisasterRecoveryProperties.ProvisioningState))
        d.Set("role", string(armDisasterRecoveryProperties.Role))
    }
    d.Set("type", resp.Type)

    return nil
}


func resourceArmDisasterRecoveryConfigDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).disasterRecoveryConfigsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["namespaces"]
    alias := id.Path["disasterRecoveryConfigs"]

    if _, err := client.Delete(ctx, resourceGroup, name, alias); err != nil {
        return fmt.Errorf("Error deleting Disaster Recovery Config %q (Alias %q / Resource Group %q): %+v", name, alias, resourceGroup, err)
    }

    return nil
}