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



func resourceArmOpenShiftManagedCluster() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmOpenShiftManagedClusterCreateUpdate,
        Read: resourceArmOpenShiftManagedClusterRead,
        Update: resourceArmOpenShiftManagedClusterCreateUpdate,
        Delete: resourceArmOpenShiftManagedClusterDelete,

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

            "open_shift_version": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "agent_pool_profiles": {
                Type: schema.TypeList,
                Optional: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "count": {
                            Type: schema.TypeInt,
                            Required: true,
                        },
                        "name": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "vm_size": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(containerservices.Standard_D2s_v3),
                                string(containerservices.Standard_D4s_v3),
                                string(containerservices.Standard_D8s_v3),
                                string(containerservices.Standard_D16s_v3),
                                string(containerservices.Standard_D32s_v3),
                                string(containerservices.Standard_D64s_v3),
                                string(containerservices.Standard_DS4_v2),
                                string(containerservices.Standard_DS5_v2),
                                string(containerservices.Standard_F8s_v2),
                                string(containerservices.Standard_F16s_v2),
                                string(containerservices.Standard_F32s_v2),
                                string(containerservices.Standard_F64s_v2),
                                string(containerservices.Standard_F72s_v2),
                                string(containerservices.Standard_F8s),
                                string(containerservices.Standard_F16s),
                                string(containerservices.Standard_E4s_v3),
                                string(containerservices.Standard_E8s_v3),
                                string(containerservices.Standard_E16s_v3),
                                string(containerservices.Standard_E20s_v3),
                                string(containerservices.Standard_E32s_v3),
                                string(containerservices.Standard_E64s_v3),
                                string(containerservices.Standard_GS2),
                                string(containerservices.Standard_GS3),
                                string(containerservices.Standard_GS4),
                                string(containerservices.Standard_GS5),
                                string(containerservices.Standard_DS12_v2),
                                string(containerservices.Standard_DS13_v2),
                                string(containerservices.Standard_DS14_v2),
                                string(containerservices.Standard_DS15_v2),
                                string(containerservices.Standard_L4s),
                                string(containerservices.Standard_L8s),
                                string(containerservices.Standard_L16s),
                                string(containerservices.Standard_L32s),
                            }, false),
                        },
                        "os_type": {
                            Type: schema.TypeString,
                            Optional: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(containerservices.Linux),
                                string(containerservices.Windows),
                            }, false),
                            Default: string(containerservices.Linux),
                        },
                        "role": {
                            Type: schema.TypeString,
                            Optional: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(containerservices.compute),
                                string(containerservices.infra),
                            }, false),
                            Default: string(containerservices.compute),
                        },
                        "subnet_cidr": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },

            "auth_profile": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "identity_providers": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "name": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                },
                            },
                        },
                    },
                },
            },

            "master_pool_profile": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "count": {
                            Type: schema.TypeInt,
                            Required: true,
                        },
                        "vm_size": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(containerservices.Standard_D2s_v3),
                                string(containerservices.Standard_D4s_v3),
                                string(containerservices.Standard_D8s_v3),
                                string(containerservices.Standard_D16s_v3),
                                string(containerservices.Standard_D32s_v3),
                                string(containerservices.Standard_D64s_v3),
                                string(containerservices.Standard_DS4_v2),
                                string(containerservices.Standard_DS5_v2),
                                string(containerservices.Standard_F8s_v2),
                                string(containerservices.Standard_F16s_v2),
                                string(containerservices.Standard_F32s_v2),
                                string(containerservices.Standard_F64s_v2),
                                string(containerservices.Standard_F72s_v2),
                                string(containerservices.Standard_F8s),
                                string(containerservices.Standard_F16s),
                                string(containerservices.Standard_E4s_v3),
                                string(containerservices.Standard_E8s_v3),
                                string(containerservices.Standard_E16s_v3),
                                string(containerservices.Standard_E20s_v3),
                                string(containerservices.Standard_E32s_v3),
                                string(containerservices.Standard_E64s_v3),
                                string(containerservices.Standard_GS2),
                                string(containerservices.Standard_GS3),
                                string(containerservices.Standard_GS4),
                                string(containerservices.Standard_GS5),
                                string(containerservices.Standard_DS12_v2),
                                string(containerservices.Standard_DS13_v2),
                                string(containerservices.Standard_DS14_v2),
                                string(containerservices.Standard_DS15_v2),
                                string(containerservices.Standard_L4s),
                                string(containerservices.Standard_L8s),
                                string(containerservices.Standard_L16s),
                                string(containerservices.Standard_L32s),
                            }, false),
                        },
                        "name": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "os_type": {
                            Type: schema.TypeString,
                            Optional: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(containerservices.Linux),
                                string(containerservices.Windows),
                            }, false),
                            Default: string(containerservices.Linux),
                        },
                        "subnet_cidr": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },

            "monitor_profile": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "enabled": {
                            Type: schema.TypeBool,
                            Optional: true,
                        },
                        "workspace_resource_id": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },

            "network_profile": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "peer_vnet_id": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "vnet_cidr": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "vnet_id": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },

            "plan": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "name": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "product": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "promotion_code": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "publisher": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },

            "router_profiles": {
                Type: schema.TypeList,
                Optional: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "name": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },

            "cluster_version": {
                Type: schema.TypeString,
                Computed: true,
            },

            "fqdn": {
                Type: schema.TypeString,
                Computed: true,
            },

            "provisioning_state": {
                Type: schema.TypeString,
                Computed: true,
            },

            "public_hostname": {
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

func resourceArmOpenShiftManagedClusterCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).openShiftManagedClustersClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Open Shift Managed Cluster %q (Resource Group %q): %+v", name, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_open_shift_managed_cluster", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    agentPoolProfiles := d.Get("agent_pool_profiles").([]interface{})
    authProfile := d.Get("auth_profile").([]interface{})
    masterPoolProfile := d.Get("master_pool_profile").([]interface{})
    monitorProfile := d.Get("monitor_profile").([]interface{})
    networkProfile := d.Get("network_profile").([]interface{})
    openShiftVersion := d.Get("open_shift_version").(string)
    plan := d.Get("plan").([]interface{})
    routerProfiles := d.Get("router_profiles").([]interface{})
    t := d.Get("tags").(map[string]interface{})

    parameters := containerservices.OpenShiftManagedCluster{
        Location: utils.String(location),
        Plan: expandArmOpenShiftManagedClusterPurchasePlan(plan),
        OpenShiftManagedClusterProperties: &containerservices.OpenShiftManagedClusterProperties{
            AgentPoolProfiles: expandArmOpenShiftManagedClusterOpenShiftManagedClusterAgentPoolProfile(agentPoolProfiles),
            AuthProfile: expandArmOpenShiftManagedClusterOpenShiftManagedClusterAuthProfile(authProfile),
            MasterPoolProfile: expandArmOpenShiftManagedClusterOpenShiftManagedClusterMasterPoolProfile(masterPoolProfile),
            MonitorProfile: expandArmOpenShiftManagedClusterOpenShiftManagedClusterMonitorProfile(monitorProfile),
            NetworkProfile: expandArmOpenShiftManagedClusterNetworkProfile(networkProfile),
            OpenShiftVersion: utils.String(openShiftVersion),
            RouterProfiles: expandArmOpenShiftManagedClusterOpenShiftRouterProfile(routerProfiles),
        },
        Tags: tags.Expand(t),
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, name, parameters)
    if err != nil {
        return fmt.Errorf("Error creating Open Shift Managed Cluster %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Open Shift Managed Cluster %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Open Shift Managed Cluster %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Open Shift Managed Cluster %q (Resource Group %q) ID", name, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmOpenShiftManagedClusterRead(d, meta)
}

func resourceArmOpenShiftManagedClusterRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).openShiftManagedClustersClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["openShiftManagedClusters"]

    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Open Shift Managed Cluster %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Open Shift Managed Cluster %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if openShiftManagedClusterProperties := resp.OpenShiftManagedClusterProperties; openShiftManagedClusterProperties != nil {
        if err := d.Set("agent_pool_profiles", flattenArmOpenShiftManagedClusterOpenShiftManagedClusterAgentPoolProfile(openShiftManagedClusterProperties.AgentPoolProfiles)); err != nil {
            return fmt.Errorf("Error setting `agent_pool_profiles`: %+v", err)
        }
        if err := d.Set("auth_profile", flattenArmOpenShiftManagedClusterOpenShiftManagedClusterAuthProfile(openShiftManagedClusterProperties.AuthProfile)); err != nil {
            return fmt.Errorf("Error setting `auth_profile`: %+v", err)
        }
        d.Set("cluster_version", openShiftManagedClusterProperties.ClusterVersion)
        d.Set("fqdn", openShiftManagedClusterProperties.Fqdn)
        if err := d.Set("master_pool_profile", flattenArmOpenShiftManagedClusterOpenShiftManagedClusterMasterPoolProfile(openShiftManagedClusterProperties.MasterPoolProfile)); err != nil {
            return fmt.Errorf("Error setting `master_pool_profile`: %+v", err)
        }
        if err := d.Set("monitor_profile", flattenArmOpenShiftManagedClusterOpenShiftManagedClusterMonitorProfile(openShiftManagedClusterProperties.MonitorProfile)); err != nil {
            return fmt.Errorf("Error setting `monitor_profile`: %+v", err)
        }
        if err := d.Set("network_profile", flattenArmOpenShiftManagedClusterNetworkProfile(openShiftManagedClusterProperties.NetworkProfile)); err != nil {
            return fmt.Errorf("Error setting `network_profile`: %+v", err)
        }
        d.Set("open_shift_version", openShiftManagedClusterProperties.OpenShiftVersion)
        d.Set("provisioning_state", openShiftManagedClusterProperties.ProvisioningState)
        d.Set("public_hostname", openShiftManagedClusterProperties.PublicHostname)
        if err := d.Set("router_profiles", flattenArmOpenShiftManagedClusterOpenShiftRouterProfile(openShiftManagedClusterProperties.RouterProfiles)); err != nil {
            return fmt.Errorf("Error setting `router_profiles`: %+v", err)
        }
    }
    if err := d.Set("plan", flattenArmOpenShiftManagedClusterPurchasePlan(resp.Plan)); err != nil {
        return fmt.Errorf("Error setting `plan`: %+v", err)
    }
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}


func resourceArmOpenShiftManagedClusterDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).openShiftManagedClustersClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["openShiftManagedClusters"]

    future, err := client.Delete(ctx, resourceGroup, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Open Shift Managed Cluster %q (Resource Group %q): %+v", name, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Open Shift Managed Cluster %q (Resource Group %q): %+v", name, resourceGroup, err)
        }
    }

    return nil
}

func expandArmOpenShiftManagedClusterPurchasePlan(input []interface{}) *containerservices.PurchasePlan {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    name := v["name"].(string)
    product := v["product"].(string)
    promotionCode := v["promotion_code"].(string)
    publisher := v["publisher"].(string)

    result := containerservices.PurchasePlan{
        Name: utils.String(name),
        Product: utils.String(product),
        PromotionCode: utils.String(promotionCode),
        Publisher: utils.String(publisher),
    }
    return &result
}

func expandArmOpenShiftManagedClusterOpenShiftManagedClusterAgentPoolProfile(input []interface{}) *[]containerservices.OpenShiftManagedClusterAgentPoolProfile {
    results := make([]containerservices.OpenShiftManagedClusterAgentPoolProfile, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        name := v["name"].(string)
        count := v["count"].(int)
        vmSize := v["vm_size"].(string)
        subnetCidr := v["subnet_cidr"].(string)
        osType := v["os_type"].(string)
        role := v["role"].(string)

        result := containerservices.OpenShiftManagedClusterAgentPoolProfile{
            Count: utils.Int32(int32(count)),
            Name: utils.String(name),
            OsType: containerservices.OSType(osType),
            Role: containerservices.OpenShiftAgentPoolProfileRole(role),
            SubnetCidr: utils.String(subnetCidr),
            VmSize: containerservices.OpenShiftContainerServiceVMSize(vmSize),
        }

        results = append(results, result)
    }
    return &results
}

