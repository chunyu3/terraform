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



func resourceArmDedicatedCloudService() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmDedicatedCloudServiceCreate,
        Read: resourceArmDedicatedCloudServiceRead,
        Update: resourceArmDedicatedCloudServiceUpdate,
        Delete: resourceArmDedicatedCloudServiceDelete,

        Importer: &schema.ResourceImporter{
            State: schema.ImportStatePassthrough,
        },


        Schema: map[string]*schema.Schema{
            "name": {
                Type: schema.TypeString,
                Computed: true,
            },

            "location": azure.SchemaLocation(),

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "dedicated_cloud_service_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "gateway_subnet": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "is_account_onboarded": {
                Type: schema.TypeString,
                Computed: true,
            },

            "nodes": {
                Type: schema.TypeInt,
                Computed: true,
            },

            "service_url": {
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

func resourceArmDedicatedCloudServiceCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).dedicatedCloudServicesClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    dedicatedCloudServiceName := d.Get("dedicated_cloud_service_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, dedicatedCloudServiceName)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Dedicated Cloud Service (Dedicated Cloud Service Name %q / Resource Group %q): %+v", dedicatedCloudServiceName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_dedicated_cloud_service", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    gatewaySubnet := d.Get("gateway_subnet").(string)
    t := d.Get("tags").(map[string]interface{})

    dedicatedCloudServiceRequest := vmwarecloudsimple.DedicatedCloudService{
        Location: utils.String(location),
        DedicatedCloudServiceProperties: &vmwarecloudsimple.DedicatedCloudServiceProperties{
            GatewaySubnet: utils.String(gatewaySubnet),
        },
        Tags: tags.Expand(t),
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroup, dedicatedCloudServiceName, dedicatedCloudServiceRequest); err != nil {
        return fmt.Errorf("Error creating Dedicated Cloud Service (Dedicated Cloud Service Name %q / Resource Group %q): %+v", dedicatedCloudServiceName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, dedicatedCloudServiceName)
    if err != nil {
        return fmt.Errorf("Error retrieving Dedicated Cloud Service (Dedicated Cloud Service Name %q / Resource Group %q): %+v", dedicatedCloudServiceName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Dedicated Cloud Service (Dedicated Cloud Service Name %q / Resource Group %q) ID", dedicatedCloudServiceName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmDedicatedCloudServiceRead(d, meta)
}

func resourceArmDedicatedCloudServiceRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).dedicatedCloudServicesClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    dedicatedCloudServiceName := id.Path["dedicatedCloudServices"]

    resp, err := client.Get(ctx, resourceGroup, dedicatedCloudServiceName)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Dedicated Cloud Service %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Dedicated Cloud Service (Dedicated Cloud Service Name %q / Resource Group %q): %+v", dedicatedCloudServiceName, resourceGroup, err)
    }


    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    d.Set("dedicated_cloud_service_name", dedicatedCloudServiceName)
    if dedicatedCloudServiceProperties := resp.DedicatedCloudServiceProperties; dedicatedCloudServiceProperties != nil {
        d.Set("gateway_subnet", dedicatedCloudServiceProperties.GatewaySubnet)
        d.Set("is_account_onboarded", string(dedicatedCloudServiceProperties.IsAccountOnboarded))
        d.Set("nodes", dedicatedCloudServiceProperties.Nodes)
        d.Set("service_url", dedicatedCloudServiceProperties.ServiceURL)
    }
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}

func resourceArmDedicatedCloudServiceUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).dedicatedCloudServicesClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    dedicatedCloudServiceName := d.Get("dedicated_cloud_service_name").(string)
    gatewaySubnet := d.Get("gateway_subnet").(string)
    t := d.Get("tags").(map[string]interface{})

    dedicatedCloudServiceRequest := vmwarecloudsimple.DedicatedCloudService{
        Location: utils.String(location),
        DedicatedCloudServiceProperties: &vmwarecloudsimple.DedicatedCloudServiceProperties{
            GatewaySubnet: utils.String(gatewaySubnet),
        },
        Tags: tags.Expand(t),
    }


    if _, err := client.Update(ctx, resourceGroup, dedicatedCloudServiceName, dedicatedCloudServiceRequest); err != nil {
        return fmt.Errorf("Error updating Dedicated Cloud Service (Dedicated Cloud Service Name %q / Resource Group %q): %+v", dedicatedCloudServiceName, resourceGroup, err)
    }

    return resourceArmDedicatedCloudServiceRead(d, meta)
}

func resourceArmDedicatedCloudServiceDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).dedicatedCloudServicesClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    dedicatedCloudServiceName := id.Path["dedicatedCloudServices"]

    future, err := client.Delete(ctx, resourceGroup, dedicatedCloudServiceName)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Dedicated Cloud Service (Dedicated Cloud Service Name %q / Resource Group %q): %+v", dedicatedCloudServiceName, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Dedicated Cloud Service (Dedicated Cloud Service Name %q / Resource Group %q): %+v", dedicatedCloudServiceName, resourceGroup, err)
        }
    }

    return nil
}
