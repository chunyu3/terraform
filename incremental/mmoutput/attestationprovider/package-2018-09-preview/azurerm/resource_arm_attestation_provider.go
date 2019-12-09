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



func resourceArmAttestationProvider() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmAttestationProviderCreateUpdate,
        Read: resourceArmAttestationProviderRead,
        Update: resourceArmAttestationProviderCreateUpdate,
        Delete: resourceArmAttestationProviderDelete,

        Importer: &schema.ResourceImporter{
            State: schema.ImportStatePassthrough,
        },


        Schema: map[string]*schema.Schema{
            "provider_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "attestation_policy": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
            },

            "policy_signing_certificates": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "keys": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "alg": {
                                        Type: schema.TypeString,
                                        Required: true,
                                        ValidateFunc: validate.NoEmptyStrings,
                                    },
                                    "kid": {
                                        Type: schema.TypeString,
                                        Required: true,
                                        ValidateFunc: validate.NoEmptyStrings,
                                    },
                                    "kty": {
                                        Type: schema.TypeString,
                                        Required: true,
                                        ValidateFunc: validate.NoEmptyStrings,
                                    },
                                    "use": {
                                        Type: schema.TypeString,
                                        Required: true,
                                        ValidateFunc: validate.NoEmptyStrings,
                                    },
                                    "crv": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                    "d": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                    "dp": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                    "dq": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                    "e": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                    "k": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                    "n": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                    "p": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                    "q": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                    "qi": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                    "x": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                    "x5c": {
                                        Type: schema.TypeList,
                                        Optional: true,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                        },
                                    },
                                    "y": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                },
                            },
                        },
                    },
                },
            },

            "attest_uri": {
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

            "status": {
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

func resourceArmAttestationProviderCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).attestationProvidersClient
    ctx, cancel := timeouts.ForCreateUpdate(meta.(*ArmClient).StopContext, d)
    defer cancel()

    resourceGroupName := d.Get("resource_group").(string)
    name := d.Get("provider_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroupName, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Attestation Provider (Provider Name %q / Resource Group %q): %+v", name, resourceGroupName, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_attestation_provider", *existing.ID)
        }
    }

    attestationPolicy := d.Get("attestation_policy").(string)
    policySigningCertificates := d.Get("policy_signing_certificates").([]interface{})

    creationParams := attestation.ServiceCreationParams{
        AttestationPolicy: utils.String(attestationPolicy),
        PolicySigningCertificates: expandArmAttestationProviderJSONWebKeySet(policySigningCertificates),
    }


    if _, err := client.Create(ctx, resourceGroupName, name, creationParams); err != nil {
        return fmt.Errorf("Error creating Attestation Provider (Provider Name %q / Resource Group %q): %+v", name, resourceGroupName, err)
    }


    resp, err := client.Get(ctx, resourceGroupName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Attestation Provider (Provider Name %q / Resource Group %q): %+v", name, resourceGroupName, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Attestation Provider (Provider Name %q / Resource Group %q) ID", name, resourceGroupName)
    }
    d.SetId(*resp.ID)

    return resourceArmAttestationProviderRead(d, meta)
}

func resourceArmAttestationProviderRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).attestationProvidersClient
    ctx, cancel := timeouts.ForRead(meta.(*ArmClient).StopContext, d)
    defer cancel()

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroupName := id.ResourceGroup
    name := id.Path["attestationProviders"]

    resp, err := client.Get(ctx, resourceGroupName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Attestation Provider %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Attestation Provider (Provider Name %q / Resource Group %q): %+v", name, resourceGroupName, err)
    }


    d.Set("resource_group", resourceGroupName)
    if statusResult := resp.StatusResult; statusResult != nil {
        d.Set("attest_uri", statusResult.AttestURI)
        d.Set("status", string(statusResult.Status))
    }
    d.Set("id", resp.ID)
    d.Set("name", resp.Name)
    d.Set("provider_name", name)
    d.Set("type", resp.Type)

    return nil
}


func resourceArmAttestationProviderDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).attestationProvidersClient
    ctx, cancel := timeouts.ForDelete(meta.(*ArmClient).StopContext, d)
    defer cancel()


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroupName := id.ResourceGroup
    name := id.Path["attestationProviders"]

    if _, err := client.Delete(ctx, resourceGroupName, name); err != nil {
        return fmt.Errorf("Error deleting Attestation Provider (Provider Name %q / Resource Group %q): %+v", name, resourceGroupName, err)
    }

    return nil
}

func expandArmAttestationProviderJSONWebKeySet(input []interface{}) *attestation.JSONWebKeySet {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    keys := v["keys"].([]interface{})

    result := attestation.JSONWebKeySet{
        Keys: expandArmAttestationProviderJSONWebKey(keys),
    }
    return &result
}

func expandArmAttestationProviderJSONWebKey(input []interface{}) *[]attestation.JSONWebKey {
    results := make([]attestation.JSONWebKey, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        alg := v["alg"].(string)
        crv := v["crv"].(string)
        d := v["d"].(string)
        dp := v["dp"].(string)
        dq := v["dq"].(string)
        e := v["e"].(string)
        k := v["k"].(string)
        kid := v["kid"].(string)
        kty := v["kty"].(string)
        n := v["n"].(string)
        p := v["p"].(string)
        q := v["q"].(string)
        qi := v["qi"].(string)
        use := v["use"].(string)
        x := v["x"].(string)
        x5c := v["x5c"].([]interface{})
        y := v["y"].(string)

        result := attestation.JSONWebKey{
            Alg: utils.String(alg),
            Crv: utils.String(crv),
            D: utils.String(d),
            Dp: utils.String(dp),
            Dq: utils.String(dq),
            E: utils.String(e),
            K: utils.String(k),
            Kid: utils.String(kid),
            Kty: utils.String(kty),
            N: utils.String(n),
            P: utils.String(p),
            Q: utils.String(q),
            Qi: utils.String(qi),
            Use: utils.String(use),
            X: utils.String(x),
            X5c: utils.ExpandStringSlice(x5c),
            Y: utils.String(y),
        }

        results = append(results, result)
    }
    return &results
}
