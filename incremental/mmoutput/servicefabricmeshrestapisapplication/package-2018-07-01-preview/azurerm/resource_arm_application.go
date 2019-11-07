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



func resourceArmApplication() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmApplicationCreateUpdate,
        Read: resourceArmApplicationRead,
        Update: resourceArmApplicationCreateUpdate,
        Delete: resourceArmApplicationDelete,

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

            "debug_params": {
                Type: schema.TypeString,
                Optional: true,
            },

            "description": {
                Type: schema.TypeString,
                Optional: true,
            },

            "diagnostics": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "default_sink_refs": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                        "enabled": {
                            Type: schema.TypeBool,
                            Optional: true,
                        },
                        "sinks": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "description": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                    "name": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                },
                            },
                        },
                    },
                },
            },

            "services": {
                Type: schema.TypeList,
                Optional: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "code_packages": {
                            Type: schema.TypeList,
                            Required: true,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "image": {
                                        Type: schema.TypeString,
                                        Required: true,
                                        ValidateFunc: validate.NoEmptyStrings,
                                    },
                                    "name": {
                                        Type: schema.TypeString,
                                        Required: true,
                                        ValidateFunc: validate.NoEmptyStrings,
                                    },
                                    "commands": {
                                        Type: schema.TypeList,
                                        Optional: true,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                        },
                                    },
                                    "entrypoint": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                },
                            },
                        },
                        "os_type": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(servicefabricmeshrestapis.Linux),
                                string(servicefabricmeshrestapis.Windows),
                            }, false),
                        },
                        "description": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "diagnostics": {
                            Type: schema.TypeList,
                            Optional: true,
                            MaxItems: 1,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "enabled": {
                                        Type: schema.TypeBool,
                                        Optional: true,
                                    },
                                    "sink_refs": {
                                        Type: schema.TypeList,
                                        Optional: true,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                        },
                                    },
                                },
                            },
                        },
                        "health_state": {
                            Type: schema.TypeString,
                            Optional: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(servicefabricmeshrestapis.Invalid),
                                string(servicefabricmeshrestapis.Ok),
                                string(servicefabricmeshrestapis.Warning),
                                string(servicefabricmeshrestapis.Error),
                                string(servicefabricmeshrestapis.Unknown),
                            }, false),
                            Default: string(servicefabricmeshrestapis.Invalid),
                        },
                        "name": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "network_refs": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "name": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                },
                            },
                        },
                        "replica_count": {
                            Type: schema.TypeInt,
                            Optional: true,
                        },
                    },
                },
            },

            "health_state": {
                Type: schema.TypeString,
                Computed: true,
            },

            "provisioning_state": {
                Type: schema.TypeString,
                Computed: true,
            },

            "service_names": {
                Type: schema.TypeList,
                Computed: true,
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
            },

            "status": {
                Type: schema.TypeString,
                Computed: true,
            },

            "status_details": {
                Type: schema.TypeString,
                Computed: true,
            },

            "type": {
                Type: schema.TypeString,
                Computed: true,
            },

            "unhealthy_evaluation": {
                Type: schema.TypeString,
                Computed: true,
            },

            "tags": tags.Schema(),
        },
    }
}

func resourceArmApplicationCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).applicationClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Application %q (Resource Group %q): %+v", name, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_application", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    debugParams := d.Get("debug_params").(string)
    description := d.Get("description").(string)
    diagnostics := d.Get("diagnostics").([]interface{})
    services := d.Get("services").([]interface{})
    t := d.Get("tags").(map[string]interface{})

    applicationResourceDescription := servicefabricmeshrestapis.ApplicationResourceDescription{
        Location: utils.String(location),
        ApplicationProperties: &servicefabricmeshrestapis.ApplicationResourceProperties{
            DebugParams: utils.String(debugParams),
            Description: utils.String(description),
            Diagnostics: expandArmApplicationDiagnosticsDescription(diagnostics),
            Services: expandArmApplicationServiceResourceDescription(services),
        },
        Tags: tags.Expand(t),
    }


    if _, err := client.Create(ctx, resourceGroup, name, applicationResourceDescription); err != nil {
        return fmt.Errorf("Error creating Application %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Application %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Application %q (Resource Group %q) ID", name, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmApplicationRead(d, meta)
}

func resourceArmApplicationRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).applicationClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["applications"]

    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Application %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Application %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if applicationProperties := resp.ApplicationProperties; applicationProperties != nil {
        d.Set("debug_params", applicationProperties.DebugParams)
        d.Set("description", applicationProperties.Description)
        if err := d.Set("diagnostics", flattenArmApplicationDiagnosticsDescription(applicationProperties.Diagnostics)); err != nil {
            return fmt.Errorf("Error setting `diagnostics`: %+v", err)
        }
        d.Set("health_state", string(applicationProperties.HealthState))
        d.Set("provisioning_state", applicationProperties.ProvisioningState)
        d.Set("service_names", utils.FlattenStringSlice(applicationProperties.ServiceNames))
        if err := d.Set("services", flattenArmApplicationServiceResourceDescription(applicationProperties.Services)); err != nil {
            return fmt.Errorf("Error setting `services`: %+v", err)
        }
        d.Set("status", string(applicationProperties.Status))
        d.Set("status_details", applicationProperties.StatusDetails)
        d.Set("unhealthy_evaluation", applicationProperties.UnhealthyEvaluation)
    }
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}


func resourceArmApplicationDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).applicationClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["applications"]

    if _, err := client.Delete(ctx, resourceGroup, name); err != nil {
        return fmt.Errorf("Error deleting Application %q (Resource Group %q): %+v", name, resourceGroup, err)
    }

    return nil
}

func expandArmApplicationDiagnosticsDescription(input []interface{}) *servicefabricmeshrestapis.DiagnosticsDescription {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    sinks := v["sinks"].([]interface{})
    enabled := v["enabled"].(bool)
    defaultSinkRefs := v["default_sink_refs"].([]interface{})

    result := servicefabricmeshrestapis.DiagnosticsDescription{
        DefaultSinkRefs: utils.ExpandStringSlice(defaultSinkRefs),
        Enabled: utils.Bool(enabled),
        Sinks: expandArmApplicationDiagnosticsSinkProperties(sinks),
    }
    return &result
}

func expandArmApplicationServiceResourceDescription(input []interface{}) *[]servicefabricmeshrestapis.ServiceResourceDescription {
    results := make([]servicefabricmeshrestapis.ServiceResourceDescription, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        name := v["name"].(string)
        osType := v["os_type"].(string)
        codePackages := v["code_packages"].([]interface{})
        networkRefs := v["network_refs"].([]interface{})
        diagnostics := v["diagnostics"].([]interface{})
        description := v["description"].(string)
        replicaCount := v["replica_count"].(int)
        healthState := v["health_state"].(string)

        result := servicefabricmeshrestapis.ServiceResourceDescription{
            Name: utils.String(name),
            ApplicationProperties: &servicefabricmeshrestapis.ServiceResourceProperties{
                CodePackages: expandArmApplicationContainerCodePackageProperties(codePackages),
                Description: utils.String(description),
                Diagnostics: expandArmApplicationDiagnosticsRef(diagnostics),
                HealthState: servicefabricmeshrestapis.HealthState(healthState),
                NetworkRefs: expandArmApplicationNetworkRef(networkRefs),
                OsType: servicefabricmeshrestapis.OperatingSystemTypes(osType),
                ReplicaCount: utils.Int(replicaCount),
            },
        }

        results = append(results, result)
    }
    return &results
}

func expandArmApplicationDiagnosticsSinkProperties(input []interface{}) *[]servicefabricmeshrestapis.DiagnosticsSinkProperties {
    results := make([]servicefabricmeshrestapis.DiagnosticsSinkProperties, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        name := v["name"].(string)
        description := v["description"].(string)

        result := servicefabricmeshrestapis.DiagnosticsSinkProperties{
            Description: utils.String(description),
            Name: utils.String(name),
        }

        results = append(results, result)
    }
    return &results
}

