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



func resourceArmService() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmServiceCreate,
        Read: resourceArmServiceRead,
        Update: resourceArmServiceUpdate,
        Delete: resourceArmServiceDelete,

        Importer: &schema.ResourceImporter{
            State: schema.ImportStatePassthrough,
        },


        Schema: map[string]*schema.Schema{
            "name": {
                Type: schema.TypeString,
                Computed: true,
            },

            "location": azure.SchemaLocation(),

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "search_service_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "hosting_mode": {
                Type: schema.TypeString,
                Optional: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(searchmanagementclient.default),
                    string(searchmanagementclient.highDensity),
                }, false),
                Default: string(searchmanagementclient.default),
            },

            "identity": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "type": {
                            Type: schema.TypeString,
                            Required: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(searchmanagementclient.None),
                                string(searchmanagementclient.SystemAssigned),
                            }, false),
                        },
                    },
                },
            },

            "partition_count": {
                Type: schema.TypeInt,
                Optional: true,
            },

            "replica_count": {
                Type: schema.TypeInt,
                Optional: true,
            },

            "sku": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "name": {
                            Type: schema.TypeString,
                            Optional: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(searchmanagementclient.free),
                                string(searchmanagementclient.basic),
                                string(searchmanagementclient.standard),
                                string(searchmanagementclient.standard2),
                                string(searchmanagementclient.standard3),
                                string(searchmanagementclient.storage_optimized_l1),
                                string(searchmanagementclient.storage_optimized_l2),
                            }, false),
                            Default: string(searchmanagementclient.free),
                        },
                    },
                },
            },

            "provisioning_state": {
                Type: schema.TypeString,
                Computed: true,
            },

            "status": {
                Type: schema.TypeString,
                Computed: true,
            },

            "status_details": {
                Type: schema.TypeString,
                Computed: true,
            },

            "type": {
                Type: schema.TypeString,
                Computed: true,
            },

            "tags": tags.Schema(),
        },
    }
}

func resourceArmServiceCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).servicesClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    searchServiceName := d.Get("search_service_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, searchServiceName)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Service (Search Service Name %q / Resource Group %q): %+v", searchServiceName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_service", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    hostingMode := d.Get("hosting_mode").(string)
    identity := d.Get("identity").([]interface{})
    partitionCount := d.Get("partition_count").(int)
    replicaCount := d.Get("replica_count").(int)
    sku := d.Get("sku").([]interface{})
    t := d.Get("tags").(map[string]interface{})

    service := searchmanagementclient.SearchService{
        Identity: expandArmServiceIdentity(identity),
        Location: utils.String(location),
        SearchServiceProperties: &searchmanagementclient.SearchServiceProperties{
            HostingMode: searchmanagementclient.HostingMode(hostingMode),
            PartitionCount: utils.Int32(int32(partitionCount)),
            ReplicaCount: utils.Int32(int32(replicaCount)),
        },
        Sku: expandArmServiceSku(sku),
        Tags: tags.Expand(t),
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, searchServiceName, service)
    if err != nil {
        return fmt.Errorf("Error creating Service (Search Service Name %q / Resource Group %q): %+v", searchServiceName, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Service (Search Service Name %q / Resource Group %q): %+v", searchServiceName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, searchServiceName)
    if err != nil {
        return fmt.Errorf("Error retrieving Service (Search Service Name %q / Resource Group %q): %+v", searchServiceName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Service (Search Service Name %q / Resource Group %q) ID", searchServiceName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmServiceRead(d, meta)
}

func resourceArmServiceRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).servicesClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    searchServiceName := id.Path["searchServices"]

    resp, err := client.Get(ctx, resourceGroup, searchServiceName)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Service %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Service (Search Service Name %q / Resource Group %q): %+v", searchServiceName, resourceGroup, err)
    }


    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if searchServiceProperties := resp.SearchServiceProperties; searchServiceProperties != nil {
        d.Set("hosting_mode", string(searchServiceProperties.HostingMode))
        d.Set("partition_count", int(*searchServiceProperties.PartitionCount))
        d.Set("provisioning_state", string(searchServiceProperties.ProvisioningState))
        d.Set("replica_count", int(*searchServiceProperties.ReplicaCount))
        d.Set("status", string(searchServiceProperties.Status))
        d.Set("status_details", searchServiceProperties.StatusDetails)
    }
    if err := d.Set("identity", flattenArmServiceIdentity(resp.Identity)); err != nil {
        return fmt.Errorf("Error setting `identity`: %+v", err)
    }
    d.Set("search_service_name", searchServiceName)
    if err := d.Set("sku", flattenArmServiceSku(resp.Sku)); err != nil {
        return fmt.Errorf("Error setting `sku`: %+v", err)
    }
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}

func resourceArmServiceUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).servicesClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    hostingMode := d.Get("hosting_mode").(string)
    identity := d.Get("identity").([]interface{})
    partitionCount := d.Get("partition_count").(int)
    replicaCount := d.Get("replica_count").(int)
    searchServiceName := d.Get("search_service_name").(string)
    sku := d.Get("sku").([]interface{})
    t := d.Get("tags").(map[string]interface{})

    service := searchmanagementclient.SearchService{
        Identity: expandArmServiceIdentity(identity),
        Location: utils.String(location),
        SearchServiceProperties: &searchmanagementclient.SearchServiceProperties{
            HostingMode: searchmanagementclient.HostingMode(hostingMode),
            PartitionCount: utils.Int32(int32(partitionCount)),
            ReplicaCount: utils.Int32(int32(replicaCount)),
        },
        Sku: expandArmServiceSku(sku),
        Tags: tags.Expand(t),
    }


    if _, err := client.Update(ctx, resourceGroup, searchServiceName, service); err != nil {
        return fmt.Errorf("Error updating Service (Search Service Name %q / Resource Group %q): %+v", searchServiceName, resourceGroup, err)
    }

    return resourceArmServiceRead(d, meta)
}

func resourceArmServiceDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).servicesClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    searchServiceName := id.Path["searchServices"]

    if _, err := client.Delete(ctx, resourceGroup, searchServiceName); err != nil {
        return fmt.Errorf("Error deleting Service (Search Service Name %q / Resource Group %q): %+v", searchServiceName, resourceGroup, err)
    }

    return nil
}

func expandArmServiceIdentity(input []interface{}) *searchmanagementclient.Identity {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    type := v["type"].(string)

    result := searchmanagementclient.Identity{
        Type: searchmanagementclient.IdentityType(type),
    }
    return &result
}

func expandArmServiceSku(input []interface{}) *searchmanagementclient.Sku {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    name := v["name"].(string)

    result := searchmanagementclient.Sku{
        Name: searchmanagementclient.SkuName(name),
    }
    return &result
}


func flattenArmServiceIdentity(input *searchmanagementclient.Identity) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    result["type"] = string(input.Type)

    return []interface{}{result}
}

func flattenArmServiceSku(input *searchmanagementclient.Sku) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    result := make(map[string]interface{})

    result["name"] = string(input.Name)

    return []interface{}{result}
}
