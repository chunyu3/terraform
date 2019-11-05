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



func resourceArmIotHubResource() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmIotHubResourceCreateUpdate,
        Read: resourceArmIotHubResourceRead,
        Update: resourceArmIotHubResourceCreateUpdate,
        Delete: resourceArmIotHubResourceDelete,

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

            "resource_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "resourcegroup": {
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
                        "capacity": {
                            Type: schema.TypeInt,
                            Required: true,
                        },
                        "name": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(iothub.F1),
                                string(iothub.S1),
                                string(iothub.S2),
                                string(iothub.S3),
                            }, false),
                        },
                    },
                },
            },

            "subscriptionid": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
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
                                string(iothub.RegistryRead),
                                string(iothub.RegistryWrite),
                                string(iothub.ServiceConnect),
                                string(iothub.DeviceConnect),
                                string(iothub.RegistryRead, RegistryWrite),
                                string(iothub.RegistryRead, ServiceConnect),
                                string(iothub.RegistryRead, DeviceConnect),
                                string(iothub.RegistryWrite, ServiceConnect),
                                string(iothub.RegistryWrite, DeviceConnect),
                                string(iothub.ServiceConnect, DeviceConnect),
                                string(iothub.RegistryRead, RegistryWrite, ServiceConnect),
                                string(iothub.RegistryRead, RegistryWrite, DeviceConnect),
                                string(iothub.RegistryRead, ServiceConnect, DeviceConnect),
                                string(iothub.RegistryWrite, ServiceConnect, DeviceConnect),
                                string(iothub.RegistryRead, RegistryWrite, ServiceConnect, DeviceConnect),
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

            "cloud_to_device": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "default_ttl_as_iso8601": {
                            Type: schema.TypeString,
                            Optional: true,
                            ValidateFunc: validateIso8601Duration(),
                        },
                        "feedback": {
                            Type: schema.TypeList,
                            Optional: true,
                            MaxItems: 1,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "lock_duration_as_iso8601": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                        ValidateFunc: validateIso8601Duration(),
                                    },
                                    "max_delivery_count": {
                                        Type: schema.TypeInt,
                                        Optional: true,
                                    },
                                    "ttl_as_iso8601": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                        ValidateFunc: validateIso8601Duration(),
                                    },
                                },
                            },
                        },
                        "max_delivery_count": {
                            Type: schema.TypeInt,
                            Optional: true,
                        },
                    },
                },
            },

            "comments": {
                Type: schema.TypeString,
                Optional: true,
            },

            "enable_file_upload_notifications": {
                Type: schema.TypeBool,
                Optional: true,
            },

            "etag": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
            },

            "event_hub_endpoints": {
                Type: schema.TypeMap,
                Optional: true,
                Elem: &schema.Schema{Type: schema.TypeString},
            },

            "features": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(iothub.None),
                    string(iothub.DeviceManagement),
                }, false),
                Default: string(iothub.None),
            },

            "ip_filter_rules": {
                Type: schema.TypeList,
                Optional: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "action": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(iothub.Accept),
                                string(iothub.Reject),
                            }, false),
                        },
                        "filter_name": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "ip_mask": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                    },
                },
            },

            "messaging_endpoints": {
                Type: schema.TypeMap,
                Optional: true,
                Elem: &schema.Schema{Type: schema.TypeString},
            },

            "operations_monitoring_properties": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "events": {
                            Type: schema.TypeMap,
                            Optional: true,
                            Elem: &schema.Schema{Type: schema.TypeString},
                        },
                    },
                },
            },

            "storage_endpoints": {
                Type: schema.TypeMap,
                Optional: true,
                Elem: &schema.Schema{Type: schema.TypeString},
            },

            "host_name": {
                Type: schema.TypeString,
                Computed: true,
            },

            "provisioning_state": {
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

func resourceArmIotHubResourceCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).iotHubResourceClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    resourceName := d.Get("resource_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, resourceName)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Iot Hub Resource (Resource Name %q / Resource Group %q): %+v", resourceName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_iot_hub_resource", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    authorizationPolicies := d.Get("authorization_policies").([]interface{})
    cloudToDevice := d.Get("cloud_to_device").([]interface{})
    comments := d.Get("comments").(string)
    enableFileUploadNotifications := d.Get("enable_file_upload_notifications").(bool)
    etag := d.Get("etag").(string)
    eventHubEndpoints := d.Get("event_hub_endpoints").(map[string]interface{})
    features := d.Get("features").(string)
    ipFilterRules := d.Get("ip_filter_rules").([]interface{})
    messagingEndpoints := d.Get("messaging_endpoints").(map[string]interface{})
    operationsMonitoringProperties := d.Get("operations_monitoring_properties").([]interface{})
    resourcegroup := d.Get("resourcegroup").(string)
    sku := d.Get("sku").([]interface{})
    storageEndpoints := d.Get("storage_endpoints").(map[string]interface{})
    subscriptionid := d.Get("subscriptionid").(string)
    t := d.Get("tags").(map[string]interface{})

    iotHubDescription := iothub.Description{
        Etag: utils.String(etag),
        Location: utils.String(location),
        Properties: &iothub.Properties{
            AuthorizationPolicies: expandArmIotHubResourceSharedAccessSignatureAuthorizationRule(authorizationPolicies),
            CloudToDevice: expandArmIotHubResourceCloudToDeviceProperties(cloudToDevice),
            Comments: utils.String(comments),
            EnableFileUploadNotifications: utils.Bool(enableFileUploadNotifications),
            EventHubEndpoints: utils.ExpandKeyValuePairs(eventHubEndpoints),
            Features: iothub.Capabilities(features),
            IpFilterRules: expandArmIotHubResourceIpFilterRule(ipFilterRules),
            MessagingEndpoints: utils.ExpandKeyValuePairs(messagingEndpoints),
            OperationsMonitoringProperties: expandArmIotHubResourceOperationsMonitoringProperties(operationsMonitoringProperties),
            StorageEndpoints: utils.ExpandKeyValuePairs(storageEndpoints),
        },
        Resourcegroup: utils.String(resourcegroup),
        Sku: expandArmIotHubResourceSkuInfo(sku),
        Subscriptionid: utils.String(subscriptionid),
        Tags: tags.Expand(t),
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, resourceName, iotHubDescription)
    if err != nil {
        return fmt.Errorf("Error creating Iot Hub Resource (Resource Name %q / Resource Group %q): %+v", resourceName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Iot Hub Resource (Resource Name %q / Resource Group %q): %+v", resourceName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, resourceName)
    if err != nil {
        return fmt.Errorf("Error retrieving Iot Hub Resource (Resource Name %q / Resource Group %q): %+v", resourceName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Iot Hub Resource (Resource Name %q / Resource Group %q) ID", resourceName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmIotHubResourceRead(d, meta)
}

func resourceArmIotHubResourceRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).iotHubResourceClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    resourceName := id.Path["IotHubs"]

    resp, err := client.Get(ctx, resourceGroup, resourceName)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Iot Hub Resource %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Iot Hub Resource (Resource Name %q / Resource Group %q): %+v", resourceName, resourceGroup, err)
    }


    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if properties := resp.Properties; properties != nil {
        if err := d.Set("authorization_policies", flattenArmIotHubResourceSharedAccessSignatureAuthorizationRule(properties.AuthorizationPolicies)); err != nil {
            return fmt.Errorf("Error setting `authorization_policies`: %+v", err)
        }
        if err := d.Set("cloud_to_device", flattenArmIotHubResourceCloudToDeviceProperties(properties.CloudToDevice)); err != nil {
            return fmt.Errorf("Error setting `cloud_to_device`: %+v", err)
        }
        d.Set("comments", properties.Comments)
        d.Set("enable_file_upload_notifications", properties.EnableFileUploadNotifications)
        d.Set("event_hub_endpoints", utils.FlattenKeyValuePairs(properties.EventHubEndpoints))
        d.Set("features", string(properties.Features))
        d.Set("host_name", properties.HostName)
        if err := d.Set("ip_filter_rules", flattenArmIotHubResourceIpFilterRule(properties.IpFilterRules)); err != nil {
            return fmt.Errorf("Error setting `ip_filter_rules`: %+v", err)
        }
        d.Set("messaging_endpoints", utils.FlattenKeyValuePairs(properties.MessagingEndpoints))
        if err := d.Set("operations_monitoring_properties", flattenArmIotHubResourceOperationsMonitoringProperties(properties.OperationsMonitoringProperties)); err != nil {
            return fmt.Errorf("Error setting `operations_monitoring_properties`: %+v", err)
        }
        d.Set("provisioning_state", properties.ProvisioningState)
        d.Set("storage_endpoints", utils.FlattenKeyValuePairs(properties.StorageEndpoints))
    }
    d.Set("etag", resp.Etag)
    d.Set("resource_name", resourceName)
    d.Set("resourcegroup", resp.Resourcegroup)
    if err := d.Set("sku", flattenArmIotHubResourceSkuInfo(resp.Sku)); err != nil {
        return fmt.Errorf("Error setting `sku`: %+v", err)
    }
    d.Set("subscriptionid", resp.Subscriptionid)
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}


func resourceArmIotHubResourceDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).iotHubResourceClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    resourceName := id.Path["IotHubs"]

    future, err := client.Delete(ctx, resourceGroup, resourceName)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Iot Hub Resource (Resource Name %q / Resource Group %q): %+v", resourceName, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Iot Hub Resource (Resource Name %q / Resource Group %q): %+v", resourceName, resourceGroup, err)
        }
    }

    return nil
}

