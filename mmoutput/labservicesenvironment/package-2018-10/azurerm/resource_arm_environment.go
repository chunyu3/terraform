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



func resourceArmEnvironment() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmEnvironmentCreate,
        Read: resourceArmEnvironmentRead,
        Update: resourceArmEnvironmentUpdate,
        Delete: resourceArmEnvironmentDelete,

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

            "environment_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "environment_setting_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "lab_account_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "lab_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "resource_sets": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "resource_setting_id": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "vm_resource_id": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },

            "unique_identifier": {
                Type: schema.TypeString,
                Optional: true,
            },

            "claimed_by_user_name": {
                Type: schema.TypeString,
                Computed: true,
            },

            "claimed_by_user_object_id": {
                Type: schema.TypeString,
                Computed: true,
            },

            "claimed_by_user_principal_id": {
                Type: schema.TypeString,
                Computed: true,
            },

            "is_claimed": {
                Type: schema.TypeBool,
                Computed: true,
            },

            "last_known_power_state": {
                Type: schema.TypeString,
                Computed: true,
            },

            "latest_operation_result": {
                Type: schema.TypeList,
                Computed: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "error_code": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "error_message": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "http_method": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "operation_url": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "request_uri": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "status": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },

            "network_interface": {
                Type: schema.TypeList,
                Computed: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "private_ip_address": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "rdp_authority": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "ssh_authority": {
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

            "password_last_reset": {
                Type: schema.TypeString,
                Computed: true,
            },

            "provisioning_state": {
                Type: schema.TypeString,
                Computed: true,
            },

            "total_usage": {
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

func resourceArmEnvironmentCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).environmentsClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    environmentName := d.Get("environment_name").(string)
    environmentSettingName := d.Get("environment_setting_name").(string)
    labAccountName := d.Get("lab_account_name").(string)
    labName := d.Get("lab_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, labAccountName, labName, environmentSettingName, environmentName)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Environment (Environment Name %q / Environment Setting Name %q / Lab Name %q / Lab Account Name %q / Resource Group %q): %+v", environmentName, environmentSettingName, labName, labAccountName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_environment", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    resourceSets := d.Get("resource_sets").([]interface{})
    uniqueIdentifier := d.Get("unique_identifier").(string)
    t := d.Get("tags").(map[string]interface{})

    environment := labservices.Environment{
        Location: utils.String(location),
        EnvironmentProperties: &labservices.EnvironmentProperties{
            ResourceSets: expandArmEnvironmentResourceSet(resourceSets),
            UniqueIdentifier: utils.String(uniqueIdentifier),
        },
        Tags: tags.Expand(t),
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroup, labAccountName, labName, environmentSettingName, environmentName, environment); err != nil {
        return fmt.Errorf("Error creating Environment (Environment Name %q / Environment Setting Name %q / Lab Name %q / Lab Account Name %q / Resource Group %q): %+v", environmentName, environmentSettingName, labName, labAccountName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, labAccountName, labName, environmentSettingName, environmentName)
    if err != nil {
        return fmt.Errorf("Error retrieving Environment (Environment Name %q / Environment Setting Name %q / Lab Name %q / Lab Account Name %q / Resource Group %q): %+v", environmentName, environmentSettingName, labName, labAccountName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Environment (Environment Name %q / Environment Setting Name %q / Lab Name %q / Lab Account Name %q / Resource Group %q) ID", environmentName, environmentSettingName, labName, labAccountName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmEnvironmentRead(d, meta)
}

func resourceArmEnvironmentRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).environmentsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    labAccountName := id.Path["labaccounts"]
    labName := id.Path["labs"]
    environmentSettingName := id.Path["environmentsettings"]
    environmentName := id.Path["environments"]

    resp, err := client.Get(ctx, resourceGroup, labAccountName, labName, environmentSettingName, environmentName)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Environment %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Environment (Environment Name %q / Environment Setting Name %q / Lab Name %q / Lab Account Name %q / Resource Group %q): %+v", environmentName, environmentSettingName, labName, labAccountName, resourceGroup, err)
    }


    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if environmentProperties := resp.EnvironmentProperties; environmentProperties != nil {
        d.Set("claimed_by_user_name", environmentProperties.ClaimedByUserName)
        d.Set("claimed_by_user_object_id", environmentProperties.ClaimedByUserObjectID)
        d.Set("claimed_by_user_principal_id", environmentProperties.ClaimedByUserPrincipalID)
        d.Set("is_claimed", environmentProperties.IsClaimed)
        d.Set("last_known_power_state", environmentProperties.LastKnownPowerState)
        if err := d.Set("latest_operation_result", flattenArmEnvironmentLatestOperationResult(environmentProperties.LatestOperationResult)); err != nil {
            return fmt.Errorf("Error setting `latest_operation_result`: %+v", err)
        }
        if err := d.Set("network_interface", flattenArmEnvironmentNetworkInterface(environmentProperties.NetworkInterface)); err != nil {
            return fmt.Errorf("Error setting `network_interface`: %+v", err)
        }
        d.Set("password_last_reset", (environmentProperties.PasswordLastReset).String())
        d.Set("provisioning_state", environmentProperties.ProvisioningState)
        if err := d.Set("resource_sets", flattenArmEnvironmentResourceSet(environmentProperties.ResourceSets)); err != nil {
            return fmt.Errorf("Error setting `resource_sets`: %+v", err)
        }
        d.Set("total_usage", environmentProperties.TotalUsage)
        d.Set("unique_identifier", environmentProperties.UniqueIdentifier)
    }
    d.Set("environment_name", environmentName)
    d.Set("environment_setting_name", environmentSettingName)
    d.Set("lab_account_name", labAccountName)
    d.Set("lab_name", labName)
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}

func resourceArmEnvironmentUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).environmentsClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    environmentName := d.Get("environment_name").(string)
    environmentSettingName := d.Get("environment_setting_name").(string)
    labAccountName := d.Get("lab_account_name").(string)
    labName := d.Get("lab_name").(string)
    resourceSets := d.Get("resource_sets").([]interface{})
    uniqueIdentifier := d.Get("unique_identifier").(string)
    t := d.Get("tags").(map[string]interface{})

    environment := labservices.Environment{
        Location: utils.String(location),
        EnvironmentProperties: &labservices.EnvironmentProperties{
            ResourceSets: expandArmEnvironmentResourceSet(resourceSets),
            UniqueIdentifier: utils.String(uniqueIdentifier),
        },
        Tags: tags.Expand(t),
    }


    if _, err := client.Update(ctx, resourceGroup, labAccountName, labName, environmentSettingName, environmentName, environment); err != nil {
        return fmt.Errorf("Error updating Environment (Environment Name %q / Environment Setting Name %q / Lab Name %q / Lab Account Name %q / Resource Group %q): %+v", environmentName, environmentSettingName, labName, labAccountName, resourceGroup, err)
    }

    return resourceArmEnvironmentRead(d, meta)
}

func resourceArmEnvironmentDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).environmentsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    labAccountName := id.Path["labaccounts"]
    labName := id.Path["labs"]
    environmentSettingName := id.Path["environmentsettings"]
    environmentName := id.Path["environments"]

    future, err := client.Delete(ctx, resourceGroup, labAccountName, labName, environmentSettingName, environmentName)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Environment (Environment Name %q / Environment Setting Name %q / Lab Name %q / Lab Account Name %q / Resource Group %q): %+v", environmentName, environmentSettingName, labName, labAccountName, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Environment (Environment Name %q / Environment Setting Name %q / Lab Name %q / Lab Account Name %q / Resource Group %q): %+v", environmentName, environmentSettingName, labName, labAccountName, resourceGroup, err)
        }
    }

    return nil
}

func expandArmEnvironmentResourceSet(input []interface{}) *labservices.ResourceSet {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    vmResourceId := v["vm_resource_id"].(string)
    resourceSettingId := v["resource_setting_id"].(string)

    result := labservices.ResourceSet{
        ResourceSettingID: utils.String(resourceSettingId),
        VmResourceID: utils.String(vmResourceId),
    }
    return &result
}


func flattenArmEnvironmentLatestOperationResult(input *labservices.LatestOperationResult) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})


    return []interface{}{result}
}

func flattenArmEnvironmentNetworkInterface(input *labservices.NetworkInterface) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})


    return []interface{}{result}
}

func flattenArmEnvironmentResourceSet(input *labservices.ResourceSet) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if resourceSettingId := input.ResourceSettingID; resourceSettingId != nil {
        result["resource_setting_id"] = *resourceSettingId
    }
    if vmResourceId := input.VmResourceID; vmResourceId != nil {
        result["vm_resource_id"] = *vmResourceId
    }

    return []interface{}{result}
}
