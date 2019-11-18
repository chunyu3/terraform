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



func resourceArmBackend() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmBackendCreate,
        Read: resourceArmBackendRead,
        Update: resourceArmBackendUpdate,
        Delete: resourceArmBackendDelete,

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

            "backendid": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "credentials": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "authorization": {
                            Type: schema.TypeList,
                            Optional: true,
                            MaxItems: 1,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "parameter": {
                                        Type: schema.TypeString,
                                        Required: true,
                                        ValidateFunc: validate.NoEmptyStrings,
                                    },
                                    "scheme": {
                                        Type: schema.TypeString,
                                        Required: true,
                                        ValidateFunc: validate.NoEmptyStrings,
                                    },
                                },
                            },
                        },
                        "certificate": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                        "header": {
                            Type: schema.TypeMap,
                            Optional: true,
                            Elem: &schema.Schema{Type: schema.TypeString},
                        },
                        "query": {
                            Type: schema.TypeMap,
                            Optional: true,
                            Elem: &schema.Schema{Type: schema.TypeString},
                        },
                    },
                },
            },

            "description": {
                Type: schema.TypeString,
                Optional: true,
            },

            "protocol": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(apimanagement.http),
                    string(apimanagement.soap),
                }, false),
                Default: string(apimanagement.http),
            },

            "proxy": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "url": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "password": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "username": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },

            "resource_id": {
                Type: schema.TypeString,
                Optional: true,
            },

            "service_fabric_cluster": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "client_certificatethumbprint": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "management_endpoints": {
                            Type: schema.TypeList,
                            Required: true,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                        "max_partition_resolution_retries": {
                            Type: schema.TypeInt,
                            Optional: true,
                        },
                        "server_certificate_thumbprints": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                        "server_x509names": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "issuer_certificate_thumbprint": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                    "name": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                },
                            },
                        },
                    },
                },
            },

            "title": {
                Type: schema.TypeString,
                Optional: true,
            },

            "tls": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "validate_certificate_chain": {
                            Type: schema.TypeBool,
                            Optional: true,
                        },
                        "validate_certificate_name": {
                            Type: schema.TypeBool,
                            Optional: true,
                        },
                    },
                },
            },

            "url": {
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

func resourceArmBackendCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).backendClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    backendid := d.Get("backendid").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, name, backendid)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Backend %q (Backendid %q / Resource Group %q): %+v", name, backendid, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_backend", *existing.ID)
        }
    }

    credentials := d.Get("credentials").([]interface{})
    description := d.Get("description").(string)
    protocol := d.Get("protocol").(string)
    proxy := d.Get("proxy").([]interface{})
    resourceId := d.Get("resource_id").(string)
    serviceFabricCluster := d.Get("service_fabric_cluster").([]interface{})
    title := d.Get("title").(string)
    tls := d.Get("tls").([]interface{})
    url := d.Get("url").(string)

    parameters := apimanagement.BackendUpdateParameters{
        BackendUpdateParameterProperties: &apimanagement.BackendUpdateParameterProperties{
            Credentials: expandArmBackendBackendCredentialsContract(credentials),
            Description: utils.String(description),
            BackendProperties: &apimanagement.BackendProperties{
                ServiceFabricCluster: expandArmBackendBackendServiceFabricClusterProperties(serviceFabricCluster),
            },
            Protocol: apimanagement.BackendProtocol(protocol),
            Proxy: expandArmBackendBackendProxyContract(proxy),
            ResourceID: utils.String(resourceId),
            Title: utils.String(title),
            Tls: expandArmBackendBackendTlsProperties(tls),
            URL: utils.String(url),
        },
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroup, name, backendid, parameters); err != nil {
        return fmt.Errorf("Error creating Backend %q (Backendid %q / Resource Group %q): %+v", name, backendid, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, name, backendid)
    if err != nil {
        return fmt.Errorf("Error retrieving Backend %q (Backendid %q / Resource Group %q): %+v", name, backendid, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Backend %q (Backendid %q / Resource Group %q) ID", name, backendid, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmBackendRead(d, meta)
}

func resourceArmBackendRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).backendClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["service"]
    backendid := id.Path["backends"]

    resp, err := client.Get(ctx, resourceGroup, name, backendid)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Backend %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Backend %q (Backendid %q / Resource Group %q): %+v", name, backendid, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    d.Set("backendid", backendid)
    if backendUpdateParameterProperties := resp.BackendUpdateParameterProperties; backendUpdateParameterProperties != nil {
        if err := d.Set("credentials", flattenArmBackendBackendCredentialsContract(backendUpdateParameterProperties.Credentials)); err != nil {
            return fmt.Errorf("Error setting `credentials`: %+v", err)
        }
        d.Set("description", backendUpdateParameterProperties.Description)
        d.Set("protocol", string(backendUpdateParameterProperties.Protocol))
        if err := d.Set("proxy", flattenArmBackendBackendProxyContract(backendUpdateParameterProperties.Proxy)); err != nil {
            return fmt.Errorf("Error setting `proxy`: %+v", err)
        }
        d.Set("resource_id", backendUpdateParameterProperties.ResourceID)
        if backendProperties := backendUpdateParameterProperties.BackendProperties; backendProperties != nil {
            if err := d.Set("service_fabric_cluster", flattenArmBackendBackendServiceFabricClusterProperties(backendProperties.ServiceFabricCluster)); err != nil {
                return fmt.Errorf("Error setting `service_fabric_cluster`: %+v", err)
            }
        }
        d.Set("title", backendUpdateParameterProperties.Title)
        if err := d.Set("tls", flattenArmBackendBackendTlsProperties(backendUpdateParameterProperties.Tls)); err != nil {
            return fmt.Errorf("Error setting `tls`: %+v", err)
        }
        d.Set("url", backendUpdateParameterProperties.URL)
    }
    d.Set("type", resp.Type)

    return nil
}

func resourceArmBackendUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).backendClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    backendid := d.Get("backendid").(string)
    credentials := d.Get("credentials").([]interface{})
    description := d.Get("description").(string)
    protocol := d.Get("protocol").(string)
    proxy := d.Get("proxy").([]interface{})
    resourceId := d.Get("resource_id").(string)
    serviceFabricCluster := d.Get("service_fabric_cluster").([]interface{})
    title := d.Get("title").(string)
    tls := d.Get("tls").([]interface{})
    url := d.Get("url").(string)

    parameters := apimanagement.BackendUpdateParameters{
        BackendUpdateParameterProperties: &apimanagement.BackendUpdateParameterProperties{
            Credentials: expandArmBackendBackendCredentialsContract(credentials),
            Description: utils.String(description),
            BackendProperties: &apimanagement.BackendProperties{
                ServiceFabricCluster: expandArmBackendBackendServiceFabricClusterProperties(serviceFabricCluster),
            },
            Protocol: apimanagement.BackendProtocol(protocol),
            Proxy: expandArmBackendBackendProxyContract(proxy),
            ResourceID: utils.String(resourceId),
            Title: utils.String(title),
            Tls: expandArmBackendBackendTlsProperties(tls),
            URL: utils.String(url),
        },
    }


    if _, err := client.Update(ctx, resourceGroup, name, backendid, parameters); err != nil {
        return fmt.Errorf("Error updating Backend %q (Backendid %q / Resource Group %q): %+v", name, backendid, resourceGroup, err)
    }

    return resourceArmBackendRead(d, meta)
}

func resourceArmBackendDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).backendClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["service"]
    backendid := id.Path["backends"]

    if _, err := client.Delete(ctx, resourceGroup, name, backendid); err != nil {
        return fmt.Errorf("Error deleting Backend %q (Backendid %q / Resource Group %q): %+v", name, backendid, resourceGroup, err)
    }

    return nil
}

func expandArmBackendBackendCredentialsContract(input []interface{}) *apimanagement.BackendCredentialsContract {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    certificate := v["certificate"].([]interface{})
    query := v["query"].(map[string]interface{})
    header := v["header"].(map[string]interface{})
    authorization := v["authorization"].([]interface{})

    result := apimanagement.BackendCredentialsContract{
        Authorization: expandArmBackendBackendAuthorizationHeaderCredentials(authorization),
        Certificate: utils.ExpandStringSlice(certificate),
        Header: utils.ExpandKeyValuePairs(header),
        Query: utils.ExpandKeyValuePairs(query),
    }
    return &result
}

func expandArmBackendBackendServiceFabricClusterProperties(input []interface{}) *apimanagement.BackendServiceFabricClusterProperties {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    clientCertificatethumbprint := v["client_certificatethumbprint"].(string)
    maxPartitionResolutionRetries := v["max_partition_resolution_retries"].(int)
    managementEndpoints := v["management_endpoints"].([]interface{})
    serverCertificateThumbprints := v["server_certificate_thumbprints"].([]interface{})
    serverX509names := v["server_x509names"].([]interface{})

    result := apimanagement.BackendServiceFabricClusterProperties{
        ClientCertificatethumbprint: utils.String(clientCertificatethumbprint),
        ManagementEndpoints: utils.ExpandStringSlice(managementEndpoints),
        MaxPartitionResolutionRetries: utils.Int32(int32(maxPartitionResolutionRetries)),
        ServerCertificateThumbprints: utils.ExpandStringSlice(serverCertificateThumbprints),
        ServerX509names: expandArmBackendX509CertificateName(serverX509names),
    }
    return &result
}

