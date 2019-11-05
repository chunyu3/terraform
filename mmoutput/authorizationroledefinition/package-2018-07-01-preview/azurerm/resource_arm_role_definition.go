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



func resourceArmRoleDefinition() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmRoleDefinitionCreateUpdate,
        Read: resourceArmRoleDefinitionRead,
        Update: resourceArmRoleDefinitionCreateUpdate,
        Delete: resourceArmRoleDefinitionDelete,

        Importer: &schema.ResourceImporter{
            State: schema.ImportStatePassthrough,
        },


        Schema: map[string]*schema.Schema{
            "name": {
                Type: schema.TypeString,
                Computed: true,
            },

            "role_definition_id": {
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

            "assignable_scopes": {
                Type: schema.TypeList,
                Optional: true,
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
            },

            "description": {
                Type: schema.TypeString,
                Optional: true,
            },

            "permissions": {
                Type: schema.TypeList,
                Optional: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "actions": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                        "data_actions": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                        "not_actions": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                        "not_data_actions": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                    },
                },
            },

            "role_name": {
                Type: schema.TypeString,
                Optional: true,
            },

            "type": {
                Type: schema.TypeString,
                Optional: true,
            },

            "type": {
                Type: schema.TypeString,
                Computed: true,
            },
        },
    }
}

func resourceArmRoleDefinitionCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).roleDefinitionsClient
    ctx := meta.(*ArmClient).StopContext

    roleDefinitionID := d.Get("role_definition_id").(string)
    scope := d.Get("scope").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, scope, roleDefinitionID)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Role Definition (Role Definition %q / Scope %q): %+v", roleDefinitionID, scope, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_role_definition", *existing.ID)
        }
    }

    assignableScopes := d.Get("assignable_scopes").([]interface{})
    description := d.Get("description").(string)
    permissions := d.Get("permissions").([]interface{})
    roleName := d.Get("role_name").(string)
    type := d.Get("type").(string)

    roleDefinition := authorization.RoleDefinition{
        RoleDefinitionProperties: &authorization.RoleDefinitionProperties{
            AssignableScopes: utils.ExpandStringSlice(assignableScopes),
            Description: utils.String(description),
            Permissions: expandArmRoleDefinitionPermission(permissions),
            RoleName: utils.String(roleName),
            Type: utils.String(type),
        },
    }


    if _, err := client.CreateOrUpdate(ctx, scope, roleDefinitionID, roleDefinition); err != nil {
        return fmt.Errorf("Error creating Role Definition (Role Definition %q / Scope %q): %+v", roleDefinitionID, scope, err)
    }


    resp, err := client.Get(ctx, scope, roleDefinitionID)
    if err != nil {
        return fmt.Errorf("Error retrieving Role Definition (Role Definition %q / Scope %q): %+v", roleDefinitionID, scope, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Role Definition (Role Definition %q / Scope %q) ID", roleDefinitionID, scope)
    }
    d.SetId(*resp.ID)

    return resourceArmRoleDefinitionRead(d, meta)
}

func resourceArmRoleDefinitionRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).roleDefinitionsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    roleDefinitionID := id.Path["roleDefinitions"]

    resp, err := client.Get(ctx, scope, roleDefinitionID)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Role Definition %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Role Definition (Role Definition %q / Scope %q): %+v", roleDefinitionID, scope, err)
    }


    d.Set("name", resp.Name)
    if roleDefinitionProperties := resp.RoleDefinitionProperties; roleDefinitionProperties != nil {
        d.Set("assignable_scopes", utils.FlattenStringSlice(roleDefinitionProperties.AssignableScopes))
        d.Set("description", roleDefinitionProperties.Description)
        if err := d.Set("permissions", flattenArmRoleDefinitionPermission(roleDefinitionProperties.Permissions)); err != nil {
            return fmt.Errorf("Error setting `permissions`: %+v", err)
        }
        d.Set("role_name", roleDefinitionProperties.RoleName)
        d.Set("type", roleDefinitionProperties.Type)
    }
    d.Set("role_definition_id", roleDefinitionID)
    d.Set("scope", scope)
    d.Set("type", resp.Type)
    d.Set("type", resp.Type)
    d.Set("type", resp.Type)
    d.Set("type", resp.Type)

    return nil
}


func resourceArmRoleDefinitionDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).roleDefinitionsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    roleDefinitionID := id.Path["roleDefinitions"]

    if _, err := client.Delete(ctx, scope, roleDefinitionID); err != nil {
        return fmt.Errorf("Error deleting Role Definition (Role Definition %q / Scope %q): %+v", roleDefinitionID, scope, err)
    }

    return nil
}

func expandArmRoleDefinitionPermission(input []interface{}) *[]authorization.Permission {
    results := make([]authorization.Permission, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        actions := v["actions"].([]interface{})
        notActions := v["not_actions"].([]interface{})
        dataActions := v["data_actions"].([]interface{})
        notDataActions := v["not_data_actions"].([]interface{})

        result := authorization.Permission{
            Actions: utils.ExpandStringSlice(actions),
            DataActions: utils.ExpandStringSlice(dataActions),
            NotActions: utils.ExpandStringSlice(notActions),
            NotDataActions: utils.ExpandStringSlice(notDataActions),
        }

        results = append(results, result)
    }
    return &results
}


func flattenArmRoleDefinitionPermission(input *[]authorization.Permission) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        v["actions"] = utils.FlattenStringSlice(item.Actions)
        v["data_actions"] = utils.FlattenStringSlice(item.DataActions)
        v["not_actions"] = utils.FlattenStringSlice(item.NotActions)
        v["not_data_actions"] = utils.FlattenStringSlice(item.NotDataActions)

        results = append(results, v)
    }

    return results
}
