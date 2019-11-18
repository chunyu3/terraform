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



func resourceArmCredential() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmCredentialCreate,
        Read: resourceArmCredentialRead,
        Update: resourceArmCredentialUpdate,
        Delete: resourceArmCredentialDelete,

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
                Optional: true,
                ForceNew: true,
            },

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "automation_account_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "description": {
                Type: schema.TypeString,
                Optional: true,
            },

            "password": {
                Type: schema.TypeString,
                Optional: true,
            },

            "user_name": {
                Type: schema.TypeString,
                Optional: true,
            },

            "creation_time": {
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

func resourceArmCredentialCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).credentialClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    automationAccountName := d.Get("automation_account_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, automationAccountName, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Credential %q (Automation Account Name %q / Resource Group %q): %+v", name, automationAccountName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_credential", *existing.ID)
        }
    }

    name := d.Get("name").(string)
    description := d.Get("description").(string)
    password := d.Get("password").(string)
    userName := d.Get("user_name").(string)

    parameters := automation.CredentialUpdateParameters{
        Name: utils.String(name),
        CredentialUpdateProperties: &automation.CredentialUpdateProperties{
            Description: utils.String(description),
            Password: utils.String(password),
            UserName: utils.String(userName),
        },
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroup, automationAccountName, name, parameters); err != nil {
        return fmt.Errorf("Error creating Credential %q (Automation Account Name %q / Resource Group %q): %+v", name, automationAccountName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, automationAccountName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Credential %q (Automation Account Name %q / Resource Group %q): %+v", name, automationAccountName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Credential %q (Automation Account Name %q / Resource Group %q) ID", name, automationAccountName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmCredentialRead(d, meta)
}

func resourceArmCredentialRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).credentialClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    automationAccountName := id.Path["automationAccounts"]
    name := id.Path["credentials"]

    resp, err := client.Get(ctx, resourceGroup, automationAccountName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Credential %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Credential %q (Automation Account Name %q / Resource Group %q): %+v", name, automationAccountName, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    d.Set("automation_account_name", automationAccountName)
    if credentialUpdateProperties := resp.CredentialUpdateProperties; credentialUpdateProperties != nil {
        d.Set("creation_time", (credentialUpdateProperties.CreationTime).String())
        d.Set("description", credentialUpdateProperties.Description)
        d.Set("last_modified_time", (credentialUpdateProperties.LastModifiedTime).String())
        d.Set("user_name", credentialUpdateProperties.UserName)
    }
    d.Set("type", resp.Type)

    return nil
}

func resourceArmCredentialUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).credentialClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    automationAccountName := d.Get("automation_account_name").(string)
    description := d.Get("description").(string)
    password := d.Get("password").(string)
    userName := d.Get("user_name").(string)

    parameters := automation.CredentialUpdateParameters{
        Name: utils.String(name),
        CredentialUpdateProperties: &automation.CredentialUpdateProperties{
            Description: utils.String(description),
            Password: utils.String(password),
            UserName: utils.String(userName),
        },
    }


    if _, err := client.Update(ctx, resourceGroup, automationAccountName, name, parameters); err != nil {
        return fmt.Errorf("Error updating Credential %q (Automation Account Name %q / Resource Group %q): %+v", name, automationAccountName, resourceGroup, err)
    }

    return resourceArmCredentialRead(d, meta)
}

func resourceArmCredentialDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).credentialClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    automationAccountName := id.Path["automationAccounts"]
    name := id.Path["credentials"]

    if _, err := client.Delete(ctx, resourceGroup, automationAccountName, name); err != nil {
        return fmt.Errorf("Error deleting Credential %q (Automation Account Name %q / Resource Group %q): %+v", name, automationAccountName, resourceGroup, err)
    }

    return nil
}
