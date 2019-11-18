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



func resourceArmService() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmServiceCreate,
        Read: resourceArmServiceRead,
        Update: resourceArmServiceUpdate,
        Delete: resourceArmServiceDelete,

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

            "application_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "cluster_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "correlation_scheme": {
                Type: schema.TypeList,
                Optional: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "scheme": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(servicefabric.Invalid),
                                string(servicefabric.Affinity),
                                string(servicefabric.AlignedAffinity),
                                string(servicefabric.NonAlignedAffinity),
                            }, false),
                        },
                        "service_name": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                    },
                },
            },

            "default_move_cost": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(servicefabric.Zero),
                    string(servicefabric.Low),
                    string(servicefabric.Medium),
                    string(servicefabric.High),
                }, false),
                Default: string(servicefabric.Zero),
            },

            "placement_constraints": {
                Type: schema.TypeString,
                Optional: true,
            },

            "service_load_metrics": {
                Type: schema.TypeList,
                Optional: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "name": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "default_load": {
                            Type: schema.TypeInt,
                            Optional: true,
                        },
                        "primary_default_load": {
                            Type: schema.TypeInt,
                            Optional: true,
                        },
                        "secondary_default_load": {
                            Type: schema.TypeInt,
                            Optional: true,
                        },
                        "weight": {
                            Type: schema.TypeString,
                            Optional: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(servicefabric.Zero),
                                string(servicefabric.Low),
                                string(servicefabric.Medium),
                                string(servicefabric.High),
                            }, false),
                            Default: string(servicefabric.Zero),
                        },
                    },
                },
            },

            "etag": {
                Type: schema.TypeString,
                Computed: true,
            },

            "provisioning_state": {
                Type: schema.TypeString,
                Computed: true,
            },

            "service_package_activation_mode": {
                Type: schema.TypeString,
                Computed: true,
            },

            "service_type_name": {
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

func resourceArmServiceCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).servicesClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    applicationName := d.Get("application_name").(string)
    clusterName := d.Get("cluster_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, clusterName, applicationName, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Service %q (Application Name %q / Cluster Name %q / Resource Group %q): %+v", name, applicationName, clusterName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_service", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    correlationScheme := d.Get("correlation_scheme").([]interface{})
    defaultMoveCost := d.Get("default_move_cost").(string)
    placementConstraints := d.Get("placement_constraints").(string)
    serviceLoadMetrics := d.Get("service_load_metrics").([]interface{})
    t := d.Get("tags").(map[string]interface{})

    parameters := servicefabric.ServiceResourceUpdate{
        Location: utils.String(location),
        ServiceResourceUpdateProperties: &servicefabric.ServiceResourceUpdateProperties{
            CorrelationScheme: expandArmServiceServiceCorrelationDescription(correlationScheme),
            DefaultMoveCost: servicefabric.MoveCost(defaultMoveCost),
            PlacementConstraints: utils.String(placementConstraints),
            ServiceLoadMetrics: expandArmServiceServiceLoadMetricDescription(serviceLoadMetrics),
        },
        Tags: tags.Expand(t),
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, clusterName, applicationName, name, parameters)
    if err != nil {
        return fmt.Errorf("Error creating Service %q (Application Name %q / Cluster Name %q / Resource Group %q): %+v", name, applicationName, clusterName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Service %q (Application Name %q / Cluster Name %q / Resource Group %q): %+v", name, applicationName, clusterName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, clusterName, applicationName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Service %q (Application Name %q / Cluster Name %q / Resource Group %q): %+v", name, applicationName, clusterName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Service %q (Application Name %q / Cluster Name %q / Resource Group %q) ID", name, applicationName, clusterName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmServiceRead(d, meta)
}

func resourceArmServiceRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).servicesClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    clusterName := id.Path["clusters"]
    applicationName := id.Path["applications"]
    name := id.Path["services"]

    resp, err := client.Get(ctx, resourceGroup, clusterName, applicationName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Service %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Service %q (Application Name %q / Cluster Name %q / Resource Group %q): %+v", name, applicationName, clusterName, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    d.Set("application_name", applicationName)
    d.Set("cluster_name", clusterName)
    if serviceResourceUpdateProperties := resp.ServiceResourceUpdateProperties; serviceResourceUpdateProperties != nil {
        if err := d.Set("correlation_scheme", flattenArmServiceServiceCorrelationDescription(serviceResourceUpdateProperties.CorrelationScheme)); err != nil {
            return fmt.Errorf("Error setting `correlation_scheme`: %+v", err)
        }
        d.Set("default_move_cost", string(serviceResourceUpdateProperties.DefaultMoveCost))
        d.Set("placement_constraints", serviceResourceUpdateProperties.PlacementConstraints)
        d.Set("provisioning_state", serviceResourceUpdateProperties.ProvisioningState)
        if err := d.Set("service_load_metrics", flattenArmServiceServiceLoadMetricDescription(serviceResourceUpdateProperties.ServiceLoadMetrics)); err != nil {
            return fmt.Errorf("Error setting `service_load_metrics`: %+v", err)
        }
        d.Set("service_package_activation_mode", string(serviceResourceUpdateProperties.ServicePackageActivationMode))
        d.Set("service_type_name", serviceResourceUpdateProperties.ServiceTypeName)
    }
    d.Set("etag", resp.Etag)
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}

func resourceArmServiceUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).servicesClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    applicationName := d.Get("application_name").(string)
    clusterName := d.Get("cluster_name").(string)
    correlationScheme := d.Get("correlation_scheme").([]interface{})
    defaultMoveCost := d.Get("default_move_cost").(string)
    placementConstraints := d.Get("placement_constraints").(string)
    serviceLoadMetrics := d.Get("service_load_metrics").([]interface{})
    t := d.Get("tags").(map[string]interface{})

    parameters := servicefabric.ServiceResourceUpdate{
        Location: utils.String(location),
        ServiceResourceUpdateProperties: &servicefabric.ServiceResourceUpdateProperties{
            CorrelationScheme: expandArmServiceServiceCorrelationDescription(correlationScheme),
            DefaultMoveCost: servicefabric.MoveCost(defaultMoveCost),
            PlacementConstraints: utils.String(placementConstraints),
            ServiceLoadMetrics: expandArmServiceServiceLoadMetricDescription(serviceLoadMetrics),
        },
        Tags: tags.Expand(t),
    }


    future, err := client.Update(ctx, resourceGroup, clusterName, applicationName, name, parameters)
    if err != nil {
        return fmt.Errorf("Error updating Service %q (Application Name %q / Cluster Name %q / Resource Group %q): %+v", name, applicationName, clusterName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for update of Service %q (Application Name %q / Cluster Name %q / Resource Group %q): %+v", name, applicationName, clusterName, resourceGroup, err)
    }

    return resourceArmServiceRead(d, meta)
}

func resourceArmServiceDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).servicesClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    clusterName := id.Path["clusters"]
    applicationName := id.Path["applications"]
    name := id.Path["services"]

    future, err := client.Delete(ctx, resourceGroup, clusterName, applicationName, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Service %q (Application Name %q / Cluster Name %q / Resource Group %q): %+v", name, applicationName, clusterName, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Service %q (Application Name %q / Cluster Name %q / Resource Group %q): %+v", name, applicationName, clusterName, resourceGroup, err)
        }
    }

    return nil
}

func expandArmServiceServiceCorrelationDescription(input []interface{}) *[]servicefabric.ServiceCorrelationDescription {
    results := make([]servicefabric.ServiceCorrelationDescription, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        scheme := v["scheme"].(string)
        serviceName := v["service_name"].(string)

        result := servicefabric.ServiceCorrelationDescription{
            Scheme: servicefabric.ServiceCorrelationScheme(scheme),
            ServiceName: utils.String(serviceName),
        }

        results = append(results, result)
    }
    return &results
}

func expandArmServiceServiceLoadMetricDescription(input []interface{}) *[]servicefabric.ServiceLoadMetricDescription {
    results := make([]servicefabric.ServiceLoadMetricDescription, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        name := v["name"].(string)
        weight := v["weight"].(string)
        primaryDefaultLoad := v["primary_default_load"].(int)
        secondaryDefaultLoad := v["secondary_default_load"].(int)
        defaultLoad := v["default_load"].(int)

        result := servicefabric.ServiceLoadMetricDescription{
            DefaultLoad: utils.Int(defaultLoad),
            Name: utils.String(name),
            PrimaryDefaultLoad: utils.Int(primaryDefaultLoad),
            SecondaryDefaultLoad: utils.Int(secondaryDefaultLoad),
            Weight: servicefabric.ServiceLoadMetricWeight(weight),
        }

        results = append(results, result)
    }
    return &results
}


func flattenArmServiceServiceCorrelationDescription(input *[]servicefabric.ServiceCorrelationDescription) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        v["scheme"] = string(item.Scheme)
        if serviceName := item.ServiceName; serviceName != nil {
            v["service_name"] = *serviceName
        }

        results = append(results, v)
    }

    return results
}

func flattenArmServiceServiceLoadMetricDescription(input *[]servicefabric.ServiceLoadMetricDescription) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        if name := item.Name; name != nil {
            v["name"] = *name
        }
        if defaultLoad := item.DefaultLoad; defaultLoad != nil {
            v["default_load"] = *defaultLoad
        }
        if primaryDefaultLoad := item.PrimaryDefaultLoad; primaryDefaultLoad != nil {
            v["primary_default_load"] = *primaryDefaultLoad
        }
        if secondaryDefaultLoad := item.SecondaryDefaultLoad; secondaryDefaultLoad != nil {
            v["secondary_default_load"] = *secondaryDefaultLoad
        }
        v["weight"] = string(item.Weight)

        results = append(results, v)
    }

    return results
}
