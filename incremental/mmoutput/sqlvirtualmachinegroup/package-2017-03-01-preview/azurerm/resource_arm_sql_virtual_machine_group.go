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



func resourceArmSqlVirtualMachineGroup() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmSqlVirtualMachineGroupCreate,
        Read: resourceArmSqlVirtualMachineGroupRead,
        Update: resourceArmSqlVirtualMachineGroupUpdate,
        Delete: resourceArmSqlVirtualMachineGroupDelete,

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

            "sql_image_offer": {
                Type: schema.TypeString,
                Optional: true,
            },

            "sql_image_sku": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(sqlvirtualmachine.Developer),
                    string(sqlvirtualmachine.Enterprise),
                }, false),
                Default: string(sqlvirtualmachine.Developer),
            },

            "wsfc_domain_profile": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "cluster_bootstrap_account": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "cluster_operator_account": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "domain_fqdn": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "file_share_witness_path": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "ou_path": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "sql_service_account": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "storage_account_primary_key": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "storage_account_url": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },

            "cluster_configuration": {
                Type: schema.TypeString,
                Computed: true,
            },

            "cluster_manager_type": {
                Type: schema.TypeString,
                Computed: true,
            },

            "provisioning_state": {
                Type: schema.TypeString,
                Computed: true,
            },

            "scale_type": {
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

func resourceArmSqlVirtualMachineGroupCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).sqlVirtualMachineGroupsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Sql Virtual Machine Group %q (Resource Group %q): %+v", name, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_sql_virtual_machine_group", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    sqlImageOffer := d.Get("sql_image_offer").(string)
    sqlImageSku := d.Get("sql_image_sku").(string)
    wsfcDomainProfile := d.Get("wsfc_domain_profile").([]interface{})
    t := d.Get("tags").(map[string]interface{})

    parameters := sqlvirtualmachine.GroupUpdate{
        Location: utils.String(location),
        GroupProperties: &sqlvirtualmachine.GroupProperties{
            SQLImageOffer: utils.String(sqlImageOffer),
            SQLImageSku: sqlvirtualmachine.SqlVmGroupImageSku(sqlImageSku),
            WsfcDomainProfile: expandArmSqlVirtualMachineGroupWsfcDomainProfile(wsfcDomainProfile),
        },
        Tags: tags.Expand(t),
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, name, parameters)
    if err != nil {
        return fmt.Errorf("Error creating Sql Virtual Machine Group %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Sql Virtual Machine Group %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Sql Virtual Machine Group %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Sql Virtual Machine Group %q (Resource Group %q) ID", name, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmSqlVirtualMachineGroupRead(d, meta)
}

func resourceArmSqlVirtualMachineGroupRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).sqlVirtualMachineGroupsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["sqlVirtualMachineGroups"]

    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Sql Virtual Machine Group %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Sql Virtual Machine Group %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if groupProperties := resp.GroupProperties; groupProperties != nil {
        d.Set("cluster_configuration", string(groupProperties.ClusterConfiguration))
        d.Set("cluster_manager_type", string(groupProperties.ClusterManagerType))
        d.Set("provisioning_state", groupProperties.ProvisioningState)
        d.Set("scale_type", string(groupProperties.ScaleType))
        d.Set("sql_image_offer", groupProperties.SQLImageOffer)
        d.Set("sql_image_sku", string(groupProperties.SQLImageSku))
        if err := d.Set("wsfc_domain_profile", flattenArmSqlVirtualMachineGroupWsfcDomainProfile(groupProperties.WsfcDomainProfile)); err != nil {
            return fmt.Errorf("Error setting `wsfc_domain_profile`: %+v", err)
        }
    }
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}

func resourceArmSqlVirtualMachineGroupUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).sqlVirtualMachineGroupsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    sqlImageOffer := d.Get("sql_image_offer").(string)
    sqlImageSku := d.Get("sql_image_sku").(string)
    wsfcDomainProfile := d.Get("wsfc_domain_profile").([]interface{})
    t := d.Get("tags").(map[string]interface{})

    parameters := sqlvirtualmachine.GroupUpdate{
        GroupProperties: &sqlvirtualmachine.GroupProperties{
            SQLImageOffer: utils.String(sqlImageOffer),
            SQLImageSku: sqlvirtualmachine.SqlVmGroupImageSku(sqlImageSku),
            WsfcDomainProfile: expandArmSqlVirtualMachineGroupWsfcDomainProfile(wsfcDomainProfile),
        },
        Tags: tags.Expand(t),
    }


    future, err := client.Update(ctx, resourceGroup, name, parameters)
    if err != nil {
        return fmt.Errorf("Error updating Sql Virtual Machine Group %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for update of Sql Virtual Machine Group %q (Resource Group %q): %+v", name, resourceGroup, err)
    }

    return resourceArmSqlVirtualMachineGroupRead(d, meta)
}

func resourceArmSqlVirtualMachineGroupDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).sqlVirtualMachineGroupsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["sqlVirtualMachineGroups"]

    future, err := client.Delete(ctx, resourceGroup, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Sql Virtual Machine Group %q (Resource Group %q): %+v", name, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Sql Virtual Machine Group %q (Resource Group %q): %+v", name, resourceGroup, err)
        }
    }

    return nil
}

func expandArmSqlVirtualMachineGroupWsfcDomainProfile(input []interface{}) *sqlvirtualmachine.WsfcDomainProfile {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    domainFqdn := v["domain_fqdn"].(string)
    ouPath := v["ou_path"].(string)
    clusterBootstrapAccount := v["cluster_bootstrap_account"].(string)
    clusterOperatorAccount := v["cluster_operator_account"].(string)
    sqlServiceAccount := v["sql_service_account"].(string)
    fileShareWitnessPath := v["file_share_witness_path"].(string)
    storageAccountUrl := v["storage_account_url"].(string)
    storageAccountPrimaryKey := v["storage_account_primary_key"].(string)

    result := sqlvirtualmachine.WsfcDomainProfile{
        ClusterBootstrapAccount: utils.String(clusterBootstrapAccount),
        ClusterOperatorAccount: utils.String(clusterOperatorAccount),
        DomainFqdn: utils.String(domainFqdn),
        FileShareWitnessPath: utils.String(fileShareWitnessPath),
        OuPath: utils.String(ouPath),
        SQLServiceAccount: utils.String(sqlServiceAccount),
        StorageAccountPrimaryKey: utils.String(storageAccountPrimaryKey),
        StorageAccountURL: utils.String(storageAccountUrl),
    }
    return &result
}


func flattenArmSqlVirtualMachineGroupWsfcDomainProfile(input *sqlvirtualmachine.WsfcDomainProfile) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if clusterBootstrapAccount := input.ClusterBootstrapAccount; clusterBootstrapAccount != nil {
        result["cluster_bootstrap_account"] = *clusterBootstrapAccount
    }
    if clusterOperatorAccount := input.ClusterOperatorAccount; clusterOperatorAccount != nil {
        result["cluster_operator_account"] = *clusterOperatorAccount
    }
    if domainFqdn := input.DomainFqdn; domainFqdn != nil {
        result["domain_fqdn"] = *domainFqdn
    }
    if fileShareWitnessPath := input.FileShareWitnessPath; fileShareWitnessPath != nil {
        result["file_share_witness_path"] = *fileShareWitnessPath
    }
    if ouPath := input.OuPath; ouPath != nil {
        result["ou_path"] = *ouPath
    }
    if sqlServiceAccount := input.SQLServiceAccount; sqlServiceAccount != nil {
        result["sql_service_account"] = *sqlServiceAccount
    }
    if storageAccountPrimaryKey := input.StorageAccountPrimaryKey; storageAccountPrimaryKey != nil {
        result["storage_account_primary_key"] = *storageAccountPrimaryKey
    }
    if storageAccountUrl := input.StorageAccountURL; storageAccountUrl != nil {
        result["storage_account_url"] = *storageAccountUrl
    }

    return []interface{}{result}
}
