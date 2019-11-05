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
                Computed: true,
            },

            "location": azure.SchemaLocation(),

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "workspace_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "creation_time": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validateRFC3339Date,
            },

            "provisioning_state": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(batchaisdk.creating),
                    string(batchaisdk.succeeded),
                    string(batchaisdk.failed),
                    string(batchaisdk.deleting),
                }, false),
                Default: string(batchaisdk.creating),
            },

            "provisioning_state_transition_time": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validateRFC3339Date,
            },

            "type": {
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

    resourceGroup := d.Get("resource_group").(string)
    workspaceName := d.Get("workspace_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, workspaceName)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Workspace (Workspace Name %q / Resource Group %q): %+v", workspaceName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_workspace", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    creationTime := d.Get("creation_time").(string)
    provisioningState := d.Get("provisioning_state").(string)
    provisioningStateTransitionTime := d.Get("provisioning_state_transition_time").(string)
    t := d.Get("tags").(map[string]interface{})

    parameters := batchaisdk.WorkspaceCreateParameters{
        Location: utils.String(location),
        // TODO: SDK Reference /properties is not supported
        Tags: tags.Expand(t),
    }


    future, err := client.Create(ctx, resourceGroup, workspaceName, parameters)
    if err != nil {
        return fmt.Errorf("Error creating Workspace (Workspace Name %q / Resource Group %q): %+v", workspaceName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Workspace (Workspace Name %q / Resource Group %q): %+v", workspaceName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, workspaceName)
    if err != nil {
        return fmt.Errorf("Error retrieving Workspace (Workspace Name %q / Resource Group %q): %+v", workspaceName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Workspace (Workspace Name %q / Resource Group %q) ID", workspaceName, resourceGroup)
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
    workspaceName := id.Path["workspaces"]

    resp, err := client.Get(ctx, resourceGroup, workspaceName)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Workspace %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Workspace (Workspace Name %q / Resource Group %q): %+v", workspaceName, resourceGroup, err)
    }


    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    d.Set("type", resp.Type)
    d.Set("workspace_name", workspaceName)

    return tags.FlattenAndSet(d, resp.Tags)
}

func resourceArmWorkspaceUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).workspacesClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    creationTime := d.Get("creation_time").(string)
    provisioningState := d.Get("provisioning_state").(string)
    provisioningStateTransitionTime := d.Get("provisioning_state_transition_time").(string)
    workspaceName := d.Get("workspace_name").(string)
    t := d.Get("tags").(map[string]interface{})

    parameters := batchaisdk.WorkspaceCreateParameters{
        Location: utils.String(location),
        // TODO: SDK Reference /properties is not supported
        Tags: tags.Expand(t),
    }


    if _, err := client.Update(ctx, resourceGroup, workspaceName, parameters); err != nil {
        return fmt.Errorf("Error updating Workspace (Workspace Name %q / Resource Group %q): %+v", workspaceName, resourceGroup, err)
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
    workspaceName := id.Path["workspaces"]

    future, err := client.Delete(ctx, resourceGroup, workspaceName)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Workspace (Workspace Name %q / Resource Group %q): %+v", workspaceName, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Workspace (Workspace Name %q / Resource Group %q): %+v", workspaceName, resourceGroup, err)
        }
    }

    return nil
}
