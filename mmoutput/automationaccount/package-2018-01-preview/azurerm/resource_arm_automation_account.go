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



func resourceArmAutomationAccount() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmAutomationAccountCreate,
        Read: resourceArmAutomationAccountRead,
        Update: resourceArmAutomationAccountUpdate,
        Delete: resourceArmAutomationAccountDelete,

        Importer: &schema.ResourceImporter{
            State: schema.ImportStatePassthrough,
        },


        Schema: map[string]*schema.Schema{
            "name": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
            },

            "location": azure.SchemaLocation(),

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "automation_account_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "sku": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "name": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(automation.Free),
                                string(automation.Basic),
                            }, false),
                        },
                        "capacity": {
                            Type: schema.TypeInt,
                            Optional: true,
                        },
                        "family": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },

            "creation_time": {
                Type: schema.TypeString,
                Computed: true,
            },

            "description": {
                Type: schema.TypeString,
                Computed: true,
            },

            "etag": {
                Type: schema.TypeString,
                Computed: true,
            },

            "last_modified_by": {
                Type: schema.TypeString,
                Computed: true,
            },

            "last_modified_time": {
                Type: schema.TypeString,
                Computed: true,
            },

            "state": {
                Type: schema.TypeString,
                Computed: true,
            },

            "type": {
                Type: schema.TypeString,
                Computed: true,
            },

            "tags": tags.Schema(),
        },
    }
}

func resourceArmAutomationAccountCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).automationAccountClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    automationAccountName := d.Get("automation_account_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, automationAccountName)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Automation Account (Automation Account Name %q / Resource Group %q): %+v", automationAccountName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_automation_account", *existing.ID)
        }
    }

    name := d.Get("name").(string)
    location := azure.NormalizeLocation(d.Get("location").(string))
    sku := d.Get("sku").([]interface{})
    t := d.Get("tags").(map[string]interface{})

    parameters := automation.AccountCreateOrUpdateParameters{
        Location: utils.String(location),
        Name: utils.String(name),
        AccountCreateOrUpdateProperties: &automation.AccountCreateOrUpdateProperties{
            Sku: expandArmAutomationAccountSku(sku),
        },
        Tags: tags.Expand(t),
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroup, automationAccountName, parameters); err != nil {
        return fmt.Errorf("Error creating Automation Account (Automation Account Name %q / Resource Group %q): %+v", automationAccountName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, automationAccountName)
    if err != nil {
        return fmt.Errorf("Error retrieving Automation Account (Automation Account Name %q / Resource Group %q): %+v", automationAccountName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Automation Account (Automation Account Name %q / Resource Group %q) ID", automationAccountName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmAutomationAccountRead(d, meta)
}

func resourceArmAutomationAccountRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).automationAccountClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    automationAccountName := id.Path["automationAccounts"]

    resp, err := client.Get(ctx, resourceGroup, automationAccountName)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Automation Account %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Automation Account (Automation Account Name %q / Resource Group %q): %+v", automationAccountName, resourceGroup, err)
    }


    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    d.Set("automation_account_name", automationAccountName)
    if accountCreateOrUpdateProperties := resp.AccountCreateOrUpdateProperties; accountCreateOrUpdateProperties != nil {
        d.Set("creation_time", (accountCreateOrUpdateProperties.CreationTime).String())
        d.Set("description", accountCreateOrUpdateProperties.Description)
        d.Set("last_modified_by", accountCreateOrUpdateProperties.LastModifiedBy)
        d.Set("last_modified_time", (accountCreateOrUpdateProperties.LastModifiedTime).String())
        if err := d.Set("sku", flattenArmAutomationAccountSku(accountCreateOrUpdateProperties.Sku)); err != nil {
            return fmt.Errorf("Error setting `sku`: %+v", err)
        }
        d.Set("state", string(accountCreateOrUpdateProperties.State))
    }
    d.Set("etag", resp.Etag)
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}

func resourceArmAutomationAccountUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).automationAccountClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    automationAccountName := d.Get("automation_account_name").(string)
    sku := d.Get("sku").([]interface{})
    t := d.Get("tags").(map[string]interface{})

    parameters := automation.AccountCreateOrUpdateParameters{
        Location: utils.String(location),
        Name: utils.String(name),
        AccountCreateOrUpdateProperties: &automation.AccountCreateOrUpdateProperties{
            Sku: expandArmAutomationAccountSku(sku),
        },
        Tags: tags.Expand(t),
    }


    if _, err := client.Update(ctx, resourceGroup, automationAccountName, parameters); err != nil {
        return fmt.Errorf("Error updating Automation Account (Automation Account Name %q / Resource Group %q): %+v", automationAccountName, resourceGroup, err)
    }

    return resourceArmAutomationAccountRead(d, meta)
}

func resourceArmAutomationAccountDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).automationAccountClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    automationAccountName := id.Path["automationAccounts"]

    if _, err := client.Delete(ctx, resourceGroup, automationAccountName); err != nil {
        return fmt.Errorf("Error deleting Automation Account (Automation Account Name %q / Resource Group %q): %+v", automationAccountName, resourceGroup, err)
    }

    return nil
}

func expandArmAutomationAccountSku(input []interface{}) *automation.Sku {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    name := v["name"].(string)
    family := v["family"].(string)
    capacity := v["capacity"].(int)

    result := automation.Sku{
        Capacity: utils.Int32(int32(capacity)),
        Family: utils.String(family),
        Name: automation.SkuNameEnum(name),
    }
    return &result
}


func flattenArmAutomationAccountSku(input *automation.Sku) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    result["name"] = string(input.Name)
    if capacity := input.Capacity; capacity != nil {
        result["capacity"] = int(*capacity)
    }
    if family := input.Family; family != nil {
        result["family"] = *family
    }

    return []interface{}{result}
}
