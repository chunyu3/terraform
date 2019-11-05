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



func resourceArmServiceTopology() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmServiceTopologyCreateUpdate,
        Read: resourceArmServiceTopologyRead,
        Update: resourceArmServiceTopologyCreateUpdate,
        Delete: resourceArmServiceTopologyDelete,

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

            "service_topology_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "artifact_source_id": {
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

func resourceArmServiceTopologyCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).serviceTopologiesClient
    ctx := meta.(*ArmClient).StopContext

    resourceGroup := d.Get("resource_group").(string)
    serviceTopologyName := d.Get("service_topology_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, serviceTopologyName)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Service Topology (Service Topology Name %q / Resource Group %q): %+v", serviceTopologyName, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_service_topology", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    artifactSourceId := d.Get("artifact_source_id").(string)
    t := d.Get("tags").(map[string]interface{})

    serviceTopologyInfo := azuredeploymentmanager.ServiceTopologyResource{
        Location: utils.String(location),
        ServiceTopologyResource_properties: &azuredeploymentmanager.ServiceTopologyResource_properties{
            ArtifactSourceID: utils.String(artifactSourceId),
        },
        Tags: tags.Expand(t),
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroup, serviceTopologyName, serviceTopologyInfo); err != nil {
        return fmt.Errorf("Error creating Service Topology (Service Topology Name %q / Resource Group %q): %+v", serviceTopologyName, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, serviceTopologyName)
    if err != nil {
        return fmt.Errorf("Error retrieving Service Topology (Service Topology Name %q / Resource Group %q): %+v", serviceTopologyName, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Service Topology (Service Topology Name %q / Resource Group %q) ID", serviceTopologyName, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmServiceTopologyRead(d, meta)
}

func resourceArmServiceTopologyRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).serviceTopologiesClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    serviceTopologyName := id.Path["serviceTopologies"]

    resp, err := client.Get(ctx, resourceGroup, serviceTopologyName)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Service Topology %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Service Topology (Service Topology Name %q / Resource Group %q): %+v", serviceTopologyName, resourceGroup, err)
    }


    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    if serviceTopologyResourceProperties := resp.ServiceTopologyResource_properties; serviceTopologyResourceProperties != nil {
        d.Set("artifact_source_id", serviceTopologyResourceProperties.ArtifactSourceID)
    }
    d.Set("service_topology_name", serviceTopologyName)
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}


func resourceArmServiceTopologyDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).serviceTopologiesClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    serviceTopologyName := id.Path["serviceTopologies"]

    if _, err := client.Delete(ctx, resourceGroup, serviceTopologyName); err != nil {
        return fmt.Errorf("Error deleting Service Topology (Service Topology Name %q / Resource Group %q): %+v", serviceTopologyName, resourceGroup, err)
    }

    return nil
}
