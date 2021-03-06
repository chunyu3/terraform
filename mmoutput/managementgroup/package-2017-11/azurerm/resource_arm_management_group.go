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



func resourceArmManagementGroup() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmManagementGroupCreate,
        Read: resourceArmManagementGroupRead,
        Update: resourceArmManagementGroupUpdate,
        Delete: resourceArmManagementGroupDelete,

        Importer: &schema.ResourceImporter{
            State: schema.ImportStatePassthrough,
        },


        Schema: map[string]*schema.Schema{
            "name": {
                Type: schema.TypeString,
                Computed: true,
            },

            "group_id": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "display_name": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
            },

            "parent_id": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
            },

            "type": {
                Type: schema.TypeString,
                Computed: true,
            },
        },
    }
}

func resourceArmManagementGroupCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).managementGroupsClient
    ctx := meta.(*ArmClient).StopContext

    groupID := d.Get("group_id").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, groupID)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Management Group (Group %q): %+v", groupID, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_management_group", *existing.ID)
        }
    }

    displayName := d.Get("display_name").(string)
    parentId := d.Get("parent_id").(string)

    createManagementGroupRequest := managementgroups.CreateManagementGroupRequest{
        DisplayName: utils.String(displayName),
        ParentID: utils.String(parentId),
    }


    if _, err := client.CreateOrUpdate(ctx, groupID, createManagementGroupRequest); err != nil {
        return fmt.Errorf("Error creating Management Group (Group %q): %+v", groupID, err)
    }


    resp, err := client.Get(ctx, groupID)
    if err != nil {
        return fmt.Errorf("Error retrieving Management Group (Group %q): %+v", groupID, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Management Group (Group %q) ID", groupID)
    }
    d.SetId(*resp.ID)

    return resourceArmManagementGroupRead(d, meta)
}

func resourceArmManagementGroupRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).managementGroupsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    groupID := id.Path["managementGroups"]

    resp, err := client.Get(ctx, groupID)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Management Group %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Management Group (Group %q): %+v", groupID, err)
    }


    d.Set("name", resp.Name)
    d.Set("group_id", groupID)
    d.Set("type", resp.Type)

    return nil
}

func resourceArmManagementGroupUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).managementGroupsClient
    ctx := meta.(*ArmClient).StopContext

    displayName := d.Get("display_name").(string)
    groupID := d.Get("group_id").(string)
    parentId := d.Get("parent_id").(string)

    createManagementGroupRequest := managementgroups.CreateManagementGroupRequest{
        DisplayName: utils.String(displayName),
        ParentID: utils.String(parentId),
    }


    if _, err := client.Update(ctx, groupID, createManagementGroupRequest); err != nil {
        return fmt.Errorf("Error updating Management Group (Group %q): %+v", groupID, err)
    }

    return resourceArmManagementGroupRead(d, meta)
}

func resourceArmManagementGroupDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).managementGroupsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    groupID := id.Path["managementGroups"]

    if _, err := client.Delete(ctx, groupID); err != nil {
        return fmt.Errorf("Error deleting Management Group (Group %q): %+v", groupID, err)
    }

    return nil
}
