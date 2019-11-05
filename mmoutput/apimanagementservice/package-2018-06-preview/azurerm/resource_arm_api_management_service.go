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



func resourceArmApiManagementService() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmApiManagementServiceCreate,
        Read: resourceArmApiManagementServiceRead,
        Update: resourceArmApiManagementServiceUpdate,
        Delete: resourceArmApiManagementServiceDelete,

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

            "publisher_email": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "publisher_name": {
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

            "sku": {
                Type: schema.TypeList,
                Required: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "name": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(apimanagement.Developer),
                                string(apimanagement.Standard),
                                string(apimanagement.Premium),
                                string(apimanagement.Basic),
                                string(apimanagement.Consumption),
                            }, false),
                        },
                        "capacity": {
                            Type: schema.TypeInt,
                            Optional: true,
                        },
                    },
                },
            },

            "additional_locations": {
                Type: schema.TypeList,
                Optional: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "location": azure.SchemaLocation(),
                        "sku": {
                            Type: schema.TypeList,
                            Required: true,
                            MaxItems: 1,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "name": {
                                        Type: schema.TypeString,
                                        Required: true,
                                        ValidateFunc: validation.StringInSlice([]string{
                                            string(apimanagement.Developer),
                                            string(apimanagement.Standard),
                                            string(apimanagement.Premium),
                                            string(apimanagement.Basic),
                                            string(apimanagement.Consumption),
                                        }, false),
                                    },
                                    "capacity": {
                                        Type: schema.TypeInt,
                                        Optional: true,
                                    },
                                },
                            },
                        },
                        "virtual_network_configuration": {
                            Type: schema.TypeList,
                            Optional: true,
                            MaxItems: 1,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "subnet_resource_id": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                },
                            },
                        },
                    },
                },
            },

            "certificates": {
                Type: schema.TypeList,
                Optional: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "store_name": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(apimanagement.CertificateAuthority),
                                string(apimanagement.Root),
                            }, false),
                        },
                        "certificate": {
                            Type: schema.TypeList,
                            Optional: true,
                            MaxItems: 1,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "expiry": {
                                        Type: schema.TypeString,
                                        Required: true,
                                        ValidateFunc: validate.NoEmptyStrings,
                                        ValidateFunc: validateRFC3339Date,
                                    },
                                    "subject": {
                                        Type: schema.TypeString,
                                        Required: true,
                                        ValidateFunc: validate.NoEmptyStrings,
                                    },
                                    "thumbprint": {
                                        Type: schema.TypeString,
                                        Required: true,
                                        ValidateFunc: validate.NoEmptyStrings,
                                    },
                                },
                            },
                        },
                        "certificate_password": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "encoded_certificate": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },

            "custom_properties": {
                Type: schema.TypeMap,
                Optional: true,
                Elem: &schema.Schema{Type: schema.TypeString},
            },

            "hostname_configurations": {
                Type: schema.TypeList,
                Optional: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "host_name": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "type": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(apimanagement.Proxy),
                                string(apimanagement.Portal),
                                string(apimanagement.Management),
                                string(apimanagement.Scm),
                            }, false),
                        },
                        "certificate": {
                            Type: schema.TypeList,
                            Optional: true,
                            MaxItems: 1,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "expiry": {
                                        Type: schema.TypeString,
                                        Required: true,
                                        ValidateFunc: validate.NoEmptyStrings,
                                        ValidateFunc: validateRFC3339Date,
                                    },
                                    "subject": {
                                        Type: schema.TypeString,
                                        Required: true,
                                        ValidateFunc: validate.NoEmptyStrings,
                                    },
                                    "thumbprint": {
                                        Type: schema.TypeString,
                                        Required: true,
                                        ValidateFunc: validate.NoEmptyStrings,
                                    },
                                },
                            },
                        },
                        "certificate_password": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "default_ssl_binding": {
                            Type: schema.TypeBool,
                            Optional: true,
                        },
                        "encoded_certificate": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "key_vault_id": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "negotiate_client_certificate": {
                            Type: schema.TypeBool,
                            Optional: true,
                        },
                    },
                },
            },

            "identity": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "type": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                    },
                },
            },

            "notification_sender_email": {
                Type: schema.TypeString,
                Optional: true,
            },

            "virtual_network_configuration": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "subnet_resource_id": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },

            "virtual_network_type": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(apimanagement.None),
                    string(apimanagement.External),
                    string(apimanagement.Internal),
                }, false),
                Default: string(apimanagement.None),
            },

            "created_at_utc": {
                Type: schema.TypeString,
                Computed: true,
            },

            "etag": {
                Type: schema.TypeString,
                Computed: true,
            },

            "gateway_regional_url": {
                Type: schema.TypeString,
                Computed: true,
            },

            "gateway_url": {
                Type: schema.TypeString,
                Computed: true,
            },

            "management_api_url": {
                Type: schema.TypeString,
                Computed: true,
            },

            "portal_url": {
                Type: schema.TypeString,
                Computed: true,
            },

            "private_ip_addresses": {
                Type: schema.TypeList,
                Computed: true,
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
            },

            "provisioning_state": {
                Type: schema.TypeString,
                Computed: true,
            },

            "public_ip_addresses": {
                Type: schema.TypeList,
                Computed: true,
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
            },

            "scm_url": {
                Type: schema.TypeString,
                Computed: true,
            },

            "target_provisioning_state": {
                Type: schema.TypeString,
                Computed: true,
            },

            "type": {
                Type: schema.TypeString,
                Computed: true,
            },

            "tags": tags.Schema(),
        },
    }
}

func resourceArmApiManagementServiceCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).apiManagementServiceClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    serviceName := d.Get("service_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, serviceName)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Api Management Service (Service Name %q / Resource Group %q): %+v", serviceName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_api_management_service", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    additionalLocations := d.Get("additional_locations").([]interface{})
    certificates := d.Get("certificates").([]interface{})
    customProperties := d.Get("custom_properties").(map[string]interface{})
    hostnameConfigurations := d.Get("hostname_configurations").([]interface{})
    identity := d.Get("identity").([]interface{})
    notificationSenderEmail := d.Get("notification_sender_email").(string)
    publisherEmail := d.Get("publisher_email").(string)
    publisherName := d.Get("publisher_name").(string)
    sku := d.Get("sku").([]interface{})
    virtualNetworkConfiguration := d.Get("virtual_network_configuration").([]interface{})
    virtualNetworkType := d.Get("virtual_network_type").(string)
    t := d.Get("tags").(map[string]interface{})

    parameters := apimanagement.ServiceResource{
        Identity: expandArmApiManagementServiceServiceIdentity(identity),
        Location: utils.String(location),
        ServiceProperties: &apimanagement.ServiceProperties{
            AdditionalLocations: expandArmApiManagementServiceAdditionalLocation(additionalLocations),
            Certificates: expandArmApiManagementServiceCertificateConfiguration(certificates),
            CustomProperties: utils.ExpandKeyValuePairs(customProperties),
            HostnameConfigurations: expandArmApiManagementServiceHostnameConfiguration(hostnameConfigurations),
            NotificationSenderEmail: utils.String(notificationSenderEmail),
            PublisherEmail: utils.String(publisherEmail),
            PublisherName: utils.String(publisherName),
            VirtualNetworkConfiguration: expandArmApiManagementServiceVirtualNetworkConfiguration(virtualNetworkConfiguration),
            VirtualNetworkType: apimanagement.VirtualNetworkType(virtualNetworkType),
        },
        Sku: expandArmApiManagementServiceServiceSkuProperties(sku),
        Tags: tags.Expand(t),
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, serviceName, parameters)
    if err != nil {
        return fmt.Errorf("Error creating Api Management Service (Service Name %q / Resource Group %q): %+v", serviceName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Api Management Service (Service Name %q / Resource Group %q): %+v", serviceName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, serviceName)
    if err != nil {
        return fmt.Errorf("Error retrieving Api Management Service (Service Name %q / Resource Group %q): %+v", serviceName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Api Management Service (Service Name %q / Resource Group %q) ID", serviceName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmApiManagementServiceRead(d, meta)
}

func resourceArmApiManagementServiceRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).apiManagementServiceClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    serviceName := id.Path["service"]

    resp, err := client.Get(ctx, resourceGroup, serviceName)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Api Management Service %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Api Management Service (Service Name %q / Resource Group %q): %+v", serviceName, resourceGroup, err)
    }


    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if serviceProperties := resp.ServiceProperties; serviceProperties != nil {
        if err := d.Set("additional_locations", flattenArmApiManagementServiceAdditionalLocation(serviceProperties.AdditionalLocations)); err != nil {
            return fmt.Errorf("Error setting `additional_locations`: %+v", err)
        }
        if err := d.Set("certificates", flattenArmApiManagementServiceCertificateConfiguration(serviceProperties.Certificates)); err != nil {
            return fmt.Errorf("Error setting `certificates`: %+v", err)
        }
        d.Set("created_at_utc", (serviceProperties.CreatedAtUtc).String())
        d.Set("custom_properties", utils.FlattenKeyValuePairs(serviceProperties.CustomProperties))
        d.Set("gateway_regional_url", serviceProperties.GatewayRegionalURL)
        d.Set("gateway_url", serviceProperties.GatewayURL)
        if err := d.Set("hostname_configurations", flattenArmApiManagementServiceHostnameConfiguration(serviceProperties.HostnameConfigurations)); err != nil {
            return fmt.Errorf("Error setting `hostname_configurations`: %+v", err)
        }
        d.Set("management_api_url", serviceProperties.ManagementApiURL)
        d.Set("notification_sender_email", serviceProperties.NotificationSenderEmail)
        d.Set("portal_url", serviceProperties.PortalURL)
        d.Set("private_ip_addresses", utils.FlattenStringSlice(serviceProperties.PrivateIpAddresses))
        d.Set("provisioning_state", serviceProperties.ProvisioningState)
        d.Set("public_ip_addresses", utils.FlattenStringSlice(serviceProperties.PublicIpAddresses))
        d.Set("publisher_email", serviceProperties.PublisherEmail)
        d.Set("publisher_name", serviceProperties.PublisherName)
        d.Set("scm_url", serviceProperties.ScmURL)
        d.Set("target_provisioning_state", serviceProperties.TargetProvisioningState)
        if err := d.Set("virtual_network_configuration", flattenArmApiManagementServiceVirtualNetworkConfiguration(serviceProperties.VirtualNetworkConfiguration)); err != nil {
            return fmt.Errorf("Error setting `virtual_network_configuration`: %+v", err)
        }
        d.Set("virtual_network_type", string(serviceProperties.VirtualNetworkType))
    }
    d.Set("etag", resp.Etag)
    if err := d.Set("identity", flattenArmApiManagementServiceServiceIdentity(resp.Identity)); err != nil {
        return fmt.Errorf("Error setting `identity`: %+v", err)
    }
    d.Set("service_name", serviceName)
    if err := d.Set("sku", flattenArmApiManagementServiceServiceSkuProperties(resp.Sku)); err != nil {
        return fmt.Errorf("Error setting `sku`: %+v", err)
    }
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}

func resourceArmApiManagementServiceUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).apiManagementServiceClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    additionalLocations := d.Get("additional_locations").([]interface{})
    certificates := d.Get("certificates").([]interface{})
    customProperties := d.Get("custom_properties").(map[string]interface{})
    hostnameConfigurations := d.Get("hostname_configurations").([]interface{})
    identity := d.Get("identity").([]interface{})
    notificationSenderEmail := d.Get("notification_sender_email").(string)
    publisherEmail := d.Get("publisher_email").(string)
    publisherName := d.Get("publisher_name").(string)
    serviceName := d.Get("service_name").(string)
    sku := d.Get("sku").([]interface{})
    virtualNetworkConfiguration := d.Get("virtual_network_configuration").([]interface{})
    virtualNetworkType := d.Get("virtual_network_type").(string)
    t := d.Get("tags").(map[string]interface{})

    parameters := apimanagement.ServiceResource{
        Identity: expandArmApiManagementServiceServiceIdentity(identity),
        Location: utils.String(location),
        ServiceProperties: &apimanagement.ServiceProperties{
            AdditionalLocations: expandArmApiManagementServiceAdditionalLocation(additionalLocations),
            Certificates: expandArmApiManagementServiceCertificateConfiguration(certificates),
            CustomProperties: utils.ExpandKeyValuePairs(customProperties),
            HostnameConfigurations: expandArmApiManagementServiceHostnameConfiguration(hostnameConfigurations),
            NotificationSenderEmail: utils.String(notificationSenderEmail),
            PublisherEmail: utils.String(publisherEmail),
            PublisherName: utils.String(publisherName),
            VirtualNetworkConfiguration: expandArmApiManagementServiceVirtualNetworkConfiguration(virtualNetworkConfiguration),
            VirtualNetworkType: apimanagement.VirtualNetworkType(virtualNetworkType),
        },
        Sku: expandArmApiManagementServiceServiceSkuProperties(sku),
        Tags: tags.Expand(t),
    }


    future, err := client.Update(ctx, resourceGroup, serviceName, parameters)
    if err != nil {
        return fmt.Errorf("Error updating Api Management Service (Service Name %q / Resource Group %q): %+v", serviceName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for update of Api Management Service (Service Name %q / Resource Group %q): %+v", serviceName, resourceGroup, err)
    }

    return resourceArmApiManagementServiceRead(d, meta)
}

func resourceArmApiManagementServiceDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).apiManagementServiceClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    serviceName := id.Path["service"]

    future, err := client.Delete(ctx, resourceGroup, serviceName)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Api Management Service (Service Name %q / Resource Group %q): %+v", serviceName, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Api Management Service (Service Name %q / Resource Group %q): %+v", serviceName, resourceGroup, err)
        }
    }

    return nil
}

func expandArmApiManagementServiceServiceIdentity(input []interface{}) *apimanagement.ServiceIdentity {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    type := v["type"].(string)

    result := apimanagement.ServiceIdentity{
        Type: utils.String(type),
    }
    return &result
}

func expandArmApiManagementServiceAdditionalLocation(input []interface{}) *[]apimanagement.AdditionalLocation {
    results := make([]apimanagement.AdditionalLocation, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        location := azure.NormalizeLocation(v["location"].(string))
        sku := v["sku"].([]interface{})
        virtualNetworkConfiguration := v["virtual_network_configuration"].([]interface{})

        result := apimanagement.AdditionalLocation{
            Location: utils.String(location),
            Sku: expandArmApiManagementServiceServiceSkuProperties(sku),
            VirtualNetworkConfiguration: expandArmApiManagementServiceVirtualNetworkConfiguration(virtualNetworkConfiguration),
        }

        results = append(results, result)
    }
    return &results
}

