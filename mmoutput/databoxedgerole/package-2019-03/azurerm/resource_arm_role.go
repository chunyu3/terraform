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



func resourceArmRole() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmRoleCreateUpdate,
        Read: resourceArmRoleRead,
        Update: resourceArmRoleCreateUpdate,
        Delete: resourceArmRoleDelete,

        Importer: &schema.ResourceImporter{
            State: schema.ImportStatePassthrough,
        },


        Schema: map[string]*schema.Schema{
            "device_name": {
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

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "id": {
                Type: schema.TypeString,
                Computed: true,
            },

            "name": {
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

func resourceArmRoleCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).rolesClient
    ctx, cancel := timeouts.ForCreateUpdate(meta.(*ArmClient).StopContext, d)
    defer cancel()

    resourceGroupName := d.Get("resource_group").(string)
    name := d.Get("device_name").(string)
    name := d.Get("name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroupName, name, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Role (Name %q / Resource Group %q / Device Name %q): %+v", name, resourceGroupName, name, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_role", *existing.ID)
        }
    }


    role := databoxedge.Role{
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroupName, name, name, role)
    if err != nil {
        return fmt.Errorf("Error creating Role (Name %q / Resource Group %q / Device Name %q): %+v", name, resourceGroupName, name, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Role (Name %q / Resource Group %q / Device Name %q): %+v", name, resourceGroupName, name, err)
    }


    resp, err := client.Get(ctx, resourceGroupName, name, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Role (Name %q / Resource Group %q / Device Name %q): %+v", name, resourceGroupName, name, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Role (Name %q / Resource Group %q / Device Name %q) ID", name, resourceGroupName, name)
    }
    d.SetId(*resp.ID)

    return resourceArmRoleRead(d, meta)
}

func resourceArmRoleRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).rolesClient
    ctx, cancel := timeouts.ForRead(meta.(*ArmClient).StopContext, d)
    defer cancel()

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroupName := id.ResourceGroup
    name := id.Path["dataBoxEdgeDevices"]
    name := id.Path["roles"]

    resp, err := client.Get(ctx, resourceGroupName, name, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Role %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Role (Name %q / Resource Group %q / Device Name %q): %+v", name, resourceGroupName, name, err)
    }


    d.Set("resource_group", resourceGroupName)
    d.Set("device_name", name)
    d.Set("id", resp.ID)
    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("type", resp.Type)

    return nil
}


func resourceArmRoleDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).rolesClient
    ctx, cancel := timeouts.ForDelete(meta.(*ArmClient).StopContext, d)
    defer cancel()


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroupName := id.ResourceGroup
    name := id.Path["dataBoxEdgeDevices"]
    name := id.Path["roles"]

    future, err := client.Delete(ctx, resourceGroupName, name, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Role (Name %q / Resource Group %q / Device Name %q): %+v", name, resourceGroupName, name, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Role (Name %q / Resource Group %q / Device Name %q): %+v", name, resourceGroupName, name, err)
        }
    }

    return nil
}
