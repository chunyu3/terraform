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



func resourceArmRosettaNetProcessConfiguration() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmRosettaNetProcessConfigurationCreateUpdate,
        Read: resourceArmRosettaNetProcessConfigurationRead,
        Update: resourceArmRosettaNetProcessConfigurationCreateUpdate,
        Delete: resourceArmRosettaNetProcessConfigurationDelete,

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

            "activity_settings": {
                Type: schema.TypeList,
                Required: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "acknowledgment_of_receipt_settings": {
                            Type: schema.TypeList,
                            Required: true,
                            MaxItems: 1,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "is_non_repudiation_required": {
                                        Type: schema.TypeBool,
                                        Required: true,
                                    },
                                    "time_to_acknowledge_in_seconds": {
                                        Type: schema.TypeInt,
                                        Required: true,
                                    },
                                },
                            },
                        },
                        "activity_behavior": {
                            Type: schema.TypeList,
                            Required: true,
                            MaxItems: 1,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "action_type": {
                                        Type: schema.TypeString,
                                        Required: true,
                                        ValidateFunc: validation.StringInSlice([]string{
                                            string(logic.NotSpecified),
                                            string(logic.SingleAction),
                                            string(logic.DoubleAction),
                                        }, false),
                                    },
                                    "is_authorization_required": {
                                        Type: schema.TypeBool,
                                        Required: true,
                                    },
                                    "is_secured_transport_required": {
                                        Type: schema.TypeBool,
                                        Required: true,
                                    },
                                    "non_repudiation_of_origin_and_content": {
                                        Type: schema.TypeBool,
                                        Required: true,
                                    },
                                    "persistent_confidentiality_scope": {
                                        Type: schema.TypeString,
                                        Required: true,
                                        ValidateFunc: validation.StringInSlice([]string{
                                            string(logic.NotSpecified),
                                            string(logic.None),
                                            string(logic.Payload),
                                            string(logic.PayloadContainer),
                                        }, false),
                                    },
                                    "response_type": {
                                        Type: schema.TypeString,
                                        Required: true,
                                        ValidateFunc: validation.StringInSlice([]string{
                                            string(logic.NotSpecified),
                                            string(logic.Sync),
                                            string(logic.Async),
                                        }, false),
                                    },
                                    "retry_count": {
                                        Type: schema.TypeInt,
                                        Required: true,
                                    },
                                    "time_to_perform_in_seconds": {
                                        Type: schema.TypeInt,
                                        Required: true,
                                    },
                                },
                            },
                        },
                        "activity_type": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(logic.NotSpecified),
                                string(logic.InformationDistribution),
                                string(logic.BusinessTransaction),
                                string(logic.Notification),
                                string(logic.QueryResponse),
                                string(logic.RequestConfirm),
                                string(logic.RequestResponse),
                            }, false),
                        },
                    },
                },
            },

            "initiator_role_settings": {
                Type: schema.TypeList,
                Required: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "action": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "business_document": {
                            Type: schema.TypeList,
                            Required: true,
                            MaxItems: 1,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "name": {
                                        Type: schema.TypeString,
                                        Required: true,
                                        ValidateFunc: validate.NoEmptyStrings,
                                    },
                                    "version": {
                                        Type: schema.TypeString,
                                        Required: true,
                                        ValidateFunc: validate.NoEmptyStrings,
                                    },
                                    "description": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                },
                            },
                        },
                        "role": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "role_type": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(logic.NotSpecified),
                                string(logic.Functional),
                                string(logic.Organizational),
                                string(logic.Employee),
                            }, false),
                        },
                        "service": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "service_classification": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "description": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },

            "integration_account_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "process_code": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "process_name": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "process_version": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "responder_role_settings": {
                Type: schema.TypeList,
                Required: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "action": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "business_document": {
                            Type: schema.TypeList,
                            Required: true,
                            MaxItems: 1,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "name": {
                                        Type: schema.TypeString,
                                        Required: true,
                                        ValidateFunc: validate.NoEmptyStrings,
                                    },
                                    "version": {
                                        Type: schema.TypeString,
                                        Required: true,
                                        ValidateFunc: validate.NoEmptyStrings,
                                    },
                                    "description": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                },
                            },
                        },
                        "role": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "role_type": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(logic.NotSpecified),
                                string(logic.Functional),
                                string(logic.Organizational),
                                string(logic.Employee),
                            }, false),
                        },
                        "service": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "service_classification": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "description": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },

            "description": {
                Type: schema.TypeString,
                Optional: true,
            },

            "metadata": {
                Type: schema.TypeMap,
                Optional: true,
                Elem: &schema.Schema{Type: schema.TypeString},
            },

            "changed_time": {
                Type: schema.TypeString,
                Computed: true,
            },

            "created_time": {
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

func resourceArmRosettaNetProcessConfigurationCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).rosettaNetProcessConfigurationsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    integrationAccountName := d.Get("integration_account_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, integrationAccountName, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Rosetta Net Process Configuration %q (Integration Account Name %q / Resource Group %q): %+v", name, integrationAccountName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_rosetta_net_process_configuration", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    activitySettings := d.Get("activity_settings").([]interface{})
    description := d.Get("description").(string)
    initiatorRoleSettings := d.Get("initiator_role_settings").([]interface{})
    metadata := d.Get("metadata").(map[string]interface{})
    processCode := d.Get("process_code").(string)
    processName := d.Get("process_name").(string)
    processVersion := d.Get("process_version").(string)
    responderRoleSettings := d.Get("responder_role_settings").([]interface{})
    t := d.Get("tags").(map[string]interface{})

    rosettaNetProcessConfiguration := logic.IntegrationAccountRosettaNetProcessConfiguration{
        Location: utils.String(location),
        IntegrationAccountRosettaNetProcessConfigurationProperties: &logic.IntegrationAccountRosettaNetProcessConfigurationProperties{
            ActivitySettings: expandArmRosettaNetProcessConfigurationRosettaNetPipActivitySettings(activitySettings),
            Description: utils.String(description),
            InitiatorRoleSettings: expandArmRosettaNetProcessConfigurationRosettaNetPipRoleSettings(initiatorRoleSettings),
            Metadata: utils.ExpandKeyValuePairs(metadata),
            ProcessCode: utils.String(processCode),
            ProcessName: utils.String(processName),
            ProcessVersion: utils.String(processVersion),
            ResponderRoleSettings: expandArmRosettaNetProcessConfigurationRosettaNetPipRoleSettings(responderRoleSettings),
        },
        Tags: tags.Expand(t),
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroup, integrationAccountName, name, rosettaNetProcessConfiguration); err != nil {
        return fmt.Errorf("Error creating Rosetta Net Process Configuration %q (Integration Account Name %q / Resource Group %q): %+v", name, integrationAccountName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, integrationAccountName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Rosetta Net Process Configuration %q (Integration Account Name %q / Resource Group %q): %+v", name, integrationAccountName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Rosetta Net Process Configuration %q (Integration Account Name %q / Resource Group %q) ID", name, integrationAccountName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmRosettaNetProcessConfigurationRead(d, meta)
}

func resourceArmRosettaNetProcessConfigurationRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).rosettaNetProcessConfigurationsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    integrationAccountName := id.Path["integrationAccounts"]
    name := id.Path["rosettanetprocessconfigurations"]

    resp, err := client.Get(ctx, resourceGroup, integrationAccountName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Rosetta Net Process Configuration %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Rosetta Net Process Configuration %q (Integration Account Name %q / Resource Group %q): %+v", name, integrationAccountName, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if integrationAccountRosettaNetProcessConfigurationProperties := resp.IntegrationAccountRosettaNetProcessConfigurationProperties; integrationAccountRosettaNetProcessConfigurationProperties != nil {
        if err := d.Set("activity_settings", flattenArmRosettaNetProcessConfigurationRosettaNetPipActivitySettings(integrationAccountRosettaNetProcessConfigurationProperties.ActivitySettings)); err != nil {
            return fmt.Errorf("Error setting `activity_settings`: %+v", err)
        }
        d.Set("changed_time", (integrationAccountRosettaNetProcessConfigurationProperties.ChangedTime).String())
        d.Set("created_time", (integrationAccountRosettaNetProcessConfigurationProperties.CreatedTime).String())
        d.Set("description", integrationAccountRosettaNetProcessConfigurationProperties.Description)
        if err := d.Set("initiator_role_settings", flattenArmRosettaNetProcessConfigurationRosettaNetPipRoleSettings(integrationAccountRosettaNetProcessConfigurationProperties.InitiatorRoleSettings)); err != nil {
            return fmt.Errorf("Error setting `initiator_role_settings`: %+v", err)
        }
        d.Set("metadata", utils.FlattenKeyValuePairs(integrationAccountRosettaNetProcessConfigurationProperties.Metadata))
        d.Set("process_code", integrationAccountRosettaNetProcessConfigurationProperties.ProcessCode)
        d.Set("process_name", integrationAccountRosettaNetProcessConfigurationProperties.ProcessName)
        d.Set("process_version", integrationAccountRosettaNetProcessConfigurationProperties.ProcessVersion)
        if err := d.Set("responder_role_settings", flattenArmRosettaNetProcessConfigurationRosettaNetPipRoleSettings(integrationAccountRosettaNetProcessConfigurationProperties.ResponderRoleSettings)); err != nil {
            return fmt.Errorf("Error setting `responder_role_settings`: %+v", err)
        }
    }
    d.Set("integration_account_name", integrationAccountName)
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}


func resourceArmRosettaNetProcessConfigurationDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).rosettaNetProcessConfigurationsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    integrationAccountName := id.Path["integrationAccounts"]
    name := id.Path["rosettanetprocessconfigurations"]

    if _, err := client.Delete(ctx, resourceGroup, integrationAccountName, name); err != nil {
        return fmt.Errorf("Error deleting Rosetta Net Process Configuration %q (Integration Account Name %q / Resource Group %q): %+v", name, integrationAccountName, resourceGroup, err)
    }

    return nil
}

func expandArmRosettaNetProcessConfigurationRosettaNetPipActivitySettings(input []interface{}) *logic.RosettaNetPipActivitySettings {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    activityType := v["activity_type"].(string)
    activityBehavior := v["activity_behavior"].([]interface{})
    acknowledgmentOfReceiptSettings := v["acknowledgment_of_receipt_settings"].([]interface{})

    result := logic.RosettaNetPipActivitySettings{
        AcknowledgmentOfReceiptSettings: expandArmRosettaNetProcessConfigurationRosettaNetPipAcknowledgmentOfReceiptSettings(acknowledgmentOfReceiptSettings),
        ActivityBehavior: expandArmRosettaNetProcessConfigurationRosettaNetPipActivityBehavior(activityBehavior),
        ActivityType: logic.RosettaNetPipActivityType(activityType),
    }
    return &result
}

