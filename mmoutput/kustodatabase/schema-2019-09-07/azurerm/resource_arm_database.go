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



func resourceArmDatabase() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmDatabaseCreate,
        Read: resourceArmDatabaseRead,
        Update: resourceArmDatabaseUpdate,
        Delete: resourceArmDatabaseDelete,

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

            "type": {
                Type: schema.TypeString,
                Computed: true,
            },
        },
    }
}

func resourceArmDatabaseCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).databasesClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    clusterName := d.Get("cluster_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, clusterName, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Database %q (Cluster Name %q / Resource Group %q): %+v", name, clusterName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_database", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))

    parameters := kusto.Database{
        Location: utils.String(location),
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, clusterName, name, parameters)
    if err != nil {
        return fmt.Errorf("Error creating Database %q (Cluster Name %q / Resource Group %q): %+v", name, clusterName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Database %q (Cluster Name %q / Resource Group %q): %+v", name, clusterName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, clusterName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Database %q (Cluster Name %q / Resource Group %q): %+v", name, clusterName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Database %q (Cluster Name %q / Resource Group %q) ID", name, clusterName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmDatabaseRead(d, meta)
}

func resourceArmDatabaseRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).databasesClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    clusterName := id.Path["clusters"]
    name := id.Path["databases"]

    resp, err := client.Get(ctx, resourceGroup, clusterName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Database %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Database %q (Cluster Name %q / Resource Group %q): %+v", name, clusterName, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    d.Set("cluster_name", clusterName)
    d.Set("type", resp.Type)

    return nil
}

func resourceArmDatabaseUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).databasesClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    clusterName := d.Get("cluster_name").(string)

    parameters := kusto.Database{
        Location: utils.String(location),
    }


    future, err := client.Update(ctx, resourceGroup, clusterName, name, parameters)
    if err != nil {
        return fmt.Errorf("Error updating Database %q (Cluster Name %q / Resource Group %q): %+v", name, clusterName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for update of Database %q (Cluster Name %q / Resource Group %q): %+v", name, clusterName, resourceGroup, err)
    }

    return resourceArmDatabaseRead(d, meta)
}

func resourceArmDatabaseDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).databasesClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    clusterName := id.Path["clusters"]
    name := id.Path["databases"]

    future, err := client.Delete(ctx, resourceGroup, clusterName, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Database %q (Cluster Name %q / Resource Group %q): %+v", name, clusterName, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Database %q (Cluster Name %q / Resource Group %q): %+v", name, clusterName, resourceGroup, err)
        }
    }

    return nil
}
