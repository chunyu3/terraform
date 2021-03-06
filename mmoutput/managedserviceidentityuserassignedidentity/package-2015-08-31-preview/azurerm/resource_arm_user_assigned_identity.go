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



func resourceArmUserAssignedIdentity() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmUserAssignedIdentityCreate,
        Read: resourceArmUserAssignedIdentityRead,
        Update: resourceArmUserAssignedIdentityUpdate,
        Delete: resourceArmUserAssignedIdentityDelete,

        Importer: &schema.ResourceImporter{
            State: schema.ImportStatePassthrough,
        },


        Schema: map[string]*schema.Schema{
            "location": azure.SchemaLocation(),

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "resource_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "tags": tags.Schema(),

            "client_id": {
                Type: schema.TypeString,
                Computed: true,
            },

            "client_secret_url": {
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

            "principal_id": {
                Type: schema.TypeString,
                Computed: true,
            },

            "tenant_id": {
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

func resourceArmUserAssignedIdentityCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).userAssignedIdentitiesClient
    ctx, cancel := timeouts.ForCreate(meta.(*ArmClient).StopContext, d)
    defer cancel()

    resourceGroupName := d.Get("resource_group").(string)
    name := d.Get("resource_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroupName, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing User Assigned Identity (Resource Name %q / Resource Group %q): %+v", name, resourceGroupName, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_user_assigned_identity", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    tags := d.Get("tags").(map[string]interface{})

    parameters := managedserviceidentity.Identity{
        Location: utils.String(location),
        Tags: tags.Expand(tags),
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroupName, name, parameters); err != nil {
        return fmt.Errorf("Error creating User Assigned Identity (Resource Name %q / Resource Group %q): %+v", name, resourceGroupName, err)
    }


    resp, err := client.Get(ctx, resourceGroupName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving User Assigned Identity (Resource Name %q / Resource Group %q): %+v", name, resourceGroupName, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read User Assigned Identity (Resource Name %q / Resource Group %q) ID", name, resourceGroupName)
    }
    d.SetId(*resp.ID)

    return resourceArmUserAssignedIdentityRead(d, meta)
}

func resourceArmUserAssignedIdentityRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).userAssignedIdentitiesClient
    ctx, cancel := timeouts.ForRead(meta.(*ArmClient).StopContext, d)
    defer cancel()

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroupName := id.ResourceGroup
    name := id.Path["userAssignedIdentities"]

    resp, err := client.Get(ctx, resourceGroupName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] User Assigned Identity %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading User Assigned Identity (Resource Name %q / Resource Group %q): %+v", name, resourceGroupName, err)
    }


    d.Set("resource_group", resourceGroupName)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if identityProperties := resp.IdentityProperties; identityProperties != nil {
        d.Set("client_id", identityProperties.ClientID)
        d.Set("client_secret_url", identityProperties.ClientSecretURL)
        d.Set("principal_id", identityProperties.PrincipalID)
        d.Set("tenant_id", identityProperties.TenantID)
    }
    d.Set("id", resp.ID)
    d.Set("name", resp.Name)
    d.Set("resource_name", name)
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}

func resourceArmUserAssignedIdentityUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).userAssignedIdentitiesClient
    ctx, cancel := timeouts.ForUpdate(meta.(*ArmClient).StopContext, d)
    defer cancel()

      resourceGroupName := d.Get("resource_group").(string)
    location := azure.NormalizeLocation(d.Get("location").(string))
    name := d.Get("resource_name").(string)
    tags := d.Get("tags").(map[string]interface{})

    parameters := managedserviceidentity.Identity{
        Location: utils.String(location),
        Tags: tags.Expand(tags),
    }


    if _, err := client.Update(ctx, resourceGroupName, name, parameters); err != nil {
        return fmt.Errorf("Error updating User Assigned Identity (Resource Name %q / Resource Group %q): %+v", name, resourceGroupName, err)
    }

    return resourceArmUserAssignedIdentityRead(d, meta)
}

func resourceArmUserAssignedIdentityDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).userAssignedIdentitiesClient
    ctx, cancel := timeouts.ForDelete(meta.(*ArmClient).StopContext, d)
    defer cancel()


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroupName := id.ResourceGroup
    name := id.Path["userAssignedIdentities"]

    if _, err := client.Delete(ctx, resourceGroupName, name); err != nil {
        return fmt.Errorf("Error deleting User Assigned Identity (Resource Name %q / Resource Group %q): %+v", name, resourceGroupName, err)
    }

    return nil
}
