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



func resourceArmWorkspace() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmWorkspaceCreate,
        Read: resourceArmWorkspaceRead,
        Update: resourceArmWorkspaceUpdate,
        Delete: resourceArmWorkspaceDelete,

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

            "key_vault_identifier_id": {
                Type: schema.TypeString,
                Optional: true,
            },

            "workspace_state": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(machinelearning.Deleted),
                    string(machinelearning.Enabled),
                    string(machinelearning.Disabled),
                    string(machinelearning.Migrated),
                    string(machinelearning.Updated),
                    string(machinelearning.Registered),
                    string(machinelearning.Unregistered),
                }, false),
                Default: string(machinelearning.Deleted),
            },

            "creation_time": {
                Type: schema.TypeString,
                Computed: true,
            },

            "owner_email": {
                Type: schema.TypeString,
                Computed: true,
            },

            "studio_endpoint": {
                Type: schema.TypeString,
                Computed: true,
            },

            "type": {
                Type: schema.TypeString,
                Computed: true,
            },

            "user_storage_account_id": {
                Type: schema.TypeString,
                Computed: true,
            },

            "workspace_id": {
                Type: schema.TypeString,
                Computed: true,
            },

            "workspace_type": {
                Type: schema.TypeString,
                Computed: true,
            },

            "tags": tags.Schema(),
        },
    }
}

func resourceArmWorkspaceCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).workspacesClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Workspace %q (Resource Group %q): %+v", name, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_workspace", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    keyVaultIdentifierId := d.Get("key_vault_identifier_id").(string)
    workspaceState := d.Get("workspace_state").(string)
    t := d.Get("tags").(map[string]interface{})

    parameters := machinelearning.WorkspaceUpdateParameters{
        Location: utils.String(location),
        WorkspacePropertiesUpdateParameters: &machinelearning.WorkspacePropertiesUpdateParameters{
            KeyVaultIdentifierID: utils.String(keyVaultIdentifierId),
            WorkspaceState: machinelearning.WorkspaceState(workspaceState),
        },
        Tags: tags.Expand(t),
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroup, name, parameters); err != nil {
        return fmt.Errorf("Error creating Workspace %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Workspace %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Workspace %q (Resource Group %q) ID", name, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmWorkspaceRead(d, meta)
}

func resourceArmWorkspaceRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).workspacesClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["workspaces"]

    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Workspace %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Workspace %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if workspacePropertiesUpdateParameters := resp.WorkspacePropertiesUpdateParameters; workspacePropertiesUpdateParameters != nil {
        d.Set("creation_time", workspacePropertiesUpdateParameters.CreationTime)
        d.Set("key_vault_identifier_id", workspacePropertiesUpdateParameters.KeyVaultIdentifierID)
        d.Set("owner_email", workspacePropertiesUpdateParameters.OwnerEmail)
        d.Set("studio_endpoint", workspacePropertiesUpdateParameters.StudioEndpoint)
        d.Set("user_storage_account_id", workspacePropertiesUpdateParameters.UserStorageAccountID)
        d.Set("workspace_id", workspacePropertiesUpdateParameters.WorkspaceID)
        d.Set("workspace_state", string(workspacePropertiesUpdateParameters.WorkspaceState))
        d.Set("workspace_type", string(workspacePropertiesUpdateParameters.WorkspaceType))
    }
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}

func resourceArmWorkspaceUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).workspacesClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    keyVaultIdentifierId := d.Get("key_vault_identifier_id").(string)
    workspaceState := d.Get("workspace_state").(string)
    t := d.Get("tags").(map[string]interface{})

    parameters := machinelearning.WorkspaceUpdateParameters{
        WorkspacePropertiesUpdateParameters: &machinelearning.WorkspacePropertiesUpdateParameters{
            KeyVaultIdentifierID: utils.String(keyVaultIdentifierId),
            WorkspaceState: machinelearning.WorkspaceState(workspaceState),
        },
        Tags: tags.Expand(t),
    }


    if _, err := client.Update(ctx, resourceGroup, name, parameters); err != nil {
        return fmt.Errorf("Error updating Workspace %q (Resource Group %q): %+v", name, resourceGroup, err)
    }

    return resourceArmWorkspaceRead(d, meta)
}

func resourceArmWorkspaceDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).workspacesClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["workspaces"]

    if _, err := client.Delete(ctx, resourceGroup, name); err != nil {
        return fmt.Errorf("Error deleting Workspace %q (Resource Group %q): %+v", name, resourceGroup, err)
    }

    return nil
}
