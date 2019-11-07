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



func resourceArmNotificationHub() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmNotificationHubCreateUpdate,
        Read: resourceArmNotificationHubRead,
        Update: resourceArmNotificationHubCreateUpdate,
        Delete: resourceArmNotificationHubDelete,

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
                Optional: true,
            },

            "name": {
                Type: schema.TypeString,
                Computed: true,
            },

            "location": azure.SchemaLocation(),

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "namespace_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "adm_credential": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "auth_token_url": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "client_id": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "client_secret": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },

            "apns_credential": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "apns_certificate": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "certificate_key": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "endpoint": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "thumbprint": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },

            "authorization_rules": {
                Type: schema.TypeList,
                Optional: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "rights": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                                ValidateFunc: validation.StringInSlice([]string{
                                    string(notificationhubs.Manage),
                                    string(notificationhubs.Send),
                                    string(notificationhubs.Listen),
                               }, false),
                            },
                        },
                    },
                },
            },

            "baidu_credential": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "baidu_api_key": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "baidu_end_point": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "baidu_secret_key": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },

            "gcm_credential": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "gcm_endpoint": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "google_api_key": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },

            "mpns_credential": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "certificate_key": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "mpns_certificate": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "thumbprint": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },

            "registration_ttl": {
                Type: schema.TypeString,
                Optional: true,
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
                            ValidateFunc: validation.StringInSlice([]string{
                                string(notificationhubs.Free),
                                string(notificationhubs.Basic),
                                string(notificationhubs.Standard),
                            }, false),
                        },
                        "capacity": {
                            Type: schema.TypeInt,
                            Optional: true,
                        },
                        "family": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "size": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "tier": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },

            "wns_credential": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "package_sid": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "secret_key": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "windows_live_endpoint": {
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

            "tags": tags.Schema(),
        },
    }
}

func resourceArmNotificationHubCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).notificationHubsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    namespaceName := d.Get("namespace_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, namespaceName, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Notification Hub %q (Namespace Name %q / Resource Group %q): %+v", name, namespaceName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_notification_hub", *existing.ID)
        }
    }

    name := d.Get("name").(string)
    location := azure.NormalizeLocation(d.Get("location").(string))
    admCredential := d.Get("adm_credential").([]interface{})
    apnsCredential := d.Get("apns_credential").([]interface{})
    authorizationRules := d.Get("authorization_rules").([]interface{})
    baiduCredential := d.Get("baidu_credential").([]interface{})
    gcmCredential := d.Get("gcm_credential").([]interface{})
    mpnsCredential := d.Get("mpns_credential").([]interface{})
    registrationTtl := d.Get("registration_ttl").(string)
    sku := d.Get("sku").([]interface{})
    wnsCredential := d.Get("wns_credential").([]interface{})
    t := d.Get("tags").(map[string]interface{})

    parameters := notificationhubs.NotificationHubCreateOrUpdateParameters{
        Location: utils.String(location),
        NotificationHubProperties: &notificationhubs.NotificationHubProperties{
            AdmCredential: expandArmNotificationHubAdmCredential(admCredential),
            ApnsCredential: expandArmNotificationHubApnsCredential(apnsCredential),
            AuthorizationRules: expandArmNotificationHubSharedAccessAuthorizationRuleProperties(authorizationRules),
            BaiduCredential: expandArmNotificationHubBaiduCredential(baiduCredential),
            GcmCredential: expandArmNotificationHubGcmCredential(gcmCredential),
            MpnsCredential: expandArmNotificationHubMpnsCredential(mpnsCredential),
            Name: utils.String(name),
            RegistrationTtl: utils.String(registrationTtl),
            WnsCredential: expandArmNotificationHubWnsCredential(wnsCredential),
        },
        Sku: expandArmNotificationHubSku(sku),
        Tags: tags.Expand(t),
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroup, namespaceName, name, parameters); err != nil {
        return fmt.Errorf("Error creating Notification Hub %q (Namespace Name %q / Resource Group %q): %+v", name, namespaceName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, namespaceName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Notification Hub %q (Namespace Name %q / Resource Group %q): %+v", name, namespaceName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Notification Hub %q (Namespace Name %q / Resource Group %q) ID", name, namespaceName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmNotificationHubRead(d, meta)
}

func resourceArmNotificationHubRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).notificationHubsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    namespaceName := id.Path["namespaces"]
    name := id.Path["notificationHubs"]

    resp, err := client.Get(ctx, resourceGroup, namespaceName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Notification Hub %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Notification Hub %q (Namespace Name %q / Resource Group %q): %+v", name, namespaceName, resourceGroup, err)
    }


    d.Set("name", name)
    if notificationHubProperties := resp.NotificationHubProperties; notificationHubProperties != nil {
        d.Set("name", notificationHubProperties.Name)
        if err := d.Set("adm_credential", flattenArmNotificationHubAdmCredential(notificationHubProperties.AdmCredential)); err != nil {
            return fmt.Errorf("Error setting `adm_credential`: %+v", err)
        }
        if err := d.Set("apns_credential", flattenArmNotificationHubApnsCredential(notificationHubProperties.ApnsCredential)); err != nil {
            return fmt.Errorf("Error setting `apns_credential`: %+v", err)
        }
        if err := d.Set("authorization_rules", flattenArmNotificationHubSharedAccessAuthorizationRuleProperties(notificationHubProperties.AuthorizationRules)); err != nil {
            return fmt.Errorf("Error setting `authorization_rules`: %+v", err)
        }
        if err := d.Set("baidu_credential", flattenArmNotificationHubBaiduCredential(notificationHubProperties.BaiduCredential)); err != nil {
            return fmt.Errorf("Error setting `baidu_credential`: %+v", err)
        }
        if err := d.Set("gcm_credential", flattenArmNotificationHubGcmCredential(notificationHubProperties.GcmCredential)); err != nil {
            return fmt.Errorf("Error setting `gcm_credential`: %+v", err)
        }
        if err := d.Set("mpns_credential", flattenArmNotificationHubMpnsCredential(notificationHubProperties.MpnsCredential)); err != nil {
            return fmt.Errorf("Error setting `mpns_credential`: %+v", err)
        }
        d.Set("registration_ttl", notificationHubProperties.RegistrationTtl)
        if err := d.Set("wns_credential", flattenArmNotificationHubWnsCredential(notificationHubProperties.WnsCredential)); err != nil {
            return fmt.Errorf("Error setting `wns_credential`: %+v", err)
        }
    }
    d.Set("name", resp.Name)
    d.Set("name", resp.Name)
    d.Set("name", resp.Name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    d.Set("namespace_name", namespaceName)
    if err := d.Set("sku", flattenArmNotificationHubSku(resp.Sku)); err != nil {
        return fmt.Errorf("Error setting `sku`: %+v", err)
    }
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}


func resourceArmNotificationHubDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).notificationHubsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    namespaceName := id.Path["namespaces"]
    name := id.Path["notificationHubs"]

    if _, err := client.Delete(ctx, resourceGroup, namespaceName, name); err != nil {
        return fmt.Errorf("Error deleting Notification Hub %q (Namespace Name %q / Resource Group %q): %+v", name, namespaceName, resourceGroup, err)
    }

    return nil
}

func expandArmNotificationHubAdmCredential(input []interface{}) *notificationhubs.AdmCredential {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    clientId := v["client_id"].(string)
    clientSecret := v["client_secret"].(string)
    authTokenUrl := v["auth_token_url"].(string)

    result := notificationhubs.AdmCredential{
        AdmCredentialProperties: &notificationhubs.AdmCredentialProperties{
            AuthTokenURL: utils.String(authTokenUrl),
            ClientID: utils.String(clientId),
            ClientSecret: utils.String(clientSecret),
        },
    }
    return &result
}

func expandArmNotificationHubApnsCredential(input []interface{}) *notificationhubs.ApnsCredential {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    apnsCertificate := v["apns_certificate"].(string)
    certificateKey := v["certificate_key"].(string)
    endpoint := v["endpoint"].(string)
    thumbprint := v["thumbprint"].(string)

    result := notificationhubs.ApnsCredential{
        ApnsCredentialProperties: &notificationhubs.ApnsCredentialProperties{
            ApnsCertificate: utils.String(apnsCertificate),
            CertificateKey: utils.String(certificateKey),
            Endpoint: utils.String(endpoint),
            Thumbprint: utils.String(thumbprint),
        },
    }
    return &result
}

