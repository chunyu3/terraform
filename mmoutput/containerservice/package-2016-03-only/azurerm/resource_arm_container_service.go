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



func resourceArmContainerService() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmContainerServiceCreateUpdate,
        Read: resourceArmContainerServiceRead,
        Update: resourceArmContainerServiceCreateUpdate,
        Delete: resourceArmContainerServiceDelete,

        Importer: &schema.ResourceImporter{
            State: schema.ImportStatePassthrough,
        },


        Schema: map[string]*schema.Schema{
            "agent_pool_profiles": {
                Type: schema.TypeList,
                Required: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "dns_prefix": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "name": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "vm_size": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(containerservices.Standard_A0),
                                string(containerservices.Standard_A1),
                                string(containerservices.Standard_A2),
                                string(containerservices.Standard_A3),
                                string(containerservices.Standard_A4),
                                string(containerservices.Standard_A5),
                                string(containerservices.Standard_A6),
                                string(containerservices.Standard_A7),
                                string(containerservices.Standard_A8),
                                string(containerservices.Standard_A9),
                                string(containerservices.Standard_A10),
                                string(containerservices.Standard_A11),
                                string(containerservices.Standard_D1),
                                string(containerservices.Standard_D2),
                                string(containerservices.Standard_D3),
                                string(containerservices.Standard_D4),
                                string(containerservices.Standard_D11),
                                string(containerservices.Standard_D12),
                                string(containerservices.Standard_D13),
                                string(containerservices.Standard_D14),
                                string(containerservices.Standard_D1_v2),
                                string(containerservices.Standard_D2_v2),
                                string(containerservices.Standard_D3_v2),
                                string(containerservices.Standard_D4_v2),
                                string(containerservices.Standard_D5_v2),
                                string(containerservices.Standard_D11_v2),
                                string(containerservices.Standard_D12_v2),
                                string(containerservices.Standard_D13_v2),
                                string(containerservices.Standard_D14_v2),
                                string(containerservices.Standard_G1),
                                string(containerservices.Standard_G2),
                                string(containerservices.Standard_G3),
                                string(containerservices.Standard_G4),
                                string(containerservices.Standard_G5),
                                string(containerservices.Standard_DS1),
                                string(containerservices.Standard_DS2),
                                string(containerservices.Standard_DS3),
                                string(containerservices.Standard_DS4),
                                string(containerservices.Standard_DS11),
                                string(containerservices.Standard_DS12),
                                string(containerservices.Standard_DS13),
                                string(containerservices.Standard_DS14),
                                string(containerservices.Standard_GS1),
                                string(containerservices.Standard_GS2),
                                string(containerservices.Standard_GS3),
                                string(containerservices.Standard_GS4),
                                string(containerservices.Standard_GS5),
                            }, false),
                        },
                        "count": {
                            Type: schema.TypeInt,
                            Optional: true,
                        },
                    },
                },
            },

            "container_service_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "linux_profile": {
                Type: schema.TypeList,
                Required: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "admin_username": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "ssh": {
                            Type: schema.TypeList,
                            Required: true,
                            MaxItems: 1,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "public_keys": {
                                        Type: schema.TypeList,
                                        Required: true,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "key_data": {
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

            "location": azure.SchemaLocation(),

            "master_profile": {
                Type: schema.TypeList,
                Required: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "dns_prefix": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "count": {
                            Type: schema.TypeInt,
                            Optional: true,
                        },
                    },
                },
            },

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "diagnostics_profile": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "vm_diagnostics": {
                            Type: schema.TypeList,
                            Required: true,
                            MaxItems: 1,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "enabled": {
                                        Type: schema.TypeBool,
                                        Required: true,
                                    },
                                },
                            },
                        },
                    },
                },
            },

            "orchestrator_profile": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "orchestrator_type": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(containerservices.Swarm),
                                string(containerservices.DCOS),
                            }, false),
                        },
                    },
                },
            },

            "tags": tags.Schema(),

            "windows_profile": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "admin_password": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "admin_username": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                    },
                },
            },

            "id": {
                Type: schema.TypeString,
                Computed: true,
            },

            "name": {
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
        },
    }
}

func resourceArmContainerServiceCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).containerServicesClient
    ctx, cancel := timeouts.ForCreateUpdate(meta.(*ArmClient).StopContext, d)
    defer cancel()

    resourceGroupName := d.Get("resource_group").(string)
    name := d.Get("container_service_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroupName, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Container Service (Container Service Name %q / Resource Group %q): %+v", name, resourceGroupName, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_container_service", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    agentPoolProfiles := d.Get("agent_pool_profiles").([]interface{})
    diagnosticsProfile := d.Get("diagnostics_profile").([]interface{})
    linuxProfile := d.Get("linux_profile").([]interface{})
    masterProfile := d.Get("master_profile").([]interface{})
    orchestratorProfile := d.Get("orchestrator_profile").([]interface{})
    windowsProfile := d.Get("windows_profile").([]interface{})
    tags := d.Get("tags").(map[string]interface{})

    parameters := containerservices.ContainerService{
        Location: utils.String(location),
        ContainerServiceProperties: &containerservices.ContainerServiceProperties{
            AgentPoolProfiles: expandArmContainerServiceContainerServiceAgentPoolProfile(agentPoolProfiles),
            DiagnosticsProfile: expandArmContainerServiceContainerServiceDiagnosticsProfile(diagnosticsProfile),
            LinuxProfile: expandArmContainerServiceContainerServiceLinuxProfile(linuxProfile),
            MasterProfile: expandArmContainerServiceContainerServiceMasterProfile(masterProfile),
            OrchestratorProfile: expandArmContainerServiceContainerServiceOrchestratorProfile(orchestratorProfile),
            WindowsProfile: expandArmContainerServiceContainerServiceWindowsProfile(windowsProfile),
        },
        Tags: tags.Expand(tags),
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroupName, name, parameters)
    if err != nil {
        return fmt.Errorf("Error creating Container Service (Container Service Name %q / Resource Group %q): %+v", name, resourceGroupName, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Container Service (Container Service Name %q / Resource Group %q): %+v", name, resourceGroupName, err)
    }


    resp, err := client.Get(ctx, resourceGroupName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Container Service (Container Service Name %q / Resource Group %q): %+v", name, resourceGroupName, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Container Service (Container Service Name %q / Resource Group %q) ID", name, resourceGroupName)
    }
    d.SetId(*resp.ID)

    return resourceArmContainerServiceRead(d, meta)
}

func resourceArmContainerServiceRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).containerServicesClient
    ctx, cancel := timeouts.ForRead(meta.(*ArmClient).StopContext, d)
    defer cancel()

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroupName := id.ResourceGroup
    name := id.Path["containerServices"]

    resp, err := client.Get(ctx, resourceGroupName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Container Service %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Container Service (Container Service Name %q / Resource Group %q): %+v", name, resourceGroupName, err)
    }


    d.Set("resource_group", resourceGroupName)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if containerServiceProperties := resp.ContainerServiceProperties; containerServiceProperties != nil {
        if err := d.Set("agent_pool_profiles", flattenArmContainerServiceContainerServiceAgentPoolProfile(containerServiceProperties.AgentPoolProfiles)); err != nil {
            return fmt.Errorf("Error setting `agent_pool_profiles`: %+v", err)
        }
        if err := d.Set("diagnostics_profile", flattenArmContainerServiceContainerServiceDiagnosticsProfile(containerServiceProperties.DiagnosticsProfile)); err != nil {
            return fmt.Errorf("Error setting `diagnostics_profile`: %+v", err)
        }
        if err := d.Set("linux_profile", flattenArmContainerServiceContainerServiceLinuxProfile(containerServiceProperties.LinuxProfile)); err != nil {
            return fmt.Errorf("Error setting `linux_profile`: %+v", err)
        }
        if err := d.Set("master_profile", flattenArmContainerServiceContainerServiceMasterProfile(containerServiceProperties.MasterProfile)); err != nil {
            return fmt.Errorf("Error setting `master_profile`: %+v", err)
        }
        if err := d.Set("orchestrator_profile", flattenArmContainerServiceContainerServiceOrchestratorProfile(containerServiceProperties.OrchestratorProfile)); err != nil {
            return fmt.Errorf("Error setting `orchestrator_profile`: %+v", err)
        }
        d.Set("provisioning_state", containerServiceProperties.ProvisioningState)
        if err := d.Set("windows_profile", flattenArmContainerServiceContainerServiceWindowsProfile(containerServiceProperties.WindowsProfile)); err != nil {
            return fmt.Errorf("Error setting `windows_profile`: %+v", err)
        }
    }
    d.Set("container_service_name", name)
    d.Set("id", resp.ID)
    d.Set("name", resp.Name)
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}


func resourceArmContainerServiceDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).containerServicesClient
    ctx, cancel := timeouts.ForDelete(meta.(*ArmClient).StopContext, d)
    defer cancel()


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroupName := id.ResourceGroup
    name := id.Path["containerServices"]

    future, err := client.Delete(ctx, resourceGroupName, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Container Service (Container Service Name %q / Resource Group %q): %+v", name, resourceGroupName, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Container Service (Container Service Name %q / Resource Group %q): %+v", name, resourceGroupName, err)
        }
    }

    return nil
}

func expandArmContainerServiceContainerServiceAgentPoolProfile(input []interface{}) *[]containerservices.ContainerServiceAgentPoolProfile {
    results := make([]containerservices.ContainerServiceAgentPoolProfile, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        name := v["name"].(string)
        count := v["count"].(int)
        vMSize := v["vm_size"].(string)
        dNSPrefix := v["dns_prefix"].(string)

        result := containerservices.ContainerServiceAgentPoolProfile{
            Count: utils.Int32(int32(count)),
            DNSPrefix: utils.String(dNSPrefix),
            Name: utils.String(name),
            VMSize: containerservices.ContainerServiceVMSizeTypes(vMSize),
        }

        results = append(results, result)
    }
    return &results
}

func expandArmContainerServiceContainerServiceDiagnosticsProfile(input []interface{}) *containerservices.ContainerServiceDiagnosticsProfile {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    vMDiagnostics := v["vm_diagnostics"].([]interface{})

    result := containerservices.ContainerServiceDiagnosticsProfile{
        VMDiagnostics: expandArmContainerServiceContainerServiceVMDiagnostics(vMDiagnostics),
    }
    return &result
}

func expandArmContainerServiceContainerServiceLinuxProfile(input []interface{}) *containerservices.ContainerServiceLinuxProfile {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    adminUsername := v["admin_username"].(string)
    sSH := v["ssh"].([]interface{})

    result := containerservices.ContainerServiceLinuxProfile{
        AdminUsername: utils.String(adminUsername),
        SSH: expandArmContainerServiceContainerServiceSshConfiguration(sSH),
    }
    return &result
}

func expandArmContainerServiceContainerServiceMasterProfile(input []interface{}) *containerservices.ContainerServiceMasterProfile {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    count := v["count"].(int)
    dNSPrefix := v["dns_prefix"].(string)

    result := containerservices.ContainerServiceMasterProfile{
        Count: utils.Int32(int32(count)),
        DNSPrefix: utils.String(dNSPrefix),
    }
    return &result
}

func expandArmContainerServiceContainerServiceOrchestratorProfile(input []interface{}) *containerservices.ContainerServiceOrchestratorProfile {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    orchestratorType := v["orchestrator_type"].(string)

    result := containerservices.ContainerServiceOrchestratorProfile{
        OrchestratorType: containerservices.ContainerServiceOchestratorTypes(orchestratorType),
    }
    return &result
}

func expandArmContainerServiceContainerServiceWindowsProfile(input []interface{}) *containerservices.ContainerServiceWindowsProfile {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    adminUsername := v["admin_username"].(string)
    adminPassword := v["admin_password"].(string)

    result := containerservices.ContainerServiceWindowsProfile{
        AdminPassword: utils.String(adminPassword),
        AdminUsername: utils.String(adminUsername),
    }
    return &result
}

