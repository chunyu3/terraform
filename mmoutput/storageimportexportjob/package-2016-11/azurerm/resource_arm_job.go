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



func resourceArmJob() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmJobCreate,
        Read: resourceArmJobRead,
        Update: resourceArmJobUpdate,
        Delete: resourceArmJobDelete,

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

            "job_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "backup_drive_manifest": {
                Type: schema.TypeBool,
                Optional: true,
            },

            "cancel_requested": {
                Type: schema.TypeBool,
                Optional: true,
            },

            "delivery_package": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "carrier_name": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "drive_count": {
                            Type: schema.TypeInt,
                            Required: true,
                        },
                        "ship_date": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "tracking_number": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                    },
                },
            },

            "diagnostics_path": {
                Type: schema.TypeString,
                Optional: true,
            },

            "drive_list": {
                Type: schema.TypeList,
                Optional: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "bit_locker_key": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "bytes_succeeded": {
                            Type: schema.TypeInt,
                            Optional: true,
                        },
                        "copy_status": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "drive_header_hash": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "drive_id": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "error_log_uri": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "manifest_file": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "manifest_hash": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "manifest_uri": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "percent_complete": {
                            Type: schema.TypeInt,
                            Optional: true,
                        },
                        "state": {
                            Type: schema.TypeString,
                            Optional: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(storageimportexport.Specified),
                                string(storageimportexport.Received),
                                string(storageimportexport.NeverReceived),
                                string(storageimportexport.Transferring),
                                string(storageimportexport.Completed),
                                string(storageimportexport.CompletedMoreInfo),
                                string(storageimportexport.ShippedBack),
                            }, false),
                            Default: string(storageimportexport.Specified),
                        },
                        "verbose_log_uri": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },

            "export": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "blob_list": {
                            Type: schema.TypeList,
                            Optional: true,
                            MaxItems: 1,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "blob_path": {
                                        Type: schema.TypeList,
                                        Optional: true,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                        },
                                    },
                                    "blob_path_prefix": {
                                        Type: schema.TypeList,
                                        Optional: true,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                        },
                                    },
                                },
                            },
                        },
                        "blob_listblob_path": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },

            "incomplete_blob_list_uri": {
                Type: schema.TypeString,
                Optional: true,
            },

            "job_type": {
                Type: schema.TypeString,
                Optional: true,
            },

            "log_level": {
                Type: schema.TypeString,
                Optional: true,
            },

            "percent_complete": {
                Type: schema.TypeInt,
                Optional: true,
            },

            "return_address": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "city": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "country_or_region": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "email": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "phone": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "postal_code": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "recipient_name": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "street_address1": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "state_or_province": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "street_address2": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },

            "return_package": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "carrier_name": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "drive_count": {
                            Type: schema.TypeInt,
                            Required: true,
                        },
                        "ship_date": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "tracking_number": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                    },
                },
            },

            "return_shipping": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "carrier_account_number": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "carrier_name": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                    },
                },
            },

            "shipping_information": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "city": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "country_or_region": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "postal_code": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "recipient_name": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "state_or_province": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "street_address1": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "phone": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "street_address2": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },

            "state": {
                Type: schema.TypeString,
                Optional: true,
            },

            "storage_account_id": {
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

func resourceArmJobCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).jobsClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    jobName := d.Get("job_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, jobName, resourceGroup)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Job (Resource Group %q / Job Name %q): %+v", resourceGroup, jobName, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_job", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    backupDriveManifest := d.Get("backup_drive_manifest").(bool)
    cancelRequested := d.Get("cancel_requested").(bool)
    deliveryPackage := d.Get("delivery_package").([]interface{})
    diagnosticsPath := d.Get("diagnostics_path").(string)
    driveList := d.Get("drive_list").([]interface{})
    export := d.Get("export").([]interface{})
    incompleteBlobListUri := d.Get("incomplete_blob_list_uri").(string)
    jobType := d.Get("job_type").(string)
    logLevel := d.Get("log_level").(string)
    percentComplete := d.Get("percent_complete").(int)
    returnAddress := d.Get("return_address").([]interface{})
    returnPackage := d.Get("return_package").([]interface{})
    returnShipping := d.Get("return_shipping").([]interface{})
    shippingInformation := d.Get("shipping_information").([]interface{})
    state := d.Get("state").(string)
    storageAccountId := d.Get("storage_account_id").(string)
    t := d.Get("tags").(map[string]interface{})

    body := storageimportexport.PutJobParameters{
        Location: utils.String(location),
        JobDetails: &storageimportexport.JobDetails{
            BackupDriveManifest: utils.Bool(backupDriveManifest),
            CancelRequested: utils.Bool(cancelRequested),
            DeliveryPackage: expandArmJobPackageInfomation(deliveryPackage),
            DiagnosticsPath: utils.String(diagnosticsPath),
            DriveList: expandArmJobDriveStatus(driveList),
            Export: expandArmJobExport(export),
            IncompleteBlobListUri: utils.String(incompleteBlobListUri),
            JobType: utils.String(jobType),
            LogLevel: utils.String(logLevel),
            PercentComplete: utils.Int(percentComplete),
            ReturnAddress: expandArmJobReturnAddress(returnAddress),
            ReturnPackage: expandArmJobPackageInfomation(returnPackage),
            ReturnShipping: expandArmJobReturnShipping(returnShipping),
            ShippingInformation: expandArmJobShippingInformation(shippingInformation),
            State: utils.String(state),
            StorageAccountID: utils.String(storageAccountId),
        },
        Tags: tags.Expand(t),
    }


    if _, err := client.Create(ctx, jobName, resourceGroup, body); err != nil {
        return fmt.Errorf("Error creating Job (Resource Group %q / Job Name %q): %+v", resourceGroup, jobName, err)
    }


    resp, err := client.Get(ctx, jobName, resourceGroup)
    if err != nil {
        return fmt.Errorf("Error retrieving Job (Resource Group %q / Job Name %q): %+v", resourceGroup, jobName, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Job (Resource Group %q / Job Name %q) ID", resourceGroup, jobName)
    }
    d.SetId(*resp.ID)

    return resourceArmJobRead(d, meta)
}

func resourceArmJobRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).jobsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    jobName := id.Path["jobs"]
    resourceGroup := id.ResourceGroup

    resp, err := client.Get(ctx, jobName, resourceGroup)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Job %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Job (Resource Group %q / Job Name %q): %+v", resourceGroup, jobName, err)
    }


    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if jobDetails := resp.JobDetails; jobDetails != nil {
        d.Set("backup_drive_manifest", jobDetails.BackupDriveManifest)
        d.Set("cancel_requested", jobDetails.CancelRequested)
        if err := d.Set("delivery_package", flattenArmJobPackageInfomation(jobDetails.DeliveryPackage)); err != nil {
            return fmt.Errorf("Error setting `delivery_package`: %+v", err)
        }
        d.Set("diagnostics_path", jobDetails.DiagnosticsPath)
        if err := d.Set("drive_list", flattenArmJobDriveStatus(jobDetails.DriveList)); err != nil {
            return fmt.Errorf("Error setting `drive_list`: %+v", err)
        }
        if err := d.Set("export", flattenArmJobExport(jobDetails.Export)); err != nil {
            return fmt.Errorf("Error setting `export`: %+v", err)
        }
        d.Set("incomplete_blob_list_uri", jobDetails.IncompleteBlobListUri)
        d.Set("job_type", jobDetails.JobType)
        d.Set("log_level", jobDetails.LogLevel)
        d.Set("percent_complete", jobDetails.PercentComplete)
        d.Set("provisioning_state", jobDetails.ProvisioningState)
        if err := d.Set("return_address", flattenArmJobReturnAddress(jobDetails.ReturnAddress)); err != nil {
            return fmt.Errorf("Error setting `return_address`: %+v", err)
        }
        if err := d.Set("return_package", flattenArmJobPackageInfomation(jobDetails.ReturnPackage)); err != nil {
            return fmt.Errorf("Error setting `return_package`: %+v", err)
        }
        if err := d.Set("return_shipping", flattenArmJobReturnShipping(jobDetails.ReturnShipping)); err != nil {
            return fmt.Errorf("Error setting `return_shipping`: %+v", err)
        }
        if err := d.Set("shipping_information", flattenArmJobShippingInformation(jobDetails.ShippingInformation)); err != nil {
            return fmt.Errorf("Error setting `shipping_information`: %+v", err)
        }
        d.Set("state", jobDetails.State)
        d.Set("storage_account_id", jobDetails.StorageAccountID)
    }
    d.Set("job_name", jobName)
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}

func resourceArmJobUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).jobsClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    backupDriveManifest := d.Get("backup_drive_manifest").(bool)
    cancelRequested := d.Get("cancel_requested").(bool)
    deliveryPackage := d.Get("delivery_package").([]interface{})
    diagnosticsPath := d.Get("diagnostics_path").(string)
    driveList := d.Get("drive_list").([]interface{})
    export := d.Get("export").([]interface{})
    incompleteBlobListUri := d.Get("incomplete_blob_list_uri").(string)
    jobName := d.Get("job_name").(string)
    jobType := d.Get("job_type").(string)
    logLevel := d.Get("log_level").(string)
    percentComplete := d.Get("percent_complete").(int)
    returnAddress := d.Get("return_address").([]interface{})
    returnPackage := d.Get("return_package").([]interface{})
    returnShipping := d.Get("return_shipping").([]interface{})
    shippingInformation := d.Get("shipping_information").([]interface{})
    state := d.Get("state").(string)
    storageAccountId := d.Get("storage_account_id").(string)
    t := d.Get("tags").(map[string]interface{})

    body := storageimportexport.PutJobParameters{
        Location: utils.String(location),
        JobDetails: &storageimportexport.JobDetails{
            BackupDriveManifest: utils.Bool(backupDriveManifest),
            CancelRequested: utils.Bool(cancelRequested),
            DeliveryPackage: expandArmJobPackageInfomation(deliveryPackage),
            DiagnosticsPath: utils.String(diagnosticsPath),
            DriveList: expandArmJobDriveStatus(driveList),
            Export: expandArmJobExport(export),
            IncompleteBlobListUri: utils.String(incompleteBlobListUri),
            JobType: utils.String(jobType),
            LogLevel: utils.String(logLevel),
            PercentComplete: utils.Int(percentComplete),
            ReturnAddress: expandArmJobReturnAddress(returnAddress),
            ReturnPackage: expandArmJobPackageInfomation(returnPackage),
            ReturnShipping: expandArmJobReturnShipping(returnShipping),
            ShippingInformation: expandArmJobShippingInformation(shippingInformation),
            State: utils.String(state),
            StorageAccountID: utils.String(storageAccountId),
        },
        Tags: tags.Expand(t),
    }


    if _, err := client.Update(ctx, jobName, resourceGroup, body); err != nil {
        return fmt.Errorf("Error updating Job (Resource Group %q / Job Name %q): %+v", resourceGroup, jobName, err)
    }

    return resourceArmJobRead(d, meta)
}

func resourceArmJobDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).jobsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    jobName := id.Path["jobs"]
    resourceGroup := id.ResourceGroup

    if _, err := client.Delete(ctx, jobName, resourceGroup); err != nil {
        return fmt.Errorf("Error deleting Job (Resource Group %q / Job Name %q): %+v", resourceGroup, jobName, err)
    }

    return nil
}

func expandArmJobPackageInfomation(input []interface{}) *storageimportexport.PackageInfomation {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    carrierName := v["carrier_name"].(string)
    trackingNumber := v["tracking_number"].(string)
    driveCount := v["drive_count"].(int)
    shipDate := v["ship_date"].(string)

    result := storageimportexport.PackageInfomation{
        CarrierName: utils.String(carrierName),
        DriveCount: utils.Int(driveCount),
        ShipDate: utils.String(shipDate),
        TrackingNumber: utils.String(trackingNumber),
    }
    return &result
}

func expandArmJobDriveStatus(input []interface{}) *[]storageimportexport.DriveStatus {
    results := make([]storageimportexport.DriveStatus, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        driveId := v["drive_id"].(string)
        bitLockerKey := v["bit_locker_key"].(string)
        manifestFile := v["manifest_file"].(string)
        manifestHash := v["manifest_hash"].(string)
        driveHeaderHash := v["drive_header_hash"].(string)
        state := v["state"].(string)
        copyStatus := v["copy_status"].(string)
        percentComplete := v["percent_complete"].(int)
        verboseLogUri := v["verbose_log_uri"].(string)
        errorLogUri := v["error_log_uri"].(string)
        manifestUri := v["manifest_uri"].(string)
        bytesSucceeded := v["bytes_succeeded"].(int)

        result := storageimportexport.DriveStatus{
            BitLockerKey: utils.String(bitLockerKey),
            BytesSucceeded: utils.Int64(int64(bytesSucceeded)),
            CopyStatus: utils.String(copyStatus),
            DriveHeaderHash: utils.String(driveHeaderHash),
            DriveID: utils.String(driveId),
            ErrorLogUri: utils.String(errorLogUri),
            ManifestFile: utils.String(manifestFile),
            ManifestHash: utils.String(manifestHash),
            ManifestUri: utils.String(manifestUri),
            PercentComplete: utils.Int(percentComplete),
            State: storageimportexport.DriveState(state),
            VerboseLogUri: utils.String(verboseLogUri),
        }

        results = append(results, result)
    }
    return &results
}

