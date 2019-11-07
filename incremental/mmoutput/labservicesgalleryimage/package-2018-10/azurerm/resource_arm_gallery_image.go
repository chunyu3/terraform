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



func resourceArmGalleryImage() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmGalleryImageCreate,
        Read: resourceArmGalleryImageRead,
        Update: resourceArmGalleryImageUpdate,
        Delete: resourceArmGalleryImageDelete,

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

            "lab_account_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "is_enabled": {
                Type: schema.TypeBool,
                Optional: true,
            },

            "is_override": {
                Type: schema.TypeBool,
                Optional: true,
            },

            "is_plan_authorized": {
                Type: schema.TypeBool,
                Optional: true,
            },

            "unique_identifier": {
                Type: schema.TypeString,
                Optional: true,
            },

            "author": {
                Type: schema.TypeString,
                Computed: true,
            },

            "created_date": {
                Type: schema.TypeString,
                Computed: true,
            },

            "description": {
                Type: schema.TypeString,
                Computed: true,
            },

            "icon": {
                Type: schema.TypeString,
                Computed: true,
            },

            "image_reference": {
                Type: schema.TypeList,
                Computed: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "offer": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "os_type": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "publisher": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "sku": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "version": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },

            "latest_operation_result": {
                Type: schema.TypeList,
                Computed: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "error_code": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "error_message": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "http_method": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "operation_url": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "request_uri": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "status": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },

            "plan_id": {
                Type: schema.TypeString,
                Computed: true,
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

func resourceArmGalleryImageCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).galleryImagesClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    labAccountName := d.Get("lab_account_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, labAccountName, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Gallery Image %q (Lab Account Name %q / Resource Group %q): %+v", name, labAccountName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_gallery_image", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    isEnabled := d.Get("is_enabled").(bool)
    isOverride := d.Get("is_override").(bool)
    isPlanAuthorized := d.Get("is_plan_authorized").(bool)
    uniqueIdentifier := d.Get("unique_identifier").(string)
    t := d.Get("tags").(map[string]interface{})

    galleryImage := labservices.GalleryImage{
        Location: utils.String(location),
        GalleryImageProperties: &labservices.GalleryImageProperties{
            IsEnabled: utils.Bool(isEnabled),
            IsOverride: utils.Bool(isOverride),
            IsPlanAuthorized: utils.Bool(isPlanAuthorized),
            UniqueIdentifier: utils.String(uniqueIdentifier),
        },
        Tags: tags.Expand(t),
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroup, labAccountName, name, galleryImage); err != nil {
        return fmt.Errorf("Error creating Gallery Image %q (Lab Account Name %q / Resource Group %q): %+v", name, labAccountName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, labAccountName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Gallery Image %q (Lab Account Name %q / Resource Group %q): %+v", name, labAccountName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Gallery Image %q (Lab Account Name %q / Resource Group %q) ID", name, labAccountName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmGalleryImageRead(d, meta)
}

func resourceArmGalleryImageRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).galleryImagesClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    labAccountName := id.Path["labaccounts"]
    name := id.Path["galleryimages"]

    resp, err := client.Get(ctx, resourceGroup, labAccountName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Gallery Image %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Gallery Image %q (Lab Account Name %q / Resource Group %q): %+v", name, labAccountName, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if galleryImageProperties := resp.GalleryImageProperties; galleryImageProperties != nil {
        d.Set("author", galleryImageProperties.Author)
        d.Set("created_date", (galleryImageProperties.CreatedDate).String())
        d.Set("description", galleryImageProperties.Description)
        d.Set("icon", galleryImageProperties.Icon)
        if err := d.Set("image_reference", flattenArmGalleryImageGalleryImageReference(galleryImageProperties.ImageReference)); err != nil {
            return fmt.Errorf("Error setting `image_reference`: %+v", err)
        }
        d.Set("is_enabled", galleryImageProperties.IsEnabled)
        d.Set("is_override", galleryImageProperties.IsOverride)
        d.Set("is_plan_authorized", galleryImageProperties.IsPlanAuthorized)
        if err := d.Set("latest_operation_result", flattenArmGalleryImageLatestOperationResult(galleryImageProperties.LatestOperationResult)); err != nil {
            return fmt.Errorf("Error setting `latest_operation_result`: %+v", err)
        }
        d.Set("plan_id", galleryImageProperties.PlanID)
        d.Set("provisioning_state", galleryImageProperties.ProvisioningState)
        d.Set("unique_identifier", galleryImageProperties.UniqueIdentifier)
    }
    d.Set("lab_account_name", labAccountName)
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}

func resourceArmGalleryImageUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).galleryImagesClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    isEnabled := d.Get("is_enabled").(bool)
    isOverride := d.Get("is_override").(bool)
    isPlanAuthorized := d.Get("is_plan_authorized").(bool)
    labAccountName := d.Get("lab_account_name").(string)
    uniqueIdentifier := d.Get("unique_identifier").(string)
    t := d.Get("tags").(map[string]interface{})

    galleryImage := labservices.GalleryImage{
        Location: utils.String(location),
        GalleryImageProperties: &labservices.GalleryImageProperties{
            IsEnabled: utils.Bool(isEnabled),
            IsOverride: utils.Bool(isOverride),
            IsPlanAuthorized: utils.Bool(isPlanAuthorized),
            UniqueIdentifier: utils.String(uniqueIdentifier),
        },
        Tags: tags.Expand(t),
    }


    if _, err := client.Update(ctx, resourceGroup, labAccountName, name, galleryImage); err != nil {
        return fmt.Errorf("Error updating Gallery Image %q (Lab Account Name %q / Resource Group %q): %+v", name, labAccountName, resourceGroup, err)
    }

    return resourceArmGalleryImageRead(d, meta)
}

func resourceArmGalleryImageDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).galleryImagesClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    labAccountName := id.Path["labaccounts"]
    name := id.Path["galleryimages"]

    if _, err := client.Delete(ctx, resourceGroup, labAccountName, name); err != nil {
        return fmt.Errorf("Error deleting Gallery Image %q (Lab Account Name %q / Resource Group %q): %+v", name, labAccountName, resourceGroup, err)
    }

    return nil
}


func flattenArmGalleryImageGalleryImageReference(input *labservices.GalleryImageReference) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})


    return []interface{}{result}
}

func flattenArmGalleryImageLatestOperationResult(input *labservices.LatestOperationResult) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})


    return []interface{}{result}
}