func expandArmContainerServiceContainerServiceVMDiagnostics(input []interface{}) *containerservices.ContainerServiceVMDiagnostics {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    enabled := v["enabled"].(bool)

    result := containerservices.ContainerServiceVMDiagnostics{
        Enabled: utils.Bool(enabled),
    }
    return &result
}

func expandArmContainerServiceContainerServiceSshConfiguration(input []interface{}) *containerservices.ContainerServiceSshConfiguration {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    publicKeys := v["public_keys"].([]interface{})

    result := containerservices.ContainerServiceSshConfiguration{
        PublicKeys: expandArmContainerServiceContainerServiceSshPublicKey(publicKeys),
    }
    return &result
}

func expandArmContainerServiceContainerServiceSshPublicKey(input []interface{}) *[]containerservices.ContainerServiceSshPublicKey {
    results := make([]containerservices.ContainerServiceSshPublicKey, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        keyData := v["key_data"].(string)

        result := containerservices.ContainerServiceSshPublicKey{
            KeyData: utils.String(keyData),
        }

        results = append(results, result)
    }
    return &results
}


func flattenArmContainerServiceContainerServiceAgentPoolProfile(input *[]containerservices.ContainerServiceAgentPoolProfile) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        if count := item.Count; count != nil {
            v["count"] = int(*count)
        }
        if dnsPrefix := item.DNSPrefix; dnsPrefix != nil {
            v["dns_prefix"] = *dnsPrefix
        }
        if name := item.Name; name != nil {
            v["name"] = *name
        }
        v["vm_size"] = string(item.VMSize)

        results = append(results, v)
    }

    return results
}

func flattenArmContainerServiceContainerServiceDiagnosticsProfile(input *containerservices.ContainerServiceDiagnosticsProfile) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    result["vm_diagnostics"] = flattenArmContainerServiceContainerServiceVMDiagnostics(input.VMDiagnostics)

    return []interface{}{result}
}

func flattenArmContainerServiceContainerServiceLinuxProfile(input *containerservices.ContainerServiceLinuxProfile) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if adminUsername := input.AdminUsername; adminUsername != nil {
        result["admin_username"] = *adminUsername
    }
    result["ssh"] = flattenArmContainerServiceContainerServiceSshConfiguration(input.SSH)

    return []interface{}{result}
}

func flattenArmContainerServiceContainerServiceMasterProfile(input *containerservices.ContainerServiceMasterProfile) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if count := input.Count; count != nil {
        result["count"] = int(*count)
    }
    if dnsPrefix := input.DNSPrefix; dnsPrefix != nil {
        result["dns_prefix"] = *dnsPrefix
    }

    return []interface{}{result}
}

func flattenArmContainerServiceContainerServiceOrchestratorProfile(input *containerservices.ContainerServiceOrchestratorProfile) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    result["orchestrator_type"] = string(input.OrchestratorType)

    return []interface{}{result}
}

func flattenArmContainerServiceContainerServiceWindowsProfile(input *containerservices.ContainerServiceWindowsProfile) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if adminPassword := input.AdminPassword; adminPassword != nil {
        result["admin_password"] = *adminPassword
    }
    if adminUsername := input.AdminUsername; adminUsername != nil {
        result["admin_username"] = *adminUsername
    }

    return []interface{}{result}
}

func flattenArmContainerServiceContainerServiceVMDiagnostics(input *containerservices.ContainerServiceVMDiagnostics) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if enabled := input.Enabled; enabled != nil {
        result["enabled"] = *enabled
    }

    return []interface{}{result}
}

func flattenArmContainerServiceContainerServiceSshConfiguration(input *containerservices.ContainerServiceSshConfiguration) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    result["public_keys"] = flattenArmContainerServiceContainerServiceSshPublicKey(input.PublicKeys)

    return []interface{}{result}
}

func flattenArmContainerServiceContainerServiceSshPublicKey(input *[]containerservices.ContainerServiceSshPublicKey) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        if keyData := item.KeyData; keyData != nil {
            v["key_data"] = *keyData
        }

        results = append(results, v)
    }

    return results
}
