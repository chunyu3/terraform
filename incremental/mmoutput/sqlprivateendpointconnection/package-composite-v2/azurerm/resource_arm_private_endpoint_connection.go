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

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "server_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

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

            "private_link_service_connection_state": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "description": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "status": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
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
        },
    }
}

func resourceArmPrivateEndpointConnectionCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).privateEndpointConnectionsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    serverName := d.Get("server_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, serverName, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Private Endpoint Connection %q (Server Name %q / Resource Group %q): %+v", name, serverName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_private_endpoint_connection", *existing.ID)
        }
    }

    privateEndpoint := d.Get("private_endpoint").([]interface{})
    privateLinkServiceConnectionState := d.Get("private_link_service_connection_state").([]interface{})

    parameters := sql.PrivateEndpointConnection{
        PrivateEndpointConnectionProperties: &sql.PrivateEndpointConnectionProperties{
            PrivateEndpoint: expandArmPrivateEndpointConnectionPrivateEndpointProperty(privateEndpoint),
            PrivateLinkServiceConnectionState: expandArmPrivateEndpointConnectionPrivateLinkServiceConnectionStateProperty(privateLinkServiceConnectionState),
        },
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, serverName, name, parameters)
    if err != nil {
        return fmt.Errorf("Error creating Private Endpoint Connection %q (Server Name %q / Resource Group %q): %+v", name, serverName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Private Endpoint Connection %q (Server Name %q / Resource Group %q): %+v", name, serverName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, serverName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Private Endpoint Connection %q (Server Name %q / Resource Group %q): %+v", name, serverName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Private Endpoint Connection %q (Server Name %q / Resource Group %q) ID", name, serverName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmPrivateEndpointConnectionRead(d, meta)
}

func resourceArmPrivateEndpointConnectionRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).privateEndpointConnectionsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    serverName := id.Path["servers"]
    name := id.Path["privateEndpointConnections"]

    resp, err := client.Get(ctx, resourceGroup, serverName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Private Endpoint Connection %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Private Endpoint Connection %q (Server Name %q / Resource Group %q): %+v", name, serverName, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if privateEndpointConnectionProperties := resp.PrivateEndpointConnectionProperties; privateEndpointConnectionProperties != nil {
        if err := d.Set("private_endpoint", flattenArmPrivateEndpointConnectionPrivateEndpointProperty(privateEndpointConnectionProperties.PrivateEndpoint)); err != nil {
            return fmt.Errorf("Error setting `private_endpoint`: %+v", err)
        }
        if err := d.Set("private_link_service_connection_state", flattenArmPrivateEndpointConnectionPrivateLinkServiceConnectionStateProperty(privateEndpointConnectionProperties.PrivateLinkServiceConnectionState)); err != nil {
            return fmt.Errorf("Error setting `private_link_service_connection_state`: %+v", err)
        }
        d.Set("provisioning_state", privateEndpointConnectionProperties.ProvisioningState)
    }
    d.Set("server_name", serverName)
    d.Set("type", resp.Type)

    return nil
}


func resourceArmPrivateEndpointConnectionDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).privateEndpointConnectionsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    serverName := id.Path["servers"]
    name := id.Path["privateEndpointConnections"]

    future, err := client.Delete(ctx, resourceGroup, serverName, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Private Endpoint Connection %q (Server Name %q / Resource Group %q): %+v", name, serverName, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Private Endpoint Connection %q (Server Name %q / Resource Group %q): %+v", name, serverName, resourceGroup, err)
        }
    }

    return nil
}

func expandArmPrivateEndpointConnectionPrivateEndpointProperty(input []interface{}) *sql.PrivateEndpointProperty {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    id := v["id"].(string)

    result := sql.PrivateEndpointProperty{
        ID: utils.String(id),
    }
    return &result
}

func expandArmPrivateEndpointConnectionPrivateLinkServiceConnectionStateProperty(input []interface{}) *sql.PrivateLinkServiceConnectionStateProperty {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    status := v["status"].(string)
    description := v["description"].(string)

    result := sql.PrivateLinkServiceConnectionStateProperty{
        Description: utils.String(description),
        Status: utils.String(status),
    }
    return &result
}


func flattenArmPrivateEndpointConnectionPrivateEndpointProperty(input *sql.PrivateEndpointProperty) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if id := input.ID; id != nil {
        result["id"] = *id
    }

    return []interface{}{result}
}

func flattenArmPrivateEndpointConnectionPrivateLinkServiceConnectionStateProperty(input *sql.PrivateLinkServiceConnectionStateProperty) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if description := input.Description; description != nil {
        result["description"] = *description
    }
    if status := input.Status; status != nil {
        result["status"] = *status
    }

    return []interface{}{result}
}
