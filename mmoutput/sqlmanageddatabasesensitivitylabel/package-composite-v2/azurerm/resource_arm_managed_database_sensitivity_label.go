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



func resourceArmManagedDatabaseSensitivityLabel() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmManagedDatabaseSensitivityLabelCreateUpdate,
        Read: resourceArmManagedDatabaseSensitivityLabelRead,
        Update: resourceArmManagedDatabaseSensitivityLabelCreateUpdate,
        Delete: resourceArmManagedDatabaseSensitivityLabelDelete,

        Importer: &schema.ResourceImporter{
            State: schema.ImportStatePassthrough,
        },


        Schema: map[string]*schema.Schema{
            "name": {
                Type: schema.TypeString,
                Computed: true,
            },

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "column_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "database_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "managed_instance_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "schema_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "sensitivity_label_source": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "table_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "information_type": {
                Type: schema.TypeString,
                Optional: true,
            },

            "information_type_id": {
                Type: schema.TypeString,
                Optional: true,
            },

            "label_id": {
                Type: schema.TypeString,
                Optional: true,
            },

            "label_name": {
                Type: schema.TypeString,
                Optional: true,
            },

            "is_disabled": {
                Type: schema.TypeBool,
                Computed: true,
            },

            "type": {
                Type: schema.TypeString,
                Computed: true,
            },
        },
    }
}

func resourceArmManagedDatabaseSensitivityLabelCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).managedDatabaseSensitivityLabelsClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    columnName := d.Get("column_name").(string)
    databaseName := d.Get("database_name").(string)
    managedInstanceName := d.Get("managed_instance_name").(string)
    schemaName := d.Get("schema_name").(string)
    sensitivityLabelSource := d.Get("sensitivity_label_source").(string)
    tableName := d.Get("table_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, managedInstanceName, databaseName, schemaName, tableName, columnName, sensitivityLabelSource)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Managed Database Sensitivity Label (Sensitivity Label Source %q / Column Name %q / Table Name %q / Schema Name %q / Database Name %q / Managed Instance Name %q / Resource Group %q): %+v", sensitivityLabelSource, columnName, tableName, schemaName, databaseName, managedInstanceName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_managed_database_sensitivity_label", *existing.ID)
        }
    }

    informationType := d.Get("information_type").(string)
    informationTypeId := d.Get("information_type_id").(string)
    labelId := d.Get("label_id").(string)
    labelName := d.Get("label_name").(string)

    parameters := sql.SensitivityLabel{
        SensitivityLabelProperties: &sql.SensitivityLabelProperties{
            InformationType: utils.String(informationType),
            InformationTypeID: utils.String(informationTypeId),
            LabelID: utils.String(labelId),
            LabelName: utils.String(labelName),
        },
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroup, managedInstanceName, databaseName, schemaName, tableName, columnName, sensitivityLabelSource, parameters); err != nil {
        return fmt.Errorf("Error creating Managed Database Sensitivity Label (Sensitivity Label Source %q / Column Name %q / Table Name %q / Schema Name %q / Database Name %q / Managed Instance Name %q / Resource Group %q): %+v", sensitivityLabelSource, columnName, tableName, schemaName, databaseName, managedInstanceName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, managedInstanceName, databaseName, schemaName, tableName, columnName, sensitivityLabelSource)
    if err != nil {
        return fmt.Errorf("Error retrieving Managed Database Sensitivity Label (Sensitivity Label Source %q / Column Name %q / Table Name %q / Schema Name %q / Database Name %q / Managed Instance Name %q / Resource Group %q): %+v", sensitivityLabelSource, columnName, tableName, schemaName, databaseName, managedInstanceName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Managed Database Sensitivity Label (Sensitivity Label Source %q / Column Name %q / Table Name %q / Schema Name %q / Database Name %q / Managed Instance Name %q / Resource Group %q) ID", sensitivityLabelSource, columnName, tableName, schemaName, databaseName, managedInstanceName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmManagedDatabaseSensitivityLabelRead(d, meta)
}

func resourceArmManagedDatabaseSensitivityLabelRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).managedDatabaseSensitivityLabelsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    managedInstanceName := id.Path["managedInstances"]
    databaseName := id.Path["databases"]
    schemaName := id.Path["schemas"]
    tableName := id.Path["tables"]
    columnName := id.Path["columns"]
    sensitivityLabelSource := id.Path["sensitivityLabels"]

    resp, err := client.Get(ctx, resourceGroup, managedInstanceName, databaseName, schemaName, tableName, columnName, sensitivityLabelSource)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Managed Database Sensitivity Label %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Managed Database Sensitivity Label (Sensitivity Label Source %q / Column Name %q / Table Name %q / Schema Name %q / Database Name %q / Managed Instance Name %q / Resource Group %q): %+v", sensitivityLabelSource, columnName, tableName, schemaName, databaseName, managedInstanceName, resourceGroup, err)
    }


    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    d.Set("column_name", columnName)
    d.Set("database_name", databaseName)
    if sensitivityLabelProperties := resp.SensitivityLabelProperties; sensitivityLabelProperties != nil {
        d.Set("information_type", sensitivityLabelProperties.InformationType)
        d.Set("information_type_id", sensitivityLabelProperties.InformationTypeID)
        d.Set("is_disabled", sensitivityLabelProperties.IsDisabled)
        d.Set("label_id", sensitivityLabelProperties.LabelID)
        d.Set("label_name", sensitivityLabelProperties.LabelName)
    }
    d.Set("managed_instance_name", managedInstanceName)
    d.Set("schema_name", schemaName)
    d.Set("sensitivity_label_source", sensitivityLabelSource)
    d.Set("table_name", tableName)
    d.Set("type", resp.Type)

    return nil
}


func resourceArmManagedDatabaseSensitivityLabelDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).managedDatabaseSensitivityLabelsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    managedInstanceName := id.Path["managedInstances"]
    databaseName := id.Path["databases"]
    schemaName := id.Path["schemas"]
    tableName := id.Path["tables"]
    columnName := id.Path["columns"]
    sensitivityLabelSource := id.Path["sensitivityLabels"]

    if _, err := client.Delete(ctx, resourceGroup, managedInstanceName, databaseName, schemaName, tableName, columnName, sensitivityLabelSource); err != nil {
        return fmt.Errorf("Error deleting Managed Database Sensitivity Label (Sensitivity Label Source %q / Column Name %q / Table Name %q / Schema Name %q / Database Name %q / Managed Instance Name %q / Resource Group %q): %+v", sensitivityLabelSource, columnName, tableName, schemaName, databaseName, managedInstanceName, resourceGroup, err)
    }

    return nil
}
