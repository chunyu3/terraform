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



func resourceArmVpnSite() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmVpnSiteCreateUpdate,
        Read: resourceArmVpnSiteRead,
        Update: resourceArmVpnSiteCreateUpdate,
        Delete: resourceArmVpnSiteDelete,

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

            "vpn_site_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "address_space": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "address_prefixes": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                    },
                },
            },

            "bgp_properties": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "asn": {
                            Type: schema.TypeInt,
                            Optional: true,
                        },
                        "bgp_peering_address": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "peer_weight": {
                            Type: schema.TypeInt,
                            Optional: true,
                        },
                    },
                },
            },

            "device_properties": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "device_model": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "device_vendor": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "link_speed_in_mbps": {
                            Type: schema.TypeInt,
                            Optional: true,
                        },
                    },
                },
            },

            "ip_address": {
                Type: schema.TypeString,
                Optional: true,
            },

            "site_key": {
                Type: schema.TypeString,
                Optional: true,
            },

            "virtual_wan": {
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

            "etag": {
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

func resourceArmVpnSiteCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).vpnSitesClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    vpnSiteName := d.Get("vpn_site_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, vpnSiteName)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Vpn Site (Vpn Site Name %q / Resource Group %q): %+v", vpnSiteName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_vpn_site", *existing.ID)
        }
    }

    id := d.Get("id").(string)
    location := azure.NormalizeLocation(d.Get("location").(string))
    addressSpace := d.Get("address_space").([]interface{})
    bgpProperties := d.Get("bgp_properties").([]interface{})
    deviceProperties := d.Get("device_properties").([]interface{})
    ipAddress := d.Get("ip_address").(string)
    siteKey := d.Get("site_key").(string)
    virtualWan := d.Get("virtual_wan").([]interface{})
    t := d.Get("tags").(map[string]interface{})

    vpnSiteParameters := network.VpnSite{
        ID: utils.String(id),
        Location: utils.String(location),
        VpnSiteProperties: &network.VpnSiteProperties{
            AddressSpace: expandArmVpnSiteAddressSpace(addressSpace),
            BgpProperties: expandArmVpnSiteBgpSettings(bgpProperties),
            DeviceProperties: expandArmVpnSiteDeviceProperties(deviceProperties),
            IpAddress: utils.String(ipAddress),
            SiteKey: utils.String(siteKey),
            VirtualWan: expandArmVpnSiteSubResource(virtualWan),
        },
        Tags: tags.Expand(t),
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, vpnSiteName, vpnSiteParameters)
    if err != nil {
        return fmt.Errorf("Error creating Vpn Site (Vpn Site Name %q / Resource Group %q): %+v", vpnSiteName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Vpn Site (Vpn Site Name %q / Resource Group %q): %+v", vpnSiteName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, vpnSiteName)
    if err != nil {
        return fmt.Errorf("Error retrieving Vpn Site (Vpn Site Name %q / Resource Group %q): %+v", vpnSiteName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Vpn Site (Vpn Site Name %q / Resource Group %q) ID", vpnSiteName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmVpnSiteRead(d, meta)
}

func resourceArmVpnSiteRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).vpnSitesClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    vpnSiteName := id.Path["vpnSites"]

    resp, err := client.Get(ctx, resourceGroup, vpnSiteName)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Vpn Site %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Vpn Site (Vpn Site Name %q / Resource Group %q): %+v", vpnSiteName, resourceGroup, err)
    }


    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if vpnSiteProperties := resp.VpnSiteProperties; vpnSiteProperties != nil {
        if err := d.Set("address_space", flattenArmVpnSiteAddressSpace(vpnSiteProperties.AddressSpace)); err != nil {
            return fmt.Errorf("Error setting `address_space`: %+v", err)
        }
        if err := d.Set("bgp_properties", flattenArmVpnSiteBgpSettings(vpnSiteProperties.BgpProperties)); err != nil {
            return fmt.Errorf("Error setting `bgp_properties`: %+v", err)
        }
        if err := d.Set("device_properties", flattenArmVpnSiteDeviceProperties(vpnSiteProperties.DeviceProperties)); err != nil {
            return fmt.Errorf("Error setting `device_properties`: %+v", err)
        }
        d.Set("ip_address", vpnSiteProperties.IpAddress)
        d.Set("provisioning_state", string(vpnSiteProperties.ProvisioningState))
        d.Set("site_key", vpnSiteProperties.SiteKey)
        if err := d.Set("virtual_wan", flattenArmVpnSiteSubResource(vpnSiteProperties.VirtualWan)); err != nil {
            return fmt.Errorf("Error setting `virtual_wan`: %+v", err)
        }
    }
    d.Set("etag", resp.Etag)
    d.Set("type", resp.Type)
    d.Set("vpn_site_name", vpnSiteName)

    return tags.FlattenAndSet(d, resp.Tags)
}


func resourceArmVpnSiteDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).vpnSitesClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    vpnSiteName := id.Path["vpnSites"]

    future, err := client.Delete(ctx, resourceGroup, vpnSiteName)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Vpn Site (Vpn Site Name %q / Resource Group %q): %+v", vpnSiteName, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Vpn Site (Vpn Site Name %q / Resource Group %q): %+v", vpnSiteName, resourceGroup, err)
        }
    }

    return nil
}

func expandArmVpnSiteAddressSpace(input []interface{}) *network.AddressSpace {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    addressPrefixes := v["address_prefixes"].([]interface{})

    result := network.AddressSpace{
        AddressPrefixes: utils.ExpandStringSlice(addressPrefixes),
    }
    return &result
}

func expandArmVpnSiteBgpSettings(input []interface{}) *network.BgpSettings {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    asn := v["asn"].(int)
    bgpPeeringAddress := v["bgp_peering_address"].(string)
    peerWeight := v["peer_weight"].(int)

    result := network.BgpSettings{
        Asn: utils.Int64(int64(asn)),
        BgpPeeringAddress: utils.String(bgpPeeringAddress),
        PeerWeight: utils.Int32(int32(peerWeight)),
    }
    return &result
}

func expandArmVpnSiteDeviceProperties(input []interface{}) *network.DeviceProperties {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    deviceVendor := v["device_vendor"].(string)
    deviceModel := v["device_model"].(string)
    linkSpeedInMbps := v["link_speed_in_mbps"].(int)

    result := network.DeviceProperties{
        DeviceModel: utils.String(deviceModel),
        DeviceVendor: utils.String(deviceVendor),
        LinkSpeedInMbps: utils.Int32(int32(linkSpeedInMbps)),
    }
    return &result
}

func expandArmVpnSiteSubResource(input []interface{}) *network.SubResource {
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


func flattenArmVpnSiteAddressSpace(input *network.AddressSpace) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    result["address_prefixes"] = utils.FlattenStringSlice(input.AddressPrefixes)

    return []interface{}{result}
}

func flattenArmVpnSiteBgpSettings(input *network.BgpSettings) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if asn := input.Asn; asn != nil {
        result["asn"] = int(*asn)
    }
    if bgpPeeringAddress := input.BgpPeeringAddress; bgpPeeringAddress != nil {
        result["bgp_peering_address"] = *bgpPeeringAddress
    }
    if peerWeight := input.PeerWeight; peerWeight != nil {
        result["peer_weight"] = int(*peerWeight)
    }

    return []interface{}{result}
}

func flattenArmVpnSiteDeviceProperties(input *network.DeviceProperties) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if deviceModel := input.DeviceModel; deviceModel != nil {
        result["device_model"] = *deviceModel
    }
    if deviceVendor := input.DeviceVendor; deviceVendor != nil {
        result["device_vendor"] = *deviceVendor
    }
    if linkSpeedInMbps := input.LinkSpeedInMbps; linkSpeedInMbps != nil {
        result["link_speed_in_mbps"] = int(*linkSpeedInMbps)
    }

    return []interface{}{result}
}

func flattenArmVpnSiteSubResource(input *network.SubResource) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if id := input.ID; id != nil {
        result["id"] = *id
    }

    return []interface{}{result}
}