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

            "resource_state": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(frontdoor.Creating),
                    string(frontdoor.Enabling),
                    string(frontdoor.Enabled),
                    string(frontdoor.Disabling),
                    string(frontdoor.Disabled),
                    string(frontdoor.Deleting),
                }, false),
                Default: string(frontdoor.Creating),
            },

            "type": {
                Type: schema.TypeString,
                Computed: true,
            },

            "tags": tags.Schema(),
        },
    }
}

func resourceArmNetworkExperimentProfileCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).networkExperimentProfilesClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Network Experiment Profile %q (Resource Group %q): %+v", name, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_network_experiment_profile", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    enabledState := d.Get("enabled_state").(string)
    etag := d.Get("etag").(string)
    resourceState := d.Get("resource_state").(string)
    t := d.Get("tags").(map[string]interface{})

    parameters := frontdoor.Profile{
        Etag: utils.String(etag),
        Location: utils.String(location),
        ProfileProperties: &frontdoor.ProfileProperties{
            EnabledState: frontdoor.State(enabledState),
            ResourceState: frontdoor.NetworkExperimentResourceState(resourceState),
        },
        Tags: tags.Expand(t),
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, name, parameters)
    if err != nil {
        return fmt.Errorf("Error creating Network Experiment Profile %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Network Experiment Profile %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Network Experiment Profile %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Network Experiment Profile %q (Resource Group %q) ID", name, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmNetworkExperimentProfileRead(d, meta)
}

func resourceArmNetworkExperimentProfileRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).networkExperimentProfilesClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["NetworkExperimentProfiles"]

    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Network Experiment Profile %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Network Experiment Profile %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if profileProperties := resp.ProfileProperties; profileProperties != nil {
        d.Set("enabled_state", string(profileProperties.EnabledState))
        d.Set("resource_state", string(profileProperties.ResourceState))
    }
    d.Set("etag", resp.Etag)
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}

func resourceArmNetworkExperimentProfileUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).networkExperimentProfilesClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    enabledState := d.Get("enabled_state").(string)
    etag := d.Get("etag").(string)
    resourceState := d.Get("resource_state").(string)
    t := d.Get("tags").(map[string]interface{})

    parameters := frontdoor.Profile{
        Etag: utils.String(etag),
        Location: utils.String(location),
        ProfileProperties: &frontdoor.ProfileProperties{
            EnabledState: frontdoor.State(enabledState),
            ResourceState: frontdoor.NetworkExperimentResourceState(resourceState),
        },
        Tags: tags.Expand(t),
    }


    future, err := client.Update(ctx, resourceGroup, name, parameters)
    if err != nil {
        return fmt.Errorf("Error updating Network Experiment Profile %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for update of Network Experiment Profile %q (Resource Group %q): %+v", name, resourceGroup, err)
    }

    return resourceArmNetworkExperimentProfileRead(d, meta)
}

func resourceArmNetworkExperimentProfileDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).networkExperimentProfilesClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["NetworkExperimentProfiles"]

    future, err := client.Delete(ctx, resourceGroup, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Network Experiment Profile %q (Resource Group %q): %+v", name, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Network Experiment Profile %q (Resource Group %q): %+v", name, resourceGroup, err)
        }
    }

    return nil
}
