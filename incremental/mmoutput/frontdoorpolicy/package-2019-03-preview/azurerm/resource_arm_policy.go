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



func resourceArmPolicy() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmPolicyCreateUpdate,
        Read: resourceArmPolicyRead,
        Update: resourceArmPolicyCreateUpdate,
        Delete: resourceArmPolicyDelete,

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

            "custom_rules": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "rules": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "action": {
                                        Type: schema.TypeString,
                                        Required: true,
                                        ValidateFunc: validation.StringInSlice([]string{
                                            string(frontdoor.Allow),
                                            string(frontdoor.Block),
                                            string(frontdoor.Log),
                                            string(frontdoor.Redirect),
                                        }, false),
                                    },
                                    "match_conditions": {
                                        Type: schema.TypeList,
                                        Required: true,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "match_value": {
                                                    Type: schema.TypeList,
                                                    Required: true,
                                                    Elem: &schema.Schema{
                                                        Type: schema.TypeString,
                                                    },
                                                },
                                                "match_variable": {
                                                    Type: schema.TypeString,
                                                    Required: true,
                                                    ValidateFunc: validation.StringInSlice([]string{
                                                        string(frontdoor.RemoteAddr),
                                                        string(frontdoor.RequestMethod),
                                                        string(frontdoor.QueryString),
                                                        string(frontdoor.PostArgs),
                                                        string(frontdoor.RequestUri),
                                                        string(frontdoor.RequestHeader),
                                                        string(frontdoor.RequestBody),
                                                        string(frontdoor.Cookies),
                                                    }, false),
                                                },
                                                "operator": {
                                                    Type: schema.TypeString,
                                                    Required: true,
                                                    ValidateFunc: validation.StringInSlice([]string{
                                                        string(frontdoor.Any),
                                                        string(frontdoor.IPMatch),
                                                        string(frontdoor.GeoMatch),
                                                        string(frontdoor.Equal),
                                                        string(frontdoor.Contains),
                                                        string(frontdoor.LessThan),
                                                        string(frontdoor.GreaterThan),
                                                        string(frontdoor.LessThanOrEqual),
                                                        string(frontdoor.GreaterThanOrEqual),
                                                        string(frontdoor.BeginsWith),
                                                        string(frontdoor.EndsWith),
                                                        string(frontdoor.RegEx),
                                                    }, false),
                                                },
                                                "negate_condition": {
                                                    Type: schema.TypeBool,
                                                    Optional: true,
                                                },
                                                "selector": {
                                                    Type: schema.TypeString,
                                                    Optional: true,
                                                },
                                                "transforms": {
                                                    Type: schema.TypeList,
                                                    Optional: true,
                                                    Elem: &schema.Schema{
                                                        Type: schema.TypeString,
                                                        ValidateFunc: validation.StringInSlice([]string{
                                                            string(frontdoor.Lowercase),
                                                            string(frontdoor.Uppercase),
                                                            string(frontdoor.Trim),
                                                            string(frontdoor.UrlDecode),
                                                            string(frontdoor.UrlEncode),
                                                            string(frontdoor.RemoveNulls),
                                                       }, false),
                                                    },
                                                },
                                            },
                                        },
                                    },
                                    "priority": {
                                        Type: schema.TypeInt,
                                        Required: true,
                                    },
                                    "rule_type": {
                                        Type: schema.TypeString,
                                        Required: true,
                                        ValidateFunc: validation.StringInSlice([]string{
                                            string(frontdoor.MatchRule),
                                            string(frontdoor.RateLimitRule),
                                        }, false),
                                    },
                                    "enabled_state": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                        ValidateFunc: validation.StringInSlice([]string{
                                            string(frontdoor.Disabled),
                                            string(frontdoor.Enabled),
                                        }, false),
                                        Default: string(frontdoor.Disabled),
                                    },
                                    "name": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                    "rate_limit_duration_in_minutes": {
                                        Type: schema.TypeInt,
                                        Optional: true,
                                    },
                                    "rate_limit_threshold": {
                                        Type: schema.TypeInt,
                                        Optional: true,
                                    },
                                },
                            },
                        },
                    },
                },
            },

            "etag": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
            },

            "managed_rules": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "managed_rule_sets": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "rule_set_type": {
                                        Type: schema.TypeString,
                                        Required: true,
                                        ValidateFunc: validate.NoEmptyStrings,
                                    },
                                    "rule_set_version": {
                                        Type: schema.TypeString,
                                        Required: true,
                                        ValidateFunc: validate.NoEmptyStrings,
                                    },
                                    "rule_group_overrides": {
                                        Type: schema.TypeList,
                                        Optional: true,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "rule_group_name": {
                                                    Type: schema.TypeString,
                                                    Required: true,
                                                    ValidateFunc: validate.NoEmptyStrings,
                                                },
                                            },
                                        },
                                    },
                                },
                            },
                        },
                    },
                },
            },

            "policy_settings": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "custom_block_response_body": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "custom_block_response_status_code": {
                            Type: schema.TypeInt,
                            Optional: true,
                        },
                        "enabled_state": {
                            Type: schema.TypeString,
                            Optional: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(frontdoor.Disabled),
                                string(frontdoor.Enabled),
                            }, false),
                            Default: string(frontdoor.Disabled),
                        },
                        "mode": {
                            Type: schema.TypeString,
                            Optional: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(frontdoor.Prevention),
                                string(frontdoor.Detection),
                            }, false),
                            Default: string(frontdoor.Prevention),
                        },
                        "redirect_url": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },

            "frontend_endpoint_links": {
                Type: schema.TypeList,
                Computed: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "id": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },

            "provisioning_state": {
                Type: schema.TypeString,
                Computed: true,
            },

            "resource_state": {
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

func resourceArmPolicyCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).policiesClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Policy %q (Resource Group %q): %+v", name, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_policy", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    customRules := d.Get("custom_rules").([]interface{})
    etag := d.Get("etag").(string)
    managedRules := d.Get("managed_rules").([]interface{})
    policySettings := d.Get("policy_settings").([]interface{})
    t := d.Get("tags").(map[string]interface{})

    parameters := frontdoor.WebApplicationFirewallPolicy{
        Etag: utils.String(etag),
        Location: utils.String(location),
        WebApplicationFirewallPolicyProperties: &frontdoor.WebApplicationFirewallPolicyProperties{
            CustomRules: expandArmPolicyCustomRuleList(customRules),
            ManagedRules: expandArmPolicyManagedRuleSetList(managedRules),
            PolicySettings: expandArmPolicyPolicySettings(policySettings),
        },
        Tags: tags.Expand(t),
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, name, parameters)
    if err != nil {
        return fmt.Errorf("Error creating Policy %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Policy %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Policy %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Policy %q (Resource Group %q) ID", name, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmPolicyRead(d, meta)
}

func resourceArmPolicyRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).policiesClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["FrontDoorWebApplicationFirewallPolicies"]

    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Policy %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Policy %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if webApplicationFirewallPolicyProperties := resp.WebApplicationFirewallPolicyProperties; webApplicationFirewallPolicyProperties != nil {
        if err := d.Set("custom_rules", flattenArmPolicyCustomRuleList(webApplicationFirewallPolicyProperties.CustomRules)); err != nil {
            return fmt.Errorf("Error setting `custom_rules`: %+v", err)
        }
        if err := d.Set("frontend_endpoint_links", flattenArmPolicyFrontendEndpointLink(webApplicationFirewallPolicyProperties.FrontendEndpointLinks)); err != nil {
            return fmt.Errorf("Error setting `frontend_endpoint_links`: %+v", err)
        }
        if err := d.Set("managed_rules", flattenArmPolicyManagedRuleSetList(webApplicationFirewallPolicyProperties.ManagedRules)); err != nil {
            return fmt.Errorf("Error setting `managed_rules`: %+v", err)
        }
        if err := d.Set("policy_settings", flattenArmPolicyPolicySettings(webApplicationFirewallPolicyProperties.PolicySettings)); err != nil {
            return fmt.Errorf("Error setting `policy_settings`: %+v", err)
        }
        d.Set("provisioning_state", webApplicationFirewallPolicyProperties.ProvisioningState)
        d.Set("resource_state", string(webApplicationFirewallPolicyProperties.ResourceState))
    }
    d.Set("etag", resp.Etag)
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}


func resourceArmPolicyDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).policiesClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["FrontDoorWebApplicationFirewallPolicies"]

    future, err := client.Delete(ctx, resourceGroup, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Policy %q (Resource Group %q): %+v", name, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Policy %q (Resource Group %q): %+v", name, resourceGroup, err)
        }
    }

    return nil
}

func expandArmPolicyCustomRuleList(input []interface{}) *frontdoor.CustomRuleList {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    rules := v["rules"].([]interface{})

    result := frontdoor.CustomRuleList{
        Rules: expandArmPolicyCustomRule(rules),
    }
    return &result
}

func expandArmPolicyManagedRuleSetList(input []interface{}) *frontdoor.ManagedRuleSetList {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    managedRuleSets := v["managed_rule_sets"].([]interface{})

    result := frontdoor.ManagedRuleSetList{
        ManagedRuleSets: expandArmPolicyManagedRuleSet(managedRuleSets),
    }
    return &result
}

func expandArmPolicyPolicySettings(input []interface{}) *frontdoor.PolicySettings {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    enabledState := v["enabled_state"].(string)
    mode := v["mode"].(string)
    redirectUrl := v["redirect_url"].(string)
    customBlockResponseStatusCode := v["custom_block_response_status_code"].(int)
    customBlockResponseBody := v["custom_block_response_body"].(string)

    result := frontdoor.PolicySettings{
        CustomBlockResponseBody: utils.String(customBlockResponseBody),
        CustomBlockResponseStatusCode: utils.Int(customBlockResponseStatusCode),
        EnabledState: frontdoor.PolicyEnabledState(enabledState),
        Mode: frontdoor.PolicyMode(mode),
        RedirectURL: utils.String(redirectUrl),
    }
    return &result
}

