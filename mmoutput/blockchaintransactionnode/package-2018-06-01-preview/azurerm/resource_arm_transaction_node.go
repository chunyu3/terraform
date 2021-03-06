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



func resourceArmTransactionNode() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmTransactionNodeCreate,
        Read: resourceArmTransactionNodeRead,
        Update: resourceArmTransactionNodeUpdate,
        Delete: resourceArmTransactionNodeDelete,

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

            "location": azure.SchemaLocation(),

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "blockchain_member_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "firewall_rules": {
                Type: schema.TypeList,
                Optional: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "end_ip_address": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "rule_name": {
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

            "key_name": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
            },

            "password": {
                Type: schema.TypeString,
                Optional: true,
            },

            "value": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
            },

            "dns": {
                Type: schema.TypeString,
                Computed: true,
            },

            "provisioning_state": {
                Type: schema.TypeString,
                Computed: true,
            },

            "public_key": {
                Type: schema.TypeString,
                Computed: true,
            },

            "type": {
                Type: schema.TypeString,
                Computed: true,
            },

            "user_name": {
                Type: schema.TypeString,
                Computed: true,
            },
        },
    }
}

func resourceArmTransactionNodeCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).transactionNodesClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    blockchainMemberName := d.Get("blockchain_member_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, blockchainMemberName, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Transaction Node %q (Resource Group %q / Blockchain Member Name %q): %+v", name, resourceGroup, blockchainMemberName, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_transaction_node", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    firewallRules := d.Get("firewall_rules").([]interface{})
    keyName := d.Get("key_name").(string)
    password := d.Get("password").(string)
    value := d.Get("value").(string)

    transactionNode := blockchain.TransactionNodeUpdate{
        KeyName: utils.String(keyName),
        Location: utils.String(location),
        TransactionNodePropertiesUpdate: &blockchain.TransactionNodePropertiesUpdate{
            FirewallRules: expandArmTransactionNodeFirewallRule(firewallRules),
            Password: utils.String(password),
        },
        Value: utils.String(value),
    }


    future, err := client.Create(ctx, resourceGroup, blockchainMemberName, name, transactionNode)
    if err != nil {
        return fmt.Errorf("Error creating Transaction Node %q (Resource Group %q / Blockchain Member Name %q): %+v", name, resourceGroup, blockchainMemberName, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Transaction Node %q (Resource Group %q / Blockchain Member Name %q): %+v", name, resourceGroup, blockchainMemberName, err)
    }


    resp, err := client.Get(ctx, resourceGroup, blockchainMemberName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Transaction Node %q (Resource Group %q / Blockchain Member Name %q): %+v", name, resourceGroup, blockchainMemberName, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Transaction Node %q (Resource Group %q / Blockchain Member Name %q) ID", name, resourceGroup, blockchainMemberName)
    }
    d.SetId(*resp.ID)

    return resourceArmTransactionNodeRead(d, meta)
}

func resourceArmTransactionNodeRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).transactionNodesClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    blockchainMemberName := id.Path["blockchainMembers"]
    name := id.Path["transactionNodes"]

    resp, err := client.Get(ctx, resourceGroup, blockchainMemberName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Transaction Node %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Transaction Node %q (Resource Group %q / Blockchain Member Name %q): %+v", name, resourceGroup, blockchainMemberName, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    d.Set("blockchain_member_name", blockchainMemberName)
    if transactionNodePropertiesUpdate := resp.TransactionNodePropertiesUpdate; transactionNodePropertiesUpdate != nil {
        d.Set("dns", transactionNodePropertiesUpdate.DNS)
        if err := d.Set("firewall_rules", flattenArmTransactionNodeFirewallRule(transactionNodePropertiesUpdate.FirewallRules)); err != nil {
            return fmt.Errorf("Error setting `firewall_rules`: %+v", err)
        }
        d.Set("password", transactionNodePropertiesUpdate.Password)
        d.Set("provisioning_state", string(transactionNodePropertiesUpdate.ProvisioningState))
        d.Set("public_key", transactionNodePropertiesUpdate.PublicKey)
        d.Set("user_name", transactionNodePropertiesUpdate.UserName)
    }
    d.Set("type", resp.Type)

    return nil
}

func resourceArmTransactionNodeUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).transactionNodesClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    blockchainMemberName := d.Get("blockchain_member_name").(string)
    firewallRules := d.Get("firewall_rules").([]interface{})
    keyName := d.Get("key_name").(string)
    password := d.Get("password").(string)
    value := d.Get("value").(string)

    transactionNode := blockchain.TransactionNodeUpdate{
        KeyName: utils.String(keyName),
        TransactionNodePropertiesUpdate: &blockchain.TransactionNodePropertiesUpdate{
            FirewallRules: expandArmTransactionNodeFirewallRule(firewallRules),
            Password: utils.String(password),
        },
        Value: utils.String(value),
    }


    if _, err := client.Update(ctx, resourceGroup, blockchainMemberName, name, transactionNode); err != nil {
        return fmt.Errorf("Error updating Transaction Node %q (Resource Group %q / Blockchain Member Name %q): %+v", name, resourceGroup, blockchainMemberName, err)
    }

    return resourceArmTransactionNodeRead(d, meta)
}

func resourceArmTransactionNodeDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).transactionNodesClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    blockchainMemberName := id.Path["blockchainMembers"]
    name := id.Path["transactionNodes"]

    future, err := client.Delete(ctx, resourceGroup, blockchainMemberName, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Transaction Node %q (Resource Group %q / Blockchain Member Name %q): %+v", name, resourceGroup, blockchainMemberName, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Transaction Node %q (Resource Group %q / Blockchain Member Name %q): %+v", name, resourceGroup, blockchainMemberName, err)
        }
    }

    return nil
}

func expandArmTransactionNodeFirewallRule(input []interface{}) *[]blockchain.FirewallRule {
    results := make([]blockchain.FirewallRule, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        ruleName := v["rule_name"].(string)
        startIpAddress := v["start_ip_address"].(string)
        endIpAddress := v["end_ip_address"].(string)

        result := blockchain.FirewallRule{
            EndIPAddress: utils.String(endIpAddress),
            RuleName: utils.String(ruleName),
            StartIPAddress: utils.String(startIpAddress),
        }

        results = append(results, result)
    }
    return &results
}


func flattenArmTransactionNodeFirewallRule(input *[]blockchain.FirewallRule) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        if endIpAddress := item.EndIPAddress; endIpAddress != nil {
            v["end_ip_address"] = *endIpAddress
        }
        if ruleName := item.RuleName; ruleName != nil {
            v["rule_name"] = *ruleName
        }
        if startIpAddress := item.StartIPAddress; startIpAddress != nil {
            v["start_ip_address"] = *startIpAddress
        }

        results = append(results, v)
    }

    return results
}
