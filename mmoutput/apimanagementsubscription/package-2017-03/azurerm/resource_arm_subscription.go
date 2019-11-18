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



func resourceArmSubscription() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmSubscriptionCreate,
        Read: resourceArmSubscriptionRead,
        Update: resourceArmSubscriptionUpdate,
        Delete: resourceArmSubscriptionDelete,

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

            "sid": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "display_name": {
                Type: schema.TypeString,
                Optional: true,
            },

            "expiration_date": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validateRFC3339Date,
            },

            "primary_key": {
                Type: schema.TypeString,
                Optional: true,
            },

            "product_id": {
                Type: schema.TypeString,
                Optional: true,
            },

            "secondary_key": {
                Type: schema.TypeString,
                Optional: true,
            },

            "state": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(apimanagement.suspended),
                    string(apimanagement.active),
                    string(apimanagement.expired),
                    string(apimanagement.submitted),
                    string(apimanagement.rejected),
                    string(apimanagement.cancelled),
                }, false),
                Default: string(apimanagement.suspended),
            },

            "state_comment": {
                Type: schema.TypeString,
                Optional: true,
            },

            "user_id": {
                Type: schema.TypeString,
                Optional: true,
            },

            "created_date": {
                Type: schema.TypeString,
                Computed: true,
            },

            "end_date": {
                Type: schema.TypeString,
                Computed: true,
            },

            "notification_date": {
                Type: schema.TypeString,
                Computed: true,
            },

            "start_date": {
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

func resourceArmSubscriptionCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).subscriptionClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    sid := d.Get("sid").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, name, sid)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Subscription %q (Sid %q / Resource Group %q): %+v", name, sid, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_subscription", *existing.ID)
        }
    }

    displayName := d.Get("display_name").(string)
    expirationDate := d.Get("expiration_date").(string)
    primaryKey := d.Get("primary_key").(string)
    productId := d.Get("product_id").(string)
    secondaryKey := d.Get("secondary_key").(string)
    state := d.Get("state").(string)
    stateComment := d.Get("state_comment").(string)
    userId := d.Get("user_id").(string)

    parameters := apimanagement.SubscriptionUpdateParameters{
        SubscriptionUpdateParameterProperties: &apimanagement.SubscriptionUpdateParameterProperties{
            DisplayName: utils.String(displayName),
            ExpirationDate: convertStringToDate(expirationDate),
            PrimaryKey: utils.String(primaryKey),
            ProductID: utils.String(productId),
            SecondaryKey: utils.String(secondaryKey),
            State: apimanagement.SubscriptionState(state),
            StateComment: utils.String(stateComment),
            UserID: utils.String(userId),
        },
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroup, name, sid, parameters); err != nil {
        return fmt.Errorf("Error creating Subscription %q (Sid %q / Resource Group %q): %+v", name, sid, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, name, sid)
    if err != nil {
        return fmt.Errorf("Error retrieving Subscription %q (Sid %q / Resource Group %q): %+v", name, sid, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Subscription %q (Sid %q / Resource Group %q) ID", name, sid, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmSubscriptionRead(d, meta)
}

func resourceArmSubscriptionRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).subscriptionClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["service"]
    sid := id.Path["subscriptions"]

    resp, err := client.Get(ctx, resourceGroup, name, sid)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Subscription %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Subscription %q (Sid %q / Resource Group %q): %+v", name, sid, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if subscriptionUpdateParameterProperties := resp.SubscriptionUpdateParameterProperties; subscriptionUpdateParameterProperties != nil {
        d.Set("created_date", (subscriptionUpdateParameterProperties.CreatedDate).String())
        d.Set("display_name", subscriptionUpdateParameterProperties.DisplayName)
        d.Set("end_date", (subscriptionUpdateParameterProperties.EndDate).String())
        d.Set("expiration_date", (subscriptionUpdateParameterProperties.ExpirationDate).String())
        d.Set("notification_date", (subscriptionUpdateParameterProperties.NotificationDate).String())
        d.Set("primary_key", subscriptionUpdateParameterProperties.PrimaryKey)
        d.Set("product_id", subscriptionUpdateParameterProperties.ProductID)
        d.Set("secondary_key", subscriptionUpdateParameterProperties.SecondaryKey)
        d.Set("start_date", (subscriptionUpdateParameterProperties.StartDate).String())
        d.Set("state", string(subscriptionUpdateParameterProperties.State))
        d.Set("state_comment", subscriptionUpdateParameterProperties.StateComment)
        d.Set("user_id", subscriptionUpdateParameterProperties.UserID)
    }
    d.Set("sid", sid)
    d.Set("type", resp.Type)

    return nil
}

func resourceArmSubscriptionUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).subscriptionClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    displayName := d.Get("display_name").(string)
    expirationDate := d.Get("expiration_date").(string)
    primaryKey := d.Get("primary_key").(string)
    productId := d.Get("product_id").(string)
    secondaryKey := d.Get("secondary_key").(string)
    sid := d.Get("sid").(string)
    state := d.Get("state").(string)
    stateComment := d.Get("state_comment").(string)
    userId := d.Get("user_id").(string)

    parameters := apimanagement.SubscriptionUpdateParameters{
        SubscriptionUpdateParameterProperties: &apimanagement.SubscriptionUpdateParameterProperties{
            DisplayName: utils.String(displayName),
            ExpirationDate: convertStringToDate(expirationDate),
            PrimaryKey: utils.String(primaryKey),
            ProductID: utils.String(productId),
            SecondaryKey: utils.String(secondaryKey),
            State: apimanagement.SubscriptionState(state),
            StateComment: utils.String(stateComment),
            UserID: utils.String(userId),
        },
    }


    if _, err := client.Update(ctx, resourceGroup, name, sid, parameters); err != nil {
        return fmt.Errorf("Error updating Subscription %q (Sid %q / Resource Group %q): %+v", name, sid, resourceGroup, err)
    }

    return resourceArmSubscriptionRead(d, meta)
}

func resourceArmSubscriptionDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).subscriptionClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["service"]
    sid := id.Path["subscriptions"]

    if _, err := client.Delete(ctx, resourceGroup, name, sid); err != nil {
        return fmt.Errorf("Error deleting Subscription %q (Sid %q / Resource Group %q): %+v", name, sid, resourceGroup, err)
    }

    return nil
}

func convertStringToDate(input interface{}) *date.Time {
  v := input.(string)

  dateTime, err := date.ParseTime(time.RFC3339, v)
  if err != nil {
      log.Printf("[ERROR] Cannot convert an invalid string to RFC3339 date %q: %+v", v, err)
      return nil
  }

  result := date.Time{
      Time: dateTime,
  }
  return &result
}
