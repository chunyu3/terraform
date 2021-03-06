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



func resourceArmExpressRouteCircuitAuthorization() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmExpressRouteCircuitAuthorizationCreateUpdate,
        Read: resourceArmExpressRouteCircuitAuthorizationRead,
        Update: resourceArmExpressRouteCircuitAuthorizationCreateUpdate,
        Delete: resourceArmExpressRouteCircuitAuthorizationDelete,

        Importer: &schema.ResourceImporter{
            State: schema.ImportStatePassthrough,
        },


        Schema: map[string]*schema.Schema{
            "name": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
            },

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "authorization_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "circuit_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "authorization_key": {
                Type: schema.TypeString,
                Optional: true,
            },

            "authorization_use_status": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(network.Available),
                    string(network.InUse),
                }, false),
                Default: string(network.Available),
            },

            "etag": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
            },

            "provisioning_state": {
                Type: schema.TypeString,
                Computed: true,
            },
        },
    }
}

func resourceArmExpressRouteCircuitAuthorizationCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).expressRouteCircuitAuthorizationsClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    authorizationName := d.Get("authorization_name").(string)
    circuitName := d.Get("circuit_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, circuitName, authorizationName)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Express Route Circuit Authorization (Authorization Name %q / Circuit Name %q / Resource Group %q): %+v", authorizationName, circuitName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_express_route_circuit_authorization", *existing.ID)
        }
    }

    id := d.Get("id").(string)
    name := d.Get("name").(string)
    authorizationKey := d.Get("authorization_key").(string)
    authorizationUseStatus := d.Get("authorization_use_status").(string)
    etag := d.Get("etag").(string)

    authorizationParameters := network.ExpressRouteCircuitAuthorization{
        Etag: utils.String(etag),
        ID: utils.String(id),
        Name: utils.String(name),
        AuthorizationPropertiesFormat: &network.AuthorizationPropertiesFormat{
            AuthorizationKey: utils.String(authorizationKey),
            AuthorizationUseStatus: network.AuthorizationUseStatus(authorizationUseStatus),
        },
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, circuitName, authorizationName, authorizationParameters)
    if err != nil {
        return fmt.Errorf("Error creating Express Route Circuit Authorization (Authorization Name %q / Circuit Name %q / Resource Group %q): %+v", authorizationName, circuitName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Express Route Circuit Authorization (Authorization Name %q / Circuit Name %q / Resource Group %q): %+v", authorizationName, circuitName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, circuitName, authorizationName)
    if err != nil {
        return fmt.Errorf("Error retrieving Express Route Circuit Authorization (Authorization Name %q / Circuit Name %q / Resource Group %q): %+v", authorizationName, circuitName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Express Route Circuit Authorization (Authorization Name %q / Circuit Name %q / Resource Group %q) ID", authorizationName, circuitName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmExpressRouteCircuitAuthorizationRead(d, meta)
}

func resourceArmExpressRouteCircuitAuthorizationRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).expressRouteCircuitAuthorizationsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    circuitName := id.Path["expressRouteCircuits"]
    authorizationName := id.Path["authorizations"]

    resp, err := client.Get(ctx, resourceGroup, circuitName, authorizationName)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Express Route Circuit Authorization %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Express Route Circuit Authorization (Authorization Name %q / Circuit Name %q / Resource Group %q): %+v", authorizationName, circuitName, resourceGroup, err)
    }


    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if authorizationPropertiesFormat := resp.AuthorizationPropertiesFormat; authorizationPropertiesFormat != nil {
        d.Set("authorization_key", authorizationPropertiesFormat.AuthorizationKey)
        d.Set("authorization_use_status", string(authorizationPropertiesFormat.AuthorizationUseStatus))
        d.Set("provisioning_state", authorizationPropertiesFormat.ProvisioningState)
    }
    d.Set("authorization_name", authorizationName)
    d.Set("circuit_name", circuitName)
    d.Set("etag", resp.Etag)

    return nil
}


func resourceArmExpressRouteCircuitAuthorizationDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).expressRouteCircuitAuthorizationsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    circuitName := id.Path["expressRouteCircuits"]
    authorizationName := id.Path["authorizations"]

    future, err := client.Delete(ctx, resourceGroup, circuitName, authorizationName)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Express Route Circuit Authorization (Authorization Name %q / Circuit Name %q / Resource Group %q): %+v", authorizationName, circuitName, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Express Route Circuit Authorization (Authorization Name %q / Circuit Name %q / Resource Group %q): %+v", authorizationName, circuitName, resourceGroup, err)
        }
    }

    return nil
}
