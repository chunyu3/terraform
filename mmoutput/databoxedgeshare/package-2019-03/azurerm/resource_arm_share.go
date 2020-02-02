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



func resourceArmShare() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmShareCreateUpdate,
        Read: resourceArmShareRead,
        Update: resourceArmShareCreateUpdate,
        Delete: resourceArmShareDelete,

        Importer: &schema.ResourceImporter{
            State: schema.ImportStatePassthrough,
        },


        Schema: map[string]*schema.Schema{
            "access_protocol": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(databoxedge.SMB),
                    string(databoxedge.NFS),
                }, false),
            },

            "device_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "monitoring_status": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(databoxedge.Enabled),
                    string(databoxedge.Disabled),
                }, false),
            },

            "name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "share_status": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(databoxedge.Online),
                    string(databoxedge.Offline),
                }, false),
            },

            "azure_container_info": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "container_name": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "data_format": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(databoxedge.BlockBlob),
                                string(databoxedge.PageBlob),
                                string(databoxedge.AzureFile),
                            }, false),
                        },
                        "storage_account_credential_id": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                    },
                },
            },

            "client_access_rights": {
                Type: schema.TypeList,
                Optional: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "access_permission": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(databoxedge.NoAccess),
                                string(databoxedge.ReadOnly),
                                string(databoxedge.ReadWrite),
                            }, false),
                        },
                        "client": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                    },
                },
            },

            "data_policy": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(databoxedge.Cloud),
                    string(databoxedge.Local),
                }, false),
                Default: string(databoxedge.Cloud),
            },

            "description": {
                Type: schema.TypeString,
                Optional: true,
            },

            "refresh_details": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "error_manifest_file": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "in_progress_refresh_job_id": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "last_completed_refresh_job_time_in_utc": {
                            Type: schema.TypeString,
                            Optional: true,
                            ValidateFunc: validateRFC3339Date,
                        },
                        "last_job": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },

            "user_access_rights": {
                Type: schema.TypeList,
                Optional: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "access_type": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(databoxedge.Change),
                                string(databoxedge.Read),
                                string(databoxedge.Custom),
                            }, false),
                        },
                        "user_id": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                    },
                },
            },

            "id": {
                Type: schema.TypeString,
                Computed: true,
            },

            "name": {
                Type: schema.TypeString,
                Computed: true,
            },

            "share_mappings": {
                Type: schema.TypeList,
                Computed: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "mount_point": {
                            Type: schema.TypeString,
                            Computed: true,
                        },
                        "role_id": {
                            Type: schema.TypeString,
                            Computed: true,
                        },
                        "role_type": {
                            Type: schema.TypeString,
                            Computed: true,
                        },
                        "share_id": {
                            Type: schema.TypeString,
                            Computed: true,
                        },
                    },
                },
            },

            "type": {
                Type: schema.TypeString,
                Computed: true,
            },
        },
    }
}

func resourceArmShareCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).sharesClient
    ctx, cancel := timeouts.ForCreateUpdate(meta.(*ArmClient).StopContext, d)
    defer cancel()

    resourceGroupName := d.Get("resource_group").(string)
    name := d.Get("device_name").(string)
    name := d.Get("name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroupName, name, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Share (Name %q / Resource Group %q / Device Name %q): %+v", name, resourceGroupName, name, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_share", *existing.ID)
        }
    }

    accessProtocol := d.Get("access_protocol").(string)
    azureContainerInfo := d.Get("azure_container_info").([]interface{})
    clientAccessRights := d.Get("client_access_rights").([]interface{})
    dataPolicy := d.Get("data_policy").(string)
    description := d.Get("description").(string)
    monitoringStatus := d.Get("monitoring_status").(string)
    refreshDetails := d.Get("refresh_details").([]interface{})
    shareStatus := d.Get("share_status").(string)
    userAccessRights := d.Get("user_access_rights").([]interface{})

    share := databoxedge.Share{
        ShareProperties: &databoxedge.ShareProperties{
            AccessProtocol: databoxedge.ShareAccessProtocol(accessProtocol),
            AzureContainerInfo: expandArmShareAzureContainerInfo(azureContainerInfo),
            ClientAccessRights: expandArmShareClientAccessRight(clientAccessRights),
            DataPolicy: databoxedge.DataPolicy(dataPolicy),
            Description: utils.String(description),
            MonitoringStatus: databoxedge.MonitoringStatus(monitoringStatus),
            RefreshDetails: expandArmShareRefreshDetails(refreshDetails),
            ShareStatus: databoxedge.ShareStatus(shareStatus),
            UserAccessRights: expandArmShareUserAccessRight(userAccessRights),
        },
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroupName, name, name, share)
    if err != nil {
        return fmt.Errorf("Error creating Share (Name %q / Resource Group %q / Device Name %q): %+v", name, resourceGroupName, name, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Share (Name %q / Resource Group %q / Device Name %q): %+v", name, resourceGroupName, name, err)
    }


    resp, err := client.Get(ctx, resourceGroupName, name, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Share (Name %q / Resource Group %q / Device Name %q): %+v", name, resourceGroupName, name, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Share (Name %q / Resource Group %q / Device Name %q) ID", name, resourceGroupName, name)
    }
    d.SetId(*resp.ID)

    return resourceArmShareRead(d, meta)
}

func resourceArmShareRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).sharesClient
    ctx, cancel := timeouts.ForRead(meta.(*ArmClient).StopContext, d)
    defer cancel()

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroupName := id.ResourceGroup
    name := id.Path["dataBoxEdgeDevices"]
    name := id.Path["shares"]

    resp, err := client.Get(ctx, resourceGroupName, name, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Share %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Share (Name %q / Resource Group %q / Device Name %q): %+v", name, resourceGroupName, name, err)
    }


    d.Set("resource_group", resourceGroupName)
    if shareProperties := resp.ShareProperties; shareProperties != nil {
        d.Set("access_protocol", string(shareProperties.AccessProtocol))
        if err := d.Set("azure_container_info", flattenArmShareAzureContainerInfo(shareProperties.AzureContainerInfo)); err != nil {
            return fmt.Errorf("Error setting `azure_container_info`: %+v", err)
        }
        if err := d.Set("client_access_rights", flattenArmShareClientAccessRight(shareProperties.ClientAccessRights)); err != nil {
            return fmt.Errorf("Error setting `client_access_rights`: %+v", err)
        }
        d.Set("data_policy", string(shareProperties.DataPolicy))
        d.Set("description", shareProperties.Description)
        d.Set("monitoring_status", string(shareProperties.MonitoringStatus))
        if err := d.Set("refresh_details", flattenArmShareRefreshDetails(shareProperties.RefreshDetails)); err != nil {
            return fmt.Errorf("Error setting `refresh_details`: %+v", err)
        }
        if err := d.Set("share_mappings", flattenArmShareMountPointMap(shareProperties.ShareMappings)); err != nil {
            return fmt.Errorf("Error setting `share_mappings`: %+v", err)
        }
        d.Set("share_status", string(shareProperties.ShareStatus))
        if err := d.Set("user_access_rights", flattenArmShareUserAccessRight(shareProperties.UserAccessRights)); err != nil {
            return fmt.Errorf("Error setting `user_access_rights`: %+v", err)
        }
    }
    d.Set("device_name", name)
    d.Set("id", resp.ID)
    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("type", resp.Type)

    return nil
}


func resourceArmShareDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).sharesClient
    ctx, cancel := timeouts.ForDelete(meta.(*ArmClient).StopContext, d)
    defer cancel()


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroupName := id.ResourceGroup
    name := id.Path["dataBoxEdgeDevices"]
    name := id.Path["shares"]

    future, err := client.Delete(ctx, resourceGroupName, name, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Share (Name %q / Resource Group %q / Device Name %q): %+v", name, resourceGroupName, name, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Share (Name %q / Resource Group %q / Device Name %q): %+v", name, resourceGroupName, name, err)
        }
    }

    return nil
}

func expandArmShareAzureContainerInfo(input []interface{}) *databoxedge.AzureContainerInfo {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    storageAccountCredentialID := v["storage_account_credential_id"].(string)
    containerName := v["container_name"].(string)
    dataFormat := v["data_format"].(string)

    result := databoxedge.AzureContainerInfo{
        ContainerName: utils.String(containerName),
        DataFormat: databoxedge.AzureContainerDataFormat(dataFormat),
        StorageAccountCredentialID: utils.String(storageAccountCredentialID),
    }
    return &result
}

