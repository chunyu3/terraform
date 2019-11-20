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



func resourceArmServer() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmServerCreate,
        Read: resourceArmServerRead,
        Update: resourceArmServerUpdate,
        Delete: resourceArmServerDelete,

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

            "type": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "administrator_login": {
                Type: schema.TypeString,
                Optional: true,
            },

            "administrator_login_password": {
                Type: schema.TypeString,
                Optional: true,
            },

            "version": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(sql.2.0),
                    string(sql.12.0),
                }, false),
                Default: string(sql.2.0),
            },

            "kind": {
                Type: schema.TypeString,
                Computed: true,
            },

            "tags": tags.Schema(),
        },
    }
}

func resourceArmServerCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).serversClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Server %q (Resource Group %q): %+v", name, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_server", *existing.ID)
        }
    }

    name := d.Get("name").(string)
    location := azure.NormalizeLocation(d.Get("location").(string))
    administratorLogin := d.Get("administrator_login").(string)
    administratorLoginPassword := d.Get("administrator_login_password").(string)
    type := d.Get("type").(string)
    version := d.Get("version").(string)
    t := d.Get("tags").(map[string]interface{})

    parameters := sql.ServerUpdate{
        Location: utils.String(location),
        Name: utils.String(name),
        ServerProperties: &sql.ServerProperties{
            AdministratorLogin: utils.String(administratorLogin),
            AdministratorLoginPassword: utils.String(administratorLoginPassword),
            Version: sql.ServerVersion(version),
        },
        Tags: tags.Expand(t),
        Type: utils.String(type),
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroup, name, parameters); err != nil {
        return fmt.Errorf("Error creating Server %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Server %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Server %q (Resource Group %q) ID", name, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmServerRead(d, meta)
}

func resourceArmServerRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).serversClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["servers"]

    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Server %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Server %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    d.Set("name", resp.Name)
    d.Set("name", name)
    d.Set("resource_group", resourceGroup)
    d.Set("kind", resp.Kind)
    d.Set("type", resp.Type)

    return nil
}

func resourceArmServerUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).serversClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    administratorLogin := d.Get("administrator_login").(string)
    administratorLoginPassword := d.Get("administrator_login_password").(string)
    type := d.Get("type").(string)
    version := d.Get("version").(string)
    t := d.Get("tags").(map[string]interface{})

    parameters := sql.ServerUpdate{
        Name: utils.String(name),
        ServerProperties: &sql.ServerProperties{
            AdministratorLogin: utils.String(administratorLogin),
            AdministratorLoginPassword: utils.String(administratorLoginPassword),
            Version: sql.ServerVersion(version),
        },
        Tags: tags.Expand(t),
        Type: utils.String(type),
    }


    if _, err := client.Update(ctx, resourceGroup, name, parameters); err != nil {
        return fmt.Errorf("Error updating Server %q (Resource Group %q): %+v", name, resourceGroup, err)
    }

    return resourceArmServerRead(d, meta)
}

func resourceArmServerDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).serversClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["servers"]

    if _, err := client.Delete(ctx, resourceGroup, name); err != nil {
        return fmt.Errorf("Error deleting Server %q (Resource Group %q): %+v", name, resourceGroup, err)
    }

    return nil
}
