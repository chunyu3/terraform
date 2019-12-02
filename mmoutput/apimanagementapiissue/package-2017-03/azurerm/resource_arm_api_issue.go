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



func resourceArmApiIssue() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmApiIssueCreateUpdate,
        Read: resourceArmApiIssueRead,
        Update: resourceArmApiIssueCreateUpdate,
        Delete: resourceArmApiIssueDelete,

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

            "api_id": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "description": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "issue_id": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "title": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "user_id": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "api_id": {
                Type: schema.TypeString,
                Optional: true,
            },

            "created_date": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validateRFC3339Date,
            },

            "state": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(apimanagement.proposed),
                    string(apimanagement.open),
                    string(apimanagement.removed),
                    string(apimanagement.resolved),
                    string(apimanagement.closed),
                }, false),
                Default: string(apimanagement.proposed),
            },

            "type": {
                Type: schema.TypeString,
                Computed: true,
            },
        },
    }
}

func resourceArmApiIssueCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).apiIssueClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    aPIID := d.Get("api_id").(string)
    issueID := d.Get("issue_id").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, name, aPIID, issueID)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Api Issue %q (Issue %q / Api %q / Resource Group %q): %+v", name, issueID, aPIID, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_api_issue", *existing.ID)
        }
    }

    apiId := d.Get("api_id").(string)
    createdDate := d.Get("created_date").(string)
    description := d.Get("description").(string)
    state := d.Get("state").(string)
    title := d.Get("title").(string)
    userId := d.Get("user_id").(string)

    parameters := apimanagement.IssueContract{
        IssueContractProperties: &apimanagement.IssueContractProperties{
            APIID: utils.String(apiId),
            CreatedDate: convertStringToDate(createdDate),
            Description: utils.String(description),
            State: apimanagement.State(state),
            Title: utils.String(title),
            UserID: utils.String(userId),
        },
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroup, name, aPIID, issueID, parameters); err != nil {
        return fmt.Errorf("Error creating Api Issue %q (Issue %q / Api %q / Resource Group %q): %+v", name, issueID, aPIID, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, name, aPIID, issueID)
    if err != nil {
        return fmt.Errorf("Error retrieving Api Issue %q (Issue %q / Api %q / Resource Group %q): %+v", name, issueID, aPIID, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Api Issue %q (Issue %q / Api %q / Resource Group %q) ID", name, issueID, aPIID, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmApiIssueRead(d, meta)
}

func resourceArmApiIssueRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).apiIssueClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["service"]
    aPIID := id.Path["apis"]
    issueID := id.Path["issues"]

    resp, err := client.Get(ctx, resourceGroup, name, aPIID, issueID)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Api Issue %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Api Issue %q (Issue %q / Api %q / Resource Group %q): %+v", name, issueID, aPIID, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    d.Set("api_id", aPIID)
    if issueContractProperties := resp.IssueContractProperties; issueContractProperties != nil {
        d.Set("api_id", issueContractProperties.APIID)
        d.Set("created_date", (issueContractProperties.CreatedDate).String())
        d.Set("description", issueContractProperties.Description)
        d.Set("state", string(issueContractProperties.State))
        d.Set("title", issueContractProperties.Title)
        d.Set("user_id", issueContractProperties.UserID)
    }
    d.Set("issue_id", issueID)
    d.Set("type", resp.Type)

    return nil
}


func resourceArmApiIssueDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).apiIssueClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["service"]
    aPIID := id.Path["apis"]
    issueID := id.Path["issues"]

    if _, err := client.Delete(ctx, resourceGroup, name, aPIID, issueID); err != nil {
        return fmt.Errorf("Error deleting Api Issue %q (Issue %q / Api %q / Resource Group %q): %+v", name, issueID, aPIID, resourceGroup, err)
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
