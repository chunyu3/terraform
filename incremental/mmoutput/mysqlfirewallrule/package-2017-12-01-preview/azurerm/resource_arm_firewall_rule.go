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



func resourceArmFirewallRule() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmFirewallRuleCreateUpdate,
        Read: resourceArmFirewallRuleRead,
        Update: resourceArmFirewallRuleCreateUpdate,
        Delete: resourceArmFirewallRuleDelete,

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

            "end_ip_address": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "server_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "start_ip_address": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "type": {
                Type: schema.TypeString,
                Computed: true,
            },
        },
    }
}

func resourceArmFirewallRuleCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).firewallRulesClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    serverName := d.Get("server_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, serverName, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Firewall Rule %q (Server Name %q / Resource Group %q): %+v", name, serverName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_firewall_rule", *existing.ID)
        }
    }

    endIpAddress := d.Get("end_ip_address").(string)
    startIpAddress := d.Get("start_ip_address").(string)

    parameters := mysql.FirewallRule{
        FirewallRuleProperties: &mysql.FirewallRuleProperties{
            EndIPAddress: utils.String(endIpAddress),
            StartIPAddress: utils.String(startIpAddress),
        },
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, serverName, name, parameters)
    if err != nil {
        return fmt.Errorf("Error creating Firewall Rule %q (Server Name %q / Resource Group %q): %+v", name, serverName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Firewall Rule %q (Server Name %q / Resource Group %q): %+v", name, serverName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, serverName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Firewall Rule %q (Server Name %q / Resource Group %q): %+v", name, serverName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Firewall Rule %q (Server Name %q / Resource Group %q) ID", name, serverName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmFirewallRuleRead(d, meta)
}

func resourceArmFirewallRuleRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).firewallRulesClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    serverName := id.Path["servers"]
    name := id.Path["firewallRules"]

    resp, err := client.Get(ctx, resourceGroup, serverName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Firewall Rule %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Firewall Rule %q (Server Name %q / Resource Group %q): %+v", name, serverName, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if firewallRuleProperties := resp.FirewallRuleProperties; firewallRuleProperties != nil {
        d.Set("end_ip_address", firewallRuleProperties.EndIPAddress)
        d.Set("start_ip_address", firewallRuleProperties.StartIPAddress)
    }
    d.Set("server_name", serverName)
    d.Set("type", resp.Type)

    return nil
}


func resourceArmFirewallRuleDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).firewallRulesClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    serverName := id.Path["servers"]
    name := id.Path["firewallRules"]

    future, err := client.Delete(ctx, resourceGroup, serverName, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Firewall Rule %q (Server Name %q / Resource Group %q): %+v", name, serverName, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Firewall Rule %q (Server Name %q / Resource Group %q): %+v", name, serverName, resourceGroup, err)
        }
    }

    return nil
}
