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



func resourceArmBinding() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmBindingCreate,
        Read: resourceArmBindingRead,
        Update: resourceArmBindingUpdate,
        Delete: resourceArmBindingDelete,

        Importer: &schema.ResourceImporter{
            State: schema.ImportStatePassthrough,
        },


        Schema: map[string]*schema.Schema{
            "name": {
                Type: schema.TypeString,
                Computed: true,
            },

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "app_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "binding_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "service_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "binding_parameters": {
                Type: schema.TypeMap,
                Optional: true,
                Elem: &schema.Schema{Type: schema.TypeString},
            },

            "key": {
                Type: schema.TypeString,
                Optional: true,
            },

            "resource_id": {
                Type: schema.TypeString,
                Optional: true,
            },

            "resource_name": {
                Type: schema.TypeString,
                Optional: true,
            },

            "resource_type": {
                Type: schema.TypeString,
                Optional: true,
            },

            "created_at": {
                Type: schema.TypeString,
                Computed: true,
            },

            "generated_properties": {
                Type: schema.TypeString,
                Computed: true,
            },

            "type": {
                Type: schema.TypeString,
                Computed: true,
            },

            "updated_at": {
                Type: schema.TypeString,
                Computed: true,
            },
        },
    }
}

func resourceArmBindingCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).bindingsClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    appName := d.Get("app_name").(string)
    bindingName := d.Get("binding_name").(string)
    serviceName := d.Get("service_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, serviceName, appName, bindingName)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Binding (Binding Name %q / App Name %q / Service Name %q / Resource Group %q): %+v", bindingName, appName, serviceName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_binding", *existing.ID)
        }
    }

    bindingParameters := d.Get("binding_parameters").(map[string]interface{})
    key := d.Get("key").(string)
    resourceId := d.Get("resource_id").(string)
    resourceName := d.Get("resource_name").(string)
    resourceType := d.Get("resource_type").(string)

    bindingResource := appplatform.BindingResource{
        BindingResourceProperties: &appplatform.BindingResourceProperties{
            BindingParameters: utils.ExpandKeyValuePairs(bindingParameters),
            Key: utils.String(key),
            ResourceID: utils.String(resourceId),
            ResourceName: utils.String(resourceName),
            ResourceType: utils.String(resourceType),
        },
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroup, serviceName, appName, bindingName, bindingResource); err != nil {
        return fmt.Errorf("Error creating Binding (Binding Name %q / App Name %q / Service Name %q / Resource Group %q): %+v", bindingName, appName, serviceName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, serviceName, appName, bindingName)
    if err != nil {
        return fmt.Errorf("Error retrieving Binding (Binding Name %q / App Name %q / Service Name %q / Resource Group %q): %+v", bindingName, appName, serviceName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Binding (Binding Name %q / App Name %q / Service Name %q / Resource Group %q) ID", bindingName, appName, serviceName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmBindingRead(d, meta)
}

func resourceArmBindingRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).bindingsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    serviceName := id.Path["Spring"]
    appName := id.Path["apps"]
    bindingName := id.Path["bindings"]

    resp, err := client.Get(ctx, resourceGroup, serviceName, appName, bindingName)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Binding %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Binding (Binding Name %q / App Name %q / Service Name %q / Resource Group %q): %+v", bindingName, appName, serviceName, resourceGroup, err)
    }


    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    d.Set("app_name", appName)
    d.Set("binding_name", bindingName)
    if bindingResourceProperties := resp.BindingResourceProperties; bindingResourceProperties != nil {
        d.Set("binding_parameters", utils.FlattenKeyValuePairs(bindingResourceProperties.BindingParameters))
        d.Set("created_at", bindingResourceProperties.CreatedAt)
        d.Set("generated_properties", bindingResourceProperties.GeneratedProperties)
        d.Set("key", bindingResourceProperties.Key)
        d.Set("resource_id", bindingResourceProperties.ResourceID)
        d.Set("resource_name", bindingResourceProperties.ResourceName)
        d.Set("resource_type", bindingResourceProperties.ResourceType)
        d.Set("updated_at", bindingResourceProperties.UpdatedAt)
    }
    d.Set("service_name", serviceName)
    d.Set("type", resp.Type)

    return nil
}

func resourceArmBindingUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).bindingsClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    appName := d.Get("app_name").(string)
    bindingName := d.Get("binding_name").(string)
    bindingParameters := d.Get("binding_parameters").(map[string]interface{})
    key := d.Get("key").(string)
    resourceId := d.Get("resource_id").(string)
    resourceName := d.Get("resource_name").(string)
    resourceType := d.Get("resource_type").(string)
    serviceName := d.Get("service_name").(string)

    bindingResource := appplatform.BindingResource{
        BindingResourceProperties: &appplatform.BindingResourceProperties{
            BindingParameters: utils.ExpandKeyValuePairs(bindingParameters),
            Key: utils.String(key),
            ResourceID: utils.String(resourceId),
            ResourceName: utils.String(resourceName),
            ResourceType: utils.String(resourceType),
        },
    }


    if _, err := client.Update(ctx, resourceGroup, serviceName, appName, bindingName, bindingResource); err != nil {
        return fmt.Errorf("Error updating Binding (Binding Name %q / App Name %q / Service Name %q / Resource Group %q): %+v", bindingName, appName, serviceName, resourceGroup, err)
    }

    return resourceArmBindingRead(d, meta)
}

func resourceArmBindingDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).bindingsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    serviceName := id.Path["Spring"]
    appName := id.Path["apps"]
    bindingName := id.Path["bindings"]

    if _, err := client.Delete(ctx, resourceGroup, serviceName, appName, bindingName); err != nil {
        return fmt.Errorf("Error deleting Binding (Binding Name %q / App Name %q / Service Name %q / Resource Group %q): %+v", bindingName, appName, serviceName, resourceGroup, err)
    }

    return nil
}
