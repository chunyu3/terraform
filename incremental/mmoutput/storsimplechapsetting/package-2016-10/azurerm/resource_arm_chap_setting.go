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



func resourceArmChapSetting() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmChapSettingCreateUpdate,
        Read: resourceArmChapSettingRead,
        Update: resourceArmChapSettingCreateUpdate,
        Delete: resourceArmChapSettingDelete,

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

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "device_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "manager_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "password": {
                Type: schema.TypeList,
                Required: true,
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
                        "encryption_certificate_thumbprint": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },

            "type": {
                Type: schema.TypeString,
                Computed: true,
            },
        },
    }
}

func resourceArmChapSettingCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).chapSettingsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    deviceName := d.Get("device_name").(string)
    managerName := d.Get("manager_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, managerName, deviceName, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Chap Setting %q (Manager Name %q / Resource Group %q / Device Name %q): %+v", name, managerName, resourceGroup, deviceName, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_chap_setting", *existing.ID)
        }
    }

    password := d.Get("password").([]interface{})

    chapSetting := storsimple.ChapSettings{
        ChapProperties: &storsimple.ChapProperties{
            Password: expandArmChapSettingAsymmetricEncryptedSecret(password),
        },
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, managerName, deviceName, name, chapSetting)
    if err != nil {
        return fmt.Errorf("Error creating Chap Setting %q (Manager Name %q / Resource Group %q / Device Name %q): %+v", name, managerName, resourceGroup, deviceName, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Chap Setting %q (Manager Name %q / Resource Group %q / Device Name %q): %+v", name, managerName, resourceGroup, deviceName, err)
    }


    resp, err := client.Get(ctx, resourceGroup, managerName, deviceName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Chap Setting %q (Manager Name %q / Resource Group %q / Device Name %q): %+v", name, managerName, resourceGroup, deviceName, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Chap Setting %q (Manager Name %q / Resource Group %q / Device Name %q) ID", name, managerName, resourceGroup, deviceName)
    }
    d.SetId(*resp.ID)

    return resourceArmChapSettingRead(d, meta)
}

func resourceArmChapSettingRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).chapSettingsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    managerName := id.Path["managers"]
    deviceName := id.Path["devices"]
    name := id.Path["chapSettings"]

    resp, err := client.Get(ctx, resourceGroup, managerName, deviceName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Chap Setting %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Chap Setting %q (Manager Name %q / Resource Group %q / Device Name %q): %+v", name, managerName, resourceGroup, deviceName, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    d.Set("device_name", deviceName)
    d.Set("manager_name", managerName)
    d.Set("type", resp.Type)

    return nil
}


func resourceArmChapSettingDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).chapSettingsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    managerName := id.Path["managers"]
    deviceName := id.Path["devices"]
    name := id.Path["chapSettings"]

    future, err := client.Delete(ctx, resourceGroup, managerName, deviceName, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Chap Setting %q (Manager Name %q / Resource Group %q / Device Name %q): %+v", name, managerName, resourceGroup, deviceName, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Chap Setting %q (Manager Name %q / Resource Group %q / Device Name %q): %+v", name, managerName, resourceGroup, deviceName, err)
        }
    }

    return nil
}

func expandArmChapSettingAsymmetricEncryptedSecret(input []interface{}) *storsimple.AsymmetricEncryptedSecret {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    value := v["value"].(string)
    encryptionCertificateThumbprint := v["encryption_certificate_thumbprint"].(string)
    encryptionAlgorithm := v["encryption_algorithm"].(string)

    result := storsimple.AsymmetricEncryptedSecret{
        EncryptionAlgorithm: storsimple.EncryptionAlgorithm(encryptionAlgorithm),
        EncryptionCertificateThumbprint: utils.String(encryptionCertificateThumbprint),
        Value: utils.String(value),
    }
    return &result
}