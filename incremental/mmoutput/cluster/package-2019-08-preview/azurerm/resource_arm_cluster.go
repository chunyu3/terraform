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



func resourceArmCluster() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmClusterCreate,
        Read: resourceArmClusterRead,
        Update: resourceArmClusterUpdate,
        Delete: resourceArmClusterDelete,

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

            "location": azure.SchemaLocation(),

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "encryption_key_uri": {
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
                            Required: true,
                            ValidateFunc: validation.StringInSlice([]string{
                                string(operationalinsights.SystemAssigned),
                                string(operationalinsights.None),
                            }, false),
                        },
                    },
                },
            },

            "type": {
                Type: schema.TypeString,
                Computed: true,
            },

            "tags": tags.Schema(),
        },
    }
}

func resourceArmClusterCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).clustersClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Cluster %q (Resource Group %q): %+v", name, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_cluster", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    encryptionKeyUri := d.Get("encryption_key_uri").(string)
    identity := d.Get("identity").([]interface{})
    t := d.Get("tags").(map[string]interface{})

    parameters := operationalinsights.ClusterPatch{
        Identity: expandArmClusterIdentity(identity),
        Location: utils.String(location),
        ClusterPatchProperties: &operationalinsights.ClusterPatchProperties{
            EncryptionKeyURI: utils.String(encryptionKeyUri),
        },
        Tags: tags.Expand(t),
    }


    future, err := client.CreateOrUpdate(ctx, resourceGroup, name, parameters)
    if err != nil {
        return fmt.Errorf("Error creating Cluster %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
        return fmt.Errorf("Error waiting for creation of Cluster %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Cluster %q (Resource Group %q): %+v", name, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Cluster %q (Resource Group %q) ID", name, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmClusterRead(d, meta)
}

func resourceArmClusterRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).clustersClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.Path["resourcegroups"]
    name := id.Path["clusters"]

    resp, err := client.Get(ctx, resourceGroup, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Cluster %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Cluster %q (Resource Group %q): %+v", name, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    d.Set("type", resp.Type)

    return nil
}

func resourceArmClusterUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).clustersClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    encryptionKeyUri := d.Get("encryption_key_uri").(string)
    identity := d.Get("identity").([]interface{})
    t := d.Get("tags").(map[string]interface{})

    parameters := operationalinsights.ClusterPatch{
        Identity: expandArmClusterIdentity(identity),
        ClusterPatchProperties: &operationalinsights.ClusterPatchProperties{
            EncryptionKeyURI: utils.String(encryptionKeyUri),
        },
        Tags: tags.Expand(t),
    }


    if _, err := client.Update(ctx, resourceGroup, name, parameters); err != nil {
        return fmt.Errorf("Error updating Cluster %q (Resource Group %q): %+v", name, resourceGroup, err)
    }

    return resourceArmClusterRead(d, meta)
}

func resourceArmClusterDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).clustersClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.Path["resourcegroups"]
    name := id.Path["clusters"]

    if _, err := client.Delete(ctx, resourceGroup, name); err != nil {
        return fmt.Errorf("Error deleting Cluster %q (Resource Group %q): %+v", name, resourceGroup, err)
    }

    return nil
}

func expandArmClusterIdentity(input []interface{}) *operationalinsights.Identity {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})

    type := v["type"].(string)

    result := operationalinsights.Identity{
        Type: operationalinsights.IdentityType(type),
    }
    return &result
}