func expandArmIotHubResourceSharedAccessSignatureAuthorizationRule(input []interface{}) *[]iothub.SharedAccessSignatureAuthorizationRule {
    results := make([]iothub.SharedAccessSignatureAuthorizationRule, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        keyName := v["key_name"].(string)
        primaryKey := v["primary_key"].(string)
        secondaryKey := v["secondary_key"].(string)
        rights := v["rights"].(string)

        result := iothub.SharedAccessSignatureAuthorizationRule{
            KeyName: utils.String(keyName),
            PrimaryKey: utils.String(primaryKey),
            Rights: iothub.AccessRights(rights),
            SecondaryKey: utils.String(secondaryKey),
        }

        results = append(results, result)
    }
    return &results
}

func expandArmIotHubResourceCloudToDeviceProperties(input []interface{}) *iothub.CloudToDeviceProperties {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    maxDeliveryCount := v["max_delivery_count"].(int)
    defaultTtlAsIso8601 := v["default_ttl_as_iso8601"].(string)
    feedback := v["feedback"].([]interface{})

    result := iothub.CloudToDeviceProperties{
        DefaultTtlAsIso8601: utils.String(defaultTtlAsIso8601),
        Feedback: expandArmIotHubResourceFeedbackProperties(feedback),
        MaxDeliveryCount: utils.Int32(int32(maxDeliveryCount)),
    }
    return &result
}

