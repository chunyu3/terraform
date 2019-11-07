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



func resourceArmIotDpsResource() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmIotDpsResourceCreateUpdate,
        Read: resourceArmIotDpsResourceRead,
        Update: resourceArmIotDpsResourceCreateUpdate,
        Delete: resourceArmIotDpsResourceDelete,

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

            "location": azure.SchemaLocation(),

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "sku": {
                Type: schema.TypeList,
                Required: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "capacity": {
                            Type: schema.TypeInt,
                            Optional: true,
                        },
                        "name": {
                            Type: schema.TypeString,
                            Optional: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(deviceprovisioningservices.S1),
                            }, false),
                            Default: string(deviceprovisioningservices.S1),
                        },
                    },
                },
            },

            "allocation_policy": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(deviceprovisioningservices.Hashed),
                    string(deviceprovisioningservices.GeoLatency),
                    string(deviceprovisioningservices.Static),
                }, false),
                Default: string(deviceprovisioningservices.Hashed),
            },

            "authorization_policies": {
                Type: schema.TypeList,
                Optional: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "key_name": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "rights": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(deviceprovisioningservices.ServiceConfig),
                                string(deviceprovisioningservices.EnrollmentRead),
                                string(deviceprovisioningservices.EnrollmentWrite),
                                string(deviceprovisioningservices.DeviceConnect),
                                string(deviceprovisioningservices.RegistrationStatusRead),
                                string(deviceprovisioningservices.RegistrationStatusWrite),
                            }, false),
                        },
                        "primary_key": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "secondary_key": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },

            "etag": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
            },

            "iot_hubs": {
                Type: schema.TypeList,
                Optional: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "connection_string": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "location": azure.SchemaLocation(),
                        "allocation_weight": {
                            Type: schema.TypeInt,
                            Optional: true,
                        },
                        "apply_allocation_policy": {
                            Type: schema.TypeBool,
                            Optional: true,
                        },
                    },
                },
            },

            "state": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(deviceprovisioningservices.Activating),
                    string(deviceprovisioningservices.Active),
                    string(deviceprovisioningservices.Deleting),
                    string(deviceprovisioningservices.Deleted),
                    string(deviceprovisioningservices.ActivationFailed),
                    string(deviceprovisioningservices.DeletionFailed),
                    string(deviceprovisioningservices.Transitioning),
                    string(deviceprovisioningservices.Suspending),
                    string(deviceprovisioningservices.Suspended),
                    string(deviceprovisioningservices.Resuming),
                    string(deviceprovisioningservices.FailingOver),
                    string(deviceprovisioningservices.FailoverFailed),
                }, false),
                Default: string(deviceprovisioningservices.Activating),
            },

            "device_provisioning_host_name": {
                Type: schema.TypeString,
                Computed: true,
            },

            "id_scope": {
                Type: schema.TypeString,
                Computed: true,
            },

            "provisioning_state": {
                Type: schema.TypeString,
                Computed: true,
            },

            "service_operations_host_name": {
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

func resourceArmIotDpsResourceCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).iotDpsResourceClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Iot Dps Resource %q (Resource Group %q): %+v", name, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_iot_dps_resource", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    allocationPolicy := d.Get("allocation_policy").(string)
    authorizationPolicies := d.Get("authorization_policies").([]interface{})
    etag := d.Get("etag").(string)
    iotHubs := d.Get("iot_hubs").([]interface{})
    sku := d.Get("sku").([]interface{})
    state := d.Get("state").(string)
    t := d.Get("tags").(map[string]interface{})

    iotDpsDescription := deviceprovisioningservices.ProvisioningServiceDescription{
        Etag: utils.String(etag),
        Location: utils.String(location),
        IotDpsPropertiesDescription: &deviceprovisioningservices.IotDpsPropertiesDescription{
            AllocationPolicy: deviceprovisioningservices.AllocationPolicy(allocationPolicy),
            AuthorizationPolicies: expandArmIotDpsResourceSharedAccessSignatureAuthorizationRule_AccessRightsDescription_(authorizationPolicies),
            IotHubs: expandArmIotDpsResourceIotHubDefinitionDescription(iotHubs),
            State: deviceprovisioningservices.State(state),
        },
        Sku: expandArmIotDpsResourceIotDpsSkuInfo(sku),
        Tags: tags.Expand(t),
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, name, iotDpsDescription)
    if err != nil {
        return fmt.Errorf("Error creating Iot Dps Resource %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Iot Dps Resource %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Iot Dps Resource %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Iot Dps Resource %q (Resource Group %q) ID", name, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmIotDpsResourceRead(d, meta)
}

func resourceArmIotDpsResourceRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).iotDpsResourceClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["provisioningServices"]

    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Iot Dps Resource %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Iot Dps Resource %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if iotDpsPropertiesDescription := resp.IotDpsPropertiesDescription; iotDpsPropertiesDescription != nil {
        d.Set("allocation_policy", string(iotDpsPropertiesDescription.AllocationPolicy))
        if err := d.Set("authorization_policies", flattenArmIotDpsResourceSharedAccessSignatureAuthorizationRule_AccessRightsDescription_(iotDpsPropertiesDescription.AuthorizationPolicies)); err != nil {
            return fmt.Errorf("Error setting `authorization_policies`: %+v", err)
        }
        d.Set("device_provisioning_host_name", iotDpsPropertiesDescription.DeviceProvisioningHostName)
        d.Set("id_scope", iotDpsPropertiesDescription.IDScope)
        if err := d.Set("iot_hubs", flattenArmIotDpsResourceIotHubDefinitionDescription(iotDpsPropertiesDescription.IotHubs)); err != nil {
            return fmt.Errorf("Error setting `iot_hubs`: %+v", err)
        }
        d.Set("provisioning_state", iotDpsPropertiesDescription.ProvisioningState)
        d.Set("service_operations_host_name", iotDpsPropertiesDescription.ServiceOperationsHostName)
        d.Set("state", string(iotDpsPropertiesDescription.State))
    }
    d.Set("etag", resp.Etag)
    if err := d.Set("sku", flattenArmIotDpsResourceIotDpsSkuInfo(resp.Sku)); err != nil {
        return fmt.Errorf("Error setting `sku`: %+v", err)
    }
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}


func resourceArmIotDpsResourceDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).iotDpsResourceClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["provisioningServices"]

    future, err := client.Delete(ctx, resourceGroup, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Iot Dps Resource %q (Resource Group %q): %+v", name, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Iot Dps Resource %q (Resource Group %q): %+v", name, resourceGroup, err)
        }
    }

    return nil
}

func expandArmIotDpsResourceSharedAccessSignatureAuthorizationRule_AccessRightsDescription_(input []interface{}) *[]deviceprovisioningservices.SharedAccessSignatureAuthorizationRule_AccessRightsDescription_ {
    results := make([]deviceprovisioningservices.SharedAccessSignatureAuthorizationRule_AccessRightsDescription_, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        keyName := v["key_name"].(string)
        primaryKey := v["primary_key"].(string)
        secondaryKey := v["secondary_key"].(string)
        rights := v["rights"].(string)

        result := deviceprovisioningservices.SharedAccessSignatureAuthorizationRule_AccessRightsDescription_{
            KeyName: utils.String(keyName),
            PrimaryKey: utils.String(primaryKey),
            Rights: deviceprovisioningservices.AccessRightsDescription(rights),
            SecondaryKey: utils.String(secondaryKey),
        }

        results = append(results, result)
    }
    return &results
}

func expandArmIotDpsResourceIotHubDefinitionDescription(input []interface{}) *[]deviceprovisioningservices.IotHubDefinitionDescription {
    results := make([]deviceprovisioningservices.IotHubDefinitionDescription, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        applyAllocationPolicy := v["apply_allocation_policy"].(bool)
        allocationWeight := v["allocation_weight"].(int)
        connectionString := v["connection_string"].(string)
        location := azure.NormalizeLocation(v["location"].(string))

        result := deviceprovisioningservices.IotHubDefinitionDescription{
            AllocationWeight: utils.Int32(int32(allocationWeight)),
            ApplyAllocationPolicy: utils.Bool(applyAllocationPolicy),
            ConnectionString: utils.String(connectionString),
            Location: utils.String(location),
        }

        results = append(results, result)
    }
    return &results
}

func expandArmIotDpsResourceIotDpsSkuInfo(input []interface{}) *deviceprovisioningservices.IotDpsSkuInfo {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    name := v["name"].(string)
    capacity := v["capacity"].(int)

    result := deviceprovisioningservices.IotDpsSkuInfo{
        Capacity: utils.Int64(int64(capacity)),
        Name: deviceprovisioningservices.IotDpsSku(name),
    }
    return &result
}


func flattenArmIotDpsResourceSharedAccessSignatureAuthorizationRule_AccessRightsDescription_(input *[]deviceprovisioningservices.SharedAccessSignatureAuthorizationRule_AccessRightsDescription_) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        if keyName := item.KeyName; keyName != nil {
            v["key_name"] = *keyName
        }
        if primaryKey := item.PrimaryKey; primaryKey != nil {
            v["primary_key"] = *primaryKey
        }
        v["rights"] = string(item.Rights)
        if secondaryKey := item.SecondaryKey; secondaryKey != nil {
            v["secondary_key"] = *secondaryKey
        }

        results = append(results, v)
    }

    return results
}

func flattenArmIotDpsResourceIotHubDefinitionDescription(input *[]deviceprovisioningservices.IotHubDefinitionDescription) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        if location := item.Location; location != nil {
            v["location"] = azure.NormalizeLocation(*location)
        }
        if allocationWeight := item.AllocationWeight; allocationWeight != nil {
            v["allocation_weight"] = int(*allocationWeight)
        }
        if applyAllocationPolicy := item.ApplyAllocationPolicy; applyAllocationPolicy != nil {
            v["apply_allocation_policy"] = *applyAllocationPolicy
        }
        if connectionString := item.ConnectionString; connectionString != nil {
            v["connection_string"] = *connectionString
        }

        results = append(results, v)
    }

    return results
}

func flattenArmIotDpsResourceIotDpsSkuInfo(input *deviceprovisioningservices.IotDpsSkuInfo) []interface{} {
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