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



func resourceArmAssignment() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmAssignmentCreateUpdate,
        Read: resourceArmAssignmentRead,
        Update: resourceArmAssignmentCreateUpdate,
        Delete: resourceArmAssignmentDelete,

        Importer: &schema.ResourceImporter{
            State: schema.ImportStatePassthrough,
        },


        Schema: map[string]*schema.Schema{
            "name": {
                Type: schema.TypeString,
                Computed: true,
            },

            "location": azure.SchemaLocation(),

            "assignment_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "identity": {
                Type: schema.TypeList,
                Required: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "type": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(blueprint.None),
                                string(blueprint.SystemAssigned),
                                string(blueprint.UserAssigned),
                            }, false),
                        },
                        "principal_id": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "tenant_id": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },

            "parameters": {
                Type: schema.TypeMap,
                Required: true,
                Elem: &schema.Schema{Type: schema.TypeString},
            },

            "resource_groups": {
                Type: schema.TypeMap,
                Required: true,
                Elem: &schema.Schema{Type: schema.TypeString},
            },

            "blueprint_id": {
                Type: schema.TypeString,
                Optional: true,
            },

            "description": {
                Type: schema.TypeString,
                Optional: true,
            },

            "display_name": {
                Type: schema.TypeString,
                Optional: true,
            },

            "locks": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "mode": {
                            Type: schema.TypeString,
                            Optional: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(blueprint.None),
                                string(blueprint.AllResources),
                            }, false),
                            Default: string(blueprint.None),
                        },
                    },
                },
            },

            "provisioning_state": {
                Type: schema.TypeString,
                Computed: true,
            },

            "status": {
                Type: schema.TypeList,
                Computed: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "last_modified": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "time_created": {
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

func resourceArmAssignmentCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).assignmentsClient
    ctx := meta.(*ArmClient).StopContext

    assignmentName := d.Get("assignment_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, assignmentName)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Assignment (Assignment Name %q): %+v", assignmentName, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_assignment", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    blueprintId := d.Get("blueprint_id").(string)
    description := d.Get("description").(string)
    displayName := d.Get("display_name").(string)
    identity := d.Get("identity").([]interface{})
    locks := d.Get("locks").([]interface{})
    parameters := d.Get("parameters").(map[string]interface{})
    resourceGroups := d.Get("resource_groups").(map[string]interface{})

    assignment := blueprint.Assignment{
        Identity: expandArmAssignmentManagedServiceIdentity(identity),
        Location: utils.String(location),
        AssignmentProperties: &blueprint.AssignmentProperties{
            BlueprintID: utils.String(blueprintId),
            Description: utils.String(description),
            DisplayName: utils.String(displayName),
            Locks: expandArmAssignmentAssignmentLockSettings(locks),
            Parameters: utils.ExpandKeyValuePairs(parameters),
            ResourceGroups: utils.ExpandKeyValuePairs(resourceGroups),
        },
    }


    if _, err := client.CreateOrUpdate(ctx, assignmentName, assignment); err != nil {
        return fmt.Errorf("Error creating Assignment (Assignment Name %q): %+v", assignmentName, err)
    }


    resp, err := client.Get(ctx, assignmentName)
    if err != nil {
        return fmt.Errorf("Error retrieving Assignment (Assignment Name %q): %+v", assignmentName, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Assignment (Assignment Name %q) ID", assignmentName)
    }
    d.SetId(*resp.ID)

    return resourceArmAssignmentRead(d, meta)
}

func resourceArmAssignmentRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).assignmentsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    assignmentName := id.Path["blueprintAssignments"]

    resp, err := client.Get(ctx, assignmentName)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Assignment %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Assignment (Assignment Name %q): %+v", assignmentName, err)
    }


    d.Set("name", resp.Name)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    d.Set("assignment_name", assignmentName)
    if assignmentProperties := resp.AssignmentProperties; assignmentProperties != nil {
        d.Set("blueprint_id", assignmentProperties.BlueprintID)
        d.Set("description", assignmentProperties.Description)
        d.Set("display_name", assignmentProperties.DisplayName)
        if err := d.Set("locks", flattenArmAssignmentAssignmentLockSettings(assignmentProperties.Locks)); err != nil {
            return fmt.Errorf("Error setting `locks`: %+v", err)
        }
        d.Set("parameters", utils.FlattenKeyValuePairs(assignmentProperties.Parameters))
        d.Set("provisioning_state", string(assignmentProperties.ProvisioningState))
        d.Set("resource_groups", utils.FlattenKeyValuePairs(assignmentProperties.ResourceGroups))
        if err := d.Set("status", flattenArmAssignmentAssignmentStatus(assignmentProperties.Status)); err != nil {
            return fmt.Errorf("Error setting `status`: %+v", err)
        }
    }
    if err := d.Set("identity", flattenArmAssignmentManagedServiceIdentity(resp.Identity)); err != nil {
        return fmt.Errorf("Error setting `identity`: %+v", err)
    }
    d.Set("type", resp.Type)

    return nil
}


func resourceArmAssignmentDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).assignmentsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    assignmentName := id.Path["blueprintAssignments"]

    if _, err := client.Delete(ctx, assignmentName); err != nil {
        return fmt.Errorf("Error deleting Assignment (Assignment Name %q): %+v", assignmentName, err)
    }

    return nil
}

func expandArmAssignmentManagedServiceIdentity(input []interface{}) *blueprint.ManagedServiceIdentity {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    type := v["type"].(string)
    principalId := v["principal_id"].(string)
    tenantId := v["tenant_id"].(string)

    result := blueprint.ManagedServiceIdentity{
        PrincipalID: utils.String(principalId),
        TenantID: utils.String(tenantId),
        Type: blueprint.ManagedServiceIdentityType(type),
    }
    return &result
}

func expandArmAssignmentAssignmentLockSettings(input []interface{}) *blueprint.AssignmentLockSettings {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    mode := v["mode"].(string)

    result := blueprint.AssignmentLockSettings{
        Mode: blueprint.AssignmentLockMode(mode),
    }
    return &result
}


func flattenArmAssignmentAssignmentLockSettings(input *blueprint.AssignmentLockSettings) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    result["mode"] = string(input.Mode)

    return []interface{}{result}
}

func flattenArmAssignmentAssignmentStatus(input *blueprint.AssignmentStatus) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})


    return []interface{}{result}
}

func flattenArmAssignmentManagedServiceIdentity(input *blueprint.ManagedServiceIdentity) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if principalId := input.PrincipalID; principalId != nil {
        result["principal_id"] = *principalId
    }
    if tenantId := input.TenantID; tenantId != nil {
        result["tenant_id"] = *tenantId
    }
    result["type"] = string(input.Type)

    return []interface{}{result}
}
