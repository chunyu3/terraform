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
        Create: resourceArmGalleryImageCreateUpdate,
        Read: resourceArmGalleryImageRead,
        Update: resourceArmGalleryImageCreateUpdate,
        Delete: resourceArmGalleryImageDelete,

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

            "gallery_image_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "gallery_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "identifier": {
                Type: schema.TypeList,
                Required: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "offer": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "publisher": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "sku": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                    },
                },
            },

            "os_state": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(compute.Generalized),
                    string(compute.Specialized),
                }, false),
            },

            "os_type": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(compute.Windows),
                    string(compute.Linux),
                }, false),
            },

            "description": {
                Type: schema.TypeString,
                Optional: true,
            },

            "disallowed": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "disk_types": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                    },
                },
            },

            "end_of_life_date": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validateRFC3339Date,
            },

            "eula": {
                Type: schema.TypeString,
                Optional: true,
            },

            "privacy_statement_uri": {
                Type: schema.TypeString,
                Optional: true,
            },

            "purchase_plan": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "name": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "product": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "publisher": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },

            "recommended": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "memory": {
                            Type: schema.TypeList,
                            Optional: true,
                            MaxItems: 1,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "max": {
                                        Type: schema.TypeInt,
                                        Optional: true,
                                    },
                                    "min": {
                                        Type: schema.TypeInt,
                                        Optional: true,
                                    },
                                },
                            },
                        },
                        "v_cpus": {
                            Type: schema.TypeList,
                            Optional: true,
                            MaxItems: 1,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "max": {
                                        Type: schema.TypeInt,
                                        Optional: true,
                                    },
                                    "min": {
                                        Type: schema.TypeInt,
                                        Optional: true,
                                    },
                                },
                            },
                        },
                    },
                },
            },

            "release_note_uri": {
                Type: schema.TypeString,
                Optional: true,
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

func resourceArmGalleryImageCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).galleryImagesClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    galleryImageName := d.Get("gallery_image_name").(string)
    galleryName := d.Get("gallery_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        resp, err := client.Get(ctx, resourceGroup, galleryName, galleryImageName)
        if err != nil {
            if !utils.ResponseWasNotFound(resp.Response) {
                return fmt.Errorf("Error checking for present of existing Gallery Image (Gallery Image Name %q / Gallery Name %q / Resource Group %q): %+v", galleryImageName, galleryName, resourceGroup, err)
            }
        }
        if !utils.ResponseWasNotFound(resp.Response) {
            return tf.ImportAsExistsError("azurerm_gallery_image", *resp.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    description := d.Get("description").(string)
    disallowed := d.Get("disallowed").([]interface{})
    endOfLifeDate := d.Get("end_of_life_date").(string)
    eula := d.Get("eula").(string)
    identifier := d.Get("identifier").([]interface{})
    osState := d.Get("os_state").(string)
    osType := d.Get("os_type").(string)
    privacyStatementUri := d.Get("privacy_statement_uri").(string)
    purchasePlan := d.Get("purchase_plan").([]interface{})
    recommended := d.Get("recommended").([]interface{})
    releaseNoteUri := d.Get("release_note_uri").(string)
    t := d.Get("tags").(map[string]interface{})

    galleryImage := compute.GalleryImage{
        Location: utils.String(location),
        GalleryImageProperties: &compute.GalleryImageProperties{
            Description: utils.String(description),
            Disallowed: expandArmGalleryImageDisallowed(disallowed),
            EndOfLifeDate: convertStringToDate(endOfLifeDate),
            Eula: utils.String(eula),
            Identifier: expandArmGalleryImageGalleryImageIdentifier(identifier),
            OsState: compute.OperatingSystemStateTypes(osState),
            OsType: compute.OperatingSystemTypes(osType),
            PrivacyStatementUri: utils.String(privacyStatementUri),
            PurchasePlan: expandArmGalleryImageImagePurchasePlan(purchasePlan),
            Recommended: expandArmGalleryImageRecommendedMachineConfiguration(recommended),
            ReleaseNoteUri: utils.String(releaseNoteUri),
        },
        Tags: tags.Expand(t),
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, galleryName, galleryImageName, galleryImage)
    if err != nil {
        return fmt.Errorf("Error creating Gallery Image (Gallery Image Name %q / Gallery Name %q / Resource Group %q): %+v", galleryImageName, galleryName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Gallery Image (Gallery Image Name %q / Gallery Name %q / Resource Group %q): %+v", galleryImageName, galleryName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, galleryName, galleryImageName)
    if err != nil {
        return fmt.Errorf("Error retrieving Gallery Image (Gallery Image Name %q / Gallery Name %q / Resource Group %q): %+v", galleryImageName, galleryName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Gallery Image (Gallery Image Name %q / Gallery Name %q / Resource Group %q) ID", galleryImageName, galleryName, resourceGroup)
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
    galleryName := id.Path["galleries"]
    galleryImageName := id.Path["images"]

    resp, err := client.Get(ctx, resourceGroup, galleryName, galleryImageName)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Gallery Image %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Gallery Image (Gallery Image Name %q / Gallery Name %q / Resource Group %q): %+v", galleryImageName, galleryName, resourceGroup, err)
    }


    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if galleryImageProperties := resp.GalleryImageProperties; galleryImageProperties != nil {
        d.Set("description", galleryImageProperties.Description)
        if err := d.Set("disallowed", flattenArmGalleryImageDisallowed(galleryImageProperties.Disallowed)); err != nil {
            return fmt.Errorf("Error setting `disallowed`: %+v", err)
        }
        d.Set("end_of_life_date", (galleryImageProperties.EndOfLifeDate).String())
        d.Set("eula", galleryImageProperties.Eula)
        if err := d.Set("identifier", flattenArmGalleryImageGalleryImageIdentifier(galleryImageProperties.Identifier)); err != nil {
            return fmt.Errorf("Error setting `identifier`: %+v", err)
        }
        d.Set("os_state", string(galleryImageProperties.OsState))
        d.Set("os_type", string(galleryImageProperties.OsType))
        d.Set("privacy_statement_uri", galleryImageProperties.PrivacyStatementUri)
        d.Set("provisioning_state", string(galleryImageProperties.ProvisioningState))
        if err := d.Set("purchase_plan", flattenArmGalleryImageImagePurchasePlan(galleryImageProperties.PurchasePlan)); err != nil {
            return fmt.Errorf("Error setting `purchase_plan`: %+v", err)
        }
        if err := d.Set("recommended", flattenArmGalleryImageRecommendedMachineConfiguration(galleryImageProperties.Recommended)); err != nil {
            return fmt.Errorf("Error setting `recommended`: %+v", err)
        }
        d.Set("release_note_uri", galleryImageProperties.ReleaseNoteUri)
    }
    d.Set("gallery_image_name", galleryImageName)
    d.Set("gallery_name", galleryName)
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}


func resourceArmGalleryImageDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).galleryImagesClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    galleryName := id.Path["galleries"]
    galleryImageName := id.Path["images"]

    future, err := client.Delete(ctx, resourceGroup, galleryName, galleryImageName)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Gallery Image (Gallery Image Name %q / Gallery Name %q / Resource Group %q): %+v", galleryImageName, galleryName, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Gallery Image (Gallery Image Name %q / Gallery Name %q / Resource Group %q): %+v", galleryImageName, galleryName, resourceGroup, err)
        }
    }

    return nil
}

