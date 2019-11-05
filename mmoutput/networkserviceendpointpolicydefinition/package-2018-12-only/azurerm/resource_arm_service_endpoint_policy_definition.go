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



func resourceArmServiceEndpointPolicyDefinition() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmServiceEndpointPolicyDefinitionCreateUpdate,
        Read: resourceArmServiceEndpointPolicyDefinitionRead,
        Update: resourceArmServiceEndpointPolicyDefinitionCreateUpdate,
        Delete: resourceArmServiceEndpointPolicyDefinitionDelete,

        Importer: &schema.ResourceImporter{
            State: schema.ImportStatePassthrough,
        },


        Schema: map[string]*schema.Schema{
            "name": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
            },

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "service_endpoint_policy_definition_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "service_endpoint_policy_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "description": {
                Type: schema.TypeString,
                Optional: true,
            },

            "etag": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
            },

            "service": {
                Type: schema.TypeString,
                Optional: true,
            },

            "service_resources": {
                Type: schema.TypeList,
                Optional: true,
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
            },

            "provisioning_state": {
                Type: schema.TypeString,
                Computed: true,
            },
        },
    }
}

func resourceArmServiceEndpointPolicyDefinitionCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).serviceEndpointPolicyDefinitionsClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    serviceEndpointPolicyDefinitionName := d.Get("service_endpoint_policy_definition_name").(string)
    serviceEndpointPolicyName := d.Get("service_endpoint_policy_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, serviceEndpointPolicyName, serviceEndpointPolicyDefinitionName)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Service Endpoint Policy Definition (Service Endpoint Policy Definition Name %q / Service Endpoint Policy Name %q / Resource Group %q): %+v", serviceEndpointPolicyDefinitionName, serviceEndpointPolicyName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_service_endpoint_policy_definition", *existing.ID)
        }
    }

    id := d.Get("id").(string)
    name := d.Get("name").(string)
    description := d.Get("description").(string)
    etag := d.Get("etag").(string)
    service := d.Get("service").(string)
    serviceResources := d.Get("service_resources").([]interface{})

    serviceEndpointPolicyDefinitions := network.ServiceEndpointPolicyDefinition{
        Etag: utils.String(etag),
        ID: utils.String(id),
        Name: utils.String(name),
        ServiceEndpointPolicyDefinitionPropertiesFormat: &network.ServiceEndpointPolicyDefinitionPropertiesFormat{
            Description: utils.String(description),
            Service: utils.String(service),
            ServiceResources: utils.ExpandStringSlice(serviceResources),
        },
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, serviceEndpointPolicyName, serviceEndpointPolicyDefinitionName, serviceEndpointPolicyDefinitions)
    if err != nil {
        return fmt.Errorf("Error creating Service Endpoint Policy Definition (Service Endpoint Policy Definition Name %q / Service Endpoint Policy Name %q / Resource Group %q): %+v", serviceEndpointPolicyDefinitionName, serviceEndpointPolicyName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Service Endpoint Policy Definition (Service Endpoint Policy Definition Name %q / Service Endpoint Policy Name %q / Resource Group %q): %+v", serviceEndpointPolicyDefinitionName, serviceEndpointPolicyName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, serviceEndpointPolicyName, serviceEndpointPolicyDefinitionName)
    if err != nil {
        return fmt.Errorf("Error retrieving Service Endpoint Policy Definition (Service Endpoint Policy Definition Name %q / Service Endpoint Policy Name %q / Resource Group %q): %+v", serviceEndpointPolicyDefinitionName, serviceEndpointPolicyName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Service Endpoint Policy Definition (Service Endpoint Policy Definition Name %q / Service Endpoint Policy Name %q / Resource Group %q) ID", serviceEndpointPolicyDefinitionName, serviceEndpointPolicyName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmServiceEndpointPolicyDefinitionRead(d, meta)
}

func resourceArmServiceEndpointPolicyDefinitionRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).serviceEndpointPolicyDefinitionsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    serviceEndpointPolicyName := id.Path["serviceEndpointPolicies"]
    serviceEndpointPolicyDefinitionName := id.Path["serviceEndpointPolicyDefinitions"]

    resp, err := client.Get(ctx, resourceGroup, serviceEndpointPolicyName, serviceEndpointPolicyDefinitionName)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Service Endpoint Policy Definition %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Service Endpoint Policy Definition (Service Endpoint Policy Definition Name %q / Service Endpoint Policy Name %q / Resource Group %q): %+v", serviceEndpointPolicyDefinitionName, serviceEndpointPolicyName, resourceGroup, err)
    }


    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if serviceEndpointPolicyDefinitionPropertiesFormat := resp.ServiceEndpointPolicyDefinitionPropertiesFormat; serviceEndpointPolicyDefinitionPropertiesFormat != nil {
        d.Set("description", serviceEndpointPolicyDefinitionPropertiesFormat.Description)
        d.Set("provisioning_state", serviceEndpointPolicyDefinitionPropertiesFormat.ProvisioningState)
        d.Set("service", serviceEndpointPolicyDefinitionPropertiesFormat.Service)
        d.Set("service_resources", utils.FlattenStringSlice(serviceEndpointPolicyDefinitionPropertiesFormat.ServiceResources))
    }
    d.Set("etag", resp.Etag)
    d.Set("service_endpoint_policy_definition_name", serviceEndpointPolicyDefinitionName)
    d.Set("service_endpoint_policy_name", serviceEndpointPolicyName)

    return nil
}


func resourceArmServiceEndpointPolicyDefinitionDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).serviceEndpointPolicyDefinitionsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    serviceEndpointPolicyName := id.Path["serviceEndpointPolicies"]
    serviceEndpointPolicyDefinitionName := id.Path["serviceEndpointPolicyDefinitions"]

    future, err := client.Delete(ctx, resourceGroup, serviceEndpointPolicyName, serviceEndpointPolicyDefinitionName)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Service Endpoint Policy Definition (Service Endpoint Policy Definition Name %q / Service Endpoint Policy Name %q / Resource Group %q): %+v", serviceEndpointPolicyDefinitionName, serviceEndpointPolicyName, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Service Endpoint Policy Definition (Service Endpoint Policy Definition Name %q / Service Endpoint Policy Name %q / Resource Group %q): %+v", serviceEndpointPolicyDefinitionName, serviceEndpointPolicyName, resourceGroup, err)
        }
    }

    return nil
}
