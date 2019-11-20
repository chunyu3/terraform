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



func resourceArmRegistration() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmRegistrationCreate,
        Read: resourceArmRegistrationRead,
        Update: resourceArmRegistrationUpdate,
        Delete: resourceArmRegistrationDelete,

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

            "location": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(azurestack.global),
                }, false),
                Default: string(azurestack.global),
            },

            "resource_group": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "registration_token": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "type": {
                Type: schema.TypeString,
                Computed: true,
            },
        },
    }
}

func resourceArmRegistrationCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).registrationsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Registration %q (Resource Group %q): %+v", name, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_registration", *existing.ID)
        }
    }

    location := d.Get("location").(string)
    registrationToken := d.Get("registration_token").(string)

    token := azurestack.RegistrationParameter{
        Location: azurestack.Location(location),
        RegistrationParameterProperties: &azurestack.RegistrationParameterProperties{
            RegistrationToken: utils.String(registrationToken),
        },
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroup, name, token); err != nil {
        return fmt.Errorf("Error creating Registration %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Registration %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Registration %q (Resource Group %q) ID", name, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmRegistrationRead(d, meta)
}

func resourceArmRegistrationRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).registrationsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["registrations"]

    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Registration %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Registration %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    d.Set("type", resp.Type)

    return nil
}

func resourceArmRegistrationUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).registrationsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    registrationToken := d.Get("registration_token").(string)
    resourceGroup := d.Get("resource_group").(string)

    token := azurestack.RegistrationParameter{
        RegistrationParameterProperties: &azurestack.RegistrationParameterProperties{
            RegistrationToken: utils.String(registrationToken),
        },
    }


    if _, err := client.Update(ctx, resourceGroup, name, token); err != nil {
        return fmt.Errorf("Error updating Registration %q (Resource Group %q): %+v", name, resourceGroup, err)
    }

    return resourceArmRegistrationRead(d, meta)
}

func resourceArmRegistrationDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).registrationsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["registrations"]

    if _, err := client.Delete(ctx, resourceGroup, name); err != nil {
        return fmt.Errorf("Error deleting Registration %q (Resource Group %q): %+v", name, resourceGroup, err)
    }

    return nil
}
