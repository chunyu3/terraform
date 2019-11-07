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



func resourceArmApiOperation() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmApiOperationCreate,
        Read: resourceArmApiOperationRead,
        Update: resourceArmApiOperationUpdate,
        Delete: resourceArmApiOperationDelete,

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

            "display_name": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "method": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "operation_id": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "url_template": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "description": {
                Type: schema.TypeString,
                Optional: true,
            },

            "policies": {
                Type: schema.TypeString,
                Optional: true,
            },

            "request": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "description": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "headers": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "name": {
                                        Type: schema.TypeString,
                                        Required: true,
                                        ValidateFunc: validate.NoEmptyStrings,
                                    },
                                    "type": {
                                        Type: schema.TypeString,
                                        Required: true,
                                        ValidateFunc: validate.NoEmptyStrings,
                                    },
                                    "default_value": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                    "description": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                    "required": {
                                        Type: schema.TypeBool,
                                        Optional: true,
                                    },
                                    "values": {
                                        Type: schema.TypeList,
                                        Optional: true,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                        },
                                    },
                                },
                            },
                        },
                        "query_parameters": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "name": {
                                        Type: schema.TypeString,
                                        Required: true,
                                        ValidateFunc: validate.NoEmptyStrings,
                                    },
                                    "type": {
                                        Type: schema.TypeString,
                                        Required: true,
                                        ValidateFunc: validate.NoEmptyStrings,
                                    },
                                    "default_value": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                    "description": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                    "required": {
                                        Type: schema.TypeBool,
                                        Optional: true,
                                    },
                                    "values": {
                                        Type: schema.TypeList,
                                        Optional: true,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                        },
                                    },
                                },
                            },
                        },
                        "representations": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "content_type": {
                                        Type: schema.TypeString,
                                        Required: true,
                                        ValidateFunc: validate.NoEmptyStrings,
                                    },
                                    "form_parameters": {
                                        Type: schema.TypeList,
                                        Optional: true,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "name": {
                                                    Type: schema.TypeString,
                                                    Required: true,
                                                    ValidateFunc: validate.NoEmptyStrings,
                                                },
                                                "type": {
                                                    Type: schema.TypeString,
                                                    Required: true,
                                                    ValidateFunc: validate.NoEmptyStrings,
                                                },
                                                "default_value": {
                                                    Type: schema.TypeString,
                                                    Optional: true,
                                                },
                                                "description": {
                                                    Type: schema.TypeString,
                                                    Optional: true,
                                                },
                                                "required": {
                                                    Type: schema.TypeBool,
                                                    Optional: true,
                                                },
                                                "values": {
                                                    Type: schema.TypeList,
                                                    Optional: true,
                                                    Elem: &schema.Schema{
                                                        Type: schema.TypeString,
                                                    },
                                                },
                                            },
                                        },
                                    },
                                    "sample": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                    "schema_id": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                    "type_name": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                },
                            },
                        },
                    },
                },
            },

            "responses": {
                Type: schema.TypeList,
                Optional: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "status_code": {
                            Type: schema.TypeInt,
                            Required: true,
                        },
                        "description": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "headers": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "name": {
                                        Type: schema.TypeString,
                                        Required: true,
                                        ValidateFunc: validate.NoEmptyStrings,
                                    },
                                    "type": {
                                        Type: schema.TypeString,
                                        Required: true,
                                        ValidateFunc: validate.NoEmptyStrings,
                                    },
                                    "default_value": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                    "description": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                    "required": {
                                        Type: schema.TypeBool,
                                        Optional: true,
                                    },
                                    "values": {
                                        Type: schema.TypeList,
                                        Optional: true,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                        },
                                    },
                                },
                            },
                        },
                        "representations": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "content_type": {
                                        Type: schema.TypeString,
                                        Required: true,
                                        ValidateFunc: validate.NoEmptyStrings,
                                    },
                                    "form_parameters": {
                                        Type: schema.TypeList,
                                        Optional: true,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "name": {
                                                    Type: schema.TypeString,
                                                    Required: true,
                                                    ValidateFunc: validate.NoEmptyStrings,
                                                },
                                                "type": {
                                                    Type: schema.TypeString,
                                                    Required: true,
                                                    ValidateFunc: validate.NoEmptyStrings,
                                                },
                                                "default_value": {
                                                    Type: schema.TypeString,
                                                    Optional: true,
                                                },
                                                "description": {
                                                    Type: schema.TypeString,
                                                    Optional: true,
                                                },
                                                "required": {
                                                    Type: schema.TypeBool,
                                                    Optional: true,
                                                },
                                                "values": {
                                                    Type: schema.TypeList,
                                                    Optional: true,
                                                    Elem: &schema.Schema{
                                                        Type: schema.TypeString,
                                                    },
                                                },
                                            },
                                        },
                                    },
                                    "sample": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                    "schema_id": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                    "type_name": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                },
                            },
                        },
                    },
                },
            },

            "template_parameters": {
                Type: schema.TypeList,
                Optional: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "name": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "type": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "default_value": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "description": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "required": {
                            Type: schema.TypeBool,
                            Optional: true,
                        },
                        "values": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
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

func resourceArmApiOperationCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).apiOperationClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    apiID := d.Get("api_id").(string)
    operationID := d.Get("operation_id").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, name, apiID, operationID)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Api Operation %q (Operation %q / Api %q / Resource Group %q): %+v", name, operationID, apiID, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_api_operation", *existing.ID)
        }
    }

    description := d.Get("description").(string)
    displayName := d.Get("display_name").(string)
    method := d.Get("method").(string)
    policies := d.Get("policies").(string)
    request := d.Get("request").([]interface{})
    responses := d.Get("responses").([]interface{})
    templateParameters := d.Get("template_parameters").([]interface{})
    urlTemplate := d.Get("url_template").(string)

    parameters := apimanagement.OperationContract{
        OperationContractProperties: &apimanagement.OperationContractProperties{
            Description: utils.String(description),
            DisplayName: utils.String(displayName),
            Method: utils.String(method),
            Policies: utils.String(policies),
            Request: expandArmApiOperationRequestContract(request),
            Responses: expandArmApiOperationResponseContract(responses),
            TemplateParameters: expandArmApiOperationParameterContract(templateParameters),
            URLTemplate: utils.String(urlTemplate),
        },
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroup, name, apiID, operationID, parameters); err != nil {
        return fmt.Errorf("Error creating Api Operation %q (Operation %q / Api %q / Resource Group %q): %+v", name, operationID, apiID, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, name, apiID, operationID)
    if err != nil {
        return fmt.Errorf("Error retrieving Api Operation %q (Operation %q / Api %q / Resource Group %q): %+v", name, operationID, apiID, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Api Operation %q (Operation %q / Api %q / Resource Group %q) ID", name, operationID, apiID, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmApiOperationRead(d, meta)
}

func resourceArmApiOperationRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).apiOperationClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["service"]
    apiID := id.Path["apis"]
    operationID := id.Path["operations"]

    resp, err := client.Get(ctx, resourceGroup, name, apiID, operationID)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Api Operation %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Api Operation %q (Operation %q / Api %q / Resource Group %q): %+v", name, operationID, apiID, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    d.Set("api_id", apiID)
    if operationContractProperties := resp.OperationContractProperties; operationContractProperties != nil {
        d.Set("description", operationContractProperties.Description)
        d.Set("display_name", operationContractProperties.DisplayName)
        d.Set("method", operationContractProperties.Method)
        d.Set("policies", operationContractProperties.Policies)
        if err := d.Set("request", flattenArmApiOperationRequestContract(operationContractProperties.Request)); err != nil {
            return fmt.Errorf("Error setting `request`: %+v", err)
        }
        if err := d.Set("responses", flattenArmApiOperationResponseContract(operationContractProperties.Responses)); err != nil {
            return fmt.Errorf("Error setting `responses`: %+v", err)
        }
        if err := d.Set("template_parameters", flattenArmApiOperationParameterContract(operationContractProperties.TemplateParameters)); err != nil {
            return fmt.Errorf("Error setting `template_parameters`: %+v", err)
        }
        d.Set("url_template", operationContractProperties.URLTemplate)
    }
    d.Set("operation_id", operationID)
    d.Set("type", resp.Type)

    return nil
}

func resourceArmApiOperationUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).apiOperationClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    apiID := d.Get("api_id").(string)
    description := d.Get("description").(string)
    displayName := d.Get("display_name").(string)
    method := d.Get("method").(string)
    operationID := d.Get("operation_id").(string)
    policies := d.Get("policies").(string)
    request := d.Get("request").([]interface{})
    responses := d.Get("responses").([]interface{})
    templateParameters := d.Get("template_parameters").([]interface{})
    urlTemplate := d.Get("url_template").(string)

    parameters := apimanagement.OperationContract{
        OperationContractProperties: &apimanagement.OperationContractProperties{
            Description: utils.String(description),
            DisplayName: utils.String(displayName),
            Method: utils.String(method),
            Policies: utils.String(policies),
            Request: expandArmApiOperationRequestContract(request),
            Responses: expandArmApiOperationResponseContract(responses),
            TemplateParameters: expandArmApiOperationParameterContract(templateParameters),
            URLTemplate: utils.String(urlTemplate),
        },
    }


    if _, err := client.Update(ctx, resourceGroup, name, apiID, operationID, parameters); err != nil {
        return fmt.Errorf("Error updating Api Operation %q (Operation %q / Api %q / Resource Group %q): %+v", name, operationID, apiID, resourceGroup, err)
    }

    return resourceArmApiOperationRead(d, meta)
}

func resourceArmApiOperationDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).apiOperationClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["service"]
    apiID := id.Path["apis"]
    operationID := id.Path["operations"]

    if _, err := client.Delete(ctx, resourceGroup, name, apiID, operationID); err != nil {
        return fmt.Errorf("Error deleting Api Operation %q (Operation %q / Api %q / Resource Group %q): %+v", name, operationID, apiID, resourceGroup, err)
    }

    return nil
}

func expandArmApiOperationRequestContract(input []interface{}) *apimanagement.RequestContract {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    description := v["description"].(string)
    queryParameters := v["query_parameters"].([]interface{})
    headers := v["headers"].([]interface{})
    representations := v["representations"].([]interface{})

    result := apimanagement.RequestContract{
        Description: utils.String(description),
        Headers: expandArmApiOperationParameterContract(headers),
        QueryParameters: expandArmApiOperationParameterContract(queryParameters),
        Representations: expandArmApiOperationRepresentationContract(representations),
    }
    return &result
}

func expandArmApiOperationResponseContract(input []interface{}) *[]apimanagement.ResponseContract {
    results := make([]apimanagement.ResponseContract, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        statusCode := v["status_code"].(int)
        description := v["description"].(string)
        representations := v["representations"].([]interface{})
        headers := v["headers"].([]interface{})

        result := apimanagement.ResponseContract{
            Description: utils.String(description),
            Headers: expandArmApiOperationParameterContract(headers),
            Representations: expandArmApiOperationRepresentationContract(representations),
            StatusCode: utils.Int32(int32(statusCode)),
        }

        results = append(results, result)
    }
    return &results
}

func expandArmApiOperationParameterContract(input []interface{}) *[]apimanagement.ParameterContract {
    results := make([]apimanagement.ParameterContract, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        name := v["name"].(string)
        description := v["description"].(string)
        type := v["type"].(string)
        defaultValue := v["default_value"].(string)
        required := v["required"].(bool)
        values := v["values"].([]interface{})

        result := apimanagement.ParameterContract{
            DefaultValue: utils.String(defaultValue),
            Description: utils.String(description),
            Name: utils.String(name),
            Required: utils.Bool(required),
            Type: utils.String(type),
            Values: utils.ExpandStringSlice(values),
        }

        results = append(results, result)
    }
    return &results
}

func expandArmApiOperationRepresentationContract(input []interface{}) *[]apimanagement.RepresentationContract {
    results := make([]apimanagement.RepresentationContract, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        contentType := v["content_type"].(string)
        sample := v["sample"].(string)
        schemaId := v["schema_id"].(string)
        typeName := v["type_name"].(string)
        formParameters := v["form_parameters"].([]interface{})

        result := apimanagement.RepresentationContract{
            ContentType: utils.String(contentType),
            FormParameters: expandArmApiOperationParameterContract(formParameters),
            Sample: utils.String(sample),
            SchemaID: utils.String(schemaId),
            TypeName: utils.String(typeName),
        }

        results = append(results, result)
    }
    return &results
}


func flattenArmApiOperationRequestContract(input *apimanagement.RequestContract) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if description := input.Description; description != nil {
        result["description"] = *description
    }
    result["headers"] = flattenArmApiOperationParameterContract(input.Headers)
    result["query_parameters"] = flattenArmApiOperationParameterContract(input.QueryParameters)
    result["representations"] = flattenArmApiOperationRepresentationContract(input.Representations)

    return []interface{}{result}
}

func flattenArmApiOperationResponseContract(input *[]apimanagement.ResponseContract) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        if description := item.Description; description != nil {
            v["description"] = *description
        }
        v["headers"] = flattenArmApiOperationParameterContract(item.Headers)
        v["representations"] = flattenArmApiOperationRepresentationContract(item.Representations)
        if statusCode := item.StatusCode; statusCode != nil {
            v["status_code"] = int(*statusCode)
        }

        results = append(results, v)
    }

    return results
}

func flattenArmApiOperationParameterContract(input *[]apimanagement.ParameterContract) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        if name := item.Name; name != nil {
            v["name"] = *name
        }
        if defaultValue := item.DefaultValue; defaultValue != nil {
            v["default_value"] = *defaultValue
        }
        if description := item.Description; description != nil {
            v["description"] = *description
        }
        if required := item.Required; required != nil {
            v["required"] = *required
        }
        if type := item.Type; type != nil {
            v["type"] = *type
        }
        v["values"] = utils.FlattenStringSlice(item.Values)

        results = append(results, v)
    }

    return results
}

func flattenArmApiOperationRepresentationContract(input *[]apimanagement.RepresentationContract) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        if contentType := item.ContentType; contentType != nil {
            v["content_type"] = *contentType
        }
        v["form_parameters"] = flattenArmApiOperationParameterContract(item.FormParameters)
        if sample := item.Sample; sample != nil {
            v["sample"] = *sample
        }
        if schemaId := item.SchemaID; schemaId != nil {
            v["schema_id"] = *schemaId
        }
        if typeName := item.TypeName; typeName != nil {
            v["type_name"] = *typeName
        }

        results = append(results, v)
    }

    return results
}
