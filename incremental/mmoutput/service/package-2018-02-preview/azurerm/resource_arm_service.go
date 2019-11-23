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



func resourceArmService() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmServiceCreate,
        Read: resourceArmServiceRead,
        Update: resourceArmServiceUpdate,
        Delete: resourceArmServiceDelete,

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

            "admin_domain_name": {
                Type: schema.TypeString,
                Optional: true,
            },

            "etag": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
            },

            "notes": {
                Type: schema.TypeString,
                Optional: true,
            },

            "quantity": {
                Type: schema.TypeInt,
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

func resourceArmServiceCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).servicesClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Service %q (Resource Group %q): %+v", name, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_service", *existing.ID)
        }
    }

    name := d.Get("name").(string)
    location := azure.NormalizeLocation(d.Get("location").(string))
    adminDomainName := d.Get("admin_domain_name").(string)
    etag := d.Get("etag").(string)
    notes := d.Get("notes").(string)
    quantity := d.Get("quantity").(int)
    t := d.Get("tags").(map[string]interface{})

    deviceService := services.DeviceService{
        Etag: utils.String(etag),
        Location: utils.String(location),
        Name: utils.String(name),
        DeviceServiceProperties: &services.DeviceServiceProperties{
            AdminDomainName: utils.String(adminDomainName),
            Notes: utils.String(notes),
            Quantity: utils.Int64(int64(quantity)),
        },
        Tags: tags.Expand(t),
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroup, name, deviceService); err != nil {
        return fmt.Errorf("Error creating Service %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Service %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Service %q (Resource Group %q) ID", name, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmServiceRead(d, meta)
}

func resourceArmServiceRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).servicesClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["deviceServices"]

    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Service %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Service %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    d.Set("type", resp.Type)

    return nil
}

func resourceArmServiceUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).servicesClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    adminDomainName := d.Get("admin_domain_name").(string)
    etag := d.Get("etag").(string)
    notes := d.Get("notes").(string)
    quantity := d.Get("quantity").(int)
    t := d.Get("tags").(map[string]interface{})

    deviceService := services.DeviceService{
        Etag: utils.String(etag),
        Name: utils.String(name),
        DeviceServiceProperties: &services.DeviceServiceProperties{
            AdminDomainName: utils.String(adminDomainName),
            Notes: utils.String(notes),
            Quantity: utils.Int64(int64(quantity)),
        },
        Tags: tags.Expand(t),
    }


    if _, err := client.Update(ctx, resourceGroup, name, deviceService); err != nil {
        return fmt.Errorf("Error updating Service %q (Resource Group %q): %+v", name, resourceGroup, err)
    }

    return resourceArmServiceRead(d, meta)
}

func resourceArmServiceDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).servicesClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["deviceServices"]

    if _, err := client.Delete(ctx, resourceGroup, name); err != nil {
        return fmt.Errorf("Error deleting Service %q (Resource Group %q): %+v", name, resourceGroup, err)
    }

    return nil
}