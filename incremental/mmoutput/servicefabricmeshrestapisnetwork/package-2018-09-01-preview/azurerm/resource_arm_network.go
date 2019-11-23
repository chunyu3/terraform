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



func resourceArmNetwork() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmNetworkCreateUpdate,
        Read: resourceArmNetworkRead,
        Update: resourceArmNetworkCreateUpdate,
        Delete: resourceArmNetworkDelete,

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

            "description": {
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

func resourceArmNetworkCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).networkClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Network %q (Resource Group %q): %+v", name, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_network", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    description := d.Get("description").(string)
    t := d.Get("tags").(map[string]interface{})

    networkResourceDescription := servicefabricmeshrestapis.NetworkResourceDescription{
        Location: utils.String(location),
        NetworkResourceProperties: &servicefabricmeshrestapis.NetworkResourceProperties{
            Description: utils.String(description),
        },
        Tags: tags.Expand(t),
    }


    if _, err := client.Create(ctx, resourceGroup, name, networkResourceDescription); err != nil {
        return fmt.Errorf("Error creating Network %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Network %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Network %q (Resource Group %q) ID", name, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmNetworkRead(d, meta)
}

func resourceArmNetworkRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).networkClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["networks"]

    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Network %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Network %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    d.Set("type", resp.Type)

    return nil
}


func resourceArmNetworkDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).networkClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["networks"]

    if _, err := client.Delete(ctx, resourceGroup, name); err != nil {
        return fmt.Errorf("Error deleting Network %q (Resource Group %q): %+v", name, resourceGroup, err)
    }

    return nil
}