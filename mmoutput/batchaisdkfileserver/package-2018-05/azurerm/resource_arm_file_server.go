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



func resourceArmFileServer() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmFileServerCreateUpdate,
        Read: resourceArmFileServerRead,
        Update: resourceArmFileServerCreateUpdate,
        Delete: resourceArmFileServerDelete,

        Importer: &schema.ResourceImporter{
            State: schema.ImportStatePassthrough,
        },


        Schema: map[string]*schema.Schema{
            "name": {
                Type: schema.TypeString,
                Computed: true,
            },

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "data_disks": {
                Type: schema.TypeList,
                Required: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "disk_count": {
                            Type: schema.TypeInt,
                            Required: true,
                        },
                        "disk_size_in_gb": {
                            Type: schema.TypeInt,
                            Required: true,
                        },
                        "storage_account_type": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(batchaisdk.Standard_LRS),
                                string(batchaisdk.Premium_LRS),
                            }, false),
                        },
                        "caching_type": {
                            Type: schema.TypeString,
                            Optional: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(batchaisdk.none),
                                string(batchaisdk.readonly),
                                string(batchaisdk.readwrite),
                            }, false),
                            Default: string(batchaisdk.none),
                        },
                    },
                },
            },

            "file_server_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "ssh_configuration": {
                Type: schema.TypeList,
                Required: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "user_account_settings": {
                            Type: schema.TypeList,
                            Required: true,
                            MaxItems: 1,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "admin_user_name": {
                                        Type: schema.TypeString,
                                        Required: true,
                                        ValidateFunc: validate.NoEmptyStrings,
                                    },
                                    "admin_user_password": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                    "admin_user_ssh_public_key": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                },
                            },
                        },
                        "public_ips_to_allow": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                    },
                },
            },

            "vm_size": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "workspace_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "subnet": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "id": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                    },
                },
            },

            "creation_time": {
                Type: schema.TypeString,
                Computed: true,
            },

            "mount_settings": {
                Type: schema.TypeList,
                Computed: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "file_server_internal_ip": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "file_server_public_ip": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "mount_point": {
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

            "provisioning_state_transition_time": {
                Type: schema.TypeString,
                Computed: true,
            },

            "type": {
                Type: schema.TypeString,
                Computed: true,
            },
        },
    }
}

func resourceArmFileServerCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).fileServersClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    fileServerName := d.Get("file_server_name").(string)
    workspaceName := d.Get("workspace_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, workspaceName, fileServerName)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing File Server (File Server Name %q / Workspace Name %q / Resource Group %q): %+v", fileServerName, workspaceName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_file_server", *existing.ID)
        }
    }

    dataDisks := d.Get("data_disks").([]interface{})
    sshConfiguration := d.Get("ssh_configuration").([]interface{})
    subnet := d.Get("subnet").([]interface{})
    vmSize := d.Get("vm_size").(string)

    parameters := batchaisdk.FileServerCreateParameters{
        FileServerBaseProperties: &batchaisdk.FileServerBaseProperties{
            DataDisks: expandArmFileServerDataDisks(dataDisks),
            SshConfiguration: expandArmFileServerSshConfiguration(sshConfiguration),
            Subnet: expandArmFileServerResourceId(subnet),
            VmSize: utils.String(vmSize),
        },
    }


    future, err := client.Create(ctx, resourceGroup, workspaceName, fileServerName, parameters)
    if err != nil {
        return fmt.Errorf("Error creating File Server (File Server Name %q / Workspace Name %q / Resource Group %q): %+v", fileServerName, workspaceName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of File Server (File Server Name %q / Workspace Name %q / Resource Group %q): %+v", fileServerName, workspaceName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, workspaceName, fileServerName)
    if err != nil {
        return fmt.Errorf("Error retrieving File Server (File Server Name %q / Workspace Name %q / Resource Group %q): %+v", fileServerName, workspaceName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read File Server (File Server Name %q / Workspace Name %q / Resource Group %q) ID", fileServerName, workspaceName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmFileServerRead(d, meta)
}

func resourceArmFileServerRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).fileServersClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    workspaceName := id.Path["workspaces"]
    fileServerName := id.Path["fileServers"]

    resp, err := client.Get(ctx, resourceGroup, workspaceName, fileServerName)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] File Server %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading File Server (File Server Name %q / Workspace Name %q / Resource Group %q): %+v", fileServerName, workspaceName, resourceGroup, err)
    }


    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if fileServerBaseProperties := resp.FileServerBaseProperties; fileServerBaseProperties != nil {
        d.Set("creation_time", (fileServerBaseProperties.CreationTime).String())
        if err := d.Set("data_disks", flattenArmFileServerDataDisks(fileServerBaseProperties.DataDisks)); err != nil {
            return fmt.Errorf("Error setting `data_disks`: %+v", err)
        }
        if err := d.Set("mount_settings", flattenArmFileServerMountSettings(fileServerBaseProperties.MountSettings)); err != nil {
            return fmt.Errorf("Error setting `mount_settings`: %+v", err)
        }
        d.Set("provisioning_state", string(fileServerBaseProperties.ProvisioningState))
        d.Set("provisioning_state_transition_time", (fileServerBaseProperties.ProvisioningStateTransitionTime).String())
        if err := d.Set("ssh_configuration", flattenArmFileServerSshConfiguration(fileServerBaseProperties.SshConfiguration)); err != nil {
            return fmt.Errorf("Error setting `ssh_configuration`: %+v", err)
        }
        if err := d.Set("subnet", flattenArmFileServerResourceId(fileServerBaseProperties.Subnet)); err != nil {
            return fmt.Errorf("Error setting `subnet`: %+v", err)
        }
        d.Set("vm_size", fileServerBaseProperties.VmSize)
    }
    d.Set("file_server_name", fileServerName)
    d.Set("type", resp.Type)
    d.Set("workspace_name", workspaceName)

    return nil
}


func resourceArmFileServerDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).fileServersClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    workspaceName := id.Path["workspaces"]
    fileServerName := id.Path["fileServers"]

    future, err := client.Delete(ctx, resourceGroup, workspaceName, fileServerName)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting File Server (File Server Name %q / Workspace Name %q / Resource Group %q): %+v", fileServerName, workspaceName, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting File Server (File Server Name %q / Workspace Name %q / Resource Group %q): %+v", fileServerName, workspaceName, resourceGroup, err)
        }
    }

    return nil
}

func expandArmFileServerDataDisks(input []interface{}) *batchaisdk.DataDisks {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    diskSizeInGb := v["disk_size_in_gb"].(int)
    cachingType := v["caching_type"].(string)
    diskCount := v["disk_count"].(int)
    storageAccountType := v["storage_account_type"].(string)

    result := batchaisdk.DataDisks{
        CachingType: batchaisdk.CachingType(cachingType),
        DiskCount: utils.Int32(int32(diskCount)),
        DiskSizeInGb: utils.Int32(int32(diskSizeInGb)),
        StorageAccountType: batchaisdk.StorageAccountType(storageAccountType),
    }
    return &result
}

func expandArmFileServerSshConfiguration(input []interface{}) *batchaisdk.SshConfiguration {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    publicIpsToAllow := v["public_ips_to_allow"].([]interface{})
    userAccountSettings := v["user_account_settings"].([]interface{})

    result := batchaisdk.SshConfiguration{
        PublicIpsToAllow: utils.ExpandStringSlice(publicIpsToAllow),
        UserAccountSettings: expandArmFileServerUserAccountSettings(userAccountSettings),
    }
    return &result
}

func expandArmFileServerResourceId(input []interface{}) *batchaisdk.ResourceId {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    id := v["id"].(string)

    result := batchaisdk.ResourceId{
        ID: utils.String(id),
    }
    return &result
}

func expandArmFileServerUserAccountSettings(input []interface{}) *batchaisdk.UserAccountSettings {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    adminUserName := v["admin_user_name"].(string)
    adminUserSshPublicKey := v["admin_user_ssh_public_key"].(string)
    adminUserPassword := v["admin_user_password"].(string)

    result := batchaisdk.UserAccountSettings{
        AdminUserName: utils.String(adminUserName),
        AdminUserPassword: utils.String(adminUserPassword),
        AdminUserSshPublicKey: utils.String(adminUserSshPublicKey),
    }
    return &result
}


func flattenArmFileServerDataDisks(input *batchaisdk.DataDisks) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    result["caching_type"] = string(input.CachingType)
    if diskCount := input.DiskCount; diskCount != nil {
        result["disk_count"] = int(*diskCount)
    }
    if diskSizeInGb := input.DiskSizeInGb; diskSizeInGb != nil {
        result["disk_size_in_gb"] = int(*diskSizeInGb)
    }
    result["storage_account_type"] = string(input.StorageAccountType)

    return []interface{}{result}
}

func flattenArmFileServerMountSettings(input *batchaisdk.MountSettings) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})


    return []interface{}{result}
}

func flattenArmFileServerSshConfiguration(input *batchaisdk.SshConfiguration) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    result["public_ips_to_allow"] = utils.FlattenStringSlice(input.PublicIpsToAllow)
    result["user_account_settings"] = flattenArmFileServerUserAccountSettings(input.UserAccountSettings)

    return []interface{}{result}
}

func flattenArmFileServerResourceId(input *batchaisdk.ResourceId) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if id := input.ID; id != nil {
        result["id"] = *id
    }

    return []interface{}{result}
}

func flattenArmFileServerUserAccountSettings(input *batchaisdk.UserAccountSettings) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if adminUserName := input.AdminUserName; adminUserName != nil {
        result["admin_user_name"] = *adminUserName
    }
    if adminUserPassword := input.AdminUserPassword; adminUserPassword != nil {
        result["admin_user_password"] = *adminUserPassword
    }
    if adminUserSshPublicKey := input.AdminUserSshPublicKey; adminUserSshPublicKey != nil {
        result["admin_user_ssh_public_key"] = *adminUserSshPublicKey
    }

    return []interface{}{result}
}
