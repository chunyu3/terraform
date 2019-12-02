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

            "identity": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "type": {
                            Type: schema.TypeString,
                            Optional: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(sql.SystemAssigned),
                            }, false),
                            Default: string(sql.SystemAssigned),
                        },
                    },
                },
            },

            "version": {
                Type: schema.TypeString,
                Optional: true,
            },

            "fully_qualified_domain_name": {
                Type: schema.TypeString,
                Computed: true,
            },

            "kind": {
                Type: schema.TypeString,
                Computed: true,
            },

            "state": {
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
    identity := d.Get("identity").([]interface{})
    type := d.Get("type").(string)
    version := d.Get("version").(string)
    t := d.Get("tags").(map[string]interface{})

    parameters := sql.ServerUpdate{
        Identity: expandArmServerResourceIdentity(identity),
        Location: utils.String(location),
        Name: utils.String(name),
        ServerProperties: &sql.ServerProperties{
            AdministratorLogin: utils.String(administratorLogin),
            AdministratorLoginPassword: utils.String(administratorLoginPassword),
            Version: utils.String(version),
        },
        Tags: tags.Expand(t),
        Type: utils.String(type),
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, name, parameters)
    if err != nil {
        return fmt.Errorf("Error creating Server %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Server %q (Resource Group %q): %+v", name, resourceGroup, err)
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
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if serverProperties := resp.ServerProperties; serverProperties != nil {
        d.Set("administrator_login", serverProperties.AdministratorLogin)
        d.Set("administrator_login_password", serverProperties.AdministratorLoginPassword)
        d.Set("fully_qualified_domain_name", serverProperties.FullyQualifiedDomainName)
        d.Set("state", serverProperties.State)
        d.Set("version", serverProperties.Version)
    }
    if err := d.Set("identity", flattenArmServerResourceIdentity(resp.Identity)); err != nil {
        return fmt.Errorf("Error setting `identity`: %+v", err)
    }
    d.Set("kind", resp.Kind)
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}

func resourceArmServerUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).serversClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    administratorLogin := d.Get("administrator_login").(string)
    administratorLoginPassword := d.Get("administrator_login_password").(string)
    identity := d.Get("identity").([]interface{})
    type := d.Get("type").(string)
    version := d.Get("version").(string)
    t := d.Get("tags").(map[string]interface{})

    parameters := sql.ServerUpdate{
        Identity: expandArmServerResourceIdentity(identity),
        Name: utils.String(name),
        ServerProperties: &sql.ServerProperties{
            AdministratorLogin: utils.String(administratorLogin),
            AdministratorLoginPassword: utils.String(administratorLoginPassword),
            Version: utils.String(version),
        },
        Tags: tags.Expand(t),
        Type: utils.String(type),
    }


    future, err := client.Update(ctx, resourceGroup, name, parameters)
    if err != nil {
        return fmt.Errorf("Error updating Server %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for update of Server %q (Resource Group %q): %+v", name, resourceGroup, err)
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

    future, err := client.Delete(ctx, resourceGroup, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Server %q (Resource Group %q): %+v", name, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Server %q (Resource Group %q): %+v", name, resourceGroup, err)
        }
    }

    return nil
}

func expandArmServerResourceIdentity(input []interface{}) *sql.ResourceIdentity {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    type := v["type"].(string)

    result := sql.ResourceIdentity{
        Type: sql.IdentityType(type),
    }
    return &result
}


func flattenArmServerResourceIdentity(input *sql.ResourceIdentity) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    result["type"] = string(input.Type)

    return []interface{}{result}
}
