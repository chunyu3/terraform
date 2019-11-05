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



func resourceArmProperty() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmPropertyCreate,
        Read: resourceArmPropertyRead,
        Update: resourceArmPropertyUpdate,
        Delete: resourceArmPropertyDelete,

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

            "prop_id": {
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

            "value": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "secret": {
                Type: schema.TypeBool,
                Optional: true,
                ForceNew: true,
            },

            "tags": {
                Type: schema.TypeList,
                Optional: true,
                ForceNew: true,
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
            },
        },
    }
}

func resourceArmPropertyCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).propertyClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    propID := d.Get("prop_id").(string)
    serviceName := d.Get("service_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, serviceName, propID)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Property (Prop %q / Service Name %q / Resource Group %q): %+v", propID, serviceName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_property", *existing.ID)
        }
    }

    name := d.Get("name").(string)
    secret := d.Get("secret").(bool)
    t := d.Get("tags").([]interface{})
    value := d.Get("value").(string)

    parameters := apimanagement.PropertyCreateParameters{
        Name: utils.String(name),
        Secret: utils.Bool(secret),
        Tags: utils.ExpandStringSlice(tags),
        Value: utils.String(value),
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroup, serviceName, propID, parameters); err != nil {
        return fmt.Errorf("Error creating Property (Prop %q / Service Name %q / Resource Group %q): %+v", propID, serviceName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, serviceName, propID)
    if err != nil {
        return fmt.Errorf("Error retrieving Property (Prop %q / Service Name %q / Resource Group %q): %+v", propID, serviceName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Property (Prop %q / Service Name %q / Resource Group %q) ID", propID, serviceName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmPropertyRead(d, meta)
}

func resourceArmPropertyRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).propertyClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    serviceName := id.Path["service"]
    propID := id.Path["properties"]

    resp, err := client.Get(ctx, resourceGroup, serviceName, propID)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Property %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Property (Prop %q / Service Name %q / Resource Group %q): %+v", propID, serviceName, resourceGroup, err)
    }


    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    d.Set("prop_id", propID)
    d.Set("secret", resp.Secret)
    d.Set("service_name", serviceName)
    d.Set("tags", utils.FlattenStringSlice(resp.Tags))
    d.Set("value", resp.Value)

    return tags.FlattenAndSet(d, resp.Tags)
}

func resourceArmPropertyUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).propertyClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    propID := d.Get("prop_id").(string)
    secret := d.Get("secret").(bool)
    serviceName := d.Get("service_name").(string)
    t := d.Get("tags").([]interface{})
    value := d.Get("value").(string)

    parameters := apimanagement.PropertyCreateParameters{
        Name: utils.String(name),
        Secret: utils.Bool(secret),
        Tags: utils.ExpandStringSlice(tags),
        Value: utils.String(value),
    }


    if _, err := client.Update(ctx, resourceGroup, serviceName, propID, parameters); err != nil {
        return fmt.Errorf("Error updating Property (Prop %q / Service Name %q / Resource Group %q): %+v", propID, serviceName, resourceGroup, err)
    }

    return resourceArmPropertyRead(d, meta)
}

func resourceArmPropertyDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).propertyClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    serviceName := id.Path["service"]
    propID := id.Path["properties"]

    if _, err := client.Delete(ctx, resourceGroup, serviceName, propID); err != nil {
        return fmt.Errorf("Error deleting Property (Prop %q / Service Name %q / Resource Group %q): %+v", propID, serviceName, resourceGroup, err)
    }

    return nil
}