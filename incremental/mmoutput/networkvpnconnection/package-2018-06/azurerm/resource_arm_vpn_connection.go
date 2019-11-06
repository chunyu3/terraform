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



func resourceArmVpnConnection() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmVpnConnectionCreateUpdate,
        Read: resourceArmVpnConnectionRead,
        Update: resourceArmVpnConnectionCreateUpdate,
        Delete: resourceArmVpnConnectionDelete,

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
                Optional: true,
                ForceNew: true,
            },

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "gateway_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "connection_status": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(network.Unknown),
                    string(network.Connecting),
                    string(network.Connected),
                    string(network.NotConnected),
                }, false),
                Default: string(network.Unknown),
            },

            "enable_bgp": {
                Type: schema.TypeBool,
                Optional: true,
            },

            "ipsec_policies": {
                Type: schema.TypeList,
                Optional: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "dh_group": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(network.None),
                                string(network.DHGroup1),
                                string(network.DHGroup2),
                                string(network.DHGroup14),
                                string(network.DHGroup2048),
                                string(network.ECP256),
                                string(network.ECP384),
                                string(network.DHGroup24),
                            }, false),
                        },
                        "ike_encryption": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(network.DES),
                                string(network.DES3),
                                string(network.AES128),
                                string(network.AES192),
                                string(network.AES256),
                                string(network.GCMAES256),
                                string(network.GCMAES128),
                            }, false),
                        },
                        "ike_integrity": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(network.MD5),
                                string(network.SHA1),
                                string(network.SHA256),
                                string(network.SHA384),
                                string(network.GCMAES256),
                                string(network.GCMAES128),
                            }, false),
                        },
                        "ipsec_encryption": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(network.None),
                                string(network.DES),
                                string(network.DES3),
                                string(network.AES128),
                                string(network.AES192),
                                string(network.AES256),
                                string(network.GCMAES128),
                                string(network.GCMAES192),
                                string(network.GCMAES256),
                            }, false),
                        },
                        "ipsec_integrity": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(network.MD5),
                                string(network.SHA1),
                                string(network.SHA256),
                                string(network.GCMAES128),
                                string(network.GCMAES192),
                                string(network.GCMAES256),
                            }, false),
                        },
                        "pfs_group": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(network.None),
                                string(network.PFS1),
                                string(network.PFS2),
                                string(network.PFS2048),
                                string(network.ECP256),
                                string(network.ECP384),
                                string(network.PFS24),
                                string(network.PFS14),
                                string(network.PFSMM),
                            }, false),
                        },
                        "sa_data_size_kilobytes": {
                            Type: schema.TypeInt,
                            Required: true,
                        },
                        "sa_life_time_seconds": {
                            Type: schema.TypeInt,
                            Required: true,
                        },
                    },
                },
            },

            "remote_vpn_site": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "id": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },

            "routing_weight": {
                Type: schema.TypeInt,
                Optional: true,
            },

            "shared_key": {
                Type: schema.TypeString,
                Optional: true,
            },

            "connection_bandwidth_in_mbps": {
                Type: schema.TypeInt,
                Computed: true,
            },

            "egress_bytes_transferred": {
                Type: schema.TypeInt,
                Computed: true,
            },

            "etag": {
                Type: schema.TypeString,
                Computed: true,
            },

            "ingress_bytes_transferred": {
                Type: schema.TypeInt,
                Computed: true,
            },

            "provisioning_state": {
                Type: schema.TypeString,
                Computed: true,
            },
        },
    }
}

func resourceArmVpnConnectionCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).vpnConnectionsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    gatewayName := d.Get("gateway_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, gatewayName, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Vpn Connection %q (Gateway Name %q / Resource Group %q): %+v", name, gatewayName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_vpn_connection", *existing.ID)
        }
    }

    id := d.Get("id").(string)
    name := d.Get("name").(string)
    connectionStatus := d.Get("connection_status").(string)
    enableBgp := d.Get("enable_bgp").(bool)
    ipsecPolicies := d.Get("ipsec_policies").([]interface{})
    remoteVpnSite := d.Get("remote_vpn_site").([]interface{})
    routingWeight := d.Get("routing_weight").(int)
    sharedKey := d.Get("shared_key").(string)

    vpnConnectionParameters := network.VpnConnection{
        ID: utils.String(id),
        Name: utils.String(name),
        VpnConnectionProperties: &network.VpnConnectionProperties{
            ConnectionStatus: network.VpnConnectionStatus(connectionStatus),
            EnableBgp: utils.Bool(enableBgp),
            IpsecPolicies: expandArmVpnConnectionIpsecPolicy(ipsecPolicies),
            RemoteVpnSite: expandArmVpnConnectionSubResource(remoteVpnSite),
            RoutingWeight: utils.Int32(int32(routingWeight)),
            SharedKey: utils.String(sharedKey),
        },
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, gatewayName, name, vpnConnectionParameters)
    if err != nil {
        return fmt.Errorf("Error creating Vpn Connection %q (Gateway Name %q / Resource Group %q): %+v", name, gatewayName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Vpn Connection %q (Gateway Name %q / Resource Group %q): %+v", name, gatewayName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, gatewayName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Vpn Connection %q (Gateway Name %q / Resource Group %q): %+v", name, gatewayName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Vpn Connection %q (Gateway Name %q / Resource Group %q) ID", name, gatewayName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmVpnConnectionRead(d, meta)
}

func resourceArmVpnConnectionRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).vpnConnectionsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    gatewayName := id.Path["vpnGateways"]
    name := id.Path["vpnConnections"]

    resp, err := client.Get(ctx, resourceGroup, gatewayName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Vpn Connection %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Vpn Connection %q (Gateway Name %q / Resource Group %q): %+v", name, gatewayName, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if vpnConnectionProperties := resp.VpnConnectionProperties; vpnConnectionProperties != nil {
        d.Set("connection_bandwidth_in_mbps", int(*vpnConnectionProperties.ConnectionBandwidthInMbps))
        d.Set("connection_status", string(vpnConnectionProperties.ConnectionStatus))
        d.Set("egress_bytes_transferred", int(*vpnConnectionProperties.EgressBytesTransferred))
        d.Set("enable_bgp", vpnConnectionProperties.EnableBgp)
        d.Set("ingress_bytes_transferred", int(*vpnConnectionProperties.IngressBytesTransferred))
        if err := d.Set("ipsec_policies", flattenArmVpnConnectionIpsecPolicy(vpnConnectionProperties.IpsecPolicies)); err != nil {
            return fmt.Errorf("Error setting `ipsec_policies`: %+v", err)
        }
        d.Set("provisioning_state", string(vpnConnectionProperties.ProvisioningState))
        if err := d.Set("remote_vpn_site", flattenArmVpnConnectionSubResource(vpnConnectionProperties.RemoteVpnSite)); err != nil {
            return fmt.Errorf("Error setting `remote_vpn_site`: %+v", err)
        }
        d.Set("routing_weight", int(*vpnConnectionProperties.RoutingWeight))
        d.Set("shared_key", vpnConnectionProperties.SharedKey)
    }
    d.Set("etag", resp.Etag)
    d.Set("gateway_name", gatewayName)

    return nil
}


func resourceArmVpnConnectionDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).vpnConnectionsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    gatewayName := id.Path["vpnGateways"]
    name := id.Path["vpnConnections"]

    future, err := client.Delete(ctx, resourceGroup, gatewayName, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Vpn Connection %q (Gateway Name %q / Resource Group %q): %+v", name, gatewayName, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Vpn Connection %q (Gateway Name %q / Resource Group %q): %+v", name, gatewayName, resourceGroup, err)
        }
    }

    return nil
}

func expandArmVpnConnectionIpsecPolicy(input []interface{}) *[]network.IpsecPolicy {
    results := make([]network.IpsecPolicy, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        saLifeTimeSeconds := v["sa_life_time_seconds"].(int)
        saDataSizeKilobytes := v["sa_data_size_kilobytes"].(int)
        ipsecEncryption := v["ipsec_encryption"].(string)
        ipsecIntegrity := v["ipsec_integrity"].(string)
        ikeEncryption := v["ike_encryption"].(string)
        ikeIntegrity := v["ike_integrity"].(string)
        dhGroup := v["dh_group"].(string)
        pfsGroup := v["pfs_group"].(string)

        result := network.IpsecPolicy{
            DhGroup: network.DhGroup(dhGroup),
            IkeEncryption: network.IkeEncryption(ikeEncryption),
            IkeIntegrity: network.IkeIntegrity(ikeIntegrity),
            IpsecEncryption: network.IpsecEncryption(ipsecEncryption),
            IpsecIntegrity: network.IpsecIntegrity(ipsecIntegrity),
            PfsGroup: network.PfsGroup(pfsGroup),
            SaDataSizeKilobytes: utils.Int32(int32(saDataSizeKilobytes)),
            SaLifeTimeSeconds: utils.Int32(int32(saLifeTimeSeconds)),
        }

        results = append(results, result)
    }
    return &results
}

func expandArmVpnConnectionSubResource(input []interface{}) *network.SubResource {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    id := v["id"].(string)

    result := network.SubResource{
        ID: utils.String(id),
    }
    return &result
}


func flattenArmVpnConnectionIpsecPolicy(input *[]network.IpsecPolicy) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        v["dh_group"] = string(item.DhGroup)
        v["ike_encryption"] = string(item.IkeEncryption)
        v["ike_integrity"] = string(item.IkeIntegrity)
        v["ipsec_encryption"] = string(item.IpsecEncryption)
        v["ipsec_integrity"] = string(item.IpsecIntegrity)
        v["pfs_group"] = string(item.PfsGroup)
        if saDataSizeKilobytes := item.SaDataSizeKilobytes; saDataSizeKilobytes != nil {
            v["sa_data_size_kilobytes"] = int(*saDataSizeKilobytes)
        }
        if saLifeTimeSeconds := item.SaLifeTimeSeconds; saLifeTimeSeconds != nil {
            v["sa_life_time_seconds"] = int(*saLifeTimeSeconds)
        }

        results = append(results, v)
    }

    return results
}

func flattenArmVpnConnectionSubResource(input *network.SubResource) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if id := input.ID; id != nil {
        result["id"] = *id
    }

    return []interface{}{result}
}
