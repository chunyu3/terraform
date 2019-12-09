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



func resourceArmZone() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmZoneCreateUpdate,
        Read: resourceArmZoneRead,
        Update: resourceArmZoneCreateUpdate,
        Delete: resourceArmZoneDelete,

        Importer: &schema.ResourceImporter{
            State: schema.ImportStatePassthrough,
        },


        Schema: map[string]*schema.Schema{
            "location": azure.SchemaLocation(),

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "zone_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "etag": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
            },

            "max_number_of_record_sets": {
                Type: schema.TypeInt,
                Optional: true,
            },

            "number_of_record_sets": {
                Type: schema.TypeInt,
                Optional: true,
            },

            "tags": tags.Schema(),

            "id": {
                Type: schema.TypeString,
                Computed: true,
            },

            "name": {
                Type: schema.TypeString,
                Computed: true,
            },

            "name_servers": {
                Type: schema.TypeList,
                Computed: true,
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
            },

            "type": {
                Type: schema.TypeString,
                Computed: true,
            },
        },
    }
}

func resourceArmZoneCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).zonesClient
    ctx, cancel := timeouts.ForCreateUpdate(meta.(*ArmClient).StopContext, d)
    defer cancel()

    resourceGroupName := d.Get("resource_group").(string)
    name := d.Get("zone_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroupName, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Zone (Zone Name %q / Resource Group %q): %+v", name, resourceGroupName, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_zone", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    etag := d.Get("etag").(string)
    maxNumberOfRecordSets := d.Get("max_number_of_record_sets").(int)
    numberOfRecordSets := d.Get("number_of_record_sets").(int)
    tags := d.Get("tags").(map[string]interface{})

    parameters := dns.Zone{
        Etag: utils.String(etag),
        Location: utils.String(location),
        ZoneProperties: &dns.ZoneProperties{
            MaxNumberOfRecordSets: utils.Int64(int64(maxNumberOfRecordSets)),
            NumberOfRecordSets: utils.Int64(int64(numberOfRecordSets)),
        },
        Tags: tags.Expand(tags),
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroupName, name, parameters); err != nil {
        return fmt.Errorf("Error creating Zone (Zone Name %q / Resource Group %q): %+v", name, resourceGroupName, err)
    }


    resp, err := client.Get(ctx, resourceGroupName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Zone (Zone Name %q / Resource Group %q): %+v", name, resourceGroupName, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Zone (Zone Name %q / Resource Group %q) ID", name, resourceGroupName)
    }
    d.SetId(*resp.ID)

    return resourceArmZoneRead(d, meta)
}

func resourceArmZoneRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).zonesClient
    ctx, cancel := timeouts.ForRead(meta.(*ArmClient).StopContext, d)
    defer cancel()

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroupName := id.ResourceGroup
    name := id.Path["dnsZones"]

    resp, err := client.Get(ctx, resourceGroupName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Zone %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Zone (Zone Name %q / Resource Group %q): %+v", name, resourceGroupName, err)
    }


    d.Set("resource_group", resourceGroupName)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    d.Set("etag", resp.Etag)
    d.Set("id", resp.ID)
    if zoneProperties := resp.ZoneProperties; zoneProperties != nil {
        d.Set("max_number_of_record_sets", int(*zoneProperties.MaxNumberOfRecordSets))
        d.Set("name_servers", utils.FlattenStringSlice(zoneProperties.NameServers))
        d.Set("number_of_record_sets", int(*zoneProperties.NumberOfRecordSets))
    }
    d.Set("name", resp.Name)
    d.Set("type", resp.Type)
    d.Set("zone_name", name)

    return tags.FlattenAndSet(d, resp.Tags)
}


func resourceArmZoneDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).zonesClient
    ctx, cancel := timeouts.ForDelete(meta.(*ArmClient).StopContext, d)
    defer cancel()


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroupName := id.ResourceGroup
    name := id.Path["dnsZones"]

    future, err := client.Delete(ctx, resourceGroupName, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Zone (Zone Name %q / Resource Group %q): %+v", name, resourceGroupName, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Zone (Zone Name %q / Resource Group %q): %+v", name, resourceGroupName, err)
        }
    }

    return nil
}