func expandArmApplicationContainerCodePackageProperties(input []interface{}) *[]servicefabricmeshrestapis.ContainerCodePackageProperties {
    results := make([]servicefabricmeshrestapis.ContainerCodePackageProperties, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        name := v["name"].(string)
        image := v["image"].(string)
        entrypoint := v["entrypoint"].(string)
        commands := v["commands"].([]interface{})

        result := servicefabricmeshrestapis.ContainerCodePackageProperties{
            Commands: utils.ExpandStringSlice(commands),
            Entrypoint: utils.String(entrypoint),
            Image: utils.String(image),
            Name: utils.String(name),
        }

        results = append(results, result)
    }
    return &results
}

func expandArmApplicationDiagnosticsRef(input []interface{}) *servicefabricmeshrestapis.DiagnosticsRef {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    enabled := v["enabled"].(bool)
    sinkRefs := v["sink_refs"].([]interface{})

    result := servicefabricmeshrestapis.DiagnosticsRef{
        Enabled: utils.Bool(enabled),
        SinkRefs: utils.ExpandStringSlice(sinkRefs),
    }
    return &result
}

func expandArmApplicationNetworkRef(input []interface{}) *[]servicefabricmeshrestapis.NetworkRef {
    results := make([]servicefabricmeshrestapis.NetworkRef, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        name := v["name"].(string)

        result := servicefabricmeshrestapis.NetworkRef{
            Name: utils.String(name),
        }

        results = append(results, result)
    }
    return &results
}


func flattenArmApplicationDiagnosticsDescription(input *servicefabricmeshrestapis.DiagnosticsDescription) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    result["default_sink_refs"] = utils.FlattenStringSlice(input.DefaultSinkRefs)
    if enabled := input.Enabled; enabled != nil {
        result["enabled"] = *enabled
    }
    result["sinks"] = flattenArmApplicationDiagnosticsSinkProperties(input.Sinks)

    return []interface{}{result}
}

func flattenArmApplicationServiceResourceDescription(input *[]servicefabricmeshrestapis.ServiceResourceDescription) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        if name := item.Name; name != nil {
            v["name"] = *name
        }
        if applicationProperties := item.ApplicationProperties; applicationProperties != nil {
            v["code_packages"] = flattenArmApplicationContainerCodePackageProperties(applicationProperties.CodePackages)
            if description := applicationProperties.Description; description != nil {
                v["description"] = *description
            }
            v["diagnostics"] = flattenArmApplicationDiagnosticsRef(applicationProperties.Diagnostics)
            v["health_state"] = string(applicationProperties.HealthState)
            v["network_refs"] = flattenArmApplicationNetworkRef(applicationProperties.NetworkRefs)
            v["os_type"] = string(applicationProperties.OsType)
            if replicaCount := applicationProperties.ReplicaCount; replicaCount != nil {
                v["replica_count"] = *replicaCount
            }
        }

        results = append(results, v)
    }

    return results
}

func flattenArmApplicationDiagnosticsSinkProperties(input *[]servicefabricmeshrestapis.DiagnosticsSinkProperties) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        if name := item.Name; name != nil {
            v["name"] = *name
        }
        if description := item.Description; description != nil {
            v["description"] = *description
        }

        results = append(results, v)
    }

    return results
}

func flattenArmApplicationContainerCodePackageProperties(input *[]servicefabricmeshrestapis.ContainerCodePackageProperties) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        if name := item.Name; name != nil {
            v["name"] = *name
        }
        v["commands"] = utils.FlattenStringSlice(item.Commands)
        if entrypoint := item.Entrypoint; entrypoint != nil {
            v["entrypoint"] = *entrypoint
        }
        if image := item.Image; image != nil {
            v["image"] = *image
        }

        results = append(results, v)
    }

    return results
}

func flattenArmApplicationDiagnosticsRef(input *servicefabricmeshrestapis.DiagnosticsRef) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if enabled := input.Enabled; enabled != nil {
        result["enabled"] = *enabled
    }
    result["sink_refs"] = utils.FlattenStringSlice(input.SinkRefs)

    return []interface{}{result}
}

func flattenArmApplicationNetworkRef(input *[]servicefabricmeshrestapis.NetworkRef) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        if name := item.Name; name != nil {
            v["name"] = *name
        }

        results = append(results, v)
    }

    return results
}