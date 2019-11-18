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
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "name": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
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

            "additional_locations": {
                Type: schema.TypeList,
                Optional: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "location": azure.SchemaLocation(),
                        "sku_type": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(apimanagement.Developer),
                                string(apimanagement.Standard),
                                string(apimanagement.Premium),
                            }, false),
                        },
                        "sku_unit_count": {
                            Type: schema.TypeInt,
                            Optional: true,
                        },
                        "vpnconfiguration": {
                            Type: schema.TypeList,
                            Optional: true,
                            MaxItems: 1,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "location": azure.SchemaLocation(),
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

            "addresser_email": {
                Type: schema.TypeString,
                Optional: true,
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
                        "certificate": {
                            Type: schema.TypeList,
                            Required: true,
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
                        "hostname": {
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
                    },
                },
            },

            "sku": {
                Type: schema.TypeList,
                Optional: true,
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
                            }, false),
                        },
                        "capacity": {
                            Type: schema.TypeInt,
                            Optional: true,
                        },
                    },
                },
            },

            "vpn_type": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(apimanagement.None),
                    string(apimanagement.External),
                    string(apimanagement.Internal),
                }, false),
                Default: string(apimanagement.None),
            },

            "vpnconfiguration": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "location": azure.SchemaLocation(),
                        "subnet_resource_id": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },

            "created_at_utc": {
                Type: schema.TypeString,
                Computed: true,
            },

            "etag": {
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

            "provisioning_state": {
                Type: schema.TypeString,
                Computed: true,
            },

            "runtime_url": {
                Type: schema.TypeString,
                Computed: true,
            },

            "scm_url": {
                Type: schema.TypeString,
                Computed: true,
            },

            "static_ips": {
                Type: schema.TypeList,
                Computed: true,
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
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
    client := meta.(*ArmClient).apiManagementServicesClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Api Management Service %q (Resource Group %q): %+v", name, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_api_management_service", *existing.ID)
        }
    }

    name := d.Get("name").(string)
    location := azure.NormalizeLocation(d.Get("location").(string))
    additionalLocations := d.Get("additional_locations").([]interface{})
    addresserEmail := d.Get("addresser_email").(string)
    customProperties := d.Get("custom_properties").(map[string]interface{})
    hostnameConfigurations := d.Get("hostname_configurations").([]interface{})
    publisherEmail := d.Get("publisher_email").(string)
    publisherName := d.Get("publisher_name").(string)
    sku := d.Get("sku").([]interface{})
    vpnType := d.Get("vpn_type").(string)
    vpnconfiguration := d.Get("vpnconfiguration").([]interface{})
    t := d.Get("tags").(map[string]interface{})

    parameters := apimanagement.ServiceUpdateParameters{
        Location: utils.String(location),
        Name: utils.String(name),
        ServiceProperties: &apimanagement.ServiceProperties{
            AdditionalLocations: expandArmApiManagementServiceAdditionalRegion(additionalLocations),
            AddresserEmail: utils.String(addresserEmail),
            CustomProperties: utils.ExpandKeyValuePairs(customProperties),
            HostnameConfigurations: expandArmApiManagementServiceHostnameConfiguration(hostnameConfigurations),
            PublisherEmail: utils.String(publisherEmail),
            PublisherName: utils.String(publisherName),
            VpnType: apimanagement.VirtualNetworkType(vpnType),
            Vpnconfiguration: expandArmApiManagementServiceVirtualNetworkConfiguration(vpnconfiguration),
        },
        Sku: expandArmApiManagementServiceServiceSkuProperties(sku),
        Tags: tags.Expand(t),
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroup, name, parameters); err != nil {
        return fmt.Errorf("Error creating Api Management Service %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Api Management Service %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Api Management Service %q (Resource Group %q) ID", name, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmApiManagementServiceRead(d, meta)
}

func resourceArmApiManagementServiceRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).apiManagementServicesClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["service"]

    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Api Management Service %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Api Management Service %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if serviceProperties := resp.ServiceProperties; serviceProperties != nil {
        if err := d.Set("additional_locations", flattenArmApiManagementServiceAdditionalRegion(serviceProperties.AdditionalLocations)); err != nil {
            return fmt.Errorf("Error setting `additional_locations`: %+v", err)
        }
        d.Set("addresser_email", serviceProperties.AddresserEmail)
        d.Set("created_at_utc", (serviceProperties.CreatedAtUtc).String())
        d.Set("custom_properties", utils.FlattenKeyValuePairs(serviceProperties.CustomProperties))
        if err := d.Set("hostname_configurations", flattenArmApiManagementServiceHostnameConfiguration(serviceProperties.HostnameConfigurations)); err != nil {
            return fmt.Errorf("Error setting `hostname_configurations`: %+v", err)
        }
        d.Set("management_api_url", serviceProperties.ManagementApiURL)
        d.Set("portal_url", serviceProperties.PortalURL)
        d.Set("provisioning_state", serviceProperties.ProvisioningState)
        d.Set("publisher_email", serviceProperties.PublisherEmail)
        d.Set("publisher_name", serviceProperties.PublisherName)
        d.Set("runtime_url", serviceProperties.RuntimeURL)
        d.Set("scm_url", serviceProperties.ScmURL)
        d.Set("static_ips", utils.FlattenStringSlice(serviceProperties.StaticIps))
        d.Set("target_provisioning_state", serviceProperties.TargetProvisioningState)
        d.Set("vpn_type", string(serviceProperties.VpnType))
        if err := d.Set("vpnconfiguration", flattenArmApiManagementServiceVirtualNetworkConfiguration(serviceProperties.Vpnconfiguration)); err != nil {
            return fmt.Errorf("Error setting `vpnconfiguration`: %+v", err)
        }
    }
    d.Set("etag", resp.Etag)
    if err := d.Set("sku", flattenArmApiManagementServiceServiceSkuProperties(resp.Sku)); err != nil {
        return fmt.Errorf("Error setting `sku`: %+v", err)
    }
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}

func resourceArmApiManagementServiceUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).apiManagementServicesClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    additionalLocations := d.Get("additional_locations").([]interface{})
    addresserEmail := d.Get("addresser_email").(string)
    customProperties := d.Get("custom_properties").(map[string]interface{})
    hostnameConfigurations := d.Get("hostname_configurations").([]interface{})
    publisherEmail := d.Get("publisher_email").(string)
    publisherName := d.Get("publisher_name").(string)
    sku := d.Get("sku").([]interface{})
    vpnType := d.Get("vpn_type").(string)
    vpnconfiguration := d.Get("vpnconfiguration").([]interface{})
    t := d.Get("tags").(map[string]interface{})

    parameters := apimanagement.ServiceUpdateParameters{
        Location: utils.String(location),
        Name: utils.String(name),
        ServiceProperties: &apimanagement.ServiceProperties{
            AdditionalLocations: expandArmApiManagementServiceAdditionalRegion(additionalLocations),
            AddresserEmail: utils.String(addresserEmail),
            CustomProperties: utils.ExpandKeyValuePairs(customProperties),
            HostnameConfigurations: expandArmApiManagementServiceHostnameConfiguration(hostnameConfigurations),
            PublisherEmail: utils.String(publisherEmail),
            PublisherName: utils.String(publisherName),
            VpnType: apimanagement.VirtualNetworkType(vpnType),
            Vpnconfiguration: expandArmApiManagementServiceVirtualNetworkConfiguration(vpnconfiguration),
        },
        Sku: expandArmApiManagementServiceServiceSkuProperties(sku),
        Tags: tags.Expand(t),
    }


    future, err := client.Update(ctx, resourceGroup, name, parameters)
    if err != nil {
        return fmt.Errorf("Error updating Api Management Service %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for update of Api Management Service %q (Resource Group %q): %+v", name, resourceGroup, err)
    }

    return resourceArmApiManagementServiceRead(d, meta)
}

func resourceArmApiManagementServiceDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).apiManagementServicesClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["service"]

    if _, err := client.Delete(ctx, resourceGroup, name); err != nil {
        return fmt.Errorf("Error deleting Api Management Service %q (Resource Group %q): %+v", name, resourceGroup, err)
    }

    return nil
}

func expandArmApiManagementServiceAdditionalRegion(input []interface{}) *[]apimanagement.AdditionalRegion {
    results := make([]apimanagement.AdditionalRegion, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        location := azure.NormalizeLocation(v["location"].(string))
        skuType := v["sku_type"].(string)
        skuUnitCount := v["sku_unit_count"].(int)
        vpnconfiguration := v["vpnconfiguration"].([]interface{})

        result := apimanagement.AdditionalRegion{
            Location: utils.String(location),
            SkuType: apimanagement.SkuType(skuType),
            SkuUnitCount: utils.Int32(int32(skuUnitCount)),
            Vpnconfiguration: expandArmApiManagementServiceVirtualNetworkConfiguration(vpnconfiguration),
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
        hostname := v["hostname"].(string)
        certificate := v["certificate"].([]interface{})

        result := apimanagement.HostnameConfiguration{
            Certificate: expandArmApiManagementServiceCertificateInformation(certificate),
            Hostname: utils.String(hostname),
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
    location := azure.NormalizeLocation(v["location"].(string))

    result := apimanagement.VirtualNetworkConfiguration{
        Location: utils.String(location),
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


func flattenArmApiManagementServiceAdditionalRegion(input *[]apimanagement.AdditionalRegion) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        if location := item.Location; location != nil {
            v["location"] = azure.NormalizeLocation(*location)
        }
        v["sku_type"] = string(item.SkuType)
        if skuUnitCount := item.SkuUnitCount; skuUnitCount != nil {
            v["sku_unit_count"] = int(*skuUnitCount)
        }
        v["vpnconfiguration"] = flattenArmApiManagementServiceVirtualNetworkConfiguration(item.Vpnconfiguration)

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
        if hostname := item.Hostname; hostname != nil {
            v["hostname"] = *hostname
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

    if location := input.Location; location != nil {
        result["location"] = azure.NormalizeLocation(*location)
    }
    if subnetResourceId := input.SubnetResourceID; subnetResourceId != nil {
        result["subnet_resource_id"] = *subnetResourceId
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
