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



func resourceArmPeeringServicePrefixe() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmPeeringServicePrefixeCreateUpdate,
        Read: resourceArmPeeringServicePrefixeRead,
        Update: resourceArmPeeringServicePrefixeCreateUpdate,
        Delete: resourceArmPeeringServicePrefixeDelete,

        Importer: &schema.ResourceImporter{
            State: schema.ImportStatePassthrough,
        },


        Schema: map[string]*schema.Schema{
            "name": {
                Type: schema.TypeString,
                Computed: true,
            },

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "peering_service_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "prefix_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "learned_type": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(peering.None),
                    string(peering.ViaPartner),
                    string(peering.ViaSession),
                }, false),
                Default: string(peering.None),
            },

            "prefix": {
                Type: schema.TypeString,
                Optional: true,
            },

            "prefix_validation_state": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(peering.None),
                    string(peering.Invalid),
                    string(peering.Verified),
                    string(peering.Failed),
                    string(peering.Pending),
                    string(peering.Unknown),
                }, false),
                Default: string(peering.None),
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

func resourceArmPeeringServicePrefixeCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).peeringServicePrefixesClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    peeringServiceName := d.Get("peering_service_name").(string)
    prefixName := d.Get("prefix_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, peeringServiceName, prefixName)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Peering Service Prefixe (Prefix Name %q / Peering Service Name %q / Resource Group %q): %+v", prefixName, peeringServiceName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_peering_service_prefixe", *existing.ID)
        }
    }

    learnedType := d.Get("learned_type").(string)
    prefix := d.Get("prefix").(string)
    prefixValidationState := d.Get("prefix_validation_state").(string)

    peeringServicePrefix := peering.ServicePrefix{
        ServicePrefixProperties: &peering.ServicePrefixProperties{
            LearnedType: peering.LearnedType(learnedType),
            Prefix: utils.String(prefix),
            PrefixValidationState: peering.PrefixValidationState(prefixValidationState),
        },
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroup, peeringServiceName, prefixName, peeringServicePrefix); err != nil {
        return fmt.Errorf("Error creating Peering Service Prefixe (Prefix Name %q / Peering Service Name %q / Resource Group %q): %+v", prefixName, peeringServiceName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, peeringServiceName, prefixName)
    if err != nil {
        return fmt.Errorf("Error retrieving Peering Service Prefixe (Prefix Name %q / Peering Service Name %q / Resource Group %q): %+v", prefixName, peeringServiceName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Peering Service Prefixe (Prefix Name %q / Peering Service Name %q / Resource Group %q) ID", prefixName, peeringServiceName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmPeeringServicePrefixeRead(d, meta)
}

func resourceArmPeeringServicePrefixeRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).peeringServicePrefixesClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    peeringServiceName := id.Path["peeringServices"]
    prefixName := id.Path["prefixes"]

    resp, err := client.Get(ctx, resourceGroup, peeringServiceName, prefixName)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Peering Service Prefixe %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Peering Service Prefixe (Prefix Name %q / Peering Service Name %q / Resource Group %q): %+v", prefixName, peeringServiceName, resourceGroup, err)
    }


    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if servicePrefixProperties := resp.ServicePrefixProperties; servicePrefixProperties != nil {
        d.Set("learned_type", string(servicePrefixProperties.LearnedType))
        d.Set("prefix", servicePrefixProperties.Prefix)
        d.Set("prefix_validation_state", string(servicePrefixProperties.PrefixValidationState))
        d.Set("provisioning_state", string(servicePrefixProperties.ProvisioningState))
    }
    d.Set("peering_service_name", peeringServiceName)
    d.Set("prefix_name", prefixName)
    d.Set("type", resp.Type)

    return nil
}


func resourceArmPeeringServicePrefixeDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).peeringServicePrefixesClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    peeringServiceName := id.Path["peeringServices"]
    prefixName := id.Path["prefixes"]

    if _, err := client.Delete(ctx, resourceGroup, peeringServiceName, prefixName); err != nil {
        return fmt.Errorf("Error deleting Peering Service Prefixe (Prefix Name %q / Peering Service Name %q / Resource Group %q): %+v", prefixName, peeringServiceName, resourceGroup, err)
    }

    return nil
}
