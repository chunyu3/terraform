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
        Create: resourceArmFirewallRuleCreate,
        Read: resourceArmFirewallRuleRead,
        Update: resourceArmFirewallRuleUpdate,
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

            "account_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "end_ip_address": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "firewall_rule_name": {
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

func resourceArmFirewallRuleCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).firewallRulesClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    accountName := d.Get("account_name").(string)
    firewallRuleName := d.Get("firewall_rule_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, accountName, firewallRuleName)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Firewall Rule (Firewall Rule Name %q / Account Name %q / Resource Group %q): %+v", firewallRuleName, accountName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_firewall_rule", *existing.ID)
        }
    }

    endIpAddress := d.Get("end_ip_address").(string)
    startIpAddress := d.Get("start_ip_address").(string)

    parameters := datalakeanalytics.CreateOrUpdateFirewallRuleParameters{
        CreateOrUpdateFirewallRuleProperties: &datalakeanalytics.CreateOrUpdateFirewallRuleProperties{
            EndIpAddress: utils.String(endIpAddress),
            StartIpAddress: utils.String(startIpAddress),
        },
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroup, accountName, firewallRuleName, parameters); err != nil {
        return fmt.Errorf("Error creating Firewall Rule (Firewall Rule Name %q / Account Name %q / Resource Group %q): %+v", firewallRuleName, accountName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, accountName, firewallRuleName)
    if err != nil {
        return fmt.Errorf("Error retrieving Firewall Rule (Firewall Rule Name %q / Account Name %q / Resource Group %q): %+v", firewallRuleName, accountName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Firewall Rule (Firewall Rule Name %q / Account Name %q / Resource Group %q) ID", firewallRuleName, accountName, resourceGroup)
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
    accountName := id.Path["accounts"]
    firewallRuleName := id.Path["firewallRules"]

    resp, err := client.Get(ctx, resourceGroup, accountName, firewallRuleName)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Firewall Rule %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Firewall Rule (Firewall Rule Name %q / Account Name %q / Resource Group %q): %+v", firewallRuleName, accountName, resourceGroup, err)
    }


    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    d.Set("account_name", accountName)
    if createOrUpdateFirewallRuleProperties := resp.CreateOrUpdateFirewallRuleProperties; createOrUpdateFirewallRuleProperties != nil {
        d.Set("end_ip_address", createOrUpdateFirewallRuleProperties.EndIpAddress)
        d.Set("start_ip_address", createOrUpdateFirewallRuleProperties.StartIpAddress)
    }
    d.Set("firewall_rule_name", firewallRuleName)
    d.Set("type", resp.Type)

    return nil
}

func resourceArmFirewallRuleUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).firewallRulesClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    accountName := d.Get("account_name").(string)
    endIpAddress := d.Get("end_ip_address").(string)
    firewallRuleName := d.Get("firewall_rule_name").(string)
    startIpAddress := d.Get("start_ip_address").(string)

    parameters := datalakeanalytics.CreateOrUpdateFirewallRuleParameters{
        CreateOrUpdateFirewallRuleProperties: &datalakeanalytics.CreateOrUpdateFirewallRuleProperties{
            EndIpAddress: utils.String(endIpAddress),
            StartIpAddress: utils.String(startIpAddress),
        },
    }


    if _, err := client.Update(ctx, resourceGroup, accountName, firewallRuleName, parameters); err != nil {
        return fmt.Errorf("Error updating Firewall Rule (Firewall Rule Name %q / Account Name %q / Resource Group %q): %+v", firewallRuleName, accountName, resourceGroup, err)
    }

    return resourceArmFirewallRuleRead(d, meta)
}

func resourceArmFirewallRuleDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).firewallRulesClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    accountName := id.Path["accounts"]
    firewallRuleName := id.Path["firewallRules"]

    if _, err := client.Delete(ctx, resourceGroup, accountName, firewallRuleName); err != nil {
        return fmt.Errorf("Error deleting Firewall Rule (Firewall Rule Name %q / Account Name %q / Resource Group %q): %+v", firewallRuleName, accountName, resourceGroup, err)
    }

    return nil
}
