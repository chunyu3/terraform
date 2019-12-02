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



func resourceArmExperiment() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmExperimentCreate,
        Read: resourceArmExperimentRead,
        Update: resourceArmExperimentUpdate,
        Delete: resourceArmExperimentDelete,

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

            "profile_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "description": {
                Type: schema.TypeString,
                Optional: true,
            },

            "enabled_state": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(frontdoor.Enabled),
                    string(frontdoor.Disabled),
                }, false),
                Default: string(frontdoor.Enabled),
            },

            "endpoint_a_endpoint": {
                Type: schema.TypeString,
                Computed: true,
            },

            "endpoint_a_name": {
                Type: schema.TypeString,
                Computed: true,
            },

            "endpoint_b_endpoint": {
                Type: schema.TypeString,
                Computed: true,
            },

            "endpoint_b_name": {
                Type: schema.TypeString,
                Computed: true,
            },

            "resource_state": {
                Type: schema.TypeString,
                Computed: true,
            },

            "script_file_uri": {
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

            "tags": tags.Schema(),
        },
    }
}

func resourceArmExperimentCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).experimentsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    profileName := d.Get("profile_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, profileName, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Experiment %q (Profile Name %q / Resource Group %q): %+v", name, profileName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_experiment", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    description := d.Get("description").(string)
    enabledState := d.Get("enabled_state").(string)
    t := d.Get("tags").(map[string]interface{})

    parameters := frontdoor.ExperimentUpdateModel{
        Location: utils.String(location),
        ExperimentUpdateProperties: &frontdoor.ExperimentUpdateProperties{
            Description: utils.String(description),
            EnabledState: frontdoor.State(enabledState),
        },
        Tags: tags.Expand(t),
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, profileName, name, parameters)
    if err != nil {
        return fmt.Errorf("Error creating Experiment %q (Profile Name %q / Resource Group %q): %+v", name, profileName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Experiment %q (Profile Name %q / Resource Group %q): %+v", name, profileName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, profileName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Experiment %q (Profile Name %q / Resource Group %q): %+v", name, profileName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Experiment %q (Profile Name %q / Resource Group %q) ID", name, profileName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmExperimentRead(d, meta)
}

func resourceArmExperimentRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).experimentsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    profileName := id.Path["NetworkExperimentProfiles"]
    name := id.Path["Experiments"]

    resp, err := client.Get(ctx, resourceGroup, profileName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Experiment %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Experiment %q (Profile Name %q / Resource Group %q): %+v", name, profileName, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if experimentUpdateProperties := resp.ExperimentUpdateProperties; experimentUpdateProperties != nil {
        d.Set("description", experimentUpdateProperties.Description)
        d.Set("enabled_state", string(experimentUpdateProperties.EnabledState))
        if endpointA := experimentUpdateProperties.EndpointA; endpointA != nil {
            d.Set("endpoint_a_endpoint", endpointA.endpointA_Endpoint)
            d.Set("endpoint_a_name", endpointA.endpointA_Name)
        }
        if endpointB := experimentUpdateProperties.EndpointB; endpointB != nil {
            d.Set("endpoint_b_endpoint", endpointB.endpointB_Endpoint)
            d.Set("endpoint_b_name", endpointB.endpointB_Name)
        }
        d.Set("resource_state", string(experimentUpdateProperties.ResourceState))
        d.Set("script_file_uri", experimentUpdateProperties.ScriptFileURI)
        d.Set("status", experimentUpdateProperties.Status)
    }
    d.Set("profile_name", profileName)
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}

func resourceArmExperimentUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).experimentsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    description := d.Get("description").(string)
    enabledState := d.Get("enabled_state").(string)
    profileName := d.Get("profile_name").(string)
    t := d.Get("tags").(map[string]interface{})

    parameters := frontdoor.ExperimentUpdateModel{
        ExperimentUpdateProperties: &frontdoor.ExperimentUpdateProperties{
            Description: utils.String(description),
            EnabledState: frontdoor.State(enabledState),
        },
        Tags: tags.Expand(t),
    }


    future, err := client.Update(ctx, resourceGroup, profileName, name, parameters)
    if err != nil {
        return fmt.Errorf("Error updating Experiment %q (Profile Name %q / Resource Group %q): %+v", name, profileName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for update of Experiment %q (Profile Name %q / Resource Group %q): %+v", name, profileName, resourceGroup, err)
    }

    return resourceArmExperimentRead(d, meta)
}

func resourceArmExperimentDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).experimentsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    profileName := id.Path["NetworkExperimentProfiles"]
    name := id.Path["Experiments"]

    future, err := client.Delete(ctx, resourceGroup, profileName, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Experiment %q (Profile Name %q / Resource Group %q): %+v", name, profileName, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Experiment %q (Profile Name %q / Resource Group %q): %+v", name, profileName, resourceGroup, err)
        }
    }

    return nil
}
