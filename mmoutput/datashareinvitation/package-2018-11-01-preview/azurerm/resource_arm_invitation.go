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



func resourceArmInvitation() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmInvitationCreateUpdate,
        Read: resourceArmInvitationRead,
        Update: resourceArmInvitationCreateUpdate,
        Delete: resourceArmInvitationDelete,

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

            "account_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "share_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "target_active_directory_id": {
                Type: schema.TypeString,
                Optional: true,
            },

            "target_email": {
                Type: schema.TypeString,
                Optional: true,
            },

            "target_object_id": {
                Type: schema.TypeString,
                Optional: true,
            },

            "invitation_id": {
                Type: schema.TypeString,
                Computed: true,
            },

            "invitation_status": {
                Type: schema.TypeString,
                Computed: true,
            },

            "responded_at": {
                Type: schema.TypeString,
                Computed: true,
            },

            "sent_at": {
                Type: schema.TypeString,
                Computed: true,
            },

            "type": {
                Type: schema.TypeString,
                Computed: true,
            },

            "user_email": {
                Type: schema.TypeString,
                Computed: true,
            },

            "user_name": {
                Type: schema.TypeString,
                Computed: true,
            },
        },
    }
}

func resourceArmInvitationCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).invitationsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    accountName := d.Get("account_name").(string)
    shareName := d.Get("share_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, accountName, shareName, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Invitation %q (Share Name %q / Account Name %q / Resource Group %q): %+v", name, shareName, accountName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_invitation", *existing.ID)
        }
    }

    targetActiveDirectoryId := d.Get("target_active_directory_id").(string)
    targetEmail := d.Get("target_email").(string)
    targetObjectId := d.Get("target_object_id").(string)

    invitation := datashare.Invitation{
        InvitationProperties: &datashare.InvitationProperties{
            TargetActiveDirectoryID: utils.String(targetActiveDirectoryId),
            TargetEmail: utils.String(targetEmail),
            TargetObjectID: utils.String(targetObjectId),
        },
    }


    if _, err := client.Create(ctx, resourceGroup, accountName, shareName, name, invitation); err != nil {
        return fmt.Errorf("Error creating Invitation %q (Share Name %q / Account Name %q / Resource Group %q): %+v", name, shareName, accountName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, accountName, shareName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Invitation %q (Share Name %q / Account Name %q / Resource Group %q): %+v", name, shareName, accountName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Invitation %q (Share Name %q / Account Name %q / Resource Group %q) ID", name, shareName, accountName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmInvitationRead(d, meta)
}

func resourceArmInvitationRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).invitationsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    accountName := id.Path["accounts"]
    shareName := id.Path["shares"]
    name := id.Path["invitations"]

    resp, err := client.Get(ctx, resourceGroup, accountName, shareName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Invitation %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Invitation %q (Share Name %q / Account Name %q / Resource Group %q): %+v", name, shareName, accountName, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    d.Set("account_name", accountName)
    if invitationProperties := resp.InvitationProperties; invitationProperties != nil {
        d.Set("invitation_id", invitationProperties.InvitationID)
        d.Set("invitation_status", string(invitationProperties.InvitationStatus))
        d.Set("responded_at", (invitationProperties.RespondedAt).String())
        d.Set("sent_at", (invitationProperties.SentAt).String())
        d.Set("target_active_directory_id", invitationProperties.TargetActiveDirectoryID)
        d.Set("target_email", invitationProperties.TargetEmail)
        d.Set("target_object_id", invitationProperties.TargetObjectID)
        d.Set("user_email", invitationProperties.UserEmail)
        d.Set("user_name", invitationProperties.UserName)
    }
    d.Set("share_name", shareName)
    d.Set("type", resp.Type)

    return nil
}


func resourceArmInvitationDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).invitationsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    accountName := id.Path["accounts"]
    shareName := id.Path["shares"]
    name := id.Path["invitations"]

    if _, err := client.Delete(ctx, resourceGroup, accountName, shareName, name); err != nil {
        return fmt.Errorf("Error deleting Invitation %q (Share Name %q / Account Name %q / Resource Group %q): %+v", name, shareName, accountName, resourceGroup, err)
    }

    return nil
}
