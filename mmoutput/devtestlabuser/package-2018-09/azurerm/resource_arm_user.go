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
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "location": azure.SchemaLocation(),

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "identity": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "app_id": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "object_id": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "principal_id": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "principal_name": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "tenant_id": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },

            "secret_store": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "key_vault_id": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "key_vault_uri": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },

            "created_date": {
                Type: schema.TypeString,
                Computed: true,
            },

            "provisioning_state": {
                Type: schema.TypeString,
                Computed: true,
            },

            "type": {
                Type: schema.TypeString,
                Computed: true,
            },

            "unique_identifier": {
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
    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, name, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing User %q (Resource Group %q): %+v", name, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_user", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    identity := d.Get("identity").([]interface{})
    secretStore := d.Get("secret_store").([]interface{})
    t := d.Get("tags").(map[string]interface{})

    user := devtestlab.UserFragment{
        Location: utils.String(location),
        UserPropertiesFragment: &devtestlab.UserPropertiesFragment{
            Identity: expandArmUserUserIdentityFragment(identity),
            SecretStore: expandArmUserUserSecretStoreFragment(secretStore),
        },
        Tags: tags.Expand(t),
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, name, name, user)
    if err != nil {
        return fmt.Errorf("Error creating User %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of User %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, name, name)
    if err != nil {
        return fmt.Errorf("Error retrieving User %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read User %q (Resource Group %q) ID", name, resourceGroup)
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
    name := id.Path["labs"]
    name := id.Path["users"]

    resp, err := client.Get(ctx, resourceGroup, name, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] User %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading User %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if userPropertiesFragment := resp.UserPropertiesFragment; userPropertiesFragment != nil {
        d.Set("created_date", (userPropertiesFragment.CreatedDate).String())
        if err := d.Set("identity", flattenArmUserUserIdentityFragment(userPropertiesFragment.Identity)); err != nil {
            return fmt.Errorf("Error setting `identity`: %+v", err)
        }
        d.Set("provisioning_state", userPropertiesFragment.ProvisioningState)
        if err := d.Set("secret_store", flattenArmUserUserSecretStoreFragment(userPropertiesFragment.SecretStore)); err != nil {
            return fmt.Errorf("Error setting `secret_store`: %+v", err)
        }
        d.Set("unique_identifier", userPropertiesFragment.UniqueIdentifier)
    }
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}

func resourceArmUserUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).usersClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    identity := d.Get("identity").([]interface{})
    secretStore := d.Get("secret_store").([]interface{})
    t := d.Get("tags").(map[string]interface{})

    user := devtestlab.UserFragment{
        Location: utils.String(location),
        UserPropertiesFragment: &devtestlab.UserPropertiesFragment{
            Identity: expandArmUserUserIdentityFragment(identity),
            SecretStore: expandArmUserUserSecretStoreFragment(secretStore),
        },
        Tags: tags.Expand(t),
    }


    if _, err := client.Update(ctx, resourceGroup, name, name, user); err != nil {
        return fmt.Errorf("Error updating User %q (Resource Group %q): %+v", name, resourceGroup, err)
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
    name := id.Path["labs"]
    name := id.Path["users"]

    future, err := client.Delete(ctx, resourceGroup, name, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting User %q (Resource Group %q): %+v", name, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting User %q (Resource Group %q): %+v", name, resourceGroup, err)
        }
    }

    return nil
}

func expandArmUserUserIdentityFragment(input []interface{}) *devtestlab.UserIdentityFragment {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    principalName := v["principal_name"].(string)
    principalId := v["principal_id"].(string)
    tenantId := v["tenant_id"].(string)
    objectId := v["object_id"].(string)
    appId := v["app_id"].(string)

    result := devtestlab.UserIdentityFragment{
        AppID: utils.String(appId),
        ObjectID: utils.String(objectId),
        PrincipalID: utils.String(principalId),
        PrincipalName: utils.String(principalName),
        TenantID: utils.String(tenantId),
    }
    return &result
}

func expandArmUserUserSecretStoreFragment(input []interface{}) *devtestlab.UserSecretStoreFragment {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    keyVaultUri := v["key_vault_uri"].(string)
    keyVaultId := v["key_vault_id"].(string)

    result := devtestlab.UserSecretStoreFragment{
        KeyVaultID: utils.String(keyVaultId),
        KeyVaultUri: utils.String(keyVaultUri),
    }
    return &result
}


func flattenArmUserUserIdentityFragment(input *devtestlab.UserIdentityFragment) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if appId := input.AppID; appId != nil {
        result["app_id"] = *appId
    }
    if objectId := input.ObjectID; objectId != nil {
        result["object_id"] = *objectId
    }
    if principalId := input.PrincipalID; principalId != nil {
        result["principal_id"] = *principalId
    }
    if principalName := input.PrincipalName; principalName != nil {
        result["principal_name"] = *principalName
    }
    if tenantId := input.TenantID; tenantId != nil {
        result["tenant_id"] = *tenantId
    }

    return []interface{}{result}
}

func flattenArmUserUserSecretStoreFragment(input *devtestlab.UserSecretStoreFragment) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if keyVaultId := input.KeyVaultID; keyVaultId != nil {
        result["key_vault_id"] = *keyVaultId
    }
    if keyVaultUri := input.KeyVaultUri; keyVaultUri != nil {
        result["key_vault_uri"] = *keyVaultUri
    }

    return []interface{}{result}
}
