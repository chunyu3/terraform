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



func resourceArmManagedInstanceAdministrator() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmManagedInstanceAdministratorCreateUpdate,
        Read: resourceArmManagedInstanceAdministratorRead,
        Update: resourceArmManagedInstanceAdministratorCreateUpdate,
        Delete: resourceArmManagedInstanceAdministratorDelete,

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

            "administrator_type": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "login": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "managed_instance_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "sid": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "tenant_id": {
                Type: schema.TypeString,
                Optional: true,
            },

            "type": {
                Type: schema.TypeString,
                Computed: true,
            },
        },
    }
}

func resourceArmManagedInstanceAdministratorCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).managedInstanceAdministratorsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    managedInstanceName := d.Get("managed_instance_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, managedInstanceName, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Managed Instance Administrator %q (Managed Instance Name %q / Resource Group %q): %+v", name, managedInstanceName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_managed_instance_administrator", *existing.ID)
        }
    }

    administratorType := d.Get("administrator_type").(string)
    login := d.Get("login").(string)
    sid := d.Get("sid").(string)
    tenantId := d.Get("tenant_id").(string)

    parameters := sql.ManagedInstanceAdministrator{
        ManagedInstanceAdministratorProperties: &sql.ManagedInstanceAdministratorProperties{
            AdministratorType: utils.String(administratorType),
            Login: utils.String(login),
            Sid: utils.String(sid),
            TenantID: utils.String(tenantId),
        },
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, managedInstanceName, name, parameters)
    if err != nil {
        return fmt.Errorf("Error creating Managed Instance Administrator %q (Managed Instance Name %q / Resource Group %q): %+v", name, managedInstanceName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Managed Instance Administrator %q (Managed Instance Name %q / Resource Group %q): %+v", name, managedInstanceName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, managedInstanceName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Managed Instance Administrator %q (Managed Instance Name %q / Resource Group %q): %+v", name, managedInstanceName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Managed Instance Administrator %q (Managed Instance Name %q / Resource Group %q) ID", name, managedInstanceName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmManagedInstanceAdministratorRead(d, meta)
}

func resourceArmManagedInstanceAdministratorRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).managedInstanceAdministratorsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    managedInstanceName := id.Path["managedInstances"]
    name := id.Path["administrators"]

    resp, err := client.Get(ctx, resourceGroup, managedInstanceName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Managed Instance Administrator %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Managed Instance Administrator %q (Managed Instance Name %q / Resource Group %q): %+v", name, managedInstanceName, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if managedInstanceAdministratorProperties := resp.ManagedInstanceAdministratorProperties; managedInstanceAdministratorProperties != nil {
        d.Set("administrator_type", managedInstanceAdministratorProperties.AdministratorType)
        d.Set("login", managedInstanceAdministratorProperties.Login)
        d.Set("sid", managedInstanceAdministratorProperties.Sid)
        d.Set("tenant_id", managedInstanceAdministratorProperties.TenantID)
    }
    d.Set("managed_instance_name", managedInstanceName)
    d.Set("type", resp.Type)

    return nil
}


func resourceArmManagedInstanceAdministratorDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).managedInstanceAdministratorsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    managedInstanceName := id.Path["managedInstances"]
    name := id.Path["administrators"]

    future, err := client.Delete(ctx, resourceGroup, managedInstanceName, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Managed Instance Administrator %q (Managed Instance Name %q / Resource Group %q): %+v", name, managedInstanceName, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Managed Instance Administrator %q (Managed Instance Name %q / Resource Group %q): %+v", name, managedInstanceName, resourceGroup, err)
        }
    }

    return nil
}
