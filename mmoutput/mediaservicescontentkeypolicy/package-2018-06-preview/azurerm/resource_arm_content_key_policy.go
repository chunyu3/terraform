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



func resourceArmContentKeyPolicy() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmContentKeyPolicyCreate,
        Read: resourceArmContentKeyPolicyRead,
        Update: resourceArmContentKeyPolicyUpdate,
        Delete: resourceArmContentKeyPolicyDelete,

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

            "account_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "options": {
                Type: schema.TypeList,
                Required: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "name": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },

            "description": {
                Type: schema.TypeString,
                Optional: true,
            },

            "created": {
                Type: schema.TypeString,
                Computed: true,
            },

            "last_modified": {
                Type: schema.TypeString,
                Computed: true,
            },

            "policy_id": {
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

func resourceArmContentKeyPolicyCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).contentKeyPoliciesClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    accountName := d.Get("account_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, accountName, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Content Key Policy %q (Account Name %q / Resource Group %q): %+v", name, accountName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_content_key_policy", *existing.ID)
        }
    }

    description := d.Get("description").(string)
    options := d.Get("options").([]interface{})

    parameters := mediaservices.ContentKeyPolicy{
        ContentKeyPolicyProperties: &mediaservices.ContentKeyPolicyProperties{
            Description: utils.String(description),
            Options: expandArmContentKeyPolicyContentKeyPolicyOption(options),
        },
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroup, accountName, name, parameters); err != nil {
        return fmt.Errorf("Error creating Content Key Policy %q (Account Name %q / Resource Group %q): %+v", name, accountName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, accountName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Content Key Policy %q (Account Name %q / Resource Group %q): %+v", name, accountName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Content Key Policy %q (Account Name %q / Resource Group %q) ID", name, accountName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmContentKeyPolicyRead(d, meta)
}

func resourceArmContentKeyPolicyRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).contentKeyPoliciesClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    accountName := id.Path["mediaServices"]
    name := id.Path["contentKeyPolicies"]

    resp, err := client.Get(ctx, resourceGroup, accountName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Content Key Policy %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Content Key Policy %q (Account Name %q / Resource Group %q): %+v", name, accountName, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    d.Set("account_name", accountName)
    if contentKeyPolicyProperties := resp.ContentKeyPolicyProperties; contentKeyPolicyProperties != nil {
        d.Set("created", (contentKeyPolicyProperties.Created).String())
        d.Set("description", contentKeyPolicyProperties.Description)
        d.Set("last_modified", (contentKeyPolicyProperties.LastModified).String())
        if err := d.Set("options", flattenArmContentKeyPolicyContentKeyPolicyOption(contentKeyPolicyProperties.Options)); err != nil {
            return fmt.Errorf("Error setting `options`: %+v", err)
        }
        d.Set("policy_id", contentKeyPolicyProperties.PolicyID)
    }
    d.Set("type", resp.Type)

    return nil
}

func resourceArmContentKeyPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).contentKeyPoliciesClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    accountName := d.Get("account_name").(string)
    description := d.Get("description").(string)
    options := d.Get("options").([]interface{})

    parameters := mediaservices.ContentKeyPolicy{
        ContentKeyPolicyProperties: &mediaservices.ContentKeyPolicyProperties{
            Description: utils.String(description),
            Options: expandArmContentKeyPolicyContentKeyPolicyOption(options),
        },
    }


    if _, err := client.Update(ctx, resourceGroup, accountName, name, parameters); err != nil {
        return fmt.Errorf("Error updating Content Key Policy %q (Account Name %q / Resource Group %q): %+v", name, accountName, resourceGroup, err)
    }

    return resourceArmContentKeyPolicyRead(d, meta)
}

func resourceArmContentKeyPolicyDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).contentKeyPoliciesClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    accountName := id.Path["mediaServices"]
    name := id.Path["contentKeyPolicies"]

    if _, err := client.Delete(ctx, resourceGroup, accountName, name); err != nil {
        return fmt.Errorf("Error deleting Content Key Policy %q (Account Name %q / Resource Group %q): %+v", name, accountName, resourceGroup, err)
    }

    return nil
}

func expandArmContentKeyPolicyContentKeyPolicyOption(input []interface{}) *[]mediaservices.ContentKeyPolicyOption {
    results := make([]mediaservices.ContentKeyPolicyOption, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        name := v["name"].(string)

        result := mediaservices.ContentKeyPolicyOption{
            Name: utils.String(name),
        }

        results = append(results, result)
    }
    return &results
}


func flattenArmContentKeyPolicyContentKeyPolicyOption(input *[]mediaservices.ContentKeyPolicyOption) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        if name := item.Name; name != nil {
            v["name"] = *name
        }

        results = append(results, v)
    }

    return results
}