func expandArmApiManagementServiceCertificateConfiguration(input []interface{}) *[]apimanagement.CertificateConfiguration {
    results := make([]apimanagement.CertificateConfiguration, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        encodedCertificate := v["encoded_certificate"].(string)
        certificatePassword := v["certificate_password"].(string)
        storeName := v["store_name"].(string)
        certificate := v["certificate"].([]interface{})

        result := apimanagement.CertificateConfiguration{
            Certificate: expandArmApiManagementServiceCertificateInformation(certificate),
            CertificatePassword: utils.String(certificatePassword),
            EncodedCertificate: utils.String(encodedCertificate),
            StoreName: apimanagement.(storeName),
        }

        results = append(results, result)
    }
    return &results
}

func expandArmApiManagementServiceHostnameConfiguration(input []interface{}) *[]apimanagement.HostnameConfiguration {
    results := make([]apimanagement.HostnameConfiguration, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        type := v["type"].(string)
        hostName := v["host_name"].(string)
        keyVaultId := v["key_vault_id"].(string)
        encodedCertificate := v["encoded_certificate"].(string)
        certificatePassword := v["certificate_password"].(string)
        defaultSslBinding := v["default_ssl_binding"].(bool)
        negotiateClientCertificate := v["negotiate_client_certificate"].(bool)
        certificate := v["certificate"].([]interface{})

        result := apimanagement.HostnameConfiguration{
            Certificate: expandArmApiManagementServiceCertificateInformation(certificate),
            CertificatePassword: utils.String(certificatePassword),
            DefaultSslBinding: utils.Bool(defaultSslBinding),
            EncodedCertificate: utils.String(encodedCertificate),
            HostName: utils.String(hostName),
            KeyVaultID: utils.String(keyVaultId),
            NegotiateClientCertificate: utils.Bool(negotiateClientCertificate),
            Type: apimanagement.HostnameType(type),
        }

        results = append(results, result)
    }
    return &results
}

func expandArmApiManagementServiceVirtualNetworkConfiguration(input []interface{}) *apimanagement.VirtualNetworkConfiguration {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    subnetResourceId := v["subnet_resource_id"].(string)

    result := apimanagement.VirtualNetworkConfiguration{
        SubnetResourceID: utils.String(subnetResourceId),
    }
    return &result
}

func expandArmApiManagementServiceServiceSkuProperties(input []interface{}) *apimanagement.ServiceSkuProperties {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    name := v["name"].(string)
    capacity := v["capacity"].(int)

    result := apimanagement.ServiceSkuProperties{
        Capacity: utils.Int32(int32(capacity)),
        Name: apimanagement.SkuType(name),
    }
    return &result
}

