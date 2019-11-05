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



func resourceArmRegistrationDefinition() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmRegistrationDefinitionCreateUpdate,
        Read: resourceArmRegistrationDefinitionRead,
        Update: resourceArmRegistrationDefinitionCreateUpdate,
        Delete: resourceArmRegistrationDefinitionDelete,

        Importer: &schema.ResourceImporter{
            State: schema.ImportStatePassthrough,
        },


        Schema: map[string]*schema.Schema{
            "name": {
                Type: schema.TypeString,
                Computed: true,
            },

            "authorizations": {
                Type: schema.TypeList,
                Required: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "principal_id": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "role_definition_id": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                    },
                },
            },

            "managed_by_tenant_id": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "registration_definition_id": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "scope": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "description": {
                Type: schema.TypeString,
                Optional: true,
            },

            "plan": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "name": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "product": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "publisher": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "version": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                    },
                },
            },

            "registration_definition_name": {
                Type: schema.TypeString,
                Optional: true,
            },

            "managed_by_tenant_name": {
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

func resourceArmRegistrationDefinitionCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).registrationDefinitionsClient
    ctx := meta.(*ArmClient).StopContext

    registrationDefinitionID := d.Get("registration_definition_id").(string)
    scope := d.Get("scope").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, scope, registrationDefinitionID)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Registration Definition (Scope %q / Registration Definition %q): %+v", scope, registrationDefinitionID, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_registration_definition", *existing.ID)
        }
    }

    authorizations := d.Get("authorizations").([]interface{})
    description := d.Get("description").(string)
    managedByTenantId := d.Get("managed_by_tenant_id").(string)
    plan := d.Get("plan").([]interface{})
    registrationDefinitionName := d.Get("registration_definition_name").(string)

    requestBody := managedservices.RegistrationDefinition{
        Plan: expandArmRegistrationDefinitionPlan(plan),
        RegistrationDefinitionProperties: &managedservices.RegistrationDefinitionProperties{
            Authorizations: expandArmRegistrationDefinitionAuthorization(authorizations),
            Description: utils.String(description),
            ManagedByTenantID: utils.String(managedByTenantId),
            RegistrationDefinitionName: utils.String(registrationDefinitionName),
        },
    }


    future, err := client.CreateOrUpdate(ctx, registrationDefinitionID, scope, requestBody)
    if err != nil {
        return fmt.Errorf("Error creating Registration Definition (Scope %q / Registration Definition %q): %+v", scope, registrationDefinitionID, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Registration Definition (Scope %q / Registration Definition %q): %+v", scope, registrationDefinitionID, err)
    }


    resp, err := client.Get(ctx, scope, registrationDefinitionID)
    if err != nil {
        return fmt.Errorf("Error retrieving Registration Definition (Scope %q / Registration Definition %q): %+v", scope, registrationDefinitionID, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Registration Definition (Scope %q / Registration Definition %q) ID", scope, registrationDefinitionID)
    }
    d.SetId(*resp.ID)

    return resourceArmRegistrationDefinitionRead(d, meta)
}

func resourceArmRegistrationDefinitionRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).registrationDefinitionsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    registrationDefinitionID := id.Path["registrationDefinitions"]

    resp, err := client.Get(ctx, scope, registrationDefinitionID)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Registration Definition %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Registration Definition (Scope %q / Registration Definition %q): %+v", scope, registrationDefinitionID, err)
    }


    d.Set("name", resp.Name)
    if registrationDefinitionProperties := resp.RegistrationDefinitionProperties; registrationDefinitionProperties != nil {
        if err := d.Set("authorizations", flattenArmRegistrationDefinitionAuthorization(registrationDefinitionProperties.Authorizations)); err != nil {
            return fmt.Errorf("Error setting `authorizations`: %+v", err)
        }
        d.Set("description", registrationDefinitionProperties.Description)
        d.Set("managed_by_tenant_id", registrationDefinitionProperties.ManagedByTenantID)
        d.Set("managed_by_tenant_name", registrationDefinitionProperties.ManagedByTenantName)
        d.Set("provisioning_state", string(registrationDefinitionProperties.ProvisioningState))
        d.Set("registration_definition_name", registrationDefinitionProperties.RegistrationDefinitionName)
    }
    if err := d.Set("plan", flattenArmRegistrationDefinitionPlan(resp.Plan)); err != nil {
        return fmt.Errorf("Error setting `plan`: %+v", err)
    }
    d.Set("registration_definition_id", registrationDefinitionID)
    d.Set("scope", scope)
    d.Set("type", resp.Type)

    return nil
}


func resourceArmRegistrationDefinitionDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).registrationDefinitionsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    registrationDefinitionID := id.Path["registrationDefinitions"]

    if _, err := client.Delete(ctx, registrationDefinitionID, scope); err != nil {
        return fmt.Errorf("Error deleting Registration Definition (Scope %q / Registration Definition %q): %+v", scope, registrationDefinitionID, err)
    }

    return nil
}

func expandArmRegistrationDefinitionPlan(input []interface{}) *managedservices.Plan {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    name := v["name"].(string)
    publisher := v["publisher"].(string)
    product := v["product"].(string)
    version := v["version"].(string)

    result := managedservices.Plan{
        Name: utils.String(name),
        Product: utils.String(product),
        Publisher: utils.String(publisher),
        Version: utils.String(version),
    }
    return &result
}

func expandArmRegistrationDefinitionAuthorization(input []interface{}) *[]managedservices.Authorization {
    results := make([]managedservices.Authorization, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        principalId := v["principal_id"].(string)
        roleDefinitionId := v["role_definition_id"].(string)

        result := managedservices.Authorization{
            PrincipalID: utils.String(principalId),
            RoleDefinitionID: utils.String(roleDefinitionId),
        }

        results = append(results, result)
    }
    return &results
}


func flattenArmRegistrationDefinitionAuthorization(input *[]managedservices.Authorization) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        if principalId := item.PrincipalID; principalId != nil {
            v["principal_id"] = *principalId
        }
        if roleDefinitionId := item.RoleDefinitionID; roleDefinitionId != nil {
            v["role_definition_id"] = *roleDefinitionId
        }

        results = append(results, v)
    }

    return results
}

func flattenArmRegistrationDefinitionPlan(input *managedservices.Plan) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if name := input.Name; name != nil {
        result["name"] = *name
    }
    if product := input.Product; product != nil {
        result["product"] = *product
    }
    if publisher := input.Publisher; publisher != nil {
        result["publisher"] = *publisher
    }
    if version := input.Version; version != nil {
        result["version"] = *version
    }

    return []interface{}{result}
}