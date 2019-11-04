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



func resourceArmShareSubscription() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmShareSubscriptionCreateUpdate,
        Read: resourceArmShareSubscriptionRead,
        Update: resourceArmShareSubscriptionCreateUpdate,
        Delete: resourceArmShareSubscriptionDelete,

        Importer: &schema.ResourceImporter{
            State: schema.ImportStatePassthrough,
        },


        Schema: map[string]*schema.Schema{
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

            "invitation_id": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "share_subscription_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "source_share_location": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "created_at": {
                Type: schema.TypeString,
                Computed: true,
            },

            "provider_email": {
                Type: schema.TypeString,
                Computed: true,
            },

            "provider_name": {
                Type: schema.TypeString,
                Computed: true,
            },

            "provider_tenant_name": {
                Type: schema.TypeString,
                Computed: true,
            },

            "provisioning_state": {
                Type: schema.TypeString,
                Computed: true,
            },

            "share_description": {
                Type: schema.TypeString,
                Computed: true,
            },

            "share_kind": {
                Type: schema.TypeString,
                Computed: true,
            },

            "share_name": {
                Type: schema.TypeString,
                Computed: true,
            },

            "share_subscription_status": {
                Type: schema.TypeString,
                Computed: true,
            },

            "share_terms": {
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

func resourceArmShareSubscriptionCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).shareSubscriptionsClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    accountName := d.Get("account_name").(string)
    shareSubscriptionName := d.Get("share_subscription_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, accountName, shareSubscriptionName)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Share Subscription (Share Subscription Name %q / Account Name %q / Resource Group %q): %+v", shareSubscriptionName, accountName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_share_subscription", *existing.ID)
        }
    }

    invitationId := d.Get("invitation_id").(string)
    sourceShareLocation := d.Get("source_share_location").(string)

    shareSubscription := datashare.ShareSubscription{
        ShareSubscriptionProperties: &datashare.ShareSubscriptionProperties{
            InvitationID: utils.String(invitationId),
            SourceShareLocation: utils.String(sourceShareLocation),
        },
    }


    if _, err := client.Create(ctx, resourceGroup, accountName, shareSubscriptionName, shareSubscription); err != nil {
        return fmt.Errorf("Error creating Share Subscription (Share Subscription Name %q / Account Name %q / Resource Group %q): %+v", shareSubscriptionName, accountName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, accountName, shareSubscriptionName)
    if err != nil {
        return fmt.Errorf("Error retrieving Share Subscription (Share Subscription Name %q / Account Name %q / Resource Group %q): %+v", shareSubscriptionName, accountName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Share Subscription (Share Subscription Name %q / Account Name %q / Resource Group %q) ID", shareSubscriptionName, accountName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmShareSubscriptionRead(d, meta)
}

func resourceArmShareSubscriptionRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).shareSubscriptionsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    accountName := id.Path["accounts"]
    shareSubscriptionName := id.Path["shareSubscriptions"]

    resp, err := client.Get(ctx, resourceGroup, accountName, shareSubscriptionName)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Share Subscription %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Share Subscription (Share Subscription Name %q / Account Name %q / Resource Group %q): %+v", shareSubscriptionName, accountName, resourceGroup, err)
    }


    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    d.Set("account_name", accountName)
    if shareSubscriptionProperties := resp.ShareSubscriptionProperties; shareSubscriptionProperties != nil {
        d.Set("created_at", (shareSubscriptionProperties.CreatedAt).String())
        d.Set("invitation_id", shareSubscriptionProperties.InvitationID)
        d.Set("provider_email", shareSubscriptionProperties.ProviderEmail)
        d.Set("provider_name", shareSubscriptionProperties.ProviderName)
        d.Set("provider_tenant_name", shareSubscriptionProperties.ProviderTenantName)
        d.Set("provisioning_state", string(shareSubscriptionProperties.ProvisioningState))
        d.Set("share_description", shareSubscriptionProperties.ShareDescription)
        d.Set("share_kind", string(shareSubscriptionProperties.ShareKind))
        d.Set("share_name", shareSubscriptionProperties.ShareName)
        d.Set("share_subscription_status", string(shareSubscriptionProperties.ShareSubscriptionStatus))
        d.Set("share_terms", shareSubscriptionProperties.ShareTerms)
        d.Set("source_share_location", shareSubscriptionProperties.SourceShareLocation)
        d.Set("user_email", shareSubscriptionProperties.UserEmail)
        d.Set("user_name", shareSubscriptionProperties.UserName)
    }
    d.Set("share_subscription_name", shareSubscriptionName)
    d.Set("type", resp.Type)

    return nil
}


func resourceArmShareSubscriptionDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).shareSubscriptionsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    accountName := id.Path["accounts"]
    shareSubscriptionName := id.Path["shareSubscriptions"]

    future, err := client.Delete(ctx, resourceGroup, accountName, shareSubscriptionName)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Share Subscription (Share Subscription Name %q / Account Name %q / Resource Group %q): %+v", shareSubscriptionName, accountName, resourceGroup, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Share Subscription (Share Subscription Name %q / Account Name %q / Resource Group %q): %+v", shareSubscriptionName, accountName, resourceGroup, err)
        }
    }

    return nil
}
