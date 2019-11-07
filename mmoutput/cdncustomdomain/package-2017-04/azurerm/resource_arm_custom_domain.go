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



func resourceArmCustomDomain() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmCustomDomainCreateUpdate,
        Read: resourceArmCustomDomainRead,
        Update: resourceArmCustomDomainCreateUpdate,
        Delete: resourceArmCustomDomainDelete,

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

            "endpoint_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "host_name": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "profile_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "custom_https_provisioning_state": {
                Type: schema.TypeString,
                Computed: true,
            },

            "custom_https_provisioning_substate": {
                Type: schema.TypeString,
                Computed: true,
            },

            "provisioning_state": {
                Type: schema.TypeString,
                Computed: true,
            },

            "resource_state": {
                Type: schema.TypeString,
                Computed: true,
            },

            "type": {
                Type: schema.TypeString,
                Computed: true,
            },

            "validation_data": {
                Type: schema.TypeString,
                Computed: true,
            },
        },
    }
}

func resourceArmCustomDomainCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).customDomainsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    endpointName := d.Get("endpoint_name").(string)
    profileName := d.Get("profile_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, profileName, endpointName, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Custom Domain %q (Endpoint Name %q / Profile Name %q / Resource Group %q): %+v", name, endpointName, profileName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_custom_domain", *existing.ID)
        }
    }

    hostName := d.Get("host_name").(string)

    customDomainProperties := cdn.CustomDomainParameters{
        CustomDomainPropertiesParameters: &cdn.CustomDomainPropertiesParameters{
            HostName: utils.String(hostName),
        },
    }


    future, err := client.Create(ctx, resourceGroup, profileName, endpointName, name, customDomainProperties)
    if err != nil {
        return fmt.Errorf("Error creating Custom Domain %q (Endpoint Name %q / Profile Name %q / Resource Group %q): %+v", name, endpointName, profileName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Custom Domain %q (Endpoint Name %q / Profile Name %q / Resource Group %q): %+v", name, endpointName, profileName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, profileName, endpointName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Custom Domain %q (Endpoint Name %q / Profile Name %q / Resource Group %q): %+v", name, endpointName, profileName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Custom Domain %q (Endpoint Name %q / Profile Name %q / Resource Group %q) ID", name, endpointName, profileName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmCustomDomainRead(d, meta)
}

func resourceArmCustomDomainRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).customDomainsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    profileName := id.Path["profiles"]
    endpointName := id.Path["endpoints"]
    name := id.Path["customDomains"]

    resp, err := client.Get(ctx, resourceGroup, profileName, endpointName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Custom Domain %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Custom Domain %q (Endpoint Name %q / Profile Name %q / Resource Group %q): %+v", name, endpointName, profileName, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if customDomainPropertiesParameters := resp.CustomDomainPropertiesParameters; customDomainPropertiesParameters != nil {
        d.Set("custom_https_provisioning_state", string(customDomainPropertiesParameters.CustomHttpsProvisioningState))
        d.Set("custom_https_provisioning_substate", string(customDomainPropertiesParameters.CustomHttpsProvisioningSubstate))
        d.Set("host_name", customDomainPropertiesParameters.HostName)
        d.Set("provisioning_state", customDomainPropertiesParameters.ProvisioningState)
        d.Set("resource_state", string(customDomainPropertiesParameters.ResourceState))
        d.Set("validation_data", customDomainPropertiesParameters.ValidationData)
    }
    d.Set("endpoint_name", endpointName)
    d.Set("profile_name", profileName)
    d.Set("type", resp.Type)

    return nil
}


func resourceArmCustomDomainDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).customDomainsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    profileName := id.Path["profiles"]
    endpointName := id.Path["endpoints"]
    name := id.Path["customDomains"]

    future, err := client.Delete(ctx, resourceGroup, profileName, endpointName, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Custom Domain %q (Endpoint Name %q / Profile Name %q / Resource Group %q): %+v", name, endpointName, profileName, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Custom Domain %q (Endpoint Name %q / Profile Name %q / Resource Group %q): %+v", name, endpointName, profileName, resourceGroup, err)
        }
    }

    return nil
}
