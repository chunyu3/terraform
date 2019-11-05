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
                Computed: true,
            },

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "cache_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "end_ip": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "rule_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "start_ip": {
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

    resourceGroup := d.Get("resource_group").(string)
    cacheName := d.Get("cache_name").(string)
    ruleName := d.Get("rule_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, cacheName, ruleName)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Firewall Rule (Rule Name %q / Cache Name %q / Resource Group %q): %+v", ruleName, cacheName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_firewall_rule", *existing.ID)
        }
    }

    endIp := d.Get("end_ip").(string)
    startIp := d.Get("start_ip").(string)

    parameters := redis.FirewallRuleCreateParameters{
        FirewallRuleProperties: &redis.FirewallRuleProperties{
            EndIp: utils.String(endIp),
            StartIp: utils.String(startIp),
        },
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroup, cacheName, ruleName, parameters); err != nil {
        return fmt.Errorf("Error creating Firewall Rule (Rule Name %q / Cache Name %q / Resource Group %q): %+v", ruleName, cacheName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, cacheName, ruleName)
    if err != nil {
        return fmt.Errorf("Error retrieving Firewall Rule (Rule Name %q / Cache Name %q / Resource Group %q): %+v", ruleName, cacheName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Firewall Rule (Rule Name %q / Cache Name %q / Resource Group %q) ID", ruleName, cacheName, resourceGroup)
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
    cacheName := id.Path["Redis"]
    ruleName := id.Path["firewallRules"]

    resp, err := client.Get(ctx, resourceGroup, cacheName, ruleName)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Firewall Rule %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Firewall Rule (Rule Name %q / Cache Name %q / Resource Group %q): %+v", ruleName, cacheName, resourceGroup, err)
    }


    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    d.Set("cache_name", cacheName)
    if firewallRuleProperties := resp.FirewallRuleProperties; firewallRuleProperties != nil {
        d.Set("end_ip", firewallRuleProperties.EndIp)
        d.Set("start_ip", firewallRuleProperties.StartIp)
    }
    d.Set("rule_name", ruleName)
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
    cacheName := id.Path["Redis"]
    ruleName := id.Path["firewallRules"]

    if _, err := client.Delete(ctx, resourceGroup, cacheName, ruleName); err != nil {
        return fmt.Errorf("Error deleting Firewall Rule (Rule Name %q / Cache Name %q / Resource Group %q): %+v", ruleName, cacheName, resourceGroup, err)
    }

    return nil
}
