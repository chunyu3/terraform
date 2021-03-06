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



func resourceArmLab() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmLabCreate,
        Read: resourceArmLabRead,
        Update: resourceArmLabUpdate,
        Delete: resourceArmLabDelete,

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

            "email_addresses": {
                Type: schema.TypeList,
                Required: true,
                ForceNew: true,
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
            },

            "lab_account_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "max_users_in_lab": {
                Type: schema.TypeInt,
                Optional: true,
            },

            "unique_identifier": {
                Type: schema.TypeString,
                Optional: true,
            },

            "usage_quota": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validateIso8601Duration(),
            },

            "user_access_mode": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(labservices.Restricted),
                    string(labservices.Open),
                }, false),
                Default: string(labservices.Restricted),
            },

            "created_by_object_id": {
                Type: schema.TypeString,
                Computed: true,
            },

            "created_by_user_principal_name": {
                Type: schema.TypeString,
                Computed: true,
            },

            "created_date": {
                Type: schema.TypeString,
                Computed: true,
            },

            "invitation_code": {
                Type: schema.TypeString,
                Computed: true,
            },

            "latest_operation_result": {
                Type: schema.TypeList,
                Computed: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "error_code": {
                            Type: schema.TypeString,
                            Computed: true,
                        },
                        "error_message": {
                            Type: schema.TypeString,
                            Computed: true,
                        },
                        "http_method": {
                            Type: schema.TypeString,
                            Computed: true,
                        },
                        "operation_url": {
                            Type: schema.TypeString,
                            Computed: true,
                        },
                        "request_uri": {
                            Type: schema.TypeString,
                            Computed: true,
                        },
                        "status": {
                            Type: schema.TypeString,
                            Computed: true,
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

            "user_quota": {
                Type: schema.TypeInt,
                Computed: true,
            },

            "tags": tags.Schema(),
        },
    }
}

func resourceArmLabCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).labsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    labAccountName := d.Get("lab_account_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, labAccountName, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Lab %q (Lab Account Name %q / Resource Group %q): %+v", name, labAccountName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_lab", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    emailAddresses := d.Get("email_addresses").([]interface{})
    maxUsersInLab := d.Get("max_users_in_lab").(int)
    uniqueIdentifier := d.Get("unique_identifier").(string)
    usageQuota := d.Get("usage_quota").(string)
    userAccessMode := d.Get("user_access_mode").(string)
    t := d.Get("tags").(map[string]interface{})

    lab := labservices.LabFragment{
        EmailAddresses: utils.ExpandStringSlice(emailAddresses),
        Location: utils.String(location),
        LabPropertiesFragment: &labservices.LabPropertiesFragment{
            MaxUsersInLab: utils.Int32(int32(maxUsersInLab)),
            UniqueIdentifier: utils.String(uniqueIdentifier),
            UsageQuota: utils.String(usageQuota),
            UserAccessMode: labservices.LabUserAccessMode(userAccessMode),
        },
        Tags: tags.Expand(t),
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroup, labAccountName, name, lab); err != nil {
        return fmt.Errorf("Error creating Lab %q (Lab Account Name %q / Resource Group %q): %+v", name, labAccountName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, labAccountName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Lab %q (Lab Account Name %q / Resource Group %q): %+v", name, labAccountName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Lab %q (Lab Account Name %q / Resource Group %q) ID", name, labAccountName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmLabRead(d, meta)
}

func resourceArmLabRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).labsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    labAccountName := id.Path["labaccounts"]
    name := id.Path["labs"]

    resp, err := client.Get(ctx, resourceGroup, labAccountName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Lab %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Lab %q (Lab Account Name %q / Resource Group %q): %+v", name, labAccountName, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if labPropertiesFragment := resp.LabPropertiesFragment; labPropertiesFragment != nil {
        d.Set("created_by_object_id", labPropertiesFragment.CreatedByObjectID)
        d.Set("created_by_user_principal_name", labPropertiesFragment.CreatedByUserPrincipalName)
        d.Set("created_date", (labPropertiesFragment.CreatedDate).String())
        d.Set("invitation_code", labPropertiesFragment.InvitationCode)
        if err := d.Set("latest_operation_result", flattenArmLabLatestOperationResult(labPropertiesFragment.LatestOperationResult)); err != nil {
            return fmt.Errorf("Error setting `latest_operation_result`: %+v", err)
        }
        d.Set("max_users_in_lab", int(*labPropertiesFragment.MaxUsersInLab))
        d.Set("provisioning_state", labPropertiesFragment.ProvisioningState)
        d.Set("unique_identifier", labPropertiesFragment.UniqueIdentifier)
        d.Set("usage_quota", labPropertiesFragment.UsageQuota)
        d.Set("user_access_mode", string(labPropertiesFragment.UserAccessMode))
        d.Set("user_quota", int(*labPropertiesFragment.UserQuota))
    }
    d.Set("lab_account_name", labAccountName)
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}

func resourceArmLabUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).labsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    emailAddresses := d.Get("email_addresses").([]interface{})
    labAccountName := d.Get("lab_account_name").(string)
    maxUsersInLab := d.Get("max_users_in_lab").(int)
    uniqueIdentifier := d.Get("unique_identifier").(string)
    usageQuota := d.Get("usage_quota").(string)
    userAccessMode := d.Get("user_access_mode").(string)
    t := d.Get("tags").(map[string]interface{})

    lab := labservices.LabFragment{
        EmailAddresses: utils.ExpandStringSlice(emailAddresses),
        LabPropertiesFragment: &labservices.LabPropertiesFragment{
            MaxUsersInLab: utils.Int32(int32(maxUsersInLab)),
            UniqueIdentifier: utils.String(uniqueIdentifier),
            UsageQuota: utils.String(usageQuota),
            UserAccessMode: labservices.LabUserAccessMode(userAccessMode),
        },
        Tags: tags.Expand(t),
    }


    if _, err := client.Update(ctx, resourceGroup, labAccountName, name, lab); err != nil {
        return fmt.Errorf("Error updating Lab %q (Lab Account Name %q / Resource Group %q): %+v", name, labAccountName, resourceGroup, err)
    }

    return resourceArmLabRead(d, meta)
}

func resourceArmLabDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).labsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    labAccountName := id.Path["labaccounts"]
    name := id.Path["labs"]

    future, err := client.Delete(ctx, resourceGroup, labAccountName, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Lab %q (Lab Account Name %q / Resource Group %q): %+v", name, labAccountName, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Lab %q (Lab Account Name %q / Resource Group %q): %+v", name, labAccountName, resourceGroup, err)
        }
    }

    return nil
}


func flattenArmLabLatestOperationResult(input *labservices.LatestOperationResult) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if errorCode := input.ErrorCode; errorCode != nil {
        result["error_code"] = *errorCode
    }
    if errorMessage := input.ErrorMessage; errorMessage != nil {
        result["error_message"] = *errorMessage
    }
    if httpMethod := input.HTTPMethod; httpMethod != nil {
        result["http_method"] = *httpMethod
    }
    if operationUrl := input.OperationURL; operationUrl != nil {
        result["operation_url"] = *operationUrl
    }
    if requestUri := input.RequestURI; requestUri != nil {
        result["request_uri"] = *requestUri
    }
    if status := input.Status; status != nil {
        result["status"] = *status
    }

    return []interface{}{result}
}
