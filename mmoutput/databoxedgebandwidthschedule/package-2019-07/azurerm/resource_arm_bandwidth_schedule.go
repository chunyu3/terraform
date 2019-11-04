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



func resourceArmBandwidthSchedule() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmBandwidthScheduleCreateUpdate,
        Read: resourceArmBandwidthScheduleRead,
        Update: resourceArmBandwidthScheduleCreateUpdate,
        Delete: resourceArmBandwidthScheduleDelete,

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

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "days": {
                Type: schema.TypeList,
                Required: true,
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                    ValidateFunc: validation.StringInSlice([]string{
                        string(databoxedge.Sunday),
                        string(databoxedge.Monday),
                        string(databoxedge.Tuesday),
                        string(databoxedge.Wednesday),
                        string(databoxedge.Thursday),
                        string(databoxedge.Friday),
                        string(databoxedge.Saturday),
                   }, false),
                },
            },

            "device_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "rate_in_mbps": {
                Type: schema.TypeInt,
                Required: true,
            },

            "start": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "stop": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "type": {
                Type: schema.TypeString,
                Computed: true,
            },
        },
    }
}

func resourceArmBandwidthScheduleCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).bandwidthSchedulesClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    deviceName := d.Get("device_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, deviceName, name, resourceGroup)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Bandwidth Schedule %q (Resource Group %q / Device Name %q): %+v", name, resourceGroup, deviceName, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_bandwidth_schedule", *existing.ID)
        }
    }

    days := d.Get("days").([]interface{})
    rateInMbps := d.Get("rate_in_mbps").(int)
    start := d.Get("start").(string)
    stop := d.Get("stop").(string)

    parameters := databoxedge.BandwidthSchedule{
        BandwidthScheduleProperties: &databoxedge.BandwidthScheduleProperties{
            Days: expandArmBandwidthSchedule(days),
            RateInMbps: utils.Int32(int32(rateInMbps)),
            Start: utils.String(start),
            Stop: utils.String(stop),
        },
    }


    future, err := client.CreateOrUpdate(ctx, deviceName, name, resourceGroup, parameters)
    if err != nil {
        return fmt.Errorf("Error creating Bandwidth Schedule %q (Resource Group %q / Device Name %q): %+v", name, resourceGroup, deviceName, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Bandwidth Schedule %q (Resource Group %q / Device Name %q): %+v", name, resourceGroup, deviceName, err)
    }


    resp, err := client.Get(ctx, deviceName, name, resourceGroup)
    if err != nil {
        return fmt.Errorf("Error retrieving Bandwidth Schedule %q (Resource Group %q / Device Name %q): %+v", name, resourceGroup, deviceName, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Bandwidth Schedule %q (Resource Group %q / Device Name %q) ID", name, resourceGroup, deviceName)
    }
    d.SetId(*resp.ID)

    return resourceArmBandwidthScheduleRead(d, meta)
}

func resourceArmBandwidthScheduleRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).bandwidthSchedulesClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    deviceName := id.Path["dataBoxEdgeDevices"]
    name := id.Path["bandwidthSchedules"]
    resourceGroup := id.ResourceGroup

    resp, err := client.Get(ctx, deviceName, name, resourceGroup)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Bandwidth Schedule %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Bandwidth Schedule %q (Resource Group %q / Device Name %q): %+v", name, resourceGroup, deviceName, err)
    }


    d.Set("name", name)
    d.Set("resource_group", resourceGroup)
    if bandwidthScheduleProperties := resp.BandwidthScheduleProperties; bandwidthScheduleProperties != nil {
        if err := d.Set("days", flattenArmBandwidthSchedule(string(bandwidthScheduleProperties.Days))); err != nil {
            return fmt.Errorf("Error setting `days`: %+v", err)
        }
        d.Set("rate_in_mbps", int(*bandwidthScheduleProperties.RateInMbps))
        d.Set("start", bandwidthScheduleProperties.Start)
        d.Set("stop", bandwidthScheduleProperties.Stop)
    }
    d.Set("device_name", deviceName)
    d.Set("type", resp.Type)

    return nil
}


func resourceArmBandwidthScheduleDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).bandwidthSchedulesClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    deviceName := id.Path["dataBoxEdgeDevices"]
    name := id.Path["bandwidthSchedules"]
    resourceGroup := id.ResourceGroup

    future, err := client.Delete(ctx, deviceName, name, resourceGroup)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Bandwidth Schedule %q (Resource Group %q / Device Name %q): %+v", name, resourceGroup, deviceName, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Bandwidth Schedule %q (Resource Group %q / Device Name %q): %+v", name, resourceGroup, deviceName, err)
        }
    }

    return nil
}

func expandArmBandwidthSchedule(input []interface{}) *[]databoxedge. {
    results := make([]databoxedge., 0)
    for _, item := range input {
        v := item.(string)
        result := databoxedge.(v)
        results = append(results, result)
    }
    return &results
}


func flattenArmBandwidthSchedule(input *[]databoxedge.) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        result := string(item)
        results = append(results, result)
    }

    return results
}