func expandArmOpenShiftManagedClusterOpenShiftManagedClusterAuthProfile(input []interface{}) *containerservices.OpenShiftManagedClusterAuthProfile {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    identityProviders := v["identity_providers"].([]interface{})

    result := containerservices.OpenShiftManagedClusterAuthProfile{
        IdentityProviders: expandArmOpenShiftManagedClusterOpenShiftManagedClusterIdentityProvider(identityProviders),
    }
    return &result
}

func expandArmOpenShiftManagedClusterOpenShiftManagedClusterMasterPoolProfile(input []interface{}) *containerservices.OpenShiftManagedClusterMasterPoolProfile {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    name := v["name"].(string)
    count := v["count"].(int)
    vmSize := v["vm_size"].(string)
    subnetCidr := v["subnet_cidr"].(string)
    osType := v["os_type"].(string)

    result := containerservices.OpenShiftManagedClusterMasterPoolProfile{
        Count: utils.Int32(int32(count)),
        Name: utils.String(name),
        OsType: containerservices.OSType(osType),
        SubnetCidr: utils.String(subnetCidr),
        VmSize: containerservices.OpenShiftContainerServiceVMSize(vmSize),
    }
    return &result
}

func expandArmOpenShiftManagedClusterOpenShiftManagedClusterMonitorProfile(input []interface{}) *containerservices.OpenShiftManagedClusterMonitorProfile {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    workspaceResourceId := v["workspace_resource_id"].(string)
    enabled := v["enabled"].(bool)

    result := containerservices.OpenShiftManagedClusterMonitorProfile{
        Enabled: utils.Bool(enabled),
        WorkspaceResourceID: utils.String(workspaceResourceId),
    }
    return &result
}

