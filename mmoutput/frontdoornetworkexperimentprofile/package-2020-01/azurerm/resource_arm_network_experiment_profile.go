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



func resourceArmNetworkExperimentProfile() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmNetworkExperimentProfileCreate,
        Read: resourceArmNetworkExperimentProfileRead,
        Update: resourceArmNetworkExperimentProfileUpdate,
        Delete: resourceArmNetworkExperimentProfileDelete,

        Importer: &schema.ResourceImporter{
            State: schema.ImportStatePassthrough,
        },


        Schema: map[string]*schema.Schema{
            "profile_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "enabled_state": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(frontdoor.Enabled),
                    string(frontdoor.Disabled),
                }, false),
                Default: string(frontdoor.Enabled),
            },

            "etag": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
            },

            "location": azure.SchemaLocation(),

            "tags": tags.Schema(),

            "id": {
                Type: schema.TypeString,
                Computed: true,
            },

            "name": {
                Type: schema.TypeString,
                Computed: true,
            },

            "resource_state": {
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

func resourceArmNetworkExperimentProfileCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).networkExperimentProfilesClient
    ctx, cancel := timeouts.ForCreate(meta.(*ArmClient).StopContext, d)
    defer cancel()

    resourceGroupName := d.Get("resource_group").(string)
    name := d.Get("profile_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroupName, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Network Experiment Profile (Profile Name %q / Resource Group %q): %+v", name, resourceGroupName, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_network_experiment_profile", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    enabledState := d.Get("enabled_state").(string)
    etag := d.Get("etag").(string)
    tags := d.Get("tags").(map[string]interface{})

    parameters := frontdoor.ProfileUpdateModel{
        Etag: utils.String(etag),
        Location: utils.String(location),
        ProfileUpdateProperties: &frontdoor.ProfileUpdateProperties{
            EnabledState: frontdoor.State(enabledState),
        },
        Tags: tags.Expand(tags),
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroupName, name, parameters)
    if err != nil {
        return fmt.Errorf("Error creating Network Experiment Profile (Profile Name %q / Resource Group %q): %+v", name, resourceGroupName, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Network Experiment Profile (Profile Name %q / Resource Group %q): %+v", name, resourceGroupName, err)
    }


    resp, err := client.Get(ctx, resourceGroupName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Network Experiment Profile (Profile Name %q / Resource Group %q): %+v", name, resourceGroupName, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Network Experiment Profile (Profile Name %q / Resource Group %q) ID", name, resourceGroupName)
    }
    d.SetId(*resp.ID)

    return resourceArmNetworkExperimentProfileRead(d, meta)
}

func resourceArmNetworkExperimentProfileRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).networkExperimentProfilesClient
    ctx, cancel := timeouts.ForRead(meta.(*ArmClient).StopContext, d)
    defer cancel()

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroupName := id.ResourceGroup
    name := id.Path["NetworkExperimentProfiles"]

    resp, err := client.Get(ctx, resourceGroupName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Network Experiment Profile %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Network Experiment Profile (Profile Name %q / Resource Group %q): %+v", name, resourceGroupName, err)
    }


    d.Set("resource_group", resourceGroupName)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if profileUpdateProperties := resp.ProfileUpdateProperties; profileUpdateProperties != nil {
        d.Set("enabled_state", string(profileUpdateProperties.EnabledState))
        d.Set("resource_state", string(profileUpdateProperties.ResourceState))
    }
    d.Set("etag", resp.Etag)
    d.Set("id", resp.ID)
    d.Set("name", resp.Name)
    d.Set("profile_name", name)
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}

func resourceArmNetworkExperimentProfileUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).networkExperimentProfilesClient
    ctx, cancel := timeouts.ForUpdate(meta.(*ArmClient).StopContext, d)
    defer cancel()

      resourceGroupName := d.Get("resource_group").(string)
    location := azure.NormalizeLocation(d.Get("location").(string))
    enabledState := d.Get("enabled_state").(string)
    etag := d.Get("etag").(string)
    name := d.Get("profile_name").(string)
    tags := d.Get("tags").(map[string]interface{})

    parameters := frontdoor.ProfileUpdateModel{
        Etag: utils.String(etag),
        Location: utils.String(location),
        ProfileUpdateProperties: &frontdoor.ProfileUpdateProperties{
            EnabledState: frontdoor.State(enabledState),
        },
        Tags: tags.Expand(tags),
    }


    future, err := client.Update(ctx, resourceGroupName, name, parameters)
    if err != nil {
        return fmt.Errorf("Error updating Network Experiment Profile (Profile Name %q / Resource Group %q): %+v", name, resourceGroupName, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for update of Network Experiment Profile (Profile Name %q / Resource Group %q): %+v", name, resourceGroupName, err)
    }

    return resourceArmNetworkExperimentProfileRead(d, meta)
}

func resourceArmNetworkExperimentProfileDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).networkExperimentProfilesClient
    ctx, cancel := timeouts.ForDelete(meta.(*ArmClient).StopContext, d)
    defer cancel()


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroupName := id.ResourceGroup
    name := id.Path["NetworkExperimentProfiles"]

    future, err := client.Delete(ctx, resourceGroupName, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Network Experiment Profile (Profile Name %q / Resource Group %q): %+v", name, resourceGroupName, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Network Experiment Profile (Profile Name %q / Resource Group %q): %+v", name, resourceGroupName, err)
        }
    }

    return nil
}
