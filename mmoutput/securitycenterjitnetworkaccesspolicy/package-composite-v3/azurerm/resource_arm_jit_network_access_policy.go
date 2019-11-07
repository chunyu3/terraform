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



func resourceArmJitNetworkAccessPolicy() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmJitNetworkAccessPolicyCreateUpdate,
        Read: resourceArmJitNetworkAccessPolicyRead,
        Update: resourceArmJitNetworkAccessPolicyCreateUpdate,
        Delete: resourceArmJitNetworkAccessPolicyDelete,

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

            "asc_location": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "virtual_machines": {
                Type: schema.TypeList,
                Required: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "id": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "ports": {
                            Type: schema.TypeList,
                            Required: true,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "max_request_access_duration": {
                                        Type: schema.TypeString,
                                        Required: true,
                                        ValidateFunc: validate.NoEmptyStrings,
                                    },
                                    "number": {
                                        Type: schema.TypeInt,
                                        Required: true,
                                    },
                                    "protocol": {
                                        Type: schema.TypeString,
                                        Required: true,
                                        ValidateFunc: validation.StringInSlice([]string{
                                            string(securitycenter.TCP),
                                            string(securitycenter.UDP),
                                            string(securitycenter.All),
                                        }, false),
                                    },
                                    "allowed_source_address_prefix": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                    "allowed_source_address_prefixes": {
                                        Type: schema.TypeList,
                                        Optional: true,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                        },
                                    },
                                },
                            },
                        },
                        "public_ip_address": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },

            "kind": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
            },

            "requests": {
                Type: schema.TypeList,
                Optional: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "requestor": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "start_time_utc": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                            ValidateFunc: validateRFC3339Date,
                        },
                        "virtual_machines": {
                            Type: schema.TypeList,
                            Required: true,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "id": {
                                        Type: schema.TypeString,
                                        Required: true,
                                        ValidateFunc: validate.NoEmptyStrings,
                                    },
                                    "ports": {
                                        Type: schema.TypeList,
                                        Required: true,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "end_time_utc": {
                                                    Type: schema.TypeString,
                                                    Required: true,
                                                    ValidateFunc: validate.NoEmptyStrings,
                                                    ValidateFunc: validateRFC3339Date,
                                                },
                                                "number": {
                                                    Type: schema.TypeInt,
                                                    Required: true,
                                                },
                                                "status": {
                                                    Type: schema.TypeString,
                                                    Required: true,
                                                    ValidateFunc: validation.StringInSlice([]string{
                                                        string(securitycenter.Revoked),
                                                        string(securitycenter.Initiated),
                                                    }, false),
                                                },
                                                "status_reason": {
                                                    Type: schema.TypeString,
                                                    Required: true,
                                                    ValidateFunc: validation.StringInSlice([]string{
                                                        string(securitycenter.Expired),
                                                        string(securitycenter.UserRequested),
                                                        string(securitycenter.NewerRequestInitiated),
                                                    }, false),
                                                },
                                                "allowed_source_address_prefix": {
                                                    Type: schema.TypeString,
                                                    Optional: true,
                                                },
                                                "allowed_source_address_prefixes": {
                                                    Type: schema.TypeList,
                                                    Optional: true,
                                                    Elem: &schema.Schema{
                                                        Type: schema.TypeString,
                                                    },
                                                },
                                                "mapped_port": {
                                                    Type: schema.TypeInt,
                                                    Optional: true,
                                                },
                                            },
                                        },
                                    },
                                },
                            },
                        },
                        "justification": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },

            "provisioning_state": {
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

func resourceArmJitNetworkAccessPolicyCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).jitNetworkAccessPoliciesClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    ascLocation := d.Get("asc_location").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, ascLocation, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Jit Network Access Policy %q (Asc Location %q / Resource Group %q): %+v", name, ascLocation, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_jit_network_access_policy", *existing.ID)
        }
    }

    kind := d.Get("kind").(string)
    requests := d.Get("requests").([]interface{})
    virtualMachines := d.Get("virtual_machines").([]interface{})

    body := securitycenter.JitNetworkAccessPolicy{
        Kind: utils.String(kind),
        JitNetworkAccessPolicyProperties: &securitycenter.JitNetworkAccessPolicyProperties{
            Requests: expandArmJitNetworkAccessPolicyJitNetworkAccessRequest(requests),
            VirtualMachines: expandArmJitNetworkAccessPolicyJitNetworkAccessPolicyVirtualMachine(virtualMachines),
        },
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroup, ascLocation, name, body); err != nil {
        return fmt.Errorf("Error creating Jit Network Access Policy %q (Asc Location %q / Resource Group %q): %+v", name, ascLocation, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, ascLocation, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Jit Network Access Policy %q (Asc Location %q / Resource Group %q): %+v", name, ascLocation, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Jit Network Access Policy %q (Asc Location %q / Resource Group %q) ID", name, ascLocation, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmJitNetworkAccessPolicyRead(d, meta)
}

func resourceArmJitNetworkAccessPolicyRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).jitNetworkAccessPoliciesClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    ascLocation := id.Path["locations"]
    name := id.Path["jitNetworkAccessPolicies"]

    resp, err := client.Get(ctx, resourceGroup, ascLocation, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Jit Network Access Policy %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Jit Network Access Policy %q (Asc Location %q / Resource Group %q): %+v", name, ascLocation, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    d.Set("asc_location", ascLocation)
    d.Set("kind", resp.Kind)
    if jitNetworkAccessPolicyProperties := resp.JitNetworkAccessPolicyProperties; jitNetworkAccessPolicyProperties != nil {
        d.Set("provisioning_state", jitNetworkAccessPolicyProperties.ProvisioningState)
        if err := d.Set("requests", flattenArmJitNetworkAccessPolicyJitNetworkAccessRequest(jitNetworkAccessPolicyProperties.Requests)); err != nil {
            return fmt.Errorf("Error setting `requests`: %+v", err)
        }
        if err := d.Set("virtual_machines", flattenArmJitNetworkAccessPolicyJitNetworkAccessPolicyVirtualMachine(jitNetworkAccessPolicyProperties.VirtualMachines)); err != nil {
            return fmt.Errorf("Error setting `virtual_machines`: %+v", err)
        }
    }
    d.Set("type", resp.Type)

    return nil
}


func resourceArmJitNetworkAccessPolicyDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).jitNetworkAccessPoliciesClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    ascLocation := id.Path["locations"]
    name := id.Path["jitNetworkAccessPolicies"]

    if _, err := client.Delete(ctx, resourceGroup, ascLocation, name); err != nil {
        return fmt.Errorf("Error deleting Jit Network Access Policy %q (Asc Location %q / Resource Group %q): %+v", name, ascLocation, resourceGroup, err)
    }

    return nil
}

func expandArmJitNetworkAccessPolicyJitNetworkAccessRequest(input []interface{}) *[]securitycenter.JitNetworkAccessRequest {
    results := make([]securitycenter.JitNetworkAccessRequest, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        virtualMachines := v["virtual_machines"].([]interface{})
        startTimeUtc := v["start_time_utc"].(string)
        requestor := v["requestor"].(string)
        justification := v["justification"].(string)

        result := securitycenter.JitNetworkAccessRequest{
            Justification: utils.String(justification),
            Requestor: utils.String(requestor),
            StartTimeUtc: convertStringToDate(startTimeUtc),
            VirtualMachines: expandArmJitNetworkAccessPolicyJitNetworkAccessRequestVirtualMachine(virtualMachines),
        }

        results = append(results, result)
    }
    return &results
}

func expandArmJitNetworkAccessPolicyJitNetworkAccessPolicyVirtualMachine(input []interface{}) *[]securitycenter.JitNetworkAccessPolicyVirtualMachine {
    results := make([]securitycenter.JitNetworkAccessPolicyVirtualMachine, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        id := v["id"].(string)
        ports := v["ports"].([]interface{})
        publicIpAddress := v["public_ip_address"].(string)

        result := securitycenter.JitNetworkAccessPolicyVirtualMachine{
            ID: utils.String(id),
            Ports: expandArmJitNetworkAccessPolicyJitNetworkAccessPortRule(ports),
            PublicIpAddress: utils.String(publicIpAddress),
        }

        results = append(results, result)
    }
    return &results
}

func convertStringToDate(input interface{}) *date.Time {
  v := input.(string)

  dateTime, err := date.ParseTime(time.RFC3339, v)
  if err != nil {
      log.Printf("[ERROR] Cannot convert an invalid string to RFC3339 date %q: %+v", v, err)
      return nil
  }

  result := date.Time{
      Time: dateTime,
  }
  return &result
}

func expandArmJitNetworkAccessPolicyJitNetworkAccessRequestVirtualMachine(input []interface{}) *[]securitycenter.JitNetworkAccessRequestVirtualMachine {
    results := make([]securitycenter.JitNetworkAccessRequestVirtualMachine, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        id := v["id"].(string)
        ports := v["ports"].([]interface{})

        result := securitycenter.JitNetworkAccessRequestVirtualMachine{
            ID: utils.String(id),
            Ports: expandArmJitNetworkAccessPolicyJitNetworkAccessRequestPort(ports),
        }

        results = append(results, result)
    }
    return &results
}

func expandArmJitNetworkAccessPolicyJitNetworkAccessPortRule(input []interface{}) *[]securitycenter.JitNetworkAccessPortRule {
    results := make([]securitycenter.JitNetworkAccessPortRule, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        number := v["number"].(int)
        protocol := v["protocol"].(string)
        allowedSourceAddressPrefix := v["allowed_source_address_prefix"].(string)
        allowedSourceAddressPrefixes := v["allowed_source_address_prefixes"].([]interface{})
        maxRequestAccessDuration := v["max_request_access_duration"].(string)

        result := securitycenter.JitNetworkAccessPortRule{
            AllowedSourceAddressPrefix: utils.String(allowedSourceAddressPrefix),
            AllowedSourceAddressPrefixes: utils.ExpandStringSlice(allowedSourceAddressPrefixes),
            MaxRequestAccessDuration: utils.String(maxRequestAccessDuration),
            Number: utils.Int(number),
            Protocol: securitycenter.Protocol(protocol),
        }

        results = append(results, result)
    }
    return &results
}