func expandArmBackendBackendProxyContract(input []interface{}) *apimanagement.BackendProxyContract {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    url := v["url"].(string)
    username := v["username"].(string)
    password := v["password"].(string)

    result := apimanagement.BackendProxyContract{
        Password: utils.String(password),
        URL: utils.String(url),
        Username: utils.String(username),
    }
    return &result
}

func expandArmBackendBackendTlsProperties(input []interface{}) *apimanagement.BackendTlsProperties {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    validateCertificateChain := v["validate_certificate_chain"].(bool)
    validateCertificateName := v["validate_certificate_name"].(bool)

    result := apimanagement.BackendTlsProperties{
        ValidateCertificateChain: utils.Bool(validateCertificateChain),
        ValidateCertificateName: utils.Bool(validateCertificateName),
    }
    return &result
}

func expandArmBackendBackendAuthorizationHeaderCredentials(input []interface{}) *apimanagement.BackendAuthorizationHeaderCredentials {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    scheme := v["scheme"].(string)
    parameter := v["parameter"].(string)

    result := apimanagement.BackendAuthorizationHeaderCredentials{
        Parameter: utils.String(parameter),
        Scheme: utils.String(scheme),
    }
    return &result
}

func expandArmBackendX509CertificateName(input []interface{}) *[]apimanagement.X509CertificateName {
    results := make([]apimanagement.X509CertificateName, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        name := v["name"].(string)
        issuerCertificateThumbprint := v["issuer_certificate_thumbprint"].(string)

        result := apimanagement.X509CertificateName{
            IssuerCertificateThumbprint: utils.String(issuerCertificateThumbprint),
            Name: utils.String(name),
        }

        results = append(results, result)
    }
    return &results
}


func flattenArmBackendBackendCredentialsContract(input *apimanagement.BackendCredentialsContract) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    result["authorization"] = flattenArmBackendBackendAuthorizationHeaderCredentials(input.Authorization)
    result["certificate"] = utils.FlattenStringSlice(input.Certificate)
    result["header"] = utils.FlattenKeyValuePairs(input.Header)
    result["query"] = utils.FlattenKeyValuePairs(input.Query)

    return []interface{}{result}
}

func flattenArmBackendBackendProxyContract(input *apimanagement.BackendProxyContract) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if password := input.Password; password != nil {
        result["password"] = *password
    }
    if url := input.URL; url != nil {
        result["url"] = *url
    }
    if username := input.Username; username != nil {
        result["username"] = *username
    }

    return []interface{}{result}
}

func flattenArmBackendBackendServiceFabricClusterProperties(input *apimanagement.BackendServiceFabricClusterProperties) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if clientCertificatethumbprint := input.ClientCertificatethumbprint; clientCertificatethumbprint != nil {
        result["client_certificatethumbprint"] = *clientCertificatethumbprint
    }
    result["management_endpoints"] = utils.FlattenStringSlice(input.ManagementEndpoints)
    if maxPartitionResolutionRetries := input.MaxPartitionResolutionRetries; maxPartitionResolutionRetries != nil {
        result["max_partition_resolution_retries"] = int(*maxPartitionResolutionRetries)
    }
    result["server_certificate_thumbprints"] = utils.FlattenStringSlice(input.ServerCertificateThumbprints)
    result["server_x509names"] = flattenArmBackendX509CertificateName(input.ServerX509names)

    return []interface{}{result}
}

func flattenArmBackendBackendTlsProperties(input *apimanagement.BackendTlsProperties) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if validateCertificateChain := input.ValidateCertificateChain; validateCertificateChain != nil {
        result["validate_certificate_chain"] = *validateCertificateChain
    }
    if validateCertificateName := input.ValidateCertificateName; validateCertificateName != nil {
        result["validate_certificate_name"] = *validateCertificateName
    }

    return []interface{}{result}
}

func flattenArmBackendBackendAuthorizationHeaderCredentials(input *apimanagement.BackendAuthorizationHeaderCredentials) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if parameter := input.Parameter; parameter != nil {
        result["parameter"] = *parameter
    }
    if scheme := input.Scheme; scheme != nil {
        result["scheme"] = *scheme
    }

    return []interface{}{result}
}

func flattenArmBackendX509CertificateName(input *[]apimanagement.X509CertificateName) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        if name := item.Name; name != nil {
            v["name"] = *name
        }
        if issuerCertificateThumbprint := item.IssuerCertificateThumbprint; issuerCertificateThumbprint != nil {
            v["issuer_certificate_thumbprint"] = *issuerCertificateThumbprint
        }

        results = append(results, v)
    }

    return results
}