func expandArmRosettaNetProcessConfigurationRosettaNetPipRoleSettings(input []interface{}) *logic.RosettaNetPipRoleSettings {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    action := v["action"].(string)
    role := v["role"].(string)
    roleType := v["role_type"].(string)
    service := v["service"].(string)
    serviceClassification := v["service_classification"].(string)
    businessDocument := v["business_document"].([]interface{})
    description := v["description"].(string)

    result := logic.RosettaNetPipRoleSettings{
        Action: utils.String(action),
        BusinessDocument: expandArmRosettaNetProcessConfigurationRosettaNetPipBusinessDocument(businessDocument),
        Description: utils.String(description),
        Role: utils.String(role),
        RoleType: logic.RosettaNetPipRoleType(roleType),
        Service: utils.String(service),
        ServiceClassification: utils.String(serviceClassification),
    }
    return &result
}

func expandArmRosettaNetProcessConfigurationRosettaNetPipAcknowledgmentOfReceiptSettings(input []interface{}) *logic.RosettaNetPipAcknowledgmentOfReceiptSettings {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    isNonRepudiationRequired := v["is_non_repudiation_required"].(bool)
    timeToAcknowledgeInSeconds := v["time_to_acknowledge_in_seconds"].(int)

    result := logic.RosettaNetPipAcknowledgmentOfReceiptSettings{
        IsNonRepudiationRequired: utils.Bool(isNonRepudiationRequired),
        TimeToAcknowledgeInSeconds: utils.Int(timeToAcknowledgeInSeconds),
    }
    return &result
}

func expandArmRosettaNetProcessConfigurationRosettaNetPipActivityBehavior(input []interface{}) *logic.RosettaNetPipActivityBehavior {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    isAuthorizationRequired := v["is_authorization_required"].(bool)
    persistentConfidentialityScope := v["persistent_confidentiality_scope"].(string)
    isSecuredTransportRequired := v["is_secured_transport_required"].(bool)
    actionType := v["action_type"].(string)
    responseType := v["response_type"].(string)
    nonRepudiationOfOriginAndContent := v["non_repudiation_of_origin_and_content"].(bool)
    retryCount := v["retry_count"].(int)
    timeToPerformInSeconds := v["time_to_perform_in_seconds"].(int)

    result := logic.RosettaNetPipActivityBehavior{
        ActionType: logic.RosettaNetActionType(actionType),
        IsAuthorizationRequired: utils.Bool(isAuthorizationRequired),
        IsSecuredTransportRequired: utils.Bool(isSecuredTransportRequired),
        NonRepudiationOfOriginAndContent: utils.Bool(nonRepudiationOfOriginAndContent),
        PersistentConfidentialityScope: logic.RosettaNetPipConfidentialityScope(persistentConfidentialityScope),
        ResponseType: logic.RosettaNetResponseType(responseType),
        RetryCount: utils.Int(retryCount),
        TimeToPerformInSeconds: utils.Int(timeToPerformInSeconds),
    }
    return &result
}