func expandArmNotificationHubSharedAccessAuthorizationRuleProperties(input []interface{}) *[]notificationhubs.SharedAccessAuthorizationRuleProperties {
    results := make([]notificationhubs.SharedAccessAuthorizationRuleProperties, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        rights := v["rights"].([]interface{})

        result := notificationhubs.SharedAccessAuthorizationRuleProperties{
            Rights: expandArmNotificationHub(rights),
        }

        results = append(results, result)
    }
    return &results
}

func expandArmNotificationHubBaiduCredential(input []interface{}) *notificationhubs.BaiduCredential {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    baiduApiKey := v["baidu_api_key"].(string)
    baiduEndPoint := v["baidu_end_point"].(string)
    baiduSecretKey := v["baidu_secret_key"].(string)

    result := notificationhubs.BaiduCredential{
        BaiduCredentialProperties: &notificationhubs.BaiduCredentialProperties{
            BaiduApiKey: utils.String(baiduApiKey),
            BaiduEndPoint: utils.String(baiduEndPoint),
            BaiduSecretKey: utils.String(baiduSecretKey),
        },
    }
    return &result
}

func expandArmNotificationHubGcmCredential(input []interface{}) *notificationhubs.GcmCredential {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    gcmEndpoint := v["gcm_endpoint"].(string)
    googleApiKey := v["google_api_key"].(string)

    result := notificationhubs.GcmCredential{
        GcmCredentialProperties: &notificationhubs.GcmCredentialProperties{
            GcmEndpoint: utils.String(gcmEndpoint),
            GoogleApiKey: utils.String(googleApiKey),
        },
    }
    return &result
}

func expandArmNotificationHubMpnsCredential(input []interface{}) *notificationhubs.MpnsCredential {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    mpnsCertificate := v["mpns_certificate"].(string)
    certificateKey := v["certificate_key"].(string)
    thumbprint := v["thumbprint"].(string)

    result := notificationhubs.MpnsCredential{
        MpnsCredentialProperties: &notificationhubs.MpnsCredentialProperties{
            CertificateKey: utils.String(certificateKey),
            MpnsCertificate: utils.String(mpnsCertificate),
            Thumbprint: utils.String(thumbprint),
        },
    }
    return &result
}

func expandArmNotificationHubWnsCredential(input []interface{}) *notificationhubs.WnsCredential {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    packageSid := v["package_sid"].(string)
    secretKey := v["secret_key"].(string)
    windowsLiveEndpoint := v["windows_live_endpoint"].(string)

    result := notificationhubs.WnsCredential{
        WnsCredentialProperties: &notificationhubs.WnsCredentialProperties{
            PackageSid: utils.String(packageSid),
            SecretKey: utils.String(secretKey),
            WindowsLiveEndpoint: utils.String(windowsLiveEndpoint),
        },
    }
    return &result
}

func expandArmNotificationHubSku(input []interface{}) *notificationhubs.Sku {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    name := v["name"].(string)
    tier := v["tier"].(string)
    size := v["size"].(string)
    family := v["family"].(string)
    capacity := v["capacity"].(int)

    result := notificationhubs.Sku{
        Capacity: utils.Int(capacity),
        Family: utils.String(family),
        Name: notificationhubs.SkuName(name),
        Size: utils.String(size),
        Tier: utils.String(tier),
    }
    return &result
}

func expandArmNotificationHub(input []interface{}) *[]notificationhubs. {
    results := make([]notificationhubs., 0)
    for _, item := range input {
        v := item.(string)
        result := notificationhubs.(v)
        results = append(results, result)
    }
    return &results
}


func flattenArmNotificationHubAdmCredential(input *notificationhubs.AdmCredential) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if admCredentialProperties := input.AdmCredentialProperties; admCredentialProperties != nil {
        if authTokenUrl := admCredentialProperties.AuthTokenURL; authTokenUrl != nil {
            result["auth_token_url"] = *authTokenUrl
        }
        if clientId := admCredentialProperties.ClientID; clientId != nil {
            result["client_id"] = *clientId
        }
        if clientSecret := admCredentialProperties.ClientSecret; clientSecret != nil {
            result["client_secret"] = *clientSecret
        }
    }

    return []interface{}{result}
}

