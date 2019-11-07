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



func resourceArmSecret() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmSecretCreate,
        Read: resourceArmSecretRead,
        Update: resourceArmSecretUpdate,
        Delete: resourceArmSecretDelete,

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

            "location": azure.SchemaLocation(),

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "lab_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "value": {
                Type: schema.TypeString,
                Optional: true,
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

func resourceArmSecretCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).secretsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    labName := d.Get("lab_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, labName, name, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Secret %q (Lab Name %q / Resource Group %q): %+v", name, labName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_secret", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    value := d.Get("value").(string)
    t := d.Get("tags").(map[string]interface{})

    secret := devtestlab.Secret{
        Location: utils.String(location),
        SecretProperties: &devtestlab.SecretProperties{
            Value: utils.String(value),
        },
        Tags: tags.Expand(t),
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, labName, name, name, secret)
    if err != nil {
        return fmt.Errorf("Error creating Secret %q (Lab Name %q / Resource Group %q): %+v", name, labName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Secret %q (Lab Name %q / Resource Group %q): %+v", name, labName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, labName, name, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Secret %q (Lab Name %q / Resource Group %q): %+v", name, labName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Secret %q (Lab Name %q / Resource Group %q) ID", name, labName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmSecretRead(d, meta)
}

func resourceArmSecretRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).secretsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    labName := id.Path["labs"]
    name := id.Path["users"]
    name := id.Path["secrets"]

    resp, err := client.Get(ctx, resourceGroup, labName, name, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Secret %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Secret %q (Lab Name %q / Resource Group %q): %+v", name, labName, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    d.Set("lab_name", labName)
    if secretProperties := resp.SecretProperties; secretProperties != nil {
        d.Set("provisioning_state", secretProperties.ProvisioningState)
        d.Set("unique_identifier", secretProperties.UniqueIdentifier)
        d.Set("value", secretProperties.Value)
    }
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}

func resourceArmSecretUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).secretsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    labName := d.Get("lab_name").(string)
    value := d.Get("value").(string)
    t := d.Get("tags").(map[string]interface{})

    secret := devtestlab.Secret{
        Location: utils.String(location),
        SecretProperties: &devtestlab.SecretProperties{
            Value: utils.String(value),
        },
        Tags: tags.Expand(t),
    }


    if _, err := client.Update(ctx, resourceGroup, labName, name, name, secret); err != nil {
        return fmt.Errorf("Error updating Secret %q (Lab Name %q / Resource Group %q): %+v", name, labName, resourceGroup, err)
    }

    return resourceArmSecretRead(d, meta)
}

func resourceArmSecretDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).secretsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    labName := id.Path["labs"]
    name := id.Path["users"]
    name := id.Path["secrets"]

    if _, err := client.Delete(ctx, resourceGroup, labName, name, name); err != nil {
        return fmt.Errorf("Error deleting Secret %q (Lab Name %q / Resource Group %q): %+v", name, labName, resourceGroup, err)
    }

    return nil
}
