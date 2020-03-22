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



func resourceArmVariable() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmVariableCreate,
        Read: resourceArmVariableRead,
        Update: resourceArmVariableUpdate,
        Delete: resourceArmVariableDelete,

        Importer: &schema.ResourceImporter{
            State: schema.ImportStatePassthrough,
        },


        Schema: map[string]*schema.Schema{
            "automation_account_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "variable_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "description": {
                Type: schema.TypeString,
                Optional: true,
            },

            "name": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
            },

            "value": {
                Type: schema.TypeString,
                Optional: true,
            },

            "creation_time": {
                Type: schema.TypeString,
                Computed: true,
            },

            "id": {
                Type: schema.TypeString,
                Computed: true,
            },

            "is_encrypted": {
                Type: schema.TypeBool,
                Computed: true,
            },

            "last_modified_time": {
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

func resourceArmVariableCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).variableClient
    ctx, cancel := timeouts.ForCreate(meta.(*ArmClient).StopContext, d)
    defer cancel()

    resourceGroupName := d.Get("resource_group").(string)
    automationAccountName := d.Get("automation_account_name").(string)
    name := d.Get("variable_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroupName, automationAccountName, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Variable (Variable Name %q / Automation Account Name %q / Resource Group %q): %+v", name, automationAccountName, resourceGroupName, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_variable", *existing.ID)
        }
    }

    description := d.Get("description").(string)
    name := d.Get("name").(string)
    value := d.Get("value").(string)

    parameters := automation.VariableUpdateParameters{
        Name: utils.String(name),
        VariableUpdateProperties: &automation.VariableUpdateProperties{
            Description: utils.String(description),
            Value: utils.String(value),
        },
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroupName, automationAccountName, name, parameters); err != nil {
        return fmt.Errorf("Error creating Variable (Variable Name %q / Automation Account Name %q / Resource Group %q): %+v", name, automationAccountName, resourceGroupName, err)
    }


    resp, err := client.Get(ctx, resourceGroupName, automationAccountName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Variable (Variable Name %q / Automation Account Name %q / Resource Group %q): %+v", name, automationAccountName, resourceGroupName, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Variable (Variable Name %q / Automation Account Name %q / Resource Group %q) ID", name, automationAccountName, resourceGroupName)
    }
    d.SetId(*resp.ID)

    return resourceArmVariableRead(d, meta)
}

func resourceArmVariableRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).variableClient
    ctx, cancel := timeouts.ForRead(meta.(*ArmClient).StopContext, d)
    defer cancel()

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroupName := id.ResourceGroup
    automationAccountName := id.Path["automationAccounts"]
    name := id.Path["variables"]

    resp, err := client.Get(ctx, resourceGroupName, automationAccountName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Variable %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Variable (Variable Name %q / Automation Account Name %q / Resource Group %q): %+v", name, automationAccountName, resourceGroupName, err)
    }


    d.Set("resource_group", resourceGroupName)
    d.Set("automation_account_name", automationAccountName)
    if variableUpdateProperties := resp.VariableUpdateProperties; variableUpdateProperties != nil {
        d.Set("creation_time", (variableUpdateProperties.CreationTime).String())
        d.Set("description", variableUpdateProperties.Description)
        d.Set("is_encrypted", variableUpdateProperties.IsEncrypted)
        d.Set("last_modified_time", (variableUpdateProperties.LastModifiedTime).String())
        d.Set("value", variableUpdateProperties.Value)
    }
    d.Set("id", resp.ID)
    d.Set("name", resp.Name)
    d.Set("type", resp.Type)
    d.Set("variable_name", name)

    return nil
}

func resourceArmVariableUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).variableClient
    ctx, cancel := timeouts.ForUpdate(meta.(*ArmClient).StopContext, d)
    defer cancel()

      resourceGroupName := d.Get("resource_group").(string)
    automationAccountName := d.Get("automation_account_name").(string)
    description := d.Get("description").(string)
    name := d.Get("name").(string)
    value := d.Get("value").(string)
    name := d.Get("variable_name").(string)

    parameters := automation.VariableUpdateParameters{
        Name: utils.String(name),
        VariableUpdateProperties: &automation.VariableUpdateProperties{
            Description: utils.String(description),
            Value: utils.String(value),
        },
    }


    if _, err := client.Update(ctx, resourceGroupName, automationAccountName, name, parameters); err != nil {
        return fmt.Errorf("Error updating Variable (Variable Name %q / Automation Account Name %q / Resource Group %q): %+v", name, automationAccountName, resourceGroupName, err)
    }

    return resourceArmVariableRead(d, meta)
}

func resourceArmVariableDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).variableClient
    ctx, cancel := timeouts.ForDelete(meta.(*ArmClient).StopContext, d)
    defer cancel()


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroupName := id.ResourceGroup
    automationAccountName := id.Path["automationAccounts"]
    name := id.Path["variables"]

    if _, err := client.Delete(ctx, resourceGroupName, automationAccountName, name); err != nil {
        return fmt.Errorf("Error deleting Variable (Variable Name %q / Automation Account Name %q / Resource Group %q): %+v", name, automationAccountName, resourceGroupName, err)
    }

    return nil
}
