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



func resourceArmProject() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmProjectCreate,
        Read: resourceArmProjectRead,
        Update: resourceArmProjectUpdate,
        Delete: resourceArmProjectDelete,

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

            "account_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "friendly_name": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "workspace_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "description": {
                Type: schema.TypeString,
                Optional: true,
            },

            "gitrepo": {
                Type: schema.TypeString,
                Optional: true,
            },

            "account_id": {
                Type: schema.TypeString,
                Computed: true,
            },

            "creation_date": {
                Type: schema.TypeString,
                Computed: true,
            },

            "project_id": {
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

            "workspace_id": {
                Type: schema.TypeString,
                Computed: true,
            },

            "tags": tags.Schema(),
        },
    }
}

func resourceArmProjectCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).projectsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    accountName := d.Get("account_name").(string)
    workspaceName := d.Get("workspace_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, accountName, workspaceName, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Project %q (Workspace Name %q / Account Name %q / Resource Group %q): %+v", name, workspaceName, accountName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_project", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    description := d.Get("description").(string)
    friendlyName := d.Get("friendly_name").(string)
    gitrepo := d.Get("gitrepo").(string)
    t := d.Get("tags").(map[string]interface{})

    parameters := machinelearningexperimentation.Project{
        Location: utils.String(location),
        ProjectProperties: &machinelearningexperimentation.ProjectProperties{
            Description: utils.String(description),
            FriendlyName: utils.String(friendlyName),
            Gitrepo: utils.String(gitrepo),
        },
        Tags: tags.Expand(t),
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroup, accountName, workspaceName, name, parameters); err != nil {
        return fmt.Errorf("Error creating Project %q (Workspace Name %q / Account Name %q / Resource Group %q): %+v", name, workspaceName, accountName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, accountName, workspaceName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Project %q (Workspace Name %q / Account Name %q / Resource Group %q): %+v", name, workspaceName, accountName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Project %q (Workspace Name %q / Account Name %q / Resource Group %q) ID", name, workspaceName, accountName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmProjectRead(d, meta)
}

func resourceArmProjectRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).projectsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    accountName := id.Path["accounts"]
    workspaceName := id.Path["workspaces"]
    name := id.Path["projects"]

    resp, err := client.Get(ctx, resourceGroup, accountName, workspaceName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Project %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Project %q (Workspace Name %q / Account Name %q / Resource Group %q): %+v", name, workspaceName, accountName, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if projectProperties := resp.ProjectProperties; projectProperties != nil {
        d.Set("account_id", projectProperties.AccountID)
        d.Set("creation_date", (projectProperties.CreationDate).String())
        d.Set("description", projectProperties.Description)
        d.Set("friendly_name", projectProperties.FriendlyName)
        d.Set("gitrepo", projectProperties.Gitrepo)
        d.Set("project_id", projectProperties.ProjectID)
        d.Set("provisioning_state", string(projectProperties.ProvisioningState))
        d.Set("workspace_id", projectProperties.WorkspaceID)
    }
    d.Set("account_name", accountName)
    d.Set("type", resp.Type)
    d.Set("workspace_name", workspaceName)

    return tags.FlattenAndSet(d, resp.Tags)
}

func resourceArmProjectUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).projectsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    accountName := d.Get("account_name").(string)
    description := d.Get("description").(string)
    friendlyName := d.Get("friendly_name").(string)
    gitrepo := d.Get("gitrepo").(string)
    workspaceName := d.Get("workspace_name").(string)
    t := d.Get("tags").(map[string]interface{})

    parameters := machinelearningexperimentation.Project{
        Location: utils.String(location),
        ProjectProperties: &machinelearningexperimentation.ProjectProperties{
            Description: utils.String(description),
            FriendlyName: utils.String(friendlyName),
            Gitrepo: utils.String(gitrepo),
        },
        Tags: tags.Expand(t),
    }


    if _, err := client.Update(ctx, resourceGroup, accountName, workspaceName, name, parameters); err != nil {
        return fmt.Errorf("Error updating Project %q (Workspace Name %q / Account Name %q / Resource Group %q): %+v", name, workspaceName, accountName, resourceGroup, err)
    }

    return resourceArmProjectRead(d, meta)
}

func resourceArmProjectDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).projectsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    accountName := id.Path["accounts"]
    workspaceName := id.Path["workspaces"]
    name := id.Path["projects"]

    if _, err := client.Delete(ctx, resourceGroup, accountName, workspaceName, name); err != nil {
        return fmt.Errorf("Error deleting Project %q (Workspace Name %q / Account Name %q / Resource Group %q): %+v", name, workspaceName, accountName, resourceGroup, err)
    }

    return nil
}
