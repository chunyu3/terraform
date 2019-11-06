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



func resourceArmDscNodeConfiguration() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmDscNodeConfigurationCreateUpdate,
        Read: resourceArmDscNodeConfigurationRead,
        Update: resourceArmDscNodeConfigurationCreateUpdate,
        Delete: resourceArmDscNodeConfigurationDelete,

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
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "automation_account_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "configuration": {
                Type: schema.TypeList,
                Required: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "name": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },

            "source": {
                Type: schema.TypeList,
                Required: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "hash": {
                            Type: schema.TypeList,
                            Optional: true,
                            MaxItems: 1,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "algorithm": {
                                        Type: schema.TypeString,
                                        Required: true,
                                        ValidateFunc: validate.NoEmptyStrings,
                                    },
                                    "value": {
                                        Type: schema.TypeString,
                                        Required: true,
                                        ValidateFunc: validate.NoEmptyStrings,
                                    },
                                },
                            },
                        },
                        "type": {
                            Type: schema.TypeString,
                            Optional: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(automation.embeddedContent),
                                string(automation.uri),
                            }, false),
                            Default: string(automation.embeddedContent),
                        },
                        "value": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "version": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },

            "increment_node_configuration_build": {
                Type: schema.TypeBool,
                Optional: true,
                ForceNew: true,
            },

            "creation_time": {
                Type: schema.TypeString,
                Computed: true,
            },

            "last_modified_time": {
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

func resourceArmDscNodeConfigurationCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).dscNodeConfigurationClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    automationAccountName := d.Get("automation_account_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, automationAccountName, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Dsc Node Configuration %q (Automation Account Name %q / Resource Group %q): %+v", name, automationAccountName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_dsc_node_configuration", *existing.ID)
        }
    }

    name := d.Get("name").(string)
    configuration := d.Get("configuration").([]interface{})
    incrementNodeConfigurationBuild := d.Get("increment_node_configuration_build").(bool)
    source := d.Get("source").([]interface{})

    parameters := automation.DscNodeConfigurationCreateOrUpdateParameters{
        Configuration: expandArmDscNodeConfigurationDscConfigurationAssociationProperty(configuration),
        IncrementNodeConfigurationBuild: utils.Bool(incrementNodeConfigurationBuild),
        Name: utils.String(name),
        Source: expandArmDscNodeConfigurationContentSource(source),
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroup, automationAccountName, name, parameters); err != nil {
        return fmt.Errorf("Error creating Dsc Node Configuration %q (Automation Account Name %q / Resource Group %q): %+v", name, automationAccountName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, automationAccountName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Dsc Node Configuration %q (Automation Account Name %q / Resource Group %q): %+v", name, automationAccountName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Dsc Node Configuration %q (Automation Account Name %q / Resource Group %q) ID", name, automationAccountName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmDscNodeConfigurationRead(d, meta)
}

func resourceArmDscNodeConfigurationRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).dscNodeConfigurationClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    automationAccountName := id.Path["automationAccounts"]
    name := id.Path["nodeConfigurations"]

    resp, err := client.Get(ctx, resourceGroup, automationAccountName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Dsc Node Configuration %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Dsc Node Configuration %q (Automation Account Name %q / Resource Group %q): %+v", name, automationAccountName, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    d.Set("automation_account_name", automationAccountName)
    if err := d.Set("configuration", flattenArmDscNodeConfigurationDscConfigurationAssociationProperty(resp.Configuration)); err != nil {
        return fmt.Errorf("Error setting `configuration`: %+v", err)
    }
    d.Set("creation_time", (resp.CreationTime).String())
    d.Set("last_modified_time", (resp.LastModifiedTime).String())
    d.Set("type", resp.Type)

    return nil
}


func resourceArmDscNodeConfigurationDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).dscNodeConfigurationClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    automationAccountName := id.Path["automationAccounts"]
    name := id.Path["nodeConfigurations"]

    if _, err := client.Delete(ctx, resourceGroup, automationAccountName, name); err != nil {
        return fmt.Errorf("Error deleting Dsc Node Configuration %q (Automation Account Name %q / Resource Group %q): %+v", name, automationAccountName, resourceGroup, err)
    }

    return nil
}

func expandArmDscNodeConfigurationDscConfigurationAssociationProperty(input []interface{}) *automation.DscConfigurationAssociationProperty {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    name := v["name"].(string)

    result := automation.DscConfigurationAssociationProperty{
        Name: utils.String(name),
    }
    return &result
}

func expandArmDscNodeConfigurationContentSource(input []interface{}) *automation.ContentSource {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    hash := v["hash"].([]interface{})
    type := v["type"].(string)
    value := v["value"].(string)
    version := v["version"].(string)

    result := automation.ContentSource{
        Hash: expandArmDscNodeConfigurationContentHash(hash),
        Type: automation.ContentSourceType(type),
        Value: utils.String(value),
        Version: utils.String(version),
    }
    return &result
}

func expandArmDscNodeConfigurationContentHash(input []interface{}) *automation.ContentHash {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    algorithm := v["algorithm"].(string)
    value := v["value"].(string)

    result := automation.ContentHash{
        Algorithm: utils.String(algorithm),
        Value: utils.String(value),
    }
    return &result
}


func flattenArmDscNodeConfigurationDscConfigurationAssociationProperty(input *automation.DscConfigurationAssociationProperty) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if name := input.Name; name != nil {
        result["name"] = *name
    }

    return []interface{}{result}
}
