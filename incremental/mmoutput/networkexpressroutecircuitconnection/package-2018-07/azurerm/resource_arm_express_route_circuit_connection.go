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



func resourceArmExpressRouteCircuitConnection() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmExpressRouteCircuitConnectionCreateUpdate,
        Read: resourceArmExpressRouteCircuitConnectionRead,
        Update: resourceArmExpressRouteCircuitConnectionCreateUpdate,
        Delete: resourceArmExpressRouteCircuitConnectionDelete,

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

            "circuit_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "peering_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "address_prefix": {
                Type: schema.TypeString,
                Optional: true,
            },

            "authorization_key": {
                Type: schema.TypeString,
                Optional: true,
            },

            "express_route_circuit_peering": {
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

            "peer_express_route_circuit_peering": {
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

            "circuit_connection_status": {
                Type: schema.TypeString,
                Computed: true,
            },

            "etag": {
                Type: schema.TypeString,
                Computed: true,
            },

            "provisioning_state": {
                Type: schema.TypeString,
                Computed: true,
            },
        },
    }
}

func resourceArmExpressRouteCircuitConnectionCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).expressRouteCircuitConnectionsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    circuitName := d.Get("circuit_name").(string)
    peeringName := d.Get("peering_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, circuitName, peeringName, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Express Route Circuit Connection %q (Peering Name %q / Circuit Name %q / Resource Group %q): %+v", name, peeringName, circuitName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_express_route_circuit_connection", *existing.ID)
        }
    }

    id := d.Get("id").(string)
    name := d.Get("name").(string)
    addressPrefix := d.Get("address_prefix").(string)
    authorizationKey := d.Get("authorization_key").(string)
    expressRouteCircuitPeering := d.Get("express_route_circuit_peering").([]interface{})
    peerExpressRouteCircuitPeering := d.Get("peer_express_route_circuit_peering").([]interface{})

    expressRouteCircuitConnectionParameters := network.ExpressRouteCircuitConnection{
        ID: utils.String(id),
        Name: utils.String(name),
        ExpressRouteCircuitConnectionPropertiesFormat: &network.ExpressRouteCircuitConnectionPropertiesFormat{
            AddressPrefix: utils.String(addressPrefix),
            AuthorizationKey: utils.String(authorizationKey),
            ExpressRouteCircuitPeering: expandArmExpressRouteCircuitConnectionSubResource(expressRouteCircuitPeering),
            PeerExpressRouteCircuitPeering: expandArmExpressRouteCircuitConnectionSubResource(peerExpressRouteCircuitPeering),
        },
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, circuitName, peeringName, name, expressRouteCircuitConnectionParameters)
    if err != nil {
        return fmt.Errorf("Error creating Express Route Circuit Connection %q (Peering Name %q / Circuit Name %q / Resource Group %q): %+v", name, peeringName, circuitName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Express Route Circuit Connection %q (Peering Name %q / Circuit Name %q / Resource Group %q): %+v", name, peeringName, circuitName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, circuitName, peeringName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Express Route Circuit Connection %q (Peering Name %q / Circuit Name %q / Resource Group %q): %+v", name, peeringName, circuitName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Express Route Circuit Connection %q (Peering Name %q / Circuit Name %q / Resource Group %q) ID", name, peeringName, circuitName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmExpressRouteCircuitConnectionRead(d, meta)
}

func resourceArmExpressRouteCircuitConnectionRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).expressRouteCircuitConnectionsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    circuitName := id.Path["expressRouteCircuits"]
    peeringName := id.Path["peerings"]
    name := id.Path["connections"]

    resp, err := client.Get(ctx, resourceGroup, circuitName, peeringName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Express Route Circuit Connection %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Express Route Circuit Connection %q (Peering Name %q / Circuit Name %q / Resource Group %q): %+v", name, peeringName, circuitName, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if expressRouteCircuitConnectionPropertiesFormat := resp.ExpressRouteCircuitConnectionPropertiesFormat; expressRouteCircuitConnectionPropertiesFormat != nil {
        d.Set("address_prefix", expressRouteCircuitConnectionPropertiesFormat.AddressPrefix)
        d.Set("authorization_key", expressRouteCircuitConnectionPropertiesFormat.AuthorizationKey)
        d.Set("circuit_connection_status", string(expressRouteCircuitConnectionPropertiesFormat.CircuitConnectionStatus))
        if err := d.Set("express_route_circuit_peering", flattenArmExpressRouteCircuitConnectionSubResource(expressRouteCircuitConnectionPropertiesFormat.ExpressRouteCircuitPeering)); err != nil {
            return fmt.Errorf("Error setting `express_route_circuit_peering`: %+v", err)
        }
        if err := d.Set("peer_express_route_circuit_peering", flattenArmExpressRouteCircuitConnectionSubResource(expressRouteCircuitConnectionPropertiesFormat.PeerExpressRouteCircuitPeering)); err != nil {
            return fmt.Errorf("Error setting `peer_express_route_circuit_peering`: %+v", err)
        }
        d.Set("provisioning_state", expressRouteCircuitConnectionPropertiesFormat.ProvisioningState)
    }
    d.Set("circuit_name", circuitName)
    d.Set("etag", resp.Etag)
    d.Set("peering_name", peeringName)

    return nil
}


func resourceArmExpressRouteCircuitConnectionDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).expressRouteCircuitConnectionsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    circuitName := id.Path["expressRouteCircuits"]
    peeringName := id.Path["peerings"]
    name := id.Path["connections"]

    future, err := client.Delete(ctx, resourceGroup, circuitName, peeringName, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Express Route Circuit Connection %q (Peering Name %q / Circuit Name %q / Resource Group %q): %+v", name, peeringName, circuitName, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Express Route Circuit Connection %q (Peering Name %q / Circuit Name %q / Resource Group %q): %+v", name, peeringName, circuitName, resourceGroup, err)
        }
    }

    return nil
}

func expandArmExpressRouteCircuitConnectionSubResource(input []interface{}) *network.SubResource {
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


func flattenArmExpressRouteCircuitConnectionSubResource(input *network.SubResource) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if id := input.ID; id != nil {
        result["id"] = *id
    }

    return []interface{}{result}
}