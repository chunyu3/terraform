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



func resourceArmApplicationTypeVersion() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmApplicationTypeVersionCreateUpdate,
        Read: resourceArmApplicationTypeVersionRead,
        Update: resourceArmApplicationTypeVersionCreateUpdate,
        Delete: resourceArmApplicationTypeVersionDelete,

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

            "app_package_url": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "cluster_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "version": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "etag": {
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

func resourceArmApplicationTypeVersionCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).applicationTypeVersionsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    clusterName := d.Get("cluster_name").(string)
    version := d.Get("version").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, clusterName, name, version)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Application Type Version %q (Version %q / Cluster Name %q / Resource Group %q): %+v", name, version, clusterName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_application_type_version", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    appPackageUrl := d.Get("app_package_url").(string)
    t := d.Get("tags").(map[string]interface{})

    parameters := servicefabric.ApplicationTypeVersionResource{
        Location: utils.String(location),
        ApplicationTypeVersionResourceProperties: &servicefabric.ApplicationTypeVersionResourceProperties{
            AppPackageURL: utils.String(appPackageUrl),
        },
        Tags: tags.Expand(t),
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, clusterName, name, version, parameters)
    if err != nil {
        return fmt.Errorf("Error creating Application Type Version %q (Version %q / Cluster Name %q / Resource Group %q): %+v", name, version, clusterName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Application Type Version %q (Version %q / Cluster Name %q / Resource Group %q): %+v", name, version, clusterName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, clusterName, name, version)
    if err != nil {
        return fmt.Errorf("Error retrieving Application Type Version %q (Version %q / Cluster Name %q / Resource Group %q): %+v", name, version, clusterName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Application Type Version %q (Version %q / Cluster Name %q / Resource Group %q) ID", name, version, clusterName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmApplicationTypeVersionRead(d, meta)
}

func resourceArmApplicationTypeVersionRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).applicationTypeVersionsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    clusterName := id.Path["clusters"]
    name := id.Path["applicationTypes"]
    version := id.Path["versions"]

    resp, err := client.Get(ctx, resourceGroup, clusterName, name, version)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Application Type Version %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Application Type Version %q (Version %q / Cluster Name %q / Resource Group %q): %+v", name, version, clusterName, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    d.Set("cluster_name", clusterName)
    d.Set("etag", resp.Etag)
    d.Set("type", resp.Type)
    d.Set("version", version)

    return nil
}


func resourceArmApplicationTypeVersionDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).applicationTypeVersionsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    clusterName := id.Path["clusters"]
    name := id.Path["applicationTypes"]
    version := id.Path["versions"]

    future, err := client.Delete(ctx, resourceGroup, clusterName, name, version)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Application Type Version %q (Version %q / Cluster Name %q / Resource Group %q): %+v", name, version, clusterName, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Application Type Version %q (Version %q / Cluster Name %q / Resource Group %q): %+v", name, version, clusterName, resourceGroup, err)
        }
    }

    return nil
}
