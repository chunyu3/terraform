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



func resourceArmSuppression() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmSuppressionCreateUpdate,
        Read: resourceArmSuppressionRead,
        Update: resourceArmSuppressionCreateUpdate,
        Delete: resourceArmSuppressionDelete,

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

            "recommendation_id": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "resource_uri": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "suppression_id": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
            },

            "ttl": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
            },

            "type": {
                Type: schema.TypeString,
                Computed: true,
            },

            "tags": tags.Schema(),
        },
    }
}

func resourceArmSuppressionCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).suppressionsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    recommendationID := d.Get("recommendation_id").(string)
    resourceURI := d.Get("resource_uri").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceURI, recommendationID, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Suppression %q (Recommendation %q / Resource Uri %q): %+v", name, recommendationID, resourceURI, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_suppression", *existing.ID)
        }
    }

    location := azure.NormalizeLocation(d.Get("location").(string))
    suppressionId := d.Get("suppression_id").(string)
    ttl := d.Get("ttl").(string)
    t := d.Get("tags").(map[string]interface{})

    suppressionContract := advisor.SuppressionContract{
        Location: utils.String(location),
        SuppressionID: utils.String(suppressionId),
        Tags: tags.Expand(t),
        TTL: utils.String(ttl),
    }


    if _, err := client.Create(ctx, resourceURI, recommendationID, name, suppressionContract); err != nil {
        return fmt.Errorf("Error creating Suppression %q (Recommendation %q / Resource Uri %q): %+v", name, recommendationID, resourceURI, err)
    }


    resp, err := client.Get(ctx, resourceURI, recommendationID, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Suppression %q (Recommendation %q / Resource Uri %q): %+v", name, recommendationID, resourceURI, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Suppression %q (Recommendation %q / Resource Uri %q) ID", name, recommendationID, resourceURI)
    }
    d.SetId(*resp.ID)

    return resourceArmSuppressionRead(d, meta)
}

func resourceArmSuppressionRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).suppressionsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    recommendationID := id.Path["recommendations"]
    name := id.Path["suppressions"]

    resp, err := client.Get(ctx, resourceURI, recommendationID, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Suppression %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Suppression %q (Recommendation %q / Resource Uri %q): %+v", name, recommendationID, resourceURI, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    if location := resp.Location; location != nil {
        d.Set("location", azure.NormalizeLocation(*location))
    }
    d.Set("recommendation_id", recommendationID)
    d.Set("resource_uri", resourceURI)
    d.Set("suppression_id", resp.SuppressionID)
    d.Set("ttl", resp.TTL)
    d.Set("type", resp.Type)

    return tags.FlattenAndSet(d, resp.Tags)
}


func resourceArmSuppressionDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).suppressionsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    recommendationID := id.Path["recommendations"]
    name := id.Path["suppressions"]

    if _, err := client.Delete(ctx, resourceURI, recommendationID, name); err != nil {
        return fmt.Errorf("Error deleting Suppression %q (Recommendation %q / Resource Uri %q): %+v", name, recommendationID, resourceURI, err)
    }

    return nil
}
