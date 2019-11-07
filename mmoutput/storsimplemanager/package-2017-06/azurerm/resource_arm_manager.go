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



func resourceArmManager() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmManagerCreate,
        Read: resourceArmManagerRead,
        Update: resourceArmManagerUpdate,
        Delete: resourceArmManagerDelete,

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

            "cis_intrinsic_settings": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "type": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(storsimple.GardaV1),
                                string(storsimple.HelsinkiV1),
                            }, false),
                        },
                    },
                },
            },

            "etag": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
            },

            "sku": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "name": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                    },
                },
            },

            "provisioning_state": {
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

func resourceArmManagerCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).managersClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Manager %q (Resource Group %q): %+v", name, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_manager", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    cisIntrinsicSettings := d.Get("cis_intrinsic_settings").([]interface{})
    etag := d.Get("etag").(string)
    sku := d.Get("sku").([]interface{})
    t := d.Get("tags").(map[string]interface{})

    parameters := storsimple.Manager{
        Etag: utils.String(etag),
        Location: utils.String(location),
        ManagerProperties: &storsimple.ManagerProperties{
            CisIntrinsicSettings: expandArmManagerManagerIntrinsicSettings(cisIntrinsicSettings),
            Sku: expandArmManagerManagerSku(sku),
        },
        Tags: tags.Expand(t),
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroup, name, parameters); err != nil {
        return fmt.Errorf("Error creating Manager %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Manager %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Manager %q (Resource Group %q) ID", name, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmManagerRead(d, meta)
}

func resourceArmManagerRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).managersClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["managers"]

    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Manager %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Manager %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if managerProperties := resp.ManagerProperties; managerProperties != nil {
        if err := d.Set("cis_intrinsic_settings", flattenArmManagerManagerIntrinsicSettings(managerProperties.CisIntrinsicSettings)); err != nil {
            return fmt.Errorf("Error setting `cis_intrinsic_settings`: %+v", err)
        }
        d.Set("provisioning_state", managerProperties.ProvisioningState)
        if err := d.Set("sku", flattenArmManagerManagerSku(managerProperties.Sku)); err != nil {
            return fmt.Errorf("Error setting `sku`: %+v", err)
        }
    }
    d.Set("etag", resp.Etag)
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}

func resourceArmManagerUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).managersClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    cisIntrinsicSettings := d.Get("cis_intrinsic_settings").([]interface{})
    etag := d.Get("etag").(string)
    sku := d.Get("sku").([]interface{})
    t := d.Get("tags").(map[string]interface{})

    parameters := storsimple.Manager{
        Etag: utils.String(etag),
        Location: utils.String(location),
        ManagerProperties: &storsimple.ManagerProperties{
            CisIntrinsicSettings: expandArmManagerManagerIntrinsicSettings(cisIntrinsicSettings),
            Sku: expandArmManagerManagerSku(sku),
        },
        Tags: tags.Expand(t),
    }


    if _, err := client.Update(ctx, resourceGroup, name, parameters); err != nil {
        return fmt.Errorf("Error updating Manager %q (Resource Group %q): %+v", name, resourceGroup, err)
    }

    return resourceArmManagerRead(d, meta)
}

func resourceArmManagerDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).managersClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["managers"]

    if _, err := client.Delete(ctx, resourceGroup, name); err != nil {
        return fmt.Errorf("Error deleting Manager %q (Resource Group %q): %+v", name, resourceGroup, err)
    }

    return nil
}

func expandArmManagerManagerIntrinsicSettings(input []interface{}) *storsimple.ManagerIntrinsicSettings {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    type := v["type"].(string)

    result := storsimple.ManagerIntrinsicSettings{
        Type: storsimple.ManagerType(type),
    }
    return &result
}

func expandArmManagerManagerSku(input []interface{}) *storsimple.ManagerSku {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    name := v["name"].(string)

    result := storsimple.ManagerSku{
        Name: utils.String(name),
    }
    return &result
}


func flattenArmManagerManagerIntrinsicSettings(input *storsimple.ManagerIntrinsicSettings) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    result["type"] = string(input.Type)

    return []interface{}{result}
}

func flattenArmManagerManagerSku(input *storsimple.ManagerSku) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if name := input.Name; name != nil {
        result["name"] = *name
    }

    return []interface{}{result}
}
