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



func resourceArmBatchAccount() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmBatchAccountCreate,
        Read: resourceArmBatchAccountRead,
        Update: resourceArmBatchAccountUpdate,
        Delete: resourceArmBatchAccountDelete,

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

            "auto_storage": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "storage_account_id": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                    },
                },
            },

            "account_endpoint": {
                Type: schema.TypeString,
                Computed: true,
            },

            "active_job_and_job_schedule_quota": {
                Type: schema.TypeInt,
                Computed: true,
            },

            "core_quota": {
                Type: schema.TypeInt,
                Computed: true,
            },

            "pool_quota": {
                Type: schema.TypeInt,
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

func resourceArmBatchAccountCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).batchAccountClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    accountName := d.Get("account_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, accountName)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Batch Account (Account Name %q / Resource Group %q): %+v", accountName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_batch_account", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    autoStorage := d.Get("auto_storage").([]interface{})
    t := d.Get("tags").(map[string]interface{})

    parameters := batch.AccountCreateParameters{
        Location: utils.String(location),
        AccountBaseProperties: &batch.AccountBaseProperties{
            AutoStorage: expandArmBatchAccountAutoStorageBaseProperties(autoStorage),
        },
        Tags: tags.Expand(t),
    }


    future, err := client.Create(ctx, resourceGroup, accountName, parameters)
    if err != nil {
        return fmt.Errorf("Error creating Batch Account (Account Name %q / Resource Group %q): %+v", accountName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Batch Account (Account Name %q / Resource Group %q): %+v", accountName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, accountName)
    if err != nil {
        return fmt.Errorf("Error retrieving Batch Account (Account Name %q / Resource Group %q): %+v", accountName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Batch Account (Account Name %q / Resource Group %q) ID", accountName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmBatchAccountRead(d, meta)
}

func resourceArmBatchAccountRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).batchAccountClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    accountName := id.Path["batchAccounts"]

    resp, err := client.Get(ctx, resourceGroup, accountName)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Batch Account %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Batch Account (Account Name %q / Resource Group %q): %+v", accountName, resourceGroup, err)
    }


    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if accountBaseProperties := resp.AccountBaseProperties; accountBaseProperties != nil {
        d.Set("account_endpoint", accountBaseProperties.AccountEndpoint)
        d.Set("active_job_and_job_schedule_quota", int(*accountBaseProperties.ActiveJobAndJobScheduleQuota))
        if err := d.Set("auto_storage", flattenArmBatchAccountAutoStorageBaseProperties(accountBaseProperties.AutoStorage)); err != nil {
            return fmt.Errorf("Error setting `auto_storage`: %+v", err)
        }
        d.Set("core_quota", int(*accountBaseProperties.CoreQuota))
        d.Set("pool_quota", int(*accountBaseProperties.PoolQuota))
        d.Set("provisioning_state", string(accountBaseProperties.ProvisioningState))
    }
    d.Set("account_name", accountName)
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}

func resourceArmBatchAccountUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).batchAccountClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    accountName := d.Get("account_name").(string)
    autoStorage := d.Get("auto_storage").([]interface{})
    t := d.Get("tags").(map[string]interface{})

    parameters := batch.AccountCreateParameters{
        Location: utils.String(location),
        AccountBaseProperties: &batch.AccountBaseProperties{
            AutoStorage: expandArmBatchAccountAutoStorageBaseProperties(autoStorage),
        },
        Tags: tags.Expand(t),
    }


    if _, err := client.Update(ctx, resourceGroup, accountName, parameters); err != nil {
        return fmt.Errorf("Error updating Batch Account (Account Name %q / Resource Group %q): %+v", accountName, resourceGroup, err)
    }

    return resourceArmBatchAccountRead(d, meta)
}

func resourceArmBatchAccountDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).batchAccountClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    accountName := id.Path["batchAccounts"]

    future, err := client.Delete(ctx, resourceGroup, accountName)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Batch Account (Account Name %q / Resource Group %q): %+v", accountName, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Batch Account (Account Name %q / Resource Group %q): %+v", accountName, resourceGroup, err)
        }
    }

    return nil
}

func expandArmBatchAccountAutoStorageBaseProperties(input []interface{}) *batch.AutoStorageBaseProperties {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    storageAccountId := v["storage_account_id"].(string)

    result := batch.AutoStorageBaseProperties{
        StorageAccountID: utils.String(storageAccountId),
    }
    return &result
}


func flattenArmBatchAccountAutoStorageBaseProperties(input *batch.AutoStorageBaseProperties) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if storageAccountId := input.StorageAccountID; storageAccountId != nil {
        result["storage_account_id"] = *storageAccountId
    }

    return []interface{}{result}
}
