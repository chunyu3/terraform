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



func resourceArmPrivateEndpointConnection() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmPrivateEndpointConnectionCreateUpdate,
        Read: resourceArmPrivateEndpointConnectionRead,
        Update: resourceArmPrivateEndpointConnectionCreateUpdate,
        Delete: resourceArmPrivateEndpointConnectionDelete,

        Importer: &schema.ResourceImporter{
            State: schema.ImportStatePassthrough,
        },


        Schema: map[string]*schema.Schema{
            "config_store_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "private_endpoint_connection_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "private_link_service_connection_state": {
                Type: schema.TypeList,
                Required: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "description": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "status": {
                            Type: schema.TypeString,
                            Optional: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(appconfiguration.Pending),
                                string(appconfiguration.Approved),
                                string(appconfiguration.Rejected),
                                string(appconfiguration.Disconnected),
                            }, false),
                            Default: string(appconfiguration.Pending),
                        },
                    },
                },
            },

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "private_endpoint": {
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

            "id": {
                Type: schema.TypeString,
                Computed: true,
            },

            "name": {
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
        },
    }
}

func resourceArmPrivateEndpointConnectionCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).privateEndpointConnectionsClient
    ctx, cancel := timeouts.ForCreateUpdate(meta.(*ArmClient).StopContext, d)
    defer cancel()

    resourceGroupName := d.Get("resource_group").(string)
    configStoreName := d.Get("config_store_name").(string)
    name := d.Get("private_endpoint_connection_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroupName, configStoreName, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Private Endpoint Connection (Private Endpoint Connection Name %q / Config Store Name %q / Resource Group %q): %+v", name, configStoreName, resourceGroupName, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_private_endpoint_connection", *existing.ID)
        }
    }

    privateEndpoint := d.Get("private_endpoint").([]interface{})
    privateLinkServiceConnectionState := d.Get("private_link_service_connection_state").([]interface{})

    privateEndpointConnection := appconfiguration.PrivateEndpointConnection{
        PrivateEndpointConnectionProperties: &appconfiguration.PrivateEndpointConnectionProperties{
            PrivateEndpoint: expandArmPrivateEndpointConnectionPrivateEndpoint(privateEndpoint),
            PrivateLinkServiceConnectionState: expandArmPrivateEndpointConnectionPrivateLinkServiceConnectionState(privateLinkServiceConnectionState),
        },
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroupName, configStoreName, name, privateEndpointConnection)
    if err != nil {
        return fmt.Errorf("Error creating Private Endpoint Connection (Private Endpoint Connection Name %q / Config Store Name %q / Resource Group %q): %+v", name, configStoreName, resourceGroupName, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Private Endpoint Connection (Private Endpoint Connection Name %q / Config Store Name %q / Resource Group %q): %+v", name, configStoreName, resourceGroupName, err)
    }


    resp, err := client.Get(ctx, resourceGroupName, configStoreName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Private Endpoint Connection (Private Endpoint Connection Name %q / Config Store Name %q / Resource Group %q): %+v", name, configStoreName, resourceGroupName, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Private Endpoint Connection (Private Endpoint Connection Name %q / Config Store Name %q / Resource Group %q) ID", name, configStoreName, resourceGroupName)
    }
    d.SetId(*resp.ID)

    return resourceArmPrivateEndpointConnectionRead(d, meta)
}

func resourceArmPrivateEndpointConnectionRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).privateEndpointConnectionsClient
    ctx, cancel := timeouts.ForRead(meta.(*ArmClient).StopContext, d)
    defer cancel()

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroupName := id.ResourceGroup
    configStoreName := id.Path["configurationStores"]
    name := id.Path["privateEndpointConnections"]

    resp, err := client.Get(ctx, resourceGroupName, configStoreName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Private Endpoint Connection %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Private Endpoint Connection (Private Endpoint Connection Name %q / Config Store Name %q / Resource Group %q): %+v", name, configStoreName, resourceGroupName, err)
    }


    d.Set("resource_group", resourceGroupName)
    d.Set("config_store_name", configStoreName)
    d.Set("id", resp.ID)
    d.Set("name", resp.Name)
    if privateEndpointConnectionProperties := resp.PrivateEndpointConnectionProperties; privateEndpointConnectionProperties != nil {
        if err := d.Set("private_endpoint", flattenArmPrivateEndpointConnectionPrivateEndpoint(privateEndpointConnectionProperties.PrivateEndpoint)); err != nil {
            return fmt.Errorf("Error setting `private_endpoint`: %+v", err)
        }
        if err := d.Set("private_link_service_connection_state", flattenArmPrivateEndpointConnectionPrivateLinkServiceConnectionState(privateEndpointConnectionProperties.PrivateLinkServiceConnectionState)); err != nil {
            return fmt.Errorf("Error setting `private_link_service_connection_state`: %+v", err)
        }
        d.Set("provisioning_state", string(privateEndpointConnectionProperties.ProvisioningState))
    }
    d.Set("private_endpoint_connection_name", name)
    d.Set("type", resp.Type)

    return nil
}


func resourceArmPrivateEndpointConnectionDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).privateEndpointConnectionsClient
    ctx, cancel := timeouts.ForDelete(meta.(*ArmClient).StopContext, d)
    defer cancel()


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroupName := id.ResourceGroup
    configStoreName := id.Path["configurationStores"]
    name := id.Path["privateEndpointConnections"]

    future, err := client.Delete(ctx, resourceGroupName, configStoreName, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Private Endpoint Connection (Private Endpoint Connection Name %q / Config Store Name %q / Resource Group %q): %+v", name, configStoreName, resourceGroupName, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Private Endpoint Connection (Private Endpoint Connection Name %q / Config Store Name %q / Resource Group %q): %+v", name, configStoreName, resourceGroupName, err)
        }
    }

    return nil
}

func expandArmPrivateEndpointConnectionPrivateEndpoint(input []interface{}) *appconfiguration.PrivateEndpoint {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    iD := v["id"].(string)

    result := appconfiguration.PrivateEndpoint{
        ID: utils.String(iD),
    }
    return &result
}

func expandArmPrivateEndpointConnectionPrivateLinkServiceConnectionState(input []interface{}) *appconfiguration.PrivateLinkServiceConnectionState {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    status := v["status"].(string)
    description := v["description"].(string)

    result := appconfiguration.PrivateLinkServiceConnectionState{
        Description: utils.String(description),
        Status: appconfiguration.ConnectionStatus(status),
    }
    return &result
}


func flattenArmPrivateEndpointConnectionPrivateEndpoint(input *appconfiguration.PrivateEndpoint) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if id := input.ID; id != nil {
        result["id"] = *id
    }

    return []interface{}{result}
}

func flattenArmPrivateEndpointConnectionPrivateLinkServiceConnectionState(input *appconfiguration.PrivateLinkServiceConnectionState) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if description := input.Description; description != nil {
        result["description"] = *description
    }
    result["status"] = string(input.Status)

    return []interface{}{result}
}