func expandArmJitNetworkAccessPolicyJitNetworkAccessRequestPort(input []interface{}) *[]securitycenter.JitNetworkAccessRequestPort {
    results := make([]securitycenter.JitNetworkAccessRequestPort, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        number := v["number"].(int)
        allowedSourceAddressPrefix := v["allowed_source_address_prefix"].(string)
        allowedSourceAddressPrefixes := v["allowed_source_address_prefixes"].([]interface{})
        endTimeUtc := v["end_time_utc"].(string)
        status := v["status"].(string)
        statusReason := v["status_reason"].(string)
        mappedPort := v["mapped_port"].(int)

        result := securitycenter.JitNetworkAccessRequestPort{
            AllowedSourceAddressPrefix: utils.String(allowedSourceAddressPrefix),
            AllowedSourceAddressPrefixes: utils.ExpandStringSlice(allowedSourceAddressPrefixes),
            EndTimeUtc: convertStringToDate(endTimeUtc),
            MappedPort: utils.Int(mappedPort),
            Number: utils.Int(number),
            Status: securitycenter.Status(status),
            StatusReason: securitycenter.StatusReason(statusReason),
        }

        results = append(results, result)
    }
    return &results
}


func flattenArmJitNetworkAccessPolicyJitNetworkAccessRequest(input *[]securitycenter.JitNetworkAccessRequest) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        if justification := item.Justification; justification != nil {
            v["justification"] = *justification
        }
        if requestor := item.Requestor; requestor != nil {
            v["requestor"] = *requestor
        }
        if startTimeUtc := item.StartTimeUtc; startTimeUtc != nil {
            v["start_time_utc"] = (*startTimeUtc).String()
        }
        v["virtual_machines"] = flattenArmJitNetworkAccessPolicyJitNetworkAccessRequestVirtualMachine(item.VirtualMachines)

        results = append(results, v)
    }

    return results
}

func flattenArmJitNetworkAccessPolicyJitNetworkAccessPolicyVirtualMachine(input *[]securitycenter.JitNetworkAccessPolicyVirtualMachine) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        if id := item.ID; id != nil {
            v["id"] = *id
        }
        v["ports"] = flattenArmJitNetworkAccessPolicyJitNetworkAccessPortRule(item.Ports)
        if publicIpAddress := item.PublicIpAddress; publicIpAddress != nil {
            v["public_ip_address"] = *publicIpAddress
        }

        results = append(results, v)
    }

    return results
}

func flattenArmJitNetworkAccessPolicyJitNetworkAccessRequestVirtualMachine(input *[]securitycenter.JitNetworkAccessRequestVirtualMachine) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        if id := item.ID; id != nil {
            v["id"] = *id
        }
        v["ports"] = flattenArmJitNetworkAccessPolicyJitNetworkAccessRequestPort(item.Ports)

        results = append(results, v)
    }

    return results
}

func flattenArmJitNetworkAccessPolicyJitNetworkAccessPortRule(input *[]securitycenter.JitNetworkAccessPortRule) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        if allowedSourceAddressPrefix := item.AllowedSourceAddressPrefix; allowedSourceAddressPrefix != nil {
            v["allowed_source_address_prefix"] = *allowedSourceAddressPrefix
        }
        v["allowed_source_address_prefixes"] = utils.FlattenStringSlice(item.AllowedSourceAddressPrefixes)
        if maxRequestAccessDuration := item.MaxRequestAccessDuration; maxRequestAccessDuration != nil {
            v["max_request_access_duration"] = *maxRequestAccessDuration
        }
        if number := item.Number; number != nil {
            v["number"] = *number
        }
        v["protocol"] = string(item.Protocol)

        results = append(results, v)
    }

    return results
}

func flattenArmJitNetworkAccessPolicyJitNetworkAccessRequestPort(input *[]securitycenter.JitNetworkAccessRequestPort) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        if allowedSourceAddressPrefix := item.AllowedSourceAddressPrefix; allowedSourceAddressPrefix != nil {
            v["allowed_source_address_prefix"] = *allowedSourceAddressPrefix
        }
        v["allowed_source_address_prefixes"] = utils.FlattenStringSlice(item.AllowedSourceAddressPrefixes)
        if endTimeUtc := item.EndTimeUtc; endTimeUtc != nil {
            v["end_time_utc"] = (*endTimeUtc).String()
        }
        if mappedPort := item.MappedPort; mappedPort != nil {
            v["mapped_port"] = *mappedPort
        }
        if number := item.Number; number != nil {
            v["number"] = *number
        }
        v["status"] = string(item.Status)
        v["status_reason"] = string(item.StatusReason)

        results = append(results, v)
    }

    return results
}
