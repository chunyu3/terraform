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



func resourceArmServiceUnit() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmServiceUnitCreateUpdate,
        Read: resourceArmServiceUnitRead,
        Update: resourceArmServiceUnitCreateUpdate,
        Delete: resourceArmServiceUnitDelete,

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

            "deployment_mode": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(azuredeploymentmanager.Incremental),
                    string(azuredeploymentmanager.Complete),
                }, false),
            },

            "service_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "service_topology_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "target_resource_group": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "artifacts": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "parameters_artifact_source_relative_path": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "parameters_uri": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "template_artifact_source_relative_path": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "template_uri": {
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

            "tags": tags.Schema(),
        },
    }
}

func resourceArmServiceUnitCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).serviceUnitsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    serviceName := d.Get("service_name").(string)
    serviceTopologyName := d.Get("service_topology_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, serviceTopologyName, serviceName, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Service Unit %q (Service Name %q / Service Topology Name %q / Resource Group %q): %+v", name, serviceName, serviceTopologyName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_service_unit", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    artifacts := d.Get("artifacts").([]interface{})
    deploymentMode := d.Get("deployment_mode").(string)
    targetResourceGroup := d.Get("target_resource_group").(string)
    t := d.Get("tags").(map[string]interface{})

    serviceUnitInfo := azuredeploymentmanager.ServiceUnitResource{
        Location: utils.String(location),
        ServiceUnitResource_properties: &azuredeploymentmanager.ServiceUnitResource_properties{
            Artifacts: expandArmServiceUnitServiceUnitArtifacts(artifacts),
            DeploymentMode: azuredeploymentmanager.DeploymentMode(deploymentMode),
            TargetResourceGroup: utils.String(targetResourceGroup),
        },
        Tags: tags.Expand(t),
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, serviceTopologyName, serviceName, name, serviceUnitInfo)
    if err != nil {
        return fmt.Errorf("Error creating Service Unit %q (Service Name %q / Service Topology Name %q / Resource Group %q): %+v", name, serviceName, serviceTopologyName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Service Unit %q (Service Name %q / Service Topology Name %q / Resource Group %q): %+v", name, serviceName, serviceTopologyName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, serviceTopologyName, serviceName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Service Unit %q (Service Name %q / Service Topology Name %q / Resource Group %q): %+v", name, serviceName, serviceTopologyName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Service Unit %q (Service Name %q / Service Topology Name %q / Resource Group %q) ID", name, serviceName, serviceTopologyName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmServiceUnitRead(d, meta)
}

func resourceArmServiceUnitRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).serviceUnitsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    serviceTopologyName := id.Path["serviceTopologies"]
    serviceName := id.Path["services"]
    name := id.Path["serviceUnits"]

    resp, err := client.Get(ctx, resourceGroup, serviceTopologyName, serviceName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Service Unit %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Service Unit %q (Service Name %q / Service Topology Name %q / Resource Group %q): %+v", name, serviceName, serviceTopologyName, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if serviceUnitResourceProperties := resp.ServiceUnitResource_properties; serviceUnitResourceProperties != nil {
        if err := d.Set("artifacts", flattenArmServiceUnitServiceUnitArtifacts(serviceUnitResourceProperties.Artifacts)); err != nil {
            return fmt.Errorf("Error setting `artifacts`: %+v", err)
        }
        d.Set("deployment_mode", string(serviceUnitResourceProperties.DeploymentMode))
        d.Set("target_resource_group", serviceUnitResourceProperties.TargetResourceGroup)
    }
    d.Set("service_name", serviceName)
    d.Set("service_topology_name", serviceTopologyName)
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}


func resourceArmServiceUnitDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).serviceUnitsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    serviceTopologyName := id.Path["serviceTopologies"]
    serviceName := id.Path["services"]
    name := id.Path["serviceUnits"]

    if _, err := client.Delete(ctx, resourceGroup, serviceTopologyName, serviceName, name); err != nil {
        return fmt.Errorf("Error deleting Service Unit %q (Service Name %q / Service Topology Name %q / Resource Group %q): %+v", name, serviceName, serviceTopologyName, resourceGroup, err)
    }

    return nil
}

func expandArmServiceUnitServiceUnitArtifacts(input []interface{}) *azuredeploymentmanager.ServiceUnitArtifacts {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    templateUri := v["template_uri"].(string)
    parametersUri := v["parameters_uri"].(string)
    templateArtifactSourceRelativePath := v["template_artifact_source_relative_path"].(string)
    parametersArtifactSourceRelativePath := v["parameters_artifact_source_relative_path"].(string)

    result := azuredeploymentmanager.ServiceUnitArtifacts{
        ParametersArtifactSourceRelativePath: utils.String(parametersArtifactSourceRelativePath),
        ParametersUri: utils.String(parametersUri),
        TemplateArtifactSourceRelativePath: utils.String(templateArtifactSourceRelativePath),
        TemplateUri: utils.String(templateUri),
    }
    return &result
}


func flattenArmServiceUnitServiceUnitArtifacts(input *azuredeploymentmanager.ServiceUnitArtifacts) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if parametersArtifactSourceRelativePath := input.ParametersArtifactSourceRelativePath; parametersArtifactSourceRelativePath != nil {
        result["parameters_artifact_source_relative_path"] = *parametersArtifactSourceRelativePath
    }
    if parametersUri := input.ParametersUri; parametersUri != nil {
        result["parameters_uri"] = *parametersUri
    }
    if templateArtifactSourceRelativePath := input.TemplateArtifactSourceRelativePath; templateArtifactSourceRelativePath != nil {
        result["template_artifact_source_relative_path"] = *templateArtifactSourceRelativePath
    }
    if templateUri := input.TemplateUri; templateUri != nil {
        result["template_uri"] = *templateUri
    }

    return []interface{}{result}
}
