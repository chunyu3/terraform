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



func resourceArmAssociation() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmAssociationCreateUpdate,
        Read: resourceArmAssociationRead,
        Update: resourceArmAssociationCreateUpdate,
        Delete: resourceArmAssociationDelete,

        Importer: &schema.ResourceImporter{
            State: schema.ImportStatePassthrough,
        },


        Schema: map[string]*schema.Schema{
            "name": {
                Type: schema.TypeString,
                Computed: true,
            },

            "association_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "scope": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "target_resource_id": {
                Type: schema.TypeString,
                Optional: true,
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

func resourceArmAssociationCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).associationsClient
    ctx := meta.(*ArmClient).StopContext

    associationName := d.Get("association_name").(string)
    scope := d.Get("scope").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, scope, associationName)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Association (Association Name %q / Scope %q): %+v", associationName, scope, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_association", *existing.ID)
        }
    }

    targetResourceId := d.Get("target_resource_id").(string)

    association := customproviders.Association{
        Association_properties: &customproviders.Association_properties{
            TargetResourceID: utils.String(targetResourceId),
        },
    }


    future, err := client.CreateOrUpdate(ctx, scope, associationName, association)
    if err != nil {
        return fmt.Errorf("Error creating Association (Association Name %q / Scope %q): %+v", associationName, scope, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Association (Association Name %q / Scope %q): %+v", associationName, scope, err)
    }


    resp, err := client.Get(ctx, scope, associationName)
    if err != nil {
        return fmt.Errorf("Error retrieving Association (Association Name %q / Scope %q): %+v", associationName, scope, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Association (Association Name %q / Scope %q) ID", associationName, scope)
    }
    d.SetId(*resp.ID)

    return resourceArmAssociationRead(d, meta)
}

func resourceArmAssociationRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).associationsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    associationName := id.Path["associations"]

    resp, err := client.Get(ctx, scope, associationName)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Association %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Association (Association Name %q / Scope %q): %+v", associationName, scope, err)
    }


    d.Set("name", resp.Name)
    d.Set("association_name", associationName)
    if associationProperties := resp.Association_properties; associationProperties != nil {
        d.Set("provisioning_state", string(associationProperties.ProvisioningState))
        d.Set("target_resource_id", associationProperties.TargetResourceID)
    }
    d.Set("scope", scope)
    d.Set("type", resp.Type)

    return nil
}


func resourceArmAssociationDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).associationsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    associationName := id.Path["associations"]

    future, err := client.Delete(ctx, scope, associationName)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Association (Association Name %q / Scope %q): %+v", associationName, scope, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Association (Association Name %q / Scope %q): %+v", associationName, scope, err)
        }
    }

    return nil
}