func expandArmRosettaNetProcessConfigurationRosettaNetPipBusinessDocument(input []interface{}) *logic.RosettaNetPipBusinessDocument {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    name := v["name"].(string)
    version := v["version"].(string)
    description := v["description"].(string)

    result := logic.RosettaNetPipBusinessDocument{
        Description: utils.String(description),
        Name: utils.String(name),
        Version: utils.String(version),
    }
    return &result
}


func flattenArmRosettaNetProcessConfigurationRosettaNetPipActivitySettings(input *logic.RosettaNetPipActivitySettings) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    result["acknowledgment_of_receipt_settings"] = flattenArmRosettaNetProcessConfigurationRosettaNetPipAcknowledgmentOfReceiptSettings(input.AcknowledgmentOfReceiptSettings)
    result["activity_behavior"] = flattenArmRosettaNetProcessConfigurationRosettaNetPipActivityBehavior(input.ActivityBehavior)
    result["activity_type"] = string(input.ActivityType)

    return []interface{}{result}
}

func flattenArmRosettaNetProcessConfigurationRosettaNetPipRoleSettings(input *logic.RosettaNetPipRoleSettings) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if action := input.Action; action != nil {
        result["action"] = *action
    }
    result["business_document"] = flattenArmRosettaNetProcessConfigurationRosettaNetPipBusinessDocument(input.BusinessDocument)
    if description := input.Description; description != nil {
        result["description"] = *description
    }
    if role := input.Role; role != nil {
        result["role"] = *role
    }
    result["role_type"] = string(input.RoleType)
    if service := input.Service; service != nil {
        result["service"] = *service
    }
    if serviceClassification := input.ServiceClassification; serviceClassification != nil {
        result["service_classification"] = *serviceClassification
    }

    return []interface{}{result}
}

func flattenArmRosettaNetProcessConfigurationRosettaNetPipAcknowledgmentOfReceiptSettings(input *logic.RosettaNetPipAcknowledgmentOfReceiptSettings) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if isNonRepudiationRequired := input.IsNonRepudiationRequired; isNonRepudiationRequired != nil {
        result["is_non_repudiation_required"] = *isNonRepudiationRequired
    }
    if timeToAcknowledgeInSeconds := input.TimeToAcknowledgeInSeconds; timeToAcknowledgeInSeconds != nil {
        result["time_to_acknowledge_in_seconds"] = *timeToAcknowledgeInSeconds
    }

    return []interface{}{result}
}

func flattenArmRosettaNetProcessConfigurationRosettaNetPipActivityBehavior(input *logic.RosettaNetPipActivityBehavior) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    result["action_type"] = string(input.ActionType)
    if isAuthorizationRequired := input.IsAuthorizationRequired; isAuthorizationRequired != nil {
        result["is_authorization_required"] = *isAuthorizationRequired
    }
    if isSecuredTransportRequired := input.IsSecuredTransportRequired; isSecuredTransportRequired != nil {
        result["is_secured_transport_required"] = *isSecuredTransportRequired
    }
    if nonRepudiationOfOriginAndContent := input.NonRepudiationOfOriginAndContent; nonRepudiationOfOriginAndContent != nil {
        result["non_repudiation_of_origin_and_content"] = *nonRepudiationOfOriginAndContent
    }
    result["persistent_confidentiality_scope"] = string(input.PersistentConfidentialityScope)
    result["response_type"] = string(input.ResponseType)
    if retryCount := input.RetryCount; retryCount != nil {
        result["retry_count"] = *retryCount
    }
    if timeToPerformInSeconds := input.TimeToPerformInSeconds; timeToPerformInSeconds != nil {
        result["time_to_perform_in_seconds"] = *timeToPerformInSeconds
    }

    return []interface{}{result}
}

func flattenArmRosettaNetProcessConfigurationRosettaNetPipBusinessDocument(input *logic.RosettaNetPipBusinessDocument) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if name := input.Name; name != nil {
        result["name"] = *name
    }
    if description := input.Description; description != nil {
        result["description"] = *description
    }
    if version := input.Version; version != nil {
        result["version"] = *version
    }

    return []interface{}{result}
}