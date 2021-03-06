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



func resourceArmRoleAssignment() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmRoleAssignmentCreateUpdate,
        Read: resourceArmRoleAssignmentRead,
        Update: resourceArmRoleAssignmentCreateUpdate,
        Delete: resourceArmRoleAssignmentDelete,

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

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "hub_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "principals": {
                Type: schema.TypeList,
                Required: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "principal_id": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "principal_type": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "principal_metadata": {
                            Type: schema.TypeMap,
                            Optional: true,
                            Elem: &schema.Schema{Type: schema.TypeString},
                        },
                    },
                },
            },

            "role": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(customerinsights.Admin),
                    string(customerinsights.Reader),
                    string(customerinsights.ManageAdmin),
                    string(customerinsights.ManageReader),
                    string(customerinsights.DataAdmin),
                    string(customerinsights.DataReader),
                }, false),
            },

            "conflation_policies": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "elements": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                        "exceptions": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                    },
                },
            },

            "connectors": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "elements": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                        "exceptions": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                    },
                },
            },

            "description": {
                Type: schema.TypeMap,
                Optional: true,
                Elem: &schema.Schema{Type: schema.TypeString},
            },

            "display_name": {
                Type: schema.TypeMap,
                Optional: true,
                Elem: &schema.Schema{Type: schema.TypeString},
            },

            "interactions": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "elements": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                        "exceptions": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                    },
                },
            },

            "kpis": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "elements": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                        "exceptions": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                    },
                },
            },

            "links": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "elements": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                        "exceptions": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                    },
                },
            },

            "profiles": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "elements": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                        "exceptions": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                    },
                },
            },

            "relationship_links": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "elements": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                        "exceptions": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                    },
                },
            },

            "relationships": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "elements": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                        "exceptions": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                    },
                },
            },

            "role_assignments": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "elements": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                        "exceptions": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                    },
                },
            },

            "sas_policies": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "elements": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                        "exceptions": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                    },
                },
            },

            "segments": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "elements": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                        "exceptions": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                    },
                },
            },

            "views": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "elements": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                        "exceptions": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                    },
                },
            },

            "widget_types": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "elements": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                        "exceptions": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                    },
                },
            },

            "assignment_name": {
                Type: schema.TypeString,
                Computed: true,
            },

            "provisioning_state": {
                Type: schema.TypeString,
                Computed: true,
            },

            "tenant_id": {
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

func resourceArmRoleAssignmentCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).roleAssignmentsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    hubName := d.Get("hub_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, hubName, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Role Assignment %q (Hub Name %q / Resource Group %q): %+v", name, hubName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_role_assignment", *existing.ID)
        }
    }

    conflationPolicies := d.Get("conflation_policies").([]interface{})
    connectors := d.Get("connectors").([]interface{})
    description := d.Get("description").(map[string]interface{})
    displayName := d.Get("display_name").(map[string]interface{})
    interactions := d.Get("interactions").([]interface{})
    kpis := d.Get("kpis").([]interface{})
    links := d.Get("links").([]interface{})
    principals := d.Get("principals").([]interface{})
    profiles := d.Get("profiles").([]interface{})
    relationshipLinks := d.Get("relationship_links").([]interface{})
    relationships := d.Get("relationships").([]interface{})
    role := d.Get("role").(string)
    roleAssignments := d.Get("role_assignments").([]interface{})
    sasPolicies := d.Get("sas_policies").([]interface{})
    segments := d.Get("segments").([]interface{})
    views := d.Get("views").([]interface{})
    widgetTypes := d.Get("widget_types").([]interface{})

    parameters := customerinsights.RoleAssignmentResourceFormat{
        RoleAssignment: &customerinsights.RoleAssignment{
            ConflationPolicies: expandArmRoleAssignmentResourceSetDescription(conflationPolicies),
            Connectors: expandArmRoleAssignmentResourceSetDescription(connectors),
            Description: utils.ExpandKeyValuePairs(description),
            DisplayName: utils.ExpandKeyValuePairs(displayName),
            Interactions: expandArmRoleAssignmentResourceSetDescription(interactions),
            Kpis: expandArmRoleAssignmentResourceSetDescription(kpis),
            Links: expandArmRoleAssignmentResourceSetDescription(links),
            Principals: expandArmRoleAssignmentAssignmentPrincipal(principals),
            Profiles: expandArmRoleAssignmentResourceSetDescription(profiles),
            RelationshipLinks: expandArmRoleAssignmentResourceSetDescription(relationshipLinks),
            Relationships: expandArmRoleAssignmentResourceSetDescription(relationships),
            Role: customerinsights.RoleTypes(role),
            RoleAssignments: expandArmRoleAssignmentResourceSetDescription(roleAssignments),
            SasPolicies: expandArmRoleAssignmentResourceSetDescription(sasPolicies),
            Segments: expandArmRoleAssignmentResourceSetDescription(segments),
            Views: expandArmRoleAssignmentResourceSetDescription(views),
            WidgetTypes: expandArmRoleAssignmentResourceSetDescription(widgetTypes),
        },
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, hubName, name, parameters)
    if err != nil {
        return fmt.Errorf("Error creating Role Assignment %q (Hub Name %q / Resource Group %q): %+v", name, hubName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Role Assignment %q (Hub Name %q / Resource Group %q): %+v", name, hubName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, hubName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Role Assignment %q (Hub Name %q / Resource Group %q): %+v", name, hubName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Role Assignment %q (Hub Name %q / Resource Group %q) ID", name, hubName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmRoleAssignmentRead(d, meta)
}

func resourceArmRoleAssignmentRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).roleAssignmentsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    hubName := id.Path["hubs"]
    name := id.Path["roleAssignments"]

    resp, err := client.Get(ctx, resourceGroup, hubName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Role Assignment %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Role Assignment %q (Hub Name %q / Resource Group %q): %+v", name, hubName, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if roleAssignment := resp.RoleAssignment; roleAssignment != nil {
        d.Set("assignment_name", roleAssignment.AssignmentName)
        if err := d.Set("conflation_policies", flattenArmRoleAssignmentResourceSetDescription(roleAssignment.ConflationPolicies)); err != nil {
            return fmt.Errorf("Error setting `conflation_policies`: %+v", err)
        }
        if err := d.Set("connectors", flattenArmRoleAssignmentResourceSetDescription(roleAssignment.Connectors)); err != nil {
            return fmt.Errorf("Error setting `connectors`: %+v", err)
        }
        d.Set("description", utils.FlattenKeyValuePairs(roleAssignment.Description))
        d.Set("display_name", utils.FlattenKeyValuePairs(roleAssignment.DisplayName))
        if err := d.Set("interactions", flattenArmRoleAssignmentResourceSetDescription(roleAssignment.Interactions)); err != nil {
            return fmt.Errorf("Error setting `interactions`: %+v", err)
        }
        if err := d.Set("kpis", flattenArmRoleAssignmentResourceSetDescription(roleAssignment.Kpis)); err != nil {
            return fmt.Errorf("Error setting `kpis`: %+v", err)
        }
        if err := d.Set("links", flattenArmRoleAssignmentResourceSetDescription(roleAssignment.Links)); err != nil {
            return fmt.Errorf("Error setting `links`: %+v", err)
        }
        if err := d.Set("principals", flattenArmRoleAssignmentAssignmentPrincipal(roleAssignment.Principals)); err != nil {
            return fmt.Errorf("Error setting `principals`: %+v", err)
        }
        if err := d.Set("profiles", flattenArmRoleAssignmentResourceSetDescription(roleAssignment.Profiles)); err != nil {
            return fmt.Errorf("Error setting `profiles`: %+v", err)
        }
        d.Set("provisioning_state", string(roleAssignment.ProvisioningState))
        if err := d.Set("relationship_links", flattenArmRoleAssignmentResourceSetDescription(roleAssignment.RelationshipLinks)); err != nil {
            return fmt.Errorf("Error setting `relationship_links`: %+v", err)
        }
        if err := d.Set("relationships", flattenArmRoleAssignmentResourceSetDescription(roleAssignment.Relationships)); err != nil {
            return fmt.Errorf("Error setting `relationships`: %+v", err)
        }
        d.Set("role", string(roleAssignment.Role))
        if err := d.Set("role_assignments", flattenArmRoleAssignmentResourceSetDescription(roleAssignment.RoleAssignments)); err != nil {
            return fmt.Errorf("Error setting `role_assignments`: %+v", err)
        }
        if err := d.Set("sas_policies", flattenArmRoleAssignmentResourceSetDescription(roleAssignment.SasPolicies)); err != nil {
            return fmt.Errorf("Error setting `sas_policies`: %+v", err)
        }
        if err := d.Set("segments", flattenArmRoleAssignmentResourceSetDescription(roleAssignment.Segments)); err != nil {
            return fmt.Errorf("Error setting `segments`: %+v", err)
        }
        d.Set("tenant_id", roleAssignment.TenantID)
        if err := d.Set("views", flattenArmRoleAssignmentResourceSetDescription(roleAssignment.Views)); err != nil {
            return fmt.Errorf("Error setting `views`: %+v", err)
        }
        if err := d.Set("widget_types", flattenArmRoleAssignmentResourceSetDescription(roleAssignment.WidgetTypes)); err != nil {
            return fmt.Errorf("Error setting `widget_types`: %+v", err)
        }
    }
    d.Set("hub_name", hubName)
    d.Set("type", resp.Type)

    return nil
}


func resourceArmRoleAssignmentDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).roleAssignmentsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    hubName := id.Path["hubs"]
    name := id.Path["roleAssignments"]

    if _, err := client.Delete(ctx, resourceGroup, hubName, name); err != nil {
        return fmt.Errorf("Error deleting Role Assignment %q (Hub Name %q / Resource Group %q): %+v", name, hubName, resourceGroup, err)
    }

    return nil
}

func expandArmRoleAssignmentResourceSetDescription(input []interface{}) *customerinsights.ResourceSetDescription {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    elements := v["elements"].([]interface{})
    exceptions := v["exceptions"].([]interface{})

    result := customerinsights.ResourceSetDescription{
        Elements: utils.ExpandStringSlice(elements),
        Exceptions: utils.ExpandStringSlice(exceptions),
    }
    return &result
}

func expandArmRoleAssignmentAssignmentPrincipal(input []interface{}) *[]customerinsights.AssignmentPrincipal {
    results := make([]customerinsights.AssignmentPrincipal, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        principalId := v["principal_id"].(string)
        principalType := v["principal_type"].(string)
        principalMetadata := v["principal_metadata"].(map[string]interface{})

        result := customerinsights.AssignmentPrincipal{
            PrincipalID: utils.String(principalId),
            PrincipalMetadata: utils.ExpandKeyValuePairs(principalMetadata),
            PrincipalType: utils.String(principalType),
        }

        results = append(results, result)
    }
    return &results
}


func flattenArmRoleAssignmentResourceSetDescription(input *customerinsights.ResourceSetDescription) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    result["elements"] = utils.FlattenStringSlice(input.Elements)
    result["exceptions"] = utils.FlattenStringSlice(input.Exceptions)

    return []interface{}{result}
}

func flattenArmRoleAssignmentAssignmentPrincipal(input *[]customerinsights.AssignmentPrincipal) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        if principalId := item.PrincipalID; principalId != nil {
            v["principal_id"] = *principalId
        }
        v["principal_metadata"] = utils.FlattenKeyValuePairs(item.PrincipalMetadata)
        if principalType := item.PrincipalType; principalType != nil {
            v["principal_type"] = *principalType
        }

        results = append(results, v)
    }

    return results
}
