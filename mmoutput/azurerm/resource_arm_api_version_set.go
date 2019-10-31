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



func resourceArmApiVersionSet() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmApiVersionSetCreate,
        Read: resourceArmApiVersionSetRead,
        Update: resourceArmApiVersionSetUpdate,
        Delete: resourceArmApiVersionSetDelete,

        Importer: &schema.ResourceImporter{
            State: schema.ImportStatePassthrough,
        },


        Schema: map[string]*schema.Schema{
            "name": {
                Type: schema.TypeString,
                Computed: true,
            },

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "display_name": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "service_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "version_set_id": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "versioning_scheme": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(apimanagement.Segment),
                    string(apimanagement.Query),
                    string(apimanagement.Header),
                }, false),
            },

            "description": {
                Type: schema.TypeString,
                Optional: true,
            },

            "version_header_name": {
                Type: schema.TypeString,
                Optional: true,
            },

            "version_query_name": {
                Type: schema.TypeString,
                Optional: true,
            },

            "type": {
                Type: schema.TypeString,
                Computed: true,
            },
        },
    }
}

func resourceArmApiVersionSetCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).apiVersionSetClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    serviceName := d.Get("service_name").(string)
    versionSetID := d.Get("version_set_id").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        resp, err := client.Get(ctx, resourceGroup, serviceName, versionSetID)
        if err != nil {
            if !utils.ResponseWasNotFound(resp.Response) {
                return fmt.Errorf("Error checking for present of existing Api Version Set (Version Set %q / Service Name %q / Resource Group %q): %+v", versionSetID, serviceName, resourceGroup, err)
            }
        }
        if !utils.ResponseWasNotFound(resp.Response) {
            return tf.ImportAsExistsError("azurerm_api_version_set", *resp.ID)
        }
    }

    description := d.Get("description").(string)
    displayName := d.Get("display_name").(string)
    versionHeaderName := d.Get("version_header_name").(string)
    versionQueryName := d.Get("version_query_name").(string)
    versioningScheme := d.Get("versioning_scheme").(string)

    parameters := apimanagement.ApiVersionSetContract{
        ApiVersionSetContractProperties: &apimanagement.ApiVersionSetContractProperties{
            Description: utils.String(description),
            DisplayName: utils.String(displayName),
            VersionHeaderName: utils.String(versionHeaderName),
            VersionQueryName: utils.String(versionQueryName),
            VersioningScheme: apimanagement.VersioningScheme(versioningScheme),
        },
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroup, serviceName, versionSetID, parameters); err != nil {
        return fmt.Errorf("Error creating Api Version Set (Version Set %q / Service Name %q / Resource Group %q): %+v", versionSetID, serviceName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, serviceName, versionSetID)
    if err != nil {
        return fmt.Errorf("Error retrieving Api Version Set (Version Set %q / Service Name %q / Resource Group %q): %+v", versionSetID, serviceName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Api Version Set (Version Set %q / Service Name %q / Resource Group %q) ID", versionSetID, serviceName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmApiVersionSetRead(d, meta)
}

func resourceArmApiVersionSetRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).apiVersionSetClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    serviceName := id.Path["service"]
    versionSetID := id.Path["apiVersionSets"]

    resp, err := client.Get(ctx, resourceGroup, serviceName, versionSetID)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Api Version Set %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Api Version Set (Version Set %q / Service Name %q / Resource Group %q): %+v", versionSetID, serviceName, resourceGroup, err)
    }


    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if apiVersionSetContractProperties := resp.ApiVersionSetContractProperties; apiVersionSetContractProperties != nil {
        d.Set("description", apiVersionSetContractProperties.Description)
        d.Set("display_name", apiVersionSetContractProperties.DisplayName)
        d.Set("version_header_name", apiVersionSetContractProperties.VersionHeaderName)
        d.Set("version_query_name", apiVersionSetContractProperties.VersionQueryName)
        d.Set("versioning_scheme", string(apiVersionSetContractProperties.VersioningScheme))
    }
    d.Set("service_name", serviceName)
    d.Set("type", resp.Type)
    d.Set("version_set_id", versionSetID)

    return nil
}

func resourceArmApiVersionSetUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).apiVersionSetClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    description := d.Get("description").(string)
    displayName := d.Get("display_name").(string)
    serviceName := d.Get("service_name").(string)
    versionHeaderName := d.Get("version_header_name").(string)
    versionQueryName := d.Get("version_query_name").(string)
    versionSetID := d.Get("version_set_id").(string)
    versioningScheme := d.Get("versioning_scheme").(string)

    parameters := apimanagement.ApiVersionSetContract{
        ApiVersionSetContractProperties: &apimanagement.ApiVersionSetContractProperties{
            Description: utils.String(description),
            DisplayName: utils.String(displayName),
            VersionHeaderName: utils.String(versionHeaderName),
            VersionQueryName: utils.String(versionQueryName),
            VersioningScheme: apimanagement.VersioningScheme(versioningScheme),
        },
    }


    if _, err := client.Update(ctx, resourceGroup, serviceName, versionSetID, parameters); err != nil {
        return fmt.Errorf("Error updating Api Version Set (Version Set %q / Service Name %q / Resource Group %q): %+v", versionSetID, serviceName, resourceGroup, err)
    }

    return resourceArmApiVersionSetRead(d, meta)
}

func resourceArmApiVersionSetDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).apiVersionSetClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    serviceName := id.Path["service"]
    versionSetID := id.Path["apiVersionSets"]

    if _, err := client.Delete(ctx, resourceGroup, serviceName, versionSetID); err != nil {
        return fmt.Errorf("Error deleting Api Version Set (Version Set %q / Service Name %q / Resource Group %q): %+v", versionSetID, serviceName, resourceGroup, err)
    }

    return nil
}
