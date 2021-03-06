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



func resourceArmArtifactSource() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmArtifactSourceCreate,
        Read: resourceArmArtifactSourceRead,
        Update: resourceArmArtifactSourceUpdate,
        Delete: resourceArmArtifactSourceDelete,

        Importer: &schema.ResourceImporter{
            State: schema.ImportStatePassthrough,
        },


        Schema: map[string]*schema.Schema{
            "lab_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "arm_template_folder_path": {
                Type: schema.TypeString,
                Optional: true,
            },

            "branch_ref": {
                Type: schema.TypeString,
                Optional: true,
            },

            "display_name": {
                Type: schema.TypeString,
                Optional: true,
            },

            "folder_path": {
                Type: schema.TypeString,
                Optional: true,
            },

            "location": azure.SchemaLocation(),

            "security_token": {
                Type: schema.TypeString,
                Optional: true,
            },

            "source_type": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(devtestlab.VsoGit),
                    string(devtestlab.GitHub),
                }, false),
                Default: string(devtestlab.VsoGit),
            },

            "status": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(devtestlab.Enabled),
                    string(devtestlab.Disabled),
                }, false),
                Default: string(devtestlab.Enabled),
            },

            "tags": tags.Schema(),

            "uri": {
                Type: schema.TypeString,
                Optional: true,
            },

            "unique_identifier": {
                Type: schema.TypeString,
                Optional: true,
            },

            "created_date": {
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

func resourceArmArtifactSourceCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).artifactSourcesClient
    ctx, cancel := timeouts.ForCreate(meta.(*ArmClient).StopContext, d)
    defer cancel()

    resourceGroupName := d.Get("resource_group").(string)
    name := d.Get("lab_name").(string)
    name := d.Get("name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroupName, name, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Artifact Source (Name %q / Lab Name %q / Resource Group %q): %+v", name, name, resourceGroupName, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_artifact_source", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    armTemplateFolderPath := d.Get("arm_template_folder_path").(string)
    branchRef := d.Get("branch_ref").(string)
    displayName := d.Get("display_name").(string)
    folderPath := d.Get("folder_path").(string)
    securityToken := d.Get("security_token").(string)
    sourceType := d.Get("source_type").(string)
    status := d.Get("status").(string)
    uRI := d.Get("uri").(string)
    uniqueIdentifier := d.Get("unique_identifier").(string)
    tags := d.Get("tags").(map[string]interface{})

    artifactSource := devtestlab.ArtifactSourceFragment{
        Location: utils.String(location),
        ArtifactSourcePropertiesFragment: &devtestlab.ArtifactSourcePropertiesFragment{
            ArmTemplateFolderPath: utils.String(armTemplateFolderPath),
            BranchRef: utils.String(branchRef),
            DisplayName: utils.String(displayName),
            FolderPath: utils.String(folderPath),
            SecurityToken: utils.String(securityToken),
            SourceType: devtestlab.SourceControlType(sourceType),
            Status: devtestlab.EnableStatus(status),
            UniqueIdentifier: utils.String(uniqueIdentifier),
            URI: utils.String(uRI),
        },
        Tags: tags.Expand(tags),
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroupName, name, name, artifactSource); err != nil {
        return fmt.Errorf("Error creating Artifact Source (Name %q / Lab Name %q / Resource Group %q): %+v", name, name, resourceGroupName, err)
    }


    resp, err := client.Get(ctx, resourceGroupName, name, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Artifact Source (Name %q / Lab Name %q / Resource Group %q): %+v", name, name, resourceGroupName, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Artifact Source (Name %q / Lab Name %q / Resource Group %q) ID", name, name, resourceGroupName)
    }
    d.SetId(*resp.ID)

    return resourceArmArtifactSourceRead(d, meta)
}

func resourceArmArtifactSourceRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).artifactSourcesClient
    ctx, cancel := timeouts.ForRead(meta.(*ArmClient).StopContext, d)
    defer cancel()

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroupName := id.ResourceGroup
    name := id.Path["labs"]
    name := id.Path["artifactsources"]

    resp, err := client.Get(ctx, resourceGroupName, name, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Artifact Source %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Artifact Source (Name %q / Lab Name %q / Resource Group %q): %+v", name, name, resourceGroupName, err)
    }


    d.Set("resource_group", resourceGroupName)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if artifactSourcePropertiesFragment := resp.ArtifactSourcePropertiesFragment; artifactSourcePropertiesFragment != nil {
        d.Set("arm_template_folder_path", artifactSourcePropertiesFragment.ArmTemplateFolderPath)
        d.Set("branch_ref", artifactSourcePropertiesFragment.BranchRef)
        d.Set("created_date", (artifactSourcePropertiesFragment.CreatedDate).String())
        d.Set("display_name", artifactSourcePropertiesFragment.DisplayName)
        d.Set("folder_path", artifactSourcePropertiesFragment.FolderPath)
        d.Set("provisioning_state", artifactSourcePropertiesFragment.ProvisioningState)
        d.Set("security_token", artifactSourcePropertiesFragment.SecurityToken)
        d.Set("source_type", string(artifactSourcePropertiesFragment.SourceType))
        d.Set("status", string(artifactSourcePropertiesFragment.Status))
        d.Set("uri", artifactSourcePropertiesFragment.URI)
        d.Set("unique_identifier", artifactSourcePropertiesFragment.UniqueIdentifier)
    }
    d.Set("id", resp.ID)
    d.Set("lab_name", name)
    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}

func resourceArmArtifactSourceUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).artifactSourcesClient
    ctx, cancel := timeouts.ForUpdate(meta.(*ArmClient).StopContext, d)
    defer cancel()

      resourceGroupName := d.Get("resource_group").(string)
    location := azure.NormalizeLocation(d.Get("location").(string))
    armTemplateFolderPath := d.Get("arm_template_folder_path").(string)
    branchRef := d.Get("branch_ref").(string)
    displayName := d.Get("display_name").(string)
    folderPath := d.Get("folder_path").(string)
    name := d.Get("lab_name").(string)
    name := d.Get("name").(string)
    securityToken := d.Get("security_token").(string)
    sourceType := d.Get("source_type").(string)
    status := d.Get("status").(string)
    uRI := d.Get("uri").(string)
    uniqueIdentifier := d.Get("unique_identifier").(string)
    tags := d.Get("tags").(map[string]interface{})

    artifactSource := devtestlab.ArtifactSourceFragment{
        Location: utils.String(location),
        ArtifactSourcePropertiesFragment: &devtestlab.ArtifactSourcePropertiesFragment{
            ArmTemplateFolderPath: utils.String(armTemplateFolderPath),
            BranchRef: utils.String(branchRef),
            DisplayName: utils.String(displayName),
            FolderPath: utils.String(folderPath),
            SecurityToken: utils.String(securityToken),
            SourceType: devtestlab.SourceControlType(sourceType),
            Status: devtestlab.EnableStatus(status),
            UniqueIdentifier: utils.String(uniqueIdentifier),
            URI: utils.String(uRI),
        },
        Tags: tags.Expand(tags),
    }


    if _, err := client.Update(ctx, resourceGroupName, name, name, artifactSource); err != nil {
        return fmt.Errorf("Error updating Artifact Source (Name %q / Lab Name %q / Resource Group %q): %+v", name, name, resourceGroupName, err)
    }

    return resourceArmArtifactSourceRead(d, meta)
}

func resourceArmArtifactSourceDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).artifactSourcesClient
    ctx, cancel := timeouts.ForDelete(meta.(*ArmClient).StopContext, d)
    defer cancel()


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroupName := id.ResourceGroup
    name := id.Path["labs"]
    name := id.Path["artifactsources"]

    if _, err := client.Delete(ctx, resourceGroupName, name, name); err != nil {
        return fmt.Errorf("Error deleting Artifact Source (Name %q / Lab Name %q / Resource Group %q): %+v", name, name, resourceGroupName, err)
    }

    return nil
}
