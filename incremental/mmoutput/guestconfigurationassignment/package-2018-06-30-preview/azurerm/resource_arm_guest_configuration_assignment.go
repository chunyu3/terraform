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



func resourceArmGuestConfigurationAssignment() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmGuestConfigurationAssignmentCreateUpdate,
        Read: resourceArmGuestConfigurationAssignmentRead,
        Update: resourceArmGuestConfigurationAssignmentCreateUpdate,
        Delete: resourceArmGuestConfigurationAssignmentDelete,

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

            "guest_configuration_assignment_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "context": {
                Type: schema.TypeString,
                Optional: true,
            },

            "guest_configuration": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "configuration_setting": {
                            Type: schema.TypeList,
                            Optional: true,
                            MaxItems: 1,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "allow_module_overwrite": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                        ValidateFunc: validation.StringInSlice([]string{
                                            string(guestconfiguration.True),
                                            string(guestconfiguration.False),
                                        }, false),
                                        Default: string(guestconfiguration.True),
                                    },
                                },
                            },
                        },
                        "kind": {
                            Type: schema.TypeString,
                            Optional: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(guestconfiguration.DSC),
                            }, false),
                            Default: string(guestconfiguration.DSC),
                        },
                        "name": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "version": {
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
        },
    }
}

func resourceArmGuestConfigurationAssignmentCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).guestConfigurationAssignmentsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    guestConfigurationAssignmentName := d.Get("guest_configuration_assignment_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, name, guestConfigurationAssignmentName)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Guest Configuration Assignment %q (Resource Group %q / Guest Configuration Assignment Name %q): %+v", name, resourceGroup, guestConfigurationAssignmentName, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_guest_configuration_assignment", *existing.ID)
        }
    }

    context := d.Get("context").(string)
    guestConfiguration := d.Get("guest_configuration").([]interface{})

    parameters := guestconfiguration.Assignment{
        AssignmentProperties: &guestconfiguration.AssignmentProperties{
            Context: utils.String(context),
            GuestConfiguration: expandArmGuestConfigurationAssignmentNavigation(guestConfiguration),
        },
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, name, guestConfigurationAssignmentName, parameters)
    if err != nil {
        return fmt.Errorf("Error creating Guest Configuration Assignment %q (Resource Group %q / Guest Configuration Assignment Name %q): %+v", name, resourceGroup, guestConfigurationAssignmentName, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Guest Configuration Assignment %q (Resource Group %q / Guest Configuration Assignment Name %q): %+v", name, resourceGroup, guestConfigurationAssignmentName, err)
    }


    resp, err := client.Get(ctx, resourceGroup, name, guestConfigurationAssignmentName)
    if err != nil {
        return fmt.Errorf("Error retrieving Guest Configuration Assignment %q (Resource Group %q / Guest Configuration Assignment Name %q): %+v", name, resourceGroup, guestConfigurationAssignmentName, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Guest Configuration Assignment %q (Resource Group %q / Guest Configuration Assignment Name %q) ID", name, resourceGroup, guestConfigurationAssignmentName)
    }
    d.SetId(*resp.ID)

    return resourceArmGuestConfigurationAssignmentRead(d, meta)
}

func resourceArmGuestConfigurationAssignmentRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).guestConfigurationAssignmentsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["virtualMachines"]
    guestConfigurationAssignmentName := id.Path["guestConfigurationAssignments"]

    resp, err := client.Get(ctx, resourceGroup, name, guestConfigurationAssignmentName)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Guest Configuration Assignment %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Guest Configuration Assignment %q (Resource Group %q / Guest Configuration Assignment Name %q): %+v", name, resourceGroup, guestConfigurationAssignmentName, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    d.Set("guest_configuration_assignment_name", guestConfigurationAssignmentName)
    d.Set("type", resp.Type)

    return nil
}


func resourceArmGuestConfigurationAssignmentDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).guestConfigurationAssignmentsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["virtualMachines"]
    guestConfigurationAssignmentName := id.Path["guestConfigurationAssignments"]

    future, err := client.Delete(ctx, resourceGroup, name, guestConfigurationAssignmentName)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Guest Configuration Assignment %q (Resource Group %q / Guest Configuration Assignment Name %q): %+v", name, resourceGroup, guestConfigurationAssignmentName, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Guest Configuration Assignment %q (Resource Group %q / Guest Configuration Assignment Name %q): %+v", name, resourceGroup, guestConfigurationAssignmentName, err)
        }
    }

    return nil
}

func expandArmGuestConfigurationAssignmentNavigation(input []interface{}) *guestconfiguration.Navigation {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    kind := v["kind"].(string)
    name := v["name"].(string)
    version := v["version"].(string)
    configurationSetting := v["configuration_setting"].([]interface{})

    result := guestconfiguration.Navigation{
        ConfigurationSetting: expandArmGuestConfigurationAssignmentConfigurationSetting(configurationSetting),
        Kind: guestconfiguration.Kind(kind),
        Name: utils.String(name),
        Version: utils.String(version),
    }
    return &result
}

func expandArmGuestConfigurationAssignmentConfigurationSetting(input []interface{}) *guestconfiguration.ConfigurationSetting {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    allowModuleOverwrite := v["allow_module_overwrite"].(string)

    result := guestconfiguration.ConfigurationSetting{
        AllowModuleOverwrite: guestconfiguration.AllowModuleOverwrite(allowModuleOverwrite),
    }
    return &result
}