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



func resourceArmProtectionPolicy() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmProtectionPolicyCreateUpdate,
        Read: resourceArmProtectionPolicyRead,
        Update: resourceArmProtectionPolicyCreateUpdate,
        Delete: resourceArmProtectionPolicyDelete,

        Importer: &schema.ResourceImporter{
            State: schema.ImportStatePassthrough,
        },


        Schema: map[string]*schema.Schema{
            "policy_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "vault_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "e_tag": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
            },

            "id": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
            },

            "location": azure.SchemaLocation(),

            "name": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
            },

            "protected_items_count": {
                Type: schema.TypeInt,
                Optional: true,
            },

            "tags": tags.Schema(),

            "type": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
            },
        },
    }
}

func resourceArmProtectionPolicyCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).protectionPoliciesClient
    ctx, cancel := timeouts.ForCreateUpdate(meta.(*ArmClient).StopContext, d)
    defer cancel()

    resourceGroupName := d.Get("resource_group").(string)
    name := d.Get("policy_name").(string)
    vaultName := d.Get("vault_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroupName, vaultName, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Protection Policy (Policy Name %q / Resource Group %q / Vault Name %q): %+v", name, resourceGroupName, vaultName, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_protection_policy", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    eTag := d.Get("e_tag").(string)
    iD := d.Get("id").(string)
    name := d.Get("name").(string)
    protectedItemsCount := d.Get("protected_items_count").(int)
    type := d.Get("type").(string)
    tags := d.Get("tags").(map[string]interface{})

    resourceProtectionPolicy := backup.ProtectionPolicyResource{
        ETag: utils.String(eTag),
        ID: utils.String(iD),
        Location: utils.String(location),
        Name: utils.String(name),
        ProtectionPolicy: &backup.ProtectionPolicy{
            ProtectedItemsCount: utils.Int32(int32(protectedItemsCount)),
        },
        Tags: tags.Expand(tags),
        Type: utils.String(type),
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroupName, vaultName, name, resourceProtectionPolicy); err != nil {
        return fmt.Errorf("Error creating Protection Policy (Policy Name %q / Resource Group %q / Vault Name %q): %+v", name, resourceGroupName, vaultName, err)
    }


    resp, err := client.Get(ctx, resourceGroupName, vaultName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Protection Policy (Policy Name %q / Resource Group %q / Vault Name %q): %+v", name, resourceGroupName, vaultName, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Protection Policy (Policy Name %q / Resource Group %q / Vault Name %q) ID", name, resourceGroupName, vaultName)
    }
    d.SetId(*resp.ID)

    return resourceArmProtectionPolicyRead(d, meta)
}

func resourceArmProtectionPolicyRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).protectionPoliciesClient
    ctx, cancel := timeouts.ForRead(meta.(*ArmClient).StopContext, d)
    defer cancel()

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroupName := id.ResourceGroup
    vaultName := id.Path["vaults"]
    name := id.Path["backupPolicies"]

    resp, err := client.Get(ctx, resourceGroupName, vaultName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Protection Policy %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Protection Policy (Policy Name %q / Resource Group %q / Vault Name %q): %+v", name, resourceGroupName, vaultName, err)
    }


    d.Set("resource_group", resourceGroupName)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    d.Set("e_tag", resp.ETag)
    d.Set("id", resp.ID)
    d.Set("name", resp.Name)
    d.Set("policy_name", name)
    if protectionPolicy := resp.ProtectionPolicy; protectionPolicy != nil {
        d.Set("protected_items_count", int(*protectionPolicy.ProtectedItemsCount))
    }
    d.Set("type", resp.Type)
    d.Set("vault_name", vaultName)

    return tags.FlattenAndSet(d, resp.Tags)
}


func resourceArmProtectionPolicyDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).protectionPoliciesClient
    ctx, cancel := timeouts.ForDelete(meta.(*ArmClient).StopContext, d)
    defer cancel()


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroupName := id.ResourceGroup
    vaultName := id.Path["vaults"]
    name := id.Path["backupPolicies"]

    if _, err := client.Delete(ctx, resourceGroupName, vaultName, name); err != nil {
        return fmt.Errorf("Error deleting Protection Policy (Policy Name %q / Resource Group %q / Vault Name %q): %+v", name, resourceGroupName, vaultName, err)
    }

    return nil
}
