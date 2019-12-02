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
                Optional: true,
                ForceNew: true,
            },

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "server_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "end_ip_address": {
                Type: schema.TypeString,
                Optional: true,
            },

            "start_ip_address": {
                Type: schema.TypeString,
                Optional: true,
            },

            "values": {
                Type: schema.TypeList,
                Optional: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "end_ip_address": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "name": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "start_ip_address": {
                            Type: schema.TypeString,
                            Optional: true,
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

    name := d.Get("name").(string)
    endIpAddress := d.Get("end_ip_address").(string)
    startIpAddress := d.Get("start_ip_address").(string)
    values := d.Get("values").([]interface{})

    parameters := sql.FirewallRuleList{
        Name: utils.String(name),
        ServerFirewallRuleProperties: &sql.ServerFirewallRuleProperties{
            EndIPAddress: utils.String(endIpAddress),
            StartIPAddress: utils.String(startIpAddress),
        },
        Values: expandArmFirewallRuleFirewallRule(values),
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroup, serverName, name, parameters); err != nil {
        return fmt.Errorf("Error creating Firewall Rule %q (Server Name %q / Resource Group %q): %+v", name, serverName, resourceGroup, err)
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
    if serverFirewallRuleProperties := resp.ServerFirewallRuleProperties; serverFirewallRuleProperties != nil {
        d.Set("end_ip_address", serverFirewallRuleProperties.EndIPAddress)
        d.Set("start_ip_address", serverFirewallRuleProperties.StartIPAddress)
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

    if _, err := client.Delete(ctx, resourceGroup, serverName, name); err != nil {
        return fmt.Errorf("Error deleting Firewall Rule %q (Server Name %q / Resource Group %q): %+v", name, serverName, resourceGroup, err)
    }

    return nil
}

func expandArmFirewallRuleFirewallRule(input []interface{}) *[]sql.FirewallRule {
    results := make([]sql.FirewallRule, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        name := v["name"].(string)
        startIpAddress := v["start_ip_address"].(string)
        endIpAddress := v["end_ip_address"].(string)

        result := sql.FirewallRule{
            Name: utils.String(name),
            ServerFirewallRuleProperties: &sql.ServerFirewallRuleProperties{
                EndIPAddress: utils.String(endIpAddress),
                StartIPAddress: utils.String(startIpAddress),
            },
        }

        results = append(results, result)
    }
    return &results
}
