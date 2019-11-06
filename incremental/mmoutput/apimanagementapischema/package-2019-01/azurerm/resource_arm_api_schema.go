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



func resourceArmApiSchema() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmApiSchemaCreateUpdate,
        Read: resourceArmApiSchemaRead,
        Update: resourceArmApiSchemaCreateUpdate,
        Delete: resourceArmApiSchemaDelete,

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

            "content_type": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "schema_id": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "document": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "value": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },

            "type": {
                Type: schema.TypeString,
                Computed: true,
            },
        },
    }
}

func resourceArmApiSchemaCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).apiSchemaClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    apiID := d.Get("api_id").(string)
    schemaID := d.Get("schema_id").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, name, apiID, schemaID)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Api Schema %q (Schema %q / Api %q / Resource Group %q): %+v", name, schemaID, apiID, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_api_schema", *existing.ID)
        }
    }

    contentType := d.Get("content_type").(string)
    document := d.Get("document").([]interface{})

    parameters := apimanagement.SchemaCreateOrUpdateContract{
        SchemaCreateOrUpdateProperties: &apimanagement.SchemaCreateOrUpdateProperties{
            ContentType: utils.String(contentType),
            Document: expandArmApiSchemaSchemaDocumentProperties(document),
        },
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroup, name, apiID, schemaID, parameters); err != nil {
        return fmt.Errorf("Error creating Api Schema %q (Schema %q / Api %q / Resource Group %q): %+v", name, schemaID, apiID, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, name, apiID, schemaID)
    if err != nil {
        return fmt.Errorf("Error retrieving Api Schema %q (Schema %q / Api %q / Resource Group %q): %+v", name, schemaID, apiID, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Api Schema %q (Schema %q / Api %q / Resource Group %q) ID", name, schemaID, apiID, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmApiSchemaRead(d, meta)
}

func resourceArmApiSchemaRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).apiSchemaClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["service"]
    apiID := id.Path["apis"]
    schemaID := id.Path["schemas"]

    resp, err := client.Get(ctx, resourceGroup, name, apiID, schemaID)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Api Schema %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Api Schema %q (Schema %q / Api %q / Resource Group %q): %+v", name, schemaID, apiID, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    d.Set("api_id", apiID)
    if schemaCreateOrUpdateProperties := resp.SchemaCreateOrUpdateProperties; schemaCreateOrUpdateProperties != nil {
        d.Set("content_type", schemaCreateOrUpdateProperties.ContentType)
        if err := d.Set("document", flattenArmApiSchemaSchemaDocumentProperties(schemaCreateOrUpdateProperties.Document)); err != nil {
            return fmt.Errorf("Error setting `document`: %+v", err)
        }
    }
    d.Set("schema_id", schemaID)
    d.Set("type", resp.Type)

    return nil
}


func resourceArmApiSchemaDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).apiSchemaClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["service"]
    apiID := id.Path["apis"]
    schemaID := id.Path["schemas"]

    if _, err := client.Delete(ctx, resourceGroup, name, apiID, schemaID); err != nil {
        return fmt.Errorf("Error deleting Api Schema %q (Schema %q / Api %q / Resource Group %q): %+v", name, schemaID, apiID, resourceGroup, err)
    }

    return nil
}

func expandArmApiSchemaSchemaDocumentProperties(input []interface{}) *apimanagement.SchemaDocumentProperties {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    value := v["value"].(string)

    result := apimanagement.SchemaDocumentProperties{
        Value: utils.String(value),
    }
    return &result
}


func flattenArmApiSchemaSchemaDocumentProperties(input *apimanagement.SchemaDocumentProperties) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})


    return []interface{}{result}
}
