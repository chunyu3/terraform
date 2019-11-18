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



func resourceArmAccount() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmAccountCreate,
        Read: resourceArmAccountRead,
        Update: resourceArmAccountUpdate,
        Delete: resourceArmAccountDelete,

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

            "compute_policies": {
                Type: schema.TypeList,
                Optional: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "name": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "max_degree_of_parallelism_per_job": {
                            Type: schema.TypeInt,
                            Optional: true,
                        },
                        "min_priority_per_job": {
                            Type: schema.TypeInt,
                            Optional: true,
                        },
                        "object_id": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "object_type": {
                            Type: schema.TypeString,
                            Optional: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(datalakeanalytics.User),
                                string(datalakeanalytics.Group),
                                string(datalakeanalytics.ServicePrincipal),
                            }, false),
                            Default: string(datalakeanalytics.User),
                        },
                    },
                },
            },

            "data_lake_store_accounts": {
                Type: schema.TypeList,
                Optional: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "name": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "suffix": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },

            "firewall_allow_azure_ips": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(datalakeanalytics.Enabled),
                    string(datalakeanalytics.Disabled),
                }, false),
                Default: string(datalakeanalytics.Enabled),
            },

            "firewall_rules": {
                Type: schema.TypeList,
                Optional: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "name": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "end_ip_address": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "start_ip_address": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },

            "firewall_state": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(datalakeanalytics.Enabled),
                    string(datalakeanalytics.Disabled),
                }, false),
                Default: string(datalakeanalytics.Enabled),
            },

            "max_degree_of_parallelism": {
                Type: schema.TypeInt,
                Optional: true,
            },

            "max_degree_of_parallelism_per_job": {
                Type: schema.TypeInt,
                Optional: true,
            },

            "max_job_count": {
                Type: schema.TypeInt,
                Optional: true,
            },

            "min_priority_per_job": {
                Type: schema.TypeInt,
                Optional: true,
            },

            "new_tier": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(datalakeanalytics.Consumption),
                    string(datalakeanalytics.Commitment_100AUHours),
                    string(datalakeanalytics.Commitment_500AUHours),
                    string(datalakeanalytics.Commitment_1000AUHours),
                    string(datalakeanalytics.Commitment_5000AUHours),
                    string(datalakeanalytics.Commitment_10000AUHours),
                    string(datalakeanalytics.Commitment_50000AUHours),
                    string(datalakeanalytics.Commitment_100000AUHours),
                    string(datalakeanalytics.Commitment_500000AUHours),
                }, false),
                Default: string(datalakeanalytics.Consumption),
            },

            "query_store_retention": {
                Type: schema.TypeInt,
                Optional: true,
            },

            "storage_accounts": {
                Type: schema.TypeList,
                Optional: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "name": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "access_key": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "suffix": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },

            "account_id": {
                Type: schema.TypeString,
                Computed: true,
            },

            "creation_time": {
                Type: schema.TypeString,
                Computed: true,
            },

            "current_tier": {
                Type: schema.TypeString,
                Computed: true,
            },

            "default_data_lake_store_account": {
                Type: schema.TypeString,
                Computed: true,
            },

            "endpoint": {
                Type: schema.TypeString,
                Computed: true,
            },

            "last_modified_time": {
                Type: schema.TypeString,
                Computed: true,
            },

            "provisioning_state": {
                Type: schema.TypeString,
                Computed: true,
            },

            "state": {
                Type: schema.TypeString,
                Computed: true,
            },

            "system_max_degree_of_parallelism": {
                Type: schema.TypeInt,
                Computed: true,
            },

            "system_max_job_count": {
                Type: schema.TypeInt,
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

func resourceArmAccountCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).accountsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Account %q (Resource Group %q): %+v", name, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_account", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    computePolicies := d.Get("compute_policies").([]interface{})
    dataLakeStoreAccounts := d.Get("data_lake_store_accounts").([]interface{})
    firewallAllowAzureIps := d.Get("firewall_allow_azure_ips").(string)
    firewallRules := d.Get("firewall_rules").([]interface{})
    firewallState := d.Get("firewall_state").(string)
    maxDegreeOfParallelism := d.Get("max_degree_of_parallelism").(int)
    maxDegreeOfParallelismPerJob := d.Get("max_degree_of_parallelism_per_job").(int)
    maxJobCount := d.Get("max_job_count").(int)
    minPriorityPerJob := d.Get("min_priority_per_job").(int)
    newTier := d.Get("new_tier").(string)
    queryStoreRetention := d.Get("query_store_retention").(int)
    storageAccounts := d.Get("storage_accounts").([]interface{})
    t := d.Get("tags").(map[string]interface{})

    parameters := datalakeanalytics.UpdateDataLakeAnalyticsAccountParameters{
        Location: utils.String(location),
        UpdateDataLakeAnalyticsAccountProperties: &datalakeanalytics.UpdateDataLakeAnalyticsAccountProperties{
            ComputePolicies: expandArmAccountUpdateComputePolicyWithAccountParameters(computePolicies),
            DataLakeStoreAccounts: expandArmAccountUpdateDataLakeStoreWithAccountParameters(dataLakeStoreAccounts),
            FirewallAllowAzureIps: datalakeanalytics.FirewallAllowAzureIpsState(firewallAllowAzureIps),
            FirewallRules: expandArmAccountUpdateFirewallRuleWithAccountParameters(firewallRules),
            FirewallState: datalakeanalytics.FirewallState(firewallState),
            MaxDegreeOfParallelism: utils.Int32(int32(maxDegreeOfParallelism)),
            MaxDegreeOfParallelismPerJob: utils.Int32(int32(maxDegreeOfParallelismPerJob)),
            MaxJobCount: utils.Int32(int32(maxJobCount)),
            MinPriorityPerJob: utils.Int32(int32(minPriorityPerJob)),
            NewTier: datalakeanalytics.TierType(newTier),
            QueryStoreRetention: utils.Int32(int32(queryStoreRetention)),
            StorageAccounts: expandArmAccountUpdateStorageAccountWithAccountParameters(storageAccounts),
        },
        Tags: tags.Expand(t),
    }


    future, err := client.Create(ctx, resourceGroup, name, parameters)
    if err != nil {
        return fmt.Errorf("Error creating Account %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Account %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Account %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Account %q (Resource Group %q) ID", name, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmAccountRead(d, meta)
}

func resourceArmAccountRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).accountsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["accounts"]

    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Account %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Account %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if updateDataLakeAnalyticsAccountProperties := resp.UpdateDataLakeAnalyticsAccountProperties; updateDataLakeAnalyticsAccountProperties != nil {
        d.Set("account_id", updateDataLakeAnalyticsAccountProperties.AccountID)
        if err := d.Set("compute_policies", flattenArmAccountUpdateComputePolicyWithAccountParameters(updateDataLakeAnalyticsAccountProperties.ComputePolicies)); err != nil {
            return fmt.Errorf("Error setting `compute_policies`: %+v", err)
        }
        d.Set("creation_time", (updateDataLakeAnalyticsAccountProperties.CreationTime).String())
        d.Set("current_tier", string(updateDataLakeAnalyticsAccountProperties.CurrentTier))
        if err := d.Set("data_lake_store_accounts", flattenArmAccountUpdateDataLakeStoreWithAccountParameters(updateDataLakeAnalyticsAccountProperties.DataLakeStoreAccounts)); err != nil {
            return fmt.Errorf("Error setting `data_lake_store_accounts`: %+v", err)
        }
        d.Set("default_data_lake_store_account", updateDataLakeAnalyticsAccountProperties.DefaultDataLakeStoreAccount)
        d.Set("endpoint", updateDataLakeAnalyticsAccountProperties.Endpoint)
        d.Set("firewall_allow_azure_ips", string(updateDataLakeAnalyticsAccountProperties.FirewallAllowAzureIps))
        if err := d.Set("firewall_rules", flattenArmAccountUpdateFirewallRuleWithAccountParameters(updateDataLakeAnalyticsAccountProperties.FirewallRules)); err != nil {
            return fmt.Errorf("Error setting `firewall_rules`: %+v", err)
        }
        d.Set("firewall_state", string(updateDataLakeAnalyticsAccountProperties.FirewallState))
        d.Set("last_modified_time", (updateDataLakeAnalyticsAccountProperties.LastModifiedTime).String())
        d.Set("max_degree_of_parallelism", int(*updateDataLakeAnalyticsAccountProperties.MaxDegreeOfParallelism))
        d.Set("max_degree_of_parallelism_per_job", int(*updateDataLakeAnalyticsAccountProperties.MaxDegreeOfParallelismPerJob))
        d.Set("max_job_count", int(*updateDataLakeAnalyticsAccountProperties.MaxJobCount))
        d.Set("min_priority_per_job", int(*updateDataLakeAnalyticsAccountProperties.MinPriorityPerJob))
        d.Set("new_tier", string(updateDataLakeAnalyticsAccountProperties.NewTier))
        d.Set("provisioning_state", string(updateDataLakeAnalyticsAccountProperties.ProvisioningState))
        d.Set("query_store_retention", int(*updateDataLakeAnalyticsAccountProperties.QueryStoreRetention))
        d.Set("state", string(updateDataLakeAnalyticsAccountProperties.State))
        if err := d.Set("storage_accounts", flattenArmAccountUpdateStorageAccountWithAccountParameters(updateDataLakeAnalyticsAccountProperties.StorageAccounts)); err != nil {
            return fmt.Errorf("Error setting `storage_accounts`: %+v", err)
        }
        d.Set("system_max_degree_of_parallelism", int(*updateDataLakeAnalyticsAccountProperties.SystemMaxDegreeOfParallelism))
        d.Set("system_max_job_count", int(*updateDataLakeAnalyticsAccountProperties.SystemMaxJobCount))
    }
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}

func resourceArmAccountUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).accountsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    computePolicies := d.Get("compute_policies").([]interface{})
    dataLakeStoreAccounts := d.Get("data_lake_store_accounts").([]interface{})
    firewallAllowAzureIps := d.Get("firewall_allow_azure_ips").(string)
    firewallRules := d.Get("firewall_rules").([]interface{})
    firewallState := d.Get("firewall_state").(string)
    maxDegreeOfParallelism := d.Get("max_degree_of_parallelism").(int)
    maxDegreeOfParallelismPerJob := d.Get("max_degree_of_parallelism_per_job").(int)
    maxJobCount := d.Get("max_job_count").(int)
    minPriorityPerJob := d.Get("min_priority_per_job").(int)
    newTier := d.Get("new_tier").(string)
    queryStoreRetention := d.Get("query_store_retention").(int)
    storageAccounts := d.Get("storage_accounts").([]interface{})
    t := d.Get("tags").(map[string]interface{})

    parameters := datalakeanalytics.UpdateDataLakeAnalyticsAccountParameters{
        Location: utils.String(location),
        UpdateDataLakeAnalyticsAccountProperties: &datalakeanalytics.UpdateDataLakeAnalyticsAccountProperties{
            ComputePolicies: expandArmAccountUpdateComputePolicyWithAccountParameters(computePolicies),
            DataLakeStoreAccounts: expandArmAccountUpdateDataLakeStoreWithAccountParameters(dataLakeStoreAccounts),
            FirewallAllowAzureIps: datalakeanalytics.FirewallAllowAzureIpsState(firewallAllowAzureIps),
            FirewallRules: expandArmAccountUpdateFirewallRuleWithAccountParameters(firewallRules),
            FirewallState: datalakeanalytics.FirewallState(firewallState),
            MaxDegreeOfParallelism: utils.Int32(int32(maxDegreeOfParallelism)),
            MaxDegreeOfParallelismPerJob: utils.Int32(int32(maxDegreeOfParallelismPerJob)),
            MaxJobCount: utils.Int32(int32(maxJobCount)),
            MinPriorityPerJob: utils.Int32(int32(minPriorityPerJob)),
            NewTier: datalakeanalytics.TierType(newTier),
            QueryStoreRetention: utils.Int32(int32(queryStoreRetention)),
            StorageAccounts: expandArmAccountUpdateStorageAccountWithAccountParameters(storageAccounts),
        },
        Tags: tags.Expand(t),
    }


    future, err := client.Update(ctx, resourceGroup, name, parameters)
    if err != nil {
        return fmt.Errorf("Error updating Account %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for update of Account %q (Resource Group %q): %+v", name, resourceGroup, err)
    }

    return resourceArmAccountRead(d, meta)
}

func resourceArmAccountDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).accountsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["accounts"]

    future, err := client.Delete(ctx, resourceGroup, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Account %q (Resource Group %q): %+v", name, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Account %q (Resource Group %q): %+v", name, resourceGroup, err)
        }
    }

    return nil
}

func expandArmAccountUpdateComputePolicyWithAccountParameters(input []interface{}) *[]datalakeanalytics.UpdateComputePolicyWithAccountParameters {
    results := make([]datalakeanalytics.UpdateComputePolicyWithAccountParameters, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        name := v["name"].(string)
        objectId := v["object_id"].(string)
        objectType := v["object_type"].(string)
        maxDegreeOfParallelismPerJob := v["max_degree_of_parallelism_per_job"].(int)
        minPriorityPerJob := v["min_priority_per_job"].(int)

        result := datalakeanalytics.UpdateComputePolicyWithAccountParameters{
            Name: utils.String(name),
            UpdateComputePolicyProperties: &datalakeanalytics.UpdateComputePolicyProperties{
                MaxDegreeOfParallelismPerJob: utils.Int32(int32(maxDegreeOfParallelismPerJob)),
                MinPriorityPerJob: utils.Int32(int32(minPriorityPerJob)),
                ObjectID: utils.String(objectId),
                ObjectType: datalakeanalytics.AADObjectType(objectType),
            },
        }

        results = append(results, result)
    }
    return &results
}

