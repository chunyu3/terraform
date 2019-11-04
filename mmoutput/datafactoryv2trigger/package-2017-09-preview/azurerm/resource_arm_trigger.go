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



func resourceArmTrigger() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmTriggerCreateUpdate,
        Read: resourceArmTriggerRead,
        Update: resourceArmTriggerCreateUpdate,
        Delete: resourceArmTriggerDelete,

        Importer: &schema.ResourceImporter{
            State: schema.ImportStatePassthrough,
        },


        Schema: map[string]*schema.Schema{
            "name": {
                Type: schema.TypeString,
                Computed: true,
            },

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "factory_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "trigger_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "additional_properties": {
                Type: schema.TypeMap,
                Optional: true,
                Elem: &schema.Schema{Type: schema.TypeString},
            },

            "description": {
                Type: schema.TypeString,
                Optional: true,
            },

            "etag": {
                Type: schema.TypeString,
                Computed: true,
            },

            "runtime_state": {
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

func resourceArmTriggerCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).triggersClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    factoryName := d.Get("factory_name").(string)
    triggerName := d.Get("trigger_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, factoryName, triggerName)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Trigger (Trigger Name %q / Factory Name %q / Resource Group %q): %+v", triggerName, factoryName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_trigger", *existing.ID)
        }
    }

    additionalProperties := d.Get("additional_properties").(map[string]interface{})
    description := d.Get("description").(string)

    trigger := datafactoryv2.TriggerResource{
        Trigger: &datafactoryv2.Trigger{
            AdditionalProperties: utils.ExpandKeyValuePairs(additionalProperties),
            Description: utils.String(description),
        },
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroup, factoryName, triggerName, trigger); err != nil {
        return fmt.Errorf("Error creating Trigger (Trigger Name %q / Factory Name %q / Resource Group %q): %+v", triggerName, factoryName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, factoryName, triggerName)
    if err != nil {
        return fmt.Errorf("Error retrieving Trigger (Trigger Name %q / Factory Name %q / Resource Group %q): %+v", triggerName, factoryName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Trigger (Trigger Name %q / Factory Name %q / Resource Group %q) ID", triggerName, factoryName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmTriggerRead(d, meta)
}

func resourceArmTriggerRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).triggersClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    factoryName := id.Path["factories"]
    triggerName := id.Path["triggers"]

    resp, err := client.Get(ctx, resourceGroup, factoryName, triggerName)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Trigger %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Trigger (Trigger Name %q / Factory Name %q / Resource Group %q): %+v", triggerName, factoryName, resourceGroup, err)
    }


    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if trigger := resp.Trigger; trigger != nil {
        d.Set("additional_properties", utils.FlattenKeyValuePairs(trigger.AdditionalProperties))
        d.Set("description", trigger.Description)
        d.Set("runtime_state", string(trigger.RuntimeState))
    }
    d.Set("etag", resp.Etag)
    d.Set("factory_name", factoryName)
    d.Set("trigger_name", triggerName)
    d.Set("type", resp.Type)

    return nil
}


func resourceArmTriggerDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).triggersClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    factoryName := id.Path["factories"]
    triggerName := id.Path["triggers"]

    if _, err := client.Delete(ctx, resourceGroup, factoryName, triggerName); err != nil {
        return fmt.Errorf("Error deleting Trigger (Trigger Name %q / Factory Name %q / Resource Group %q): %+v", triggerName, factoryName, resourceGroup, err)
    }

    return nil
}