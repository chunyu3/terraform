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



func resourceArmOpenIdConnectProvider() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmOpenIdConnectProviderCreate,
        Read: resourceArmOpenIdConnectProviderRead,
        Update: resourceArmOpenIdConnectProviderUpdate,
        Delete: resourceArmOpenIdConnectProviderDelete,

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

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "client_id": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "metadata_endpoint": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "opid": {
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

            "client_secret": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
            },

            "description": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
            },
        },
    }
}

func resourceArmOpenIdConnectProviderCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).openIdConnectProvidersClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    opid := d.Get("opid").(string)
    serviceName := d.Get("service_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, serviceName, opid)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Open Id Connect Provider (Opid %q / Service Name %q / Resource Group %q): %+v", opid, serviceName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_open_id_connect_provider", *existing.ID)
        }
    }

    name := d.Get("name").(string)
    clientId := d.Get("client_id").(string)
    clientSecret := d.Get("client_secret").(string)
    description := d.Get("description").(string)
    metadataEndpoint := d.Get("metadata_endpoint").(string)

    parameters := apimanagement.OpenidConnectProviderCreateContract{
        ClientID: utils.String(clientId),
        ClientSecret: utils.String(clientSecret),
        Description: utils.String(description),
        MetadataEndpoint: utils.String(metadataEndpoint),
        Name: utils.String(name),
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroup, serviceName, opid, parameters); err != nil {
        return fmt.Errorf("Error creating Open Id Connect Provider (Opid %q / Service Name %q / Resource Group %q): %+v", opid, serviceName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, serviceName, opid)
    if err != nil {
        return fmt.Errorf("Error retrieving Open Id Connect Provider (Opid %q / Service Name %q / Resource Group %q): %+v", opid, serviceName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Open Id Connect Provider (Opid %q / Service Name %q / Resource Group %q) ID", opid, serviceName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmOpenIdConnectProviderRead(d, meta)
}

func resourceArmOpenIdConnectProviderRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).openIdConnectProvidersClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    serviceName := id.Path["service"]
    opid := id.Path["openidConnectProviders"]

    resp, err := client.Get(ctx, resourceGroup, serviceName, opid)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Open Id Connect Provider %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Open Id Connect Provider (Opid %q / Service Name %q / Resource Group %q): %+v", opid, serviceName, resourceGroup, err)
    }


    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    d.Set("client_id", resp.ClientID)
    d.Set("client_secret", resp.ClientSecret)
    d.Set("description", resp.Description)
    d.Set("metadata_endpoint", resp.MetadataEndpoint)
    d.Set("opid", opid)
    d.Set("service_name", serviceName)

    return nil
}

func resourceArmOpenIdConnectProviderUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).openIdConnectProvidersClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    clientId := d.Get("client_id").(string)
    clientSecret := d.Get("client_secret").(string)
    description := d.Get("description").(string)
    metadataEndpoint := d.Get("metadata_endpoint").(string)
    opid := d.Get("opid").(string)
    serviceName := d.Get("service_name").(string)

    parameters := apimanagement.OpenidConnectProviderCreateContract{
        ClientID: utils.String(clientId),
        ClientSecret: utils.String(clientSecret),
        Description: utils.String(description),
        MetadataEndpoint: utils.String(metadataEndpoint),
        Name: utils.String(name),
    }


    if _, err := client.Update(ctx, resourceGroup, serviceName, opid, parameters); err != nil {
        return fmt.Errorf("Error updating Open Id Connect Provider (Opid %q / Service Name %q / Resource Group %q): %+v", opid, serviceName, resourceGroup, err)
    }

    return resourceArmOpenIdConnectProviderRead(d, meta)
}

func resourceArmOpenIdConnectProviderDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).openIdConnectProvidersClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    serviceName := id.Path["service"]
    opid := id.Path["openidConnectProviders"]

    if _, err := client.Delete(ctx, resourceGroup, serviceName, opid); err != nil {
        return fmt.Errorf("Error deleting Open Id Connect Provider (Opid %q / Service Name %q / Resource Group %q): %+v", opid, serviceName, resourceGroup, err)
    }

    return nil
}