func expandArmAccountUpdateDataLakeStoreWithAccountParameters(input []interface{}) *[]datalakeanalytics.UpdateDataLakeStoreWithAccountParameters {
    results := make([]datalakeanalytics.UpdateDataLakeStoreWithAccountParameters, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        name := v["name"].(string)
        suffix := v["suffix"].(string)

        result := datalakeanalytics.UpdateDataLakeStoreWithAccountParameters{
            Name: utils.String(name),
            UpdateDataLakeStoreProperties: &datalakeanalytics.UpdateDataLakeStoreProperties{
                Suffix: utils.String(suffix),
            },
        }

        results = append(results, result)
    }
    return &results
}

func expandArmAccountUpdateFirewallRuleWithAccountParameters(input []interface{}) *[]datalakeanalytics.UpdateFirewallRuleWithAccountParameters {
    results := make([]datalakeanalytics.UpdateFirewallRuleWithAccountParameters, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        name := v["name"].(string)
        startIpAddress := v["start_ip_address"].(string)
        endIpAddress := v["end_ip_address"].(string)

        result := datalakeanalytics.UpdateFirewallRuleWithAccountParameters{
            Name: utils.String(name),
            UpdateFirewallRuleProperties: &datalakeanalytics.UpdateFirewallRuleProperties{
                EndIpAddress: utils.String(endIpAddress),
                StartIpAddress: utils.String(startIpAddress),
            },
        }

        results = append(results, result)
    }
    return &results
}

func expandArmAccountUpdateStorageAccountWithAccountParameters(input []interface{}) *[]datalakeanalytics.UpdateStorageAccountWithAccountParameters {
    results := make([]datalakeanalytics.UpdateStorageAccountWithAccountParameters, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        name := v["name"].(string)
        accessKey := v["access_key"].(string)
        suffix := v["suffix"].(string)

        result := datalakeanalytics.UpdateStorageAccountWithAccountParameters{
            Name: utils.String(name),
            UpdateStorageAccountProperties: &datalakeanalytics.UpdateStorageAccountProperties{
                AccessKey: utils.String(accessKey),
                Suffix: utils.String(suffix),
            },
        }

        results = append(results, result)
    }
    return &results
}


func flattenArmAccountUpdateComputePolicyWithAccountParameters(input *[]datalakeanalytics.UpdateComputePolicyWithAccountParameters) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        if name := item.Name; name != nil {
            v["name"] = *name
        }
        if updateComputePolicyProperties := item.UpdateComputePolicyProperties; updateComputePolicyProperties != nil {
            if maxDegreeOfParallelismPerJob := updateComputePolicyProperties.MaxDegreeOfParallelismPerJob; maxDegreeOfParallelismPerJob != nil {
                v["max_degree_of_parallelism_per_job"] = int(*maxDegreeOfParallelismPerJob)
            }
            if minPriorityPerJob := updateComputePolicyProperties.MinPriorityPerJob; minPriorityPerJob != nil {
                v["min_priority_per_job"] = int(*minPriorityPerJob)
            }
            if objectId := updateComputePolicyProperties.ObjectID; objectId != nil {
                v["object_id"] = *objectId
            }
            v["object_type"] = string(updateComputePolicyProperties.ObjectType)
        }

        results = append(results, v)
    }

    return results
}

func flattenArmAccountUpdateDataLakeStoreWithAccountParameters(input *[]datalakeanalytics.UpdateDataLakeStoreWithAccountParameters) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        if name := item.Name; name != nil {
            v["name"] = *name
        }
        if updateDataLakeStoreProperties := item.UpdateDataLakeStoreProperties; updateDataLakeStoreProperties != nil {
            if suffix := updateDataLakeStoreProperties.Suffix; suffix != nil {
                v["suffix"] = *suffix
            }
        }

        results = append(results, v)
    }

    return results
}

func flattenArmAccountUpdateFirewallRuleWithAccountParameters(input *[]datalakeanalytics.UpdateFirewallRuleWithAccountParameters) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        if name := item.Name; name != nil {
            v["name"] = *name
        }
        if updateFirewallRuleProperties := item.UpdateFirewallRuleProperties; updateFirewallRuleProperties != nil {
            if endIpAddress := updateFirewallRuleProperties.EndIpAddress; endIpAddress != nil {
                v["end_ip_address"] = *endIpAddress
            }
            if startIpAddress := updateFirewallRuleProperties.StartIpAddress; startIpAddress != nil {
                v["start_ip_address"] = *startIpAddress
            }
        }

        results = append(results, v)
    }

    return results
}

func flattenArmAccountUpdateStorageAccountWithAccountParameters(input *[]datalakeanalytics.UpdateStorageAccountWithAccountParameters) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        if name := item.Name; name != nil {
            v["name"] = *name
        }
        if updateStorageAccountProperties := item.UpdateStorageAccountProperties; updateStorageAccountProperties != nil {
            if suffix := updateStorageAccountProperties.Suffix; suffix != nil {
                v["suffix"] = *suffix
            }
        }

        results = append(results, v)
    }

    return results
}
