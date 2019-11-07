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

            "location": azure.SchemaLocation(),

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "client_id": {
                Type: schema.TypeString,
                Computed: true,
            },

            "client_secret_url": {
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

            "tags": tags.Schema(),
        },
    }
}

func resourceArmUserAssignedIdentityCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).userAssignedIdentitiesClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing User Assigned Identity %q (Resource Group %q): %+v", name, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_user_assigned_identity", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    t := d.Get("tags").(map[string]interface{})

    parameters := managedserviceidentity.Identity{
        Location: utils.String(location),
        Tags: tags.Expand(t),
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroup, name, parameters); err != nil {
        return fmt.Errorf("Error creating User Assigned Identity %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        return fmt.Errorf("Error retrieving User Assigned Identity %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read User Assigned Identity %q (Resource Group %q) ID", name, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmUserAssignedIdentityRead(d, meta)
}

func resourceArmUserAssignedIdentityRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).userAssignedIdentitiesClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["userAssignedIdentities"]

    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] User Assigned Identity %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading User Assigned Identity %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if identityProperties := resp.IdentityProperties; identityProperties != nil {
        d.Set("client_id", identityProperties.ClientID)
        d.Set("client_secret_url", identityProperties.ClientSecretURL)
        d.Set("principal_id", identityProperties.PrincipalID)
        d.Set("tenant_id", identityProperties.TenantID)
    }
    d.Set("type", string(resp.Type))

    return tags.FlattenAndSet(d, resp.Tags)
}

func resourceArmUserAssignedIdentityUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).userAssignedIdentitiesClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    t := d.Get("tags").(map[string]interface{})

    parameters := managedserviceidentity.Identity{
        Location: utils.String(location),
        Tags: tags.Expand(t),
    }


    if _, err := client.Update(ctx, resourceGroup, name, parameters); err != nil {
        return fmt.Errorf("Error updating User Assigned Identity %q (Resource Group %q): %+v", name, resourceGroup, err)
    }

    return resourceArmUserAssignedIdentityRead(d, meta)
}

func resourceArmUserAssignedIdentityDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).userAssignedIdentitiesClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["userAssignedIdentities"]

    if _, err := client.Delete(ctx, resourceGroup, name); err != nil {
        return fmt.Errorf("Error deleting User Assigned Identity %q (Resource Group %q): %+v", name, resourceGroup, err)
    }

    return nil
}