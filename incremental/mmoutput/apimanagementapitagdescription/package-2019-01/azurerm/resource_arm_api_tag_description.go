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



func resourceArmApiTagDescription() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmApiTagDescriptionCreateUpdate,
        Read: resourceArmApiTagDescriptionRead,
        Update: resourceArmApiTagDescriptionCreateUpdate,
        Delete: resourceArmApiTagDescriptionDelete,

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

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "api_id": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "tag_id": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "description": {
                Type: schema.TypeString,
                Optional: true,
            },

            "external_docs_description": {
                Type: schema.TypeString,
                Optional: true,
            },

            "external_docs_url": {
                Type: schema.TypeString,
                Optional: true,
            },

            "display_name": {
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

func resourceArmApiTagDescriptionCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).apiTagDescriptionClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    apiID := d.Get("api_id").(string)
    tagID := d.Get("tag_id").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, name, apiID, tagID)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Api Tag Description %q (Tag %q / Api %q / Resource Group %q): %+v", name, tagID, apiID, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_api_tag_description", *existing.ID)
        }
    }

    description := d.Get("description").(string)
    externalDocsDescription := d.Get("external_docs_description").(string)
    externalDocsUrl := d.Get("external_docs_url").(string)

    parameters := apimanagement.TagDescriptionCreateParameters{
        TagDescriptionBaseProperties: &apimanagement.TagDescriptionBaseProperties{
            Description: utils.String(description),
            ExternalDocsDescription: utils.String(externalDocsDescription),
            ExternalDocsURL: utils.String(externalDocsUrl),
        },
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroup, name, apiID, tagID, parameters); err != nil {
        return fmt.Errorf("Error creating Api Tag Description %q (Tag %q / Api %q / Resource Group %q): %+v", name, tagID, apiID, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, name, apiID, tagID)
    if err != nil {
        return fmt.Errorf("Error retrieving Api Tag Description %q (Tag %q / Api %q / Resource Group %q): %+v", name, tagID, apiID, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Api Tag Description %q (Tag %q / Api %q / Resource Group %q) ID", name, tagID, apiID, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmApiTagDescriptionRead(d, meta)
}

func resourceArmApiTagDescriptionRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).apiTagDescriptionClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["service"]
    apiID := id.Path["apis"]
    tagID := id.Path["tagDescriptions"]

    resp, err := client.Get(ctx, resourceGroup, name, apiID, tagID)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Api Tag Description %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Api Tag Description %q (Tag %q / Api %q / Resource Group %q): %+v", name, tagID, apiID, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    d.Set("api_id", apiID)
    if tagDescriptionBaseProperties := resp.TagDescriptionBaseProperties; tagDescriptionBaseProperties != nil {
        d.Set("description", tagDescriptionBaseProperties.Description)
        d.Set("display_name", tagDescriptionBaseProperties.DisplayName)
        d.Set("external_docs_description", tagDescriptionBaseProperties.ExternalDocsDescription)
        d.Set("external_docs_url", tagDescriptionBaseProperties.ExternalDocsURL)
    }
    d.Set("tag_id", tagID)
    d.Set("type", resp.Type)

    return nil
}


func resourceArmApiTagDescriptionDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).apiTagDescriptionClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["service"]
    apiID := id.Path["apis"]
    tagID := id.Path["tagDescriptions"]

    if _, err := client.Delete(ctx, resourceGroup, name, apiID, tagID); err != nil {
        return fmt.Errorf("Error deleting Api Tag Description %q (Tag %q / Api %q / Resource Group %q): %+v", name, tagID, apiID, resourceGroup, err)
    }

    return nil
}