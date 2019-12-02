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

            "backend_id": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "after": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validateIso8601Duration(),
            },

            "credentials": {
                Type: schema.TypeList,
                Computed: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "authorization": {
                            Type: schema.TypeList,
                            Computed: true,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "parameter": {
                                        Type: schema.TypeString,
                                        Computed: true,
                                    },
                                    "scheme": {
                                        Type: schema.TypeString,
                                        Computed: true,
                                    },
                                },
                            },
                        },
                        "certificate": {
                            Type: schema.TypeList,
                            Computed: true,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                        "header": {
                            Type: schema.TypeMap,
                            Computed: true,
                            Elem: &schema.Schema{Type: schema.TypeString},
                        },
                        "query": {
                            Type: schema.TypeMap,
                            Computed: true,
                            Elem: &schema.Schema{Type: schema.TypeString},
                        },
                    },
                },
            },

            "description": {
                Type: schema.TypeString,
                Computed: true,
            },

            "protocol": {
                Type: schema.TypeString,
                Computed: true,
            },

            "proxy": {
                Type: schema.TypeList,
                Computed: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "password": {
                            Type: schema.TypeString,
                            Computed: true,
                        },
                        "url": {
                            Type: schema.TypeString,
                            Computed: true,
                        },
                        "username": {
                            Type: schema.TypeString,
                            Computed: true,
                        },
                    },
                },
            },

            "resource_id": {
                Type: schema.TypeString,
                Computed: true,
            },

            "service_fabric_cluster": {
                Type: schema.TypeList,
                Computed: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "client_certificatethumbprint": {
                            Type: schema.TypeString,
                            Computed: true,
                        },
                        "management_endpoints": {
                            Type: schema.TypeList,
                            Computed: true,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                        "max_partition_resolution_retries": {
                            Type: schema.TypeInt,
                            Computed: true,
                        },
                        "server_certificate_thumbprints": {
                            Type: schema.TypeList,
                            Computed: true,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                        "server_x509names": {
                            Type: schema.TypeList,
                            Computed: true,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "issuer_certificate_thumbprint": {
                                        Type: schema.TypeString,
                                        Computed: true,
                                    },
                                    "name": {
                                        Type: schema.TypeString,
                                        Computed: true,
                                    },
                                },
                            },
                        },
                    },
                },
            },

            "title": {
                Type: schema.TypeString,
                Computed: true,
            },

            "tls": {
                Type: schema.TypeList,
                Computed: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "validate_certificate_chain": {
                            Type: schema.TypeBool,
                            Computed: true,
                        },
                        "validate_certificate_name": {
                            Type: schema.TypeBool,
                            Computed: true,
                        },
                    },
                },
            },

            "type": {
                Type: schema.TypeString,
                Computed: true,
            },

            "url": {
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
    backendID := d.Get("backend_id").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, name, backendID)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Backend %q (Backend %q / Resource Group %q): %+v", name, backendID, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_backend", *existing.ID)
        }
    }

    after := d.Get("after").(string)

    parameters := apimanagement.BackendReconnectContract{
        BackendReconnectProperties: &apimanagement.BackendReconnectProperties{
            After: utils.String(after),
        },
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroup, name, backendID, parameters); err != nil {
        return fmt.Errorf("Error creating Backend %q (Backend %q / Resource Group %q): %+v", name, backendID, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, name, backendID)
    if err != nil {
        return fmt.Errorf("Error retrieving Backend %q (Backend %q / Resource Group %q): %+v", name, backendID, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Backend %q (Backend %q / Resource Group %q) ID", name, backendID, resourceGroup)
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
    backendID := id.Path["backends"]

    resp, err := client.Get(ctx, resourceGroup, name, backendID)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Backend %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Backend %q (Backend %q / Resource Group %q): %+v", name, backendID, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    d.Set("backend_id", backendID)
    if backendReconnectProperties := resp.BackendReconnectProperties; backendReconnectProperties != nil {
        if err := d.Set("credentials", flattenArmBackendBackendCredentialsContract(backendReconnectProperties.Credentials)); err != nil {
            return fmt.Errorf("Error setting `credentials`: %+v", err)
        }
        d.Set("description", backendReconnectProperties.Description)
        d.Set("protocol", string(backendReconnectProperties.Protocol))
        if err := d.Set("proxy", flattenArmBackendBackendProxyContract(backendReconnectProperties.Proxy)); err != nil {
            return fmt.Errorf("Error setting `proxy`: %+v", err)
        }
        d.Set("resource_id", backendReconnectProperties.ResourceID)
        if backendProperties := backendReconnectProperties.BackendProperties; backendProperties != nil {
            if err := d.Set("service_fabric_cluster", flattenArmBackendBackendServiceFabricClusterProperties(backendProperties.ServiceFabricCluster)); err != nil {
                return fmt.Errorf("Error setting `service_fabric_cluster`: %+v", err)
            }
        }
        d.Set("title", backendReconnectProperties.Title)
        if err := d.Set("tls", flattenArmBackendBackendTlsProperties(backendReconnectProperties.TLS)); err != nil {
            return fmt.Errorf("Error setting `tls`: %+v", err)
        }
        d.Set("url", backendReconnectProperties.URL)
    }
    d.Set("type", resp.Type)

    return nil
}

func resourceArmBackendUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).backendClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    after := d.Get("after").(string)
    backendID := d.Get("backend_id").(string)

    parameters := apimanagement.BackendReconnectContract{
        BackendReconnectProperties: &apimanagement.BackendReconnectProperties{
            After: utils.String(after),
        },
    }


    if _, err := client.Update(ctx, resourceGroup, name, backendID, parameters); err != nil {
        return fmt.Errorf("Error updating Backend %q (Backend %q / Resource Group %q): %+v", name, backendID, resourceGroup, err)
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
    backendID := id.Path["backends"]

    if _, err := client.Delete(ctx, resourceGroup, name, backendID); err != nil {
        return fmt.Errorf("Error deleting Backend %q (Backend %q / Resource Group %q): %+v", name, backendID, resourceGroup, err)
    }

    return nil
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
    result["server_x509names"] = flattenArmBackendX509CertificateName(input.ServerX509Names)

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
