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



func resourceArmReplicationFabric() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmReplicationFabricCreateUpdate,
        Read: resourceArmReplicationFabricRead,
        Update: resourceArmReplicationFabricCreateUpdate,
        Delete: resourceArmReplicationFabricDelete,

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

            "fabric_name": {
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

            "bcdr_state": {
                Type: schema.TypeString,
                Computed: true,
            },

            "encryption_details": {
                Type: schema.TypeList,
                Computed: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "kek_cert_expiry_date": {
                            Type: schema.TypeString,
                            Optional: true,
                            ValidateFunc: validateRFC3339Date,
                        },
                        "kek_cert_thumbprint": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "kek_state": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },

            "friendly_name": {
                Type: schema.TypeString,
                Computed: true,
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

            "internal_identifier": {
                Type: schema.TypeString,
                Computed: true,
            },

            "rollover_encryption_details": {
                Type: schema.TypeList,
                Computed: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "kek_cert_expiry_date": {
                            Type: schema.TypeString,
                            Optional: true,
                            ValidateFunc: validateRFC3339Date,
                        },
                        "kek_cert_thumbprint": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "kek_state": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },

            "type": {
                Type: schema.TypeString,
                Computed: true,
            },
        },
    }
}

func resourceArmReplicationFabricCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).replicationFabricsClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    fabricName := d.Get("fabric_name").(string)
    resourceName := d.Get("resource_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceName, resourceGroup, fabricName)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Replication Fabric (Fabric Name %q / Resource Group %q / Resource Name %q): %+v", fabricName, resourceGroup, resourceName, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_replication_fabric", *existing.ID)
        }
    }


    input := recoveryservicessiterecovery.FabricCreationInput{
    }


    future, err := client.Create(ctx, resourceName, resourceGroup, fabricName, input)
    if err != nil {
        return fmt.Errorf("Error creating Replication Fabric (Fabric Name %q / Resource Group %q / Resource Name %q): %+v", fabricName, resourceGroup, resourceName, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Replication Fabric (Fabric Name %q / Resource Group %q / Resource Name %q): %+v", fabricName, resourceGroup, resourceName, err)
    }


    resp, err := client.Get(ctx, resourceName, resourceGroup, fabricName)
    if err != nil {
        return fmt.Errorf("Error retrieving Replication Fabric (Fabric Name %q / Resource Group %q / Resource Name %q): %+v", fabricName, resourceGroup, resourceName, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Replication Fabric (Fabric Name %q / Resource Group %q / Resource Name %q) ID", fabricName, resourceGroup, resourceName)
    }
    d.SetId(*resp.ID)

    return resourceArmReplicationFabricRead(d, meta)
}

func resourceArmReplicationFabricRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).replicationFabricsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceName := id.Path["vaults"]
    resourceGroup := id.ResourceGroup
    fabricName := id.Path["replicationFabrics"]

    resp, err := client.Get(ctx, resourceName, resourceGroup, fabricName)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Replication Fabric %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Replication Fabric (Fabric Name %q / Resource Group %q / Resource Name %q): %+v", fabricName, resourceGroup, resourceName, err)
    }


    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if fabricCreationInputProperties := resp.FabricCreationInputProperties; fabricCreationInputProperties != nil {
        d.Set("bcdr_state", fabricCreationInputProperties.BcdrState)
        if err := d.Set("encryption_details", flattenArmReplicationFabricEncryptionDetails(fabricCreationInputProperties.EncryptionDetails)); err != nil {
            return fmt.Errorf("Error setting `encryption_details`: %+v", err)
        }
        d.Set("friendly_name", fabricCreationInputProperties.FriendlyName)
        d.Set("health", fabricCreationInputProperties.Health)
        if err := d.Set("health_error_details", flattenArmReplicationFabricHealthError(fabricCreationInputProperties.HealthErrorDetails)); err != nil {
            return fmt.Errorf("Error setting `health_error_details`: %+v", err)
        }
        d.Set("internal_identifier", fabricCreationInputProperties.InternalIdentifier)
        if err := d.Set("rollover_encryption_details", flattenArmReplicationFabricEncryptionDetails(fabricCreationInputProperties.RolloverEncryptionDetails)); err != nil {
            return fmt.Errorf("Error setting `rollover_encryption_details`: %+v", err)
        }
    }
    d.Set("fabric_name", fabricName)
    d.Set("resource_name", resourceName)
    d.Set("type", resp.Type)

    return nil
}


func resourceArmReplicationFabricDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).replicationFabricsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceName := id.Path["vaults"]
    resourceGroup := id.ResourceGroup
    fabricName := id.Path["replicationFabrics"]

    future, err := client.Delete(ctx, resourceName, resourceGroup, fabricName)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Replication Fabric (Fabric Name %q / Resource Group %q / Resource Name %q): %+v", fabricName, resourceGroup, resourceName, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Replication Fabric (Fabric Name %q / Resource Group %q / Resource Name %q): %+v", fabricName, resourceGroup, resourceName, err)
        }
    }

    return nil
}


func flattenArmReplicationFabricEncryptionDetails(input *recoveryservicessiterecovery.EncryptionDetails) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})


    return []interface{}{result}
}

func flattenArmReplicationFabricHealthError(input *[]recoveryservicessiterecovery.HealthError) []interface{} {
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
