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
                Optional: true,
                ForceNew: true,
            },

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "sid": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "expiration_date": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
                ValidateFunc: validateRFC3339Date,
            },

            "primary_key": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
            },

            "product_id": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
            },

            "secondary_key": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
            },

            "state": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(apimanagement.Suspended),
                    string(apimanagement.Active),
                    string(apimanagement.Expired),
                    string(apimanagement.Submitted),
                    string(apimanagement.Rejected),
                    string(apimanagement.Cancelled),
                }, false),
                Default: string(apimanagement.Suspended),
            },

            "state_comment": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
            },

            "user_id": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
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
        },
    }
}

func resourceArmSubscriptionCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).subscriptionsClient
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

    name := d.Get("name").(string)
    expirationDate := d.Get("expiration_date").(string)
    primaryKey := d.Get("primary_key").(string)
    productId := d.Get("product_id").(string)
    secondaryKey := d.Get("secondary_key").(string)
    state := d.Get("state").(string)
    stateComment := d.Get("state_comment").(string)
    userId := d.Get("user_id").(string)

    parameters := apimanagement.SubscriptionUpdateParameters{
        ExpirationDate: convertStringToDate(expirationDate),
        Name: utils.String(name),
        PrimaryKey: utils.String(primaryKey),
        ProductID: utils.String(productId),
        SecondaryKey: utils.String(secondaryKey),
        State: apimanagement.SubscriptionStateContract(state),
        StateComment: utils.String(stateComment),
        UserID: utils.String(userId),
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
    client := meta.(*ArmClient).subscriptionsClient
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
    d.Set("created_date", (resp.CreatedDate).String())
    d.Set("end_date", (resp.EndDate).String())
    d.Set("expiration_date", (resp.ExpirationDate).String())
    d.Set("notification_date", (resp.NotificationDate).String())
    d.Set("primary_key", resp.PrimaryKey)
    d.Set("product_id", resp.ProductID)
    d.Set("secondary_key", resp.SecondaryKey)
    d.Set("sid", sid)
    d.Set("start_date", (resp.StartDate).String())
    d.Set("state", string(resp.State))
    d.Set("state_comment", resp.StateComment)
    d.Set("user_id", resp.UserID)

    return nil
}

func resourceArmSubscriptionUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).subscriptionsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    expirationDate := d.Get("expiration_date").(string)
    primaryKey := d.Get("primary_key").(string)
    productId := d.Get("product_id").(string)
    secondaryKey := d.Get("secondary_key").(string)
    sid := d.Get("sid").(string)
    state := d.Get("state").(string)
    stateComment := d.Get("state_comment").(string)
    userId := d.Get("user_id").(string)

    parameters := apimanagement.SubscriptionUpdateParameters{
        ExpirationDate: convertStringToDate(expirationDate),
        Name: utils.String(name),
        PrimaryKey: utils.String(primaryKey),
        ProductID: utils.String(productId),
        SecondaryKey: utils.String(secondaryKey),
        State: apimanagement.SubscriptionStateContract(state),
        StateComment: utils.String(stateComment),
        UserID: utils.String(userId),
    }


    if _, err := client.Update(ctx, resourceGroup, name, sid, parameters); err != nil {
        return fmt.Errorf("Error updating Subscription %q (Sid %q / Resource Group %q): %+v", name, sid, resourceGroup, err)
    }

    return resourceArmSubscriptionRead(d, meta)
}

func resourceArmSubscriptionDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).subscriptionsClient
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
