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



func resourceArmTag() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmTagCreate,
        Read: resourceArmTagRead,
        Update: resourceArmTagUpdate,
        Delete: resourceArmTagDelete,

        Importer: &schema.ResourceImporter{
            State: schema.ImportStatePassthrough,
        },


        Schema: map[string]*schema.Schema{
            "name": {
                Type: schema.TypeString,
                Computed: true,
            },

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "display_name": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "service_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "tag_id": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "type": {
                Type: schema.TypeString,
                Computed: true,
            },
        },
    }
}

func resourceArmTagCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).tagClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    serviceName := d.Get("service_name").(string)
    tagID := d.Get("tag_id").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, serviceName, tagID)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Tag (Tag %q / Service Name %q / Resource Group %q): %+v", tagID, serviceName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_tag", *existing.ID)
        }
    }

    displayName := d.Get("display_name").(string)

    parameters := apimanagement.TagCreateUpdateParameters{
        TagContractProperties: &apimanagement.TagContractProperties{
            DisplayName: utils.String(displayName),
        },
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroup, serviceName, tagID, parameters); err != nil {
        return fmt.Errorf("Error creating Tag (Tag %q / Service Name %q / Resource Group %q): %+v", tagID, serviceName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, serviceName, tagID)
    if err != nil {
        return fmt.Errorf("Error retrieving Tag (Tag %q / Service Name %q / Resource Group %q): %+v", tagID, serviceName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Tag (Tag %q / Service Name %q / Resource Group %q) ID", tagID, serviceName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmTagRead(d, meta)
}

func resourceArmTagRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).tagClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    serviceName := id.Path["service"]
    tagID := id.Path["tags"]

    resp, err := client.Get(ctx, resourceGroup, serviceName, tagID)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Tag %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Tag (Tag %q / Service Name %q / Resource Group %q): %+v", tagID, serviceName, resourceGroup, err)
    }


    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if tagContractProperties := resp.TagContractProperties; tagContractProperties != nil {
        d.Set("display_name", tagContractProperties.DisplayName)
    }
    d.Set("service_name", serviceName)
    d.Set("tag_id", tagID)
    d.Set("type", resp.Type)

    return nil
}

func resourceArmTagUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).tagClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    displayName := d.Get("display_name").(string)
    serviceName := d.Get("service_name").(string)
    tagID := d.Get("tag_id").(string)

    parameters := apimanagement.TagCreateUpdateParameters{
        TagContractProperties: &apimanagement.TagContractProperties{
            DisplayName: utils.String(displayName),
        },
    }


    if _, err := client.Update(ctx, resourceGroup, serviceName, tagID, parameters); err != nil {
        return fmt.Errorf("Error updating Tag (Tag %q / Service Name %q / Resource Group %q): %+v", tagID, serviceName, resourceGroup, err)
    }

    return resourceArmTagRead(d, meta)
}

func resourceArmTagDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).tagClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    serviceName := id.Path["service"]
    tagID := id.Path["tags"]

    if _, err := client.Delete(ctx, resourceGroup, serviceName, tagID); err != nil {
        return fmt.Errorf("Error deleting Tag (Tag %q / Service Name %q / Resource Group %q): %+v", tagID, serviceName, resourceGroup, err)
    }

    return nil
}