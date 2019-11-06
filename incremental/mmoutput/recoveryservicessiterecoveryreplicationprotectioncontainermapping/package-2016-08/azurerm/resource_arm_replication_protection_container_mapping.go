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



func resourceArmReplicationProtectionContainerMapping() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmReplicationProtectionContainerMappingCreateUpdate,
        Read: resourceArmReplicationProtectionContainerMappingRead,
        Update: resourceArmReplicationProtectionContainerMappingCreateUpdate,
        Delete: resourceArmReplicationProtectionContainerMappingDelete,

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

            "fabric_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "protection_container_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "resource_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "policy_id": {
                Type: schema.TypeString,
                Optional: true,
            },

            "provider_specific_input": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "instance_type": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },

            "target_protection_container_id": {
                Type: schema.TypeString,
                Optional: true,
            },

            "health": {
                Type: schema.TypeString,
                Computed: true,
            },

            "health_error_details": {
                Type: schema.TypeList,
                Computed: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "child_errors": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "child_errors": {
                                        Type: schema.TypeList,
                                        Optional: true,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "creation_time_utc": {
                                                    Type: schema.TypeString,
                                                    Optional: true,
                                                    ValidateFunc: validateRFC3339Date,
                                                },
                                                "entity_id": {
                                                    Type: schema.TypeString,
                                                    Optional: true,
                                                },
                                                "error_code": {
                                                    Type: schema.TypeString,
                                                    Optional: true,
                                                },
                                                "error_level": {
                                                    Type: schema.TypeString,
                                                    Optional: true,
                                                },
                                                "error_message": {
                                                    Type: schema.TypeString,
                                                    Optional: true,
                                                },
                                                "error_source": {
                                                    Type: schema.TypeString,
                                                    Optional: true,
                                                },
                                                "error_type": {
                                                    Type: schema.TypeString,
                                                    Optional: true,
                                                },
                                                "possible_causes": {
                                                    Type: schema.TypeString,
                                                    Optional: true,
                                                },
                                                "recommended_action": {
                                                    Type: schema.TypeString,
                                                    Optional: true,
                                                },
                                                "recovery_provider_error_message": {
                                                    Type: schema.TypeString,
                                                    Optional: true,
                                                },
                                            },
                                        },
                                    },
                                    "creation_time_utc": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                        ValidateFunc: validateRFC3339Date,
                                    },
                                    "entity_id": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                    "error_code": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                    "error_level": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                    "error_message": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                    "error_source": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                    "error_type": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                    "possible_causes": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                    "recommended_action": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                    "recovery_provider_error_message": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                },
                            },
                        },
                        "creation_time_utc": {
                            Type: schema.TypeString,
                            Optional: true,
                            ValidateFunc: validateRFC3339Date,
                        },
                        "entity_id": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "error_code": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "error_level": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "error_message": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "error_source": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "error_type": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "possible_causes": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "recommended_action": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "recovery_provider_error_message": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },

            "policy_friendly_name": {
                Type: schema.TypeString,
                Computed: true,
            },

            "policy_id": {
                Type: schema.TypeString,
                Computed: true,
            },

            "provider_specific_details": {
                Type: schema.TypeList,
                Computed: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "instance_type": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },

            "source_fabric_friendly_name": {
                Type: schema.TypeString,
                Computed: true,
            },

            "source_protection_container_friendly_name": {
                Type: schema.TypeString,
                Computed: true,
            },

            "state": {
                Type: schema.TypeString,
                Computed: true,
            },

            "target_fabric_friendly_name": {
                Type: schema.TypeString,
                Computed: true,
            },

            "target_protection_container_friendly_name": {
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

func resourceArmReplicationProtectionContainerMappingCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).replicationProtectionContainerMappingsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    fabricName := d.Get("fabric_name").(string)
    protectionContainerName := d.Get("protection_container_name").(string)
    resourceName := d.Get("resource_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, resourceName, fabricName, protectionContainerName, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Replication Protection Container Mapping %q (Protection Container Name %q / Fabric Name %q / Resource Group %q / Resource Name %q): %+v", name, protectionContainerName, fabricName, resourceGroup, resourceName, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_replication_protection_container_mapping", *existing.ID)
        }
    }

    policyId := d.Get("policy_id").(string)
    providerSpecificInput := d.Get("provider_specific_input").([]interface{})
    targetProtectionContainerId := d.Get("target_protection_container_id").(string)

    creationInput := recoveryservicessiterecovery.CreateProtectionContainerMappingInput{
        CreateProtectionContainerMappingInputProperties: &recoveryservicessiterecovery.CreateProtectionContainerMappingInputProperties{
            PolicyID: utils.String(policyId),
            ProviderSpecificInput: expandArmReplicationProtectionContainerMappingReplicationProviderSpecificContainerMappingInput(providerSpecificInput),
            TargetProtectionContainerID: utils.String(targetProtectionContainerId),
        },
    }


    future, err := client.Create(ctx, resourceGroup, resourceName, fabricName, protectionContainerName, name, creationInput)
    if err != nil {
        return fmt.Errorf("Error creating Replication Protection Container Mapping %q (Protection Container Name %q / Fabric Name %q / Resource Group %q / Resource Name %q): %+v", name, protectionContainerName, fabricName, resourceGroup, resourceName, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Replication Protection Container Mapping %q (Protection Container Name %q / Fabric Name %q / Resource Group %q / Resource Name %q): %+v", name, protectionContainerName, fabricName, resourceGroup, resourceName, err)
    }


    resp, err := client.Get(ctx, resourceGroup, resourceName, fabricName, protectionContainerName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Replication Protection Container Mapping %q (Protection Container Name %q / Fabric Name %q / Resource Group %q / Resource Name %q): %+v", name, protectionContainerName, fabricName, resourceGroup, resourceName, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Replication Protection Container Mapping %q (Protection Container Name %q / Fabric Name %q / Resource Group %q / Resource Name %q) ID", name, protectionContainerName, fabricName, resourceGroup, resourceName)
    }
    d.SetId(*resp.ID)

    return resourceArmReplicationProtectionContainerMappingRead(d, meta)
}

func resourceArmReplicationProtectionContainerMappingRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).replicationProtectionContainerMappingsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    resourceName := id.Path["vaults"]
    fabricName := id.Path["replicationFabrics"]
    protectionContainerName := id.Path["replicationProtectionContainers"]
    name := id.Path["replicationProtectionContainerMappings"]

    resp, err := client.Get(ctx, resourceGroup, resourceName, fabricName, protectionContainerName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Replication Protection Container Mapping %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Replication Protection Container Mapping %q (Protection Container Name %q / Fabric Name %q / Resource Group %q / Resource Name %q): %+v", name, protectionContainerName, fabricName, resourceGroup, resourceName, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    d.Set("fabric_name", fabricName)
    if createProtectionContainerMappingInputProperties := resp.CreateProtectionContainerMappingInputProperties; createProtectionContainerMappingInputProperties != nil {
        d.Set("health", createProtectionContainerMappingInputProperties.Health)
        if err := d.Set("health_error_details", flattenArmReplicationProtectionContainerMappingHealthError(createProtectionContainerMappingInputProperties.HealthErrorDetails)); err != nil {
            return fmt.Errorf("Error setting `health_error_details`: %+v", err)
        }
        d.Set("policy_friendly_name", createProtectionContainerMappingInputProperties.PolicyFriendlyName)
        d.Set("policy_id", createProtectionContainerMappingInputProperties.PolicyID)
        if err := d.Set("provider_specific_details", flattenArmReplicationProtectionContainerMappingProtectionContainerMappingProviderSpecificDetails(createProtectionContainerMappingInputProperties.ProviderSpecificDetails)); err != nil {
            return fmt.Errorf("Error setting `provider_specific_details`: %+v", err)
        }
        d.Set("source_fabric_friendly_name", createProtectionContainerMappingInputProperties.SourceFabricFriendlyName)
        d.Set("source_protection_container_friendly_name", createProtectionContainerMappingInputProperties.SourceProtectionContainerFriendlyName)
        d.Set("state", createProtectionContainerMappingInputProperties.State)
        d.Set("target_fabric_friendly_name", createProtectionContainerMappingInputProperties.TargetFabricFriendlyName)
        d.Set("target_protection_container_friendly_name", createProtectionContainerMappingInputProperties.TargetProtectionContainerFriendlyName)
        d.Set("target_protection_container_id", createProtectionContainerMappingInputProperties.TargetProtectionContainerID)
    }
    d.Set("protection_container_name", protectionContainerName)
    d.Set("resource_name", resourceName)
    d.Set("type", resp.Type)

    return nil
}


func resourceArmReplicationProtectionContainerMappingDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).replicationProtectionContainerMappingsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    resourceName := id.Path["vaults"]
    fabricName := id.Path["replicationFabrics"]
    protectionContainerName := id.Path["replicationProtectionContainers"]
    name := id.Path["replicationProtectionContainerMappings"]

    future, err := client.Delete(ctx, resourceGroup, resourceName, fabricName, protectionContainerName, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Replication Protection Container Mapping %q (Protection Container Name %q / Fabric Name %q / Resource Group %q / Resource Name %q): %+v", name, protectionContainerName, fabricName, resourceGroup, resourceName, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Replication Protection Container Mapping %q (Protection Container Name %q / Fabric Name %q / Resource Group %q / Resource Name %q): %+v", name, protectionContainerName, fabricName, resourceGroup, resourceName, err)
        }
    }

    return nil
}

func expandArmReplicationProtectionContainerMappingReplicationProviderSpecificContainerMappingInput(input []interface{}) *recoveryservicessiterecovery.ReplicationProviderSpecificContainerMappingInput {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    instanceType := v["instance_type"].(string)

    result := recoveryservicessiterecovery.ReplicationProviderSpecificContainerMappingInput{
        InstanceType: utils.String(instanceType),
    }
    return &result
}


func flattenArmReplicationProtectionContainerMappingHealthError(input *[]recoveryservicessiterecovery.HealthError) []interface{} {
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

func flattenArmReplicationProtectionContainerMappingProtectionContainerMappingProviderSpecificDetails(input *recoveryservicessiterecovery.ProtectionContainerMappingProviderSpecificDetails) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})


    return []interface{}{result}
}
