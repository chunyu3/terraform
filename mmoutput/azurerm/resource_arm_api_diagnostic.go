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



func resourceArmApiDiagnostic() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmApiDiagnosticCreate,
        Read: resourceArmApiDiagnosticRead,
        Update: resourceArmApiDiagnosticUpdate,
        Delete: resourceArmApiDiagnosticDelete,

        Importer: &schema.ResourceImporter{
            State: schema.ImportStatePassthrough,
        },


        Schema: map[string]*schema.Schema{
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

            "diagnostic_id": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "logger_id": {
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

            "always_log": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(apimanagement.allErrors),
                }, false),
                Default: string(apimanagement.allErrors),
            },

            "backend": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "request": {
                            Type: schema.TypeList,
                            Optional: true,
                            MaxItems: 1,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "body": {
                                        Type: schema.TypeList,
                                        Optional: true,
                                        MaxItems: 1,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "bytes": {
                                                    Type: schema.TypeInt,
                                                    Optional: true,
                                                },
                                            },
                                        },
                                    },
                                    "headers": {
                                        Type: schema.TypeList,
                                        Optional: true,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                        },
                                    },
                                },
                            },
                        },
                        "response": {
                            Type: schema.TypeList,
                            Optional: true,
                            MaxItems: 1,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "body": {
                                        Type: schema.TypeList,
                                        Optional: true,
                                        MaxItems: 1,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "bytes": {
                                                    Type: schema.TypeInt,
                                                    Optional: true,
                                                },
                                            },
                                        },
                                    },
                                    "headers": {
                                        Type: schema.TypeList,
                                        Optional: true,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                        },
                                    },
                                },
                            },
                        },
                    },
                },
            },

            "enable_http_correlation_headers": {
                Type: schema.TypeBool,
                Optional: true,
            },

            "frontend": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "request": {
                            Type: schema.TypeList,
                            Optional: true,
                            MaxItems: 1,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "body": {
                                        Type: schema.TypeList,
                                        Optional: true,
                                        MaxItems: 1,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "bytes": {
                                                    Type: schema.TypeInt,
                                                    Optional: true,
                                                },
                                            },
                                        },
                                    },
                                    "headers": {
                                        Type: schema.TypeList,
                                        Optional: true,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                        },
                                    },
                                },
                            },
                        },
                        "response": {
                            Type: schema.TypeList,
                            Optional: true,
                            MaxItems: 1,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "body": {
                                        Type: schema.TypeList,
                                        Optional: true,
                                        MaxItems: 1,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "bytes": {
                                                    Type: schema.TypeInt,
                                                    Optional: true,
                                                },
                                            },
                                        },
                                    },
                                    "headers": {
                                        Type: schema.TypeList,
                                        Optional: true,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                        },
                                    },
                                },
                            },
                        },
                    },
                },
            },

            "sampling": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "percentage": {
                            Type: schema.TypeFloat,
                            Optional: true,
                        },
                        "sampling_type": {
                            Type: schema.TypeString,
                            Optional: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(apimanagement.fixed),
                            }, false),
                            Default: string(apimanagement.fixed),
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

func resourceArmApiDiagnosticCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).apiDiagnosticClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    apiID := d.Get("api_id").(string)
    diagnosticID := d.Get("diagnostic_id").(string)
    serviceName := d.Get("service_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        resp, err := client.Get(ctx, resourceGroup, serviceName, apiID, diagnosticID)
        if err != nil {
            if !utils.ResponseWasNotFound(resp.Response) {
                return fmt.Errorf("Error checking for present of existing Api Diagnostic (Diagnostic %q / Api %q / Service Name %q / Resource Group %q): %+v", diagnosticID, apiID, serviceName, resourceGroup, err)
            }
        }
        if !utils.ResponseWasNotFound(resp.Response) {
            return tf.ImportAsExistsError("azurerm_api_diagnostic", *resp.ID)
        }
    }

    alwaysLog := d.Get("always_log").(string)
    backend := d.Get("backend").([]interface{})
    enableHttpCorrelationHeaders := d.Get("enable_http_correlation_headers").(bool)
    frontend := d.Get("frontend").([]interface{})
    loggerId := d.Get("logger_id").(string)
    sampling := d.Get("sampling").([]interface{})

    parameters := apimanagement.DiagnosticContract{
        DiagnosticContractProperties: &apimanagement.DiagnosticContractProperties{
            AlwaysLog: apimanagement.AlwaysLog(alwaysLog),
            Backend: expandArmApiDiagnosticPipelineDiagnosticSettings(backend),
            EnableHttpCorrelationHeaders: utils.Bool(enableHttpCorrelationHeaders),
            Frontend: expandArmApiDiagnosticPipelineDiagnosticSettings(frontend),
            LoggerID: utils.String(loggerId),
            Sampling: expandArmApiDiagnosticSamplingSettings(sampling),
        },
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroup, serviceName, apiID, diagnosticID, parameters); err != nil {
        return fmt.Errorf("Error creating Api Diagnostic (Diagnostic %q / Api %q / Service Name %q / Resource Group %q): %+v", diagnosticID, apiID, serviceName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, serviceName, apiID, diagnosticID)
    if err != nil {
        return fmt.Errorf("Error retrieving Api Diagnostic (Diagnostic %q / Api %q / Service Name %q / Resource Group %q): %+v", diagnosticID, apiID, serviceName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Api Diagnostic (Diagnostic %q / Api %q / Service Name %q / Resource Group %q) ID", diagnosticID, apiID, serviceName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmApiDiagnosticRead(d, meta)
}

func resourceArmApiDiagnosticRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).apiDiagnosticClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    serviceName := id.Path["service"]
    apiID := id.Path["apis"]
    diagnosticID := id.Path["diagnostics"]

    resp, err := client.Get(ctx, resourceGroup, serviceName, apiID, diagnosticID)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Api Diagnostic %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Api Diagnostic (Diagnostic %q / Api %q / Service Name %q / Resource Group %q): %+v", diagnosticID, apiID, serviceName, resourceGroup, err)
    }


    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if diagnosticContractProperties := resp.DiagnosticContractProperties; diagnosticContractProperties != nil {
        d.Set("always_log", string(diagnosticContractProperties.AlwaysLog))
        if err := d.Set("backend", flattenArmApiDiagnosticPipelineDiagnosticSettings(diagnosticContractProperties.Backend)); err != nil {
            return fmt.Errorf("Error setting `backend`: %+v", err)
        }
        d.Set("enable_http_correlation_headers", diagnosticContractProperties.EnableHttpCorrelationHeaders)
        if err := d.Set("frontend", flattenArmApiDiagnosticPipelineDiagnosticSettings(diagnosticContractProperties.Frontend)); err != nil {
            return fmt.Errorf("Error setting `frontend`: %+v", err)
        }
        d.Set("logger_id", diagnosticContractProperties.LoggerID)
        if err := d.Set("sampling", flattenArmApiDiagnosticSamplingSettings(diagnosticContractProperties.Sampling)); err != nil {
            return fmt.Errorf("Error setting `sampling`: %+v", err)
        }
    }
    d.Set("api_id", apiID)
    d.Set("diagnostic_id", diagnosticID)
    d.Set("service_name", serviceName)
    d.Set("type", resp.Type)

    return nil
}

func resourceArmApiDiagnosticUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).apiDiagnosticClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    alwaysLog := d.Get("always_log").(string)
    apiID := d.Get("api_id").(string)
    backend := d.Get("backend").([]interface{})
    diagnosticID := d.Get("diagnostic_id").(string)
    enableHttpCorrelationHeaders := d.Get("enable_http_correlation_headers").(bool)
    frontend := d.Get("frontend").([]interface{})
    loggerId := d.Get("logger_id").(string)
    sampling := d.Get("sampling").([]interface{})
    serviceName := d.Get("service_name").(string)

    parameters := apimanagement.DiagnosticContract{
        DiagnosticContractProperties: &apimanagement.DiagnosticContractProperties{
            AlwaysLog: apimanagement.AlwaysLog(alwaysLog),
            Backend: expandArmApiDiagnosticPipelineDiagnosticSettings(backend),
            EnableHttpCorrelationHeaders: utils.Bool(enableHttpCorrelationHeaders),
            Frontend: expandArmApiDiagnosticPipelineDiagnosticSettings(frontend),
            LoggerID: utils.String(loggerId),
            Sampling: expandArmApiDiagnosticSamplingSettings(sampling),
        },
    }


    if _, err := client.Update(ctx, resourceGroup, serviceName, apiID, diagnosticID, parameters); err != nil {
        return fmt.Errorf("Error updating Api Diagnostic (Diagnostic %q / Api %q / Service Name %q / Resource Group %q): %+v", diagnosticID, apiID, serviceName, resourceGroup, err)
    }

    return resourceArmApiDiagnosticRead(d, meta)
}

func resourceArmApiDiagnosticDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).apiDiagnosticClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    serviceName := id.Path["service"]
    apiID := id.Path["apis"]
    diagnosticID := id.Path["diagnostics"]

    if _, err := client.Delete(ctx, resourceGroup, serviceName, apiID, diagnosticID); err != nil {
        return fmt.Errorf("Error deleting Api Diagnostic (Diagnostic %q / Api %q / Service Name %q / Resource Group %q): %+v", diagnosticID, apiID, serviceName, resourceGroup, err)
    }

    return nil
}

func expandArmApiDiagnosticPipelineDiagnosticSettings(input []interface{}) *apimanagement.PipelineDiagnosticSettings {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    request := v["request"].([]interface{})
    response := v["response"].([]interface{})

    result := apimanagement.PipelineDiagnosticSettings{
        Request: expandArmApiDiagnosticHttpMessageDiagnostic(request),
        Response: expandArmApiDiagnosticHttpMessageDiagnostic(response),
    }
    return &result
}

func expandArmApiDiagnosticSamplingSettings(input []interface{}) *apimanagement.SamplingSettings {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    samplingType := v["sampling_type"].(string)
    percentage := v["percentage"].(float64)

    result := apimanagement.SamplingSettings{
        Percentage: utils.Float(percentage),
        SamplingType: apimanagement.SamplingType(samplingType),
    }
    return &result
}

func expandArmApiDiagnosticHttpMessageDiagnostic(input []interface{}) *apimanagement.HttpMessageDiagnostic {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    headers := v["headers"].([]interface{})
    body := v["body"].([]interface{})

    result := apimanagement.HttpMessageDiagnostic{
        Body: expandArmApiDiagnosticBodyDiagnosticSettings(body),
        Headers: utils.ExpandStringSlice(headers),
    }
    return &result
}

func expandArmApiDiagnosticBodyDiagnosticSettings(input []interface{}) *apimanagement.BodyDiagnosticSettings {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    bytes := v["bytes"].(int)

    result := apimanagement.BodyDiagnosticSettings{
        Bytes: utils.Int32(int32(bytes)),
    }
    return &result
}


func flattenArmApiDiagnosticPipelineDiagnosticSettings(input *apimanagement.PipelineDiagnosticSettings) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    result["request"] = flattenArmApiDiagnosticHttpMessageDiagnostic(input.Request)
    result["response"] = flattenArmApiDiagnosticHttpMessageDiagnostic(input.Response)

    return []interface{}{result}
}

func flattenArmApiDiagnosticSamplingSettings(input *apimanagement.SamplingSettings) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if percentage := input.Percentage; percentage != nil {
        result["percentage"] = *percentage
    }
    result["sampling_type"] = string(input.SamplingType)

    return []interface{}{result}
}

func flattenArmApiDiagnosticHttpMessageDiagnostic(input *apimanagement.HttpMessageDiagnostic) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    result["body"] = flattenArmApiDiagnosticBodyDiagnosticSettings(input.Body)
    result["headers"] = utils.FlattenStringSlice(input.Headers)

    return []interface{}{result}
}

func flattenArmApiDiagnosticBodyDiagnosticSettings(input *apimanagement.BodyDiagnosticSettings) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if bytes := input.Bytes; bytes != nil {
        result["bytes"] = int(*bytes)
    }

    return []interface{}{result}
}
