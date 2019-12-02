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



func resourceArmProductPolicy() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmProductPolicyCreateUpdate,
        Read: resourceArmProductPolicyRead,
        Update: resourceArmProductPolicyCreateUpdate,
        Delete: resourceArmProductPolicyDelete,

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

            "policy_content": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "policy_id": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "product_id": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "type": {
                Type: schema.TypeString,
                Computed: true,
            },
        },
    }
}

func resourceArmProductPolicyCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).productPolicyClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    policyID := d.Get("policy_id").(string)
    productID := d.Get("product_id").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, name, productID, policyID)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Product Policy %q (Policy %q / Product %q / Resource Group %q): %+v", name, policyID, productID, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_product_policy", *existing.ID)
        }
    }

    policyContent := d.Get("policy_content").(string)

    parameters := apimanagement.PolicyContract{
        PolicyContractProperties: &apimanagement.PolicyContractProperties{
            PolicyContent: utils.String(policyContent),
        },
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroup, name, productID, policyID, parameters); err != nil {
        return fmt.Errorf("Error creating Product Policy %q (Policy %q / Product %q / Resource Group %q): %+v", name, policyID, productID, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, name, productID, policyID)
    if err != nil {
        return fmt.Errorf("Error retrieving Product Policy %q (Policy %q / Product %q / Resource Group %q): %+v", name, policyID, productID, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Product Policy %q (Policy %q / Product %q / Resource Group %q) ID", name, policyID, productID, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmProductPolicyRead(d, meta)
}

func resourceArmProductPolicyRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).productPolicyClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["service"]
    productID := id.Path["products"]
    policyID := id.Path["policies"]

    resp, err := client.Get(ctx, resourceGroup, name, productID, policyID)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Product Policy %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Product Policy %q (Policy %q / Product %q / Resource Group %q): %+v", name, policyID, productID, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if policyContractProperties := resp.PolicyContractProperties; policyContractProperties != nil {
        d.Set("policy_content", policyContractProperties.PolicyContent)
    }
    d.Set("policy_id", policyID)
    d.Set("product_id", productID)
    d.Set("type", resp.Type)

    return nil
}


func resourceArmProductPolicyDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).productPolicyClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["service"]
    productID := id.Path["products"]
    policyID := id.Path["policies"]

    if _, err := client.Delete(ctx, resourceGroup, name, productID, policyID); err != nil {
        return fmt.Errorf("Error deleting Product Policy %q (Policy %q / Product %q / Resource Group %q): %+v", name, policyID, productID, resourceGroup, err)
    }

    return nil
}
