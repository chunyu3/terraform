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



func resourceArmIdentityProvider() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmIdentityProviderCreate,
        Read: resourceArmIdentityProviderRead,
        Update: resourceArmIdentityProviderUpdate,
        Delete: resourceArmIdentityProviderDelete,

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

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "client_id": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "client_secret": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "service_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "allowed_tenants": {
                Type: schema.TypeList,
                Optional: true,
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
            },

            "authority": {
                Type: schema.TypeString,
                Optional: true,
            },

            "password_reset_policy_name": {
                Type: schema.TypeString,
                Optional: true,
            },

            "profile_editing_policy_name": {
                Type: schema.TypeString,
                Optional: true,
            },

            "signin_policy_name": {
                Type: schema.TypeString,
                Optional: true,
            },

            "signup_policy_name": {
                Type: schema.TypeString,
                Optional: true,
            },

            "type": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(apimanagement.facebook),
                    string(apimanagement.google),
                    string(apimanagement.microsoft),
                    string(apimanagement.twitter),
                    string(apimanagement.aad),
                    string(apimanagement.aadB2C),
                }, false),
                Default: string(apimanagement.facebook),
            },

            "type": {
                Type: schema.TypeString,
                Computed: true,
            },
        },
    }
}

func resourceArmIdentityProviderCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).identityProviderClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    serviceName := d.Get("service_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, serviceName, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Identity Provider %q (Service Name %q / Resource Group %q): %+v", name, serviceName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_identity_provider", *existing.ID)
        }
    }

    allowedTenants := d.Get("allowed_tenants").([]interface{})
    authority := d.Get("authority").(string)
    clientId := d.Get("client_id").(string)
    clientSecret := d.Get("client_secret").(string)
    passwordResetPolicyName := d.Get("password_reset_policy_name").(string)
    profileEditingPolicyName := d.Get("profile_editing_policy_name").(string)
    signinPolicyName := d.Get("signin_policy_name").(string)
    signupPolicyName := d.Get("signup_policy_name").(string)
    type := d.Get("type").(string)

    parameters := apimanagement.IdentityProviderContract{
        IdentityProviderContractProperties: &apimanagement.IdentityProviderContractProperties{
            AllowedTenants: utils.ExpandStringSlice(allowedTenants),
            Authority: utils.String(authority),
            ClientID: utils.String(clientId),
            ClientSecret: utils.String(clientSecret),
            PasswordResetPolicyName: utils.String(passwordResetPolicyName),
            ProfileEditingPolicyName: utils.String(profileEditingPolicyName),
            SigninPolicyName: utils.String(signinPolicyName),
            SignupPolicyName: utils.String(signupPolicyName),
            Type: apimanagement.IdentityProviderType(type),
        },
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroup, serviceName, name, parameters); err != nil {
        return fmt.Errorf("Error creating Identity Provider %q (Service Name %q / Resource Group %q): %+v", name, serviceName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, serviceName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Identity Provider %q (Service Name %q / Resource Group %q): %+v", name, serviceName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Identity Provider %q (Service Name %q / Resource Group %q) ID", name, serviceName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmIdentityProviderRead(d, meta)
}

func resourceArmIdentityProviderRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).identityProviderClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    serviceName := id.Path["service"]
    name := id.Path["identityProviders"]

    resp, err := client.Get(ctx, resourceGroup, serviceName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Identity Provider %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Identity Provider %q (Service Name %q / Resource Group %q): %+v", name, serviceName, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if identityProviderContractProperties := resp.IdentityProviderContractProperties; identityProviderContractProperties != nil {
        d.Set("allowed_tenants", utils.FlattenStringSlice(identityProviderContractProperties.AllowedTenants))
        d.Set("authority", identityProviderContractProperties.Authority)
        d.Set("client_id", identityProviderContractProperties.ClientID)
        d.Set("client_secret", identityProviderContractProperties.ClientSecret)
        d.Set("password_reset_policy_name", identityProviderContractProperties.PasswordResetPolicyName)
        d.Set("profile_editing_policy_name", identityProviderContractProperties.ProfileEditingPolicyName)
        d.Set("signin_policy_name", identityProviderContractProperties.SigninPolicyName)
        d.Set("signup_policy_name", identityProviderContractProperties.SignupPolicyName)
        d.Set("type", string(identityProviderContractProperties.Type))
    }
    d.Set("service_name", serviceName)
    d.Set("type", resp.Type)
    d.Set("type", resp.Type)
    d.Set("type", resp.Type)
    d.Set("type", resp.Type)

    return nil
}

func resourceArmIdentityProviderUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).identityProviderClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    allowedTenants := d.Get("allowed_tenants").([]interface{})
    authority := d.Get("authority").(string)
    clientId := d.Get("client_id").(string)
    clientSecret := d.Get("client_secret").(string)
    passwordResetPolicyName := d.Get("password_reset_policy_name").(string)
    profileEditingPolicyName := d.Get("profile_editing_policy_name").(string)
    serviceName := d.Get("service_name").(string)
    signinPolicyName := d.Get("signin_policy_name").(string)
    signupPolicyName := d.Get("signup_policy_name").(string)
    type := d.Get("type").(string)

    parameters := apimanagement.IdentityProviderContract{
        IdentityProviderContractProperties: &apimanagement.IdentityProviderContractProperties{
            AllowedTenants: utils.ExpandStringSlice(allowedTenants),
            Authority: utils.String(authority),
            ClientID: utils.String(clientId),
            ClientSecret: utils.String(clientSecret),
            PasswordResetPolicyName: utils.String(passwordResetPolicyName),
            ProfileEditingPolicyName: utils.String(profileEditingPolicyName),
            SigninPolicyName: utils.String(signinPolicyName),
            SignupPolicyName: utils.String(signupPolicyName),
            Type: apimanagement.IdentityProviderType(type),
        },
    }


    if _, err := client.Update(ctx, resourceGroup, serviceName, name, parameters); err != nil {
        return fmt.Errorf("Error updating Identity Provider %q (Service Name %q / Resource Group %q): %+v", name, serviceName, resourceGroup, err)
    }

    return resourceArmIdentityProviderRead(d, meta)
}

func resourceArmIdentityProviderDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).identityProviderClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    serviceName := id.Path["service"]
    name := id.Path["identityProviders"]

    if _, err := client.Delete(ctx, resourceGroup, serviceName, name); err != nil {
        return fmt.Errorf("Error deleting Identity Provider %q (Service Name %q / Resource Group %q): %+v", name, serviceName, resourceGroup, err)
    }

    return nil
}
