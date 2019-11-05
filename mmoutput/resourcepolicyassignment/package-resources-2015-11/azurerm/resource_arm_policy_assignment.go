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



func resourceArmPolicyAssignment() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmPolicyAssignmentCreateUpdate,
        Read: resourceArmPolicyAssignmentRead,
        Update: resourceArmPolicyAssignmentCreateUpdate,
        Delete: resourceArmPolicyAssignmentDelete,

        Importer: &schema.ResourceImporter{
            State: schema.ImportStatePassthrough,
        },


        Schema: map[string]*schema.Schema{
            "name": {
                Type: schema.TypeString,
                Computed: true,
            },

            "policy_assignment_name": {
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

            "display_name": {
                Type: schema.TypeString,
                Optional: true,
            },

            "policy_definition_id": {
                Type: schema.TypeString,
                Optional: true,
            },

            "scope": {
                Type: schema.TypeString,
                Optional: true,
            },
        },
    }
}

func resourceArmPolicyAssignmentCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).policyAssignmentsClient
    ctx := meta.(*ArmClient).StopContext

    policyAssignmentName := d.Get("policy_assignment_name").(string)
    scope := d.Get("scope").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, scope, policyAssignmentName)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Policy Assignment (Policy Assignment Name %q / Scope %q): %+v", policyAssignmentName, scope, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_policy_assignment", *existing.ID)
        }
    }

    displayName := d.Get("display_name").(string)
    policyDefinitionId := d.Get("policy_definition_id").(string)
    scope := d.Get("scope").(string)

    parameters := resource.PolicyAssignment{
        PolicyAssignmentProperties: &resource.PolicyAssignmentProperties{
            DisplayName: utils.String(displayName),
            PolicyDefinitionID: utils.String(policyDefinitionId),
            Scope: utils.String(scope),
        },
    }


    if _, err := client.Create(ctx, scope, policyAssignmentName, parameters); err != nil {
        return fmt.Errorf("Error creating Policy Assignment (Policy Assignment Name %q / Scope %q): %+v", policyAssignmentName, scope, err)
    }


    resp, err := client.Get(ctx, scope, policyAssignmentName)
    if err != nil {
        return fmt.Errorf("Error retrieving Policy Assignment (Policy Assignment Name %q / Scope %q): %+v", policyAssignmentName, scope, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Policy Assignment (Policy Assignment Name %q / Scope %q) ID", policyAssignmentName, scope)
    }
    d.SetId(*resp.ID)

    return resourceArmPolicyAssignmentRead(d, meta)
}

func resourceArmPolicyAssignmentRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).policyAssignmentsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    policyAssignmentName := id.Path["policyAssignments"]

    resp, err := client.Get(ctx, scope, policyAssignmentName)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Policy Assignment %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Policy Assignment (Policy Assignment Name %q / Scope %q): %+v", policyAssignmentName, scope, err)
    }


    d.Set("name", resp.Name)
    if policyAssignmentProperties := resp.PolicyAssignmentProperties; policyAssignmentProperties != nil {
        d.Set("display_name", policyAssignmentProperties.DisplayName)
        d.Set("policy_definition_id", policyAssignmentProperties.PolicyDefinitionID)
        d.Set("scope", policyAssignmentProperties.Scope)
    }
    d.Set("policy_assignment_name", policyAssignmentName)
    d.Set("scope", scope)

    return nil
}


func resourceArmPolicyAssignmentDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).policyAssignmentsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    policyAssignmentName := id.Path["policyAssignments"]

    if _, err := client.Delete(ctx, scope, policyAssignmentName); err != nil {
        return fmt.Errorf("Error deleting Policy Assignment (Policy Assignment Name %q / Scope %q): %+v", policyAssignmentName, scope, err)
    }

    return nil
}