func flattenArmNotificationHubApnsCredential(input *notificationhubs.ApnsCredential) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if apnsCredentialProperties := input.ApnsCredentialProperties; apnsCredentialProperties != nil {
        if apnsCertificate := apnsCredentialProperties.ApnsCertificate; apnsCertificate != nil {
            result["apns_certificate"] = *apnsCertificate
        }
        if certificateKey := apnsCredentialProperties.CertificateKey; certificateKey != nil {
            result["certificate_key"] = *certificateKey
        }
        if endpoint := apnsCredentialProperties.Endpoint; endpoint != nil {
            result["endpoint"] = *endpoint
        }
        if thumbprint := apnsCredentialProperties.Thumbprint; thumbprint != nil {
            result["thumbprint"] = *thumbprint
        }
    }

    return []interface{}{result}
}

func flattenArmNotificationHubSharedAccessAuthorizationRuleProperties(input *[]notificationhubs.SharedAccessAuthorizationRuleProperties) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        v["rights"] = flattenArmNotificationHub(string(item.Rights))

        results = append(results, v)
    }

    return results
}

func flattenArmNotificationHubBaiduCredential(input *notificationhubs.BaiduCredential) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if baiduCredentialProperties := input.BaiduCredentialProperties; baiduCredentialProperties != nil {
        if baiduApiKey := baiduCredentialProperties.BaiduApiKey; baiduApiKey != nil {
            result["baidu_api_key"] = *baiduApiKey
        }
        if baiduEndPoint := baiduCredentialProperties.BaiduEndPoint; baiduEndPoint != nil {
            result["baidu_end_point"] = *baiduEndPoint
        }
        if baiduSecretKey := baiduCredentialProperties.BaiduSecretKey; baiduSecretKey != nil {
            result["baidu_secret_key"] = *baiduSecretKey
        }
    }

    return []interface{}{result}
}

func flattenArmNotificationHubGcmCredential(input *notificationhubs.GcmCredential) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if gcmCredentialProperties := input.GcmCredentialProperties; gcmCredentialProperties != nil {
        if gcmEndpoint := gcmCredentialProperties.GcmEndpoint; gcmEndpoint != nil {
            result["gcm_endpoint"] = *gcmEndpoint
        }
        if googleApiKey := gcmCredentialProperties.GoogleApiKey; googleApiKey != nil {
            result["google_api_key"] = *googleApiKey
        }
    }

    return []interface{}{result}
}

func flattenArmNotificationHubMpnsCredential(input *notificationhubs.MpnsCredential) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if mpnsCredentialProperties := input.MpnsCredentialProperties; mpnsCredentialProperties != nil {
        if certificateKey := mpnsCredentialProperties.CertificateKey; certificateKey != nil {
            result["certificate_key"] = *certificateKey
        }
        if mpnsCertificate := mpnsCredentialProperties.MpnsCertificate; mpnsCertificate != nil {
            result["mpns_certificate"] = *mpnsCertificate
        }
        if thumbprint := mpnsCredentialProperties.Thumbprint; thumbprint != nil {
            result["thumbprint"] = *thumbprint
        }
    }

    return []interface{}{result}
}

func flattenArmNotificationHubWnsCredential(input *notificationhubs.WnsCredential) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if wnsCredentialProperties := input.WnsCredentialProperties; wnsCredentialProperties != nil {
        if packageSid := wnsCredentialProperties.PackageSid; packageSid != nil {
            result["package_sid"] = *packageSid
        }
        if secretKey := wnsCredentialProperties.SecretKey; secretKey != nil {
            result["secret_key"] = *secretKey
        }
        if windowsLiveEndpoint := wnsCredentialProperties.WindowsLiveEndpoint; windowsLiveEndpoint != nil {
            result["windows_live_endpoint"] = *windowsLiveEndpoint
        }
    }

    return []interface{}{result}
}

func flattenArmNotificationHubSku(input *notificationhubs.Sku) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    result["name"] = string(input.Name)
    if capacity := input.Capacity; capacity != nil {
        result["capacity"] = *capacity
    }
    if family := input.Family; family != nil {
        result["family"] = *family
    }
    if size := input.Size; size != nil {
        result["size"] = *size
    }
    if tier := input.Tier; tier != nil {
        result["tier"] = *tier
    }

    return []interface{}{result}
}

func flattenArmNotificationHub(input *[]notificationhubs.) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        result := string(item)
        results = append(results, result)
    }

    return results
}