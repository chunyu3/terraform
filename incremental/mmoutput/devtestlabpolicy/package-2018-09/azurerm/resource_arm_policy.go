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



func resourceArmPolicy() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmPolicyCreate,
        Read: resourceArmPolicyRead,
        Update: resourceArmPolicyUpdate,
        Delete: resourceArmPolicyDelete,

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

            "name": {
                Type: schema.TypeString,
                Computed: true,
            },

            "location": azure.SchemaLocation(),

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "lab_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "description": {
                Type: schema.TypeString,
                Optional: true,
            },

            "evaluator_type": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(devtestlab.AllowedValuesPolicy),
                    string(devtestlab.MaxValuePolicy),
                }, false),
                Default: string(devtestlab.AllowedValuesPolicy),
            },

            "fact_data": {
                Type: schema.TypeString,
                Optional: true,
            },

            "fact_name": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(devtestlab.UserOwnedLabVmCount),
                    string(devtestlab.UserOwnedLabPremiumVmCount),
                    string(devtestlab.LabVmCount),
                    string(devtestlab.LabPremiumVmCount),
                    string(devtestlab.LabVmSize),
                    string(devtestlab.GalleryImage),
                    string(devtestlab.UserOwnedLabVmCountInSubnet),
                    string(devtestlab.LabTargetCost),
                    string(devtestlab.EnvironmentTemplate),
                    string(devtestlab.ScheduleEditPermission),
                }, false),
                Default: string(devtestlab.UserOwnedLabVmCount),
            },

            "status": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(devtestlab.Enabled),
                    string(devtestlab.Disabled),
                }, false),
                Default: string(devtestlab.Enabled),
            },

            "threshold": {
                Type: schema.TypeString,
                Optional: true,
            },

            "created_date": {
                Type: schema.TypeString,
                Computed: true,
            },

            "provisioning_state": {
                Type: schema.TypeString,
                Computed: true,
            },

            "type": {
                Type: schema.TypeString,
                Computed: true,
            },

            "unique_identifier": {
                Type: schema.TypeString,
                Computed: true,
            },

            "tags": tags.Schema(),
        },
    }
}

func resourceArmPolicyCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).policiesClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    labName := d.Get("lab_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, labName, name, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Policy %q (Lab Name %q / Resource Group %q): %+v", name, labName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_policy", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    description := d.Get("description").(string)
    evaluatorType := d.Get("evaluator_type").(string)
    factData := d.Get("fact_data").(string)
    factName := d.Get("fact_name").(string)
    status := d.Get("status").(string)
    threshold := d.Get("threshold").(string)
    t := d.Get("tags").(map[string]interface{})

    policy := devtestlab.PolicyFragment{
        Location: utils.String(location),
        PolicyPropertiesFragment: &devtestlab.PolicyPropertiesFragment{
            Description: utils.String(description),
            EvaluatorType: devtestlab.PolicyEvaluatorType(evaluatorType),
            FactData: utils.String(factData),
            FactName: devtestlab.PolicyFactName(factName),
            Status: devtestlab.PolicyStatus(status),
            Threshold: utils.String(threshold),
        },
        Tags: tags.Expand(t),
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroup, labName, name, name, policy); err != nil {
        return fmt.Errorf("Error creating Policy %q (Lab Name %q / Resource Group %q): %+v", name, labName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, labName, name, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Policy %q (Lab Name %q / Resource Group %q): %+v", name, labName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Policy %q (Lab Name %q / Resource Group %q) ID", name, labName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmPolicyRead(d, meta)
}

func resourceArmPolicyRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).policiesClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    labName := id.Path["labs"]
    name := id.Path["policysets"]
    name := id.Path["policies"]

    resp, err := client.Get(ctx, resourceGroup, labName, name, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Policy %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Policy %q (Lab Name %q / Resource Group %q): %+v", name, labName, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if policyPropertiesFragment := resp.PolicyPropertiesFragment; policyPropertiesFragment != nil {
        d.Set("created_date", (policyPropertiesFragment.CreatedDate).String())
        d.Set("description", policyPropertiesFragment.Description)
        d.Set("evaluator_type", string(policyPropertiesFragment.EvaluatorType))
        d.Set("fact_data", policyPropertiesFragment.FactData)
        d.Set("fact_name", string(policyPropertiesFragment.FactName))
        d.Set("provisioning_state", policyPropertiesFragment.ProvisioningState)
        d.Set("status", string(policyPropertiesFragment.Status))
        d.Set("threshold", policyPropertiesFragment.Threshold)
        d.Set("unique_identifier", policyPropertiesFragment.UniqueIdentifier)
    }
    d.Set("lab_name", labName)
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}

func resourceArmPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).policiesClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    description := d.Get("description").(string)
    evaluatorType := d.Get("evaluator_type").(string)
    factData := d.Get("fact_data").(string)
    factName := d.Get("fact_name").(string)
    labName := d.Get("lab_name").(string)
    status := d.Get("status").(string)
    threshold := d.Get("threshold").(string)
    t := d.Get("tags").(map[string]interface{})

    policy := devtestlab.PolicyFragment{
        PolicyPropertiesFragment: &devtestlab.PolicyPropertiesFragment{
            Description: utils.String(description),
            EvaluatorType: devtestlab.PolicyEvaluatorType(evaluatorType),
            FactData: utils.String(factData),
            FactName: devtestlab.PolicyFactName(factName),
            Status: devtestlab.PolicyStatus(status),
            Threshold: utils.String(threshold),
        },
        Tags: tags.Expand(t),
    }


    if _, err := client.Update(ctx, resourceGroup, labName, name, name, policy); err != nil {
        return fmt.Errorf("Error updating Policy %q (Lab Name %q / Resource Group %q): %+v", name, labName, resourceGroup, err)
    }

    return resourceArmPolicyRead(d, meta)
}

func resourceArmPolicyDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).policiesClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    labName := id.Path["labs"]
    name := id.Path["policysets"]
    name := id.Path["policies"]

    if _, err := client.Delete(ctx, resourceGroup, labName, name, name); err != nil {
        return fmt.Errorf("Error deleting Policy %q (Lab Name %q / Resource Group %q): %+v", name, labName, resourceGroup, err)
    }

    return nil
}
