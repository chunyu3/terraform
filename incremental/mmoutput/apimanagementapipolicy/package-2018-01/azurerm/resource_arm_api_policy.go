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
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

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

            "content_format": {
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

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    aPIID := d.Get("api_id").(string)
    policyID := d.Get("policy_id").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, name, aPIID, policyID)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Api Policy %q (Policy %q / Api %q / Resource Group %q): %+v", name, policyID, aPIID, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_api_policy", *existing.ID)
        }
    }

    contentFormat := d.Get("content_format").(string)
    policyContent := d.Get("policy_content").(string)

    parameters := apimanagement.PolicyContract{
        PolicyContractProperties: &apimanagement.PolicyContractProperties{
            ContentFormat: apimanagement.PolicyContentFormat(contentFormat),
            PolicyContent: utils.String(policyContent),
        },
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroup, name, aPIID, policyID, parameters); err != nil {
        return fmt.Errorf("Error creating Api Policy %q (Policy %q / Api %q / Resource Group %q): %+v", name, policyID, aPIID, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, name, aPIID, policyID)
    if err != nil {
        return fmt.Errorf("Error retrieving Api Policy %q (Policy %q / Api %q / Resource Group %q): %+v", name, policyID, aPIID, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Api Policy %q (Policy %q / Api %q / Resource Group %q) ID", name, policyID, aPIID, resourceGroup)
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
    name := id.Path["service"]
    aPIID := id.Path["apis"]
    policyID := id.Path["policies"]

    resp, err := client.Get(ctx, resourceGroup, name, aPIID, policyID)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Api Policy %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Api Policy %q (Policy %q / Api %q / Resource Group %q): %+v", name, policyID, aPIID, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    d.Set("api_id", aPIID)
    d.Set("policy_id", policyID)
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
    name := id.Path["service"]
    aPIID := id.Path["apis"]
    policyID := id.Path["policies"]

    if _, err := client.Delete(ctx, resourceGroup, name, aPIID, policyID); err != nil {
        return fmt.Errorf("Error deleting Api Policy %q (Policy %q / Api %q / Resource Group %q): %+v", name, policyID, aPIID, resourceGroup, err)
    }

    return nil
}
