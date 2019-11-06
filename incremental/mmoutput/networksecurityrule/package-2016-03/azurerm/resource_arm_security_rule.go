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



func resourceArmSecurityRule() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmSecurityRuleCreateUpdate,
        Read: resourceArmSecurityRuleRead,
        Update: resourceArmSecurityRuleCreateUpdate,
        Delete: resourceArmSecurityRuleDelete,

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

            "access": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(network.Allow),
                    string(network.Deny),
                }, false),
            },

            "destination_address_prefix": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "direction": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(network.Inbound),
                    string(network.Outbound),
                }, false),
            },

            "network_security_group_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "protocol": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(network.Tcp),
                    string(network.Udp),
                    string(network.*),
                }, false),
            },

            "source_address_prefix": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "description": {
                Type: schema.TypeString,
                Optional: true,
            },

            "destination_port_range": {
                Type: schema.TypeString,
                Optional: true,
            },

            "etag": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
            },

            "priority": {
                Type: schema.TypeInt,
                Optional: true,
            },

            "source_port_range": {
                Type: schema.TypeString,
                Optional: true,
            },

            "provisioning_state": {
                Type: schema.TypeString,
                Computed: true,
            },
        },
    }
}

func resourceArmSecurityRuleCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).securityRulesClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    networkSecurityGroupName := d.Get("network_security_group_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, networkSecurityGroupName, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Security Rule %q (Network Security Group Name %q / Resource Group %q): %+v", name, networkSecurityGroupName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_security_rule", *existing.ID)
        }
    }

    id := d.Get("id").(string)
    name := d.Get("name").(string)
    access := d.Get("access").(string)
    description := d.Get("description").(string)
    destinationAddressPrefix := d.Get("destination_address_prefix").(string)
    destinationPortRange := d.Get("destination_port_range").(string)
    direction := d.Get("direction").(string)
    etag := d.Get("etag").(string)
    priority := d.Get("priority").(int)
    protocol := d.Get("protocol").(string)
    sourceAddressPrefix := d.Get("source_address_prefix").(string)
    sourcePortRange := d.Get("source_port_range").(string)

    securityRuleParameters := network.SecurityRule{
        Etag: utils.String(etag),
        ID: utils.String(id),
        Name: utils.String(name),
        SecurityRulePropertiesFormat: &network.SecurityRulePropertiesFormat{
            Access: network.SecurityRuleAccess(access),
            Description: utils.String(description),
            DestinationAddressPrefix: utils.String(destinationAddressPrefix),
            DestinationPortRange: utils.String(destinationPortRange),
            Direction: network.SecurityRuleDirection(direction),
            Priority: utils.Int32(int32(priority)),
            Protocol: network.SecurityRuleProtocol(protocol),
            SourceAddressPrefix: utils.String(sourceAddressPrefix),
            SourcePortRange: utils.String(sourcePortRange),
        },
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, networkSecurityGroupName, name, securityRuleParameters)
    if err != nil {
        return fmt.Errorf("Error creating Security Rule %q (Network Security Group Name %q / Resource Group %q): %+v", name, networkSecurityGroupName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Security Rule %q (Network Security Group Name %q / Resource Group %q): %+v", name, networkSecurityGroupName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, networkSecurityGroupName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Security Rule %q (Network Security Group Name %q / Resource Group %q): %+v", name, networkSecurityGroupName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Security Rule %q (Network Security Group Name %q / Resource Group %q) ID", name, networkSecurityGroupName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmSecurityRuleRead(d, meta)
}

func resourceArmSecurityRuleRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).securityRulesClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    networkSecurityGroupName := id.Path["networkSecurityGroups"]
    name := id.Path["securityRules"]

    resp, err := client.Get(ctx, resourceGroup, networkSecurityGroupName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Security Rule %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Security Rule %q (Network Security Group Name %q / Resource Group %q): %+v", name, networkSecurityGroupName, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if securityRulePropertiesFormat := resp.SecurityRulePropertiesFormat; securityRulePropertiesFormat != nil {
        d.Set("access", string(securityRulePropertiesFormat.Access))
        d.Set("description", securityRulePropertiesFormat.Description)
        d.Set("destination_address_prefix", securityRulePropertiesFormat.DestinationAddressPrefix)
        d.Set("destination_port_range", securityRulePropertiesFormat.DestinationPortRange)
        d.Set("direction", string(securityRulePropertiesFormat.Direction))
        d.Set("priority", int(*securityRulePropertiesFormat.Priority))
        d.Set("protocol", string(securityRulePropertiesFormat.Protocol))
        d.Set("provisioning_state", securityRulePropertiesFormat.ProvisioningState)
        d.Set("source_address_prefix", securityRulePropertiesFormat.SourceAddressPrefix)
        d.Set("source_port_range", securityRulePropertiesFormat.SourcePortRange)
    }
    d.Set("etag", resp.Etag)
    d.Set("network_security_group_name", networkSecurityGroupName)

    return nil
}


func resourceArmSecurityRuleDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).securityRulesClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    networkSecurityGroupName := id.Path["networkSecurityGroups"]
    name := id.Path["securityRules"]

    future, err := client.Delete(ctx, resourceGroup, networkSecurityGroupName, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Security Rule %q (Network Security Group Name %q / Resource Group %q): %+v", name, networkSecurityGroupName, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Security Rule %q (Network Security Group Name %q / Resource Group %q): %+v", name, networkSecurityGroupName, resourceGroup, err)
        }
    }

    return nil
}
