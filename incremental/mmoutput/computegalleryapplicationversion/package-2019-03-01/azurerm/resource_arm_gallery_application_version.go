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



func resourceArmGalleryApplicationVersion() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmGalleryApplicationVersionCreateUpdate,
        Read: resourceArmGalleryApplicationVersionRead,
        Update: resourceArmGalleryApplicationVersionCreateUpdate,
        Delete: resourceArmGalleryApplicationVersionDelete,

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

            "gallery_application_name": {
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

            "publishing_profile": {
                Type: schema.TypeList,
                Required: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "source": {
                            Type: schema.TypeList,
                            Required: true,
                            MaxItems: 1,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "file_name": {
                                        Type: schema.TypeString,
                                        Required: true,
                                        ValidateFunc: validate.NoEmptyStrings,
                                    },
                                    "media_link": {
                                        Type: schema.TypeString,
                                        Required: true,
                                        ValidateFunc: validate.NoEmptyStrings,
                                    },
                                },
                            },
                        },
                        "content_type": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "enable_health_check": {
                            Type: schema.TypeBool,
                            Optional: true,
                        },
                        "end_of_life_date": {
                            Type: schema.TypeString,
                            Optional: true,
                            ValidateFunc: validateRFC3339Date,
                        },
                        "exclude_from_latest": {
                            Type: schema.TypeBool,
                            Optional: true,
                        },
                        "replica_count": {
                            Type: schema.TypeInt,
                            Optional: true,
                        },
                        "storage_account_type": {
                            Type: schema.TypeString,
                            Optional: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(compute.Standard_LRS),
                                string(compute.Standard_ZRS),
                            }, false),
                            Default: string(compute.Standard_LRS),
                        },
                        "target_regions": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "name": {
                                        Type: schema.TypeString,
                                        Required: true,
                                        ValidateFunc: validate.NoEmptyStrings,
                                    },
                                    "regional_replica_count": {
                                        Type: schema.TypeInt,
                                        Optional: true,
                                    },
                                    "storage_account_type": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                        ValidateFunc: validation.StringInSlice([]string{
                                            string(compute.Standard_LRS),
                                            string(compute.Standard_ZRS),
                                        }, false),
                                        Default: string(compute.Standard_LRS),
                                    },
                                },
                            },
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

func resourceArmGalleryApplicationVersionCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).galleryApplicationVersionsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    galleryApplicationName := d.Get("gallery_application_name").(string)
    galleryName := d.Get("gallery_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, galleryName, galleryApplicationName, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Gallery Application Version %q (Gallery Application Name %q / Gallery Name %q / Resource Group %q): %+v", name, galleryApplicationName, galleryName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_gallery_application_version", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    publishingProfile := d.Get("publishing_profile").([]interface{})
    t := d.Get("tags").(map[string]interface{})

    galleryApplicationVersion := compute.GalleryApplicationVersion{
        Location: utils.String(location),
        GalleryApplicationVersionProperties: &compute.GalleryApplicationVersionProperties{
            PublishingProfile: expandArmGalleryApplicationVersionGalleryApplicationVersionPublishingProfile(publishingProfile),
        },
        Tags: tags.Expand(t),
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, galleryName, galleryApplicationName, name, galleryApplicationVersion)
    if err != nil {
        return fmt.Errorf("Error creating Gallery Application Version %q (Gallery Application Name %q / Gallery Name %q / Resource Group %q): %+v", name, galleryApplicationName, galleryName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Gallery Application Version %q (Gallery Application Name %q / Gallery Name %q / Resource Group %q): %+v", name, galleryApplicationName, galleryName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, galleryName, galleryApplicationName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Gallery Application Version %q (Gallery Application Name %q / Gallery Name %q / Resource Group %q): %+v", name, galleryApplicationName, galleryName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Gallery Application Version %q (Gallery Application Name %q / Gallery Name %q / Resource Group %q) ID", name, galleryApplicationName, galleryName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmGalleryApplicationVersionRead(d, meta)
}

func resourceArmGalleryApplicationVersionRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).galleryApplicationVersionsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    galleryName := id.Path["galleries"]
    galleryApplicationName := id.Path["applications"]
    name := id.Path["versions"]

    resp, err := client.Get(ctx, resourceGroup, galleryName, galleryApplicationName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Gallery Application Version %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Gallery Application Version %q (Gallery Application Name %q / Gallery Name %q / Resource Group %q): %+v", name, galleryApplicationName, galleryName, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    d.Set("gallery_application_name", galleryApplicationName)
    d.Set("gallery_name", galleryName)
    d.Set("type", resp.Type)

    return nil
}


func resourceArmGalleryApplicationVersionDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).galleryApplicationVersionsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    galleryName := id.Path["galleries"]
    galleryApplicationName := id.Path["applications"]
    name := id.Path["versions"]

    future, err := client.Delete(ctx, resourceGroup, galleryName, galleryApplicationName, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Gallery Application Version %q (Gallery Application Name %q / Gallery Name %q / Resource Group %q): %+v", name, galleryApplicationName, galleryName, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Gallery Application Version %q (Gallery Application Name %q / Gallery Name %q / Resource Group %q): %+v", name, galleryApplicationName, galleryName, resourceGroup, err)
        }
    }

    return nil
}

func expandArmGalleryApplicationVersionGalleryApplicationVersionPublishingProfile(input []interface{}) *compute.GalleryApplicationVersionPublishingProfile {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    targetRegions := v["target_regions"].([]interface{})
    replicaCount := v["replica_count"].(int)
    excludeFromLatest := v["exclude_from_latest"].(bool)
    endOfLifeDate := v["end_of_life_date"].(string)
    storageAccountType := v["storage_account_type"].(string)
    source := v["source"].([]interface{})
    contentType := v["content_type"].(string)
    enableHealthCheck := v["enable_health_check"].(bool)

    result := compute.GalleryApplicationVersionPublishingProfile{
        ContentType: utils.String(contentType),
        EnableHealthCheck: utils.Bool(enableHealthCheck),
        EndOfLifeDate: convertStringToDate(endOfLifeDate),
        ExcludeFromLatest: utils.Bool(excludeFromLatest),
        ReplicaCount: utils.Int32(int32(replicaCount)),
        Source: expandArmGalleryApplicationVersionUserArtifactSource(source),
        StorageAccountType: compute.StorageAccountType(storageAccountType),
        TargetRegions: expandArmGalleryApplicationVersionTargetRegion(targetRegions),
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

func expandArmGalleryApplicationVersionUserArtifactSource(input []interface{}) *compute.UserArtifactSource {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    fileName := v["file_name"].(string)
    mediaLink := v["media_link"].(string)

    result := compute.UserArtifactSource{
        FileName: utils.String(fileName),
        MediaLink: utils.String(mediaLink),
    }
    return &result
}

func expandArmGalleryApplicationVersionTargetRegion(input []interface{}) *[]compute.TargetRegion {
    results := make([]compute.TargetRegion, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        name := v["name"].(string)
        regionalReplicaCount := v["regional_replica_count"].(int)
        storageAccountType := v["storage_account_type"].(string)

        result := compute.TargetRegion{
            Name: utils.String(name),
            RegionalReplicaCount: utils.Int32(int32(regionalReplicaCount)),
            StorageAccountType: compute.StorageAccountType(storageAccountType),
        }

        results = append(results, result)
    }
    return &results
}
