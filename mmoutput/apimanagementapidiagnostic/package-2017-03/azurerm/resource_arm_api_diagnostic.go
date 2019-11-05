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



func resourceArmApiDiagnostic() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmApiDiagnosticCreate,
        Read: resourceArmApiDiagnosticRead,
        Update: resourceArmApiDiagnosticUpdate,
        Delete: resourceArmApiDiagnosticDelete,

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

            "diagnostic_id": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "enabled": {
                Type: schema.TypeBool,
                Required: true,
            },

            "service_name": {
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

func resourceArmApiDiagnosticCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).apiDiagnosticClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    apiID := d.Get("api_id").(string)
    diagnosticID := d.Get("diagnostic_id").(string)
    serviceName := d.Get("service_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, serviceName, apiID, diagnosticID)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Api Diagnostic (Diagnostic %q / Api %q / Service Name %q / Resource Group %q): %+v", diagnosticID, apiID, serviceName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_api_diagnostic", *existing.ID)
        }
    }

    enabled := d.Get("enabled").(bool)

    parameters := apimanagement.DiagnosticContract{
        DiagnosticContractProperties: &apimanagement.DiagnosticContractProperties{
            Enabled: utils.Bool(enabled),
        },
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroup, serviceName, apiID, diagnosticID, parameters); err != nil {
        return fmt.Errorf("Error creating Api Diagnostic (Diagnostic %q / Api %q / Service Name %q / Resource Group %q): %+v", diagnosticID, apiID, serviceName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, serviceName, apiID, diagnosticID)
    if err != nil {
        return fmt.Errorf("Error retrieving Api Diagnostic (Diagnostic %q / Api %q / Service Name %q / Resource Group %q): %+v", diagnosticID, apiID, serviceName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Api Diagnostic (Diagnostic %q / Api %q / Service Name %q / Resource Group %q) ID", diagnosticID, apiID, serviceName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmApiDiagnosticRead(d, meta)
}

func resourceArmApiDiagnosticRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).apiDiagnosticClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    serviceName := id.Path["service"]
    apiID := id.Path["apis"]
    diagnosticID := id.Path["diagnostics"]

    resp, err := client.Get(ctx, resourceGroup, serviceName, apiID, diagnosticID)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Api Diagnostic %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Api Diagnostic (Diagnostic %q / Api %q / Service Name %q / Resource Group %q): %+v", diagnosticID, apiID, serviceName, resourceGroup, err)
    }


    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    d.Set("api_id", apiID)
    d.Set("diagnostic_id", diagnosticID)
    if diagnosticContractProperties := resp.DiagnosticContractProperties; diagnosticContractProperties != nil {
        d.Set("enabled", diagnosticContractProperties.Enabled)
    }
    d.Set("service_name", serviceName)
    d.Set("type", resp.Type)

    return nil
}

func resourceArmApiDiagnosticUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).apiDiagnosticClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    apiID := d.Get("api_id").(string)
    diagnosticID := d.Get("diagnostic_id").(string)
    enabled := d.Get("enabled").(bool)
    serviceName := d.Get("service_name").(string)

    parameters := apimanagement.DiagnosticContract{
        DiagnosticContractProperties: &apimanagement.DiagnosticContractProperties{
            Enabled: utils.Bool(enabled),
        },
    }


    if _, err := client.Update(ctx, resourceGroup, serviceName, apiID, diagnosticID, parameters); err != nil {
        return fmt.Errorf("Error updating Api Diagnostic (Diagnostic %q / Api %q / Service Name %q / Resource Group %q): %+v", diagnosticID, apiID, serviceName, resourceGroup, err)
    }

    return resourceArmApiDiagnosticRead(d, meta)
}

func resourceArmApiDiagnosticDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).apiDiagnosticClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    serviceName := id.Path["service"]
    apiID := id.Path["apis"]
    diagnosticID := id.Path["diagnostics"]

    if _, err := client.Delete(ctx, resourceGroup, serviceName, apiID, diagnosticID); err != nil {
        return fmt.Errorf("Error deleting Api Diagnostic (Diagnostic %q / Api %q / Service Name %q / Resource Group %q): %+v", diagnosticID, apiID, serviceName, resourceGroup, err)
    }

    return nil
}