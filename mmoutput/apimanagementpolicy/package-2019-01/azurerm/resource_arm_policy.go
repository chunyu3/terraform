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



func resourceArmPolicy() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmPolicyCreateUpdate,
        Read: resourceArmPolicyRead,
        Update: resourceArmPolicyCreateUpdate,
        Delete: resourceArmPolicyDelete,

        Importer: &schema.ResourceImporter{
            State: schema.ImportStatePassthrough,
        },


        Schema: map[string]*schema.Schema{
            "name": {
                Type: schema.TypeString,
                Computed: true,
            },

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "policy_id": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "service_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "value": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "format": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(apimanagement.xml),
                    string(apimanagement.xml-link),
                    string(apimanagement.rawxml),
                    string(apimanagement.rawxml-link),
                }, false),
                Default: string(apimanagement.xml),
            },

            "type": {
                Type: schema.TypeString,
                Computed: true,
            },
        },
    }
}

func resourceArmPolicyCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).policyClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    policyID := d.Get("policy_id").(string)
    serviceName := d.Get("service_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, serviceName, policyID)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Policy (Policy %q / Service Name %q / Resource Group %q): %+v", policyID, serviceName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_policy", *existing.ID)
        }
    }

    format := d.Get("format").(string)
    value := d.Get("value").(string)

    parameters := apimanagement.PolicyContract{
        PolicyContractProperties: &apimanagement.PolicyContractProperties{
            Format: apimanagement.PolicyContentFormat(format),
            Value: utils.String(value),
        },
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroup, serviceName, policyID, parameters); err != nil {
        return fmt.Errorf("Error creating Policy (Policy %q / Service Name %q / Resource Group %q): %+v", policyID, serviceName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, serviceName, policyID)
    if err != nil {
        return fmt.Errorf("Error retrieving Policy (Policy %q / Service Name %q / Resource Group %q): %+v", policyID, serviceName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Policy (Policy %q / Service Name %q / Resource Group %q) ID", policyID, serviceName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmPolicyRead(d, meta)
}

func resourceArmPolicyRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).policyClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    serviceName := id.Path["service"]
    policyID := id.Path["policies"]

    resp, err := client.Get(ctx, resourceGroup, serviceName, policyID)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Policy %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Policy (Policy %q / Service Name %q / Resource Group %q): %+v", policyID, serviceName, resourceGroup, err)
    }


    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if policyContractProperties := resp.PolicyContractProperties; policyContractProperties != nil {
        d.Set("format", string(policyContractProperties.Format))
        d.Set("value", policyContractProperties.Value)
    }
    d.Set("policy_id", policyID)
    d.Set("service_name", serviceName)
    d.Set("type", resp.Type)

    return nil
}


func resourceArmPolicyDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).policyClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    serviceName := id.Path["service"]
    policyID := id.Path["policies"]

    if _, err := client.Delete(ctx, resourceGroup, serviceName, policyID); err != nil {
        return fmt.Errorf("Error deleting Policy (Policy %q / Service Name %q / Resource Group %q): %+v", policyID, serviceName, resourceGroup, err)
    }

    return nil
}
