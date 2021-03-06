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



func resourceArmOrder() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmOrderCreateUpdate,
        Read: resourceArmOrderRead,
        Update: resourceArmOrderCreateUpdate,
        Delete: resourceArmOrderDelete,

        Importer: &schema.ResourceImporter{
            State: schema.ImportStatePassthrough,
        },


        Schema: map[string]*schema.Schema{
            "contact_information": {
                Type: schema.TypeList,
                Required: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "company_name": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "contact_person": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "email_list": {
                            Type: schema.TypeList,
                            Required: true,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                        "phone": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                    },
                },
            },

            "device_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "shipping_address": {
                Type: schema.TypeList,
                Required: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "address_line1": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "city": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "country": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "postal_code": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "state": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validate.NoEmptyStrings,
                        },
                        "address_line2": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "address_line3": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },

            "current_status": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "status": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(databoxedge.Untracked),
                                string(databoxedge.AwaitingFulfilment),
                                string(databoxedge.AwaitingPreparation),
                                string(databoxedge.AwaitingShipment),
                                string(databoxedge.Shipped),
                                string(databoxedge.Arriving),
                                string(databoxedge.Delivered),
                                string(databoxedge.ReplacementRequested),
                                string(databoxedge.LostDevice),
                                string(databoxedge.Declined),
                                string(databoxedge.ReturnInitiated),
                                string(databoxedge.AwaitingReturnShipment),
                                string(databoxedge.ShippedBack),
                                string(databoxedge.CollectedAtMicrosoft),
                            }, false),
                        },
                        "comments": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },

            "delivery_tracking_info": {
                Type: schema.TypeList,
                Computed: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "carrier_name": {
                            Type: schema.TypeString,
                            Computed: true,
                        },
                        "serial_number": {
                            Type: schema.TypeString,
                            Computed: true,
                        },
                        "tracking_id": {
                            Type: schema.TypeString,
                            Computed: true,
                        },
                        "tracking_url": {
                            Type: schema.TypeString,
                            Computed: true,
                        },
                    },
                },
            },

            "id": {
                Type: schema.TypeString,
                Computed: true,
            },

            "name": {
                Type: schema.TypeString,
                Computed: true,
            },

            "order_history": {
                Type: schema.TypeList,
                Computed: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "additional_order_details": {
                            Type: schema.TypeMap,
                            Computed: true,
                            Elem: &schema.Schema{Type: schema.TypeString},
                        },
                        "comments": {
                            Type: schema.TypeString,
                            Computed: true,
                        },
                        "status": {
                            Type: schema.TypeString,
                            Computed: true,
                        },
                        "update_date_time": {
                            Type: schema.TypeString,
                            Computed: true,
                        },
                    },
                },
            },

            "return_tracking_info": {
                Type: schema.TypeList,
                Computed: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "carrier_name": {
                            Type: schema.TypeString,
                            Computed: true,
                        },
                        "serial_number": {
                            Type: schema.TypeString,
                            Computed: true,
                        },
                        "tracking_id": {
                            Type: schema.TypeString,
                            Computed: true,
                        },
                        "tracking_url": {
                            Type: schema.TypeString,
                            Computed: true,
                        },
                    },
                },
            },

            "serial_number": {
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

func resourceArmOrderCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).ordersClient
    ctx, cancel := timeouts.ForCreateUpdate(meta.(*ArmClient).StopContext, d)
    defer cancel()

    resourceGroupName := d.Get("resource_group").(string)
    name := d.Get("device_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroupName, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Order (Resource Group %q / Device Name %q): %+v", resourceGroupName, name, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_order", *existing.ID)
        }
    }

    contactInformation := d.Get("contact_information").([]interface{})
    currentStatus := d.Get("current_status").([]interface{})
    shippingAddress := d.Get("shipping_address").([]interface{})

    order := databoxedge.Order{
        OrderProperties: &databoxedge.OrderProperties{
            ContactInformation: expandArmOrderContactDetails(contactInformation),
            CurrentStatus: expandArmOrderOrderStatus(currentStatus),
            ShippingAddress: expandArmOrderAddress(shippingAddress),
        },
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroupName, name, order)
    if err != nil {
        return fmt.Errorf("Error creating Order (Resource Group %q / Device Name %q): %+v", resourceGroupName, name, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Order (Resource Group %q / Device Name %q): %+v", resourceGroupName, name, err)
    }


    resp, err := client.Get(ctx, resourceGroupName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Order (Resource Group %q / Device Name %q): %+v", resourceGroupName, name, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Order (Resource Group %q / Device Name %q) ID", resourceGroupName, name)
    }
    d.SetId(*resp.ID)

    return resourceArmOrderRead(d, meta)
}

func resourceArmOrderRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).ordersClient
    ctx, cancel := timeouts.ForRead(meta.(*ArmClient).StopContext, d)
    defer cancel()

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroupName := id.ResourceGroup
    name := id.Path["dataBoxEdgeDevices"]

    resp, err := client.Get(ctx, resourceGroupName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Order %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Order (Resource Group %q / Device Name %q): %+v", resourceGroupName, name, err)
    }


    d.Set("resource_group", resourceGroupName)
    if orderProperties := resp.OrderProperties; orderProperties != nil {
        if err := d.Set("contact_information", flattenArmOrderContactDetails(orderProperties.ContactInformation)); err != nil {
            return fmt.Errorf("Error setting `contact_information`: %+v", err)
        }
        if err := d.Set("current_status", flattenArmOrderOrderStatus(orderProperties.CurrentStatus)); err != nil {
            return fmt.Errorf("Error setting `current_status`: %+v", err)
        }
        if err := d.Set("delivery_tracking_info", flattenArmOrderTrackingInfo(orderProperties.DeliveryTrackingInfo)); err != nil {
            return fmt.Errorf("Error setting `delivery_tracking_info`: %+v", err)
        }
        if err := d.Set("order_history", flattenArmOrderOrderStatus(orderProperties.OrderHistory)); err != nil {
            return fmt.Errorf("Error setting `order_history`: %+v", err)
        }
        if err := d.Set("return_tracking_info", flattenArmOrderTrackingInfo(orderProperties.ReturnTrackingInfo)); err != nil {
            return fmt.Errorf("Error setting `return_tracking_info`: %+v", err)
        }
        d.Set("serial_number", orderProperties.SerialNumber)
        if err := d.Set("shipping_address", flattenArmOrderAddress(orderProperties.ShippingAddress)); err != nil {
            return fmt.Errorf("Error setting `shipping_address`: %+v", err)
        }
    }
    d.Set("device_name", name)
    d.Set("id", resp.ID)
    d.Set("name", resp.Name)
    d.Set("type", resp.Type)

    return nil
}


func resourceArmOrderDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).ordersClient
    ctx, cancel := timeouts.ForDelete(meta.(*ArmClient).StopContext, d)
    defer cancel()


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroupName := id.ResourceGroup
    name := id.Path["dataBoxEdgeDevices"]

    future, err := client.Delete(ctx, resourceGroupName, name)
    if err != nil {
        if response.WasNotFound(future.Response()) {
            return nil
        }
        return fmt.Errorf("Error deleting Order (Resource Group %q / Device Name %q): %+v", resourceGroupName, name, err)
    }

    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        if !response.WasNotFound(future.Response()) {
            return fmt.Errorf("Error waiting for deleting Order (Resource Group %q / Device Name %q): %+v", resourceGroupName, name, err)
        }
    }

    return nil
}