func expandArmIotHubResourceIpFilterRule(input []interface{}) *[]iothub.IpFilterRule {
    results := make([]iothub.IpFilterRule, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        filterName := v["filter_name"].(string)
        action := v["action"].(string)
        ipMask := v["ip_mask"].(string)

        result := iothub.IpFilterRule{
            Action: iothub.IpFilterActionType(action),
            FilterName: utils.String(filterName),
            IpMask: utils.String(ipMask),
        }

        results = append(results, result)
    }
    return &results
}

func expandArmIotHubResourceOperationsMonitoringProperties(input []interface{}) *iothub.OperationsMonitoringProperties {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    events := v["events"].(map[string]interface{})

    result := iothub.OperationsMonitoringProperties{
        Events: utils.ExpandKeyValuePairs(events),
    }
    return &result
}

func expandArmIotHubResourceSkuInfo(input []interface{}) *iothub.SkuInfo {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    name := v["name"].(string)
    capacity := v["capacity"].(int)

    result := iothub.SkuInfo{
        Capacity: utils.Int64(int64(capacity)),
        Name: iothub.Sku(name),
    }
    return &result
}

func expandArmIotHubResourceFeedbackProperties(input []interface{}) *iothub.FeedbackProperties {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    lockDurationAsIso8601 := v["lock_duration_as_iso8601"].(string)
    ttlAsIso8601 := v["ttl_as_iso8601"].(string)
    maxDeliveryCount := v["max_delivery_count"].(int)

    result := iothub.FeedbackProperties{
        LockDurationAsIso8601: utils.String(lockDurationAsIso8601),
        MaxDeliveryCount: utils.Int32(int32(maxDeliveryCount)),
        TtlAsIso8601: utils.String(ttlAsIso8601),
    }
    return &result
}


func flattenArmIotHubResourceSharedAccessSignatureAuthorizationRule(input *[]iothub.SharedAccessSignatureAuthorizationRule) []interface{} {
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

func flattenArmIotHubResourceCloudToDeviceProperties(input *iothub.CloudToDeviceProperties) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if defaultTtlAsIso8601 := input.DefaultTtlAsIso8601; defaultTtlAsIso8601 != nil {
        result["default_ttl_as_iso8601"] = *defaultTtlAsIso8601
    }
    result["feedback"] = flattenArmIotHubResourceFeedbackProperties(input.Feedback)
    if maxDeliveryCount := input.MaxDeliveryCount; maxDeliveryCount != nil {
        result["max_delivery_count"] = int(*maxDeliveryCount)
    }

    return []interface{}{result}
}

func flattenArmIotHubResourceIpFilterRule(input *[]iothub.IpFilterRule) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        v["action"] = string(item.Action)
        if filterName := item.FilterName; filterName != nil {
            v["filter_name"] = *filterName
        }
        if ipMask := item.IpMask; ipMask != nil {
            v["ip_mask"] = *ipMask
        }

        results = append(results, v)
    }

    return results
}

func flattenArmIotHubResourceOperationsMonitoringProperties(input *iothub.OperationsMonitoringProperties) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    result["events"] = utils.FlattenKeyValuePairs(input.Events)

    return []interface{}{result}
}

func flattenArmIotHubResourceSkuInfo(input *iothub.SkuInfo) []interface{} {
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

func flattenArmIotHubResourceFeedbackProperties(input *iothub.FeedbackProperties) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if lockDurationAsIso8601 := input.LockDurationAsIso8601; lockDurationAsIso8601 != nil {
        result["lock_duration_as_iso8601"] = *lockDurationAsIso8601
    }
    if maxDeliveryCount := input.MaxDeliveryCount; maxDeliveryCount != nil {
        result["max_delivery_count"] = int(*maxDeliveryCount)
    }
    if ttlAsIso8601 := input.TtlAsIso8601; ttlAsIso8601 != nil {
        result["ttl_as_iso8601"] = *ttlAsIso8601
    }

    return []interface{}{result}
}