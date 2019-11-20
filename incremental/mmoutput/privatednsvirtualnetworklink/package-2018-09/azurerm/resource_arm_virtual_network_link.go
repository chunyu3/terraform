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



func resourceArmVirtualNetworkLink() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmVirtualNetworkLinkCreate,
        Read: resourceArmVirtualNetworkLinkRead,
        Update: resourceArmVirtualNetworkLinkUpdate,
        Delete: resourceArmVirtualNetworkLinkDelete,

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

            "private_zone_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "etag": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
            },

            "registration_enabled": {
                Type: schema.TypeBool,
                Optional: true,
            },

            "virtual_network": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "id": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },

            "type": {
                Type: schema.TypeString,
                Computed: true,
            },

            "tags": tags.Schema(),
        },
    }
}

func resourceArmVirtualNetworkLinkCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).virtualNetworkLinksClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    privateZoneName := d.Get("private_zone_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, privateZoneName, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Virtual Network Link %q (Private Zone Name %q / Resource Group %q): %+v", name, privateZoneName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_virtual_network_link", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    etag := d.Get("etag").(string)
    registrationEnabled := d.Get("registration_enabled").(bool)
    virtualNetwork := d.Get("virtual_network").([]interface{})
    t := d.Get("tags").(map[string]interface{})

    parameters := privatedns.VirtualNetworkLink{
        Etag: utils.String(etag),
        Location: utils.String(location),
        VirtualNetworkLinkProperties: &privatedns.VirtualNetworkLinkProperties{
            RegistrationEnabled: utils.Bool(registrationEnabled),
            VirtualNetwork: expandArmVirtualNetworkLinkSubResource(virtualNetwork),
        },
        Tags: tags.Expand(t),
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, privateZoneName, name, parameters)
    if err != nil {
        return fmt.Errorf("Error creating Virtual Network Link %q (Private Zone Name %q / Resource Group %q): %+v", name, privateZoneName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Virtual Network Link %q (Private Zone Name %q / Resource Group %q): %+v", name, privateZoneName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, privateZoneName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Virtual Network Link %q (Private Zone Name %q / Resource Group %q): %+v", name, privateZoneName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Virtual Network Link %q (Private Zone Name %q / Resource Group %q) ID", name, privateZoneName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmVirtualNetworkLinkRead(d, meta)
}

func resourceArmVirtualNetworkLinkRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).virtualNetworkLinksClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    privateZoneName := id.Path["privateDnsZones"]
    name := id.Path["virtualNetworkLinks"]

    resp, err := client.Get(ctx, resourceGroup, privateZoneName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Virtual Network Link %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Virtual Network Link %q (Private Zone Name %q / Resource Group %q): %+v", name, privateZoneName, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    d.Set("private_zone_name", privateZoneName)
    d.Set("type", resp.Type)

    return nil
}

func resourceArmVirtualNetworkLinkUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).virtualNetworkLinksClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    etag := d.Get("etag").(string)
    privateZoneName := d.Get("private_zone_name").(string)
    registrationEnabled := d.Get("registration_enabled").(bool)
    virtualNetwork := d.Get("virtual_network").([]interface{})
    t := d.Get("tags").(map[string]interface{})

    parameters := privatedns.VirtualNetworkLink{
        Etag: utils.String(etag),
        VirtualNetworkLinkProperties: &privatedns.VirtualNetworkLinkProperties{
            RegistrationEnabled: utils.Bool(registrationEnabled),
            VirtualNetwork: expandArmVirtualNetworkLinkSubResource(virtualNetwork),
        },
        Tags: tags.Expand(t),
    }


    future, err := client.Update(ctx, resourceGroup, privateZoneName, name, parameters)
    if err != nil {
        return fmt.Errorf("Error updating Virtual Network Link %q (Private Zone Name %q / Resource Group %q): %+v", name, privateZoneName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for update of Virtual Network Link %q (Private Zone Name %q / Resource Group %q): %+v", name, privateZoneName, resourceGroup, err)
    }

    return resourceArmVirtualNetworkLinkRead(d, meta)
}

func resourceArmVirtualNetworkLinkDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).virtualNetworkLinksClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    privateZoneName := id.Path["privateDnsZones"]
    name := id.Path["virtualNetworkLinks"]

    future, err := client.Delete(ctx, resourceGroup, privateZoneName, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Virtual Network Link %q (Private Zone Name %q / Resource Group %q): %+v", name, privateZoneName, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Virtual Network Link %q (Private Zone Name %q / Resource Group %q): %+v", name, privateZoneName, resourceGroup, err)
        }
    }

    return nil
}

func expandArmVirtualNetworkLinkSubResource(input []interface{}) *privatedns.SubResource {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    id := v["id"].(string)

    result := privatedns.SubResource{
        ID: utils.String(id),
    }
    return &result
}
