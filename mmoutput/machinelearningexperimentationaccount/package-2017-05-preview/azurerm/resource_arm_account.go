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



func resourceArmAccount() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmAccountCreate,
        Read: resourceArmAccountRead,
        Update: resourceArmAccountUpdate,
        Delete: resourceArmAccountDelete,

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

            "account_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "key_vault_id": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "storage_account": {
                Type: schema.TypeList,
                Required: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "access_key": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "storage_account_id": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                    },
                },
            },

            "vso_account_id": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "description": {
                Type: schema.TypeString,
                Optional: true,
            },

            "friendly_name": {
                Type: schema.TypeString,
                Optional: true,
            },

            "seats": {
                Type: schema.TypeString,
                Optional: true,
            },

            "account_id": {
                Type: schema.TypeString,
                Computed: true,
            },

            "creation_date": {
                Type: schema.TypeString,
                Computed: true,
            },

            "discovery_uri": {
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

            "tags": tags.Schema(),
        },
    }
}

func resourceArmAccountCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).accountsClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    accountName := d.Get("account_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, accountName)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Account (Account Name %q / Resource Group %q): %+v", accountName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_account", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    description := d.Get("description").(string)
    friendlyName := d.Get("friendly_name").(string)
    keyVaultId := d.Get("key_vault_id").(string)
    seats := d.Get("seats").(string)
    storageAccount := d.Get("storage_account").([]interface{})
    vsoAccountId := d.Get("vso_account_id").(string)
    t := d.Get("tags").(map[string]interface{})

    parameters := machinelearningexperimentation.Account{
        Location: utils.String(location),
        AccountProperties: &machinelearningexperimentation.AccountProperties{
            Description: utils.String(description),
            FriendlyName: utils.String(friendlyName),
            KeyVaultID: utils.String(keyVaultId),
            Seats: utils.String(seats),
            StorageAccount: expandArmAccountStorageAccountProperties(storageAccount),
            VsoAccountID: utils.String(vsoAccountId),
        },
        Tags: tags.Expand(t),
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroup, accountName, parameters); err != nil {
        return fmt.Errorf("Error creating Account (Account Name %q / Resource Group %q): %+v", accountName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, accountName)
    if err != nil {
        return fmt.Errorf("Error retrieving Account (Account Name %q / Resource Group %q): %+v", accountName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Account (Account Name %q / Resource Group %q) ID", accountName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmAccountRead(d, meta)
}

func resourceArmAccountRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).accountsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    accountName := id.Path["accounts"]

    resp, err := client.Get(ctx, resourceGroup, accountName)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Account %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Account (Account Name %q / Resource Group %q): %+v", accountName, resourceGroup, err)
    }


    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if accountProperties := resp.AccountProperties; accountProperties != nil {
        d.Set("account_id", accountProperties.AccountID)
        d.Set("creation_date", (accountProperties.CreationDate).String())
        d.Set("description", accountProperties.Description)
        d.Set("discovery_uri", accountProperties.DiscoveryUri)
        d.Set("friendly_name", accountProperties.FriendlyName)
        d.Set("key_vault_id", accountProperties.KeyVaultID)
        d.Set("provisioning_state", string(accountProperties.ProvisioningState))
        d.Set("seats", accountProperties.Seats)
        if err := d.Set("storage_account", flattenArmAccountStorageAccountProperties(accountProperties.StorageAccount)); err != nil {
            return fmt.Errorf("Error setting `storage_account`: %+v", err)
        }
        d.Set("vso_account_id", accountProperties.VsoAccountID)
    }
    d.Set("account_name", accountName)
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}

func resourceArmAccountUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).accountsClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    accountName := d.Get("account_name").(string)
    description := d.Get("description").(string)
    friendlyName := d.Get("friendly_name").(string)
    keyVaultId := d.Get("key_vault_id").(string)
    seats := d.Get("seats").(string)
    storageAccount := d.Get("storage_account").([]interface{})
    vsoAccountId := d.Get("vso_account_id").(string)
    t := d.Get("tags").(map[string]interface{})

    parameters := machinelearningexperimentation.Account{
        Location: utils.String(location),
        AccountProperties: &machinelearningexperimentation.AccountProperties{
            Description: utils.String(description),
            FriendlyName: utils.String(friendlyName),
            KeyVaultID: utils.String(keyVaultId),
            Seats: utils.String(seats),
            StorageAccount: expandArmAccountStorageAccountProperties(storageAccount),
            VsoAccountID: utils.String(vsoAccountId),
        },
        Tags: tags.Expand(t),
    }


    if _, err := client.Update(ctx, resourceGroup, accountName, parameters); err != nil {
        return fmt.Errorf("Error updating Account (Account Name %q / Resource Group %q): %+v", accountName, resourceGroup, err)
    }

    return resourceArmAccountRead(d, meta)
}

func resourceArmAccountDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).accountsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    accountName := id.Path["accounts"]

    if _, err := client.Delete(ctx, resourceGroup, accountName); err != nil {
        return fmt.Errorf("Error deleting Account (Account Name %q / Resource Group %q): %+v", accountName, resourceGroup, err)
    }

    return nil
}

func expandArmAccountStorageAccountProperties(input []interface{}) *machinelearningexperimentation.StorageAccountProperties {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    storageAccountId := v["storage_account_id"].(string)
    accessKey := v["access_key"].(string)

    result := machinelearningexperimentation.StorageAccountProperties{
        AccessKey: utils.String(accessKey),
        StorageAccountID: utils.String(storageAccountId),
    }
    return &result
}


func flattenArmAccountStorageAccountProperties(input *machinelearningexperimentation.StorageAccountProperties) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if accessKey := input.AccessKey; accessKey != nil {
        result["access_key"] = *accessKey
    }
    if storageAccountId := input.StorageAccountID; storageAccountId != nil {
        result["storage_account_id"] = *storageAccountId
    }

    return []interface{}{result}
}