func expandArmJobExport(input []interface{}) *storageimportexport.Export {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    blobList := v["blob_list"].([]interface{})
    blobListblobPath := v["blob_listblob_path"].(string)

    result := storageimportexport.Export{
        BlobList: expandArmJobExport_blobList(blobList),
        BlobListblobPath: utils.String(blobListblobPath),
    }
    return &result
}

func expandArmJobReturnAddress(input []interface{}) *storageimportexport.ReturnAddress {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    recipientName := v["recipient_name"].(string)
    streetAddress1 := v["street_address1"].(string)
    streetAddress2 := v["street_address2"].(string)
    city := v["city"].(string)
    stateOrProvince := v["state_or_province"].(string)
    postalCode := v["postal_code"].(string)
    countryOrRegion := v["country_or_region"].(string)
    phone := v["phone"].(string)
    email := v["email"].(string)

    result := storageimportexport.ReturnAddress{
        City: utils.String(city),
        CountryOrRegion: utils.String(countryOrRegion),
        Email: utils.String(email),
        Phone: utils.String(phone),
        PostalCode: utils.String(postalCode),
        RecipientName: utils.String(recipientName),
        StateOrProvince: utils.String(stateOrProvince),
        StreetAddress1: utils.String(streetAddress1),
        StreetAddress2: utils.String(streetAddress2),
    }
    return &result
}

func expandArmJobReturnShipping(input []interface{}) *storageimportexport.ReturnShipping {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    carrierName := v["carrier_name"].(string)
    carrierAccountNumber := v["carrier_account_number"].(string)

    result := storageimportexport.ReturnShipping{
        CarrierAccountNumber: utils.String(carrierAccountNumber),
        CarrierName: utils.String(carrierName),
    }
    return &result
}

func expandArmJobShippingInformation(input []interface{}) *storageimportexport.ShippingInformation {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    recipientName := v["recipient_name"].(string)
    streetAddress1 := v["street_address1"].(string)
    streetAddress2 := v["street_address2"].(string)
    city := v["city"].(string)
    stateOrProvince := v["state_or_province"].(string)
    postalCode := v["postal_code"].(string)
    countryOrRegion := v["country_or_region"].(string)
    phone := v["phone"].(string)

    result := storageimportexport.ShippingInformation{
        City: utils.String(city),
        CountryOrRegion: utils.String(countryOrRegion),
        Phone: utils.String(phone),
        PostalCode: utils.String(postalCode),
        RecipientName: utils.String(recipientName),
        StateOrProvince: utils.String(stateOrProvince),
        StreetAddress1: utils.String(streetAddress1),
        StreetAddress2: utils.String(streetAddress2),
    }
    return &result
}

func expandArmJobExport_blobList(input []interface{}) *storageimportexport.Export_blobList {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    blobPath := v["blob_path"].([]interface{})
    blobPathPrefix := v["blob_path_prefix"].([]interface{})

    result := storageimportexport.Export_blobList{
        BlobPath: utils.ExpandStringSlice(blobPath),
        BlobPathPrefix: utils.ExpandStringSlice(blobPathPrefix),
    }
    return &result
}


func flattenArmJobPackageInfomation(input *storageimportexport.PackageInfomation) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if carrierName := input.CarrierName; carrierName != nil {
        result["carrier_name"] = *carrierName
    }
    if driveCount := input.DriveCount; driveCount != nil {
        result["drive_count"] = *driveCount
    }
    if shipDate := input.ShipDate; shipDate != nil {
        result["ship_date"] = *shipDate
    }
    if trackingNumber := input.TrackingNumber; trackingNumber != nil {
        result["tracking_number"] = *trackingNumber
    }

    return []interface{}{result}
}

