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
            "name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "location": azure.SchemaLocation(),

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "lab_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

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

            "unique_identifier": {
                Type: schema.TypeString,
                Optional: true,
            },

            "created_by_user": {
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

            "tags": tags.Schema(),
        },
    }
}

func resourceArmEnvironmentCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).environmentsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    labName := d.Get("lab_name").(string)
    userName := d.Get("user_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, labName, userName, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Environment %q (User Name %q / Lab Name %q / Resource Group %q): %+v", name, userName, labName, resourceGroup, err)
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
    t := d.Get("tags").(map[string]interface{})

    dtlEnvironment := devtestlab.DtlEnvironment{
        Location: utils.String(location),
        EnvironmentProperties: &devtestlab.EnvironmentProperties{
            ArmTemplateDisplayName: utils.String(armTemplateDisplayName),
            DeploymentProperties: expandArmEnvironmentEnvironmentDeploymentProperties(deploymentProperties),
            UniqueIdentifier: utils.String(uniqueIdentifier),
        },
        Tags: tags.Expand(t),
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, labName, userName, name, dtlEnvironment)
    if err != nil {
        return fmt.Errorf("Error creating Environment %q (User Name %q / Lab Name %q / Resource Group %q): %+v", name, userName, labName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Environment %q (User Name %q / Lab Name %q / Resource Group %q): %+v", name, userName, labName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, labName, userName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Environment %q (User Name %q / Lab Name %q / Resource Group %q): %+v", name, userName, labName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Environment %q (User Name %q / Lab Name %q / Resource Group %q) ID", name, userName, labName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmEnvironmentRead(d, meta)
}

func resourceArmEnvironmentRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).environmentsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    labName := id.Path["labs"]
    userName := id.Path["users"]
    name := id.Path["environments"]

    resp, err := client.Get(ctx, resourceGroup, labName, userName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Environment %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Environment %q (User Name %q / Lab Name %q / Resource Group %q): %+v", name, userName, labName, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("resource_group", resourceGroup)
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
    d.Set("lab_name", labName)
    d.Set("type", resp.Type)
    d.Set("user_name", userName)

    return tags.FlattenAndSet(d, resp.Tags)
}


func resourceArmEnvironmentDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).environmentsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    labName := id.Path["labs"]
    userName := id.Path["users"]
    name := id.Path["environments"]

    future, err := client.Delete(ctx, resourceGroup, labName, userName, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Environment %q (User Name %q / Lab Name %q / Resource Group %q): %+v", name, userName, labName, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Environment %q (User Name %q / Lab Name %q / Resource Group %q): %+v", name, userName, labName, resourceGroup, err)
        }
    }

    return nil
}

func expandArmEnvironmentEnvironmentDeploymentProperties(input []interface{}) *devtestlab.EnvironmentDeploymentProperties {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    armTemplateId := v["arm_template_id"].(string)
    parameters := v["parameters"].([]interface{})

    result := devtestlab.EnvironmentDeploymentProperties{
        ArmTemplateID: utils.String(armTemplateId),
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