func expandArmPolicyCustomRule(input []interface{}) *[]frontdoor.CustomRule {
    results := make([]frontdoor.CustomRule, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        name := v["name"].(string)
        priority := v["priority"].(int)
        enabledState := v["enabled_state"].(string)
        ruleType := v["rule_type"].(string)
        rateLimitDurationInMinutes := v["rate_limit_duration_in_minutes"].(int)
        rateLimitThreshold := v["rate_limit_threshold"].(int)
        matchConditions := v["match_conditions"].([]interface{})
        action := v["action"].(string)

        result := frontdoor.CustomRule{
            Action: frontdoor.ActionType(action),
            EnabledState: frontdoor.CustomRuleEnabledState(enabledState),
            MatchConditions: expandArmPolicyMatchCondition(matchConditions),
            Name: utils.String(name),
            Priority: utils.Int(priority),
            RateLimitDurationInMinutes: utils.Int(rateLimitDurationInMinutes),
            RateLimitThreshold: utils.Int(rateLimitThreshold),
            RuleType: frontdoor.RuleType(ruleType),
        }

        results = append(results, result)
    }
    return &results
}

func expandArmPolicyManagedRuleSet(input []interface{}) *[]frontdoor.ManagedRuleSet {
    results := make([]frontdoor.ManagedRuleSet, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        ruleSetType := v["rule_set_type"].(string)
        ruleSetVersion := v["rule_set_version"].(string)
        ruleGroupOverrides := v["rule_group_overrides"].([]interface{})

        result := frontdoor.ManagedRuleSet{
            RuleGroupOverrides: expandArmPolicyManagedRuleGroupOverride(ruleGroupOverrides),
            RuleSetType: utils.String(ruleSetType),
            RuleSetVersion: utils.String(ruleSetVersion),
        }

        results = append(results, result)
    }
    return &results
}

func expandArmPolicyMatchCondition(input []interface{}) *[]frontdoor.MatchCondition {
    results := make([]frontdoor.MatchCondition, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        matchVariable := v["match_variable"].(string)
        selector := v["selector"].(string)
        operator := v["operator"].(string)
        negateCondition := v["negate_condition"].(bool)
        matchValue := v["match_value"].([]interface{})
        transforms := v["transforms"].([]interface{})

        result := frontdoor.MatchCondition{
            MatchValue: utils.ExpandStringSlice(matchValue),
            MatchVariable: frontdoor.MatchVariable(matchVariable),
            NegateCondition: utils.Bool(negateCondition),
            Operator: frontdoor.Operator(operator),
            Selector: utils.String(selector),
            Transforms: expandArmPolicy(transforms),
        }

        results = append(results, result)
    }
    return &results
}

