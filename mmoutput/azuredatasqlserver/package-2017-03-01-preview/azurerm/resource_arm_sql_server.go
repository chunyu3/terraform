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



func resourceArmSqlServer() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmSqlServerCreateUpdate,
        Read: resourceArmSqlServerRead,
        Update: resourceArmSqlServerCreateUpdate,
        Delete: resourceArmSqlServerDelete,

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

            "sql_server_registration_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "cores": {
                Type: schema.TypeInt,
                Optional: true,
            },

            "edition": {
                Type: schema.TypeString,
                Optional: true,
            },

            "property_bag": {
                Type: schema.TypeString,
                Optional: true,
            },

            "registration_id": {
                Type: schema.TypeString,
                Optional: true,
            },

            "version": {
                Type: schema.TypeString,
                Optional: true,
            },

            "type": {
                Type: schema.TypeString,
                Computed: true,
            },
        },
    }
}

func resourceArmSqlServerCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).sqlServersClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    sQLServerRegistrationName := d.Get("sql_server_registration_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, sQLServerRegistrationName, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Sql Server %q (Sql Server Registration Name %q / Resource Group %q): %+v", name, sQLServerRegistrationName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_sql_server", *existing.ID)
        }
    }

    cores := d.Get("cores").(int)
    edition := d.Get("edition").(string)
    propertyBag := d.Get("property_bag").(string)
    registrationId := d.Get("registration_id").(string)
    version := d.Get("version").(string)

    parameters := azuredata.SqlServer{
        SqlServerProperties: &azuredata.SqlServerProperties{
            Cores: utils.Int32(int32(cores)),
            Edition: utils.String(edition),
            PropertyBag: utils.String(propertyBag),
            RegistrationID: utils.String(registrationId),
            Version: utils.String(version),
        },
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroup, sQLServerRegistrationName, name, parameters); err != nil {
        return fmt.Errorf("Error creating Sql Server %q (Sql Server Registration Name %q / Resource Group %q): %+v", name, sQLServerRegistrationName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, sQLServerRegistrationName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Sql Server %q (Sql Server Registration Name %q / Resource Group %q): %+v", name, sQLServerRegistrationName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Sql Server %q (Sql Server Registration Name %q / Resource Group %q) ID", name, sQLServerRegistrationName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmSqlServerRead(d, meta)
}

func resourceArmSqlServerRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).sqlServersClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    sQLServerRegistrationName := id.Path["sqlServerRegistrations"]
    name := id.Path["sqlServers"]

    resp, err := client.Get(ctx, resourceGroup, sQLServerRegistrationName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Sql Server %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Sql Server %q (Sql Server Registration Name %q / Resource Group %q): %+v", name, sQLServerRegistrationName, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if sqlServerProperties := resp.SqlServerProperties; sqlServerProperties != nil {
        d.Set("cores", int(*sqlServerProperties.Cores))
        d.Set("edition", sqlServerProperties.Edition)
        d.Set("property_bag", sqlServerProperties.PropertyBag)
        d.Set("registration_id", sqlServerProperties.RegistrationID)
        d.Set("version", sqlServerProperties.Version)
    }
    d.Set("sql_server_registration_name", sQLServerRegistrationName)
    d.Set("type", resp.Type)

    return nil
}


func resourceArmSqlServerDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).sqlServersClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    sQLServerRegistrationName := id.Path["sqlServerRegistrations"]
    name := id.Path["sqlServers"]

    if _, err := client.Delete(ctx, resourceGroup, sQLServerRegistrationName, name); err != nil {
        return fmt.Errorf("Error deleting Sql Server %q (Sql Server Registration Name %q / Resource Group %q): %+v", name, sQLServerRegistrationName, resourceGroup, err)
    }

    return nil
}