func expandArmGalleryImageDisallowed(input []interface{}) *compute.Disallowed {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    diskTypes := v["disk_types"].([]interface{})

    result := compute.Disallowed{
        DiskTypes: utils.ExpandStringSlice(diskTypes),
    }
    return &result
}

func convertStringToDate(input interface{}) *date.Time {
  v := input.(string)

  dateTime, err := date.ParseTime(time.RFC3339, v)
  if err != nil {
      log.Printf("[ERROR] Cannot convert an invalid string to RFC3339 date %q: %+v", v, err)
      return nil
  }

  result := date.Time{
      Time: dateTime,
  }
  return &result
}

func expandArmGalleryImageGalleryImageIdentifier(input []interface{}) *compute.GalleryImageIdentifier {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    publisher := v["publisher"].(string)
    offer := v["offer"].(string)
    sku := v["sku"].(string)

    result := compute.GalleryImageIdentifier{
        Offer: utils.String(offer),
        Publisher: utils.String(publisher),
        Sku: utils.String(sku),
    }
    return &result
}

func expandArmGalleryImageImagePurchasePlan(input []interface{}) *compute.ImagePurchasePlan {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    name := v["name"].(string)
    publisher := v["publisher"].(string)
    product := v["product"].(string)

    result := compute.ImagePurchasePlan{
        Name: utils.String(name),
        Product: utils.String(product),
        Publisher: utils.String(publisher),
    }
    return &result
}

func expandArmGalleryImageRecommendedMachineConfiguration(input []interface{}) *compute.RecommendedMachineConfiguration {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    vCpus := v["v_cpus"].([]interface{})
    memory := v["memory"].([]interface{})

    result := compute.RecommendedMachineConfiguration{
        Memory: expandArmGalleryImageResourceRange(memory),
        VCpus: expandArmGalleryImageResourceRange(vCpus),
    }
    return &result
}

func expandArmGalleryImageResourceRange(input []interface{}) *compute.ResourceRange {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    min := v["min"].(int)
    max := v["max"].(int)

    result := compute.ResourceRange{
        Max: utils.Int32(int32(max)),
        Min: utils.Int32(int32(min)),
    }
    return &result
}


func flattenArmGalleryImageDisallowed(input *compute.Disallowed) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    result["disk_types"] = utils.FlattenStringSlice(input.DiskTypes)

    return []interface{}{result}
}

func flattenArmGalleryImageGalleryImageIdentifier(input *compute.GalleryImageIdentifier) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if offer := input.Offer; offer != nil {
        result["offer"] = *offer
    }
    if publisher := input.Publisher; publisher != nil {
        result["publisher"] = *publisher
    }
    if sku := input.Sku; sku != nil {
        result["sku"] = *sku
    }

    return []interface{}{result}
}

func flattenArmGalleryImageImagePurchasePlan(input *compute.ImagePurchasePlan) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if name := input.Name; name != nil {
        result["name"] = *name
    }
    if product := input.Product; product != nil {
        result["product"] = *product
    }
    if publisher := input.Publisher; publisher != nil {
        result["publisher"] = *publisher
    }

    return []interface{}{result}
}

func flattenArmGalleryImageRecommendedMachineConfiguration(input *compute.RecommendedMachineConfiguration) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    result["memory"] = flattenArmGalleryImageResourceRange(input.Memory)
    result["v_cpus"] = flattenArmGalleryImageResourceRange(input.VCpus)

    return []interface{}{result}
}

func flattenArmGalleryImageResourceRange(input *compute.ResourceRange) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if max := input.Max; max != nil {
        result["max"] = int(*max)
    }
    if min := input.Min; min != nil {
        result["min"] = int(*min)
    }

    return []interface{}{result}
}