func expandArmPolicyManagedRuleGroupOverride(input []interface{}) *[]frontdoor.ManagedRuleGroupOverride {
    results := make([]frontdoor.ManagedRuleGroupOverride, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        ruleGroupName := v["rule_group_name"].(string)

        result := frontdoor.ManagedRuleGroupOverride{
            RuleGroupName: utils.String(ruleGroupName),
        }

        results = append(results, result)
    }
    return &results
}

func expandArmPolicy(input []interface{}) *[]frontdoor. {
    results := make([]frontdoor., 0)
    for _, item := range input {
        v := item.(string)
        result := frontdoor.(v)
        results = append(results, result)
    }
    return &results
}


func flattenArmPolicyCustomRuleList(input *frontdoor.CustomRuleList) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    result["rules"] = flattenArmPolicyCustomRule(input.Rules)

    return []interface{}{result}
}

func flattenArmPolicyFrontendEndpointLink(input *[]frontdoor.FrontendEndpointLink) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})


        results = append(results, v)
    }

    return results
}

func flattenArmPolicyManagedRuleSetList(input *frontdoor.ManagedRuleSetList) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    result["managed_rule_sets"] = flattenArmPolicyManagedRuleSet(input.ManagedRuleSets)

    return []interface{}{result}
}

func flattenArmPolicyPolicySettings(input *frontdoor.PolicySettings) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if customBlockResponseBody := input.CustomBlockResponseBody; customBlockResponseBody != nil {
        result["custom_block_response_body"] = *customBlockResponseBody
    }
    if customBlockResponseStatusCode := input.CustomBlockResponseStatusCode; customBlockResponseStatusCode != nil {
        result["custom_block_response_status_code"] = *customBlockResponseStatusCode
    }
    result["enabled_state"] = string(input.EnabledState)
    result["mode"] = string(input.Mode)
    if redirectUrl := input.RedirectURL; redirectUrl != nil {
        result["redirect_url"] = *redirectUrl
    }

    return []interface{}{result}
}

func flattenArmPolicyCustomRule(input *[]frontdoor.CustomRule) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        if name := item.Name; name != nil {
            v["name"] = *name
        }
        v["action"] = string(item.Action)
        v["enabled_state"] = string(item.EnabledState)
        v["match_conditions"] = flattenArmPolicyMatchCondition(item.MatchConditions)
        if priority := item.Priority; priority != nil {
            v["priority"] = *priority
        }
        if rateLimitDurationInMinutes := item.RateLimitDurationInMinutes; rateLimitDurationInMinutes != nil {
            v["rate_limit_duration_in_minutes"] = *rateLimitDurationInMinutes
        }
        if rateLimitThreshold := item.RateLimitThreshold; rateLimitThreshold != nil {
            v["rate_limit_threshold"] = *rateLimitThreshold
        }
        v["rule_type"] = string(item.RuleType)

        results = append(results, v)
    }

    return results
}

func flattenArmPolicyManagedRuleSet(input *[]frontdoor.ManagedRuleSet) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        v["rule_group_overrides"] = flattenArmPolicyManagedRuleGroupOverride(item.RuleGroupOverrides)
        if ruleSetType := item.RuleSetType; ruleSetType != nil {
            v["rule_set_type"] = *ruleSetType
        }
        if ruleSetVersion := item.RuleSetVersion; ruleSetVersion != nil {
            v["rule_set_version"] = *ruleSetVersion
        }

        results = append(results, v)
    }

    return results
}

func flattenArmPolicyMatchCondition(input *[]frontdoor.MatchCondition) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        v["match_value"] = utils.FlattenStringSlice(item.MatchValue)
        v["match_variable"] = string(item.MatchVariable)
        if negateCondition := item.NegateCondition; negateCondition != nil {
            v["negate_condition"] = *negateCondition
        }
        v["operator"] = string(item.Operator)
        if selector := item.Selector; selector != nil {
            v["selector"] = *selector
        }
        v["transforms"] = flattenArmPolicy(string(item.Transforms))

        results = append(results, v)
    }

    return results
}

func flattenArmPolicyManagedRuleGroupOverride(input *[]frontdoor.ManagedRuleGroupOverride) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        if ruleGroupName := item.RuleGroupName; ruleGroupName != nil {
            v["rule_group_name"] = *ruleGroupName
        }

        results = append(results, v)
    }

    return results
}

func flattenArmPolicy(input *[]frontdoor.) []interface{} {
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
