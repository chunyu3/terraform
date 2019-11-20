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



func resourceArmServiceFabric() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmServiceFabricCreate,
        Read: resourceArmServiceFabricRead,
        Update: resourceArmServiceFabricUpdate,
        Delete: resourceArmServiceFabricDelete,

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
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "location": azure.SchemaLocation(),

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "lab_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "environment_id": {
                Type: schema.TypeString,
                Optional: true,
            },

            "external_service_fabric_id": {
                Type: schema.TypeString,
                Optional: true,
            },

            "type": {
                Type: schema.TypeString,
                Computed: true,
            },

            "tags": tags.Schema(),
        },
    }
}

func resourceArmServiceFabricCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).serviceFabricsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    labName := d.Get("lab_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, labName, name, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Service Fabric %q (Lab Name %q / Resource Group %q): %+v", name, labName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_service_fabric", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    environmentId := d.Get("environment_id").(string)
    externalServiceFabricId := d.Get("external_service_fabric_id").(string)
    t := d.Get("tags").(map[string]interface{})

    serviceFabric := devtestlab.ServiceFabricFragment{
        Location: utils.String(location),
        ServiceFabricPropertiesFragment: &devtestlab.ServiceFabricPropertiesFragment{
            EnvironmentID: utils.String(environmentId),
            ExternalServiceFabricID: utils.String(externalServiceFabricId),
        },
        Tags: tags.Expand(t),
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, labName, name, name, serviceFabric)
    if err != nil {
        return fmt.Errorf("Error creating Service Fabric %q (Lab Name %q / Resource Group %q): %+v", name, labName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Service Fabric %q (Lab Name %q / Resource Group %q): %+v", name, labName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, labName, name, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Service Fabric %q (Lab Name %q / Resource Group %q): %+v", name, labName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Service Fabric %q (Lab Name %q / Resource Group %q) ID", name, labName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmServiceFabricRead(d, meta)
}

func resourceArmServiceFabricRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).serviceFabricsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    labName := id.Path["labs"]
    name := id.Path["users"]
    name := id.Path["servicefabrics"]

    resp, err := client.Get(ctx, resourceGroup, labName, name, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Service Fabric %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Service Fabric %q (Lab Name %q / Resource Group %q): %+v", name, labName, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", name)
    d.Set("resource_group", resourceGroup)
    d.Set("lab_name", labName)
    d.Set("type", resp.Type)

    return nil
}

func resourceArmServiceFabricUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).serviceFabricsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    environmentId := d.Get("environment_id").(string)
    externalServiceFabricId := d.Get("external_service_fabric_id").(string)
    labName := d.Get("lab_name").(string)
    t := d.Get("tags").(map[string]interface{})

    serviceFabric := devtestlab.ServiceFabricFragment{
        ServiceFabricPropertiesFragment: &devtestlab.ServiceFabricPropertiesFragment{
            EnvironmentID: utils.String(environmentId),
            ExternalServiceFabricID: utils.String(externalServiceFabricId),
        },
        Tags: tags.Expand(t),
    }


    if _, err := client.Update(ctx, resourceGroup, labName, name, name, serviceFabric); err != nil {
        return fmt.Errorf("Error updating Service Fabric %q (Lab Name %q / Resource Group %q): %+v", name, labName, resourceGroup, err)
    }

    return resourceArmServiceFabricRead(d, meta)
}

func resourceArmServiceFabricDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).serviceFabricsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    labName := id.Path["labs"]
    name := id.Path["users"]
    name := id.Path["servicefabrics"]

    future, err := client.Delete(ctx, resourceGroup, labName, name, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Service Fabric %q (Lab Name %q / Resource Group %q): %+v", name, labName, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Service Fabric %q (Lab Name %q / Resource Group %q): %+v", name, labName, resourceGroup, err)
        }
    }

    return nil
}
