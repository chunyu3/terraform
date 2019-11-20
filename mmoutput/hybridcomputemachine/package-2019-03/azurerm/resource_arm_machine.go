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



func resourceArmMachine() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmMachineCreate,
        Read: resourceArmMachineRead,
        Update: resourceArmMachineUpdate,
        Delete: resourceArmMachineDelete,

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

            "location": azure.SchemaLocation(),

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "client_public_key": {
                Type: schema.TypeString,
                Optional: true,
            },

            "identity": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "type": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },

            "vm_id": {
                Type: schema.TypeString,
                Optional: true,
            },

            "type": {
                Type: schema.TypeString,
                Computed: true,
            },

            "tags": tags.Schema(),
        },
    }
}

func resourceArmMachineCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).machinesClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Machine %q (Resource Group %q): %+v", name, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_machine", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    clientPublicKey := d.Get("client_public_key").(string)
    identity := d.Get("identity").([]interface{})
    vmId := d.Get("vm_id").(string)
    t := d.Get("tags").(map[string]interface{})

    parameters := hybridcompute.MachineReconnect{
        Identity: expandArmMachineIdentity(identity),
        Location: utils.String(location),
        MachineReconnectProperties: &hybridcompute.MachineReconnectProperties{
            ClientPublicKey: utils.String(clientPublicKey),
            VMID: utils.String(vmId),
        },
        Tags: tags.Expand(t),
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroup, name, parameters); err != nil {
        return fmt.Errorf("Error creating Machine %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Machine %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Machine %q (Resource Group %q) ID", name, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmMachineRead(d, meta)
}

func resourceArmMachineRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).machinesClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["machines"]

    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Machine %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Machine %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("resource_group", resourceGroup)
    d.Set("type", resp.Type)

    return nil
}

func resourceArmMachineUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).machinesClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    clientPublicKey := d.Get("client_public_key").(string)
    identity := d.Get("identity").([]interface{})
    vmId := d.Get("vm_id").(string)
    t := d.Get("tags").(map[string]interface{})

    parameters := hybridcompute.MachineReconnect{
        Identity: expandArmMachineIdentity(identity),
        MachineReconnectProperties: &hybridcompute.MachineReconnectProperties{
            ClientPublicKey: utils.String(clientPublicKey),
            VMID: utils.String(vmId),
        },
        Tags: tags.Expand(t),
    }


    if _, err := client.Update(ctx, resourceGroup, name, parameters); err != nil {
        return fmt.Errorf("Error updating Machine %q (Resource Group %q): %+v", name, resourceGroup, err)
    }

    return resourceArmMachineRead(d, meta)
}

func resourceArmMachineDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).machinesClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["machines"]

    if _, err := client.Delete(ctx, resourceGroup, name); err != nil {
        return fmt.Errorf("Error deleting Machine %q (Resource Group %q): %+v", name, resourceGroup, err)
    }

    return nil
}

func expandArmMachineIdentity(input []interface{}) *hybridcompute.Identity {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    type := v["type"].(string)

    result := hybridcompute.Identity{
        Type: utils.String(type),
    }
    return &result
}