func expandArmOrderContactDetails(input []interface{}) *databoxedge.ContactDetails {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    contactPerson := v["contact_person"].(string)
    companyName := v["company_name"].(string)
    phone := v["phone"].(string)
    emailList := v["email_list"].([]interface{})

    result := databoxedge.ContactDetails{
        CompanyName: utils.String(companyName),
        ContactPerson: utils.String(contactPerson),
        EmailList: utils.ExpandStringSlice(emailList),
        Phone: utils.String(phone),
    }
    return &result
}

func expandArmOrderOrderStatus(input []interface{}) *databoxedge.OrderStatus {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    status := v["status"].(string)
    comments := v["comments"].(string)

    result := databoxedge.OrderStatus{
        Comments: utils.String(comments),
        Status: databoxedge.OrderState(status),
    }
    return &result
}

func expandArmOrderAddress(input []interface{}) *databoxedge.Address {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    addressLine1 := v["address_line1"].(string)
    addressLine2 := v["address_line2"].(string)
    addressLine3 := v["address_line3"].(string)
    postalCode := v["postal_code"].(string)
    city := v["city"].(string)
    state := v["state"].(string)
    country := v["country"].(string)

    result := databoxedge.Address{
        AddressLine1: utils.String(addressLine1),
        AddressLine2: utils.String(addressLine2),
        AddressLine3: utils.String(addressLine3),
        City: utils.String(city),
        Country: utils.String(country),
        PostalCode: utils.String(postalCode),
        State: utils.String(state),
    }
    return &result
}


func flattenArmOrderContactDetails(input *databoxedge.ContactDetails) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if companyName := input.CompanyName; companyName != nil {
        result["company_name"] = *companyName
    }
    if contactPerson := input.ContactPerson; contactPerson != nil {
        result["contact_person"] = *contactPerson
    }
    result["email_list"] = utils.FlattenStringSlice(input.EmailList)
    if phone := input.Phone; phone != nil {
        result["phone"] = *phone
    }

    return []interface{}{result}
}

func flattenArmOrderOrderStatus(input *databoxedge.OrderStatus) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if comments := input.Comments; comments != nil {
        result["comments"] = *comments
    }
    result["status"] = string(input.Status)

    return []interface{}{result}
}

func flattenArmOrderTrackingInfo(input *[]databoxedge.TrackingInfo) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        if carrierName := item.CarrierName; carrierName != nil {
            v["carrier_name"] = *carrierName
        }
        if serialNumber := item.SerialNumber; serialNumber != nil {
            v["serial_number"] = *serialNumber
        }
        if trackingId := item.TrackingID; trackingId != nil {
            v["tracking_id"] = *trackingId
        }
        if trackingUrl := item.TrackingURL; trackingUrl != nil {
            v["tracking_url"] = *trackingUrl
        }

        results = append(results, v)
    }

    return results
}

func flattenArmOrderOrderStatus(input *[]databoxedge.OrderStatus) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})

        v["additional_order_details"] = utils.FlattenKeyValuePairs(item.AdditionalOrderDetails)
        if comments := item.Comments; comments != nil {
            v["comments"] = *comments
        }
        v["status"] = string(item.Status)
        if updateDateTime := item.UpdateDateTime; updateDateTime != nil {
            v["update_date_time"] = (*updateDateTime).String()
        }

        results = append(results, v)
    }

    return results
}

func flattenArmOrderAddress(input *databoxedge.Address) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    if addressLine1 := input.AddressLine1; addressLine1 != nil {
        result["address_line1"] = *addressLine1
    }
    if addressLine2 := input.AddressLine2; addressLine2 != nil {
        result["address_line2"] = *addressLine2
    }
    if addressLine3 := input.AddressLine3; addressLine3 != nil {
        result["address_line3"] = *addressLine3
    }
    if city := input.City; city != nil {
        result["city"] = *city
    }
    if country := input.Country; country != nil {
        result["country"] = *country
    }
    if postalCode := input.PostalCode; postalCode != nil {
        result["postal_code"] = *postalCode
    }
    if state := input.State; state != nil {
        result["state"] = *state
    }

    return []interface{}{result}
}