func flattenArmJobDriveStatus(input *[]storageimportexport.DriveStatus) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        if bitLockerKey := item.BitLockerKey; bitLockerKey != nil {
            v["bit_locker_key"] = *bitLockerKey
        }
        if bytesSucceeded := item.BytesSucceeded; bytesSucceeded != nil {
            v["bytes_succeeded"] = int(*bytesSucceeded)
        }
        if copyStatus := item.CopyStatus; copyStatus != nil {
            v["copy_status"] = *copyStatus
        }
        if driveHeaderHash := item.DriveHeaderHash; driveHeaderHash != nil {
            v["drive_header_hash"] = *driveHeaderHash
        }
        if driveId := item.DriveID; driveId != nil {
            v["drive_id"] = *driveId
        }
        if errorLogUri := item.ErrorLogUri; errorLogUri != nil {
            v["error_log_uri"] = *errorLogUri
        }
        if manifestFile := item.ManifestFile; manifestFile != nil {
            v["manifest_file"] = *manifestFile
        }
        if manifestHash := item.ManifestHash; manifestHash != nil {
            v["manifest_hash"] = *manifestHash
        }
        if manifestUri := item.ManifestUri; manifestUri != nil {
            v["manifest_uri"] = *manifestUri
        }
        if percentComplete := item.PercentComplete; percentComplete != nil {
            v["percent_complete"] = *percentComplete
        }
        v["state"] = string(item.State)
        if verboseLogUri := item.VerboseLogUri; verboseLogUri != nil {
            v["verbose_log_uri"] = *verboseLogUri
        }

        results = append(results, v)
    }

    return results
}

func flattenArmJobExport(input *storageimportexport.Export) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    result["blob_list"] = flattenArmJobExport_blobList(input.BlobList)
    if blobListblobPath := input.BlobListblobPath; blobListblobPath != nil {
        result["blob_listblob_path"] = *blobListblobPath
    }

    return []interface{}{result}
}

func flattenArmJobReturnAddress(input *storageimportexport.ReturnAddress) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if city := input.City; city != nil {
        result["city"] = *city
    }
    if countryOrRegion := input.CountryOrRegion; countryOrRegion != nil {
        result["country_or_region"] = *countryOrRegion
    }
    if email := input.Email; email != nil {
        result["email"] = *email
    }
    if phone := input.Phone; phone != nil {
        result["phone"] = *phone
    }
    if postalCode := input.PostalCode; postalCode != nil {
        result["postal_code"] = *postalCode
    }
    if recipientName := input.RecipientName; recipientName != nil {
        result["recipient_name"] = *recipientName
    }
    if stateOrProvince := input.StateOrProvince; stateOrProvince != nil {
        result["state_or_province"] = *stateOrProvince
    }
    if streetAddress1 := input.StreetAddress1; streetAddress1 != nil {
        result["street_address1"] = *streetAddress1
    }
    if streetAddress2 := input.StreetAddress2; streetAddress2 != nil {
        result["street_address2"] = *streetAddress2
    }

    return []interface{}{result}
}

func flattenArmJobReturnShipping(input *storageimportexport.ReturnShipping) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if carrierAccountNumber := input.CarrierAccountNumber; carrierAccountNumber != nil {
        result["carrier_account_number"] = *carrierAccountNumber
    }
    if carrierName := input.CarrierName; carrierName != nil {
        result["carrier_name"] = *carrierName
    }

    return []interface{}{result}
}

func flattenArmJobShippingInformation(input *storageimportexport.ShippingInformation) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if city := input.City; city != nil {
        result["city"] = *city
    }
    if countryOrRegion := input.CountryOrRegion; countryOrRegion != nil {
        result["country_or_region"] = *countryOrRegion
    }
    if phone := input.Phone; phone != nil {
        result["phone"] = *phone
    }
    if postalCode := input.PostalCode; postalCode != nil {
        result["postal_code"] = *postalCode
    }
    if recipientName := input.RecipientName; recipientName != nil {
        result["recipient_name"] = *recipientName
    }
    if stateOrProvince := input.StateOrProvince; stateOrProvince != nil {
        result["state_or_province"] = *stateOrProvince
    }
    if streetAddress1 := input.StreetAddress1; streetAddress1 != nil {
        result["street_address1"] = *streetAddress1
    }
    if streetAddress2 := input.StreetAddress2; streetAddress2 != nil {
        result["street_address2"] = *streetAddress2
    }

    return []interface{}{result}
}

func flattenArmJobExport_blobList(input *storageimportexport.Export_blobList) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    result["blob_path"] = utils.FlattenStringSlice(input.BlobPath)
    result["blob_path_prefix"] = utils.FlattenStringSlice(input.BlobPathPrefix)

    return []interface{}{result}
}
