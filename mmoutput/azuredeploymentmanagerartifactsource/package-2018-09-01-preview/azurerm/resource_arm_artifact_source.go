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
        Create: resourceArmArtifactSourceCreateUpdate,
        Read: resourceArmArtifactSourceRead,
        Update: resourceArmArtifactSourceCreateUpdate,
        Delete: resourceArmArtifactSourceDelete,

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

            "artifact_source_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "source_type": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "artifact_root": {
                Type: schema.TypeString,
                Optional: true,
            },

            "type": {
                Type: schema.TypeString,
                Computed: true,
            },

            "tags": tags.Schema(),
        },
    }
}

func resourceArmArtifactSourceCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).artifactSourcesClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    artifactSourceName := d.Get("artifact_source_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, artifactSourceName)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Artifact Source (Artifact Source Name %q / Resource Group %q): %+v", artifactSourceName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_artifact_source", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    artifactRoot := d.Get("artifact_root").(string)
    sourceType := d.Get("source_type").(string)
    t := d.Get("tags").(map[string]interface{})

    artifactSourceInfo := azuredeploymentmanager.ArtifactSource{
        Location: utils.String(location),
        ArtifactSource_properties: &azuredeploymentmanager.ArtifactSource_properties{
            ArtifactRoot: utils.String(artifactRoot),
            SourceType: utils.String(sourceType),
        },
        Tags: tags.Expand(t),
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroup, artifactSourceName, artifactSourceInfo); err != nil {
        return fmt.Errorf("Error creating Artifact Source (Artifact Source Name %q / Resource Group %q): %+v", artifactSourceName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, artifactSourceName)
    if err != nil {
        return fmt.Errorf("Error retrieving Artifact Source (Artifact Source Name %q / Resource Group %q): %+v", artifactSourceName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Artifact Source (Artifact Source Name %q / Resource Group %q) ID", artifactSourceName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmArtifactSourceRead(d, meta)
}

func resourceArmArtifactSourceRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).artifactSourcesClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    artifactSourceName := id.Path["artifactSources"]

    resp, err := client.Get(ctx, resourceGroup, artifactSourceName)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Artifact Source %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Artifact Source (Artifact Source Name %q / Resource Group %q): %+v", artifactSourceName, resourceGroup, err)
    }


    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if artifactSourceProperties := resp.ArtifactSource_properties; artifactSourceProperties != nil {
        d.Set("artifact_root", artifactSourceProperties.ArtifactRoot)
        d.Set("source_type", artifactSourceProperties.SourceType)
    }
    d.Set("artifact_source_name", artifactSourceName)
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}


func resourceArmArtifactSourceDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).artifactSourcesClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    artifactSourceName := id.Path["artifactSources"]

    if _, err := client.Delete(ctx, resourceGroup, artifactSourceName); err != nil {
        return fmt.Errorf("Error deleting Artifact Source (Artifact Source Name %q / Resource Group %q): %+v", artifactSourceName, resourceGroup, err)
    }

    return nil
}