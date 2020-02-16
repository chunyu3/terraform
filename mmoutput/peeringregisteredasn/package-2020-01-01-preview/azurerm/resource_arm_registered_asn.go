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



func resourceArmRegisteredAsn() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmRegisteredAsnCreateUpdate,
        Read: resourceArmRegisteredAsnRead,
        Update: resourceArmRegisteredAsnCreateUpdate,
        Delete: resourceArmRegisteredAsnDelete,

        Importer: &schema.ResourceImporter{
            State: schema.ImportStatePassthrough,
        },


        Schema: map[string]*schema.Schema{
            "peering_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "registered_asn_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "asn": {
                Type: schema.TypeInt,
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

            "peering_service_prefix_key": {
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

func resourceArmRegisteredAsnCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).registeredAsnsClient
    ctx, cancel := timeouts.ForCreateUpdate(meta.(*ArmClient).StopContext, d)
    defer cancel()

    resourceGroupName := d.Get("resource_group").(string)
    peeringName := d.Get("peering_name").(string)
    name := d.Get("registered_asn_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroupName, peeringName, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Registered Asn (Registered Asn Name %q / Peering Name %q / Resource Group %q): %+v", name, peeringName, resourceGroupName, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_registered_asn", *existing.ID)
        }
    }

    asn := d.Get("asn").(int)

    registeredAsn := peering.RegisteredAsn{
        RegisteredAsnProperties: &peering.RegisteredAsnProperties{
            Asn: utils.Int32(int32(asn)),
        },
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroupName, peeringName, name, registeredAsn); err != nil {
        return fmt.Errorf("Error creating Registered Asn (Registered Asn Name %q / Peering Name %q / Resource Group %q): %+v", name, peeringName, resourceGroupName, err)
    }


    resp, err := client.Get(ctx, resourceGroupName, peeringName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Registered Asn (Registered Asn Name %q / Peering Name %q / Resource Group %q): %+v", name, peeringName, resourceGroupName, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Registered Asn (Registered Asn Name %q / Peering Name %q / Resource Group %q) ID", name, peeringName, resourceGroupName)
    }
    d.SetId(*resp.ID)

    return resourceArmRegisteredAsnRead(d, meta)
}

func resourceArmRegisteredAsnRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).registeredAsnsClient
    ctx, cancel := timeouts.ForRead(meta.(*ArmClient).StopContext, d)
    defer cancel()

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroupName := id.ResourceGroup
    peeringName := id.Path["peerings"]
    name := id.Path["registeredAsns"]

    resp, err := client.Get(ctx, resourceGroupName, peeringName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Registered Asn %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Registered Asn (Registered Asn Name %q / Peering Name %q / Resource Group %q): %+v", name, peeringName, resourceGroupName, err)
    }


    d.Set("resource_group", resourceGroupName)
    if registeredAsnProperties := resp.RegisteredAsnProperties; registeredAsnProperties != nil {
        d.Set("asn", int(*registeredAsnProperties.Asn))
        d.Set("peering_service_prefix_key", registeredAsnProperties.PeeringServicePrefixKey)
        d.Set("provisioning_state", string(registeredAsnProperties.ProvisioningState))
    }
    d.Set("id", resp.ID)
    d.Set("name", resp.Name)
    d.Set("peering_name", peeringName)
    d.Set("registered_asn_name", name)
    d.Set("type", resp.Type)

    return nil
}


func resourceArmRegisteredAsnDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).registeredAsnsClient
    ctx, cancel := timeouts.ForDelete(meta.(*ArmClient).StopContext, d)
    defer cancel()


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroupName := id.ResourceGroup
    peeringName := id.Path["peerings"]
    name := id.Path["registeredAsns"]

    if _, err := client.Delete(ctx, resourceGroupName, peeringName, name); err != nil {
        return fmt.Errorf("Error deleting Registered Asn (Registered Asn Name %q / Peering Name %q / Resource Group %q): %+v", name, peeringName, resourceGroupName, err)
    }

    return nil
}