func expandArmShareClientAccessRight(input []interface{}) *[]databoxedge.ClientAccessRight {
    results := make([]databoxedge.ClientAccessRight, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        client := v["client"].(string)
        accessPermission := v["access_permission"].(string)

        result := databoxedge.ClientAccessRight{
            AccessPermission: databoxedge.ClientPermissionType(accessPermission),
            Client: utils.String(client),
        }

        results = append(results, result)
    }
    return &results
}

func expandArmShareRefreshDetails(input []interface{}) *databoxedge.RefreshDetails {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    inProgressRefreshJobID := v["in_progress_refresh_job_id"].(string)
    lastCompletedRefreshJobTimeInUTC := v["last_completed_refresh_job_time_in_utc"].(string)
    errorManifestFile := v["error_manifest_file"].(string)
    lastJob := v["last_job"].(string)

    result := databoxedge.RefreshDetails{
        ErrorManifestFile: utils.String(errorManifestFile),
        InProgressRefreshJobID: utils.String(inProgressRefreshJobID),
        LastCompletedRefreshJobTimeInUTC: convertStringToDate(lastCompletedRefreshJobTimeInUTC),
        LastJob: utils.String(lastJob),
    }
    return &result
}

func expandArmShareUserAccessRight(input []interface{}) *[]databoxedge.UserAccessRight {
    results := make([]databoxedge.UserAccessRight, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        userID := v["user_id"].(string)
        accessType := v["access_type"].(string)

        result := databoxedge.UserAccessRight{
            AccessType: databoxedge.ShareAccessType(accessType),
            UserID: utils.String(userID),
        }

        results = append(results, result)
    }
    return &results
}

func convertStringToDate(input interface{}) *date.Time {
  v := input.(string)

  dateTime, err := date.ParseTime(time.RFC3339, v)
  if err != nil {
      log.Printf("[ERROR] Cannot convert an invalid string to RFC3339 date %q: %+v", v, err)
      return nil
  }

  result := date.Time{
      Time: dateTime,
  }
  return &result
}


func flattenArmShareAzureContainerInfo(input *databoxedge.AzureContainerInfo) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if containerName := input.ContainerName; containerName != nil {
        result["container_name"] = *containerName
    }
    result["data_format"] = string(input.DataFormat)
    if storageAccountCredentialId := input.StorageAccountCredentialID; storageAccountCredentialId != nil {
        result["storage_account_credential_id"] = *storageAccountCredentialId
    }

    return []interface{}{result}
}

func flattenArmShareClientAccessRight(input *[]databoxedge.ClientAccessRight) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        v["access_permission"] = string(item.AccessPermission)
        if client := item.Client; client != nil {
            v["client"] = *client
        }

        results = append(results, v)
    }

    return results
}

func flattenArmShareRefreshDetails(input *databoxedge.RefreshDetails) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if errorManifestFile := input.ErrorManifestFile; errorManifestFile != nil {
        result["error_manifest_file"] = *errorManifestFile
    }
    if inProgressRefreshJobId := input.InProgressRefreshJobID; inProgressRefreshJobId != nil {
        result["in_progress_refresh_job_id"] = *inProgressRefreshJobId
    }
    if lastCompletedRefreshJobTimeInUtc := input.LastCompletedRefreshJobTimeInUTC; lastCompletedRefreshJobTimeInUtc != nil {
        result["last_completed_refresh_job_time_in_utc"] = (*lastCompletedRefreshJobTimeInUtc).String()
    }
    if lastJob := input.LastJob; lastJob != nil {
        result["last_job"] = *lastJob
    }

    return []interface{}{result}
}

func flattenArmShareMountPointMap(input *[]databoxedge.MountPointMap) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        if mountPoint := item.MountPoint; mountPoint != nil {
            v["mount_point"] = *mountPoint
        }
        if roleId := item.RoleID; roleId != nil {
            v["role_id"] = *roleId
        }
        v["role_type"] = string(item.RoleType)
        if shareId := item.ShareID; shareId != nil {
            v["share_id"] = *shareId
        }

        results = append(results, v)
    }

    return results
}

func flattenArmShareUserAccessRight(input *[]databoxedge.UserAccessRight) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        v["access_type"] = string(item.AccessType)
        if userId := item.UserID; userId != nil {
            v["user_id"] = *userId
        }

        results = append(results, v)
    }

    return results
}
