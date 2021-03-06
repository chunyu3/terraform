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



func resourceArmCache() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmCacheCreate,
        Read: resourceArmCacheRead,
        Update: resourceArmCacheUpdate,
        Delete: resourceArmCacheDelete,

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

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "cache_id": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "connection_string": {
                Type: schema.TypeString,
                Optional: true,
            },

            "description": {
                Type: schema.TypeString,
                Optional: true,
            },

            "resource_id": {
                Type: schema.TypeString,
                Optional: true,
            },

            "type": {
                Type: schema.TypeString,
                Computed: true,
            },
        },
    }
}

func resourceArmCacheCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).cacheClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    cacheID := d.Get("cache_id").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, name, cacheID)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Cache %q (Cache %q / Resource Group %q): %+v", name, cacheID, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_cache", *existing.ID)
        }
    }

    connectionString := d.Get("connection_string").(string)
    description := d.Get("description").(string)
    resourceId := d.Get("resource_id").(string)

    parameters := apimanagement.CacheUpdateParameters{
        CacheUpdateProperties: &apimanagement.CacheUpdateProperties{
            ConnectionString: utils.String(connectionString),
            Description: utils.String(description),
            ResourceID: utils.String(resourceId),
        },
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroup, name, cacheID, parameters); err != nil {
        return fmt.Errorf("Error creating Cache %q (Cache %q / Resource Group %q): %+v", name, cacheID, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, name, cacheID)
    if err != nil {
        return fmt.Errorf("Error retrieving Cache %q (Cache %q / Resource Group %q): %+v", name, cacheID, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Cache %q (Cache %q / Resource Group %q) ID", name, cacheID, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmCacheRead(d, meta)
}

func resourceArmCacheRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).cacheClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["service"]
    cacheID := id.Path["caches"]

    resp, err := client.Get(ctx, resourceGroup, name, cacheID)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Cache %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Cache %q (Cache %q / Resource Group %q): %+v", name, cacheID, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    d.Set("resource_group", resourceGroup)
    d.Set("cache_id", cacheID)
    if cacheUpdateProperties := resp.CacheUpdateProperties; cacheUpdateProperties != nil {
        d.Set("connection_string", cacheUpdateProperties.ConnectionString)
        d.Set("description", cacheUpdateProperties.Description)
        d.Set("resource_id", cacheUpdateProperties.ResourceID)
    }
    d.Set("type", resp.Type)

    return nil
}

func resourceArmCacheUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).cacheClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    cacheID := d.Get("cache_id").(string)
    connectionString := d.Get("connection_string").(string)
    description := d.Get("description").(string)
    resourceId := d.Get("resource_id").(string)

    parameters := apimanagement.CacheUpdateParameters{
        CacheUpdateProperties: &apimanagement.CacheUpdateProperties{
            ConnectionString: utils.String(connectionString),
            Description: utils.String(description),
            ResourceID: utils.String(resourceId),
        },
    }


    if _, err := client.Update(ctx, resourceGroup, name, cacheID, parameters); err != nil {
        return fmt.Errorf("Error updating Cache %q (Cache %q / Resource Group %q): %+v", name, cacheID, resourceGroup, err)
    }

    return resourceArmCacheRead(d, meta)
}

func resourceArmCacheDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).cacheClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["service"]
    cacheID := id.Path["caches"]

    if _, err := client.Delete(ctx, resourceGroup, name, cacheID); err != nil {
        return fmt.Errorf("Error deleting Cache %q (Cache %q / Resource Group %q): %+v", name, cacheID, resourceGroup, err)
    }

    return nil
}
