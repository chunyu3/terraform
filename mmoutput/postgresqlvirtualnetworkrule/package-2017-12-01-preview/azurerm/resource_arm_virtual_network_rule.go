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



func resourceArmVirtualNetworkRule() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmVirtualNetworkRuleCreateUpdate,
        Read: resourceArmVirtualNetworkRuleRead,
        Update: resourceArmVirtualNetworkRuleCreateUpdate,
        Delete: resourceArmVirtualNetworkRuleDelete,

        Importer: &schema.ResourceImporter{
            State: schema.ImportStatePassthrough,
        },


        Schema: map[string]*schema.Schema{
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

            "virtual_network_rule_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "virtual_network_subnet_id": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "ignore_missing_vnet_service_endpoint": {
                Type: schema.TypeBool,
                Optional: true,
            },

            "state": {
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

func resourceArmVirtualNetworkRuleCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).virtualNetworkRulesClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    serverName := d.Get("server_name").(string)
    virtualNetworkRuleName := d.Get("virtual_network_rule_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, serverName, virtualNetworkRuleName)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Virtual Network Rule (Virtual Network Rule Name %q / Server Name %q / Resource Group %q): %+v", virtualNetworkRuleName, serverName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_virtual_network_rule", *existing.ID)
        }
    }

    ignoreMissingVnetServiceEndpoint := d.Get("ignore_missing_vnet_service_endpoint").(bool)
    virtualNetworkSubnetId := d.Get("virtual_network_subnet_id").(string)

    parameters := postgresql.VirtualNetworkRule{
        VirtualNetworkRuleProperties: &postgresql.VirtualNetworkRuleProperties{
            IgnoreMissingVnetServiceEndpoint: utils.Bool(ignoreMissingVnetServiceEndpoint),
            VirtualNetworkSubnetID: utils.String(virtualNetworkSubnetId),
        },
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, serverName, virtualNetworkRuleName, parameters)
    if err != nil {
        return fmt.Errorf("Error creating Virtual Network Rule (Virtual Network Rule Name %q / Server Name %q / Resource Group %q): %+v", virtualNetworkRuleName, serverName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Virtual Network Rule (Virtual Network Rule Name %q / Server Name %q / Resource Group %q): %+v", virtualNetworkRuleName, serverName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, serverName, virtualNetworkRuleName)
    if err != nil {
        return fmt.Errorf("Error retrieving Virtual Network Rule (Virtual Network Rule Name %q / Server Name %q / Resource Group %q): %+v", virtualNetworkRuleName, serverName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Virtual Network Rule (Virtual Network Rule Name %q / Server Name %q / Resource Group %q) ID", virtualNetworkRuleName, serverName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmVirtualNetworkRuleRead(d, meta)
}

func resourceArmVirtualNetworkRuleRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).virtualNetworkRulesClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    serverName := id.Path["servers"]
    virtualNetworkRuleName := id.Path["virtualNetworkRules"]

    resp, err := client.Get(ctx, resourceGroup, serverName, virtualNetworkRuleName)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Virtual Network Rule %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Virtual Network Rule (Virtual Network Rule Name %q / Server Name %q / Resource Group %q): %+v", virtualNetworkRuleName, serverName, resourceGroup, err)
    }


    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if virtualNetworkRuleProperties := resp.VirtualNetworkRuleProperties; virtualNetworkRuleProperties != nil {
        d.Set("ignore_missing_vnet_service_endpoint", virtualNetworkRuleProperties.IgnoreMissingVnetServiceEndpoint)
        d.Set("state", string(virtualNetworkRuleProperties.State))
        d.Set("virtual_network_subnet_id", virtualNetworkRuleProperties.VirtualNetworkSubnetID)
    }
    d.Set("server_name", serverName)
    d.Set("type", resp.Type)
    d.Set("virtual_network_rule_name", virtualNetworkRuleName)

    return nil
}


func resourceArmVirtualNetworkRuleDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).virtualNetworkRulesClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    serverName := id.Path["servers"]
    virtualNetworkRuleName := id.Path["virtualNetworkRules"]

    future, err := client.Delete(ctx, resourceGroup, serverName, virtualNetworkRuleName)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Virtual Network Rule (Virtual Network Rule Name %q / Server Name %q / Resource Group %q): %+v", virtualNetworkRuleName, serverName, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Virtual Network Rule (Virtual Network Rule Name %q / Server Name %q / Resource Group %q): %+v", virtualNetworkRuleName, serverName, resourceGroup, err)
        }
    }

    return nil
}
