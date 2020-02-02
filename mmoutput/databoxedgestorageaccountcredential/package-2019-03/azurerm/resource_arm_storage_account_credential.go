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
            "account_type": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(databoxedge.GeneralPurposeStorage),
                    string(databoxedge.BlobStorage),
                }, false),
            },

            "alias": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "device_name": {
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

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "ssl_status": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(databoxedge.Enabled),
                    string(databoxedge.Disabled),
                }, false),
            },

            "account_key": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "encryption_algorithm": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(databoxedge.None),
                                string(databoxedge.AES256),
                                string(databoxedge.RSAES_PKCS1_v_1_5),
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

            "blob_domain_name": {
                Type: schema.TypeString,
                Optional: true,
            },

            "connection_string": {
                Type: schema.TypeString,
                Optional: true,
            },

            "user_name": {
                Type: schema.TypeString,
                Optional: true,
            },

            "id": {
                Type: schema.TypeString,
                Computed: true,
            },

            "name": {
                Type: schema.TypeString,
                Computed: true,
            },

            "type": {
                Type: schema.TypeString,
                Computed: true,
            },
        },
    }
}

func resourceArmStorageAccountCredentialCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).storageAccountCredentialsClient
    ctx, cancel := timeouts.ForCreateUpdate(meta.(*ArmClient).StopContext, d)
    defer cancel()

    resourceGroupName := d.Get("resource_group").(string)
    name := d.Get("device_name").(string)
    name := d.Get("name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroupName, name, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Storage Account Credential (Name %q / Resource Group %q / Device Name %q): %+v", name, resourceGroupName, name, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_storage_account_credential", *existing.ID)
        }
    }

    accountKey := d.Get("account_key").([]interface{})
    accountType := d.Get("account_type").(string)
    alias := d.Get("alias").(string)
    blobDomainName := d.Get("blob_domain_name").(string)
    connectionString := d.Get("connection_string").(string)
    sslStatus := d.Get("ssl_status").(string)
    userName := d.Get("user_name").(string)

    storageAccountCredential := databoxedge.StorageAccountCredential{
        StorageAccountCredentialProperties: &databoxedge.StorageAccountCredentialProperties{
            AccountKey: expandArmStorageAccountCredentialAsymmetricEncryptedSecret(accountKey),
            AccountType: databoxedge.AccountType(accountType),
            Alias: utils.String(alias),
            BlobDomainName: utils.String(blobDomainName),
            ConnectionString: utils.String(connectionString),
            SslStatus: databoxedge.SSLStatus(sslStatus),
            UserName: utils.String(userName),
        },
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroupName, name, name, storageAccountCredential)
    if err != nil {
        return fmt.Errorf("Error creating Storage Account Credential (Name %q / Resource Group %q / Device Name %q): %+v", name, resourceGroupName, name, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Storage Account Credential (Name %q / Resource Group %q / Device Name %q): %+v", name, resourceGroupName, name, err)
    }


    resp, err := client.Get(ctx, resourceGroupName, name, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Storage Account Credential (Name %q / Resource Group %q / Device Name %q): %+v", name, resourceGroupName, name, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Storage Account Credential (Name %q / Resource Group %q / Device Name %q) ID", name, resourceGroupName, name)
    }
    d.SetId(*resp.ID)

    return resourceArmStorageAccountCredentialRead(d, meta)
}

func resourceArmStorageAccountCredentialRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).storageAccountCredentialsClient
    ctx, cancel := timeouts.ForRead(meta.(*ArmClient).StopContext, d)
    defer cancel()

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroupName := id.ResourceGroup
    name := id.Path["dataBoxEdgeDevices"]
    name := id.Path["storageAccountCredentials"]

    resp, err := client.Get(ctx, resourceGroupName, name, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Storage Account Credential %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Storage Account Credential (Name %q / Resource Group %q / Device Name %q): %+v", name, resourceGroupName, name, err)
    }


    d.Set("resource_group", resourceGroupName)
    if storageAccountCredentialProperties := resp.StorageAccountCredentialProperties; storageAccountCredentialProperties != nil {
        if err := d.Set("account_key", flattenArmStorageAccountCredentialAsymmetricEncryptedSecret(storageAccountCredentialProperties.AccountKey)); err != nil {
            return fmt.Errorf("Error setting `account_key`: %+v", err)
        }
        d.Set("account_type", string(storageAccountCredentialProperties.AccountType))
        d.Set("alias", storageAccountCredentialProperties.Alias)
        d.Set("blob_domain_name", storageAccountCredentialProperties.BlobDomainName)
        d.Set("connection_string", storageAccountCredentialProperties.ConnectionString)
        d.Set("ssl_status", string(storageAccountCredentialProperties.SslStatus))
        d.Set("user_name", storageAccountCredentialProperties.UserName)
    }
    d.Set("device_name", name)
    d.Set("id", resp.ID)
    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("type", resp.Type)

    return nil
}


func resourceArmStorageAccountCredentialDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).storageAccountCredentialsClient
    ctx, cancel := timeouts.ForDelete(meta.(*ArmClient).StopContext, d)
    defer cancel()


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroupName := id.ResourceGroup
    name := id.Path["dataBoxEdgeDevices"]
    name := id.Path["storageAccountCredentials"]

    future, err := client.Delete(ctx, resourceGroupName, name, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Storage Account Credential (Name %q / Resource Group %q / Device Name %q): %+v", name, resourceGroupName, name, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Storage Account Credential (Name %q / Resource Group %q / Device Name %q): %+v", name, resourceGroupName, name, err)
        }
    }

    return nil
}

func expandArmStorageAccountCredentialAsymmetricEncryptedSecret(input []interface{}) *databoxedge.AsymmetricEncryptedSecret {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    value := v["value"].(string)
    encryptionCertThumbprint := v["encryption_cert_thumbprint"].(string)
    encryptionAlgorithm := v["encryption_algorithm"].(string)

    result := databoxedge.AsymmetricEncryptedSecret{
        EncryptionAlgorithm: databoxedge.EncryptionAlgorithm(encryptionAlgorithm),
        EncryptionCertThumbprint: utils.String(encryptionCertThumbprint),
        Value: utils.String(value),
    }
    return &result
}


func flattenArmStorageAccountCredentialAsymmetricEncryptedSecret(input *databoxedge.AsymmetricEncryptedSecret) []interface{} {
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
