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



func resourceArmRegisteredPrefix() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmRegisteredPrefixCreateUpdate,
        Read: resourceArmRegisteredPrefixRead,
        Update: resourceArmRegisteredPrefixCreateUpdate,
        Delete: resourceArmRegisteredPrefixDelete,

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

            "registered_prefix_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "prefix": {
                Type: schema.TypeString,
                Optional: true,
            },

            "error_message": {
                Type: schema.TypeString,
                Computed: true,
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

            "prefix_validation_state": {
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

func resourceArmRegisteredPrefixCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).registeredPrefixesClient
    ctx, cancel := timeouts.ForCreateUpdate(meta.(*ArmClient).StopContext, d)
    defer cancel()

    resourceGroupName := d.Get("resource_group").(string)
    peeringName := d.Get("peering_name").(string)
    name := d.Get("registered_prefix_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroupName, peeringName, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Registered Prefix (Registered Prefix Name %q / Peering Name %q / Resource Group %q): %+v", name, peeringName, resourceGroupName, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_registered_prefix", *existing.ID)
        }
    }

    prefix := d.Get("prefix").(string)

    registeredPrefix := peering.RegisteredPrefix{
        RegisteredPrefixProperties: &peering.RegisteredPrefixProperties{
            Prefix: utils.String(prefix),
        },
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroupName, peeringName, name, registeredPrefix); err != nil {
        return fmt.Errorf("Error creating Registered Prefix (Registered Prefix Name %q / Peering Name %q / Resource Group %q): %+v", name, peeringName, resourceGroupName, err)
    }


    resp, err := client.Get(ctx, resourceGroupName, peeringName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Registered Prefix (Registered Prefix Name %q / Peering Name %q / Resource Group %q): %+v", name, peeringName, resourceGroupName, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Registered Prefix (Registered Prefix Name %q / Peering Name %q / Resource Group %q) ID", name, peeringName, resourceGroupName)
    }
    d.SetId(*resp.ID)

    return resourceArmRegisteredPrefixRead(d, meta)
}

func resourceArmRegisteredPrefixRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).registeredPrefixesClient
    ctx, cancel := timeouts.ForRead(meta.(*ArmClient).StopContext, d)
    defer cancel()

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroupName := id.ResourceGroup
    peeringName := id.Path["peerings"]
    name := id.Path["registeredPrefixes"]

    resp, err := client.Get(ctx, resourceGroupName, peeringName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Registered Prefix %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Registered Prefix (Registered Prefix Name %q / Peering Name %q / Resource Group %q): %+v", name, peeringName, resourceGroupName, err)
    }


    d.Set("resource_group", resourceGroupName)
    if registeredPrefixProperties := resp.RegisteredPrefixProperties; registeredPrefixProperties != nil {
        d.Set("error_message", registeredPrefixProperties.ErrorMessage)
        d.Set("peering_service_prefix_key", registeredPrefixProperties.PeeringServicePrefixKey)
        d.Set("prefix", registeredPrefixProperties.Prefix)
        d.Set("prefix_validation_state", string(registeredPrefixProperties.PrefixValidationState))
        d.Set("provisioning_state", string(registeredPrefixProperties.ProvisioningState))
    }
    d.Set("id", resp.ID)
    d.Set("name", resp.Name)
    d.Set("peering_name", peeringName)
    d.Set("registered_prefix_name", name)
    d.Set("type", resp.Type)

    return nil
}


func resourceArmRegisteredPrefixDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).registeredPrefixesClient
    ctx, cancel := timeouts.ForDelete(meta.(*ArmClient).StopContext, d)
    defer cancel()


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroupName := id.ResourceGroup
    peeringName := id.Path["peerings"]
    name := id.Path["registeredPrefixes"]

    if _, err := client.Delete(ctx, resourceGroupName, peeringName, name); err != nil {
        return fmt.Errorf("Error deleting Registered Prefix (Registered Prefix Name %q / Peering Name %q / Resource Group %q): %+v", name, peeringName, resourceGroupName, err)
    }

    return nil
}
