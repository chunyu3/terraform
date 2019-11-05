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



func resourceArmConnectionType() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmConnectionTypeCreateUpdate,
        Read: resourceArmConnectionTypeRead,
        Update: resourceArmConnectionTypeCreateUpdate,
        Delete: resourceArmConnectionTypeDelete,

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

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "automation_account_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "connection_type_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "field_definitions": {
                Type: schema.TypeMap,
                Required: true,
                Elem: &schema.Schema{Type: schema.TypeString},
            },

            "is_global": {
                Type: schema.TypeBool,
                Optional: true,
            },

            "creation_time": {
                Type: schema.TypeString,
                Computed: true,
            },

            "description": {
                Type: schema.TypeString,
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

func resourceArmConnectionTypeCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).connectionTypeClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    automationAccountName := d.Get("automation_account_name").(string)
    connectionTypeName := d.Get("connection_type_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, automationAccountName, connectionTypeName)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Connection Type (Connection Type Name %q / Automation Account Name %q / Resource Group %q): %+v", connectionTypeName, automationAccountName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_connection_type", *existing.ID)
        }
    }

    name := d.Get("name").(string)
    fieldDefinitions := d.Get("field_definitions").(map[string]interface{})
    isGlobal := d.Get("is_global").(bool)

    parameters := automation.ConnectionTypeCreateOrUpdateParameters{
        Name: utils.String(name),
        ConnectionTypeCreateOrUpdateProperties: &automation.ConnectionTypeCreateOrUpdateProperties{
            FieldDefinitions: utils.ExpandKeyValuePairs(fieldDefinitions),
            IsGlobal: utils.Bool(isGlobal),
        },
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroup, automationAccountName, connectionTypeName, parameters); err != nil {
        return fmt.Errorf("Error creating Connection Type (Connection Type Name %q / Automation Account Name %q / Resource Group %q): %+v", connectionTypeName, automationAccountName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, automationAccountName, connectionTypeName)
    if err != nil {
        return fmt.Errorf("Error retrieving Connection Type (Connection Type Name %q / Automation Account Name %q / Resource Group %q): %+v", connectionTypeName, automationAccountName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Connection Type (Connection Type Name %q / Automation Account Name %q / Resource Group %q) ID", connectionTypeName, automationAccountName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmConnectionTypeRead(d, meta)
}

func resourceArmConnectionTypeRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).connectionTypeClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    automationAccountName := id.Path["automationAccounts"]
    connectionTypeName := id.Path["connectionTypes"]

    resp, err := client.Get(ctx, resourceGroup, automationAccountName, connectionTypeName)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Connection Type %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Connection Type (Connection Type Name %q / Automation Account Name %q / Resource Group %q): %+v", connectionTypeName, automationAccountName, resourceGroup, err)
    }


    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    d.Set("automation_account_name", automationAccountName)
    d.Set("connection_type_name", connectionTypeName)
    if connectionTypeCreateOrUpdateProperties := resp.ConnectionTypeCreateOrUpdateProperties; connectionTypeCreateOrUpdateProperties != nil {
        d.Set("creation_time", (connectionTypeCreateOrUpdateProperties.CreationTime).String())
        d.Set("description", connectionTypeCreateOrUpdateProperties.Description)
        d.Set("field_definitions", utils.FlattenKeyValuePairs(connectionTypeCreateOrUpdateProperties.FieldDefinitions))
        d.Set("is_global", connectionTypeCreateOrUpdateProperties.IsGlobal)
        d.Set("last_modified_time", (connectionTypeCreateOrUpdateProperties.LastModifiedTime).String())
    }
    d.Set("type", resp.Type)

    return nil
}


func resourceArmConnectionTypeDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).connectionTypeClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    automationAccountName := id.Path["automationAccounts"]
    connectionTypeName := id.Path["connectionTypes"]

    if _, err := client.Delete(ctx, resourceGroup, automationAccountName, connectionTypeName); err != nil {
        return fmt.Errorf("Error deleting Connection Type (Connection Type Name %q / Automation Account Name %q / Resource Group %q): %+v", connectionTypeName, automationAccountName, resourceGroup, err)
    }

    return nil
}
