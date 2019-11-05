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



func resourceArmApi() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmApiCreate,
        Read: resourceArmApiRead,
        Update: resourceArmApiUpdate,
        Delete: resourceArmApiDelete,

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

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "api_id": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "path": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "protocols": {
                Type: schema.TypeList,
                Required: true,
                ForceNew: true,
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                    ValidateFunc: validation.StringInSlice([]string{
                        string(apimanagement.Http),
                        string(apimanagement.Https),
                   }, false),
                },
            },

            "service_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "service_url": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "authentication_settings": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "o_auth2": {
                            Type: schema.TypeList,
                            Optional: true,
                            MaxItems: 1,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "authorization_server_id": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                    "scope": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                },
                            },
                        },
                    },
                },
            },

            "description": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
            },

            "subscription_key_parameter_names": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "header": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "query": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },

            "type": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(apimanagement.Http),
                    string(apimanagement.Soap),
                }, false),
                Default: string(apimanagement.Http),
            },
        },
    }
}

func resourceArmApiCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).apisClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    apiID := d.Get("api_id").(string)
    serviceName := d.Get("service_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, serviceName, apiID)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Api (Api %q / Service Name %q / Resource Group %q): %+v", apiID, serviceName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_api", *existing.ID)
        }
    }

    name := d.Get("name").(string)
    authenticationSettings := d.Get("authentication_settings").([]interface{})
    description := d.Get("description").(string)
    path := d.Get("path").(string)
    protocols := d.Get("protocols").([]interface{})
    serviceUrl := d.Get("service_url").(string)
    subscriptionKeyParameterNames := d.Get("subscription_key_parameter_names").([]interface{})
    type := d.Get("type").(string)

    parameters := apimanagement.ApiContract{
        AuthenticationSettings: expandArmApiAuthenticationSettingsContract(authenticationSettings),
        Description: utils.String(description),
        Name: utils.String(name),
        Path: utils.String(path),
        Protocols: expandArmApi(protocols),
        ServiceURL: utils.String(serviceUrl),
        SubscriptionKeyParameterNames: expandArmApiSubscriptionKeyParameterNamesContract(subscriptionKeyParameterNames),
        Type: apimanagement.ApiTypeContract(type),
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroup, serviceName, apiID, parameters); err != nil {
        return fmt.Errorf("Error creating Api (Api %q / Service Name %q / Resource Group %q): %+v", apiID, serviceName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, serviceName, apiID)
    if err != nil {
        return fmt.Errorf("Error retrieving Api (Api %q / Service Name %q / Resource Group %q): %+v", apiID, serviceName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Api (Api %q / Service Name %q / Resource Group %q) ID", apiID, serviceName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmApiRead(d, meta)
}

func resourceArmApiRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).apisClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    serviceName := id.Path["service"]
    apiID := id.Path["apis"]

    resp, err := client.Get(ctx, resourceGroup, serviceName, apiID)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Api %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Api (Api %q / Service Name %q / Resource Group %q): %+v", apiID, serviceName, resourceGroup, err)
    }


    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    d.Set("api_id", apiID)
    if err := d.Set("authentication_settings", flattenArmApiAuthenticationSettingsContract(resp.AuthenticationSettings)); err != nil {
        return fmt.Errorf("Error setting `authentication_settings`: %+v", err)
    }
    d.Set("description", resp.Description)
    d.Set("path", resp.Path)
    if err := d.Set("protocols", flattenArmApi(string(resp.Protocols))); err != nil {
        return fmt.Errorf("Error setting `protocols`: %+v", err)
    }
    d.Set("service_name", serviceName)
    d.Set("service_url", resp.ServiceURL)
    if err := d.Set("subscription_key_parameter_names", flattenArmApiSubscriptionKeyParameterNamesContract(resp.SubscriptionKeyParameterNames)); err != nil {
        return fmt.Errorf("Error setting `subscription_key_parameter_names`: %+v", err)
    }
    d.Set("type", string(resp.Type))

    return nil
}

func resourceArmApiUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).apisClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    apiID := d.Get("api_id").(string)
    authenticationSettings := d.Get("authentication_settings").([]interface{})
    description := d.Get("description").(string)
    path := d.Get("path").(string)
    protocols := d.Get("protocols").([]interface{})
    serviceName := d.Get("service_name").(string)
    serviceUrl := d.Get("service_url").(string)
    subscriptionKeyParameterNames := d.Get("subscription_key_parameter_names").([]interface{})
    type := d.Get("type").(string)

    parameters := apimanagement.ApiContract{
        AuthenticationSettings: expandArmApiAuthenticationSettingsContract(authenticationSettings),
        Description: utils.String(description),
        Name: utils.String(name),
        Path: utils.String(path),
        Protocols: expandArmApi(protocols),
        ServiceURL: utils.String(serviceUrl),
        SubscriptionKeyParameterNames: expandArmApiSubscriptionKeyParameterNamesContract(subscriptionKeyParameterNames),
        Type: apimanagement.ApiTypeContract(type),
    }


    if _, err := client.Update(ctx, resourceGroup, serviceName, apiID, parameters); err != nil {
        return fmt.Errorf("Error updating Api (Api %q / Service Name %q / Resource Group %q): %+v", apiID, serviceName, resourceGroup, err)
    }

    return resourceArmApiRead(d, meta)
}

func resourceArmApiDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).apisClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    serviceName := id.Path["service"]
    apiID := id.Path["apis"]

    if _, err := client.Delete(ctx, resourceGroup, serviceName, apiID); err != nil {
        return fmt.Errorf("Error deleting Api (Api %q / Service Name %q / Resource Group %q): %+v", apiID, serviceName, resourceGroup, err)
    }

    return nil
}

func expandArmApiAuthenticationSettingsContract(input []interface{}) *apimanagement.AuthenticationSettingsContract {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    oAuth2 := v["o_auth2"].([]interface{})

    result := apimanagement.AuthenticationSettingsContract{
        OAuth2: expandArmApiOAuth2AuthenticationSettingsContract(oAuth2),
    }
    return &result
}

func expandArmApi(input []interface{}) *[]apimanagement. {
    results := make([]apimanagement., 0)
    for _, item := range input {
        v := item.(string)
        result := apimanagement.(v)
        results = append(results, result)
    }
    return &results
}

func expandArmApiSubscriptionKeyParameterNamesContract(input []interface{}) *apimanagement.SubscriptionKeyParameterNamesContract {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    header := v["header"].(string)
    query := v["query"].(string)

    result := apimanagement.SubscriptionKeyParameterNamesContract{
        Header: utils.String(header),
        Query: utils.String(query),
    }
    return &result
}

func expandArmApiOAuth2AuthenticationSettingsContract(input []interface{}) *apimanagement.OAuth2AuthenticationSettingsContract {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    authorizationServerId := v["authorization_server_id"].(string)
    scope := v["scope"].(string)

    result := apimanagement.OAuth2AuthenticationSettingsContract{
        AuthorizationServerID: utils.String(authorizationServerId),
        Scope: utils.String(scope),
    }
    return &result
}


func flattenArmApiAuthenticationSettingsContract(input *apimanagement.AuthenticationSettingsContract) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    result["o_auth2"] = flattenArmApiOAuth2AuthenticationSettingsContract(input.OAuth2)

    return []interface{}{result}
}

func flattenArmApi(input *[]apimanagement.) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        result := string(item)
        results = append(results, result)
    }

    return results
}

func flattenArmApiSubscriptionKeyParameterNamesContract(input *apimanagement.SubscriptionKeyParameterNamesContract) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if header := input.Header; header != nil {
        result["header"] = *header
    }
    if query := input.Query; query != nil {
        result["query"] = *query
    }

    return []interface{}{result}
}

func flattenArmApiOAuth2AuthenticationSettingsContract(input *apimanagement.OAuth2AuthenticationSettingsContract) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if authorizationServerId := input.AuthorizationServerID; authorizationServerId != nil {
        result["authorization_server_id"] = *authorizationServerId
    }
    if scope := input.Scope; scope != nil {
        result["scope"] = *scope
    }

    return []interface{}{result}
}
