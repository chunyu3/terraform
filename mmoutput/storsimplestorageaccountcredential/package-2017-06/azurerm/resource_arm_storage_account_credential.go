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



func resourceArmStorageAccountCredential() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmStorageAccountCredentialCreateUpdate,
        Read: resourceArmStorageAccountCredentialRead,
        Update: resourceArmStorageAccountCredentialCreateUpdate,
        Delete: resourceArmStorageAccountCredentialDelete,

        Importer: &schema.ResourceImporter{
            State: schema.ImportStatePassthrough,
        },


        Schema: map[string]*schema.Schema{
            "name": {
                Type: schema.TypeString,
                Computed: true,
            },

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "end_point": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "manager_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "ssl_status": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(storsimple.Enabled),
                    string(storsimple.Disabled),
                }, false),
            },

            "storage_account_credential_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "access_key": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "encryption_algorithm": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(storsimple.None),
                                string(storsimple.AES256),
                                string(storsimple.RSAES_PKCS1_v_1_5),
                            }, false),
                        },
                        "value": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "encryption_cert_thumbprint": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },

            "kind": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(storsimple.Series8000),
                }, false),
                Default: string(storsimple.Series8000),
            },

            "type": {
                Type: schema.TypeString,
                Computed: true,
            },

            "volumes_count": {
                Type: schema.TypeInt,
                Computed: true,
            },
        },
    }
}

func resourceArmStorageAccountCredentialCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).storageAccountCredentialsClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    managerName := d.Get("manager_name").(string)
    storageAccountCredentialName := d.Get("storage_account_credential_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, storageAccountCredentialName, resourceGroup, managerName)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Storage Account Credential (Manager Name %q / Resource Group %q / Storage Account Credential Name %q): %+v", managerName, resourceGroup, storageAccountCredentialName, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_storage_account_credential", *existing.ID)
        }
    }

    accessKey := d.Get("access_key").([]interface{})
    endPoint := d.Get("end_point").(string)
    kind := d.Get("kind").(string)
    sslStatus := d.Get("ssl_status").(string)

    parameters := storsimple.StorageAccountCredential{
        Kind: storsimple.Kind(kind),
        StorageAccountCredentialProperties: &storsimple.StorageAccountCredentialProperties{
            AccessKey: expandArmStorageAccountCredentialAsymmetricEncryptedSecret(accessKey),
            EndPoint: utils.String(endPoint),
            SslStatus: storsimple.SslStatus(sslStatus),
        },
    }


    future, err := client.CreateOrUpdate(ctx, storageAccountCredentialName, resourceGroup, managerName, parameters)
    if err != nil {
        return fmt.Errorf("Error creating Storage Account Credential (Manager Name %q / Resource Group %q / Storage Account Credential Name %q): %+v", managerName, resourceGroup, storageAccountCredentialName, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Storage Account Credential (Manager Name %q / Resource Group %q / Storage Account Credential Name %q): %+v", managerName, resourceGroup, storageAccountCredentialName, err)
    }


    resp, err := client.Get(ctx, storageAccountCredentialName, resourceGroup, managerName)
    if err != nil {
        return fmt.Errorf("Error retrieving Storage Account Credential (Manager Name %q / Resource Group %q / Storage Account Credential Name %q): %+v", managerName, resourceGroup, storageAccountCredentialName, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Storage Account Credential (Manager Name %q / Resource Group %q / Storage Account Credential Name %q) ID", managerName, resourceGroup, storageAccountCredentialName)
    }
    d.SetId(*resp.ID)

    return resourceArmStorageAccountCredentialRead(d, meta)
}

func resourceArmStorageAccountCredentialRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).storageAccountCredentialsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    storageAccountCredentialName := id.Path["storageAccountCredentials"]
    resourceGroup := id.ResourceGroup
    managerName := id.Path["managers"]

    resp, err := client.Get(ctx, storageAccountCredentialName, resourceGroup, managerName)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Storage Account Credential %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Storage Account Credential (Manager Name %q / Resource Group %q / Storage Account Credential Name %q): %+v", managerName, resourceGroup, storageAccountCredentialName, err)
    }


    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if storageAccountCredentialProperties := resp.StorageAccountCredentialProperties; storageAccountCredentialProperties != nil {
        if err := d.Set("access_key", flattenArmStorageAccountCredentialAsymmetricEncryptedSecret(storageAccountCredentialProperties.AccessKey)); err != nil {
            return fmt.Errorf("Error setting `access_key`: %+v", err)
        }
        d.Set("end_point", storageAccountCredentialProperties.EndPoint)
        d.Set("ssl_status", string(storageAccountCredentialProperties.SslStatus))
        d.Set("volumes_count", int(*storageAccountCredentialProperties.VolumesCount))
    }
    d.Set("kind", string(resp.Kind))
    d.Set("manager_name", managerName)
    d.Set("storage_account_credential_name", storageAccountCredentialName)
    d.Set("type", resp.Type)

    return nil
}


func resourceArmStorageAccountCredentialDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).storageAccountCredentialsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    storageAccountCredentialName := id.Path["storageAccountCredentials"]
    resourceGroup := id.ResourceGroup
    managerName := id.Path["managers"]

    future, err := client.Delete(ctx, storageAccountCredentialName, resourceGroup, managerName)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Storage Account Credential (Manager Name %q / Resource Group %q / Storage Account Credential Name %q): %+v", managerName, resourceGroup, storageAccountCredentialName, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Storage Account Credential (Manager Name %q / Resource Group %q / Storage Account Credential Name %q): %+v", managerName, resourceGroup, storageAccountCredentialName, err)
        }
    }

    return nil
}

func expandArmStorageAccountCredentialAsymmetricEncryptedSecret(input []interface{}) *storsimple.AsymmetricEncryptedSecret {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    value := v["value"].(string)
    encryptionCertThumbprint := v["encryption_cert_thumbprint"].(string)
    encryptionAlgorithm := v["encryption_algorithm"].(string)

    result := storsimple.AsymmetricEncryptedSecret{
        EncryptionAlgorithm: storsimple.EncryptionAlgorithm(encryptionAlgorithm),
        EncryptionCertThumbprint: utils.String(encryptionCertThumbprint),
        Value: utils.String(value),
    }
    return &result
}


func flattenArmStorageAccountCredentialAsymmetricEncryptedSecret(input *storsimple.AsymmetricEncryptedSecret) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    result["encryption_algorithm"] = string(input.EncryptionAlgorithm)
    if encryptionCertThumbprint := input.EncryptionCertThumbprint; encryptionCertThumbprint != nil {
        result["encryption_cert_thumbprint"] = *encryptionCertThumbprint
    }
    if value := input.Value; value != nil {
        result["value"] = *value
    }

    return []interface{}{result}
}
