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
        Create: resourceArmUserCreate,
        Read: resourceArmUserRead,
        Update: resourceArmUserUpdate,
        Delete: resourceArmUserDelete,

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

            "lab_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "unique_identifier": {
                Type: schema.TypeString,
                Optional: true,
            },

            "email": {
                Type: schema.TypeString,
                Computed: true,
            },

            "family_name": {
                Type: schema.TypeString,
                Computed: true,
            },

            "given_name": {
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

            "provisioning_state": {
                Type: schema.TypeString,
                Computed: true,
            },

            "tenant_id": {
                Type: schema.TypeString,
                Computed: true,
            },

            "total_usage": {
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

func resourceArmUserCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).usersClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    labAccountName := d.Get("lab_account_name").(string)
    labName := d.Get("lab_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, labAccountName, labName, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing User %q (Lab Name %q / Lab Account Name %q / Resource Group %q): %+v", name, labName, labAccountName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_user", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    uniqueIdentifier := d.Get("unique_identifier").(string)
    t := d.Get("tags").(map[string]interface{})

    user := labservices.User{
        Location: utils.String(location),
        UserProperties: &labservices.UserProperties{
            UniqueIdentifier: utils.String(uniqueIdentifier),
        },
        Tags: tags.Expand(t),
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroup, labAccountName, labName, name, user); err != nil {
        return fmt.Errorf("Error creating User %q (Lab Name %q / Lab Account Name %q / Resource Group %q): %+v", name, labName, labAccountName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, labAccountName, labName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving User %q (Lab Name %q / Lab Account Name %q / Resource Group %q): %+v", name, labName, labAccountName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read User %q (Lab Name %q / Lab Account Name %q / Resource Group %q) ID", name, labName, labAccountName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmUserRead(d, meta)
}

func resourceArmUserRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).usersClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    labAccountName := id.Path["labaccounts"]
    labName := id.Path["labs"]
    name := id.Path["users"]

    resp, err := client.Get(ctx, resourceGroup, labAccountName, labName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] User %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading User %q (Lab Name %q / Lab Account Name %q / Resource Group %q): %+v", name, labName, labAccountName, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if userProperties := resp.UserProperties; userProperties != nil {
        d.Set("email", userProperties.Email)
        d.Set("family_name", userProperties.FamilyName)
        d.Set("given_name", userProperties.GivenName)
        if err := d.Set("latest_operation_result", flattenArmUserLatestOperationResult(userProperties.LatestOperationResult)); err != nil {
            return fmt.Errorf("Error setting `latest_operation_result`: %+v", err)
        }
        d.Set("provisioning_state", userProperties.ProvisioningState)
        d.Set("tenant_id", userProperties.TenantID)
        d.Set("total_usage", userProperties.TotalUsage)
        d.Set("unique_identifier", userProperties.UniqueIdentifier)
    }
    d.Set("lab_account_name", labAccountName)
    d.Set("lab_name", labName)
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}

func resourceArmUserUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).usersClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    labAccountName := d.Get("lab_account_name").(string)
    labName := d.Get("lab_name").(string)
    uniqueIdentifier := d.Get("unique_identifier").(string)
    t := d.Get("tags").(map[string]interface{})

    user := labservices.User{
        Location: utils.String(location),
        UserProperties: &labservices.UserProperties{
            UniqueIdentifier: utils.String(uniqueIdentifier),
        },
        Tags: tags.Expand(t),
    }


    if _, err := client.Update(ctx, resourceGroup, labAccountName, labName, name, user); err != nil {
        return fmt.Errorf("Error updating User %q (Lab Name %q / Lab Account Name %q / Resource Group %q): %+v", name, labName, labAccountName, resourceGroup, err)
    }

    return resourceArmUserRead(d, meta)
}

func resourceArmUserDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).usersClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    labAccountName := id.Path["labaccounts"]
    labName := id.Path["labs"]
    name := id.Path["users"]

    future, err := client.Delete(ctx, resourceGroup, labAccountName, labName, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting User %q (Lab Name %q / Lab Account Name %q / Resource Group %q): %+v", name, labName, labAccountName, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting User %q (Lab Name %q / Lab Account Name %q / Resource Group %q): %+v", name, labName, labAccountName, resourceGroup, err)
        }
    }

    return nil
}


func flattenArmUserLatestOperationResult(input *labservices.LatestOperationResult) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})


    return []interface{}{result}
}
