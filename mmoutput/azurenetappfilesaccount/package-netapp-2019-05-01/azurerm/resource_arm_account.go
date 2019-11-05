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

            "active_directories": {
                Type: schema.TypeList,
                Optional: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "active_directory_id": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "dns": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "domain": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "organizational_unit": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "password": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "smb_server_name": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "status": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "username": {
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
    activeDirectories := d.Get("active_directories").([]interface{})
    t := d.Get("tags").(map[string]interface{})

    body := azurenetappfiles.NetAppAccount{
        Location: utils.String(location),
        AccountProperties: &azurenetappfiles.AccountProperties{
            ActiveDirectories: expandArmAccountActiveDirectory(activeDirectories),
        },
        Tags: tags.Expand(t),
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, accountName, body)
    if err != nil {
        return fmt.Errorf("Error creating Account (Account Name %q / Resource Group %q): %+v", accountName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Account (Account Name %q / Resource Group %q): %+v", accountName, resourceGroup, err)
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
    accountName := id.Path["netAppAccounts"]

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
    d.Set("account_name", accountName)
    if accountProperties := resp.AccountProperties; accountProperties != nil {
        if err := d.Set("active_directories", flattenArmAccountActiveDirectory(accountProperties.ActiveDirectories)); err != nil {
            return fmt.Errorf("Error setting `active_directories`: %+v", err)
        }
        d.Set("provisioning_state", accountProperties.ProvisioningState)
    }
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}

func resourceArmAccountUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).accountsClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    accountName := d.Get("account_name").(string)
    activeDirectories := d.Get("active_directories").([]interface{})
    t := d.Get("tags").(map[string]interface{})

    body := azurenetappfiles.NetAppAccount{
        Location: utils.String(location),
        AccountProperties: &azurenetappfiles.AccountProperties{
            ActiveDirectories: expandArmAccountActiveDirectory(activeDirectories),
        },
        Tags: tags.Expand(t),
    }


    if _, err := client.Update(ctx, resourceGroup, accountName, body); err != nil {
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
    accountName := id.Path["netAppAccounts"]

    future, err := client.Delete(ctx, resourceGroup, accountName)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Account (Account Name %q / Resource Group %q): %+v", accountName, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Account (Account Name %q / Resource Group %q): %+v", accountName, resourceGroup, err)
        }
    }

    return nil
}

func expandArmAccountActiveDirectory(input []interface{}) *[]azurenetappfiles.ActiveDirectory {
    results := make([]azurenetappfiles.ActiveDirectory, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        activeDirectoryId := v["active_directory_id"].(string)
        username := v["username"].(string)
        password := v["password"].(string)
        domain := v["domain"].(string)
        dns := v["dns"].(string)
        status := v["status"].(string)
        smbServerName := v["smb_server_name"].(string)
        organizationalUnit := v["organizational_unit"].(string)

        result := azurenetappfiles.ActiveDirectory{
            ActiveDirectoryID: utils.String(activeDirectoryId),
            Dns: utils.String(dns),
            Domain: utils.String(domain),
            OrganizationalUnit: utils.String(organizationalUnit),
            Password: utils.String(password),
            SmbServerName: utils.String(smbServerName),
            Status: utils.String(status),
            Username: utils.String(username),
        }

        results = append(results, result)
    }
    return &results
}


func flattenArmAccountActiveDirectory(input *[]azurenetappfiles.ActiveDirectory) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        if activeDirectoryId := item.ActiveDirectoryID; activeDirectoryId != nil {
            v["active_directory_id"] = *activeDirectoryId
        }
        if dns := item.Dns; dns != nil {
            v["dns"] = *dns
        }
        if domain := item.Domain; domain != nil {
            v["domain"] = *domain
        }
        if organizationalUnit := item.OrganizationalUnit; organizationalUnit != nil {
            v["organizational_unit"] = *organizationalUnit
        }
        if password := item.Password; password != nil {
            v["password"] = *password
        }
        if smbServerName := item.SmbServerName; smbServerName != nil {
            v["smb_server_name"] = *smbServerName
        }
        if status := item.Status; status != nil {
            v["status"] = *status
        }
        if username := item.Username; username != nil {
            v["username"] = *username
        }

        results = append(results, v)
    }

    return results
}