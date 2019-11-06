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



func resourceArmReplicationMigrationItem() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmReplicationMigrationItemCreate,
        Read: resourceArmReplicationMigrationItemRead,
        Update: resourceArmReplicationMigrationItemUpdate,
        Delete: resourceArmReplicationMigrationItemDelete,

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

            "policy_id": {
                Type: schema.TypeString,
                Required: true,
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

            "allowed_operations": {
                Type: schema.TypeList,
                Computed: true,
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
            },

            "current_job": {
                Type: schema.TypeList,
                Computed: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "job_id": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "job_name": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "start_time": {
                            Type: schema.TypeString,
                            Optional: true,
                            ValidateFunc: validateRFC3339Date,
                        },
                    },
                },
            },

            "health": {
                Type: schema.TypeString,
                Computed: true,
            },

            "health_errors": {
                Type: schema.TypeList,
                Computed: true,
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
                        "error_category": {
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
                        "inner_health_errors": {
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
                                    "error_category": {
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
                                    "summary_message": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                },
                            },
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
                        "summary_message": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },

            "machine_name": {
                Type: schema.TypeString,
                Computed: true,
            },

            "migration_state": {
                Type: schema.TypeString,
                Computed: true,
            },

            "migration_state_description": {
                Type: schema.TypeString,
                Computed: true,
            },

            "policy_friendly_name": {
                Type: schema.TypeString,
                Computed: true,
            },

            "recovery_services_provider_id": {
                Type: schema.TypeString,
                Computed: true,
            },

            "test_migrate_state": {
                Type: schema.TypeString,
                Computed: true,
            },

            "test_migrate_state_description": {
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

func resourceArmReplicationMigrationItemCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).replicationMigrationItemsClient
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
                return fmt.Errorf("Error checking for present of existing Replication Migration Item %q (Protection Container Name %q / Fabric Name %q / Resource Group %q / Resource Name %q): %+v", name, protectionContainerName, fabricName, resourceGroup, resourceName, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_replication_migration_item", *existing.ID)
        }
    }

    policyId := d.Get("policy_id").(string)

    input := recoveryservicessiterecovery.EnableMigrationInput{
        EnableMigrationInputProperties: &recoveryservicessiterecovery.EnableMigrationInputProperties{
            PolicyID: utils.String(policyId),
        },
    }


    future, err := client.Create(ctx, resourceGroup, resourceName, fabricName, protectionContainerName, name, input)
    if err != nil {
        return fmt.Errorf("Error creating Replication Migration Item %q (Protection Container Name %q / Fabric Name %q / Resource Group %q / Resource Name %q): %+v", name, protectionContainerName, fabricName, resourceGroup, resourceName, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Replication Migration Item %q (Protection Container Name %q / Fabric Name %q / Resource Group %q / Resource Name %q): %+v", name, protectionContainerName, fabricName, resourceGroup, resourceName, err)
    }


    resp, err := client.Get(ctx, resourceGroup, resourceName, fabricName, protectionContainerName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Replication Migration Item %q (Protection Container Name %q / Fabric Name %q / Resource Group %q / Resource Name %q): %+v", name, protectionContainerName, fabricName, resourceGroup, resourceName, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Replication Migration Item %q (Protection Container Name %q / Fabric Name %q / Resource Group %q / Resource Name %q) ID", name, protectionContainerName, fabricName, resourceGroup, resourceName)
    }
    d.SetId(*resp.ID)

    return resourceArmReplicationMigrationItemRead(d, meta)
}

func resourceArmReplicationMigrationItemRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).replicationMigrationItemsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    resourceName := id.Path["vaults"]
    fabricName := id.Path["replicationFabrics"]
    protectionContainerName := id.Path["replicationProtectionContainers"]
    name := id.Path["replicationMigrationItems"]

    resp, err := client.Get(ctx, resourceGroup, resourceName, fabricName, protectionContainerName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Replication Migration Item %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Replication Migration Item %q (Protection Container Name %q / Fabric Name %q / Resource Group %q / Resource Name %q): %+v", name, protectionContainerName, fabricName, resourceGroup, resourceName, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if enableMigrationInputProperties := resp.EnableMigrationInputProperties; enableMigrationInputProperties != nil {
        if err := d.Set("allowed_operations", flattenArmReplicationMigrationItem(string(enableMigrationInputProperties.AllowedOperations))); err != nil {
            return fmt.Errorf("Error setting `allowed_operations`: %+v", err)
        }
        if err := d.Set("current_job", flattenArmReplicationMigrationItemCurrentJobDetails(enableMigrationInputProperties.CurrentJob)); err != nil {
            return fmt.Errorf("Error setting `current_job`: %+v", err)
        }
        d.Set("health", enableMigrationInputProperties.Health)
        if err := d.Set("health_errors", flattenArmReplicationMigrationItemHealthError(enableMigrationInputProperties.HealthErrors)); err != nil {
            return fmt.Errorf("Error setting `health_errors`: %+v", err)
        }
        d.Set("machine_name", enableMigrationInputProperties.MachineName)
        d.Set("migration_state", string(enableMigrationInputProperties.MigrationState))
        d.Set("migration_state_description", enableMigrationInputProperties.MigrationStateDescription)
        d.Set("policy_friendly_name", enableMigrationInputProperties.PolicyFriendlyName)
        d.Set("policy_id", enableMigrationInputProperties.PolicyID)
        d.Set("recovery_services_provider_id", enableMigrationInputProperties.RecoveryServicesProviderID)
        d.Set("test_migrate_state", string(enableMigrationInputProperties.TestMigrateState))
        d.Set("test_migrate_state_description", enableMigrationInputProperties.TestMigrateStateDescription)
    }
    d.Set("fabric_name", fabricName)
    d.Set("protection_container_name", protectionContainerName)
    d.Set("resource_name", resourceName)
    d.Set("type", resp.Type)

    return nil
}

func resourceArmReplicationMigrationItemUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).replicationMigrationItemsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    fabricName := d.Get("fabric_name").(string)
    policyId := d.Get("policy_id").(string)
    protectionContainerName := d.Get("protection_container_name").(string)
    resourceName := d.Get("resource_name").(string)

    input := recoveryservicessiterecovery.EnableMigrationInput{
        EnableMigrationInputProperties: &recoveryservicessiterecovery.EnableMigrationInputProperties{
            PolicyID: utils.String(policyId),
        },
    }


    future, err := client.Update(ctx, resourceGroup, resourceName, fabricName, protectionContainerName, name, input)
    if err != nil {
        return fmt.Errorf("Error updating Replication Migration Item %q (Protection Container Name %q / Fabric Name %q / Resource Group %q / Resource Name %q): %+v", name, protectionContainerName, fabricName, resourceGroup, resourceName, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for update of Replication Migration Item %q (Protection Container Name %q / Fabric Name %q / Resource Group %q / Resource Name %q): %+v", name, protectionContainerName, fabricName, resourceGroup, resourceName, err)
    }

    return resourceArmReplicationMigrationItemRead(d, meta)
}

func resourceArmReplicationMigrationItemDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).replicationMigrationItemsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    resourceName := id.Path["vaults"]
    fabricName := id.Path["replicationFabrics"]
    protectionContainerName := id.Path["replicationProtectionContainers"]
    name := id.Path["replicationMigrationItems"]

    future, err := client.Delete(ctx, resourceGroup, resourceName, fabricName, protectionContainerName, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Replication Migration Item %q (Protection Container Name %q / Fabric Name %q / Resource Group %q / Resource Name %q): %+v", name, protectionContainerName, fabricName, resourceGroup, resourceName, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Replication Migration Item %q (Protection Container Name %q / Fabric Name %q / Resource Group %q / Resource Name %q): %+v", name, protectionContainerName, fabricName, resourceGroup, resourceName, err)
        }
    }

    return nil
}


func flattenArmReplicationMigrationItem(input *[]recoveryservicessiterecovery.) []interface{} {
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

func flattenArmReplicationMigrationItemCurrentJobDetails(input *recoveryservicessiterecovery.CurrentJobDetails) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})


    return []interface{}{result}
}

func flattenArmReplicationMigrationItemHealthError(input *[]recoveryservicessiterecovery.HealthError) []interface{} {
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
