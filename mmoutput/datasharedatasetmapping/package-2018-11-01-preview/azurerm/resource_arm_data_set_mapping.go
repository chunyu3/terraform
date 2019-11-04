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



func resourceArmDataSetMapping() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmDataSetMappingCreateUpdate,
        Read: resourceArmDataSetMappingRead,
        Update: resourceArmDataSetMappingCreateUpdate,
        Delete: resourceArmDataSetMappingDelete,

        Importer: &schema.ResourceImporter{
            State: schema.ImportStatePassthrough,
        },


        Schema: map[string]*schema.Schema{
            "name": {
                Type: schema.TypeString,
                Computed: true,
            },

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "account_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "data_set_mapping_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "share_subscription_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "type": {
                Type: schema.TypeString,
                Computed: true,
            },
        },
    }
}

func resourceArmDataSetMappingCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).dataSetMappingsClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    accountName := d.Get("account_name").(string)
    dataSetMappingName := d.Get("data_set_mapping_name").(string)
    shareSubscriptionName := d.Get("share_subscription_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, accountName, shareSubscriptionName, dataSetMappingName)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Data Set Mapping (Data Set Mapping Name %q / Share Subscription Name %q / Account Name %q / Resource Group %q): %+v", dataSetMappingName, shareSubscriptionName, accountName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_data_set_mapping", *existing.ID)
        }
    }


    dataSetMapping := datashare.DataSetMapping{
    }


    if _, err := client.Create(ctx, resourceGroup, accountName, shareSubscriptionName, dataSetMappingName, dataSetMapping); err != nil {
        return fmt.Errorf("Error creating Data Set Mapping (Data Set Mapping Name %q / Share Subscription Name %q / Account Name %q / Resource Group %q): %+v", dataSetMappingName, shareSubscriptionName, accountName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, accountName, shareSubscriptionName, dataSetMappingName)
    if err != nil {
        return fmt.Errorf("Error retrieving Data Set Mapping (Data Set Mapping Name %q / Share Subscription Name %q / Account Name %q / Resource Group %q): %+v", dataSetMappingName, shareSubscriptionName, accountName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Data Set Mapping (Data Set Mapping Name %q / Share Subscription Name %q / Account Name %q / Resource Group %q) ID", dataSetMappingName, shareSubscriptionName, accountName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmDataSetMappingRead(d, meta)
}

func resourceArmDataSetMappingRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).dataSetMappingsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    accountName := id.Path["accounts"]
    shareSubscriptionName := id.Path["shareSubscriptions"]
    dataSetMappingName := id.Path["dataSetMappings"]

    resp, err := client.Get(ctx, resourceGroup, accountName, shareSubscriptionName, dataSetMappingName)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Data Set Mapping %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Data Set Mapping (Data Set Mapping Name %q / Share Subscription Name %q / Account Name %q / Resource Group %q): %+v", dataSetMappingName, shareSubscriptionName, accountName, resourceGroup, err)
    }


    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    d.Set("account_name", accountName)
    d.Set("data_set_mapping_name", dataSetMappingName)
    d.Set("share_subscription_name", shareSubscriptionName)
    d.Set("type", resp.Type)

    return nil
}


func resourceArmDataSetMappingDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).dataSetMappingsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    accountName := id.Path["accounts"]
    shareSubscriptionName := id.Path["shareSubscriptions"]
    dataSetMappingName := id.Path["dataSetMappings"]

    if _, err := client.Delete(ctx, resourceGroup, accountName, shareSubscriptionName, dataSetMappingName); err != nil {
        return fmt.Errorf("Error deleting Data Set Mapping (Data Set Mapping Name %q / Share Subscription Name %q / Account Name %q / Resource Group %q): %+v", dataSetMappingName, shareSubscriptionName, accountName, resourceGroup, err)
    }

    return nil
}