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



func resourceArmUser() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmUserCreateUpdate,
        Read: resourceArmUserRead,
        Update: resourceArmUserCreateUpdate,
        Delete: resourceArmUserDelete,

        Importer: &schema.ResourceImporter{
            State: schema.ImportStatePassthrough,
        },


        Schema: map[string]*schema.Schema{
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

            "encrypted_password": {
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

            "share_access_rights": {
                Type: schema.TypeList,
                Optional: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "access_type": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(databoxedge.Change),
                                string(databoxedge.Read),
                                string(databoxedge.Custom),
                            }, false),
                        },
                        "share_id": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                    },
                },
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

func resourceArmUserCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).usersClient
    ctx, cancel := timeouts.ForCreateUpdate(meta.(*ArmClient).StopContext, d)
    defer cancel()

    resourceGroupName := d.Get("resource_group").(string)
    name := d.Get("device_name").(string)
    name := d.Get("name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroupName, name, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing User (Name %q / Resource Group %q / Device Name %q): %+v", name, resourceGroupName, name, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_user", *existing.ID)
        }
    }

    encryptedPassword := d.Get("encrypted_password").([]interface{})
    shareAccessRights := d.Get("share_access_rights").([]interface{})

    user := databoxedge.User{
        UserProperties: &databoxedge.UserProperties{
            EncryptedPassword: expandArmUserAsymmetricEncryptedSecret(encryptedPassword),
            ShareAccessRights: expandArmUserShareAccessRight(shareAccessRights),
        },
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroupName, name, name, user)
    if err != nil {
        return fmt.Errorf("Error creating User (Name %q / Resource Group %q / Device Name %q): %+v", name, resourceGroupName, name, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of User (Name %q / Resource Group %q / Device Name %q): %+v", name, resourceGroupName, name, err)
    }


    resp, err := client.Get(ctx, resourceGroupName, name, name)
    if err != nil {
        return fmt.Errorf("Error retrieving User (Name %q / Resource Group %q / Device Name %q): %+v", name, resourceGroupName, name, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read User (Name %q / Resource Group %q / Device Name %q) ID", name, resourceGroupName, name)
    }
    d.SetId(*resp.ID)

    return resourceArmUserRead(d, meta)
}

func resourceArmUserRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).usersClient
    ctx, cancel := timeouts.ForRead(meta.(*ArmClient).StopContext, d)
    defer cancel()

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroupName := id.ResourceGroup
    name := id.Path["dataBoxEdgeDevices"]
    name := id.Path["users"]

    resp, err := client.Get(ctx, resourceGroupName, name, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] User %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading User (Name %q / Resource Group %q / Device Name %q): %+v", name, resourceGroupName, name, err)
    }


    d.Set("resource_group", resourceGroupName)
    d.Set("device_name", name)
    if userProperties := resp.UserProperties; userProperties != nil {
        if err := d.Set("encrypted_password", flattenArmUserAsymmetricEncryptedSecret(userProperties.EncryptedPassword)); err != nil {
            return fmt.Errorf("Error setting `encrypted_password`: %+v", err)
        }
        if err := d.Set("share_access_rights", flattenArmUserShareAccessRight(userProperties.ShareAccessRights)); err != nil {
            return fmt.Errorf("Error setting `share_access_rights`: %+v", err)
        }
    }
    d.Set("id", resp.ID)
    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("type", resp.Type)

    return nil
}


func resourceArmUserDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).usersClient
    ctx, cancel := timeouts.ForDelete(meta.(*ArmClient).StopContext, d)
    defer cancel()


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroupName := id.ResourceGroup
    name := id.Path["dataBoxEdgeDevices"]
    name := id.Path["users"]

    future, err := client.Delete(ctx, resourceGroupName, name, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting User (Name %q / Resource Group %q / Device Name %q): %+v", name, resourceGroupName, name, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting User (Name %q / Resource Group %q / Device Name %q): %+v", name, resourceGroupName, name, err)
        }
    }

    return nil
}

func expandArmUserAsymmetricEncryptedSecret(input []interface{}) *databoxedge.AsymmetricEncryptedSecret {
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

func expandArmUserShareAccessRight(input []interface{}) *[]databoxedge.ShareAccessRight {
    results := make([]databoxedge.ShareAccessRight, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        shareID := v["share_id"].(string)
        accessType := v["access_type"].(string)

        result := databoxedge.ShareAccessRight{
            AccessType: databoxedge.ShareAccessType(accessType),
            ShareID: utils.String(shareID),
        }

        results = append(results, result)
    }
    return &results
}


func flattenArmUserAsymmetricEncryptedSecret(input *databoxedge.AsymmetricEncryptedSecret) []interface{} {
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

func flattenArmUserShareAccessRight(input *[]databoxedge.ShareAccessRight) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        v["access_type"] = string(item.AccessType)
        if shareId := item.ShareID; shareId != nil {
            v["share_id"] = *shareId
        }

        results = append(results, v)
    }

    return results
}