func expandArmOpenShiftManagedClusterNetworkProfile(input []interface{}) *containerservices.NetworkProfile {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    vnetCidr := v["vnet_cidr"].(string)
    peerVnetId := v["peer_vnet_id"].(string)
    vnetId := v["vnet_id"].(string)

    result := containerservices.NetworkProfile{
        PeerVnetID: utils.String(peerVnetId),
        VnetCidr: utils.String(vnetCidr),
        VnetID: utils.String(vnetId),
    }
    return &result
}

func expandArmOpenShiftManagedClusterOpenShiftRouterProfile(input []interface{}) *[]containerservices.OpenShiftRouterProfile {
    results := make([]containerservices.OpenShiftRouterProfile, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        name := v["name"].(string)

        result := containerservices.OpenShiftRouterProfile{
            Name: utils.String(name),
        }

        results = append(results, result)
    }
    return &results
}

func expandArmOpenShiftManagedClusterOpenShiftManagedClusterIdentityProvider(input []interface{}) *[]containerservices.OpenShiftManagedClusterIdentityProvider {
    results := make([]containerservices.OpenShiftManagedClusterIdentityProvider, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        name := v["name"].(string)

        result := containerservices.OpenShiftManagedClusterIdentityProvider{
            Name: utils.String(name),
        }

        results = append(results, result)
    }
    return &results
}


func flattenArmOpenShiftManagedClusterOpenShiftManagedClusterAgentPoolProfile(input *[]containerservices.OpenShiftManagedClusterAgentPoolProfile) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        if name := item.Name; name != nil {
            v["name"] = *name
        }
        if count := item.Count; count != nil {
            v["count"] = int(*count)
        }
        v["os_type"] = string(item.OsType)
        v["role"] = string(item.Role)
        if subnetCidr := item.SubnetCidr; subnetCidr != nil {
            v["subnet_cidr"] = *subnetCidr
        }
        v["vm_size"] = string(item.VmSize)

        results = append(results, v)
    }

    return results
}

func flattenArmOpenShiftManagedClusterOpenShiftManagedClusterAuthProfile(input *containerservices.OpenShiftManagedClusterAuthProfile) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    result["identity_providers"] = flattenArmOpenShiftManagedClusterOpenShiftManagedClusterIdentityProvider(input.IdentityProviders)

    return []interface{}{result}
}

func flattenArmOpenShiftManagedClusterOpenShiftManagedClusterMasterPoolProfile(input *containerservices.OpenShiftManagedClusterMasterPoolProfile) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if name := input.Name; name != nil {
        result["name"] = *name
    }
    if count := input.Count; count != nil {
        result["count"] = int(*count)
    }
    result["os_type"] = string(input.OsType)
    if subnetCidr := input.SubnetCidr; subnetCidr != nil {
        result["subnet_cidr"] = *subnetCidr
    }
    result["vm_size"] = string(input.VmSize)

    return []interface{}{result}
}

func flattenArmOpenShiftManagedClusterOpenShiftManagedClusterMonitorProfile(input *containerservices.OpenShiftManagedClusterMonitorProfile) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if enabled := input.Enabled; enabled != nil {
        result["enabled"] = *enabled
    }
    if workspaceResourceId := input.WorkspaceResourceID; workspaceResourceId != nil {
        result["workspace_resource_id"] = *workspaceResourceId
    }

    return []interface{}{result}
}

func flattenArmOpenShiftManagedClusterNetworkProfile(input *containerservices.NetworkProfile) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if peerVnetId := input.PeerVnetID; peerVnetId != nil {
        result["peer_vnet_id"] = *peerVnetId
    }
    if vnetCidr := input.VnetCidr; vnetCidr != nil {
        result["vnet_cidr"] = *vnetCidr
    }
    if vnetId := input.VnetID; vnetId != nil {
        result["vnet_id"] = *vnetId
    }

    return []interface{}{result}
}

func flattenArmOpenShiftManagedClusterOpenShiftRouterProfile(input *[]containerservices.OpenShiftRouterProfile) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        if name := item.Name; name != nil {
            v["name"] = *name
        }

        results = append(results, v)
    }

    return results
}

func flattenArmOpenShiftManagedClusterPurchasePlan(input *containerservices.PurchasePlan) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if name := input.Name; name != nil {
        result["name"] = *name
    }
    if product := input.Product; product != nil {
        result["product"] = *product
    }
    if promotionCode := input.PromotionCode; promotionCode != nil {
        result["promotion_code"] = *promotionCode
    }
    if publisher := input.Publisher; publisher != nil {
        result["publisher"] = *publisher
    }

    return []interface{}{result}
}

func flattenArmOpenShiftManagedClusterOpenShiftManagedClusterIdentityProvider(input *[]containerservices.OpenShiftManagedClusterIdentityProvider) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        if name := item.Name; name != nil {
            v["name"] = *name
        }

        results = append(results, v)
    }

    return results
}
