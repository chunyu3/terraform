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



func resourceArmRouteFilterRule() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmRouteFilterRuleCreate,
        Read: resourceArmRouteFilterRuleRead,
        Update: resourceArmRouteFilterRuleUpdate,
        Delete: resourceArmRouteFilterRuleDelete,

        Importer: &schema.ResourceImporter{
            State: schema.ImportStatePassthrough,
        },


        Schema: map[string]*schema.Schema{
            "name": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
            },

            "location": azure.SchemaLocation(),

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "access": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(network.Allow),
                    string(network.Deny),
                }, false),
            },

            "communities": {
                Type: schema.TypeList,
                Required: true,
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
            },

            "route_filter_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "route_filter_rule_type": {
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

            "etag": {
                Type: schema.TypeString,
                Computed: true,
            },

            "provisioning_state": {
                Type: schema.TypeString,
                Computed: true,
            },

            "tags": tags.Schema(),
        },
    }
}

func resourceArmRouteFilterRuleCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).routeFilterRulesClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    routeFilterName := d.Get("route_filter_name").(string)
    ruleName := d.Get("rule_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, routeFilterName, ruleName)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Route Filter Rule (Rule Name %q / Route Filter Name %q / Resource Group %q): %+v", ruleName, routeFilterName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_route_filter_rule", *existing.ID)
        }
    }

    id := d.Get("id").(string)
    name := d.Get("name").(string)
    location := azure.NormalizeLocation(d.Get("location").(string))
    access := d.Get("access").(string)
    communities := d.Get("communities").([]interface{})
    routeFilterRuleType := d.Get("route_filter_rule_type").(string)
    t := d.Get("tags").(map[string]interface{})

    routeFilterRuleParameters := network.RouteFilterRule{
        ID: utils.String(id),
        Location: utils.String(location),
        Name: utils.String(name),
        RouteFilterRulePropertiesFormat: &network.RouteFilterRulePropertiesFormat{
            Access: network.Access(access),
            Communities: utils.ExpandStringSlice(communities),
            RouteFilterRuleType: utils.String(routeFilterRuleType),
        },
        Tags: tags.Expand(t),
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, routeFilterName, ruleName, routeFilterRuleParameters)
    if err != nil {
        return fmt.Errorf("Error creating Route Filter Rule (Rule Name %q / Route Filter Name %q / Resource Group %q): %+v", ruleName, routeFilterName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Route Filter Rule (Rule Name %q / Route Filter Name %q / Resource Group %q): %+v", ruleName, routeFilterName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, routeFilterName, ruleName)
    if err != nil {
        return fmt.Errorf("Error retrieving Route Filter Rule (Rule Name %q / Route Filter Name %q / Resource Group %q): %+v", ruleName, routeFilterName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Route Filter Rule (Rule Name %q / Route Filter Name %q / Resource Group %q) ID", ruleName, routeFilterName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmRouteFilterRuleRead(d, meta)
}

func resourceArmRouteFilterRuleRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).routeFilterRulesClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    routeFilterName := id.Path["routeFilters"]
    ruleName := id.Path["routeFilterRules"]

    resp, err := client.Get(ctx, resourceGroup, routeFilterName, ruleName)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Route Filter Rule %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Route Filter Rule (Rule Name %q / Route Filter Name %q / Resource Group %q): %+v", ruleName, routeFilterName, resourceGroup, err)
    }


    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if routeFilterRulePropertiesFormat := resp.RouteFilterRulePropertiesFormat; routeFilterRulePropertiesFormat != nil {
        d.Set("access", string(routeFilterRulePropertiesFormat.Access))
        d.Set("communities", utils.FlattenStringSlice(routeFilterRulePropertiesFormat.Communities))
        d.Set("provisioning_state", routeFilterRulePropertiesFormat.ProvisioningState)
        d.Set("route_filter_rule_type", routeFilterRulePropertiesFormat.RouteFilterRuleType)
    }
    d.Set("etag", resp.Etag)
    d.Set("route_filter_name", routeFilterName)
    d.Set("rule_name", ruleName)

    return tags.FlattenAndSet(d, resp.Tags)
}

func resourceArmRouteFilterRuleUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).routeFilterRulesClient
    ctx := meta.(*ArmClient).StopContext

    id := d.Get("id").(string)
    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    access := d.Get("access").(string)
    communities := d.Get("communities").([]interface{})
    routeFilterName := d.Get("route_filter_name").(string)
    routeFilterRuleType := d.Get("route_filter_rule_type").(string)
    ruleName := d.Get("rule_name").(string)
    t := d.Get("tags").(map[string]interface{})

    routeFilterRuleParameters := network.RouteFilterRule{
        ID: utils.String(id),
        Location: utils.String(location),
        Name: utils.String(name),
        RouteFilterRulePropertiesFormat: &network.RouteFilterRulePropertiesFormat{
            Access: network.Access(access),
            Communities: utils.ExpandStringSlice(communities),
            RouteFilterRuleType: utils.String(routeFilterRuleType),
        },
        Tags: tags.Expand(t),
    }


    future, err := client.Update(ctx, resourceGroup, routeFilterName, ruleName, routeFilterRuleParameters)
    if err != nil {
        return fmt.Errorf("Error updating Route Filter Rule (Rule Name %q / Route Filter Name %q / Resource Group %q): %+v", ruleName, routeFilterName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for update of Route Filter Rule (Rule Name %q / Route Filter Name %q / Resource Group %q): %+v", ruleName, routeFilterName, resourceGroup, err)
    }

    return resourceArmRouteFilterRuleRead(d, meta)
}

func resourceArmRouteFilterRuleDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).routeFilterRulesClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    routeFilterName := id.Path["routeFilters"]
    ruleName := id.Path["routeFilterRules"]

    future, err := client.Delete(ctx, resourceGroup, routeFilterName, ruleName)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Route Filter Rule (Rule Name %q / Route Filter Name %q / Resource Group %q): %+v", ruleName, routeFilterName, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Route Filter Rule (Rule Name %q / Route Filter Name %q / Resource Group %q): %+v", ruleName, routeFilterName, resourceGroup, err)
        }
    }

    return nil
}
