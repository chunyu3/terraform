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



func resourceArmApiPolicy() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmApiPolicyCreateUpdate,
        Read: resourceArmApiPolicyRead,
        Update: resourceArmApiPolicyCreateUpdate,
        Delete: resourceArmApiPolicyDelete,

        Importer: &schema.ResourceImporter{
            State: schema.ImportStatePassthrough,
        },


        Schema: map[string]*schema.Schema{
            "name": {
                Type: schema.TypeString,
                Computed: true,
            },

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "api_id": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

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

func resourceArmApiPolicyCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).apiPolicyClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    apiID := d.Get("api_id").(string)
    policyID := d.Get("policy_id").(string)
    serviceName := d.Get("service_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, serviceName, apiID, policyID)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Api Policy (Policy %q / Api %q / Service Name %q / Resource Group %q): %+v", policyID, apiID, serviceName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_api_policy", *existing.ID)
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


    if _, err := client.CreateOrUpdate(ctx, resourceGroup, serviceName, apiID, policyID, parameters); err != nil {
        return fmt.Errorf("Error creating Api Policy (Policy %q / Api %q / Service Name %q / Resource Group %q): %+v", policyID, apiID, serviceName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, serviceName, apiID, policyID)
    if err != nil {
        return fmt.Errorf("Error retrieving Api Policy (Policy %q / Api %q / Service Name %q / Resource Group %q): %+v", policyID, apiID, serviceName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Api Policy (Policy %q / Api %q / Service Name %q / Resource Group %q) ID", policyID, apiID, serviceName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmApiPolicyRead(d, meta)
}

func resourceArmApiPolicyRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).apiPolicyClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    serviceName := id.Path["service"]
    apiID := id.Path["apis"]
    policyID := id.Path["policies"]

    resp, err := client.Get(ctx, resourceGroup, serviceName, apiID, policyID)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Api Policy %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Api Policy (Policy %q / Api %q / Service Name %q / Resource Group %q): %+v", policyID, apiID, serviceName, resourceGroup, err)
    }


    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    d.Set("api_id", apiID)
    if policyContractProperties := resp.PolicyContractProperties; policyContractProperties != nil {
        d.Set("format", string(policyContractProperties.Format))
        d.Set("value", policyContractProperties.Value)
    }
    d.Set("policy_id", policyID)
    d.Set("service_name", serviceName)
    d.Set("type", resp.Type)

    return nil
}


func resourceArmApiPolicyDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).apiPolicyClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    serviceName := id.Path["service"]
    apiID := id.Path["apis"]
    policyID := id.Path["policies"]

    if _, err := client.Delete(ctx, resourceGroup, serviceName, apiID, policyID); err != nil {
        return fmt.Errorf("Error deleting Api Policy (Policy %q / Api %q / Service Name %q / Resource Group %q): %+v", policyID, apiID, serviceName, resourceGroup, err)
    }

    return nil
}
