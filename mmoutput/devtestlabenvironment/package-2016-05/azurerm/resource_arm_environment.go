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



func resourceArmEnvironment() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmEnvironmentCreateUpdate,
        Read: resourceArmEnvironmentRead,
        Update: resourceArmEnvironmentCreateUpdate,
        Delete: resourceArmEnvironmentDelete,

        Importer: &schema.ResourceImporter{
            State: schema.ImportStatePassthrough,
        },


        Schema: map[string]*schema.Schema{
            "lab_name": {
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

            "user_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "arm_template_display_name": {
                Type: schema.TypeString,
                Optional: true,
            },

            "deployment_properties": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "arm_template_id": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "parameters": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "name": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                    "value": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                },
                            },
                        },
                    },
                },
            },

            "location": azure.SchemaLocation(),

            "tags": tags.Schema(),

            "unique_identifier": {
                Type: schema.TypeString,
                Optional: true,
            },

            "created_by_user": {
                Type: schema.TypeString,
                Computed: true,
            },

            "id": {
                Type: schema.TypeString,
                Computed: true,
            },

            "name": {
                Type: schema.TypeString,
                Computed: true,
            },

            "provisioning_state": {
                Type: schema.TypeString,
                Computed: true,
            },

            "resource_group_id": {
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

func resourceArmEnvironmentCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).environmentsClient
    ctx, cancel := timeouts.ForCreateUpdate(meta.(*ArmClient).StopContext, d)
    defer cancel()

    resourceGroupName := d.Get("resource_group").(string)
    labName := d.Get("lab_name").(string)
    name := d.Get("name").(string)
    name := d.Get("user_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroupName, labName, name, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Environment (Name %q / User Name %q / Lab Name %q / Resource Group %q): %+v", name, name, labName, resourceGroupName, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_environment", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    armTemplateDisplayName := d.Get("arm_template_display_name").(string)
    deploymentProperties := d.Get("deployment_properties").([]interface{})
    uniqueIdentifier := d.Get("unique_identifier").(string)
    tags := d.Get("tags").(map[string]interface{})

    dtlEnvironment := devtestlab.DtlEnvironment{
        Location: utils.String(location),
        EnvironmentProperties: &devtestlab.EnvironmentProperties{
            ArmTemplateDisplayName: utils.String(armTemplateDisplayName),
            DeploymentProperties: expandArmEnvironmentEnvironmentDeploymentProperties(deploymentProperties),
            UniqueIdentifier: utils.String(uniqueIdentifier),
        },
        Tags: tags.Expand(tags),
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroupName, labName, name, name, dtlEnvironment)
    if err != nil {
        return fmt.Errorf("Error creating Environment (Name %q / User Name %q / Lab Name %q / Resource Group %q): %+v", name, name, labName, resourceGroupName, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Environment (Name %q / User Name %q / Lab Name %q / Resource Group %q): %+v", name, name, labName, resourceGroupName, err)
    }


    resp, err := client.Get(ctx, resourceGroupName, labName, name, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Environment (Name %q / User Name %q / Lab Name %q / Resource Group %q): %+v", name, name, labName, resourceGroupName, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Environment (Name %q / User Name %q / Lab Name %q / Resource Group %q) ID", name, name, labName, resourceGroupName)
    }
    d.SetId(*resp.ID)

    return resourceArmEnvironmentRead(d, meta)
}

func resourceArmEnvironmentRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).environmentsClient
    ctx, cancel := timeouts.ForRead(meta.(*ArmClient).StopContext, d)
    defer cancel()

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroupName := id.ResourceGroup
    labName := id.Path["labs"]
    name := id.Path["users"]
    name := id.Path["environments"]

    resp, err := client.Get(ctx, resourceGroupName, labName, name, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Environment %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Environment (Name %q / User Name %q / Lab Name %q / Resource Group %q): %+v", name, name, labName, resourceGroupName, err)
    }


    d.Set("resource_group", resourceGroupName)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if environmentProperties := resp.EnvironmentProperties; environmentProperties != nil {
        d.Set("arm_template_display_name", environmentProperties.ArmTemplateDisplayName)
        d.Set("created_by_user", environmentProperties.CreatedByUser)
        if err := d.Set("deployment_properties", flattenArmEnvironmentEnvironmentDeploymentProperties(environmentProperties.DeploymentProperties)); err != nil {
            return fmt.Errorf("Error setting `deployment_properties`: %+v", err)
        }
        d.Set("provisioning_state", environmentProperties.ProvisioningState)
        d.Set("resource_group_id", environmentProperties.ResourceGroupID)
        d.Set("unique_identifier", environmentProperties.UniqueIdentifier)
    }
    d.Set("id", resp.ID)
    d.Set("lab_name", labName)
    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("type", resp.Type)
    d.Set("user_name", name)

    return tags.FlattenAndSet(d, resp.Tags)
}


func resourceArmEnvironmentDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).environmentsClient
    ctx, cancel := timeouts.ForDelete(meta.(*ArmClient).StopContext, d)
    defer cancel()


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroupName := id.ResourceGroup
    labName := id.Path["labs"]
    name := id.Path["users"]
    name := id.Path["environments"]

    future, err := client.Delete(ctx, resourceGroupName, labName, name, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Environment (Name %q / User Name %q / Lab Name %q / Resource Group %q): %+v", name, name, labName, resourceGroupName, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Environment (Name %q / User Name %q / Lab Name %q / Resource Group %q): %+v", name, name, labName, resourceGroupName, err)
        }
    }

    return nil
}

func expandArmEnvironmentEnvironmentDeploymentProperties(input []interface{}) *devtestlab.EnvironmentDeploymentProperties {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    armTemplateID := v["arm_template_id"].(string)
    parameters := v["parameters"].([]interface{})

    result := devtestlab.EnvironmentDeploymentProperties{
        ArmTemplateID: utils.String(armTemplateID),
        Parameters: expandArmEnvironmentArmTemplateParameterProperties(parameters),
    }
    return &result
}

func expandArmEnvironmentArmTemplateParameterProperties(input []interface{}) *[]devtestlab.ArmTemplateParameterProperties {
    results := make([]devtestlab.ArmTemplateParameterProperties, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        name := v["name"].(string)
        value := v["value"].(string)

        result := devtestlab.ArmTemplateParameterProperties{
            Name: utils.String(name),
            Value: utils.String(value),
        }

        results = append(results, result)
    }
    return &results
}


func flattenArmEnvironmentEnvironmentDeploymentProperties(input *devtestlab.EnvironmentDeploymentProperties) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if armTemplateId := input.ArmTemplateID; armTemplateId != nil {
        result["arm_template_id"] = *armTemplateId
    }
    result["parameters"] = flattenArmEnvironmentArmTemplateParameterProperties(input.Parameters)

    return []interface{}{result}
}

func flattenArmEnvironmentArmTemplateParameterProperties(input *[]devtestlab.ArmTemplateParameterProperties) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        if name := item.Name; name != nil {
            v["name"] = *name
        }
        if value := item.Value; value != nil {
            v["value"] = *value
        }

        results = append(results, v)
    }

    return results
}
