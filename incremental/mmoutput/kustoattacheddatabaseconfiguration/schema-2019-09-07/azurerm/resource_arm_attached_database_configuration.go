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



func resourceArmAttachedDatabaseConfiguration() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmAttachedDatabaseConfigurationCreateUpdate,
        Read: resourceArmAttachedDatabaseConfigurationRead,
        Update: resourceArmAttachedDatabaseConfigurationCreateUpdate,
        Delete: resourceArmAttachedDatabaseConfigurationDelete,

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

            "location": azure.SchemaLocation(),

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "cluster_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "cluster_resource_id": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "database_name": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "default_principals_modification_kind": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(kusto.Union),
                    string(kusto.Replace),
                    string(kusto.None),
                }, false),
            },

            "attached_database_names": {
                Type: schema.TypeList,
                Computed: true,
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
            },

            "provisioning_state": {
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

func resourceArmAttachedDatabaseConfigurationCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).attachedDatabaseConfigurationsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    clusterName := d.Get("cluster_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, clusterName, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Attached Database Configuration %q (Cluster Name %q / Resource Group %q): %+v", name, clusterName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_attached_database_configuration", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    clusterResourceId := d.Get("cluster_resource_id").(string)
    databaseName := d.Get("database_name").(string)
    defaultPrincipalsModificationKind := d.Get("default_principals_modification_kind").(string)

    parameters := kusto.AttachedDatabaseConfiguration{
        Location: utils.String(location),
        AttachedDatabaseConfigurationProperties: &kusto.AttachedDatabaseConfigurationProperties{
            ClusterResourceID: utils.String(clusterResourceId),
            DatabaseName: utils.String(databaseName),
            DefaultPrincipalsModificationKind: kusto.DefaultPrincipalsModificationKind(defaultPrincipalsModificationKind),
        },
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, clusterName, name, parameters)
    if err != nil {
        return fmt.Errorf("Error creating Attached Database Configuration %q (Cluster Name %q / Resource Group %q): %+v", name, clusterName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Attached Database Configuration %q (Cluster Name %q / Resource Group %q): %+v", name, clusterName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, clusterName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Attached Database Configuration %q (Cluster Name %q / Resource Group %q): %+v", name, clusterName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Attached Database Configuration %q (Cluster Name %q / Resource Group %q) ID", name, clusterName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmAttachedDatabaseConfigurationRead(d, meta)
}

func resourceArmAttachedDatabaseConfigurationRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).attachedDatabaseConfigurationsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    clusterName := id.Path["clusters"]
    name := id.Path["attachedDatabaseConfigurations"]

    resp, err := client.Get(ctx, resourceGroup, clusterName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Attached Database Configuration %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Attached Database Configuration %q (Cluster Name %q / Resource Group %q): %+v", name, clusterName, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if attachedDatabaseConfigurationProperties := resp.AttachedDatabaseConfigurationProperties; attachedDatabaseConfigurationProperties != nil {
        d.Set("attached_database_names", utils.FlattenStringSlice(attachedDatabaseConfigurationProperties.AttachedDatabaseNames))
        d.Set("cluster_resource_id", attachedDatabaseConfigurationProperties.ClusterResourceID)
        d.Set("database_name", attachedDatabaseConfigurationProperties.DatabaseName)
        d.Set("default_principals_modification_kind", string(attachedDatabaseConfigurationProperties.DefaultPrincipalsModificationKind))
        d.Set("provisioning_state", string(attachedDatabaseConfigurationProperties.ProvisioningState))
    }
    d.Set("cluster_name", clusterName)
    d.Set("type", resp.Type)

    return nil
}


func resourceArmAttachedDatabaseConfigurationDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).attachedDatabaseConfigurationsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    clusterName := id.Path["clusters"]
    name := id.Path["attachedDatabaseConfigurations"]

    future, err := client.Delete(ctx, resourceGroup, clusterName, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Attached Database Configuration %q (Cluster Name %q / Resource Group %q): %+v", name, clusterName, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Attached Database Configuration %q (Cluster Name %q / Resource Group %q): %+v", name, clusterName, resourceGroup, err)
        }
    }

    return nil
}
