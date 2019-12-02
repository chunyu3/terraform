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



func resourceArmLogger() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmLoggerCreate,
        Read: resourceArmLoggerRead,
        Update: resourceArmLoggerUpdate,
        Delete: resourceArmLoggerDelete,

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

            "loggerid": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "credentials": {
                Type: schema.TypeMap,
                Optional: true,
                Elem: &schema.Schema{Type: schema.TypeString},
            },

            "description": {
                Type: schema.TypeString,
                Optional: true,
            },

            "is_buffered": {
                Type: schema.TypeBool,
                Optional: true,
            },

            "logger_type": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(apimanagement.azureEventHub),
                    string(apimanagement.applicationInsights),
                }, false),
                Default: string(apimanagement.azureEventHub),
            },

            "type": {
                Type: schema.TypeString,
                Computed: true,
            },
        },
    }
}

func resourceArmLoggerCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).loggerClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    loggerid := d.Get("loggerid").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, name, loggerid)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Logger %q (Loggerid %q / Resource Group %q): %+v", name, loggerid, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_logger", *existing.ID)
        }
    }

    credentials := d.Get("credentials").(map[string]interface{})
    description := d.Get("description").(string)
    isBuffered := d.Get("is_buffered").(bool)
    loggerType := d.Get("logger_type").(string)

    parameters := apimanagement.LoggerUpdateContract{
        LoggerUpdateParameters: &apimanagement.LoggerUpdateParameters{
            Credentials: utils.ExpandKeyValuePairs(credentials),
            Description: utils.String(description),
            IsBuffered: utils.Bool(isBuffered),
            LoggerType: apimanagement.LoggerType(loggerType),
        },
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroup, name, loggerid, parameters); err != nil {
        return fmt.Errorf("Error creating Logger %q (Loggerid %q / Resource Group %q): %+v", name, loggerid, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, name, loggerid)
    if err != nil {
        return fmt.Errorf("Error retrieving Logger %q (Loggerid %q / Resource Group %q): %+v", name, loggerid, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Logger %q (Loggerid %q / Resource Group %q) ID", name, loggerid, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmLoggerRead(d, meta)
}

func resourceArmLoggerRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).loggerClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["service"]
    loggerid := id.Path["loggers"]

    resp, err := client.Get(ctx, resourceGroup, name, loggerid)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Logger %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Logger %q (Loggerid %q / Resource Group %q): %+v", name, loggerid, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if loggerUpdateParameters := resp.LoggerUpdateParameters; loggerUpdateParameters != nil {
        d.Set("credentials", utils.FlattenKeyValuePairs(loggerUpdateParameters.Credentials))
        d.Set("description", loggerUpdateParameters.Description)
        d.Set("is_buffered", loggerUpdateParameters.IsBuffered)
        d.Set("logger_type", string(loggerUpdateParameters.LoggerType))
    }
    d.Set("loggerid", loggerid)
    d.Set("type", resp.Type)

    return nil
}

func resourceArmLoggerUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).loggerClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    credentials := d.Get("credentials").(map[string]interface{})
    description := d.Get("description").(string)
    isBuffered := d.Get("is_buffered").(bool)
    loggerType := d.Get("logger_type").(string)
    loggerid := d.Get("loggerid").(string)

    parameters := apimanagement.LoggerUpdateContract{
        LoggerUpdateParameters: &apimanagement.LoggerUpdateParameters{
            Credentials: utils.ExpandKeyValuePairs(credentials),
            Description: utils.String(description),
            IsBuffered: utils.Bool(isBuffered),
            LoggerType: apimanagement.LoggerType(loggerType),
        },
    }


    if _, err := client.Update(ctx, resourceGroup, name, loggerid, parameters); err != nil {
        return fmt.Errorf("Error updating Logger %q (Loggerid %q / Resource Group %q): %+v", name, loggerid, resourceGroup, err)
    }

    return resourceArmLoggerRead(d, meta)
}

func resourceArmLoggerDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).loggerClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["service"]
    loggerid := id.Path["loggers"]

    if _, err := client.Delete(ctx, resourceGroup, name, loggerid); err != nil {
        return fmt.Errorf("Error deleting Logger %q (Loggerid %q / Resource Group %q): %+v", name, loggerid, resourceGroup, err)
    }

    return nil
}