func expandArmApiManagementServiceCertificateInformation(input []interface{}) *apimanagement.CertificateInformation {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    expiry := v["expiry"].(string)
    thumbprint := v["thumbprint"].(string)
    subject := v["subject"].(string)

    result := apimanagement.CertificateInformation{
        Expiry: convertStringToDate(expiry),
        Subject: utils.String(subject),
        Thumbprint: utils.String(thumbprint),
    }
    return &result
}

func convertStringToDate(input interface{}) *date.Time {
  v := input.(string)

  dateTime, err := date.ParseTime(time.RFC3339, v)
  if err != nil {
      log.Printf("[ERROR] Cannot convert an invalid string to RFC3339 date %q: %+v", v, err)
      return nil
  }

  result := date.Time{
      Time: dateTime,
  }
  return &result
}


func flattenArmApiManagementServiceAdditionalLocation(input *[]apimanagement.AdditionalLocation) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        if location := item.Location; location != nil {
            v["location"] = azure.NormalizeLocation(*location)
        }
        v["sku"] = flattenArmApiManagementServiceServiceSkuProperties(item.Sku)
        v["virtual_network_configuration"] = flattenArmApiManagementServiceVirtualNetworkConfiguration(item.VirtualNetworkConfiguration)

        results = append(results, v)
    }

    return results
}

func flattenArmApiManagementServiceCertificateConfiguration(input *[]apimanagement.CertificateConfiguration) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        v["certificate"] = flattenArmApiManagementServiceCertificateInformation(item.Certificate)
        if certificatePassword := item.CertificatePassword; certificatePassword != nil {
            v["certificate_password"] = *certificatePassword
        }
        if encodedCertificate := item.EncodedCertificate; encodedCertificate != nil {
            v["encoded_certificate"] = *encodedCertificate
        }
        v["store_name"] = string(item.StoreName)

        results = append(results, v)
    }

    return results
}

func flattenArmApiManagementServiceHostnameConfiguration(input *[]apimanagement.HostnameConfiguration) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        v["certificate"] = flattenArmApiManagementServiceCertificateInformation(item.Certificate)
        if certificatePassword := item.CertificatePassword; certificatePassword != nil {
            v["certificate_password"] = *certificatePassword
        }
        if defaultSslBinding := item.DefaultSslBinding; defaultSslBinding != nil {
            v["default_ssl_binding"] = *defaultSslBinding
        }
        if encodedCertificate := item.EncodedCertificate; encodedCertificate != nil {
            v["encoded_certificate"] = *encodedCertificate
        }
        if hostName := item.HostName; hostName != nil {
            v["host_name"] = *hostName
        }
        if keyVaultId := item.KeyVaultID; keyVaultId != nil {
            v["key_vault_id"] = *keyVaultId
        }
        if negotiateClientCertificate := item.NegotiateClientCertificate; negotiateClientCertificate != nil {
            v["negotiate_client_certificate"] = *negotiateClientCertificate
        }
        v["type"] = string(item.Type)

        results = append(results, v)
    }

    return results
}

func flattenArmApiManagementServiceVirtualNetworkConfiguration(input *apimanagement.VirtualNetworkConfiguration) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if subnetResourceId := input.SubnetResourceID; subnetResourceId != nil {
        result["subnet_resource_id"] = *subnetResourceId
    }

    return []interface{}{result}
}

func flattenArmApiManagementServiceServiceIdentity(input *apimanagement.ServiceIdentity) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if type := input.Type; type != nil {
        result["type"] = *type
    }

    return []interface{}{result}
}

func flattenArmApiManagementServiceServiceSkuProperties(input *apimanagement.ServiceSkuProperties) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    result["name"] = string(input.Name)
    if capacity := input.Capacity; capacity != nil {
        result["capacity"] = int(*capacity)
    }

    return []interface{}{result}
}

func flattenArmApiManagementServiceCertificateInformation(input *apimanagement.CertificateInformation) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if expiry := input.Expiry; expiry != nil {
        result["expiry"] = (*expiry).String()
    }
    if subject := input.Subject; subject != nil {
        result["subject"] = *subject
    }
    if thumbprint := input.Thumbprint; thumbprint != nil {
        result["thumbprint"] = *thumbprint
    }

    return []interface{}{result}
}
